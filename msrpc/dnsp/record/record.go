// The record package implements the RECORD client protocol.
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
package record

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
	GoPackage = "dnsp/record"
)

// TypeRecord type represents DNS_RPC_TYPE_RECORD RPC enumeration.
type TypeRecord uint16

var (
	TypeRecordZERO       TypeRecord = 0
	TypeRecordA          TypeRecord = 1
	TypeRecordNS         TypeRecord = 2
	TypeRecordMD         TypeRecord = 3
	TypeRecordMF         TypeRecord = 4
	TypeRecordCNAME      TypeRecord = 5
	TypeRecordSOA        TypeRecord = 6
	TypeRecordMB         TypeRecord = 7
	TypeRecordMG         TypeRecord = 8
	TypeRecordMR         TypeRecord = 9
	TypeRecordNULL       TypeRecord = 10
	TypeRecordWKS        TypeRecord = 11
	TypeRecordPTR        TypeRecord = 12
	TypeRecordHINFO      TypeRecord = 13
	TypeRecordMINFO      TypeRecord = 14
	TypeRecordMX         TypeRecord = 15
	TypeRecordTXT        TypeRecord = 16
	TypeRecordRP         TypeRecord = 17
	TypeRecordAFSDB      TypeRecord = 18
	TypeRecordX25        TypeRecord = 19
	TypeRecordISDN       TypeRecord = 20
	TypeRecordRT         TypeRecord = 21
	TypeRecordSIG        TypeRecord = 24
	TypeRecordKEY        TypeRecord = 25
	TypeRecordAAAA       TypeRecord = 28
	TypeRecordNXT        TypeRecord = 30
	TypeRecordSRV        TypeRecord = 33
	TypeRecordATMA       TypeRecord = 34
	TypeRecordNAPTR      TypeRecord = 35
	TypeRecordDNAME      TypeRecord = 39
	TypeRecordDS         TypeRecord = 43
	TypeRecordRRSIG      TypeRecord = 46
	TypeRecordNSEC       TypeRecord = 47
	TypeRecordDNSKEY     TypeRecord = 48
	TypeRecordDHCID      TypeRecord = 49
	TypeRecordNSEC3      TypeRecord = 50
	TypeRecordNSEC3PARAM TypeRecord = 51
	TypeRecordTLSA       TypeRecord = 52
	TypeRecordWINS       TypeRecord = 65281
	TypeRecordWINSR      TypeRecord = 65282
)

func (o TypeRecord) String() string {
	switch o {
	case TypeRecordZERO:
		return "TypeRecordZERO"
	case TypeRecordA:
		return "TypeRecordA"
	case TypeRecordNS:
		return "TypeRecordNS"
	case TypeRecordMD:
		return "TypeRecordMD"
	case TypeRecordMF:
		return "TypeRecordMF"
	case TypeRecordCNAME:
		return "TypeRecordCNAME"
	case TypeRecordSOA:
		return "TypeRecordSOA"
	case TypeRecordMB:
		return "TypeRecordMB"
	case TypeRecordMG:
		return "TypeRecordMG"
	case TypeRecordMR:
		return "TypeRecordMR"
	case TypeRecordNULL:
		return "TypeRecordNULL"
	case TypeRecordWKS:
		return "TypeRecordWKS"
	case TypeRecordPTR:
		return "TypeRecordPTR"
	case TypeRecordHINFO:
		return "TypeRecordHINFO"
	case TypeRecordMINFO:
		return "TypeRecordMINFO"
	case TypeRecordMX:
		return "TypeRecordMX"
	case TypeRecordTXT:
		return "TypeRecordTXT"
	case TypeRecordRP:
		return "TypeRecordRP"
	case TypeRecordAFSDB:
		return "TypeRecordAFSDB"
	case TypeRecordX25:
		return "TypeRecordX25"
	case TypeRecordISDN:
		return "TypeRecordISDN"
	case TypeRecordRT:
		return "TypeRecordRT"
	case TypeRecordSIG:
		return "TypeRecordSIG"
	case TypeRecordKEY:
		return "TypeRecordKEY"
	case TypeRecordAAAA:
		return "TypeRecordAAAA"
	case TypeRecordNXT:
		return "TypeRecordNXT"
	case TypeRecordSRV:
		return "TypeRecordSRV"
	case TypeRecordATMA:
		return "TypeRecordATMA"
	case TypeRecordNAPTR:
		return "TypeRecordNAPTR"
	case TypeRecordDNAME:
		return "TypeRecordDNAME"
	case TypeRecordDS:
		return "TypeRecordDS"
	case TypeRecordRRSIG:
		return "TypeRecordRRSIG"
	case TypeRecordNSEC:
		return "TypeRecordNSEC"
	case TypeRecordDNSKEY:
		return "TypeRecordDNSKEY"
	case TypeRecordDHCID:
		return "TypeRecordDHCID"
	case TypeRecordNSEC3:
		return "TypeRecordNSEC3"
	case TypeRecordNSEC3PARAM:
		return "TypeRecordNSEC3PARAM"
	case TypeRecordTLSA:
		return "TypeRecordTLSA"
	case TypeRecordWINS:
		return "TypeRecordWINS"
	case TypeRecordWINSR:
		return "TypeRecordWINSR"
	}
	return "Invalid"
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
	Type TypeRecord `idl:"name:wType" json:"type"`
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
	Buffer []byte         `idl:"name:Buffer;size_is:(wDataLength)" json:"buffer"`
	Record *Record_Record `idl:"name:Record;switch_is:wType" json:"record"`
}

func (o *Record) xxx_PreparePayload(ctx context.Context) error {
	if o.Buffer != nil && o.DataLength == 0 {
		o.DataLength = uint16(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
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
	if err := w.WriteEnum(uint16(o.Type)); err != nil {
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
	if err := w.ReadEnum((*uint16)(&o.Type)); err != nil {
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
	_layout_Record := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Record == nil {
			o.Record = &Record_Record{}
		}
		_swRecord := uint16(o.Type)
		if err := o.Record.UnmarshalUnionNDR(ctx, w, _swRecord); err != nil {
			return err
		}
		return nil
	})
	if len(o.Buffer) > 0 {
		if err := w.WithBytes(o.Buffer).Unmarshal(ctx, _layout_Record); err != nil {
			return err
		}
	}
	return nil
}

// Record_Record structure represents DNS_RPC_RECORD union anonymous member.
//
// The DNS_RPC_RECORD structure is used to specify a single DNS record's parameters
// and data. This structure is returned by the DNS server in response to an R_DnssrvEnumRecords2
// (section 3.1.4.9) method call.
type Record_Record struct {
	// Types that are assignable to Value
	//
	// *Record_ZERO
	// *Record_A
	// *Record_NS
	// *Record_MD
	// *Record_MF
	// *Record_CNAME
	// *Record_SOA
	// *Record_MB
	// *Record_MG
	// *Record_MR
	// *Record_NULL
	// *Record_WKS
	// *Record_PTR
	// *Record_HINFO
	// *Record_MINFO
	// *Record_MX
	// *Record_TXT
	// *Record_RP
	// *Record_AFSDB
	// *Record_X25
	// *Record_ISDN
	// *Record_RT
	// *Record_SIG
	// *Record_KEY
	// *Record_AAAA
	// *Record_NXT
	// *Record_SRV
	// *Record_ATMA
	// *Record_NAPTR
	// *Record_DNAME
	// *Record_DS
	// *Record_RRSIG
	// *Record_NSEC
	// *Record_DNSKEY
	// *Record_DHCID
	// *Record_NSEC3
	// *Record_NSEC3PARAM
	// *Record_TLSA
	// *Record_WINS
	// *Record_WINSR
	// *Record_Unknown
	Value is_Record_Record `json:"value"`
}

func (o *Record_Record) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *Record_ZERO:
		if value != nil {
			return value.RecordZERO
		}
	case *Record_A:
		if value != nil {
			return value.RecordA
		}
	case *Record_NS:
		if value != nil {
			return value.RecordNS
		}
	case *Record_MD:
		if value != nil {
			return value.RecordMD
		}
	case *Record_MF:
		if value != nil {
			return value.RecordMF
		}
	case *Record_CNAME:
		if value != nil {
			return value.RecordCNAME
		}
	case *Record_SOA:
		if value != nil {
			return value.RecordSOA
		}
	case *Record_MB:
		if value != nil {
			return value.RecordMB
		}
	case *Record_MG:
		if value != nil {
			return value.RecordMG
		}
	case *Record_MR:
		if value != nil {
			return value.RecordMR
		}
	case *Record_NULL:
		if value != nil {
			return value.RecordNULL
		}
	case *Record_WKS:
		if value != nil {
			return value.RecordWKS
		}
	case *Record_PTR:
		if value != nil {
			return value.RecordPTR
		}
	case *Record_HINFO:
		if value != nil {
			return value.RecordHINFO
		}
	case *Record_MINFO:
		if value != nil {
			return value.RecordMINFO
		}
	case *Record_MX:
		if value != nil {
			return value.RecordMX
		}
	case *Record_TXT:
		if value != nil {
			return value.RecordTXT
		}
	case *Record_RP:
		if value != nil {
			return value.RecordRP
		}
	case *Record_AFSDB:
		if value != nil {
			return value.RecordAFSDB
		}
	case *Record_X25:
		if value != nil {
			return value.RecordX25
		}
	case *Record_ISDN:
		if value != nil {
			return value.RecordISDN
		}
	case *Record_RT:
		if value != nil {
			return value.RecordRT
		}
	case *Record_SIG:
		if value != nil {
			return value.RecordSIG
		}
	case *Record_KEY:
		if value != nil {
			return value.RecordKEY
		}
	case *Record_AAAA:
		if value != nil {
			return value.RecordAAAA
		}
	case *Record_NXT:
		if value != nil {
			return value.RecordNXT
		}
	case *Record_SRV:
		if value != nil {
			return value.RecordSRV
		}
	case *Record_ATMA:
		if value != nil {
			return value.RecordATMA
		}
	case *Record_NAPTR:
		if value != nil {
			return value.RecordNAPTR
		}
	case *Record_DNAME:
		if value != nil {
			return value.RecordDNAME
		}
	case *Record_DS:
		if value != nil {
			return value.RecordDS
		}
	case *Record_RRSIG:
		if value != nil {
			return value.RecordRRSIG
		}
	case *Record_NSEC:
		if value != nil {
			return value.RecordNSEC
		}
	case *Record_DNSKEY:
		if value != nil {
			return value.RecordDNSKEY
		}
	case *Record_DHCID:
		if value != nil {
			return value.RecordDHCID
		}
	case *Record_NSEC3:
		if value != nil {
			return value.RecordNSEC3
		}
	case *Record_NSEC3PARAM:
		if value != nil {
			return value.RecordNSEC3PARAM
		}
	case *Record_TLSA:
		if value != nil {
			return value.RecordTLSA
		}
	case *Record_WINS:
		if value != nil {
			return value.RecordWINS
		}
	case *Record_WINSR:
		if value != nil {
			return value.RecordWINSR
		}
	case *Record_Unknown:
		if value != nil {
			return value.RecordUnknown
		}
	}
	return nil
}

type is_Record_Record interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_Record_Record()
}

func (o *Record_Record) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *Record_ZERO:
		return uint16(0)
	case *Record_A:
		return uint16(1)
	case *Record_NS:
		return uint16(2)
	case *Record_MD:
		return uint16(3)
	case *Record_MF:
		return uint16(4)
	case *Record_CNAME:
		return uint16(5)
	case *Record_SOA:
		return uint16(6)
	case *Record_MB:
		return uint16(7)
	case *Record_MG:
		return uint16(8)
	case *Record_MR:
		return uint16(9)
	case *Record_NULL:
		return uint16(10)
	case *Record_WKS:
		return uint16(11)
	case *Record_PTR:
		return uint16(12)
	case *Record_HINFO:
		return uint16(13)
	case *Record_MINFO:
		return uint16(14)
	case *Record_MX:
		return uint16(15)
	case *Record_TXT:
		return uint16(16)
	case *Record_RP:
		return uint16(17)
	case *Record_AFSDB:
		return uint16(18)
	case *Record_X25:
		return uint16(19)
	case *Record_ISDN:
		return uint16(20)
	case *Record_RT:
		return uint16(21)
	case *Record_SIG:
		return uint16(24)
	case *Record_KEY:
		return uint16(25)
	case *Record_AAAA:
		return uint16(28)
	case *Record_NXT:
		return uint16(30)
	case *Record_SRV:
		return uint16(33)
	case *Record_ATMA:
		return uint16(34)
	case *Record_NAPTR:
		return uint16(35)
	case *Record_DNAME:
		return uint16(39)
	case *Record_DS:
		return uint16(43)
	case *Record_RRSIG:
		return uint16(46)
	case *Record_NSEC:
		return uint16(47)
	case *Record_DNSKEY:
		return uint16(48)
	case *Record_DHCID:
		return uint16(49)
	case *Record_NSEC3:
		return uint16(50)
	case *Record_NSEC3PARAM:
		return uint16(51)
	case *Record_TLSA:
		return uint16(52)
	case *Record_WINS:
		return uint16(65281)
	case *Record_WINSR:
		return uint16(65282)
	}
	return uint16(0)
}
func (o *Record_Record) NDRLayout() {}

func (o *Record_Record) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteUnionAlign(7); err != nil {
		return err
	}
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(7); err != nil {
		return err
	}
	switch sw {
	case uint16(0):
		_o, _ := o.Value.(*Record_ZERO)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_ZERO{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(1):
		_o, _ := o.Value.(*Record_A)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_A{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*Record_NS)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_NS{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(3):
		_o, _ := o.Value.(*Record_MD)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_MD{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(4):
		_o, _ := o.Value.(*Record_MF)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_MF{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(5):
		_o, _ := o.Value.(*Record_CNAME)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_CNAME{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(6):
		_o, _ := o.Value.(*Record_SOA)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_SOA{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(7):
		_o, _ := o.Value.(*Record_MB)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_MB{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(8):
		_o, _ := o.Value.(*Record_MG)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_MG{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(9):
		_o, _ := o.Value.(*Record_MR)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_MR{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(10):
		_o, _ := o.Value.(*Record_NULL)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_NULL{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(11):
		_o, _ := o.Value.(*Record_WKS)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_WKS{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(12):
		_o, _ := o.Value.(*Record_PTR)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_PTR{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(13):
		_o, _ := o.Value.(*Record_HINFO)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_HINFO{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(14):
		_o, _ := o.Value.(*Record_MINFO)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_MINFO{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(15):
		_o, _ := o.Value.(*Record_MX)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_MX{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(16):
		_o, _ := o.Value.(*Record_TXT)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_TXT{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(17):
		_o, _ := o.Value.(*Record_RP)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_RP{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(18):
		_o, _ := o.Value.(*Record_AFSDB)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_AFSDB{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(19):
		_o, _ := o.Value.(*Record_X25)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_X25{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(20):
		_o, _ := o.Value.(*Record_ISDN)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_ISDN{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(21):
		_o, _ := o.Value.(*Record_RT)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_RT{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(24):
		_o, _ := o.Value.(*Record_SIG)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_SIG{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(25):
		_o, _ := o.Value.(*Record_KEY)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_KEY{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(28):
		_o, _ := o.Value.(*Record_AAAA)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_AAAA{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(30):
		_o, _ := o.Value.(*Record_NXT)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_NXT{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(33):
		_o, _ := o.Value.(*Record_SRV)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_SRV{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(34):
		_o, _ := o.Value.(*Record_ATMA)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_ATMA{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(35):
		_o, _ := o.Value.(*Record_NAPTR)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_NAPTR{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(39):
		_o, _ := o.Value.(*Record_DNAME)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_DNAME{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(43):
		_o, _ := o.Value.(*Record_DS)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_DS{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(46):
		_o, _ := o.Value.(*Record_RRSIG)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_RRSIG{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(47):
		_o, _ := o.Value.(*Record_NSEC)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_NSEC{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(48):
		_o, _ := o.Value.(*Record_DNSKEY)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_DNSKEY{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(49):
		_o, _ := o.Value.(*Record_DHCID)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_DHCID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(50):
		_o, _ := o.Value.(*Record_NSEC3)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_NSEC3{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(51):
		_o, _ := o.Value.(*Record_NSEC3PARAM)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_NSEC3PARAM{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(52):
		_o, _ := o.Value.(*Record_TLSA)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_TLSA{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(65281):
		_o, _ := o.Value.(*Record_WINS)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_WINS{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(65282):
		_o, _ := o.Value.(*Record_WINSR)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_WINSR{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		_o, _ := o.Value.(*Record_Unknown)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Record_Unknown{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *Record_Record) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadUnionAlign(7); err != nil {
		return err
	}
	if err := w.ReadSwitch((*uint16)(&sw)); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(7); err != nil {
		return err
	}
	switch sw {
	case uint16(0):
		o.Value = &Record_ZERO{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(1):
		o.Value = &Record_A{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &Record_NS{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(3):
		o.Value = &Record_MD{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(4):
		o.Value = &Record_MF{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(5):
		o.Value = &Record_CNAME{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(6):
		o.Value = &Record_SOA{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(7):
		o.Value = &Record_MB{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(8):
		o.Value = &Record_MG{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(9):
		o.Value = &Record_MR{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(10):
		o.Value = &Record_NULL{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(11):
		o.Value = &Record_WKS{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(12):
		o.Value = &Record_PTR{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(13):
		o.Value = &Record_HINFO{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(14):
		o.Value = &Record_MINFO{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(15):
		o.Value = &Record_MX{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(16):
		o.Value = &Record_TXT{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(17):
		o.Value = &Record_RP{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(18):
		o.Value = &Record_AFSDB{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(19):
		o.Value = &Record_X25{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(20):
		o.Value = &Record_ISDN{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(21):
		o.Value = &Record_RT{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(24):
		o.Value = &Record_SIG{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(25):
		o.Value = &Record_KEY{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(28):
		o.Value = &Record_AAAA{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(30):
		o.Value = &Record_NXT{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(33):
		o.Value = &Record_SRV{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(34):
		o.Value = &Record_ATMA{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(35):
		o.Value = &Record_NAPTR{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(39):
		o.Value = &Record_DNAME{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(43):
		o.Value = &Record_DS{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(46):
		o.Value = &Record_RRSIG{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(47):
		o.Value = &Record_NSEC{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(48):
		o.Value = &Record_DNSKEY{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(49):
		o.Value = &Record_DHCID{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(50):
		o.Value = &Record_NSEC3{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(51):
		o.Value = &Record_NSEC3PARAM{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(52):
		o.Value = &Record_TLSA{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(65281):
		o.Value = &Record_WINS{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(65282):
		o.Value = &Record_WINSR{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		o.Value = &Record_Unknown{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// Record_ZERO structure represents Record_Record RPC union arm.
//
// It has following labels: 0
type Record_ZERO struct {
	RecordZERO *RecordZERO `idl:"name:recordZERO" json:"record_zero"`
}

func (*Record_ZERO) is_Record_Record() {}

func (o *Record_ZERO) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordZERO != nil {
		_ptr_recordZERO := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordZERO != nil {
				if err := o.RecordZERO.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordZERO{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordZERO, _ptr_recordZERO); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_ZERO) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordZERO := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordZERO == nil {
			o.RecordZERO = &RecordZERO{}
		}
		if err := o.RecordZERO.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordZERO := func(ptr interface{}) { o.RecordZERO = *ptr.(**RecordZERO) }
	if err := w.ReadPointer(&o.RecordZERO, _s_recordZERO, _ptr_recordZERO); err != nil {
		return err
	}
	return nil
}

// Record_A structure represents Record_Record RPC union arm.
//
// It has following labels: 1
type Record_A struct {
	RecordA *RecordA `idl:"name:recordA" json:"record_a"`
}

func (*Record_A) is_Record_Record() {}

func (o *Record_A) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordA != nil {
		_ptr_recordA := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordA != nil {
				if err := o.RecordA.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordA{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordA, _ptr_recordA); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_A) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordA := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordA == nil {
			o.RecordA = &RecordA{}
		}
		if err := o.RecordA.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordA := func(ptr interface{}) { o.RecordA = *ptr.(**RecordA) }
	if err := w.ReadPointer(&o.RecordA, _s_recordA, _ptr_recordA); err != nil {
		return err
	}
	return nil
}

// Record_NS structure represents Record_Record RPC union arm.
//
// It has following labels: 2
type Record_NS struct {
	RecordNS *RecordNS `idl:"name:recordNS" json:"record_ns"`
}

func (*Record_NS) is_Record_Record() {}

func (o *Record_NS) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordNS != nil {
		_ptr_recordNS := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordNS != nil {
				if err := o.RecordNS.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordNS{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordNS, _ptr_recordNS); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_NS) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordNS := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordNS == nil {
			o.RecordNS = &RecordNS{}
		}
		if err := o.RecordNS.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordNS := func(ptr interface{}) { o.RecordNS = *ptr.(**RecordNS) }
	if err := w.ReadPointer(&o.RecordNS, _s_recordNS, _ptr_recordNS); err != nil {
		return err
	}
	return nil
}

// Record_MD structure represents Record_Record RPC union arm.
//
// It has following labels: 3
type Record_MD struct {
	RecordMD *RecordMD `idl:"name:recordMD" json:"record_md"`
}

func (*Record_MD) is_Record_Record() {}

func (o *Record_MD) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordMD != nil {
		_ptr_recordMD := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordMD != nil {
				if err := o.RecordMD.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordMD{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordMD, _ptr_recordMD); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_MD) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordMD := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordMD == nil {
			o.RecordMD = &RecordMD{}
		}
		if err := o.RecordMD.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordMD := func(ptr interface{}) { o.RecordMD = *ptr.(**RecordMD) }
	if err := w.ReadPointer(&o.RecordMD, _s_recordMD, _ptr_recordMD); err != nil {
		return err
	}
	return nil
}

// Record_MF structure represents Record_Record RPC union arm.
//
// It has following labels: 4
type Record_MF struct {
	RecordMF *RecordMF `idl:"name:recordMF" json:"record_mf"`
}

func (*Record_MF) is_Record_Record() {}

func (o *Record_MF) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordMF != nil {
		_ptr_recordMF := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordMF != nil {
				if err := o.RecordMF.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordMF{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordMF, _ptr_recordMF); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_MF) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordMF := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordMF == nil {
			o.RecordMF = &RecordMF{}
		}
		if err := o.RecordMF.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordMF := func(ptr interface{}) { o.RecordMF = *ptr.(**RecordMF) }
	if err := w.ReadPointer(&o.RecordMF, _s_recordMF, _ptr_recordMF); err != nil {
		return err
	}
	return nil
}

// Record_CNAME structure represents Record_Record RPC union arm.
//
// It has following labels: 5
type Record_CNAME struct {
	RecordCNAME *RecordCNAME `idl:"name:recordCNAME" json:"record_cname"`
}

func (*Record_CNAME) is_Record_Record() {}

func (o *Record_CNAME) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordCNAME != nil {
		_ptr_recordCNAME := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordCNAME != nil {
				if err := o.RecordCNAME.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordCNAME{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordCNAME, _ptr_recordCNAME); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_CNAME) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordCNAME := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordCNAME == nil {
			o.RecordCNAME = &RecordCNAME{}
		}
		if err := o.RecordCNAME.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordCNAME := func(ptr interface{}) { o.RecordCNAME = *ptr.(**RecordCNAME) }
	if err := w.ReadPointer(&o.RecordCNAME, _s_recordCNAME, _ptr_recordCNAME); err != nil {
		return err
	}
	return nil
}

// Record_SOA structure represents Record_Record RPC union arm.
//
// It has following labels: 6
type Record_SOA struct {
	RecordSOA *RecordSOA `idl:"name:recordSOA" json:"record_soa"`
}

func (*Record_SOA) is_Record_Record() {}

func (o *Record_SOA) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordSOA != nil {
		_ptr_recordSOA := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordSOA != nil {
				if err := o.RecordSOA.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordSOA{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordSOA, _ptr_recordSOA); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_SOA) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordSOA := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordSOA == nil {
			o.RecordSOA = &RecordSOA{}
		}
		if err := o.RecordSOA.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordSOA := func(ptr interface{}) { o.RecordSOA = *ptr.(**RecordSOA) }
	if err := w.ReadPointer(&o.RecordSOA, _s_recordSOA, _ptr_recordSOA); err != nil {
		return err
	}
	return nil
}

// Record_MB structure represents Record_Record RPC union arm.
//
// It has following labels: 7
type Record_MB struct {
	RecordMB *RecordMB `idl:"name:recordMB" json:"record_mb"`
}

func (*Record_MB) is_Record_Record() {}

func (o *Record_MB) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordMB != nil {
		_ptr_recordMB := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordMB != nil {
				if err := o.RecordMB.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordMB{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordMB, _ptr_recordMB); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_MB) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordMB := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordMB == nil {
			o.RecordMB = &RecordMB{}
		}
		if err := o.RecordMB.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordMB := func(ptr interface{}) { o.RecordMB = *ptr.(**RecordMB) }
	if err := w.ReadPointer(&o.RecordMB, _s_recordMB, _ptr_recordMB); err != nil {
		return err
	}
	return nil
}

// Record_MG structure represents Record_Record RPC union arm.
//
// It has following labels: 8
type Record_MG struct {
	RecordMG *RecordMG `idl:"name:recordMG" json:"record_mg"`
}

func (*Record_MG) is_Record_Record() {}

func (o *Record_MG) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordMG != nil {
		_ptr_recordMG := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordMG != nil {
				if err := o.RecordMG.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordMG{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordMG, _ptr_recordMG); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_MG) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordMG := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordMG == nil {
			o.RecordMG = &RecordMG{}
		}
		if err := o.RecordMG.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordMG := func(ptr interface{}) { o.RecordMG = *ptr.(**RecordMG) }
	if err := w.ReadPointer(&o.RecordMG, _s_recordMG, _ptr_recordMG); err != nil {
		return err
	}
	return nil
}

// Record_MR structure represents Record_Record RPC union arm.
//
// It has following labels: 9
type Record_MR struct {
	RecordMR *RecordMR `idl:"name:recordMR" json:"record_mr"`
}

func (*Record_MR) is_Record_Record() {}

func (o *Record_MR) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordMR != nil {
		_ptr_recordMR := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordMR != nil {
				if err := o.RecordMR.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordMR{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordMR, _ptr_recordMR); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_MR) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordMR := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordMR == nil {
			o.RecordMR = &RecordMR{}
		}
		if err := o.RecordMR.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordMR := func(ptr interface{}) { o.RecordMR = *ptr.(**RecordMR) }
	if err := w.ReadPointer(&o.RecordMR, _s_recordMR, _ptr_recordMR); err != nil {
		return err
	}
	return nil
}

// Record_NULL structure represents Record_Record RPC union arm.
//
// It has following labels: 10
type Record_NULL struct {
	RecordNULL *RecordNULL `idl:"name:recordNULL" json:"record_null"`
}

func (*Record_NULL) is_Record_Record() {}

func (o *Record_NULL) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordNULL != nil {
		_ptr_recordNULL := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordNULL != nil {
				if err := o.RecordNULL.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordNULL{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordNULL, _ptr_recordNULL); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_NULL) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordNULL := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordNULL == nil {
			o.RecordNULL = &RecordNULL{}
		}
		if err := o.RecordNULL.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordNULL := func(ptr interface{}) { o.RecordNULL = *ptr.(**RecordNULL) }
	if err := w.ReadPointer(&o.RecordNULL, _s_recordNULL, _ptr_recordNULL); err != nil {
		return err
	}
	return nil
}

// Record_WKS structure represents Record_Record RPC union arm.
//
// It has following labels: 11
type Record_WKS struct {
	RecordWKS *RecordWKS `idl:"name:recordWKS" json:"record_wks"`
}

func (*Record_WKS) is_Record_Record() {}

func (o *Record_WKS) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordWKS != nil {
		_ptr_recordWKS := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordWKS != nil {
				if err := o.RecordWKS.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordWKS{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordWKS, _ptr_recordWKS); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_WKS) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordWKS := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordWKS == nil {
			o.RecordWKS = &RecordWKS{}
		}
		if err := o.RecordWKS.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordWKS := func(ptr interface{}) { o.RecordWKS = *ptr.(**RecordWKS) }
	if err := w.ReadPointer(&o.RecordWKS, _s_recordWKS, _ptr_recordWKS); err != nil {
		return err
	}
	return nil
}

// Record_PTR structure represents Record_Record RPC union arm.
//
// It has following labels: 12
type Record_PTR struct {
	RecordPTR *RecordPTR `idl:"name:recordPTR" json:"record_ptr"`
}

func (*Record_PTR) is_Record_Record() {}

func (o *Record_PTR) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordPTR != nil {
		_ptr_recordPTR := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordPTR != nil {
				if err := o.RecordPTR.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordPTR{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordPTR, _ptr_recordPTR); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_PTR) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordPTR := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordPTR == nil {
			o.RecordPTR = &RecordPTR{}
		}
		if err := o.RecordPTR.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordPTR := func(ptr interface{}) { o.RecordPTR = *ptr.(**RecordPTR) }
	if err := w.ReadPointer(&o.RecordPTR, _s_recordPTR, _ptr_recordPTR); err != nil {
		return err
	}
	return nil
}

// Record_HINFO structure represents Record_Record RPC union arm.
//
// It has following labels: 13
type Record_HINFO struct {
	RecordHINFO *RecordHINFO `idl:"name:recordHINFO" json:"record_hinfo"`
}

func (*Record_HINFO) is_Record_Record() {}

func (o *Record_HINFO) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordHINFO != nil {
		_ptr_recordHINFO := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordHINFO != nil {
				if err := o.RecordHINFO.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordHINFO{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordHINFO, _ptr_recordHINFO); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_HINFO) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordHINFO := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordHINFO == nil {
			o.RecordHINFO = &RecordHINFO{}
		}
		if err := o.RecordHINFO.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordHINFO := func(ptr interface{}) { o.RecordHINFO = *ptr.(**RecordHINFO) }
	if err := w.ReadPointer(&o.RecordHINFO, _s_recordHINFO, _ptr_recordHINFO); err != nil {
		return err
	}
	return nil
}

// Record_MINFO structure represents Record_Record RPC union arm.
//
// It has following labels: 14
type Record_MINFO struct {
	RecordMINFO *RecordMINFO `idl:"name:recordMINFO" json:"record_minfo"`
}

func (*Record_MINFO) is_Record_Record() {}

func (o *Record_MINFO) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordMINFO != nil {
		_ptr_recordMINFO := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordMINFO != nil {
				if err := o.RecordMINFO.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordMINFO{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordMINFO, _ptr_recordMINFO); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_MINFO) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordMINFO := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordMINFO == nil {
			o.RecordMINFO = &RecordMINFO{}
		}
		if err := o.RecordMINFO.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordMINFO := func(ptr interface{}) { o.RecordMINFO = *ptr.(**RecordMINFO) }
	if err := w.ReadPointer(&o.RecordMINFO, _s_recordMINFO, _ptr_recordMINFO); err != nil {
		return err
	}
	return nil
}

// Record_MX structure represents Record_Record RPC union arm.
//
// It has following labels: 15
type Record_MX struct {
	RecordMX *RecordMX `idl:"name:recordMX" json:"record_mx"`
}

func (*Record_MX) is_Record_Record() {}

func (o *Record_MX) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordMX != nil {
		_ptr_recordMX := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordMX != nil {
				if err := o.RecordMX.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordMX{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordMX, _ptr_recordMX); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_MX) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordMX := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordMX == nil {
			o.RecordMX = &RecordMX{}
		}
		if err := o.RecordMX.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordMX := func(ptr interface{}) { o.RecordMX = *ptr.(**RecordMX) }
	if err := w.ReadPointer(&o.RecordMX, _s_recordMX, _ptr_recordMX); err != nil {
		return err
	}
	return nil
}

// Record_TXT structure represents Record_Record RPC union arm.
//
// It has following labels: 16
type Record_TXT struct {
	RecordTXT *RecordTXT `idl:"name:recordTXT" json:"record_txt"`
}

func (*Record_TXT) is_Record_Record() {}

func (o *Record_TXT) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordTXT != nil {
		_ptr_recordTXT := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordTXT != nil {
				if err := o.RecordTXT.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordTXT{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordTXT, _ptr_recordTXT); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_TXT) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordTXT := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordTXT == nil {
			o.RecordTXT = &RecordTXT{}
		}
		if err := o.RecordTXT.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordTXT := func(ptr interface{}) { o.RecordTXT = *ptr.(**RecordTXT) }
	if err := w.ReadPointer(&o.RecordTXT, _s_recordTXT, _ptr_recordTXT); err != nil {
		return err
	}
	return nil
}

// Record_RP structure represents Record_Record RPC union arm.
//
// It has following labels: 17
type Record_RP struct {
	RecordRP *RecordRP `idl:"name:recordRP" json:"record_rp"`
}

func (*Record_RP) is_Record_Record() {}

func (o *Record_RP) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordRP != nil {
		_ptr_recordRP := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordRP != nil {
				if err := o.RecordRP.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordRP{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordRP, _ptr_recordRP); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_RP) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordRP := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordRP == nil {
			o.RecordRP = &RecordRP{}
		}
		if err := o.RecordRP.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordRP := func(ptr interface{}) { o.RecordRP = *ptr.(**RecordRP) }
	if err := w.ReadPointer(&o.RecordRP, _s_recordRP, _ptr_recordRP); err != nil {
		return err
	}
	return nil
}

// Record_AFSDB structure represents Record_Record RPC union arm.
//
// It has following labels: 18
type Record_AFSDB struct {
	RecordAFSDB *RecordAFSDB `idl:"name:recordAFSDB" json:"record_afsdb"`
}

func (*Record_AFSDB) is_Record_Record() {}

func (o *Record_AFSDB) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordAFSDB != nil {
		_ptr_recordAFSDB := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordAFSDB != nil {
				if err := o.RecordAFSDB.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordAFSDB{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordAFSDB, _ptr_recordAFSDB); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_AFSDB) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordAFSDB := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordAFSDB == nil {
			o.RecordAFSDB = &RecordAFSDB{}
		}
		if err := o.RecordAFSDB.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordAFSDB := func(ptr interface{}) { o.RecordAFSDB = *ptr.(**RecordAFSDB) }
	if err := w.ReadPointer(&o.RecordAFSDB, _s_recordAFSDB, _ptr_recordAFSDB); err != nil {
		return err
	}
	return nil
}

// Record_X25 structure represents Record_Record RPC union arm.
//
// It has following labels: 19
type Record_X25 struct {
	RecordX25 *RecordX25 `idl:"name:recordX25" json:"record_x25"`
}

func (*Record_X25) is_Record_Record() {}

func (o *Record_X25) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordX25 != nil {
		_ptr_recordX25 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordX25 != nil {
				if err := o.RecordX25.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordX25{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordX25, _ptr_recordX25); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_X25) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordX25 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordX25 == nil {
			o.RecordX25 = &RecordX25{}
		}
		if err := o.RecordX25.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordX25 := func(ptr interface{}) { o.RecordX25 = *ptr.(**RecordX25) }
	if err := w.ReadPointer(&o.RecordX25, _s_recordX25, _ptr_recordX25); err != nil {
		return err
	}
	return nil
}

// Record_ISDN structure represents Record_Record RPC union arm.
//
// It has following labels: 20
type Record_ISDN struct {
	RecordISDN *RecordISDN `idl:"name:recordISDN" json:"record_isdn"`
}

func (*Record_ISDN) is_Record_Record() {}

func (o *Record_ISDN) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordISDN != nil {
		_ptr_recordISDN := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordISDN != nil {
				if err := o.RecordISDN.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordISDN{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordISDN, _ptr_recordISDN); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_ISDN) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordISDN := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordISDN == nil {
			o.RecordISDN = &RecordISDN{}
		}
		if err := o.RecordISDN.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordISDN := func(ptr interface{}) { o.RecordISDN = *ptr.(**RecordISDN) }
	if err := w.ReadPointer(&o.RecordISDN, _s_recordISDN, _ptr_recordISDN); err != nil {
		return err
	}
	return nil
}

// Record_RT structure represents Record_Record RPC union arm.
//
// It has following labels: 21
type Record_RT struct {
	RecordRT *RecordRT `idl:"name:recordRT" json:"record_rt"`
}

func (*Record_RT) is_Record_Record() {}

func (o *Record_RT) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordRT != nil {
		_ptr_recordRT := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordRT != nil {
				if err := o.RecordRT.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordRT{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordRT, _ptr_recordRT); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_RT) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordRT := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordRT == nil {
			o.RecordRT = &RecordRT{}
		}
		if err := o.RecordRT.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordRT := func(ptr interface{}) { o.RecordRT = *ptr.(**RecordRT) }
	if err := w.ReadPointer(&o.RecordRT, _s_recordRT, _ptr_recordRT); err != nil {
		return err
	}
	return nil
}

// Record_SIG structure represents Record_Record RPC union arm.
//
// It has following labels: 24
type Record_SIG struct {
	RecordSIG *RecordSIG `idl:"name:recordSIG" json:"record_sig"`
}

func (*Record_SIG) is_Record_Record() {}

func (o *Record_SIG) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordSIG != nil {
		_ptr_recordSIG := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordSIG != nil {
				if err := o.RecordSIG.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordSIG{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordSIG, _ptr_recordSIG); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_SIG) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordSIG := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordSIG == nil {
			o.RecordSIG = &RecordSIG{}
		}
		if err := o.RecordSIG.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordSIG := func(ptr interface{}) { o.RecordSIG = *ptr.(**RecordSIG) }
	if err := w.ReadPointer(&o.RecordSIG, _s_recordSIG, _ptr_recordSIG); err != nil {
		return err
	}
	return nil
}

// Record_KEY structure represents Record_Record RPC union arm.
//
// It has following labels: 25
type Record_KEY struct {
	RecordKEY *RecordKEY `idl:"name:recordKEY" json:"record_key"`
}

func (*Record_KEY) is_Record_Record() {}

func (o *Record_KEY) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordKEY != nil {
		_ptr_recordKEY := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordKEY != nil {
				if err := o.RecordKEY.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordKEY{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordKEY, _ptr_recordKEY); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_KEY) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordKEY := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordKEY == nil {
			o.RecordKEY = &RecordKEY{}
		}
		if err := o.RecordKEY.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordKEY := func(ptr interface{}) { o.RecordKEY = *ptr.(**RecordKEY) }
	if err := w.ReadPointer(&o.RecordKEY, _s_recordKEY, _ptr_recordKEY); err != nil {
		return err
	}
	return nil
}

// Record_AAAA structure represents Record_Record RPC union arm.
//
// It has following labels: 28
type Record_AAAA struct {
	RecordAAAA *RecordAAAA `idl:"name:recordAAAA" json:"record_aaaa"`
}

func (*Record_AAAA) is_Record_Record() {}

func (o *Record_AAAA) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordAAAA != nil {
		_ptr_recordAAAA := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordAAAA != nil {
				if err := o.RecordAAAA.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordAAAA{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordAAAA, _ptr_recordAAAA); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_AAAA) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordAAAA := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordAAAA == nil {
			o.RecordAAAA = &RecordAAAA{}
		}
		if err := o.RecordAAAA.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordAAAA := func(ptr interface{}) { o.RecordAAAA = *ptr.(**RecordAAAA) }
	if err := w.ReadPointer(&o.RecordAAAA, _s_recordAAAA, _ptr_recordAAAA); err != nil {
		return err
	}
	return nil
}

// Record_NXT structure represents Record_Record RPC union arm.
//
// It has following labels: 30
type Record_NXT struct {
	RecordNXT *RecordNXT `idl:"name:recordNXT" json:"record_nxt"`
}

func (*Record_NXT) is_Record_Record() {}

func (o *Record_NXT) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordNXT != nil {
		_ptr_recordNXT := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordNXT != nil {
				if err := o.RecordNXT.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordNXT{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordNXT, _ptr_recordNXT); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_NXT) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordNXT := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordNXT == nil {
			o.RecordNXT = &RecordNXT{}
		}
		if err := o.RecordNXT.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordNXT := func(ptr interface{}) { o.RecordNXT = *ptr.(**RecordNXT) }
	if err := w.ReadPointer(&o.RecordNXT, _s_recordNXT, _ptr_recordNXT); err != nil {
		return err
	}
	return nil
}

// Record_SRV structure represents Record_Record RPC union arm.
//
// It has following labels: 33
type Record_SRV struct {
	RecordSRV *RecordSRV `idl:"name:recordSRV" json:"record_srv"`
}

func (*Record_SRV) is_Record_Record() {}

func (o *Record_SRV) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordSRV != nil {
		_ptr_recordSRV := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordSRV != nil {
				if err := o.RecordSRV.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordSRV{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordSRV, _ptr_recordSRV); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_SRV) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordSRV := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordSRV == nil {
			o.RecordSRV = &RecordSRV{}
		}
		if err := o.RecordSRV.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordSRV := func(ptr interface{}) { o.RecordSRV = *ptr.(**RecordSRV) }
	if err := w.ReadPointer(&o.RecordSRV, _s_recordSRV, _ptr_recordSRV); err != nil {
		return err
	}
	return nil
}

// Record_ATMA structure represents Record_Record RPC union arm.
//
// It has following labels: 34
type Record_ATMA struct {
	RecordATMA *RecordATMA `idl:"name:recordATMA" json:"record_atma"`
}

func (*Record_ATMA) is_Record_Record() {}

func (o *Record_ATMA) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordATMA != nil {
		_ptr_recordATMA := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordATMA != nil {
				if err := o.RecordATMA.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordATMA{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordATMA, _ptr_recordATMA); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_ATMA) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordATMA := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordATMA == nil {
			o.RecordATMA = &RecordATMA{}
		}
		if err := o.RecordATMA.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordATMA := func(ptr interface{}) { o.RecordATMA = *ptr.(**RecordATMA) }
	if err := w.ReadPointer(&o.RecordATMA, _s_recordATMA, _ptr_recordATMA); err != nil {
		return err
	}
	return nil
}

// Record_NAPTR structure represents Record_Record RPC union arm.
//
// It has following labels: 35
type Record_NAPTR struct {
	RecordNAPTR *RecordNAPTR `idl:"name:recordNAPTR" json:"record_naptr"`
}

func (*Record_NAPTR) is_Record_Record() {}

func (o *Record_NAPTR) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordNAPTR != nil {
		_ptr_recordNAPTR := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordNAPTR != nil {
				if err := o.RecordNAPTR.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordNAPTR{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordNAPTR, _ptr_recordNAPTR); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_NAPTR) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordNAPTR := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordNAPTR == nil {
			o.RecordNAPTR = &RecordNAPTR{}
		}
		if err := o.RecordNAPTR.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordNAPTR := func(ptr interface{}) { o.RecordNAPTR = *ptr.(**RecordNAPTR) }
	if err := w.ReadPointer(&o.RecordNAPTR, _s_recordNAPTR, _ptr_recordNAPTR); err != nil {
		return err
	}
	return nil
}

// Record_DNAME structure represents Record_Record RPC union arm.
//
// It has following labels: 39
type Record_DNAME struct {
	RecordDNAME *RecordDNAME `idl:"name:recordDNAME" json:"record_dname"`
}

func (*Record_DNAME) is_Record_Record() {}

func (o *Record_DNAME) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordDNAME != nil {
		_ptr_recordDNAME := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordDNAME != nil {
				if err := o.RecordDNAME.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordDNAME{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordDNAME, _ptr_recordDNAME); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_DNAME) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordDNAME := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordDNAME == nil {
			o.RecordDNAME = &RecordDNAME{}
		}
		if err := o.RecordDNAME.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordDNAME := func(ptr interface{}) { o.RecordDNAME = *ptr.(**RecordDNAME) }
	if err := w.ReadPointer(&o.RecordDNAME, _s_recordDNAME, _ptr_recordDNAME); err != nil {
		return err
	}
	return nil
}

// Record_DS structure represents Record_Record RPC union arm.
//
// It has following labels: 43
type Record_DS struct {
	RecordDS *RecordDS `idl:"name:recordDS" json:"record_ds"`
}

func (*Record_DS) is_Record_Record() {}

func (o *Record_DS) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordDS != nil {
		_ptr_recordDS := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordDS != nil {
				if err := o.RecordDS.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordDS{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordDS, _ptr_recordDS); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_DS) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordDS := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordDS == nil {
			o.RecordDS = &RecordDS{}
		}
		if err := o.RecordDS.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordDS := func(ptr interface{}) { o.RecordDS = *ptr.(**RecordDS) }
	if err := w.ReadPointer(&o.RecordDS, _s_recordDS, _ptr_recordDS); err != nil {
		return err
	}
	return nil
}

// Record_RRSIG structure represents Record_Record RPC union arm.
//
// It has following labels: 46
type Record_RRSIG struct {
	RecordRRSIG *RecordRRSIG `idl:"name:recordRRSIG" json:"record_rrsig"`
}

func (*Record_RRSIG) is_Record_Record() {}

func (o *Record_RRSIG) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordRRSIG != nil {
		_ptr_recordRRSIG := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordRRSIG != nil {
				if err := o.RecordRRSIG.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordRRSIG{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordRRSIG, _ptr_recordRRSIG); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_RRSIG) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordRRSIG := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordRRSIG == nil {
			o.RecordRRSIG = &RecordRRSIG{}
		}
		if err := o.RecordRRSIG.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordRRSIG := func(ptr interface{}) { o.RecordRRSIG = *ptr.(**RecordRRSIG) }
	if err := w.ReadPointer(&o.RecordRRSIG, _s_recordRRSIG, _ptr_recordRRSIG); err != nil {
		return err
	}
	return nil
}

// Record_NSEC structure represents Record_Record RPC union arm.
//
// It has following labels: 47
type Record_NSEC struct {
	RecordNSEC *RecordNSEC `idl:"name:recordNSEC" json:"record_nsec"`
}

func (*Record_NSEC) is_Record_Record() {}

func (o *Record_NSEC) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordNSEC != nil {
		_ptr_recordNSEC := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordNSEC != nil {
				if err := o.RecordNSEC.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordNSEC{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordNSEC, _ptr_recordNSEC); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_NSEC) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordNSEC := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordNSEC == nil {
			o.RecordNSEC = &RecordNSEC{}
		}
		if err := o.RecordNSEC.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordNSEC := func(ptr interface{}) { o.RecordNSEC = *ptr.(**RecordNSEC) }
	if err := w.ReadPointer(&o.RecordNSEC, _s_recordNSEC, _ptr_recordNSEC); err != nil {
		return err
	}
	return nil
}

// Record_DNSKEY structure represents Record_Record RPC union arm.
//
// It has following labels: 48
type Record_DNSKEY struct {
	RecordDNSKEY *RecordDNSKEY `idl:"name:recordDNSKEY" json:"record_dnskey"`
}

func (*Record_DNSKEY) is_Record_Record() {}

func (o *Record_DNSKEY) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordDNSKEY != nil {
		_ptr_recordDNSKEY := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordDNSKEY != nil {
				if err := o.RecordDNSKEY.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordDNSKEY{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordDNSKEY, _ptr_recordDNSKEY); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_DNSKEY) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordDNSKEY := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordDNSKEY == nil {
			o.RecordDNSKEY = &RecordDNSKEY{}
		}
		if err := o.RecordDNSKEY.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordDNSKEY := func(ptr interface{}) { o.RecordDNSKEY = *ptr.(**RecordDNSKEY) }
	if err := w.ReadPointer(&o.RecordDNSKEY, _s_recordDNSKEY, _ptr_recordDNSKEY); err != nil {
		return err
	}
	return nil
}

// Record_DHCID structure represents Record_Record RPC union arm.
//
// It has following labels: 49
type Record_DHCID struct {
	RecordDHCID *RecordDHCID `idl:"name:recordDHCID" json:"record_dhcid"`
}

func (*Record_DHCID) is_Record_Record() {}

func (o *Record_DHCID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordDHCID != nil {
		_ptr_recordDHCID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordDHCID != nil {
				if err := o.RecordDHCID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordDHCID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordDHCID, _ptr_recordDHCID); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_DHCID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordDHCID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordDHCID == nil {
			o.RecordDHCID = &RecordDHCID{}
		}
		if err := o.RecordDHCID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordDHCID := func(ptr interface{}) { o.RecordDHCID = *ptr.(**RecordDHCID) }
	if err := w.ReadPointer(&o.RecordDHCID, _s_recordDHCID, _ptr_recordDHCID); err != nil {
		return err
	}
	return nil
}

// Record_NSEC3 structure represents Record_Record RPC union arm.
//
// It has following labels: 50
type Record_NSEC3 struct {
	RecordNSEC3 *RecordNSEC3 `idl:"name:recordNSEC3" json:"record_nsec3"`
}

func (*Record_NSEC3) is_Record_Record() {}

func (o *Record_NSEC3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordNSEC3 != nil {
		_ptr_recordNSEC3 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordNSEC3 != nil {
				if err := o.RecordNSEC3.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordNSEC3{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordNSEC3, _ptr_recordNSEC3); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_NSEC3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordNSEC3 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordNSEC3 == nil {
			o.RecordNSEC3 = &RecordNSEC3{}
		}
		if err := o.RecordNSEC3.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordNSEC3 := func(ptr interface{}) { o.RecordNSEC3 = *ptr.(**RecordNSEC3) }
	if err := w.ReadPointer(&o.RecordNSEC3, _s_recordNSEC3, _ptr_recordNSEC3); err != nil {
		return err
	}
	return nil
}

// Record_NSEC3PARAM structure represents Record_Record RPC union arm.
//
// It has following labels: 51
type Record_NSEC3PARAM struct {
	RecordNSEC3PARAM *RecordNSEC3PARAM `idl:"name:recordNSEC3PARAM" json:"record_nsec3param"`
}

func (*Record_NSEC3PARAM) is_Record_Record() {}

func (o *Record_NSEC3PARAM) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordNSEC3PARAM != nil {
		_ptr_recordNSEC3PARAM := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordNSEC3PARAM != nil {
				if err := o.RecordNSEC3PARAM.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordNSEC3PARAM{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordNSEC3PARAM, _ptr_recordNSEC3PARAM); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_NSEC3PARAM) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordNSEC3PARAM := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordNSEC3PARAM == nil {
			o.RecordNSEC3PARAM = &RecordNSEC3PARAM{}
		}
		if err := o.RecordNSEC3PARAM.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordNSEC3PARAM := func(ptr interface{}) { o.RecordNSEC3PARAM = *ptr.(**RecordNSEC3PARAM) }
	if err := w.ReadPointer(&o.RecordNSEC3PARAM, _s_recordNSEC3PARAM, _ptr_recordNSEC3PARAM); err != nil {
		return err
	}
	return nil
}

// Record_TLSA structure represents Record_Record RPC union arm.
//
// It has following labels: 52
type Record_TLSA struct {
	RecordTLSA *RecordTLSA `idl:"name:recordTLSA" json:"record_tlsa"`
}

func (*Record_TLSA) is_Record_Record() {}

func (o *Record_TLSA) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordTLSA != nil {
		_ptr_recordTLSA := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordTLSA != nil {
				if err := o.RecordTLSA.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordTLSA{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordTLSA, _ptr_recordTLSA); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_TLSA) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordTLSA := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordTLSA == nil {
			o.RecordTLSA = &RecordTLSA{}
		}
		if err := o.RecordTLSA.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordTLSA := func(ptr interface{}) { o.RecordTLSA = *ptr.(**RecordTLSA) }
	if err := w.ReadPointer(&o.RecordTLSA, _s_recordTLSA, _ptr_recordTLSA); err != nil {
		return err
	}
	return nil
}

// Record_WINS structure represents Record_Record RPC union arm.
//
// It has following labels: 65281
type Record_WINS struct {
	RecordWINS *RecordWINS `idl:"name:recordWINS" json:"record_wins"`
}

func (*Record_WINS) is_Record_Record() {}

func (o *Record_WINS) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordWINS != nil {
		_ptr_recordWINS := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordWINS != nil {
				if err := o.RecordWINS.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordWINS{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordWINS, _ptr_recordWINS); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_WINS) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordWINS := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordWINS == nil {
			o.RecordWINS = &RecordWINS{}
		}
		if err := o.RecordWINS.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordWINS := func(ptr interface{}) { o.RecordWINS = *ptr.(**RecordWINS) }
	if err := w.ReadPointer(&o.RecordWINS, _s_recordWINS, _ptr_recordWINS); err != nil {
		return err
	}
	return nil
}

// Record_WINSR structure represents Record_Record RPC union arm.
//
// It has following labels: 65282
type Record_WINSR struct {
	RecordWINSR *RecordWINSR `idl:"name:recordWINSR" json:"record_winsr"`
}

func (*Record_WINSR) is_Record_Record() {}

func (o *Record_WINSR) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordWINSR != nil {
		_ptr_recordWINSR := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordWINSR != nil {
				if err := o.RecordWINSR.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordWINSR{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordWINSR, _ptr_recordWINSR); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_WINSR) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordWINSR := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordWINSR == nil {
			o.RecordWINSR = &RecordWINSR{}
		}
		if err := o.RecordWINSR.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordWINSR := func(ptr interface{}) { o.RecordWINSR = *ptr.(**RecordWINSR) }
	if err := w.ReadPointer(&o.RecordWINSR, _s_recordWINSR, _ptr_recordWINSR); err != nil {
		return err
	}
	return nil
}

// Record_Unknown structure represents Record_Record RPC default union arm.
type Record_Unknown struct {
	RecordUnknown *RecordUnknown `idl:"name:recordUnknown" json:"record_unknown"`
}

func (*Record_Unknown) is_Record_Record() {}

func (o *Record_Unknown) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RecordUnknown != nil {
		_ptr_recordUnknown := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordUnknown != nil {
				if err := o.RecordUnknown.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RecordUnknown{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordUnknown, _ptr_recordUnknown); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record_Unknown) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_recordUnknown := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordUnknown == nil {
			o.RecordUnknown = &RecordUnknown{}
		}
		if err := o.RecordUnknown.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_recordUnknown := func(ptr interface{}) { o.RecordUnknown = *ptr.(**RecordUnknown) }
	if err := w.ReadPointer(&o.RecordUnknown, _s_recordUnknown, _ptr_recordUnknown); err != nil {
		return err
	}
	return nil
}

// NodeList structure represents DNS_RPC_NODE_LIST RPC structure.
type NodeList struct {
	DNSNodes []*Node `idl:"name:dnsNodes" json:"dns_nodes"`
}

func (o *NodeList) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *NodeList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	for i1 := range o.DNSNodes {
		i1 := i1
		if o.DNSNodes[i1] != nil {
			_ptr_dnsNodes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DNSNodes[i1] != nil {
					if err := o.DNSNodes[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&Node{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DNSNodes[i1], _ptr_dnsNodes); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	}
	return nil
}
func (o *NodeList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	for i1 := 0; w.Len() > 0; i1++ {
		i1 := i1
		o.DNSNodes = append(o.DNSNodes, nil)
		_ptr_dnsNodes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DNSNodes[i1] == nil {
				o.DNSNodes[i1] = &Node{}
			}
			if err := o.DNSNodes[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_dnsNodes := func(ptr interface{}) { o.DNSNodes[i1] = *ptr.(**Node) }
		if err := w.ReadPointer(&o.DNSNodes[i1], _s_dnsNodes, _ptr_dnsNodes); err != nil {
			return err
		}
	}
	return nil
}

// Node structure represents DNS_RPC_NODE RPC structure.
//
// The DNS_RPC_NODE defines a structure that is used as a header for a list of DNS_RPC_RECORD
// structures (section 2.2.2.2.5) returned by the DNS server inside a DNS_RPC_BUFFER
// (section 2.2.1.2.2) while processing the R_DnssrvEnumRecords2 (section 3.1.4.9).
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| wLength                                                       | wRecordCount                                                  |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| dwFlags                                                                                                                       |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| dwChildCount                                                                                                                  |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| dnsNodeName (variable)                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type Node struct {
	// wLength (2 bytes): The length of this structure, in bytes, including the fixed size
	// elements plus the length of the dnsNodeName element. The length of this structure
	// is always 4-byte aligned, so it is possible to have 0-3 bytes of padding at the end
	// of the structure. The pad bytes are included in the wLength count.
	Length uint16 `idl:"name:wLength" json:"length"`
	// wRecordCount (2 bytes): The number of DNS_RPC_RECORD structures that follow this
	// node structure.
	RecordCount uint16 `idl:"name:wRecordCount" json:"record_count"`
	// dwFlags (4 bytes): The properties of the DNS_RPC_NODE structure.
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
	// dwChildCount (4 bytes): The total number of children nodes below this node in the
	// DNS record database.
	ChildCount uint32 `idl:"name:dwChildCount" json:"child_count"`
	// dnsNodeName (variable): The name of this node in DNS_RPC_NAME (section 2.2.2.2.1)
	// format.
	DNSNodeName *NodeName `idl:"name:dnsNodeName" json:"dns_node_name"`
	DNSRecords  []*Record `idl:"name:dnsRecords;size_is:(wRecordCount)" json:"dns_records"`
}

func (o *Node) xxx_PreparePayload(ctx context.Context) error {
	if o.DNSRecords != nil && o.RecordCount == 0 {
		o.RecordCount = uint16(len(o.DNSRecords))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Node) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if err := w.WriteData(o.RecordCount); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.ChildCount); err != nil {
		return err
	}
	if o.DNSNodeName != nil {
		_ptr_dnsNodeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.DNSNodeName != nil {
				if err := o.DNSNodeName.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.DNSNodeName, _ptr_dnsNodeName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DNSRecords != nil || o.RecordCount > 0 {
		_ptr_dnsRecords := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.RecordCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.DNSRecords {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.DNSRecords[i1] != nil {
					if err := o.DNSRecords[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&Record{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.DNSRecords); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&Record{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.DNSRecords, _ptr_dnsRecords); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	// pad 4
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	return nil
}
func (o *Node) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	if err := w.ReadData(&o.RecordCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.ChildCount); err != nil {
		return err
	}
	_ptr_dnsNodeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.DNSNodeName == nil {
			o.DNSNodeName = &NodeName{}
		}
		if err := o.DNSNodeName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_dnsNodeName := func(ptr interface{}) { o.DNSNodeName = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.DNSNodeName, _s_dnsNodeName, _ptr_dnsNodeName); err != nil {
		return err
	}
	_ptr_dnsRecords := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.RecordCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.RecordCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.DNSRecords", sizeInfo[0])
		}
		o.DNSRecords = make([]*Record, sizeInfo[0])
		for i1 := range o.DNSRecords {
			i1 := i1
			if o.DNSRecords[i1] == nil {
				o.DNSRecords[i1] = &Record{}
			}
			if err := o.DNSRecords[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_dnsRecords := func(ptr interface{}) { o.DNSRecords = *ptr.(*[]*Record) }
	if err := w.ReadPointer(&o.DNSRecords, _s_dnsRecords, _ptr_dnsRecords); err != nil {
		return err
	}
	// pad 4
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	return nil
}

// NodeName structure represents DNS_NODE_NAME RPC structure.
type NodeName struct {
	NameLength uint8  `idl:"name:cchNameLength" json:"name_length"`
	DNSName    []byte `idl:"name:dnsName;size_is:(cchNameLength)" json:"dns_name"`
}

func (o *NodeName) xxx_PreparePayload(ctx context.Context) error {
	if o.DNSName != nil && o.NameLength == 0 {
		o.NameLength = uint8(len(o.DNSName))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *NodeName) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if err := w.WriteData(o.NameLength); err != nil {
		return err
	}
	if o.DNSName != nil || o.NameLength > 0 {
		_ptr_dnsName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.NameLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.DNSName {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.DNSName[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.DNSName); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.DNSName, _ptr_dnsName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *NodeName) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	if err := w.ReadData(&o.NameLength); err != nil {
		return err
	}
	_ptr_dnsName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.NameLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.NameLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.DNSName", sizeInfo[0])
		}
		o.DNSName = make([]byte, sizeInfo[0])
		for i1 := range o.DNSName {
			i1 := i1
			if err := w.ReadData(&o.DNSName[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_dnsName := func(ptr interface{}) { o.DNSName = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.DNSName, _s_dnsName, _ptr_dnsName); err != nil {
		return err
	}
	return nil
}

// NodeText structure represents DNS_NODE_TEXT RPC structure.
type NodeText struct {
	TextLength uint8  `idl:"name:cchTextLength" json:"text_length"`
	Text       []byte `idl:"name:text;size_is:(cchTextLength)" json:"text"`
}

func (o *NodeText) xxx_PreparePayload(ctx context.Context) error {
	if o.Text != nil && o.TextLength == 0 {
		o.TextLength = uint8(len(o.Text))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *NodeText) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if err := w.WriteData(o.TextLength); err != nil {
		return err
	}
	if o.Text != nil || o.TextLength > 0 {
		_ptr_text := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.TextLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Text {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Text[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Text); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Text, _ptr_text); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *NodeText) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	if err := w.ReadData(&o.TextLength); err != nil {
		return err
	}
	_ptr_text := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.TextLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.TextLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Text", sizeInfo[0])
		}
		o.Text = make([]byte, sizeInfo[0])
		for i1 := range o.Text {
			i1 := i1
			if err := w.ReadData(&o.Text[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_text := func(ptr interface{}) { o.Text = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Text, _s_text, _ptr_text); err != nil {
		return err
	}
	return nil
}

// RecordUnknown structure represents DNS_RPC_RECORD_UNKNOWN RPC structure.
type RecordUnknown struct {
	Data []byte `idl:"name:bData" json:"data"`
}

func (o *RecordUnknown) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordUnknown) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordUnknown) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	for i1 := 0; w.Len() > 0; i1++ {
		i1 := i1
		o.Data = append(o.Data, uint8(0))
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// RecordNULL structure represents DNS_RPC_RECORD_NULL RPC structure.
//
// The DNS_RPC_RECORD_NULL structure contains information for any record for which there
// is no more specific DNS_RPC_RECORD structure. This record MUST be formatted as follows:
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| bData (variable)                                                                                                              |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type RecordNULL struct {
	// bData (variable): An array of data. The sender can provide any data in this array.
	Data []byte `idl:"name:bData" json:"data"`
}

func (o *RecordNULL) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordNULL) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordNULL) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	for i1 := 0; w.Len() > 0; i1++ {
		i1 := i1
		o.Data = append(o.Data, uint8(0))
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// RecordMailError structure represents DNS_RPC_RECORD_MAIL_ERROR RPC structure.
//
// The DNS_RPC_RECORD_MAIL_ERROR structure contains information about a DNS record of
// any of the following types:
//
// * DNS_TYPE_MINFO
//
// * DNS_TYPE_RP
//
// This record MUST be formatted as follows:
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| nameMailBox (variable)                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ErrorMailBox (variable)                                                                                                       |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type RecordMailError struct {
	// nameMailBox (variable): A structure of type DNS_RPC_NAME (section 2.2.2.2.1) containing
	// the RMAILBX value specified in section 3.3.7 of [RFC1035] for an MINFO record, or
	// the mbox-dname value specified in section 2.2 of [RFC1183] for an RP record.
	MailBox *NodeName `idl:"name:nameMailBox" json:"mail_box"`
	// ErrorMailBox (variable): A structure of type DNS_RPC_NAME (section 2.2.2.2.1) containing
	// the EMAILBX value specified in section 3.3.7 of [RFC1035] for an MINFO record, or
	// the txt-dname value specified in section 2.2 of [RFC1183] for an RP record.
	ErrorMailBox *NodeName `idl:"name:ErrorMailBox" json:"error_mail_box"`
}

func (o *RecordMailError) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordMailError) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.MailBox != nil {
		_ptr_nameMailBox := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.MailBox != nil {
				if err := o.MailBox.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.MailBox, _ptr_nameMailBox); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ErrorMailBox != nil {
		_ptr_ErrorMailBox := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ErrorMailBox != nil {
				if err := o.ErrorMailBox.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ErrorMailBox, _ptr_ErrorMailBox); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordMailError) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_nameMailBox := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.MailBox == nil {
			o.MailBox = &NodeName{}
		}
		if err := o.MailBox.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_nameMailBox := func(ptr interface{}) { o.MailBox = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.MailBox, _s_nameMailBox, _ptr_nameMailBox); err != nil {
		return err
	}
	_ptr_ErrorMailBox := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ErrorMailBox == nil {
			o.ErrorMailBox = &NodeName{}
		}
		if err := o.ErrorMailBox.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ErrorMailBox := func(ptr interface{}) { o.ErrorMailBox = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.ErrorMailBox, _s_ErrorMailBox, _ptr_ErrorMailBox); err != nil {
		return err
	}
	return nil
}

// RecordMINFO structure represents DNS_RPC_RECORD_MINFO RPC structure.
type RecordMINFO RecordMailError

func (o *RecordMINFO) RecordMailError() *RecordMailError { return (*RecordMailError)(o) }

func (o *RecordMINFO) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordMINFO) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.MailBox != nil {
		_ptr_nameMailBox := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.MailBox != nil {
				if err := o.MailBox.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.MailBox, _ptr_nameMailBox); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ErrorMailBox != nil {
		_ptr_ErrorMailBox := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ErrorMailBox != nil {
				if err := o.ErrorMailBox.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ErrorMailBox, _ptr_ErrorMailBox); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordMINFO) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_nameMailBox := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.MailBox == nil {
			o.MailBox = &NodeName{}
		}
		if err := o.MailBox.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_nameMailBox := func(ptr interface{}) { o.MailBox = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.MailBox, _s_nameMailBox, _ptr_nameMailBox); err != nil {
		return err
	}
	_ptr_ErrorMailBox := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ErrorMailBox == nil {
			o.ErrorMailBox = &NodeName{}
		}
		if err := o.ErrorMailBox.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ErrorMailBox := func(ptr interface{}) { o.ErrorMailBox = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.ErrorMailBox, _s_ErrorMailBox, _ptr_ErrorMailBox); err != nil {
		return err
	}
	return nil
}

// RecordRP structure represents DNS_RPC_RECORD_RP RPC structure.
type RecordRP RecordMailError

func (o *RecordRP) RecordMailError() *RecordMailError { return (*RecordMailError)(o) }

func (o *RecordRP) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordRP) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.MailBox != nil {
		_ptr_nameMailBox := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.MailBox != nil {
				if err := o.MailBox.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.MailBox, _ptr_nameMailBox); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ErrorMailBox != nil {
		_ptr_ErrorMailBox := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ErrorMailBox != nil {
				if err := o.ErrorMailBox.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ErrorMailBox, _ptr_ErrorMailBox); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordRP) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_nameMailBox := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.MailBox == nil {
			o.MailBox = &NodeName{}
		}
		if err := o.MailBox.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_nameMailBox := func(ptr interface{}) { o.MailBox = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.MailBox, _s_nameMailBox, _ptr_nameMailBox); err != nil {
		return err
	}
	_ptr_ErrorMailBox := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ErrorMailBox == nil {
			o.ErrorMailBox = &NodeName{}
		}
		if err := o.ErrorMailBox.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ErrorMailBox := func(ptr interface{}) { o.ErrorMailBox = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.ErrorMailBox, _s_ErrorMailBox, _ptr_ErrorMailBox); err != nil {
		return err
	}
	return nil
}

// RecordZERO structure represents DNS_RPC_RECORD_ZERO RPC structure.
type RecordZERO struct {
	EntombedTime *dtyp.Filetime `idl:"name:EntombedTime" json:"entombed_time"`
}

func (o *RecordZERO) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordZERO) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.EntombedTime != nil {
		if err := o.EntombedTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordZERO) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.EntombedTime == nil {
		o.EntombedTime = &dtyp.Filetime{}
	}
	if err := o.EntombedTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RecordRRSIG structure represents DNS_RPC_RECORD_RRSIG RPC structure.
//
// The DNS_RPC_RECORD_RRSIG structure contains information about cryptographic public
// key signatures as specified in section 3 of [RFC4034].<13> This record MUST be formatted
// as follows.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| wTypeCovered                                                  | chAlgorithm                   | chLabelCount                  |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| dwOriginalTtl                                                                                                                 |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| dwSigExpiration                                                                                                               |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| dwSigInception                                                                                                                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| wKeyTag                                                       | nameSigner (variable)                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| SignatureInfo (variable)                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type RecordRRSIG struct {
	// wTypeCovered (2 bytes): The type covered value for RRSIG RR as specified in section
	// 3.1 of [RFC4034].
	TypeCovered uint16 `idl:"name:wTypeCovered" json:"type_covered"`
	// chAlgorithm  (1 byte): The algorithm value for RRSIG RR as specified in section
	// 3.1 of [RFC4034].
	Algorithm uint8 `idl:"name:chAlgorithm" json:"algorithm"`
	// chLabelCount (1 byte): The total number of labels present in the name of the record
	// signed by the RRSIG RR as specified in section 3.1 of [RFC4034].
	LabelCount uint8 `idl:"name:chLabelCount" json:"label_count"`
	// dwOriginalTtl (4 bytes): The original TTL value of the record signed by the RRSIG
	// RR as specified in section 3.1 of [RFC4034].
	OriginalTTL uint32 `idl:"name:dwOriginalTtl" json:"original_ttl"`
	// dwSigExpiration  (4 bytes): The signature expiration time as specified in section
	// 3.1 of [RFC4034].
	SigExpiration uint32 `idl:"name:dwSigExpiration" json:"sig_expiration"`
	// dwSigInception  (4 bytes): The signature inception time as specified in section
	// 3.1 of [RFC4034].
	SigInception uint32 `idl:"name:dwSigInception" json:"sig_inception"`
	// wKeyTag  (2 bytes): The tag value for RRSIG RR as specified in section 3.1 of [RFC4034].
	KeyTag uint16 `idl:"name:wKeyTag" json:"key_tag"`
	// nameSigner (variable): A structure of type DNS_RPC_NAME (section 2.2.2.2.1) containing
	// the FQDN of the originating host for this record.
	Signer *NodeName `idl:"name:nameSigner" json:"signer"`
	// SignatureInfo  (variable): Binary signature information as specified in section
	// 3.1 of [RFC4034].
	SignatureInfo []byte `idl:"name:SignatureInfo" json:"signature_info"`
}

func (o *RecordRRSIG) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordRRSIG) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.TypeCovered); err != nil {
		return err
	}
	if err := w.WriteData(o.Algorithm); err != nil {
		return err
	}
	if err := w.WriteData(o.LabelCount); err != nil {
		return err
	}
	if err := w.WriteData(o.OriginalTTL); err != nil {
		return err
	}
	if err := w.WriteData(o.SigExpiration); err != nil {
		return err
	}
	if err := w.WriteData(o.SigInception); err != nil {
		return err
	}
	if err := w.WriteData(o.KeyTag); err != nil {
		return err
	}
	if o.Signer != nil {
		_ptr_nameSigner := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Signer != nil {
				if err := o.Signer.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Signer, _ptr_nameSigner); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	for i1 := range o.SignatureInfo {
		i1 := i1
		if err := w.WriteData(o.SignatureInfo[i1]); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *RecordRRSIG) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.TypeCovered); err != nil {
		return err
	}
	if err := w.ReadData(&o.Algorithm); err != nil {
		return err
	}
	if err := w.ReadData(&o.LabelCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.OriginalTTL); err != nil {
		return err
	}
	if err := w.ReadData(&o.SigExpiration); err != nil {
		return err
	}
	if err := w.ReadData(&o.SigInception); err != nil {
		return err
	}
	if err := w.ReadData(&o.KeyTag); err != nil {
		return err
	}
	_ptr_nameSigner := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Signer == nil {
			o.Signer = &NodeName{}
		}
		if err := o.Signer.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_nameSigner := func(ptr interface{}) { o.Signer = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.Signer, _s_nameSigner, _ptr_nameSigner); err != nil {
		return err
	}
	for i1 := 0; w.Len() > 0; i1++ {
		i1 := i1
		o.SignatureInfo = append(o.SignatureInfo, uint8(0))
		if err := w.ReadData(&o.SignatureInfo[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// RecordSIG structure represents DNS_RPC_RECORD_SIG RPC structure.
//
// The DNS_RPC_RECORD_SIG structure contains information about cryptographic public
// key signatures as specified in section 4 of [RFC2535].<12> This record MUST be formatted
// as follows:
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| wTypeCovered                                                  | chAlgorithm                   | chLabelCount                  |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| dwOriginalTtl                                                                                                                 |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| dwSigExpiration                                                                                                               |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| dwSigInception                                                                                                                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| wKeyTag                                                       | nameSigner (variable)                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| SignatureInfo (variable)                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type RecordSIG struct {
	// wTypeCovered (2 bytes): The type covered value for SIG RR as specified in section
	// 4.1 of [RFC2535].
	TypeCovered uint16 `idl:"name:wTypeCovered" json:"type_covered"`
	// chAlgorithm (1 byte): The algorithm value for SIG RR as specified in section 4.1
	// of [RFC2535].
	Algorithm uint8 `idl:"name:chAlgorithm" json:"algorithm"`
	// chLabelCount (1 byte): The total number of labels present in the name of the record
	// signed by the SIG RR as specified in section 4.1 of [RFC2535].
	LabelCount uint8 `idl:"name:chLabelCount" json:"label_count"`
	// dwOriginalTtl (4 bytes): The original TTL value of the record signed by the SIG RR
	// as specified in section 4.1 of [RFC2535].
	OriginalTTL uint32 `idl:"name:dwOriginalTtl" json:"original_ttl"`
	// dwSigExpiration (4 bytes): The signature expiration time as specified in section
	// 4.1 of [RFC2535].
	SigExpiration uint32 `idl:"name:dwSigExpiration" json:"sig_expiration"`
	// dwSigInception (4 bytes): The signature inception time as specified in section 4.1
	// of [RFC2535].
	SigInception uint32 `idl:"name:dwSigInception" json:"sig_inception"`
	// wKeyTag (2 bytes): The key tag value for SIG RR as specified in section 4.1 of [RFC2535].
	KeyTag uint16 `idl:"name:wKeyTag" json:"key_tag"`
	// nameSigner (variable): Pointer to a structure of type DNS_RPC_NAME (section 2.2.2.2.1)
	// containing the FQDN of the originating host for this record.
	Signer *NodeName `idl:"name:nameSigner" json:"signer"`
	// SignatureInfo (variable): Binary signature information as specified in section 4.1
	// of [RFC2535].
	SignatureInfo []byte `idl:"name:SignatureInfo" json:"signature_info"`
}

func (o *RecordSIG) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordSIG) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.TypeCovered); err != nil {
		return err
	}
	if err := w.WriteData(o.Algorithm); err != nil {
		return err
	}
	if err := w.WriteData(o.LabelCount); err != nil {
		return err
	}
	if err := w.WriteData(o.OriginalTTL); err != nil {
		return err
	}
	if err := w.WriteData(o.SigExpiration); err != nil {
		return err
	}
	if err := w.WriteData(o.SigInception); err != nil {
		return err
	}
	if err := w.WriteData(o.KeyTag); err != nil {
		return err
	}
	if o.Signer != nil {
		_ptr_nameSigner := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Signer != nil {
				if err := o.Signer.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Signer, _ptr_nameSigner); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	for i1 := range o.SignatureInfo {
		i1 := i1
		if err := w.WriteData(o.SignatureInfo[i1]); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *RecordSIG) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.TypeCovered); err != nil {
		return err
	}
	if err := w.ReadData(&o.Algorithm); err != nil {
		return err
	}
	if err := w.ReadData(&o.LabelCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.OriginalTTL); err != nil {
		return err
	}
	if err := w.ReadData(&o.SigExpiration); err != nil {
		return err
	}
	if err := w.ReadData(&o.SigInception); err != nil {
		return err
	}
	if err := w.ReadData(&o.KeyTag); err != nil {
		return err
	}
	_ptr_nameSigner := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Signer == nil {
			o.Signer = &NodeName{}
		}
		if err := o.Signer.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_nameSigner := func(ptr interface{}) { o.Signer = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.Signer, _s_nameSigner, _ptr_nameSigner); err != nil {
		return err
	}
	for i1 := 0; w.Len() > 0; i1++ {
		i1 := i1
		o.SignatureInfo = append(o.SignatureInfo, uint8(0))
		if err := w.ReadData(&o.SignatureInfo[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// RecordKEY structure represents DNS_RPC_RECORD_KEY RPC structure.
//
// The DNS_RPC_RECORD_KEY structure contains a public key associated with an FQDN as
// specified in section 3 of [RFC2535].<16> This record MUST be formatted as follows:
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| wFlags                                                        | chProtocol                    | chAlgorithm                   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| bKey (variable)                                                                                                               |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type RecordKEY struct {
	// wFlags (2 bytes): Flags value for the key RR as specified in section 3.1 of [RFC2535].
	Flags uint16 `idl:"name:wFlags" json:"flags"`
	// chProtocol (1 byte): Protocol value for the key RR as specified in section 3.1 of
	// [RFC2535].
	Protocol uint8 `idl:"name:chProtocol" json:"protocol"`
	// chAlgorithm (1 byte): Algorithm value for the key RR as specified in section 3.1
	// of [RFC2535].
	Algorithm uint8 `idl:"name:chAlgorithm" json:"algorithm"`
	// bKey (variable): An array of bytes containing the key value as specified of section
	// 3.1 in [RFC2535].
	Key []byte `idl:"name:bKey" json:"key"`
}

func (o *RecordKEY) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordKEY) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(2); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.Protocol); err != nil {
		return err
	}
	if err := w.WriteData(o.Algorithm); err != nil {
		return err
	}
	for i1 := range o.Key {
		i1 := i1
		if err := w.WriteData(o.Key[i1]); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(2); err != nil {
		return err
	}
	return nil
}
func (o *RecordKEY) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(2); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.Protocol); err != nil {
		return err
	}
	if err := w.ReadData(&o.Algorithm); err != nil {
		return err
	}
	for i1 := 0; w.Len() > 0; i1++ {
		i1 := i1
		o.Key = append(o.Key, uint8(0))
		if err := w.ReadData(&o.Key[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(2); err != nil {
		return err
	}
	return nil
}

// RecordDNSKEY structure represents DNS_RPC_RECORD_DNSKEY RPC structure.
//
// The DNS_RPC_RECORD_DNSKEY structure contains a public key associated with an FQDN
// as specified in section 2 of [RFC4034].<18> This record MUST be formatted as follows.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| wFlags                                                        | chProtocol                    | chAlgorithm                   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| bKey (variable)                                                                                                               |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type RecordDNSKEY struct {
	// wFlags (2 bytes): Flags value for the key RR as specified in section 2.1 of [RFC4034].
	Flags uint16 `idl:"name:wFlags" json:"flags"`
	// chProtocol (1 byte): Protocol value for the key RR as specified in section 2.1 of
	// [RFC4034].
	Protocol uint8 `idl:"name:chProtocol" json:"protocol"`
	// chAlgorithm (1 byte): Algorithm value for the key RR as specified in section 2.1
	// of [RFC4034].
	Algorithm uint8 `idl:"name:chAlgorithm" json:"algorithm"`
	// bKey (variable): An array of bytes containing the key value as specified of section
	// 2.1 in [RFC4034].
	Key []byte `idl:"name:bKey" json:"key"`
}

func (o *RecordDNSKEY) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordDNSKEY) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(2); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.Protocol); err != nil {
		return err
	}
	if err := w.WriteData(o.Algorithm); err != nil {
		return err
	}
	for i1 := range o.Key {
		i1 := i1
		if err := w.WriteData(o.Key[i1]); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(2); err != nil {
		return err
	}
	return nil
}
func (o *RecordDNSKEY) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(2); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.Protocol); err != nil {
		return err
	}
	if err := w.ReadData(&o.Algorithm); err != nil {
		return err
	}
	for i1 := 0; w.Len() > 0; i1++ {
		i1 := i1
		o.Key = append(o.Key, uint8(0))
		if err := w.ReadData(&o.Key[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(2); err != nil {
		return err
	}
	return nil
}

// RecordNXT structure represents DNS_RPC_RECORD_NXT RPC structure.
//
// The DNS_RPC_RECORD_NXT specifies a NXT resource record as specified in section 5.1
// of [RFC2535].<19> This record MUST be formatted as follows:
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| wNumRecordTypes                                               | wTypeWords (variable)                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| nextName (variable)                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type RecordNXT struct {
	// wNumRecordTypes (2 bytes): The number of 16-bit unsigned integers in the variable
	// sized wTypeWords array. This value MUST be greater than 1.
	RecordTypesLength uint16 `idl:"name:wNumRecordTypes" json:"record_types_length"`
	// wTypeWords (variable): An array for of 16-bit unsigned integers in little-endian
	// byte order for that contains a variable sized bit-mask value for types present in
	// this record, as specified in section 5.2 of [RFC2535]. The most significant bit of
	// the first integer corresponds to type zero and MUST be zero. If there is a second
	// 16-bit unsigned integer present in the array, the most significant bit of the second
	// integer corresponds to type 16, and so on.
	TypeWords []uint16 `idl:"name:wTypeWords;size_is:(wNumRecordTypes)" json:"type_words"`
	// nextName (variable): A DNS_RPC_NAME (section 2.2.2.2.1) containing next name information,
	// as specified in section 5.2 of [RFC2535].
	NextName *NodeName `idl:"name:nextName" json:"next_name"`
}

func (o *RecordNXT) xxx_PreparePayload(ctx context.Context) error {
	if o.TypeWords != nil && o.RecordTypesLength == 0 {
		o.RecordTypesLength = uint16(len(o.TypeWords))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordNXT) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteData(o.RecordTypesLength); err != nil {
		return err
	}
	if o.TypeWords != nil || o.RecordTypesLength > 0 {
		_ptr_wTypeWords := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.RecordTypesLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.TypeWords {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.TypeWords[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.TypeWords); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.TypeWords, _ptr_wTypeWords); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.NextName != nil {
		_ptr_nextName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.NextName != nil {
				if err := o.NextName.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.NextName, _ptr_nextName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordNXT) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadData(&o.RecordTypesLength); err != nil {
		return err
	}
	_ptr_wTypeWords := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.RecordTypesLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.RecordTypesLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.TypeWords", sizeInfo[0])
		}
		o.TypeWords = make([]uint16, sizeInfo[0])
		for i1 := range o.TypeWords {
			i1 := i1
			if err := w.ReadData(&o.TypeWords[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_wTypeWords := func(ptr interface{}) { o.TypeWords = *ptr.(*[]uint16) }
	if err := w.ReadPointer(&o.TypeWords, _s_wTypeWords, _ptr_wTypeWords); err != nil {
		return err
	}
	_ptr_nextName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.NextName == nil {
			o.NextName = &NodeName{}
		}
		if err := o.NextName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_nextName := func(ptr interface{}) { o.NextName = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.NextName, _s_nextName, _ptr_nextName); err != nil {
		return err
	}
	return nil
}

// RecordATMA structure represents DNS_RPC_RECORD_ATMA RPC structure.
//
// The DNS_RPC_RECORD_ATMA specifies a resource record that contains ATM address information
// as specified in [ATMA]. This record MUST be formatted as follows:
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| chFormat                      | bAddress (variable)                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type RecordATMA struct {
	// chFormat (1 byte): The format of the address as specified in section 5.2 of [ATMA].
	Format uint8 `idl:"name:chFormat" json:"format"`
	// bAddress (variable): The ATM address of the node to which this resource record pertains
	// (see section 5.2 of [ATMA]).
	Address []byte `idl:"name:bAddress" json:"address"`
}

func (o *RecordATMA) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordATMA) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteData(o.Format); err != nil {
		return err
	}
	for i1 := range o.Address {
		i1 := i1
		if err := w.WriteData(o.Address[i1]); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordATMA) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Format); err != nil {
		return err
	}
	for i1 := 0; w.Len() > 0; i1++ {
		i1 := i1
		o.Address = append(o.Address, uint8(0))
		if err := w.ReadData(&o.Address[i1]); err != nil {
			return err
		}
	}
	return nil
}

// RecordNSEC structure represents DNS_RPC_RECORD_NSEC RPC structure.
//
// The DNS_RPC_RECORD_NSEC structure contains the next FQDN in the zone as specified
// in section 4 of [RFC4034].<14> This record MUST be formatted as follows.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| nameSigner (variable)                                                                                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| NSECBitmap (variable)                                                                                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type RecordNSEC struct {
	// nameSigner (variable): A structure of type DNS_RPC_NAME (section 2.2.2.2.1) containing
	// the FQDN of the originating host for this record.
	Signer *NodeName `idl:"name:nameSigner" json:"signer"`
	// NSECBitmap (variable): Bitmap of types present at this node as specified in section
	// 4 of [RFC4034].
	NSECBitmap []uint16 `idl:"name:NSECBitmap" json:"nsec_bitmap"`
}

func (o *RecordNSEC) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordNSEC) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if o.Signer != nil {
		_ptr_nameSigner := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Signer != nil {
				if err := o.Signer.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Signer, _ptr_nameSigner); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	for i1 := range o.NSECBitmap {
		i1 := i1
		if err := w.WriteData(o.NSECBitmap[i1]); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(7); err != nil {
		return err
	}
	return nil
}
func (o *RecordNSEC) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	_ptr_nameSigner := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Signer == nil {
			o.Signer = &NodeName{}
		}
		if err := o.Signer.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_nameSigner := func(ptr interface{}) { o.Signer = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.Signer, _s_nameSigner, _ptr_nameSigner); err != nil {
		return err
	}
	for i1 := 0; w.Len() > 0; i1++ {
		i1 := i1
		o.NSECBitmap = append(o.NSECBitmap, uint16(0))
		if err := w.ReadData(&o.NSECBitmap[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(7); err != nil {
		return err
	}
	return nil
}

// RecordDHCID structure represents DNS_RPC_RECORD_DHCID RPC structure.
//
// The DNS_RPC_RECORD_DHCID structure contains a public key associated with an FQDN
// as specified in section 3 of [RFC2535].<17> This record MUST be formatted as follows.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| bDHCID (variable)                                                                                                             |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type RecordDHCID struct {
	// bDHCID (variable): An opaque DHCID record as specified in section 3 in [RFC4701].
	DHCID []byte `idl:"name:bDHCID" json:"dhcid"`
}

func (o *RecordDHCID) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordDHCID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	for i1 := range o.DHCID {
		i1 := i1
		if err := w.WriteData(o.DHCID[i1]); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordDHCID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	for i1 := 0; w.Len() > 0; i1++ {
		i1 := i1
		o.DHCID = append(o.DHCID, uint8(0))
		if err := w.ReadData(&o.DHCID[i1]); err != nil {
			return err
		}
	}
	return nil
}

// RecordNSEC3 structure represents DNS_RPC_RECORD_NSEC3 RPC structure.
//
// The DNS_RPC_RECORD_NSEC3 structure SHOULD<21> specify an NSEC3 resource record as
// specified in [RFC5155] section 3. This record MUST be formatted as follows.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| chAlgorithm                   | bFlags                        | wIterations                                                   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| bSaltLength                   | bHashLength                   | salt (variable)                                               |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| nextHashedOwnerName (variable)                                                                                                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| bitmaps (variable)                                                                                                            |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type RecordNSEC3 struct {
	// chAlgorithm (1 byte): A value that specifies the cryptographic hash algorithm used
	// to construct the hash value, as specified in [RFC5155] section 3.1.
	Algorithm uint8 `idl:"name:chAlgorithm" json:"algorithm"`
	// bFlags (1 byte): A value that specifies NSEC3 processing flags, as specified in [RFC5155]
	// section 3.1.
	Flags uint8 `idl:"name:bFlags" json:"flags"`
	// wIterations (2 bytes): A value that specifies the number of additional times the
	// hash function has been performed, as specified in [RFC5155] section 3.1.
	Iterations uint16 `idl:"name:wIterations" json:"iterations"`
	// bSaltLength (1 byte): A value that specifies the length of the salt field, in octets.
	SaltLength uint8 `idl:"name:bSaltLength" json:"salt_length"`
	// bHashLength (1 byte): A value that specifies the length of the nextHashedOwnerName
	// field, in octets.
	HashLength uint8 `idl:"name:bHashLength" json:"hash_length"`
	// salt (variable): A value that specifies the salt to be appended to the original owner
	// name before hashing, as specified in [RFC5155] section 3.1.
	Salt []byte `idl:"name:salt;size_is:(bSaltLength)" json:"salt"`
	// nextHashedOwnerName (variable): A value that specifies the next hashed owner name
	// in hash order, as specified in [RFC5155] section 3.1.
	NextHashedOwnerName []byte `idl:"name:nextHashedOwnerName;size_is:(bHashLength)" json:"next_hashed_owner_name"`
	// bitmaps (variable): A value that specifies the DNS types that exist at the original
	// owner name of the NSEC3 record, as specified in [RFC5155] section 3.1.
	Bitmaps []uint16 `idl:"name:bitmaps" json:"bitmaps"`
}

func (o *RecordNSEC3) xxx_PreparePayload(ctx context.Context) error {
	if o.Salt != nil && o.SaltLength == 0 {
		o.SaltLength = uint8(len(o.Salt))
	}
	if o.NextHashedOwnerName != nil && o.HashLength == 0 {
		o.HashLength = uint8(len(o.NextHashedOwnerName))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordNSEC3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteData(o.Algorithm); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.Iterations); err != nil {
		return err
	}
	if err := w.WriteData(o.SaltLength); err != nil {
		return err
	}
	if err := w.WriteData(o.HashLength); err != nil {
		return err
	}
	if o.Salt != nil || o.SaltLength > 0 {
		_ptr_salt := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.SaltLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Salt {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Salt[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Salt); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Salt, _ptr_salt); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.NextHashedOwnerName != nil || o.HashLength > 0 {
		_ptr_nextHashedOwnerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.HashLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.NextHashedOwnerName {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.NextHashedOwnerName[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.NextHashedOwnerName); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.NextHashedOwnerName, _ptr_nextHashedOwnerName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	for i1 := range o.Bitmaps {
		i1 := i1
		if err := w.WriteData(o.Bitmaps[i1]); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(7); err != nil {
		return err
	}
	return nil
}
func (o *RecordNSEC3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadData(&o.Algorithm); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.Iterations); err != nil {
		return err
	}
	if err := w.ReadData(&o.SaltLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.HashLength); err != nil {
		return err
	}
	_ptr_salt := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.SaltLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.SaltLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Salt", sizeInfo[0])
		}
		o.Salt = make([]byte, sizeInfo[0])
		for i1 := range o.Salt {
			i1 := i1
			if err := w.ReadData(&o.Salt[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_salt := func(ptr interface{}) { o.Salt = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Salt, _s_salt, _ptr_salt); err != nil {
		return err
	}
	_ptr_nextHashedOwnerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.HashLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.HashLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.NextHashedOwnerName", sizeInfo[0])
		}
		o.NextHashedOwnerName = make([]byte, sizeInfo[0])
		for i1 := range o.NextHashedOwnerName {
			i1 := i1
			if err := w.ReadData(&o.NextHashedOwnerName[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_nextHashedOwnerName := func(ptr interface{}) { o.NextHashedOwnerName = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.NextHashedOwnerName, _s_nextHashedOwnerName, _ptr_nextHashedOwnerName); err != nil {
		return err
	}
	for i1 := 0; w.Len() > 0; i1++ {
		i1 := i1
		o.Bitmaps = append(o.Bitmaps, uint16(0))
		if err := w.ReadData(&o.Bitmaps[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(7); err != nil {
		return err
	}
	return nil
}

// RecordNSEC3PARAM structure represents DNS_RPC_RECORD_NSEC3PARAM RPC structure.
//
// The DNS_RPC_RECORD_NSEC3PARAM structure SHOULD<22> specify an NSEC3PARAM resource
// record as specified in [RFC5155] section 3. This record MUST be formatted as follows.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| chAlgorithm                   | bFlags                        | wIterations                                                   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| bSaltLength                   | salt (variable)                                                                               |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type RecordNSEC3PARAM struct {
	// chAlgorithm (1 byte): A value that specifies the cryptographic hash algorithm used
	// to construct the hash value, as specified in [RFC5155] section 4.1.
	Algorithm uint8 `idl:"name:chAlgorithm" json:"algorithm"`
	// bFlags (1 byte): A value that specifies NSEC3 processing flags, as specified in [RFC5155]
	// section 3.1.
	Flags uint8 `idl:"name:bFlags" json:"flags"`
	// wIterations (2 bytes): A value that specifies the number of additional times the
	// hash function has been performed, as specified in [RFC5155] section 4.1.
	Iterations uint16 `idl:"name:wIterations" json:"iterations"`
	// bSaltLength (1 byte): A value that specifies the length of the salt field, in octets.
	SaltLength uint8 `idl:"name:bSaltLength" json:"salt_length"`
	// salt (variable): A value that specifies the salt to be appended to the original owner
	// name before hashing, as specified in [RFC5155].
	Salt []byte `idl:"name:salt;size_is:(bSaltLength)" json:"salt"`
}

func (o *RecordNSEC3PARAM) xxx_PreparePayload(ctx context.Context) error {
	if o.Salt != nil && o.SaltLength == 0 {
		o.SaltLength = uint8(len(o.Salt))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordNSEC3PARAM) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteData(o.Algorithm); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.Iterations); err != nil {
		return err
	}
	if err := w.WriteData(o.SaltLength); err != nil {
		return err
	}
	if o.Salt != nil || o.SaltLength > 0 {
		_ptr_salt := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.SaltLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Salt {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Salt[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Salt); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Salt, _ptr_salt); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordNSEC3PARAM) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadData(&o.Algorithm); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.Iterations); err != nil {
		return err
	}
	if err := w.ReadData(&o.SaltLength); err != nil {
		return err
	}
	_ptr_salt := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.SaltLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.SaltLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Salt", sizeInfo[0])
		}
		o.Salt = make([]byte, sizeInfo[0])
		for i1 := range o.Salt {
			i1 := i1
			if err := w.ReadData(&o.Salt[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_salt := func(ptr interface{}) { o.Salt = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Salt, _s_salt, _ptr_salt); err != nil {
		return err
	}
	return nil
}

// RecordTLSA structure represents DNS_RPC_RECORD_TLSA RPC structure.
type RecordTLSA struct {
	CertUsage                  uint8  `idl:"name:bCertUsage" json:"cert_usage"`
	Selector                   uint8  `idl:"name:bSelector" json:"selector"`
	MatchingType               uint8  `idl:"name:bMatchingType" json:"matching_type"`
	CertificateAssociationData []byte `idl:"name:bCertificateAssociationData" json:"certificate_association_data"`
}

func (o *RecordTLSA) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordTLSA) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteData(o.CertUsage); err != nil {
		return err
	}
	if err := w.WriteData(o.Selector); err != nil {
		return err
	}
	if err := w.WriteData(o.MatchingType); err != nil {
		return err
	}
	for i1 := range o.CertificateAssociationData {
		i1 := i1
		if err := w.WriteData(o.CertificateAssociationData[i1]); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordTLSA) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.CertUsage); err != nil {
		return err
	}
	if err := w.ReadData(&o.Selector); err != nil {
		return err
	}
	if err := w.ReadData(&o.MatchingType); err != nil {
		return err
	}
	for i1 := 0; w.Len() > 0; i1++ {
		i1 := i1
		o.CertificateAssociationData = append(o.CertificateAssociationData, uint8(0))
		if err := w.ReadData(&o.CertificateAssociationData[i1]); err != nil {
			return err
		}
	}
	return nil
}

// RecordWINS structure represents DNS_RPC_RECORD_WINS RPC structure.
//
// The DNS_RPC_RECORD_WINS specifies a WINS resource record. This record MUST be formatted
// as follows:
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| dwMappingFlag                                                                                                                 |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| dwLookupTimeout                                                                                                               |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| dwCacheTimeout                                                                                                                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| cWinsServerCount                                                                                                              |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| aipWinsServers (variable)                                                                                                     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type RecordWINS struct {
	// dwMappingFlag (4 bytes): The scope of the WINS record lookups. This value MUST be
	// set to 0x00000000 or any combination of the following:
	//
	//	+--------------------------------+---------------------------------------------------------+
	//	|                                |                                                         |
	//	|             VALUE              |                         MEANING                         |
	//	|                                |                                                         |
	//	+--------------------------------+---------------------------------------------------------+
	//	+--------------------------------+---------------------------------------------------------+
	//	| DNS_WINS_FLAG_SCOPE 0x80000000 | Server forwards lookup requests to remote WINS servers. |
	//	+--------------------------------+---------------------------------------------------------+
	//	| DNS_WINS_FLAG_LOCAL 0x00010000 | Server performs WINS lookups locally.                   |
	//	+--------------------------------+---------------------------------------------------------+
	MappingFlag uint32 `idl:"name:dwMappingFlag" json:"mapping_flag"`
	// dwLookupTimeout (4 bytes): The duration, in seconds, for which the server waits to
	// receive a response from a remote DNS server.
	LookupTimeout uint32 `idl:"name:dwLookupTimeout" json:"lookup_timeout"`
	// dwCacheTimeout (4 bytes): The duration, in seconds, for which the server keeps this
	// record in its cache before considering it stale.
	CacheTimeout uint32 `idl:"name:dwCacheTimeout" json:"cache_timeout"`
	// cWinsServerCount (4 bytes): The number of WINS server addresses in this record. The
	// value of this field MUST be at least one.
	WINSServerCount uint32 `idl:"name:cWinsServerCount" json:"wins_server_count"`
	// aipWinsServers (variable): An array of IPv4 addresses in network byte order with
	// length given by cWinsServerCount.
	WINSServers [][]byte `idl:"name:aipWinsServers;size_is:(cWinsServerCount, )" json:"wins_servers"`
}

func (o *RecordWINS) xxx_PreparePayload(ctx context.Context) error {
	if o.WINSServers != nil && o.WINSServerCount == 0 {
		o.WINSServerCount = uint32(len(o.WINSServers))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordWINS) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.MappingFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.LookupTimeout); err != nil {
		return err
	}
	if err := w.WriteData(o.CacheTimeout); err != nil {
		return err
	}
	if err := w.WriteData(o.WINSServerCount); err != nil {
		return err
	}
	if o.WINSServers != nil || o.WINSServerCount > 0 {
		_ptr_aipWinsServers := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.WINSServerCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.WINSServers {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				for i2 := range o.WINSServers[i1] {
					i2 := i2
					if uint64(i2) >= 4 {
						break
					}
					if err := w.WriteData(o.WINSServers[i1][i2]); err != nil {
						return err
					}
				}
				for i2 := len(o.WINSServers[i1]); uint64(i2) < 4; i2++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.WINSServers); uint64(i1) < sizeInfo[0]; i1++ {
				for i2 := 0; uint64(i2) < 4; i2++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.WINSServers, _ptr_aipWinsServers); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordWINS) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.MappingFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.LookupTimeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.CacheTimeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.WINSServerCount); err != nil {
		return err
	}
	_ptr_aipWinsServers := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.WINSServerCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.WINSServerCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.WINSServers", sizeInfo[0])
		}
		o.WINSServers = make([][]byte, sizeInfo[0])
		for i1 := range o.WINSServers {
			i1 := i1
			o.WINSServers[i1] = make([]byte, 4)
			for i2 := range o.WINSServers[i1] {
				i2 := i2
				if err := w.ReadData(&o.WINSServers[i1][i2]); err != nil {
					return err
				}
			}
		}
		return nil
	})
	_s_aipWinsServers := func(ptr interface{}) { o.WINSServers = *ptr.(*[][]byte) }
	if err := w.ReadPointer(&o.WINSServers, _s_aipWinsServers, _ptr_aipWinsServers); err != nil {
		return err
	}
	return nil
}

// RecordWINSR structure represents DNS_RPC_RECORD_WINSR RPC structure.
//
// The DNS_RPC_RECORD_WINSR specifies a Windows Internet Name Service Reverse Lookup
// (WINS-R) resource record. This record MUST be formatted as follows:
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| dwMappingFlag                                                                                                                 |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| dwLookupTimeout                                                                                                               |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| dwCacheTimeout                                                                                                                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| nameResultDomain (variable)                                                                                                   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type RecordWINSR struct {
	// dwMappingFlag (4 bytes): The scope of the WINS-R record lookups. This value MUST
	// be set to zero or any combination of the following:
	//
	//	+--------------------------------+---------------------------------------------------------+
	//	|                                |                                                         |
	//	|             VALUE              |                         MEANING                         |
	//	|                                |                                                         |
	//	+--------------------------------+---------------------------------------------------------+
	//	+--------------------------------+---------------------------------------------------------+
	//	| DNS_WINS_FLAG_SCOPE 0x80000000 | Server forwards lookup requests to remote WINS servers. |
	//	+--------------------------------+---------------------------------------------------------+
	//	| DNS_WINS_FLAG_LOCAL 0x00010000 | Server performs WINS lookups locally.                   |
	//	+--------------------------------+---------------------------------------------------------+
	MappingFlag uint32 `idl:"name:dwMappingFlag" json:"mapping_flag"`
	// dwLookupTimeout (4 bytes): The duration, in seconds, for which server waits to receive
	// a response from a remote DNS server.
	LookupTimeout uint32 `idl:"name:dwLookupTimeout" json:"lookup_timeout"`
	// dwCacheTimeout (4 bytes): The duration, in seconds, for which server keeps this record
	// in its cache before considering it stale.
	CacheTimeout uint32 `idl:"name:dwCacheTimeout" json:"cache_timeout"`
	// nameResultDomain (variable): Pointer to a structure of type DNS_RPC_NAME (section
	// 2.2.2.2.1) containing a domain name suffix that will be appended to a single-label
	// name [RFC1034] obtained from a WINS-R lookup.
	ResultDomain *NodeName `idl:"name:nameResultDomain" json:"result_domain"`
}

func (o *RecordWINSR) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordWINSR) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.MappingFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.LookupTimeout); err != nil {
		return err
	}
	if err := w.WriteData(o.CacheTimeout); err != nil {
		return err
	}
	if o.ResultDomain != nil {
		_ptr_nameResultDomain := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ResultDomain != nil {
				if err := o.ResultDomain.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ResultDomain, _ptr_nameResultDomain); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordWINSR) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.MappingFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.LookupTimeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.CacheTimeout); err != nil {
		return err
	}
	_ptr_nameResultDomain := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ResultDomain == nil {
			o.ResultDomain = &NodeName{}
		}
		if err := o.ResultDomain.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_nameResultDomain := func(ptr interface{}) { o.ResultDomain = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.ResultDomain, _s_nameResultDomain, _ptr_nameResultDomain); err != nil {
		return err
	}
	return nil
}

// RecordDS structure represents DNS_RPC_RECORD_DS RPC structure.
//
// The DNS_RPC_RECORD_DS structure contains a public key associated with an FQDN as
// specified in section 5 of [RFC4034].<15> This record MUST be formatted as follows.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| wKeyTag                                                       | chAlgorithm                   | chDigestType                  |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| bDigest (variable)                                                                                                            |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type RecordDS struct {
	// wKeyTag (2 bytes): The key tag of the DNSKEY record referred to by this DS record,
	// as specified in section 5 of [RFC4034].
	KeyTag uint16 `idl:"name:wKeyTag" json:"key_tag"`
	// chAlgorithm (1 byte): Algorithm value of the DNSKEY record referred to by this DS
	// record, as specified in section 5 of [RFC4034].
	Algorithm uint8 `idl:"name:chAlgorithm" json:"algorithm"`
	// chDigestType (1 byte): The digest algorithm that was used to construct this DS record,
	// as specified in section 5 of [RFC4034].
	DigestType uint8 `idl:"name:chDigestType" json:"digest_type"`
	// bDigest (variable): An array of bytes containing the digest value as specified of
	// section 5 in [RFC4034].
	Digest []byte `idl:"name:bDigest" json:"digest"`
}

func (o *RecordDS) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordDS) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(2); err != nil {
		return err
	}
	if err := w.WriteData(o.KeyTag); err != nil {
		return err
	}
	if err := w.WriteData(o.Algorithm); err != nil {
		return err
	}
	if err := w.WriteData(o.DigestType); err != nil {
		return err
	}
	for i1 := range o.Digest {
		i1 := i1
		if err := w.WriteData(o.Digest[i1]); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(2); err != nil {
		return err
	}
	return nil
}
func (o *RecordDS) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(2); err != nil {
		return err
	}
	if err := w.ReadData(&o.KeyTag); err != nil {
		return err
	}
	if err := w.ReadData(&o.Algorithm); err != nil {
		return err
	}
	if err := w.ReadData(&o.DigestType); err != nil {
		return err
	}
	for i1 := 0; w.Len() > 0; i1++ {
		i1 := i1
		o.Digest = append(o.Digest, uint8(0))
		if err := w.ReadData(&o.Digest[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(2); err != nil {
		return err
	}
	return nil
}

// RecordSRV structure represents DNS_RPC_RECORD_SRV RPC structure.
//
// The DNS_RPC_RECORD_SRV specifies an SRV resource record as specified in [RFC2782].
// This record MUST be formatted as follows:
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| wPriority                                                     | wWeight                                                       |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| wPort                                                         | nameTarget (variable)                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type RecordSRV struct {
	// wPriority (2 bytes): The priority of the target host as specified in [RFC2782].
	Priority uint16 `idl:"name:wPriority" json:"priority"`
	// wWeight (2 bytes): The relative weight for the target host as specified in [RFC2782].
	Weight uint16 `idl:"name:wWeight" json:"weight"`
	// wPort (2 bytes): The port number for the service on the target host as specified
	// in [RFC2782].
	Port uint16 `idl:"name:wPort" json:"port"`
	// nameTarget (variable): The FDQN of the server that hosts this service in DNS_RPC_NAME
	// (section 2.2.2.2.1) format, as specified in [RFC2782].
	Target *NodeName `idl:"name:nameTarget" json:"target"`
}

func (o *RecordSRV) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordSRV) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteData(o.Priority); err != nil {
		return err
	}
	if err := w.WriteData(o.Weight); err != nil {
		return err
	}
	if err := w.WriteData(o.Port); err != nil {
		return err
	}
	if o.Target != nil {
		_ptr_nameTarget := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Target != nil {
				if err := o.Target.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Target, _ptr_nameTarget); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordSRV) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadData(&o.Priority); err != nil {
		return err
	}
	if err := w.ReadData(&o.Weight); err != nil {
		return err
	}
	if err := w.ReadData(&o.Port); err != nil {
		return err
	}
	_ptr_nameTarget := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Target == nil {
			o.Target = &NodeName{}
		}
		if err := o.Target.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_nameTarget := func(ptr interface{}) { o.Target = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.Target, _s_nameTarget, _ptr_nameTarget); err != nil {
		return err
	}
	return nil
}

// RecordSOA structure represents DNS_RPC_RECORD_SOA RPC structure.
//
// The DNS_RPC_RECORD_SOA structure contains information about an SOA record (section
// 3.3.13 in [RFC1035]). This record MUST be formatted as follows:
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| dwSerialNo                                                                                                                    |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| dwRefresh                                                                                                                     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| dwRetry                                                                                                                       |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| dwExpire                                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| dwMinimumTtl                                                                                                                  |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| namePrimaryServer (variable)                                                                                                  |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Zone Administrator Email (variable)                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type RecordSOA struct {
	// dwSerialNo (4 bytes): The serial number of the SOA record.
	SerialNo uint32 `idl:"name:dwSerialNo" json:"serial_no"`
	// dwRefresh (4 bytes): The interval, in seconds, at which a secondary DNS server attempts
	// to contact the primary DNS server for getting an update.
	Refresh uint32 `idl:"name:dwRefresh" json:"refresh"`
	// dwRetry (4 bytes): The interval, in seconds, at which a secondary DNS server retries
	// to check with the primary DNS server in case of failure.
	Retry uint32 `idl:"name:dwRetry" json:"retry"`
	// dwExpire (4 bytes): The duration, in seconds, that a secondary DNS server continues
	// attempts to get updates from the primary DNS server and if still unsuccessful assumes
	// that the primary DNS server is unreachable.
	Expire uint32 `idl:"name:dwExpire" json:"expire"`
	// dwMinimumTtl (4 bytes): The minimum duration, in seconds, for which record data in
	// the zone is valid.
	MinimumTTL uint32 `idl:"name:dwMinimumTtl" json:"minimum_ttl"`
	// namePrimaryServer (variable): The FQDN of the primary DNS server for this zone in
	// DNS_RPC_NAME (section 2.2.2.2.1) format.
	PrimaryServer          *NodeName `idl:"name:namePrimaryServer" json:"primary_server"`
	ZoneAdministratorEmail *NodeName `idl:"name:nameZoneAdministratorEmail" json:"zone_administrator_email"`
}

func (o *RecordSOA) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordSOA) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.SerialNo); err != nil {
		return err
	}
	if err := w.WriteData(o.Refresh); err != nil {
		return err
	}
	if err := w.WriteData(o.Retry); err != nil {
		return err
	}
	if err := w.WriteData(o.Expire); err != nil {
		return err
	}
	if err := w.WriteData(o.MinimumTTL); err != nil {
		return err
	}
	if o.PrimaryServer != nil {
		_ptr_namePrimaryServer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.PrimaryServer != nil {
				if err := o.PrimaryServer.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.PrimaryServer, _ptr_namePrimaryServer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ZoneAdministratorEmail != nil {
		_ptr_nameZoneAdministratorEmail := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ZoneAdministratorEmail != nil {
				if err := o.ZoneAdministratorEmail.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ZoneAdministratorEmail, _ptr_nameZoneAdministratorEmail); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordSOA) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.SerialNo); err != nil {
		return err
	}
	if err := w.ReadData(&o.Refresh); err != nil {
		return err
	}
	if err := w.ReadData(&o.Retry); err != nil {
		return err
	}
	if err := w.ReadData(&o.Expire); err != nil {
		return err
	}
	if err := w.ReadData(&o.MinimumTTL); err != nil {
		return err
	}
	_ptr_namePrimaryServer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.PrimaryServer == nil {
			o.PrimaryServer = &NodeName{}
		}
		if err := o.PrimaryServer.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_namePrimaryServer := func(ptr interface{}) { o.PrimaryServer = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.PrimaryServer, _s_namePrimaryServer, _ptr_namePrimaryServer); err != nil {
		return err
	}
	_ptr_nameZoneAdministratorEmail := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ZoneAdministratorEmail == nil {
			o.ZoneAdministratorEmail = &NodeName{}
		}
		if err := o.ZoneAdministratorEmail.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_nameZoneAdministratorEmail := func(ptr interface{}) { o.ZoneAdministratorEmail = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.ZoneAdministratorEmail, _s_nameZoneAdministratorEmail, _ptr_nameZoneAdministratorEmail); err != nil {
		return err
	}
	return nil
}

// RecordPreference structure represents DNS_RPC_RECORD_PREFERENCE RPC structure.
type RecordPreference struct {
	Preference uint16    `idl:"name:wPreference" json:"preference"`
	Exchange   *NodeName `idl:"name:nameExchange" json:"exchange"`
}

func (o *RecordPreference) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordPreference) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteData(o.Preference); err != nil {
		return err
	}
	if o.Exchange != nil {
		_ptr_nameExchange := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Exchange != nil {
				if err := o.Exchange.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Exchange, _ptr_nameExchange); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordPreference) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadData(&o.Preference); err != nil {
		return err
	}
	_ptr_nameExchange := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Exchange == nil {
			o.Exchange = &NodeName{}
		}
		if err := o.Exchange.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_nameExchange := func(ptr interface{}) { o.Exchange = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.Exchange, _s_nameExchange, _ptr_nameExchange); err != nil {
		return err
	}
	return nil
}

// RecordMX structure represents DNS_RPC_RECORD_MX RPC structure.
type RecordMX RecordPreference

func (o *RecordMX) RecordPreference() *RecordPreference { return (*RecordPreference)(o) }

func (o *RecordMX) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordMX) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteData(o.Preference); err != nil {
		return err
	}
	if o.Exchange != nil {
		_ptr_nameExchange := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Exchange != nil {
				if err := o.Exchange.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Exchange, _ptr_nameExchange); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordMX) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadData(&o.Preference); err != nil {
		return err
	}
	_ptr_nameExchange := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Exchange == nil {
			o.Exchange = &NodeName{}
		}
		if err := o.Exchange.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_nameExchange := func(ptr interface{}) { o.Exchange = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.Exchange, _s_nameExchange, _ptr_nameExchange); err != nil {
		return err
	}
	return nil
}

// RecordAFSDB structure represents DNS_RPC_RECORD_AFSDB RPC structure.
type RecordAFSDB RecordPreference

func (o *RecordAFSDB) RecordPreference() *RecordPreference { return (*RecordPreference)(o) }

func (o *RecordAFSDB) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordAFSDB) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteData(o.Preference); err != nil {
		return err
	}
	if o.Exchange != nil {
		_ptr_nameExchange := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Exchange != nil {
				if err := o.Exchange.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Exchange, _ptr_nameExchange); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordAFSDB) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadData(&o.Preference); err != nil {
		return err
	}
	_ptr_nameExchange := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Exchange == nil {
			o.Exchange = &NodeName{}
		}
		if err := o.Exchange.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_nameExchange := func(ptr interface{}) { o.Exchange = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.Exchange, _s_nameExchange, _ptr_nameExchange); err != nil {
		return err
	}
	return nil
}

// RecordRT structure represents DNS_RPC_RECORD_RT RPC structure.
type RecordRT RecordPreference

func (o *RecordRT) RecordPreference() *RecordPreference { return (*RecordPreference)(o) }

func (o *RecordRT) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordRT) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteData(o.Preference); err != nil {
		return err
	}
	if o.Exchange != nil {
		_ptr_nameExchange := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Exchange != nil {
				if err := o.Exchange.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Exchange, _ptr_nameExchange); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordRT) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadData(&o.Preference); err != nil {
		return err
	}
	_ptr_nameExchange := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Exchange == nil {
			o.Exchange = &NodeName{}
		}
		if err := o.Exchange.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_nameExchange := func(ptr interface{}) { o.Exchange = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.Exchange, _s_nameExchange, _ptr_nameExchange); err != nil {
		return err
	}
	return nil
}

// RecordNAPTR structure represents DNS_RPC_RECORD_NAPTR RPC structure.
//
// The DNS_RPC_RECORD_NAPTR specifies a NAPTR resource record as specified in section
// 4 of [RFC3403].<20> This record MUST be formatted as follows:
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| wOrder                                                        | wPreference                                                   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| nameFlags (variable)                                                                                                          |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| nameService (variable)                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| nameSubstitution (variable)                                                                                                   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| nameReplacement (variable)                                                                                                    |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type RecordNAPTR struct {
	// wOrder (2 bytes): A value that specifies the order in which the NAPTR record is processed,
	// as specified in section 4.1 in [RFC3403].
	Order uint16 `idl:"name:wOrder" json:"order"`
	// wPreference (2 bytes): The preference given to this NAPTR record, as specified in
	// section 4.1 in [RFC3403].
	Preference uint16 `idl:"name:wPreference" json:"preference"`
	// nameFlags (variable): Pointer to a structure of type DNS_RPC_NAME (section 2.2.2.2.1)
	// containing the string flags value as specified in section 4.1 in [RFC3403].
	Flags *NodeName `idl:"name:nameFlags" json:"flags"`
	// nameService (variable): Pointer to a structure of type DNS_RPC_NAME (section 2.2.2.2.1)
	// containing service parameters value for NAPTR to control the rewriting and interpretation
	// of the field in the record, as specified in section 4.1 in [RFC3403].
	Service *NodeName `idl:"name:nameService" json:"service"`
	// nameSubstitution (variable): Pointer to a structure of type DNS_RPC_NAME (section
	// 2.2.2.2.1) containing a substitution expression value for the NAPTR record, as specified
	// in section 4.1 in [RFC3403].
	Substitution *NodeName `idl:"name:nameSubstitution" json:"substitution"`
	// nameReplacement (variable): Pointer to a structure of type DNS_RPC_NAME (section
	// 2.2.2.2.1) containing a replacement expression value for the NAPTR record, as specified
	// in section 4.1 in [RFC3403].
	Replacement *NodeName `idl:"name:nameReplacement" json:"replacement"`
}

func (o *RecordNAPTR) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordNAPTR) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteData(o.Order); err != nil {
		return err
	}
	if err := w.WriteData(o.Preference); err != nil {
		return err
	}
	if o.Flags != nil {
		_ptr_nameFlags := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Flags != nil {
				if err := o.Flags.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Flags, _ptr_nameFlags); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Service != nil {
		_ptr_nameService := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Service != nil {
				if err := o.Service.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Service, _ptr_nameService); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Substitution != nil {
		_ptr_nameSubstitution := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Substitution != nil {
				if err := o.Substitution.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Substitution, _ptr_nameSubstitution); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Replacement != nil {
		_ptr_nameReplacement := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Replacement != nil {
				if err := o.Replacement.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Replacement, _ptr_nameReplacement); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordNAPTR) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadData(&o.Order); err != nil {
		return err
	}
	if err := w.ReadData(&o.Preference); err != nil {
		return err
	}
	_ptr_nameFlags := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Flags == nil {
			o.Flags = &NodeName{}
		}
		if err := o.Flags.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_nameFlags := func(ptr interface{}) { o.Flags = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.Flags, _s_nameFlags, _ptr_nameFlags); err != nil {
		return err
	}
	_ptr_nameService := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Service == nil {
			o.Service = &NodeName{}
		}
		if err := o.Service.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_nameService := func(ptr interface{}) { o.Service = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.Service, _s_nameService, _ptr_nameService); err != nil {
		return err
	}
	_ptr_nameSubstitution := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Substitution == nil {
			o.Substitution = &NodeName{}
		}
		if err := o.Substitution.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_nameSubstitution := func(ptr interface{}) { o.Substitution = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.Substitution, _s_nameSubstitution, _ptr_nameSubstitution); err != nil {
		return err
	}
	_ptr_nameReplacement := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Replacement == nil {
			o.Replacement = &NodeName{}
		}
		if err := o.Replacement.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_nameReplacement := func(ptr interface{}) { o.Replacement = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.Replacement, _s_nameReplacement, _ptr_nameReplacement); err != nil {
		return err
	}
	return nil
}

// RecordNodeName structure represents DNS_RPC_RECORD_NODE_NAME RPC structure.
//
// The DNS_RPC_RECORD_NODE_NAME structure contains information about a DNS record of
// any of the following types:<11>
//
// * DNS_TYPE_PTR
//
// * DNS_TYPE_NS
//
// * DNS_TYPE_CNAME
//
// * DNS_TYPE_DNAME
//
// * DNS_TYPE_MB
//
// * DNS_TYPE_MR
//
// * DNS_TYPE_MG
//
// * DNS_TYPE_MD
//
// * DNS_TYPE_MF
//
// This record MUST be formatted as follows:
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| nameNode (variable)                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type RecordNodeName struct {
	Name *NodeName `idl:"name:Name" json:"name"`
}

func (o *RecordNodeName) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordNodeName) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.Name != nil {
		_ptr_Name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Name != nil {
				if err := o.Name.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
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
func (o *RecordNodeName) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_Name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Name == nil {
			o.Name = &NodeName{}
		}
		if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Name := func(ptr interface{}) { o.Name = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.Name, _s_Name, _ptr_Name); err != nil {
		return err
	}
	return nil
}

// RecordIPAddress structure represents DNS_RPC_RECORD_IP_ADDRESS RPC structure.
type RecordIPAddress struct {
	IPAddress []byte `idl:"name:pIpAddress" json:"ip_address"`
}

func (o *RecordIPAddress) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordIPAddress) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	for i1 := range o.IPAddress {
		i1 := i1
		if err := w.WriteData(o.IPAddress[i1]); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordIPAddress) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	for i1 := 0; w.Len() > 0; i1++ {
		i1 := i1
		o.IPAddress = append(o.IPAddress, uint8(0))
		if err := w.ReadData(&o.IPAddress[i1]); err != nil {
			return err
		}
	}
	return nil
}

// RecordA structure represents DNS_RPC_RECORD_A RPC structure.
//
// The DNS_RPC_RECORD_A structure contains an IPv4 address. This record MUST be formatted
// as follows:
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| IPv4 Address                                                                                                                  |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type RecordA RecordIPAddress

func (o *RecordA) RecordIPAddress() *RecordIPAddress { return (*RecordIPAddress)(o) }

func (o *RecordA) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordA) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	for i1 := range o.IPAddress {
		i1 := i1
		if err := w.WriteData(o.IPAddress[i1]); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordA) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	for i1 := 0; w.Len() > 0; i1++ {
		i1 := i1
		o.IPAddress = append(o.IPAddress, uint8(0))
		if err := w.ReadData(&o.IPAddress[i1]); err != nil {
			return err
		}
	}
	return nil
}

// RecordAAAA structure represents DNS_RPC_RECORD_AAAA RPC structure.
//
// The DNS_RPC_RECORD_AAAA structure contains IPv6 address information. This record
// MUST be formatted as follows:
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ipv6Address (16 bytes)                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type RecordAAAA RecordIPAddress

func (o *RecordAAAA) RecordIPAddress() *RecordIPAddress { return (*RecordIPAddress)(o) }

func (o *RecordAAAA) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordAAAA) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	for i1 := range o.IPAddress {
		i1 := i1
		if err := w.WriteData(o.IPAddress[i1]); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordAAAA) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	for i1 := 0; w.Len() > 0; i1++ {
		i1 := i1
		o.IPAddress = append(o.IPAddress, uint8(0))
		if err := w.ReadData(&o.IPAddress[i1]); err != nil {
			return err
		}
	}
	return nil
}

// RecordPTR structure represents DNS_RPC_RECORD_PTR RPC structure.
type RecordPTR RecordNodeName

func (o *RecordPTR) RecordNodeName() *RecordNodeName { return (*RecordNodeName)(o) }

func (o *RecordPTR) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordPTR) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.Name != nil {
		_ptr_Name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Name != nil {
				if err := o.Name.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
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
func (o *RecordPTR) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_Name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Name == nil {
			o.Name = &NodeName{}
		}
		if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Name := func(ptr interface{}) { o.Name = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.Name, _s_Name, _ptr_Name); err != nil {
		return err
	}
	return nil
}

// RecordNS structure represents DNS_RPC_RECORD_NS RPC structure.
type RecordNS RecordNodeName

func (o *RecordNS) RecordNodeName() *RecordNodeName { return (*RecordNodeName)(o) }

func (o *RecordNS) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordNS) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.Name != nil {
		_ptr_Name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Name != nil {
				if err := o.Name.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
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
func (o *RecordNS) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_Name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Name == nil {
			o.Name = &NodeName{}
		}
		if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Name := func(ptr interface{}) { o.Name = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.Name, _s_Name, _ptr_Name); err != nil {
		return err
	}
	return nil
}

// RecordCNAME structure represents DNS_RPC_RECORD_CNAME RPC structure.
type RecordCNAME RecordNodeName

func (o *RecordCNAME) RecordNodeName() *RecordNodeName { return (*RecordNodeName)(o) }

func (o *RecordCNAME) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordCNAME) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.Name != nil {
		_ptr_Name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Name != nil {
				if err := o.Name.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
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
func (o *RecordCNAME) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_Name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Name == nil {
			o.Name = &NodeName{}
		}
		if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Name := func(ptr interface{}) { o.Name = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.Name, _s_Name, _ptr_Name); err != nil {
		return err
	}
	return nil
}

// RecordDNAME structure represents DNS_RPC_RECORD_DNAME RPC structure.
type RecordDNAME RecordNodeName

func (o *RecordDNAME) RecordNodeName() *RecordNodeName { return (*RecordNodeName)(o) }

func (o *RecordDNAME) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordDNAME) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.Name != nil {
		_ptr_Name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Name != nil {
				if err := o.Name.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
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
func (o *RecordDNAME) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_Name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Name == nil {
			o.Name = &NodeName{}
		}
		if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Name := func(ptr interface{}) { o.Name = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.Name, _s_Name, _ptr_Name); err != nil {
		return err
	}
	return nil
}

// RecordMB structure represents DNS_RPC_RECORD_MB RPC structure.
type RecordMB RecordNodeName

func (o *RecordMB) RecordNodeName() *RecordNodeName { return (*RecordNodeName)(o) }

func (o *RecordMB) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordMB) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.Name != nil {
		_ptr_Name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Name != nil {
				if err := o.Name.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
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
func (o *RecordMB) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_Name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Name == nil {
			o.Name = &NodeName{}
		}
		if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Name := func(ptr interface{}) { o.Name = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.Name, _s_Name, _ptr_Name); err != nil {
		return err
	}
	return nil
}

// RecordMR structure represents DNS_RPC_RECORD_MR RPC structure.
type RecordMR RecordNodeName

func (o *RecordMR) RecordNodeName() *RecordNodeName { return (*RecordNodeName)(o) }

func (o *RecordMR) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordMR) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.Name != nil {
		_ptr_Name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Name != nil {
				if err := o.Name.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
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
func (o *RecordMR) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_Name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Name == nil {
			o.Name = &NodeName{}
		}
		if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Name := func(ptr interface{}) { o.Name = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.Name, _s_Name, _ptr_Name); err != nil {
		return err
	}
	return nil
}

// RecordMG structure represents DNS_RPC_RECORD_MG RPC structure.
type RecordMG RecordNodeName

func (o *RecordMG) RecordNodeName() *RecordNodeName { return (*RecordNodeName)(o) }

func (o *RecordMG) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordMG) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.Name != nil {
		_ptr_Name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Name != nil {
				if err := o.Name.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
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
func (o *RecordMG) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_Name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Name == nil {
			o.Name = &NodeName{}
		}
		if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Name := func(ptr interface{}) { o.Name = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.Name, _s_Name, _ptr_Name); err != nil {
		return err
	}
	return nil
}

// RecordMD structure represents DNS_RPC_RECORD_MD RPC structure.
type RecordMD RecordNodeName

func (o *RecordMD) RecordNodeName() *RecordNodeName { return (*RecordNodeName)(o) }

func (o *RecordMD) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordMD) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.Name != nil {
		_ptr_Name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Name != nil {
				if err := o.Name.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
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
func (o *RecordMD) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_Name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Name == nil {
			o.Name = &NodeName{}
		}
		if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Name := func(ptr interface{}) { o.Name = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.Name, _s_Name, _ptr_Name); err != nil {
		return err
	}
	return nil
}

// RecordMF structure represents DNS_RPC_RECORD_MF RPC structure.
type RecordMF RecordNodeName

func (o *RecordMF) RecordNodeName() *RecordNodeName { return (*RecordNodeName)(o) }

func (o *RecordMF) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordMF) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.Name != nil {
		_ptr_Name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Name != nil {
				if err := o.Name.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeName{}).MarshalNDR(ctx, w); err != nil {
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
func (o *RecordMF) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_Name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Name == nil {
			o.Name = &NodeName{}
		}
		if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Name := func(ptr interface{}) { o.Name = *ptr.(**NodeName) }
	if err := w.ReadPointer(&o.Name, _s_Name, _ptr_Name); err != nil {
		return err
	}
	return nil
}

// RecordString structure represents DNS_RPC_RECORD_STRING RPC structure.
//
// The DNS_RPC_RECORD_STRING structure contains information about a DNS record of any
// of the following types:
//
// * DNS_TYPE_HINFO
//
// * DNS_TYPE_ISDN
//
// * DNS_TYPE_TXT
//
// * DNS_TYPE_X25
//
// * DNS_TYPE_LOC
//
// This packet contains one or more instances of stringData, depending upon the type
// listed above. This record MUST be formatted as follows:
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| stringData (variable)                                                                                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type RecordString struct {
	Data []*NodeText `idl:"name:Data" json:"data"`
}

func (o *RecordString) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordString) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if o.Data[i1] != nil {
			_ptr_Data := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Data[i1] != nil {
					if err := o.Data[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&NodeText{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Data[i1], _ptr_Data); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	}
	return nil
}
func (o *RecordString) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	for i1 := 0; w.Len() > 0; i1++ {
		i1 := i1
		o.Data = append(o.Data, nil)
		_ptr_Data := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Data[i1] == nil {
				o.Data[i1] = &NodeText{}
			}
			if err := o.Data[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Data := func(ptr interface{}) { o.Data[i1] = *ptr.(**NodeText) }
		if err := w.ReadPointer(&o.Data[i1], _s_Data, _ptr_Data); err != nil {
			return err
		}
	}
	return nil
}

// RecordTXT structure represents DNS_RPC_RECORD_TXT RPC structure.
type RecordTXT RecordString

func (o *RecordTXT) RecordString() *RecordString { return (*RecordString)(o) }

func (o *RecordTXT) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordTXT) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if o.Data[i1] != nil {
			_ptr_Data := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Data[i1] != nil {
					if err := o.Data[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&NodeText{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Data[i1], _ptr_Data); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	}
	return nil
}
func (o *RecordTXT) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	for i1 := 0; w.Len() > 0; i1++ {
		i1 := i1
		o.Data = append(o.Data, nil)
		_ptr_Data := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Data[i1] == nil {
				o.Data[i1] = &NodeText{}
			}
			if err := o.Data[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Data := func(ptr interface{}) { o.Data[i1] = *ptr.(**NodeText) }
		if err := w.ReadPointer(&o.Data[i1], _s_Data, _ptr_Data); err != nil {
			return err
		}
	}
	return nil
}

// RecordISDN structure represents DNS_RPC_RECORD_ISDN RPC structure.
type RecordISDN RecordString

func (o *RecordISDN) RecordString() *RecordString { return (*RecordString)(o) }

func (o *RecordISDN) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordISDN) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if o.Data[i1] != nil {
			_ptr_Data := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Data[i1] != nil {
					if err := o.Data[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&NodeText{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Data[i1], _ptr_Data); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	}
	return nil
}
func (o *RecordISDN) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	for i1 := 0; w.Len() > 0; i1++ {
		i1 := i1
		o.Data = append(o.Data, nil)
		_ptr_Data := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Data[i1] == nil {
				o.Data[i1] = &NodeText{}
			}
			if err := o.Data[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Data := func(ptr interface{}) { o.Data[i1] = *ptr.(**NodeText) }
		if err := w.ReadPointer(&o.Data[i1], _s_Data, _ptr_Data); err != nil {
			return err
		}
	}
	return nil
}

// RecordHINFO structure represents DNS_RPC_RECORD_HINFO RPC structure.
type RecordHINFO RecordString

func (o *RecordHINFO) RecordString() *RecordString { return (*RecordString)(o) }

func (o *RecordHINFO) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordHINFO) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if o.Data[i1] != nil {
			_ptr_Data := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Data[i1] != nil {
					if err := o.Data[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&NodeText{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Data[i1], _ptr_Data); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	}
	return nil
}
func (o *RecordHINFO) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	for i1 := 0; w.Len() > 0; i1++ {
		i1 := i1
		o.Data = append(o.Data, nil)
		_ptr_Data := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Data[i1] == nil {
				o.Data[i1] = &NodeText{}
			}
			if err := o.Data[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Data := func(ptr interface{}) { o.Data[i1] = *ptr.(**NodeText) }
		if err := w.ReadPointer(&o.Data[i1], _s_Data, _ptr_Data); err != nil {
			return err
		}
	}
	return nil
}

// RecordX25 structure represents DNS_RPC_RECORD_X25 RPC structure.
type RecordX25 RecordString

func (o *RecordX25) RecordString() *RecordString { return (*RecordString)(o) }

func (o *RecordX25) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordX25) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if o.Data[i1] != nil {
			_ptr_Data := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Data[i1] != nil {
					if err := o.Data[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&NodeText{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Data[i1], _ptr_Data); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	}
	return nil
}
func (o *RecordX25) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	for i1 := 0; w.Len() > 0; i1++ {
		i1 := i1
		o.Data = append(o.Data, nil)
		_ptr_Data := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Data[i1] == nil {
				o.Data[i1] = &NodeText{}
			}
			if err := o.Data[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Data := func(ptr interface{}) { o.Data[i1] = *ptr.(**NodeText) }
		if err := w.ReadPointer(&o.Data[i1], _s_Data, _ptr_Data); err != nil {
			return err
		}
	}
	return nil
}

// RecordLoc structure represents DNS_RPC_RECORD_LOC RPC structure.
type RecordLoc RecordString

func (o *RecordLoc) RecordString() *RecordString { return (*RecordString)(o) }

func (o *RecordLoc) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordLoc) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if o.Data[i1] != nil {
			_ptr_Data := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Data[i1] != nil {
					if err := o.Data[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&NodeText{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Data[i1], _ptr_Data); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	}
	return nil
}
func (o *RecordLoc) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	for i1 := 0; w.Len() > 0; i1++ {
		i1 := i1
		o.Data = append(o.Data, nil)
		_ptr_Data := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Data[i1] == nil {
				o.Data[i1] = &NodeText{}
			}
			if err := o.Data[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Data := func(ptr interface{}) { o.Data[i1] = *ptr.(**NodeText) }
		if err := w.ReadPointer(&o.Data[i1], _s_Data, _ptr_Data); err != nil {
			return err
		}
	}
	return nil
}

// RecordWKS structure represents DNS_RPC_RECORD_WKS RPC structure.
//
// The DNS_RPC_RECORD_WKS structure contains the information for the well known services
// supported by a host, as defined in section 3.4.2 [RFC1035]. This record MUST be formatted
// as follows:
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ipAddress                                                                                                                     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| chProtocol                    | bBitMask (variable)                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type RecordWKS struct {
	// ipAddress (4 bytes): The IPv4 address of the server that provides the service.
	IPAddress []byte `idl:"name:ipAddress" json:"ip_address"`
	// chProtocol (1 byte): The IP protocol number as specified in [IANA-PROTO-NUM].
	Protocol uint8     `idl:"name:chProtocol" json:"protocol"`
	Services *NodeText `idl:"name:nameServices" json:"services"`
}

func (o *RecordWKS) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordWKS) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	for i1 := range o.IPAddress {
		i1 := i1
		if uint64(i1) >= 4 {
			break
		}
		if err := w.WriteData(o.IPAddress[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.IPAddress); uint64(i1) < 4; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Protocol); err != nil {
		return err
	}
	if o.Services != nil {
		_ptr_nameServices := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Services != nil {
				if err := o.Services.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NodeText{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Services, _ptr_nameServices); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordWKS) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	o.IPAddress = make([]byte, 4)
	for i1 := range o.IPAddress {
		i1 := i1
		if err := w.ReadData(&o.IPAddress[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.Protocol); err != nil {
		return err
	}
	_ptr_nameServices := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Services == nil {
			o.Services = &NodeText{}
		}
		if err := o.Services.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_nameServices := func(ptr interface{}) { o.Services = *ptr.(**NodeText) }
	if err := w.ReadPointer(&o.Services, _s_nameServices, _ptr_nameServices); err != nil {
		return err
	}
	return nil
}
