// The dnsp package implements the DNSP client protocol.
//
// # Introduction
//
// The Domain Name Service (DNS) Server Management Protocol defines RPC interfaces that
// provide methods for remotely accessing and administering a DNS server. It is a client/server
// protocol based on RPC that is used in the configuration, management, and monitoring
// of a DNS server.
//
// # Overview
//
// The DNS Server Management Protocol is a client/server protocol that is used to remotely
// query, monitor and configure DNS server settings, its zones, and resource records.
// The protocol allows a client to access DNS server settings and properties and also
// to enumerate all DNS data stored on the server (DNS zones and DNS records).
//
// The DNS Server Management Protocol is a simple protocol with no state shared across
// RPC method calls. Each RPC method call contains one complete request. Output from
// one method call can be used as input to another call but the DNS Server Management
// Protocol does not provide for locking of DNS data across method calls. For example,
// a client can enumerate DNS zones with one call and retrieve the properties of one
// or more of the enumerated DNS zones with another call. However, no guarantee is made
// that the zone has not been deleted by another DNS Server Management Protocol client
// between these two method calls.
//
// When the DNS server is directory server-integrated, some client requests can require
// or trigger Lightweight Directory Access Protocol (LDAP) requests from the DNS server
// to the local directory server or another directory server.
//
// In particular, the DNS server can use the defaultNamingContext of the directory server's
// rootDSE, a DNS domain partition named DnsDomainZones, or a DNS forest partition named
// DnsForestZones to store zone information and zone records. (See section 2.3 for a
// discussion of the schemas used to store this information.) A DNS Server integrated
// with a directory server creates and automatically enlists itself in these default
// application directory partitions. Alternatively, zone information and zone records
// can be stored in additional application directory partitions, which can be created
// (and removed) by the DNS Server Management Protocol client in order to control the
// granularity of zone replication. Zones created in these additional application directory
// partitions will only be visible to directory servers enlisted in those partitions,
// thus allowing for granular control over replication.
//
// A typical remote management involves the client querying or setting the configuration
// parameters of the DNS server. The client can also enumerate DNS zones and the DNS
// records stored in one or more zones. The client can modify the configuration of the
// DNS server as required. The client can also add, delete, or modify DNS zones or the
// DNS records held in zones as required. For example, a remote management client can:
//
// * Set or retrieve the server's forwarders ( a95b05da-f1fd-4db3-94b4-817fdaa1f642#gt_025cfacf-ebc5-4659-971a-ee2ab5903575
// ).
//
// * Set or retrieve various DNS server settings.
//
// * Create or modify zones.
//
// * Create or modify zone records.
//
// This usually involves sending a request to the DNS server specifying the type of
// operation (get, set and execute are examples of types of operations) to perform and
// any specific parameters that are associated with that operation. The DNS server responds
// to the client with the result of the operation.
//
// The following diagram shows an example of a remote client creating a zone on the
// DNS server using the DNS server Management Protocol. The client sends a request to
// the server with the operation type and parameters. The server responds with a success
// or an error.
package dnsp

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
	GoPackage = "dnsp"
)

// DPAutocreated represents the DNS_DP_AUTOCREATED RPC constant
var DPAutocreated = 1

// DPLegacy represents the DNS_DP_LEGACY RPC constant
var DPLegacy = 2

// DPDomainDefault represents the DNS_DP_DOMAIN_DEFAULT RPC constant
var DPDomainDefault = 4

// DPForestDefault represents the DNS_DP_FOREST_DEFAULT RPC constant
var DPForestDefault = 8

// DPEnlisted represents the DNS_DP_ENLISTED RPC constant
var DPEnlisted = 16

// DPDeleted represents the DNS_DP_DELETED RPC constant
var DPDeleted = 32

// DPStateOkay represents the DNS_DP_STATE_OKAY RPC constant
var DPStateOkay = 0

// DPStateReplicationIncoming represents the DNS_DP_STATE_REPL_INCOMING RPC constant
var DPStateReplicationIncoming = 1

// DPStateReplicationOutgoing represents the DNS_DP_STATE_REPL_OUTGOING RPC constant
var DPStateReplicationOutgoing = 2

// DPStateUnknown represents the DNS_DP_STATE_UNKNOWN RPC constant
var DPStateUnknown = 3

// DPOperationCreate represents the DNS_DP_OP_CREATE RPC constant
var DPOperationCreate = 1

// DPOperationDelete represents the DNS_DP_OP_DELETE RPC constant
var DPOperationDelete = 2

// DPOperationEnlist represents the DNS_DP_OP_ENLIST RPC constant
var DPOperationEnlist = 3

// DPOperationUnenlist represents the DNS_DP_OP_UNENLIST RPC constant
var DPOperationUnenlist = 4

// DPOperationCreateDomain represents the DNS_DP_OP_CREATE_DOMAIN RPC constant
var DPOperationCreateDomain = 5

// DPOperationCreateForest represents the DNS_DP_OP_CREATE_FOREST RPC constant
var DPOperationCreateForest = 6

// ZoneStatsVer represents the DNS_RPC_ZONE_STATS_VER RPC constant
var ZoneStatsVer = 1

// ZoneScopeCreateInfoVer represents the DNS_RPC_ZONE_SCOPE_CREATE_INFO_VER RPC constant
var ZoneScopeCreateInfoVer = 1

// ZoneScopeInfoVer represents the DNS_RPC_ZONE_SCOPE_INFO_VER RPC constant
var ZoneScopeInfoVer = 1

// EnumZoneScopeListVer represents the DNS_RPC_ENUM_ZONE_SCOPE_LIST_VER RPC constant
var EnumZoneScopeListVer = 1

// MaxRecordTypes represents the MAX_RECORD_TYPES RPC constant
var MaxRecordTypes = 32

// MaxZoneTransferTypes represents the MAX_ZONE_TRANSFER_TYPES RPC constant
var MaxZoneTransferTypes = 2

// ZoneRequestFilter type represents ZONE_REQUEST_FILTER RPC enumeration.
type ZoneRequestFilter uint16

var (
	ZoneRequestFilterPrimary   ZoneRequestFilter = 1
	ZoneRequestFilterSecondary ZoneRequestFilter = 2
	ZoneRequestFilterCache     ZoneRequestFilter = 4
	ZoneRequestFilterAuto      ZoneRequestFilter = 8
	ZoneRequestFilterForward   ZoneRequestFilter = 16
	ZoneRequestFilterReverse   ZoneRequestFilter = 32
	ZoneRequestFilterForwarder ZoneRequestFilter = 64
	ZoneRequestFilterStub      ZoneRequestFilter = 128
	ZoneRequestFilterDS        ZoneRequestFilter = 256
	ZoneRequestFilterNonDS     ZoneRequestFilter = 512
	ZoneRequestFilterDomainDP  ZoneRequestFilter = 1024
	ZoneRequestFilterForestDP  ZoneRequestFilter = 2048
	ZoneRequestFilterCustomDP  ZoneRequestFilter = 4096
	ZoneRequestFilterLegacyDP  ZoneRequestFilter = 8192
	ZoneRequestFilterAll       ZoneRequestFilter = 16383
)

func (o ZoneRequestFilter) String() string {
	switch o {
	case ZoneRequestFilterPrimary:
		return "ZoneRequestFilterPrimary"
	case ZoneRequestFilterSecondary:
		return "ZoneRequestFilterSecondary"
	case ZoneRequestFilterCache:
		return "ZoneRequestFilterCache"
	case ZoneRequestFilterAuto:
		return "ZoneRequestFilterAuto"
	case ZoneRequestFilterForward:
		return "ZoneRequestFilterForward"
	case ZoneRequestFilterReverse:
		return "ZoneRequestFilterReverse"
	case ZoneRequestFilterForwarder:
		return "ZoneRequestFilterForwarder"
	case ZoneRequestFilterStub:
		return "ZoneRequestFilterStub"
	case ZoneRequestFilterDS:
		return "ZoneRequestFilterDS"
	case ZoneRequestFilterNonDS:
		return "ZoneRequestFilterNonDS"
	case ZoneRequestFilterDomainDP:
		return "ZoneRequestFilterDomainDP"
	case ZoneRequestFilterForestDP:
		return "ZoneRequestFilterForestDP"
	case ZoneRequestFilterCustomDP:
		return "ZoneRequestFilterCustomDP"
	case ZoneRequestFilterLegacyDP:
		return "ZoneRequestFilterLegacyDP"
	case ZoneRequestFilterAll:
		return "ZoneRequestFilterAll"
	}
	return "Invalid"
}

// StatID type represents DNSSRV_STATID RPC enumeration.
type StatID uint32

var (
	StatIDTime          StatID = 1
	StatIDQuery         StatID = 2
	StatIDQuery2        StatID = 4
	StatIDRecurse       StatID = 8
	StatIDMaster        StatID = 16
	StatIDSecondary     StatID = 32
	StatIDWINS          StatID = 64
	StatIDWireUpdate    StatID = 256
	StatIDSKWANSEC      StatID = 512
	StatIDDS            StatID = 1024
	StatIDNonWireUpdate StatID = 2048
	StatIDMemory        StatID = 65536
	StatIDTimeout       StatID = 131072
	StatIDDatabase      StatID = 262144
	StatIDRecord        StatID = 524288
	StatIDPacket        StatID = 1048576
	StatIDNBSTAT        StatID = 2097152
	StatIDErrors        StatID = 4194304
	StatIDCache         StatID = 8388608
	StatIDDNSSEC        StatID = 16777216
	StatIDPrivate       StatID = 268435456
	StatIDRRL           StatID = 536870912
	StatIDAll           StatID = 838799231
)

func (o StatID) String() string {
	switch o {
	case StatIDTime:
		return "StatIDTime"
	case StatIDQuery:
		return "StatIDQuery"
	case StatIDQuery2:
		return "StatIDQuery2"
	case StatIDRecurse:
		return "StatIDRecurse"
	case StatIDMaster:
		return "StatIDMaster"
	case StatIDSecondary:
		return "StatIDSecondary"
	case StatIDWINS:
		return "StatIDWINS"
	case StatIDWireUpdate:
		return "StatIDWireUpdate"
	case StatIDSKWANSEC:
		return "StatIDSKWANSEC"
	case StatIDDS:
		return "StatIDDS"
	case StatIDNonWireUpdate:
		return "StatIDNonWireUpdate"
	case StatIDMemory:
		return "StatIDMemory"
	case StatIDTimeout:
		return "StatIDTimeout"
	case StatIDDatabase:
		return "StatIDDatabase"
	case StatIDRecord:
		return "StatIDRecord"
	case StatIDPacket:
		return "StatIDPacket"
	case StatIDNBSTAT:
		return "StatIDNBSTAT"
	case StatIDErrors:
		return "StatIDErrors"
	case StatIDCache:
		return "StatIDCache"
	case StatIDDNSSEC:
		return "StatIDDNSSEC"
	case StatIDPrivate:
		return "StatIDPrivate"
	case StatIDRRL:
		return "StatIDRRL"
	case StatIDAll:
		return "StatIDAll"
	}
	return "Invalid"
}

// StatHeader structure represents DNSSRV_STAT_HEADER RPC structure.
//
// The DNSSRV_STAT_HEADER structure precedes each DNSSRV_STATS structure (section 2.2.10.2.2)
// which provides DNS server runtime statistics. This structure MUST be formatted as
// follows:
type StatHeader struct {
	// StatId: The type of statistics contained in the DNSSRV_STATS structure. This value
	// MUST be set to one of the allowed values specified in section 2.2.10.1.1.
	StatID uint32 `idl:"name:StatId" json:"stat_id"`
	// wLength: The length, in bytes, of the Buffer member in the DNSSRV_STATS structure
	// (section 2.2.10.2.2).
	Length uint16 `idl:"name:wLength" json:"length"`
	// fClear: A Boolean value that indicates whether the server is to clear the statistics
	// buffer for the server attribute indicated at by StatId.
	Clear bool `idl:"name:fClear" json:"clear"`
	// fReserved: MUST be set to zero when sent and MUST be ignored on receipt.
	_ uint8 `idl:"name:fReserved"`
}

func (o *StatHeader) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *StatHeader) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.StatID); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if err := w.WriteData(o.Clear); err != nil {
		return err
	}
	// reserved fReserved
	if err := w.WriteData(uint8(0)); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *StatHeader) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.StatID); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	if err := w.ReadData(&o.Clear); err != nil {
		return err
	}
	// reserved fReserved
	var _fReserved uint8
	if err := w.ReadData(&_fReserved); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// Stat structure represents DNSSRV_STAT RPC structure.
//
// The DNSSRV_STATS structure carries server statistics information. This structure
// MUST be interpreted as one of the more specific statistics structures specified in
// sections 2.2.10.2.4 through 2.2.10.2.24, depending upon the StatId value in the Header
// member. This structure MUST be formatted as follows:
type Stat struct {
	StatID uint32 `idl:"name:StatId" json:"stat_id"`
	Length uint16 `idl:"name:wLength" json:"length"`
	Clear  bool   `idl:"name:fClear" json:"clear"`
	_      uint8  `idl:"name:fReserved"`
	// Buffer: A variable length array of bytes that contains information specific to the
	// type of DNS server statistics, as specified by the StatId value in the Header.
	Buffer []byte `idl:"name:Buffer;size_is:(wLength)" json:"buffer"`
}

func (o *Stat) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Buffer != nil && o.Length == 0 {
		o.Length = uint16(len(o.Buffer))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Stat) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.StatID); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if err := w.WriteData(o.Clear); err != nil {
		return err
	}
	// reserved fReserved
	if err := w.WriteData(uint8(0)); err != nil {
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
func (o *Stat) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.StatID); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	if err := w.ReadData(&o.Clear); err != nil {
		return err
	}
	// reserved fReserved
	var _fReserved uint8
	if err := w.ReadData(&_fReserved); err != nil {
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

// IPv4Array structure represents IP4_ARRAY RPC structure.
//
// The IP4_ARRAY structure is used to represent an array of IPv4 addresses. This structure
// cannot represent IPv6 addresses.
type IPv4Array struct {
	// AddrCount: The number of IPv4 addresses present in the AddrArray member.
	AddrCount uint32 `idl:"name:AddrCount" json:"addr_count"`
	// AddrArray: An array of IPv4 addresses. An IPv4 address is represented as a 32-bit
	// unsigned integer in network byte order.
	//
	// An empty IP4_ARRAY is represented by AddrCount set to zero and AddrArray unused.
	// Senders of an empty IP4_ARRAY MUST set AddrArray to a single entry containing binary
	// zeros, and receivers MUST ignore it.
	AddrArray []uint32 `idl:"name:AddrArray;size_is:(AddrCount)" json:"addr_array"`
}

func (o *IPv4Array) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.AddrArray != nil && o.AddrCount == 0 {
		o.AddrCount = uint32(len(o.AddrArray))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *IPv4Array) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.AddrCount)
	return []uint64{
		dimSize1,
	}
}
func (o *IPv4Array) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.AddrCount); err != nil {
		return err
	}
	for i1 := range o.AddrArray {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.AddrArray[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.AddrArray); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPv4Array) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.AddrCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.AddrCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.AddrCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.AddrArray", sizeInfo[0])
	}
	o.AddrArray = make([]uint32, sizeInfo[0])
	for i1 := range o.AddrArray {
		i1 := i1
		if err := w.ReadData(&o.AddrArray[i1]); err != nil {
			return err
		}
	}
	return nil
}

// Addr structure represents DNS_ADDR RPC structure.
//
// This DNS_ADDR structure is used to represent an IP address. The IP address is either
// IPv4 or IPv6.
//
// The DNS ADDR structure is an IPv4 or IPv6 address used by DNS_ADDR MaxSa field (section
// 2.2.3.2.2).
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Address Family                                                | Port Number                                                   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| IPv4 Address                                                                                                                  |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| IPv6 Address (16 bytes)                                                                                                       |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Padding                                                                                                                       |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type Addr struct {
	// MaxSa: This field MUST be constructed as specified in DNS ADDR.
	MaxSA []byte `idl:"name:MaxSa" json:"max_sa"`
	// DnsAddrUserDword: This field MUST be constructed as specified in DNS ADD USER.
	//
	// Any field not specified above MUST be set to zero by the sender and ignored by the
	// receiver.
	DNSAddrUser []uint32 `idl:"name:DnsAddrUserDword" json:"dns_addr_user"`
}

func (o *Addr) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Addr) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	for i1 := range o.MaxSA {
		i1 := i1
		if uint64(i1) >= 32 {
			break
		}
		if err := w.WriteData(o.MaxSA[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.MaxSA); uint64(i1) < 32; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.DNSAddrUser {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.DNSAddrUser[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.DNSAddrUser); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *Addr) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	o.MaxSA = make([]byte, 32)
	for i1 := range o.MaxSA {
		i1 := i1
		if err := w.ReadData(&o.MaxSA[i1]); err != nil {
			return err
		}
	}
	o.DNSAddrUser = make([]uint32, 8)
	for i1 := range o.DNSAddrUser {
		i1 := i1
		if err := w.ReadData(&o.DNSAddrUser[i1]); err != nil {
			return err
		}
	}
	return nil
}

// AddrArray structure represents DNS_ADDR_ARRAY RPC structure.
//
// The DNS_ADDR_ARRAY structure is used to represent an array of DNS_ADDR (section 2.2.3.2.2)
// structures. The DNS Server Management Protocol uses this structure to exchange lists
// of mixed IPv4 and IPv6 addresses between client and server.
type AddrArray struct {
	// MaxCount: The actual number of IP addresses that are present in the AddrArray member.
	MaxCount uint32 `idl:"name:MaxCount" json:"max_count"`
	// AddrCount: Must be set to the same value as MaxCount.
	AddrCount uint32 `idl:"name:AddrCount" json:"addr_count"`
	// Tag: This field is unused. Senders MUST set the value to zero and receivers MUST
	// ignore it.
	Tag uint32 `idl:"name:Tag" json:"tag"`
	// Family: The family of addresses present in the array, such as AF_INET or AF_INET6.
	// If this field is not specified, addresses with all families can be present.
	Family uint16 `idl:"name:Family" json:"family"`
	// WordReserved: This field is unused. Senders MUST set the value to zero and receivers
	// MUST ignore it.
	_ uint16 `idl:"name:WordReserved"`
	// Flags: This field is unused. Senders MUST set the value to zero and receivers MUST
	// ignore it.
	Flags uint32 `idl:"name:Flags" json:"flags"`
	// MatchFlag: This field is unused. Senders MUST set the value to zero and receivers
	// MUST ignore it.
	MatchFlag uint32 `idl:"name:MatchFlag" json:"match_flag"`
	// Reserved1: This field is unused. Senders MUST set the value to zero and receivers
	// MUST ignore it.
	_ uint32 `idl:"name:Reserved1"`
	// Reserved2: This field is unused. Senders MUST set the value to zero and receivers
	// MUST ignore it.
	_ uint32 `idl:"name:Reserved2"`
	// AddrArray: An array of DNS_ADDR (section 2.2.3.2.2) structures. The number of elements
	// in this array is specified by the AddrCount member.
	//
	// An empty DNS_ADDR_ARRAY is represented by AddrCount set to zero. Senders of an empty
	// DNS_ADR_ARRAY MUST set the other fields' values to zero (including a single entry
	// in AddrArray, which is set to binary zeros), and receivers MUST ignore them.
	AddrArray []*Addr `idl:"name:AddrArray;size_is:(AddrCount)" json:"addr_array"`
}

func (o *AddrArray) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.AddrArray != nil && o.AddrCount == 0 {
		o.AddrCount = uint32(len(o.AddrArray))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *AddrArray) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.AddrCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AddrArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.MaxCount); err != nil {
		return err
	}
	if err := w.WriteData(o.AddrCount); err != nil {
		return err
	}
	if err := w.WriteData(o.Tag); err != nil {
		return err
	}
	if err := w.WriteData(o.Family); err != nil {
		return err
	}
	// reserved WordReserved
	if err := w.WriteData(uint16(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.MatchFlag); err != nil {
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
	for i1 := range o.AddrArray {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.AddrArray[i1] != nil {
			if err := o.AddrArray[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Addr{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.AddrArray); uint64(i1) < sizeInfo[0]; i1++ {
		if err := (&Addr{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *AddrArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.MaxCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddrCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.Tag); err != nil {
		return err
	}
	if err := w.ReadData(&o.Family); err != nil {
		return err
	}
	// reserved WordReserved
	var _WordReserved uint16
	if err := w.ReadData(&_WordReserved); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.MatchFlag); err != nil {
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
	// XXX: for opaque unmarshaling
	if o.AddrCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.AddrCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.AddrArray", sizeInfo[0])
	}
	o.AddrArray = make([]*Addr, sizeInfo[0])
	for i1 := range o.AddrArray {
		i1 := i1
		if o.AddrArray[i1] == nil {
			o.AddrArray[i1] = &Addr{}
		}
		if err := o.AddrArray[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// Buffer structure represents DNS_RPC_BUFFER RPC structure.
//
// The DNS_RPC_BUFFER structure contains a set of a specific type of structures. The
// DNS Server Management Protocol uses this structure to return information from the
// server, while processing R_DnssrvComplexOperation2 (section 3.1.4.8) method calls
// with operation type "Statistics".
type Buffer struct {
	// dwLength: The length, in bytes, of the data stored in Buffer.
	Length uint32 `idl:"name:dwLength" json:"length"`
	// Buffer: A variable length array of bytes of length specified by dwLength. The buffer
	// can contain one or more DNS_RPC_NODE structures (section 2.2.2.2.3). Each DNS_RPC_NODE
	// contains the length of that node, so the DNS_RPC_BUFFER dwLength can be larger to
	// indicate multiple DNS_RPC_NODE structures.
	Buffer []byte `idl:"name:Buffer;size_is:(dwLength)" json:"buffer"`
}

func (o *Buffer) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Buffer != nil && o.Length == 0 {
		o.Length = uint32(len(o.Buffer))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *Buffer) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.Length)
	return []uint64{
		dimSize1,
	}
}
func (o *Buffer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.Length); err != nil {
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
}
func (o *Buffer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.Length); err != nil {
		return err
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
}

// ServerExtension structure represents DNS_RPC_SERVER_EXTENSION RPC structure.
type ServerExtension struct {
	Extension uint32 `idl:"name:Extension" json:"extension"`
}

func (o *ServerExtension) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ServerExtension) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Extension); err != nil {
		return err
	}
	return nil
}
func (o *ServerExtension) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Extension); err != nil {
		return err
	}
	return nil
}

// ServerInfoW2K structure represents DNS_RPC_SERVER_INFO_W2K RPC structure.
//
// The DNS_RPC_SERVER_INFO_W2K structure is used to specify general DNS server state
// and configuration.
type ServerInfoW2K struct {
	// dwVersion: The operating system version of the DNS server in DNSSRV_VERSION (section
	// 2.2.4.2.1).
	Version uint32 `idl:"name:dwVersion" json:"version"`
	// fBootMethod: The method by which the DNS server obtains information at the start
	// time. This MUST be set to one of the possible values as specified in DNS_BOOT_METHODS
	// (section 2.2.4.1.1).
	BootMethod uint8 `idl:"name:fBootMethod" json:"boot_method"`
	// fAdminConfigured: A Boolean field that specifies whether the DNS server has been
	// configured by an administrator. On a fresh installed server this value MUST be set
	// to FALSE. This value MUST be set to TRUE whenever a zone is created, or a record
	// is modified, or an Active Directory domain controller promotion (DCPROMO) configures
	// the DNS server.
	AdminConfigured bool `idl:"name:fAdminConfigured" json:"admin_configured"`
	// fAllowUpdate: A Boolean field that indicates whether the DNS server allows dynamic
	// DNS updates. This field MUST be set to FALSE if the server does not allow dynamic
	// zone-updates, otherwise set to TRUE.
	AllowUpdate bool `idl:"name:fAllowUpdate" json:"allow_update"`
	// fDsAvailable: A Boolean field that specifies whether a directory server is available
	// to the DNS server. It MUST be set to FALSE, if the server does not have access to
	// a directory server.
	DSAvailable bool `idl:"name:fDsAvailable" json:"ds_available"`
	// pszServerName: A pointer to a null-terminated UTF-8 string that contains the FQDN
	// of the DNS server.
	ServerName string `idl:"name:pszServerName;string" json:"server_name"`
	// pszDsContainer: A pointer to a null-terminated Unicode string that points to the
	// DNS server's container path as a distinguished name (DN) in the directory server.
	// If no directory server is configured, this value MUST be set to NULL. This value
	// is synthesized by the server by concatenating a constant container relative distinguished
	// name (RDN) and the result of an LDAP search operation to retrieve the defaultNamingContext
	// of the Active Directory server's rootDSE.<31>
	DSContainer string `idl:"name:pszDsContainer;string" json:"ds_container"`
	// aipServerAddrs: The list of IP addresses that are available on the server.
	ServerAddrs *IPv4Array `idl:"name:aipServerAddrs" json:"server_addrs"`
	// aipListenAddrs: The list of IP addresses that are explicitly configured by the administrator
	// on the DNS server that listens for the DNS requests. If this value is set to NULL
	// then the server listens to all available IP addresses.
	ListenAddrs *IPv4Array `idl:"name:aipListenAddrs" json:"listen_addrs"`
	// aipForwarders: The list of remote DNS servers to which this DNS server will forward
	// unresolved DNS requests.
	Forwarders *IPv4Array `idl:"name:aipForwarders" json:"forwarders"`
	// pExtension1: Reserved for future use and MUST be ignored by receiver.
	Extension1 *ServerExtension `idl:"name:pExtension1" json:"extension1"`
	// pExtension2: Reserved for future use and MUST be ignored by receiver.
	Extension2 *ServerExtension `idl:"name:pExtension2" json:"extension2"`
	// pExtension3: Reserved for future use and MUST be ignored by receiver.
	Extension3 *ServerExtension `idl:"name:pExtension3" json:"extension3"`
	// pExtension4: Reserved for future use and MUST be ignored by receiver.
	Extension4 *ServerExtension `idl:"name:pExtension4" json:"extension4"`
	// pExtension5: Reserved for future use and MUST be ignored by receiver.
	Extension5 *ServerExtension `idl:"name:pExtension5" json:"extension5"`
	// dwLogLevel: This indicates which DNS packets will be logged and how they will be
	// logged. This field MUST be set to either zero or a combination (by bitwise OR) of
	// the possible values as specified under DNS_LOG_LEVELS (section 2.2.9.1.1). If this
	// value is set to zero, then no logging will be performed for DNS packets.
	LogLevel uint32 `idl:"name:dwLogLevel" json:"log_level"`
	// dwDebugLevel: Unused. Receivers MUST ignore.
	DebugLevel uint32 `idl:"name:dwDebugLevel" json:"debug_level"`
	// dwForwardTimeout: The time interval, in seconds, for which the DNS server waits for
	// a response from each server in the forwarders list.
	ForwardTimeout uint32 `idl:"name:dwForwardTimeout" json:"forward_timeout"`
	// dwRpcProtocol: This value indicates what RPC protocols this DNS server will accept
	// connections on. This value MUST be set to any combination of values specified in
	// DNS_RPC_PROTOCOLS (section 2.2.1.1.2).
	RPCProtocol uint32 `idl:"name:dwRpcProtocol" json:"rpc_protocol"`
	// dwNameCheckFlag: The level of domain name checking and validation enforced by the
	// DNS server. This value MUST be set one of the allowed values that are specified in
	// DNS_NAME_CHECK_FLAGS (section 2.2.4.1.2).
	NameCheckFlag uint32 `idl:"name:dwNameCheckFlag" json:"name_check_flag"`
	// cAddressAnswerLimit: The configured value for the maximum number of type A IP address
	// resource records that the DNS server can insert in the answer section of a response
	// to a UDP query of type A. If this value is set to 0x00000000 then the DNS server
	// MUST NOT enforce any artificial limit on number of records in a response and if response
	// becomes larger than the DNS UDP packet size then the truncation bit MUST be set [RFC1035].
	// If this property value is not 0x00000000 and the DNS server is unable to add the
	// specified number of records to the response message due to message size limitations,
	// it MUST return as many records as will fit in the message and the truncation bit
	// MUST NOT be set. The DNS server MUST NOT enforce this limit if the query is not of
	// type A. If the value of this property is not 0x00000000 the DNS server MUST enforce
	// this limit for UDP queries and MUST NOT enforce this limit for TCP queries. If the
	// LocalNetPriority property value is set to TRUE, the DNS server first orders the address
	// records as per the LocalNetPriority property and then MUST select the first cAddressAnswerLimit
	// type A records in this sorted list for inclusion in the response. The value MUST
	// be either zero or between 0x00000005 and 0x0000001C inclusive.
	AddressAnswerLimitCount uint32 `idl:"name:cAddressAnswerLimit" json:"address_answer_limit_count"`
	// dwRecursionRetry: The time-interval, in seconds, for which the DNS server waits before
	// it retries a recursive query to the remote DNS server from which it did not receive
	// a response. The values MUST be between 1 and 15 seconds inclusive.
	RecursionRetry uint32 `idl:"name:dwRecursionRetry" json:"recursion_retry"`
	// dwRecursionTimeout: The time-interval, in seconds, for which the DNS server waits
	// for a recursive query-response from a remote DNS server. The values MUST be between
	// 1 and 15 seconds inclusive.
	RecursionTimeout uint32 `idl:"name:dwRecursionTimeout" json:"recursion_timeout"`
	// dwMaxCacheTtl: The maximum time duration, in seconds, for which the DNS server will
	// cache a resource record obtained from a remote server in a successful query response.
	// The values for this MUST be between 0 to 30 days (but specified in seconds) inclusive.
	MaxCacheTTL uint32 `idl:"name:dwMaxCacheTtl" json:"max_cache_ttl"`
	// dwDsPollingInterval: The interval, in seconds, at which the DNS server will poll
	// a directory server to obtain updated information for any changes that have occurred
	// to zones loaded in the server. The values MUST be between 30 and 3600 seconds inclusive.
	DSPollingInterval uint32 `idl:"name:dwDsPollingInterval" json:"ds_polling_interval"`
	// dwScavengingInterval: The scavenging interval, in hours, on the DNS server. This
	// is the interval at which the server will execute the cleanup of stale DNS records.
	// The value MUST be between 0 and 8760 hours (1 year). If this value is zero then scavenging
	// is disabled.
	ScavengingInterval uint32 `idl:"name:dwScavengingInterval" json:"scavenging_interval"`
	// dwDefaultRefreshInterval: The default value of the refresh interval, in hours, for
	// new zones created on the DNS server. For any primary zone created on the server by
	// default this value is used as the refresh interval.
	DefaultRefreshInterval uint32 `idl:"name:dwDefaultRefreshInterval" json:"default_refresh_interval"`
	// dwDefaultNoRefreshInterval: The default value of the NoRefresh interval, in hours,
	// for new zones created on the DNS server. For any primary zone created on the server
	// by default this value is used as the NoRefresh interval.
	DefaultNoRefreshInterval uint32 `idl:"name:dwDefaultNoRefreshInterval" json:"default_no_refresh_interval"`
	// dwReserveArray: This value is reserved for future use and MUST be ignored by the
	// receiver. Senders MUST set this to zero and receivers MUST ignore it.
	_ []uint32 `idl:"name:dwReserveArray"`
	// fAutoReverseZones: A Boolean value that indicates whether the DNS server is configured
	// to automatically create standard reverse lookup zones at boot time.
	AutoReverseZones bool `idl:"name:fAutoReverseZones" json:"auto_reverse_zones"`
	// fAutoCacheUpdate: A Boolean value that indicates whether the DNS server is configured
	// to automatically write-back cached root hints and delegation data to persistent storage.
	AutoCacheUpdate bool `idl:"name:fAutoCacheUpdate" json:"auto_cache_update"`
	// fRecurseAfterForwarding: A Boolean value that indicates whether the DNS server is
	// configured to use recursion in addition to forwarding. If this value is TRUE (0x01)
	// then if the DNS server does not have any forwarders configured or if forwarders are
	// unreachable then it MUST return failure, otherwise it MUST perform normal recursive
	// processing for this query as specified in [RFC1034] section 4.3.1.
	RecurseAfterForwarding bool `idl:"name:fRecurseAfterForwarding" json:"recurse_after_forwarding"`
	// fForwardDelegations: A Boolean value indicates whether or not the DNS server will
	// forward queries about delegated subzones to servers outside of its authoritative
	// zone. If this value is set to TRUE, then the DNS server forwards all name queries
	// about delegated subzones to forwarding servers in other zones; otherwise it will
	// send such queries within its authoritative zone to the corresponding subzone only.
	ForwardDelegations bool `idl:"name:fForwardDelegations" json:"forward_delegations"`
	// fNoRecursion: A Boolean value that indicates whether the DNS server will perform
	// recursive resolution for queries. If this value is TRUE then recursion MUST NOT be
	// performed even if the Recursion Desired (RD) bit ([RFC1035] section 4.1.1) is set
	// in the DNS query packet header. If this value is FALSE then recursion will be performed
	// as per [RFC1035].
	NoRecursion bool `idl:"name:fNoRecursion" json:"no_recursion"`
	// fSecureResponses: A Boolean value that indicates if the DNS server needs to screen
	// DNS records received in remote query responses against the zone of authority for
	// the remote server to prevent cache pollution. If it is set to TRUE, the DNS server
	// caches only the records that are in zone of authority for the remote server that
	// was queried. When set to FALSE, all records in the cache are saved.
	SecureResponses bool `idl:"name:fSecureResponses" json:"secure_responses"`
	// fRoundRobin: A Boolean value that indicates if the DNS server is configured to rotate
	// the order of DNS records it returns for a given name. If this value is set to FALSE
	// no round robin will be performed and DNS records will be returned in static, arbitrary
	// order.
	RoundRobin bool `idl:"name:fRoundRobin" json:"round_robin"`
	// fLocalNetPriority: A Boolean value that indicates if the DNS server is configured
	// to prioritize address records in a response based on the IP address of the DNS client
	// that submitted the query. If this is set to TRUE the DNS server will return address
	// records in the order of their closeness to the querying client's IP address by applying
	// the network mask pointed to by LocalNetPriorityNetMask. If this value is set to 0
	// then the DNS server returns records in the order in which they are obtained from
	// the database.
	LocalNetPriority bool `idl:"name:fLocalNetPriority" json:"local_net_priority"`
	// fBindSecondaries: A Boolean value that indicates if the DNS server allows zone transfers
	// to secondary DNS servers running older non-Microsoft software. If this value is set
	// to TRUE the DNS server sends zone transfer to secondaries via a slower mechanism,
	// with one resource record in each message.
	BindSecondaries bool `idl:"name:fBindSecondaries" json:"bind_secondaries"`
	// fWriteAuthorityNs: A Boolean value that indicates if the DNS server is enabled to
	// write NS records in the authority section of all successful authoritative responses.
	// If this value is TRUE then NS records will be included in the authority section of
	// responses, otherwise NS records will only be included in referral responses.
	WriteAuthorityNS bool `idl:"name:fWriteAuthorityNs" json:"write_authority_ns"`
	// fStrictFileParsing: A Boolean value that indicates if the DNS server is configured
	// to perform strict file parsing. When this value is set to TRUE and a record parsing
	// error is detected server will quit after indicating error. If this value is FALSE
	// parsing errors will cause that specific record to be ignored and the server will
	// continue to load the rest of the database.
	StrictFileParsing bool `idl:"name:fStrictFileParsing" json:"strict_file_parsing"`
	// fLooseWildcarding: A Boolean value that indicates if the DNS server is configured
	// to perform loose wildcarding [RFC1035], otherwise it returns FALSE. When a server
	// does not find a resource record that matches the name and type specified in the query
	// in the authoritative zone, then it searches for related wildcard records, ([RFC1034]
	// section 4.3.3), if configured to perform loose wildcarding will return the first
	// node it finds that has matching resource-record type, whereas if it is not then it
	// will return the first node that has any resource record.
	LooseWildcarding bool `idl:"name:fLooseWildcarding" json:"loose_wildcarding"`
	// fDefaultAgingState: A Boolean value that indicates if the default value of ageing
	// state for new primary zones created on the DNS server. For any primary zone created
	// on the server this value is used as its default aging state. If this is FALSE then
	// timestamps of records in the zone will not be tracked whereas when this value is
	// TRUE then the timestamps of records in the zone will be tracked.
	DefaultAgingState bool `idl:"name:fDefaultAgingState" json:"default_aging_state"`
	// fReserveArray: Reserved for future use. These values MUST be ignored by receiver.
	_ []bool `idl:"name:fReserveArray"`
}

func (o *ServerInfoW2K) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ServerInfoW2K) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	if err := w.WriteData(o.BootMethod); err != nil {
		return err
	}
	if err := w.WriteData(o.AdminConfigured); err != nil {
		return err
	}
	if err := w.WriteData(o.AllowUpdate); err != nil {
		return err
	}
	if err := w.WriteData(o.DSAvailable); err != nil {
		return err
	}
	if o.ServerName != "" {
		_ptr_pszServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.ServerName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ServerName, _ptr_pszServerName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DSContainer != "" {
		_ptr_pszDsContainer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DSContainer); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DSContainer, _ptr_pszDsContainer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ServerAddrs != nil {
		_ptr_aipServerAddrs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ServerAddrs != nil {
				if err := o.ServerAddrs.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPv4Array{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ServerAddrs, _ptr_aipServerAddrs); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ListenAddrs != nil {
		_ptr_aipListenAddrs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ListenAddrs != nil {
				if err := o.ListenAddrs.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPv4Array{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ListenAddrs, _ptr_aipListenAddrs); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Forwarders != nil {
		_ptr_aipForwarders := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Forwarders != nil {
				if err := o.Forwarders.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPv4Array{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Forwarders, _ptr_aipForwarders); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Extension1 != nil {
		_ptr_pExtension1 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Extension1 != nil {
				if err := o.Extension1.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ServerExtension{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Extension1, _ptr_pExtension1); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Extension2 != nil {
		_ptr_pExtension2 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Extension2 != nil {
				if err := o.Extension2.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ServerExtension{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Extension2, _ptr_pExtension2); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Extension3 != nil {
		_ptr_pExtension3 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Extension3 != nil {
				if err := o.Extension3.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ServerExtension{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Extension3, _ptr_pExtension3); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Extension4 != nil {
		_ptr_pExtension4 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Extension4 != nil {
				if err := o.Extension4.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ServerExtension{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Extension4, _ptr_pExtension4); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Extension5 != nil {
		_ptr_pExtension5 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Extension5 != nil {
				if err := o.Extension5.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ServerExtension{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Extension5, _ptr_pExtension5); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LogLevel); err != nil {
		return err
	}
	if err := w.WriteData(o.DebugLevel); err != nil {
		return err
	}
	if err := w.WriteData(o.ForwardTimeout); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCProtocol); err != nil {
		return err
	}
	if err := w.WriteData(o.NameCheckFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.AddressAnswerLimitCount); err != nil {
		return err
	}
	if err := w.WriteData(o.RecursionRetry); err != nil {
		return err
	}
	if err := w.WriteData(o.RecursionTimeout); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxCacheTTL); err != nil {
		return err
	}
	if err := w.WriteData(o.DSPollingInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.ScavengingInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.DefaultRefreshInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.DefaultNoRefreshInterval); err != nil {
		return err
	}
	// reserved dwReserveArray
	for i1 := 0; uint64(i1) < 10; i1++ {
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.AutoReverseZones); err != nil {
		return err
	}
	if err := w.WriteData(o.AutoCacheUpdate); err != nil {
		return err
	}
	if err := w.WriteData(o.RecurseAfterForwarding); err != nil {
		return err
	}
	if err := w.WriteData(o.ForwardDelegations); err != nil {
		return err
	}
	if err := w.WriteData(o.NoRecursion); err != nil {
		return err
	}
	if err := w.WriteData(o.SecureResponses); err != nil {
		return err
	}
	if err := w.WriteData(o.RoundRobin); err != nil {
		return err
	}
	if err := w.WriteData(o.LocalNetPriority); err != nil {
		return err
	}
	if err := w.WriteData(o.BindSecondaries); err != nil {
		return err
	}
	if err := w.WriteData(o.WriteAuthorityNS); err != nil {
		return err
	}
	if err := w.WriteData(o.StrictFileParsing); err != nil {
		return err
	}
	if err := w.WriteData(o.LooseWildcarding); err != nil {
		return err
	}
	if err := w.WriteData(o.DefaultAgingState); err != nil {
		return err
	}
	// reserved fReserveArray
	for i1 := 0; uint64(i1) < 15; i1++ {
		if err := w.WriteData(false); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *ServerInfoW2K) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	if err := w.ReadData(&o.BootMethod); err != nil {
		return err
	}
	if err := w.ReadData(&o.AdminConfigured); err != nil {
		return err
	}
	if err := w.ReadData(&o.AllowUpdate); err != nil {
		return err
	}
	if err := w.ReadData(&o.DSAvailable); err != nil {
		return err
	}
	_ptr_pszServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.ServerName); err != nil {
			return err
		}
		return nil
	})
	_s_pszServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ServerName, _s_pszServerName, _ptr_pszServerName); err != nil {
		return err
	}
	_ptr_pszDsContainer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DSContainer); err != nil {
			return err
		}
		return nil
	})
	_s_pszDsContainer := func(ptr interface{}) { o.DSContainer = *ptr.(*string) }
	if err := w.ReadPointer(&o.DSContainer, _s_pszDsContainer, _ptr_pszDsContainer); err != nil {
		return err
	}
	_ptr_aipServerAddrs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ServerAddrs == nil {
			o.ServerAddrs = &IPv4Array{}
		}
		if err := o.ServerAddrs.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipServerAddrs := func(ptr interface{}) { o.ServerAddrs = *ptr.(**IPv4Array) }
	if err := w.ReadPointer(&o.ServerAddrs, _s_aipServerAddrs, _ptr_aipServerAddrs); err != nil {
		return err
	}
	_ptr_aipListenAddrs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ListenAddrs == nil {
			o.ListenAddrs = &IPv4Array{}
		}
		if err := o.ListenAddrs.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipListenAddrs := func(ptr interface{}) { o.ListenAddrs = *ptr.(**IPv4Array) }
	if err := w.ReadPointer(&o.ListenAddrs, _s_aipListenAddrs, _ptr_aipListenAddrs); err != nil {
		return err
	}
	_ptr_aipForwarders := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Forwarders == nil {
			o.Forwarders = &IPv4Array{}
		}
		if err := o.Forwarders.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipForwarders := func(ptr interface{}) { o.Forwarders = *ptr.(**IPv4Array) }
	if err := w.ReadPointer(&o.Forwarders, _s_aipForwarders, _ptr_aipForwarders); err != nil {
		return err
	}
	_ptr_pExtension1 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Extension1 == nil {
			o.Extension1 = &ServerExtension{}
		}
		if err := o.Extension1.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pExtension1 := func(ptr interface{}) { o.Extension1 = *ptr.(**ServerExtension) }
	if err := w.ReadPointer(&o.Extension1, _s_pExtension1, _ptr_pExtension1); err != nil {
		return err
	}
	_ptr_pExtension2 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Extension2 == nil {
			o.Extension2 = &ServerExtension{}
		}
		if err := o.Extension2.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pExtension2 := func(ptr interface{}) { o.Extension2 = *ptr.(**ServerExtension) }
	if err := w.ReadPointer(&o.Extension2, _s_pExtension2, _ptr_pExtension2); err != nil {
		return err
	}
	_ptr_pExtension3 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Extension3 == nil {
			o.Extension3 = &ServerExtension{}
		}
		if err := o.Extension3.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pExtension3 := func(ptr interface{}) { o.Extension3 = *ptr.(**ServerExtension) }
	if err := w.ReadPointer(&o.Extension3, _s_pExtension3, _ptr_pExtension3); err != nil {
		return err
	}
	_ptr_pExtension4 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Extension4 == nil {
			o.Extension4 = &ServerExtension{}
		}
		if err := o.Extension4.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pExtension4 := func(ptr interface{}) { o.Extension4 = *ptr.(**ServerExtension) }
	if err := w.ReadPointer(&o.Extension4, _s_pExtension4, _ptr_pExtension4); err != nil {
		return err
	}
	_ptr_pExtension5 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Extension5 == nil {
			o.Extension5 = &ServerExtension{}
		}
		if err := o.Extension5.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pExtension5 := func(ptr interface{}) { o.Extension5 = *ptr.(**ServerExtension) }
	if err := w.ReadPointer(&o.Extension5, _s_pExtension5, _ptr_pExtension5); err != nil {
		return err
	}
	if err := w.ReadData(&o.LogLevel); err != nil {
		return err
	}
	if err := w.ReadData(&o.DebugLevel); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwardTimeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCProtocol); err != nil {
		return err
	}
	if err := w.ReadData(&o.NameCheckFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddressAnswerLimitCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.RecursionRetry); err != nil {
		return err
	}
	if err := w.ReadData(&o.RecursionTimeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxCacheTTL); err != nil {
		return err
	}
	if err := w.ReadData(&o.DSPollingInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.ScavengingInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.DefaultRefreshInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.DefaultNoRefreshInterval); err != nil {
		return err
	}
	// reserved dwReserveArray
	var _dwReserveArray []uint32
	_dwReserveArray = make([]uint32, 10)
	for i1 := range _dwReserveArray {
		i1 := i1
		if err := w.ReadData(&_dwReserveArray[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.AutoReverseZones); err != nil {
		return err
	}
	if err := w.ReadData(&o.AutoCacheUpdate); err != nil {
		return err
	}
	if err := w.ReadData(&o.RecurseAfterForwarding); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwardDelegations); err != nil {
		return err
	}
	if err := w.ReadData(&o.NoRecursion); err != nil {
		return err
	}
	if err := w.ReadData(&o.SecureResponses); err != nil {
		return err
	}
	if err := w.ReadData(&o.RoundRobin); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocalNetPriority); err != nil {
		return err
	}
	if err := w.ReadData(&o.BindSecondaries); err != nil {
		return err
	}
	if err := w.ReadData(&o.WriteAuthorityNS); err != nil {
		return err
	}
	if err := w.ReadData(&o.StrictFileParsing); err != nil {
		return err
	}
	if err := w.ReadData(&o.LooseWildcarding); err != nil {
		return err
	}
	if err := w.ReadData(&o.DefaultAgingState); err != nil {
		return err
	}
	// reserved fReserveArray
	var _fReserveArray []bool
	_fReserveArray = make([]bool, 15)
	for i1 := range _fReserveArray {
		i1 := i1
		if err := w.ReadData(&_fReserveArray[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// ServerInfo_NET structure represents DNS_RPC_SERVER_INFO_DOTNET RPC structure.
//
// In the DNS_RPC_SERVER_INFO_DOTNET structure all fields have same definition as specified
// in DNS_RPC_SERVER_INFO_W2K (section 2.2.4.2.2.1), with the following exceptions:<32>
type ServerInfo_NET struct {
	// dwRpcStructureVersion: The DNS management structure version number. The following
	// are possible values:
	//
	//	+------------+--------------------------------------------------------------------------+
	//	|            |                                                                          |
	//	|   VALUE    |                                 MEANING                                  |
	//	|            |                                                                          |
	//	+------------+--------------------------------------------------------------------------+
	//	+------------+--------------------------------------------------------------------------+
	//	| 0x00000001 | Structure is of type DNS_RPC_SERVER_INFO_DOTNET (section 2.2.4.2.2.2).   |
	//	+------------+--------------------------------------------------------------------------+
	//	| 0x00000002 | Structure is of type DNS_RPC_SERVER_INFO_LONGHORN (section 2.2.4.2.2.3). |
	//	+------------+--------------------------------------------------------------------------+
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// dwReserved0: This field is reserved for future use. Senders MUST set this to zero
	// and it MUST be ignored by receiver.
	_               uint32     `idl:"name:dwReserved0"`
	Version         uint32     `idl:"name:dwVersion" json:"version"`
	BootMethod      uint8      `idl:"name:fBootMethod" json:"boot_method"`
	AdminConfigured bool       `idl:"name:fAdminConfigured" json:"admin_configured"`
	AllowUpdate     bool       `idl:"name:fAllowUpdate" json:"allow_update"`
	DSAvailable     bool       `idl:"name:fDsAvailable" json:"ds_available"`
	ServerName      string     `idl:"name:pszServerName;string" json:"server_name"`
	DSContainer     string     `idl:"name:pszDsContainer;string" json:"ds_container"`
	ServerAddrs     *IPv4Array `idl:"name:aipServerAddrs" json:"server_addrs"`
	ListenAddrs     *IPv4Array `idl:"name:aipListenAddrs" json:"listen_addrs"`
	Forwarders      *IPv4Array `idl:"name:aipForwarders" json:"forwarders"`
	// aipLogFilter: The list of IP addresses used for debug log filtering. The DNS server
	// will log DNS traffic that is sent to or received from for these IP addresses and
	// will not log DNS traffic to or from other IP addresses. If this value is set to NULL
	// then the DNS server will not perform IP filtering when logging DNS traffic.
	LogFilter *IPv4Array `idl:"name:aipLogFilter" json:"log_filter"`
	// pwszLogFilePath: A pointer to a null-terminated Unicode string that contains an absolute
	// pathname or relative pathname or filename for the operational log file on the DNS
	// server. If this value is set to NULL, the log SHOULD be logged to an implementation
	// specific log file.
	LogFilePath string `idl:"name:pwszLogFilePath;string" json:"log_file_path"`
	// pszDomainName: A pointer to a null-terminated UTF-8 string that contains the name
	// of the directory server domain to which the DNS server belongs if directory server
	// is available. This value will be NULL if no directory server is available.
	DomainName string `idl:"name:pszDomainName;string" json:"domain_name"`
	// pszForestName: A pointer to a null-terminated UTF-8 string that contains the name
	// of the directory server forest to which the DNS server belongs if directory server
	// is available. This value will be NULL if no directory server is available.
	ForestName string `idl:"name:pszForestName;string" json:"forest_name"`
	// pszDomainDirectoryPartition: A pointer to a null-terminated UTF-8 string that contains
	// the base name for the domain wide DNS application directory partition.
	DomainDirectoryPartition string `idl:"name:pszDomainDirectoryPartition;string" json:"domain_directory_partition"`
	// pszForestDirectoryPartition: A pointer to a null-terminated UTF-8 string that contains
	// the base name for the forest-wide DNS application directory partition.
	ForestDirectoryPartition string `idl:"name:pszForestDirectoryPartition;string" json:"forest_directory_partition"`
	// pExtensions: Reserved for future use. Senders MUST set this to NULL and it MUST be
	// ignored by the receiver.
	Extensions              string `idl:"name:pExtensions;string" json:"extensions"`
	LogLevel                uint32 `idl:"name:dwLogLevel" json:"log_level"`
	DebugLevel              uint32 `idl:"name:dwDebugLevel" json:"debug_level"`
	ForwardTimeout          uint32 `idl:"name:dwForwardTimeout" json:"forward_timeout"`
	RPCProtocol             uint32 `idl:"name:dwRpcProtocol" json:"rpc_protocol"`
	NameCheckFlag           uint32 `idl:"name:dwNameCheckFlag" json:"name_check_flag"`
	AddressAnswerLimitCount uint32 `idl:"name:cAddressAnswerLimit" json:"address_answer_limit_count"`
	RecursionRetry          uint32 `idl:"name:dwRecursionRetry" json:"recursion_retry"`
	RecursionTimeout        uint32 `idl:"name:dwRecursionTimeout" json:"recursion_timeout"`
	MaxCacheTTL             uint32 `idl:"name:dwMaxCacheTtl" json:"max_cache_ttl"`
	DSPollingInterval       uint32 `idl:"name:dwDsPollingInterval" json:"ds_polling_interval"`
	// dwLocalNetPriorityNetMask: The net mask used by the DNS server to prioritize address
	// records in responses when the server is configured to enforce LocalNetPriority as
	// mentioned above.
	LocalNetPriorityNetMask  uint32 `idl:"name:dwLocalNetPriorityNetMask" json:"local_net_priority_net_mask"`
	ScavengingInterval       uint32 `idl:"name:dwScavengingInterval" json:"scavenging_interval"`
	DefaultRefreshInterval   uint32 `idl:"name:dwDefaultRefreshInterval" json:"default_refresh_interval"`
	DefaultNoRefreshInterval uint32 `idl:"name:dwDefaultNoRefreshInterval" json:"default_no_refresh_interval"`
	// dwLastScavengeTime: The timestamp at which the last scavenging cycle was executed
	// on the DNS server. If this value is set to 0 then no scavenging cycle has been run
	// since the server was last started.
	LastScavengeTime uint32 `idl:"name:dwLastScavengeTime" json:"last_scavenge_time"`
	// dwEventLogLevel: This value indicates what level of events will be logged by the
	// DNS server. This value MUST be set to one of the combination of the possible values
	// for this defined in DNS_EVENTLOG_TYPES (section 2.2.9.1.2).
	EventLogLevel uint32 `idl:"name:dwEventLogLevel" json:"event_log_level"`
	// dwLogFileMaxSize: The maximum allowed size, in bytes, for the log file.
	LogFileMaxSize uint32 `idl:"name:dwLogFileMaxSize" json:"log_file_max_size"`
	// dwDsForestVersion: This value indicates the directory server forest version being
	// used by the DNS server, stored in the ForceForestBehaviorVersion property.
	DSForestVersion uint32 `idl:"name:dwDsForestVersion" json:"ds_forest_version"`
	// dwDsDomainVersion: This value indicates the directory server domain version being
	// used by the DNS server, stored in the ForceDomainBehaviorVersion property.
	DSDomainVersion uint32 `idl:"name:dwDsDomainVersion" json:"ds_domain_version"`
	// dwDsDsaVersion: This value indicates the directory server local server version being
	// used by the DNS server, stored in the ForceDsaBehaviorVersion property.
	DSDSAVersion           uint32   `idl:"name:dwDsDsaVersion" json:"ds_dsa_version"`
	_                      []uint32 `idl:"name:dwReserveArray"`
	AutoReverseZones       bool     `idl:"name:fAutoReverseZones" json:"auto_reverse_zones"`
	AutoCacheUpdate        bool     `idl:"name:fAutoCacheUpdate" json:"auto_cache_update"`
	RecurseAfterForwarding bool     `idl:"name:fRecurseAfterForwarding" json:"recurse_after_forwarding"`
	ForwardDelegations     bool     `idl:"name:fForwardDelegations" json:"forward_delegations"`
	NoRecursion            bool     `idl:"name:fNoRecursion" json:"no_recursion"`
	SecureResponses        bool     `idl:"name:fSecureResponses" json:"secure_responses"`
	RoundRobin             bool     `idl:"name:fRoundRobin" json:"round_robin"`
	LocalNetPriority       bool     `idl:"name:fLocalNetPriority" json:"local_net_priority"`
	BindSecondaries        bool     `idl:"name:fBindSecondaries" json:"bind_secondaries"`
	WriteAuthorityNS       bool     `idl:"name:fWriteAuthorityNs" json:"write_authority_ns"`
	StrictFileParsing      bool     `idl:"name:fStrictFileParsing" json:"strict_file_parsing"`
	LooseWildcarding       bool     `idl:"name:fLooseWildcarding" json:"loose_wildcarding"`
	DefaultAgingState      bool     `idl:"name:fDefaultAgingState" json:"default_aging_state"`
	_                      []bool   `idl:"name:fReserveArray"`
}

func (o *ServerInfo_NET) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ServerInfo_NET) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	if err := w.WriteData(o.BootMethod); err != nil {
		return err
	}
	if err := w.WriteData(o.AdminConfigured); err != nil {
		return err
	}
	if err := w.WriteData(o.AllowUpdate); err != nil {
		return err
	}
	if err := w.WriteData(o.DSAvailable); err != nil {
		return err
	}
	if o.ServerName != "" {
		_ptr_pszServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.ServerName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ServerName, _ptr_pszServerName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DSContainer != "" {
		_ptr_pszDsContainer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DSContainer); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DSContainer, _ptr_pszDsContainer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ServerAddrs != nil {
		_ptr_aipServerAddrs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ServerAddrs != nil {
				if err := o.ServerAddrs.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPv4Array{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ServerAddrs, _ptr_aipServerAddrs); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ListenAddrs != nil {
		_ptr_aipListenAddrs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ListenAddrs != nil {
				if err := o.ListenAddrs.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPv4Array{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ListenAddrs, _ptr_aipListenAddrs); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Forwarders != nil {
		_ptr_aipForwarders := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Forwarders != nil {
				if err := o.Forwarders.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPv4Array{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Forwarders, _ptr_aipForwarders); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LogFilter != nil {
		_ptr_aipLogFilter := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.LogFilter != nil {
				if err := o.LogFilter.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPv4Array{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.LogFilter, _ptr_aipLogFilter); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LogFilePath != "" {
		_ptr_pwszLogFilePath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LogFilePath); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LogFilePath, _ptr_pwszLogFilePath); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DomainName != "" {
		_ptr_pszDomainName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.DomainName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DomainName, _ptr_pszDomainName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ForestName != "" {
		_ptr_pszForestName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.ForestName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ForestName, _ptr_pszForestName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DomainDirectoryPartition != "" {
		_ptr_pszDomainDirectoryPartition := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.DomainDirectoryPartition); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DomainDirectoryPartition, _ptr_pszDomainDirectoryPartition); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ForestDirectoryPartition != "" {
		_ptr_pszForestDirectoryPartition := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.ForestDirectoryPartition); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ForestDirectoryPartition, _ptr_pszForestDirectoryPartition); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Extensions != "" {
		_ptr_pExtensions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.Extensions); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Extensions, _ptr_pExtensions); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LogLevel); err != nil {
		return err
	}
	if err := w.WriteData(o.DebugLevel); err != nil {
		return err
	}
	if err := w.WriteData(o.ForwardTimeout); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCProtocol); err != nil {
		return err
	}
	if err := w.WriteData(o.NameCheckFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.AddressAnswerLimitCount); err != nil {
		return err
	}
	if err := w.WriteData(o.RecursionRetry); err != nil {
		return err
	}
	if err := w.WriteData(o.RecursionTimeout); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxCacheTTL); err != nil {
		return err
	}
	if err := w.WriteData(o.DSPollingInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.LocalNetPriorityNetMask); err != nil {
		return err
	}
	if err := w.WriteData(o.ScavengingInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.DefaultRefreshInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.DefaultNoRefreshInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.LastScavengeTime); err != nil {
		return err
	}
	if err := w.WriteData(o.EventLogLevel); err != nil {
		return err
	}
	if err := w.WriteData(o.LogFileMaxSize); err != nil {
		return err
	}
	if err := w.WriteData(o.DSForestVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.DSDomainVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.DSDSAVersion); err != nil {
		return err
	}
	// reserved dwReserveArray
	for i1 := 0; uint64(i1) < 4; i1++ {
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.AutoReverseZones); err != nil {
		return err
	}
	if err := w.WriteData(o.AutoCacheUpdate); err != nil {
		return err
	}
	if err := w.WriteData(o.RecurseAfterForwarding); err != nil {
		return err
	}
	if err := w.WriteData(o.ForwardDelegations); err != nil {
		return err
	}
	if err := w.WriteData(o.NoRecursion); err != nil {
		return err
	}
	if err := w.WriteData(o.SecureResponses); err != nil {
		return err
	}
	if err := w.WriteData(o.RoundRobin); err != nil {
		return err
	}
	if err := w.WriteData(o.LocalNetPriority); err != nil {
		return err
	}
	if err := w.WriteData(o.BindSecondaries); err != nil {
		return err
	}
	if err := w.WriteData(o.WriteAuthorityNS); err != nil {
		return err
	}
	if err := w.WriteData(o.StrictFileParsing); err != nil {
		return err
	}
	if err := w.WriteData(o.LooseWildcarding); err != nil {
		return err
	}
	if err := w.WriteData(o.DefaultAgingState); err != nil {
		return err
	}
	// reserved fReserveArray
	for i1 := 0; uint64(i1) < 15; i1++ {
		if err := w.WriteData(false); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *ServerInfo_NET) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	if err := w.ReadData(&o.BootMethod); err != nil {
		return err
	}
	if err := w.ReadData(&o.AdminConfigured); err != nil {
		return err
	}
	if err := w.ReadData(&o.AllowUpdate); err != nil {
		return err
	}
	if err := w.ReadData(&o.DSAvailable); err != nil {
		return err
	}
	_ptr_pszServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.ServerName); err != nil {
			return err
		}
		return nil
	})
	_s_pszServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ServerName, _s_pszServerName, _ptr_pszServerName); err != nil {
		return err
	}
	_ptr_pszDsContainer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DSContainer); err != nil {
			return err
		}
		return nil
	})
	_s_pszDsContainer := func(ptr interface{}) { o.DSContainer = *ptr.(*string) }
	if err := w.ReadPointer(&o.DSContainer, _s_pszDsContainer, _ptr_pszDsContainer); err != nil {
		return err
	}
	_ptr_aipServerAddrs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ServerAddrs == nil {
			o.ServerAddrs = &IPv4Array{}
		}
		if err := o.ServerAddrs.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipServerAddrs := func(ptr interface{}) { o.ServerAddrs = *ptr.(**IPv4Array) }
	if err := w.ReadPointer(&o.ServerAddrs, _s_aipServerAddrs, _ptr_aipServerAddrs); err != nil {
		return err
	}
	_ptr_aipListenAddrs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ListenAddrs == nil {
			o.ListenAddrs = &IPv4Array{}
		}
		if err := o.ListenAddrs.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipListenAddrs := func(ptr interface{}) { o.ListenAddrs = *ptr.(**IPv4Array) }
	if err := w.ReadPointer(&o.ListenAddrs, _s_aipListenAddrs, _ptr_aipListenAddrs); err != nil {
		return err
	}
	_ptr_aipForwarders := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Forwarders == nil {
			o.Forwarders = &IPv4Array{}
		}
		if err := o.Forwarders.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipForwarders := func(ptr interface{}) { o.Forwarders = *ptr.(**IPv4Array) }
	if err := w.ReadPointer(&o.Forwarders, _s_aipForwarders, _ptr_aipForwarders); err != nil {
		return err
	}
	_ptr_aipLogFilter := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.LogFilter == nil {
			o.LogFilter = &IPv4Array{}
		}
		if err := o.LogFilter.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipLogFilter := func(ptr interface{}) { o.LogFilter = *ptr.(**IPv4Array) }
	if err := w.ReadPointer(&o.LogFilter, _s_aipLogFilter, _ptr_aipLogFilter); err != nil {
		return err
	}
	_ptr_pwszLogFilePath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LogFilePath); err != nil {
			return err
		}
		return nil
	})
	_s_pwszLogFilePath := func(ptr interface{}) { o.LogFilePath = *ptr.(*string) }
	if err := w.ReadPointer(&o.LogFilePath, _s_pwszLogFilePath, _ptr_pwszLogFilePath); err != nil {
		return err
	}
	_ptr_pszDomainName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.DomainName); err != nil {
			return err
		}
		return nil
	})
	_s_pszDomainName := func(ptr interface{}) { o.DomainName = *ptr.(*string) }
	if err := w.ReadPointer(&o.DomainName, _s_pszDomainName, _ptr_pszDomainName); err != nil {
		return err
	}
	_ptr_pszForestName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.ForestName); err != nil {
			return err
		}
		return nil
	})
	_s_pszForestName := func(ptr interface{}) { o.ForestName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ForestName, _s_pszForestName, _ptr_pszForestName); err != nil {
		return err
	}
	_ptr_pszDomainDirectoryPartition := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.DomainDirectoryPartition); err != nil {
			return err
		}
		return nil
	})
	_s_pszDomainDirectoryPartition := func(ptr interface{}) { o.DomainDirectoryPartition = *ptr.(*string) }
	if err := w.ReadPointer(&o.DomainDirectoryPartition, _s_pszDomainDirectoryPartition, _ptr_pszDomainDirectoryPartition); err != nil {
		return err
	}
	_ptr_pszForestDirectoryPartition := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.ForestDirectoryPartition); err != nil {
			return err
		}
		return nil
	})
	_s_pszForestDirectoryPartition := func(ptr interface{}) { o.ForestDirectoryPartition = *ptr.(*string) }
	if err := w.ReadPointer(&o.ForestDirectoryPartition, _s_pszForestDirectoryPartition, _ptr_pszForestDirectoryPartition); err != nil {
		return err
	}
	_ptr_pExtensions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.Extensions); err != nil {
			return err
		}
		return nil
	})
	_s_pExtensions := func(ptr interface{}) { o.Extensions = *ptr.(*string) }
	if err := w.ReadPointer(&o.Extensions, _s_pExtensions, _ptr_pExtensions); err != nil {
		return err
	}
	if err := w.ReadData(&o.LogLevel); err != nil {
		return err
	}
	if err := w.ReadData(&o.DebugLevel); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwardTimeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCProtocol); err != nil {
		return err
	}
	if err := w.ReadData(&o.NameCheckFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddressAnswerLimitCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.RecursionRetry); err != nil {
		return err
	}
	if err := w.ReadData(&o.RecursionTimeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxCacheTTL); err != nil {
		return err
	}
	if err := w.ReadData(&o.DSPollingInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocalNetPriorityNetMask); err != nil {
		return err
	}
	if err := w.ReadData(&o.ScavengingInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.DefaultRefreshInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.DefaultNoRefreshInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastScavengeTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.EventLogLevel); err != nil {
		return err
	}
	if err := w.ReadData(&o.LogFileMaxSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.DSForestVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.DSDomainVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.DSDSAVersion); err != nil {
		return err
	}
	// reserved dwReserveArray
	var _dwReserveArray []uint32
	_dwReserveArray = make([]uint32, 4)
	for i1 := range _dwReserveArray {
		i1 := i1
		if err := w.ReadData(&_dwReserveArray[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.AutoReverseZones); err != nil {
		return err
	}
	if err := w.ReadData(&o.AutoCacheUpdate); err != nil {
		return err
	}
	if err := w.ReadData(&o.RecurseAfterForwarding); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwardDelegations); err != nil {
		return err
	}
	if err := w.ReadData(&o.NoRecursion); err != nil {
		return err
	}
	if err := w.ReadData(&o.SecureResponses); err != nil {
		return err
	}
	if err := w.ReadData(&o.RoundRobin); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocalNetPriority); err != nil {
		return err
	}
	if err := w.ReadData(&o.BindSecondaries); err != nil {
		return err
	}
	if err := w.ReadData(&o.WriteAuthorityNS); err != nil {
		return err
	}
	if err := w.ReadData(&o.StrictFileParsing); err != nil {
		return err
	}
	if err := w.ReadData(&o.LooseWildcarding); err != nil {
		return err
	}
	if err := w.ReadData(&o.DefaultAgingState); err != nil {
		return err
	}
	// reserved fReserveArray
	var _fReserveArray []bool
	_fReserveArray = make([]bool, 15)
	for i1 := range _fReserveArray {
		i1 := i1
		if err := w.ReadData(&_fReserveArray[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// ServerInfoLonghorn structure represents DNS_RPC_SERVER_INFO_LONGHORN RPC structure.
//
// In the DNS_RPC_SERVER_INFO_LONGHORN structure all fields have same definition as
// specified in section DNS_RPC_SERVER_INFO_DOTNET (section 2.2.4.2.2.2 ), with the
// following exceptions:<33>
type ServerInfoLonghorn struct {
	RPCStructureVersion      uint32     `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	_                        uint32     `idl:"name:dwReserved0"`
	Version                  uint32     `idl:"name:dwVersion" json:"version"`
	BootMethod               uint8      `idl:"name:fBootMethod" json:"boot_method"`
	AdminConfigured          bool       `idl:"name:fAdminConfigured" json:"admin_configured"`
	AllowUpdate              bool       `idl:"name:fAllowUpdate" json:"allow_update"`
	DSAvailable              bool       `idl:"name:fDsAvailable" json:"ds_available"`
	ServerName               string     `idl:"name:pszServerName;string" json:"server_name"`
	DSContainer              string     `idl:"name:pszDsContainer;string" json:"ds_container"`
	ServerAddrs              *AddrArray `idl:"name:aipServerAddrs" json:"server_addrs"`
	ListenAddrs              *AddrArray `idl:"name:aipListenAddrs" json:"listen_addrs"`
	Forwarders               *AddrArray `idl:"name:aipForwarders" json:"forwarders"`
	LogFilter                *AddrArray `idl:"name:aipLogFilter" json:"log_filter"`
	LogFilePath              string     `idl:"name:pwszLogFilePath;string" json:"log_file_path"`
	DomainName               string     `idl:"name:pszDomainName;string" json:"domain_name"`
	ForestName               string     `idl:"name:pszForestName;string" json:"forest_name"`
	DomainDirectoryPartition string     `idl:"name:pszDomainDirectoryPartition;string" json:"domain_directory_partition"`
	ForestDirectoryPartition string     `idl:"name:pszForestDirectoryPartition;string" json:"forest_directory_partition"`
	Extensions               string     `idl:"name:pExtensions;string" json:"extensions"`
	LogLevel                 uint32     `idl:"name:dwLogLevel" json:"log_level"`
	DebugLevel               uint32     `idl:"name:dwDebugLevel" json:"debug_level"`
	ForwardTimeout           uint32     `idl:"name:dwForwardTimeout" json:"forward_timeout"`
	RPCProtocol              uint32     `idl:"name:dwRpcProtocol" json:"rpc_protocol"`
	NameCheckFlag            uint32     `idl:"name:dwNameCheckFlag" json:"name_check_flag"`
	AddressAnswerLimitCount  uint32     `idl:"name:cAddressAnswerLimit" json:"address_answer_limit_count"`
	RecursionRetry           uint32     `idl:"name:dwRecursionRetry" json:"recursion_retry"`
	RecursionTimeout         uint32     `idl:"name:dwRecursionTimeout" json:"recursion_timeout"`
	MaxCacheTTL              uint32     `idl:"name:dwMaxCacheTtl" json:"max_cache_ttl"`
	DSPollingInterval        uint32     `idl:"name:dwDsPollingInterval" json:"ds_polling_interval"`
	LocalNetPriorityNetMask  uint32     `idl:"name:dwLocalNetPriorityNetMask" json:"local_net_priority_net_mask"`
	ScavengingInterval       uint32     `idl:"name:dwScavengingInterval" json:"scavenging_interval"`
	DefaultRefreshInterval   uint32     `idl:"name:dwDefaultRefreshInterval" json:"default_refresh_interval"`
	DefaultNoRefreshInterval uint32     `idl:"name:dwDefaultNoRefreshInterval" json:"default_no_refresh_interval"`
	LastScavengeTime         uint32     `idl:"name:dwLastScavengeTime" json:"last_scavenge_time"`
	EventLogLevel            uint32     `idl:"name:dwEventLogLevel" json:"event_log_level"`
	LogFileMaxSize           uint32     `idl:"name:dwLogFileMaxSize" json:"log_file_max_size"`
	DSForestVersion          uint32     `idl:"name:dwDsForestVersion" json:"ds_forest_version"`
	DSDomainVersion          uint32     `idl:"name:dwDsDomainVersion" json:"ds_domain_version"`
	DSDSAVersion             uint32     `idl:"name:dwDsDsaVersion" json:"ds_dsa_version"`
	// fReadOnlyDC: A Boolean value that indicates whether the DNS server has access to
	// a directory server that is running in read-only mode, that is, whether the server
	// does not accept directory server write operations. The DNS server detects whether
	// this is the case by reading the supportedCapabilities attribute of the server's rootDse
	// object, looking for LDAP_CAP_ACTIVE_DIRECTORY_PARTIAL_SECRETS_OID (See [MS-ADTS]
	// section 3.1.1.3.4.3.6).
	ReadOnlyDC             bool     `idl:"name:fReadOnlyDC" json:"read_only_dc"`
	_                      []uint32 `idl:"name:dwReserveArray"`
	AutoReverseZones       bool     `idl:"name:fAutoReverseZones" json:"auto_reverse_zones"`
	AutoCacheUpdate        bool     `idl:"name:fAutoCacheUpdate" json:"auto_cache_update"`
	RecurseAfterForwarding bool     `idl:"name:fRecurseAfterForwarding" json:"recurse_after_forwarding"`
	ForwardDelegations     bool     `idl:"name:fForwardDelegations" json:"forward_delegations"`
	NoRecursion            bool     `idl:"name:fNoRecursion" json:"no_recursion"`
	SecureResponses        bool     `idl:"name:fSecureResponses" json:"secure_responses"`
	RoundRobin             bool     `idl:"name:fRoundRobin" json:"round_robin"`
	LocalNetPriority       bool     `idl:"name:fLocalNetPriority" json:"local_net_priority"`
	BindSecondaries        bool     `idl:"name:fBindSecondaries" json:"bind_secondaries"`
	WriteAuthorityNS       bool     `idl:"name:fWriteAuthorityNs" json:"write_authority_ns"`
	StrictFileParsing      bool     `idl:"name:fStrictFileParsing" json:"strict_file_parsing"`
	LooseWildcarding       bool     `idl:"name:fLooseWildcarding" json:"loose_wildcarding"`
	DefaultAgingState      bool     `idl:"name:fDefaultAgingState" json:"default_aging_state"`
	_                      []bool   `idl:"name:fReserveArray"`
}

func (o *ServerInfoLonghorn) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ServerInfoLonghorn) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	if err := w.WriteData(o.BootMethod); err != nil {
		return err
	}
	if err := w.WriteData(o.AdminConfigured); err != nil {
		return err
	}
	if err := w.WriteData(o.AllowUpdate); err != nil {
		return err
	}
	if err := w.WriteData(o.DSAvailable); err != nil {
		return err
	}
	if o.ServerName != "" {
		_ptr_pszServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.ServerName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ServerName, _ptr_pszServerName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DSContainer != "" {
		_ptr_pszDsContainer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DSContainer); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DSContainer, _ptr_pszDsContainer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ServerAddrs != nil {
		_ptr_aipServerAddrs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ServerAddrs != nil {
				if err := o.ServerAddrs.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AddrArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ServerAddrs, _ptr_aipServerAddrs); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ListenAddrs != nil {
		_ptr_aipListenAddrs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ListenAddrs != nil {
				if err := o.ListenAddrs.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AddrArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ListenAddrs, _ptr_aipListenAddrs); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Forwarders != nil {
		_ptr_aipForwarders := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Forwarders != nil {
				if err := o.Forwarders.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AddrArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Forwarders, _ptr_aipForwarders); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LogFilter != nil {
		_ptr_aipLogFilter := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.LogFilter != nil {
				if err := o.LogFilter.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AddrArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.LogFilter, _ptr_aipLogFilter); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LogFilePath != "" {
		_ptr_pwszLogFilePath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LogFilePath); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LogFilePath, _ptr_pwszLogFilePath); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DomainName != "" {
		_ptr_pszDomainName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.DomainName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DomainName, _ptr_pszDomainName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ForestName != "" {
		_ptr_pszForestName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.ForestName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ForestName, _ptr_pszForestName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DomainDirectoryPartition != "" {
		_ptr_pszDomainDirectoryPartition := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.DomainDirectoryPartition); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DomainDirectoryPartition, _ptr_pszDomainDirectoryPartition); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ForestDirectoryPartition != "" {
		_ptr_pszForestDirectoryPartition := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.ForestDirectoryPartition); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ForestDirectoryPartition, _ptr_pszForestDirectoryPartition); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Extensions != "" {
		_ptr_pExtensions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.Extensions); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Extensions, _ptr_pExtensions); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LogLevel); err != nil {
		return err
	}
	if err := w.WriteData(o.DebugLevel); err != nil {
		return err
	}
	if err := w.WriteData(o.ForwardTimeout); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCProtocol); err != nil {
		return err
	}
	if err := w.WriteData(o.NameCheckFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.AddressAnswerLimitCount); err != nil {
		return err
	}
	if err := w.WriteData(o.RecursionRetry); err != nil {
		return err
	}
	if err := w.WriteData(o.RecursionTimeout); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxCacheTTL); err != nil {
		return err
	}
	if err := w.WriteData(o.DSPollingInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.LocalNetPriorityNetMask); err != nil {
		return err
	}
	if err := w.WriteData(o.ScavengingInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.DefaultRefreshInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.DefaultNoRefreshInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.LastScavengeTime); err != nil {
		return err
	}
	if err := w.WriteData(o.EventLogLevel); err != nil {
		return err
	}
	if err := w.WriteData(o.LogFileMaxSize); err != nil {
		return err
	}
	if err := w.WriteData(o.DSForestVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.DSDomainVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.DSDSAVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.ReadOnlyDC); err != nil {
		return err
	}
	// reserved dwReserveArray
	for i1 := 0; uint64(i1) < 3; i1++ {
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.AutoReverseZones); err != nil {
		return err
	}
	if err := w.WriteData(o.AutoCacheUpdate); err != nil {
		return err
	}
	if err := w.WriteData(o.RecurseAfterForwarding); err != nil {
		return err
	}
	if err := w.WriteData(o.ForwardDelegations); err != nil {
		return err
	}
	if err := w.WriteData(o.NoRecursion); err != nil {
		return err
	}
	if err := w.WriteData(o.SecureResponses); err != nil {
		return err
	}
	if err := w.WriteData(o.RoundRobin); err != nil {
		return err
	}
	if err := w.WriteData(o.LocalNetPriority); err != nil {
		return err
	}
	if err := w.WriteData(o.BindSecondaries); err != nil {
		return err
	}
	if err := w.WriteData(o.WriteAuthorityNS); err != nil {
		return err
	}
	if err := w.WriteData(o.StrictFileParsing); err != nil {
		return err
	}
	if err := w.WriteData(o.LooseWildcarding); err != nil {
		return err
	}
	if err := w.WriteData(o.DefaultAgingState); err != nil {
		return err
	}
	// reserved fReserveArray
	for i1 := 0; uint64(i1) < 15; i1++ {
		if err := w.WriteData(false); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *ServerInfoLonghorn) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	if err := w.ReadData(&o.BootMethod); err != nil {
		return err
	}
	if err := w.ReadData(&o.AdminConfigured); err != nil {
		return err
	}
	if err := w.ReadData(&o.AllowUpdate); err != nil {
		return err
	}
	if err := w.ReadData(&o.DSAvailable); err != nil {
		return err
	}
	_ptr_pszServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.ServerName); err != nil {
			return err
		}
		return nil
	})
	_s_pszServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ServerName, _s_pszServerName, _ptr_pszServerName); err != nil {
		return err
	}
	_ptr_pszDsContainer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DSContainer); err != nil {
			return err
		}
		return nil
	})
	_s_pszDsContainer := func(ptr interface{}) { o.DSContainer = *ptr.(*string) }
	if err := w.ReadPointer(&o.DSContainer, _s_pszDsContainer, _ptr_pszDsContainer); err != nil {
		return err
	}
	_ptr_aipServerAddrs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ServerAddrs == nil {
			o.ServerAddrs = &AddrArray{}
		}
		if err := o.ServerAddrs.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipServerAddrs := func(ptr interface{}) { o.ServerAddrs = *ptr.(**AddrArray) }
	if err := w.ReadPointer(&o.ServerAddrs, _s_aipServerAddrs, _ptr_aipServerAddrs); err != nil {
		return err
	}
	_ptr_aipListenAddrs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ListenAddrs == nil {
			o.ListenAddrs = &AddrArray{}
		}
		if err := o.ListenAddrs.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipListenAddrs := func(ptr interface{}) { o.ListenAddrs = *ptr.(**AddrArray) }
	if err := w.ReadPointer(&o.ListenAddrs, _s_aipListenAddrs, _ptr_aipListenAddrs); err != nil {
		return err
	}
	_ptr_aipForwarders := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Forwarders == nil {
			o.Forwarders = &AddrArray{}
		}
		if err := o.Forwarders.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipForwarders := func(ptr interface{}) { o.Forwarders = *ptr.(**AddrArray) }
	if err := w.ReadPointer(&o.Forwarders, _s_aipForwarders, _ptr_aipForwarders); err != nil {
		return err
	}
	_ptr_aipLogFilter := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.LogFilter == nil {
			o.LogFilter = &AddrArray{}
		}
		if err := o.LogFilter.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipLogFilter := func(ptr interface{}) { o.LogFilter = *ptr.(**AddrArray) }
	if err := w.ReadPointer(&o.LogFilter, _s_aipLogFilter, _ptr_aipLogFilter); err != nil {
		return err
	}
	_ptr_pwszLogFilePath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LogFilePath); err != nil {
			return err
		}
		return nil
	})
	_s_pwszLogFilePath := func(ptr interface{}) { o.LogFilePath = *ptr.(*string) }
	if err := w.ReadPointer(&o.LogFilePath, _s_pwszLogFilePath, _ptr_pwszLogFilePath); err != nil {
		return err
	}
	_ptr_pszDomainName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.DomainName); err != nil {
			return err
		}
		return nil
	})
	_s_pszDomainName := func(ptr interface{}) { o.DomainName = *ptr.(*string) }
	if err := w.ReadPointer(&o.DomainName, _s_pszDomainName, _ptr_pszDomainName); err != nil {
		return err
	}
	_ptr_pszForestName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.ForestName); err != nil {
			return err
		}
		return nil
	})
	_s_pszForestName := func(ptr interface{}) { o.ForestName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ForestName, _s_pszForestName, _ptr_pszForestName); err != nil {
		return err
	}
	_ptr_pszDomainDirectoryPartition := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.DomainDirectoryPartition); err != nil {
			return err
		}
		return nil
	})
	_s_pszDomainDirectoryPartition := func(ptr interface{}) { o.DomainDirectoryPartition = *ptr.(*string) }
	if err := w.ReadPointer(&o.DomainDirectoryPartition, _s_pszDomainDirectoryPartition, _ptr_pszDomainDirectoryPartition); err != nil {
		return err
	}
	_ptr_pszForestDirectoryPartition := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.ForestDirectoryPartition); err != nil {
			return err
		}
		return nil
	})
	_s_pszForestDirectoryPartition := func(ptr interface{}) { o.ForestDirectoryPartition = *ptr.(*string) }
	if err := w.ReadPointer(&o.ForestDirectoryPartition, _s_pszForestDirectoryPartition, _ptr_pszForestDirectoryPartition); err != nil {
		return err
	}
	_ptr_pExtensions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.Extensions); err != nil {
			return err
		}
		return nil
	})
	_s_pExtensions := func(ptr interface{}) { o.Extensions = *ptr.(*string) }
	if err := w.ReadPointer(&o.Extensions, _s_pExtensions, _ptr_pExtensions); err != nil {
		return err
	}
	if err := w.ReadData(&o.LogLevel); err != nil {
		return err
	}
	if err := w.ReadData(&o.DebugLevel); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwardTimeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCProtocol); err != nil {
		return err
	}
	if err := w.ReadData(&o.NameCheckFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddressAnswerLimitCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.RecursionRetry); err != nil {
		return err
	}
	if err := w.ReadData(&o.RecursionTimeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxCacheTTL); err != nil {
		return err
	}
	if err := w.ReadData(&o.DSPollingInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocalNetPriorityNetMask); err != nil {
		return err
	}
	if err := w.ReadData(&o.ScavengingInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.DefaultRefreshInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.DefaultNoRefreshInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastScavengeTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.EventLogLevel); err != nil {
		return err
	}
	if err := w.ReadData(&o.LogFileMaxSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.DSForestVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.DSDomainVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.DSDSAVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReadOnlyDC); err != nil {
		return err
	}
	// reserved dwReserveArray
	var _dwReserveArray []uint32
	_dwReserveArray = make([]uint32, 3)
	for i1 := range _dwReserveArray {
		i1 := i1
		if err := w.ReadData(&_dwReserveArray[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.AutoReverseZones); err != nil {
		return err
	}
	if err := w.ReadData(&o.AutoCacheUpdate); err != nil {
		return err
	}
	if err := w.ReadData(&o.RecurseAfterForwarding); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwardDelegations); err != nil {
		return err
	}
	if err := w.ReadData(&o.NoRecursion); err != nil {
		return err
	}
	if err := w.ReadData(&o.SecureResponses); err != nil {
		return err
	}
	if err := w.ReadData(&o.RoundRobin); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocalNetPriority); err != nil {
		return err
	}
	if err := w.ReadData(&o.BindSecondaries); err != nil {
		return err
	}
	if err := w.ReadData(&o.WriteAuthorityNS); err != nil {
		return err
	}
	if err := w.ReadData(&o.StrictFileParsing); err != nil {
		return err
	}
	if err := w.ReadData(&o.LooseWildcarding); err != nil {
		return err
	}
	if err := w.ReadData(&o.DefaultAgingState); err != nil {
		return err
	}
	// reserved fReserveArray
	var _fReserveArray []bool
	_fReserveArray = make([]bool, 15)
	for i1 := range _fReserveArray {
		i1 := i1
		if err := w.ReadData(&_fReserveArray[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// ServerInfo structure represents DNS_RPC_SERVER_INFO RPC structure.
//
// The DNS_RPC_SERVER_INFO structure contains information about the DNS server's configuration
// and state. There are different versions of the DNS_RPC_SERVER_INFO structure. The
// DNS server MUST use one of the structures corresponding to the value of the dwClientVersion
// field in DNS Server Management Protocol method calls (section 3.1.4) as shown in
// the following table. If the method call does not specify the value of dwClientVersion,
// the DNS_RPC_SERVER_INFO_W2K version of the structure MUST be used.
//
//	+------------+----------------------------------------------------+
//	|            |                                                    |
//	|   VALUE    |                     STRUCTURE                      |
//	|            |                                                    |
//	+------------+----------------------------------------------------+
//	+------------+----------------------------------------------------+
//	| 0x00000000 | DNS_RPC_SERVER_INFO_W2K (section 2.2.4.2.2.1)      |
//	+------------+----------------------------------------------------+
//	| 0x00060000 | DNS_RPC_SERVER_INFO_DOTNET (section 2.2.4.2.2.2)   |
//	+------------+----------------------------------------------------+
//	| 0x00070000 | DNS_RPC_SERVER_INFO_LONGHORN (section 2.2.4.2.2.3) |
//	+------------+----------------------------------------------------+
type ServerInfo struct {
	RPCStructureVersion      uint32     `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	_                        uint32     `idl:"name:dwReserved0"`
	Version                  uint32     `idl:"name:dwVersion" json:"version"`
	BootMethod               uint8      `idl:"name:fBootMethod" json:"boot_method"`
	AdminConfigured          bool       `idl:"name:fAdminConfigured" json:"admin_configured"`
	AllowUpdate              bool       `idl:"name:fAllowUpdate" json:"allow_update"`
	DSAvailable              bool       `idl:"name:fDsAvailable" json:"ds_available"`
	ServerName               string     `idl:"name:pszServerName;string" json:"server_name"`
	DSContainer              string     `idl:"name:pszDsContainer;string" json:"ds_container"`
	ServerAddrs              *AddrArray `idl:"name:aipServerAddrs" json:"server_addrs"`
	ListenAddrs              *AddrArray `idl:"name:aipListenAddrs" json:"listen_addrs"`
	Forwarders               *AddrArray `idl:"name:aipForwarders" json:"forwarders"`
	LogFilter                *AddrArray `idl:"name:aipLogFilter" json:"log_filter"`
	LogFilePath              string     `idl:"name:pwszLogFilePath;string" json:"log_file_path"`
	DomainName               string     `idl:"name:pszDomainName;string" json:"domain_name"`
	ForestName               string     `idl:"name:pszForestName;string" json:"forest_name"`
	DomainDirectoryPartition string     `idl:"name:pszDomainDirectoryPartition;string" json:"domain_directory_partition"`
	ForestDirectoryPartition string     `idl:"name:pszForestDirectoryPartition;string" json:"forest_directory_partition"`
	Extensions               string     `idl:"name:pExtensions;string" json:"extensions"`
	LogLevel                 uint32     `idl:"name:dwLogLevel" json:"log_level"`
	DebugLevel               uint32     `idl:"name:dwDebugLevel" json:"debug_level"`
	ForwardTimeout           uint32     `idl:"name:dwForwardTimeout" json:"forward_timeout"`
	RPCProtocol              uint32     `idl:"name:dwRpcProtocol" json:"rpc_protocol"`
	NameCheckFlag            uint32     `idl:"name:dwNameCheckFlag" json:"name_check_flag"`
	AddressAnswerLimitCount  uint32     `idl:"name:cAddressAnswerLimit" json:"address_answer_limit_count"`
	RecursionRetry           uint32     `idl:"name:dwRecursionRetry" json:"recursion_retry"`
	RecursionTimeout         uint32     `idl:"name:dwRecursionTimeout" json:"recursion_timeout"`
	MaxCacheTTL              uint32     `idl:"name:dwMaxCacheTtl" json:"max_cache_ttl"`
	DSPollingInterval        uint32     `idl:"name:dwDsPollingInterval" json:"ds_polling_interval"`
	LocalNetPriorityNetMask  uint32     `idl:"name:dwLocalNetPriorityNetMask" json:"local_net_priority_net_mask"`
	ScavengingInterval       uint32     `idl:"name:dwScavengingInterval" json:"scavenging_interval"`
	DefaultRefreshInterval   uint32     `idl:"name:dwDefaultRefreshInterval" json:"default_refresh_interval"`
	DefaultNoRefreshInterval uint32     `idl:"name:dwDefaultNoRefreshInterval" json:"default_no_refresh_interval"`
	LastScavengeTime         uint32     `idl:"name:dwLastScavengeTime" json:"last_scavenge_time"`
	EventLogLevel            uint32     `idl:"name:dwEventLogLevel" json:"event_log_level"`
	LogFileMaxSize           uint32     `idl:"name:dwLogFileMaxSize" json:"log_file_max_size"`
	DSForestVersion          uint32     `idl:"name:dwDsForestVersion" json:"ds_forest_version"`
	DSDomainVersion          uint32     `idl:"name:dwDsDomainVersion" json:"ds_domain_version"`
	DSDSAVersion             uint32     `idl:"name:dwDsDsaVersion" json:"ds_dsa_version"`
	ReadOnlyDC               bool       `idl:"name:fReadOnlyDC" json:"read_only_dc"`
	_                        []uint32   `idl:"name:dwReserveArray"`
	AutoReverseZones         bool       `idl:"name:fAutoReverseZones" json:"auto_reverse_zones"`
	AutoCacheUpdate          bool       `idl:"name:fAutoCacheUpdate" json:"auto_cache_update"`
	RecurseAfterForwarding   bool       `idl:"name:fRecurseAfterForwarding" json:"recurse_after_forwarding"`
	ForwardDelegations       bool       `idl:"name:fForwardDelegations" json:"forward_delegations"`
	NoRecursion              bool       `idl:"name:fNoRecursion" json:"no_recursion"`
	SecureResponses          bool       `idl:"name:fSecureResponses" json:"secure_responses"`
	RoundRobin               bool       `idl:"name:fRoundRobin" json:"round_robin"`
	LocalNetPriority         bool       `idl:"name:fLocalNetPriority" json:"local_net_priority"`
	BindSecondaries          bool       `idl:"name:fBindSecondaries" json:"bind_secondaries"`
	WriteAuthorityNS         bool       `idl:"name:fWriteAuthorityNs" json:"write_authority_ns"`
	StrictFileParsing        bool       `idl:"name:fStrictFileParsing" json:"strict_file_parsing"`
	LooseWildcarding         bool       `idl:"name:fLooseWildcarding" json:"loose_wildcarding"`
	DefaultAgingState        bool       `idl:"name:fDefaultAgingState" json:"default_aging_state"`
	_                        []bool     `idl:"name:fReserveArray"`
}

func (o *ServerInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ServerInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	if err := w.WriteData(o.BootMethod); err != nil {
		return err
	}
	if err := w.WriteData(o.AdminConfigured); err != nil {
		return err
	}
	if err := w.WriteData(o.AllowUpdate); err != nil {
		return err
	}
	if err := w.WriteData(o.DSAvailable); err != nil {
		return err
	}
	if o.ServerName != "" {
		_ptr_pszServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.ServerName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ServerName, _ptr_pszServerName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DSContainer != "" {
		_ptr_pszDsContainer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DSContainer); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DSContainer, _ptr_pszDsContainer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ServerAddrs != nil {
		_ptr_aipServerAddrs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ServerAddrs != nil {
				if err := o.ServerAddrs.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AddrArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ServerAddrs, _ptr_aipServerAddrs); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ListenAddrs != nil {
		_ptr_aipListenAddrs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ListenAddrs != nil {
				if err := o.ListenAddrs.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AddrArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ListenAddrs, _ptr_aipListenAddrs); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Forwarders != nil {
		_ptr_aipForwarders := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Forwarders != nil {
				if err := o.Forwarders.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AddrArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Forwarders, _ptr_aipForwarders); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LogFilter != nil {
		_ptr_aipLogFilter := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.LogFilter != nil {
				if err := o.LogFilter.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AddrArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.LogFilter, _ptr_aipLogFilter); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LogFilePath != "" {
		_ptr_pwszLogFilePath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LogFilePath); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LogFilePath, _ptr_pwszLogFilePath); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DomainName != "" {
		_ptr_pszDomainName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.DomainName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DomainName, _ptr_pszDomainName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ForestName != "" {
		_ptr_pszForestName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.ForestName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ForestName, _ptr_pszForestName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DomainDirectoryPartition != "" {
		_ptr_pszDomainDirectoryPartition := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.DomainDirectoryPartition); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DomainDirectoryPartition, _ptr_pszDomainDirectoryPartition); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ForestDirectoryPartition != "" {
		_ptr_pszForestDirectoryPartition := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.ForestDirectoryPartition); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ForestDirectoryPartition, _ptr_pszForestDirectoryPartition); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Extensions != "" {
		_ptr_pExtensions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.Extensions); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Extensions, _ptr_pExtensions); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LogLevel); err != nil {
		return err
	}
	if err := w.WriteData(o.DebugLevel); err != nil {
		return err
	}
	if err := w.WriteData(o.ForwardTimeout); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCProtocol); err != nil {
		return err
	}
	if err := w.WriteData(o.NameCheckFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.AddressAnswerLimitCount); err != nil {
		return err
	}
	if err := w.WriteData(o.RecursionRetry); err != nil {
		return err
	}
	if err := w.WriteData(o.RecursionTimeout); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxCacheTTL); err != nil {
		return err
	}
	if err := w.WriteData(o.DSPollingInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.LocalNetPriorityNetMask); err != nil {
		return err
	}
	if err := w.WriteData(o.ScavengingInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.DefaultRefreshInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.DefaultNoRefreshInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.LastScavengeTime); err != nil {
		return err
	}
	if err := w.WriteData(o.EventLogLevel); err != nil {
		return err
	}
	if err := w.WriteData(o.LogFileMaxSize); err != nil {
		return err
	}
	if err := w.WriteData(o.DSForestVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.DSDomainVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.DSDSAVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.ReadOnlyDC); err != nil {
		return err
	}
	// reserved dwReserveArray
	for i1 := 0; uint64(i1) < 3; i1++ {
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.AutoReverseZones); err != nil {
		return err
	}
	if err := w.WriteData(o.AutoCacheUpdate); err != nil {
		return err
	}
	if err := w.WriteData(o.RecurseAfterForwarding); err != nil {
		return err
	}
	if err := w.WriteData(o.ForwardDelegations); err != nil {
		return err
	}
	if err := w.WriteData(o.NoRecursion); err != nil {
		return err
	}
	if err := w.WriteData(o.SecureResponses); err != nil {
		return err
	}
	if err := w.WriteData(o.RoundRobin); err != nil {
		return err
	}
	if err := w.WriteData(o.LocalNetPriority); err != nil {
		return err
	}
	if err := w.WriteData(o.BindSecondaries); err != nil {
		return err
	}
	if err := w.WriteData(o.WriteAuthorityNS); err != nil {
		return err
	}
	if err := w.WriteData(o.StrictFileParsing); err != nil {
		return err
	}
	if err := w.WriteData(o.LooseWildcarding); err != nil {
		return err
	}
	if err := w.WriteData(o.DefaultAgingState); err != nil {
		return err
	}
	// reserved fReserveArray
	for i1 := 0; uint64(i1) < 15; i1++ {
		if err := w.WriteData(false); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *ServerInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	if err := w.ReadData(&o.BootMethod); err != nil {
		return err
	}
	if err := w.ReadData(&o.AdminConfigured); err != nil {
		return err
	}
	if err := w.ReadData(&o.AllowUpdate); err != nil {
		return err
	}
	if err := w.ReadData(&o.DSAvailable); err != nil {
		return err
	}
	_ptr_pszServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.ServerName); err != nil {
			return err
		}
		return nil
	})
	_s_pszServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ServerName, _s_pszServerName, _ptr_pszServerName); err != nil {
		return err
	}
	_ptr_pszDsContainer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DSContainer); err != nil {
			return err
		}
		return nil
	})
	_s_pszDsContainer := func(ptr interface{}) { o.DSContainer = *ptr.(*string) }
	if err := w.ReadPointer(&o.DSContainer, _s_pszDsContainer, _ptr_pszDsContainer); err != nil {
		return err
	}
	_ptr_aipServerAddrs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ServerAddrs == nil {
			o.ServerAddrs = &AddrArray{}
		}
		if err := o.ServerAddrs.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipServerAddrs := func(ptr interface{}) { o.ServerAddrs = *ptr.(**AddrArray) }
	if err := w.ReadPointer(&o.ServerAddrs, _s_aipServerAddrs, _ptr_aipServerAddrs); err != nil {
		return err
	}
	_ptr_aipListenAddrs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ListenAddrs == nil {
			o.ListenAddrs = &AddrArray{}
		}
		if err := o.ListenAddrs.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipListenAddrs := func(ptr interface{}) { o.ListenAddrs = *ptr.(**AddrArray) }
	if err := w.ReadPointer(&o.ListenAddrs, _s_aipListenAddrs, _ptr_aipListenAddrs); err != nil {
		return err
	}
	_ptr_aipForwarders := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Forwarders == nil {
			o.Forwarders = &AddrArray{}
		}
		if err := o.Forwarders.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipForwarders := func(ptr interface{}) { o.Forwarders = *ptr.(**AddrArray) }
	if err := w.ReadPointer(&o.Forwarders, _s_aipForwarders, _ptr_aipForwarders); err != nil {
		return err
	}
	_ptr_aipLogFilter := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.LogFilter == nil {
			o.LogFilter = &AddrArray{}
		}
		if err := o.LogFilter.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipLogFilter := func(ptr interface{}) { o.LogFilter = *ptr.(**AddrArray) }
	if err := w.ReadPointer(&o.LogFilter, _s_aipLogFilter, _ptr_aipLogFilter); err != nil {
		return err
	}
	_ptr_pwszLogFilePath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LogFilePath); err != nil {
			return err
		}
		return nil
	})
	_s_pwszLogFilePath := func(ptr interface{}) { o.LogFilePath = *ptr.(*string) }
	if err := w.ReadPointer(&o.LogFilePath, _s_pwszLogFilePath, _ptr_pwszLogFilePath); err != nil {
		return err
	}
	_ptr_pszDomainName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.DomainName); err != nil {
			return err
		}
		return nil
	})
	_s_pszDomainName := func(ptr interface{}) { o.DomainName = *ptr.(*string) }
	if err := w.ReadPointer(&o.DomainName, _s_pszDomainName, _ptr_pszDomainName); err != nil {
		return err
	}
	_ptr_pszForestName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.ForestName); err != nil {
			return err
		}
		return nil
	})
	_s_pszForestName := func(ptr interface{}) { o.ForestName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ForestName, _s_pszForestName, _ptr_pszForestName); err != nil {
		return err
	}
	_ptr_pszDomainDirectoryPartition := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.DomainDirectoryPartition); err != nil {
			return err
		}
		return nil
	})
	_s_pszDomainDirectoryPartition := func(ptr interface{}) { o.DomainDirectoryPartition = *ptr.(*string) }
	if err := w.ReadPointer(&o.DomainDirectoryPartition, _s_pszDomainDirectoryPartition, _ptr_pszDomainDirectoryPartition); err != nil {
		return err
	}
	_ptr_pszForestDirectoryPartition := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.ForestDirectoryPartition); err != nil {
			return err
		}
		return nil
	})
	_s_pszForestDirectoryPartition := func(ptr interface{}) { o.ForestDirectoryPartition = *ptr.(*string) }
	if err := w.ReadPointer(&o.ForestDirectoryPartition, _s_pszForestDirectoryPartition, _ptr_pszForestDirectoryPartition); err != nil {
		return err
	}
	_ptr_pExtensions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.Extensions); err != nil {
			return err
		}
		return nil
	})
	_s_pExtensions := func(ptr interface{}) { o.Extensions = *ptr.(*string) }
	if err := w.ReadPointer(&o.Extensions, _s_pExtensions, _ptr_pExtensions); err != nil {
		return err
	}
	if err := w.ReadData(&o.LogLevel); err != nil {
		return err
	}
	if err := w.ReadData(&o.DebugLevel); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwardTimeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCProtocol); err != nil {
		return err
	}
	if err := w.ReadData(&o.NameCheckFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddressAnswerLimitCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.RecursionRetry); err != nil {
		return err
	}
	if err := w.ReadData(&o.RecursionTimeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxCacheTTL); err != nil {
		return err
	}
	if err := w.ReadData(&o.DSPollingInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocalNetPriorityNetMask); err != nil {
		return err
	}
	if err := w.ReadData(&o.ScavengingInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.DefaultRefreshInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.DefaultNoRefreshInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastScavengeTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.EventLogLevel); err != nil {
		return err
	}
	if err := w.ReadData(&o.LogFileMaxSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.DSForestVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.DSDomainVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.DSDSAVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReadOnlyDC); err != nil {
		return err
	}
	// reserved dwReserveArray
	var _dwReserveArray []uint32
	_dwReserveArray = make([]uint32, 3)
	for i1 := range _dwReserveArray {
		i1 := i1
		if err := w.ReadData(&_dwReserveArray[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.AutoReverseZones); err != nil {
		return err
	}
	if err := w.ReadData(&o.AutoCacheUpdate); err != nil {
		return err
	}
	if err := w.ReadData(&o.RecurseAfterForwarding); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwardDelegations); err != nil {
		return err
	}
	if err := w.ReadData(&o.NoRecursion); err != nil {
		return err
	}
	if err := w.ReadData(&o.SecureResponses); err != nil {
		return err
	}
	if err := w.ReadData(&o.RoundRobin); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocalNetPriority); err != nil {
		return err
	}
	if err := w.ReadData(&o.BindSecondaries); err != nil {
		return err
	}
	if err := w.ReadData(&o.WriteAuthorityNS); err != nil {
		return err
	}
	if err := w.ReadData(&o.StrictFileParsing); err != nil {
		return err
	}
	if err := w.ReadData(&o.LooseWildcarding); err != nil {
		return err
	}
	if err := w.ReadData(&o.DefaultAgingState); err != nil {
		return err
	}
	// reserved fReserveArray
	var _fReserveArray []bool
	_fReserveArray = make([]bool, 15)
	for i1 := range _fReserveArray {
		i1 := i1
		if err := w.ReadData(&_fReserveArray[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// ForwardersW2K structure represents DNS_RPC_FORWARDERS_W2K RPC structure.
//
// The DNS_RPC_FORWARDERS_W2K structure specifies the set of DNS servers this DNS server
// will forward unresolved queries to.
type ForwardersW2K struct {
	// fRecurseAfterForwarding: A value of 0x00000001 indicates that the DNS server is configured
	// to use normal recursion for name resolution if forwarders are not configured or are
	// unreachable; a value of 0x00000000 indicates it is not.
	RecurseAfterForwarding uint32 `idl:"name:fRecurseAfterForwarding" json:"recurse_after_forwarding"`
	// dwForwardTimeout: The time-interval, in seconds, for which the DNS server waits for
	// a response from each server in the forwarders list. No restrictions are applied to
	// the range for the dwForwardTimeout member when modifying its value through this structure.
	// If dwForwardTimeout is set to zero, then the server SHOULD<45> reset the forward
	// timeout to the default value, 3 minutes (180 seconds).
	ForwardTimeout uint32 `idl:"name:dwForwardTimeout" json:"forward_timeout"`
	// aipForwarders: The list of IP addresses that will be used as forwarders by the DNS
	// server.
	Forwarders *IPv4Array `idl:"name:aipForwarders" json:"forwarders"`
}

func (o *ForwardersW2K) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ForwardersW2K) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RecurseAfterForwarding); err != nil {
		return err
	}
	if err := w.WriteData(o.ForwardTimeout); err != nil {
		return err
	}
	if o.Forwarders != nil {
		_ptr_aipForwarders := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Forwarders != nil {
				if err := o.Forwarders.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPv4Array{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Forwarders, _ptr_aipForwarders); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ForwardersW2K) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RecurseAfterForwarding); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwardTimeout); err != nil {
		return err
	}
	_ptr_aipForwarders := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Forwarders == nil {
			o.Forwarders = &IPv4Array{}
		}
		if err := o.Forwarders.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipForwarders := func(ptr interface{}) { o.Forwarders = *ptr.(**IPv4Array) }
	if err := w.ReadPointer(&o.Forwarders, _s_aipForwarders, _ptr_aipForwarders); err != nil {
		return err
	}
	return nil
}

// Forwarders_NET structure represents DNS_RPC_FORWARDERS_DOTNET RPC structure.
//
// In the DNS_RPC_FORWARDERS_DOTNET structure all fields have same definition as specified
// in section DNS_RPC_FORWARDERS_W2K (section 2.2.5.2.10.1), with the following exceptions:
type Forwarders_NET struct {
	// dwRpcStructureVersion: The structure version number. It MUST be set to 0x00000001.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// dwReserved0: MUST be set to zero when sent and MUST be ignored on receipt.
	_                      uint32     `idl:"name:dwReserved0"`
	RecurseAfterForwarding uint32     `idl:"name:fRecurseAfterForwarding" json:"recurse_after_forwarding"`
	ForwardTimeout         uint32     `idl:"name:dwForwardTimeout" json:"forward_timeout"`
	Forwarders             *IPv4Array `idl:"name:aipForwarders" json:"forwarders"`
}

func (o *Forwarders_NET) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Forwarders_NET) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.RecurseAfterForwarding); err != nil {
		return err
	}
	if err := w.WriteData(o.ForwardTimeout); err != nil {
		return err
	}
	if o.Forwarders != nil {
		_ptr_aipForwarders := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Forwarders != nil {
				if err := o.Forwarders.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPv4Array{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Forwarders, _ptr_aipForwarders); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Forwarders_NET) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	if err := w.ReadData(&o.RecurseAfterForwarding); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwardTimeout); err != nil {
		return err
	}
	_ptr_aipForwarders := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Forwarders == nil {
			o.Forwarders = &IPv4Array{}
		}
		if err := o.Forwarders.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipForwarders := func(ptr interface{}) { o.Forwarders = *ptr.(**IPv4Array) }
	if err := w.ReadPointer(&o.Forwarders, _s_aipForwarders, _ptr_aipForwarders); err != nil {
		return err
	}
	return nil
}

// ForwardersLonghorn structure represents DNS_RPC_FORWARDERS_LONGHORN RPC structure.
//
// In the DNS_RPC_FORWARDERS_LONGHORN structure all fields have same definition as specified
// in section DNS_RPC_FORWARDERS_DOTNET (section 2.2.5.2.10.2), with the following exceptions:
type ForwardersLonghorn struct {
	// dwRpcStructureVersion: The structure version number. It MUST be set to 0x00000002.
	RPCStructureVersion    uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	_                      uint32 `idl:"name:dwReserved0"`
	RecurseAfterForwarding uint32 `idl:"name:fRecurseAfterForwarding" json:"recurse_after_forwarding"`
	ForwardTimeout         uint32 `idl:"name:dwForwardTimeout" json:"forward_timeout"`
	// aipForwarders: A pointer to a structure of type DNS_ADDR_ARRAY (section 2.2.3.2.3)
	// specifies a list of IP addresses that will be used as forwarders by the DNS server.
	Forwarders *AddrArray `idl:"name:aipForwarders" json:"forwarders"`
}

func (o *ForwardersLonghorn) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ForwardersLonghorn) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.RecurseAfterForwarding); err != nil {
		return err
	}
	if err := w.WriteData(o.ForwardTimeout); err != nil {
		return err
	}
	if o.Forwarders != nil {
		_ptr_aipForwarders := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Forwarders != nil {
				if err := o.Forwarders.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AddrArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Forwarders, _ptr_aipForwarders); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ForwardersLonghorn) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	if err := w.ReadData(&o.RecurseAfterForwarding); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwardTimeout); err != nil {
		return err
	}
	_ptr_aipForwarders := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Forwarders == nil {
			o.Forwarders = &AddrArray{}
		}
		if err := o.Forwarders.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipForwarders := func(ptr interface{}) { o.Forwarders = *ptr.(**AddrArray) }
	if err := w.ReadPointer(&o.Forwarders, _s_aipForwarders, _ptr_aipForwarders); err != nil {
		return err
	}
	return nil
}

// Forwarders structure represents DNS_RPC_FORWARDERS RPC structure.
//
// The DNS_RPC_FORWARDERS structure contains information about forwarders configured
// on the DNS server. There are different versions of the DNS_RPC_FORWARDERS structure.
// The DNS server MUST use the structures corresponding to the value of dwClientVersion
// in DNS Server Management Protocol method calls (section 3.1.4) in the following table,
// or if the method call does not specify the value of dwClientVersion, the DNS_RPC_FORWARDERS_W2K
// version of the structure MUST be used.
//
//	+------------+----------------------------------------------------+
//	|            |                                                    |
//	|   VALUE    |                     STRUCTURE                      |
//	|            |                                                    |
//	+------------+----------------------------------------------------+
//	+------------+----------------------------------------------------+
//	| 0x00000000 | DNS_RPC_FORWARDERS_W2K (section 2.2.5.2.10.1)      |
//	+------------+----------------------------------------------------+
//	| 0x00060000 | DNS_RPC_FORWARDERS_DOTNET (section 2.2.5.2.10.2)   |
//	+------------+----------------------------------------------------+
//	| 0x00070000 | DNS_RPC_FORWARDERS_LONGHORN (section 2.2.5.2.10.3) |
//	+------------+----------------------------------------------------+
type Forwarders struct {
	RPCStructureVersion    uint32     `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	_                      uint32     `idl:"name:dwReserved0"`
	RecurseAfterForwarding uint32     `idl:"name:fRecurseAfterForwarding" json:"recurse_after_forwarding"`
	ForwardTimeout         uint32     `idl:"name:dwForwardTimeout" json:"forward_timeout"`
	Forwarders             *AddrArray `idl:"name:aipForwarders" json:"forwarders"`
}

func (o *Forwarders) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Forwarders) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.RecurseAfterForwarding); err != nil {
		return err
	}
	if err := w.WriteData(o.ForwardTimeout); err != nil {
		return err
	}
	if o.Forwarders != nil {
		_ptr_aipForwarders := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Forwarders != nil {
				if err := o.Forwarders.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AddrArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Forwarders, _ptr_aipForwarders); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Forwarders) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	if err := w.ReadData(&o.RecurseAfterForwarding); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwardTimeout); err != nil {
		return err
	}
	_ptr_aipForwarders := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Forwarders == nil {
			o.Forwarders = &AddrArray{}
		}
		if err := o.Forwarders.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipForwarders := func(ptr interface{}) { o.Forwarders = *ptr.(**AddrArray) }
	if err := w.ReadPointer(&o.Forwarders, _s_aipForwarders, _ptr_aipForwarders); err != nil {
		return err
	}
	return nil
}

// ZoneW2K structure represents DNS_RPC_ZONE_W2K RPC structure.
//
// The DNS_RPC_ZONE_W2K structure is used to specify basic information about a DNS zone.
type ZoneW2K struct {
	// pszZoneName: A pointer to a null-terminated Unicode string that contains zone-name.
	ZoneName string `idl:"name:pszZoneName;string" json:"zone_name"`
	// Flags: Zone flags as specified in section 2.2.5.2.2.
	Flags uint32 `idl:"name:Flags" json:"flags"`
	// ZoneType: The zone type. This MUST be set to one of the allowed DNS_ZONE_TYPE section
	// 2.2.5.1.1.
	ZoneType uint8 `idl:"name:ZoneType" json:"zone_type"`
	// Version: The RPC protocol version. It MUST be set to 0x32.
	Version uint8 `idl:"name:Version" json:"version"`
}

func (o *ZoneW2K) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ZoneW2K) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.ZoneName != "" {
		_ptr_pszZoneName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ZoneName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ZoneName, _ptr_pszZoneName); err != nil {
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
	if err := w.WriteData(o.ZoneType); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *ZoneW2K) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_pszZoneName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ZoneName); err != nil {
			return err
		}
		return nil
	})
	_s_pszZoneName := func(ptr interface{}) { o.ZoneName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ZoneName, _s_pszZoneName, _ptr_pszZoneName); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.ZoneType); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// Zone_NET structure represents DNS_RPC_ZONE_DOTNET RPC structure.
//
// In the DNS_RPC_ZONE_DOTNET structure all fields have same definition as specified
// in section DNS_RPC_ZONE_W2K (section 2.2.5.2.1.1), with the following exceptions:
type Zone_NET struct {
	// dwRpcStructureVersion: The DNS management structure version number. This value MUST
	// be set to 0x00000001.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// dwReserved0: This field is reserved for future use. Senders set this to an arbitrary
	// value and receivers MUST ignore it.
	_        uint32 `idl:"name:dwReserved0"`
	ZoneName string `idl:"name:pszZoneName;string" json:"zone_name"`
	Flags    uint32 `idl:"name:Flags" json:"flags"`
	ZoneType uint8  `idl:"name:ZoneType" json:"zone_type"`
	Version  uint8  `idl:"name:Version" json:"version"`
	// dwDpFlags: Application directory partition flags for this zone. This MUST be set
	// to one of the combination of the DNS_RPC_DP_FLAGS (section 2.2.7.1.1). If this zone
	// is not stored in the directory server this value MUST be 0x00000000.
	DPFlags uint32 `idl:"name:dwDpFlags" json:"dp_flags"`
	// pszDpFqdn: A pointer to a null-terminated UTF-8 string that specifies the FQDN of
	// the application directory partition in which this zone is stored in the directory
	// server. If this zone is not stored in the directory server this value MUST be NULL.
	//
	// If the DNS RPC client sends an older version of DNS_RPC_ZONE structure, that is,
	// DNS_RPC_ZONE_W2K (section 2.2.5.2.1.1), then the DNS RPC server MUST construct a
	// current version of DNS_RPC_ZONE structure, that is, DNS_RPC_ZONE_DOTNET, using the
	// following steps:
	//
	// *
	//
	// Copy the same value for fields that are common to input and the current version of
	// the DNS_RPC_ZONE structures.
	//
	// *
	//
	// The dwRpcStructureVersion field MUST be set to "1".
	//
	// *
	//
	// All other fields that are defined only in DNS_RPC_ZONE_DOTNET and are not defined
	// in DNS_RPC_ZONE (section 2.2.5.2.1 ( f22773ca-bfbf-46a5-8042-fbda1f3c4ad2 ) ), MUST
	// be set to "0".
	DPFQDN string `idl:"name:pszDpFqdn;string" json:"dp_fqdn"`
}

func (o *Zone_NET) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Zone_NET) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if o.ZoneName != "" {
		_ptr_pszZoneName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ZoneName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ZoneName, _ptr_pszZoneName); err != nil {
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
	if err := w.WriteData(o.ZoneType); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	if err := w.WriteData(o.DPFlags); err != nil {
		return err
	}
	if o.DPFQDN != "" {
		_ptr_pszDpFqdn := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.DPFQDN); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DPFQDN, _ptr_pszDpFqdn); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Zone_NET) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	_ptr_pszZoneName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ZoneName); err != nil {
			return err
		}
		return nil
	})
	_s_pszZoneName := func(ptr interface{}) { o.ZoneName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ZoneName, _s_pszZoneName, _ptr_pszZoneName); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.ZoneType); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	if err := w.ReadData(&o.DPFlags); err != nil {
		return err
	}
	_ptr_pszDpFqdn := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.DPFQDN); err != nil {
			return err
		}
		return nil
	})
	_s_pszDpFqdn := func(ptr interface{}) { o.DPFQDN = *ptr.(*string) }
	if err := w.ReadPointer(&o.DPFQDN, _s_pszDpFqdn, _ptr_pszDpFqdn); err != nil {
		return err
	}
	return nil
}

// Zone structure represents DNS_RPC_ZONE RPC structure.
//
// The DNS_RPC_ZONE structure contains basic information about a zone present on the
// DNS server. There are different versions of the DNS_RPC_ZONE structure. The DNS server
// MAY<38> decide to use one of these structures depending upon the value of the dwClientVersion
// field in DNS Server Management Protocol method calls (section 3.1.4) as follows in
// the table provided. If the method call does not specify the value of dwClientVersion,
// the DNS_RPC_ZONE_W2K version of the structure MUST be used.
//
//	+------------+-------------------------------------------+
//	|            |                                           |
//	|   VALUE    |                 STRUCTURE                 |
//	|            |                                           |
//	+------------+-------------------------------------------+
//	+------------+-------------------------------------------+
//	| 0x00000000 | DNS_RPC_ZONE_W2K (section 2.2.5.2.1.1)    |
//	+------------+-------------------------------------------+
//	| 0x00060000 | DNS_RPC_ZONE_DOTNET (section 2.2.5.2.1.2) |
//	+------------+-------------------------------------------+
type Zone struct {
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	_                   uint32 `idl:"name:dwReserved0"`
	ZoneName            string `idl:"name:pszZoneName;string" json:"zone_name"`
	Flags               uint32 `idl:"name:Flags" json:"flags"`
	ZoneType            uint8  `idl:"name:ZoneType" json:"zone_type"`
	Version             uint8  `idl:"name:Version" json:"version"`
	DPFlags             uint32 `idl:"name:dwDpFlags" json:"dp_flags"`
	DPFQDN              string `idl:"name:pszDpFqdn;string" json:"dp_fqdn"`
}

func (o *Zone) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Zone) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if o.ZoneName != "" {
		_ptr_pszZoneName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ZoneName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ZoneName, _ptr_pszZoneName); err != nil {
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
	if err := w.WriteData(o.ZoneType); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	if err := w.WriteData(o.DPFlags); err != nil {
		return err
	}
	if o.DPFQDN != "" {
		_ptr_pszDpFqdn := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.DPFQDN); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DPFQDN, _ptr_pszDpFqdn); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Zone) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	_ptr_pszZoneName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ZoneName); err != nil {
			return err
		}
		return nil
	})
	_s_pszZoneName := func(ptr interface{}) { o.ZoneName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ZoneName, _s_pszZoneName, _ptr_pszZoneName); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.ZoneType); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	if err := w.ReadData(&o.DPFlags); err != nil {
		return err
	}
	_ptr_pszDpFqdn := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.DPFQDN); err != nil {
			return err
		}
		return nil
	})
	_s_pszDpFqdn := func(ptr interface{}) { o.DPFQDN = *ptr.(*string) }
	if err := w.ReadPointer(&o.DPFQDN, _s_pszDpFqdn, _ptr_pszDpFqdn); err != nil {
		return err
	}
	return nil
}

// ZoneListW2K structure represents DNS_RPC_ZONE_LIST_W2K RPC structure.
//
// The DNS_RPC_ZONE_LIST_W2K structure is used to enumerate zones.
type ZoneListW2K struct {
	// dwZoneCount: The number of zones present in the array of zones pointed to by ZoneArray.
	ZoneCount uint32 `idl:"name:dwZoneCount" json:"zone_count"`
	// ZoneArray: An array of structures of type DNS_RPC_ZONE_W2K (section 2.2.5.2.1.1).
	// Each element of the array represents one zone.
	ZoneArray []*ZoneW2K `idl:"name:ZoneArray;size_is:(dwZoneCount)" json:"zone_array"`
}

func (o *ZoneListW2K) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.ZoneArray != nil && o.ZoneCount == 0 {
		o.ZoneCount = uint32(len(o.ZoneArray))
	}
	if o.ZoneCount > uint32(500000) {
		return fmt.Errorf("ZoneCount is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *ZoneListW2K) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.ZoneCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ZoneListW2K) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.ZoneCount); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	for i1 := range o.ZoneArray {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.ZoneArray[i1] != nil {
			_ptr_ZoneArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ZoneArray[i1] != nil {
					if err := o.ZoneArray[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ZoneW2K{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ZoneArray[i1], _ptr_ZoneArray); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.ZoneArray); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ZoneListW2K) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.ZoneCount); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.ZoneCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.ZoneCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.ZoneArray", sizeInfo[0])
	}
	o.ZoneArray = make([]*ZoneW2K, sizeInfo[0])
	for i1 := range o.ZoneArray {
		i1 := i1
		_ptr_ZoneArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ZoneArray[i1] == nil {
				o.ZoneArray[i1] = &ZoneW2K{}
			}
			if err := o.ZoneArray[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ZoneArray := func(ptr interface{}) { o.ZoneArray[i1] = *ptr.(**ZoneW2K) }
		if err := w.ReadPointer(&o.ZoneArray[i1], _s_ZoneArray, _ptr_ZoneArray); err != nil {
			return err
		}
	}
	return nil
}

// ZoneList_NET structure represents DNS_RPC_ZONE_LIST_DOTNET RPC structure.
//
// In the DNS_RPC_ZONE_LIST_DOTNET structure all fields have same definition as specified
// in section DNS_RPC_ZONE_LIST_W2K (section 2.2.5.2.3.1), with the following exceptions:
type ZoneList_NET struct {
	// dwRpcStructureVersion: The DNS management structure version number. This MUST be
	// set to 0x00000001.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// dwReserved0: This field is reserved for future use. Senders can send an arbitrary
	// value and receivers MUST ignore it.
	//
	// If the DNS RPC client sends an older version of DNS_RPC_ZONE_LIST structure, that
	// is, DNS_RPC_ZONE_LIST_W2K (section 2.2.5.2.3.1), then the DNS RPC server MUST construct
	// a current version of DNS_RPC_ZONE_LIST structure, that is, DNS_RPC_ZONE_LIST_DOTNET,
	// using the following steps:
	//
	// *
	//
	// Copy the same value for the fields that are common to input and the current version
	// of DNS_RPC_ZONE_LIST structures.
	//
	// *
	//
	// The dwRpcStructureVersion field MUST be set to "1".
	//
	// *
	//
	// The value for the ZoneArray field MUST be obtained from the input structure as DNS_RPC_ZONE_W2K
	// (section 2.2.5.2.1.1 ( 6b6e834a-e08e-432c-a9a7-8ced1422a688 ) ) array elements and
	// each MUST be converted using the steps specified in section DNS_RPC_ZONE_DOTNET (section
	// 2.2.5.2.1.2 ( 64c5005d-2a55-46cc-8453-af393bb27bb9 ) ), and then assigned to the
	// ZoneArray field in the DNS_RPC_ZONE_LIST_DOTNET structure.
	//
	// *
	//
	// All other fields that are defined only in DNS_RPC_ZONE_LIST_DOTNET and are not defined
	// in DNS_RPC_ZONE_LIST_W2K (section 2.2.5.2.3.1), MUST be set to "0".
	_         uint32      `idl:"name:dwReserved0"`
	ZoneCount uint32      `idl:"name:dwZoneCount" json:"zone_count"`
	ZoneArray []*Zone_NET `idl:"name:ZoneArray;size_is:(dwZoneCount)" json:"zone_array"`
}

func (o *ZoneList_NET) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.ZoneArray != nil && o.ZoneCount == 0 {
		o.ZoneCount = uint32(len(o.ZoneArray))
	}
	if o.ZoneCount > uint32(500000) {
		return fmt.Errorf("ZoneCount is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *ZoneList_NET) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.ZoneCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ZoneList_NET) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.ZoneCount); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	for i1 := range o.ZoneArray {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.ZoneArray[i1] != nil {
			_ptr_ZoneArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ZoneArray[i1] != nil {
					if err := o.ZoneArray[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&Zone_NET{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ZoneArray[i1], _ptr_ZoneArray); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.ZoneArray); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ZoneList_NET) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	if err := w.ReadData(&o.ZoneCount); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.ZoneCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.ZoneCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.ZoneArray", sizeInfo[0])
	}
	o.ZoneArray = make([]*Zone_NET, sizeInfo[0])
	for i1 := range o.ZoneArray {
		i1 := i1
		_ptr_ZoneArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ZoneArray[i1] == nil {
				o.ZoneArray[i1] = &Zone_NET{}
			}
			if err := o.ZoneArray[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ZoneArray := func(ptr interface{}) { o.ZoneArray[i1] = *ptr.(**Zone_NET) }
		if err := w.ReadPointer(&o.ZoneArray[i1], _s_ZoneArray, _ptr_ZoneArray); err != nil {
			return err
		}
	}
	return nil
}

// ZoneList structure represents DNS_RPC_ZONE_LIST RPC structure.
//
// The DNS_RPC_ZONE_LIST structure contains the information about a list of zones present
// on the DNS server. There are different versions of the DNS_RPC_ZONE_LIST structure.
// The DNS server SHOULD<41> use one of these structures depending upon the passed-in
// value for the dwClientVersion field in DNS Server Management Protocol method calls
// (section 3.1.4) as shown in the following table. If the method call does not specify
// the value of dwClientVersion, the DNS_RPC_ZONE_LIST_W2K version of the structure
// MUST be used.
//
//	+------------+------------------------------------------------+
//	|            |                                                |
//	|   VALUE    |                   STRUCTURE                    |
//	|            |                                                |
//	+------------+------------------------------------------------+
//	+------------+------------------------------------------------+
//	| 0x00000000 | DNS_RPC_ZONE_LIST_W2K (section 2.2.5.2.3.1)    |
//	+------------+------------------------------------------------+
//	| 0x00060000 | DNS_RPC_ZONE_LIST_DOTNET (section 2.2.5.2.3.2) |
//	+------------+------------------------------------------------+
type ZoneList struct {
	RPCStructureVersion uint32      `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	_                   uint32      `idl:"name:dwReserved0"`
	ZoneCount           uint32      `idl:"name:dwZoneCount" json:"zone_count"`
	ZoneArray           []*Zone_NET `idl:"name:ZoneArray;size_is:(dwZoneCount)" json:"zone_array"`
}

func (o *ZoneList) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.ZoneArray != nil && o.ZoneCount == 0 {
		o.ZoneCount = uint32(len(o.ZoneArray))
	}
	if o.ZoneCount > uint32(500000) {
		return fmt.Errorf("ZoneCount is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *ZoneList) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.ZoneCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ZoneList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.ZoneCount); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	for i1 := range o.ZoneArray {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.ZoneArray[i1] != nil {
			_ptr_ZoneArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ZoneArray[i1] != nil {
					if err := o.ZoneArray[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&Zone_NET{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ZoneArray[i1], _ptr_ZoneArray); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.ZoneArray); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ZoneList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	if err := w.ReadData(&o.ZoneCount); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.ZoneCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.ZoneCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.ZoneArray", sizeInfo[0])
	}
	o.ZoneArray = make([]*Zone_NET, sizeInfo[0])
	for i1 := range o.ZoneArray {
		i1 := i1
		_ptr_ZoneArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ZoneArray[i1] == nil {
				o.ZoneArray[i1] = &Zone_NET{}
			}
			if err := o.ZoneArray[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ZoneArray := func(ptr interface{}) { o.ZoneArray[i1] = *ptr.(**Zone_NET) }
		if err := w.ReadPointer(&o.ZoneArray[i1], _s_ZoneArray, _ptr_ZoneArray); err != nil {
			return err
		}
	}
	return nil
}

// TrustPointState type represents TRUSTPOINT_STATE RPC enumeration.
//
// The TRUSTPOINT_STATE enumeration identifies the current state of the life cycle of
// a DNSSEC trust point. The state of the trust point is determined by the states of
// its trust anchors and is given for informational purposes only; it does not affect
// the DNS or LDAP protocol.
type TrustPointState uint16

var (
	// TRUSTPOINT_STATE_INITIALIZED: Indicates a trust point that has no trust anchors but
	// has child nodes that do have trust anchors.
	TrustPointStateInitialized TrustPointState = 0
	// TRUSTPOINT_STATE_DSPENDING: Indicates a trust point that has only DS trust anchors
	// and is therefore ineffective for DNSSEC operations.
	TrustPointStateDspending TrustPointState = 1
	// TRUSTPOINT_STATE_ACTIVE: Indicates a trust point that has one or more DNSKEY trust
	// anchors. Queries covered by this trust point will be validated using the DNSKEY trust
	// anchors.
	TrustPointStateActive TrustPointState = 2
	// TRUSTPOINT_STATE_DELETE_PENDING: Indicates a trust point containing only revoked
	// DNSKEY trust anchors. When the [RFC5011] remove-hold-down time is exceeded for all
	// revoked trust anchors, this trust point will be deleted.
	TrustPointStateDeletePending TrustPointState = 3
	// TRUSTPOINT_STATE_DELETED: Reserved. MUST NOT be sent by the server, MUST be ignored
	// by the receiver.
	TrustPointStateDeleted TrustPointState = 4
)

func (o TrustPointState) String() string {
	switch o {
	case TrustPointStateInitialized:
		return "TrustPointStateInitialized"
	case TrustPointStateDspending:
		return "TrustPointStateDspending"
	case TrustPointStateActive:
		return "TrustPointStateActive"
	case TrustPointStateDeletePending:
		return "TrustPointStateDeletePending"
	case TrustPointStateDeleted:
		return "TrustPointStateDeleted"
	}
	return "Invalid"
}

// TrustAnchorState type represents TRUSTANCHOR_STATE RPC enumeration.
//
// The TRUSTANCHOR_STATE enumeration identifies the current state of the life cycle
// of a DNSSEC trust anchor. All states other than TRUSTANCHOR_STATE_DSPENDING and TRUSTANCHOR_STATE_DSINVALID
// correspond to the states given in [RFC5011] section 4.2. Only the TRUSTANCHOR_STATE_VALID
// and TRUSTANCHOR_STATE_MISSING trust anchors affect the DNS protocol (by acting as
// a DNSSEC trust anchor). None of the states affect the LDAP protocol [RFC4511].
type TrustAnchorState uint16

var (
	// TRUSTANCHOR_STATE_INITIALIZED: Reserved. MUST NOT be sent by the server, MUST be
	// ignored by the receiver.
	TrustAnchorStateInitialized TrustAnchorState = 0
	// TRUSTANCHOR_STATE_DSPENDING: Indicates a DS trust anchor that can be replaced by
	// the DNSKEY trust anchor that matches the hash of the DS trust anchor after the next
	// [RFC5011] active refresh.
	TrustAnchorStateDspending TrustAnchorState = 1
	// TRUSTANCHOR_STATE_DSINVALID: Indicates a DS trust anchor that matches a DNSKEY record
	// that is not valid as a trust anchor because the SEP flag is not set or because the
	// algorithm is unsupported.
	TrustAnchorStateDsinvalid TrustAnchorState = 2
	// TRUSTANCHOR_STATE_ADDPEND: Indicates a DNSKEY trust anchor that has not exceeded
	// the [RFC5011] add-hold-down time. This trust anchor can become a valid trust anchor
	// after the hold-down time has been exceeded.
	TrustAnchorStateAddpend TrustAnchorState = 3
	// TRUSTANCHOR_STATE_VALID: Indicates a DNSKEY trust anchor that has exceeded the [RFC5011]
	// add-hold-down time, has been matched to (and has replaced) a DS trust anchor, or
	// has been added by the administrator. This trust anchor is trusted for DNSSEC operations,
	// and a copy exists in the TrustAnchors zone.
	TrustAnchorStateValid TrustAnchorState = 4
	// TRUSTANCHOR_STATE_MISSING: Indicates a DNSKEY trust anchor that is valid but was
	// not seen in the last [RFC5011] active refresh. This trust anchor is trusted for DNSSEC
	// operations, and a copy exists in the TrustAnchors zone.
	TrustAnchorStateMissing TrustAnchorState = 5
	// TRUSTANCHOR_STATE_REVOKED: Indicates a DNSKEY trust anchor that has been revoked
	// by the authoritative zone administrator according to [RFC5011]. This trust anchor
	// will be deleted when the [RFC5011] remove-hold-down time is exceeded.
	TrustAnchorStateRevoked TrustAnchorState = 6
	// TRUSTANCHOR_STATE_DELETED: Reserved. MUST NOT be sent by the server, MUST be ignored
	// by the receiver.
	TrustAnchorStateDeleted TrustAnchorState = 7
)

func (o TrustAnchorState) String() string {
	switch o {
	case TrustAnchorStateInitialized:
		return "TrustAnchorStateInitialized"
	case TrustAnchorStateDspending:
		return "TrustAnchorStateDspending"
	case TrustAnchorStateDsinvalid:
		return "TrustAnchorStateDsinvalid"
	case TrustAnchorStateAddpend:
		return "TrustAnchorStateAddpend"
	case TrustAnchorStateValid:
		return "TrustAnchorStateValid"
	case TrustAnchorStateMissing:
		return "TrustAnchorStateMissing"
	case TrustAnchorStateRevoked:
		return "TrustAnchorStateRevoked"
	case TrustAnchorStateDeleted:
		return "TrustAnchorStateDeleted"
	}
	return "Invalid"
}

// TrustPoint structure represents DNS_RPC_TRUST_POINT RPC structure.
//
// The DNS_RPC_TRUST_POINT structure contains information about a trust point or a node
// in the TrustAnchors zone ([RFC5011]).
type TrustPoint struct {
	// dwRpcStructureVersion: The structure version number; it MUST be set to 0x00000001.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// dwReserved0: MUST be set to zero when sent and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved0"`
	// pszTrustPointName: The FQDN of the trust point or node in the TrustAnchors zone represented
	// by this structure. This MUST be a non-empty, non-NULL string.
	TrustPointName string `idl:"name:pszTrustPointName;string" json:"trust_point_name"`
	// eTrustPointState: A TRUSTPOINT_STATE enumeration value (section 2.2.1.1.3) containing
	// the current trust point state. This MUST be set to one of the following values. For
	// the TRUSTANCHOR_STATE enumeration values see section 2.2.1.1.4.
	//
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                            |                                                                                  |
	//	|                   VALUE                    |                                     MEANING                                      |
	//	|                                            |                                                                                  |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| TRUSTPOINT_STATE_INITIALIZED 0x00000000    | This structure represents a node in the TrustAnchors zone that does not contain  |
	//	|                                            | any trust anchors. This node is not a trust point.                               |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| TRUSTPOINT_STATE_DSPENDING 0x00000001      | This trust point contains only DS trust anchors (trust anchors in the            |
	//	|                                            | TRUSTANCHOR_STATE_DSPENDING or TRUSTANCHOR_STATE_DSINVALID state), rendering it  |
	//	|                                            | unusable for DNSSEC proofs.                                                      |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| TRUSTPOINT_STATE_ACTIVE 0x00000002         | This trust point contains at least one trust anchor in the                       |
	//	|                                            | TRUSTANCHOR_STATE_VALID or TRUSTANCHOR_STATE_MISSING state.                      |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| TRUSTPOINT_STATE_DELETE_PENDING 0x00000003 | This trust point contains only trust anchors in the TRUSTANCHOR_STATE_REVOKED    |
	//	|                                            | state.                                                                           |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	TrustPointState TrustPointState `idl:"name:eTrustPointState" json:"trust_point_state"`
	// i64LastActiveRefreshTime: The time of the last active refresh. This is set to zero
	// to indicate that no active refresh has occurred. This is a 64-bit value representing
	// the number of 100-nanosecond intervals since January 1, 1601 Coordinated Universal
	// Time (UTC).
	LastActiveRefreshTime int64 `idl:"name:i64LastActiveRefreshTime" json:"last_active_refresh_time"`
	// i64NextActiveRefreshTime: The scheduled time of the next active refresh, or zero
	// to indicate that no active refresh is scheduled. This is a 64-bit value representing
	// the number of 100-nanosecond intervals since January 1, 1601 (UTC).
	NextActiveRefreshTime int64 `idl:"name:i64NextActiveRefreshTime" json:"next_active_refresh_time"`
	// i64LastSuccessfulActiveRefreshTime: The time of the last successful active refresh,
	// or zero to indicate that no active refresh has occurred. This is a 64-bit value representing
	// the number of 100-nanosecond intervals since January 1, 1601 (UTC). A successful
	// active refresh is defined as an active refresh resulting in retrieval of one or more
	// DNSKEY records for the trust point and, if this trust point has trust anchors in
	// the TRUSTANCHOR_STATE_VALID state, signifies that one or more of the retrieved DNSKEY
	// records was validated by DNSSEC.
	LastSuccessfulActiveRefreshTime int64 `idl:"name:i64LastSuccessfulActiveRefreshTime" json:"last_successful_active_refresh_time"`
	// dwLastActiveRefreshResult: The result of the last active refresh, either ERROR_SUCCESS
	// or a nonzero value to indicate that an error has occurred.
	LastActiveRefreshResult uint32 `idl:"name:dwLastActiveRefreshResult" json:"last_active_refresh_result"`
	// dwReserved: MUST be set to zero when sent and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved"`
}

func (o *TrustPoint) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *TrustPoint) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if o.TrustPointName != "" {
		_ptr_pszTrustPointName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.TrustPointName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.TrustPointName, _ptr_pszTrustPointName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteEnum(uint16(o.TrustPointState)); err != nil {
		return err
	}
	if err := w.WriteData(o.LastActiveRefreshTime); err != nil {
		return err
	}
	if err := w.WriteData(o.NextActiveRefreshTime); err != nil {
		return err
	}
	if err := w.WriteData(o.LastSuccessfulActiveRefreshTime); err != nil {
		return err
	}
	if err := w.WriteData(o.LastActiveRefreshResult); err != nil {
		return err
	}
	// reserved dwReserved
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *TrustPoint) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	_ptr_pszTrustPointName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.TrustPointName); err != nil {
			return err
		}
		return nil
	})
	_s_pszTrustPointName := func(ptr interface{}) { o.TrustPointName = *ptr.(*string) }
	if err := w.ReadPointer(&o.TrustPointName, _s_pszTrustPointName, _ptr_pszTrustPointName); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.TrustPointState)); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastActiveRefreshTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.NextActiveRefreshTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastSuccessfulActiveRefreshTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastActiveRefreshResult); err != nil {
		return err
	}
	// reserved dwReserved
	var _dwReserved uint32
	if err := w.ReadData(&_dwReserved); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// TrustPointList structure represents DNS_RPC_TRUST_POINT_LIST RPC structure.
//
// The DNS_RPC_TRUST_POINT_LIST structure contains zero or more DNS_RPC_TRUST_POINT
// structures.
type TrustPointList struct {
	// dwRpcStructureVersion: The structure version number; this MUST be set to 0x00000001.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// dwReserved0: MUST be set to zero when sent and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved0"`
	// dwTrustPointCount: The size of the TrustPointArray array.
	TrustPointCount uint32 `idl:"name:dwTrustPointCount" json:"trust_point_count"`
	// TrustPointArray: An array of pointers to DNS_RPC_TRUST_POINT structures.
	TrustPointArray []*TrustPoint `idl:"name:TrustPointArray;size_is:(dwTrustPointCount)" json:"trust_point_array"`
}

func (o *TrustPointList) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.TrustPointArray != nil && o.TrustPointCount == 0 {
		o.TrustPointCount = uint32(len(o.TrustPointArray))
	}
	if o.TrustPointCount > uint32(500000) {
		return fmt.Errorf("TrustPointCount is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *TrustPointList) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.TrustPointCount)
	return []uint64{
		dimSize1,
	}
}
func (o *TrustPointList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.TrustPointCount); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	for i1 := range o.TrustPointArray {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.TrustPointArray[i1] != nil {
			_ptr_TrustPointArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.TrustPointArray[i1] != nil {
					if err := o.TrustPointArray[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&TrustPoint{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.TrustPointArray[i1], _ptr_TrustPointArray); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.TrustPointArray); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustPointList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	if err := w.ReadData(&o.TrustPointCount); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.TrustPointCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.TrustPointCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.TrustPointArray", sizeInfo[0])
	}
	o.TrustPointArray = make([]*TrustPoint, sizeInfo[0])
	for i1 := range o.TrustPointArray {
		i1 := i1
		_ptr_TrustPointArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.TrustPointArray[i1] == nil {
				o.TrustPointArray[i1] = &TrustPoint{}
			}
			if err := o.TrustPointArray[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_TrustPointArray := func(ptr interface{}) { o.TrustPointArray[i1] = *ptr.(**TrustPoint) }
		if err := w.ReadPointer(&o.TrustPointArray[i1], _s_TrustPointArray, _ptr_TrustPointArray); err != nil {
			return err
		}
	}
	return nil
}

// SKD structure represents DNS_RPC_SKD RPC structure.
//
// The DNS_RPC_SKD structure specifies a signing key descriptor.
type SKD struct {
	// dwRpcStructureVersion: The structure version number; this MUST be set to 0x00000001.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// dwReserved0: MUST be set to zero when sent and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved0"`
	// Guid: A unique identifier for this signing key descriptor.
	GUID *dtyp.GUID `idl:"name:Guid" json:"guid"`
	// pwszKeyStorageProvider: The Key Storage Provider that will be used to generate and
	// store keys for this signing key descriptor.
	KeyStorageProvider string `idl:"name:pwszKeyStorageProvider;string" json:"key_storage_provider"`
	// fStoreKeysInDirectory: A value of 0x00000001 indicates that the DNS server exports
	// keys generated for this signing key descriptor and stores them on the DNS zone object
	// in Active Directory. A value of 0x00000000 indicates that the DNS server stores keys
	// in self-signed certificates in the local machine certificate store.
	StoreKeysInDirectory bool `idl:"name:fStoreKeysInDirectory" json:"store_keys_in_directory"`
	// fIsKSK: A value of 0x00000001 indicates that this signing key descriptor describes
	// a key signing key (KSK); a value of 0x00000000 indicates a zone signing key [RFC4641].
	IsKSK bool `idl:"name:fIsKSK" json:"is_ksk"`
	// bSigningAlgorithm: The cryptographic algorithm used to generate signing keys. The
	// DNS server SHOULD<48> support all values given by [IANA-DNSSECAN] and [DRAFT-DNSSEC-ECDSA].
	SigningAlgorithm uint8 `idl:"name:bSigningAlgorithm" json:"signing_algorithm"`
	// dwKeyLength: The length, in bits, of cryptographic signing keys. This value MUST
	// be ignored if the signing algorithm does not have variable key length.
	KeyLength uint32 `idl:"name:dwKeyLength" json:"key_length"`
	// dwInitialRolloverOffset: The amount of time, in seconds, to delay the first scheduled
	// key rollover for this signing key descriptor. The value SHOULD be limited to the
	// range 0x00000000 to 0x00278D00 (30 days), inclusive, but it can be any value. The
	// default value SHOULD be 0x00000000, and the value zero MUST be allowed and treated
	// literally.
	InitialRolloverOffset uint32 `idl:"name:dwInitialRolloverOffset" json:"initial_rollover_offset"`
	// dwDNSKEYSignatureValidityPeriod: The number of seconds that signatures covering DNSKEY
	// record sets generated for this signing key descriptor's keys are valid, as defined
	// by [RFC4034] section 3.1.5. The value SHOULD be limited to the range 0x00005460 (6
	// hours) to 0x00093A80 (7 days), inclusive, but it can be any value. The default value
	// SHOULD be 0x0003F480 (3 days).
	DNSKEYSignatureValidityPeriod uint32 `idl:"name:dwDNSKEYSignatureValidityPeriod" json:"dnskey_signature_validity_period"`
	// dwDSSignatureValidityPeriod: The number of seconds that signatures covering DS record
	// sets generated for this signing key descriptor's keys are valid, as defined by [RFC4034]
	// section 3.1.5. The value SHOULD be limited to the range 0x00005460 (6 hours) to 0x00093A80
	// (7 days), inclusive, but it can be any value. The default value SHOULD be 0x0003F480
	// (3 days).
	DSSignatureValidityPeriod uint32 `idl:"name:dwDSSignatureValidityPeriod" json:"ds_signature_validity_period"`
	// dwStandardSignatureValidityPeriod: The number of seconds that signatures covering
	// record sets not of type DNSKEY or DS generated for this signing key descriptor's
	// keys are valid, as defined by [RFC4034] section 3.1.5. The value SHOULD be limited
	// to the range 0x00005460 (6 hours) to 0x00278D00 (30 days), inclusive, but it can
	// be any value. The default value SHOULD be 0x000D2F00 (10 days).
	StandardSignatureValidityPeriod uint32 `idl:"name:dwStandardSignatureValidityPeriod" json:"standard_signature_validity_period"`
	// dwRolloverType: When sent, this value MUST be set to one of the allowed values specified
	// in ZONE_SKD_ROLLOVER_TYPE (section 2.2.5.1.5) as follows: this value MUST be DNS_ROLLOVER_TYPE_PREPUBLISH
	// if fIsKSK is 0x00000000 and MUST be DNS_ROLLOVER_TYPE_DOUBLE_SIGNATURE if fIsKSK
	// is 0x00000001. On receipt, this value MUST be ignored.
	RolloverType uint32 `idl:"name:dwRolloverType" json:"rollover_type"`
	// dwRolloverPeriod: The number of seconds between scheduled key rollovers, or 0xFFFFFFFF
	// to disable automatic key rollovers. This value SHOULD be limited to the range 0x00093A80
	// (1 week) to 0x25980600 (20 years), inclusive, in addition to 0xFFFFFFFF, when fIsKSK
	// is 0x00000001, and the range 0x00093A80 (1 week) to 0x09660180 (5 years), inclusive,
	// in addition to 0xFFFFFFFF, when fIsKSK is 0x00000000. The default SHOULD be 0x02022900
	// (13 months) when fIsKSK is 0x00000001, and 0x0x00278D00 (1 month) when fIsKSK is
	// 0x00000000.
	RolloverPeriod uint32 `idl:"name:dwRolloverPeriod" json:"rollover_period"`
	// dwNextRolloverAction: This value describes the next key rollover action that the
	// DNS server will take for this signing key descriptor. This value MUST be set to one
	// of the allowed values specified in ZONE_SKD_ROLLOVER_ACTION (section 2.2.5.1.6).
	NextRolloverAction uint32 `idl:"name:dwNextRolloverAction" json:"next_rollover_action"`
	// dwReserved: This value MUST be set to 0x00000000 when sent by the client and ignored
	// on receipt by the server.
	_ uint32 `idl:"name:dwReserved"`
}

func (o *SKD) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SKD) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
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
	if o.KeyStorageProvider != "" {
		_ptr_pwszKeyStorageProvider := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.KeyStorageProvider); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.KeyStorageProvider, _ptr_pwszKeyStorageProvider); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if !o.StoreKeysInDirectory {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.IsKSK {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SigningAlgorithm); err != nil {
		return err
	}
	if err := w.WriteData(o.KeyLength); err != nil {
		return err
	}
	if err := w.WriteData(o.InitialRolloverOffset); err != nil {
		return err
	}
	if err := w.WriteData(o.DNSKEYSignatureValidityPeriod); err != nil {
		return err
	}
	if err := w.WriteData(o.DSSignatureValidityPeriod); err != nil {
		return err
	}
	if err := w.WriteData(o.StandardSignatureValidityPeriod); err != nil {
		return err
	}
	if err := w.WriteData(o.RolloverType); err != nil {
		return err
	}
	if err := w.WriteData(o.RolloverPeriod); err != nil {
		return err
	}
	if err := w.WriteData(o.NextRolloverAction); err != nil {
		return err
	}
	// reserved dwReserved
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *SKD) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	if o.GUID == nil {
		o.GUID = &dtyp.GUID{}
	}
	if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_pwszKeyStorageProvider := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.KeyStorageProvider); err != nil {
			return err
		}
		return nil
	})
	_s_pwszKeyStorageProvider := func(ptr interface{}) { o.KeyStorageProvider = *ptr.(*string) }
	if err := w.ReadPointer(&o.KeyStorageProvider, _s_pwszKeyStorageProvider, _ptr_pwszKeyStorageProvider); err != nil {
		return err
	}
	var _bStoreKeysInDirectory int32
	if err := w.ReadData(&_bStoreKeysInDirectory); err != nil {
		return err
	}
	o.StoreKeysInDirectory = _bStoreKeysInDirectory != 0
	var _bIsKSK int32
	if err := w.ReadData(&_bIsKSK); err != nil {
		return err
	}
	o.IsKSK = _bIsKSK != 0
	if err := w.ReadData(&o.SigningAlgorithm); err != nil {
		return err
	}
	if err := w.ReadData(&o.KeyLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.InitialRolloverOffset); err != nil {
		return err
	}
	if err := w.ReadData(&o.DNSKEYSignatureValidityPeriod); err != nil {
		return err
	}
	if err := w.ReadData(&o.DSSignatureValidityPeriod); err != nil {
		return err
	}
	if err := w.ReadData(&o.StandardSignatureValidityPeriod); err != nil {
		return err
	}
	if err := w.ReadData(&o.RolloverType); err != nil {
		return err
	}
	if err := w.ReadData(&o.RolloverPeriod); err != nil {
		return err
	}
	if err := w.ReadData(&o.NextRolloverAction); err != nil {
		return err
	}
	// reserved dwReserved
	var _dwReserved uint32
	if err := w.ReadData(&_dwReserved); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// KeySignScope type represents KeySignScope RPC enumeration.
//
// The KeySignScope enumeration defines the scope of the key in a signed zone. It SHOULD<46>
// be used to indicate different signing and publishing scopes of the key.
type KeySignScope uint16

var (
	// SIGN_SCOPE_DEFAULT: The key is used for its default purpose: If the signing key descriptor's
	// fIsKSK field is set to 0x00000001, the key is used to sign only DNSKEY records in
	// the zone. If the signing key descriptor's fIsKSK field is set to 0x00000000, the
	// key is used to sign all records in the zone.
	KeySignScopeDefault KeySignScope = 0
	// SIGN_SCOPE_DNSKEY_ONLY: The key is used to sign only DNSKEY records in the zone.
	KeySignScopeDNSKEYOnly KeySignScope = 1
	// SIGN_SCOPE_ALL_RECORDS: The key is used to sign all records in the zone.
	KeySignScopeAllRecords KeySignScope = 2
	// SIGN_SCOPE_ADD_ONLY: The key is published as a DNSKEY in the zone, but it is not
	// used to sign any records.
	KeySignScopeAddOnly KeySignScope = 3
	// SIGN_SCOPE_DO_NOT_PUBLISH: The key is not published to the zone and is not used to
	// sign any records.
	KeySignScopeDoNotPublish KeySignScope = 4
	// SIGN_SCOPE_REVOKED: The key is published as a DNSKEY in the zone with its Revoked
	// bit ([RFC5011] section 2.1) set. It is used to sign DNSKEY records.
	KeySignScopeRevoked KeySignScope = 5
)

func (o KeySignScope) String() string {
	switch o {
	case KeySignScopeDefault:
		return "KeySignScopeDefault"
	case KeySignScopeDNSKEYOnly:
		return "KeySignScopeDNSKEYOnly"
	case KeySignScopeAllRecords:
		return "KeySignScopeAllRecords"
	case KeySignScopeAddOnly:
		return "KeySignScopeAddOnly"
	case KeySignScopeDoNotPublish:
		return "KeySignScopeDoNotPublish"
	case KeySignScopeRevoked:
		return "KeySignScopeRevoked"
	}
	return "Invalid"
}

// SKDStateEx structure represents DNS_RPC_SKD_STATE_EX RPC structure.
//
// The DNS_RPC_SKD_STATE_EX structure <52> represents the collection of extended dynamic
// configuration information of a signing key descriptor state.
type SKDStateEx struct {
	// dwRpcStructureVersion: The structure version number; this MUST be set to 0x00000001.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// dwReserved0: This MUST be set to zero when sent and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved0"`
	// Guid: A unique identifier for this signing key descriptor.
	GUID *dtyp.GUID `idl:"name:Guid" json:"guid"`
	// dwCurrentRollState: The current rollover status of this signing key descriptor. Note
	// that these are DNS Server Management Protocol implementations of the protocol specified
	// rollover states of SKDs in DNS_RPC_SKD_STATE.
	//
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                                             |                                    MAPPING IN                                    |
	//	|                            VALUE                            |                                DNS RPC SKD STATE                                 |
	//	|                                                             |                                                                                  |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_SKD_ROLL_STATE_NOT_STARTED 0x00000000                   | DNS_SKD_STATUS_NOT_ROLLING                                                       |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_SKD_ROLL_STATE_ZSK_SWAP_ACTIVE_WITH_ROLLOVER 0x00000001 | DNS_SKD_STATUS_ZSK_WAITING_FOR_DNSKEY_TTL                                        |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_SKD_ROLL_STATE_ZSK_FINISH 0x00000002                    | DNS_SKD_STATUS_ZSK_WAITING_FOR_MAXZONE_TTL                                       |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_SKD_ROLL_STATE_KSK_WAITING_FOR_DS 0x00000003            | DNS_SKD_STATUS_KSK_WAITING_FOR_DS_UPDATE                                         |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_SKD_ROLL_STATE_KSK_REVOKE 0x00000004                    | DNS_SKD_STATUS_KSK_WAITING_FOR_DS_TTL The key is revoked once the rollover is    |
	//	|                                                             | completed.                                                                       |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_SKD_ROLL_STATE_KSK_FINISH 0x00000005                    | DNS_SKD_STATUS_KSK_WAITING_FOR_DS_TTL The key is not revoked and is used with    |
	//	|                                                             | the rolled over key-set.                                                         |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_SKD_ROLL_STATE_MAX 0x00000005                           | DNS_SKD_STATUS_KSK_WAITING_FOR_DS_TTL                                            |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_SKD_STATUS_ROLL_ERROR 0x00000009                        | This signing key descriptor experienced an unrecoverable error during the key    |
	//	|                                                             | rollover.                                                                        |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	CurrentRollState uint32 `idl:"name:dwCurrentRollState" json:"current_roll_state"`
	// fManualTrigger: TRUE if the rollover was triggered manually.
	ManualTrigger uint32 `idl:"name:fManualTrigger" json:"manual_trigger"`
	// dwPreRollEventFired: Specifies which prerollover event has been fired.
	PreRollEventFired uint32 `idl:"name:dwPreRollEventFired" json:"pre_roll_event_fired"`
	// ftNextKeyGenerationTime: The time at which the next key was added to the zone.
	NextKeyGenerationTime *dtyp.Filetime `idl:"name:ftNextKeyGenerationTime" json:"next_key_generation_time"`
	// dwRevokedOrSwappedDnskeysLength: The length of the revoked or swapped DNSKEY RRSet.
	RevokedOrSwappedDNSKeysLength uint32 `idl:"name:dwRevokedOrSwappedDnskeysLength" json:"revoked_or_swapped_dns_keys_length"`
	// pRevokedOrSwappedDnskeysBuffer: Presigned DNSKEY RRSet for ZSK swap. This is a pointer
	// to a pointer that points to a buffer containing the enumerated records. The buffer
	// is a series of structures beginning with a DNS_RPC_NODE structure (section 2.2.2.2.3).
	// The records for the node will be represented by a series of DNS_RPC_RECORD structures
	// (section 2.2.2.2.5). The number of DNS_RPC_RECORD structures following a DNS_RPC_NODE
	// structure is given by the wRecordCount member of DNS_RPC_NODE.
	RevokedOrSwappedDNSKeysBuffer []byte `idl:"name:pRevokedOrSwappedDnskeysBuffer;size_is:(dwRevokedOrSwappedDnskeysLength)" json:"revoked_or_swapped_dns_keys_buffer"`
	// dwFinalDnskeysLength: Length of presigned DNSKEY RRSet.
	FinalDNSKeysLength uint32 `idl:"name:dwFinalDnskeysLength" json:"final_dns_keys_length"`
	// pFinalDnskeys: Presigned DNSKEY RRSet Post Rollover. This is a pointer to a pointer
	// that points to a buffer containing the enumerated records. The buffer is a series
	// of structures beginning with a DNS_RPC_NODE structure. The records for the node will
	// be represented by a series of DNS_RPC_RECORD structures. The number of DNS_RPC_RECORD
	// structures following a DNS_RPC_NODE structure is given by the wRecordCount member
	// of DNS_RPC_NODE.
	FinalDNSKeys []byte `idl:"name:pFinalDnskeys;size_is:(dwFinalDnskeysLength)" json:"final_dns_keys"`
	// eActiveKeyScope: Signing key scope for the SKD's active key. The signing key scope
	// is defined in KeySignScope (section 2.2.6.1.2).
	ActiveKeyScope KeySignScope `idl:"name:eActiveKeyScope" json:"active_key_scope"`
	// eStandByKeyScope: Signing key scope for the SKD's standby key. The signing key scope
	// is defined in KeySignScope.
	StandByKeyScope KeySignScope `idl:"name:eStandByKeyScope" json:"stand_by_key_scope"`
	// eNextKeyScope: Signing key scope for the SKD's next key. The signing key scope is
	// defined in KeySignScope.
	NextKeyScope KeySignScope `idl:"name:eNextKeyScope" json:"next_key_scope"`
}

func (o *SKDStateEx) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.RevokedOrSwappedDNSKeysBuffer != nil && o.RevokedOrSwappedDNSKeysLength == 0 {
		o.RevokedOrSwappedDNSKeysLength = uint32(len(o.RevokedOrSwappedDNSKeysBuffer))
	}
	if o.FinalDNSKeys != nil && o.FinalDNSKeysLength == 0 {
		o.FinalDNSKeysLength = uint32(len(o.FinalDNSKeys))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SKDStateEx) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
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
	if err := w.WriteData(o.CurrentRollState); err != nil {
		return err
	}
	if err := w.WriteData(o.ManualTrigger); err != nil {
		return err
	}
	if err := w.WriteData(o.PreRollEventFired); err != nil {
		return err
	}
	if o.NextKeyGenerationTime != nil {
		if err := o.NextKeyGenerationTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.RevokedOrSwappedDNSKeysLength); err != nil {
		return err
	}
	if o.RevokedOrSwappedDNSKeysBuffer != nil || o.RevokedOrSwappedDNSKeysLength > 0 {
		_ptr_pRevokedOrSwappedDnskeysBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.RevokedOrSwappedDNSKeysLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.RevokedOrSwappedDNSKeysBuffer {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.RevokedOrSwappedDNSKeysBuffer[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.RevokedOrSwappedDNSKeysBuffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RevokedOrSwappedDNSKeysBuffer, _ptr_pRevokedOrSwappedDnskeysBuffer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.FinalDNSKeysLength); err != nil {
		return err
	}
	if o.FinalDNSKeys != nil || o.FinalDNSKeysLength > 0 {
		_ptr_pFinalDnskeys := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.FinalDNSKeysLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.FinalDNSKeys {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.FinalDNSKeys[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.FinalDNSKeys); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.FinalDNSKeys, _ptr_pFinalDnskeys); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteEnum(uint16(o.ActiveKeyScope)); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.StandByKeyScope)); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.NextKeyScope)); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *SKDStateEx) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	if o.GUID == nil {
		o.GUID = &dtyp.GUID{}
	}
	if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.CurrentRollState); err != nil {
		return err
	}
	if err := w.ReadData(&o.ManualTrigger); err != nil {
		return err
	}
	if err := w.ReadData(&o.PreRollEventFired); err != nil {
		return err
	}
	if o.NextKeyGenerationTime == nil {
		o.NextKeyGenerationTime = &dtyp.Filetime{}
	}
	if err := o.NextKeyGenerationTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.RevokedOrSwappedDNSKeysLength); err != nil {
		return err
	}
	_ptr_pRevokedOrSwappedDnskeysBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.RevokedOrSwappedDNSKeysLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.RevokedOrSwappedDNSKeysLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.RevokedOrSwappedDNSKeysBuffer", sizeInfo[0])
		}
		o.RevokedOrSwappedDNSKeysBuffer = make([]byte, sizeInfo[0])
		for i1 := range o.RevokedOrSwappedDNSKeysBuffer {
			i1 := i1
			if err := w.ReadData(&o.RevokedOrSwappedDNSKeysBuffer[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pRevokedOrSwappedDnskeysBuffer := func(ptr interface{}) { o.RevokedOrSwappedDNSKeysBuffer = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.RevokedOrSwappedDNSKeysBuffer, _s_pRevokedOrSwappedDnskeysBuffer, _ptr_pRevokedOrSwappedDnskeysBuffer); err != nil {
		return err
	}
	if err := w.ReadData(&o.FinalDNSKeysLength); err != nil {
		return err
	}
	_ptr_pFinalDnskeys := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.FinalDNSKeysLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.FinalDNSKeysLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.FinalDNSKeys", sizeInfo[0])
		}
		o.FinalDNSKeys = make([]byte, sizeInfo[0])
		for i1 := range o.FinalDNSKeys {
			i1 := i1
			if err := w.ReadData(&o.FinalDNSKeys[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pFinalDnskeys := func(ptr interface{}) { o.FinalDNSKeys = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.FinalDNSKeys, _s_pFinalDnskeys, _ptr_pFinalDnskeys); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.ActiveKeyScope)); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.StandByKeyScope)); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.NextKeyScope)); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// SKDState structure represents DNS_RPC_SKD_STATE RPC structure.
//
// The DNS_RPC_SKD_STATE structure contains information about the current state of a
// signing key descriptor.
type SKDState struct {
	// dwRpcStructureVersion: The structure version number; this MUST be set to 0x00000001.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// dwReserved0: MUST be set to zero when sent and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved0"`
	// Guid: A unique identifier for this signing key descriptor.
	GUID *dtyp.GUID `idl:"name:Guid" json:"guid"`
	// ftLastRolloverTime: The time at which the last successful rollover event was performed
	// for this signing key descriptor. This value MUST be 0 if no rollover has been performed
	// on this signing key descriptor.
	LastRolloverTime *dtyp.Filetime `idl:"name:ftLastRolloverTime" json:"last_rollover_time"`
	// ftNextRolloverTime: The time at which the next rollover for this signing key descriptor
	// is scheduled. This MAY be 0 if no rollover event is scheduled. When the time comes
	// for a key rollover to start, the signing key descriptor is added to the Rollover
	// Queue, and its dwCurrentRolloverStatus is changed to DNS_SKD_STATUS_QUEUED. If another
	// signing key descriptor in the zone is in the process of rolling, ftNextRolloverTime
	// MAY represent a time in the past.
	NextRolloverTime *dtyp.Filetime `idl:"name:ftNextRolloverTime" json:"next_rollover_time"`
	// dwState: The current state of this signing key descriptor. This MUST be set to one
	// of the following values.<49>
	//
	//	+----------------------------------+---------------------------------------------------------------------------------+
	//	|                                  |                                                                                 |
	//	|              VALUE               |                                     MEANING                                     |
	//	|                                  |                                                                                 |
	//	+----------------------------------+---------------------------------------------------------------------------------+
	//	+----------------------------------+---------------------------------------------------------------------------------+
	//	| DNS_SKD_STATE_ACTIVE 0x00000000  | The signing key descriptor is active and in use for online signing of the zone. |
	//	+----------------------------------+---------------------------------------------------------------------------------+
	//	| DNS_SKD_STATE_RETIRED 0x00000001 | The signing key descriptor is no longer in use for online signing.              |
	//	+----------------------------------+---------------------------------------------------------------------------------+
	State uint32 `idl:"name:dwState" json:"state"`
	// dwCurrentRolloverStatus: The current rollover status of this signing key descriptor.
	// This MUST be set to one of the following values, representing the various stages
	// of key rollover as described in [RFC4641] and [RFC5011]:
	//
	//	+-----------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                                                 |                                                                                  |
	//	|                              VALUE                              |                                     MEANING                                      |
	//	|                                                                 |                                                                                  |
	//	+-----------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_SKD_STATUS_NOT_ROLLING 0x00000000                           | The signing key descriptor is not currently in the process of rolling over keys. |
	//	+-----------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_SKD_STATUS_QUEUED 0x00000001                                | This signing key descriptor is waiting for another rollover to complete          |
	//	|                                                                 | before its rollover can begin. After the signing key descriptor reaches the      |
	//	|                                                                 | head of the Rollover Queue, the signing key descriptor will move into the        |
	//	|                                                                 | DNS_SKD_STATUS_ROLL_STARTED state.                                               |
	//	+-----------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_SKD_STATUS_ROLL_STARTED 0x00000002                          | This signing key descriptor has begun the process of key rollover.               |
	//	|                                                                 | Signing key descriptors representing ZSKs will move from this state              |
	//	|                                                                 | to the DNS_SKD_STATUS_ZSK_WAITING_FOR_DNSKEY_TTL state, and signing              |
	//	|                                                                 | key descriptors representing KSKs will move from this state to the               |
	//	|                                                                 | DNS_SKD_STATUS_KSK_WAITING_FOR_DS_UPDATE state.                                  |
	//	+-----------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_SKD_STATUS_ZSK_WAITING_FOR_DNSKEY_TTL 0x00000003            | This ZSK signing key descriptor is waiting for the previous key to expire in all |
	//	|                                                                 | caching resolvers (it is waiting for the DNSKEY TTL to expire). The signing key  |
	//	|                                                                 | descriptor will next move into the DNS_SKD_STATUS_ZSK_WAITING_FOR_MAXZONE_TTL    |
	//	|                                                                 | state.                                                                           |
	//	+-----------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_SKD_STATUS_ZSK_WAITING_FOR_MAXZONE_TTL 0x00000004           | This ZSK signing key descriptor is waiting for the signatures using the previous |
	//	|                                                                 | key to expire in all caching resolvers (it is waiting for the largest record     |
	//	|                                                                 | TTL to expire). This is the final state for ZSK rollovers. The signing key       |
	//	|                                                                 | descriptor will move into the DNS_SKD_STATUS_NOT_ROLLING state when this portion |
	//	|                                                                 | of key rollover is complete.                                                     |
	//	+-----------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_SKD_STATUS_KSK_WAITING_FOR_DS_UPDATE 0x00000005             | This KSK signing key descriptor is waiting for a DS record corresponding to      |
	//	|                                                                 | the new key to appear in the parent zone. If DS records are found in the parent  |
	//	|                                                                 | zone, the server MUST set the zone's "ParentHasSecureDelegation" property        |
	//	|                                                                 | to 0x00000001 and continue to wait until the parent's DS record set includes     |
	//	|                                                                 | the new key, at which point the signing key descriptor will move into the        |
	//	|                                                                 | DNS_SKD_STATUS_KSK_WAITING_FOR_DS_TTL state. If no DS records are found in       |
	//	|                                                                 | the parent zone, the server MUST set the zone's "ParentHasSecureDelegation"      |
	//	|                                                                 | property to 0x00000000 and transition the signing key descriptor into the        |
	//	|                                                                 | DNS_SKD_STATUS_KSK_WAITING_FOR_DNSKEY_TTL state. If there is an error or if the  |
	//	|                                                                 | presence of a DS record set in the parent zone cannot be determined, the server  |
	//	|                                                                 | MUST continue to attempt to query for this record set for up to 15 minutes       |
	//	|                                                                 | if the zone's "ParentHasSecureDelegation" property is 0x00000000 or until the    |
	//	|                                                                 | PokeZoneKeyRollover command is received if it is 0x00000001.                     |
	//	+-----------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_SKD_STATUS_KSK_WAITING_FOR_DS_TTL 0x00000006                | This KSK signing key descriptor is waiting for the DS record set in              |
	//	|                                                                 | the parent zone to expire in all caching resolvers (it is waiting                |
	//	|                                                                 | for the parent DS TTL to expire). If the zone's "RFC5011KeyRollovers"            |
	//	|                                                                 | property is TRUE, the signing key descriptor will next move into the             |
	//	|                                                                 | DNS_SKD_STATUS_WAITING_FOR_5011_REMOVE_HOLD_DOWN state. Otherwise, this is a     |
	//	|                                                                 | final state for KSK rollovers, and signing key descriptors will move into the    |
	//	|                                                                 | DNS_SKD_STATUS_NOT_ROLLING state when this portion of key rollover is complete.  |
	//	+-----------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_SKD_STATUS_KSK_WAITING_FOR_DNSKEY_TTL 0x00000007            | This KSK signing key descriptor is waiting for the previous key to expire in     |
	//	|                                                                 | all caching resolvers (it is waiting for the DNSKEY TTL to expire). This is a    |
	//	|                                                                 | final state for KSK rollovers, and signing key descriptors will move into the    |
	//	|                                                                 | DNS_SKD_STATUS_NOT_ROLLING state when this portion of key rollover is complete.  |
	//	+-----------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_SKD_STATUS_KSK_WAITING_FOR_5011_REMOVE_HOLD_DOWN 0x00000008 | This KSK signing key descriptor is waiting for the RFC5011 remove hold-down      |
	//	|                                                                 | time before the revoked previous key can be removed. This is a final             |
	//	|                                                                 | state for KSK rollovers, and signing key descriptors will move into the          |
	//	|                                                                 | DNS_SKD_STATUS_NOT_ROLLING state when this portion of key rollover is complete.  |
	//	+-----------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_SKD_STATUS_ROLL_ERROR 0x00000009                            | This signing key descriptor experienced an unrecoverable error during the key    |
	//	|                                                                 | rollover.                                                                        |
	//	+-----------------------------------------------------------------+----------------------------------------------------------------------------------+
	CurrentRolloverStatus uint32 `idl:"name:dwCurrentRolloverStatus" json:"current_rollover_status"`
	// pwszActiveKey: Signing key pointer string for the signing key descriptor's currently
	// active key. This is the key that is currently being used to sign zone records.
	ActiveKey string `idl:"name:pwszActiveKey;string" json:"active_key"`
	// pwszStandbyKey: Signing key pointer string for the signing key descriptor's standby
	// key. The key represented by this string has several different roles depending on
	// the associated signing key descriptor's fIsKSK flag and its dwRolloverStatus:
	//
	// * If the signing key descriptor's fIsKSK flag is TRUE, pwszStandbyKey represents
	// the "double signature" key as described in [RFC4641] section 4.2.2, also depicted
	// as the "standby" key in [RFC5011] section 6.
	//
	// * If the signing key descriptor's fIsKSK flag is FALSE, *pwszStandbyKey* will generally
	// be NULL unless the SKD is in the process of key rollover:
	//
	// * If the signing key descriptor's dwRolloverStatus is DNS_SKD_STATUS_ZSK_WAITING_FOR_DNSKEY_TTL,
	// *pwszStandbyKey* represents the "pre-publish" key as described in [RFC4641] section
	// 4.2.1.1.
	//
	// * If the signing key descriptor's dwRolloverStatus is DNS_SKD_STATUS_ZSK_WAITING_FOR_MAXZONE_TTL,
	// *pwszStandbyKey* represents the previously active key during the "new RRSIGs" phase
	// of Pre-Publish Key Rollover, as described in [RFC4641] section 4.2.1.1.
	StandbyKey string `idl:"name:pwszStandbyKey;string" json:"standby_key"`
	// pwszNextKey: Signing key pointer string for the signing key descriptor's next key.
	// The key represented by this string has several different roles depending on the associated
	// signing key descriptor's fIsKSK flag:
	//
	// * If the signing key descriptor's fIsKSK flag is TRUE, *pwszNextKey* represents the
	// next key that will be consumed during key rollover. It is not published in the zone
	// and is not used to sign any other records. If the zone's "RFC5011KeyRollovers" property
	// is TRUE, this key becomes the next "standby" key according to [RFC5011]. Otherwise,
	// this key is used as the "double signature" key according to [RFC4641] as the signing
	// key descriptor's key rollover process begins.
	//
	// * If the signing key descriptor's fIsKSK flag is FALSE, *pwszNextKey* represents
	// the "pre-publish" key as described in [RFC4641] section 4.2.1.1. When the SKD is
	// in the process of key rollover, *pwszNextKey* can be populated with a newly generated
	// post-rollover "pre-publish" key.
	NextKey string `idl:"name:pwszNextKey;string" json:"next_key"`
	// dwReserved: MUST be set to zero when sent and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved"`
}

func (o *SKDState) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SKDState) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
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
	if o.LastRolloverTime != nil {
		if err := o.LastRolloverTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.NextRolloverTime != nil {
		if err := o.NextRolloverTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.CurrentRolloverStatus); err != nil {
		return err
	}
	if o.ActiveKey != "" {
		_ptr_pwszActiveKey := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ActiveKey); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ActiveKey, _ptr_pwszActiveKey); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.StandbyKey != "" {
		_ptr_pwszStandbyKey := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.StandbyKey); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.StandbyKey, _ptr_pwszStandbyKey); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.NextKey != "" {
		_ptr_pwszNextKey := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.NextKey); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.NextKey, _ptr_pwszNextKey); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	// reserved dwReserved
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *SKDState) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	if o.GUID == nil {
		o.GUID = &dtyp.GUID{}
	}
	if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LastRolloverTime == nil {
		o.LastRolloverTime = &dtyp.Filetime{}
	}
	if err := o.LastRolloverTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.NextRolloverTime == nil {
		o.NextRolloverTime = &dtyp.Filetime{}
	}
	if err := o.NextRolloverTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if err := w.ReadData(&o.CurrentRolloverStatus); err != nil {
		return err
	}
	_ptr_pwszActiveKey := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ActiveKey); err != nil {
			return err
		}
		return nil
	})
	_s_pwszActiveKey := func(ptr interface{}) { o.ActiveKey = *ptr.(*string) }
	if err := w.ReadPointer(&o.ActiveKey, _s_pwszActiveKey, _ptr_pwszActiveKey); err != nil {
		return err
	}
	_ptr_pwszStandbyKey := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.StandbyKey); err != nil {
			return err
		}
		return nil
	})
	_s_pwszStandbyKey := func(ptr interface{}) { o.StandbyKey = *ptr.(*string) }
	if err := w.ReadPointer(&o.StandbyKey, _s_pwszStandbyKey, _ptr_pwszStandbyKey); err != nil {
		return err
	}
	_ptr_pwszNextKey := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.NextKey); err != nil {
			return err
		}
		return nil
	})
	_s_pwszNextKey := func(ptr interface{}) { o.NextKey = *ptr.(*string) }
	if err := w.ReadPointer(&o.NextKey, _s_pwszNextKey, _ptr_pwszNextKey); err != nil {
		return err
	}
	// reserved dwReserved
	var _dwReserved uint32
	if err := w.ReadData(&_dwReserved); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// ZoneSKD structure represents DNS_RPC_ZONE_SKD RPC structure.
//
// The DNS_RPC_ZONE_SKD structure<51> groups all the properties of a signing key descriptor
// of a zone.
type ZoneSKD struct {
	// dwRpcStructureVersion: The structure version number; this MUST be set to 0x00000001.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// dwReserved0: MUST be set to zero when sent and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved0"`
	// pSkd: The static properties of a zone SKD. This is a DNS_RPC_SKD structure (section
	// 2.2.6.2.1).
	SKD *SKD `idl:"name:pSkd" json:"skd"`
	// pSkdState: The dynamic properties of a zone SKD. This is a DNS_RPC_SKD_STATE structure
	// (section 2.2.6.2.3).
	SKDState *SKDState `idl:"name:pSkdState" json:"skd_state"`
	// pSkdStateEx: The extended dynamic properties of a zone SKD. This is a DNS_RPC_SKD_STATE_EX
	// structure (section 2.2.6.2.11).
	SKDStateEx *SKDStateEx `idl:"name:pSkdStateEx" json:"skd_state_ex"`
}

func (o *ZoneSKD) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ZoneSKD) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if o.SKD != nil {
		_ptr_pSkd := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SKD != nil {
				if err := o.SKD.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SKD{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SKD, _ptr_pSkd); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.SKDState != nil {
		_ptr_pSkdState := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SKDState != nil {
				if err := o.SKDState.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SKDState{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SKDState, _ptr_pSkdState); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.SKDStateEx != nil {
		_ptr_pSkdStateEx := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SKDStateEx != nil {
				if err := o.SKDStateEx.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SKDStateEx{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SKDStateEx, _ptr_pSkdStateEx); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ZoneSKD) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	_ptr_pSkd := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SKD == nil {
			o.SKD = &SKD{}
		}
		if err := o.SKD.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pSkd := func(ptr interface{}) { o.SKD = *ptr.(**SKD) }
	if err := w.ReadPointer(&o.SKD, _s_pSkd, _ptr_pSkd); err != nil {
		return err
	}
	_ptr_pSkdState := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SKDState == nil {
			o.SKDState = &SKDState{}
		}
		if err := o.SKDState.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pSkdState := func(ptr interface{}) { o.SKDState = *ptr.(**SKDState) }
	if err := w.ReadPointer(&o.SKDState, _s_pSkdState, _ptr_pSkdState); err != nil {
		return err
	}
	_ptr_pSkdStateEx := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SKDStateEx == nil {
			o.SKDStateEx = &SKDStateEx{}
		}
		if err := o.SKDStateEx.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pSkdStateEx := func(ptr interface{}) { o.SKDStateEx = *ptr.(**SKDStateEx) }
	if err := w.ReadPointer(&o.SKDStateEx, _s_pSkdStateEx, _ptr_pSkdStateEx); err != nil {
		return err
	}
	return nil
}

// ZoneDNSSECSettings structure represents DNS_RPC_ZONE_DNSSEC_SETTINGS RPC structure.
//
// The DNS_RPC_ZONE_DNSSEC_SETTINGS structure SHOULD<50>  represent the DNSSEC properties
// of a zone.
type ZoneDNSSECSettings struct {
	// dwRpcStructureVersion: The structure version number; this MUST be set to 0x00000001.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// dwReserved0: MUST be set to zero when sent and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved0"`
	// fIsSigned: States whether or not the zone is signed.
	IsSigned uint32 `idl:"name:fIsSigned" json:"is_signed"`
	// fSignWithNSEC3: States whether the zone has NSEC3 Authenticated Denial of Existence
	// support.
	SignWithNSEC3 uint32 `idl:"name:fSignWithNSEC3" json:"sign_with_nsec3"`
	// fNSEC3OptOut: States whether the zone has unsecure delegations.
	NSEC3OptOut uint32 `idl:"name:fNSEC3OptOut" json:"nsec3_opt_out"`
	// dwMaintainTrustAnchor: States whether Trust Anchors are maintained for this zone.
	MaintainTrustAnchor uint32 `idl:"name:dwMaintainTrustAnchor" json:"maintain_trust_anchor"`
	// fParentHasSecureDelegation: Delegation Status (Secure\UnSecure) for this zone from
	// the parent.
	ParentHasSecureDelegation uint32 `idl:"name:fParentHasSecureDelegation" json:"parent_has_secure_delegation"`
	// dwDSRecordAlgorithms: Algorithms used for generating a hash of the DNSKEY record.
	DSRecordAlgorithms uint32 `idl:"name:dwDSRecordAlgorithms" json:"ds_record_algorithms"`
	// fRFC5011KeyRollovers: 5011 rollover status (Enabled\Disabled) for this zone.
	Rfc5011KeyRollovers uint32 `idl:"name:fRFC5011KeyRollovers" json:"rfc5011_key_rollovers"`
	// bNSEC3HashAlgorithm: Algorithm used for generating NSEC3 hash (see [RFC5155] section
	// 5).
	NSEC3HashAlgorithm uint8 `idl:"name:bNSEC3HashAlgorithm" json:"nsec3_hash_algorithm"`
	// bNSEC3RandomSaltLength: Length of Salt used in generating NSEC3 records for this
	// zone (see [RFC5155] section 5).
	NSEC3RandomSaltLength uint8 `idl:"name:bNSEC3RandomSaltLength" json:"nsec3_random_salt_length"`
	// wNSEC3IterationCount: Iteration count for generating NSEC3 records for this zone
	// (see [RFC5155] section 5).
	NSEC3IterationCount uint16 `idl:"name:wNSEC3IterationCount" json:"nsec3_iteration_count"`
	// pwszNSEC3UserSalt: User defined salt used for generating NSEC3 records for this zone
	// (see [RFC5155] section 5).
	NSEC3UserSalt string `idl:"name:pwszNSEC3UserSalt" json:"nsec3_user_salt"`
	// dwDNSKEYRecordSetTtl: Time-to-live (TTL) for the DNSKEY resource record.
	DNSKEYRecordSetTTL uint32 `idl:"name:dwDNSKEYRecordSetTtl" json:"dnskey_record_set_ttl"`
	// dwDSRecordSetTtl: TTL for the DS Resource Record.
	DSRecordSetTTL uint32 `idl:"name:dwDSRecordSetTtl" json:"ds_record_set_ttl"`
	// dwSignatureInceptionOffset: Time in seconds for Inception of Signatures for RRSIGs
	// as defined in [RFC4034] section 3.1.5.
	SignatureInceptionOffset uint32 `idl:"name:dwSignatureInceptionOffset" json:"signature_inception_offset"`
	// dwSecureDelegationPollingPeriod: The interval, in seconds, to refresh the set of
	// delegation signer (DS) records in a secure delegation.
	SecureDelegationPollingPeriod uint32 `idl:"name:dwSecureDelegationPollingPeriod" json:"secure_delegation_polling_period"`
	// dwPropagationTime: The time, in seconds, that it takes for zone data changes to propagate
	// to other copies of the zone.
	PropagationTime uint32 `idl:"name:dwPropagationTime" json:"propagation_time"`
	// cbNSEC3CurrentSaltLength: Length of the Current User salt for building an NSEC3 chain
	// of zone records.
	NSEC3CurrentSaltLength uint32 `idl:"name:cbNSEC3CurrentSaltLength" json:"nsec3_current_salt_length"`
	// pbNSEC3CurrentSalt: Pointer to the pwszNSEC3UserSalt for building an NSEC3 chain
	// of zone records.
	NSEC3CurrentSalt []byte `idl:"name:pbNSEC3CurrentSalt;size_is:(cbNSEC3CurrentSaltLength)" json:"nsec3_current_salt"`
	// CurrentRollingSKDGuid: Unique identifier of a rolling SKD of a zone, if any.
	CurrentRollingSKDGUID *dtyp.GUID `idl:"name:CurrentRollingSKDGuid" json:"current_rolling_skd_guid"`
	BufferLength          uint32     `idl:"name:dwBufferLength" json:"buffer_length"`
	Buffer                []byte     `idl:"name:pBuffer;size_is:(dwBufferLength)" json:"buffer"`
	// dwCount: The number of signing key descriptors present in the array of signing key
	// descriptors pointed to by SkdArray.
	Count uint32 `idl:"name:dwCount" json:"count"`
	// pZoneSkdArray: A list of SKDs for a zone. This is the array of the DNS_RPC_ZONE_SKD
	// structure.
	ZoneSKDArray []*ZoneSKD `idl:"name:pZoneSkdArray;size_is:(dwCount)" json:"zone_skd_array"`
}

func (o *ZoneDNSSECSettings) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.NSEC3CurrentSalt != nil && o.NSEC3CurrentSaltLength == 0 {
		o.NSEC3CurrentSaltLength = uint32(len(o.NSEC3CurrentSalt))
	}
	if o.Buffer != nil && o.BufferLength == 0 {
		o.BufferLength = uint32(len(o.Buffer))
	}
	if o.ZoneSKDArray != nil && o.Count == 0 {
		o.Count = uint32(len(o.ZoneSKDArray))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *ZoneDNSSECSettings) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.Count)
	return []uint64{
		dimSize1,
	}
}
func (o *ZoneDNSSECSettings) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.IsSigned); err != nil {
		return err
	}
	if err := w.WriteData(o.SignWithNSEC3); err != nil {
		return err
	}
	if err := w.WriteData(o.NSEC3OptOut); err != nil {
		return err
	}
	if err := w.WriteData(o.MaintainTrustAnchor); err != nil {
		return err
	}
	if err := w.WriteData(o.ParentHasSecureDelegation); err != nil {
		return err
	}
	if err := w.WriteData(o.DSRecordAlgorithms); err != nil {
		return err
	}
	if err := w.WriteData(o.Rfc5011KeyRollovers); err != nil {
		return err
	}
	if err := w.WriteData(o.NSEC3HashAlgorithm); err != nil {
		return err
	}
	if err := w.WriteData(o.NSEC3RandomSaltLength); err != nil {
		return err
	}
	if err := w.WriteData(o.NSEC3IterationCount); err != nil {
		return err
	}
	if o.NSEC3UserSalt != "" {
		_ptr_pwszNSEC3UserSalt := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.NSEC3UserSalt); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.NSEC3UserSalt, _ptr_pwszNSEC3UserSalt); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.DNSKEYRecordSetTTL); err != nil {
		return err
	}
	if err := w.WriteData(o.DSRecordSetTTL); err != nil {
		return err
	}
	if err := w.WriteData(o.SignatureInceptionOffset); err != nil {
		return err
	}
	if err := w.WriteData(o.SecureDelegationPollingPeriod); err != nil {
		return err
	}
	if err := w.WriteData(o.PropagationTime); err != nil {
		return err
	}
	if err := w.WriteData(o.NSEC3CurrentSaltLength); err != nil {
		return err
	}
	if o.NSEC3CurrentSalt != nil || o.NSEC3CurrentSaltLength > 0 {
		_ptr_pbNSEC3CurrentSalt := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.NSEC3CurrentSaltLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.NSEC3CurrentSalt {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.NSEC3CurrentSalt[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.NSEC3CurrentSalt); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.NSEC3CurrentSalt, _ptr_pbNSEC3CurrentSalt); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.CurrentRollingSKDGUID != nil {
		if err := o.CurrentRollingSKDGUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.BufferLength); err != nil {
		return err
	}
	if o.Buffer != nil || o.BufferLength > 0 {
		_ptr_pBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.BufferLength)
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
		if err := w.WritePointer(&o.Buffer, _ptr_pBuffer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Count); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	for i1 := range o.ZoneSKDArray {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.ZoneSKDArray[i1] != nil {
			_ptr_pZoneSkdArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ZoneSKDArray[i1] != nil {
					if err := o.ZoneSKDArray[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ZoneSKD{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ZoneSKDArray[i1], _ptr_pZoneSkdArray); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.ZoneSKDArray); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ZoneDNSSECSettings) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsSigned); err != nil {
		return err
	}
	if err := w.ReadData(&o.SignWithNSEC3); err != nil {
		return err
	}
	if err := w.ReadData(&o.NSEC3OptOut); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaintainTrustAnchor); err != nil {
		return err
	}
	if err := w.ReadData(&o.ParentHasSecureDelegation); err != nil {
		return err
	}
	if err := w.ReadData(&o.DSRecordAlgorithms); err != nil {
		return err
	}
	if err := w.ReadData(&o.Rfc5011KeyRollovers); err != nil {
		return err
	}
	if err := w.ReadData(&o.NSEC3HashAlgorithm); err != nil {
		return err
	}
	if err := w.ReadData(&o.NSEC3RandomSaltLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.NSEC3IterationCount); err != nil {
		return err
	}
	_ptr_pwszNSEC3UserSalt := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.NSEC3UserSalt); err != nil {
			return err
		}
		return nil
	})
	_s_pwszNSEC3UserSalt := func(ptr interface{}) { o.NSEC3UserSalt = *ptr.(*string) }
	if err := w.ReadPointer(&o.NSEC3UserSalt, _s_pwszNSEC3UserSalt, _ptr_pwszNSEC3UserSalt); err != nil {
		return err
	}
	if err := w.ReadData(&o.DNSKEYRecordSetTTL); err != nil {
		return err
	}
	if err := w.ReadData(&o.DSRecordSetTTL); err != nil {
		return err
	}
	if err := w.ReadData(&o.SignatureInceptionOffset); err != nil {
		return err
	}
	if err := w.ReadData(&o.SecureDelegationPollingPeriod); err != nil {
		return err
	}
	if err := w.ReadData(&o.PropagationTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.NSEC3CurrentSaltLength); err != nil {
		return err
	}
	_ptr_pbNSEC3CurrentSalt := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.NSEC3CurrentSaltLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.NSEC3CurrentSaltLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.NSEC3CurrentSalt", sizeInfo[0])
		}
		o.NSEC3CurrentSalt = make([]byte, sizeInfo[0])
		for i1 := range o.NSEC3CurrentSalt {
			i1 := i1
			if err := w.ReadData(&o.NSEC3CurrentSalt[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pbNSEC3CurrentSalt := func(ptr interface{}) { o.NSEC3CurrentSalt = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.NSEC3CurrentSalt, _s_pbNSEC3CurrentSalt, _ptr_pbNSEC3CurrentSalt); err != nil {
		return err
	}
	if o.CurrentRollingSKDGUID == nil {
		o.CurrentRollingSKDGUID = &dtyp.GUID{}
	}
	if err := o.CurrentRollingSKDGUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.BufferLength); err != nil {
		return err
	}
	_ptr_pBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.BufferLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.BufferLength)
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
	_s_pBuffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Buffer, _s_pBuffer, _ptr_pBuffer); err != nil {
		return err
	}
	if err := w.ReadData(&o.Count); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.Count > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.Count)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.ZoneSKDArray", sizeInfo[0])
	}
	o.ZoneSKDArray = make([]*ZoneSKD, sizeInfo[0])
	for i1 := range o.ZoneSKDArray {
		i1 := i1
		_ptr_pZoneSkdArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ZoneSKDArray[i1] == nil {
				o.ZoneSKDArray[i1] = &ZoneSKD{}
			}
			if err := o.ZoneSKDArray[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pZoneSkdArray := func(ptr interface{}) { o.ZoneSKDArray[i1] = *ptr.(**ZoneSKD) }
		if err := w.ReadPointer(&o.ZoneSKDArray[i1], _s_pZoneSkdArray, _ptr_pZoneSkdArray); err != nil {
			return err
		}
	}
	return nil
}

// TrustAnchor structure represents DNS_RPC_TRUST_ANCHOR RPC structure.
//
// The DNS_RPC_TRUST_ANCHOR structure contains information about a trust anchor.
type TrustAnchor struct {
	// dwRpcStructureVersion: The structure version number; this MUST be set to 0x00000001.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// dwReserved0: MUST be set to zero when sent and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved0"`
	// wTrustAnchorType: The DNS record type corresponding to the trust anchor. This MUST
	// be set to one of the following values.
	//
	//	+------------------------+---------------------------------+
	//	|                        |                                 |
	//	|         VALUE          |             MEANING             |
	//	|                        |                                 |
	//	+------------------------+---------------------------------+
	//	+------------------------+---------------------------------+
	//	| DNS_TYPE_DS 0x002B     | A DS record type [RFC4034].     |
	//	+------------------------+---------------------------------+
	//	| DNS_TYPE_DNSKEY 0x0030 | A DNSKEY record type [RFC4034]. |
	//	+------------------------+---------------------------------+
	TrustAnchorType uint16 `idl:"name:wTrustAnchorType" json:"trust_anchor_type"`
	// wKeyTag: The DNSSEC key tag for this trust anchor. The key tag for a DS record trust
	// anchor MUST match the value of the record’s "Key Tag" field (see [RFC4034]). The
	// key tag for a DNSKEY record trust anchor MUST match the value calculated for the
	// DNSKEY record (see [RFC4034] Appendix B), with the exception that the REVOKE bit
	// of the DNSKEY flags field MUST be set to zero before the calculation.
	KeyTag uint16 `idl:"name:wKeyTag" json:"key_tag"`
	// wRRLength: The length of the RRData array.
	RRLength uint16 `idl:"name:wRRLength" json:"rr_length"`
	// eTrustAnchorState: The current state of the trust anchor. This MUST be one of the
	// following TRUSTANCHOR_STATE enumeration values (section 2.2.1.1.4).
	//
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	|                                        |                                                                                  |
	//	|                 VALUE                  |                                     MEANING                                      |
	//	|                                        |                                                                                  |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| TRUSTANCHOR_STATE_DSPENDING 0x00000001 | This trust anchor can be replaced with a matching DNSKEY trust anchor when the   |
	//	|                                        | associated trust point  has had a successful active refresh. If this state is    |
	//	|                                        | set, wTrustAnchorType MUST be DNS_TYPE_DS.                                       |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| TRUSTANCHOR_STATE_DSINVALID 0x00000002 | The associated trust point has had a successful active refresh, but no           |
	//	|                                        | DNSKEY record was found that matches this trust anchor. If this state is set,    |
	//	|                                        | wTrustAnchorType MUST be DNS_TYPE_DS.                                            |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| TRUSTANCHOR_STATE_ADDPEND 0x00000003   | This trust anchor will become valid after the expiration of the RFC 5011 add     |
	//	|                                        | hold-down time (see [RFC5011]). This corresponds to the "AddPend" state in RFC   |
	//	|                                        | 5011.                                                                            |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| TRUSTANCHOR_STATE_VALID 0x00000004     | This trust anchor is trusted for DNSSEC proofs because it was either explicitly  |
	//	|                                        | added by the administrator or became valid after the expiration of the RFC 5011  |
	//	|                                        | add hold-down time (see [RFC5011]). This corresponds to the Valid state in RFC   |
	//	|                                        | 5011.                                                                            |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| TRUSTANCHOR_STATE_MISSING 0x00000005   | This trust anchor was in the TRUSTANCHOR_STATE_VALID state but was missing in    |
	//	|                                        | the last successful active refresh. It is still trusted for DNSSEC proofs. This  |
	//	|                                        | corresponds to the Valid state in [RFC5011].                                     |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| TRUSTANCHOR_STATE_REVOKED 0x00000006   | This trust anchor has been marked as revoked by the administrator for the trust  |
	//	|                                        | point's zone. It can never again be trusted for DNSSEC proofs. This corresponds  |
	//	|                                        | to the Revoked state in [RFC5011].                                               |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	TrustAnchorState TrustAnchorState `idl:"name:eTrustAnchorState" json:"trust_anchor_state"`
	// i64EnteredStateTime: The time at which this trust anchor entered its current state.
	// This is a 64-bit value representing the number of 100-nanosecond intervals since
	// January 1, 1601 (UTC).
	EnteredStateTime int64 `idl:"name:i64EnteredStateTime" json:"entered_state_time"`
	// i64NextStateTime: The time at which this trust anchor is scheduled to exit the current
	// state. This is a 64-bit value representing the number of 100-nanosecond intervals
	// since January 1, 1601 (UTC). The meaning is dependent on the value of eTrustAnchorState.
	//
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	|                VALUE OF                |                                                                                  |
	//	|           ETRUSTANCHORSTATE            |                                     MEANING                                      |
	//	|                                        |                                                                                  |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| TRUSTANCHOR_STATE_DSPENDING 0x00000001 | Reserved. The value of i64NextStateTime MUST be set to zero when sent and MUST   |
	//	|                                        | be ignored on receipt.                                                           |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| TRUSTANCHOR_STATE_DSINVALID 0x00000002 | Reserved. The value of i64NextStateTime MUST be set to zero when sent and MUST   |
	//	|                                        | be ignored on receipt.                                                           |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| TRUSTANCHOR_STATE_ADDPEND 0x00000003   | This trust anchor is scheduled to enter the TRUSTANCHOR_STATE_VALID state on     |
	//	|                                        | or after the value of i64NextStateTime. This MUST be equivalent to the value     |
	//	|                                        | of i64EnteredStateTime added to the value of the add hold-down time (see         |
	//	|                                        | [RFC5011]).                                                                      |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| TRUSTANCHOR_STATE_VALID 0x00000004     | Reserved. The value of i64NextStateTime MUST be set to zero when sent and MUST   |
	//	|                                        | be ignored on receipt.                                                           |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| TRUSTANCHOR_STATE_MISSING 0x00000005   | Reserved. The value of i64NextStateTime MUST be set to zero when sent and MUST   |
	//	|                                        | be ignored on receipt.                                                           |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| TRUSTANCHOR_STATE_REVOKED 0x00000006   | This trust anchor will become eligible for deletion on or after the value of     |
	//	|                                        | i64NextStateTime. This MUST be equivalent to the value of i64EnteredStateTime    |
	//	|                                        | added to the value of the remove hold-down time (see [RFC5011]).                 |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	NextStateTime int64 `idl:"name:i64NextStateTime" json:"next_state_time"`
	// dwReserved: MUST be set to zero when sent and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved"`
	// RRData: Binary data in the same format as DNS_RPC_RECORD_DNSKEY (section 2.2.2.2.4.15)
	// if wTrustAnchorType is DNS_TYPE_DNSKEY, or binary data in the same format as DNS_RPC_RECORD_DS
	// (section 2.2.2.2.4.12) if wTrustAnchorType is DNS_TYPE_DS.
	RRData []byte `idl:"name:RRData;size_is:(wRRLength)" json:"rr_data"`
}

func (o *TrustAnchor) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.RRData != nil && o.RRLength == 0 {
		o.RRLength = uint16(len(o.RRData))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *TrustAnchor) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.RRLength)
	return []uint64{
		dimSize1,
	}
}
func (o *TrustAnchor) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.TrustAnchorType); err != nil {
		return err
	}
	if err := w.WriteData(o.KeyTag); err != nil {
		return err
	}
	if err := w.WriteData(o.RRLength); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.TrustAnchorState)); err != nil {
		return err
	}
	if err := w.WriteData(o.EnteredStateTime); err != nil {
		return err
	}
	if err := w.WriteData(o.NextStateTime); err != nil {
		return err
	}
	// reserved dwReserved
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	for i1 := range o.RRData {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.RRData[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.RRData); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustAnchor) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	if err := w.ReadData(&o.TrustAnchorType); err != nil {
		return err
	}
	if err := w.ReadData(&o.KeyTag); err != nil {
		return err
	}
	if err := w.ReadData(&o.RRLength); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.TrustAnchorState)); err != nil {
		return err
	}
	if err := w.ReadData(&o.EnteredStateTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.NextStateTime); err != nil {
		return err
	}
	// reserved dwReserved
	var _dwReserved uint32
	if err := w.ReadData(&_dwReserved); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.RRLength > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.RRLength)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.RRData", sizeInfo[0])
	}
	o.RRData = make([]byte, sizeInfo[0])
	for i1 := range o.RRData {
		i1 := i1
		if err := w.ReadData(&o.RRData[i1]); err != nil {
			return err
		}
	}
	return nil
}

// TrustAnchorList structure represents DNS_RPC_TRUST_ANCHOR_LIST RPC structure.
//
// The DNS_RPC_TRUST_ANCHOR_LIST structure contains zero or more DNS_RPC_TRUST_ANCHOR
// structures.
type TrustAnchorList struct {
	// dwRpcStructureVersion: The structure version number; this MUST be set to 0x00000001.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// dwReserved0: MUST be set to zero when sent and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved0"`
	// dwTrustAnchorCount: The size of the TrustAnchorArray array.
	TrustAnchorCount uint32 `idl:"name:dwTrustAnchorCount" json:"trust_anchor_count"`
	// TrustAnchorArray: An array of pointers to DNS_RPC_TRUST_ANCHOR structures (section
	// 2.2.6.2.6).
	TrustAnchorArray []*TrustAnchor `idl:"name:TrustAnchorArray;size_is:(dwTrustAnchorCount)" json:"trust_anchor_array"`
}

func (o *TrustAnchorList) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.TrustAnchorArray != nil && o.TrustAnchorCount == 0 {
		o.TrustAnchorCount = uint32(len(o.TrustAnchorArray))
	}
	if o.TrustAnchorCount > uint32(500000) {
		return fmt.Errorf("TrustAnchorCount is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *TrustAnchorList) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.TrustAnchorCount)
	return []uint64{
		dimSize1,
	}
}
func (o *TrustAnchorList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.TrustAnchorCount); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	for i1 := range o.TrustAnchorArray {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.TrustAnchorArray[i1] != nil {
			_ptr_TrustAnchorArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.TrustAnchorArray[i1] != nil {
					if err := o.TrustAnchorArray[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&TrustAnchor{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.TrustAnchorArray[i1], _ptr_TrustAnchorArray); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.TrustAnchorArray); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustAnchorList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	if err := w.ReadData(&o.TrustAnchorCount); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.TrustAnchorCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.TrustAnchorCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.TrustAnchorArray", sizeInfo[0])
	}
	o.TrustAnchorArray = make([]*TrustAnchor, sizeInfo[0])
	for i1 := range o.TrustAnchorArray {
		i1 := i1
		_ptr_TrustAnchorArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.TrustAnchorArray[i1] == nil {
				o.TrustAnchorArray[i1] = &TrustAnchor{}
			}
			if err := o.TrustAnchorArray[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_TrustAnchorArray := func(ptr interface{}) { o.TrustAnchorArray[i1] = *ptr.(**TrustAnchor) }
		if err := w.ReadPointer(&o.TrustAnchorArray[i1], _s_TrustAnchorArray, _ptr_TrustAnchorArray); err != nil {
			return err
		}
	}
	return nil
}

// DPEnum structure represents DNS_RPC_DP_ENUM RPC structure.
//
// The DNS_RPC_DP_ENUM structure contains abbreviated information about an application
// directory partition.<58>
type DPEnum struct {
	// dwRpcStructureVersion: As specified in section 2.2.7.2.1.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// dwReserved0: As specified in section 2.2.7.2.1.
	_ uint32 `idl:"name:dwReserved0"`
	// pszDpFqdn: As specified in section 2.2.7.2.1.
	DPFQDN string `idl:"name:pszDpFqdn;string" json:"dp_fqdn"`
	// dwFlags: As specified in section 2.2.7.2.1.
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
	// dwZoneCount: As specified in section 2.2.7.2.1.
	ZoneCount uint32 `idl:"name:dwZoneCount" json:"zone_count"`
}

func (o *DPEnum) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *DPEnum) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if o.DPFQDN != "" {
		_ptr_pszDpFqdn := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.DPFQDN); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DPFQDN, _ptr_pszDpFqdn); err != nil {
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
	if err := w.WriteData(o.ZoneCount); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *DPEnum) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	_ptr_pszDpFqdn := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.DPFQDN); err != nil {
			return err
		}
		return nil
	})
	_s_pszDpFqdn := func(ptr interface{}) { o.DPFQDN = *ptr.(*string) }
	if err := w.ReadPointer(&o.DPFQDN, _s_pszDpFqdn, _ptr_pszDpFqdn); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.ZoneCount); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// DPList structure represents DNS_RPC_DP_LIST RPC structure.
//
// The DNS_RPC_DP_LIST structure contains a list of application directory partition
// information structures.<59>
type DPList struct {
	// dwRpcStructureVersion: As specified in section 2.2.7.2.1.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// dwReserved0: As specified in section 2.2.7.2.1.
	_ uint32 `idl:"name:dwReserved0"`
	// dwDpCount: The number of DNS_RPC_DP_ENUM (section 2.2.7.2.3) structures present in
	// the array pointed to by DpArray.
	DPCount uint32 `idl:"name:dwDpCount" json:"dp_count"`
	// DpArray: An array of DNS_RPC_DP_ENUM structures (section 2.2.7.2.3), containing information
	// about the application directory partitions available to the DNS server.
	DPArray []*DPEnum `idl:"name:DpArray;size_is:(dwDpCount)" json:"dp_array"`
}

func (o *DPList) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.DPArray != nil && o.DPCount == 0 {
		o.DPCount = uint32(len(o.DPArray))
	}
	if o.DPCount > uint32(5000) {
		return fmt.Errorf("DPCount is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *DPList) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DPCount)
	return []uint64{
		dimSize1,
	}
}
func (o *DPList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.DPCount); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	for i1 := range o.DPArray {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.DPArray[i1] != nil {
			_ptr_DpArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DPArray[i1] != nil {
					if err := o.DPArray[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&DPEnum{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DPArray[i1], _ptr_DpArray); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.DPArray); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *DPList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	if err := w.ReadData(&o.DPCount); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DPCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DPCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.DPArray", sizeInfo[0])
	}
	o.DPArray = make([]*DPEnum, sizeInfo[0])
	for i1 := range o.DPArray {
		i1 := i1
		_ptr_DpArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DPArray[i1] == nil {
				o.DPArray[i1] = &DPEnum{}
			}
			if err := o.DPArray[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_DpArray := func(ptr interface{}) { o.DPArray[i1] = *ptr.(**DPEnum) }
		if err := w.ReadPointer(&o.DPArray[i1], _s_DpArray, _ptr_DpArray); err != nil {
			return err
		}
	}
	return nil
}

// DPReplica structure represents DNS_RPC_DP_REPLICA RPC structure.
//
// The DNS_RPC_DP_REPLICA structure contains information about an application directory
// partition replica by giving a distinguished name which can be used to uniquely identify
// the replica.<57>
type DPReplica struct {
	// pszReplicaDn: A pointer to a null-terminated Unicode string that specifies the distinguished
	// name that identifies a specific directory server.
	ReplicaDN string `idl:"name:pszReplicaDn;string" json:"replica_dn"`
}

func (o *DPReplica) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *DPReplica) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.ReplicaDN != "" {
		_ptr_pszReplicaDn := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ReplicaDN); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ReplicaDN, _ptr_pszReplicaDn); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *DPReplica) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_pszReplicaDn := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ReplicaDN); err != nil {
			return err
		}
		return nil
	})
	_s_pszReplicaDn := func(ptr interface{}) { o.ReplicaDN = *ptr.(*string) }
	if err := w.ReadPointer(&o.ReplicaDN, _s_pszReplicaDn, _ptr_pszReplicaDn); err != nil {
		return err
	}
	return nil
}

// DPInfo structure represents DNS_RPC_DP_INFO RPC structure.
//
// The DNS_RPC_DP_INFO structure SHOULD<55> represent the current state of an application
// directory partition on the directory server.
type DPInfo struct {
	// dwRpcStructureVersion: The structure version number; this value MUST be set to 0x00000000.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// dwReserved0: MUST be set to zero when sent and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved0"`
	// pszDpFqdn: A pointer to a null-terminated UTF-8 string that specifies the FQDN of
	// the application directory partition. This value is read from the dnsRoot attribute
	// of the partition crossRef object (see pszCrDn) converted to UTF-8 for this application
	// directory partition.
	DPFQDN string `idl:"name:pszDpFqdn;string" json:"dp_fqdn"`
	// pszDpDn: A pointer to a null-terminated Unicode string that specifies the distinguished
	// name for the application directory partition naming context root object. This is
	// the value of the nCName attribute of the application directory partition crossRef
	// object (see pszCrDn).
	DPDN string `idl:"name:pszDpDn;string" json:"dp_dn"`
	// pszCrDn: A pointer to a null-terminated Unicode string that specifies the distinguished
	// name for the application directory partition crossRef object (located beneath "CN=Partitions,
	// CN=Configuration, <Forest DN>").
	CRDN string `idl:"name:pszCrDn;string" json:"cr_dn"`
	// dwFlags: The application directory partition properties; this MUST be set to a combination
	// of allowed values for DNS_RPC_DP_FLAGS (section 2.2.7.1.1).
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
	// dwZoneCount: The number of zones from this application directory partition that are
	// loaded in the DNS server's memory. This value is incremented or decremented in the
	// Application Directory Partition Table whenever a DNS Zone Table entry corresponding
	// to a zone in this application directory partition is initialized or deleted, respectively.
	ZoneCount uint32 `idl:"name:dwZoneCount" json:"zone_count"`
	// dwState: The current state of this application directory partition. This MUST be
	// set to one of the following values:
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                       |                                                                                  |                                                                                  |
	//	|                 VALUE                 |                                     MEANING                                      |                                      SOURCE                                      |
	//	|                                       |                                                                                  |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_DP_OKAY 0x00000000                | The application directory partition is running and ready for all operations.     | The Application Directory Partition naming context root object's instanceType    |
	//	|                                       |                                                                                  | attribute ([MS-ADA1] section 2.309) has neither DS_INSTANCETYPE_NC_COMING        |
	//	|                                       |                                                                                  | (0x00000010), nor the DS_INSTANCETYPE_NC_GOING ( 0x00000020) bit set.            |
	//	+---------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_DP_STATE_REPL_INCOMING 0x00000001 | The application directory partition is replicating onto the directory server     | The Application Directory Partition naming context root object's instanceType    |
	//	|                                       | but has not completed an initial synchronization so will be ignored for the time | attribute has the (DS_INSTANCETYPE_NC_COMING ( 0x00000010) bit set.              |
	//	|                                       | being.                                                                           |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_DP_STATE_REPL_OUTGOING 0x00000002 | The application directory partition is being deleted from the directory server   | The Application Directory Partition naming context root object's instanceType    |
	//	|                                       | and so will be ignored.                                                          | attribute has the DS_INSTANCETYPE_NC_GOING ( 0x00000020) bit set.                |
	//	+---------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_DP_STATE_UNKNOWN 0x00000003       | The application directory partition state is unavailable for unknown reasons.    | The Application Directory Partition naming context root object's instanceType    |
	//	|                                       |                                                                                  | attribute is unavailable due to an error condition.                              |
	//	+---------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	State uint32 `idl:"name:dwState" json:"state"`
	// dwReserved: MUST be set to zero when sent and MUST be ignored on receipt.
	_ []uint32 `idl:"name:dwReserved"`
	// pwszReserved: MUST be set to zero when sent and MUST be ignored on receipt.
	_ string `idl:"name:pwszReserved;string"`
	// dwReplicaCount: The number of replication locations for the application directory
	// partition. This value MUST be between 0 and 10000. This value is calculated from
	// the values of the msDS-NC-Replica-Locations and msDS-NC-RO-Replica-Locations<56>attributes
	// of the application directory partition crossRef object (see pszCrDn), as the sum
	// of the number of DNs listed in each attribute.
	ReplicaCount uint32 `idl:"name:dwReplicaCount" json:"replica_count"`
	// ReplicaArray: Array of DNS_RPC_DP_REPLICA (section 2.2.7.2.2), that contains information
	// about replication locations for this application directory partition. This structure
	// is populated from the values of the msDS-NC-Replica-Locations (section 2.382) and
	// msDS-NC-RO-Replica-Locations (section 2.383) attributes of the application directory
	// partition crossRef object (see pszCrDn). Failure to read any of those attributes
	// will be treated as if no replica exists for that attribute.
	ArrayReplica []*DPReplica `idl:"name:ReplicaArray;size_is:(dwReplicaCount)" json:"array_replica"`
}

func (o *DPInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.ArrayReplica != nil && o.ReplicaCount == 0 {
		o.ReplicaCount = uint32(len(o.ArrayReplica))
	}
	if o.ReplicaCount > uint32(10000) {
		return fmt.Errorf("ReplicaCount is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *DPInfo) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.ReplicaCount)
	return []uint64{
		dimSize1,
	}
}
func (o *DPInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if o.DPFQDN != "" {
		_ptr_pszDpFqdn := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.DPFQDN); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DPFQDN, _ptr_pszDpFqdn); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DPDN != "" {
		_ptr_pszDpDn := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DPDN); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DPDN, _ptr_pszDpDn); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.CRDN != "" {
		_ptr_pszCrDn := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.CRDN); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.CRDN, _ptr_pszCrDn); err != nil {
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
	if err := w.WriteData(o.ZoneCount); err != nil {
		return err
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	// reserved dwReserved
	for i1 := 0; uint64(i1) < 3; i1++ {
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	// reserved pwszReserved
	if err := w.WritePointer(nil); err != nil {
		return err
	}
	if err := w.WriteData(o.ReplicaCount); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	for i1 := range o.ArrayReplica {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.ArrayReplica[i1] != nil {
			_ptr_ReplicaArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ArrayReplica[i1] != nil {
					if err := o.ArrayReplica[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&DPReplica{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ArrayReplica[i1], _ptr_ReplicaArray); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.ArrayReplica); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *DPInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	_ptr_pszDpFqdn := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.DPFQDN); err != nil {
			return err
		}
		return nil
	})
	_s_pszDpFqdn := func(ptr interface{}) { o.DPFQDN = *ptr.(*string) }
	if err := w.ReadPointer(&o.DPFQDN, _s_pszDpFqdn, _ptr_pszDpFqdn); err != nil {
		return err
	}
	_ptr_pszDpDn := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DPDN); err != nil {
			return err
		}
		return nil
	})
	_s_pszDpDn := func(ptr interface{}) { o.DPDN = *ptr.(*string) }
	if err := w.ReadPointer(&o.DPDN, _s_pszDpDn, _ptr_pszDpDn); err != nil {
		return err
	}
	_ptr_pszCrDn := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.CRDN); err != nil {
			return err
		}
		return nil
	})
	_s_pszCrDn := func(ptr interface{}) { o.CRDN = *ptr.(*string) }
	if err := w.ReadPointer(&o.CRDN, _s_pszCrDn, _ptr_pszCrDn); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.ZoneCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	// reserved dwReserved
	var _dwReserved []uint32
	_dwReserved = make([]uint32, 3)
	for i1 := range _dwReserved {
		i1 := i1
		if err := w.ReadData(&_dwReserved[i1]); err != nil {
			return err
		}
	}
	// reserved pwszReserved
	var _pwszReserved string
	_ptr_pwszReserved := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &_pwszReserved); err != nil {
			return err
		}
		return nil
	})
	_s_pwszReserved := func(ptr interface{}) { _pwszReserved = *ptr.(*string) }
	if err := w.ReadPointer(&_pwszReserved, _s_pwszReserved, _ptr_pwszReserved); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReplicaCount); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.ReplicaCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.ReplicaCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.ArrayReplica", sizeInfo[0])
	}
	o.ArrayReplica = make([]*DPReplica, sizeInfo[0])
	for i1 := range o.ArrayReplica {
		i1 := i1
		_ptr_ReplicaArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ArrayReplica[i1] == nil {
				o.ArrayReplica[i1] = &DPReplica{}
			}
			if err := o.ArrayReplica[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReplicaArray := func(ptr interface{}) { o.ArrayReplica[i1] = *ptr.(**DPReplica) }
		if err := w.ReadPointer(&o.ArrayReplica[i1], _s_ReplicaArray, _ptr_ReplicaArray); err != nil {
			return err
		}
	}
	return nil
}

// EnlistDP structure represents DNS_RPC_ENLIST_DP RPC structure.
//
// The DNS_RPC_ENLIST_DP structure contains the information required to create, delete,
// or enumerate application directory partitions.<60>
type EnlistDP struct {
	// dwRpcStructureVersion: The DNS management structure version number; this value MUST
	// be set to 0x00000001.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// dwReserved0: As specified in section 2.2.7.2.1.
	_ uint32 `idl:"name:dwReserved0"`
	// pszDpFqdn: As specified in section 2.2.7.2.1.
	DPFQDN string `idl:"name:pszDpFqdn;string" json:"dp_fqdn"`
	// dwOperation: The application directory partition operation to be performed by the
	// DNS server; this MUST be set to one of the following values:
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|                                    |                                                                                  |
	//	|               VALUE                |                                     MEANING                                      |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_DP_OP_CREATE 0x00000001        | Create and enlist (DNS_DP_OP_ENLIST) a new application directory partition.      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_DP_OP_DELETE 0x00000002        | Delete an existing application directory partition. If the application           |
	//	|                                    | directory partition has been marked DNS_DP_AUTOCREATED, DNS_DP_LEGACY,           |
	//	|                                    | DNS_DP_DOMAIN_DEFAULT, DNS_DP_FOREST_DEFAULT, or DNS_DP_DELETED, as specified in |
	//	|                                    | section 2.2.7.1.1, or if the DNS server cannot connect and bind to the FSMO role |
	//	|                                    | owner, then the server MUST return an error.                                     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_DP_OP_ENLIST 0x00000003        | Enlist this DNS server in an existing application directory partition.           |
	//	|                                    | If the application directory partition has been marked DNS_DP_ENLISTED or        |
	//	|                                    | DNS_DP_DELETED, as specified in section 2.2.7.1.1, then the DNS server MUST      |
	//	|                                    | return an error.                                                                 |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_DP_OP_UNENLIST 0x00000004      | Un-enlist this DNS server from an existing application directory partition. If   |
	//	|                                    | the application directory partition has been marked DNS_DP_DELETED, as specified |
	//	|                                    | in section 2.2.7.1.1, then the DNS server MUST return an error.                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_DP_OP_CREATE_DOMAIN 0x00000005 | Create a domain partition on the directory server if one does not already exist. |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_DP_OP_CREATE_FOREST 0x00000006 | Create a forest partition on the directory server if it does not already exist.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	Operation uint32 `idl:"name:dwOperation" json:"operation"`
}

func (o *EnlistDP) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *EnlistDP) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if o.DPFQDN != "" {
		_ptr_pszDpFqdn := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.DPFQDN); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DPFQDN, _ptr_pszDpFqdn); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Operation); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *EnlistDP) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	_ptr_pszDpFqdn := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.DPFQDN); err != nil {
			return err
		}
		return nil
	})
	_s_pszDpFqdn := func(ptr interface{}) { o.DPFQDN = *ptr.(*string) }
	if err := w.ReadPointer(&o.DPFQDN, _s_pszDpFqdn, _ptr_pszDpFqdn); err != nil {
		return err
	}
	if err := w.ReadData(&o.Operation); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// ZoneExportInfo structure represents DNS_RPC_ZONE_EXPORT_INFO RPC structure.
//
// The DNS_RPC_ZONE_EXPORT_INFO structure contains the information file to which a zone
// is exported on the DNS server.<44>
type ZoneExportInfo struct {
	// dwRpcStructureVersion: The structure version number; this MUST be set to 0x00000001.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// dwReserved0: MUST be set to zero when sent and MUST be ignored on receipt
	_ uint32 `idl:"name:dwReserved0"`
	// pszZoneExportFile: A pointer to a null-terminated UTF-8 string that specifies the
	// name of the file to which a zone is exported by the DNS server.
	ZoneExportFile string `idl:"name:pszZoneExportFile;string" json:"zone_export_file"`
}

func (o *ZoneExportInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ZoneExportInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if o.ZoneExportFile != "" {
		_ptr_pszZoneExportFile := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.ZoneExportFile); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ZoneExportFile, _ptr_pszZoneExportFile); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ZoneExportInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	_ptr_pszZoneExportFile := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.ZoneExportFile); err != nil {
			return err
		}
		return nil
	})
	_s_pszZoneExportFile := func(ptr interface{}) { o.ZoneExportFile = *ptr.(*string) }
	if err := w.ReadPointer(&o.ZoneExportFile, _s_pszZoneExportFile, _ptr_pszZoneExportFile); err != nil {
		return err
	}
	return nil
}

// ZoneSecondariesW2K structure represents DNS_RPC_ZONE_SECONDARIES_W2K RPC structure.
//
// The DNS_RPC_ZONE_SECONDARIES_W2K structure is used to specify information about the
// secondary servers for a primary DNS zone.
type ZoneSecondariesW2K struct {
	// fSecureSecondaries: The secondary security settings configured for this zone. The
	// DNS server MUST respond to zone transfer requests from a secondary server according
	// to the behavior corresponding to the value of the flag, as described in DNS_ZONE_SECONDARY_SECURITY
	// (section 2.2.5.1.2). This value MUST be set to one of the allowed values as specified
	// in DNS_ZONE_SECONDARY_SECURITY (section 2.2.5.1.2).
	SecureSecondaries uint32 `idl:"name:fSecureSecondaries" json:"secure_secondaries"`
	// fNotifyLevel: The settings for sending zone notifications for this zone. The DNS
	// server MUST send notify messages to secondary servers as specified by the entry corresponding
	// to the value of this flag, as shown in the table in section 2.2.5.1.3. This flag's
	// value MUST be set to one of the allowed values as specified in DNS_ZONE_ NOTIFY_LEVEL
	// (section 2.2.5.1.3).
	NotifyLevel uint32 `idl:"name:fNotifyLevel" json:"notify_level"`
	// aipSecondaries: The list of IPv4 addresses of remote DNS servers that are permitted
	// to perform zone transfers for this zone. The DNS server will honor zone transfer
	// requests from these secondary servers, as specified by fSecureSecondaries above.
	Secondaries *IPv4Array `idl:"name:aipSecondaries" json:"secondaries"`
	// aipNotify: The list of IPv4 addresses of the remote DNS servers that will be sent
	// notification messages when records in this zone change, as directed by fNotifyLevel
	// above.
	Notify *IPv4Array `idl:"name:aipNotify" json:"notify"`
}

func (o *ZoneSecondariesW2K) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ZoneSecondariesW2K) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.SecureSecondaries); err != nil {
		return err
	}
	if err := w.WriteData(o.NotifyLevel); err != nil {
		return err
	}
	if o.Secondaries != nil {
		_ptr_aipSecondaries := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Secondaries != nil {
				if err := o.Secondaries.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPv4Array{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Secondaries, _ptr_aipSecondaries); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Notify != nil {
		_ptr_aipNotify := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Notify != nil {
				if err := o.Notify.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPv4Array{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Notify, _ptr_aipNotify); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ZoneSecondariesW2K) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.SecureSecondaries); err != nil {
		return err
	}
	if err := w.ReadData(&o.NotifyLevel); err != nil {
		return err
	}
	_ptr_aipSecondaries := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Secondaries == nil {
			o.Secondaries = &IPv4Array{}
		}
		if err := o.Secondaries.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipSecondaries := func(ptr interface{}) { o.Secondaries = *ptr.(**IPv4Array) }
	if err := w.ReadPointer(&o.Secondaries, _s_aipSecondaries, _ptr_aipSecondaries); err != nil {
		return err
	}
	_ptr_aipNotify := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Notify == nil {
			o.Notify = &IPv4Array{}
		}
		if err := o.Notify.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipNotify := func(ptr interface{}) { o.Notify = *ptr.(**IPv4Array) }
	if err := w.ReadPointer(&o.Notify, _s_aipNotify, _ptr_aipNotify); err != nil {
		return err
	}
	return nil
}

// ZoneSecondaries_NET structure represents DNS_RPC_ZONE_SECONDARIES_DOTNET RPC structure.
//
// In the DNS_RPC_ZONE_SECONDARIES_DOTNET structure all fields have same definition
// as specified in section DNS_RPC_ZONE_SECONDARIES_W2K (section 2.2.5.2.5.1), with
// the following exceptions:
type ZoneSecondaries_NET struct {
	// dwRpcStructureVersion: The DNS server management structure version number. It MUST
	// be set to 0x00000001
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// dwReserved0: Reserved for future use. Sender MUST set to zero and receiver MUST ignore
	// this value.
	_                 uint32     `idl:"name:dwReserved0"`
	SecureSecondaries uint32     `idl:"name:fSecureSecondaries" json:"secure_secondaries"`
	NotifyLevel       uint32     `idl:"name:fNotifyLevel" json:"notify_level"`
	Secondaries       *IPv4Array `idl:"name:aipSecondaries" json:"secondaries"`
	Notify            *IPv4Array `idl:"name:aipNotify" json:"notify"`
}

func (o *ZoneSecondaries_NET) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ZoneSecondaries_NET) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.SecureSecondaries); err != nil {
		return err
	}
	if err := w.WriteData(o.NotifyLevel); err != nil {
		return err
	}
	if o.Secondaries != nil {
		_ptr_aipSecondaries := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Secondaries != nil {
				if err := o.Secondaries.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPv4Array{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Secondaries, _ptr_aipSecondaries); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Notify != nil {
		_ptr_aipNotify := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Notify != nil {
				if err := o.Notify.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPv4Array{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Notify, _ptr_aipNotify); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ZoneSecondaries_NET) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	if err := w.ReadData(&o.SecureSecondaries); err != nil {
		return err
	}
	if err := w.ReadData(&o.NotifyLevel); err != nil {
		return err
	}
	_ptr_aipSecondaries := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Secondaries == nil {
			o.Secondaries = &IPv4Array{}
		}
		if err := o.Secondaries.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipSecondaries := func(ptr interface{}) { o.Secondaries = *ptr.(**IPv4Array) }
	if err := w.ReadPointer(&o.Secondaries, _s_aipSecondaries, _ptr_aipSecondaries); err != nil {
		return err
	}
	_ptr_aipNotify := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Notify == nil {
			o.Notify = &IPv4Array{}
		}
		if err := o.Notify.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipNotify := func(ptr interface{}) { o.Notify = *ptr.(**IPv4Array) }
	if err := w.ReadPointer(&o.Notify, _s_aipNotify, _ptr_aipNotify); err != nil {
		return err
	}
	return nil
}

// ZoneSecondariesLonghorn structure represents DNS_RPC_ZONE_SECONDARIES_LONGHORN RPC structure.
//
// In the DNS_RPC_ZONE_SECONDARIES_LONGHORN structure all fields have same definition
// as specified in section DNS_RPC_ZONE_SECONDARIES_DOTNET (section 2.2.5.2.5.2), with
// the following exceptions:
type ZoneSecondariesLonghorn struct {
	// dwRpcStructureVersion: The DNS server management structure version number. It MUST
	// be set to 0x00000002.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	_                   uint32 `idl:"name:dwReserved0"`
	SecureSecondaries   uint32 `idl:"name:fSecureSecondaries" json:"secure_secondaries"`
	NotifyLevel         uint32 `idl:"name:fNotifyLevel" json:"notify_level"`
	// aipSecondaries: A pointer to a structure of type DNS_ADDR_ARRAY (section 2.2.3.2.3)
	// that specifies a list of IP addresses of remote DNS servers that are permitted to
	// perform zone transfers for this zone. The DNS server will honor zone transfer requests
	// from these secondary servers, as directed by the value of fSecureSecondaries above.
	Secondaries *AddrArray `idl:"name:aipSecondaries" json:"secondaries"`
	// aipNotify: A pointer to a structure of type DNS_ADDR_ARRAY (section 2.2.3.2.3) that
	// specifies a list of IP addresses of the remote DNS servers that be sent notification
	// messages when records in this zone change, as directed by fNotifyLevel above.
	Notify *AddrArray `idl:"name:aipNotify" json:"notify"`
}

func (o *ZoneSecondariesLonghorn) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ZoneSecondariesLonghorn) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.SecureSecondaries); err != nil {
		return err
	}
	if err := w.WriteData(o.NotifyLevel); err != nil {
		return err
	}
	if o.Secondaries != nil {
		_ptr_aipSecondaries := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Secondaries != nil {
				if err := o.Secondaries.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AddrArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Secondaries, _ptr_aipSecondaries); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Notify != nil {
		_ptr_aipNotify := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Notify != nil {
				if err := o.Notify.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AddrArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Notify, _ptr_aipNotify); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ZoneSecondariesLonghorn) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	if err := w.ReadData(&o.SecureSecondaries); err != nil {
		return err
	}
	if err := w.ReadData(&o.NotifyLevel); err != nil {
		return err
	}
	_ptr_aipSecondaries := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Secondaries == nil {
			o.Secondaries = &AddrArray{}
		}
		if err := o.Secondaries.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipSecondaries := func(ptr interface{}) { o.Secondaries = *ptr.(**AddrArray) }
	if err := w.ReadPointer(&o.Secondaries, _s_aipSecondaries, _ptr_aipSecondaries); err != nil {
		return err
	}
	_ptr_aipNotify := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Notify == nil {
			o.Notify = &AddrArray{}
		}
		if err := o.Notify.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipNotify := func(ptr interface{}) { o.Notify = *ptr.(**AddrArray) }
	if err := w.ReadPointer(&o.Notify, _s_aipNotify, _ptr_aipNotify); err != nil {
		return err
	}
	return nil
}

// ZoneSecondaries structure represents DNS_RPC_ZONE_SECONDARIES RPC structure.
//
// The DNS_RPC_ZONE_SECONDARIES structure contains the information about the secondary
// DNS servers for a zone. There are different versions of the DNS_RPC_ZONE_SECONDARIES
// structure. The DNS server MUST use the structure corresponding to the value of the
// dwClientVersion in DNS Server Management Protocol method calls (section 3.1.4) as
// in the following table, or if the method call does not specify the value of dwClientVersion,
// the DNS_RPC_ZONE_SECONDARIES_W2K version of the structure MUST be used.
//
//	+------------+----------------------------------------------------------+
//	|            |                                                          |
//	|   VALUE    |                        STRUCTURE                         |
//	|            |                                                          |
//	+------------+----------------------------------------------------------+
//	+------------+----------------------------------------------------------+
//	| 0x00000000 | DNS_RPC_ZONE_SECONDARIES_W2K (section 2.2.5.2.5.1)       |
//	+------------+----------------------------------------------------------+
//	| 0x00060000 | DNS_RPC_ ZONE_SECONDARIES_DOTNET (section 2.2.5.2.5.2)   |
//	+------------+----------------------------------------------------------+
//	| 0x00070000 | DNS_RPC_ ZONE_SECONDARIES_LONGHORN (section 2.2.5.2.5.3) |
//	+------------+----------------------------------------------------------+
type ZoneSecondaries struct {
	RPCStructureVersion uint32     `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	_                   uint32     `idl:"name:dwReserved0"`
	SecureSecondaries   uint32     `idl:"name:fSecureSecondaries" json:"secure_secondaries"`
	NotifyLevel         uint32     `idl:"name:fNotifyLevel" json:"notify_level"`
	Secondaries         *AddrArray `idl:"name:aipSecondaries" json:"secondaries"`
	Notify              *AddrArray `idl:"name:aipNotify" json:"notify"`
}

func (o *ZoneSecondaries) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ZoneSecondaries) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.SecureSecondaries); err != nil {
		return err
	}
	if err := w.WriteData(o.NotifyLevel); err != nil {
		return err
	}
	if o.Secondaries != nil {
		_ptr_aipSecondaries := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Secondaries != nil {
				if err := o.Secondaries.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AddrArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Secondaries, _ptr_aipSecondaries); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Notify != nil {
		_ptr_aipNotify := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Notify != nil {
				if err := o.Notify.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AddrArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Notify, _ptr_aipNotify); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ZoneSecondaries) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	if err := w.ReadData(&o.SecureSecondaries); err != nil {
		return err
	}
	if err := w.ReadData(&o.NotifyLevel); err != nil {
		return err
	}
	_ptr_aipSecondaries := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Secondaries == nil {
			o.Secondaries = &AddrArray{}
		}
		if err := o.Secondaries.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipSecondaries := func(ptr interface{}) { o.Secondaries = *ptr.(**AddrArray) }
	if err := w.ReadPointer(&o.Secondaries, _s_aipSecondaries, _ptr_aipSecondaries); err != nil {
		return err
	}
	_ptr_aipNotify := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Notify == nil {
			o.Notify = &AddrArray{}
		}
		if err := o.Notify.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipNotify := func(ptr interface{}) { o.Notify = *ptr.(**AddrArray) }
	if err := w.ReadPointer(&o.Notify, _s_aipNotify, _ptr_aipNotify); err != nil {
		return err
	}
	return nil
}

// ZoneDatabaseW2K structure represents DNS_RPC_ZONE_DATABASE_W2K RPC structure.
//
// The DNS_RPC_ZONE_DATABASE_W2K structure specifies how a DNS zone is stored in persistent
// storage.
type ZoneDatabaseW2K struct {
	// fDsIntegrated: This MUST be set TRUE (0x00000001) if the zone is stored in a directory
	// server, or FALSE (0x00000000) if not.
	DSIntegrated uint32 `idl:"name:fDsIntegrated" json:"ds_integrated"`
	// pszFileName: A pointer to a null-terminated UTF-8 string that specifies the name
	// of the file in which this zone is stored, or NULL if this zone is to be stored in
	// a directory server or in a file with the default file name for the zone.
	FileName string `idl:"name:pszFileName;string" json:"file_name"`
}

func (o *ZoneDatabaseW2K) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ZoneDatabaseW2K) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.DSIntegrated); err != nil {
		return err
	}
	if o.FileName != "" {
		_ptr_pszFileName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.FileName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.FileName, _ptr_pszFileName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ZoneDatabaseW2K) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.DSIntegrated); err != nil {
		return err
	}
	_ptr_pszFileName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.FileName); err != nil {
			return err
		}
		return nil
	})
	_s_pszFileName := func(ptr interface{}) { o.FileName = *ptr.(*string) }
	if err := w.ReadPointer(&o.FileName, _s_pszFileName, _ptr_pszFileName); err != nil {
		return err
	}
	return nil
}

// ZoneDatabase_NET structure represents DNS_RPC_ZONE_DATABASE_DOTNET RPC structure.
//
// In the DNS_RPC_ZONE_DATABASE_DOTNET structure all fields have same definition as
// specified in section DNS_RPC_ZONE_DATABASE_W2K (section 2.2.5.2.6.1), with the following
// exceptions:
type ZoneDatabase_NET struct {
	// dwRpcStructureVersion: The DNS management structure version number. This MUST be
	// set to 0x00000001.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// dwReserved0: Reserved for future use. The sender MUST set this value to 0x00000000
	// and it MUST be ignored by the receiver.
	_            uint32 `idl:"name:dwReserved0"`
	DSIntegrated uint32 `idl:"name:fDsIntegrated" json:"ds_integrated"`
	FileName     string `idl:"name:pszFileName;string" json:"file_name"`
}

func (o *ZoneDatabase_NET) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ZoneDatabase_NET) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.DSIntegrated); err != nil {
		return err
	}
	if o.FileName != "" {
		_ptr_pszFileName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.FileName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.FileName, _ptr_pszFileName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ZoneDatabase_NET) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	if err := w.ReadData(&o.DSIntegrated); err != nil {
		return err
	}
	_ptr_pszFileName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.FileName); err != nil {
			return err
		}
		return nil
	})
	_s_pszFileName := func(ptr interface{}) { o.FileName = *ptr.(*string) }
	if err := w.ReadPointer(&o.FileName, _s_pszFileName, _ptr_pszFileName); err != nil {
		return err
	}
	return nil
}

// ZoneDatabase structure represents DNS_RPC_ZONE_DATABASE RPC structure.
//
// The DNS_RPC_ZONE_DATABASE structure contains information about the persistent data
// store for a zone on the DNS server. There are different versions of the DNS_RPC_ZONE_DATABASE
// structure. The DNS server MUST use the structure corresponding to the value of dwClientVersion
// in DNS Server Management Protocol method calls (section 3.1.4) as shown in the following
// table, or if the method call does not specify the value of dwClientVersion, the DNS_RPC_ZONE_DATABASE_W2K
// (section 2.2.5.2.6.1) version of the structure MUST be used.
//
//	+------------+----------------------------------------------------+
//	|            |                                                    |
//	|   VALUE    |                     STRUCTURE                      |
//	|            |                                                    |
//	+------------+----------------------------------------------------+
//	+------------+----------------------------------------------------+
//	| 0x00000000 | DNS_RPC_ZONE_DATABASE_W2K (section 2.2.5.2.6.1)    |
//	+------------+----------------------------------------------------+
//	| 0x00060000 | DNS_RPC_ZONE_DATABASE_DOTNET (section 2.2.5.2.6.2) |
//	+------------+----------------------------------------------------+
type ZoneDatabase struct {
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	_                   uint32 `idl:"name:dwReserved0"`
	DSIntegrated        uint32 `idl:"name:fDsIntegrated" json:"ds_integrated"`
	FileName            string `idl:"name:pszFileName;string" json:"file_name"`
}

func (o *ZoneDatabase) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ZoneDatabase) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.DSIntegrated); err != nil {
		return err
	}
	if o.FileName != "" {
		_ptr_pszFileName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.FileName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.FileName, _ptr_pszFileName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ZoneDatabase) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	if err := w.ReadData(&o.DSIntegrated); err != nil {
		return err
	}
	_ptr_pszFileName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.FileName); err != nil {
			return err
		}
		return nil
	})
	_s_pszFileName := func(ptr interface{}) { o.FileName = *ptr.(*string) }
	if err := w.ReadPointer(&o.FileName, _s_pszFileName, _ptr_pszFileName); err != nil {
		return err
	}
	return nil
}

// ZoneChangeDP structure represents DNS_RPC_ZONE_CHANGE_DP RPC structure.
//
// The DNS_RPC_ZONE_CHANGE_DP structure contains information required to move a zone
// to a different application directory partition on the DNS server.<61>
type ZoneChangeDP struct {
	// dwRpcStructureVersion: As specified in section 2.2.7.2.5.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// dwReserved0: As specified in section 2.2.7.2.1.
	_ uint32 `idl:"name:dwReserved0"`
	// pszDestPartition: A pointer to a null-terminated UTF-8 string that specifies the
	// distinguished name for a new application directory partition to which a zone is to
	// be moved.
	DestinationPartition string `idl:"name:pszDestPartition;string" json:"destination_partition"`
}

func (o *ZoneChangeDP) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ZoneChangeDP) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if o.DestinationPartition != "" {
		_ptr_pszDestPartition := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.DestinationPartition); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DestinationPartition, _ptr_pszDestPartition); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ZoneChangeDP) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	_ptr_pszDestPartition := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.DestinationPartition); err != nil {
			return err
		}
		return nil
	})
	_s_pszDestPartition := func(ptr interface{}) { o.DestinationPartition = *ptr.(*string) }
	if err := w.ReadPointer(&o.DestinationPartition, _s_pszDestPartition, _ptr_pszDestPartition); err != nil {
		return err
	}
	return nil
}

// ZoneInfoW2K structure represents DNS_RPC_ZONE_INFO_W2K RPC structure.
//
// The DNS_RPC_ZONE_INFO_W2K structure is used to specify detailed DNS zone information.
type ZoneInfoW2K struct {
	// pszZoneName: A pointer to a null-terminated Unicode string that contains a zone name.
	ZoneName string `idl:"name:pszZoneName;string" json:"zone_name"`
	// dwZoneType: The zone type. This MUST be set to one of the allowed values as specified
	// in DNS_ZONE_TYPE (section 2.2.5.1.1).
	ZoneType uint32 `idl:"name:dwZoneType" json:"zone_type"`
	// fReverse: A Boolean value where TRUE (0x00000001) indicates this is a reverse lookup
	// zone and FALSE (0x00000000) indicates this is a forward lookup zone.
	Reverse uint32 `idl:"name:fReverse" json:"reverse"`
	// fAllowUpdate: A value that indicates what kind dynamic updates, as specified in [RFC2136],
	// are allowed for this zone. This MUST be set to one of the following values:
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|                                 |                                                                                  |
	//	|         CONSTANT/VALUE          |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| ZONE_UPDATE_OFF 0x00000000      | No updates are allowed for the zone.                                             |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| ZONE_UPDATE_UNSECURE 0x00000001 | All updates (secure and unsecure) are allowed for the zone.                      |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| ZONE_UPDATE_SECURE 0x00000002   | The zone only allows secure updates, that is, DNS packet MUST have a TSIG        |
	//	|                                 | [RFC2845] present in the additional section.                                     |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	AllowUpdate uint32 `idl:"name:fAllowUpdate" json:"allow_update"`
	// fPaused: A Boolean value indicates whether zone operations are currently paused.
	// TRUE indicates that the DNS server does not use this zone to answer queries or accept
	// updates. FALSE indicates that the zone is handled normally.
	Paused uint32 `idl:"name:fPaused" json:"paused"`
	// fShutdown: A Boolean value that indicates whether this zone is currently shut down.
	Shutdown uint32 `idl:"name:fShutdown" json:"shutdown"`
	// fAutoCreated: A Boolean value that indicates whether this zone was autocreated by
	// the DNS server at boot time or when a virtualization instance is created via the
	// CreateVirtualizationInstance RPC operation (section 3.1.4.1).
	AutoCreated uint32 `idl:"name:fAutoCreated" json:"auto_created"`
	// fUseDatabase: A Boolean value that indicates whether this zone is stored in a directory
	// server.
	UseDatabase uint32 `idl:"name:fUseDatabase" json:"use_database"`
	// pszDataFile: A pointer to a null-terminated UTF-8 character string that specifies
	// the name (with no path) of the zone file for a file-based zone or NULL if this zone
	// is not stored in a file.
	DataFile string `idl:"name:pszDataFile;string" json:"data_file"`
	// aipMasters: A pointer to a structure of type IP4_ARRAY (section 2.2.3.2.1) that specifies
	// a list of IPv4 addresses of the remote DNS servers that can be sources of information
	// for this zone to perform zone transfers by a secondary. This value is applicable
	// for secondary, stub and forwarder zones only and MUST be NULL for all other zone
	// types.
	Masters *IPv4Array `idl:"name:aipMasters" json:"masters"`
	// fSecureSecondaries: The secondary security settings configured for a zone on the
	// master DNS server. The DNS server MUST respond to zone transfer requests from a secondary
	// server according to the behavior description corresponding to this flag's value as
	// specified in DNS_ZONE_SECONDARY_SECURITY (section 2.2.5.1.2). This value MUST be
	// set to one of the allowed values as specified in DNS_ZONE_SECONDARY_SECURITY (section
	// 2.2.5.1.2).
	SecureSecondaries uint32 `idl:"name:fSecureSecondaries" json:"secure_secondaries"`
	// fNotifyLevel: This parameter is ignored, and for dwZoneType parameter setting DNS_ZONE_TYPE_PRIMARY
	// (specified in section 2.2.5.1.1) and for DNS servers that are not directory service-integrated,
	// the zone notification setting is set to ZONE_NOTIFY_ALL_SECONDARIES. Otherwise, it
	// is set to ZONE_NOTIFY_LIST_ONLY, as specified in section 2.2.5.1.3.
	NotifyLevel uint32 `idl:"name:fNotifyLevel" json:"notify_level"`
	// aipSecondaries: A pointer to a structure of type IP4_ARRAY (section 2.2.3.2.1) that
	// specifies a list of IPv4 addresses of the remote DNS servers that are secondary DNS
	// servers for this zone, or NULL if there are no secondary DNS servers. If fSecureSecondaries
	// is set to ZONE_SECSECURE_LIST_ONLY then only zone transfers from IP addresses in
	// this list will be honored.
	Secondaries *IPv4Array `idl:"name:aipSecondaries" json:"secondaries"`
	// aipNotify: A pointer to a structure of type IP4_ARRAY (section 2.2.3.2.1) that specifies
	// a list of IPv4 addresses of the remote DNS servers that are secondaries for this
	// zone, for which this DNS server is acting as master and the DNS server will send
	// zone notifications to these secondary servers, as directed by the value of fNotifyLevel
	// above.
	Notify *IPv4Array `idl:"name:aipNotify" json:"notify"`
	// fUseWins: A Boolean value that indicates whether WINS resource record lookup is enabled
	// for this forward lookup zones.
	UseWINS uint32 `idl:"name:fUseWins" json:"use_wins"`
	// fUseNbstat: A Boolean value that indicates whether WINS-R resource record lookup
	// is enabled for this reverse lookup zones.
	UseNBSTAT uint32 `idl:"name:fUseNbstat" json:"use_nbstat"`
	// fAging: A Boolean value where TRUE (0x00000001) indicates that aging is enabled for
	// resource records in this zone, so the time-stamps of records in the zone will be
	// updated when server receives dynamic update request for that record; whereas FALSE(0x00000000)
	// indicates, that the time-stamps of the records in the zone will not be updated.
	Aging uint32 `idl:"name:fAging" json:"aging"`
	// dwNoRefreshInterval: The time interval, in hours, that is configured as NoRefresh
	// interval value for this zone. This value determines the time interval between the
	// last update of a record's time-stamp and the earliest instance when that time-stamp
	// can be refreshed.
	NoRefreshInterval uint32 `idl:"name:dwNoRefreshInterval" json:"no_refresh_interval"`
	// dwRefreshInterval: The time interval, in hours, that is configured as the refresh
	// interval value for this zone. Records that have not been refreshed by the expiration
	// of this interval are eligible to be removed during the next scavenging cycle performed
	// by the DNS server.
	RefreshInterval uint32 `idl:"name:dwRefreshInterval" json:"refresh_interval"`
	// dwAvailForScavengeTime: The time interval, in hours, that is available before the
	// scheduled next scavenging cycle for this zone.
	AvailForScavengeTime uint32 `idl:"name:dwAvailForScavengeTime" json:"avail_for_scavenge_time"`
	// aipScavengeServers: A pointer to a structure of type IP4_ARRAY (section 2.2.3.2.1)
	// that specifies a list of IPv4 addresses of the DNS servers that will perform scavenging
	// for this zone. This value is applicable for zones of type DNS_ZONE_TYPE_PRIMARY (section
	// DNS_ZONE_TYPE) only. If this value is NULL,  there are no restrictions on which
	// DNS server can perform scavenging for this zone.
	ScavengeServers *IPv4Array `idl:"name:aipScavengeServers" json:"scavenge_servers"`
	// pvReserved1: Reserved for future use. Server MUST set to zero and receiver MUST ignore
	// this value.
	_ uint32 `idl:"name:pvReserved1"`
	// pvReserved2: Reserved for future use. Server MUST set to zero and receiver MUST ignore
	// this value.
	_ uint32 `idl:"name:pvReserved2"`
	// pvReserved3: Reserved for future use. Server MUST set to zero and receiver MUST ignore
	// this value.
	_ uint32 `idl:"name:pvReserved3"`
	// pvReserved4: Reserved for future use. Server MUST set to zero and receiver MUST ignore
	// this value.
	_ uint32 `idl:"name:pvReserved4"`
}

func (o *ZoneInfoW2K) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ZoneInfoW2K) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.ZoneName != "" {
		_ptr_pszZoneName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.ZoneName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ZoneName, _ptr_pszZoneName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ZoneType); err != nil {
		return err
	}
	if err := w.WriteData(o.Reverse); err != nil {
		return err
	}
	if err := w.WriteData(o.AllowUpdate); err != nil {
		return err
	}
	if err := w.WriteData(o.Paused); err != nil {
		return err
	}
	if err := w.WriteData(o.Shutdown); err != nil {
		return err
	}
	if err := w.WriteData(o.AutoCreated); err != nil {
		return err
	}
	if err := w.WriteData(o.UseDatabase); err != nil {
		return err
	}
	if o.DataFile != "" {
		_ptr_pszDataFile := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.DataFile); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DataFile, _ptr_pszDataFile); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Masters != nil {
		_ptr_aipMasters := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Masters != nil {
				if err := o.Masters.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPv4Array{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Masters, _ptr_aipMasters); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SecureSecondaries); err != nil {
		return err
	}
	if err := w.WriteData(o.NotifyLevel); err != nil {
		return err
	}
	if o.Secondaries != nil {
		_ptr_aipSecondaries := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Secondaries != nil {
				if err := o.Secondaries.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPv4Array{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Secondaries, _ptr_aipSecondaries); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Notify != nil {
		_ptr_aipNotify := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Notify != nil {
				if err := o.Notify.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPv4Array{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Notify, _ptr_aipNotify); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.UseWINS); err != nil {
		return err
	}
	if err := w.WriteData(o.UseNBSTAT); err != nil {
		return err
	}
	if err := w.WriteData(o.Aging); err != nil {
		return err
	}
	if err := w.WriteData(o.NoRefreshInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.RefreshInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.AvailForScavengeTime); err != nil {
		return err
	}
	if o.ScavengeServers != nil {
		_ptr_aipScavengeServers := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ScavengeServers != nil {
				if err := o.ScavengeServers.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPv4Array{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ScavengeServers, _ptr_aipScavengeServers); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	// reserved pvReserved1
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	// reserved pvReserved2
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	// reserved pvReserved3
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	// reserved pvReserved4
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *ZoneInfoW2K) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_pszZoneName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.ZoneName); err != nil {
			return err
		}
		return nil
	})
	_s_pszZoneName := func(ptr interface{}) { o.ZoneName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ZoneName, _s_pszZoneName, _ptr_pszZoneName); err != nil {
		return err
	}
	if err := w.ReadData(&o.ZoneType); err != nil {
		return err
	}
	if err := w.ReadData(&o.Reverse); err != nil {
		return err
	}
	if err := w.ReadData(&o.AllowUpdate); err != nil {
		return err
	}
	if err := w.ReadData(&o.Paused); err != nil {
		return err
	}
	if err := w.ReadData(&o.Shutdown); err != nil {
		return err
	}
	if err := w.ReadData(&o.AutoCreated); err != nil {
		return err
	}
	if err := w.ReadData(&o.UseDatabase); err != nil {
		return err
	}
	_ptr_pszDataFile := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.DataFile); err != nil {
			return err
		}
		return nil
	})
	_s_pszDataFile := func(ptr interface{}) { o.DataFile = *ptr.(*string) }
	if err := w.ReadPointer(&o.DataFile, _s_pszDataFile, _ptr_pszDataFile); err != nil {
		return err
	}
	_ptr_aipMasters := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Masters == nil {
			o.Masters = &IPv4Array{}
		}
		if err := o.Masters.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipMasters := func(ptr interface{}) { o.Masters = *ptr.(**IPv4Array) }
	if err := w.ReadPointer(&o.Masters, _s_aipMasters, _ptr_aipMasters); err != nil {
		return err
	}
	if err := w.ReadData(&o.SecureSecondaries); err != nil {
		return err
	}
	if err := w.ReadData(&o.NotifyLevel); err != nil {
		return err
	}
	_ptr_aipSecondaries := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Secondaries == nil {
			o.Secondaries = &IPv4Array{}
		}
		if err := o.Secondaries.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipSecondaries := func(ptr interface{}) { o.Secondaries = *ptr.(**IPv4Array) }
	if err := w.ReadPointer(&o.Secondaries, _s_aipSecondaries, _ptr_aipSecondaries); err != nil {
		return err
	}
	_ptr_aipNotify := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Notify == nil {
			o.Notify = &IPv4Array{}
		}
		if err := o.Notify.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipNotify := func(ptr interface{}) { o.Notify = *ptr.(**IPv4Array) }
	if err := w.ReadPointer(&o.Notify, _s_aipNotify, _ptr_aipNotify); err != nil {
		return err
	}
	if err := w.ReadData(&o.UseWINS); err != nil {
		return err
	}
	if err := w.ReadData(&o.UseNBSTAT); err != nil {
		return err
	}
	if err := w.ReadData(&o.Aging); err != nil {
		return err
	}
	if err := w.ReadData(&o.NoRefreshInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.RefreshInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.AvailForScavengeTime); err != nil {
		return err
	}
	_ptr_aipScavengeServers := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ScavengeServers == nil {
			o.ScavengeServers = &IPv4Array{}
		}
		if err := o.ScavengeServers.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipScavengeServers := func(ptr interface{}) { o.ScavengeServers = *ptr.(**IPv4Array) }
	if err := w.ReadPointer(&o.ScavengeServers, _s_aipScavengeServers, _ptr_aipScavengeServers); err != nil {
		return err
	}
	// reserved pvReserved1
	var _pvReserved1 uint32
	if err := w.ReadData(&_pvReserved1); err != nil {
		return err
	}
	// reserved pvReserved2
	var _pvReserved2 uint32
	if err := w.ReadData(&_pvReserved2); err != nil {
		return err
	}
	// reserved pvReserved3
	var _pvReserved3 uint32
	if err := w.ReadData(&_pvReserved3); err != nil {
		return err
	}
	// reserved pvReserved4
	var _pvReserved4 uint32
	if err := w.ReadData(&_pvReserved4); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// ZoneInfo_NET structure represents DNS_RPC_ZONE_INFO_DOTNET RPC structure.
//
// In the DNS_RPC_ZONE_INFO_DOTNET structure all fields have same definition as specified
// in section DNS_RPC_ZONE_INFO_W2K (section 2.2.5.2.4.1), with the following exceptions:
type ZoneInfo_NET struct {
	// dwRpcStructureVersion: The DNS server management structure version number. This value
	// SHOULD<42> be set to 0x00000001.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// dwReserved0: Reserved for future use. Server will set to zero and receiver MUST ignore
	// this value.
	_                    uint32     `idl:"name:dwReserved0"`
	ZoneName             string     `idl:"name:pszZoneName;string" json:"zone_name"`
	ZoneType             uint32     `idl:"name:dwZoneType" json:"zone_type"`
	Reverse              uint32     `idl:"name:fReverse" json:"reverse"`
	AllowUpdate          uint32     `idl:"name:fAllowUpdate" json:"allow_update"`
	Paused               uint32     `idl:"name:fPaused" json:"paused"`
	Shutdown             uint32     `idl:"name:fShutdown" json:"shutdown"`
	AutoCreated          uint32     `idl:"name:fAutoCreated" json:"auto_created"`
	UseDatabase          uint32     `idl:"name:fUseDatabase" json:"use_database"`
	DataFile             string     `idl:"name:pszDataFile;string" json:"data_file"`
	Masters              *IPv4Array `idl:"name:aipMasters" json:"masters"`
	SecureSecondaries    uint32     `idl:"name:fSecureSecondaries" json:"secure_secondaries"`
	NotifyLevel          uint32     `idl:"name:fNotifyLevel" json:"notify_level"`
	Secondaries          *IPv4Array `idl:"name:aipSecondaries" json:"secondaries"`
	Notify               *IPv4Array `idl:"name:aipNotify" json:"notify"`
	UseWINS              uint32     `idl:"name:fUseWins" json:"use_wins"`
	UseNBSTAT            uint32     `idl:"name:fUseNbstat" json:"use_nbstat"`
	Aging                uint32     `idl:"name:fAging" json:"aging"`
	NoRefreshInterval    uint32     `idl:"name:dwNoRefreshInterval" json:"no_refresh_interval"`
	RefreshInterval      uint32     `idl:"name:dwRefreshInterval" json:"refresh_interval"`
	AvailForScavengeTime uint32     `idl:"name:dwAvailForScavengeTime" json:"avail_for_scavenge_time"`
	ScavengeServers      *IPv4Array `idl:"name:aipScavengeServers" json:"scavenge_servers"`
	// dwForwarderTimeout (4 bytes): The forwarder timeout value for a zone, in seconds.
	// This is the number of seconds the DNS server waits for response from a remote DNS
	// server for a forwarded query. This value is applicable for zones of type DNS_ZONE_TYPE_FORWARDER
	// (section 2.2.5.1.1). For all other zone types, senders MUST set this value to zero
	// and receivers MUST ignore it.
	ForwarderTimeout uint32 `idl:"name:dwForwarderTimeout" json:"forwarder_timeout"`
	// fForwarderSlave: A Boolean value indicating whether the DNS server is not allowed
	// to perform recursion while resolving names for this zone. The DNS server MUST fail
	// queries after getting failure from all forwarded servers, if the value is TRUE (0x00000001).
	// Otherwise normal recursion MUST be performed. This value is applicable for zones
	// of type DNS_ZONE_TYPE_FORWARDER (section 2.2.5.1.1). For all other zone types, senders
	// MUST set this value to zero and receivers MUST ignore it.
	ForwarderSlave uint32 `idl:"name:fForwarderSlave" json:"forwarder_slave"`
	// aipLocalMasters: A pointer to a structure of type IP4_ARRAY (section 2.2.3.2.1) that
	// specifies a list of IPv4 addresses of primary DNS servers for this zone. If this
	// value is not NULL then it overrides the master servers list configured in the directory
	// server.
	LocalMasters *IPv4Array `idl:"name:aipLocalMasters" json:"local_masters"`
	// dwDpFlags: Flag value that specifies information about the application directory
	// partition in which this zone is stored. This MUST be set to any combination of the
	// DNS_RPC_DP_FLAGS (section 2.2.7.1.1) or zero if this zone is not stored in a directory
	// server.
	DPFlags uint32 `idl:"name:dwDpFlags" json:"dp_flags"`
	// pszDpFqdn: A pointer to a null-terminated UTF-8 string that specifies the FQDN of
	// the application directory partition in which this zone is stored. If the zone is
	// not stored in an application directory partition this value MUST be NULL.
	DPFQDN string `idl:"name:pszDpFqdn;string" json:"dp_fqdn"`
	// pwszZoneDn: Pointer to a null-terminated Unicode string that specifies the distinguished
	// name for the zone if the zone is stored the directory server. This field is applicable
	// for directory server integrated zones only. The value MUST be NULL if this zone is
	// not stored in the directory server.
	ZoneDN string `idl:"name:pwszZoneDn;string" json:"zone_dn"`
	// dwLastSuccessfulSoaCheck: The time-stamp at which last SOA record was received successfully
	// from the primary DNS server for this zone. This field is applicable only for zones
	// which are secondary or non-authoritative. For all other zones this value MUST be
	// set to zero by senders and MUST be ignored by receivers.
	LastSuccessfulSOACheck uint32 `idl:"name:dwLastSuccessfulSoaCheck" json:"last_successful_soa_check"`
	// dwLastSuccessfulXfr: The time-stamp at which last zone transfer was completed successfully
	// from the primary DNS server for this zone. This field is applicable only for zones
	// which are secondary or non-authoritative. For all other zones this value MUST be
	// set to zero by senders and MUST be ignored by receivers.
	LastSuccessfulXFR uint32 `idl:"name:dwLastSuccessfulXfr" json:"last_successful_xfr"`
	// dwReserved1: Reserved for future use. Server MUST set to zero and receiver MUST ignore
	// this value.
	_ uint32 `idl:"name:dwReserved1"`
	// dwReserved2: Reserved for future use. Server MUST set to zero and receiver MUST ignore
	// this value.
	_ uint32 `idl:"name:dwReserved2"`
	// dwReserved3: Reserved for future use. Server MUST set to zero and receiver MUST ignore
	// this value.
	_ uint32 `idl:"name:dwReserved3"`
	// dwReserved4: Reserved for future use. Server MUST set to zero and receiver MUST ignore
	// this value.
	_ uint32 `idl:"name:dwReserved4"`
	// dwReserved5: Reserved for future use. Server MUST set to zero and receiver MUST ignore
	// this value.
	_ uint32 `idl:"name:dwReserved5"`
	// pReserved1: Reserved for future use. Server MUST set to zero and receiver MUST ignore
	// this value.
	_ string `idl:"name:pReserved1;string"`
	// pReserved2: Reserved for future use. Server MUST set to zero and receiver MUST ignore
	// this value.
	_ string `idl:"name:pReserved2;string"`
	// pReserved3: Reserved for future use. Server MUST set to zero and receiver MUST ignore
	// this value.
	_ string `idl:"name:pReserved3;string"`
	// pReserved4: Reserved for future use. Server MUST set to zero and receiver MUST ignore
	// this value.
	_ string `idl:"name:pReserved4;string"`
}

func (o *ZoneInfo_NET) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ZoneInfo_NET) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if o.ZoneName != "" {
		_ptr_pszZoneName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.ZoneName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ZoneName, _ptr_pszZoneName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ZoneType); err != nil {
		return err
	}
	if err := w.WriteData(o.Reverse); err != nil {
		return err
	}
	if err := w.WriteData(o.AllowUpdate); err != nil {
		return err
	}
	if err := w.WriteData(o.Paused); err != nil {
		return err
	}
	if err := w.WriteData(o.Shutdown); err != nil {
		return err
	}
	if err := w.WriteData(o.AutoCreated); err != nil {
		return err
	}
	if err := w.WriteData(o.UseDatabase); err != nil {
		return err
	}
	if o.DataFile != "" {
		_ptr_pszDataFile := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.DataFile); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DataFile, _ptr_pszDataFile); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Masters != nil {
		_ptr_aipMasters := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Masters != nil {
				if err := o.Masters.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPv4Array{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Masters, _ptr_aipMasters); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SecureSecondaries); err != nil {
		return err
	}
	if err := w.WriteData(o.NotifyLevel); err != nil {
		return err
	}
	if o.Secondaries != nil {
		_ptr_aipSecondaries := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Secondaries != nil {
				if err := o.Secondaries.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPv4Array{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Secondaries, _ptr_aipSecondaries); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Notify != nil {
		_ptr_aipNotify := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Notify != nil {
				if err := o.Notify.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPv4Array{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Notify, _ptr_aipNotify); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.UseWINS); err != nil {
		return err
	}
	if err := w.WriteData(o.UseNBSTAT); err != nil {
		return err
	}
	if err := w.WriteData(o.Aging); err != nil {
		return err
	}
	if err := w.WriteData(o.NoRefreshInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.RefreshInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.AvailForScavengeTime); err != nil {
		return err
	}
	if o.ScavengeServers != nil {
		_ptr_aipScavengeServers := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ScavengeServers != nil {
				if err := o.ScavengeServers.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPv4Array{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ScavengeServers, _ptr_aipScavengeServers); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ForwarderTimeout); err != nil {
		return err
	}
	if err := w.WriteData(o.ForwarderSlave); err != nil {
		return err
	}
	if o.LocalMasters != nil {
		_ptr_aipLocalMasters := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.LocalMasters != nil {
				if err := o.LocalMasters.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPv4Array{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalMasters, _ptr_aipLocalMasters); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.DPFlags); err != nil {
		return err
	}
	if o.DPFQDN != "" {
		_ptr_pszDpFqdn := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.DPFQDN); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DPFQDN, _ptr_pszDpFqdn); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ZoneDN != "" {
		_ptr_pwszZoneDn := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ZoneDN); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ZoneDN, _ptr_pwszZoneDn); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LastSuccessfulSOACheck); err != nil {
		return err
	}
	if err := w.WriteData(o.LastSuccessfulXFR); err != nil {
		return err
	}
	// reserved dwReserved1
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	// reserved dwReserved2
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	// reserved dwReserved3
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	// reserved dwReserved4
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	// reserved dwReserved5
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	// reserved pReserved1
	if err := w.WritePointer(nil); err != nil {
		return err
	}
	// reserved pReserved2
	if err := w.WritePointer(nil); err != nil {
		return err
	}
	// reserved pReserved3
	if err := w.WritePointer(nil); err != nil {
		return err
	}
	// reserved pReserved4
	if err := w.WritePointer(nil); err != nil {
		return err
	}
	return nil
}
func (o *ZoneInfo_NET) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	_ptr_pszZoneName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.ZoneName); err != nil {
			return err
		}
		return nil
	})
	_s_pszZoneName := func(ptr interface{}) { o.ZoneName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ZoneName, _s_pszZoneName, _ptr_pszZoneName); err != nil {
		return err
	}
	if err := w.ReadData(&o.ZoneType); err != nil {
		return err
	}
	if err := w.ReadData(&o.Reverse); err != nil {
		return err
	}
	if err := w.ReadData(&o.AllowUpdate); err != nil {
		return err
	}
	if err := w.ReadData(&o.Paused); err != nil {
		return err
	}
	if err := w.ReadData(&o.Shutdown); err != nil {
		return err
	}
	if err := w.ReadData(&o.AutoCreated); err != nil {
		return err
	}
	if err := w.ReadData(&o.UseDatabase); err != nil {
		return err
	}
	_ptr_pszDataFile := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.DataFile); err != nil {
			return err
		}
		return nil
	})
	_s_pszDataFile := func(ptr interface{}) { o.DataFile = *ptr.(*string) }
	if err := w.ReadPointer(&o.DataFile, _s_pszDataFile, _ptr_pszDataFile); err != nil {
		return err
	}
	_ptr_aipMasters := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Masters == nil {
			o.Masters = &IPv4Array{}
		}
		if err := o.Masters.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipMasters := func(ptr interface{}) { o.Masters = *ptr.(**IPv4Array) }
	if err := w.ReadPointer(&o.Masters, _s_aipMasters, _ptr_aipMasters); err != nil {
		return err
	}
	if err := w.ReadData(&o.SecureSecondaries); err != nil {
		return err
	}
	if err := w.ReadData(&o.NotifyLevel); err != nil {
		return err
	}
	_ptr_aipSecondaries := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Secondaries == nil {
			o.Secondaries = &IPv4Array{}
		}
		if err := o.Secondaries.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipSecondaries := func(ptr interface{}) { o.Secondaries = *ptr.(**IPv4Array) }
	if err := w.ReadPointer(&o.Secondaries, _s_aipSecondaries, _ptr_aipSecondaries); err != nil {
		return err
	}
	_ptr_aipNotify := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Notify == nil {
			o.Notify = &IPv4Array{}
		}
		if err := o.Notify.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipNotify := func(ptr interface{}) { o.Notify = *ptr.(**IPv4Array) }
	if err := w.ReadPointer(&o.Notify, _s_aipNotify, _ptr_aipNotify); err != nil {
		return err
	}
	if err := w.ReadData(&o.UseWINS); err != nil {
		return err
	}
	if err := w.ReadData(&o.UseNBSTAT); err != nil {
		return err
	}
	if err := w.ReadData(&o.Aging); err != nil {
		return err
	}
	if err := w.ReadData(&o.NoRefreshInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.RefreshInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.AvailForScavengeTime); err != nil {
		return err
	}
	_ptr_aipScavengeServers := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ScavengeServers == nil {
			o.ScavengeServers = &IPv4Array{}
		}
		if err := o.ScavengeServers.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipScavengeServers := func(ptr interface{}) { o.ScavengeServers = *ptr.(**IPv4Array) }
	if err := w.ReadPointer(&o.ScavengeServers, _s_aipScavengeServers, _ptr_aipScavengeServers); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwarderTimeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwarderSlave); err != nil {
		return err
	}
	_ptr_aipLocalMasters := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.LocalMasters == nil {
			o.LocalMasters = &IPv4Array{}
		}
		if err := o.LocalMasters.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipLocalMasters := func(ptr interface{}) { o.LocalMasters = *ptr.(**IPv4Array) }
	if err := w.ReadPointer(&o.LocalMasters, _s_aipLocalMasters, _ptr_aipLocalMasters); err != nil {
		return err
	}
	if err := w.ReadData(&o.DPFlags); err != nil {
		return err
	}
	_ptr_pszDpFqdn := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.DPFQDN); err != nil {
			return err
		}
		return nil
	})
	_s_pszDpFqdn := func(ptr interface{}) { o.DPFQDN = *ptr.(*string) }
	if err := w.ReadPointer(&o.DPFQDN, _s_pszDpFqdn, _ptr_pszDpFqdn); err != nil {
		return err
	}
	_ptr_pwszZoneDn := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ZoneDN); err != nil {
			return err
		}
		return nil
	})
	_s_pwszZoneDn := func(ptr interface{}) { o.ZoneDN = *ptr.(*string) }
	if err := w.ReadPointer(&o.ZoneDN, _s_pwszZoneDn, _ptr_pwszZoneDn); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastSuccessfulSOACheck); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastSuccessfulXFR); err != nil {
		return err
	}
	// reserved dwReserved1
	var _dwReserved1 uint32
	if err := w.ReadData(&_dwReserved1); err != nil {
		return err
	}
	// reserved dwReserved2
	var _dwReserved2 uint32
	if err := w.ReadData(&_dwReserved2); err != nil {
		return err
	}
	// reserved dwReserved3
	var _dwReserved3 uint32
	if err := w.ReadData(&_dwReserved3); err != nil {
		return err
	}
	// reserved dwReserved4
	var _dwReserved4 uint32
	if err := w.ReadData(&_dwReserved4); err != nil {
		return err
	}
	// reserved dwReserved5
	var _dwReserved5 uint32
	if err := w.ReadData(&_dwReserved5); err != nil {
		return err
	}
	// reserved pReserved1
	var _pReserved1 string
	_ptr_pReserved1 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &_pReserved1); err != nil {
			return err
		}
		return nil
	})
	_s_pReserved1 := func(ptr interface{}) { _pReserved1 = *ptr.(*string) }
	if err := w.ReadPointer(&_pReserved1, _s_pReserved1, _ptr_pReserved1); err != nil {
		return err
	}
	// reserved pReserved2
	var _pReserved2 string
	_ptr_pReserved2 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &_pReserved2); err != nil {
			return err
		}
		return nil
	})
	_s_pReserved2 := func(ptr interface{}) { _pReserved2 = *ptr.(*string) }
	if err := w.ReadPointer(&_pReserved2, _s_pReserved2, _ptr_pReserved2); err != nil {
		return err
	}
	// reserved pReserved3
	var _pReserved3 string
	_ptr_pReserved3 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &_pReserved3); err != nil {
			return err
		}
		return nil
	})
	_s_pReserved3 := func(ptr interface{}) { _pReserved3 = *ptr.(*string) }
	if err := w.ReadPointer(&_pReserved3, _s_pReserved3, _ptr_pReserved3); err != nil {
		return err
	}
	// reserved pReserved4
	var _pReserved4 string
	_ptr_pReserved4 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &_pReserved4); err != nil {
			return err
		}
		return nil
	})
	_s_pReserved4 := func(ptr interface{}) { _pReserved4 = *ptr.(*string) }
	if err := w.ReadPointer(&_pReserved4, _s_pReserved4, _ptr_pReserved4); err != nil {
		return err
	}
	return nil
}

// ZoneInfoLonghorn structure represents DNS_RPC_ZONE_INFO_LONGHORN RPC structure.
//
// In the DNS_RPC_ZONE_INFO_LONGHORN structure all fields have same definition as specified
// in section DNS_RPC_ZONE_INFO_DOTNET (section 2.2.5.2.4.2), with the following exceptions:
type ZoneInfoLonghorn struct {
	// dwRpcStructureVersion: The DNS server management structure version number. It SHOULD<43>
	// be set to 0x00000002.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	_                   uint32 `idl:"name:dwReserved0"`
	ZoneName            string `idl:"name:pszZoneName;string" json:"zone_name"`
	ZoneType            uint32 `idl:"name:dwZoneType" json:"zone_type"`
	Reverse             uint32 `idl:"name:fReverse" json:"reverse"`
	AllowUpdate         uint32 `idl:"name:fAllowUpdate" json:"allow_update"`
	Paused              uint32 `idl:"name:fPaused" json:"paused"`
	Shutdown            uint32 `idl:"name:fShutdown" json:"shutdown"`
	AutoCreated         uint32 `idl:"name:fAutoCreated" json:"auto_created"`
	UseDatabase         uint32 `idl:"name:fUseDatabase" json:"use_database"`
	DataFile            string `idl:"name:pszDataFile;string" json:"data_file"`
	// aipMasters: A pointer to a structure of type DNS_ADDR_ARRAY (section 2.2.3.2.3) that
	// specifies a list of IP addresses of the remote DNS servers that can be sources of
	// information for this zone on which to perform zone transfers by a secondary. This
	// value is applicable for secondary, stub and forwarder zones only and MUST be NULL
	// for all other zone types.
	Masters           *AddrArray `idl:"name:aipMasters" json:"masters"`
	SecureSecondaries uint32     `idl:"name:fSecureSecondaries" json:"secure_secondaries"`
	NotifyLevel       uint32     `idl:"name:fNotifyLevel" json:"notify_level"`
	// aipSecondaries: A pointer to a structure of type DNS_ADDR_ARRAY (section 2.2.3.2.3)
	// that specifies a list of IP addresses of the remote DNS servers that are secondary
	// DNS servers for this zone, or NULL if there are no secondary DNS servers. If fSecureSecondaries
	// is set to ZONE_SECSECURE_LIST_ONLY, then only zone transfers from IP addresses in
	// this list will be honored.
	Secondaries *AddrArray `idl:"name:aipSecondaries" json:"secondaries"`
	// aipNotify: A pointer to a structure of type DNS_ADDR_ARRAY (section 2.2.3.2.3) that
	// specifies a list of IP addresses of the remote DNS servers that are secondaries for
	// this zone, for which this DNS server is acting as master and the DNS server will
	// send zone notifications to these secondary servers, as directed by the value of fNotifyLevel
	// above.
	Notify               *AddrArray `idl:"name:aipNotify" json:"notify"`
	UseWINS              uint32     `idl:"name:fUseWins" json:"use_wins"`
	UseNBSTAT            uint32     `idl:"name:fUseNbstat" json:"use_nbstat"`
	Aging                uint32     `idl:"name:fAging" json:"aging"`
	NoRefreshInterval    uint32     `idl:"name:dwNoRefreshInterval" json:"no_refresh_interval"`
	RefreshInterval      uint32     `idl:"name:dwRefreshInterval" json:"refresh_interval"`
	AvailForScavengeTime uint32     `idl:"name:dwAvailForScavengeTime" json:"avail_for_scavenge_time"`
	// aipScavengeServers: A pointer to a structure of type DNS_ADDR_ARRAY (section 2.2.3.2.3)
	// that specifies a list of IP addresses of the DNS servers that will perform scavenging
	// for this zone. This value is applicable for zones of type DNS_ZONE_TYPE_PRIMARY (section
	// 2.2.5.1.1) only. If this value is NULL, there are no restrictions on which DNS server
	// can perform scavenging for this zone.
	ScavengeServers  *AddrArray `idl:"name:aipScavengeServers" json:"scavenge_servers"`
	ForwarderTimeout uint32     `idl:"name:dwForwarderTimeout" json:"forwarder_timeout"`
	ForwarderSlave   uint32     `idl:"name:fForwarderSlave" json:"forwarder_slave"`
	// aipLocalMasters: A pointer to a structure of type DNS_ADDR_ARRAY (section 2.2.3.2.3)
	// that specifies a list of IP addresses of primary DNS servers for this zone. If this
	// value is not NULL then it overrides the masters servers list configured in the directory
	// server.
	LocalMasters           *AddrArray `idl:"name:aipLocalMasters" json:"local_masters"`
	DPFlags                uint32     `idl:"name:dwDpFlags" json:"dp_flags"`
	DPFQDN                 string     `idl:"name:pszDpFqdn;string" json:"dp_fqdn"`
	ZoneDN                 string     `idl:"name:pwszZoneDn;string" json:"zone_dn"`
	LastSuccessfulSOACheck uint32     `idl:"name:dwLastSuccessfulSoaCheck" json:"last_successful_soa_check"`
	LastSuccessfulXFR      uint32     `idl:"name:dwLastSuccessfulXfr" json:"last_successful_xfr"`
	// fQueuedForBackgroundLoad: This MUST be set to 0x00000001 if this zone is has not
	// yet been loaded from persistent storage.
	QueuedForBackgroundLoad uint32 `idl:"name:fQueuedForBackgroundLoad" json:"queued_for_background_load"`
	// fBackgroundLoadInProgress: This MUST be set to 0x00000001 if this zone is currently
	// being loaded from persistent storage, or 0x00000000 if it has been loaded.
	BackgroundLoadInProgress uint32 `idl:"name:fBackgroundLoadInProgress" json:"background_load_in_progress"`
	// fReadOnlyZone: This MUST be set to 0x00000001 if this zone is backed by a read-only
	// store that will not accept write operations, or 0x00000000 if not.
	ReadOnlyZone bool `idl:"name:fReadOnlyZone" json:"read_only_zone"`
	// dwLastXfrAttempt: The time-stamp at which last zone transfer was attempted by a DNS
	// server. This field is applicable only for zones which are secondary or not-authoritative.
	// For all other zones senders MUST set this value to zero and receivers MUST ignore
	// it.
	LastXFRAttempt uint32 `idl:"name:dwLastXfrAttempt" json:"last_xfr_attempt"`
	// dwLastXfrResult: The result of the last zone transfer attempted by server. This field
	// is applicable only for zones which are secondary or not-authoritative, and in this
	// case it MUST be either a Win32 error code, or 0xFFFFFFFF to indicate that a zone
	// transfer is currently in progress. For all other zones senders MUST set this value
	// to zero and receivers MUST ignore it.
	LastXFRResult uint32 `idl:"name:dwLastXfrResult" json:"last_xfr_result"`
}

func (o *ZoneInfoLonghorn) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ZoneInfoLonghorn) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if o.ZoneName != "" {
		_ptr_pszZoneName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.ZoneName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ZoneName, _ptr_pszZoneName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ZoneType); err != nil {
		return err
	}
	if err := w.WriteData(o.Reverse); err != nil {
		return err
	}
	if err := w.WriteData(o.AllowUpdate); err != nil {
		return err
	}
	if err := w.WriteData(o.Paused); err != nil {
		return err
	}
	if err := w.WriteData(o.Shutdown); err != nil {
		return err
	}
	if err := w.WriteData(o.AutoCreated); err != nil {
		return err
	}
	if err := w.WriteData(o.UseDatabase); err != nil {
		return err
	}
	if o.DataFile != "" {
		_ptr_pszDataFile := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.DataFile); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DataFile, _ptr_pszDataFile); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Masters != nil {
		_ptr_aipMasters := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Masters != nil {
				if err := o.Masters.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AddrArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Masters, _ptr_aipMasters); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SecureSecondaries); err != nil {
		return err
	}
	if err := w.WriteData(o.NotifyLevel); err != nil {
		return err
	}
	if o.Secondaries != nil {
		_ptr_aipSecondaries := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Secondaries != nil {
				if err := o.Secondaries.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AddrArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Secondaries, _ptr_aipSecondaries); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Notify != nil {
		_ptr_aipNotify := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Notify != nil {
				if err := o.Notify.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AddrArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Notify, _ptr_aipNotify); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.UseWINS); err != nil {
		return err
	}
	if err := w.WriteData(o.UseNBSTAT); err != nil {
		return err
	}
	if err := w.WriteData(o.Aging); err != nil {
		return err
	}
	if err := w.WriteData(o.NoRefreshInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.RefreshInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.AvailForScavengeTime); err != nil {
		return err
	}
	if o.ScavengeServers != nil {
		_ptr_aipScavengeServers := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ScavengeServers != nil {
				if err := o.ScavengeServers.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AddrArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ScavengeServers, _ptr_aipScavengeServers); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ForwarderTimeout); err != nil {
		return err
	}
	if err := w.WriteData(o.ForwarderSlave); err != nil {
		return err
	}
	if o.LocalMasters != nil {
		_ptr_aipLocalMasters := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.LocalMasters != nil {
				if err := o.LocalMasters.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AddrArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalMasters, _ptr_aipLocalMasters); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.DPFlags); err != nil {
		return err
	}
	if o.DPFQDN != "" {
		_ptr_pszDpFqdn := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.DPFQDN); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DPFQDN, _ptr_pszDpFqdn); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ZoneDN != "" {
		_ptr_pwszZoneDn := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ZoneDN); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ZoneDN, _ptr_pwszZoneDn); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LastSuccessfulSOACheck); err != nil {
		return err
	}
	if err := w.WriteData(o.LastSuccessfulXFR); err != nil {
		return err
	}
	if err := w.WriteData(o.QueuedForBackgroundLoad); err != nil {
		return err
	}
	if err := w.WriteData(o.BackgroundLoadInProgress); err != nil {
		return err
	}
	if !o.ReadOnlyZone {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LastXFRAttempt); err != nil {
		return err
	}
	if err := w.WriteData(o.LastXFRResult); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *ZoneInfoLonghorn) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	_ptr_pszZoneName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.ZoneName); err != nil {
			return err
		}
		return nil
	})
	_s_pszZoneName := func(ptr interface{}) { o.ZoneName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ZoneName, _s_pszZoneName, _ptr_pszZoneName); err != nil {
		return err
	}
	if err := w.ReadData(&o.ZoneType); err != nil {
		return err
	}
	if err := w.ReadData(&o.Reverse); err != nil {
		return err
	}
	if err := w.ReadData(&o.AllowUpdate); err != nil {
		return err
	}
	if err := w.ReadData(&o.Paused); err != nil {
		return err
	}
	if err := w.ReadData(&o.Shutdown); err != nil {
		return err
	}
	if err := w.ReadData(&o.AutoCreated); err != nil {
		return err
	}
	if err := w.ReadData(&o.UseDatabase); err != nil {
		return err
	}
	_ptr_pszDataFile := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.DataFile); err != nil {
			return err
		}
		return nil
	})
	_s_pszDataFile := func(ptr interface{}) { o.DataFile = *ptr.(*string) }
	if err := w.ReadPointer(&o.DataFile, _s_pszDataFile, _ptr_pszDataFile); err != nil {
		return err
	}
	_ptr_aipMasters := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Masters == nil {
			o.Masters = &AddrArray{}
		}
		if err := o.Masters.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipMasters := func(ptr interface{}) { o.Masters = *ptr.(**AddrArray) }
	if err := w.ReadPointer(&o.Masters, _s_aipMasters, _ptr_aipMasters); err != nil {
		return err
	}
	if err := w.ReadData(&o.SecureSecondaries); err != nil {
		return err
	}
	if err := w.ReadData(&o.NotifyLevel); err != nil {
		return err
	}
	_ptr_aipSecondaries := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Secondaries == nil {
			o.Secondaries = &AddrArray{}
		}
		if err := o.Secondaries.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipSecondaries := func(ptr interface{}) { o.Secondaries = *ptr.(**AddrArray) }
	if err := w.ReadPointer(&o.Secondaries, _s_aipSecondaries, _ptr_aipSecondaries); err != nil {
		return err
	}
	_ptr_aipNotify := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Notify == nil {
			o.Notify = &AddrArray{}
		}
		if err := o.Notify.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipNotify := func(ptr interface{}) { o.Notify = *ptr.(**AddrArray) }
	if err := w.ReadPointer(&o.Notify, _s_aipNotify, _ptr_aipNotify); err != nil {
		return err
	}
	if err := w.ReadData(&o.UseWINS); err != nil {
		return err
	}
	if err := w.ReadData(&o.UseNBSTAT); err != nil {
		return err
	}
	if err := w.ReadData(&o.Aging); err != nil {
		return err
	}
	if err := w.ReadData(&o.NoRefreshInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.RefreshInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.AvailForScavengeTime); err != nil {
		return err
	}
	_ptr_aipScavengeServers := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ScavengeServers == nil {
			o.ScavengeServers = &AddrArray{}
		}
		if err := o.ScavengeServers.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipScavengeServers := func(ptr interface{}) { o.ScavengeServers = *ptr.(**AddrArray) }
	if err := w.ReadPointer(&o.ScavengeServers, _s_aipScavengeServers, _ptr_aipScavengeServers); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwarderTimeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwarderSlave); err != nil {
		return err
	}
	_ptr_aipLocalMasters := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.LocalMasters == nil {
			o.LocalMasters = &AddrArray{}
		}
		if err := o.LocalMasters.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipLocalMasters := func(ptr interface{}) { o.LocalMasters = *ptr.(**AddrArray) }
	if err := w.ReadPointer(&o.LocalMasters, _s_aipLocalMasters, _ptr_aipLocalMasters); err != nil {
		return err
	}
	if err := w.ReadData(&o.DPFlags); err != nil {
		return err
	}
	_ptr_pszDpFqdn := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.DPFQDN); err != nil {
			return err
		}
		return nil
	})
	_s_pszDpFqdn := func(ptr interface{}) { o.DPFQDN = *ptr.(*string) }
	if err := w.ReadPointer(&o.DPFQDN, _s_pszDpFqdn, _ptr_pszDpFqdn); err != nil {
		return err
	}
	_ptr_pwszZoneDn := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ZoneDN); err != nil {
			return err
		}
		return nil
	})
	_s_pwszZoneDn := func(ptr interface{}) { o.ZoneDN = *ptr.(*string) }
	if err := w.ReadPointer(&o.ZoneDN, _s_pwszZoneDn, _ptr_pwszZoneDn); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastSuccessfulSOACheck); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastSuccessfulXFR); err != nil {
		return err
	}
	if err := w.ReadData(&o.QueuedForBackgroundLoad); err != nil {
		return err
	}
	if err := w.ReadData(&o.BackgroundLoadInProgress); err != nil {
		return err
	}
	var _bReadOnlyZone int32
	if err := w.ReadData(&_bReadOnlyZone); err != nil {
		return err
	}
	o.ReadOnlyZone = _bReadOnlyZone != 0
	if err := w.ReadData(&o.LastXFRAttempt); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastXFRResult); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// ZoneInfo structure represents DNS_RPC_ZONE_INFO RPC structure.
//
// The DNS_RPC_ZONE_INFO structure contains the detailed information about a zone present
// on the DNS server. The DNS server uses this structure to return information about
// a zone while responding to R_DnssrvQuery2 (section 3.1.4.7) method calls with operation
// type "ZoneInfo". There are different versions of DNS_RPC_ZONE_INFO. The DNS server
// MUST use the structures corresponding to the passed-in value for the dwClientVersion
// field in DNS Server Management Protocol method calls (section 3.1.4) in the following
// table, or if the method call does not specify the value of dwClientVersion, the DNS_RPC_ZONE_
// INFO_W2K version of the structure MUST be used.
//
//	+------------+---------------------------------------------------+
//	|            |                                                   |
//	|   VALUE    |                     STRUCTURE                     |
//	|            |                                                   |
//	+------------+---------------------------------------------------+
//	+------------+---------------------------------------------------+
//	| 0x00000000 | DNS_RPC_ZONE_INFO_W2K (section 2.2.5.2.4.1)       |
//	+------------+---------------------------------------------------+
//	| 0x00060000 | DNS_RPC_ ZONE_INFO_DOTNET (section 2.2.5.2.4.2)   |
//	+------------+---------------------------------------------------+
//	| 0x00070000 | DNS_RPC_ ZONE_INFO_LONGHORN (section 2.2.5.2.4.3) |
//	+------------+---------------------------------------------------+
type ZoneInfo struct {
	RPCStructureVersion      uint32     `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	_                        uint32     `idl:"name:dwReserved0"`
	ZoneName                 string     `idl:"name:pszZoneName;string" json:"zone_name"`
	ZoneType                 uint32     `idl:"name:dwZoneType" json:"zone_type"`
	Reverse                  uint32     `idl:"name:fReverse" json:"reverse"`
	AllowUpdate              uint32     `idl:"name:fAllowUpdate" json:"allow_update"`
	Paused                   uint32     `idl:"name:fPaused" json:"paused"`
	Shutdown                 uint32     `idl:"name:fShutdown" json:"shutdown"`
	AutoCreated              uint32     `idl:"name:fAutoCreated" json:"auto_created"`
	UseDatabase              uint32     `idl:"name:fUseDatabase" json:"use_database"`
	DataFile                 string     `idl:"name:pszDataFile;string" json:"data_file"`
	Masters                  *AddrArray `idl:"name:aipMasters" json:"masters"`
	SecureSecondaries        uint32     `idl:"name:fSecureSecondaries" json:"secure_secondaries"`
	NotifyLevel              uint32     `idl:"name:fNotifyLevel" json:"notify_level"`
	Secondaries              *AddrArray `idl:"name:aipSecondaries" json:"secondaries"`
	Notify                   *AddrArray `idl:"name:aipNotify" json:"notify"`
	UseWINS                  uint32     `idl:"name:fUseWins" json:"use_wins"`
	UseNBSTAT                uint32     `idl:"name:fUseNbstat" json:"use_nbstat"`
	Aging                    uint32     `idl:"name:fAging" json:"aging"`
	NoRefreshInterval        uint32     `idl:"name:dwNoRefreshInterval" json:"no_refresh_interval"`
	RefreshInterval          uint32     `idl:"name:dwRefreshInterval" json:"refresh_interval"`
	AvailForScavengeTime     uint32     `idl:"name:dwAvailForScavengeTime" json:"avail_for_scavenge_time"`
	ScavengeServers          *AddrArray `idl:"name:aipScavengeServers" json:"scavenge_servers"`
	ForwarderTimeout         uint32     `idl:"name:dwForwarderTimeout" json:"forwarder_timeout"`
	ForwarderSlave           uint32     `idl:"name:fForwarderSlave" json:"forwarder_slave"`
	LocalMasters             *AddrArray `idl:"name:aipLocalMasters" json:"local_masters"`
	DPFlags                  uint32     `idl:"name:dwDpFlags" json:"dp_flags"`
	DPFQDN                   string     `idl:"name:pszDpFqdn;string" json:"dp_fqdn"`
	ZoneDN                   string     `idl:"name:pwszZoneDn;string" json:"zone_dn"`
	LastSuccessfulSOACheck   uint32     `idl:"name:dwLastSuccessfulSoaCheck" json:"last_successful_soa_check"`
	LastSuccessfulXFR        uint32     `idl:"name:dwLastSuccessfulXfr" json:"last_successful_xfr"`
	QueuedForBackgroundLoad  uint32     `idl:"name:fQueuedForBackgroundLoad" json:"queued_for_background_load"`
	BackgroundLoadInProgress uint32     `idl:"name:fBackgroundLoadInProgress" json:"background_load_in_progress"`
	ReadOnlyZone             bool       `idl:"name:fReadOnlyZone" json:"read_only_zone"`
	LastXFRAttempt           uint32     `idl:"name:dwLastXfrAttempt" json:"last_xfr_attempt"`
	LastXFRResult            uint32     `idl:"name:dwLastXfrResult" json:"last_xfr_result"`
}

func (o *ZoneInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ZoneInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if o.ZoneName != "" {
		_ptr_pszZoneName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.ZoneName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ZoneName, _ptr_pszZoneName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ZoneType); err != nil {
		return err
	}
	if err := w.WriteData(o.Reverse); err != nil {
		return err
	}
	if err := w.WriteData(o.AllowUpdate); err != nil {
		return err
	}
	if err := w.WriteData(o.Paused); err != nil {
		return err
	}
	if err := w.WriteData(o.Shutdown); err != nil {
		return err
	}
	if err := w.WriteData(o.AutoCreated); err != nil {
		return err
	}
	if err := w.WriteData(o.UseDatabase); err != nil {
		return err
	}
	if o.DataFile != "" {
		_ptr_pszDataFile := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.DataFile); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DataFile, _ptr_pszDataFile); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Masters != nil {
		_ptr_aipMasters := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Masters != nil {
				if err := o.Masters.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AddrArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Masters, _ptr_aipMasters); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SecureSecondaries); err != nil {
		return err
	}
	if err := w.WriteData(o.NotifyLevel); err != nil {
		return err
	}
	if o.Secondaries != nil {
		_ptr_aipSecondaries := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Secondaries != nil {
				if err := o.Secondaries.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AddrArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Secondaries, _ptr_aipSecondaries); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Notify != nil {
		_ptr_aipNotify := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Notify != nil {
				if err := o.Notify.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AddrArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Notify, _ptr_aipNotify); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.UseWINS); err != nil {
		return err
	}
	if err := w.WriteData(o.UseNBSTAT); err != nil {
		return err
	}
	if err := w.WriteData(o.Aging); err != nil {
		return err
	}
	if err := w.WriteData(o.NoRefreshInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.RefreshInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.AvailForScavengeTime); err != nil {
		return err
	}
	if o.ScavengeServers != nil {
		_ptr_aipScavengeServers := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ScavengeServers != nil {
				if err := o.ScavengeServers.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AddrArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ScavengeServers, _ptr_aipScavengeServers); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ForwarderTimeout); err != nil {
		return err
	}
	if err := w.WriteData(o.ForwarderSlave); err != nil {
		return err
	}
	if o.LocalMasters != nil {
		_ptr_aipLocalMasters := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.LocalMasters != nil {
				if err := o.LocalMasters.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AddrArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalMasters, _ptr_aipLocalMasters); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.DPFlags); err != nil {
		return err
	}
	if o.DPFQDN != "" {
		_ptr_pszDpFqdn := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.DPFQDN); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DPFQDN, _ptr_pszDpFqdn); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ZoneDN != "" {
		_ptr_pwszZoneDn := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ZoneDN); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ZoneDN, _ptr_pwszZoneDn); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LastSuccessfulSOACheck); err != nil {
		return err
	}
	if err := w.WriteData(o.LastSuccessfulXFR); err != nil {
		return err
	}
	if err := w.WriteData(o.QueuedForBackgroundLoad); err != nil {
		return err
	}
	if err := w.WriteData(o.BackgroundLoadInProgress); err != nil {
		return err
	}
	if !o.ReadOnlyZone {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LastXFRAttempt); err != nil {
		return err
	}
	if err := w.WriteData(o.LastXFRResult); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *ZoneInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	_ptr_pszZoneName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.ZoneName); err != nil {
			return err
		}
		return nil
	})
	_s_pszZoneName := func(ptr interface{}) { o.ZoneName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ZoneName, _s_pszZoneName, _ptr_pszZoneName); err != nil {
		return err
	}
	if err := w.ReadData(&o.ZoneType); err != nil {
		return err
	}
	if err := w.ReadData(&o.Reverse); err != nil {
		return err
	}
	if err := w.ReadData(&o.AllowUpdate); err != nil {
		return err
	}
	if err := w.ReadData(&o.Paused); err != nil {
		return err
	}
	if err := w.ReadData(&o.Shutdown); err != nil {
		return err
	}
	if err := w.ReadData(&o.AutoCreated); err != nil {
		return err
	}
	if err := w.ReadData(&o.UseDatabase); err != nil {
		return err
	}
	_ptr_pszDataFile := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.DataFile); err != nil {
			return err
		}
		return nil
	})
	_s_pszDataFile := func(ptr interface{}) { o.DataFile = *ptr.(*string) }
	if err := w.ReadPointer(&o.DataFile, _s_pszDataFile, _ptr_pszDataFile); err != nil {
		return err
	}
	_ptr_aipMasters := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Masters == nil {
			o.Masters = &AddrArray{}
		}
		if err := o.Masters.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipMasters := func(ptr interface{}) { o.Masters = *ptr.(**AddrArray) }
	if err := w.ReadPointer(&o.Masters, _s_aipMasters, _ptr_aipMasters); err != nil {
		return err
	}
	if err := w.ReadData(&o.SecureSecondaries); err != nil {
		return err
	}
	if err := w.ReadData(&o.NotifyLevel); err != nil {
		return err
	}
	_ptr_aipSecondaries := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Secondaries == nil {
			o.Secondaries = &AddrArray{}
		}
		if err := o.Secondaries.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipSecondaries := func(ptr interface{}) { o.Secondaries = *ptr.(**AddrArray) }
	if err := w.ReadPointer(&o.Secondaries, _s_aipSecondaries, _ptr_aipSecondaries); err != nil {
		return err
	}
	_ptr_aipNotify := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Notify == nil {
			o.Notify = &AddrArray{}
		}
		if err := o.Notify.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipNotify := func(ptr interface{}) { o.Notify = *ptr.(**AddrArray) }
	if err := w.ReadPointer(&o.Notify, _s_aipNotify, _ptr_aipNotify); err != nil {
		return err
	}
	if err := w.ReadData(&o.UseWINS); err != nil {
		return err
	}
	if err := w.ReadData(&o.UseNBSTAT); err != nil {
		return err
	}
	if err := w.ReadData(&o.Aging); err != nil {
		return err
	}
	if err := w.ReadData(&o.NoRefreshInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.RefreshInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.AvailForScavengeTime); err != nil {
		return err
	}
	_ptr_aipScavengeServers := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ScavengeServers == nil {
			o.ScavengeServers = &AddrArray{}
		}
		if err := o.ScavengeServers.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipScavengeServers := func(ptr interface{}) { o.ScavengeServers = *ptr.(**AddrArray) }
	if err := w.ReadPointer(&o.ScavengeServers, _s_aipScavengeServers, _ptr_aipScavengeServers); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwarderTimeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwarderSlave); err != nil {
		return err
	}
	_ptr_aipLocalMasters := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.LocalMasters == nil {
			o.LocalMasters = &AddrArray{}
		}
		if err := o.LocalMasters.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipLocalMasters := func(ptr interface{}) { o.LocalMasters = *ptr.(**AddrArray) }
	if err := w.ReadPointer(&o.LocalMasters, _s_aipLocalMasters, _ptr_aipLocalMasters); err != nil {
		return err
	}
	if err := w.ReadData(&o.DPFlags); err != nil {
		return err
	}
	_ptr_pszDpFqdn := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.DPFQDN); err != nil {
			return err
		}
		return nil
	})
	_s_pszDpFqdn := func(ptr interface{}) { o.DPFQDN = *ptr.(*string) }
	if err := w.ReadPointer(&o.DPFQDN, _s_pszDpFqdn, _ptr_pszDpFqdn); err != nil {
		return err
	}
	_ptr_pwszZoneDn := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ZoneDN); err != nil {
			return err
		}
		return nil
	})
	_s_pwszZoneDn := func(ptr interface{}) { o.ZoneDN = *ptr.(*string) }
	if err := w.ReadPointer(&o.ZoneDN, _s_pwszZoneDn, _ptr_pwszZoneDn); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastSuccessfulSOACheck); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastSuccessfulXFR); err != nil {
		return err
	}
	if err := w.ReadData(&o.QueuedForBackgroundLoad); err != nil {
		return err
	}
	if err := w.ReadData(&o.BackgroundLoadInProgress); err != nil {
		return err
	}
	var _bReadOnlyZone int32
	if err := w.ReadData(&_bReadOnlyZone); err != nil {
		return err
	}
	o.ReadOnlyZone = _bReadOnlyZone != 0
	if err := w.ReadData(&o.LastXFRAttempt); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastXFRResult); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// ZoneCreateInfoW2K structure represents DNS_RPC_ZONE_CREATE_INFO_W2K RPC structure.
//
// The DNS_RPC_ZONE_CREATE_INFO_W2K structure is used to specify parameters required
// when creating a new DNS zone.
type ZoneCreateInfoW2K struct {
	// pszZoneName: As specified in section 2.2.5.2.4.1.
	ZoneName string `idl:"name:pszZoneName;string" json:"zone_name"`
	// dwZoneType: The zone type. This MUST be set to one of the allowed values specified
	// in DNS_ZONE_TYPE (section 2.2.5.1.1), and it MUST NOT be either DNS_ZONE_TYPE_CACHE
	// or DNS_ZONE_TYPE_SECONDARY_CACHE.
	ZoneType uint32 `idl:"name:dwZoneType" json:"zone_type"`
	// fAllowUpdate: As specified in section 2.2.5.2.4.1.
	AllowUpdate uint32 `idl:"name:fAllowUpdate" json:"allow_update"`
	// fAging: As specified in section 2.2.5.2.4.1.
	Aging uint32 `idl:"name:fAging" json:"aging"`
	// dwFlags: The zone creation behavior that the DNS server is to follow while creating
	// the zone. This field is only used when the operation is ZoneCreate. The DNS server
	// MUST ignore the value of this field when the operation is ZoneTypeReset. This field
	// MUST be set to any combination of the following values:
	//
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                               |                                                                                  |
	//	|                     VALUE                     |                                     MEANING                                      |
	//	|                                               |                                                                                  |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_ZONE_LOAD_OVERWRITE_MEMORY 0x00000010     | If dwZoneType is not set to the value DNS_ZONE_TYPE_PRIMARY (section 2.2.5.1.1), |
	//	|                                               | then this flag MUST be ignored. Otherwise, the DNS server MUST attempt to find   |
	//	|                                               | and load the zone database from persistent storage instead of creating a new     |
	//	|                                               | empty zone database. If the value of fDsIntegrated is 0x00000001 then the DNS    |
	//	|                                               | server MUST search for a pre-existing zone database in the directory server,     |
	//	|                                               | otherwise the DNS server MUST search for a pre-existing zone database in a       |
	//	|                                               | file. If a pre-existing zone database is not found then it continues with zone   |
	//	|                                               | creation, however if a pre-existing zone database is found but could not be      |
	//	|                                               | loaded, the DNS server MUST fail the operation and return an error.              |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_ZONE_CREATE_FOR_DCPROMO 0x00001000        | If this flag is set, the DNS server MUST create the zone such that it is         |
	//	|                                               | directory server-integrated and stored in the DNS domain partition.              |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_ZONE_CREATE_FOR_DCPROMO_FOREST 0x00004000 | If this flag is set, the DNS server MUST create the zone such that it is         |
	//	|                                               | directory server-integrated and stored in the DNS forest partition.              |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
	// pszDataFile: As specified in section 2.2.5.2.4.1.
	DataFile string `idl:"name:pszDataFile;string" json:"data_file"`
	// fDsIntegrated: A value of 0x00000001 indicates that the zone is to be created to
	// use the directory server for persistent storage, and 0x00000000 indicates it is not.
	// If this is set to 0x00000001 the caller MUST specify the application directory partition
	// information in pszDpFqdn (section 2.2.5.2.7.2); in this case the DNS server MUST
	// ignore the value of pszDataFile.
	DSIntegrated uint32 `idl:"name:fDsIntegrated" json:"ds_integrated"`
	// fLoadExisting: If the operation is ZoneCreate this field is interpreted as a Boolean
	// value. If set to TRUE this has the same effect as specifying DNS_ZONE_LOAD_OVERWRITE_MEMORY
	// in the dwFlags field. If the operation is ZoneTypeReset, this field can be set to
	// one of the following values; however, the DNS server MUST ignore the value of this
	// field if fDsIntegrated is not TRUE or dwZoneType is not DNS_ZONE_TYPE_PRIMARY (section
	// 2.2.5.1.1).
	//
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                           |                                                                                  |
	//	|                   VALUE                   |                                     MEANING                                      |
	//	|                                           |                                                                                  |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_ZONE_LOAD_OVERWRITE_MEMORY 0x00000010 | The server MUST attempt to find and load the zone database from persistent       |
	//	|                                           | storage instead of retaining the existing in-memory zone database by searching   |
	//	|                                           | for a pre-existing zone database in the directory server. If a pre-existing zone |
	//	|                                           | database is not found, then the server MUST fail the operation and return an     |
	//	|                                           | error.                                                                           |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_ZONE_LOAD_OVERWRITE_DS 0x00000020     | If this flag is set and if the zone already exists in the database, then the     |
	//	|                                           | server MUST overwrite the existing zone database with current in-memory zone     |
	//	|                                           | database.                                                                        |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	LoadExisting uint32 `idl:"name:fLoadExisting" json:"load_existing"`
	// pszAdmin: A pointer to a null-terminated UTF-8 string containing the administrator's
	// email name (in the format specified in [RFC1035] section 8) or NULL to cause the
	// DNS server to use a default value "hostmaster", followed by the name of the zone.
	// This value MUST be used to populate the zone administrator email field in the SOA
	// record in the new zone.
	Admin string `idl:"name:pszAdmin;string" json:"admin"`
	// aipMasters: As specified in section 2.2.5.2.4.1.
	Masters *IPv4Array `idl:"name:aipMasters" json:"masters"`
	// aipSecondaries: As specified in section 2.2.5.2.4.1.
	Secondaries *IPv4Array `idl:"name:aipSecondaries" json:"secondaries"`
	// fSecureSecondaries: As specified in section 2.2.5.2.4.1.
	SecureSecondaries uint32 `idl:"name:fSecureSecondaries" json:"secure_secondaries"`
	// fNotifyLevel: As specified in section 2.2.5.2.4.1.
	NotifyLevel uint32 `idl:"name:fNotifyLevel" json:"notify_level"`
	// pvReserved1: MUST be set to zero when sent and MUST be ignored on receipt.
	_ string `idl:"name:pvReserved1;string"`
	// pvReserved2: MUST be set to zero when sent and MUST be ignored on receipt.
	_ string `idl:"name:pvReserved2;string"`
	// pvReserved3:MUST be set to zero when sent and MUST be ignored on receipt.
	_ string `idl:"name:pvReserved3;string"`
	// pvReserved4: MUST be set to zero when sent and MUST be ignored on receipt.
	_ string `idl:"name:pvReserved4;string"`
	// pvReserved5: MUST be set to zero when sent and MUST be ignored on receipt.
	_ string `idl:"name:pvReserved5;string"`
	// pvReserved6: MUST be set to zero when sent and MUST be ignored on receipt.
	_ string `idl:"name:pvReserved6;string"`
	// pvReserved7: MUST be set to zero when sent and MUST be ignored on receipt.
	_ string `idl:"name:pvReserved7;string"`
	// pvReserved8: MUST be set to zero when sent and MUST be ignored on receipt.
	_ string `idl:"name:pvReserved8;string"`
	// dwReserved1: MUST be set to zero when sent and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved1"`
	// dwReserved2: MUST be set to zero when sent and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved2"`
	// dwReserved3: MUST be set to zero when sent and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved3"`
	// dwReserved4: MUST be set to zero when sent and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved4"`
	// dwReserved5: MUST be set to zero when sent and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved5"`
	// dwReserved6: MUST be set to zero when sent and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved6"`
	// dwReserved7: MUST be set to zero when sent and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved7"`
	// dwReserved8: MUST be set to zero when sent and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved8"`
}

func (o *ZoneCreateInfoW2K) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ZoneCreateInfoW2K) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.ZoneName != "" {
		_ptr_pszZoneName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.ZoneName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ZoneName, _ptr_pszZoneName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ZoneType); err != nil {
		return err
	}
	if err := w.WriteData(o.AllowUpdate); err != nil {
		return err
	}
	if err := w.WriteData(o.Aging); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.DataFile != "" {
		_ptr_pszDataFile := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.DataFile); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DataFile, _ptr_pszDataFile); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.DSIntegrated); err != nil {
		return err
	}
	if err := w.WriteData(o.LoadExisting); err != nil {
		return err
	}
	if o.Admin != "" {
		_ptr_pszAdmin := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.Admin); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Admin, _ptr_pszAdmin); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Masters != nil {
		_ptr_aipMasters := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Masters != nil {
				if err := o.Masters.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPv4Array{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Masters, _ptr_aipMasters); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Secondaries != nil {
		_ptr_aipSecondaries := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Secondaries != nil {
				if err := o.Secondaries.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPv4Array{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Secondaries, _ptr_aipSecondaries); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SecureSecondaries); err != nil {
		return err
	}
	if err := w.WriteData(o.NotifyLevel); err != nil {
		return err
	}
	// reserved pvReserved1
	if err := w.WritePointer(nil); err != nil {
		return err
	}
	// reserved pvReserved2
	if err := w.WritePointer(nil); err != nil {
		return err
	}
	// reserved pvReserved3
	if err := w.WritePointer(nil); err != nil {
		return err
	}
	// reserved pvReserved4
	if err := w.WritePointer(nil); err != nil {
		return err
	}
	// reserved pvReserved5
	if err := w.WritePointer(nil); err != nil {
		return err
	}
	// reserved pvReserved6
	if err := w.WritePointer(nil); err != nil {
		return err
	}
	// reserved pvReserved7
	if err := w.WritePointer(nil); err != nil {
		return err
	}
	// reserved pvReserved8
	if err := w.WritePointer(nil); err != nil {
		return err
	}
	// reserved dwReserved1
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	// reserved dwReserved2
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	// reserved dwReserved3
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	// reserved dwReserved4
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	// reserved dwReserved5
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	// reserved dwReserved6
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	// reserved dwReserved7
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	// reserved dwReserved8
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *ZoneCreateInfoW2K) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_pszZoneName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.ZoneName); err != nil {
			return err
		}
		return nil
	})
	_s_pszZoneName := func(ptr interface{}) { o.ZoneName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ZoneName, _s_pszZoneName, _ptr_pszZoneName); err != nil {
		return err
	}
	if err := w.ReadData(&o.ZoneType); err != nil {
		return err
	}
	if err := w.ReadData(&o.AllowUpdate); err != nil {
		return err
	}
	if err := w.ReadData(&o.Aging); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_pszDataFile := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.DataFile); err != nil {
			return err
		}
		return nil
	})
	_s_pszDataFile := func(ptr interface{}) { o.DataFile = *ptr.(*string) }
	if err := w.ReadPointer(&o.DataFile, _s_pszDataFile, _ptr_pszDataFile); err != nil {
		return err
	}
	if err := w.ReadData(&o.DSIntegrated); err != nil {
		return err
	}
	if err := w.ReadData(&o.LoadExisting); err != nil {
		return err
	}
	_ptr_pszAdmin := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.Admin); err != nil {
			return err
		}
		return nil
	})
	_s_pszAdmin := func(ptr interface{}) { o.Admin = *ptr.(*string) }
	if err := w.ReadPointer(&o.Admin, _s_pszAdmin, _ptr_pszAdmin); err != nil {
		return err
	}
	_ptr_aipMasters := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Masters == nil {
			o.Masters = &IPv4Array{}
		}
		if err := o.Masters.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipMasters := func(ptr interface{}) { o.Masters = *ptr.(**IPv4Array) }
	if err := w.ReadPointer(&o.Masters, _s_aipMasters, _ptr_aipMasters); err != nil {
		return err
	}
	_ptr_aipSecondaries := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Secondaries == nil {
			o.Secondaries = &IPv4Array{}
		}
		if err := o.Secondaries.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipSecondaries := func(ptr interface{}) { o.Secondaries = *ptr.(**IPv4Array) }
	if err := w.ReadPointer(&o.Secondaries, _s_aipSecondaries, _ptr_aipSecondaries); err != nil {
		return err
	}
	if err := w.ReadData(&o.SecureSecondaries); err != nil {
		return err
	}
	if err := w.ReadData(&o.NotifyLevel); err != nil {
		return err
	}
	// reserved pvReserved1
	var _pvReserved1 string
	_ptr_pvReserved1 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &_pvReserved1); err != nil {
			return err
		}
		return nil
	})
	_s_pvReserved1 := func(ptr interface{}) { _pvReserved1 = *ptr.(*string) }
	if err := w.ReadPointer(&_pvReserved1, _s_pvReserved1, _ptr_pvReserved1); err != nil {
		return err
	}
	// reserved pvReserved2
	var _pvReserved2 string
	_ptr_pvReserved2 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &_pvReserved2); err != nil {
			return err
		}
		return nil
	})
	_s_pvReserved2 := func(ptr interface{}) { _pvReserved2 = *ptr.(*string) }
	if err := w.ReadPointer(&_pvReserved2, _s_pvReserved2, _ptr_pvReserved2); err != nil {
		return err
	}
	// reserved pvReserved3
	var _pvReserved3 string
	_ptr_pvReserved3 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &_pvReserved3); err != nil {
			return err
		}
		return nil
	})
	_s_pvReserved3 := func(ptr interface{}) { _pvReserved3 = *ptr.(*string) }
	if err := w.ReadPointer(&_pvReserved3, _s_pvReserved3, _ptr_pvReserved3); err != nil {
		return err
	}
	// reserved pvReserved4
	var _pvReserved4 string
	_ptr_pvReserved4 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &_pvReserved4); err != nil {
			return err
		}
		return nil
	})
	_s_pvReserved4 := func(ptr interface{}) { _pvReserved4 = *ptr.(*string) }
	if err := w.ReadPointer(&_pvReserved4, _s_pvReserved4, _ptr_pvReserved4); err != nil {
		return err
	}
	// reserved pvReserved5
	var _pvReserved5 string
	_ptr_pvReserved5 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &_pvReserved5); err != nil {
			return err
		}
		return nil
	})
	_s_pvReserved5 := func(ptr interface{}) { _pvReserved5 = *ptr.(*string) }
	if err := w.ReadPointer(&_pvReserved5, _s_pvReserved5, _ptr_pvReserved5); err != nil {
		return err
	}
	// reserved pvReserved6
	var _pvReserved6 string
	_ptr_pvReserved6 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &_pvReserved6); err != nil {
			return err
		}
		return nil
	})
	_s_pvReserved6 := func(ptr interface{}) { _pvReserved6 = *ptr.(*string) }
	if err := w.ReadPointer(&_pvReserved6, _s_pvReserved6, _ptr_pvReserved6); err != nil {
		return err
	}
	// reserved pvReserved7
	var _pvReserved7 string
	_ptr_pvReserved7 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &_pvReserved7); err != nil {
			return err
		}
		return nil
	})
	_s_pvReserved7 := func(ptr interface{}) { _pvReserved7 = *ptr.(*string) }
	if err := w.ReadPointer(&_pvReserved7, _s_pvReserved7, _ptr_pvReserved7); err != nil {
		return err
	}
	// reserved pvReserved8
	var _pvReserved8 string
	_ptr_pvReserved8 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &_pvReserved8); err != nil {
			return err
		}
		return nil
	})
	_s_pvReserved8 := func(ptr interface{}) { _pvReserved8 = *ptr.(*string) }
	if err := w.ReadPointer(&_pvReserved8, _s_pvReserved8, _ptr_pvReserved8); err != nil {
		return err
	}
	// reserved dwReserved1
	var _dwReserved1 uint32
	if err := w.ReadData(&_dwReserved1); err != nil {
		return err
	}
	// reserved dwReserved2
	var _dwReserved2 uint32
	if err := w.ReadData(&_dwReserved2); err != nil {
		return err
	}
	// reserved dwReserved3
	var _dwReserved3 uint32
	if err := w.ReadData(&_dwReserved3); err != nil {
		return err
	}
	// reserved dwReserved4
	var _dwReserved4 uint32
	if err := w.ReadData(&_dwReserved4); err != nil {
		return err
	}
	// reserved dwReserved5
	var _dwReserved5 uint32
	if err := w.ReadData(&_dwReserved5); err != nil {
		return err
	}
	// reserved dwReserved6
	var _dwReserved6 uint32
	if err := w.ReadData(&_dwReserved6); err != nil {
		return err
	}
	// reserved dwReserved7
	var _dwReserved7 uint32
	if err := w.ReadData(&_dwReserved7); err != nil {
		return err
	}
	// reserved dwReserved8
	var _dwReserved8 uint32
	if err := w.ReadData(&_dwReserved8); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// ZoneCreateInfo_NET structure represents DNS_RPC_ZONE_CREATE_INFO_DOTNET RPC structure.
//
// In the DNS_RPC_ZONE_CREATE_INFO_DOTNET structure all fields have same definition
// as specified in section DNS_RPC_ZONE_CREATE_INFO_W2K (section 2.2.5.2.7.1), with
// the following exceptions:
type ZoneCreateInfo_NET struct {
	// dwRpcStructureVersion: As specified in section 2.2.5.2.4.2.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// dwReserved0: As specified in section 2.2.5.2.4.2.
	_                 uint32     `idl:"name:dwReserved0"`
	ZoneName          string     `idl:"name:pszZoneName;string" json:"zone_name"`
	ZoneType          uint32     `idl:"name:dwZoneType" json:"zone_type"`
	AllowUpdate       uint32     `idl:"name:fAllowUpdate" json:"allow_update"`
	Aging             uint32     `idl:"name:fAging" json:"aging"`
	Flags             uint32     `idl:"name:dwFlags" json:"flags"`
	DataFile          string     `idl:"name:pszDataFile;string" json:"data_file"`
	DSIntegrated      uint32     `idl:"name:fDsIntegrated" json:"ds_integrated"`
	LoadExisting      uint32     `idl:"name:fLoadExisting" json:"load_existing"`
	Admin             string     `idl:"name:pszAdmin;string" json:"admin"`
	Masters           *IPv4Array `idl:"name:aipMasters" json:"masters"`
	Secondaries       *IPv4Array `idl:"name:aipSecondaries" json:"secondaries"`
	SecureSecondaries uint32     `idl:"name:fSecureSecondaries" json:"secure_secondaries"`
	NotifyLevel       uint32     `idl:"name:fNotifyLevel" json:"notify_level"`
	// dwTimeout: Equivalent to dwForwarderTimeout specified in section 2.2.5.2.4.2.
	Timeout uint32 `idl:"name:dwTimeout" json:"timeout"`
	// fRecurseAfterForwarding: Equivalent to fForwarderSlave specified in section 2.2.5.2.4.2.
	RecurseAfterForwarding uint32 `idl:"name:fRecurseAfterForwarding" json:"recurse_after_forwarding"`
	// dwDpFlags: As specified in section 2.2.5.2.4.2. However, only the following values
	// can be used and of these values more than one MUST NOT be specified: DNS_DP_LEGACY,
	// DNS_DP_DOMAIN_DEFAULT, and DNS_DP_FOREST_DEFAULT. This field is set to zero if the
	// zone is not to be created to use the directory server for persistent storage. The
	// DNS server returns an error if any value not listed above is specified or if more
	// than one of the allowable values are specified.
	DPFlags uint32 `idl:"name:dwDpFlags" json:"dp_flags"`
	// pszDpFqdn: As specified in section 2.2.5.2.4.2.
	DPFQDN string `idl:"name:pszDpFqdn;string" json:"dp_fqdn"`
	// dwReserved: MUST be set to zero when sent and MUST be ignored on receipt.
	//
	// If the DNS RPC client sends an older version of DNS_RPC_ZONE_CREATE_INFO structure
	// such as DNS_RPC_ZONE_CREATE_INFO_W2K (section 2.2.5.2.7.1), then the DNS RPC server
	// MUST construct a current version of DNS_RPC_ZONE_CREATE_INFO structure such as DNS_RPC_ZONE_CREATE_INFO_DOTNET,
	// using steps as specified below:
	//
	// *
	//
	// Copy the same value for fields that are common to input and current version of DNS_RPC_ZONE_CREATE_INFO
	// structures.
	//
	// *
	//
	// dwRpcStructureVersion field MUST be set to 1.
	//
	// *
	//
	// All other fields that are defined only in DNS_RPC_ZONE_CREATE_INFO_DOTNET and are
	// not defined in DNS_RPC_ZONE_CREATE_INFO_W2K (section 2.2.5.2.7.1), MUST be set to
	// 0.
	_ []uint32 `idl:"name:dwReserved"`
}

func (o *ZoneCreateInfo_NET) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ZoneCreateInfo_NET) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if o.ZoneName != "" {
		_ptr_pszZoneName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.ZoneName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ZoneName, _ptr_pszZoneName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ZoneType); err != nil {
		return err
	}
	if err := w.WriteData(o.AllowUpdate); err != nil {
		return err
	}
	if err := w.WriteData(o.Aging); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.DataFile != "" {
		_ptr_pszDataFile := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.DataFile); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DataFile, _ptr_pszDataFile); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.DSIntegrated); err != nil {
		return err
	}
	if err := w.WriteData(o.LoadExisting); err != nil {
		return err
	}
	if o.Admin != "" {
		_ptr_pszAdmin := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.Admin); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Admin, _ptr_pszAdmin); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Masters != nil {
		_ptr_aipMasters := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Masters != nil {
				if err := o.Masters.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPv4Array{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Masters, _ptr_aipMasters); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Secondaries != nil {
		_ptr_aipSecondaries := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Secondaries != nil {
				if err := o.Secondaries.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPv4Array{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Secondaries, _ptr_aipSecondaries); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SecureSecondaries); err != nil {
		return err
	}
	if err := w.WriteData(o.NotifyLevel); err != nil {
		return err
	}
	if err := w.WriteData(o.Timeout); err != nil {
		return err
	}
	if err := w.WriteData(o.RecurseAfterForwarding); err != nil {
		return err
	}
	if err := w.WriteData(o.DPFlags); err != nil {
		return err
	}
	if o.DPFQDN != "" {
		_ptr_pszDpFqdn := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.DPFQDN); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DPFQDN, _ptr_pszDpFqdn); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	// reserved dwReserved
	for i1 := 0; uint64(i1) < 32; i1++ {
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *ZoneCreateInfo_NET) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	_ptr_pszZoneName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.ZoneName); err != nil {
			return err
		}
		return nil
	})
	_s_pszZoneName := func(ptr interface{}) { o.ZoneName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ZoneName, _s_pszZoneName, _ptr_pszZoneName); err != nil {
		return err
	}
	if err := w.ReadData(&o.ZoneType); err != nil {
		return err
	}
	if err := w.ReadData(&o.AllowUpdate); err != nil {
		return err
	}
	if err := w.ReadData(&o.Aging); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_pszDataFile := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.DataFile); err != nil {
			return err
		}
		return nil
	})
	_s_pszDataFile := func(ptr interface{}) { o.DataFile = *ptr.(*string) }
	if err := w.ReadPointer(&o.DataFile, _s_pszDataFile, _ptr_pszDataFile); err != nil {
		return err
	}
	if err := w.ReadData(&o.DSIntegrated); err != nil {
		return err
	}
	if err := w.ReadData(&o.LoadExisting); err != nil {
		return err
	}
	_ptr_pszAdmin := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.Admin); err != nil {
			return err
		}
		return nil
	})
	_s_pszAdmin := func(ptr interface{}) { o.Admin = *ptr.(*string) }
	if err := w.ReadPointer(&o.Admin, _s_pszAdmin, _ptr_pszAdmin); err != nil {
		return err
	}
	_ptr_aipMasters := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Masters == nil {
			o.Masters = &IPv4Array{}
		}
		if err := o.Masters.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipMasters := func(ptr interface{}) { o.Masters = *ptr.(**IPv4Array) }
	if err := w.ReadPointer(&o.Masters, _s_aipMasters, _ptr_aipMasters); err != nil {
		return err
	}
	_ptr_aipSecondaries := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Secondaries == nil {
			o.Secondaries = &IPv4Array{}
		}
		if err := o.Secondaries.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipSecondaries := func(ptr interface{}) { o.Secondaries = *ptr.(**IPv4Array) }
	if err := w.ReadPointer(&o.Secondaries, _s_aipSecondaries, _ptr_aipSecondaries); err != nil {
		return err
	}
	if err := w.ReadData(&o.SecureSecondaries); err != nil {
		return err
	}
	if err := w.ReadData(&o.NotifyLevel); err != nil {
		return err
	}
	if err := w.ReadData(&o.Timeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.RecurseAfterForwarding); err != nil {
		return err
	}
	if err := w.ReadData(&o.DPFlags); err != nil {
		return err
	}
	_ptr_pszDpFqdn := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.DPFQDN); err != nil {
			return err
		}
		return nil
	})
	_s_pszDpFqdn := func(ptr interface{}) { o.DPFQDN = *ptr.(*string) }
	if err := w.ReadPointer(&o.DPFQDN, _s_pszDpFqdn, _ptr_pszDpFqdn); err != nil {
		return err
	}
	// reserved dwReserved
	var _dwReserved []uint32
	_dwReserved = make([]uint32, 32)
	for i1 := range _dwReserved {
		i1 := i1
		if err := w.ReadData(&_dwReserved[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// ZoneCreateInfoLonghorn structure represents DNS_RPC_ZONE_CREATE_INFO_LONGHORN RPC structure.
//
// In the DNS_RPC_ZONE_CREATE_INFO_LONGHORN structure all fields have same definition
// as specified in section DNS_RPC_ZONE_CREATE_INFO_DOTNET (section 2.2.5.2.7.2), with
// the following exceptions:
type ZoneCreateInfoLonghorn struct {
	// dwRpcStructureVersion: As specified in section 2.2.5.2.4.3.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	_                   uint32 `idl:"name:dwReserved0"`
	ZoneName            string `idl:"name:pszZoneName;string" json:"zone_name"`
	ZoneType            uint32 `idl:"name:dwZoneType" json:"zone_type"`
	AllowUpdate         uint32 `idl:"name:fAllowUpdate" json:"allow_update"`
	Aging               uint32 `idl:"name:fAging" json:"aging"`
	Flags               uint32 `idl:"name:dwFlags" json:"flags"`
	DataFile            string `idl:"name:pszDataFile;string" json:"data_file"`
	DSIntegrated        uint32 `idl:"name:fDsIntegrated" json:"ds_integrated"`
	LoadExisting        uint32 `idl:"name:fLoadExisting" json:"load_existing"`
	Admin               string `idl:"name:pszAdmin;string" json:"admin"`
	// aipMasters: As specified in section 2.2.5.2.4.3.
	Masters *AddrArray `idl:"name:aipMasters" json:"masters"`
	// aipSecondaries: As specified in section 2.2.5.2.4.3.
	//
	// If the DNS RPC client sends an older version of DNS_RPC_ZONE_CREATE_INFO structure
	// such as DNS_RPC_ZONE_CREATE_INFO_W2K (section 2.2.5.2.7.1) or DNS_RPC_ZONE_CREATE_INFO_DOTNET
	// (section 2.2.5.2.7.2), then the DNS RPC server MUST construct a current version of
	// DNS_RPC_ZONE_CREATE_INFO structure such as DNS_RPC_ZONE_CREATE_INFO_LONGHORN, using
	// steps as specified below:
	//
	// *
	//
	// Copy the same value for fields that are common to input and current version of DNS_RPC_ZONE_CREATE_INFO
	// structures.
	//
	// *
	//
	// dwRpcStructureVersion field MUST be set to 2.
	//
	// *
	//
	// The values for aipMasters and aipSecondaries fields MUST be obtained from input structure
	// as IP4_ARRAY type and MUST be converted to DNS_ADDR_ARRAY type, and then assigned
	// to aipMasters and aipSecondaries fields in the DNS_RPC_ZONE_CREATE_INFO_LONGHORN
	// structure.
	//
	// *Note* DNS_RPC_ZONE_CREATE_INFO_W2K (section 2.2.5.2.7.1) and DNS_RPC_ZONE_CREATE_INFO_DOTNET
	// (section 2.2.5.2.7.2), do not support IPv6 ( a95b05da-f1fd-4db3-94b4-817fdaa1f642#gt_64c29bb6-c8b2-4281-9f3a-c1eb5d2288aa
	// ) address list of *aipMasters* and *aipSecondaries*.
	//
	// *
	//
	// All other fields that are defined only in DNS_RPC_ZONE_CREATE_INFO_LONGHORN and are
	// not defined in DNS_RPC_ZONE_CREATE_INFO_W2K (section 2.2.5.2.7.1) or DNS_RPC_ZONE_CREATE_INFO_DOTNET
	// (section 2.2.5.2.7.2) structure MUST be set to 0.
	Secondaries            *AddrArray `idl:"name:aipSecondaries" json:"secondaries"`
	SecureSecondaries      uint32     `idl:"name:fSecureSecondaries" json:"secure_secondaries"`
	NotifyLevel            uint32     `idl:"name:fNotifyLevel" json:"notify_level"`
	Timeout                uint32     `idl:"name:dwTimeout" json:"timeout"`
	RecurseAfterForwarding uint32     `idl:"name:fRecurseAfterForwarding" json:"recurse_after_forwarding"`
	DPFlags                uint32     `idl:"name:dwDpFlags" json:"dp_flags"`
	DPFQDN                 string     `idl:"name:pszDpFqdn;string" json:"dp_fqdn"`
	_                      []uint32   `idl:"name:dwReserved"`
}

func (o *ZoneCreateInfoLonghorn) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ZoneCreateInfoLonghorn) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if o.ZoneName != "" {
		_ptr_pszZoneName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.ZoneName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ZoneName, _ptr_pszZoneName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ZoneType); err != nil {
		return err
	}
	if err := w.WriteData(o.AllowUpdate); err != nil {
		return err
	}
	if err := w.WriteData(o.Aging); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.DataFile != "" {
		_ptr_pszDataFile := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.DataFile); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DataFile, _ptr_pszDataFile); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.DSIntegrated); err != nil {
		return err
	}
	if err := w.WriteData(o.LoadExisting); err != nil {
		return err
	}
	if o.Admin != "" {
		_ptr_pszAdmin := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.Admin); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Admin, _ptr_pszAdmin); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Masters != nil {
		_ptr_aipMasters := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Masters != nil {
				if err := o.Masters.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AddrArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Masters, _ptr_aipMasters); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Secondaries != nil {
		_ptr_aipSecondaries := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Secondaries != nil {
				if err := o.Secondaries.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AddrArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Secondaries, _ptr_aipSecondaries); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SecureSecondaries); err != nil {
		return err
	}
	if err := w.WriteData(o.NotifyLevel); err != nil {
		return err
	}
	if err := w.WriteData(o.Timeout); err != nil {
		return err
	}
	if err := w.WriteData(o.RecurseAfterForwarding); err != nil {
		return err
	}
	if err := w.WriteData(o.DPFlags); err != nil {
		return err
	}
	if o.DPFQDN != "" {
		_ptr_pszDpFqdn := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.DPFQDN); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DPFQDN, _ptr_pszDpFqdn); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	// reserved dwReserved
	for i1 := 0; uint64(i1) < 32; i1++ {
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *ZoneCreateInfoLonghorn) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	_ptr_pszZoneName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.ZoneName); err != nil {
			return err
		}
		return nil
	})
	_s_pszZoneName := func(ptr interface{}) { o.ZoneName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ZoneName, _s_pszZoneName, _ptr_pszZoneName); err != nil {
		return err
	}
	if err := w.ReadData(&o.ZoneType); err != nil {
		return err
	}
	if err := w.ReadData(&o.AllowUpdate); err != nil {
		return err
	}
	if err := w.ReadData(&o.Aging); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_pszDataFile := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.DataFile); err != nil {
			return err
		}
		return nil
	})
	_s_pszDataFile := func(ptr interface{}) { o.DataFile = *ptr.(*string) }
	if err := w.ReadPointer(&o.DataFile, _s_pszDataFile, _ptr_pszDataFile); err != nil {
		return err
	}
	if err := w.ReadData(&o.DSIntegrated); err != nil {
		return err
	}
	if err := w.ReadData(&o.LoadExisting); err != nil {
		return err
	}
	_ptr_pszAdmin := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.Admin); err != nil {
			return err
		}
		return nil
	})
	_s_pszAdmin := func(ptr interface{}) { o.Admin = *ptr.(*string) }
	if err := w.ReadPointer(&o.Admin, _s_pszAdmin, _ptr_pszAdmin); err != nil {
		return err
	}
	_ptr_aipMasters := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Masters == nil {
			o.Masters = &AddrArray{}
		}
		if err := o.Masters.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipMasters := func(ptr interface{}) { o.Masters = *ptr.(**AddrArray) }
	if err := w.ReadPointer(&o.Masters, _s_aipMasters, _ptr_aipMasters); err != nil {
		return err
	}
	_ptr_aipSecondaries := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Secondaries == nil {
			o.Secondaries = &AddrArray{}
		}
		if err := o.Secondaries.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipSecondaries := func(ptr interface{}) { o.Secondaries = *ptr.(**AddrArray) }
	if err := w.ReadPointer(&o.Secondaries, _s_aipSecondaries, _ptr_aipSecondaries); err != nil {
		return err
	}
	if err := w.ReadData(&o.SecureSecondaries); err != nil {
		return err
	}
	if err := w.ReadData(&o.NotifyLevel); err != nil {
		return err
	}
	if err := w.ReadData(&o.Timeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.RecurseAfterForwarding); err != nil {
		return err
	}
	if err := w.ReadData(&o.DPFlags); err != nil {
		return err
	}
	_ptr_pszDpFqdn := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.DPFQDN); err != nil {
			return err
		}
		return nil
	})
	_s_pszDpFqdn := func(ptr interface{}) { o.DPFQDN = *ptr.(*string) }
	if err := w.ReadPointer(&o.DPFQDN, _s_pszDpFqdn, _ptr_pszDpFqdn); err != nil {
		return err
	}
	// reserved dwReserved
	var _dwReserved []uint32
	_dwReserved = make([]uint32, 32)
	for i1 := range _dwReserved {
		i1 := i1
		if err := w.ReadData(&_dwReserved[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// ZoneCreateInfo structure represents DNS_RPC_ZONE_CREATE_INFO RPC structure.
//
// The DNS_RPC_ZONE_CREATE_INFO structure contains information required to create a
// zone or reset a zone's information on the DNS server. There are different versions
// of the DNS_RPC_ZONE_CREATE_INFO structure. The DNS server MUST use the structure
// corresponding to the value of dwClientVersion in DNS Server Management Protocol method
// calls (section 3.1.4) as shown in the following table, or if the method call does
// not specify the value of dwClientVersion, the DNS_RPC_ZONE_CREATE_INFO_W2K version
// of the structure MUST be used.
//
//	+------------+---------------------------------------------------------+
//	|            |                                                         |
//	|   VALUE    |                        STRUCTURE                        |
//	|            |                                                         |
//	+------------+---------------------------------------------------------+
//	+------------+---------------------------------------------------------+
//	| 0x00000000 | DNS_RPC_ZONE_CREATE_INFO_W2K (section 2.2.5.2.7.1)      |
//	+------------+---------------------------------------------------------+
//	| 0x00060000 | DNS_RPC_ZONE_CREATE_INFO_DOTNET (section 2.2.5.2.7.2)   |
//	+------------+---------------------------------------------------------+
//	| 0x00070000 | DNS_RPC_ZONE_CREATE_INFO_LONGHORN (section 2.2.5.2.7.3) |
//	+------------+---------------------------------------------------------+
type ZoneCreateInfo struct {
	RPCStructureVersion    uint32     `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	_                      uint32     `idl:"name:dwReserved0"`
	ZoneName               string     `idl:"name:pszZoneName;string" json:"zone_name"`
	ZoneType               uint32     `idl:"name:dwZoneType" json:"zone_type"`
	AllowUpdate            uint32     `idl:"name:fAllowUpdate" json:"allow_update"`
	Aging                  uint32     `idl:"name:fAging" json:"aging"`
	Flags                  uint32     `idl:"name:dwFlags" json:"flags"`
	DataFile               string     `idl:"name:pszDataFile;string" json:"data_file"`
	DSIntegrated           uint32     `idl:"name:fDsIntegrated" json:"ds_integrated"`
	LoadExisting           uint32     `idl:"name:fLoadExisting" json:"load_existing"`
	Admin                  string     `idl:"name:pszAdmin;string" json:"admin"`
	Masters                *AddrArray `idl:"name:aipMasters" json:"masters"`
	Secondaries            *AddrArray `idl:"name:aipSecondaries" json:"secondaries"`
	SecureSecondaries      uint32     `idl:"name:fSecureSecondaries" json:"secure_secondaries"`
	NotifyLevel            uint32     `idl:"name:fNotifyLevel" json:"notify_level"`
	Timeout                uint32     `idl:"name:dwTimeout" json:"timeout"`
	RecurseAfterForwarding uint32     `idl:"name:fRecurseAfterForwarding" json:"recurse_after_forwarding"`
	DPFlags                uint32     `idl:"name:dwDpFlags" json:"dp_flags"`
	DPFQDN                 string     `idl:"name:pszDpFqdn;string" json:"dp_fqdn"`
	_                      []uint32   `idl:"name:dwReserved"`
}

func (o *ZoneCreateInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ZoneCreateInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if o.ZoneName != "" {
		_ptr_pszZoneName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.ZoneName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ZoneName, _ptr_pszZoneName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ZoneType); err != nil {
		return err
	}
	if err := w.WriteData(o.AllowUpdate); err != nil {
		return err
	}
	if err := w.WriteData(o.Aging); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.DataFile != "" {
		_ptr_pszDataFile := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.DataFile); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DataFile, _ptr_pszDataFile); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.DSIntegrated); err != nil {
		return err
	}
	if err := w.WriteData(o.LoadExisting); err != nil {
		return err
	}
	if o.Admin != "" {
		_ptr_pszAdmin := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.Admin); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Admin, _ptr_pszAdmin); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Masters != nil {
		_ptr_aipMasters := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Masters != nil {
				if err := o.Masters.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AddrArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Masters, _ptr_aipMasters); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Secondaries != nil {
		_ptr_aipSecondaries := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Secondaries != nil {
				if err := o.Secondaries.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AddrArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Secondaries, _ptr_aipSecondaries); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SecureSecondaries); err != nil {
		return err
	}
	if err := w.WriteData(o.NotifyLevel); err != nil {
		return err
	}
	if err := w.WriteData(o.Timeout); err != nil {
		return err
	}
	if err := w.WriteData(o.RecurseAfterForwarding); err != nil {
		return err
	}
	if err := w.WriteData(o.DPFlags); err != nil {
		return err
	}
	if o.DPFQDN != "" {
		_ptr_pszDpFqdn := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.DPFQDN); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DPFQDN, _ptr_pszDpFqdn); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	// reserved dwReserved
	for i1 := 0; uint64(i1) < 32; i1++ {
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *ZoneCreateInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	_ptr_pszZoneName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.ZoneName); err != nil {
			return err
		}
		return nil
	})
	_s_pszZoneName := func(ptr interface{}) { o.ZoneName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ZoneName, _s_pszZoneName, _ptr_pszZoneName); err != nil {
		return err
	}
	if err := w.ReadData(&o.ZoneType); err != nil {
		return err
	}
	if err := w.ReadData(&o.AllowUpdate); err != nil {
		return err
	}
	if err := w.ReadData(&o.Aging); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_pszDataFile := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.DataFile); err != nil {
			return err
		}
		return nil
	})
	_s_pszDataFile := func(ptr interface{}) { o.DataFile = *ptr.(*string) }
	if err := w.ReadPointer(&o.DataFile, _s_pszDataFile, _ptr_pszDataFile); err != nil {
		return err
	}
	if err := w.ReadData(&o.DSIntegrated); err != nil {
		return err
	}
	if err := w.ReadData(&o.LoadExisting); err != nil {
		return err
	}
	_ptr_pszAdmin := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.Admin); err != nil {
			return err
		}
		return nil
	})
	_s_pszAdmin := func(ptr interface{}) { o.Admin = *ptr.(*string) }
	if err := w.ReadPointer(&o.Admin, _s_pszAdmin, _ptr_pszAdmin); err != nil {
		return err
	}
	_ptr_aipMasters := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Masters == nil {
			o.Masters = &AddrArray{}
		}
		if err := o.Masters.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipMasters := func(ptr interface{}) { o.Masters = *ptr.(**AddrArray) }
	if err := w.ReadPointer(&o.Masters, _s_aipMasters, _ptr_aipMasters); err != nil {
		return err
	}
	_ptr_aipSecondaries := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Secondaries == nil {
			o.Secondaries = &AddrArray{}
		}
		if err := o.Secondaries.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipSecondaries := func(ptr interface{}) { o.Secondaries = *ptr.(**AddrArray) }
	if err := w.ReadPointer(&o.Secondaries, _s_aipSecondaries, _ptr_aipSecondaries); err != nil {
		return err
	}
	if err := w.ReadData(&o.SecureSecondaries); err != nil {
		return err
	}
	if err := w.ReadData(&o.NotifyLevel); err != nil {
		return err
	}
	if err := w.ReadData(&o.Timeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.RecurseAfterForwarding); err != nil {
		return err
	}
	if err := w.ReadData(&o.DPFlags); err != nil {
		return err
	}
	_ptr_pszDpFqdn := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.DPFQDN); err != nil {
			return err
		}
		return nil
	})
	_s_pszDpFqdn := func(ptr interface{}) { o.DPFQDN = *ptr.(*string) }
	if err := w.ReadPointer(&o.DPFQDN, _s_pszDpFqdn, _ptr_pszDpFqdn); err != nil {
		return err
	}
	// reserved dwReserved
	var _dwReserved []uint32
	_dwReserved = make([]uint32, 32)
	for i1 := range _dwReserved {
		i1 := i1
		if err := w.ReadData(&_dwReserved[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// SKDList structure represents DNS_RPC_SKD_LIST RPC structure.
//
// The DNS_RPC_SKD_LIST structure contains the information about a list of signing key
// descriptors that are present for a particular zone on the DNS server.
type SKDList struct {
	// dwRpcStructureVersion: The structure version number; this MUST be set to 0x00000001.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// dwReserved0: MUST be set to zero when sent and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved0"`
	// dwCount: The number of signing key descriptors present in the array of signing key
	// descriptors pointed to by SkdArray.
	Count uint32 `idl:"name:dwCount" json:"count"`
	// SkdArray: An array of structures of type DNS_RPC_SKD (section 2.2.6.2.1). Each element
	// of the array represents one signing key descriptor.
	SKDArray []*SKD `idl:"name:SkdArray;size_is:(dwCount)" json:"skd_array"`
}

func (o *SKDList) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.SKDArray != nil && o.Count == 0 {
		o.Count = uint32(len(o.SKDArray))
	}
	if o.Count > uint32(1000) {
		return fmt.Errorf("Count is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *SKDList) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.Count)
	return []uint64{
		dimSize1,
	}
}
func (o *SKDList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.Count); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	for i1 := range o.SKDArray {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.SKDArray[i1] != nil {
			_ptr_SkdArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SKDArray[i1] != nil {
					if err := o.SKDArray[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&SKD{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SKDArray[i1], _ptr_SkdArray); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.SKDArray); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SKDList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	if err := w.ReadData(&o.Count); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.Count > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.Count)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.SKDArray", sizeInfo[0])
	}
	o.SKDArray = make([]*SKD, sizeInfo[0])
	for i1 := range o.SKDArray {
		i1 := i1
		_ptr_SkdArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SKDArray[i1] == nil {
				o.SKDArray[i1] = &SKD{}
			}
			if err := o.SKDArray[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_SkdArray := func(ptr interface{}) { o.SKDArray[i1] = *ptr.(**SKD) }
		if err := w.ReadPointer(&o.SKDArray[i1], _s_SkdArray, _ptr_SkdArray); err != nil {
			return err
		}
	}
	return nil
}

// SigningValidationError structure represents DNS_RPC_SIGNING_VALIDATION_ERROR RPC structure.
//
// The DNS_RPC_SIGNING_VALIDATION_ERROR structure describes an error that occurred during
// the use of an SKD.
type SigningValidationError struct {
	// dwRpcStructureVersion: The structure version number; this MUST be set to 0x00000001.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// dwReserved0: MUST be set to zero when sent and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved0"`
	// guidSKD: The unique identifier for the signing key descriptor that caused the validation
	// error.
	SKD *dtyp.GUID `idl:"name:guidSKD" json:"skd"`
	// pwszSigningKeyPointerString: The signing key pointer string of the signing key that
	// caused the validation error.
	SigningKeyPointerString string `idl:"name:pwszSigningKeyPointerString;string" json:"signing_key_pointer_string"`
	// dwExtendedError: MUST be set to zero when sent.
	ExtendedError uint32 `idl:"name:dwExtendedError" json:"extended_error"`
	// dwReserved: MUST be set to zero when sent and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved"`
}

func (o *SigningValidationError) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SigningValidationError) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if o.SKD != nil {
		if err := o.SKD.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SigningKeyPointerString != "" {
		_ptr_pwszSigningKeyPointerString := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.SigningKeyPointerString); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SigningKeyPointerString, _ptr_pwszSigningKeyPointerString); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ExtendedError); err != nil {
		return err
	}
	// reserved dwReserved
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *SigningValidationError) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	if o.SKD == nil {
		o.SKD = &dtyp.GUID{}
	}
	if err := o.SKD.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_pwszSigningKeyPointerString := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.SigningKeyPointerString); err != nil {
			return err
		}
		return nil
	})
	_s_pwszSigningKeyPointerString := func(ptr interface{}) { o.SigningKeyPointerString = *ptr.(*string) }
	if err := w.ReadPointer(&o.SigningKeyPointerString, _s_pwszSigningKeyPointerString, _ptr_pwszSigningKeyPointerString); err != nil {
		return err
	}
	if err := w.ReadData(&o.ExtendedError); err != nil {
		return err
	}
	// reserved dwReserved
	var _dwReserved uint32
	if err := w.ReadData(&_dwReserved); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// AutoConfigure structure represents DNS_RPC_AUTOCONFIGURE RPC structure.
//
// The DNS_RPC_AUTOCONFIGURE structure contains the information required to auto-configure
// the DNS server.
type AutoConfigure struct {
	// dwRpcStructureVersion: The structure version number; this value MUST be set to 0x00000000.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// dwReserved0: MUST be set to zero when sent and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved0"`
	// dwAutoConfigFlags: The autoconfiguration operation being requested by the client
	// as specified in DNS_RPC_AUTOCONFIG (section 2.2.8.1.1).
	AutoConfigFlags uint32 `idl:"name:dwAutoConfigFlags" json:"auto_config_flags"`
	// dwReserved1: MUST be set to zero when sent and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved1"`
	// pszNewDomainName: A pointer to a null-terminated UTF-8 string which contains the
	// name of the directory server domain that this DNS server is about to join.
	NewDomainName string `idl:"name:pszNewDomainName;string" json:"new_domain_name"`
}

func (o *AutoConfigure) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *AutoConfigure) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.AutoConfigFlags); err != nil {
		return err
	}
	// reserved dwReserved1
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if o.NewDomainName != "" {
		_ptr_pszNewDomainName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.NewDomainName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.NewDomainName, _ptr_pszNewDomainName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *AutoConfigure) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	if err := w.ReadData(&o.AutoConfigFlags); err != nil {
		return err
	}
	// reserved dwReserved1
	var _dwReserved1 uint32
	if err := w.ReadData(&_dwReserved1); err != nil {
		return err
	}
	_ptr_pszNewDomainName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.NewDomainName); err != nil {
			return err
		}
		return nil
	})
	_s_pszNewDomainName := func(ptr interface{}) { o.NewDomainName = *ptr.(*string) }
	if err := w.ReadPointer(&o.NewDomainName, _s_pszNewDomainName, _ptr_pszNewDomainName); err != nil {
		return err
	}
	return nil
}

// EnumZonesFilter structure represents DNS_RPC_ENUM_ZONES_FILTER RPC structure.
//
// The DNS_RPC_ENUM_ZONES_FILTER structure specifies zone filtering criteria.
type EnumZonesFilter struct {
	// dwRpcStructureVersion: The structure version number; this MUST be set to 0x00000001.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// dwReserved0: MUST be set to zero when sent and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved0"`
	// dwFilter: A filter value that specifies the zone types that are to be selected as
	// part of the output zone list. This value MUST be set to any combination of the ZONE_REQUEST_FILTERS
	// (section 2.2.5.1.4).
	Filter uint32 `idl:"name:dwFilter" json:"filter"`
	// pszPartitionFqdn: A pointer to a null-terminated UTF-8 string that specifies the
	// distinguished name for an application directory partition location from which the
	// server is to enumerate zones; if this is NULL then zone enumeration is not restricted
	// based on the application directory partition.
	PartitionFQDN string `idl:"name:pszPartitionFqdn;string" json:"partition_fqdn"`
	// pszQueryString: Reserved for future use. Server MUST set to zero and receiver MUST
	// ignore this value.
	QueryString string `idl:"name:pszQueryString;string" json:"query_string"`
	// pszReserved: Reserved for future use. Server MUST set to zero and receiver MUST ignore
	// this value.
	_ string `idl:"name:pszReserved;string"`
}

func (o *EnumZonesFilter) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *EnumZonesFilter) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.Filter); err != nil {
		return err
	}
	if o.PartitionFQDN != "" {
		_ptr_pszPartitionFqdn := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.PartitionFQDN); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.PartitionFQDN, _ptr_pszPartitionFqdn); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.QueryString != "" {
		_ptr_pszQueryString := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.QueryString); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.QueryString, _ptr_pszQueryString); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	// reserved pszReserved
	if err := w.WritePointer(nil); err != nil {
		return err
	}
	return nil
}
func (o *EnumZonesFilter) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	if err := w.ReadData(&o.Filter); err != nil {
		return err
	}
	_ptr_pszPartitionFqdn := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.PartitionFQDN); err != nil {
			return err
		}
		return nil
	})
	_s_pszPartitionFqdn := func(ptr interface{}) { o.PartitionFQDN = *ptr.(*string) }
	if err := w.ReadPointer(&o.PartitionFQDN, _s_pszPartitionFqdn, _ptr_pszPartitionFqdn); err != nil {
		return err
	}
	_ptr_pszQueryString := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.QueryString); err != nil {
			return err
		}
		return nil
	})
	_s_pszQueryString := func(ptr interface{}) { o.QueryString = *ptr.(*string) }
	if err := w.ReadPointer(&o.QueryString, _s_pszQueryString, _ptr_pszQueryString); err != nil {
		return err
	}
	// reserved pszReserved
	var _pszReserved string
	_ptr_pszReserved := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &_pszReserved); err != nil {
			return err
		}
		return nil
	})
	_s_pszReserved := func(ptr interface{}) { _pszReserved = *ptr.(*string) }
	if err := w.ReadPointer(&_pszReserved, _s_pszReserved, _ptr_pszReserved); err != nil {
		return err
	}
	return nil
}

// Record structure represents DNS_RPC_RECORD RPC structure.
//
// The DNS_RPC_RECORD structure is used to specify a single DNS record's parameters
// and data. This structure is returned by the DNS server in response to an R_DnssrvEnumRecords2
// (section 3.1.4.9) method call.
type Record struct {
	// wDataLength: The total size of the variable buffer, in bytes. Note that the DNS_RPC_RECORD
	// structure is always 4-byte aligned, which means there can be 0-3 bytes of padding
	// at the end of the structure. The pad bytes are not included in the wDataLength count.
	DataLength uint16 `idl:"name:wDataLength" json:"data_length"`
	// wType: The type of the resource record, as specified in section 2.2.2.1.1 DNS_RECORD_TYPE.
	Type uint16 `idl:"name:wType" json:"type"`
	// dwFlags: Resource record properties. This field can contain one of the RANK* flags
	// in the low-order bits and one of the DNS_RPC_FLAGS* in the high-order bits.
	//
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                            |                                                                                  |
	//	|                   VALUE                    |                                     MEANING                                      |
	//	|                                            |                                                                                  |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| RANK_CACHE_BIT 0x00000001                  | The record came from the cache.                                                  |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| RANK_ROOT_HINT 0x00000008                  | The record is a preconfigured root hint.                                         |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| RANK_OUTSIDE_GLUE 0x00000020               | This value is not used.                                                          |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| RANK_CACHE_NA_ADDITIONAL 0x00000031        | The record was cached from the additional section of a non-authoritative         |
	//	|                                            | response.                                                                        |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| RANK_CACHE_NA_AUTHORITY 0x00000041         | The record was cached from the authority section of a non-authoritative          |
	//	|                                            | response.                                                                        |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| RANK_CACHE_A_ADDITIONAL 0x00000051         | The record was cached from the additional section of an authoritative response.  |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| RANK_CACHE_NA_ANSWER 0x00000061            | The record was cached from the answer section of a non-authoritative response.   |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| RANK_CACHE_A_AUTHORITY 0x00000071          | The record was cached from the authority section of an authoritative response.   |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| RANK_GLUE 0x00000080                       | The record is a glue record in an authoritative zone.                            |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| RANK_NS_GLUE 0x00000082                    | The record is a delegation  (type NS) record in an authoritative zone.           |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| RANK_CACHE_A_ANSWER 0x000000C1             | The record was cached from the answer section of an authoritative response.      |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| RANK_ZONE 0x000000F0                       | The record comes from an authoritative zone.                                     |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_RPC_FLAG_ZONE_ROOT 0x40000000          | The record is at the root of a zone (not necessarily a zone hosted by this       |
	//	|                                            | server; the record could have come from the cache).                              |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_RPC_FLAG_AUTH_ZONE_ROOT 0x20000000     | The record is at the root of a zone that is locally hosted on this server.       |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_RPC_FLAG_CACHE_DATA 0x80000000         | The record came from the cache.                                                  |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_RPC_FLAG_RECORD_WIRE_FORMAT 0x00100000 | The record SHOULD<25> be treated as a resource record of unknown type ([RFC3597] |
	//	|                                            | section 2) by the DNS server.                                                    |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
	// dwSerial: This MUST be set to 0x00000000 when sent by the client or server, and ignored
	// on receipt by the server or client.
	Serial uint32 `idl:"name:dwSerial" json:"serial"`
	// dwTtlSeconds: The duration, in seconds, after which this record will expire.
	TTLSeconds uint32 `idl:"name:dwTtlSeconds" json:"ttl_seconds"`
	// dwTimeStamp: The time stamp, in hours, for the record when it received the last update.
	Timestamp uint32 `idl:"name:dwTimeStamp" json:"timestamp"`
	// dwReserved: This value MUST be set to 0x00000000 when sent by the client and ignored
	// on receipt by the server.
	_ uint32 `idl:"name:dwReserved"`
	// Buffer: Record data in DNS_RPC_RECORD_DATA (section 2.2.2.2.4) format where type
	// is specified by the value wType.<26>
	//
	//	+----------------------------+--------------------------------+
	//	|                            |                                |
	//	|           VALUE            |            MEANING             |
	//	|                            |                                |
	//	+----------------------------+--------------------------------+
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_ZERO 0x0000       | DNS_RPC_RECORD_TS              |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_A 0x0001          | DNS_RPC_RECORD_A               |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_NS 0x0002         | DNS_RPC_RECORD_NODE_NAME       |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_MD 0x0003         | DNS_RPC_RECORD_NODE_NAME       |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_MF 0x0004         | DNS_RPC_RECORD_NODE_NAME       |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_CNAME 0x0005      | DNS_RPC_RECORD_NODE_NAME       |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_SOA 0x0006        | DNS_RPC_RECORD_SOA             |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_MB 0x0007         | DNS_RPC_RECORD_NODE_NAME       |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_MG 0x0008         | DNS_RPC_RECORD_NODE_NAME       |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_MR 0x0009         | DNS_RPC_RECORD_NODE_NAME       |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_NULL 0x000A       | DNS_RPC_RECORD_NULL            |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_WKS 0x000B        | DNS_RPC_RECORD_WKS             |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_PTR 0x000C        | DNS_RPC_RECORD_NODE_NAME       |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_HINFO 0x000D      | DNS_RPC_RECORD_STRING          |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_MINFO 0x000E      | DNS_RPC_RECORD_MAIL_ERROR      |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_MX 0x000F         | DNS_RPC_RECORD_NAME_PREFERENCE |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_TXT 0x0010        | DNS_RPC_RECORD_STRING          |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_RP 0x0011         | DNS_RPC_RECORD_MAIL_ERROR      |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_AFSDB 0x0012      | DNS_RPC_RECORD_NAME_PREFERENCE |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_X25 0x0013        | DNS_RPC_RECORD_STRING          |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_ISDN 0x0014       | DNS_RPC_RECORD_STRING          |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_RT 0x0015         | DNS_RPC_RECORD_NAME_PREFERENCE |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_SIG 0x0018        | DNS_RPC_RECORD_SIG             |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_KEY 0x0019        | DNS_RPC_RECORD_KEY             |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_AAAA 0x001C       | DNS_RPC_RECORD_AAAA            |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_NXT 0x001E        | DNS_RPC_RECORD_NXT             |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_SRV 0x0021        | DNS_RPC_RECORD_SRV             |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_ATMA 0x0022       | DNS_RPC_RECORD_ATMA            |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_NAPTR 0x0023      | DNS_RPC_RECORD_NAPTR           |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_DNAME 0x0027      | DNS_RPC_RECORD_NODE_NAME       |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_DS 0x002B         | DNS_RPC_RECORD_DS              |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_RRSIG 0x002E      | DNS_RPC_RECORD_RRSIG           |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_NSEC 0x002F       | DNS_RPC_RECORD_NSEC            |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_DNSKEY 0x0030     | DNS_RPC_RECORD_DNSKEY          |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_DHCID 0x0031      | DNS_RPC_RECORD_DHCID           |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_NSEC3 0x0032      | DNS_RPC_RECORD_NSEC3           |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_NSEC3PARAM 0x0033 | DNS_RPC_RECORD_NSEC3PARAM      |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_TLSA 0x0034       | DNS_RPC_RECORD_TLSA            |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_WINS 0xFF01       | DNS_RPC_RECORD_WINS            |
	//	+----------------------------+--------------------------------+
	//	| DNS_TYPE_WINSR 0xFF02      | DNS_RPC_RECORD_WINSR           |
	//	+----------------------------+--------------------------------+
	Buffer []byte `idl:"name:Buffer;size_is:(wDataLength)" json:"buffer"`
}

func (o *Record) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Buffer != nil && o.DataLength == 0 {
		o.DataLength = uint16(len(o.Buffer))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *Record) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataLength)
	return []uint64{
		dimSize1,
	}
}
func (o *Record) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.DataLength); err != nil {
		return err
	}
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.Serial); err != nil {
		return err
	}
	if err := w.WriteData(o.TTLSeconds); err != nil {
		return err
	}
	if err := w.WriteData(o.Timestamp); err != nil {
		return err
	}
	// reserved dwReserved
	if err := w.WriteData(uint32(0)); err != nil {
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
}
func (o *Record) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.DataLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.Serial); err != nil {
		return err
	}
	if err := w.ReadData(&o.TTLSeconds); err != nil {
		return err
	}
	if err := w.ReadData(&o.Timestamp); err != nil {
		return err
	}
	// reserved dwReserved
	var _dwReserved uint32
	if err := w.ReadData(&_dwReserved); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataLength > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataLength)
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
}

// FlatRecord structure represents DNS_FLAT_RECORD RPC structure.
type FlatRecord struct {
	DataLength uint16 `idl:"name:wDataLength" json:"data_length"`
	Type       uint16 `idl:"name:wType" json:"type"`
	Flags      uint32 `idl:"name:dwFlags" json:"flags"`
	Serial     uint32 `idl:"name:dwSerial" json:"serial"`
	TTLSeconds uint32 `idl:"name:dwTtlSeconds" json:"ttl_seconds"`
	Timestamp  uint32 `idl:"name:dwTimeStamp" json:"timestamp"`
	_          uint32 `idl:"name:dwReserved"`
	Buffer     []byte `idl:"name:Buffer;size_is:(wDataLength)" json:"buffer"`
}

func (o *FlatRecord) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Buffer != nil && o.DataLength == 0 {
		o.DataLength = uint16(len(o.Buffer))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *FlatRecord) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataLength)
	return []uint64{
		dimSize1,
	}
}
func (o *FlatRecord) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.DataLength); err != nil {
		return err
	}
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.Serial); err != nil {
		return err
	}
	if err := w.WriteData(o.TTLSeconds); err != nil {
		return err
	}
	if err := w.WriteData(o.Timestamp); err != nil {
		return err
	}
	// reserved dwReserved
	if err := w.WriteData(uint32(0)); err != nil {
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
}
func (o *FlatRecord) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.DataLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.Serial); err != nil {
		return err
	}
	if err := w.ReadData(&o.TTLSeconds); err != nil {
		return err
	}
	if err := w.ReadData(&o.Timestamp); err != nil {
		return err
	}
	// reserved dwReserved
	var _dwReserved uint32
	if err := w.ReadData(&_dwReserved); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataLength > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataLength)
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
}

// NameAndParam structure represents DNS_RPC_NAME_AND_PARAM RPC structure.
//
// The DNS_RPC_NAME_AND_PARAM structure contains information about a simple server property
// that takes a DWORD value. The DNS Server Management Protocol uses this structure
// to exchange information about various properties that take an integer value, while
// processing the R_DnssrvOperation2 (section 3.1.4.6) method call with operation types
// "ResetDwordProperty", "DeleteNode", and "DeleteRecordSet".
type NameAndParam struct {
	// dwParam: The requested new value for the server property specified by pszNodeName.
	Param uint32 `idl:"name:dwParam" json:"param"`
	// pszNodeName: Pointer to a null-terminated UTF-8 string that specifies the name of
	// the server property.
	NodeName string `idl:"name:pszNodeName;string" json:"node_name"`
}

func (o *NameAndParam) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NameAndParam) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Param); err != nil {
		return err
	}
	if o.NodeName != "" {
		_ptr_pszNodeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.NodeName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.NodeName, _ptr_pszNodeName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *NameAndParam) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Param); err != nil {
		return err
	}
	_ptr_pszNodeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.NodeName); err != nil {
			return err
		}
		return nil
	})
	_s_pszNodeName := func(ptr interface{}) { o.NodeName = *ptr.(*string) }
	if err := w.ReadPointer(&o.NodeName, _s_pszNodeName, _ptr_pszNodeName); err != nil {
		return err
	}
	return nil
}

// IPValidate structure represents DNS_RPC_IP_VALIDATE RPC structure.
//
// The DNS_RPC_IP_VALIDATE structure is used to request that the DNS server validate
// a number of IP addresses. This can be used by clients to determine if an IP address
// is suitable for use as a DNS server in the context specified by the dwContext member
// (see below). This structure is to request IP validation while processing the R_DnssrvComplexOperation2
// (section 3.1.4.8) method call with operation type IpValidate.
type IPValidate struct {
	// dwRpcStructureVersion: The DNS management structure version number. This value MUST
	// be set to 0x00000001.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// dwReserved0: This field is unused. The client MUST set this value to zero and the
	// server MUST ignore it.
	_ uint32 `idl:"name:dwReserved0"`
	// dwContext: The context or purpose for which addresses present in aipValidateAddrs
	// MUST be validated by the DNS server. This field MUST be set to one of the following
	// values:
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                                       |                                                                                  |
	//	|                 VALUE                 |                                     MEANING                                      |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_IPVAL_DNS_SERVERS 0x00000000      | Validate that IP addresses are reachable and operational by the DNS servers.     |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_IPVAL_DNS_ROOTHINTS 0x00000001    | Validate that IP addresses are suitable as root hints.                           |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_IPVAL_DNS_FORWARDERS 0x00000002   | Validate that IP addresses are server level forwarders.                          |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_IPVAL_DNS_ZONE_MASTERS 0x00000003 | Validate that IP addresses are remote DNS servers hosting a zone, named as       |
	//	|                                       | pointed to by pszContextName.                                                    |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_IPVAL_DNS_DELEGATIONS 0x00000004  | Validate that IP addresses are remote DNS servers are name-server for the        |
	//	|                                       | delegated zone, named as pointed to by pszContextName.                           |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	Context uint32 `idl:"name:dwContext" json:"context"`
	// dwReserved1: This field is unused. The client MUST set this to zero and the server
	// MUST ignore it.
	_ uint32 `idl:"name:dwReserved1"`
	// pszContextName: A pointer to a null-terminated ASCII character string that specifies
	// a zone name. The use of this zone name is specified by the dwContext member.
	ContextName string `idl:"name:pszContextName;string" json:"context_name"`
	// aipValidateAddrs: A pointer to a DNS_ADDR_ARRAY structure (section 2.2.3.2.3) contains
	// a list of IP addresses to be validated by the DNS server.
	ValidateAddrs *AddrArray `idl:"name:aipValidateAddrs" json:"validate_addrs"`
}

func (o *IPValidate) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPValidate) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.Context); err != nil {
		return err
	}
	// reserved dwReserved1
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if o.ContextName != "" {
		_ptr_pszContextName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.ContextName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ContextName, _ptr_pszContextName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ValidateAddrs != nil {
		_ptr_aipValidateAddrs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ValidateAddrs != nil {
				if err := o.ValidateAddrs.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AddrArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ValidateAddrs, _ptr_aipValidateAddrs); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPValidate) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved0
	var _dwReserved0 uint32
	if err := w.ReadData(&_dwReserved0); err != nil {
		return err
	}
	if err := w.ReadData(&o.Context); err != nil {
		return err
	}
	// reserved dwReserved1
	var _dwReserved1 uint32
	if err := w.ReadData(&_dwReserved1); err != nil {
		return err
	}
	_ptr_pszContextName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.ContextName); err != nil {
			return err
		}
		return nil
	})
	_s_pszContextName := func(ptr interface{}) { o.ContextName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ContextName, _s_pszContextName, _ptr_pszContextName); err != nil {
		return err
	}
	_ptr_aipValidateAddrs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ValidateAddrs == nil {
			o.ValidateAddrs = &AddrArray{}
		}
		if err := o.ValidateAddrs.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_aipValidateAddrs := func(ptr interface{}) { o.ValidateAddrs = *ptr.(**AddrArray) }
	if err := w.ReadPointer(&o.ValidateAddrs, _s_aipValidateAddrs, _ptr_aipValidateAddrs); err != nil {
		return err
	}
	return nil
}

// UTF8StringList structure represents DNS_RPC_UTF8_STRING_LIST RPC structure.
//
// The DNS_RPC_UTF8_STRING_LIST structure contains a list of null-terminated UTF-8 strings.
// This structure is used by the DNS Server Management Protocol while processing R_DnssrvOperations2
// (section 3.1.4.6) and R_DnssrvQuery2 (section 3.1.4.7) method calls, with operations
// type "GlobalQueryBlockList".
type UTF8StringList struct {
	// dwCount: The number of strings present in the pszStrings array.
	Count uint32 `idl:"name:dwCount" json:"count"`
	// pszStrings: A variable length array of pointers to null-terminated UTF-8 strings.
	Strings string `idl:"name:pszStrings;size_is:(dwCount);string" json:"strings"`
}

func (o *UTF8StringList) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Strings != "" && o.Count == 0 {
		o.Count = uint32(ndr.CharNLen(o.Strings))
	}
	if o.Count > uint32(10000) {
		return fmt.Errorf("Count is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *UTF8StringList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Count); err != nil {
		return err
	}
	if o.Strings != "" || o.Count > 0 {
		_ptr_pszStrings := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Count)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			dimLength1 := ndr.CharNLen(o.Strings)
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
			_Strings_buf := []byte(o.Strings)
			if uint64(len(_Strings_buf)) > sizeInfo[0]-1 {
				_Strings_buf = _Strings_buf[:sizeInfo[0]-1]
			}
			if o.Strings != ndr.ZeroString {
				_Strings_buf = append(_Strings_buf, byte(0))
			}
			for i1 := range _Strings_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_Strings_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_Strings_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Strings, _ptr_pszStrings); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *UTF8StringList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Count); err != nil {
		return err
	}
	_ptr_pszStrings := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
		var _Strings_buf []byte
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Strings_buf", sizeInfo[0])
		}
		_Strings_buf = make([]byte, sizeInfo[0])
		for i1 := range _Strings_buf {
			i1 := i1
			if err := w.ReadData(&_Strings_buf[i1]); err != nil {
				return err
			}
		}
		o.Strings = strings.TrimRight(string(_Strings_buf), ndr.ZeroString)
		return nil
	})
	_s_pszStrings := func(ptr interface{}) { o.Strings = *ptr.(*string) }
	if err := w.ReadPointer(&o.Strings, _s_pszStrings, _ptr_pszStrings); err != nil {
		return err
	}
	return nil
}

// UnicodeStringList structure represents DNS_RPC_UNICODE_STRING_LIST RPC structure.
//
// The DNS_RPC_UNICODE_STRING_LIST structure contains a list of null-terminated Unicode
// strings. This structure is used by the DNS Server Management Protocol while processing
// the R_DnssrvComplexOperation2 (section 3.1.4.8) method call, with operations of type
// "EnumerateKeyStorageProviders".
type UnicodeStringList struct {
	// dwCount: The number of strings present in the pwszStrings array.
	Count uint32 `idl:"name:dwCount" json:"count"`
	// pwszStrings: A variable-length array of pointers to null-terminated Unicode strings.
	Strings string `idl:"name:pwszStrings;size_is:(dwCount);string" json:"strings"`
}

func (o *UnicodeStringList) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Strings != "" && o.Count == 0 {
		o.Count = uint32(ndr.UTF16NLen(o.Strings))
	}
	if o.Count > uint32(10000) {
		return fmt.Errorf("Count is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *UnicodeStringList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Count); err != nil {
		return err
	}
	if o.Strings != "" || o.Count > 0 {
		_ptr_pwszStrings := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Count)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			dimLength1 := ndr.UTF16NLen(o.Strings)
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
			_Strings_buf := utf16.Encode([]rune(o.Strings))
			if uint64(len(_Strings_buf)) > sizeInfo[0]-1 {
				_Strings_buf = _Strings_buf[:sizeInfo[0]-1]
			}
			if o.Strings != ndr.ZeroString {
				_Strings_buf = append(_Strings_buf, uint16(0))
			}
			for i1 := range _Strings_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_Strings_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_Strings_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Strings, _ptr_pwszStrings); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *UnicodeStringList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Count); err != nil {
		return err
	}
	_ptr_pwszStrings := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
		var _Strings_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Strings_buf", sizeInfo[0])
		}
		_Strings_buf = make([]uint16, sizeInfo[0])
		for i1 := range _Strings_buf {
			i1 := i1
			if err := w.ReadData(&_Strings_buf[i1]); err != nil {
				return err
			}
		}
		o.Strings = strings.TrimRight(string(utf16.Decode(_Strings_buf)), ndr.ZeroString)
		return nil
	})
	_s_pwszStrings := func(ptr interface{}) { o.Strings = *ptr.(*string) }
	if err := w.ReadPointer(&o.Strings, _s_pwszStrings, _ptr_pwszStrings); err != nil {
		return err
	}
	return nil
}

// CriteriaComparator type represents DNS_RPC_CRITERIA_COMPARATOR RPC enumeration.
//
// The DNS_RPC_CRITERIA_COMPARATOR enumeration specifies the logical condition which
// connects the DNS Policy Criteria evaluation during DNS Operations.
type CriteriaComparator uint16

var (
	// Equals: Specifies that the values in a criteria (DNS_RPC_POLICY section 2.2.15.2.5)
	// of a DNS policy are connected by a logical OR condition.
	CriteriaComparatorEquals CriteriaComparator = 1
	// NotEquals: Specifies that the values in a criteria (DNS_RPC_POLICY section 2.2.15.2.5)
	// of a DNS Policy are connected by logical AND condition.
	CriteriaComparatorNotEquals CriteriaComparator = 2
)

func (o CriteriaComparator) String() string {
	switch o {
	case CriteriaComparatorEquals:
		return "CriteriaComparatorEquals"
	case CriteriaComparatorNotEquals:
		return "CriteriaComparatorNotEquals"
	}
	return "Invalid"
}

// PolicyCondition type represents DNS_RPC_POLICY_CONDITION RPC enumeration.
//
// The DNS_RPC_POLICY_CONDITION enumeration specifies the logical condition that connects
// the DNS Policy Criteria evaluation during DNS operations.
type PolicyCondition uint16

var (
	// DNS_AND: While evaluating a DNS policy (section 2.2.15.2.5) during a DNS Operation,
	// all criteria (section 2.2.15.2.4) configured for that particular DNS Policy MUST
	// match.
	PolicyConditionAnd PolicyCondition = 0
	// DNS_OR: While evaluating a DNS Policy during a DNS Operation, at least one configured
	// criterion MUST match.
	PolicyConditionOr PolicyCondition = 1
)

func (o PolicyCondition) String() string {
	switch o {
	case PolicyConditionAnd:
		return "PolicyConditionAnd"
	case PolicyConditionOr:
		return "PolicyConditionOr"
	}
	return "Invalid"
}

// PolicyLevel type represents DNS_RPC_POLICY_LEVEL RPC enumeration.
//
// The DNS_RPC_POLICY_LEVEL enumeration specifies that the DNS Policy is applied for
// a zone, cache zone, or at server level.
type PolicyLevel uint16

var (
	// DnsPolicyServerLevel: The DNS Policy is applicable at the server level for all DNS
	// operations.
	PolicyLevelServerLevel PolicyLevel = 0
	// DnsPolicyZoneLevel: The DNS Policy is applicable only for a specified zone. It is
	// applicable for all DNS operations allowed for any zone.
	PolicyLevelZoneLevel PolicyLevel = 1
	// DnsPolicyLevelMax: Shows the maximum level types supported.
	PolicyLevelMax PolicyLevel = 2
)

func (o PolicyLevel) String() string {
	switch o {
	case PolicyLevelServerLevel:
		return "PolicyLevelServerLevel"
	case PolicyLevelZoneLevel:
		return "PolicyLevelZoneLevel"
	case PolicyLevelMax:
		return "PolicyLevelMax"
	}
	return "Invalid"
}

// PolicyActionType type represents DNS_RPC_POLICY_ACTION_TYPE RPC enumeration.
//
// The DNS_RPC_POLICY_ACTION_TYPE enumeration specifies the action to be applied when
// a DNS Policy is matched.
type PolicyActionType uint16

var (
	// DnsPolicyDeny: Denies the specific DNS operation.
	PolicyActionTypeDeny PolicyActionType = 0
	// DnsPolicyAllow: Allows a specific DNS operation.
	PolicyActionTypeAllow  PolicyActionType = 1
	PolicyActionTypeIgnore PolicyActionType = 2
	// DnsPolicyActionMax: Shows the maximum action types supported.
	PolicyActionTypeMax PolicyActionType = 3
)

func (o PolicyActionType) String() string {
	switch o {
	case PolicyActionTypeDeny:
		return "PolicyActionTypeDeny"
	case PolicyActionTypeAllow:
		return "PolicyActionTypeAllow"
	case PolicyActionTypeIgnore:
		return "PolicyActionTypeIgnore"
	case PolicyActionTypeMax:
		return "PolicyActionTypeMax"
	}
	return "Invalid"
}

// PolicyType type represents DNS_RPC_POLICY_TYPE RPC enumeration.
//
// The DNS_RPC_POLICY_TYPE enumeration specifies the DNS operation for which the DNS
// policy is applicable.
//
//	+----------------------------+----------------------------------------------------------------------------------+
//	|         ENUMERATOR         |                                                                                  |
//	|           VALUE            |                                   DESCRIPTION                                    |
//	|                            |                                                                                  |
//	+----------------------------+----------------------------------------------------------------------------------+
//	+----------------------------+----------------------------------------------------------------------------------+
//	| DnsPolicyQueryProcessing=0 | The DNS policy is applicable for DNS queries.                                    |
//	+----------------------------+----------------------------------------------------------------------------------+
//	| DnsPolicyZoneTransfer      | The DNS policy is applicable for zone transfer.                                  |
//	+----------------------------+----------------------------------------------------------------------------------+
//	| DnsPolicyDynamicUpdate     | The DNS policy is applicable for dynamic DNS updates.                            |
//	+----------------------------+----------------------------------------------------------------------------------+
//	| DnsPolicyRecursion         | The DNS policy is applicable for recursive queries.                              |
//	+----------------------------+----------------------------------------------------------------------------------+
//	| DnsPolicyRRLExceptionList  | The DNS policy is applicable for RRL Exceptionlisting. Exceptionlisting exempts  |
//	|                            | queries from response rate limiting. The queries can be exempted based on name,  |
//	|                            | source subnet, or the DNS server interface address where the query was received. |
//	+----------------------------+----------------------------------------------------------------------------------+
//	| DnsPolicyMax               | Shows the maximum DNS policies supported.                                        |
//	+----------------------------+----------------------------------------------------------------------------------+
type PolicyType uint16

var (
	PolicyTypeQueryProcessing PolicyType = 0
	PolicyTypeZoneTransfer    PolicyType = 1
	PolicyTypeDynamicUpdate   PolicyType = 2
	PolicyTypeRecursion       PolicyType = 3
	PolicyTypeMax             PolicyType = 4
)

func (o PolicyType) String() string {
	switch o {
	case PolicyTypeQueryProcessing:
		return "PolicyTypeQueryProcessing"
	case PolicyTypeZoneTransfer:
		return "PolicyTypeZoneTransfer"
	case PolicyTypeDynamicUpdate:
		return "PolicyTypeDynamicUpdate"
	case PolicyTypeRecursion:
		return "PolicyTypeRecursion"
	case PolicyTypeMax:
		return "PolicyTypeMax"
	}
	return "Invalid"
}

// CriteriaEnum type represents DNS_RPC_CRITERIA_ENUM RPC enumeration.
//
// The DNS_RPC_CRITERIA_ENUM enumeration specifies the various criteria that a DNS server
// needs to match during a DNS operation to have a successful DNS policy match. For
// more information, see section 3.1.8.2.
type CriteriaEnum uint16

var (
	// DnsPolicyCriteriaSubnet: This specifies that the criteria for match of a DNS policy
	// is based on the client subnet record, derived from the IP address of a resolver [RFC1034]
	// section 2.4 in a DNS Operation.
	CriteriaEnumPolicyCriteriaSubnet CriteriaEnum = 0
	// DnsPolicyCriteriaTransportProtocol: This specifies that the criteria for match of
	// a DNS policy is based on the transport layer protocol (TCP or UDP) used to send a
	// DNS query to the DNS server during a DNS operation.
	CriteriaEnumPolicyCriteriaTransportProtocol CriteriaEnum = 1
	// DnsPolicyCriteriaNetworkProtocol: This specifies that the criteria for match of a
	// DNS policy is based on the Internet protocol used (IPv4 or IPv6) to send a DNS query
	// to the DNS server during a DNS operation.
	CriteriaEnumPolicyCriteriaNetworkProtocol CriteriaEnum = 2
	// DnsPolicyCriteriaInterface: This specifies that the criteria for match of a DNS policy
	// is based on the IP address on which the DNS server receives the DNS query in a DNS
	// operation.
	CriteriaEnumPolicyCriteriaInterface CriteriaEnum = 3
	// DnsPolicyCriteriaFqdn: This specifies that the criteria for match of a DNS policy
	// is based on the query name in the question section [RFC1034] section 3.7 of a DNS
	// query, in a DNS operation.
	CriteriaEnumPolicyCriteriaFQDN CriteriaEnum = 4
	// DnsPolicyCriteriaQtype: This specifies that the criteria for match of a DNS policy
	// is based on the QTYPE [RFC1035] section 3.2.3 of a DNS query in a DNS operation.
	CriteriaEnumPolicyCriteriaQTYPE CriteriaEnum = 5
	// DnsPolicyCriteriaTime: This specifies that the criteria for match of a DNS policy
	// is based on the time of day, in minutes, when the DNS server receives a DNS query
	// during a DNS operation.
	CriteriaEnumPolicyCriteriaTime CriteriaEnum = 6
	// DnsPolicyCriteriaMax: This shows the maximum number of criteria type supported by
	// DNS policy.
	CriteriaEnumPolicyCriteriaMax CriteriaEnum = 7
)

func (o CriteriaEnum) String() string {
	switch o {
	case CriteriaEnumPolicyCriteriaSubnet:
		return "CriteriaEnumPolicyCriteriaSubnet"
	case CriteriaEnumPolicyCriteriaTransportProtocol:
		return "CriteriaEnumPolicyCriteriaTransportProtocol"
	case CriteriaEnumPolicyCriteriaNetworkProtocol:
		return "CriteriaEnumPolicyCriteriaNetworkProtocol"
	case CriteriaEnumPolicyCriteriaInterface:
		return "CriteriaEnumPolicyCriteriaInterface"
	case CriteriaEnumPolicyCriteriaFQDN:
		return "CriteriaEnumPolicyCriteriaFQDN"
	case CriteriaEnumPolicyCriteriaQTYPE:
		return "CriteriaEnumPolicyCriteriaQTYPE"
	case CriteriaEnumPolicyCriteriaTime:
		return "CriteriaEnumPolicyCriteriaTime"
	case CriteriaEnumPolicyCriteriaMax:
		return "CriteriaEnumPolicyCriteriaMax"
	}
	return "Invalid"
}

// ClientSubnetRecord structure represents DNS_RPC_CLIENT_SUBNET_RECORD RPC structure.
//
// The DNS_RPC_CLIENT_SUBNET_RECORD structure contains the IPv4 and IPv6 subnets that
// are grouped together.
type ClientSubnetRecord struct {
	// pwszClientSubnetName: (variable) A NULL-terminated Unicode string containing the
	// name of the client subnet record.
	ClientSubnetName string `idl:"name:pwszClientSubnetName" json:"client_subnet_name"`
	// pIPAddr: A pointer to a DNS_ADDR_ARRAY (section 2.2.3.2.3) structure that contains
	// a list of IP subnets contained in this client subnet record along with the SubnetLength
	// as defined in DNS ADD USER (section 2.2.3.2.2.2).
	IPAddr *AddrArray `idl:"name:pIPAddr" json:"ip_addr"`
	// pIPv6Addr: A pointer to a DNS_ADDR_ARRAY structure that contains a list of IPv6 subnets
	// contained in this client subnet record along with the SubnetLength as defined in
	// DNS ADD USER.
	IPv6Addr *AddrArray `idl:"name:pIPv6Addr" json:"ipv6_addr"`
}

func (o *ClientSubnetRecord) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ClientSubnetRecord) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.ClientSubnetName != "" {
		_ptr_pwszClientSubnetName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.ClientSubnetName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ClientSubnetName, _ptr_pwszClientSubnetName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.IPAddr != nil {
		_ptr_pIPAddr := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.IPAddr != nil {
				if err := o.IPAddr.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AddrArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.IPAddr, _ptr_pIPAddr); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.IPv6Addr != nil {
		_ptr_pIPv6Addr := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.IPv6Addr != nil {
				if err := o.IPv6Addr.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AddrArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.IPv6Addr, _ptr_pIPv6Addr); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClientSubnetRecord) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_pwszClientSubnetName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.ClientSubnetName); err != nil {
			return err
		}
		return nil
	})
	_s_pwszClientSubnetName := func(ptr interface{}) { o.ClientSubnetName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ClientSubnetName, _s_pwszClientSubnetName, _ptr_pwszClientSubnetName); err != nil {
		return err
	}
	_ptr_pIPAddr := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.IPAddr == nil {
			o.IPAddr = &AddrArray{}
		}
		if err := o.IPAddr.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pIPAddr := func(ptr interface{}) { o.IPAddr = *ptr.(**AddrArray) }
	if err := w.ReadPointer(&o.IPAddr, _s_pIPAddr, _ptr_pIPAddr); err != nil {
		return err
	}
	_ptr_pIPv6Addr := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.IPv6Addr == nil {
			o.IPv6Addr = &AddrArray{}
		}
		if err := o.IPv6Addr.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pIPv6Addr := func(ptr interface{}) { o.IPv6Addr = *ptr.(**AddrArray) }
	if err := w.ReadPointer(&o.IPv6Addr, _s_pIPv6Addr, _ptr_pIPv6Addr); err != nil {
		return err
	}
	return nil
}

// PolicyContent structure represents DNS_RPC_POLICY_CONTENT RPC structure.
//
// The DNS_RPC_POLICY_CONTENT structure contains the zone scope, cache scope, or server
// scope details that are used to service the DNS operation if a DNS policy's DNS_RPC_POLICY_ACTION_TYPE
// (section 2.2.15.1.1.4) is DnsPolicyAllow.
type PolicyContent struct {
	// pwszScopeName: A NULL-terminated Unicode string containing the name of the zone scope,
	// cache scope or server scope.
	ScopeName string `idl:"name:pwszScopeName" json:"scope_name"`
	// dwWeight: A 32-bit unsigned integer that specifies the weight associated with the
	// scope name. The weight of a scope determines the number of responses that are sent
	// from that scope. For example, if a DNS_RPC_POLICY has pContentList with two DNS_RPC_POLICY_CONTENT
	// items having scope name and weight as {Scope1,2} and {Scope2,3}, then when the DNS
	// policy is a match, Scope1 is used to respond to the DNS operation for the first two
	// times the DNS policy is a match. Scope2 is used to respond to the DNS operation for
	// the next three times the DNS policy is a match. The sequence is repeated for further
	// matches of the DNS policy. If a dwWeight of a DNS_RPC_POLICY_CONTENT is not given.
	// its default weight is 1. So if two or more DNS_RPC_POLICY_CONTENT structures are
	// given during DNS_RPC_POLICY creation without a value for dwWeight, all of them will
	// get weight of 1 and the DNS operation will be performed from each scope in a round-robin
	// fashion. Allowed values for weight are any positive number from 1 and higher to 0xFFFFFFFF.
	// If 0 is sent as a weight, the DNS server returns the error DNS_ERROR_POLICY_INVALID_WEIGHT
	// (9981).
	Weight uint32 `idl:"name:dwWeight" json:"weight"`
}

func (o *PolicyContent) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PolicyContent) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.ScopeName != "" {
		_ptr_pwszScopeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.ScopeName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ScopeName, _ptr_pwszScopeName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Weight); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *PolicyContent) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_pwszScopeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.ScopeName); err != nil {
			return err
		}
		return nil
	})
	_s_pwszScopeName := func(ptr interface{}) { o.ScopeName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ScopeName, _s_pwszScopeName, _ptr_pwszScopeName); err != nil {
		return err
	}
	if err := w.ReadData(&o.Weight); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// PolicyContentList structure represents DNS_RPC_POLICY_CONTENT_LIST RPC structure.
//
// The DNS_RPC_POLICY_CONTENT_LIST structure contains a list of DNS_RPC_POLICY_CONTENT
// structures (section 2.2.15.2.2) configured for a DNS policy.
type PolicyContentList struct {
	// dwContentCount: An unsigned integer specifying the number of DNS_RPC_POLICY_CONTENT
	// elements present.
	ContentCount uint32 `idl:"name:dwContentCount" json:"content_count"`
	// pContent: An array of pointers to DNS_RPC_POLICY_CONTENT structures.
	Content []*PolicyContent `idl:"name:pContent;size_is:(dwContentCount)" json:"content"`
}

func (o *PolicyContentList) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Content != nil && o.ContentCount == 0 {
		o.ContentCount = uint32(len(o.Content))
	}
	if o.ContentCount > uint32(50000) {
		return fmt.Errorf("ContentCount is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *PolicyContentList) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.ContentCount)
	return []uint64{
		dimSize1,
	}
}
func (o *PolicyContentList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.ContentCount); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	for i1 := range o.Content {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.Content[i1] != nil {
			_ptr_pContent := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Content[i1] != nil {
					if err := o.Content[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PolicyContent{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Content[i1], _ptr_pContent); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Content); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyContentList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.ContentCount); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.ContentCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.ContentCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Content", sizeInfo[0])
	}
	o.Content = make([]*PolicyContent, sizeInfo[0])
	for i1 := range o.Content {
		i1 := i1
		_ptr_pContent := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Content[i1] == nil {
				o.Content[i1] = &PolicyContent{}
			}
			if err := o.Content[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pContent := func(ptr interface{}) { o.Content[i1] = *ptr.(**PolicyContent) }
		if err := w.ReadPointer(&o.Content[i1], _s_pContent, _ptr_pContent); err != nil {
			return err
		}
	}
	return nil
}

// Criteria structure represents DNS_RPC_CRITERIA RPC structure.
//
// The DNS_RPC_CRITERIA structure contains the DNS policy criteria that is associated
// with a DNS policy.
type Criteria struct {
	// type: This specifies the type of criteria associated with the DNS policy. The value
	// is of type DNS_RPC_CRITERIA_ENUM (section 2.2.15.1.1.6).
	Type CriteriaEnum `idl:"name:type" json:"type"`
	// pCriteria: A NULL-terminated Unicode string containing the DNS policy criteria details.
	Criteria string `idl:"name:pCriteria" json:"criteria"`
}

func (o *Criteria) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Criteria) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Type)); err != nil {
		return err
	}
	if o.Criteria != "" {
		_ptr_pCriteria := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.Criteria); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Criteria, _ptr_pCriteria); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Criteria) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Type)); err != nil {
		return err
	}
	_ptr_pCriteria := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.Criteria); err != nil {
			return err
		}
		return nil
	})
	_s_pCriteria := func(ptr interface{}) { o.Criteria = *ptr.(*string) }
	if err := w.ReadPointer(&o.Criteria, _s_pCriteria, _ptr_pCriteria); err != nil {
		return err
	}
	return nil
}

// Policy structure represents DNS_RPC_POLICY RPC structure.
//
// The DNS_RPC_POLICY structure contains all details associated with a DNS policy.
type Policy struct {
	// pwszPolicyName: A NULL-terminated Unicode string containing the name of the DNS policy.
	// The name is unique across a level as specified in DNS_RPC_POLICY_LEVEL (section 2.2.15.1.1.3).
	PolicyName string `idl:"name:pwszPolicyName" json:"policy_name"`
	// level: This specifies whether the policy is configured for a zone (including a cache
	// zone) or is at server level. The value is of type DNS_RPC_POLICY_LEVEL.
	Level PolicyLevel `idl:"name:level" json:"level"`
	// appliesOn: This specifies the DNS operation for which the DNS policy is applicable.
	// The value is of type DNS_RPC_POLICY_TYPE (section 2.2.15.1.1.5).
	AppliesOn PolicyType `idl:"name:appliesOn" json:"applies_on"`
	// action: This specifies the action to be applied on a DNS operation when there is
	// a DNS policy match. The value is of type DNS_RPC_POLICY_ACTION_TYPE (section 2.2.15.1.1.4).
	Action PolicyActionType `idl:"name:action" json:"action"`
	// condition: This specifies the logical operation across the DNS policy criteria of
	// a DNS policy, which determines whether the DNS policy match has succeeded or failed.
	// The value is of type DNS_RPC_POLICY_CONDITION (section 2.2.15.1.1.2).
	Condition PolicyCondition `idl:"name:condition" json:"condition"`
	// isEnabled: A BOOL variable that specifies whether a DNS policy is enabled. TRUE indicates
	// that the DNS policy MUST be applied to the DNS operation; FALSE indicates that the
	// DNS policy is disabled and MUST NOT be applied to the DNS operation.
	IsEnabled bool `idl:"name:isEnabled" json:"is_enabled"`
	// dwProcessingOrder: An unsigned integer that specifies the order in which this DNS
	// policy is processed, along with 0 or more other DNS policies for a given level 2.2.15.1.1.3
	// and DNS_RPC_POLICY_TYPE (section 2.2.15.1.1.5) during a DNS operation.
	ProcessingOrder uint32 `idl:"name:dwProcessingOrder" json:"processing_order"`
	// pszZoneName: A NULL-terminated string containing the name of the zone for which this
	// DNS policy is configured.
	ZoneName string `idl:"name:pszZoneName" json:"zone_name"`
	// pContentList: An object of type DNS_RPC_POLICY_CONTENT_LIST (section 2.2.15.2.3)
	// that has a list of zone scope, cache scope, or server scope with weight as defined
	// in DNS_RPC_POLICY_CONTENT (section 2.2.15.2.2). If the DNS policy matches the criteria
	// in pCriteriaList and its action is DnsPolicyAllow, then the zone scopes, cache scopes,
	// or server scopes contained in pContentList are used in the DNS operation as per their
	// dwWeight.
	ContentList *PolicyContentList `idl:"name:pContentList" json:"content_list"`
	// flags: This is used during update of a DNS policy. The various bits of flags show
	// which members of the DNS Policy are to be updated. For possible values of the bits
	// of flags. see section 2.2.15.1.1. For details on how this is used during a policy
	// update, see the operation UpdatePolicy in section 3.1.4.1.
	Flags uint64 `idl:"name:flags" json:"flags"`
	// dwCriteriaCount: An unsigned integer containing the number of DNS policy criteria
	// that are configured for this DNS policy.
	CriteriaCount uint32 `idl:"name:dwCriteriaCount" json:"criteria_count"`
	// pCriteriaList: An array of DNS_RPC_CRITERIA (section 2.2.15.2.4) that are compared
	// with each other using condition DNS_RPC_POLICY_CONDITION (section 2.2.15.1.1.2) to
	// match a DNS policy.
	CriteriaList []*Criteria `idl:"name:pCriteriaList;size_is:(dwCriteriaCount)" json:"criteria_list"`
}

func (o *Policy) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.CriteriaList != nil && o.CriteriaCount == 0 {
		o.CriteriaCount = uint32(len(o.CriteriaList))
	}
	if o.CriteriaCount > uint32(50000) {
		return fmt.Errorf("CriteriaCount is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *Policy) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.CriteriaCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Policy) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.PolicyName != "" {
		_ptr_pwszPolicyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.PolicyName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.PolicyName, _ptr_pwszPolicyName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteEnum(uint16(o.Level)); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.AppliesOn)); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Action)); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Condition)); err != nil {
		return err
	}
	if !o.IsEnabled {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ProcessingOrder); err != nil {
		return err
	}
	if o.ZoneName != "" {
		_ptr_pszZoneName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharString(ctx, w, o.ZoneName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ZoneName, _ptr_pszZoneName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ContentList != nil {
		_ptr_pContentList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ContentList != nil {
				if err := o.ContentList.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&PolicyContentList{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ContentList, _ptr_pContentList); err != nil {
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
	if err := w.WriteData(o.CriteriaCount); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	for i1 := range o.CriteriaList {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.CriteriaList[i1] != nil {
			_ptr_pCriteriaList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CriteriaList[i1] != nil {
					if err := o.CriteriaList[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&Criteria{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CriteriaList[i1], _ptr_pCriteriaList); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.CriteriaList); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Policy) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	_ptr_pwszPolicyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.PolicyName); err != nil {
			return err
		}
		return nil
	})
	_s_pwszPolicyName := func(ptr interface{}) { o.PolicyName = *ptr.(*string) }
	if err := w.ReadPointer(&o.PolicyName, _s_pwszPolicyName, _ptr_pwszPolicyName); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Level)); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.AppliesOn)); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Action)); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Condition)); err != nil {
		return err
	}
	var _bIsEnabled int32
	if err := w.ReadData(&_bIsEnabled); err != nil {
		return err
	}
	o.IsEnabled = _bIsEnabled != 0
	if err := w.ReadData(&o.ProcessingOrder); err != nil {
		return err
	}
	_ptr_pszZoneName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharString(ctx, w, &o.ZoneName); err != nil {
			return err
		}
		return nil
	})
	_s_pszZoneName := func(ptr interface{}) { o.ZoneName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ZoneName, _s_pszZoneName, _ptr_pszZoneName); err != nil {
		return err
	}
	_ptr_pContentList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ContentList == nil {
			o.ContentList = &PolicyContentList{}
		}
		if err := o.ContentList.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pContentList := func(ptr interface{}) { o.ContentList = *ptr.(**PolicyContentList) }
	if err := w.ReadPointer(&o.ContentList, _s_pContentList, _ptr_pContentList); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.CriteriaCount); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.CriteriaCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.CriteriaCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.CriteriaList", sizeInfo[0])
	}
	o.CriteriaList = make([]*Criteria, sizeInfo[0])
	for i1 := range o.CriteriaList {
		i1 := i1
		_ptr_pCriteriaList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.CriteriaList[i1] == nil {
				o.CriteriaList[i1] = &Criteria{}
			}
			if err := o.CriteriaList[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pCriteriaList := func(ptr interface{}) { o.CriteriaList[i1] = *ptr.(**Criteria) }
		if err := w.ReadPointer(&o.CriteriaList[i1], _s_pCriteriaList, _ptr_pCriteriaList); err != nil {
			return err
		}
	}
	return nil
}

// PolicyName structure represents DNS_RPC_POLICY_NAME RPC structure.
//
// The DNS_RPC_POLICY_NAME structure contains the details of the DNS policy when the
// DNS policies are enumerated at a specified level (section 2.2.15.1.1.3).
type PolicyName struct {
	// pwszPolicyName: A NULL-terminated Unicode string containing the name of the DNS policy.
	// The name is unique across a level as specified in DNS_RPC_POLICY_LEVEL (section 2.2.15.1.1.3).
	PolicyName string `idl:"name:pwszPolicyName" json:"policy_name"`
	// appliesOn: This specifies the DNS operation to which the DNS policy applies. The
	// value is of type DNS_RPC_POLICY_TYPE (section 2.2.15.1.1.5).
	AppliesOn PolicyType `idl:"name:appliesOn" json:"applies_on"`
	// fEnabled: A BOOL variable that specifies whether a DNS policy is enabled.
	Enabled bool `idl:"name:fEnabled" json:"enabled"`
	// processingOrder: An unsigned integer that specifies the order in which this DNS policy
	// is processed, along with 0 or more other DNS policies for a given level (section
	// 2.2.15.1.1.3) and DNS_RPC_POLICY_TYPE (section 2.2.15.1.1.5) during a DNS operation.
	ProcessingOrder uint32 `idl:"name:processingOrder" json:"processing_order"`
}

func (o *PolicyName) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PolicyName) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.PolicyName != "" {
		_ptr_pwszPolicyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.PolicyName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.PolicyName, _ptr_pwszPolicyName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteEnum(uint16(o.AppliesOn)); err != nil {
		return err
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
	if err := w.WriteData(o.ProcessingOrder); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *PolicyName) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_pwszPolicyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.PolicyName); err != nil {
			return err
		}
		return nil
	})
	_s_pwszPolicyName := func(ptr interface{}) { o.PolicyName = *ptr.(*string) }
	if err := w.ReadPointer(&o.PolicyName, _s_pwszPolicyName, _ptr_pwszPolicyName); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.AppliesOn)); err != nil {
		return err
	}
	var _bEnabled int32
	if err := w.ReadData(&_bEnabled); err != nil {
		return err
	}
	o.Enabled = _bEnabled != 0
	if err := w.ReadData(&o.ProcessingOrder); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// EnumeratePolicyList structure represents DNS_RPC_ENUMERATE_POLICY_LIST RPC structure.
//
// The DNS_RPC_ENUMERATE_POLICY_LIST structure contains all the DNS policies at a specified
// level (section 2.2.15.1.1.3).
type EnumeratePolicyList struct {
	// dwPolicyCount: An unsigned integer that specifies the total number of DNS_RPC_POLICY_NAME
	// structures (section 2.2.15.2.6).
	PolicyCount uint32 `idl:"name:dwPolicyCount" json:"policy_count"`
	// pPolicyArray: An array of pointers of DNS_RPC_POLICY_NAME (section 2.2.15.2.6) structures.
	PolicyArray []*PolicyName `idl:"name:pPolicyArray;size_is:(dwPolicyCount)" json:"policy_array"`
}

func (o *EnumeratePolicyList) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.PolicyArray != nil && o.PolicyCount == 0 {
		o.PolicyCount = uint32(len(o.PolicyArray))
	}
	if o.PolicyCount > uint32(50000) {
		return fmt.Errorf("PolicyCount is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *EnumeratePolicyList) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.PolicyCount)
	return []uint64{
		dimSize1,
	}
}
func (o *EnumeratePolicyList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.PolicyCount); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	for i1 := range o.PolicyArray {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.PolicyArray[i1] != nil {
			_ptr_pPolicyArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PolicyArray[i1] != nil {
					if err := o.PolicyArray[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PolicyName{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PolicyArray[i1], _ptr_pPolicyArray); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.PolicyArray); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *EnumeratePolicyList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.PolicyCount); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.PolicyCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.PolicyCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.PolicyArray", sizeInfo[0])
	}
	o.PolicyArray = make([]*PolicyName, sizeInfo[0])
	for i1 := range o.PolicyArray {
		i1 := i1
		_ptr_pPolicyArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PolicyArray[i1] == nil {
				o.PolicyArray[i1] = &PolicyName{}
			}
			if err := o.PolicyArray[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pPolicyArray := func(ptr interface{}) { o.PolicyArray[i1] = *ptr.(**PolicyName) }
		if err := w.ReadPointer(&o.PolicyArray[i1], _s_pPolicyArray, _ptr_pPolicyArray); err != nil {
			return err
		}
	}
	return nil
}

// RRLModeEnum type represents DNS_RRL_MODE_ENUM RPC enumeration.
//
// The DNS_RRL_MODE_ENUM enumeration controls the behavior and application of DNS Response
// Rate Limiting at the server level.
type RRLModeEnum uint16

var (
	// DnsRRLLogOnly: This is the simulation mode for RRL, where the impact of RRL is logged
	// in analytical channel logs but no actual action is taken (that is, responses are
	// not dropped).
	RRLModeEnumLogOnly RRLModeEnum = 0
	// DnsRRLEnabled: In this mode, the RRL algorithm is applied to every incoming query.
	RRLModeEnumEnabled RRLModeEnum = 1
	// DnsRRLDisabled: This mode ensures RRL is disabled and not applied to any incoming
	// query.
	RRLModeEnumDisabled RRLModeEnum = 2
)

func (o RRLModeEnum) String() string {
	switch o {
	case RRLModeEnumLogOnly:
		return "RRLModeEnumLogOnly"
	case RRLModeEnumEnabled:
		return "RRLModeEnumEnabled"
	case RRLModeEnumDisabled:
		return "RRLModeEnumDisabled"
	}
	return "Invalid"
}

// RRLParams structure represents DNS_RPC_RRL_PARAMS RPC structure.
//
// The DNS_RPC_RRL_PARAMS structure SHOULD<93> contain the configuration parameters
// for the Response Rate Limiting capability configured on the DNS server.
type RRLParams struct {
	// dwResponsesPerSecond: The maximum number of responses a DNS server can give for each
	// successful "unique response" in one-second intervals. A DNS response is considered
	// a unique response if the combination of the following parameters is unique: the requestor's
	// IP address, masked according to either dwIPv4PrefixLength or dwIPv6PrefixLength;
	// an imputed domain name that is either a wildcard (if a wildcard match occurred),
	// the zone name (if no match occurred), or the query name; and a Boolean error indicator
	// (response code Refused, FormError, or ServFail).
	//
	// This parameter can be set any positive integer; the default value is 5.
	ResponsesPerSecond uint32 `idl:"name:dwResponsesPerSecond" json:"responses_per_second"`
	// dwErrorsPerSecond: The maximum number of responses a DNS server can give for queries
	// resulting in error (ServFail, FormError, Refused) in one-second intervals. This parameter
	// can be set to any positive integer; the default value is 5.
	ErrorsPerSecond uint32 `idl:"name:dwErrorsPerSecond" json:"errors_per_second"`
	// dwLeakRate: When a query would be dropped due to rate limiting, the DNS server responds
	// once per Leak Rate query. The default value for this parameter is 3. This parameter
	// can be set to any positive integer greater than 1. If the value is set to 0, this
	// behavior is disabled.
	LeakRate uint32 `idl:"name:dwLeakRate" json:"leak_rate"`
	// dwTCRate: When a query would be dropped due to rate limiting, the DNS server returns
	// a truncated response once per TC rate query. The default value for this parameter
	// is 2. The parameter can be set to any positive integer greater than 1. If the value
	// is set to 0, the behavior is disabled (see [RRL] section 2.2.8).
	TCRate uint32 `idl:"name:dwTCRate" json:"tc_rate"`
	// dwTotalResponsesInWindow: The maximum number of responses the DNS server can give
	// for each "unique response" in the defined window duration. See dwResponsesPerSecond
	// for the definition of "unique response". This includes responses given from leak
	// rate, truncation rate, and responses per second / errors per second. The default
	// value for this parameter is 1024. This parameter can be set any positive integer.
	TotalResponsesInWindow uint32 `idl:"name:dwTotalResponsesInWindow" json:"total_responses_in_window"`
	// dwWindowSize: The duration, in seconds, where the state of dwTotalResponsesInWindow
	// is maintained for each "unique response". See dwResponsesPerSecond for the definition
	// of "unique response". After this duration, the value for dwTotalResponsesInWindow
	// is reset to 0. The default value for this parameter is 5. The parameter can be set
	// to any positive integer (see [RRL] section 2.2.4).
	WindowSize uint32 `idl:"name:dwWindowSize" json:"window_size"`
	// dwIPv4PrefixLength: Controls how the DNS query source IPv4 addresses are grouped
	// into buckets of size (32 – dwIPv4PrefixLength) ^ 2. The default value for this
	// parameter is 24. The parameter can be set to any positive integer between 0 and 32.
	IPv4PrefixLength uint32 `idl:"name:dwIPv4PrefixLength" json:"ipv4_prefix_length"`
	// dwIPv6PrefixLength: Controls how DNS query source IPv6 addresses are grouped into
	// buckets of size (32 – dwIPv6PrefixLength) ^ 2. The default value for this parameter
	// is 56. This parameter can be set any positive integer between 0 and 128.
	IPv6PrefixLength uint32 `idl:"name:dwIPv6PrefixLength" json:"ipv6_prefix_length"`
	// eMode: The mode in which RRL functions on a DNS server.
	Mode RRLModeEnum `idl:"name:eMode" json:"mode"`
	// dwFlags: This is used during the update of a DNS Response Rate Limiting. The various
	// bits show which members of the DNS RRL are to be updated. For possible values, see
	// section 2.2.16.1.1. For details on how this is used for setting RRL, see the operation
	// SetRRL in section 3.1.4.1. If dwFlags is not set for a configuration parameter, default
	// values are applied as shown in the following table:
	//
	//	+-----------------------------+-----------------+
	//	|      RRL CONFIGURATION      |     DEFAULT     |
	//	|          PARAMETER          |     VALUES      |
	//	+-----------------------------+-----------------+
	//	+-----------------------------+-----------------+
	//	| dwResponsesPerSecond        |               5 |
	//	+-----------------------------+-----------------+
	//	| dwErrorsPerSecond           |               5 |
	//	+-----------------------------+-----------------+
	//	| dwLeakRate                  |               3 |
	//	+-----------------------------+-----------------+
	//	| dwTCRate                    |               2 |
	//	+-----------------------------+-----------------+
	//	| dwTotalResponsesInWindow    |            1024 |
	//	+-----------------------------+-----------------+
	//	| dwWindowSize                |               5 |
	//	+-----------------------------+-----------------+
	//	| dwIPv4PrefixLength          |              24 |
	//	+-----------------------------+-----------------+
	//	| dwIPv6PrefixLength          |              56 |
	//	+-----------------------------+-----------------+
	//	| eMode                       | DnsRRLDisabled  |
	//	+-----------------------------+-----------------+
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
	// fSetDefault: Set this value to TRUE to set RRL parameters to their default values.
	// This parameter does not affect the eMode parameter of RRL.
	SetDefault bool `idl:"name:fSetDefault" json:"set_default"`
}

func (o *RRLParams) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RRLParams) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ResponsesPerSecond); err != nil {
		return err
	}
	if err := w.WriteData(o.ErrorsPerSecond); err != nil {
		return err
	}
	if err := w.WriteData(o.LeakRate); err != nil {
		return err
	}
	if err := w.WriteData(o.TCRate); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalResponsesInWindow); err != nil {
		return err
	}
	if err := w.WriteData(o.WindowSize); err != nil {
		return err
	}
	if err := w.WriteData(o.IPv4PrefixLength); err != nil {
		return err
	}
	if err := w.WriteData(o.IPv6PrefixLength); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Mode)); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if !o.SetDefault {
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
func (o *RRLParams) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ResponsesPerSecond); err != nil {
		return err
	}
	if err := w.ReadData(&o.ErrorsPerSecond); err != nil {
		return err
	}
	if err := w.ReadData(&o.LeakRate); err != nil {
		return err
	}
	if err := w.ReadData(&o.TCRate); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalResponsesInWindow); err != nil {
		return err
	}
	if err := w.ReadData(&o.WindowSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPv4PrefixLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPv6PrefixLength); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Mode)); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	var _bSetDefault int32
	if err := w.ReadData(&_bSetDefault); err != nil {
		return err
	}
	o.SetDefault = _bSetDefault != 0
	return nil
}

// VirtualizationInstance structure represents DNS_RPC_VIRTUALIZATION_INSTANCE RPC structure.
//
// The following example describes how to create a new virtualization instance named
// "example_vi1" in the DNS server.
//
// The client calls R_DnssrvOperation4 (section 3.1.4.16) and provides the following
// parameters:
//
// * *DNS_RPC_CURRENT_CLIENT_VER* (section 2.2.1.2.1 ( a237781b-b048-495d-844b-10ffb38df2a8
// ) ) as the client version.
//
// * Zero as the settings flag.
//
// * A Unicode string containing the FQDN of the DNS server on which the virtualization
// instance is to be created.
//
// * Zero as the context.
//
// * CreateVirtualizationInstance as the operation.
//
// * DNSSRV_TYPEID_VIRTUALIZATION_INSTANCE as the type ID.
//
// * Create a DNS_RPC_VIRTUALIZATION_INSTANCE (section 2.2.17.1.1 ( 44d991d5-392a-48e7-827f-a1ccbe22ddef
// ) ) structure with dwFlags as zero and pwszVirtualizationID as "example_vi1". Pass
// this structure as RPC data.
//
// The DNS server returns ERROR_SUCCESS if the operation was successful or a Windows
// error code if the operation fails.
type VirtualizationInstance struct {
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	_                   uint32 `idl:"name:dwReserved"`
	Flags               uint32 `idl:"name:dwFlags" json:"flags"`
	VirtualizationID    string `idl:"name:pwszVirtualizationID" json:"virtualization_id"`
	FriendlyName        string `idl:"name:pwszFriendlyName" json:"friendly_name"`
	Description         string `idl:"name:pwszDescription" json:"description"`
}

func (o *VirtualizationInstance) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *VirtualizationInstance) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.VirtualizationID != "" {
		_ptr_pwszVirtualizationID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.VirtualizationID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.VirtualizationID, _ptr_pwszVirtualizationID); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.FriendlyName != "" {
		_ptr_pwszFriendlyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.FriendlyName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.FriendlyName, _ptr_pwszFriendlyName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Description != "" {
		_ptr_pwszDescription := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.Description); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Description, _ptr_pwszDescription); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *VirtualizationInstance) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	// reserved dwReserved
	var _dwReserved uint32
	if err := w.ReadData(&_dwReserved); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_pwszVirtualizationID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.VirtualizationID); err != nil {
			return err
		}
		return nil
	})
	_s_pwszVirtualizationID := func(ptr interface{}) { o.VirtualizationID = *ptr.(*string) }
	if err := w.ReadPointer(&o.VirtualizationID, _s_pwszVirtualizationID, _ptr_pwszVirtualizationID); err != nil {
		return err
	}
	_ptr_pwszFriendlyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.FriendlyName); err != nil {
			return err
		}
		return nil
	})
	_s_pwszFriendlyName := func(ptr interface{}) { o.FriendlyName = *ptr.(*string) }
	if err := w.ReadPointer(&o.FriendlyName, _s_pwszFriendlyName, _ptr_pwszFriendlyName); err != nil {
		return err
	}
	_ptr_pwszDescription := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.Description); err != nil {
			return err
		}
		return nil
	})
	_s_pwszDescription := func(ptr interface{}) { o.Description = *ptr.(*string) }
	if err := w.ReadPointer(&o.Description, _s_pwszDescription, _ptr_pwszDescription); err != nil {
		return err
	}
	return nil
}

// VirtualizationInstanceInfo structure represents DNS_RPC_VIRTUALIZATION_INSTANCE_INFO RPC structure.
//
// The DNS_RPC_VIRTUALIZATION_INSTANCE_INFO structure contains the details of a virtualization
// instance when the virtualization instances in the DNS server are enumerated (section
// 3.1.4.15).
type VirtualizationInstanceInfo struct {
	// pwszVirtualizationID: A NULL-terminated Unicode string that uniquely identifies a
	// particular virtualization instance in the DNS server. The maximum length of this
	// identifier is limited to 64 characters.
	VirtualizationID string `idl:"name:pwszVirtualizationID" json:"virtualization_id"`
	// pwszFriendlyName: A NULL-terminated Unicode string that contains a user-friendly
	// name of the DNS virtualization instance.
	FriendlyName string `idl:"name:pwszFriendlyName" json:"friendly_name"`
	// pwszDescription: A NULL-terminated Unicode string that contains a description of
	// the DNS virtualization instance.
	Description string `idl:"name:pwszDescription" json:"description"`
}

func (o *VirtualizationInstanceInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *VirtualizationInstanceInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.VirtualizationID != "" {
		_ptr_pwszVirtualizationID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.VirtualizationID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.VirtualizationID, _ptr_pwszVirtualizationID); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.FriendlyName != "" {
		_ptr_pwszFriendlyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.FriendlyName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.FriendlyName, _ptr_pwszFriendlyName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Description != "" {
		_ptr_pwszDescription := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.Description); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Description, _ptr_pwszDescription); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *VirtualizationInstanceInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_pwszVirtualizationID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.VirtualizationID); err != nil {
			return err
		}
		return nil
	})
	_s_pwszVirtualizationID := func(ptr interface{}) { o.VirtualizationID = *ptr.(*string) }
	if err := w.ReadPointer(&o.VirtualizationID, _s_pwszVirtualizationID, _ptr_pwszVirtualizationID); err != nil {
		return err
	}
	_ptr_pwszFriendlyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.FriendlyName); err != nil {
			return err
		}
		return nil
	})
	_s_pwszFriendlyName := func(ptr interface{}) { o.FriendlyName = *ptr.(*string) }
	if err := w.ReadPointer(&o.FriendlyName, _s_pwszFriendlyName, _ptr_pwszFriendlyName); err != nil {
		return err
	}
	_ptr_pwszDescription := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.Description); err != nil {
			return err
		}
		return nil
	})
	_s_pwszDescription := func(ptr interface{}) { o.Description = *ptr.(*string) }
	if err := w.ReadPointer(&o.Description, _s_pwszDescription, _ptr_pwszDescription); err != nil {
		return err
	}
	return nil
}

// EnumVirtualizationInstanceList structure represents DNS_RPC_ENUM_VIRTUALIZATION_INSTANCE_LIST RPC structure.
//
// The DNS_RPC_ENUM_VIRTUALIZATION_INSTANCE_LIST structure contains information about
// all the virtualization instances in a DNS server (section 3.1.4.15).
type EnumVirtualizationInstanceList struct {
	// dwRpcStructureVersion: The DNS management structure version number. This value MUST
	// be set to 0x00000000.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// dwVirtualizationInstanceCount: An unsigned integer that specifies the total number
	// of DNS_RPC_VIRTUALIZATION_INSTANCE_INFO (section 2.2.17.1.2) elements.
	VirtualizationInstanceCount uint32 `idl:"name:dwVirtualizationInstanceCount" json:"virtualization_instance_count"`
	// VirtualizationInstanceArray: An array of pointers to DNS_RPC_VIRTUALIZATION_INSTANCE_INFO
	// (section 2.2.17.1.2) structures.
	VirtualizationInstanceArray []*VirtualizationInstanceInfo `idl:"name:VirtualizationInstanceArray;size_is:(dwVirtualizationInstanceCount)" json:"virtualization_instance_array"`
}

func (o *EnumVirtualizationInstanceList) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.VirtualizationInstanceArray != nil && o.VirtualizationInstanceCount == 0 {
		o.VirtualizationInstanceCount = uint32(len(o.VirtualizationInstanceArray))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *EnumVirtualizationInstanceList) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.VirtualizationInstanceCount)
	return []uint64{
		dimSize1,
	}
}
func (o *EnumVirtualizationInstanceList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.VirtualizationInstanceCount); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	for i1 := range o.VirtualizationInstanceArray {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.VirtualizationInstanceArray[i1] != nil {
			_ptr_VirtualizationInstanceArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.VirtualizationInstanceArray[i1] != nil {
					if err := o.VirtualizationInstanceArray[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&VirtualizationInstanceInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.VirtualizationInstanceArray[i1], _ptr_VirtualizationInstanceArray); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.VirtualizationInstanceArray); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *EnumVirtualizationInstanceList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.VirtualizationInstanceCount); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.VirtualizationInstanceCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.VirtualizationInstanceCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.VirtualizationInstanceArray", sizeInfo[0])
	}
	o.VirtualizationInstanceArray = make([]*VirtualizationInstanceInfo, sizeInfo[0])
	for i1 := range o.VirtualizationInstanceArray {
		i1 := i1
		_ptr_VirtualizationInstanceArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.VirtualizationInstanceArray[i1] == nil {
				o.VirtualizationInstanceArray[i1] = &VirtualizationInstanceInfo{}
			}
			if err := o.VirtualizationInstanceArray[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_VirtualizationInstanceArray := func(ptr interface{}) { o.VirtualizationInstanceArray[i1] = *ptr.(**VirtualizationInstanceInfo) }
		if err := w.ReadPointer(&o.VirtualizationInstanceArray[i1], _s_VirtualizationInstanceArray, _ptr_VirtualizationInstanceArray); err != nil {
			return err
		}
	}
	return nil
}

// TypeID type represents DNS_RPC_TYPEID RPC enumeration.
//
// The DNS Server Management Protocol RPC methods use a generic and extensible data
// structure of type DNSSRV_RPC_UNION structure (section 2.2.1.2.6), which is a union
// of pointers to different data types. A DNS_RPC_TYPEID enumeration value is used to
// specify what data is being stored in an instance of the DNSSRV_RPC_UNION structure.
// When combined with a DNSSRV_RPC_UNION structure, the DNS_RPC_TYPEID enumeration allows
// the DNS RPC interface to communicate many different types of DNS server configuration
// and data in a single structure.
type TypeID uint16

var (
	// DNSSRV_TYPEID_ANY: Type is invalid.
	TypeIDAny TypeID = 1
	// DNSSRV_TYPEID_NULL: No data is provided.
	TypeIDNull TypeID = 0
	// DNSSRV_TYPEID_DWORD: A DWORD value.
	TypeIDDword TypeID = 1
	// DNSSRV_TYPEID_LPSTR: A pointer to a null-terminated UTF-8 [RFC3629] string.
	TypeIDString TypeID = 2
	// DNSSRV_TYPEID_LPWSTR: A pointer to a null-terminated Unicode string.
	TypeIDUnicodeString TypeID = 3
	// DNSSRV_TYPEID_IPARRAY: A pointer to an IP4_ARRAY. This structure is used to specify
	// a list of IPv4 addresses.
	TypeIDIPArray TypeID = 4
	// DNSSRV_TYPEID_BUFFER: A pointer to a DNS_RPC_BUFFER (section 2.2.1.2.2). This structure
	// is used to hold a generic buffer of the DNS server information. Interpretation of
	// the buffer depends upon the request.
	TypeIDBuffer TypeID = 5
	// DNSSRV_TYPEID_SERVER_INFO_W2K: A pointer to a structure of type DNS_RPC_SERVER_INFO_W2K
	// (section 2.2.4.2.2.1). This structure is used to specify general DNS server state
	// and configuration.
	TypeIDServerInfoW2K TypeID = 6
	// DNSSRV_TYPEID_STATS: A pointer to a structure of type DNSSRV_STATS (section 2.2.10.2.2).
	// The structure exposes internal statistics and counters.
	TypeIDStats TypeID = 7
	// DNSSRV_TYPEID_FORWARDERS_W2K: A pointer to a structure of type DNS_RPC_FORWARDERS_W2K
	// (section 2.2.5.2.10.1). This structure specifies the set of DNS servers this DNS
	// server will forward unresolved queries to.
	TypeIDForwardersW2K TypeID = 8
	// DNSSRV_TYPEID_ZONE_W2K: A pointer to a structure of type DNS_RPC_ZONE_W2K (section
	// 2.2.5.2.1.1). This structure is used to specify basic information about a DNS zone.
	TypeIDZoneW2K TypeID = 9
	// DNSSRV_TYPEID_ZONE_INFO_W2K: A pointer to a structure of type DNS_RPC_ZONE_INFO_W2K
	// (section 2.2.5.2.4.1). This structure is used to specify detailed DNS zone information.
	TypeIDZoneInfoW2K TypeID = 10
	// DNSSRV_TYPEID_ZONE_SECONDARIES_W2K: A pointer to a structure of type DNS_RPC_ZONE_SECONDARIES_W2K
	// (section 2.2.5.2.5.1). This structure is used to specify information about the secondary
	// servers for a primary DNS zone.
	TypeIDZoneSecondariesW2K TypeID = 11
	// DNSSRV_TYPEID_ZONE_DATABASE_W2K: A pointer to a structure of type DNS_RPC_ZONE_DATABASE_W2K
	// (section 2.2.5.2.6.1). This structure specifies how a DNS zone is stored in persistent
	// storage.
	TypeIDZoneDatabaseW2K TypeID = 12
	// DNSSRV_TYPEID_ZONE_TYPE_RESET_W2K: This value is not used.
	TypeIDZoneTypeResetW2K TypeID = 13
	// DNSSRV_TYPEID_ZONE_CREATE_W2K: A pointer to a structure of type DNS_RPC_ZONE_CREATE_INFO_W2K
	// (section 2.2.5.2.7.1). This structure is used to specify parameters required when
	// creating a new DNS zone.
	TypeIDZoneCreateW2K TypeID = 14
	// DNSSRV_TYPEID_NAME_AND_PARAM: A pointer to a structure of type DNS_RPC_NAME_AND_PARAM
	// (section 2.2.1.2.5). This is a general purpose structure used to associate a parameter
	// name with a DWORD value.
	TypeIDNameAndParam TypeID = 15
	// DNSSRV_TYPEID_ZONE_LIST_W2K: A pointer to a structure of type DNS_RPC_ZONE_LIST_W2K
	// (section 2.2.5.2.3.1). This structure is used to enumerate zones.
	TypeIDZoneListW2K TypeID = 16
	// DNSSRV_TYPEID_ZONE_RENAME: This value is not used.
	TypeIDZoneRename TypeID = 17
	// DNSSRV_TYPEID_ZONE_EXPORT: A pointer to a structure of type DNS_RPC_ZONE_EXPORT_INFO
	// (section 2.2.5.2.8). This structure is used to specify how to export a zone to a
	// file.
	TypeIDZoneExport TypeID = 18
	// DNSSRV_TYPEID_SERVER_INFO_DOTNET: A pointer to a structure of type DNS_RPC_SERVER_INFO_DOTNET
	// (section 2.2.4.2.2.2). This structure is used to specify general DNS server state
	// and configuration.
	TypeIDServerInfoNet TypeID = 19
	// DNSSRV_TYPEID_FORWARDERS_DOTNET: A pointer to a structure of type DNS_RPC_FORWARDERS_DOTNET
	// (section 2.2.5.2.10.2). This structure specifies the set of DNS servers this DNS
	// server will forward unresolved queries to.
	TypeIDForwardersNet TypeID = 20
	// DNSSRV_TYPEID_ZONE: A pointer to a structure of type DNS_RPC_ZONE (section 2.2.5.2.1).
	// This structure is used to specify basic information and a DNS zone.
	TypeIDZone TypeID = 21
	// DNSSRV_TYPEID_ZONE_INFO_DOTNET: A pointer to a structure of type DNS_RPC_ZONE_INFO_DOTNET
	// (section 2.2.5.2.4.2). This structure is used to specify detailed information about
	// a DNS zone.
	TypeIDZoneInfoNet TypeID = 22
	// DNSSRV_TYPEID_ZONE_SECONDARIES_DOTNET: A pointer to a structure of type DNS_RPC_ZONE_SECONDARIES_DOTNET
	// (section 2.2.5.2.5.2). This structure is used to specify information about the secondary
	// servers for a primary DNS zone.
	TypeIDZoneSecondariesNet TypeID = 23
	// DNSSRV_TYPEID_ZONE_DATABASE: A pointer to a structure of type DNS_RPC_ZONE_DATABASE
	// (section 2.2.5.2.6). This structure specifies how a DNS zone is stored in persistent
	// storage.
	TypeIDZoneDatabase TypeID = 24
	// DNSSRV_TYPEID_ZONE_TYPE_RESET_DOTNET: This value is not used.
	TypeIDZoneTypeResetNet TypeID = 25
	// DNSSRV_TYPEID_ZONE_CREATE_DOTNET: A pointer to a structure of type DNS_RPC_ZONE_CREATE_INFO_DOTNET.
	// This structure is used to specify parameters required when creating a new DNS zone.
	TypeIDZoneCreateNet TypeID = 26
	// DNSSRV_TYPEID_ZONE_LIST: A pointer to a structure of type DNS_RPC_ZONE_LIST (section
	// 2.2.5.2.3). This structure is used to enumerate zones.
	TypeIDZoneList TypeID = 27
	// DNSSRV_TYPEID_DP_ENUM: A pointer to a structure of type DNS_RPC_DP_ENUM (section
	// 2.2.7.2.3). This structure is used to specify basic information about an application
	// directory partition.
	TypeIDIddpEnum TypeID = 28
	// DNSSRV_TYPEID_DP_INFO: A pointer to a structure of type DNS_RPC_DP_INFO (section
	// 2.2.7.2.1). This structure specifies detailed information about a single application
	// directory partition.
	TypeIDIddpInfo TypeID = 29
	// DNSSRV_TYPEID_DP_LIST: A pointer to a structure of type DNS_RPC_DP_LIST (section
	// 2.2.7.2.4). This structure is used to enumerate application directory partitions.
	TypeIDIddpList TypeID = 30
	// DNSSRV_TYPEID_ENLIST_DP: A pointer to a structure of type DNS_RPC_ENLIST_DP (section
	// 2.2.7.2.5). This structure is used to request enlistment changes for an application
	// directory partition.
	TypeIDEnlistDP TypeID = 31
	// DNSSRV_TYPEID_ZONE_CHANGE_DP: A pointer to a structure of type DNS_RPC_ZONE_CHANGE_DP
	// (section 2.2.7.2.6). This structure is used to request that a DNS zone be moved from
	// one application directory partition to another.
	TypeIDZoneChangeDP TypeID = 32
	// DNSSRV_TYPEID_ENUM_ZONES_FILTER: A pointer to a structure of type DNS_RPC_ENUM_ZONES_FILTER
	// (section 2.2.5.2.9). This structure is used to filter DNS zones during enumeration.
	TypeIDEnumZonesFilter TypeID = 33
	// DNSSRV_TYPEID_ADDRARRAY: A pointer to a structure of type DNS_ADDR_ARRAY (section
	// 2.2.3.2.3). This structure is used to specify a list of IPv4 or IPv6 addresses.
	TypeIDAddrArray TypeID = 34
	// DNSSRV_TYPEID_SERVER_INFO: A pointer to a structure of type DNS_RPC_SERVER_INFO (section
	// 2.2.4.2.2). This structure is used to specify general DNS server state and configuration.
	TypeIDServerInfo TypeID = 35
	// DNSSRV_TYPEID_ZONE_INFO: A pointer to a structure of type DNS_RPC_ZONE_INFO (section
	// 2.2.5.2.4). This structure is used to specify detailed information about a DNS zone.
	TypeIDZoneInfo TypeID = 36
	// DNSSRV_TYPEID_FORWARDERS: A pointer to a structure of type DNS_RPC_FORWARDERS (section
	// 2.2.5.2.10). This structure specifies the set of DNS servers this DNS server will
	// forward unresolved queries to.
	TypeIDForwarders TypeID = 37
	// DNSSRV_TYPEID_ZONE_SECONDARIES: A pointer to a structure of type DNS_RPC_ZONE_SECONDARIES
	// (section 2.2.5.2.5). This structure is used to specify information about the secondary
	// servers for a primary DNS zone.
	TypeIDZoneSecondaries TypeID = 38
	// DNSSRV_TYPEID_ZONE_TYPE_RESET: This value is not used.
	TypeIDZoneTypeReset TypeID = 39
	// DNSSRV_TYPEID_ZONE_CREATE: A pointer to a structure of type DNS_RPC_ZONE_CREATE_INFO
	// (section 2.2.5.2.7). This structure is used to specify parameters required when creating
	// a new DNS zone.
	TypeIDZoneCreate TypeID = 40
	// DNSSRV_TYPEID_IP_VALIDATE: A pointer to a structure of type DNS_RPC_IP_VALIDATE (section
	// 2.2.3.2.4). This structure is used to request IP validation and to return the results
	// of IP validation.
	TypeIDIPValidate TypeID = 41
	// DNSSRV_TYPEID_AUTOCONFIGURE: A pointer to a structure of type DNS_RPC_AUTOCONFIGURE
	// (section 2.2.8.2.1). This structure is used to request DNS server autoconfiguration.
	TypeIDAutoConfigure TypeID = 42
	// DNSSRV_TYPEID_UTF8_STRING_LIST: A pointer to a structure of type DNS_RPC_UTF8_STRING_LIST
	// (section 2.2.1.2.3). This structure is used to represent a list of UTF-8 [RFC3629]
	// strings.
	TypeIDUTF8StringList TypeID = 43
	// DNSSRV_TYPEID_UNICODE_STRING_LIST: A pointer to a structure of type DNS_RPC_UNICODE_STRING_LIST
	// (section 2.2.1.2.4). This structure is used to represent a list of Unicode strings.
	TypeIDUnicodeStringList TypeID = 44
	// DNSSRV_TYPEID_SKD: A pointer to a structure of type DNS_RPC_SKD (section 2.2.6.2.1).
	// This structure is used to specify detailed signing key descriptor (SKD) information.
	TypeIDSKD TypeID = 45
	// DNSSRV_TYPEID_SKD_LIST: A pointer to a structure of type DNS_RPC_SKD_LIST (section
	// 2.2.6.2.2). This structure is used to enumerate signing key descriptors.
	TypeIDSKDList TypeID = 46
	// DNSSRV_TYPEID_SKD_STATE: A pointer to a structure of type DNS_RPC_SKD_STATE (section
	// 2.2.6.2.3). This structure is used to specify detailed signing key descriptor state
	// information.
	TypeIDSKDState TypeID = 47
	// DNSSRV_TYPEID_SIGNING_VALIDATION_ERROR: A pointer to a structure of type DNS_RPC_SIGNING_VALIDATION_ERROR
	// (section 2.2.6.2.8). This structure is used to specify signing key descriptor validation
	// error information.
	TypeIDSigningValidationError TypeID = 48
	// DNSSRV_TYPEID_TRUST_POINT_LIST: A pointer to a structure of type DNS_RPC_TRUST_POINT_LIST
	// (section 2.2.6.2.5). This structure is used to enumerate trust points.
	TypeIDTrustPointList TypeID = 49
	// DNSSRV_TYPEID_TRUST_ANCHOR_LIST: A pointer to a structure of type DNS_RPC_TRUST_ANCHOR_LIST
	// (section 2.2.6.2.7). This structure is used to enumerate trust anchors.
	TypeIDTrustAnchorList TypeID = 50
	// DNSSRV_TYPEID_ZONE_SIGNING_SETTINGS: A pointer to a structure of type DNS_RPC_ZONE_DNSSEC_SETTINGS
	// (section 2.2.6.2.9). This structure is used to specify the DNSSEC settings for file-backed
	// zones.
	TypeIDZoneSigningSettings TypeID = 51
	// DNSSRV_TYPEID_ZONE_SCOPE_ENUM: A pointer to a structure of type DNS_RPC_ENUM_ZONE_SCOPE_LIST
	// (section 2.2.13.1.1). This structure is used to enumerate zone scopes or cache scopes.
	TypeIDZoneScopeEnum TypeID = 52
	// DNSSRV_TYPEID_ZONE_STATS: A pointer to a structure of type DNS_RPC_ZONE_STATS_V1
	// (section 2.2.12.2.5). This structure is used to enumerate the zone statistics.
	TypeIDZoneStats TypeID = 53
	// DNSSRV_TYPEID_ZONE_SCOPE_CREATE: A pointer to a structure of type DNS_RPC_ZONE_SCOPE_CREATE_INFO_V1
	// (section 2.2.13.1.2.1). This structure is used to create the zone scope or cache
	// scope.
	TypeIDZoneScopeCreate TypeID = 54
	// DNSSRV_TYPEID_ZONE_SCOPE_INFO: A pointer to a structure of type DNS_RPC_ZONE_SCOPE_INFO_V1
	// (section 2.2.13.1.3.1). This structure is used to get the zone scope or cache scope
	// information.
	TypeIDZoneScopeInfo TypeID = 55
	// DNSSRV_TYPEID_SCOPE_ENUM: A pointer to a structure of type DNS_RPC_ENUM_SCOPE_LIST
	// (section 2.2.14.1.1). The structure is used to enumerate the server scopes configured
	// on the DNS server.
	TypeIDScopeEnum TypeID = 56
	// DNSSRV_TYPEID_CLIENT_SUBNET_RECORD: A pointer to a structure of type DNS_RPC_CLIENT_SUBNET_RECORD
	// (section 2.2.15.2.1). The structure is used to enumerate the client subnet records
	// configured on the DNS server.
	TypeIDClientSubnetRecord TypeID = 57
	// DNSSRV_TYPEID_POLICY: A pointer to a structure of type DNS_RPC_POLICY (section 2.2.15.2.5).
	TypeIDPolicy TypeID = 58
	// DNSSRV_TYPEID_POLICY_NAME: A pointer to a structure of type DNS_RPC_POLICY_NAME (section
	// 2.2.15.2.6).
	TypeIDPolicyName TypeID = 59
	// DNSSRV_TYPEID_POLICY_ENUM: A pointer to a structure of type DNS_RPC_ENUMERATE_POLICY_LIST
	// (section 2.2.15.2.7). The structure is used to enumerate the DNS policy configured
	// on the DNS server.
	TypeIDPolicyEnum TypeID = 60
	// DNSSRV_TYPEID_RRL: A pointer to a structure of type DNS_RPC_RRL_PARAMS (section 2.2.16.2.1).
	// This structure is used to configure parameters for Response Rate Limiting (RRL).
	TypeIDRRL TypeID = 61
	// DNSSRV_TYPEID_VIRTUALIZATION_INSTANCE:A pointer to a structure of type DNS_RPC_VIRTUALIZATION_INSTANCE
	// (section 2.2.17.1.1).
	TypeIDVirtualizationInstance TypeID = 62
	// DNSSRV_TYPEID_VIRTUALIZATION_INSTANCE_ENUM: A pointer to a structure of type DNS_RPC_ENUM_VIRTUALIZATION_INSTANCE_LIST
	// (section 2.2.17.1.3). This structure is used to enumerate the virtualization instances
	// in the DNS Server.
	//
	// Clients and servers of the DNS Server Management Protocol SHOULD<5> support all values
	// above.
	TypeIDVirtualizationInstanceEnum TypeID = 63
)

func (o TypeID) String() string {
	switch o {
	case TypeIDAny:
		return "TypeIDAny"
	case TypeIDNull:
		return "TypeIDNull"
	case TypeIDDword:
		return "TypeIDDword"
	case TypeIDString:
		return "TypeIDString"
	case TypeIDUnicodeString:
		return "TypeIDUnicodeString"
	case TypeIDIPArray:
		return "TypeIDIPArray"
	case TypeIDBuffer:
		return "TypeIDBuffer"
	case TypeIDServerInfoW2K:
		return "TypeIDServerInfoW2K"
	case TypeIDStats:
		return "TypeIDStats"
	case TypeIDForwardersW2K:
		return "TypeIDForwardersW2K"
	case TypeIDZoneW2K:
		return "TypeIDZoneW2K"
	case TypeIDZoneInfoW2K:
		return "TypeIDZoneInfoW2K"
	case TypeIDZoneSecondariesW2K:
		return "TypeIDZoneSecondariesW2K"
	case TypeIDZoneDatabaseW2K:
		return "TypeIDZoneDatabaseW2K"
	case TypeIDZoneTypeResetW2K:
		return "TypeIDZoneTypeResetW2K"
	case TypeIDZoneCreateW2K:
		return "TypeIDZoneCreateW2K"
	case TypeIDNameAndParam:
		return "TypeIDNameAndParam"
	case TypeIDZoneListW2K:
		return "TypeIDZoneListW2K"
	case TypeIDZoneRename:
		return "TypeIDZoneRename"
	case TypeIDZoneExport:
		return "TypeIDZoneExport"
	case TypeIDServerInfoNet:
		return "TypeIDServerInfoNet"
	case TypeIDForwardersNet:
		return "TypeIDForwardersNet"
	case TypeIDZone:
		return "TypeIDZone"
	case TypeIDZoneInfoNet:
		return "TypeIDZoneInfoNet"
	case TypeIDZoneSecondariesNet:
		return "TypeIDZoneSecondariesNet"
	case TypeIDZoneDatabase:
		return "TypeIDZoneDatabase"
	case TypeIDZoneTypeResetNet:
		return "TypeIDZoneTypeResetNet"
	case TypeIDZoneCreateNet:
		return "TypeIDZoneCreateNet"
	case TypeIDZoneList:
		return "TypeIDZoneList"
	case TypeIDIddpEnum:
		return "TypeIDIddpEnum"
	case TypeIDIddpInfo:
		return "TypeIDIddpInfo"
	case TypeIDIddpList:
		return "TypeIDIddpList"
	case TypeIDEnlistDP:
		return "TypeIDEnlistDP"
	case TypeIDZoneChangeDP:
		return "TypeIDZoneChangeDP"
	case TypeIDEnumZonesFilter:
		return "TypeIDEnumZonesFilter"
	case TypeIDAddrArray:
		return "TypeIDAddrArray"
	case TypeIDServerInfo:
		return "TypeIDServerInfo"
	case TypeIDZoneInfo:
		return "TypeIDZoneInfo"
	case TypeIDForwarders:
		return "TypeIDForwarders"
	case TypeIDZoneSecondaries:
		return "TypeIDZoneSecondaries"
	case TypeIDZoneTypeReset:
		return "TypeIDZoneTypeReset"
	case TypeIDZoneCreate:
		return "TypeIDZoneCreate"
	case TypeIDIPValidate:
		return "TypeIDIPValidate"
	case TypeIDAutoConfigure:
		return "TypeIDAutoConfigure"
	case TypeIDUTF8StringList:
		return "TypeIDUTF8StringList"
	case TypeIDUnicodeStringList:
		return "TypeIDUnicodeStringList"
	case TypeIDSKD:
		return "TypeIDSKD"
	case TypeIDSKDList:
		return "TypeIDSKDList"
	case TypeIDSKDState:
		return "TypeIDSKDState"
	case TypeIDSigningValidationError:
		return "TypeIDSigningValidationError"
	case TypeIDTrustPointList:
		return "TypeIDTrustPointList"
	case TypeIDTrustAnchorList:
		return "TypeIDTrustAnchorList"
	case TypeIDZoneSigningSettings:
		return "TypeIDZoneSigningSettings"
	case TypeIDZoneScopeEnum:
		return "TypeIDZoneScopeEnum"
	case TypeIDZoneStats:
		return "TypeIDZoneStats"
	case TypeIDZoneScopeCreate:
		return "TypeIDZoneScopeCreate"
	case TypeIDZoneScopeInfo:
		return "TypeIDZoneScopeInfo"
	case TypeIDScopeEnum:
		return "TypeIDScopeEnum"
	case TypeIDClientSubnetRecord:
		return "TypeIDClientSubnetRecord"
	case TypeIDPolicy:
		return "TypeIDPolicy"
	case TypeIDPolicyName:
		return "TypeIDPolicyName"
	case TypeIDPolicyEnum:
		return "TypeIDPolicyEnum"
	case TypeIDRRL:
		return "TypeIDRRL"
	case TypeIDVirtualizationInstance:
		return "TypeIDVirtualizationInstance"
	case TypeIDVirtualizationInstanceEnum:
		return "TypeIDVirtualizationInstanceEnum"
	}
	return "Invalid"
}

// ZoneStats structure represents DNS_RPC_ZONE_STATS RPC structure.
type ZoneStats ZoneStatsV1

func (o *ZoneStats) ZoneStatsV1() *ZoneStatsV1 { return (*ZoneStatsV1)(o) }

func (o *ZoneStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ZoneStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	if o.ZoneTimeStats != nil {
		if err := o.ZoneTimeStats.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ZoneTimeStats{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	for i1 := range o.ZoneQueryStats {
		i1 := i1
		if uint64(i1) >= 32 {
			break
		}
		if o.ZoneQueryStats[i1] != nil {
			if err := o.ZoneQueryStats[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ZoneQueryStats{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.ZoneQueryStats); uint64(i1) < 32; i1++ {
		if err := (&ZoneQueryStats{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	for i1 := range o.ZoneTransferStats {
		i1 := i1
		if uint64(i1) >= 2 {
			break
		}
		if o.ZoneTransferStats[i1] != nil {
			if err := o.ZoneTransferStats[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ZoneTransferStats{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.ZoneTransferStats); uint64(i1) < 2; i1++ {
		if err := (&ZoneTransferStats{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ZoneUpdateStats != nil {
		if err := o.ZoneUpdateStats.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ZoneUpdateStats{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ZoneRRLStats != nil {
		if err := o.ZoneRRLStats.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ZoneRRLStats{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *ZoneStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	if o.ZoneTimeStats == nil {
		o.ZoneTimeStats = &ZoneTimeStats{}
	}
	if err := o.ZoneTimeStats.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	o.ZoneQueryStats = make([]*ZoneQueryStats, 32)
	for i1 := range o.ZoneQueryStats {
		i1 := i1
		if o.ZoneQueryStats[i1] == nil {
			o.ZoneQueryStats[i1] = &ZoneQueryStats{}
		}
		if err := o.ZoneQueryStats[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	o.ZoneTransferStats = make([]*ZoneTransferStats, 2)
	for i1 := range o.ZoneTransferStats {
		i1 := i1
		if o.ZoneTransferStats[i1] == nil {
			o.ZoneTransferStats[i1] = &ZoneTransferStats{}
		}
		if err := o.ZoneTransferStats[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ZoneUpdateStats == nil {
		o.ZoneUpdateStats = &ZoneUpdateStats{}
	}
	if err := o.ZoneUpdateStats.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ZoneRRLStats == nil {
		o.ZoneRRLStats = &ZoneRRLStats{}
	}
	if err := o.ZoneRRLStats.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// ZoneScopeCreateInfo structure represents DNS_RPC_ZONE_SCOPE_CREATE_INFO RPC structure.
type ZoneScopeCreateInfo ZoneScopeCreateInfoV1

func (o *ZoneScopeCreateInfo) ZoneScopeCreateInfoV1() *ZoneScopeCreateInfoV1 {
	return (*ZoneScopeCreateInfoV1)(o)
}

func (o *ZoneScopeCreateInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ZoneScopeCreateInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.ScopeName != "" {
		_ptr_pwszScopeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.ScopeName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ScopeName, _ptr_pwszScopeName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ZoneScopeCreateInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_pwszScopeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.ScopeName); err != nil {
			return err
		}
		return nil
	})
	_s_pwszScopeName := func(ptr interface{}) { o.ScopeName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ScopeName, _s_pwszScopeName, _ptr_pwszScopeName); err != nil {
		return err
	}
	return nil
}

// ZoneScopeInfo structure represents DNS_RPC_ZONE_SCOPE_INFO RPC structure.
type ZoneScopeInfo ZoneScopeInfoV1

func (o *ZoneScopeInfo) ZoneScopeInfoV1() *ZoneScopeInfoV1 { return (*ZoneScopeInfoV1)(o) }

func (o *ZoneScopeInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ZoneScopeInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	if o.ScopeName != "" {
		_ptr_pwszScopeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.ScopeName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ScopeName, _ptr_pwszScopeName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DataFile != "" {
		_ptr_pwszDataFile := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.DataFile); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DataFile, _ptr_pwszDataFile); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ZoneScopeInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	_ptr_pwszScopeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.ScopeName); err != nil {
			return err
		}
		return nil
	})
	_s_pwszScopeName := func(ptr interface{}) { o.ScopeName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ScopeName, _s_pwszScopeName, _ptr_pwszScopeName); err != nil {
		return err
	}
	_ptr_pwszDataFile := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.DataFile); err != nil {
			return err
		}
		return nil
	})
	_s_pwszDataFile := func(ptr interface{}) { o.DataFile = *ptr.(*string) }
	if err := w.ReadPointer(&o.DataFile, _s_pwszDataFile, _ptr_pwszDataFile); err != nil {
		return err
	}
	return nil
}

// EnumZoneScopeList structure represents DNS_RPC_ENUM_ZONE_SCOPE_LIST RPC structure.
//
// The DNS_RPC_ENUM_ZONE_SCOPE_LIST structure<86> contains a list of zone scopes or
// cache scopes to be enumerated.
type EnumZoneScopeList struct {
	// dwRpcStructureVersion: The DNS management structure version number. This value MUST
	// be set to 0x00000001.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// dwZoneScopeCount: The number of zone scopes or cache scopes.
	ZoneScopeCount uint32 `idl:"name:dwZoneScopeCount" json:"zone_scope_count"`
	// ZoneScopeArray: An array containing the names of zone scopes or cache scopes.
	ZoneScopeArray []string `idl:"name:ZoneScopeArray;size_is:(dwZoneScopeCount)" json:"zone_scope_array"`
}

func (o *EnumZoneScopeList) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.ZoneScopeArray != nil && o.ZoneScopeCount == 0 {
		o.ZoneScopeCount = uint32(len(o.ZoneScopeArray))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *EnumZoneScopeList) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.ZoneScopeCount)
	return []uint64{
		dimSize1,
	}
}
func (o *EnumZoneScopeList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.ZoneScopeCount); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	for i1 := range o.ZoneScopeArray {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.ZoneScopeArray[i1] != "" {
			_ptr_ZoneScopeArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16String(ctx, w, o.ZoneScopeArray[i1]); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ZoneScopeArray[i1], _ptr_ZoneScopeArray); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.ZoneScopeArray); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *EnumZoneScopeList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.ZoneScopeCount); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.ZoneScopeCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.ZoneScopeCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.ZoneScopeArray", sizeInfo[0])
	}
	o.ZoneScopeArray = make([]string, sizeInfo[0])
	for i1 := range o.ZoneScopeArray {
		i1 := i1
		_ptr_ZoneScopeArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16String(ctx, w, &o.ZoneScopeArray[i1]); err != nil {
				return err
			}
			return nil
		})
		_s_ZoneScopeArray := func(ptr interface{}) { o.ZoneScopeArray[i1] = *ptr.(*string) }
		if err := w.ReadPointer(&o.ZoneScopeArray[i1], _s_ZoneScopeArray, _ptr_ZoneScopeArray); err != nil {
			return err
		}
	}
	return nil
}

// SystemTime structure represents DNS_SYSTEMTIME RPC structure.
//
// The DNS_SYSTEMTIME structure stores time values for DNS statistics. It is always
// populated by the server, which MUST supply valid values. This structure MUST be formatted
// as follows:
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| wYear                                                         | wMonth                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| wDayOfWeek                                                    | wDay                                                          |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| wHour                                                         | wMinute                                                       |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| wSecond                                                       | wMillisecond                                                  |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type SystemTime struct {
	// wYear (2 bytes): The year, as a 16-bit, unsigned integer. Valid values are from 1601
	// to 30827.
	Year uint16 `idl:"name:wYear" json:"year"`
	// wMonth (2 bytes): The month from 1 (January) to 12 (December), as a 16-bit, unsigned
	// integer.
	Month uint16 `idl:"name:wMonth" json:"month"`
	// wDayOfWeek (2 bytes): The day of the week from 0 (Sunday) to 6 (Saturday), as a 16-bit,
	// unsigned integer.
	DayOfWeek uint16 `idl:"name:wDayOfWeek" json:"day_of_week"`
	// wDay (2 bytes): The day of the month from 1 to 31, as a 16-bit, unsigned integer.
	Day uint16 `idl:"name:wDay" json:"day"`
	// wHour (2 bytes): The hour from 0 to 23, as a 16-bit, unsigned integer.
	Hour uint16 `idl:"name:wHour" json:"hour"`
	// wMinute (2 bytes): The minute from 0 to 59, as a 16-bit, unsigned integer.
	Minute uint16 `idl:"name:wMinute" json:"minute"`
	// wSecond (2 bytes): The second from 0 to 59, as a 16-bit, unsigned integer.
	Second       uint16 `idl:"name:wSecond" json:"second"`
	Milliseconds uint16 `idl:"name:wMilliseconds" json:"milliseconds"`
}

func (o *SystemTime) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SystemTime) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(2); err != nil {
		return err
	}
	if err := w.WriteData(o.Year); err != nil {
		return err
	}
	if err := w.WriteData(o.Month); err != nil {
		return err
	}
	if err := w.WriteData(o.DayOfWeek); err != nil {
		return err
	}
	if err := w.WriteData(o.Day); err != nil {
		return err
	}
	if err := w.WriteData(o.Hour); err != nil {
		return err
	}
	if err := w.WriteData(o.Minute); err != nil {
		return err
	}
	if err := w.WriteData(o.Second); err != nil {
		return err
	}
	if err := w.WriteData(o.Milliseconds); err != nil {
		return err
	}
	return nil
}
func (o *SystemTime) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(2); err != nil {
		return err
	}
	if err := w.ReadData(&o.Year); err != nil {
		return err
	}
	if err := w.ReadData(&o.Month); err != nil {
		return err
	}
	if err := w.ReadData(&o.DayOfWeek); err != nil {
		return err
	}
	if err := w.ReadData(&o.Day); err != nil {
		return err
	}
	if err := w.ReadData(&o.Hour); err != nil {
		return err
	}
	if err := w.ReadData(&o.Minute); err != nil {
		return err
	}
	if err := w.ReadData(&o.Second); err != nil {
		return err
	}
	if err := w.ReadData(&o.Milliseconds); err != nil {
		return err
	}
	return nil
}

// ZoneTimeStats structure represents DNSSRV_ZONE_TIME_STATS RPC structure.
//
// The DNSSRV_ZONE_TIME_STATS structure<80> contains zone time statistics information.
type ZoneTimeStats struct {
	// StatsCollectionStartTime: The time when zone statistics collection started.
	StatsCollectionStartTime *SystemTime `idl:"name:StatsCollectionStartTime" json:"stats_collection_start_time"`
}

func (o *ZoneTimeStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ZoneTimeStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(2); err != nil {
		return err
	}
	if o.StatsCollectionStartTime != nil {
		if err := o.StatsCollectionStartTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SystemTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ZoneTimeStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(2); err != nil {
		return err
	}
	if o.StatsCollectionStartTime == nil {
		o.StatsCollectionStartTime = &SystemTime{}
	}
	if err := o.StatsCollectionStartTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ZoneStatsType type represents DNS_ZONE_STATS_TYPE RPC enumeration.
//
// The DNS_ZONE_STATS_TYPE enumeration is the enumerator for all zone statistics types.<79>
type ZoneStatsType uint16

var (
	// ZONE_STATS_TYPE_RECORD_A: A record type.
	ZoneStatsTypeRecordA ZoneStatsType = 0
	// ZONE_STATS_TYPE_RECORD_AAAA: AAAA record type.
	ZoneStatsTypeRecordAAAA ZoneStatsType = 1
	// ZONE_STATS_TYPE_RECORD_PTR: PTR record type.
	ZoneStatsTypeRecordPTR ZoneStatsType = 2
	// ZONE_STATS_TYPE_RECORD_CNAME: CNAME record type.
	ZoneStatsTypeRecordCNAME ZoneStatsType = 3
	// ZONE_STATS_TYPE_RECORD_MX: MX record type.
	ZoneStatsTypeRecordMX ZoneStatsType = 4
	// ZONE_STATS_TYPE_RECORD_AFSDB: AFDSB record type.
	ZoneStatsTypeRecordAFSDB ZoneStatsType = 5
	// ZONE_STATS_TYPE_RECORD_ATMA: ATMA record type.
	ZoneStatsTypeRecordATMA ZoneStatsType = 6
	// ZONE_STATS_TYPE_RECORD_DHCID: DHCID record type.
	ZoneStatsTypeRecordDHCID ZoneStatsType = 7
	// ZONE_STATS_TYPE_RECORD_DNAME: DNAME record type.
	ZoneStatsTypeRecordDNAME ZoneStatsType = 8
	// ZONE_STATS_TYPE_RECORD_HINFO: HINFO record type.
	ZoneStatsTypeRecordHINFO ZoneStatsType = 9
	// ZONE_STATS_TYPE_RECORD_ISDN: ISDN record type.
	ZoneStatsTypeRecordISDN ZoneStatsType = 10
	// ZONE_STATS_TYPE_RECORD_MG: MG record type.
	ZoneStatsTypeRecordMG ZoneStatsType = 11
	// ZONE_STATS_TYPE_RECORD_MB: MB record type.
	ZoneStatsTypeRecordMB ZoneStatsType = 12
	// ZONE_STATS_TYPE_RECORD_MINFO: MINFO record type.
	ZoneStatsTypeRecordMINFO ZoneStatsType = 13
	// ZONE_STATS_TYPE_RECORD_NAPTR: NAPTR record type.
	ZoneStatsTypeRecordNAPTR ZoneStatsType = 14
	// ZONE_STATS_TYPE_RECORD_NXT: NXT record type.
	ZoneStatsTypeRecordNXT ZoneStatsType = 15
	// ZONE_STATS_TYPE_RECORD_KEY: KEY record type.
	ZoneStatsTypeRecordKEY ZoneStatsType = 16
	// ZONE_STATS_TYPE_RECORD_MR: MR record type.
	ZoneStatsTypeRecordMR ZoneStatsType = 17
	// ZONE_STATS_TYPE_RECORD_RP: RP record type.
	ZoneStatsTypeRecordRP ZoneStatsType = 18
	// ZONE_STATS_TYPE_RECORD_RT: RT record type.
	ZoneStatsTypeRecordRT ZoneStatsType = 19
	// ZONE_STATS_TYPE_RECORD_SRV: SRV record type.
	ZoneStatsTypeRecordSRV ZoneStatsType = 20
	// ZONE_STATS_TYPE_RECORD_SIG: SIG record type.
	ZoneStatsTypeRecordSIG ZoneStatsType = 21
	// ZONE_STATS_TYPE_RECORD_TEXT: TXT record type.
	ZoneStatsTypeRecordTXT ZoneStatsType = 22
	// ZONE_STATS_TYPE_RECORD_WKS: WKS record type.
	ZoneStatsTypeRecordWKS ZoneStatsType = 23
	// ZONE_STATS_TYPE_RECORD_X25: X25 record type.
	ZoneStatsTypeRecordX25 ZoneStatsType = 24
	// ZONE_STATS_TYPE_RECORD_DNSKEY: DNSKEY record type.
	ZoneStatsTypeRecordDNSKEY ZoneStatsType = 25
	// ZONE_STATS_TYPE_RECORD_DS: DS record type.
	ZoneStatsTypeRecordDS ZoneStatsType = 26
	// ZONE_STATS_TYPE_RECORD_NS: NS record type.
	ZoneStatsTypeRecordNS ZoneStatsType = 27
	// ZONE_STATS_TYPE_RECORD_SOA: SOA record type.
	ZoneStatsTypeRecordSOA ZoneStatsType = 28
	// ZONE_STATS_TYPE_RECORD_TLSA: TLSA record type.
	ZoneStatsTypeRecordTLSA ZoneStatsType = 29
	// ZONE_STATS_TYPE_RECORD_ALL: All record types.
	ZoneStatsTypeRecordAll ZoneStatsType = 30
	// ZONE_STATS_TYPE_RECORD_OTHERS: Resource records not specified in this list.
	ZoneStatsTypeRecordOthers ZoneStatsType = 31
	// ZONE_STATS_TYPE_TRANSFER_AXFR: Axft transfer type.
	ZoneStatsTypeTransferAXFR ZoneStatsType = 32
	// ZONE_STATS_TYPE_TRANSFER_IXFR: Ixfr transfer type.
	ZoneStatsTypeTransferIXFR ZoneStatsType = 33
	// ZONE_STATS_TYPE_UPDATE: Zone Update.
	ZoneStatsTypeUpdate ZoneStatsType = 34
	// ZONE_STATS_TYPE_RRL: Response Rate Limiting.
	ZoneStatsTypeRRL ZoneStatsType = 35
	// MAX_ZONE_STATS_TYPES: Specifies the maximum zone statistics types supported.
	ZoneStatsTypeMaxZoneStatsTypes ZoneStatsType = 36
)

func (o ZoneStatsType) String() string {
	switch o {
	case ZoneStatsTypeRecordA:
		return "ZoneStatsTypeRecordA"
	case ZoneStatsTypeRecordAAAA:
		return "ZoneStatsTypeRecordAAAA"
	case ZoneStatsTypeRecordPTR:
		return "ZoneStatsTypeRecordPTR"
	case ZoneStatsTypeRecordCNAME:
		return "ZoneStatsTypeRecordCNAME"
	case ZoneStatsTypeRecordMX:
		return "ZoneStatsTypeRecordMX"
	case ZoneStatsTypeRecordAFSDB:
		return "ZoneStatsTypeRecordAFSDB"
	case ZoneStatsTypeRecordATMA:
		return "ZoneStatsTypeRecordATMA"
	case ZoneStatsTypeRecordDHCID:
		return "ZoneStatsTypeRecordDHCID"
	case ZoneStatsTypeRecordDNAME:
		return "ZoneStatsTypeRecordDNAME"
	case ZoneStatsTypeRecordHINFO:
		return "ZoneStatsTypeRecordHINFO"
	case ZoneStatsTypeRecordISDN:
		return "ZoneStatsTypeRecordISDN"
	case ZoneStatsTypeRecordMG:
		return "ZoneStatsTypeRecordMG"
	case ZoneStatsTypeRecordMB:
		return "ZoneStatsTypeRecordMB"
	case ZoneStatsTypeRecordMINFO:
		return "ZoneStatsTypeRecordMINFO"
	case ZoneStatsTypeRecordNAPTR:
		return "ZoneStatsTypeRecordNAPTR"
	case ZoneStatsTypeRecordNXT:
		return "ZoneStatsTypeRecordNXT"
	case ZoneStatsTypeRecordKEY:
		return "ZoneStatsTypeRecordKEY"
	case ZoneStatsTypeRecordMR:
		return "ZoneStatsTypeRecordMR"
	case ZoneStatsTypeRecordRP:
		return "ZoneStatsTypeRecordRP"
	case ZoneStatsTypeRecordRT:
		return "ZoneStatsTypeRecordRT"
	case ZoneStatsTypeRecordSRV:
		return "ZoneStatsTypeRecordSRV"
	case ZoneStatsTypeRecordSIG:
		return "ZoneStatsTypeRecordSIG"
	case ZoneStatsTypeRecordTXT:
		return "ZoneStatsTypeRecordTXT"
	case ZoneStatsTypeRecordWKS:
		return "ZoneStatsTypeRecordWKS"
	case ZoneStatsTypeRecordX25:
		return "ZoneStatsTypeRecordX25"
	case ZoneStatsTypeRecordDNSKEY:
		return "ZoneStatsTypeRecordDNSKEY"
	case ZoneStatsTypeRecordDS:
		return "ZoneStatsTypeRecordDS"
	case ZoneStatsTypeRecordNS:
		return "ZoneStatsTypeRecordNS"
	case ZoneStatsTypeRecordSOA:
		return "ZoneStatsTypeRecordSOA"
	case ZoneStatsTypeRecordTLSA:
		return "ZoneStatsTypeRecordTLSA"
	case ZoneStatsTypeRecordAll:
		return "ZoneStatsTypeRecordAll"
	case ZoneStatsTypeRecordOthers:
		return "ZoneStatsTypeRecordOthers"
	case ZoneStatsTypeTransferAXFR:
		return "ZoneStatsTypeTransferAXFR"
	case ZoneStatsTypeTransferIXFR:
		return "ZoneStatsTypeTransferIXFR"
	case ZoneStatsTypeUpdate:
		return "ZoneStatsTypeUpdate"
	case ZoneStatsTypeRRL:
		return "ZoneStatsTypeRRL"
	case ZoneStatsTypeMaxZoneStatsTypes:
		return "ZoneStatsTypeMaxZoneStatsTypes"
	}
	return "Invalid"
}

// ZoneQueryStats structure represents DNSSRV_ZONE_QUERY_STATS RPC structure.
//
// The DNSSRV_ZONE_QUERY_STATS structure<81> contains per-zone per-record-type statistics.
type ZoneQueryStats struct {
	// RecordType: The type of record for which the query was received. The value SHOULD
	// be of type DNS_ZONE_STATS_TYPE (section 2.2.12.1.1).
	RecordType ZoneStatsType `idl:"name:RecordType" json:"record_type"`
	// QueriesResponded: The total number of queries to which the server responded for a
	// specific zone.
	QueriesResponded uint64 `idl:"name:QueriesResponded" json:"queries_responded"`
	// QueriesReceived: The total number of queries received by the server for a specific
	// zone.
	QueriesReceived uint64 `idl:"name:QueriesReceived" json:"queries_received"`
	// QueriesFailure: The total number of queries for which the server responded with server
	// failure for a specific zone.
	QueriesFailure uint64 `idl:"name:QueriesFailure" json:"queries_failure"`
	// QueriesNameError: The total number of queries for which the server responded with
	// a name error for a specific zone.
	QueriesNameError uint64 `idl:"name:QueriesNameError" json:"queries_name_error"`
}

func (o *ZoneQueryStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ZoneQueryStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.RecordType)); err != nil {
		return err
	}
	if err := w.WriteData(o.QueriesResponded); err != nil {
		return err
	}
	if err := w.WriteData(o.QueriesReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.QueriesFailure); err != nil {
		return err
	}
	if err := w.WriteData(o.QueriesNameError); err != nil {
		return err
	}
	return nil
}
func (o *ZoneQueryStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.RecordType)); err != nil {
		return err
	}
	if err := w.ReadData(&o.QueriesResponded); err != nil {
		return err
	}
	if err := w.ReadData(&o.QueriesReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.QueriesFailure); err != nil {
		return err
	}
	if err := w.ReadData(&o.QueriesNameError); err != nil {
		return err
	}
	return nil
}

// RRLStats structure represents DNSSRV_RRL_STATS RPC structure.
//
// The DNSSRV_RRL_STATS structure SHOULD<78> contain DNS server statistics related to
// the Response Rate Limiting (RRL). This structure is formatted as follows:
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Header                                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| TotalResponsesSent                                                                                                            |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| TotalResponsesDropped                                                                                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| TotalResponsesTruncated                                                                                                       |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| TotalResponsesLeaked                                                                                                          |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type RRLStats struct {
	// Header (8 bytes): A structure of type DNSSRV_STAT_HEADER (section 2.2.10.2.1).
	Header *StatHeader `idl:"name:Header" json:"header"`
	// TotalResponsesSent (4 bytes): The number of times the server responded to a valid
	// query since it last started. This counter is maintained and updated only when RRL
	// is enabled. After reaching 0xFFFFFFFF, the value increments to 0x00000000.
	TotalResponsesSent uint32 `idl:"name:TotalResponsesSent" json:"total_responses_sent"`
	// TotalResponsesDropped (4 bytes): The number of times the server dropped a valid query
	// due to Response Rate Limiting, since it last started. This counter is maintained
	// and updated only when RRL is enabled. After reaching 0xFFFFFFFF, the value increments
	// to 0x00000000.
	TotalResponsesDropped uint32 `idl:"name:TotalResponsesDropped" json:"total_responses_dropped"`
	// TotalResponsesTruncated (4 bytes): The number of times the server has responded to
	// a valid query with a truncation bit set, since it last started. This counter is maintained
	// and updated only when RRL is enabled. After reaching 0xFFFFFFFF, the value increments
	// to 0x00000000.
	TotalResponsesTruncated uint32 `idl:"name:TotalResponsesTruncated" json:"total_responses_truncated"`
	// TotalResponsesLeaked (4 bytes): The number of times the server has responded to a
	// valid query after using the total responses available within a given timeframe since
	// it last started. For details see DNS_RPC_RRL_PARAMS (section 2.2.16.2.1). This counter
	// is maintained and updated only when RRL is enabled. After reaching 0xFFFFFFFF, the
	// value increments to 0x00000000.
	TotalResponsesLeaked uint32 `idl:"name:TotalResponsesLeaked" json:"total_responses_leaked"`
}

func (o *RRLStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RRLStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.Header != nil {
		if err := o.Header.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&StatHeader{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.TotalResponsesSent); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalResponsesDropped); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalResponsesTruncated); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalResponsesLeaked); err != nil {
		return err
	}
	return nil
}
func (o *RRLStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.Header == nil {
		o.Header = &StatHeader{}
	}
	if err := o.Header.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalResponsesSent); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalResponsesDropped); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalResponsesTruncated); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalResponsesLeaked); err != nil {
		return err
	}
	return nil
}

// ZoneTransferStats structure represents DNSSRV_ZONE_TRANSFER_STATS RPC structure.
//
// The DNSSRV_ZONE_TRANSFER_STATS structure<82> contains zone transfer statistics per
// zone.
type ZoneTransferStats struct {
	// TransferType: The type of zone transfer request for which the query was received.
	// The value is of type DNS_ZONE_STATS_TYPE (section 2.2.12.1.1).
	TransferType ZoneStatsType `idl:"name:TransferType" json:"transfer_type"`
	// RequestReceived: The total number of zone transfer requests received by the server
	// for a specific zone.
	RequestReceived uint64 `idl:"name:RequestReceived" json:"request_received"`
	// RequestSent: The total number of zone transfer requests sent by the server for a
	// specific zone.
	RequestSent uint64 `idl:"name:RequestSent" json:"request_sent"`
	// ResponseReceived: The total number of zone transfer responses received by the server
	// for a specific zone.
	ResponseReceived uint64 `idl:"name:ResponseReceived" json:"response_received"`
	// SuccessReceived: The total number of zone transfer requests received successfully
	// by the receiver for a specific zone.
	SuccessReceived uint64 `idl:"name:SuccessReceived" json:"success_received"`
	// SuccessSent: The total number of zone transfer responses sent by the server for a
	// specific zone.
	SuccessSent uint64 `idl:"name:SuccessSent" json:"success_sent"`
}

func (o *ZoneTransferStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ZoneTransferStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.TransferType)); err != nil {
		return err
	}
	if err := w.WriteData(o.RequestReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.RequestSent); err != nil {
		return err
	}
	if err := w.WriteData(o.ResponseReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.SuccessReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.SuccessSent); err != nil {
		return err
	}
	return nil
}
func (o *ZoneTransferStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.TransferType)); err != nil {
		return err
	}
	if err := w.ReadData(&o.RequestReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.RequestSent); err != nil {
		return err
	}
	if err := w.ReadData(&o.ResponseReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.SuccessReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.SuccessSent); err != nil {
		return err
	}
	return nil
}

// ZoneUpdateStats structure represents DNSSRV_ZONE_UPDATE_STATS RPC structure.
//
// The DNSSRV_ZONE_UPDATE_STATS structure<83> contains statistics about the dynamic
// updates per zone.
type ZoneUpdateStats struct {
	// Type: The type of zone update request for which statistics are required. The value
	// is of type DNS_ZONE_STATS_TYPE (section 2.2.12.1.1).
	Type ZoneStatsType `idl:"name:Type" json:"type"`
	// DynamicUpdateReceived: The total number of dynamic updates accepted on a specified
	// zone.
	DynamicUpdateReceived uint64 `idl:"name:DynamicUpdateReceived" json:"dynamic_update_received"`
	// DynamicUpdateRejected: The total number of dynamic updates rejected on a specified
	// zone.
	DynamicUpdateRejected uint64 `idl:"name:DynamicUpdateRejected" json:"dynamic_update_rejected"`
}

func (o *ZoneUpdateStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ZoneUpdateStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Type)); err != nil {
		return err
	}
	if err := w.WriteData(o.DynamicUpdateReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.DynamicUpdateRejected); err != nil {
		return err
	}
	return nil
}
func (o *ZoneUpdateStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Type)); err != nil {
		return err
	}
	if err := w.ReadData(&o.DynamicUpdateReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.DynamicUpdateRejected); err != nil {
		return err
	}
	return nil
}

// ZoneRRLStats structure represents DNSSRV_ZONE_RRL_STATS RPC structure.
//
// The DNSSRV_ZONE_RRL_STATS structure SHOULD<85> contain zone statistics about Response
// Rate Limiting.
type ZoneRRLStats struct {
	// Type: The DNS query type for which statistics are required. The value is of type
	// DNS_ZONE_STATS_TYPE (section 2.2.12.1.1).
	Type ZoneStatsType `idl:"name:Type" json:"type"`
	// TotalResponsesSent: The total number of responses sent on a specified zone when RRL
	// is enabled.
	TotalResponsesSent uint32 `idl:"name:TotalResponsesSent" json:"total_responses_sent"`
	// TotalResponsesDropped: The total number responses dropped on a specified zone when
	// RRL is enabled.
	TotalResponsesDropped uint32 `idl:"name:TotalResponsesDropped" json:"total_responses_dropped"`
	// TotalResponsesTruncated: The total number of responses sent with a truncated bit
	// set on a specified zone when RRL is enabled.
	TotalResponsesTruncated uint32 `idl:"name:TotalResponsesTruncated" json:"total_responses_truncated"`
	// TotalResponsesLeaked: The total number of responses sent after exhaustion of available
	// responses on a specified zone when RRL is enabled.
	TotalResponsesLeaked uint32 `idl:"name:TotalResponsesLeaked" json:"total_responses_leaked"`
}

func (o *ZoneRRLStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ZoneRRLStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Type)); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalResponsesSent); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalResponsesDropped); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalResponsesTruncated); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalResponsesLeaked); err != nil {
		return err
	}
	return nil
}
func (o *ZoneRRLStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Type)); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalResponsesSent); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalResponsesDropped); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalResponsesTruncated); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalResponsesLeaked); err != nil {
		return err
	}
	return nil
}

// ZoneStatsV1 structure represents DNS_RPC_ZONE_STATS_V1 RPC structure.
//
// The DNS_RPC_ZONE_STATS_V1 structure SHOULD<84> contain all the statistics about a
// zone.
type ZoneStatsV1 struct {
	// dwRpcStructureVersion: The DNS management structure version number. This value MUST
	// be set to 0x00000001.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// ZoneTimeStats: Information about the zone time statistics in the DNSSRV_ZONE_TIME_STATS
	// structure (section 2.2.12.2.1).
	ZoneTimeStats *ZoneTimeStats `idl:"name:ZoneTimeStats" json:"zone_time_stats"`
	// ZoneQueryStats: An array of DNSSRV_ZONE_QUERY_STATS structures (section 2.2.12.2.2).
	ZoneQueryStats []*ZoneQueryStats `idl:"name:ZoneQueryStats" json:"zone_query_stats"`
	// ZoneTransferStats: An array of DNSSRV_ZONE_TRANSFER_STATS structures (section 2.2.12.2.3).
	ZoneTransferStats []*ZoneTransferStats `idl:"name:ZoneTransferStats" json:"zone_transfer_stats"`
	// ZoneUpdateStats: Information about the zone dynamic update statistics in the DNSSRV_ZONE_UPDATE_STATS
	// structure (section 2.2.12.2.4).
	ZoneUpdateStats *ZoneUpdateStats `idl:"name:ZoneUpdateStats" json:"zone_update_stats"`
	// ZoneRRLStats: Information about the zone RRL statistics in the DNSSRV_ZONE_RRL_STATS
	// structure (section 2.2.12.2.6).
	ZoneRRLStats *ZoneRRLStats `idl:"name:ZoneRRLStats" json:"zone_rrl_stats"`
}

func (o *ZoneStatsV1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ZoneStatsV1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	if o.ZoneTimeStats != nil {
		if err := o.ZoneTimeStats.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ZoneTimeStats{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	for i1 := range o.ZoneQueryStats {
		i1 := i1
		if uint64(i1) >= 32 {
			break
		}
		if o.ZoneQueryStats[i1] != nil {
			if err := o.ZoneQueryStats[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ZoneQueryStats{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.ZoneQueryStats); uint64(i1) < 32; i1++ {
		if err := (&ZoneQueryStats{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	for i1 := range o.ZoneTransferStats {
		i1 := i1
		if uint64(i1) >= 2 {
			break
		}
		if o.ZoneTransferStats[i1] != nil {
			if err := o.ZoneTransferStats[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ZoneTransferStats{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.ZoneTransferStats); uint64(i1) < 2; i1++ {
		if err := (&ZoneTransferStats{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ZoneUpdateStats != nil {
		if err := o.ZoneUpdateStats.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ZoneUpdateStats{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ZoneRRLStats != nil {
		if err := o.ZoneRRLStats.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ZoneRRLStats{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *ZoneStatsV1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	if o.ZoneTimeStats == nil {
		o.ZoneTimeStats = &ZoneTimeStats{}
	}
	if err := o.ZoneTimeStats.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	o.ZoneQueryStats = make([]*ZoneQueryStats, 32)
	for i1 := range o.ZoneQueryStats {
		i1 := i1
		if o.ZoneQueryStats[i1] == nil {
			o.ZoneQueryStats[i1] = &ZoneQueryStats{}
		}
		if err := o.ZoneQueryStats[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	o.ZoneTransferStats = make([]*ZoneTransferStats, 2)
	for i1 := range o.ZoneTransferStats {
		i1 := i1
		if o.ZoneTransferStats[i1] == nil {
			o.ZoneTransferStats[i1] = &ZoneTransferStats{}
		}
		if err := o.ZoneTransferStats[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ZoneUpdateStats == nil {
		o.ZoneUpdateStats = &ZoneUpdateStats{}
	}
	if err := o.ZoneUpdateStats.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ZoneRRLStats == nil {
		o.ZoneRRLStats = &ZoneRRLStats{}
	}
	if err := o.ZoneRRLStats.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// ZoneScopeCreateInfoV1 structure represents DNS_RPC_ZONE_SCOPE_CREATE_INFO_V1 RPC structure.
//
// The DNS_RPC_ZONE_SCOPE_CREATE_INFO_V1 structure<87> contains the name of the zone
// scope or cache scope.
type ZoneScopeCreateInfoV1 struct {
	// dwFlags: Specifies the behavior that the DNS server SHOULD follow when creating the
	// scope. This field is used only when the operation is CreateZoneScope. For any other
	// value than the following, a new empty zone scope is created:
	//
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                           |                                                                                  |
	//	|                   VALUE                   |                                     MEANING                                      |
	//	|                                           |                                                                                  |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_ZONE_LOAD_OVERWRITE_MEMORY 0x00000010 | If this value is specified, the DNS server MUST attempt to find and load the     |
	//	|                                           | zone scope database from persistent storage instead of creating a new empty      |
	//	|                                           | zone scope. If a pre-existing zone scope database is not found, or if there is a |
	//	|                                           | failure in loading the existing database, the DNS server MUST fail the operation |
	//	|                                           | and return an error. If a cache scope is being created, this operation fails on  |
	//	|                                           | the DNS server and an ERROR_INVALID_PARAMETER error is returned.                 |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
	// pwszScopeName: The name of the zone scope or cache scope.
	ScopeName string `idl:"name:pwszScopeName" json:"scope_name"`
}

func (o *ZoneScopeCreateInfoV1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ZoneScopeCreateInfoV1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.ScopeName != "" {
		_ptr_pwszScopeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.ScopeName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ScopeName, _ptr_pwszScopeName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ZoneScopeCreateInfoV1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_pwszScopeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.ScopeName); err != nil {
			return err
		}
		return nil
	})
	_s_pwszScopeName := func(ptr interface{}) { o.ScopeName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ScopeName, _s_pwszScopeName, _ptr_pwszScopeName); err != nil {
		return err
	}
	return nil
}

// ZoneScopeInfoV1 structure represents DNS_RPC_ZONE_SCOPE_INFO_V1 RPC structure.
//
// The DNS_RPC_ZONE_SCOPE_INFO_V1 structure<88> contains the details of the zone scope
// or cache scope.
type ZoneScopeInfoV1 struct {
	// dwRpcStructureVersion: The DNS management structure version number. This value MUST
	// be set to 0x00000001.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// pwszScopeName: The name of the zone scope or cache scope.
	ScopeName string `idl:"name:pwszScopeName" json:"scope_name"`
	// pwszDataFile: The name of the zone scope or cache scope data file.
	DataFile string `idl:"name:pwszDataFile" json:"data_file"`
}

func (o *ZoneScopeInfoV1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ZoneScopeInfoV1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	if o.ScopeName != "" {
		_ptr_pwszScopeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.ScopeName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ScopeName, _ptr_pwszScopeName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DataFile != "" {
		_ptr_pwszDataFile := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.DataFile); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DataFile, _ptr_pwszDataFile); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ZoneScopeInfoV1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	_ptr_pwszScopeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.ScopeName); err != nil {
			return err
		}
		return nil
	})
	_s_pwszScopeName := func(ptr interface{}) { o.ScopeName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ScopeName, _s_pwszScopeName, _ptr_pwszScopeName); err != nil {
		return err
	}
	_ptr_pwszDataFile := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.DataFile); err != nil {
			return err
		}
		return nil
	})
	_s_pwszDataFile := func(ptr interface{}) { o.DataFile = *ptr.(*string) }
	if err := w.ReadPointer(&o.DataFile, _s_pwszDataFile, _ptr_pwszDataFile); err != nil {
		return err
	}
	return nil
}

// EnumScopeList structure represents DNS_RPC_ENUM_SCOPE_LIST RPC structure.
//
// The DNS_RPC_ENUM_SCOPE_LIST structure<89> contains the list of server scopes configured
// on the DNS server.
type EnumScopeList struct {
	// dwRpcStructureVersion: The DNS management structure version number. This value MUST
	// be set to 0x00000000.
	RPCStructureVersion uint32 `idl:"name:dwRpcStructureVersion" json:"rpc_structure_version"`
	// dwScopeCount: The number of server scopes.
	ScopeCount uint32 `idl:"name:dwScopeCount" json:"scope_count"`
	// ScopeArray: An array containing the names of server scopes.
	ScopeArray []string `idl:"name:ScopeArray;size_is:(dwScopeCount)" json:"scope_array"`
}

func (o *EnumScopeList) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.ScopeArray != nil && o.ScopeCount == 0 {
		o.ScopeCount = uint32(len(o.ScopeArray))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *EnumScopeList) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.ScopeCount)
	return []uint64{
		dimSize1,
	}
}
func (o *EnumScopeList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.RPCStructureVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.ScopeCount); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	for i1 := range o.ScopeArray {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.ScopeArray[i1] != "" {
			_ptr_ScopeArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16String(ctx, w, o.ScopeArray[i1]); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ScopeArray[i1], _ptr_ScopeArray); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.ScopeArray); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *EnumScopeList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.RPCStructureVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.ScopeCount); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.ScopeCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.ScopeCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.ScopeArray", sizeInfo[0])
	}
	o.ScopeArray = make([]string, sizeInfo[0])
	for i1 := range o.ScopeArray {
		i1 := i1
		_ptr_ScopeArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16String(ctx, w, &o.ScopeArray[i1]); err != nil {
				return err
			}
			return nil
		})
		_s_ScopeArray := func(ptr interface{}) { o.ScopeArray[i1] = *ptr.(*string) }
		if err := w.ReadPointer(&o.ScopeArray[i1], _s_ScopeArray, _ptr_ScopeArray); err != nil {
			return err
		}
	}
	return nil
}

// Union structure represents DNSSRV_RPC_UNION RPC union.
//
// The DNSSRV_RPC_UNION specifies a collection of all possible messages structures that
// can be exchanged between a client and server communicating using the DNS Server Management
// Protocol. This is used by the R_DnssrvOperation2 (section 3.1.4.6), R_DnssrvQuery2
// (section 3.1.4.7) and R_DnssrvComplexOperation2 (section 3.1.4.8) method calls. The
// exact message format inside DNSSRV_RPC_UNION is identified by an accompanying DNS_RPC_TYPEID
// (section 2.2.1.1.1) value. Clients and servers of the DNS Server Management Protocol
// SHOULD<7> support all members of DNSSRV_RPC_UNION.
type Union struct {
	// Types that are assignable to Value
	//
	// *Union_Null
	// *Union_Dword
	// *Union_String
	// *Union_UnicodeString
	// *Union_IPArray
	// *Union_Buffer
	// *Union_ServerInfoW2K
	// *Union_Stats
	// *Union_ForwardersW2K
	// *Union_ZoneW2K
	// *Union_ZoneInfoW2K
	// *Union_SecondariesW2K
	// *Union_DatabaseW2K
	// *Union_ZoneCreateW2K
	// *Union_NameAndParam
	// *Union_ZoneListW2K
	// *Union_ServerInfoDotNet
	// *Union_ForwardersDotNet
	// *Union_Zone
	// *Union_ZoneInfoDotNet
	// *Union_SecondariesDotNet
	// *Union_Database
	// *Union_ZoneCreateDotNet
	// *Union_ZoneList
	// *Union_ZoneExport
	// *Union_DirectoryPartition
	// *Union_DirectoryPartitionEnum
	// *Union_DirectoryPartitionList
	// *Union_EnlistDirectoryPartition
	// *Union_ZoneChangeDirectoryPartition
	// *Union_EnumZonesFilter
	// *Union_AddrArray
	// *Union_ServerInfo
	// *Union_ZoneCreate
	// *Union_Forwarders
	// *Union_Secondaries
	// *Union_IPValidate
	// *Union_ZoneInfo
	// *Union_AutoConfigure
	// *Union_UTF8StringList
	// *Union_UnicodeStringList
	// *Union_SKD
	// *Union_SKDList
	// *Union_SKDState
	// *Union_SigningValidationError
	// *Union_TrustPointList
	// *Union_TrustAnchorList
	// *Union_ZoneDNSSecuritySettings
	// *Union_ZoneScopeList
	// *Union_ZoneStats
	// *Union_ScopeCreate
	// *Union_ScopeInfo
	// *Union_ScopeList
	// *Union_SubnetList
	// *Union_Policy
	// *Union_PolicyName
	// *Union_PolicyList
	// *Union_RRLParams
	// *Union_VirtualizationInstance
	// *Union_VirtualizationInstanceList
	Value is_Union `json:"value"`
}

func (o *Union) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *Union_Null:
		if value != nil {
			return value.Null
		}
	case *Union_Dword:
		if value != nil {
			return value.Dword
		}
	case *Union_String:
		if value != nil {
			return value.String
		}
	case *Union_UnicodeString:
		if value != nil {
			return value.UnicodeString
		}
	case *Union_IPArray:
		if value != nil {
			return value.IPArray
		}
	case *Union_Buffer:
		if value != nil {
			return value.Buffer
		}
	case *Union_ServerInfoW2K:
		if value != nil {
			return value.ServerInfoW2K
		}
	case *Union_Stats:
		if value != nil {
			return value.Stats
		}
	case *Union_ForwardersW2K:
		if value != nil {
			return value.ForwardersW2K
		}
	case *Union_ZoneW2K:
		if value != nil {
			return value.ZoneW2K
		}
	case *Union_ZoneInfoW2K:
		if value != nil {
			return value.ZoneInfoW2K
		}
	case *Union_SecondariesW2K:
		if value != nil {
			return value.SecondariesW2K
		}
	case *Union_DatabaseW2K:
		if value != nil {
			return value.DatabaseW2K
		}
	case *Union_ZoneCreateW2K:
		if value != nil {
			return value.ZoneCreateW2K
		}
	case *Union_NameAndParam:
		if value != nil {
			return value.NameAndParam
		}
	case *Union_ZoneListW2K:
		if value != nil {
			return value.ZoneListW2K
		}
	case *Union_ServerInfoDotNet:
		if value != nil {
			return value.ServerInfoDotNet
		}
	case *Union_ForwardersDotNet:
		if value != nil {
			return value.ForwardersDotNet
		}
	case *Union_Zone:
		if value != nil {
			return value.Zone
		}
	case *Union_ZoneInfoDotNet:
		if value != nil {
			return value.ZoneInfoDotNet
		}
	case *Union_SecondariesDotNet:
		if value != nil {
			return value.SecondariesDotNet
		}
	case *Union_Database:
		if value != nil {
			return value.Database
		}
	case *Union_ZoneCreateDotNet:
		if value != nil {
			return value.ZoneCreateDotNet
		}
	case *Union_ZoneList:
		if value != nil {
			return value.ZoneList
		}
	case *Union_ZoneExport:
		if value != nil {
			return value.ZoneExport
		}
	case *Union_DirectoryPartition:
		if value != nil {
			return value.DirectoryPartition
		}
	case *Union_DirectoryPartitionEnum:
		if value != nil {
			return value.DirectoryPartitionEnum
		}
	case *Union_DirectoryPartitionList:
		if value != nil {
			return value.DirectoryPartitionList
		}
	case *Union_EnlistDirectoryPartition:
		if value != nil {
			return value.EnlistDirectoryPartition
		}
	case *Union_ZoneChangeDirectoryPartition:
		if value != nil {
			return value.ZoneChangeDirectoryPartition
		}
	case *Union_EnumZonesFilter:
		if value != nil {
			return value.EnumZonesFilter
		}
	case *Union_AddrArray:
		if value != nil {
			return value.AddrArray
		}
	case *Union_ServerInfo:
		if value != nil {
			return value.ServerInfo
		}
	case *Union_ZoneCreate:
		if value != nil {
			return value.ZoneCreate
		}
	case *Union_Forwarders:
		if value != nil {
			return value.Forwarders
		}
	case *Union_Secondaries:
		if value != nil {
			return value.Secondaries
		}
	case *Union_IPValidate:
		if value != nil {
			return value.IPValidate
		}
	case *Union_ZoneInfo:
		if value != nil {
			return value.ZoneInfo
		}
	case *Union_AutoConfigure:
		if value != nil {
			return value.AutoConfigure
		}
	case *Union_UTF8StringList:
		if value != nil {
			return value.UTF8StringList
		}
	case *Union_UnicodeStringList:
		if value != nil {
			return value.UnicodeStringList
		}
	case *Union_SKD:
		if value != nil {
			return value.SKD
		}
	case *Union_SKDList:
		if value != nil {
			return value.SKDList
		}
	case *Union_SKDState:
		if value != nil {
			return value.SKDState
		}
	case *Union_SigningValidationError:
		if value != nil {
			return value.SigningValidationError
		}
	case *Union_TrustPointList:
		if value != nil {
			return value.TrustPointList
		}
	case *Union_TrustAnchorList:
		if value != nil {
			return value.TrustAnchorList
		}
	case *Union_ZoneDNSSecuritySettings:
		if value != nil {
			return value.ZoneDNSSecuritySettings
		}
	case *Union_ZoneScopeList:
		if value != nil {
			return value.ZoneScopeList
		}
	case *Union_ZoneStats:
		if value != nil {
			return value.ZoneStats
		}
	case *Union_ScopeCreate:
		if value != nil {
			return value.ScopeCreate
		}
	case *Union_ScopeInfo:
		if value != nil {
			return value.ScopeInfo
		}
	case *Union_ScopeList:
		if value != nil {
			return value.ScopeList
		}
	case *Union_SubnetList:
		if value != nil {
			return value.SubnetList
		}
	case *Union_Policy:
		if value != nil {
			return value.Policy
		}
	case *Union_PolicyName:
		if value != nil {
			return value.PolicyName
		}
	case *Union_PolicyList:
		if value != nil {
			return value.PolicyList
		}
	case *Union_RRLParams:
		if value != nil {
			return value.RRLParams
		}
	case *Union_VirtualizationInstance:
		if value != nil {
			return value.VirtualizationInstance
		}
	case *Union_VirtualizationInstanceList:
		if value != nil {
			return value.VirtualizationInstanceList
		}
	}
	return nil
}

type is_Union interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_Union()
}

func (o *Union) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *Union_Null:
		return uint32(0)
	case *Union_Dword:
		return uint32(1)
	case *Union_String:
		return uint32(2)
	case *Union_UnicodeString:
		return uint32(3)
	case *Union_IPArray:
		return uint32(4)
	case *Union_Buffer:
		return uint32(5)
	case *Union_ServerInfoW2K:
		return uint32(6)
	case *Union_Stats:
		return uint32(7)
	case *Union_ForwardersW2K:
		return uint32(8)
	case *Union_ZoneW2K:
		return uint32(9)
	case *Union_ZoneInfoW2K:
		return uint32(10)
	case *Union_SecondariesW2K:
		return uint32(11)
	case *Union_DatabaseW2K:
		return uint32(12)
	case *Union_ZoneCreateW2K:
		return uint32(14)
	case *Union_NameAndParam:
		return uint32(15)
	case *Union_ZoneListW2K:
		return uint32(16)
	case *Union_ServerInfoDotNet:
		return uint32(19)
	case *Union_ForwardersDotNet:
		return uint32(20)
	case *Union_Zone:
		return uint32(21)
	case *Union_ZoneInfoDotNet:
		return uint32(22)
	case *Union_SecondariesDotNet:
		return uint32(23)
	case *Union_Database:
		return uint32(24)
	case *Union_ZoneCreateDotNet:
		return uint32(26)
	case *Union_ZoneList:
		return uint32(27)
	case *Union_ZoneExport:
		return uint32(18)
	case *Union_DirectoryPartition:
		return uint32(29)
	case *Union_DirectoryPartitionEnum:
		return uint32(28)
	case *Union_DirectoryPartitionList:
		return uint32(30)
	case *Union_EnlistDirectoryPartition:
		return uint32(31)
	case *Union_ZoneChangeDirectoryPartition:
		return uint32(32)
	case *Union_EnumZonesFilter:
		return uint32(33)
	case *Union_AddrArray:
		return uint32(34)
	case *Union_ServerInfo:
		return uint32(35)
	case *Union_ZoneCreate:
		return uint32(40)
	case *Union_Forwarders:
		return uint32(37)
	case *Union_Secondaries:
		return uint32(38)
	case *Union_IPValidate:
		return uint32(41)
	case *Union_ZoneInfo:
		return uint32(36)
	case *Union_AutoConfigure:
		return uint32(42)
	case *Union_UTF8StringList:
		return uint32(43)
	case *Union_UnicodeStringList:
		return uint32(44)
	case *Union_SKD:
		return uint32(45)
	case *Union_SKDList:
		return uint32(46)
	case *Union_SKDState:
		return uint32(47)
	case *Union_SigningValidationError:
		return uint32(48)
	case *Union_TrustPointList:
		return uint32(49)
	case *Union_TrustAnchorList:
		return uint32(50)
	case *Union_ZoneDNSSecuritySettings:
		return uint32(51)
	case *Union_ZoneScopeList:
		return uint32(52)
	case *Union_ZoneStats:
		return uint32(53)
	case *Union_ScopeCreate:
		return uint32(54)
	case *Union_ScopeInfo:
		return uint32(55)
	case *Union_ScopeList:
		return uint32(56)
	case *Union_SubnetList:
		return uint32(57)
	case *Union_Policy:
		return uint32(58)
	case *Union_PolicyName:
		return uint32(59)
	case *Union_PolicyList:
		return uint32(60)
	case *Union_RRLParams:
		return uint32(61)
	case *Union_VirtualizationInstance:
		return uint32(62)
	case *Union_VirtualizationInstanceList:
		return uint32(63)
	}
	return uint32(0)
}

func (o *Union) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	if err := w.WriteSwitch(uint32(sw)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint32(0):
		_o, _ := o.Value.(*Union_Null)
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
	case uint32(1):
		_o, _ := o.Value.(*Union_Dword)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_Dword{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(2):
		_o, _ := o.Value.(*Union_String)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_String{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(3):
		_o, _ := o.Value.(*Union_UnicodeString)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(4):
		_o, _ := o.Value.(*Union_IPArray)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_IPArray{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(5):
		_o, _ := o.Value.(*Union_Buffer)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_Buffer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(6):
		_o, _ := o.Value.(*Union_ServerInfoW2K)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_ServerInfoW2K{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(7):
		_o, _ := o.Value.(*Union_Stats)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_Stats{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(8):
		_o, _ := o.Value.(*Union_ForwardersW2K)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_ForwardersW2K{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(9):
		_o, _ := o.Value.(*Union_ZoneW2K)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_ZoneW2K{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(10):
		_o, _ := o.Value.(*Union_ZoneInfoW2K)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_ZoneInfoW2K{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(11):
		_o, _ := o.Value.(*Union_SecondariesW2K)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_SecondariesW2K{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(12):
		_o, _ := o.Value.(*Union_DatabaseW2K)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_DatabaseW2K{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(14):
		_o, _ := o.Value.(*Union_ZoneCreateW2K)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_ZoneCreateW2K{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(15):
		_o, _ := o.Value.(*Union_NameAndParam)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_NameAndParam{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(16):
		_o, _ := o.Value.(*Union_ZoneListW2K)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_ZoneListW2K{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(19):
		_o, _ := o.Value.(*Union_ServerInfoDotNet)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_ServerInfoDotNet{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(20):
		_o, _ := o.Value.(*Union_ForwardersDotNet)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_ForwardersDotNet{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(21):
		_o, _ := o.Value.(*Union_Zone)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_Zone{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(22):
		_o, _ := o.Value.(*Union_ZoneInfoDotNet)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_ZoneInfoDotNet{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(23):
		_o, _ := o.Value.(*Union_SecondariesDotNet)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_SecondariesDotNet{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(24):
		_o, _ := o.Value.(*Union_Database)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_Database{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(26):
		_o, _ := o.Value.(*Union_ZoneCreateDotNet)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_ZoneCreateDotNet{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(27):
		_o, _ := o.Value.(*Union_ZoneList)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_ZoneList{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(18):
		_o, _ := o.Value.(*Union_ZoneExport)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_ZoneExport{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(29):
		_o, _ := o.Value.(*Union_DirectoryPartition)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_DirectoryPartition{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(28):
		_o, _ := o.Value.(*Union_DirectoryPartitionEnum)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_DirectoryPartitionEnum{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(30):
		_o, _ := o.Value.(*Union_DirectoryPartitionList)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_DirectoryPartitionList{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(31):
		_o, _ := o.Value.(*Union_EnlistDirectoryPartition)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_EnlistDirectoryPartition{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(32):
		_o, _ := o.Value.(*Union_ZoneChangeDirectoryPartition)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_ZoneChangeDirectoryPartition{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(33):
		_o, _ := o.Value.(*Union_EnumZonesFilter)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_EnumZonesFilter{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(34):
		_o, _ := o.Value.(*Union_AddrArray)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_AddrArray{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(35):
		_o, _ := o.Value.(*Union_ServerInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_ServerInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(40):
		_o, _ := o.Value.(*Union_ZoneCreate)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_ZoneCreate{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(37):
		_o, _ := o.Value.(*Union_Forwarders)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_Forwarders{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(38):
		_o, _ := o.Value.(*Union_Secondaries)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_Secondaries{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(41):
		_o, _ := o.Value.(*Union_IPValidate)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_IPValidate{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(36):
		_o, _ := o.Value.(*Union_ZoneInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_ZoneInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(42):
		_o, _ := o.Value.(*Union_AutoConfigure)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_AutoConfigure{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(43):
		_o, _ := o.Value.(*Union_UTF8StringList)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_UTF8StringList{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(44):
		_o, _ := o.Value.(*Union_UnicodeStringList)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_UnicodeStringList{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(45):
		_o, _ := o.Value.(*Union_SKD)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_SKD{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(46):
		_o, _ := o.Value.(*Union_SKDList)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_SKDList{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(47):
		_o, _ := o.Value.(*Union_SKDState)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_SKDState{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(48):
		_o, _ := o.Value.(*Union_SigningValidationError)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_SigningValidationError{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(49):
		_o, _ := o.Value.(*Union_TrustPointList)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_TrustPointList{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(50):
		_o, _ := o.Value.(*Union_TrustAnchorList)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_TrustAnchorList{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(51):
		_o, _ := o.Value.(*Union_ZoneDNSSecuritySettings)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_ZoneDNSSecuritySettings{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(52):
		_o, _ := o.Value.(*Union_ZoneScopeList)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_ZoneScopeList{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(53):
		_o, _ := o.Value.(*Union_ZoneStats)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_ZoneStats{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(54):
		_o, _ := o.Value.(*Union_ScopeCreate)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_ScopeCreate{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(55):
		_o, _ := o.Value.(*Union_ScopeInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_ScopeInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(56):
		_o, _ := o.Value.(*Union_ScopeList)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_ScopeList{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(57):
		_o, _ := o.Value.(*Union_SubnetList)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_SubnetList{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(58):
		_o, _ := o.Value.(*Union_Policy)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_Policy{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(59):
		_o, _ := o.Value.(*Union_PolicyName)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_PolicyName{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(60):
		_o, _ := o.Value.(*Union_PolicyList)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_PolicyList{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(61):
		_o, _ := o.Value.(*Union_RRLParams)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_RRLParams{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(62):
		_o, _ := o.Value.(*Union_VirtualizationInstance)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_VirtualizationInstance{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(63):
		_o, _ := o.Value.(*Union_VirtualizationInstanceList)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Union_VirtualizationInstanceList{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *Union) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	if err := w.ReadSwitch((*uint32)(&sw)); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint32(0):
		_ptr_o := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			o.Value = &Union_Null{}
			if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_o := func(ptr interface{}) { o.Value = *ptr.(**Union_Null) }
		if err := w.ReadPointer(&o.Value, _s_o, _ptr_o); err != nil {
			return err
		}
	case uint32(1):
		o.Value = &Union_Dword{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(2):
		o.Value = &Union_String{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(3):
		o.Value = &Union_UnicodeString{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(4):
		o.Value = &Union_IPArray{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(5):
		o.Value = &Union_Buffer{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(6):
		o.Value = &Union_ServerInfoW2K{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(7):
		o.Value = &Union_Stats{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(8):
		o.Value = &Union_ForwardersW2K{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(9):
		o.Value = &Union_ZoneW2K{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(10):
		o.Value = &Union_ZoneInfoW2K{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(11):
		o.Value = &Union_SecondariesW2K{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(12):
		o.Value = &Union_DatabaseW2K{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(14):
		o.Value = &Union_ZoneCreateW2K{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(15):
		o.Value = &Union_NameAndParam{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(16):
		o.Value = &Union_ZoneListW2K{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(19):
		o.Value = &Union_ServerInfoDotNet{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(20):
		o.Value = &Union_ForwardersDotNet{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(21):
		o.Value = &Union_Zone{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(22):
		o.Value = &Union_ZoneInfoDotNet{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(23):
		o.Value = &Union_SecondariesDotNet{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(24):
		o.Value = &Union_Database{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(26):
		o.Value = &Union_ZoneCreateDotNet{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(27):
		o.Value = &Union_ZoneList{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(18):
		o.Value = &Union_ZoneExport{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(29):
		o.Value = &Union_DirectoryPartition{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(28):
		o.Value = &Union_DirectoryPartitionEnum{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(30):
		o.Value = &Union_DirectoryPartitionList{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(31):
		o.Value = &Union_EnlistDirectoryPartition{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(32):
		o.Value = &Union_ZoneChangeDirectoryPartition{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(33):
		o.Value = &Union_EnumZonesFilter{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(34):
		o.Value = &Union_AddrArray{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(35):
		o.Value = &Union_ServerInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(40):
		o.Value = &Union_ZoneCreate{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(37):
		o.Value = &Union_Forwarders{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(38):
		o.Value = &Union_Secondaries{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(41):
		o.Value = &Union_IPValidate{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(36):
		o.Value = &Union_ZoneInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(42):
		o.Value = &Union_AutoConfigure{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(43):
		o.Value = &Union_UTF8StringList{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(44):
		o.Value = &Union_UnicodeStringList{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(45):
		o.Value = &Union_SKD{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(46):
		o.Value = &Union_SKDList{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(47):
		o.Value = &Union_SKDState{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(48):
		o.Value = &Union_SigningValidationError{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(49):
		o.Value = &Union_TrustPointList{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(50):
		o.Value = &Union_TrustAnchorList{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(51):
		o.Value = &Union_ZoneDNSSecuritySettings{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(52):
		o.Value = &Union_ZoneScopeList{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(53):
		o.Value = &Union_ZoneStats{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(54):
		o.Value = &Union_ScopeCreate{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(55):
		o.Value = &Union_ScopeInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(56):
		o.Value = &Union_ScopeList{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(57):
		o.Value = &Union_SubnetList{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(58):
		o.Value = &Union_Policy{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(59):
		o.Value = &Union_PolicyName{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(60):
		o.Value = &Union_PolicyList{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(61):
		o.Value = &Union_RRLParams{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(62):
		o.Value = &Union_VirtualizationInstance{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(63):
		o.Value = &Union_VirtualizationInstanceList{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// Union_Null structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 0
type Union_Null struct {
	// Null: No data is provided.
	Null uint8 `idl:"name:Null" json:"null"`
}

func (*Union_Null) is_Union() {}

func (o *Union_Null) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Null); err != nil {
		return err
	}
	return nil
}
func (o *Union_Null) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Null); err != nil {
		return err
	}
	return nil
}

// Union_Dword structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 1
type Union_Dword struct {
	// Dword: Data is a DWORD value.
	Dword uint32 `idl:"name:Dword" json:"dword"`
}

func (*Union_Dword) is_Union() {}

func (o *Union_Dword) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Dword); err != nil {
		return err
	}
	return nil
}
func (o *Union_Dword) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Dword); err != nil {
		return err
	}
	return nil
}

// Union_String structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 2
type Union_String struct {
	// String: A pointer to a null-terminated UTF-8 string or a NULL pointer.
	String string `idl:"name:String;string" json:"string"`
}

func (*Union_String) is_Union() {}

func (o *Union_String) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.String != "" {
		_ptr_String := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.String); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.String, _ptr_String); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_String) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_String := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.String); err != nil {
			return err
		}
		return nil
	})
	_s_String := func(ptr interface{}) { o.String = *ptr.(*string) }
	if err := w.ReadPointer(&o.String, _s_String, _ptr_String); err != nil {
		return err
	}
	return nil
}

// Union_UnicodeString structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 3
type Union_UnicodeString struct {
	// WideString: A pointer to a null-terminated Unicode string or a NULL pointer.
	UnicodeString string `idl:"name:WideString;string" json:"unicode_string"`
}

func (*Union_UnicodeString) is_Union() {}

func (o *Union_UnicodeString) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.UnicodeString != "" {
		_ptr_WideString := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.UnicodeString); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.UnicodeString, _ptr_WideString); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_UnicodeString) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_WideString := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.UnicodeString); err != nil {
			return err
		}
		return nil
	})
	_s_WideString := func(ptr interface{}) { o.UnicodeString = *ptr.(*string) }
	if err := w.ReadPointer(&o.UnicodeString, _s_WideString, _ptr_WideString); err != nil {
		return err
	}
	return nil
}

// Union_IPArray structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 4
type Union_IPArray struct {
	// IpArray: An array of IPv4 addresses in IP4_ARRAY (section 2.2.3.2.1) format.
	IPArray *IPv4Array `idl:"name:IpArray" json:"ip_array"`
}

func (*Union_IPArray) is_Union() {}

func (o *Union_IPArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.IPArray != nil {
		_ptr_IpArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.IPArray != nil {
				if err := o.IPArray.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPv4Array{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.IPArray, _ptr_IpArray); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_IPArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_IpArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.IPArray == nil {
			o.IPArray = &IPv4Array{}
		}
		if err := o.IPArray.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_IpArray := func(ptr interface{}) { o.IPArray = *ptr.(**IPv4Array) }
	if err := w.ReadPointer(&o.IPArray, _s_IpArray, _ptr_IpArray); err != nil {
		return err
	}
	return nil
}

// Union_Buffer structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 5
type Union_Buffer struct {
	// Buffer: A pointer to a DNS_RPC_BUFFER (section 2.2.1.2.2).
	Buffer *Buffer `idl:"name:Buffer" json:"buffer"`
}

func (*Union_Buffer) is_Union() {}

func (o *Union_Buffer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Buffer != nil {
		_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Buffer != nil {
				if err := o.Buffer.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Buffer{}).MarshalNDR(ctx, w); err != nil {
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
func (o *Union_Buffer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Buffer == nil {
			o.Buffer = &Buffer{}
		}
		if err := o.Buffer.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(**Buffer) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// Union_ServerInfoW2K structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 6
type Union_ServerInfoW2K struct {
	// ServerInfoW2K: A pointer to a structure of type DNS_RPC_SERVER_INFO_W2K (section
	// 2.2.4.2.2.1). This structure is used to specify the general DNS server state and
	// configuration.
	ServerInfoW2K *ServerInfoW2K `idl:"name:ServerInfoW2K" json:"server_info_w2k"`
}

func (*Union_ServerInfoW2K) is_Union() {}

func (o *Union_ServerInfoW2K) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ServerInfoW2K != nil {
		_ptr_ServerInfoW2K := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ServerInfoW2K != nil {
				if err := o.ServerInfoW2K.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ServerInfoW2K{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ServerInfoW2K, _ptr_ServerInfoW2K); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_ServerInfoW2K) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ServerInfoW2K := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ServerInfoW2K == nil {
			o.ServerInfoW2K = &ServerInfoW2K{}
		}
		if err := o.ServerInfoW2K.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ServerInfoW2K := func(ptr interface{}) { o.ServerInfoW2K = *ptr.(**ServerInfoW2K) }
	if err := w.ReadPointer(&o.ServerInfoW2K, _s_ServerInfoW2K, _ptr_ServerInfoW2K); err != nil {
		return err
	}
	return nil
}

// Union_Stats structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 7
type Union_Stats struct {
	// Stats: A pointer to a structure of type DNSSRV_STATS (section 2.2.10.2.2). The structure
	// exposes internal statistics and counters.
	Stats *Stat `idl:"name:Stats" json:"stats"`
}

func (*Union_Stats) is_Union() {}

func (o *Union_Stats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Stats != nil {
		_ptr_Stats := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Stats != nil {
				if err := o.Stats.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Stat{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Stats, _ptr_Stats); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_Stats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_Stats := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Stats == nil {
			o.Stats = &Stat{}
		}
		if err := o.Stats.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Stats := func(ptr interface{}) { o.Stats = *ptr.(**Stat) }
	if err := w.ReadPointer(&o.Stats, _s_Stats, _ptr_Stats); err != nil {
		return err
	}
	return nil
}

// Union_ForwardersW2K structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 8
type Union_ForwardersW2K struct {
	// ForwardersW2K: A pointer to a structure of type DNS_RPC_FORWARDERS_W2K (section 2.2.5.2.10.1).
	// This structure specifies the set of DNS servers this DNS server will forward unresolved
	// queries to.
	ForwardersW2K *ForwardersW2K `idl:"name:ForwardersW2K" json:"forwarders_w2k"`
}

func (*Union_ForwardersW2K) is_Union() {}

func (o *Union_ForwardersW2K) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ForwardersW2K != nil {
		_ptr_ForwardersW2K := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ForwardersW2K != nil {
				if err := o.ForwardersW2K.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ForwardersW2K{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ForwardersW2K, _ptr_ForwardersW2K); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_ForwardersW2K) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ForwardersW2K := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ForwardersW2K == nil {
			o.ForwardersW2K = &ForwardersW2K{}
		}
		if err := o.ForwardersW2K.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ForwardersW2K := func(ptr interface{}) { o.ForwardersW2K = *ptr.(**ForwardersW2K) }
	if err := w.ReadPointer(&o.ForwardersW2K, _s_ForwardersW2K, _ptr_ForwardersW2K); err != nil {
		return err
	}
	return nil
}

// Union_ZoneW2K structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 9
type Union_ZoneW2K struct {
	// ZoneW2K: A pointer to a structure of type DNS_RPC_ZONE_W2K (section 2.2.5.2.1.1).
	// This structure is used to specify basic information about a DNS zone.
	ZoneW2K *ZoneW2K `idl:"name:ZoneW2K" json:"zone_w2k"`
}

func (*Union_ZoneW2K) is_Union() {}

func (o *Union_ZoneW2K) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ZoneW2K != nil {
		_ptr_ZoneW2K := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ZoneW2K != nil {
				if err := o.ZoneW2K.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ZoneW2K{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ZoneW2K, _ptr_ZoneW2K); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_ZoneW2K) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ZoneW2K := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ZoneW2K == nil {
			o.ZoneW2K = &ZoneW2K{}
		}
		if err := o.ZoneW2K.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ZoneW2K := func(ptr interface{}) { o.ZoneW2K = *ptr.(**ZoneW2K) }
	if err := w.ReadPointer(&o.ZoneW2K, _s_ZoneW2K, _ptr_ZoneW2K); err != nil {
		return err
	}
	return nil
}

// Union_ZoneInfoW2K structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 10
type Union_ZoneInfoW2K struct {
	// ZoneInfoW2K: A pointer to a structure of type DNS_RPC_ZONE_INFO_W2K (section 2.2.5.2.4.1).
	// This structure is used to specify detailed DNS zone information.
	ZoneInfoW2K *ZoneInfoW2K `idl:"name:ZoneInfoW2K" json:"zone_info_w2k"`
}

func (*Union_ZoneInfoW2K) is_Union() {}

func (o *Union_ZoneInfoW2K) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ZoneInfoW2K != nil {
		_ptr_ZoneInfoW2K := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ZoneInfoW2K != nil {
				if err := o.ZoneInfoW2K.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ZoneInfoW2K{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ZoneInfoW2K, _ptr_ZoneInfoW2K); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_ZoneInfoW2K) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ZoneInfoW2K := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ZoneInfoW2K == nil {
			o.ZoneInfoW2K = &ZoneInfoW2K{}
		}
		if err := o.ZoneInfoW2K.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ZoneInfoW2K := func(ptr interface{}) { o.ZoneInfoW2K = *ptr.(**ZoneInfoW2K) }
	if err := w.ReadPointer(&o.ZoneInfoW2K, _s_ZoneInfoW2K, _ptr_ZoneInfoW2K); err != nil {
		return err
	}
	return nil
}

// Union_SecondariesW2K structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 11
type Union_SecondariesW2K struct {
	// SecondariesW2K: A pointer to a structure of type DNS_RPC_ZONE_SECONDARIES_W2K (section
	// 2.2.5.2.5.1). This structure is used to specify information about the secondary servers
	// for a primary DNS zone.
	SecondariesW2K *ZoneSecondariesW2K `idl:"name:SecondariesW2K" json:"secondaries_w2k"`
}

func (*Union_SecondariesW2K) is_Union() {}

func (o *Union_SecondariesW2K) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.SecondariesW2K != nil {
		_ptr_SecondariesW2K := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SecondariesW2K != nil {
				if err := o.SecondariesW2K.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ZoneSecondariesW2K{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SecondariesW2K, _ptr_SecondariesW2K); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_SecondariesW2K) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_SecondariesW2K := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SecondariesW2K == nil {
			o.SecondariesW2K = &ZoneSecondariesW2K{}
		}
		if err := o.SecondariesW2K.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_SecondariesW2K := func(ptr interface{}) { o.SecondariesW2K = *ptr.(**ZoneSecondariesW2K) }
	if err := w.ReadPointer(&o.SecondariesW2K, _s_SecondariesW2K, _ptr_SecondariesW2K); err != nil {
		return err
	}
	return nil
}

// Union_DatabaseW2K structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 12
type Union_DatabaseW2K struct {
	// DatabaseW2K: A pointer to a structure of type DNS_RPC_ZONE_DATABASE_W2K (section
	// 2.2.5.2.6.1). This structure specifies how a DNS zone is stored in persistent storage.
	DatabaseW2K *ZoneDatabaseW2K `idl:"name:DatabaseW2K" json:"database_w2k"`
}

func (*Union_DatabaseW2K) is_Union() {}

func (o *Union_DatabaseW2K) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.DatabaseW2K != nil {
		_ptr_DatabaseW2K := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.DatabaseW2K != nil {
				if err := o.DatabaseW2K.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ZoneDatabaseW2K{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.DatabaseW2K, _ptr_DatabaseW2K); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_DatabaseW2K) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DatabaseW2K := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.DatabaseW2K == nil {
			o.DatabaseW2K = &ZoneDatabaseW2K{}
		}
		if err := o.DatabaseW2K.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DatabaseW2K := func(ptr interface{}) { o.DatabaseW2K = *ptr.(**ZoneDatabaseW2K) }
	if err := w.ReadPointer(&o.DatabaseW2K, _s_DatabaseW2K, _ptr_DatabaseW2K); err != nil {
		return err
	}
	return nil
}

// Union_ZoneCreateW2K structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 14
type Union_ZoneCreateW2K struct {
	// ZoneCreateW2K: A pointer to a structure of type DNS_RPC_ZONE_CREATE_INFO_W2K (section
	// 2.2.5.2.7.1). This structure is used to specify the parameters required when creating
	// a new DNS zone.
	ZoneCreateW2K *ZoneCreateInfoW2K `idl:"name:ZoneCreateW2K" json:"zone_create_w2k"`
}

func (*Union_ZoneCreateW2K) is_Union() {}

func (o *Union_ZoneCreateW2K) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ZoneCreateW2K != nil {
		_ptr_ZoneCreateW2K := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ZoneCreateW2K != nil {
				if err := o.ZoneCreateW2K.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ZoneCreateInfoW2K{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ZoneCreateW2K, _ptr_ZoneCreateW2K); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_ZoneCreateW2K) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ZoneCreateW2K := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ZoneCreateW2K == nil {
			o.ZoneCreateW2K = &ZoneCreateInfoW2K{}
		}
		if err := o.ZoneCreateW2K.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ZoneCreateW2K := func(ptr interface{}) { o.ZoneCreateW2K = *ptr.(**ZoneCreateInfoW2K) }
	if err := w.ReadPointer(&o.ZoneCreateW2K, _s_ZoneCreateW2K, _ptr_ZoneCreateW2K); err != nil {
		return err
	}
	return nil
}

// Union_NameAndParam structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 15
type Union_NameAndParam struct {
	// NameAndParam: A pointer to a structure of type DNS_RPC_NAME_AND_PARAM (section 2.2.1.2.5).
	// This is a general purpose structure used to associate a parameter name with a DWORD
	// value.
	NameAndParam *NameAndParam `idl:"name:NameAndParam" json:"name_and_param"`
}

func (*Union_NameAndParam) is_Union() {}

func (o *Union_NameAndParam) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.NameAndParam != nil {
		_ptr_NameAndParam := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.NameAndParam != nil {
				if err := o.NameAndParam.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NameAndParam{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.NameAndParam, _ptr_NameAndParam); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_NameAndParam) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_NameAndParam := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.NameAndParam == nil {
			o.NameAndParam = &NameAndParam{}
		}
		if err := o.NameAndParam.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_NameAndParam := func(ptr interface{}) { o.NameAndParam = *ptr.(**NameAndParam) }
	if err := w.ReadPointer(&o.NameAndParam, _s_NameAndParam, _ptr_NameAndParam); err != nil {
		return err
	}
	return nil
}

// Union_ZoneListW2K structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 16
type Union_ZoneListW2K struct {
	// ZoneListW2K: A pointer to a structure of type DNS_RPC_ZONE_LIST_W2K (section 2.2.5.2.3.1).
	// This structure is used to enumerate zones.
	ZoneListW2K *ZoneListW2K `idl:"name:ZoneListW2K" json:"zone_list_w2k"`
}

func (*Union_ZoneListW2K) is_Union() {}

func (o *Union_ZoneListW2K) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ZoneListW2K != nil {
		_ptr_ZoneListW2K := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ZoneListW2K != nil {
				if err := o.ZoneListW2K.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ZoneListW2K{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ZoneListW2K, _ptr_ZoneListW2K); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_ZoneListW2K) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ZoneListW2K := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ZoneListW2K == nil {
			o.ZoneListW2K = &ZoneListW2K{}
		}
		if err := o.ZoneListW2K.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ZoneListW2K := func(ptr interface{}) { o.ZoneListW2K = *ptr.(**ZoneListW2K) }
	if err := w.ReadPointer(&o.ZoneListW2K, _s_ZoneListW2K, _ptr_ZoneListW2K); err != nil {
		return err
	}
	return nil
}

// Union_ServerInfoDotNet structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 19
type Union_ServerInfoDotNet struct {
	// ServerInfoDotNet: A pointer to a structure of type DNS_RPC_SERVER_INFO_DOTNET (section
	// 2.2.4.2.2.2). This structure is used to specify the general DNS server state and
	// configuration.
	ServerInfoDotNet *ServerInfo_NET `idl:"name:ServerInfoDotNet" json:"server_info_dot_net"`
}

func (*Union_ServerInfoDotNet) is_Union() {}

func (o *Union_ServerInfoDotNet) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ServerInfoDotNet != nil {
		_ptr_ServerInfoDotNet := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ServerInfoDotNet != nil {
				if err := o.ServerInfoDotNet.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ServerInfo_NET{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ServerInfoDotNet, _ptr_ServerInfoDotNet); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_ServerInfoDotNet) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ServerInfoDotNet := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ServerInfoDotNet == nil {
			o.ServerInfoDotNet = &ServerInfo_NET{}
		}
		if err := o.ServerInfoDotNet.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ServerInfoDotNet := func(ptr interface{}) { o.ServerInfoDotNet = *ptr.(**ServerInfo_NET) }
	if err := w.ReadPointer(&o.ServerInfoDotNet, _s_ServerInfoDotNet, _ptr_ServerInfoDotNet); err != nil {
		return err
	}
	return nil
}

// Union_ForwardersDotNet structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 20
type Union_ForwardersDotNet struct {
	// ForwardersDotNet: A pointer to a structure of type DNS_RPC_FORWARDERS_DOTNET (section
	// 2.2.5.2.10.2). This structure specifies the set of DNS servers this DNS server will
	// forward unresolved queries to.
	ForwardersDotNet *Forwarders_NET `idl:"name:ForwardersDotNet" json:"forwarders_dot_net"`
}

func (*Union_ForwardersDotNet) is_Union() {}

func (o *Union_ForwardersDotNet) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ForwardersDotNet != nil {
		_ptr_ForwardersDotNet := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ForwardersDotNet != nil {
				if err := o.ForwardersDotNet.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Forwarders_NET{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ForwardersDotNet, _ptr_ForwardersDotNet); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_ForwardersDotNet) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ForwardersDotNet := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ForwardersDotNet == nil {
			o.ForwardersDotNet = &Forwarders_NET{}
		}
		if err := o.ForwardersDotNet.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ForwardersDotNet := func(ptr interface{}) { o.ForwardersDotNet = *ptr.(**Forwarders_NET) }
	if err := w.ReadPointer(&o.ForwardersDotNet, _s_ForwardersDotNet, _ptr_ForwardersDotNet); err != nil {
		return err
	}
	return nil
}

// Union_Zone structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 21
type Union_Zone struct {
	// Zone: A pointer to a structure of type DNS_RPC_ZONE (section 2.2.5.2.1). This structure
	// is used to specify basic information about a DNS zone.
	Zone *Zone `idl:"name:Zone" json:"zone"`
}

func (*Union_Zone) is_Union() {}

func (o *Union_Zone) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Zone != nil {
		_ptr_Zone := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Zone != nil {
				if err := o.Zone.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Zone{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Zone, _ptr_Zone); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_Zone) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_Zone := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Zone == nil {
			o.Zone = &Zone{}
		}
		if err := o.Zone.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Zone := func(ptr interface{}) { o.Zone = *ptr.(**Zone) }
	if err := w.ReadPointer(&o.Zone, _s_Zone, _ptr_Zone); err != nil {
		return err
	}
	return nil
}

// Union_ZoneInfoDotNet structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 22
type Union_ZoneInfoDotNet struct {
	// ZoneInfoDotNet: A pointer to a structure of type DNS_RPC_ZONE_INFO_DOTNET (section
	// 2.2.5.2.4.2). This structure is used to specify detailed DNS zone information.
	ZoneInfoDotNet *ZoneInfo_NET `idl:"name:ZoneInfoDotNet" json:"zone_info_dot_net"`
}

func (*Union_ZoneInfoDotNet) is_Union() {}

func (o *Union_ZoneInfoDotNet) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ZoneInfoDotNet != nil {
		_ptr_ZoneInfoDotNet := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ZoneInfoDotNet != nil {
				if err := o.ZoneInfoDotNet.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ZoneInfo_NET{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ZoneInfoDotNet, _ptr_ZoneInfoDotNet); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_ZoneInfoDotNet) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ZoneInfoDotNet := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ZoneInfoDotNet == nil {
			o.ZoneInfoDotNet = &ZoneInfo_NET{}
		}
		if err := o.ZoneInfoDotNet.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ZoneInfoDotNet := func(ptr interface{}) { o.ZoneInfoDotNet = *ptr.(**ZoneInfo_NET) }
	if err := w.ReadPointer(&o.ZoneInfoDotNet, _s_ZoneInfoDotNet, _ptr_ZoneInfoDotNet); err != nil {
		return err
	}
	return nil
}

// Union_SecondariesDotNet structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 23
type Union_SecondariesDotNet struct {
	// SecondariesDotNet: A pointer to a structure of type DNS_RPC_ZONE_SECONDARIES_DOTNET
	// (section 2.2.5.2.5.2). This structure is used to specify information about the secondary
	// servers for a primary DNS zone.
	SecondariesDotNet *ZoneSecondaries_NET `idl:"name:SecondariesDotNet" json:"secondaries_dot_net"`
}

func (*Union_SecondariesDotNet) is_Union() {}

func (o *Union_SecondariesDotNet) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.SecondariesDotNet != nil {
		_ptr_SecondariesDotNet := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SecondariesDotNet != nil {
				if err := o.SecondariesDotNet.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ZoneSecondaries_NET{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SecondariesDotNet, _ptr_SecondariesDotNet); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_SecondariesDotNet) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_SecondariesDotNet := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SecondariesDotNet == nil {
			o.SecondariesDotNet = &ZoneSecondaries_NET{}
		}
		if err := o.SecondariesDotNet.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_SecondariesDotNet := func(ptr interface{}) { o.SecondariesDotNet = *ptr.(**ZoneSecondaries_NET) }
	if err := w.ReadPointer(&o.SecondariesDotNet, _s_SecondariesDotNet, _ptr_SecondariesDotNet); err != nil {
		return err
	}
	return nil
}

// Union_Database structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 24
type Union_Database struct {
	// Database: A pointer to a structure of type DNS_RPC_ZONE_DATABASE (section 2.2.5.2.6).
	// This structure specifies how a DNS zone is stored in persistent storage.
	Database *ZoneDatabase `idl:"name:Database" json:"database"`
}

func (*Union_Database) is_Union() {}

func (o *Union_Database) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Database != nil {
		_ptr_Database := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Database != nil {
				if err := o.Database.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ZoneDatabase{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Database, _ptr_Database); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_Database) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_Database := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Database == nil {
			o.Database = &ZoneDatabase{}
		}
		if err := o.Database.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Database := func(ptr interface{}) { o.Database = *ptr.(**ZoneDatabase) }
	if err := w.ReadPointer(&o.Database, _s_Database, _ptr_Database); err != nil {
		return err
	}
	return nil
}

// Union_ZoneCreateDotNet structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 26
type Union_ZoneCreateDotNet struct {
	// ZoneCreateDotNet: A pointer to a structure of type DNS_RPC_ZONE_CREATE_INFO_DOTNET
	// (section 2.2.5.2.7.2). This structure is used to specify the parameters required
	// when creating a new DNS zone.
	ZoneCreateDotNet *ZoneCreateInfo_NET `idl:"name:ZoneCreateDotNet" json:"zone_create_dot_net"`
}

func (*Union_ZoneCreateDotNet) is_Union() {}

func (o *Union_ZoneCreateDotNet) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ZoneCreateDotNet != nil {
		_ptr_ZoneCreateDotNet := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ZoneCreateDotNet != nil {
				if err := o.ZoneCreateDotNet.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ZoneCreateInfo_NET{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ZoneCreateDotNet, _ptr_ZoneCreateDotNet); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_ZoneCreateDotNet) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ZoneCreateDotNet := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ZoneCreateDotNet == nil {
			o.ZoneCreateDotNet = &ZoneCreateInfo_NET{}
		}
		if err := o.ZoneCreateDotNet.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ZoneCreateDotNet := func(ptr interface{}) { o.ZoneCreateDotNet = *ptr.(**ZoneCreateInfo_NET) }
	if err := w.ReadPointer(&o.ZoneCreateDotNet, _s_ZoneCreateDotNet, _ptr_ZoneCreateDotNet); err != nil {
		return err
	}
	return nil
}

// Union_ZoneList structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 27
type Union_ZoneList struct {
	// ZoneList: A pointer to a structure of type DNS_RPC_ZONE_LIST (section 2.2.5.2.3).
	// This structure is used to enumerate zones.
	ZoneList *ZoneList `idl:"name:ZoneList" json:"zone_list"`
}

func (*Union_ZoneList) is_Union() {}

func (o *Union_ZoneList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ZoneList != nil {
		_ptr_ZoneList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ZoneList != nil {
				if err := o.ZoneList.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ZoneList{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ZoneList, _ptr_ZoneList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_ZoneList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ZoneList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ZoneList == nil {
			o.ZoneList = &ZoneList{}
		}
		if err := o.ZoneList.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ZoneList := func(ptr interface{}) { o.ZoneList = *ptr.(**ZoneList) }
	if err := w.ReadPointer(&o.ZoneList, _s_ZoneList, _ptr_ZoneList); err != nil {
		return err
	}
	return nil
}

// Union_ZoneExport structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 18
type Union_ZoneExport struct {
	// ZoneExport: A pointer to a structure of type DNS_RPC_ZONE_EXPORT_INFO (section 2.2.5.2.8).
	// This structure is used to specify how a zone is exported to file.
	ZoneExport *ZoneExportInfo `idl:"name:ZoneExport" json:"zone_export"`
}

func (*Union_ZoneExport) is_Union() {}

func (o *Union_ZoneExport) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ZoneExport != nil {
		_ptr_ZoneExport := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ZoneExport != nil {
				if err := o.ZoneExport.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ZoneExportInfo{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ZoneExport, _ptr_ZoneExport); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_ZoneExport) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ZoneExport := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ZoneExport == nil {
			o.ZoneExport = &ZoneExportInfo{}
		}
		if err := o.ZoneExport.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ZoneExport := func(ptr interface{}) { o.ZoneExport = *ptr.(**ZoneExportInfo) }
	if err := w.ReadPointer(&o.ZoneExport, _s_ZoneExport, _ptr_ZoneExport); err != nil {
		return err
	}
	return nil
}

// Union_DirectoryPartition structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 29
type Union_DirectoryPartition struct {
	// DirectoryPartition: A pointer to a structure of type DNS_RPC_DP_INFO (section 2.2.7.2.1).
	// This structure specifies detailed information about a single application directory
	// partition.
	DirectoryPartition *DPInfo `idl:"name:DirectoryPartition" json:"directory_partition"`
}

func (*Union_DirectoryPartition) is_Union() {}

func (o *Union_DirectoryPartition) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.DirectoryPartition != nil {
		_ptr_DirectoryPartition := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.DirectoryPartition != nil {
				if err := o.DirectoryPartition.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&DPInfo{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.DirectoryPartition, _ptr_DirectoryPartition); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_DirectoryPartition) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DirectoryPartition := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.DirectoryPartition == nil {
			o.DirectoryPartition = &DPInfo{}
		}
		if err := o.DirectoryPartition.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DirectoryPartition := func(ptr interface{}) { o.DirectoryPartition = *ptr.(**DPInfo) }
	if err := w.ReadPointer(&o.DirectoryPartition, _s_DirectoryPartition, _ptr_DirectoryPartition); err != nil {
		return err
	}
	return nil
}

// Union_DirectoryPartitionEnum structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 28
type Union_DirectoryPartitionEnum struct {
	// DirectoryPartitionEnum: A pointer to a structure of type DNS_RPC_DP_ENUM (section
	// 2.2.7.2.3). This structure is used to specify basic information about an application
	// directory partition.
	DirectoryPartitionEnum *DPEnum `idl:"name:DirectoryPartitionEnum" json:"directory_partition_enum"`
}

func (*Union_DirectoryPartitionEnum) is_Union() {}

func (o *Union_DirectoryPartitionEnum) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.DirectoryPartitionEnum != nil {
		_ptr_DirectoryPartitionEnum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.DirectoryPartitionEnum != nil {
				if err := o.DirectoryPartitionEnum.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&DPEnum{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.DirectoryPartitionEnum, _ptr_DirectoryPartitionEnum); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_DirectoryPartitionEnum) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DirectoryPartitionEnum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.DirectoryPartitionEnum == nil {
			o.DirectoryPartitionEnum = &DPEnum{}
		}
		if err := o.DirectoryPartitionEnum.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DirectoryPartitionEnum := func(ptr interface{}) { o.DirectoryPartitionEnum = *ptr.(**DPEnum) }
	if err := w.ReadPointer(&o.DirectoryPartitionEnum, _s_DirectoryPartitionEnum, _ptr_DirectoryPartitionEnum); err != nil {
		return err
	}
	return nil
}

// Union_DirectoryPartitionList structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 30
type Union_DirectoryPartitionList struct {
	// DirectoryPartitionList: A pointer to a structure of type DNS_RPC_DP_LIST (section
	// 2.2.7.2.4). This structure is used to enumerate the Application Directory Partition
	// Table.
	DirectoryPartitionList *DPList `idl:"name:DirectoryPartitionList" json:"directory_partition_list"`
}

func (*Union_DirectoryPartitionList) is_Union() {}

func (o *Union_DirectoryPartitionList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.DirectoryPartitionList != nil {
		_ptr_DirectoryPartitionList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.DirectoryPartitionList != nil {
				if err := o.DirectoryPartitionList.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&DPList{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.DirectoryPartitionList, _ptr_DirectoryPartitionList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_DirectoryPartitionList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DirectoryPartitionList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.DirectoryPartitionList == nil {
			o.DirectoryPartitionList = &DPList{}
		}
		if err := o.DirectoryPartitionList.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DirectoryPartitionList := func(ptr interface{}) { o.DirectoryPartitionList = *ptr.(**DPList) }
	if err := w.ReadPointer(&o.DirectoryPartitionList, _s_DirectoryPartitionList, _ptr_DirectoryPartitionList); err != nil {
		return err
	}
	return nil
}

// Union_EnlistDirectoryPartition structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 31
type Union_EnlistDirectoryPartition struct {
	// EnlistDirectoryPartition: A pointer to a structure of type DNS_RPC_ENLIST_DP (section
	// 2.2.7.2.5). This structure is used to request enlistment changes for an application
	// directory partition.
	EnlistDirectoryPartition *EnlistDP `idl:"name:EnlistDirectoryPartition" json:"enlist_directory_partition"`
}

func (*Union_EnlistDirectoryPartition) is_Union() {}

func (o *Union_EnlistDirectoryPartition) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.EnlistDirectoryPartition != nil {
		_ptr_EnlistDirectoryPartition := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.EnlistDirectoryPartition != nil {
				if err := o.EnlistDirectoryPartition.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&EnlistDP{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.EnlistDirectoryPartition, _ptr_EnlistDirectoryPartition); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_EnlistDirectoryPartition) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_EnlistDirectoryPartition := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.EnlistDirectoryPartition == nil {
			o.EnlistDirectoryPartition = &EnlistDP{}
		}
		if err := o.EnlistDirectoryPartition.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_EnlistDirectoryPartition := func(ptr interface{}) { o.EnlistDirectoryPartition = *ptr.(**EnlistDP) }
	if err := w.ReadPointer(&o.EnlistDirectoryPartition, _s_EnlistDirectoryPartition, _ptr_EnlistDirectoryPartition); err != nil {
		return err
	}
	return nil
}

// Union_ZoneChangeDirectoryPartition structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 32
type Union_ZoneChangeDirectoryPartition struct {
	// ZoneChangeDirectoryPartition: A pointer to a structure of type DNS_RPC_ZONE_CHANGE_DP
	// (section 2.2.7.2.6). This structure is used to request that a DNS zone be moved from
	// one application directory partition to another.
	ZoneChangeDirectoryPartition *ZoneChangeDP `idl:"name:ZoneChangeDirectoryPartition" json:"zone_change_directory_partition"`
}

func (*Union_ZoneChangeDirectoryPartition) is_Union() {}

func (o *Union_ZoneChangeDirectoryPartition) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ZoneChangeDirectoryPartition != nil {
		_ptr_ZoneChangeDirectoryPartition := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ZoneChangeDirectoryPartition != nil {
				if err := o.ZoneChangeDirectoryPartition.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ZoneChangeDP{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ZoneChangeDirectoryPartition, _ptr_ZoneChangeDirectoryPartition); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_ZoneChangeDirectoryPartition) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ZoneChangeDirectoryPartition := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ZoneChangeDirectoryPartition == nil {
			o.ZoneChangeDirectoryPartition = &ZoneChangeDP{}
		}
		if err := o.ZoneChangeDirectoryPartition.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ZoneChangeDirectoryPartition := func(ptr interface{}) { o.ZoneChangeDirectoryPartition = *ptr.(**ZoneChangeDP) }
	if err := w.ReadPointer(&o.ZoneChangeDirectoryPartition, _s_ZoneChangeDirectoryPartition, _ptr_ZoneChangeDirectoryPartition); err != nil {
		return err
	}
	return nil
}

// Union_EnumZonesFilter structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 33
type Union_EnumZonesFilter struct {
	// EnumZonesFilter: A pointer to a structure of type DNS_RPC_ENUM_ZONES_FILTER (section
	// 2.2.5.2.9). This structure is used to filter DNS zones during enumeration.
	EnumZonesFilter *EnumZonesFilter `idl:"name:EnumZonesFilter" json:"enum_zones_filter"`
}

func (*Union_EnumZonesFilter) is_Union() {}

func (o *Union_EnumZonesFilter) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.EnumZonesFilter != nil {
		_ptr_EnumZonesFilter := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.EnumZonesFilter != nil {
				if err := o.EnumZonesFilter.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&EnumZonesFilter{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.EnumZonesFilter, _ptr_EnumZonesFilter); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_EnumZonesFilter) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_EnumZonesFilter := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.EnumZonesFilter == nil {
			o.EnumZonesFilter = &EnumZonesFilter{}
		}
		if err := o.EnumZonesFilter.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_EnumZonesFilter := func(ptr interface{}) { o.EnumZonesFilter = *ptr.(**EnumZonesFilter) }
	if err := w.ReadPointer(&o.EnumZonesFilter, _s_EnumZonesFilter, _ptr_EnumZonesFilter); err != nil {
		return err
	}
	return nil
}

// Union_AddrArray structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 34
type Union_AddrArray struct {
	// AddrArray: A pointer to a structure of type DNS_ADDR_ARRAY (section 2.2.3.2.3). This
	// structure is used to specify a list of IPv4 or IPv6 addresses.
	AddrArray *AddrArray `idl:"name:AddrArray" json:"addr_array"`
}

func (*Union_AddrArray) is_Union() {}

func (o *Union_AddrArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.AddrArray != nil {
		_ptr_AddrArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.AddrArray != nil {
				if err := o.AddrArray.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AddrArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.AddrArray, _ptr_AddrArray); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_AddrArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_AddrArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.AddrArray == nil {
			o.AddrArray = &AddrArray{}
		}
		if err := o.AddrArray.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_AddrArray := func(ptr interface{}) { o.AddrArray = *ptr.(**AddrArray) }
	if err := w.ReadPointer(&o.AddrArray, _s_AddrArray, _ptr_AddrArray); err != nil {
		return err
	}
	return nil
}

// Union_ServerInfo structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 35
type Union_ServerInfo struct {
	// ServerInfo: A pointer to a structure of type DNS_RPC_SERVER_INFO (section 2.2.4.2.2).
	// This structure is used to specify general DNS server state and configuration.
	ServerInfo *ServerInfo `idl:"name:ServerInfo" json:"server_info"`
}

func (*Union_ServerInfo) is_Union() {}

func (o *Union_ServerInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ServerInfo != nil {
		_ptr_ServerInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ServerInfo != nil {
				if err := o.ServerInfo.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ServerInfo{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ServerInfo, _ptr_ServerInfo); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_ServerInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ServerInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ServerInfo == nil {
			o.ServerInfo = &ServerInfo{}
		}
		if err := o.ServerInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ServerInfo := func(ptr interface{}) { o.ServerInfo = *ptr.(**ServerInfo) }
	if err := w.ReadPointer(&o.ServerInfo, _s_ServerInfo, _ptr_ServerInfo); err != nil {
		return err
	}
	return nil
}

// Union_ZoneCreate structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 40
type Union_ZoneCreate struct {
	// ZoneCreate: A pointer to a structure of type DNS_RPC_ZONE_CREATE_INFO (section 2.2.5.2.7).
	// This structure is used to specify the parameters required when creating a new DNS
	// zone.
	ZoneCreate *ZoneCreateInfo `idl:"name:ZoneCreate" json:"zone_create"`
}

func (*Union_ZoneCreate) is_Union() {}

func (o *Union_ZoneCreate) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ZoneCreate != nil {
		_ptr_ZoneCreate := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ZoneCreate != nil {
				if err := o.ZoneCreate.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ZoneCreateInfo{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ZoneCreate, _ptr_ZoneCreate); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_ZoneCreate) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ZoneCreate := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ZoneCreate == nil {
			o.ZoneCreate = &ZoneCreateInfo{}
		}
		if err := o.ZoneCreate.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ZoneCreate := func(ptr interface{}) { o.ZoneCreate = *ptr.(**ZoneCreateInfo) }
	if err := w.ReadPointer(&o.ZoneCreate, _s_ZoneCreate, _ptr_ZoneCreate); err != nil {
		return err
	}
	return nil
}

// Union_Forwarders structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 37
type Union_Forwarders struct {
	// Forwarders: A pointer to a structure of type DNS_RPC_FORWARDERS (section 2.2.5.2.10).
	// This structure specifies the set of DNS servers this DNS server will forward unresolved
	// queries to.
	Forwarders *Forwarders `idl:"name:Forwarders" json:"forwarders"`
}

func (*Union_Forwarders) is_Union() {}

func (o *Union_Forwarders) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Forwarders != nil {
		_ptr_Forwarders := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Forwarders != nil {
				if err := o.Forwarders.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Forwarders{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Forwarders, _ptr_Forwarders); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_Forwarders) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_Forwarders := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Forwarders == nil {
			o.Forwarders = &Forwarders{}
		}
		if err := o.Forwarders.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Forwarders := func(ptr interface{}) { o.Forwarders = *ptr.(**Forwarders) }
	if err := w.ReadPointer(&o.Forwarders, _s_Forwarders, _ptr_Forwarders); err != nil {
		return err
	}
	return nil
}

// Union_Secondaries structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 38
type Union_Secondaries struct {
	// Secondaries: A pointer to a structure of type DNS_RPC_ZONE_SECONDARIES (section 2.2.5.2.5).
	// This structure is used to specify information about the secondary servers for a primary
	// DNS zone.
	Secondaries *ZoneSecondaries `idl:"name:Secondaries" json:"secondaries"`
}

func (*Union_Secondaries) is_Union() {}

func (o *Union_Secondaries) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Secondaries != nil {
		_ptr_Secondaries := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Secondaries != nil {
				if err := o.Secondaries.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ZoneSecondaries{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Secondaries, _ptr_Secondaries); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_Secondaries) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_Secondaries := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Secondaries == nil {
			o.Secondaries = &ZoneSecondaries{}
		}
		if err := o.Secondaries.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Secondaries := func(ptr interface{}) { o.Secondaries = *ptr.(**ZoneSecondaries) }
	if err := w.ReadPointer(&o.Secondaries, _s_Secondaries, _ptr_Secondaries); err != nil {
		return err
	}
	return nil
}

// Union_IPValidate structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 41
type Union_IPValidate struct {
	// IpValidate: A pointer to a structure of type DNS_RPC_IP_VALIDATE (section 2.2.3.2.4).
	// This structure is used to request IP validation and to return the results of IP validation.
	IPValidate *IPValidate `idl:"name:IpValidate" json:"ip_validate"`
}

func (*Union_IPValidate) is_Union() {}

func (o *Union_IPValidate) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.IPValidate != nil {
		_ptr_IpValidate := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.IPValidate != nil {
				if err := o.IPValidate.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPValidate{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.IPValidate, _ptr_IpValidate); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_IPValidate) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_IpValidate := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.IPValidate == nil {
			o.IPValidate = &IPValidate{}
		}
		if err := o.IPValidate.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_IpValidate := func(ptr interface{}) { o.IPValidate = *ptr.(**IPValidate) }
	if err := w.ReadPointer(&o.IPValidate, _s_IpValidate, _ptr_IpValidate); err != nil {
		return err
	}
	return nil
}

// Union_ZoneInfo structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 36
type Union_ZoneInfo struct {
	// ZoneInfo: A pointer to a structure of type DNS_RPC_ZONE_INFO (section 2.2.5.2.4).
	// This structure is used to specify detailed DNS zone information.
	ZoneInfo *ZoneInfo `idl:"name:ZoneInfo" json:"zone_info"`
}

func (*Union_ZoneInfo) is_Union() {}

func (o *Union_ZoneInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ZoneInfo != nil {
		_ptr_ZoneInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ZoneInfo != nil {
				if err := o.ZoneInfo.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ZoneInfo{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ZoneInfo, _ptr_ZoneInfo); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_ZoneInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ZoneInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ZoneInfo == nil {
			o.ZoneInfo = &ZoneInfo{}
		}
		if err := o.ZoneInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ZoneInfo := func(ptr interface{}) { o.ZoneInfo = *ptr.(**ZoneInfo) }
	if err := w.ReadPointer(&o.ZoneInfo, _s_ZoneInfo, _ptr_ZoneInfo); err != nil {
		return err
	}
	return nil
}

// Union_AutoConfigure structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 42
type Union_AutoConfigure struct {
	// AutoConfigure: A pointer to a structure of type DNS_RPC_AUTOCONFIGURE (section 2.2.8.2.1).
	// This structure is used to request DNS server autoconfiguration.
	AutoConfigure *AutoConfigure `idl:"name:AutoConfigure" json:"auto_configure"`
}

func (*Union_AutoConfigure) is_Union() {}

func (o *Union_AutoConfigure) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.AutoConfigure != nil {
		_ptr_AutoConfigure := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.AutoConfigure != nil {
				if err := o.AutoConfigure.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AutoConfigure{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.AutoConfigure, _ptr_AutoConfigure); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_AutoConfigure) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_AutoConfigure := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.AutoConfigure == nil {
			o.AutoConfigure = &AutoConfigure{}
		}
		if err := o.AutoConfigure.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_AutoConfigure := func(ptr interface{}) { o.AutoConfigure = *ptr.(**AutoConfigure) }
	if err := w.ReadPointer(&o.AutoConfigure, _s_AutoConfigure, _ptr_AutoConfigure); err != nil {
		return err
	}
	return nil
}

// Union_UTF8StringList structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 43
type Union_UTF8StringList struct {
	// Utf8StringList: A pointer to a structure of type DNS_RPC_UTF8_STRING_LIST (section
	// 2.2.1.2.3). This structure is used to represent a list of UTF-8 [RFC3629] strings.
	UTF8StringList *UTF8StringList `idl:"name:Utf8StringList" json:"utf8_string_list"`
}

func (*Union_UTF8StringList) is_Union() {}

func (o *Union_UTF8StringList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.UTF8StringList != nil {
		_ptr_Utf8StringList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.UTF8StringList != nil {
				if err := o.UTF8StringList.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&UTF8StringList{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.UTF8StringList, _ptr_Utf8StringList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_UTF8StringList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_Utf8StringList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.UTF8StringList == nil {
			o.UTF8StringList = &UTF8StringList{}
		}
		if err := o.UTF8StringList.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Utf8StringList := func(ptr interface{}) { o.UTF8StringList = *ptr.(**UTF8StringList) }
	if err := w.ReadPointer(&o.UTF8StringList, _s_Utf8StringList, _ptr_Utf8StringList); err != nil {
		return err
	}
	return nil
}

// Union_UnicodeStringList structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 44
type Union_UnicodeStringList struct {
	// UnicodeStringList: A pointer to a structure of type DNS_RPC_UNICODE_STRING_LIST (section
	// 2.2.1.2.4). This structure is used to represent a list of Unicode strings.
	UnicodeStringList *UnicodeStringList `idl:"name:UnicodeStringList" json:"unicode_string_list"`
}

func (*Union_UnicodeStringList) is_Union() {}

func (o *Union_UnicodeStringList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.UnicodeStringList != nil {
		_ptr_UnicodeStringList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.UnicodeStringList != nil {
				if err := o.UnicodeStringList.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&UnicodeStringList{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.UnicodeStringList, _ptr_UnicodeStringList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_UnicodeStringList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_UnicodeStringList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.UnicodeStringList == nil {
			o.UnicodeStringList = &UnicodeStringList{}
		}
		if err := o.UnicodeStringList.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_UnicodeStringList := func(ptr interface{}) { o.UnicodeStringList = *ptr.(**UnicodeStringList) }
	if err := w.ReadPointer(&o.UnicodeStringList, _s_UnicodeStringList, _ptr_UnicodeStringList); err != nil {
		return err
	}
	return nil
}

// Union_SKD structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 45
type Union_SKD struct {
	// Skd: A pointer to a structure of type DNS_RPC_SKD (section 2.2.6.2.1). This structure
	// is used to specify detailed signing key descriptor (SKD) information.
	SKD *SKD `idl:"name:Skd" json:"skd"`
}

func (*Union_SKD) is_Union() {}

func (o *Union_SKD) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.SKD != nil {
		_ptr_Skd := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SKD != nil {
				if err := o.SKD.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SKD{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SKD, _ptr_Skd); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_SKD) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_Skd := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SKD == nil {
			o.SKD = &SKD{}
		}
		if err := o.SKD.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Skd := func(ptr interface{}) { o.SKD = *ptr.(**SKD) }
	if err := w.ReadPointer(&o.SKD, _s_Skd, _ptr_Skd); err != nil {
		return err
	}
	return nil
}

// Union_SKDList structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 46
type Union_SKDList struct {
	// SkdList: A pointer to a structure of type DNS_RPC_SKD_LIST (section 2.2.6.2.2). This
	// structure is used to enumerate signing key descriptors.
	SKDList *SKDList `idl:"name:SkdList" json:"skd_list"`
}

func (*Union_SKDList) is_Union() {}

func (o *Union_SKDList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.SKDList != nil {
		_ptr_SkdList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SKDList != nil {
				if err := o.SKDList.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SKDList{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SKDList, _ptr_SkdList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_SKDList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_SkdList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SKDList == nil {
			o.SKDList = &SKDList{}
		}
		if err := o.SKDList.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_SkdList := func(ptr interface{}) { o.SKDList = *ptr.(**SKDList) }
	if err := w.ReadPointer(&o.SKDList, _s_SkdList, _ptr_SkdList); err != nil {
		return err
	}
	return nil
}

// Union_SKDState structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 47
type Union_SKDState struct {
	// SkdState: A pointer to a structure of type DNS_RPC_SKD_STATE (section 2.2.6.2.3).
	// This structure is used to specify detailed signing key descriptor state information.
	SKDState *SKDState `idl:"name:SkdState" json:"skd_state"`
}

func (*Union_SKDState) is_Union() {}

func (o *Union_SKDState) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.SKDState != nil {
		_ptr_SkdState := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SKDState != nil {
				if err := o.SKDState.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SKDState{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SKDState, _ptr_SkdState); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_SKDState) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_SkdState := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SKDState == nil {
			o.SKDState = &SKDState{}
		}
		if err := o.SKDState.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_SkdState := func(ptr interface{}) { o.SKDState = *ptr.(**SKDState) }
	if err := w.ReadPointer(&o.SKDState, _s_SkdState, _ptr_SkdState); err != nil {
		return err
	}
	return nil
}

// Union_SigningValidationError structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 48
type Union_SigningValidationError struct {
	// SigningValidationError: A pointer to a structure of type DNS_RPC_SIGNING_VALIDATION_ERROR
	// (section 2.2.6.2.8). This structure is used to specify signing key descriptor validation
	// error information.
	SigningValidationError *SigningValidationError `idl:"name:SigningValidationError" json:"signing_validation_error"`
}

func (*Union_SigningValidationError) is_Union() {}

func (o *Union_SigningValidationError) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.SigningValidationError != nil {
		_ptr_SigningValidationError := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SigningValidationError != nil {
				if err := o.SigningValidationError.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SigningValidationError{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SigningValidationError, _ptr_SigningValidationError); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_SigningValidationError) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_SigningValidationError := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SigningValidationError == nil {
			o.SigningValidationError = &SigningValidationError{}
		}
		if err := o.SigningValidationError.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_SigningValidationError := func(ptr interface{}) { o.SigningValidationError = *ptr.(**SigningValidationError) }
	if err := w.ReadPointer(&o.SigningValidationError, _s_SigningValidationError, _ptr_SigningValidationError); err != nil {
		return err
	}
	return nil
}

// Union_TrustPointList structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 49
type Union_TrustPointList struct {
	// TrustPointList: A pointer to a structure of type DNS_RPC_TRUST_POINT_LIST (section
	// 2.2.6.2.5). This structure is used to enumerate trust points.
	TrustPointList *TrustPointList `idl:"name:TrustPointList" json:"trust_point_list"`
}

func (*Union_TrustPointList) is_Union() {}

func (o *Union_TrustPointList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TrustPointList != nil {
		_ptr_TrustPointList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.TrustPointList != nil {
				if err := o.TrustPointList.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&TrustPointList{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.TrustPointList, _ptr_TrustPointList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_TrustPointList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_TrustPointList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.TrustPointList == nil {
			o.TrustPointList = &TrustPointList{}
		}
		if err := o.TrustPointList.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_TrustPointList := func(ptr interface{}) { o.TrustPointList = *ptr.(**TrustPointList) }
	if err := w.ReadPointer(&o.TrustPointList, _s_TrustPointList, _ptr_TrustPointList); err != nil {
		return err
	}
	return nil
}

// Union_TrustAnchorList structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 50
type Union_TrustAnchorList struct {
	// TrustAnchorList: A pointer to a structure of type DNS_RPC_TRUST_ANCHOR_LIST (section
	// 2.2.6.2.7). This structure is used to enumerate trust anchors.
	TrustAnchorList *TrustAnchorList `idl:"name:TrustAnchorList" json:"trust_anchor_list"`
}

func (*Union_TrustAnchorList) is_Union() {}

func (o *Union_TrustAnchorList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TrustAnchorList != nil {
		_ptr_TrustAnchorList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.TrustAnchorList != nil {
				if err := o.TrustAnchorList.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&TrustAnchorList{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.TrustAnchorList, _ptr_TrustAnchorList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_TrustAnchorList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_TrustAnchorList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.TrustAnchorList == nil {
			o.TrustAnchorList = &TrustAnchorList{}
		}
		if err := o.TrustAnchorList.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_TrustAnchorList := func(ptr interface{}) { o.TrustAnchorList = *ptr.(**TrustAnchorList) }
	if err := w.ReadPointer(&o.TrustAnchorList, _s_TrustAnchorList, _ptr_TrustAnchorList); err != nil {
		return err
	}
	return nil
}

// Union_ZoneDNSSecuritySettings structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 51
type Union_ZoneDNSSecuritySettings struct {
	// ZoneDnsSecSettings: A pointer to a structure of type DNS_RPC_ZONE_DNSSEC_SETTINGS
	// (section 2.2.6.2.9). This structure is used to list the DNSSEC settings of a zone.
	ZoneDNSSecuritySettings *ZoneDNSSECSettings `idl:"name:ZoneDnsSecSettings" json:"zone_dns_security_settings"`
}

func (*Union_ZoneDNSSecuritySettings) is_Union() {}

func (o *Union_ZoneDNSSecuritySettings) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ZoneDNSSecuritySettings != nil {
		_ptr_ZoneDnsSecSettings := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ZoneDNSSecuritySettings != nil {
				if err := o.ZoneDNSSecuritySettings.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ZoneDNSSECSettings{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ZoneDNSSecuritySettings, _ptr_ZoneDnsSecSettings); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_ZoneDNSSecuritySettings) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ZoneDnsSecSettings := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ZoneDNSSecuritySettings == nil {
			o.ZoneDNSSecuritySettings = &ZoneDNSSECSettings{}
		}
		if err := o.ZoneDNSSecuritySettings.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ZoneDnsSecSettings := func(ptr interface{}) { o.ZoneDNSSecuritySettings = *ptr.(**ZoneDNSSECSettings) }
	if err := w.ReadPointer(&o.ZoneDNSSecuritySettings, _s_ZoneDnsSecSettings, _ptr_ZoneDnsSecSettings); err != nil {
		return err
	}
	return nil
}

// Union_ZoneScopeList structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 52
type Union_ZoneScopeList struct {
	// ZoneScopeList: A pointer to a structure of type DNS_RPC_ENUM_ZONE_SCOPE_LIST (section
	// 2.2.13.1.1). This structure is used to enumerate zone scopes in a specified zone
	// or cache scopes in a cache zone.
	ZoneScopeList *EnumZoneScopeList `idl:"name:ZoneScopeList" json:"zone_scope_list"`
}

func (*Union_ZoneScopeList) is_Union() {}

func (o *Union_ZoneScopeList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ZoneScopeList != nil {
		_ptr_ZoneScopeList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ZoneScopeList != nil {
				if err := o.ZoneScopeList.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&EnumZoneScopeList{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ZoneScopeList, _ptr_ZoneScopeList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_ZoneScopeList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ZoneScopeList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ZoneScopeList == nil {
			o.ZoneScopeList = &EnumZoneScopeList{}
		}
		if err := o.ZoneScopeList.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ZoneScopeList := func(ptr interface{}) { o.ZoneScopeList = *ptr.(**EnumZoneScopeList) }
	if err := w.ReadPointer(&o.ZoneScopeList, _s_ZoneScopeList, _ptr_ZoneScopeList); err != nil {
		return err
	}
	return nil
}

// Union_ZoneStats structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 53
type Union_ZoneStats struct {
	// ZoneStats: A pointer to a structure of type DNS_RPC_ZONE_STATS_V1 (section 2.2.12.2.5).
	// This structure is used to enumerate the zone statistics.
	ZoneStats *ZoneStatsV1 `idl:"name:ZoneStats" json:"zone_stats"`
}

func (*Union_ZoneStats) is_Union() {}

func (o *Union_ZoneStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ZoneStats != nil {
		_ptr_ZoneStats := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ZoneStats != nil {
				if err := o.ZoneStats.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ZoneStatsV1{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ZoneStats, _ptr_ZoneStats); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_ZoneStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ZoneStats := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ZoneStats == nil {
			o.ZoneStats = &ZoneStatsV1{}
		}
		if err := o.ZoneStats.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ZoneStats := func(ptr interface{}) { o.ZoneStats = *ptr.(**ZoneStatsV1) }
	if err := w.ReadPointer(&o.ZoneStats, _s_ZoneStats, _ptr_ZoneStats); err != nil {
		return err
	}
	return nil
}

// Union_ScopeCreate structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 54
type Union_ScopeCreate struct {
	// ScopeCreate: A pointer to a structure of type DNS_RPC_ZONE_SCOPE_CREATE_INFO_V1 (section
	// 2.2.13.1.2.1). This structure is used to create a zone scope or cache scope.
	ScopeCreate *ZoneScopeCreateInfoV1 `idl:"name:ScopeCreate" json:"scope_create"`
}

func (*Union_ScopeCreate) is_Union() {}

func (o *Union_ScopeCreate) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ScopeCreate != nil {
		_ptr_ScopeCreate := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ScopeCreate != nil {
				if err := o.ScopeCreate.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ZoneScopeCreateInfoV1{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ScopeCreate, _ptr_ScopeCreate); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_ScopeCreate) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ScopeCreate := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ScopeCreate == nil {
			o.ScopeCreate = &ZoneScopeCreateInfoV1{}
		}
		if err := o.ScopeCreate.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ScopeCreate := func(ptr interface{}) { o.ScopeCreate = *ptr.(**ZoneScopeCreateInfoV1) }
	if err := w.ReadPointer(&o.ScopeCreate, _s_ScopeCreate, _ptr_ScopeCreate); err != nil {
		return err
	}
	return nil
}

// Union_ScopeInfo structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 55
type Union_ScopeInfo struct {
	// ScopeInfo: A pointer to a structure of type DNS_RPC_ZONE_SCOPE_INFO_V1 (section 2.2.13.1.3.1).
	// This structure is used to specify detailed DNS zone scope or cache scope information.
	ScopeInfo *ZoneScopeInfoV1 `idl:"name:ScopeInfo" json:"scope_info"`
}

func (*Union_ScopeInfo) is_Union() {}

func (o *Union_ScopeInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ScopeInfo != nil {
		_ptr_ScopeInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ScopeInfo != nil {
				if err := o.ScopeInfo.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ZoneScopeInfoV1{}).MarshalNDR(ctx, w); err != nil {
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
func (o *Union_ScopeInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ScopeInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ScopeInfo == nil {
			o.ScopeInfo = &ZoneScopeInfoV1{}
		}
		if err := o.ScopeInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ScopeInfo := func(ptr interface{}) { o.ScopeInfo = *ptr.(**ZoneScopeInfoV1) }
	if err := w.ReadPointer(&o.ScopeInfo, _s_ScopeInfo, _ptr_ScopeInfo); err != nil {
		return err
	}
	return nil
}

// Union_ScopeList structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 56
type Union_ScopeList struct {
	// ScopeList: A pointer to a structure of type DNS_RPC_ENUM_SCOPE_LIST (section 2.2.14.1.1).
	// This structure is used to retrieve all the server scopes configured on the DNS server.
	ScopeList *EnumScopeList `idl:"name:ScopeList" json:"scope_list"`
}

func (*Union_ScopeList) is_Union() {}

func (o *Union_ScopeList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ScopeList != nil {
		_ptr_ScopeList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ScopeList != nil {
				if err := o.ScopeList.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&EnumScopeList{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ScopeList, _ptr_ScopeList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_ScopeList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ScopeList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ScopeList == nil {
			o.ScopeList = &EnumScopeList{}
		}
		if err := o.ScopeList.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ScopeList := func(ptr interface{}) { o.ScopeList = *ptr.(**EnumScopeList) }
	if err := w.ReadPointer(&o.ScopeList, _s_ScopeList, _ptr_ScopeList); err != nil {
		return err
	}
	return nil
}

// Union_SubnetList structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 57
type Union_SubnetList struct {
	// SubnetList: A pointer to a structure of type DNS_RPC_CLIENT_SUBNET_RECORD (section
	// 2.2.15.2.1). This structure is used to retrieve details of a client subnet record
	// in the DNS server.
	SubnetList *ClientSubnetRecord `idl:"name:SubnetList" json:"subnet_list"`
}

func (*Union_SubnetList) is_Union() {}

func (o *Union_SubnetList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.SubnetList != nil {
		_ptr_SubnetList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SubnetList != nil {
				if err := o.SubnetList.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ClientSubnetRecord{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SubnetList, _ptr_SubnetList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_SubnetList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_SubnetList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SubnetList == nil {
			o.SubnetList = &ClientSubnetRecord{}
		}
		if err := o.SubnetList.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_SubnetList := func(ptr interface{}) { o.SubnetList = *ptr.(**ClientSubnetRecord) }
	if err := w.ReadPointer(&o.SubnetList, _s_SubnetList, _ptr_SubnetList); err != nil {
		return err
	}
	return nil
}

// Union_Policy structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 58
type Union_Policy struct {
	// pPolicy: A pointer to a structure of type DNS_RPC_POLICY (section 2.2.15.2.5). This
	// structure is used to retrieve details of a DNS Policy in the DNS server.
	Policy *Policy `idl:"name:pPolicy" json:"policy"`
}

func (*Union_Policy) is_Union() {}

func (o *Union_Policy) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Policy != nil {
		_ptr_pPolicy := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Policy != nil {
				if err := o.Policy.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Policy{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Policy, _ptr_pPolicy); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_Policy) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_pPolicy := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Policy == nil {
			o.Policy = &Policy{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pPolicy := func(ptr interface{}) { o.Policy = *ptr.(**Policy) }
	if err := w.ReadPointer(&o.Policy, _s_pPolicy, _ptr_pPolicy); err != nil {
		return err
	}
	return nil
}

// Union_PolicyName structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 59
type Union_PolicyName struct {
	// pPolicyName: A pointer to a structure of type DNS_RPC_POLICY_NAME (section 2.2.15.2.6).
	// This structure is used while enumerating DNS Policies in a DNS server per level.
	PolicyName *PolicyName `idl:"name:pPolicyName" json:"policy_name"`
}

func (*Union_PolicyName) is_Union() {}

func (o *Union_PolicyName) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PolicyName != nil {
		_ptr_pPolicyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.PolicyName != nil {
				if err := o.PolicyName.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&PolicyName{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.PolicyName, _ptr_pPolicyName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_PolicyName) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_pPolicyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.PolicyName == nil {
			o.PolicyName = &PolicyName{}
		}
		if err := o.PolicyName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pPolicyName := func(ptr interface{}) { o.PolicyName = *ptr.(**PolicyName) }
	if err := w.ReadPointer(&o.PolicyName, _s_pPolicyName, _ptr_pPolicyName); err != nil {
		return err
	}
	return nil
}

// Union_PolicyList structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 60
type Union_PolicyList struct {
	// pPolicyList: A pointer to a structure of type DNS_RPC_ENUMERATE_POLICY_LIST (section
	// 2.2.15.2.7). This structure contains a list of DNS_RPC_POLICY_NAME structures.
	PolicyList *EnumeratePolicyList `idl:"name:pPolicyList" json:"policy_list"`
}

func (*Union_PolicyList) is_Union() {}

func (o *Union_PolicyList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PolicyList != nil {
		_ptr_pPolicyList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.PolicyList != nil {
				if err := o.PolicyList.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&EnumeratePolicyList{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.PolicyList, _ptr_pPolicyList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_PolicyList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_pPolicyList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.PolicyList == nil {
			o.PolicyList = &EnumeratePolicyList{}
		}
		if err := o.PolicyList.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pPolicyList := func(ptr interface{}) { o.PolicyList = *ptr.(**EnumeratePolicyList) }
	if err := w.ReadPointer(&o.PolicyList, _s_pPolicyList, _ptr_pPolicyList); err != nil {
		return err
	}
	return nil
}

// Union_RRLParams structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 61
type Union_RRLParams struct {
	// pRRLParams: A pointer to a structure of type DNS_RPC_RRL_PARAMS (section 2.2.16.2.1).
	// This structure contains the configuration parameters for Response Rate Limiting (RRL).
	RRLParams *RRLParams `idl:"name:pRRLParams" json:"rrl_params"`
}

func (*Union_RRLParams) is_Union() {}

func (o *Union_RRLParams) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RRLParams != nil {
		_ptr_pRRLParams := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RRLParams != nil {
				if err := o.RRLParams.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RRLParams{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RRLParams, _ptr_pRRLParams); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_RRLParams) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_pRRLParams := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RRLParams == nil {
			o.RRLParams = &RRLParams{}
		}
		if err := o.RRLParams.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pRRLParams := func(ptr interface{}) { o.RRLParams = *ptr.(**RRLParams) }
	if err := w.ReadPointer(&o.RRLParams, _s_pRRLParams, _ptr_pRRLParams); err != nil {
		return err
	}
	return nil
}

// Union_VirtualizationInstance structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 62
type Union_VirtualizationInstance struct {
	// VirtualizationInstance: A pointer to a structure of type DNS_RPC_VIRTUALIZATION_INSTANCE
	// (section 2.2.17.1.1). This structure is used to retrieve details of a virtualization
	// instance in the DNS server.
	VirtualizationInstance *VirtualizationInstance `idl:"name:VirtualizationInstance" json:"virtualization_instance"`
}

func (*Union_VirtualizationInstance) is_Union() {}

func (o *Union_VirtualizationInstance) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.VirtualizationInstance != nil {
		_ptr_VirtualizationInstance := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.VirtualizationInstance != nil {
				if err := o.VirtualizationInstance.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&VirtualizationInstance{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.VirtualizationInstance, _ptr_VirtualizationInstance); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_VirtualizationInstance) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_VirtualizationInstance := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.VirtualizationInstance == nil {
			o.VirtualizationInstance = &VirtualizationInstance{}
		}
		if err := o.VirtualizationInstance.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_VirtualizationInstance := func(ptr interface{}) { o.VirtualizationInstance = *ptr.(**VirtualizationInstance) }
	if err := w.ReadPointer(&o.VirtualizationInstance, _s_VirtualizationInstance, _ptr_VirtualizationInstance); err != nil {
		return err
	}
	return nil
}

// Union_VirtualizationInstanceList structure represents DNSSRV_RPC_UNION RPC union arm.
//
// It has following labels: 63
type Union_VirtualizationInstanceList struct {
	// VirtualizationInstanceList: A pointer to a structure of type DNS_RPC_ENUM_VIRTUALIZATION_INSTANCE_LIST
	// (section 2.2.17.1.3). This structure is used to enumerate virtualization instances
	// in a DNS server.
	VirtualizationInstanceList *EnumVirtualizationInstanceList `idl:"name:VirtualizationInstanceList" json:"virtualization_instance_list"`
}

func (*Union_VirtualizationInstanceList) is_Union() {}

func (o *Union_VirtualizationInstanceList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.VirtualizationInstanceList != nil {
		_ptr_VirtualizationInstanceList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.VirtualizationInstanceList != nil {
				if err := o.VirtualizationInstanceList.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&EnumVirtualizationInstanceList{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.VirtualizationInstanceList, _ptr_VirtualizationInstanceList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Union_VirtualizationInstanceList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_VirtualizationInstanceList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.VirtualizationInstanceList == nil {
			o.VirtualizationInstanceList = &EnumVirtualizationInstanceList{}
		}
		if err := o.VirtualizationInstanceList.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_VirtualizationInstanceList := func(ptr interface{}) { o.VirtualizationInstanceList = *ptr.(**EnumVirtualizationInstanceList) }
	if err := w.ReadPointer(&o.VirtualizationInstanceList, _s_VirtualizationInstanceList, _ptr_VirtualizationInstanceList); err != nil {
		return err
	}
	return nil
}

// ImportOperationResult type represents ImportOpResult RPC enumeration.
//
// The ImportOpResult enumeration SHOULD<47> define the operations to be executed when
// the DNS_RPC_ZONE_DNSSEC_SETTINGS structure (section 2.2.6.2.9) is imported on a zone.
type ImportOperationResult uint16

var (
	// IMPORT_STATUS_NOOP: No change was detected in the imported signing settings or the
	// signing settings of the zone.
	ImportOperationResultStatusNoop ImportOperationResult = 0
	// IMPORT_STATUS_SIGNING_READY: The zone is marked for re-signing.
	ImportOperationResultStatusSigningReady ImportOperationResult = 1
	// IMPORT_STATUS_UNSIGNING_READY: The zone is marked for unsigning.
	ImportOperationResultStatusUnsigningReady ImportOperationResult = 2
	// IMPORT_STATUS_CHANGED: The change was detected in signing settings imported and were
	// incorporated, but no re-signing or unsigning is required.
	ImportOperationResultStatusChanged ImportOperationResult = 3
)

func (o ImportOperationResult) String() string {
	switch o {
	case ImportOperationResultStatusNoop:
		return "ImportOperationResultStatusNoop"
	case ImportOperationResultStatusSigningReady:
		return "ImportOperationResultStatusSigningReady"
	case ImportOperationResultStatusUnsigningReady:
		return "ImportOperationResultStatusUnsigningReady"
	case ImportOperationResultStatusChanged:
		return "ImportOperationResultStatusChanged"
	}
	return "Invalid"
}
