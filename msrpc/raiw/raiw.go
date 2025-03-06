// The raiw package implements the RAIW client protocol.
//
// # Introduction
//
// This is a specification of the Remote Administrative Interface: WINS protocol. This
// protocol defines remote procedure call (RPC) interfaces that provide methods for
// remotely accessing and administering a server for the Windows Internet Name Service
// (WINS). This protocol is a client/server protocol that is based on RPC and is used
// in the configuration, management, and monitoring of a WINS server.
//
// An application implementing this protocol can remotely perform service monitoring
// of a WINS server as well as creating, updating, querying, or deleting database records,
// performing database scavenging, and replicating the database records with other WINS
// servers.
//
// # Overview
//
// The Remote Administrative Interface: WINS protocol is a client/server protocol that
// is used to remotely configure, manage, and monitor the WINS server. This protocol
// allows a client to view and update the server configuration settings as well as to
// create, modify, and delete WINS database records. It also allows clients to trigger
// scavenging and replicating operations and to query the database.
//
// The Remote Administrative Interface: WINS protocol is stateless with no state shared
// across RPC method calls. Each RPC method call contains one complete request. Output
// from one method call can be used as an input to another call, but the protocol does
// not provide methods for locking the WINS server con figuration or state data across
// method calls. For example, a client can pull a range of records from the database
// and delete some of them using another RPC call. However, the protocol does not guarantee
// that the specified record has not been modified by another client between the two
// method calls.
//
//	+---------------------------------------+
//	|   REMOTE ADMINISTRATIVE INTERFACE:    |
//	|                 WINS                  |
//	+---------------------------------------+
//	+---------------------------------------+
//	| Remote Procedure Call (RPC)           |
//	+---------------------------------------+
//	|
//	+---------------------------------------+
package raiw

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
	GoPackage = "raiw"
)

// MaxNoReplicationPullPartners represents the WINSINTF_MAX_NO_RPL_PNRS RPC constant
var MaxNoReplicationPullPartners = 25

// VersNo structure represents WINSINTF_VERS_NO_T RPC structure.
type VersNo dtyp.LargeInteger

func (o *VersNo) LargeInteger() *dtyp.LargeInteger { return (*dtyp.LargeInteger)(o) }

func (o *VersNo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *VersNo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.QuadPart); err != nil {
		return err
	}
	return nil
}
func (o *VersNo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.QuadPart); err != nil {
		return err
	}
	return nil
}

// Addr structure represents WINSINTF_ADDR_T RPC structure.
type Addr struct {
	Type   uint8  `idl:"name:Type" json:"type"`
	Length uint32 `idl:"name:Len" json:"length"`
	IPAddr uint32 `idl:"name:IPAddr" json:"ip_addr"`
}

func (o *Addr) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
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
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if err := w.WriteData(o.IPAddr); err != nil {
		return err
	}
	return nil
}
func (o *Addr) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPAddr); err != nil {
		return err
	}
	return nil
}

// PriorityClass type represents WINSINTF_PRIORITY_CLASS_E RPC enumeration.
//
// The WINSINTF_PRIORITY_CLASS_E enumeration defines the priority class of a WINS process.
// It is used by the RPC method R_WinsSetPriorityClass.
type PriorityClass uint16

var (
	// WINSINTF_E_NORMAL:  WINS process is assigned normal priority class.
	PriorityClassNormal PriorityClass = 0
	// WINSINTF_E_HIGH:  WINS process is assigned high priority class.
	PriorityClassHigh PriorityClass = 1
)

func (o PriorityClass) String() string {
	switch o {
	case PriorityClassNormal:
		return "PriorityClassNormal"
	case PriorityClassHigh:
		return "PriorityClassHigh"
	}
	return "Invalid"
}

// Action type represents WINSINTF_ACT_E RPC enumeration.
//
// The WINSINTF_ACT_E enumeration indicates an action type requested by the RPC method
// R_WinsRecordAction for a record contained in the WINSINTF_RECORD_ACTION_T structure.
type Action uint16

var (
	// WINSINTF_E_INSERT:  Insert a record into the WINS database.
	ActionInsert Action = 0
	// WINSINTF_E_DELETE:  Delete a matching record from the WINS database.
	ActionDelete Action = 1
	// WINSINTF_E_RELEASE:  Release a matching record from the WINS database.
	ActionRelease Action = 2
	// WINSINTF_E_MODIFY:  Modify the attributes of the matching record.
	ActionModify Action = 3
	// WINSINTF_E_QUERY:  Query the database for a given name.
	ActionQuery Action = 4
)

func (o Action) String() string {
	switch o {
	case ActionInsert:
		return "ActionInsert"
	case ActionDelete:
		return "ActionDelete"
	case ActionRelease:
		return "ActionRelease"
	case ActionModify:
		return "ActionModify"
	case ActionQuery:
		return "ActionQuery"
	}
	return "Invalid"
}

// TriggerTypeE type represents WINSINTF_TRIG_TYPE_E RPC enumeration.
//
// The WINSINTF_TRIG_TYPE_E enumeration defines the type of replication to be done.
// It is used by the RPC method R_WinsTrigger.
type TriggerTypeE uint16

var (
	// WINSINTF_E_PULL:  The target WINS server performs pull replication with the specified
	// WINS server.
	TriggerTypeEPull TriggerTypeE = 0
	// WINSINTF_E_PUSH:  The target WINS server performs push replication with the specified
	// WINS server.
	TriggerTypeEPush TriggerTypeE = 1
	// WINSINTF_E_PUSH_PROP:  The target WINS server performs propagating push replication
	// with the specified WINS server.
	TriggerTypeEPushProperty TriggerTypeE = 2
)

func (o TriggerTypeE) String() string {
	switch o {
	case TriggerTypeEPull:
		return "TriggerTypeEPull"
	case TriggerTypeEPush:
		return "TriggerTypeEPush"
	case TriggerTypeEPushProperty:
		return "TriggerTypeEPushProperty"
	}
	return "Invalid"
}

// RecordAction structure represents WINSINTF_RECORD_ACTION_T RPC structure.
//
// The WINSINTF_RECORD_ACTION_T structure defines a WINS database record and the action
// to be performed on it. The structure WINSINTF_RECS_T (section 2.2.2.8) and the RPC
// method R_WinsRecordAction (section 3.1.4.1) both use this structure.
type RecordAction struct {
	// Cmd_e:  A WINSINTF_ACT_E enumeration (section 2.2.1.4) value that specifies the action
	// to be performed on the specified record.
	Cmd Action `idl:"name:Cmd_e" json:"cmd"`
	// pName:  A pointer to a NULL-terminated string that contains the NetBIOS name and
	// optionally the NetBIOS scope name of the record. The NetBIOS scope name, if present,
	// is appended to the NetBIOS name with a dot character ".".
	Name []byte `idl:"name:pName;size_is:((NameLen+1))" json:"name"`
	// NameLen:  The length of the string that pName points to. It has the following possible
	// values:
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	|         16 | The pName value points to a string that contains only the NetBIOS name of the    |
	//	|            | record. The NameLen value does not include the terminating NULL character.       |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 18 < value | The pName value points to a string that contains the NetBIOS name, a dot         |
	//	|            | character ".", and the NULL-terminated NetBIOS scope name of the record. The     |
	//	|            | NameLen value includes the terminating NULL character. If the NameLen value is   |
	//	|            | greater than 255, the pName string SHOULD be truncated to 254 characters plus a  |
	//	|            | terminating NULL character.                                                      |
	//	+------------+----------------------------------------------------------------------------------+
	NameLength uint32 `idl:"name:NameLen" json:"name_length"`
	// TypOfRec_e:  The record type. Only the two least-significant bits of the member value
	// are considered valid. All other bits are masked with zero. The following values are
	// allowed.
	//
	//	+-------+-------------------------+
	//	|       |                         |
	//	| VALUE |         MEANING         |
	//	|       |                         |
	//	+-------+-------------------------+
	//	+-------+-------------------------+
	//	|     0 | Unique name             |
	//	+-------+-------------------------+
	//	|     1 | Normal group name       |
	//	+-------+-------------------------+
	//	|     2 | Special group name      |
	//	+-------+-------------------------+
	//	|     3 | Multihomed machine name |
	//	+-------+-------------------------+
	TypeOfRecord  uint32  `idl:"name:TypOfRec_e" json:"type_of_record"`
	NumberOfAddrs uint32  `idl:"name:NoOfAddrs" json:"number_of_addrs"`
	Addrs         []*Addr `idl:"name:pAddrs;size_is:(NoOfAddrs);pointer:unique" json:"addrs"`
	Addr          *Addr   `idl:"name:Addr" json:"addr"`
	// VersNo:  The version number of the record.
	VersNo *dtyp.LargeInteger `idl:"name:VersNo" json:"vers_no"`
	// NodeTyp:  The NetBT node type. Only the two least-significant bits of the member
	// value are considered valid. All other bits are masked with zero. This member MUST
	// have one of the following values:
	//
	//	+-------+---------+
	//	|       |         |
	//	| VALUE | MEANING |
	//	|       |         |
	//	+-------+---------+
	//	+-------+---------+
	//	|     0 | B-node  |
	//	+-------+---------+
	//	|     1 | P-node  |
	//	+-------+---------+
	//	|     2 | M-node  |
	//	+-------+---------+
	//	|     3 | H-node  |
	//	+-------+---------+
	NodeType uint8 `idl:"name:NodeTyp" json:"node_type"`
	// OwnerId:  The owner IP address of the record, in little-endian byte order.
	OwnerID uint32 `idl:"name:OwnerId" json:"owner_id"`
	// State_e:  The state of the record. Only the two least-significant bits of the member
	// value are considered valid. All other bits are masked with zero. This member MUST
	// have one of the following values:
	//
	//	+-------+-------------------+
	//	|       |                   |
	//	| VALUE |      MEANING      |
	//	|       |                   |
	//	+-------+-------------------+
	//	+-------+-------------------+
	//	|     0 | Active record     |
	//	+-------+-------------------+
	//	|     1 | Released record   |
	//	+-------+-------------------+
	//	|     2 | Tombstoned record |
	//	+-------+-------------------+
	//	|     3 | Deleted record    |
	//	+-------+-------------------+
	State uint32 `idl:"name:State_e" json:"state"`
	// fStatic:  A value that indicates whether the record is static or dynamic. A value
	// of 0 indicates a dynamic record, and 1 indicates a static record. Only the least-significant
	// bit is considered valid. All other bits are masked with zero.
	Static uint32 `idl:"name:fStatic" json:"static"`
	// TimeStamp:  The time stamp [ISO-8601] of the record.
	Timestamp uint64 `idl:"name:TimeStamp" json:"timestamp"`
}

func (o *RecordAction) xxx_PreparePayload(ctx context.Context) error {
	if o.Name != nil && o.NameLength == 0 {
		o.NameLength = uint32((len(o.Name) - 1))
		if len(o.Name) < 1 {
			o.NameLength = uint32(0)
		}
	}
	if o.Addrs != nil && o.NumberOfAddrs == 0 {
		o.NumberOfAddrs = uint32(len(o.Addrs))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordAction) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Cmd)); err != nil {
		return err
	}
	if o.Name != nil || (o.NameLength+1) > 0 {
		_ptr_pName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64((o.NameLength + 1))
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Name {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Name[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Name); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_pName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.NameLength); err != nil {
		return err
	}
	if err := w.WriteData(o.TypeOfRecord); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfAddrs); err != nil {
		return err
	}
	if o.Addrs != nil || o.NumberOfAddrs > 0 {
		_ptr_pAddrs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.NumberOfAddrs)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Addrs {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Addrs[i1] != nil {
					if err := o.Addrs[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&Addr{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Addrs); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&Addr{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Addrs, _ptr_pAddrs); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Addr != nil {
		if err := o.Addr.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Addr{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.VersNo != nil {
		if err := o.VersNo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.NodeType); err != nil {
		return err
	}
	if err := w.WriteData(o.OwnerID); err != nil {
		return err
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.Static); err != nil {
		return err
	}
	if err := w.WriteData(ndr.Uint3264(o.Timestamp)); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *RecordAction) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Cmd)); err != nil {
		return err
	}
	_ptr_pName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if (o.NameLength+1) > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64((o.NameLength + 1))
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Name", sizeInfo[0])
		}
		o.Name = make([]byte, sizeInfo[0])
		for i1 := range o.Name {
			i1 := i1
			if err := w.ReadData(&o.Name[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pName := func(ptr interface{}) { o.Name = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Name, _s_pName, _ptr_pName); err != nil {
		return err
	}
	if err := w.ReadData(&o.NameLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.TypeOfRecord); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfAddrs); err != nil {
		return err
	}
	_ptr_pAddrs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.NumberOfAddrs > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.NumberOfAddrs)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Addrs", sizeInfo[0])
		}
		o.Addrs = make([]*Addr, sizeInfo[0])
		for i1 := range o.Addrs {
			i1 := i1
			if o.Addrs[i1] == nil {
				o.Addrs[i1] = &Addr{}
			}
			if err := o.Addrs[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pAddrs := func(ptr interface{}) { o.Addrs = *ptr.(*[]*Addr) }
	if err := w.ReadPointer(&o.Addrs, _s_pAddrs, _ptr_pAddrs); err != nil {
		return err
	}
	if o.Addr == nil {
		o.Addr = &Addr{}
	}
	if err := o.Addr.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.VersNo == nil {
		o.VersNo = &dtyp.LargeInteger{}
	}
	if err := o.VersNo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.NodeType); err != nil {
		return err
	}
	if err := w.ReadData(&o.OwnerID); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if err := w.ReadData(&o.Static); err != nil {
		return err
	}
	if err := w.ReadData((*ndr.Uint3264)(&o.Timestamp)); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// ReplicationCounters structure represents WINSINTF_RPL_COUNTERS_T RPC structure.
//
// The WINSINTF_RPL_COUNTERS_T structure defines counters that contain the number of
// successful pull replications and the number of communication failures for a given
// replication partner. It is used in the structure WINSINTF_STAT_T.
type ReplicationCounters struct {
	Addr *Addr `idl:"name:Addr" json:"addr"`
	// NoOfRpls:  The number of successful pull replications that have been performed with
	// the replication partner. The target WINS server stores the replication partner's
	// IP address in the Add member.
	NumberOfRpls uint32 `idl:"name:NoOfRpls" json:"number_of_rpls"`
	// NoOfCommFails:  The number of communication failures that have occurred in pull replications
	// between the WINS server whose IP address is given in Add and the target WINS server.
	NumberOfCommFails uint32 `idl:"name:NoOfCommFails" json:"number_of_comm_fails"`
}

func (o *ReplicationCounters) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReplicationCounters) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.Addr != nil {
		if err := o.Addr.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Addr{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.NumberOfRpls); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfCommFails); err != nil {
		return err
	}
	return nil
}
func (o *ReplicationCounters) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.Addr == nil {
		o.Addr = &Addr{}
	}
	if err := o.Addr.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfRpls); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfCommFails); err != nil {
		return err
	}
	return nil
}

// Stat structure represents WINSINTF_STAT_T RPC structure.
//
// The WINSINTF_STAT_T structure defines counters, configured timestamps, the pull replication
// statistics for a given WINS server. This structure is used by the structure WINSINTF_RESULTS_T
// (section 2.2.2.7).
type Stat struct {
	// Counters:  A structure that contains 32-bit unsigned integer counters, which measure
	// various statistics on a WINS server.
	Counters *Stat_Counters `idl:"name:Counters" json:"counters"`
	// TimeStamps:   A structure that contains data in SYSTEMTIME structures ([MS-DTYP]
	// section 2.3.13), which reflect the local time zone of the target WINS server.
	Timestamps *Stat_Timestamps `idl:"name:TimeStamps" json:"timestamps"`
	// NoOfPnrs:  The number of pull partners configured for the target WINS server.
	NumberOfPullPartners uint32 `idl:"name:NoOfPnrs" json:"number_of_pull_partners"`
	// pRplPnrs:  A pointer to structures that contain the details of successful and failed
	// replication counters of configured pull partners at the target WINS server, since
	// the time service was started; or, the time at which the last reset happened by a
	// call to the method R_WinsResetCounters (section 3.1.4.12). The number of structures
	// is specified by NoOfPnrs.
	PeriodicReplicationPullPartners []*ReplicationCounters `idl:"name:pRplPnrs;size_is:(NoOfPnrs);pointer:unique" json:"periodic_replication_pull_partners"`
}

func (o *Stat) xxx_PreparePayload(ctx context.Context) error {
	if o.PeriodicReplicationPullPartners != nil && o.NumberOfPullPartners == 0 {
		o.NumberOfPullPartners = uint32(len(o.PeriodicReplicationPullPartners))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
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
	if o.Counters != nil {
		if err := o.Counters.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Stat_Counters{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Timestamps != nil {
		if err := o.Timestamps.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Stat_Timestamps{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.NumberOfPullPartners); err != nil {
		return err
	}
	if o.PeriodicReplicationPullPartners != nil || o.NumberOfPullPartners > 0 {
		_ptr_pRplPnrs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.NumberOfPullPartners)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.PeriodicReplicationPullPartners {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.PeriodicReplicationPullPartners[i1] != nil {
					if err := o.PeriodicReplicationPullPartners[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ReplicationCounters{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.PeriodicReplicationPullPartners); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&ReplicationCounters{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.PeriodicReplicationPullPartners, _ptr_pRplPnrs); err != nil {
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
	if o.Counters == nil {
		o.Counters = &Stat_Counters{}
	}
	if err := o.Counters.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Timestamps == nil {
		o.Timestamps = &Stat_Timestamps{}
	}
	if err := o.Timestamps.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfPullPartners); err != nil {
		return err
	}
	_ptr_pRplPnrs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.NumberOfPullPartners > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.NumberOfPullPartners)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.PeriodicReplicationPullPartners", sizeInfo[0])
		}
		o.PeriodicReplicationPullPartners = make([]*ReplicationCounters, sizeInfo[0])
		for i1 := range o.PeriodicReplicationPullPartners {
			i1 := i1
			if o.PeriodicReplicationPullPartners[i1] == nil {
				o.PeriodicReplicationPullPartners[i1] = &ReplicationCounters{}
			}
			if err := o.PeriodicReplicationPullPartners[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pRplPnrs := func(ptr interface{}) { o.PeriodicReplicationPullPartners = *ptr.(*[]*ReplicationCounters) }
	if err := w.ReadPointer(&o.PeriodicReplicationPullPartners, _s_pRplPnrs, _ptr_pRplPnrs); err != nil {
		return err
	}
	return nil
}

// Stat_Counters structure represents WINSINTF_STAT_T structure anonymous member.
//
// The WINSINTF_STAT_T structure defines counters, configured timestamps, the pull replication
// statistics for a given WINS server. This structure is used by the structure WINSINTF_RESULTS_T
// (section 2.2.2.7).
type Stat_Counters struct {
	// NoOfUniqueReg:  The number of unique registrations on the target WINS server since
	// the service was started.
	NumberOfUniqueReg uint32 `idl:"name:NoOfUniqueReg" json:"number_of_unique_reg"`
	// NoOfGroupReg:  The number of group registrations at the target WINS server since
	// the service was started.
	NumberOfGroupReg uint32 `idl:"name:NoOfGroupReg" json:"number_of_group_reg"`
	// NoOfQueries:  The number of queries that clients have performed on the target WINS
	// server to resolve NetBIOS names since the service was started. This value is the
	// sum of the values maintained in NoOfSuccQueries and NoOfFailQueries.
	NumberOfQueries uint32 `idl:"name:NoOfQueries" json:"number_of_queries"`
	// NoOfSuccQueries:  The number of successful name resolution queries on the target
	// WINS server since the service was started.
	NumberOfSuccessQueries uint32 `idl:"name:NoOfSuccQueries" json:"number_of_success_queries"`
	// NoOfFailQueries:  The number of failed name resolution queries on the target WINS
	// server since the service was started.
	NumberOfFailQueries uint32 `idl:"name:NoOfFailQueries" json:"number_of_fail_queries"`
	// NoOfUniqueRef:  The number of unique name refreshes on the target WINS server since
	// the service was started.
	NumberOfUniqueReference uint32 `idl:"name:NoOfUniqueRef" json:"number_of_unique_reference"`
	// NoOfGroupRef:  The number of group name refreshes on the target WINS server since
	// the service was started.
	NumberOfGroupReference uint32 `idl:"name:NoOfGroupRef" json:"number_of_group_reference"`
	// NoOfRel:  The number of name releases on the target WINS server since the service
	// was started. This value is the sum of the values maintained in NoOfSuccRel and NoOfFailRel.
	NumberOfRelation uint32 `idl:"name:NoOfRel" json:"number_of_relation"`
	// NoOfSuccRel:  The number of successful name releases on the target WINS server since
	// the service was started.
	NumberOfSuccessRelation uint32 `idl:"name:NoOfSuccRel" json:"number_of_success_relation"`
	// NoOfFailRel:  The number of failed name releases on the target WINS server since
	// the service was started.
	NumberOfFailRelation uint32 `idl:"name:NoOfFailRel" json:"number_of_fail_relation"`
	// NoOfUniqueCnf:  The number of unique name conflicts on the target WINS server since
	// the service was started. Unique name conflicts can occur in the following cases:
	//
	// § The server is registering or refreshing unique name requests from clients.
	NumberOfUniqueConflict uint32 `idl:"name:NoOfUniqueCnf" json:"number_of_unique_conflict"`
	// NoOfGroupCnf:  The number of group name conflicts on the target WINS server since
	// the service was started. Group name conflicts can occur in the following cases:
	//
	// § The server is registering or refreshing unique name requests from clients.
	NumberOfGroupConflict uint32 `idl:"name:NoOfGroupCnf" json:"number_of_group_conflict"`
}

func (o *Stat_Counters) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Stat_Counters) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfUniqueReg); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfGroupReg); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfQueries); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfSuccessQueries); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfFailQueries); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfUniqueReference); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfGroupReference); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfRelation); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfSuccessRelation); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfFailRelation); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfUniqueConflict); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfGroupConflict); err != nil {
		return err
	}
	return nil
}
func (o *Stat_Counters) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfUniqueReg); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfGroupReg); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfQueries); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfSuccessQueries); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfFailQueries); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfUniqueReference); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfGroupReference); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfRelation); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfSuccessRelation); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfFailRelation); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfUniqueConflict); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfGroupConflict); err != nil {
		return err
	}
	return nil
}

// Stat_Timestamps structure represents WINSINTF_STAT_T structure anonymous member.
//
// The WINSINTF_STAT_T structure defines counters, configured timestamps, the pull replication
// statistics for a given WINS server. This structure is used by the structure WINSINTF_RESULTS_T
// (section 2.2.2.7).
type Stat_Timestamps struct {
	// WINSStartTime:  The time at which the WINS service was started on the target WINS
	// server.
	WINSStartTime *dtyp.SystemTime `idl:"name:WINSStartTime" json:"wins_start_time"`
	// LastPScvTime:  The time at which the last periodic scavenging operation was done
	// on the target WINS server.
	LastPeriodicScavengingTime *dtyp.SystemTime `idl:"name:LastPScvTime" json:"last_periodic_scavenging_time"`
	// LastATScvTime:  The time at which the last administrator-triggered scavenging operation
	// was done on the target WINS server.
	LastATScavengingTime *dtyp.SystemTime `idl:"name:LastATScvTime" json:"last_at_scavenging_time"`
	// LastTombScvTime:  The time at which the last scavenging operation was done for the
	// replicated tombstone records on the target WINS server.
	LastTombScavengingTime *dtyp.SystemTime `idl:"name:LastTombScvTime" json:"last_tomb_scavenging_time"`
	// LastVerifyScvTime:  The time at which the last verification scavenging operation
	// was done for the replicated active records on the target WINS server.
	LastVerifyScavengingTime *dtyp.SystemTime `idl:"name:LastVerifyScvTime" json:"last_verify_scavenging_time"`
	// LastPRplTime:  The time at which the last periodic pull replication was done on the
	// target WINS server.
	LastPeriodicReplicationTime *dtyp.SystemTime `idl:"name:LastPRplTime" json:"last_periodic_replication_time"`
	// LastATRplTime:  The time at which the last administrator-triggered pull replication
	// was done on the target WINS server.
	LastATReplicationTime *dtyp.SystemTime `idl:"name:LastATRplTime" json:"last_at_replication_time"`
	// LastNTRplTime:  This member is not set and MUST be ignored on receipt.
	LastNTReplicationTime *dtyp.SystemTime `idl:"name:LastNTRplTime" json:"last_nt_replication_time"`
	// LastACTRplTime:  This member is not set and MUST be ignored on receipt.
	LastActReplicationTime *dtyp.SystemTime `idl:"name:LastACTRplTime" json:"last_act_replication_time"`
	// LastInitDbTime:  The time at which the last static database initialization was done
	// on the target WINS server.
	LastInitDBTime *dtyp.SystemTime `idl:"name:LastInitDbTime" json:"last_init_db_time"`
	// CounterResetTime:  The last time at which the administrator has cleared the success
	// and failure replication counters of the target WINS server.
	CounterResetTime *dtyp.SystemTime `idl:"name:CounterResetTime" json:"counter_reset_time"`
}

func (o *Stat_Timestamps) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Stat_Timestamps) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(2); err != nil {
		return err
	}
	if o.WINSStartTime != nil {
		if err := o.WINSStartTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.SystemTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LastPeriodicScavengingTime != nil {
		if err := o.LastPeriodicScavengingTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.SystemTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LastATScavengingTime != nil {
		if err := o.LastATScavengingTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.SystemTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LastTombScavengingTime != nil {
		if err := o.LastTombScavengingTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.SystemTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LastVerifyScavengingTime != nil {
		if err := o.LastVerifyScavengingTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.SystemTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LastPeriodicReplicationTime != nil {
		if err := o.LastPeriodicReplicationTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.SystemTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LastATReplicationTime != nil {
		if err := o.LastATReplicationTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.SystemTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LastNTReplicationTime != nil {
		if err := o.LastNTReplicationTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.SystemTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LastActReplicationTime != nil {
		if err := o.LastActReplicationTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.SystemTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LastInitDBTime != nil {
		if err := o.LastInitDBTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.SystemTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.CounterResetTime != nil {
		if err := o.CounterResetTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.SystemTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Stat_Timestamps) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(2); err != nil {
		return err
	}
	if o.WINSStartTime == nil {
		o.WINSStartTime = &dtyp.SystemTime{}
	}
	if err := o.WINSStartTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LastPeriodicScavengingTime == nil {
		o.LastPeriodicScavengingTime = &dtyp.SystemTime{}
	}
	if err := o.LastPeriodicScavengingTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LastATScavengingTime == nil {
		o.LastATScavengingTime = &dtyp.SystemTime{}
	}
	if err := o.LastATScavengingTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LastTombScavengingTime == nil {
		o.LastTombScavengingTime = &dtyp.SystemTime{}
	}
	if err := o.LastTombScavengingTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LastVerifyScavengingTime == nil {
		o.LastVerifyScavengingTime = &dtyp.SystemTime{}
	}
	if err := o.LastVerifyScavengingTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LastPeriodicReplicationTime == nil {
		o.LastPeriodicReplicationTime = &dtyp.SystemTime{}
	}
	if err := o.LastPeriodicReplicationTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LastATReplicationTime == nil {
		o.LastATReplicationTime = &dtyp.SystemTime{}
	}
	if err := o.LastATReplicationTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LastNTReplicationTime == nil {
		o.LastNTReplicationTime = &dtyp.SystemTime{}
	}
	if err := o.LastNTReplicationTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LastActReplicationTime == nil {
		o.LastActReplicationTime = &dtyp.SystemTime{}
	}
	if err := o.LastActReplicationTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LastInitDBTime == nil {
		o.LastInitDBTime = &dtyp.SystemTime{}
	}
	if err := o.LastInitDBTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.CounterResetTime == nil {
		o.CounterResetTime = &dtyp.SystemTime{}
	}
	if err := o.CounterResetTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// AddrVersMap structure represents WINSINTF_ADDR_VERS_MAP_T RPC structure.
type AddrVersMap struct {
	Addr   *Addr              `idl:"name:Addr" json:"addr"`
	VersNo *dtyp.LargeInteger `idl:"name:VersNo" json:"vers_no"`
}

func (o *AddrVersMap) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AddrVersMap) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.Addr != nil {
		if err := o.Addr.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Addr{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.VersNo != nil {
		if err := o.VersNo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *AddrVersMap) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.Addr == nil {
		o.Addr = &Addr{}
	}
	if err := o.Addr.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.VersNo == nil {
		o.VersNo = &dtyp.LargeInteger{}
	}
	if err := o.VersNo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Results structure represents WINSINTF_RESULTS_T RPC structure.
//
// The WINSINTF_RESULTS_T structure defines information related to the configuration
// and statistics of a target WINS server. This is used by RPC method R_WinsStatus.
type Results struct {
	// NoOfOwners:  The number of owners whose records are part of the target WINS server
	// database. The value of this member MUST be less than or equal to 25.
	NumberOfOwners uint32         `idl:"name:NoOfOwners" json:"number_of_owners"`
	AddrVersMaps   []*AddrVersMap `idl:"name:AddrVersMaps" json:"addr_vers_maps"`
	// MyMaxVersNo:  This member is not set and MUST be ignored on receipt.
	MyMaxVersNo *dtyp.LargeInteger `idl:"name:MyMaxVersNo" json:"my_max_vers_no"`
	// RefreshInterval:  The refresh time interval configured on the target WINS server,
	// in seconds.
	RefreshInterval uint32 `idl:"name:RefreshInterval" json:"refresh_interval"`
	// TombstoneInterval:  The tombstone interval configured on the target WINS server,
	// in seconds.
	TombstoneInterval uint32 `idl:"name:TombstoneInterval" json:"tombstone_interval"`
	// TombstoneTimeout:  The tombstone timeout configured on the target WINS server, in
	// seconds.
	TombstoneTimeout uint32 `idl:"name:TombstoneTimeout" json:"tombstone_timeout"`
	// VerifyInterval:  The verify time interval configured on the target WINS server, in
	// seconds.
	VerifyInterval uint32 `idl:"name:VerifyInterval" json:"verify_interval"`
	// WINSPriorityClass:  The priority class of the WINS process running on the target
	// WINS server. It SHOULD<2> have one of the following values:
	//
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	|                                  |                                                                                  |
	//	|              VALUE               |                                     MEANING                                      |
	//	|                                  |                                                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| NORMAL_PRIORITY_CLASS 0x00000020 | The process has no special scheduling requirements.                              |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| HIGH_PRIORITY_CLASS 0x00000080   | The process performs time-critical tasks that MUST be executed immediately       |
	//	|                                  | for the process to run correctly. The threads of a high-priority class process   |
	//	|                                  | preempt the threads of normal-priority class processes.                          |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	WINSPriorityClass uint32 `idl:"name:WINSPriorityClass" json:"wins_priority_class"`
	// NoOfWorkerThds:  The number of threads created in the WINS process for serving the
	// NetBIOS name requests.
	NumberOfWorkerThreads uint32 `idl:"name:NoOfWorkerThds" json:"number_of_worker_threads"`
	// WINSStat:  A WINSINTF_STAT_T structure (section 2.2.2.6) containing timing parameters
	// configured on the target WINS server and the pull replication statistics of partner
	// WINS servers.
	WINSStat *Stat `idl:"name:WINSStat" json:"wins_stat"`
}

func (o *Results) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Results) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfOwners); err != nil {
		return err
	}
	for i1 := range o.AddrVersMaps {
		i1 := i1
		if uint64(i1) >= 25 {
			break
		}
		if o.AddrVersMaps[i1] != nil {
			if err := o.AddrVersMaps[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&AddrVersMap{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.AddrVersMaps); uint64(i1) < 25; i1++ {
		if err := (&AddrVersMap{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.MyMaxVersNo != nil {
		if err := o.MyMaxVersNo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.RefreshInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.TombstoneInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.TombstoneTimeout); err != nil {
		return err
	}
	if err := w.WriteData(o.VerifyInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.WINSPriorityClass); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfWorkerThreads); err != nil {
		return err
	}
	if o.WINSStat != nil {
		if err := o.WINSStat.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Stat{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *Results) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfOwners); err != nil {
		return err
	}
	o.AddrVersMaps = make([]*AddrVersMap, 25)
	for i1 := range o.AddrVersMaps {
		i1 := i1
		if o.AddrVersMaps[i1] == nil {
			o.AddrVersMaps[i1] = &AddrVersMap{}
		}
		if err := o.AddrVersMaps[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.MyMaxVersNo == nil {
		o.MyMaxVersNo = &dtyp.LargeInteger{}
	}
	if err := o.MyMaxVersNo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.RefreshInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.TombstoneInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.TombstoneTimeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.VerifyInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.WINSPriorityClass); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfWorkerThreads); err != nil {
		return err
	}
	if o.WINSStat == nil {
		o.WINSStat = &Stat{}
	}
	if err := o.WINSStat.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// ResultsNew structure represents WINSINTF_RESULTS_NEW_T RPC structure.
//
// The WINSINTF_RESULTS_NEW_T structure defines configuration information and statistics
// for a target WINS server. This structure is used by the RPC method R_WinsStatusNew
// (section 3.1.4.20).
type ResultsNew struct {
	// NoOfOwners:  The number of owners whose records are part of the target WINS server
	// database.
	NumberOfOwners uint32         `idl:"name:NoOfOwners" json:"number_of_owners"`
	AddrVersMaps   []*AddrVersMap `idl:"name:pAddrVersMaps;size_is:(NoOfOwners);pointer:unique" json:"addr_vers_maps"`
	// MyMaxVersNo:  This member is not set and MUST be ignored on receipt.
	MyMaxVersNo *dtyp.LargeInteger `idl:"name:MyMaxVersNo" json:"my_max_vers_no"`
	// RefreshInterval:  The refresh time interval configured on the target WINS server,
	// in seconds.
	RefreshInterval uint32 `idl:"name:RefreshInterval" json:"refresh_interval"`
	// TombstoneInterval:  The tombstone time interval configured on the target WINS server,
	// in seconds.
	TombstoneInterval uint32 `idl:"name:TombstoneInterval" json:"tombstone_interval"`
	// TombstoneTimeout:  The tombstone timeout configured on the target WINS server, in
	// seconds.
	TombstoneTimeout uint32 `idl:"name:TombstoneTimeout" json:"tombstone_timeout"`
	// VerifyInterval:  The verify time interval configured on the target WINS server, in
	// seconds.
	VerifyInterval uint32 `idl:"name:VerifyInterval" json:"verify_interval"`
	// WINSPriorityClass:  The priority class of the WINS process running on the target
	// WINS server. It can have one of the following values:
	//
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	|                                  |                                                                                  |
	//	|              VALUE               |                                     MEANING                                      |
	//	|                                  |                                                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| NORMAL_PRIORITY_CLASS 0x00000020 | The process has no special scheduling requirements.                              |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| HIGH_PRIORITY_CLASS 0x00000080   | The process performs time-critical tasks that MUST be executed immediately       |
	//	|                                  | for the process to run correctly. The threads of a high-priority class process   |
	//	|                                  | preempt the threads of normal-priority class processes.                          |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	WINSPriorityClass uint32 `idl:"name:WINSPriorityClass" json:"wins_priority_class"`
	// NoOfWorkerThds:  The number of threads created in the WINS process to serve NetBIOS
	// name requests.
	NumberOfWorkerThreads uint32 `idl:"name:NoOfWorkerThds" json:"number_of_worker_threads"`
	// WINSStat:  A WINSINTF_STAT_T structure (section 2.2.2.6) containing timing parameters
	// configured on the target WINS server and pull replication statistics of partner WINS
	// servers.
	WINSStat *Stat `idl:"name:WINSStat" json:"wins_stat"`
}

func (o *ResultsNew) xxx_PreparePayload(ctx context.Context) error {
	if o.AddrVersMaps != nil && o.NumberOfOwners == 0 {
		o.NumberOfOwners = uint32(len(o.AddrVersMaps))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ResultsNew) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfOwners); err != nil {
		return err
	}
	if o.AddrVersMaps != nil || o.NumberOfOwners > 0 {
		_ptr_pAddrVersMaps := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.NumberOfOwners)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.AddrVersMaps {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.AddrVersMaps[i1] != nil {
					if err := o.AddrVersMaps[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&AddrVersMap{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.AddrVersMaps); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&AddrVersMap{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.AddrVersMaps, _ptr_pAddrVersMaps); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.MyMaxVersNo != nil {
		if err := o.MyMaxVersNo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.RefreshInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.TombstoneInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.TombstoneTimeout); err != nil {
		return err
	}
	if err := w.WriteData(o.VerifyInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.WINSPriorityClass); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfWorkerThreads); err != nil {
		return err
	}
	if o.WINSStat != nil {
		if err := o.WINSStat.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Stat{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *ResultsNew) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfOwners); err != nil {
		return err
	}
	_ptr_pAddrVersMaps := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.NumberOfOwners > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.NumberOfOwners)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.AddrVersMaps", sizeInfo[0])
		}
		o.AddrVersMaps = make([]*AddrVersMap, sizeInfo[0])
		for i1 := range o.AddrVersMaps {
			i1 := i1
			if o.AddrVersMaps[i1] == nil {
				o.AddrVersMaps[i1] = &AddrVersMap{}
			}
			if err := o.AddrVersMaps[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pAddrVersMaps := func(ptr interface{}) { o.AddrVersMaps = *ptr.(*[]*AddrVersMap) }
	if err := w.ReadPointer(&o.AddrVersMaps, _s_pAddrVersMaps, _ptr_pAddrVersMaps); err != nil {
		return err
	}
	if o.MyMaxVersNo == nil {
		o.MyMaxVersNo = &dtyp.LargeInteger{}
	}
	if err := o.MyMaxVersNo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.RefreshInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.TombstoneInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.TombstoneTimeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.VerifyInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.WINSPriorityClass); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfWorkerThreads); err != nil {
		return err
	}
	if o.WINSStat == nil {
		o.WINSStat = &Stat{}
	}
	if err := o.WINSStat.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// Cmd type represents WINSINTF_CMD_E RPC enumeration.
//
// The WINSINTF_CMD_E enumeration is used by the RPC methods to retrieve the configuration
// of a particular WINS server. This enumeration is used in conjunction with the WINSINTF_RESULTS_T
// and WINSINTF_RESULTS_NEW_T structures.
type Cmd uint16

var (
	// WINSINTF_E_ADDVERSMAP:  Gets an entry from the owner version map of the target WINS
	// server.
	CmdAddVersMap Cmd = 0
	// WINSINTF_E_CONFIG:  Get the configuration details of the target WINS server.
	CmdConfig Cmd = 1
	// WINSINTF_E_STAT:   Get statistics for the target WINS server.
	CmdStat Cmd = 2
	// WINSINTF_E_CONFIG_ALL_MAPS:  Get all owner version map entries from the target WINS
	// server.
	CmdConfigAllMaps Cmd = 3
)

func (o Cmd) String() string {
	switch o {
	case CmdAddVersMap:
		return "CmdAddVersMap"
	case CmdConfig:
		return "CmdConfig"
	case CmdStat:
		return "CmdStat"
	case CmdConfigAllMaps:
		return "CmdConfigAllMaps"
	}
	return "Invalid"
}

// Records structure represents WINSINTF_RECS_T RPC structure.
//
// The structure WINSINTF_RECS_T defines an array of WINSINTF_RECORD_ACTION_T (section
// 2.2.2.3) elements. The R_WinsGetDbRecs (section 3.1.4.6) and R_WinsGetDbRecsByName
// (section 3.1.4.19) methods use this structure.
type Records struct {
	// BuffSize:  The number of bytes allocated for the pointer pRow.
	BufferSize uint32 `idl:"name:BuffSize" json:"buffer_size"`
	// pRow:  A pointer to an array of WINSINTF_RECORD_ACTION_T elements.
	Row []*RecordAction `idl:"name:pRow;size_is:(NoOfRecs);pointer:unique" json:"row"`
	// NoOfRecs:  The number of records stored in the array pointed to by pRow.
	NumberOfRecords uint32 `idl:"name:NoOfRecs" json:"number_of_records"`
	// TotalNoOfRecs:  This member is not set and MUST be ignored on receipt.
	TotalNumberOfRecords uint32 `idl:"name:TotalNoOfRecs" json:"total_number_of_records"`
}

func (o *Records) xxx_PreparePayload(ctx context.Context) error {
	if o.Row != nil && o.NumberOfRecords == 0 {
		o.NumberOfRecords = uint32(len(o.Row))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Records) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.BufferSize); err != nil {
		return err
	}
	if o.Row != nil || o.NumberOfRecords > 0 {
		_ptr_pRow := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.NumberOfRecords)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Row {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Row[i1] != nil {
					if err := o.Row[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&RecordAction{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Row); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&RecordAction{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Row, _ptr_pRow); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.NumberOfRecords); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalNumberOfRecords); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *Records) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.BufferSize); err != nil {
		return err
	}
	_ptr_pRow := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.NumberOfRecords > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.NumberOfRecords)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Row", sizeInfo[0])
		}
		o.Row = make([]*RecordAction, sizeInfo[0])
		for i1 := range o.Row {
			i1 := i1
			if o.Row[i1] == nil {
				o.Row[i1] = &RecordAction{}
			}
			if err := o.Row[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pRow := func(ptr interface{}) { o.Row = *ptr.(*[]*RecordAction) }
	if err := w.ReadPointer(&o.Row, _s_pRow, _ptr_pRow); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfRecords); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalNumberOfRecords); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// PullRangeInfo structure represents WINSINTF_PULL_RANGE_INFO_T RPC structure.
type PullRangeInfo struct {
	PullPartner []byte  `idl:"name:pPnr" json:"pull_partner"`
	OwnAddr     *Addr   `idl:"name:OwnAddr" json:"own_addr"`
	MinVersNo   *VersNo `idl:"name:MinVersNo" json:"min_vers_no"`
	MaxVersNo   *VersNo `idl:"name:MaxVersNo" json:"max_vers_no"`
}

func (o *PullRangeInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PullRangeInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.PullPartner != nil {
		_ptr_pPnr := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			// FIXME unknown type pPnr
			return nil
		})
		if err := w.WritePointer(&o.PullPartner, _ptr_pPnr); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.OwnAddr != nil {
		if err := o.OwnAddr.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Addr{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.MinVersNo != nil {
		if err := o.MinVersNo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&VersNo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.MaxVersNo != nil {
		if err := o.MaxVersNo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&VersNo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PullRangeInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	_ptr_pPnr := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		// FIXME: unknown type pPnr
		return nil
	})
	_s_pPnr := func(ptr interface{}) { o.PullPartner = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.PullPartner, _s_pPnr, _ptr_pPnr); err != nil {
		return err
	}
	if o.OwnAddr == nil {
		o.OwnAddr = &Addr{}
	}
	if err := o.OwnAddr.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.MinVersNo == nil {
		o.MinVersNo = &VersNo{}
	}
	if err := o.MinVersNo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.MaxVersNo == nil {
		o.MaxVersNo = &VersNo{}
	}
	if err := o.MaxVersNo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// BrowserInfo structure represents WINSINTF_BROWSER_INFO_T RPC structure.
//
// The WINSINTF_BROWSER_INFO_T structure defines information about browser names. It
// is used by the structure WINSINTF_BROWSER_NAMES_T.
type BrowserInfo struct {
	// dwNameLen:  The length of the name that pName points to, in bytes. The length includes
	// the terminating NULL character.
	NameLength uint32 `idl:"name:dwNameLen" json:"name_length"`
	// pName:  A pointer to a NULL-terminated string that contains the browser name.
	Name string `idl:"name:pName;string" json:"name"`
}

func (o *BrowserInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *BrowserInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.NameLength); err != nil {
		return err
	}
	if o.Name != "" {
		_ptr_pName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_pName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *BrowserInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.NameLength); err != nil {
		return err
	}
	_ptr_pName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_pName := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_pName, _ptr_pName); err != nil {
		return err
	}
	return nil
}

// BrowserNames structure represents WINSINTF_BROWSER_NAMES_T RPC structure.
//
// The WINSINTF_BROWSER_NAMES_T structure defines an array of browser names. This structure
// is used by the RPC method R_WinsGetBrowserNames.
type BrowserNames struct {
	// EntriesRead:  The number of entries in the array that pInfo points to.
	EntriesRead uint32 `idl:"name:EntriesRead" json:"entries_read"`
	// pInfo:  A pointer to an array of browser names. EntriesRead contains the length of
	// this array.
	Info []*BrowserInfo `idl:"name:pInfo;size_is:(EntriesRead);pointer:unique" json:"info"`
}

func (o *BrowserNames) xxx_PreparePayload(ctx context.Context) error {
	if o.Info != nil && o.EntriesRead == 0 {
		o.EntriesRead = uint32(len(o.Info))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *BrowserNames) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesRead); err != nil {
		return err
	}
	if o.Info != nil || o.EntriesRead > 0 {
		_ptr_pInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesRead)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Info {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Info[i1] != nil {
					if err := o.Info[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&BrowserInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Info); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&BrowserInfo{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Info, _ptr_pInfo); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *BrowserNames) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesRead); err != nil {
		return err
	}
	_ptr_pInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
			return fmt.Errorf("buffer overflow for size %d of array o.Info", sizeInfo[0])
		}
		o.Info = make([]*BrowserInfo, sizeInfo[0])
		for i1 := range o.Info {
			i1 := i1
			if o.Info[i1] == nil {
				o.Info[i1] = &BrowserInfo{}
			}
			if err := o.Info[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pInfo := func(ptr interface{}) { o.Info = *ptr.(*[]*BrowserInfo) }
	if err := w.ReadPointer(&o.Info, _s_pInfo, _ptr_pInfo); err != nil {
		return err
	}
	return nil
}

// ScavengingType type represents WINSINTF_SCV_OPC_E RPC enumeration.
//
// The WINSINTF_SCV_OPC_E enumeration specifies the type of scavenging to be done on
// the target WINS server. This enumeration is used in the structure WINSINTF_SCV_REQ_T.
type ScavengingType uint16

var (
	// WINSINTF_E_SCV_GENERAL:  Requests normal scavenging operation.
	ScavengingTypeGeneral ScavengingType = 0
	// WINSINTF_E_SCV_VERIFY:  Verifies only the replicated active records with their owner
	// NBNS servers for their validity.
	ScavengingTypeVerify ScavengingType = 1
)

func (o ScavengingType) String() string {
	switch o {
	case ScavengingTypeGeneral:
		return "ScavengingTypeGeneral"
	case ScavengingTypeVerify:
		return "ScavengingTypeVerify"
	}
	return "Invalid"
}

// ScavengingRequest structure represents WINSINTF_SCV_REQ_T RPC structure.
//
// The WINSINTF_SCV_REQ_T structure defines the type of scavenging that needs to be
// done on the target WINS server. This is used by the RPC method R_WinsDoScavengingNew
// (section 3.1.4.22).
type ScavengingRequest struct {
	// Opcode_e:  A WINSINTF_SCV_OPC_E enumeration (section 2.2.1.8) value describing the
	// type of scavenging operation to be performed on the target WINS server.
	OpcodeE ScavengingType `idl:"name:Opcode_e" json:"opcode_e"`
	// Age:  This member is not set and MUST be ignored on receipt.
	Age uint32 `idl:"name:Age" json:"age"`
	// fForce:  Specifies whether a forceful scavenging is required.
	//
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	|                         |                                                                                  |
	//	|          VALUE          |                                     MEANING                                      |
	//	|                         |                                                                                  |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000              | The internal state and configuration of the WINS server determine whether        |
	//	|                         | scavenging is performed.                                                         |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000001 — 0xFFFFFFFF | The target WINS server performs scavenging.                                      |
	//	+-------------------------+----------------------------------------------------------------------------------+
	Force uint32 `idl:"name:fForce" json:"force"`
}

func (o *ScavengingRequest) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ScavengingRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.OpcodeE)); err != nil {
		return err
	}
	if err := w.WriteData(o.Age); err != nil {
		return err
	}
	if err := w.WriteData(o.Force); err != nil {
		return err
	}
	return nil
}
func (o *ScavengingRequest) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.OpcodeE)); err != nil {
		return err
	}
	if err := w.ReadData(&o.Age); err != nil {
		return err
	}
	if err := w.ReadData(&o.Force); err != nil {
		return err
	}
	return nil
}

// BindData structure represents WINSINTF_BIND_DATA_T RPC structure.
//
// The WINSINTF_BIND_DATA_T structure defines the binding information of the WINS server
// to which the client connects.
type BindData struct {
	// fTcpIp:  The transport mechanism to be used. If this value is 0x00000001, then TCP/IP
	// is selected; otherwise, the named pipe is selected.
	TCPIP      uint32 `idl:"name:fTcpIp" json:"tcp_ip"`
	ServerAddr string `idl:"name:pServerAddr;string" json:"server_addr"`
	// pPipeName:  A NULL-terminated string that specifies the pipe name. This value MUST
	// be NULL when fTcpIP is 0x00000001.
	PipeName string `idl:"name:pPipeName;string" json:"pipe_name"`
}

func (o *BindData) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *BindData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.TCPIP); err != nil {
		return err
	}
	if o.ServerAddr != "" {
		_ptr_pServerAddr := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.ServerAddr); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ServerAddr, _ptr_pServerAddr); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.PipeName != "" {
		_ptr_pPipeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.PipeName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.PipeName, _ptr_pPipeName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *BindData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.TCPIP); err != nil {
		return err
	}
	_ptr_pServerAddr := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.ServerAddr); err != nil {
			return err
		}
		return nil
	})
	_s_pServerAddr := func(ptr interface{}) { o.ServerAddr = *ptr.(*string) }
	if err := w.ReadPointer(&o.ServerAddr, _s_pServerAddr, _ptr_pServerAddr); err != nil {
		return err
	}
	_ptr_pPipeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.PipeName); err != nil {
			return err
		}
		return nil
	})
	_s_pPipeName := func(ptr interface{}) { o.PipeName = *ptr.(*string) }
	if err := w.ReadPointer(&o.PipeName, _s_pPipeName, _ptr_pPipeName); err != nil {
		return err
	}
	return nil
}
