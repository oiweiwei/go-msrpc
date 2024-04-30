// The mqds package implements the MQDS client protocol.
//
// # Introduction
//
// This document specifies the Message Queuing (MSMQ): Directory Service Protocol, a
// remote procedure call (RPC)-based protocol that is used by a client to remotely access
// and maintain Message Queuing (MSMQ) objects in a directory service through an MSMQ
// Directory Service server.
//
// # Overview
//
// Message Queuing is a communications service that provides asynchronous and reliable
// message passing between client applications, including between client applications
// running on different hosts. In Message Queuing, clients send messages to queues and
// consume messages from queues. Queues provide message persistence, enabling the sending
// and receiving of client applications to operate asynchronously from each other.
//
// Because Message Queuing involves message passing between nodes, a directory service
// can be useful to Message Queuing services in several ways. First, a directory service
// can provide network topology information that the Message Queuing services can use
// to route messages between nodes. Second, a directory service can be used as a key
// distribution mechanism for security services that are used by Message Queuing to
// secure messages and authenticate clients. Third, a directory service can provide
// clients with discovery capabilities, allowing clients to discover the queues available
// within the network.
//
// The Message Queuing (MSMQ): Directory Service Protocol provides a set of procedure
// calls that can be made between a client and an MSMQ Directory Server. The client
// can use these calls to access the Directory Service remotely. For example, a client
// can use this protocol to create queue objects in a directory. This protocol is intended
// for use by the Message Queuing system.
//
// The directory defined by the Message Queuing (MSMQ): Directory Service Protocol is
// composed of eight types of directory objects representing enterprises, sites, routing
// links, machines, users, queues, connected networks, and deleted objects.
//
// Each directory object is composed of a collection of properties. Each property has
// an integer property identifier and a variant property value. Properties are specific
// to the directory object type. Most directory object types include a GUID property
// to identify a particular object instance, a directory service pathname property specifying
// where in the directory the object is stored, and security properties. Some MSMQ object
// properties are assigned by the directory service while other MSMQ object properties
// are specified by the client. Some properties are immutable; some properties are mutable
// by the directory server but not by the client while other properties are mutable
// by both.
//
// The directory objects either are maintained by the MSMQ Directory Service servers
// or are stored in a directory that supports [MS-ADTS]. In the latter case, the Message
// Queuing (MSMQ): Directory Service Protocol server interprets the client RPC calls
// and uses the [MS-MQDSSM] algorithm to access the [MS-ADTS]-based directory to satisfy
// client requests.
//
// The Message Queuing (MSMQ): Directory Service Protocol provides methods to create,
// update, retrieve, and delete objects from the directory service by using either the
// object name or the unique object GUID as a key to identify the object. Separate interface
// methods are implemented to manipulate object security properties.
//
// The Message Queuing (MSMQ): Directory Service Protocol also provides a simple query
// mechanism that allows the enumeration of directory objects through comparison with
// client-supplied values. The client can specify the matching criteria, the properties
// to be returned, and the sort order for the results. The server computes the result
// set. Thereafter, the client retrieves the results in order, in an iterative manner
// through repeated calls to the server, each call returning the next portion of the
// result set.
//
// The Message Queuing (MSMQ): Directory Service Protocol includes a method for RPC
// endpoint port negotiation. Through this, the client can determine the RPC endpoint
// port to use for this protocol.
//
// Generally, for methods that create, update, or delete information in the directory
// service, the Message Queuing (MSMQ): Directory Service Protocol relies on security
// mechanisms of the underlying RPC transport to provide client authentication information
// to the server. There are two exceptions to this. When setting properties on a machine
// object, and when retrieving the security properties of a machine object, the server
// invokes a challenge/response callback to the client to authenticate the client. This
// client signs the challenge by using a private key, and the server validates the signature
// by using a corresponding public key stored with the machine object in the directory.
//
// Because the directory service provides network topology information and security
// key distribution, clients can trust the source of this data. Therefore, this protocol
// includes methods for a security handshake to allow mutual authentication and to establish
// cryptographic keys that are used to compute digital signatures. These handshake methods
// tunnel Generic Security Service API (GSS-API), as specified in [RFC2743], operations
// to establish a security context. See [RFC2743] section 2.2.
//
// All methods that return data to the client include signed hashes over returned data,
// allowing the client to authenticate the source of the data and verify that the data
// has not been tampered with en route. The signed hashes are computed by using the
// established security context.
//
// This is an RPC-based protocol consisting of simple request-response exchanges. For
// every method request that the server receives, it executes the method and returns
// a completion. The client simply returns the completion status to the caller.
package mqds

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	mqmq "github.com/oiweiwei/go-msrpc/msrpc/mqmq"
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
	_ = mqmq.GoPackage
)

var (
	// import guard
	GoPackage = "mqds"
)

// LT represents the PRLT RPC constant
var LT = 0

// LE represents the PRLE RPC constant
var LE = 1

// GT represents the PRGT RPC constant
var GT = 2

// GE represents the PRGE RPC constant
var GE = 3

// EQ represents the PREQ RPC constant
var EQ = 4

// NE represents the PRNE RPC constant
var NE = 5

// QuerySortAscend represents the QUERY_SORTASCEND RPC constant
var QuerySortAscend = 0

// QuerySortDescend represents the QUERY_SORTDESCEND RPC constant
var QuerySortDescend = 1

// PropertyRestriction structure represents MQPROPERTYRESTRICTION RPC structure.
//
// The MQPROPERTYRESTRICTION structure specifies a logical expression over an MSMQ object
// property. The logical expression evaluates to TRUE or FALSE. The logical expression
// is defined through a relational comparison operation between an MSMQ object property
// and a specified constant value.
type PropertyRestriction struct {
	// rel:   Specifies the binary relation to be computed between the MSMQ object property
	// identified by prop and the constant specified by prval. The value of this field MUST
	// be one of the values as defined as follows.
	//
	//	+-----------------+---------------------------+
	//	|                 |                           |
	//	|      VALUE      |          MEANING          |
	//	|                 |                           |
	//	+-----------------+---------------------------+
	//	+-----------------+---------------------------+
	//	| PRLT 0x00000000 | Less than.                |
	//	+-----------------+---------------------------+
	//	| PRLE 0x00000001 | Less than or equal to.    |
	//	+-----------------+---------------------------+
	//	| PRGT 0x00000002 | Greater than.             |
	//	+-----------------+---------------------------+
	//	| PRGE 0x00000003 | Greater than or equal to. |
	//	+-----------------+---------------------------+
	//	| PREQ 0x00000004 | Equal to.                 |
	//	+-----------------+---------------------------+
	//	| PRNE 0x00000005 | Not equal to.             |
	//	+-----------------+---------------------------+
	Relation uint32 `idl:"name:rel" json:"relation"`
	// prop:   A property identifier specifying what MSMQ object property to be used as
	// the left operand in the binary relation specified in rel. MUST be one of the values
	// specified in the object property identifier table, as specified in section 2.2.10.1.
	Property uint32 `idl:"name:prop" json:"property"`
	// prval:   A constant value to be used as the right operand in the binary relation
	// specified in rel. The variant type of prval MUST match the variant type of the MSMQ
	// object property identified by prop as specified in the property identifier tables
	// in section 2.2.10.1.
	//
	// This structure is used in directory query operations to define a single constraint
	// over the set of directory objects to be returned. An object is deemed to satisfy
	// the constraint if the binary expression, as specified by the MQPROPERTYRESTRICTION
	// structure, evaluates to TRUE, and is deemed not to satisfy the constraint otherwise.
	// See section 3.1.4.17.
	Value *mqmq.PropertyVariant `idl:"name:prval" json:"value"`
}

func (o *PropertyRestriction) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyRestriction) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.Relation); err != nil {
		return err
	}
	if err := w.WriteData(o.Property); err != nil {
		return err
	}
	if o.Value != nil {
		if err := o.Value.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&mqmq.PropertyVariant{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyRestriction) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.Relation); err != nil {
		return err
	}
	if err := w.ReadData(&o.Property); err != nil {
		return err
	}
	if o.Value == nil {
		o.Value = &mqmq.PropertyVariant{}
	}
	if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Restriction structure represents MQRESTRICTION RPC structure.
//
// The MQRESTRICTION structure specifies a set of MQPROPERTYRESTRICTION structures.
// This structure is used in directory query operations to define a set of constraints
// over the set of directory objects to be returned. An object is deemed to satisfy
// the constraint if all the binary expressions specified by the MQPROPERTYRESTRICTION
// array evaluate to TRUE, and is deemed not to satisfy the constraint otherwise. See
// section 3.1.4.17.
type Restriction struct {
	// cRes:   MUST be set to the count of MQPROPERTYRESTRICTION structures in the paPropRes
	// array.
	RestrictionCount uint32 `idl:"name:cRes" json:"restriction_count"`
	// paPropRes:   A pointer to an array of MQPROPERTYRESTRICTION structures.
	PropertyRestriction []*PropertyRestriction `idl:"name:paPropRes;size_is:(cRes)" json:"property_restriction"`
}

func (o *Restriction) xxx_PreparePayload(ctx context.Context) error {
	if o.PropertyRestriction != nil && o.RestrictionCount == 0 {
		o.RestrictionCount = uint32(len(o.PropertyRestriction))
	}
	if o.RestrictionCount > uint32(128) {
		return fmt.Errorf("RestrictionCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Restriction) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RestrictionCount); err != nil {
		return err
	}
	if o.PropertyRestriction != nil || o.RestrictionCount > 0 {
		_ptr_paPropRes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.RestrictionCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.PropertyRestriction {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.PropertyRestriction[i1] != nil {
					if err := o.PropertyRestriction[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PropertyRestriction{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.PropertyRestriction); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&PropertyRestriction{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.PropertyRestriction, _ptr_paPropRes); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Restriction) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RestrictionCount); err != nil {
		return err
	}
	_ptr_paPropRes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.RestrictionCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.RestrictionCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.PropertyRestriction", sizeInfo[0])
		}
		o.PropertyRestriction = make([]*PropertyRestriction, sizeInfo[0])
		for i1 := range o.PropertyRestriction {
			i1 := i1
			if o.PropertyRestriction[i1] == nil {
				o.PropertyRestriction[i1] = &PropertyRestriction{}
			}
			if err := o.PropertyRestriction[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_paPropRes := func(ptr interface{}) { o.PropertyRestriction = *ptr.(*[]*PropertyRestriction) }
	if err := w.ReadPointer(&o.PropertyRestriction, _s_paPropRes, _ptr_paPropRes); err != nil {
		return err
	}
	return nil
}

// ColumnSet structure represents MQCOLUMNSET RPC structure.
//
// The MQCOLUMNSET structure specifies a list of MSMQ object property identifiers. This
// structure is used in directory query operations to define the set of MSMQ object
// properties to be returned. See section 3.1.4.17.
type ColumnSet struct {
	// cCol:   MUST be set to the count of property identifiers in aCol.
	ColumnCount uint32 `idl:"name:cCol" json:"column_count"`
	// aCol:   A pointer to an array of property identifiers. Each element of the array
	// MUST be one of the values specified in the object property identifier table in section
	// 2.2.10.1.
	Columns []uint32 `idl:"name:aCol;size_is:(cCol)" json:"columns"`
}

func (o *ColumnSet) xxx_PreparePayload(ctx context.Context) error {
	if o.Columns != nil && o.ColumnCount == 0 {
		o.ColumnCount = uint32(len(o.Columns))
	}
	if o.ColumnCount > uint32(128) {
		return fmt.Errorf("ColumnCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ColumnSet) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ColumnCount); err != nil {
		return err
	}
	if o.Columns != nil || o.ColumnCount > 0 {
		_ptr_aCol := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ColumnCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Columns {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Columns[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Columns); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint32(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Columns, _ptr_aCol); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ColumnSet) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ColumnCount); err != nil {
		return err
	}
	_ptr_aCol := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ColumnCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ColumnCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Columns", sizeInfo[0])
		}
		o.Columns = make([]uint32, sizeInfo[0])
		for i1 := range o.Columns {
			i1 := i1
			if err := w.ReadData(&o.Columns[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_aCol := func(ptr interface{}) { o.Columns = *ptr.(*[]uint32) }
	if err := w.ReadPointer(&o.Columns, _s_aCol, _ptr_aCol); err != nil {
		return err
	}
	return nil
}

// SortKey structure represents MQSORTKEY RPC structure.
//
// The MQSORTKEY structure specifies a sort key and sort order. This structure is used
// in directory query operations to identify an MSMQ object property to be used as a
// sort key by which to sort the result set, and to define the sort order for that key.
type SortKey struct {
	// propColumn:   The property identifier to be used as the sort key. MUST be one of
	// the values specified in the object property identifier table in section 2.2.10.1.
	PropertyColumn uint32 `idl:"name:propColumn" json:"property_column"`
	// dwOrder:   A symbolic constant specifying whether the sort is to be done in ascending
	// or descending order. MUST be set to one of the following values.
	//
	//	+------------------------------+------------------+
	//	|                              |                  |
	//	|            VALUE             |     MEANING      |
	//	|                              |                  |
	//	+------------------------------+------------------+
	//	+------------------------------+------------------+
	//	| QUERY_SORTASCEND 0x00000000  | Ascending sort.  |
	//	+------------------------------+------------------+
	//	| QUERY_SORTDESCEND 0x00000001 | Descending sort. |
	//	+------------------------------+------------------+
	Order uint32 `idl:"name:dwOrder" json:"order"`
}

func (o *SortKey) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SortKey) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.PropertyColumn); err != nil {
		return err
	}
	if err := w.WriteData(o.Order); err != nil {
		return err
	}
	return nil
}
func (o *SortKey) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.PropertyColumn); err != nil {
		return err
	}
	if err := w.ReadData(&o.Order); err != nil {
		return err
	}
	return nil
}

// SortSet structure represents MQSORTSET RPC structure.
//
// The MQSORTSET structure specifies a multipart sort key. This structure is used in
// directory query operations to define a collection of sort keys and sort orders by
// which to sort the result set. See section 3.1.4.17.
type SortSet struct {
	// cCol:   MUST be set to the count of MQSORTKEY structures referenced by aCol.
	ColumnCount uint32 `idl:"name:cCol" json:"column_count"`
	// aCol:   A pointer to an array of MQSORTKEY structures.
	Columns []*SortKey `idl:"name:aCol;size_is:(cCol)" json:"columns"`
}

func (o *SortSet) xxx_PreparePayload(ctx context.Context) error {
	if o.Columns != nil && o.ColumnCount == 0 {
		o.ColumnCount = uint32(len(o.Columns))
	}
	if o.ColumnCount > uint32(128) {
		return fmt.Errorf("ColumnCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SortSet) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ColumnCount); err != nil {
		return err
	}
	if o.Columns != nil || o.ColumnCount > 0 {
		_ptr_aCol := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ColumnCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Columns {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Columns[i1] != nil {
					if err := o.Columns[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&SortKey{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Columns); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&SortKey{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Columns, _ptr_aCol); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SortSet) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ColumnCount); err != nil {
		return err
	}
	_ptr_aCol := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ColumnCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ColumnCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Columns", sizeInfo[0])
		}
		o.Columns = make([]*SortKey, sizeInfo[0])
		for i1 := range o.Columns {
			i1 := i1
			if o.Columns[i1] == nil {
				o.Columns[i1] = &SortKey{}
			}
			if err := o.Columns[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_aCol := func(ptr interface{}) { o.Columns = *ptr.(*[]*SortKey) }
	if err := w.ReadPointer(&o.Columns, _s_aCol, _ptr_aCol); err != nil {
		return err
	}
	return nil
}
