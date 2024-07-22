// The dtyp package implements the DTYP client protocol.
//
// # Introduction
//
// This document provides a collection of commonly used data types, which are categorized
// into two basic types: common base types and common data types. The common base types
// are those types that Microsoft compilers natively support. The common data types
// are data types that are frequently used by many protocols. These data types are user-defined
// types.
//
// # Overview
//
// Two types of data structures are specified in this document: data structures that
// are specified in terms of the wire format and data structures that are RPC-marshaled
// as specified in [MS-RPCE]. The latter are specified by using the Interface Definition
// Language (IDL) that is defined in [MS-RPCE] section 2.2.4.
//
// For some types of data, both formats are shown. For example, both formats are shown
// if some protocols use the raw wire format but other protocols use the RPC-marshaled
// format. Any protocol that uses a data structure name in its IDL necessarily implies
// the use of the IDL version of the data structure. Any other use implies the use of
// the wire format version unless otherwise specified by the protocol that uses the
// data structure.
package dtyp

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
	GoPackage = "dtyp"
)

// AccessMaskGenericRead represents the ACCESS_MASK_GENERIC_READ RPC constant
var AccessMaskGenericRead = 2147483648

// AccessMaskGenericWrite represents the ACCESS_MASK_GENERIC_WRITE RPC constant
var AccessMaskGenericWrite = 1073741824

// AccessMaskGenericExecute represents the ACCESS_MASK_GENERIC_EXECUTE RPC constant
var AccessMaskGenericExecute = 536870912

// AccessMaskGenericAll represents the ACCESS_MASK_GENERIC_ALL RPC constant
var AccessMaskGenericAll = 268435456

// AccessMaskMaximumAllowed represents the ACCESS_MASK_MAXIMUM_ALLOWED RPC constant
var AccessMaskMaximumAllowed = 33554432

// AccessMaskAccessSystemSecurity represents the ACCESS_MASK_ACCESS_SYSTEM_SECURITY RPC constant
var AccessMaskAccessSystemSecurity = 16777216

// AccessMaskSynchronize represents the ACCESS_MASK_SYNCHRONIZE RPC constant
var AccessMaskSynchronize = 1048576

// AccessMaskWriteOwner represents the ACCESS_MASK_WRITE_OWNER RPC constant
var AccessMaskWriteOwner = 524288

// AccessMaskWriteDACL represents the ACCESS_MASK_WRITE_DACL RPC constant
var AccessMaskWriteDACL = 262144

// AccessMaskReadControl represents the ACCESS_MASK_READ_CONTROL RPC constant
var AccessMaskReadControl = 131072

// AccessMaskDelete represents the ACCESS_MASK_DELETE RPC constant
var AccessMaskDelete = 65536

// ACEFlagContainerInheritACE represents the ACE_FLAG_CONTAINER_INHERIT_ACE RPC constant
var ACEFlagContainerInheritACE = 2

// ACEFlagFailedAccessACEFlag represents the ACE_FLAG_FAILED_ACCESS_ACE_FLAG RPC constant
var ACEFlagFailedAccessACEFlag = 128

// ACEFlagInheritOnlyACE represents the ACE_FLAG_INHERIT_ONLY_ACE RPC constant
var ACEFlagInheritOnlyACE = 8

// ACEFlagInheritedACE represents the ACE_FLAG_INHERITED_ACE RPC constant
var ACEFlagInheritedACE = 16

// ACEFlagNoPropagateInheritACE represents the ACE_FLAG_NO_PROPAGATE_INHERIT_ACE RPC constant
var ACEFlagNoPropagateInheritACE = 4

// ACEFlagObjectInheritACE represents the ACE_FLAG_OBJECT_INHERIT_ACE RPC constant
var ACEFlagObjectInheritACE = 1

// ACEFlagSuccessfulAccessACEFlag represents the ACE_FLAG_SUCCESSFUL_ACCESS_ACE_FLAG RPC constant
var ACEFlagSuccessfulAccessACEFlag = 64

// Filetime structure represents FILETIME RPC structure.
//
// The FILETIME structure is a 64-bit value that represents the number of 100-nanosecond
// intervals that have elapsed since January 1, 1601, Coordinated Universal Time (UTC).
type Filetime struct {
	// dwLowDateTime:  A 32-bit unsigned integer that contains the low-order bits of the
	// file time.
	LowDateTime uint32 `idl:"name:dwLowDateTime" json:"low_date_time"`
	// dwHighDateTime:  A 32-bit unsigned integer that contains the high-order bits of the
	// file time.
	HighDateTime uint32 `idl:"name:dwHighDateTime" json:"high_date_time"`
}

func (o *Filetime) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Filetime) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Filetime) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// GUID structure represents GUID RPC structure.
//
// The packet version is used within block protocols. The following diagram represents
// a GUID as an opaque sequence of bytes.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Data1                                                                                                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Data2                                                         | Data3                                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Data4                                                                                                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//
// The following structure is an IDL representation of GUID equivalent to and compatible
// with a DCE UUID ([C706] section A.1) according to the following mappings.
type GUID struct {
	// Data1 (4 bytes): The value of the Data1 member (section 2.3.4), in little-endian
	// byte order.
	//
	// Data1:  This member is generally treated as an opaque value. This member is equivalent
	// to the time_low field of a DCE UUID ([C706] section A.1).
	Data1 uint32 `idl:"name:Data1" json:"data1"`
	// Data2 (2 bytes): The value of the Data2 member (section 2.3.4), in little-endian
	// byte order.
	//
	// Data2:  This member is generally treated as an opaque value. This member is equivalent
	// to the time_mid field of a DCE UUID ([C706] section A.1).
	Data2 uint16 `idl:"name:Data2" json:"data2"`
	// Data3 (2 bytes): The value of the Data3 member (section 2.3.4), in little-endian
	// byte order.
	//
	// Data3:  This member is generally treated as an opaque value. This member is equivalent
	// to the time_hi_and_version field of a DCE UUID ([C706] section A.1).
	Data3 uint16 `idl:"name:Data3" json:"data3"`
	// Data4 (8 bytes): The value of the Data4 member (section 2.3.4), in little-endian
	// byte order.
	//
	// Data4:  This array is generally treated as a sequence of opaque values. This member
	// is equivalent to the following sequence of fields of a DCE UUID ([C706] section A.1)
	// in this order: clock_seq_hi_and_reserved, clock_seq_low, and the sequence of bytes
	// in the node field.
	Data4 []byte `idl:"name:Data4" json:"data4"`
}

func (o *GUID) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *GUID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Data1); err != nil {
		return err
	}
	if err := w.WriteData(o.Data2); err != nil {
		return err
	}
	if err := w.WriteData(o.Data3); err != nil {
		return err
	}
	for i1 := range o.Data4 {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.Data4[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data4); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *GUID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Data1); err != nil {
		return err
	}
	if err := w.ReadData(&o.Data2); err != nil {
		return err
	}
	if err := w.ReadData(&o.Data3); err != nil {
		return err
	}
	o.Data4 = make([]byte, 8)
	for i1 := range o.Data4 {
		i1 := i1
		if err := w.ReadData(&o.Data4[i1]); err != nil {
			return err
		}
	}
	return nil
}

// UUID structure represents UUID RPC structure.
type UUID struct {
	Data1 uint32 `idl:"name:Data1" json:"data1"`
	Data2 uint16 `idl:"name:Data2" json:"data2"`
	Data3 uint16 `idl:"name:Data3" json:"data3"`
	Data4 []byte `idl:"name:Data4" json:"data4"`
}

func (o *UUID) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UUID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Data1); err != nil {
		return err
	}
	if err := w.WriteData(o.Data2); err != nil {
		return err
	}
	if err := w.WriteData(o.Data3); err != nil {
		return err
	}
	for i1 := range o.Data4 {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.Data4[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data4); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *UUID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Data1); err != nil {
		return err
	}
	if err := w.ReadData(&o.Data2); err != nil {
		return err
	}
	if err := w.ReadData(&o.Data3); err != nil {
		return err
	}
	o.Data4 = make([]byte, 8)
	for i1 := range o.Data4 {
		i1 := i1
		if err := w.ReadData(&o.Data4[i1]); err != nil {
			return err
		}
	}
	return nil
}

// ClassID structure represents CLSID RPC structure.
type ClassID GUID

func (o *ClassID) GUID() *GUID { return (*GUID)(o) }

func (o *ClassID) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClassID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Data1); err != nil {
		return err
	}
	if err := w.WriteData(o.Data2); err != nil {
		return err
	}
	if err := w.WriteData(o.Data3); err != nil {
		return err
	}
	for i1 := range o.Data4 {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.Data4[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data4); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClassID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Data1); err != nil {
		return err
	}
	if err := w.ReadData(&o.Data2); err != nil {
		return err
	}
	if err := w.ReadData(&o.Data3); err != nil {
		return err
	}
	o.Data4 = make([]byte, 8)
	for i1 := range o.Data4 {
		i1 := i1
		if err := w.ReadData(&o.Data4[i1]); err != nil {
			return err
		}
	}
	return nil
}

// LargeInteger structure represents LARGE_INTEGER RPC structure.
//
// The LARGE_INTEGER structure is used to represent a 64-bit signed integer value.
type LargeInteger struct {
	QuadPart int64 `idl:"name:QuadPart" json:"quad_part"`
}

func (o *LargeInteger) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *LargeInteger) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *LargeInteger) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.QuadPart); err != nil {
		return err
	}
	return nil
}

// EventDescriptor structure represents EVENT_DESCRIPTOR RPC structure.
//
// The EVENT_DESCRIPTOR structure specifies the metadata that defines an event.
type EventDescriptor struct {
	// Id:  The event identifier.
	ID uint16 `idl:"name:Id" json:"id"`
	// Version:  The version of the event, which indicates a revision to the event definition.
	// The Version and Id members uniquely identify the event within the scope of a provider.
	Version uint8 `idl:"name:Version" json:"version"`
	// Channel:  Defines the audience for the event (for example, administrator or developer).
	Channel uint8 `idl:"name:Channel" json:"channel"`
	// Level:  Specifies the severity or level of detail included in the event (for example,
	// informational or fatal).
	Level uint8 `idl:"name:Level" json:"level"`
	// Opcode:  Identifies a step in a sequence of operations being performed within a Task.
	Opcode uint8 `idl:"name:Opcode" json:"opcode"`
	// Task:  Identifies a larger unit of work within an application or component (broader
	// in scope than the Opcode).
	Task uint16 `idl:"name:Task" json:"task"`
	// Keyword:  A bitmask that specifies a logical group of related events. Each bit corresponds
	// to one group. An event can belong to one or more groups. The keyword can contain
	// one or more provider-defined keywords, standard keywords, or both.
	//
	// This structure represents an event defined in a manifest and is included in the EVENT_HEADER
	// structure.
	Keyword uint64 `idl:"name:Keyword" json:"keyword"`
}

func (o *EventDescriptor) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *EventDescriptor) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.ID); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	if err := w.WriteData(o.Channel); err != nil {
		return err
	}
	if err := w.WriteData(o.Level); err != nil {
		return err
	}
	if err := w.WriteData(o.Opcode); err != nil {
		return err
	}
	if err := w.WriteData(o.Task); err != nil {
		return err
	}
	if err := w.WriteData(o.Keyword); err != nil {
		return err
	}
	return nil
}
func (o *EventDescriptor) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.ID); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	if err := w.ReadData(&o.Channel); err != nil {
		return err
	}
	if err := w.ReadData(&o.Level); err != nil {
		return err
	}
	if err := w.ReadData(&o.Opcode); err != nil {
		return err
	}
	if err := w.ReadData(&o.Task); err != nil {
		return err
	}
	if err := w.ReadData(&o.Keyword); err != nil {
		return err
	}
	return nil
}

// EventHeader structure represents EVENT_HEADER RPC structure.
//
// The EVENT_HEADER structure defines the main parameters of an event.
type EventHeader struct {
	// Size:  Size of the event record, in bytes.
	Size uint16 `idl:"name:Size" json:"size"`
	// HeaderType:  Reserved.
	HeaderType uint16 `idl:"name:HeaderType" json:"header_type"`
	// Flags:  Flags that provide information about the event such as the type of session
	// it was logged to and whether the event contains extended data. This member can contain
	// one or more of the following flags.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|                                   |                                                                                  |
	//	|               VALUE               |                                     MEANING                                      |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| EVENT_HEADER_FLAG_EXTENDED_INFO   | The ExtendedData member of the event record contains data.                       |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| EVENT_HEADER_FLAG_PRIVATE_SESSION | The event was logged to a private session.                                       |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| EVENT_HEADER_FLAG_STRING_ONLY     | The event data is a null-terminated Unicode string.                              |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| EVENT_HEADER_FLAG_TRACE_MESSAGE   | The provider used an implementation-specific trace message function to log the   |
	//	|                                   | event. Typically indicates that the event was written by the Windows software    |
	//	|                                   | trace preprocessor.                                                              |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| EVENT_HEADER_FLAG_NO_CPUTIME      | Indicates that elapsed execution time was not recorded; the ProcessorTime member |
	//	|                                   | can be used to determine the elapsed execution time.                             |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| EVENT_HEADER_FLAG_32_BIT_HEADER   | Indicates that the provider was running on a 32-bit computer or in a WOW64       |
	//	|                                   | session.                                                                         |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| EVENT_HEADER_FLAG_64_BIT_HEADER   | Indicates that the provider was running on a 64-bit computer.                    |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| EVENT_HEADER_FLAG_DECODE_GUID     | Indicates that the ProviderId member of the event record is a decode GUID rather |
	//	|                                   | than a control GUID.<2>                                                          |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| EVENT_HEADER_FLAG_CLASSIC_HEADER  | Indicates that provider used a trace event function to log the event.            |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| EVENT_HEADER_FLAG_PROCESSOR_INDEX | If this flag is set, the identifier for the CPU that logged the event MUST be    |
	//	|                                   | accessed using the ProcessorIndex member of the BufferContext member of the      |
	//	|                                   | event record. If this flag is not set, the identifier for the CPU that logged    |
	//	|                                   | the event MUST be read from the ProcessorNumber member of the BufferContext      |
	//	|                                   | member of the event record.<3>                                                   |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	Flags uint16 `idl:"name:Flags" json:"flags"`
	// EventProperty:  Indicates the source to use for parsing the event data.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                                       |                                                                                  |
	//	|                 VALUE                 |                                     MEANING                                      |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| EVENT_HEADER_PROPERTY_XML             | Indicates that you need a manifest to parse the event data.                      |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| EVENT_HEADER_PROPERTY_FORWARDED_XML   | Indicates that the event data contains within itself a fully rendered XML        |
	//	|                                       | description of the data, so you do not need a manifest to parse the event data.  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| EVENT_HEADER_PROPERTY_LEGACY_EVENTLOG | Indicates that you need a WMI MOF class to parse the event data.                 |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	EventProperty uint16 `idl:"name:EventProperty" json:"event_property"`
	// ThreadId:  Identifies the thread that generated the event.
	ThreadID uint32 `idl:"name:ThreadId" json:"thread_id"`
	// ProcessId:  Identifies the process that generated the event.
	ProcessID uint32 `idl:"name:ProcessId" json:"process_id"`
	// TimeStamp:  Contains the time that the event occurred. The resolution is system time
	// unless the ProcessTraceMode member of EVENT_TRACE_LOGFILE contains the PROCESS_TRACE_MODE_RAW_TIMESTAMP
	// flag, in which case the resolution depends on the value of the Wnode.ClientContext
	// member of EVENT_TRACE_PROPERTIES at the time the controller created the session.
	Timestamp *LargeInteger `idl:"name:TimeStamp" json:"timestamp"`
	// ProviderId:  GUID that uniquely identifies the provider that logged the event.
	ProviderID *GUID `idl:"name:ProviderId" json:"provider_id"`
	// EventDescriptor:  Defines information about the event such as the event identifier
	// and severity level.
	EventDescriptor *EventDescriptor     `idl:"name:EventDescriptor" json:"event_descriptor"`
	Field10         *EventHeader_Field10 `idl:"name:" json:""`
	// ActivityId:  Identifier that relates two events.
	//
	// The KernelTime and UserTime members can be used to determine the CPU cost in units
	// for a set of instructions (the values indicate the CPU usage charged to that thread
	// at the time of logging). For example, if Event A and Event B are consecutively logged
	// by the same thread and they have CPU usage numbers 150 and 175, then the activity
	// that was performed by that thread between events A and B cost 25 CPU time units (175
	// – 150).
	ActivityID *GUID `idl:"name:ActivityId" json:"activity_id"`
}

func (o *EventHeader) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *EventHeader) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if err := w.WriteData(o.HeaderType); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.EventProperty); err != nil {
		return err
	}
	if err := w.WriteData(o.ThreadID); err != nil {
		return err
	}
	if err := w.WriteData(o.ProcessID); err != nil {
		return err
	}
	if o.Timestamp != nil {
		if err := o.Timestamp.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ProviderID != nil {
		if err := o.ProviderID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.EventDescriptor != nil {
		if err := o.EventDescriptor.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&EventDescriptor{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// FIXME unknown type
	if o.ActivityID != nil {
		if err := o.ActivityID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *EventHeader) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if err := w.ReadData(&o.HeaderType); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.EventProperty); err != nil {
		return err
	}
	if err := w.ReadData(&o.ThreadID); err != nil {
		return err
	}
	if err := w.ReadData(&o.ProcessID); err != nil {
		return err
	}
	if o.Timestamp == nil {
		o.Timestamp = &LargeInteger{}
	}
	if err := o.Timestamp.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ProviderID == nil {
		o.ProviderID = &GUID{}
	}
	if err := o.ProviderID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.EventDescriptor == nil {
		o.EventDescriptor = &EventDescriptor{}
	}
	if err := o.EventDescriptor.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	// FIXME: unknown type
	if o.ActivityID == nil {
		o.ActivityID = &GUID{}
	}
	if err := o.ActivityID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

type EventHeader_Field10 struct {
	Field1 *EventHeader_Field10_Field1 `idl:"name:" json:""`
	// ProcessorTime:  For private sessions, the elapsed execution time for user-mode instructions,
	// in CPU ticks.
	ProcessorTime uint64 `idl:"name:ProcessorTime" json:"processor_time"`
}

// EventHeader_Field10_Field1 structure represents EVENT_HEADER structure anonymous member.
//
// The EVENT_HEADER structure defines the main parameters of an event.
type EventHeader_Field10_Field1 struct {
	// KernelTime:  Elapsed execution time for kernel-mode instructions, in CPU time units.
	// For private sessions, the value in the ProcessorTime member can be used instead.
	KernelTime uint32 `idl:"name:KernelTime" json:"kernel_time"`
	// UserTime:  Elapsed execution time for user-mode instructions, in CPU time units.
	// For private sessions, the value in the ProcessorTime member can be used instead.
	UserTime uint32 `idl:"name:UserTime" json:"user_time"`
}

func (o *EventHeader_Field10_Field1) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *EventHeader_Field10_Field1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.KernelTime); err != nil {
		return err
	}
	if err := w.WriteData(o.UserTime); err != nil {
		return err
	}
	return nil
}
func (o *EventHeader_Field10_Field1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.KernelTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.UserTime); err != nil {
		return err
	}
	return nil
}

// LUID structure represents LUID RPC structure.
//
// The LUID structure is 64-bit value guaranteed to be unique only on the system on
// which it was generated. The uniqueness of a locally unique identifier (LUID) is guaranteed
// only until the system is restarted.
type LUID struct {
	// LowPart:  The low-order bits of the structure.
	LowPart uint32 `idl:"name:LowPart" json:"low_part"`
	// HighPart:  The high-order bits of the structure.
	HighPart int32 `idl:"name:HighPart" json:"high_part"`
}

func (o *LUID) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *LUID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *LUID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Multistring structure represents MULTI_SZ RPC structure.
//
// The MULTI_SZ structure defines an implementation-specific<4> type that contains a
// sequence of null-terminated strings, terminated by an empty string (\0) so that the
// last two characters are both null terminators.
type Multistring struct {
	// Value:  A data buffer, which is a string literal containing multiple null-terminated
	// strings serially.
	Value string `idl:"name:Value" json:"value"`
	// nChar:  The length, in characters, including the two terminating nulls.
	CharCount uint32 `idl:"name:nChar" json:"char_count"`
}

func (o *Multistring) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Multistring) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Value != "" {
		_ptr_Value := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.Value); err != nil {
				return err
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
	if err := w.WriteData(o.CharCount); err != nil {
		return err
	}
	return nil
}
func (o *Multistring) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_Value := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.Value); err != nil {
			return err
		}
		return nil
	})
	_s_Value := func(ptr interface{}) { o.Value = *ptr.(*string) }
	if err := w.ReadPointer(&o.Value, _s_Value, _ptr_Value); err != nil {
		return err
	}
	if err := w.ReadData(&o.CharCount); err != nil {
		return err
	}
	return nil
}

// UnicodeString structure represents RPC_UNICODE_STRING RPC structure.
//
// The RPC_UNICODE_STRING structure specifies a Unicode string. This structure is defined
// in IDL as follows:
type UnicodeString struct {
	// Length:  The length, in bytes, of the string pointed to by the Buffer member. The
	// length MUST be a multiple of 2. The length MUST equal the entire size of the buffer.
	Length uint16 `idl:"name:Length" json:"length"`
	// MaximumLength:  The maximum size, in bytes, of the string pointed to by Buffer. The
	// size MUST be a multiple of 2. If not, the size MUST be decremented by 1 prior to
	// use. This value MUST not be less than Length.
	MaximumLength uint16 `idl:"name:MaximumLength" json:"maximum_length"`
	// Buffer:  A pointer to a string buffer. The string pointed to by the buffer member
	// MUST NOT include a terminating null character.
	Buffer string `idl:"name:Buffer;size_is:((MaximumLength/2));length_is:((Length/2))" json:"buffer"`
}

func (o *UnicodeString) xxx_PreparePayload(ctx context.Context) error {
	if o.Buffer != "" && o.MaximumLength == 0 {
		o.MaximumLength = uint16((len(o.Buffer) * 2))
	}
	if o.Buffer != "" && o.Length == 0 {
		o.Length = uint16((len(o.Buffer) * 2))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UnicodeString) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if o.Buffer != "" || (o.MaximumLength/2) > 0 {
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
			_Buffer_buf := utf16.Encode([]rune(o.Buffer))
			if uint64(len(_Buffer_buf)) > sizeInfo[0] {
				_Buffer_buf = _Buffer_buf[:sizeInfo[0]]
			}
			for i1 := range _Buffer_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_Buffer_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_Buffer_buf); uint64(i1) < sizeInfo[0]; i1++ {
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
func (o *UnicodeString) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		var _Buffer_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Buffer_buf", sizeInfo[0])
		}
		_Buffer_buf = make([]uint16, sizeInfo[0])
		for i1 := range _Buffer_buf {
			i1 := i1
			if err := w.ReadData(&_Buffer_buf[i1]); err != nil {
				return err
			}
		}
		o.Buffer = strings.TrimRight(string(utf16.Decode(_Buffer_buf)), ndr.ZeroString)
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*string) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// ServerInfo100 structure represents SERVER_INFO_100 RPC structure.
//
// The SERVER_INFO_100 structure contains information about the specified server, including
// the name and platform.
type ServerInfo100 struct {
	// sv100_platform_id:  Specifies the information level to use for platform-specific
	// information.
	//
	//	+-----------------+-------+
	//	|                 |       |
	//	|      NAME       | VALUE |
	//	|                 |       |
	//	+-----------------+-------+
	//	+-----------------+-------+
	//	| PLATFORM_ID_DOS |   300 |
	//	+-----------------+-------+
	//	| PLATFORM_ID_OS2 |   400 |
	//	+-----------------+-------+
	//	| PLATFORM_ID_NT  |   500 |
	//	+-----------------+-------+
	//	| PLATFORM_ID_OSF |   600 |
	//	+-----------------+-------+
	//	| PLATFORM_ID_VMS |   700 |
	//	+-----------------+-------+
	PlatformID uint32 `idl:"name:sv100_platform_id" json:"platform_id"`
	// sv100_name:  A pointer to a null-terminated Unicode UTF-16 Internet host name or
	// NetBIOS host name of a server.
	Name string `idl:"name:sv100_name;string" json:"name"`
}

func (o *ServerInfo100) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ServerInfo100) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.PlatformID); err != nil {
		return err
	}
	if o.Name != "" {
		_ptr_sv100_name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_sv100_name); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ServerInfo100) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.PlatformID); err != nil {
		return err
	}
	_ptr_sv100_name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_sv100_name := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_sv100_name, _ptr_sv100_name); err != nil {
		return err
	}
	return nil
}

// ServerInfo101 structure represents SERVER_INFO_101 RPC structure.
//
// The SERVER_INFO_101 structure contains information about the specified server, including
// the name, platform, type of server, and associated software.
type ServerInfo101 struct {
	// sv101_platform_id:  Specifies the information level to use for platform-specific
	// information.
	//
	//	+-----------------+-------+
	//	|                 |       |
	//	|      NAME       | VALUE |
	//	|                 |       |
	//	+-----------------+-------+
	//	+-----------------+-------+
	//	| PLATFORM_ID_DOS |   300 |
	//	+-----------------+-------+
	//	| PLATFORM_ID_OS2 |   400 |
	//	+-----------------+-------+
	//	| PLATFORM_ID_NT  |   500 |
	//	+-----------------+-------+
	//	| PLATFORM_ID_OSF |   600 |
	//	+-----------------+-------+
	//	| PLATFORM_ID_VMS |   700 |
	//	+-----------------+-------+
	PlatformID uint32 `idl:"name:sv101_platform_id" json:"platform_id"`
	// sv101_name:  A pointer to a null-terminated Unicode UTF-16 Internet host name or
	// NetBIOS host name of a server.
	Name string `idl:"name:sv101_name;string" json:"name"`
	// sv101_version_major:  Specifies the major release version number of the operating
	// system. The server MUST set this field to an implementation-specific major release
	// version number that corresponds to the host operating system as specified in the
	// following table.
	//
	//	+-----------------------------------------+---------------+
	//	|                OPERATING                |     MAJOR     |
	//	|                 SYSTEM                  |    VERSION    |
	//	+-----------------------------------------+---------------+
	//	+-----------------------------------------+---------------+
	//	| Windows NT 4.0 operating system         |             4 |
	//	+-----------------------------------------+---------------+
	//	| Windows 2000 operating system           |             5 |
	//	+-----------------------------------------+---------------+
	//	| Windows XP operating system             |             5 |
	//	+-----------------------------------------+---------------+
	//	| Windows Server 2003 operating system    |             5 |
	//	+-----------------------------------------+---------------+
	//	| Windows Vista operating system          |             6 |
	//	+-----------------------------------------+---------------+
	//	| Windows Server 2008 operating system    |             6 |
	//	+-----------------------------------------+---------------+
	//	| Windows Server 2008 R2 operating system |             6 |
	//	+-----------------------------------------+---------------+
	VersionMajor uint32 `idl:"name:sv101_version_major" json:"version_major"`
	// sv101_version_minor:  Specifies the minor release version number of the operating
	// system. The server MUST set this field to an implementation-specific minor release
	// version number that corresponds to the host operating system as specified in the
	// following table.
	//
	//	+------------------------+---------------+
	//	|       OPERATING        |     MINOR     |
	//	|         SYSTEM         |    VERSION    |
	//	+------------------------+---------------+
	//	+------------------------+---------------+
	//	| Windows NT 4.0         |             0 |
	//	+------------------------+---------------+
	//	| Windows 2000           |             0 |
	//	+------------------------+---------------+
	//	| Windows XP             |             1 |
	//	+------------------------+---------------+
	//	| Windows Server 2003    |             2 |
	//	+------------------------+---------------+
	//	| Windows Vista          |             0 |
	//	+------------------------+---------------+
	//	| Windows Server 2008    |             0 |
	//	+------------------------+---------------+
	//	| Windows Server 2008 R2 |             1 |
	//	+------------------------+---------------+
	VersionMinor uint32 `idl:"name:sv101_version_minor" json:"version_minor"`
	// sv101_version_type:  The sv101_version_type field specifies the SV_TYPE flags, which
	// indicate the software services that are available (but not necessarily running) on
	// the server. This member MUST be a combination of one or more of the following values.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                                      |                                                                                  |
	//	|            CONSTANT/VALUE            |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| SV_TYPE_WORKSTATION 0x00000001       | A server running the WorkStation Service.                                        |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| SV_TYPE_SERVER 0x00000002            | A server running the Server Service.                                             |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| SV_TYPE_SQLSERVER 0x00000004         | A server running SQL Server.                                                     |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| SV_TYPE_DOMAIN_CTRL 0x00000008       | A primary domain controller.                                                     |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| SV_TYPE_DOMAIN_BAKCTRL 0x00000010    | A backup domain controller.                                                      |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| SV_TYPE_TIME_SOURCE 0x00000020       | A server is available as a time source for network time synchronization.         |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| SV_TYPE_AFP 0x00000040               | An Apple File Protocol server.                                                   |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| SV_TYPE_NOVELL 0x00000080            | A Novell server.                                                                 |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| SV_TYPE_DOMAIN_MEMBER 0x00000100     | A LAN Manager 2.x domain member.                                                 |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| SV_TYPE_PRINTQ_SERVER 0x00000200     | A server sharing print queue.                                                    |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| SV_TYPE_DIALIN_SERVER 0x00000400     | A server running a dial-in service.                                              |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| SV_TYPE_XENIX_SERVER 0x00000800      | A Xenix server.                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| SV_TYPE_NT 0x00001000                | Windows Server 2003, Windows XP, Windows 2000, or Windows NT operating system.   |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| SV_TYPE_WFW 0x00002000               | A server running Windows for Workgroups.                                         |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| SV_TYPE_SERVER_MFPN 0x00004000       | Microsoft File and Print for NetWare.                                            |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| SV_TYPE_SERVER_NT 0x00008000         | Windows Server 2003, Windows 2000 Server operating system, or a server that is   |
	//	|                                      | not a domain controller.                                                         |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| SV_TYPE_POTENTIAL_BROWSER 0x00010000 | A server that can run the browser service.                                       |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| SV_TYPE_BACKUP_BROWSER 0x00020000    | A server running a browser service as backup.                                    |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| SV_TYPE_MASTER_BROWSER 0x00040000    | A server running the master browser service.                                     |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| SV_TYPE_DOMAIN_MASTER 0x00080000     | A server running the domain master browser.                                      |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| SV_TYPE_WINDOWS 0x00400000           | Windows Millennium Edition operating system, Windows 98 operating system, or     |
	//	|                                      | Windows 95.                                                                      |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| SV_TYPE_TERMINALSERVER 0x02000000    | Terminal Server.                                                                 |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| SV_TYPE_CLUSTER_VS_NT 0x04000000     | Cluster virtual servers available in the domain.                                 |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| SV_TYPE_LOCAL_LIST_ONLY 0x40000000   | Servers maintained by the browser.                                               |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| SV_TYPE_DOMAIN_ENUM 0x80000000       | Primary domain.                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| SV_TYPE_ALL 0xFFFFFFFF               | All servers.                                                                     |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	VersionType uint32 `idl:"name:sv101_version_type" json:"version_type"`
	// sv101_comment:  A pointer to a null-terminated Unicode UTF-16 string that specifies
	// a comment that describes the server.
	Comment string `idl:"name:sv101_comment;string" json:"comment"`
}

func (o *ServerInfo101) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ServerInfo101) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.PlatformID); err != nil {
		return err
	}
	if o.Name != "" {
		_ptr_sv101_name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_sv101_name); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.VersionMajor); err != nil {
		return err
	}
	if err := w.WriteData(o.VersionMinor); err != nil {
		return err
	}
	if err := w.WriteData(o.VersionType); err != nil {
		return err
	}
	if o.Comment != "" {
		_ptr_sv101_comment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Comment); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Comment, _ptr_sv101_comment); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ServerInfo101) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.PlatformID); err != nil {
		return err
	}
	_ptr_sv101_name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_sv101_name := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_sv101_name, _ptr_sv101_name); err != nil {
		return err
	}
	if err := w.ReadData(&o.VersionMajor); err != nil {
		return err
	}
	if err := w.ReadData(&o.VersionMinor); err != nil {
		return err
	}
	if err := w.ReadData(&o.VersionType); err != nil {
		return err
	}
	_ptr_sv101_comment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Comment); err != nil {
			return err
		}
		return nil
	})
	_s_sv101_comment := func(ptr interface{}) { o.Comment = *ptr.(*string) }
	if err := w.ReadPointer(&o.Comment, _s_sv101_comment, _ptr_sv101_comment); err != nil {
		return err
	}
	return nil
}

// SystemTime structure represents SYSTEMTIME RPC structure.
//
// The SYSTEMTIME structure is a date and time, in Coordinated Universal Time (UTC),
// represented by using individual WORD-sized structure members for the month, day,
// year, day of week, hour, minute, second, and millisecond.
type SystemTime struct {
	Year         uint16 `idl:"name:wYear" json:"year"`
	Month        uint16 `idl:"name:wMonth" json:"month"`
	DayOfWeek    uint16 `idl:"name:wDayOfWeek" json:"day_of_week"`
	Day          uint16 `idl:"name:wDay" json:"day"`
	Hour         uint16 `idl:"name:wHour" json:"hour"`
	Minute       uint16 `idl:"name:wMinute" json:"minute"`
	Second       uint16 `idl:"name:wSecond" json:"second"`
	Milliseconds uint16 `idl:"name:wMilliseconds" json:"milliseconds"`
}

func (o *SystemTime) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
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

// Uint128 structure represents UINT128 RPC structure.
//
// The UINT128 structure is intended to hold 128-bit unsigned integers, such as an IPv6
// destination address.
type Uint128 struct {
	Lower uint64 `idl:"name:lower" json:"lower"`
	Upper uint64 `idl:"name:upper" json:"upper"`
}

func (o *Uint128) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Uint128) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.Lower); err != nil {
		return err
	}
	if err := w.WriteData(o.Upper); err != nil {
		return err
	}
	return nil
}
func (o *Uint128) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.Lower); err != nil {
		return err
	}
	if err := w.ReadData(&o.Upper); err != nil {
		return err
	}
	return nil
}

// UlargeInteger structure represents ULARGE_INTEGER RPC structure.
//
// The ULARGE_INTEGER structure is used to represent a 64-bit unsigned integer value.
type UlargeInteger struct {
	QuadPart uint64 `idl:"name:QuadPart" json:"quad_part"`
}

func (o *UlargeInteger) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UlargeInteger) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *UlargeInteger) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.QuadPart); err != nil {
		return err
	}
	return nil
}

// SIDIDAuthority structure represents RPC_SID_IDENTIFIER_AUTHORITY RPC structure.
//
// The RPC_SID_IDENTIFIER_AUTHORITY structure is a representation of a security identifier
// (SID) authority, as specified by the SID_IDENTIFIER_AUTHORITY structure. This structure
// is defined in IDL as follows.
//
// For individual member semantics of the SID_IDENTIFIER_AUTHORITY structure, see section
// 2.4.1.
type SIDIDAuthority struct {
	Value []byte `idl:"name:Value" json:"value"`
}

func (o *SIDIDAuthority) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SIDIDAuthority) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	for i1 := range o.Value {
		i1 := i1
		if uint64(i1) >= 6 {
			break
		}
		if err := w.WriteData(o.Value[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Value); uint64(i1) < 6; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *SIDIDAuthority) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	o.Value = make([]byte, 6)
	for i1 := range o.Value {
		i1 := i1
		if err := w.ReadData(&o.Value[i1]); err != nil {
			return err
		}
	}
	return nil
}

// ObjectTypeList structure represents OBJECT_TYPE_LIST RPC structure.
//
// The OBJECT_TYPE_LIST structure identifies an object type element in a hierarchy of
// object types. The Access Check Algorithm Pseudocode functions (section 2.5.3.2) use
// an array of OBJECT_TYPE_LIST structures to define a hierarchy of an object and its
// sub-objects, such as property sets and properties.
type ObjectTypeList struct {
	// Level:  Specifies the level of the object type in the hierarchy of an object and
	// its sub-objects. Level zero indicates the object itself. Level one indicates a sub-object
	// of the object, such as a property set. Level two indicates a sub-object of the level
	// one sub-object, such as a property. There can be a maximum of five levels numbered
	// zero through four.
	//
	//	+------------------------------+--------------------------------------------+
	//	|                              |                                            |
	//	|            VALUE             |                  MEANING                   |
	//	|                              |                                            |
	//	+------------------------------+--------------------------------------------+
	//	+------------------------------+--------------------------------------------+
	//	| ACCESS_OBJECT_GUID 0x0       | Indicates the object itself at level zero. |
	//	+------------------------------+--------------------------------------------+
	//	| ACCESS_PROPERTY_SET_GUID 0x1 | Indicates a property set at level one.     |
	//	+------------------------------+--------------------------------------------+
	//	| ACCESS_PROPERTY_GUID 0x2     | Indicates a property at level two.         |
	//	+------------------------------+--------------------------------------------+
	//	| ACCESS_MAX_LEVEL 0x4         | Maximum level.                             |
	//	+------------------------------+--------------------------------------------+
	Level uint16 `idl:"name:Level" json:"level"`
	// Remaining:  Remaining access bits for this element, used by the access check algorithm,
	// as specified in section 2.5.3.2.
	Remaining uint32 `idl:"name:Remaining" json:"remaining"`
	// ObjectType:  A pointer to the GUID for the object or sub-object.
	ObjectType *GUID `idl:"name:ObjectType" json:"object_type"`
}

func (o *ObjectTypeList) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectTypeList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Level); err != nil {
		return err
	}
	if err := w.WriteData(o.Remaining); err != nil {
		return err
	}
	if o.ObjectType != nil {
		_ptr_ObjectType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ObjectType != nil {
				if err := o.ObjectType.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ObjectType, _ptr_ObjectType); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectTypeList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Level); err != nil {
		return err
	}
	if err := w.ReadData(&o.Remaining); err != nil {
		return err
	}
	_ptr_ObjectType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ObjectType == nil {
			o.ObjectType = &GUID{}
		}
		if err := o.ObjectType.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ObjectType := func(ptr interface{}) { o.ObjectType = *ptr.(**GUID) }
	if err := w.ReadPointer(&o.ObjectType, _s_ObjectType, _ptr_ObjectType); err != nil {
		return err
	}
	return nil
}

// ACEHeader structure represents ACE_HEADER RPC structure.
//
// The ACE_HEADER structure defines the type and size of an access control entry (ACE).
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| AceType                       | AceFlags                      | AceSize                                                       |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//
// The RPC representation of the ACE_HEADER defines the type and size of an ACE. The
// members and values are as specified in section 2.4.4.1.
type ACEHeader struct {
	// AceType (1 byte): An unsigned 8-bit integer that specifies the ACE types. This field
	// MUST be one of the following values.
	//
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                              |                                                                                  |
	//	|                    VALUE                     |                                     MEANING                                      |
	//	|                                              |                                                                                  |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ACCESS_ALLOWED_ACE_TYPE 0x00                 | Access-allowed ACE that uses the ACCESS_ALLOWED_ACE (section 2.4.4.2) structure. |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ACCESS_DENIED_ACE_TYPE 0x01                  | Access-denied ACE that uses the ACCESS_DENIED_ACE (section 2.4.4.4) structure.   |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| SYSTEM_AUDIT_ACE_TYPE 0x02                   | System-audit ACE that uses the SYSTEM_AUDIT_ACE (section 2.4.4.10) structure.    |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| SYSTEM_ALARM_ACE_TYPE 0x03                   | Reserved for future use.                                                         |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ACCESS_ALLOWED_COMPOUND_ACE_TYPE 0x04        | Reserved for future use.                                                         |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ACCESS_ALLOWED_OBJECT_ACE_TYPE 0x05          | Object-specific access-allowed ACE that uses the ACCESS_ALLOWED_OBJECT_ACE       |
	//	|                                              | (section 2.4.4.3) structure.<45>                                                 |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ACCESS_DENIED_OBJECT_ACE_TYPE 0x06           | Object-specific access-denied ACE that uses the ACCESS_DENIED_OBJECT_ACE         |
	//	|                                              | (section 2.4.4.5) structure.<46>                                                 |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| SYSTEM_AUDIT_OBJECT_ACE_TYPE 0x07            | Object-specific system-audit ACE that uses the SYSTEM_AUDIT_OBJECT_ACE (section  |
	//	|                                              | 2.4.4.11) structure.<47>                                                         |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| SYSTEM_ALARM_OBJECT_ACE_TYPE 0x08            | Reserved for future use.                                                         |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ACCESS_ALLOWED_CALLBACK_ACE_TYPE 0x09        | Access-allowed callback ACE that uses the ACCESS_ALLOWED_CALLBACK_ACE (section   |
	//	|                                              | 2.4.4.6) structure.<48>                                                          |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ACCESS_DENIED_CALLBACK_ACE_TYPE 0x0A         | Access-denied callback ACE that uses the ACCESS_DENIED_CALLBACK_ACE (section     |
	//	|                                              | 2.4.4.7) structure.<49>                                                          |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ACCESS_ALLOWED_CALLBACK_OBJECT_ACE_TYPE 0x0B | Object-specific access-allowed callback ACE that uses the                        |
	//	|                                              | ACCESS_ALLOWED_CALLBACK_OBJECT_ACE (section 2.4.4.8) structure.<50>              |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ACCESS_DENIED_CALLBACK_OBJECT_ACE_TYPE 0x0C  | Object-specific access-denied callback ACE that uses the                         |
	//	|                                              | ACCESS_DENIED_CALLBACK_OBJECT_ACE (section 2.4.4.9) structure.<51>               |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| SYSTEM_AUDIT_CALLBACK_ACE_TYPE 0x0D          | System-audit callback ACE that uses the SYSTEM_AUDIT_CALLBACK_ACE (section       |
	//	|                                              | 2.4.4.12) structure.<52>                                                         |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| SYSTEM_ALARM_CALLBACK_ACE_TYPE 0x0E          | Reserved for future use.                                                         |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| SYSTEM_AUDIT_CALLBACK_OBJECT_ACE_TYPE 0x0F   | Object-specific system-audit callback ACE that uses the                          |
	//	|                                              | SYSTEM_AUDIT_CALLBACK_OBJECT_ACE (section 2.4.4.14) structure.                   |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| SYSTEM_ALARM_CALLBACK_OBJECT_ACE_TYPE 0x10   | Reserved for future use.                                                         |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| SYSTEM_MANDATORY_LABEL_ACE_TYPE 0x11         | Mandatory label ACE that uses the SYSTEM_MANDATORY_LABEL_ACE (section 2.4.4.13)  |
	//	|                                              | structure.                                                                       |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| SYSTEM_RESOURCE_ATTRIBUTE_ACE_TYPE 0x12      | Resource attribute ACE that uses the SYSTEM_RESOURCE_ATTRIBUTE_ACE (section      |
	//	|                                              | 2.4.4.15)                                                                        |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| SYSTEM_SCOPED_POLICY_ID_ACE_TYPE 0x13        | A central policy ID ACE that uses the SYSTEM_SCOPED_POLICY_ID_ACE (section       |
	//	|                                              | 2.4.4.16)                                                                        |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	ACEType uint8 `idl:"name:AceType" json:"ace_type"`
	// AceFlags (1 byte): An unsigned 8-bit integer that specifies a set of ACE type-specific
	// control flags. This field can be a combination of the following values.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|                                 |                                                                                  |
	//	|              VALUE              |                                     MEANING                                      |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| CONTAINER_INHERIT_ACE 0x02      | Child objects that are containers, such as directories, inherit the              |
	//	|                                 | ACE as an effective ACE. The inherited ACE is inheritable unless the             |
	//	|                                 | NO_PROPAGATE_INHERIT_ACE bit flag is also set.                                   |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| FAILED_ACCESS_ACE_FLAG 0x80     | Used with system-audit ACEs in a system access control list (SACL) to generate   |
	//	|                                 | audit messages for failed access attempts.                                       |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| INHERIT_ONLY_ACE 0x08           | Indicates an inherit-only ACE, which does not control access to the object       |
	//	|                                 | to which it is attached. If this flag is not set, the ACE is an effective ACE    |
	//	|                                 | that controls access to the object to which it is attached. Both effective       |
	//	|                                 | and inherit-only ACEs can be inherited depending on the state of the other       |
	//	|                                 | inheritance flags.                                                               |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| INHERITED_ACE 0x10              | Used to indicate that the ACE was inherited.<54> See section 2.5.3.5 for         |
	//	|                                 | processing rules for setting this flag.                                          |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| NO_PROPAGATE_INHERIT_ACE 0x04   | If the ACE is inherited by a child object, the system clears the                 |
	//	|                                 | OBJECT_INHERIT_ACE and CONTAINER_INHERIT_ACE flags in the inherited ACE. This    |
	//	|                                 | prevents the ACE from being inherited by subsequent generations of objects.      |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| OBJECT_INHERIT_ACE 0x01         | Noncontainer child objects inherit the ACE as an effective ACE. For child        |
	//	|                                 | objects that are containers, the ACE is inherited as an inherit-only ACE unless  |
	//	|                                 | the NO_PROPAGATE_INHERIT_ACE bit flag is also set.                               |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| SUCCESSFUL_ACCESS_ACE_FLAG 0x40 | Used with system-audit ACEs in a SACL to generate audit messages for successful  |
	//	|                                 | access attempts.                                                                 |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	ACEFlags uint8 `idl:"name:AceFlags" json:"ace_flags"`
	// AceSize (2 bytes): An unsigned 16-bit integer that specifies the size, in bytes,
	// of the ACE. The AceSize field can be greater than the sum of the individual fields,
	// but MUST be a multiple of 4 to ensure alignment on a DWORD boundary. In cases where
	// the AceSize field encompasses additional data for the callback ACEs types, that data
	// is implementation-specific. Otherwise, this additional data is not interpreted and
	// MUST be ignored.
	ACESize uint16 `idl:"name:AceSize" json:"ace_size"`
}

func (o *ACEHeader) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ACEHeader) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(2); err != nil {
		return err
	}
	if err := w.WriteData(o.ACEType); err != nil {
		return err
	}
	if err := w.WriteData(o.ACEFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.ACESize); err != nil {
		return err
	}
	return nil
}
func (o *ACEHeader) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(2); err != nil {
		return err
	}
	if err := w.ReadData(&o.ACEType); err != nil {
		return err
	}
	if err := w.ReadData(&o.ACEFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.ACESize); err != nil {
		return err
	}
	return nil
}

// AccessAllowedACE structure represents ACCESS_ALLOWED_ACE RPC structure.
//
// The ACCESS_ALLOWED_ACE structure defines an ACE for the discretionary access control
// list (DACL) that controls access to an object. An access-allowed ACE allows access
// to an object for a specific trustee identified by a security identifier (SID).
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Header                                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Mask                                                                                                                          |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Sid (variable)                                                                                                                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type AccessAllowedACE struct {
	// Mask (4 bytes): An ACCESS_MASK that specifies the user rights allowed by this ACE.
	Mask uint32 `idl:"name:Mask" json:"mask"`
	// Sid (variable): The SID of a trustee. The length of the SID MUST be a multiple of
	// 4.
	SID *SID `idl:"name:Sid" json:"sid"`
}

func (o *AccessAllowedACE) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AccessAllowedACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Mask); err != nil {
		return err
	}
	if o.SID != nil {
		_ptr_Sid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SID != nil {
				if err := o.SID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SID{}).MarshalNDR(ctx, w); err != nil {
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
	return nil
}
func (o *AccessAllowedACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Mask); err != nil {
		return err
	}
	_ptr_Sid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SID == nil {
			o.SID = &SID{}
		}
		if err := o.SID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Sid := func(ptr interface{}) { o.SID = *ptr.(**SID) }
	if err := w.ReadPointer(&o.SID, _s_Sid, _ptr_Sid); err != nil {
		return err
	}
	return nil
}

// AccessAllowedObjectACE structure represents ACCESS_ALLOWED_OBJECT_ACE RPC structure.
//
// The ACCESS_ALLOWED_OBJECT_ACE structure defines an ACE that controls allowed access
// to an object, a property set, or property. The ACE contains a set of access rights,
// a GUID that identifies the type of object, and a SID that identifies the trustee
// to whom the system will grant access. The ACE also contains a GUID and a set of flags
// that control inheritance of the ACE by child objects.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Header                                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Mask                                                                                                                          |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Flags                                                                                                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ObjectType (16 bytes)                                                                                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| InheritedObjectType (16 bytes)                                                                                                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Sid (variable)                                                                                                                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type AccessAllowedObjectACE struct {
	// Mask (4 bytes): An ACCESS_MASK that specifies the user rights allowed by this ACE.
	//
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	|                                        |                                                                                  |
	//	|                 VALUE                  |                                     MEANING                                      |
	//	|                                        |                                                                                  |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ADS_RIGHT_DS_CONTROL_ACCESS 0X00000100 | The ObjectType GUID identifies an extended access right.                         |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ADS_RIGHT_DS_CREATE_CHILD 0X00000001   | The ObjectType GUID identifies a type of child object. The ACE controls the      |
	//	|                                        | trustee's right to create this type of child object.                             |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ADS_RIGHT_DS_DELETE_CHILD 0X00000002   | The ObjectType GUID identifies a type of child object. The ACE controls the      |
	//	|                                        | trustee's right to delete this type of child object.                             |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ADS_RIGHT_DS_READ_PROP 0x00000010      | The ObjectType GUID identifies a property set or property of the object. The ACE |
	//	|                                        | controls the trustee's right to read the property or property set.               |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ADS_RIGHT_DS_WRITE_PROP 0x00000020     | The ObjectType GUID identifies a property set or property of the object. The ACE |
	//	|                                        | controls the trustee's right to write the property or property set.              |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ADS_RIGHT_DS_SELF 0x00000008           | The ObjectType GUID identifies a validated write.                                |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	Mask uint32 `idl:"name:Mask" json:"mask"`
	// Flags  (4 bytes): A 32-bit unsigned integer that specifies a set of bit flags that
	// indicate whether the ObjectType and InheritedObjectType fields contain valid data.
	// This parameter can be one or more of the following values.
	//
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                              |                                                                                  |
	//	|                    VALUE                     |                                     MEANING                                      |
	//	|                                              |                                                                                  |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000                                   | Neither ObjectType nor InheritedObjectType are valid.                            |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ACE_OBJECT_TYPE_PRESENT 0x00000001           | ObjectType is valid.                                                             |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ACE_INHERITED_OBJECT_TYPE_PRESENT 0x00000002 | InheritedObjectType is valid. If this value is not specified, all types of child |
	//	|                                              | objects can inherit the ACE.                                                     |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
	// ObjectType (16 bytes): A GUID that identifies a property set, property, extended
	// right, or type of child object. The purpose of this GUID depends on the user rights
	// specified in the Mask field. This field is valid only if the ACE _OBJECT_TYPE_PRESENT
	// bit is set in the Flags field. Otherwise, the ObjectType field is ignored. For information
	// on access rights and for a mapping of the control access rights to the corresponding
	// GUID value that identifies each right, see [MS-ADTS] sections 5.1.3.2 and 5.1.3.2.1.
	//
	// ACCESS_MASK bits are not mutually exclusive. Therefore, the ObjectType field can
	// be set in an ACE with any ACCESS_MASK. If the AccessCheck algorithm calls this ACE
	// and does not find an appropriate GUID, then that ACE will be ignored. For more information
	// on access checks and object access, see [MS-ADTS] section 5.1.3.3.3.
	ObjectType *ACEGUID `idl:"name:ObjectType;switch_is:(Flags 1 &)" json:"object_type"`
	// InheritedObjectType (16 bytes): A GUID that identifies the type of child object that
	// can inherit the ACE. Inheritance is also controlled by the inheritance flags in the
	// ACE_HEADER, as well as by any protection against inheritance placed on the child
	// objects. This field is valid only if the ACE_INHERITED_OBJECT_TYPE_PRESENT bit is
	// set in the Flags member. Otherwise, the InheritedObjectType field is ignored.
	InheritedObjectType *ACEGUID `idl:"name:InheritedObjectType;switch_is:((Flags 2 &) 1 >>)" json:"inherited_object_type"`
	// Sid (variable): The SID of a trustee. The length of the SID MUST be a multiple of
	// 4.
	SID *SID `idl:"name:Sid" json:"sid"`
}

func (o *AccessAllowedObjectACE) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AccessAllowedObjectACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Mask); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.ObjectType != nil {
		_ptr_ObjectType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			_swObjectType := uint32((o.Flags & 1))
			if o.ObjectType != nil {
				if err := o.ObjectType.MarshalUnionNDR(ctx, w, _swObjectType); err != nil {
					return err
				}
			} else {
				if err := (&ACEGUID{}).MarshalUnionNDR(ctx, w, _swObjectType); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ObjectType, _ptr_ObjectType); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.InheritedObjectType != nil {
		_ptr_InheritedObjectType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			_swInheritedObjectType := uint32(((o.Flags & 2) >> 1))
			if o.InheritedObjectType != nil {
				if err := o.InheritedObjectType.MarshalUnionNDR(ctx, w, _swInheritedObjectType); err != nil {
					return err
				}
			} else {
				if err := (&ACEGUID{}).MarshalUnionNDR(ctx, w, _swInheritedObjectType); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.InheritedObjectType, _ptr_InheritedObjectType); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.SID != nil {
		_ptr_Sid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SID != nil {
				if err := o.SID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SID{}).MarshalNDR(ctx, w); err != nil {
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
	return nil
}
func (o *AccessAllowedObjectACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Mask); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_ObjectType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ObjectType == nil {
			o.ObjectType = &ACEGUID{}
		}
		_swObjectType := uint32((o.Flags & 1))
		if err := o.ObjectType.UnmarshalUnionNDR(ctx, w, _swObjectType); err != nil {
			return err
		}
		return nil
	})
	_s_ObjectType := func(ptr interface{}) { o.ObjectType = *ptr.(**ACEGUID) }
	if err := w.ReadPointer(&o.ObjectType, _s_ObjectType, _ptr_ObjectType); err != nil {
		return err
	}
	_ptr_InheritedObjectType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.InheritedObjectType == nil {
			o.InheritedObjectType = &ACEGUID{}
		}
		_swInheritedObjectType := uint32(((o.Flags & 2) >> 1))
		if err := o.InheritedObjectType.UnmarshalUnionNDR(ctx, w, _swInheritedObjectType); err != nil {
			return err
		}
		return nil
	})
	_s_InheritedObjectType := func(ptr interface{}) { o.InheritedObjectType = *ptr.(**ACEGUID) }
	if err := w.ReadPointer(&o.InheritedObjectType, _s_InheritedObjectType, _ptr_InheritedObjectType); err != nil {
		return err
	}
	_ptr_Sid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SID == nil {
			o.SID = &SID{}
		}
		if err := o.SID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Sid := func(ptr interface{}) { o.SID = *ptr.(**SID) }
	if err := w.ReadPointer(&o.SID, _s_Sid, _ptr_Sid); err != nil {
		return err
	}
	return nil
}

// ACEGUID structure represents ACE_GUID RPC union.
type ACEGUID struct {
	// Types that are assignable to Value
	//
	// *ACEGUID_GUID
	// *ACEGUID_DefaultACEGUID
	Value is_ACEGUID `json:"value"`
}

func (o *ACEGUID) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *ACEGUID_GUID:
		if value != nil {
			return value.GUID
		}
	}
	return nil
}

type is_ACEGUID interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_ACEGUID()
}

func (o *ACEGUID) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *ACEGUID_GUID:
		return uint32(1)
	}
	return uint32(0)
}

func (o *ACEGUID) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteSwitch(uint32(sw)); err != nil {
		return err
	}
	switch sw {
	case uint32(1):
		_o, _ := o.Value.(*ACEGUID_GUID)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ACEGUID_GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		// no-op
	}
	return nil
}

func (o *ACEGUID) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadSwitch((*uint32)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint32(1):
		o.Value = &ACEGUID_GUID{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		o.Value = &ACEGUID_DefaultACEGUID{}
	}
	return nil
}

// ACEGUID_GUID structure represents ACE_GUID RPC union arm.
//
// It has following labels: 1
type ACEGUID_GUID struct {
	GUID *GUID `idl:"name:GUID" json:"guid"`
}

func (*ACEGUID_GUID) is_ACEGUID() {}

func (o *ACEGUID_GUID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.GUID != nil {
		if err := o.GUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ACEGUID_GUID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.GUID == nil {
		o.GUID = &GUID{}
	}
	if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ACEGUID_DefaultACEGUID structure represents ACE_GUID RPC default union arm.
type ACEGUID_DefaultACEGUID struct {
}

func (*ACEGUID_DefaultACEGUID) is_ACEGUID() {}

func (o *ACEGUID_DefaultACEGUID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return nil
}
func (o *ACEGUID_DefaultACEGUID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	return nil
}

// AccessDeniedACE structure represents ACCESS_DENIED_ACE RPC structure.
//
// The ACCESS_DENIED_ACE structure defines an ACE for the DACL that controls access
// to an object. An access-denied ACE denies access to an object for a specific trustee
// identified by a SID.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Header                                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Mask                                                                                                                          |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Sid (variable)                                                                                                                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type AccessDeniedACE struct {
	// Mask (4 bytes): An ACCESS_MASK that specifies the user rights denied by this ACE.
	Mask uint32 `idl:"name:Mask" json:"mask"`
	// Sid (variable): The SID of a trustee. The length of the SID MUST be a multiple of
	// 4.
	SID *SID `idl:"name:Sid" json:"sid"`
}

func (o *AccessDeniedACE) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AccessDeniedACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Mask); err != nil {
		return err
	}
	if o.SID != nil {
		_ptr_Sid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SID != nil {
				if err := o.SID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SID{}).MarshalNDR(ctx, w); err != nil {
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
	return nil
}
func (o *AccessDeniedACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Mask); err != nil {
		return err
	}
	_ptr_Sid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SID == nil {
			o.SID = &SID{}
		}
		if err := o.SID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Sid := func(ptr interface{}) { o.SID = *ptr.(**SID) }
	if err := w.ReadPointer(&o.SID, _s_Sid, _ptr_Sid); err != nil {
		return err
	}
	return nil
}

// AccessDeniedObjectACE structure represents ACCESS_DENIED_OBJECT_ACE RPC structure.
//
// The ACCESS_DENIED_OBJECT_ACE structure defines an ACE that controls denied access
// to an object, a property set, or a property. The ACE contains a set of access rights,
// a GUID that identifies the type of object, and a SID that identifies the trustee
// to whom the system will deny access. The ACE also contains a GUID and a set of flags
// that control inheritance of the ACE by child objects.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Header                                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Mask                                                                                                                          |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Flags                                                                                                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ObjectType (16 bytes)                                                                                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| InheritedObjectType (16 bytes)                                                                                                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Sid (variable)                                                                                                                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type AccessDeniedObjectACE struct {
	// Mask (4 bytes): An ACCESS_MASK that specifies the user rights allowed by this ACE.
	//
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	|                                        |                                                                                  |
	//	|                 VALUE                  |                                     MEANING                                      |
	//	|                                        |                                                                                  |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ADS_RIGHT_DS_CONTROL_ACCESS 0X00000100 | The ObjectType GUID identifies an extended access right.                         |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ADS_RIGHT_DS_CREATE_CHILD 0X00000001   | The ObjectType GUID identifies a type of child object. The ACE controls the      |
	//	|                                        | trustee's right to create this type of child object.                             |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ADS_RIGHT_DS_DELETE_CHILD 0X00000002   | The ObjectType GUID identifies a type of child object. The ACE controls the      |
	//	|                                        | trustee's right to delete this type of child object.                             |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ADS_RIGHT_DS_READ_PROP 0x00000010      | The ObjectType GUID identifies a property set or property of the object. The ACE |
	//	|                                        | controls the trustee's right to read the property or property set.               |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ADS_RIGHT_DS_WRITE_PROP 0x00000020     | The ObjectType GUID identifies a property set or property of the object. The ACE |
	//	|                                        | controls the trustee's right to write the property or property set.              |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ADS_RIGHT_DS_SELF 0x00000008           | The ObjectType GUID identifies a validated write.                                |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	Mask uint32 `idl:"name:Mask" json:"mask"`
	// Flags  (4 bytes): A 32-bit unsigned integer that specifies a set of bit flags that
	// indicate whether the ObjectType and InheritedObjectType fields contain valid data.
	// This parameter can be one or more of the following values.
	//
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                              |                                                                                  |
	//	|                    VALUE                     |                                     MEANING                                      |
	//	|                                              |                                                                                  |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000                                   | Neither ObjectType nor InheritedObjectType is valid.                             |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ACE_OBJECT_TYPE_PRESENT 0x00000001           | ObjectType is valid.                                                             |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ACE_INHERITED_OBJECT_TYPE_PRESENT 0x00000002 | InheritedObjectType is valid. If this value is not specified, all types of child |
	//	|                                              | objects can inherit the ACE.                                                     |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
	// ObjectType (16 bytes): A GUID that identifies a property set, a property, an extended
	// right, or a type of child object. The purpose of this GUID depends on the user rights
	// specified in the Mask field. This field is valid only if the ACE _OBJECT_TYPE_PRESENT
	// bit is set in the Flags field. Otherwise, the ObjectType field is ignored. For information
	// about access rights and for a mapping of the control access rights to the corresponding
	// GUID value that identifies each right, see [MS-ADTS] sections 5.1.3.2 and 5.1.3.2.1.
	ObjectType *ACEGUID `idl:"name:ObjectType;switch_is:(Flags 1 &)" json:"object_type"`
	// InheritedObjectType (16 bytes): A GUID that identifies the type of child object that
	// can inherit the ACE. Inheritance is also controlled by the inheritance flags in the
	// ACE_HEADER, as well as by any protection against inheritance placed on the child
	// objects. This field is valid only if the ACE_INHERITED_OBJECT_TYPE_PRESENT bit is
	// set in the Flags member. Otherwise, the InheritedObjectType field is ignored.
	InheritedObjectType *ACEGUID `idl:"name:InheritedObjectType;switch_is:((Flags 2 &) 1 >>)" json:"inherited_object_type"`
	// Sid (variable): The SID of a trustee. The length of the SID MUST be a multiple of
	// 4.
	SID *SID `idl:"name:Sid" json:"sid"`
}

func (o *AccessDeniedObjectACE) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AccessDeniedObjectACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Mask); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.ObjectType != nil {
		_ptr_ObjectType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			_swObjectType := uint32((o.Flags & 1))
			if o.ObjectType != nil {
				if err := o.ObjectType.MarshalUnionNDR(ctx, w, _swObjectType); err != nil {
					return err
				}
			} else {
				if err := (&ACEGUID{}).MarshalUnionNDR(ctx, w, _swObjectType); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ObjectType, _ptr_ObjectType); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.InheritedObjectType != nil {
		_ptr_InheritedObjectType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			_swInheritedObjectType := uint32(((o.Flags & 2) >> 1))
			if o.InheritedObjectType != nil {
				if err := o.InheritedObjectType.MarshalUnionNDR(ctx, w, _swInheritedObjectType); err != nil {
					return err
				}
			} else {
				if err := (&ACEGUID{}).MarshalUnionNDR(ctx, w, _swInheritedObjectType); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.InheritedObjectType, _ptr_InheritedObjectType); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.SID != nil {
		_ptr_Sid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SID != nil {
				if err := o.SID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SID{}).MarshalNDR(ctx, w); err != nil {
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
	return nil
}
func (o *AccessDeniedObjectACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Mask); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_ObjectType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ObjectType == nil {
			o.ObjectType = &ACEGUID{}
		}
		_swObjectType := uint32((o.Flags & 1))
		if err := o.ObjectType.UnmarshalUnionNDR(ctx, w, _swObjectType); err != nil {
			return err
		}
		return nil
	})
	_s_ObjectType := func(ptr interface{}) { o.ObjectType = *ptr.(**ACEGUID) }
	if err := w.ReadPointer(&o.ObjectType, _s_ObjectType, _ptr_ObjectType); err != nil {
		return err
	}
	_ptr_InheritedObjectType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.InheritedObjectType == nil {
			o.InheritedObjectType = &ACEGUID{}
		}
		_swInheritedObjectType := uint32(((o.Flags & 2) >> 1))
		if err := o.InheritedObjectType.UnmarshalUnionNDR(ctx, w, _swInheritedObjectType); err != nil {
			return err
		}
		return nil
	})
	_s_InheritedObjectType := func(ptr interface{}) { o.InheritedObjectType = *ptr.(**ACEGUID) }
	if err := w.ReadPointer(&o.InheritedObjectType, _s_InheritedObjectType, _ptr_InheritedObjectType); err != nil {
		return err
	}
	_ptr_Sid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SID == nil {
			o.SID = &SID{}
		}
		if err := o.SID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Sid := func(ptr interface{}) { o.SID = *ptr.(**SID) }
	if err := w.ReadPointer(&o.SID, _s_Sid, _ptr_Sid); err != nil {
		return err
	}
	return nil
}

// AccessAllowedCallbackACE structure represents ACCESS_ALLOWED_CALLBACK_ACE RPC structure.
//
// The ACCESS_ALLOWED_CALLBACK_ACE structure defines an ACE for the DACL that controls
// access to an object. An access-allowed ACE allows access to an object for a specific
// trustee identified by a SID.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Header                                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Mask                                                                                                                          |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Sid (variable)                                                                                                                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ApplicationData (variable)                                                                                                    |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type AccessAllowedCallbackACE struct {
	// Mask (4 bytes): An ACCESS_MASK that specifies the user rights allowed by this ACE.
	Mask uint32 `idl:"name:Mask" json:"mask"`
	// Sid (variable): The SID of a trustee. The length of the SID MUST be a multiple of
	// 4.
	SID *SID `idl:"name:Sid" json:"sid"`
	// ApplicationData (variable): Optional application data. The size of the application
	// data is determined by the AceSize field of the ACE_HEADER.
	ApplicationData []byte `idl:"name:ApplicationData" json:"application_data"`
}

func (o *AccessAllowedCallbackACE) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AccessAllowedCallbackACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Mask); err != nil {
		return err
	}
	if o.SID != nil {
		_ptr_Sid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SID != nil {
				if err := o.SID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SID{}).MarshalNDR(ctx, w); err != nil {
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
	if o.ApplicationData != nil {
		_ptr_ApplicationData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			for i1 := range o.ApplicationData {
				i1 := i1
				if err := w.WriteData(o.ApplicationData[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ApplicationData, _ptr_ApplicationData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *AccessAllowedCallbackACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Mask); err != nil {
		return err
	}
	_ptr_Sid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SID == nil {
			o.SID = &SID{}
		}
		if err := o.SID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Sid := func(ptr interface{}) { o.SID = *ptr.(**SID) }
	if err := w.ReadPointer(&o.SID, _s_Sid, _ptr_Sid); err != nil {
		return err
	}
	_ptr_ApplicationData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		for i1 := 0; w.Len() > 0; i1++ {
			i1 := i1
			o.ApplicationData = append(o.ApplicationData, uint8(0))
			if err := w.ReadData(&o.ApplicationData[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ApplicationData := func(ptr interface{}) { o.ApplicationData = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.ApplicationData, _s_ApplicationData, _ptr_ApplicationData); err != nil {
		return err
	}
	return nil
}

// AccessDeniedCallbackACE structure represents ACCESS_DENIED_CALLBACK_ACE RPC structure.
//
// The ACCESS_DENIED_CALLBACK_ACE structure defines an ACE for the DACL that controls
// access to an object. An access-denied ACE denies access to an object for a specific
// trustee identified by a SID.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Header                                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Mask                                                                                                                          |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Sid (variable)                                                                                                                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ApplicationData (variable)                                                                                                    |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type AccessDeniedCallbackACE struct {
	// Mask (4 bytes): An ACCESS_MASK that specifies the user rights denied by this ACE.
	Mask uint32 `idl:"name:Mask" json:"mask"`
	// Sid (variable): The SID of a trustee. The length of the SID MUST be a multiple of
	// 4.
	SID *SID `idl:"name:Sid" json:"sid"`
	// ApplicationData (variable): Optional application data. The size of the application
	// data is determined by the AceSize field of the ACE_HEADER.
	ApplicationData []byte `idl:"name:ApplicationData" json:"application_data"`
}

func (o *AccessDeniedCallbackACE) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AccessDeniedCallbackACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Mask); err != nil {
		return err
	}
	if o.SID != nil {
		_ptr_Sid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SID != nil {
				if err := o.SID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SID{}).MarshalNDR(ctx, w); err != nil {
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
	if o.ApplicationData != nil {
		_ptr_ApplicationData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			for i1 := range o.ApplicationData {
				i1 := i1
				if err := w.WriteData(o.ApplicationData[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ApplicationData, _ptr_ApplicationData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *AccessDeniedCallbackACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Mask); err != nil {
		return err
	}
	_ptr_Sid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SID == nil {
			o.SID = &SID{}
		}
		if err := o.SID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Sid := func(ptr interface{}) { o.SID = *ptr.(**SID) }
	if err := w.ReadPointer(&o.SID, _s_Sid, _ptr_Sid); err != nil {
		return err
	}
	_ptr_ApplicationData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		for i1 := 0; w.Len() > 0; i1++ {
			i1 := i1
			o.ApplicationData = append(o.ApplicationData, uint8(0))
			if err := w.ReadData(&o.ApplicationData[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ApplicationData := func(ptr interface{}) { o.ApplicationData = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.ApplicationData, _s_ApplicationData, _ptr_ApplicationData); err != nil {
		return err
	}
	return nil
}

// AccessAllowedCallbackObjectACE structure represents ACCESS_ALLOWED_CALLBACK_OBJECT_ACE RPC structure.
//
// The ACCESS_ALLOWED_CALLBACK_OBJECT_ACE structure defines an ACE that controls allowed
// access to an object, property set, or property. The ACE contains a set of user rights,
// a GUID that identifies the type of object, and a SID that identifies the trustee
// to whom the system will grant access. The ACE also contains a GUID and a set of flags
// that control inheritance of the ACE by child objects.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Header                                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Mask                                                                                                                          |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Flags                                                                                                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ObjectType (16 bytes)                                                                                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| InheritedObjectType (16 bytes)                                                                                                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Sid (variable)                                                                                                                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ApplicationData (variable)                                                                                                    |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type AccessAllowedCallbackObjectACE struct {
	// Mask (4 bytes): An ACCESS_MASK structure that specifies the user rights allowed by
	// this ACE.
	//
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	|                                        |                                                                                  |
	//	|                 VALUE                  |                                     MEANING                                      |
	//	|                                        |                                                                                  |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ADS_RIGHT_DS_CONTROL_ACCESS 0X00000100 | The ObjectType GUID identifies an extended access right.                         |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ADS_RIGHT_DS_CREATE_CHILD 0X00000001   | The ObjectType GUID identifies a type of child object. The ACE controls the      |
	//	|                                        | trustee's right to create this type of child object.                             |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ADS_RIGHT_DS_READ_PROP 0x00000010      | The ObjectType GUID identifies a property set or property of the object. The ACE |
	//	|                                        | controls the trustee's right to read the property or property set.               |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ADS_RIGHT_DS_WRITE_PROP 0x00000020     | The ObjectType GUID identifies a property set or property of the object. The ACE |
	//	|                                        | controls the trustee's right to write the property or property set.              |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ADS_RIGHT_DS_SELF 0x00000008           | The ObjectType GUID identifies a validated write.                                |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	Mask uint32 `idl:"name:Mask" json:"mask"`
	// Flags (4 bytes): A 32-bit unsigned integer that specifies a set of bit flags that
	// indicate whether the ObjectType and InheritedObjectType fields contain valid data.
	// This parameter can be one or more of the following values.
	//
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                              |                                                                                  |
	//	|                    VALUE                     |                                     MEANING                                      |
	//	|                                              |                                                                                  |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000                                   | Neither ObjectType nor InheritedObjectType are valid.                            |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ACE_OBJECT_TYPE_PRESENT 0x00000001           | ObjectType is valid.                                                             |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ACE_INHERITED_OBJECT_TYPE_PRESENT 0x00000002 | InheritedObjectType is valid. If this value is not specified, all types of child |
	//	|                                              | objects can inherit the ACE.                                                     |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
	// ObjectType (16 bytes): A GUID that identifies a property set, property, extended
	// right, or type of child object. The purpose of this GUID depends on the user rights
	// specified in the Mask field. This field is valid only if the ACE _OBJECT_TYPE_PRESENT
	// bit is set in the Flags field. Otherwise, the ObjectType field is ignored.
	ObjectType *ACEGUID `idl:"name:ObjectType;switch_is:(Flags 1 &)" json:"object_type"`
	// InheritedObjectType (16 bytes): A GUID that identifies the type of child object that
	// can inherit the ACE. Inheritance is also controlled by the inheritance flags in the
	// ACE_HEADER, as well as by any protection against inheritance placed on the child
	// objects. This field is valid only if the ACE_INHERITED_OBJECT_TYPE_PRESENT bit is
	// set in the Flags member. Otherwise, the InheritedObjectType field is ignored.
	InheritedObjectType *ACEGUID `idl:"name:InheritedObjectType;switch_is:((Flags 2 &) 1 >>)" json:"inherited_object_type"`
	// Sid (variable): The SID of a trustee. The length of the SID MUST be a multiple of
	// 4.
	SID *SID `idl:"name:Sid" json:"sid"`
	// ApplicationData (variable): Optional application data. The size of the application
	// data is determined by the AceSize field of the ACE_HEADER.
	ApplicationData []byte `idl:"name:ApplicationData" json:"application_data"`
}

func (o *AccessAllowedCallbackObjectACE) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AccessAllowedCallbackObjectACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Mask); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.ObjectType != nil {
		_ptr_ObjectType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			_swObjectType := uint32((o.Flags & 1))
			if o.ObjectType != nil {
				if err := o.ObjectType.MarshalUnionNDR(ctx, w, _swObjectType); err != nil {
					return err
				}
			} else {
				if err := (&ACEGUID{}).MarshalUnionNDR(ctx, w, _swObjectType); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ObjectType, _ptr_ObjectType); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.InheritedObjectType != nil {
		_ptr_InheritedObjectType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			_swInheritedObjectType := uint32(((o.Flags & 2) >> 1))
			if o.InheritedObjectType != nil {
				if err := o.InheritedObjectType.MarshalUnionNDR(ctx, w, _swInheritedObjectType); err != nil {
					return err
				}
			} else {
				if err := (&ACEGUID{}).MarshalUnionNDR(ctx, w, _swInheritedObjectType); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.InheritedObjectType, _ptr_InheritedObjectType); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.SID != nil {
		_ptr_Sid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SID != nil {
				if err := o.SID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SID{}).MarshalNDR(ctx, w); err != nil {
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
	if o.ApplicationData != nil {
		_ptr_ApplicationData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			for i1 := range o.ApplicationData {
				i1 := i1
				if err := w.WriteData(o.ApplicationData[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ApplicationData, _ptr_ApplicationData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *AccessAllowedCallbackObjectACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Mask); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_ObjectType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ObjectType == nil {
			o.ObjectType = &ACEGUID{}
		}
		_swObjectType := uint32((o.Flags & 1))
		if err := o.ObjectType.UnmarshalUnionNDR(ctx, w, _swObjectType); err != nil {
			return err
		}
		return nil
	})
	_s_ObjectType := func(ptr interface{}) { o.ObjectType = *ptr.(**ACEGUID) }
	if err := w.ReadPointer(&o.ObjectType, _s_ObjectType, _ptr_ObjectType); err != nil {
		return err
	}
	_ptr_InheritedObjectType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.InheritedObjectType == nil {
			o.InheritedObjectType = &ACEGUID{}
		}
		_swInheritedObjectType := uint32(((o.Flags & 2) >> 1))
		if err := o.InheritedObjectType.UnmarshalUnionNDR(ctx, w, _swInheritedObjectType); err != nil {
			return err
		}
		return nil
	})
	_s_InheritedObjectType := func(ptr interface{}) { o.InheritedObjectType = *ptr.(**ACEGUID) }
	if err := w.ReadPointer(&o.InheritedObjectType, _s_InheritedObjectType, _ptr_InheritedObjectType); err != nil {
		return err
	}
	_ptr_Sid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SID == nil {
			o.SID = &SID{}
		}
		if err := o.SID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Sid := func(ptr interface{}) { o.SID = *ptr.(**SID) }
	if err := w.ReadPointer(&o.SID, _s_Sid, _ptr_Sid); err != nil {
		return err
	}
	_ptr_ApplicationData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		for i1 := 0; w.Len() > 0; i1++ {
			i1 := i1
			o.ApplicationData = append(o.ApplicationData, uint8(0))
			if err := w.ReadData(&o.ApplicationData[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ApplicationData := func(ptr interface{}) { o.ApplicationData = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.ApplicationData, _s_ApplicationData, _ptr_ApplicationData); err != nil {
		return err
	}
	return nil
}

// AccessDeniedCallbackObjectACE structure represents ACCESS_DENIED_CALLBACK_OBJECT_ACE RPC structure.
//
// The ACCESS_DENIED_CALLBACK_OBJECT_ACE structure defines an ACE that controls denied
// access to an object, a property set, or property. The ACE contains a set of user
// rights, a GUID that identifies the type of object, and a SID that identifies the
// trustee to whom the system will deny access. The ACE also contains a GUID and a set
// of flags that control inheritance of the ACE by child objects.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Header                                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Mask                                                                                                                          |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Flags                                                                                                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ObjectType (16 bytes)                                                                                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| InheritedObjectType (16 bytes)                                                                                                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Sid (variable)                                                                                                                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ApplicationData (variable)                                                                                                    |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type AccessDeniedCallbackObjectACE struct {
	// Mask (4 bytes): An ACCESS_MASK structure that specifies the user rights denied by
	// this ACE.
	//
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	|                                        |                                                                                  |
	//	|                 VALUE                  |                                     MEANING                                      |
	//	|                                        |                                                                                  |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ADS_RIGHT_DS_CONTROL_ACCESS 0X00000100 | The ObjectType GUID identifies an extended access right.                         |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ADS_RIGHT_DS_CREATE_CHILD 0X00000001   | The ObjectType GUID identifies a type of child object. The ACE controls the      |
	//	|                                        | trustee's right to create this type of child object.                             |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ADS_RIGHT_DS_READ_PROP 0x00000010      | The ObjectType GUID identifies a property set or property of the object. The ACE |
	//	|                                        | controls the trustee's right to read the property or property set.               |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ADS_RIGHT_DS_WRITE_PROP 0x00000020     | The ObjectType GUID identifies a property set or property of the object. The ACE |
	//	|                                        | controls the trustee's right to write the property or property set.              |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ADS_RIGHT_DS_SELF 0x00000008           | The ObjectType GUID identifies a validated write.                                |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	Mask uint32 `idl:"name:Mask" json:"mask"`
	// Flags (4 bytes): A 32-bit unsigned integer that specifies a set of bit flags that
	// indicate whether the ObjectType and InheritedObjectType fields contain valid data.
	// This parameter can be one or more of the following values.
	//
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                              |                                                                                  |
	//	|                    VALUE                     |                                     MEANING                                      |
	//	|                                              |                                                                                  |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000                                   | Neither ObjectType nor InheritedObjectType are valid.                            |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ACE_OBJECT_TYPE_PRESENT 0x00000001           | ObjectType is valid.                                                             |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ACE_INHERITED_OBJECT_TYPE_PRESENT 0x00000002 | InheritedObjectType is valid. If this value is not specified, all types of child |
	//	|                                              | objects can inherit the ACE.                                                     |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
	// ObjectType (16 bytes): A GUID that identifies a property set, property, extended
	// right, or type of child object. The purpose of this GUID depends on the user rights
	// specified in the Mask field. This field is valid only if the ACE _OBJECT_TYPE_PRESENT
	// bit is set in the Flags field. Otherwise, the ObjectType field is ignored.
	ObjectType *ACEGUID `idl:"name:ObjectType;switch_is:(Flags 1 &)" json:"object_type"`
	// InheritedObjectType (16 bytes): A GUID that identifies the type of child object that
	// can inherit the ACE. Inheritance is also controlled by the inheritance flags in the
	// ACE_HEADER, as well as by any protection against inheritance placed on the child
	// objects. This field is valid only if the ACE_INHERITED_OBJECT_TYPE_PRESENT bit is
	// set in the Flags member. Otherwise, the InheritedObjectType field is ignored.
	InheritedObjectType *ACEGUID `idl:"name:InheritedObjectType;switch_is:((Flags 2 &) 1 >>)" json:"inherited_object_type"`
	// Sid (variable): The SID of a trustee. The length of the SID MUST be a multiple of
	// 4.
	SID *SID `idl:"name:Sid" json:"sid"`
	// ApplicationData (variable): Optional application data. The size of the application
	// data is determined by the AceSize field of the ACE_HEADER.
	ApplicationData []byte `idl:"name:ApplicationData" json:"application_data"`
}

func (o *AccessDeniedCallbackObjectACE) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AccessDeniedCallbackObjectACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Mask); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.ObjectType != nil {
		_ptr_ObjectType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			_swObjectType := uint32((o.Flags & 1))
			if o.ObjectType != nil {
				if err := o.ObjectType.MarshalUnionNDR(ctx, w, _swObjectType); err != nil {
					return err
				}
			} else {
				if err := (&ACEGUID{}).MarshalUnionNDR(ctx, w, _swObjectType); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ObjectType, _ptr_ObjectType); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.InheritedObjectType != nil {
		_ptr_InheritedObjectType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			_swInheritedObjectType := uint32(((o.Flags & 2) >> 1))
			if o.InheritedObjectType != nil {
				if err := o.InheritedObjectType.MarshalUnionNDR(ctx, w, _swInheritedObjectType); err != nil {
					return err
				}
			} else {
				if err := (&ACEGUID{}).MarshalUnionNDR(ctx, w, _swInheritedObjectType); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.InheritedObjectType, _ptr_InheritedObjectType); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.SID != nil {
		_ptr_Sid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SID != nil {
				if err := o.SID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SID{}).MarshalNDR(ctx, w); err != nil {
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
	if o.ApplicationData != nil {
		_ptr_ApplicationData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			for i1 := range o.ApplicationData {
				i1 := i1
				if err := w.WriteData(o.ApplicationData[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ApplicationData, _ptr_ApplicationData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *AccessDeniedCallbackObjectACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Mask); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_ObjectType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ObjectType == nil {
			o.ObjectType = &ACEGUID{}
		}
		_swObjectType := uint32((o.Flags & 1))
		if err := o.ObjectType.UnmarshalUnionNDR(ctx, w, _swObjectType); err != nil {
			return err
		}
		return nil
	})
	_s_ObjectType := func(ptr interface{}) { o.ObjectType = *ptr.(**ACEGUID) }
	if err := w.ReadPointer(&o.ObjectType, _s_ObjectType, _ptr_ObjectType); err != nil {
		return err
	}
	_ptr_InheritedObjectType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.InheritedObjectType == nil {
			o.InheritedObjectType = &ACEGUID{}
		}
		_swInheritedObjectType := uint32(((o.Flags & 2) >> 1))
		if err := o.InheritedObjectType.UnmarshalUnionNDR(ctx, w, _swInheritedObjectType); err != nil {
			return err
		}
		return nil
	})
	_s_InheritedObjectType := func(ptr interface{}) { o.InheritedObjectType = *ptr.(**ACEGUID) }
	if err := w.ReadPointer(&o.InheritedObjectType, _s_InheritedObjectType, _ptr_InheritedObjectType); err != nil {
		return err
	}
	_ptr_Sid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SID == nil {
			o.SID = &SID{}
		}
		if err := o.SID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Sid := func(ptr interface{}) { o.SID = *ptr.(**SID) }
	if err := w.ReadPointer(&o.SID, _s_Sid, _ptr_Sid); err != nil {
		return err
	}
	_ptr_ApplicationData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		for i1 := 0; w.Len() > 0; i1++ {
			i1 := i1
			o.ApplicationData = append(o.ApplicationData, uint8(0))
			if err := w.ReadData(&o.ApplicationData[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ApplicationData := func(ptr interface{}) { o.ApplicationData = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.ApplicationData, _s_ApplicationData, _ptr_ApplicationData); err != nil {
		return err
	}
	return nil
}

// SystemAuditACE structure represents SYSTEM_AUDIT_ACE RPC structure.
//
// The SYSTEM_AUDIT_ACE structure defines an access ACE for the system access control
// list (SACL) that specifies what types of access cause system-level notifications.
// A system-audit ACE causes an audit message to be logged when a specified trustee
// attempts to gain access to an object. The trustee is identified by a SID.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Header                                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Mask                                                                                                                          |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Sid (variable)                                                                                                                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type SystemAuditACE struct {
	// Mask (4 bytes): An ACCESS_MASK structure that specifies the user rights that cause
	// audit messages to be generated.
	Mask uint32 `idl:"name:Mask" json:"mask"`
	// Sid (variable): The SID of a trustee. The length of the SID MUST be a multiple of
	// 4. An access attempt of a kind specified by the Mask field by any trustee whose SID
	// matches the Sid field causes the system to generate an audit message. If an application
	// does not specify a SID for this field, audit messages are generated for the specified
	// access rights for all trustees.
	SID *SID `idl:"name:Sid" json:"sid"`
}

func (o *SystemAuditACE) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SystemAuditACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Mask); err != nil {
		return err
	}
	if o.SID != nil {
		_ptr_Sid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SID != nil {
				if err := o.SID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SID{}).MarshalNDR(ctx, w); err != nil {
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
	return nil
}
func (o *SystemAuditACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Mask); err != nil {
		return err
	}
	_ptr_Sid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SID == nil {
			o.SID = &SID{}
		}
		if err := o.SID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Sid := func(ptr interface{}) { o.SID = *ptr.(**SID) }
	if err := w.ReadPointer(&o.SID, _s_Sid, _ptr_Sid); err != nil {
		return err
	}
	return nil
}

// SystemAuditObjectACE structure represents SYSTEM_AUDIT_OBJECT_ACE RPC structure.
type SystemAuditObjectACE struct {
	Mask                uint32   `idl:"name:Mask" json:"mask"`
	Flags               uint32   `idl:"name:Flags" json:"flags"`
	ObjectType          *ACEGUID `idl:"name:ObjectType;switch_is:(Flags 1 &)" json:"object_type"`
	InheritedObjectType *ACEGUID `idl:"name:InheritedObjectType;switch_is:((Flags 2 &) 1 >>)" json:"inherited_object_type"`
	SID                 *SID     `idl:"name:Sid" json:"sid"`
	ApplicationData     []byte   `idl:"name:ApplicationData" json:"application_data"`
}

func (o *SystemAuditObjectACE) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SystemAuditObjectACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Mask); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.ObjectType != nil {
		_ptr_ObjectType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			_swObjectType := uint32((o.Flags & 1))
			if o.ObjectType != nil {
				if err := o.ObjectType.MarshalUnionNDR(ctx, w, _swObjectType); err != nil {
					return err
				}
			} else {
				if err := (&ACEGUID{}).MarshalUnionNDR(ctx, w, _swObjectType); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ObjectType, _ptr_ObjectType); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.InheritedObjectType != nil {
		_ptr_InheritedObjectType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			_swInheritedObjectType := uint32(((o.Flags & 2) >> 1))
			if o.InheritedObjectType != nil {
				if err := o.InheritedObjectType.MarshalUnionNDR(ctx, w, _swInheritedObjectType); err != nil {
					return err
				}
			} else {
				if err := (&ACEGUID{}).MarshalUnionNDR(ctx, w, _swInheritedObjectType); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.InheritedObjectType, _ptr_InheritedObjectType); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.SID != nil {
		_ptr_Sid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SID != nil {
				if err := o.SID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SID{}).MarshalNDR(ctx, w); err != nil {
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
	if o.ApplicationData != nil {
		_ptr_ApplicationData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			for i1 := range o.ApplicationData {
				i1 := i1
				if err := w.WriteData(o.ApplicationData[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ApplicationData, _ptr_ApplicationData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SystemAuditObjectACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Mask); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_ObjectType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ObjectType == nil {
			o.ObjectType = &ACEGUID{}
		}
		_swObjectType := uint32((o.Flags & 1))
		if err := o.ObjectType.UnmarshalUnionNDR(ctx, w, _swObjectType); err != nil {
			return err
		}
		return nil
	})
	_s_ObjectType := func(ptr interface{}) { o.ObjectType = *ptr.(**ACEGUID) }
	if err := w.ReadPointer(&o.ObjectType, _s_ObjectType, _ptr_ObjectType); err != nil {
		return err
	}
	_ptr_InheritedObjectType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.InheritedObjectType == nil {
			o.InheritedObjectType = &ACEGUID{}
		}
		_swInheritedObjectType := uint32(((o.Flags & 2) >> 1))
		if err := o.InheritedObjectType.UnmarshalUnionNDR(ctx, w, _swInheritedObjectType); err != nil {
			return err
		}
		return nil
	})
	_s_InheritedObjectType := func(ptr interface{}) { o.InheritedObjectType = *ptr.(**ACEGUID) }
	if err := w.ReadPointer(&o.InheritedObjectType, _s_InheritedObjectType, _ptr_InheritedObjectType); err != nil {
		return err
	}
	_ptr_Sid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SID == nil {
			o.SID = &SID{}
		}
		if err := o.SID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Sid := func(ptr interface{}) { o.SID = *ptr.(**SID) }
	if err := w.ReadPointer(&o.SID, _s_Sid, _ptr_Sid); err != nil {
		return err
	}
	_ptr_ApplicationData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		for i1 := 0; w.Len() > 0; i1++ {
			i1 := i1
			o.ApplicationData = append(o.ApplicationData, uint8(0))
			if err := w.ReadData(&o.ApplicationData[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ApplicationData := func(ptr interface{}) { o.ApplicationData = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.ApplicationData, _s_ApplicationData, _ptr_ApplicationData); err != nil {
		return err
	}
	return nil
}

// SystemAuditCallbackACE structure represents SYSTEM_AUDIT_CALLBACK_ACE RPC structure.
//
// The SYSTEM_AUDIT_CALLBACK_ACE structure defines an ACE for the SACL that specifies
// what types of access cause system-level notifications. A system-audit ACE causes
// an audit message to be logged when a specified trustee attempts to gain access to
// an object. The trustee is identified by a SID.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Header                                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Mask                                                                                                                          |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Sid (variable)                                                                                                                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ApplicationData (variable)                                                                                                    |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type SystemAuditCallbackACE struct {
	// Mask (4 bytes): An ACCESS_MASK structure that specifies the user rights that cause
	// audit messages to be generated.
	Mask uint32 `idl:"name:Mask" json:"mask"`
	// Sid (variable): The SID of a trustee. The length of the SID MUST be a multiple of
	// 4. An access attempt of a kind specified by the Mask field by any trustee whose SID
	// matches the Sid field causes the system to generate an audit message. If an application
	// does not specify a SID for this field, audit messages are generated for the specified
	// access rights for all trustees.
	SID *SID `idl:"name:Sid" json:"sid"`
	// ApplicationData (variable): Optional application data. The size of the application
	// data is determined by the AceSize field of the ACE_HEADER.
	ApplicationData []byte `idl:"name:ApplicationData" json:"application_data"`
}

func (o *SystemAuditCallbackACE) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SystemAuditCallbackACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Mask); err != nil {
		return err
	}
	if o.SID != nil {
		_ptr_Sid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SID != nil {
				if err := o.SID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SID{}).MarshalNDR(ctx, w); err != nil {
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
	if o.ApplicationData != nil {
		_ptr_ApplicationData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			for i1 := range o.ApplicationData {
				i1 := i1
				if err := w.WriteData(o.ApplicationData[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ApplicationData, _ptr_ApplicationData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SystemAuditCallbackACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Mask); err != nil {
		return err
	}
	_ptr_Sid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SID == nil {
			o.SID = &SID{}
		}
		if err := o.SID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Sid := func(ptr interface{}) { o.SID = *ptr.(**SID) }
	if err := w.ReadPointer(&o.SID, _s_Sid, _ptr_Sid); err != nil {
		return err
	}
	_ptr_ApplicationData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		for i1 := 0; w.Len() > 0; i1++ {
			i1 := i1
			o.ApplicationData = append(o.ApplicationData, uint8(0))
			if err := w.ReadData(&o.ApplicationData[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ApplicationData := func(ptr interface{}) { o.ApplicationData = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.ApplicationData, _s_ApplicationData, _ptr_ApplicationData); err != nil {
		return err
	}
	return nil
}

// SystemMandatoryLabelACE structure represents SYSTEM_MANDATORY_LABEL_ACE RPC structure.
//
// The SYSTEM_MANDATORY_LABEL_ACE structure defines an ACE for the SACL that specifies
// the mandatory access level and policy for a securable object.<55>
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Header                                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Mask                                                                                                                          |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Sid (variable)                                                                                                                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//
// The RPC representation of the SYSTEM_MANDATORY_LABEL_ACE type defines an access control
// entry (ACE) for the system access control list (SACL) that specifies the mandatory
// access level and policy for a securable object.
type SystemMandatoryLabelACE struct {
	// Mask (4 bytes): An ACCESS_MASK structure that specifies the access policy for principals
	// with a mandatory integrity level lower than the object associated with the SACL that
	// contains this ACE.
	//
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                                 |                                                                                  |
	//	|                      VALUE                      |                                     MEANING                                      |
	//	|                                                 |                                                                                  |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| SYSTEM_MANDATORY_LABEL_NO_WRITE_UP 0x00000001   | A principal with a lower mandatory level than the object cannot write to the     |
	//	|                                                 | object.                                                                          |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| SYSTEM_MANDATORY_LABEL_NO_READ_UP 0x00000002    | A principal with a lower mandatory level than the object cannot read the object. |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| SYSTEM_MANDATORY_LABEL_NO_EXECUTE_UP 0x00000004 | A principal with a lower mandatory level than the object cannot execute the      |
	//	|                                                 | object.                                                                          |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Mask:  An ACCESS_MASK as specified in section 2.4.4.13.
	Mask uint32 `idl:"name:Mask" json:"mask"`
	// Sid (variable): The SID of a trustee. The length of the SID MUST be a multiple of
	// 4. The identifier authority of the SID must be SECURITY_MANDATORY_LABEL_AUTHORITY.
	// The RID of the SID specifies the mandatory integrity level of the object associated
	// with the SACL that contains this ACE. The RID must be one of the following values.
	//
	//	+------------+------------------------------------+
	//	|            |                                    |
	//	|   VALUE    |              MEANING               |
	//	|            |                                    |
	//	+------------+------------------------------------+
	//	+------------+------------------------------------+
	//	| 0x00000000 | Untrusted integrity level.         |
	//	+------------+------------------------------------+
	//	| 0x00001000 | Low integrity level.               |
	//	+------------+------------------------------------+
	//	| 0x00002000 | Medium integrity level.            |
	//	+------------+------------------------------------+
	//	| 0x00003000 | High integrity level.              |
	//	+------------+------------------------------------+
	//	| 0x00004000 | System integrity level.            |
	//	+------------+------------------------------------+
	//	| 0x00005000 | Protected process integrity level. |
	//	+------------+------------------------------------+
	SID *SID `idl:"name:Sid" json:"sid"`
}

func (o *SystemMandatoryLabelACE) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SystemMandatoryLabelACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Mask); err != nil {
		return err
	}
	if o.SID != nil {
		_ptr_Sid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SID != nil {
				if err := o.SID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SID{}).MarshalNDR(ctx, w); err != nil {
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
	return nil
}
func (o *SystemMandatoryLabelACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Mask); err != nil {
		return err
	}
	_ptr_Sid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SID == nil {
			o.SID = &SID{}
		}
		if err := o.SID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Sid := func(ptr interface{}) { o.SID = *ptr.(**SID) }
	if err := w.ReadPointer(&o.SID, _s_Sid, _ptr_Sid); err != nil {
		return err
	}
	return nil
}

// SystemAuditCallbackObjectACE structure represents SYSTEM_AUDIT_CALLBACK_OBJECT_ACE RPC structure.
type SystemAuditCallbackObjectACE struct {
	Mask                uint32   `idl:"name:Mask" json:"mask"`
	Flags               uint32   `idl:"name:Flags" json:"flags"`
	ObjectType          *ACEGUID `idl:"name:ObjectType;switch_is:(Flags 1 &)" json:"object_type"`
	InheritedObjectType *ACEGUID `idl:"name:InheritedObjectType;switch_is:((Flags 2 &) 1 >>)" json:"inherited_object_type"`
	SID                 *SID     `idl:"name:Sid" json:"sid"`
	ApplicationData     []byte   `idl:"name:ApplicationData" json:"application_data"`
}

func (o *SystemAuditCallbackObjectACE) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SystemAuditCallbackObjectACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Mask); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.ObjectType != nil {
		_ptr_ObjectType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			_swObjectType := uint32((o.Flags & 1))
			if o.ObjectType != nil {
				if err := o.ObjectType.MarshalUnionNDR(ctx, w, _swObjectType); err != nil {
					return err
				}
			} else {
				if err := (&ACEGUID{}).MarshalUnionNDR(ctx, w, _swObjectType); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ObjectType, _ptr_ObjectType); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.InheritedObjectType != nil {
		_ptr_InheritedObjectType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			_swInheritedObjectType := uint32(((o.Flags & 2) >> 1))
			if o.InheritedObjectType != nil {
				if err := o.InheritedObjectType.MarshalUnionNDR(ctx, w, _swInheritedObjectType); err != nil {
					return err
				}
			} else {
				if err := (&ACEGUID{}).MarshalUnionNDR(ctx, w, _swInheritedObjectType); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.InheritedObjectType, _ptr_InheritedObjectType); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.SID != nil {
		_ptr_Sid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SID != nil {
				if err := o.SID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SID{}).MarshalNDR(ctx, w); err != nil {
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
	if o.ApplicationData != nil {
		_ptr_ApplicationData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			for i1 := range o.ApplicationData {
				i1 := i1
				if err := w.WriteData(o.ApplicationData[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ApplicationData, _ptr_ApplicationData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SystemAuditCallbackObjectACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Mask); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_ObjectType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ObjectType == nil {
			o.ObjectType = &ACEGUID{}
		}
		_swObjectType := uint32((o.Flags & 1))
		if err := o.ObjectType.UnmarshalUnionNDR(ctx, w, _swObjectType); err != nil {
			return err
		}
		return nil
	})
	_s_ObjectType := func(ptr interface{}) { o.ObjectType = *ptr.(**ACEGUID) }
	if err := w.ReadPointer(&o.ObjectType, _s_ObjectType, _ptr_ObjectType); err != nil {
		return err
	}
	_ptr_InheritedObjectType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.InheritedObjectType == nil {
			o.InheritedObjectType = &ACEGUID{}
		}
		_swInheritedObjectType := uint32(((o.Flags & 2) >> 1))
		if err := o.InheritedObjectType.UnmarshalUnionNDR(ctx, w, _swInheritedObjectType); err != nil {
			return err
		}
		return nil
	})
	_s_InheritedObjectType := func(ptr interface{}) { o.InheritedObjectType = *ptr.(**ACEGUID) }
	if err := w.ReadPointer(&o.InheritedObjectType, _s_InheritedObjectType, _ptr_InheritedObjectType); err != nil {
		return err
	}
	_ptr_Sid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SID == nil {
			o.SID = &SID{}
		}
		if err := o.SID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Sid := func(ptr interface{}) { o.SID = *ptr.(**SID) }
	if err := w.ReadPointer(&o.SID, _s_Sid, _ptr_Sid); err != nil {
		return err
	}
	_ptr_ApplicationData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		for i1 := 0; w.Len() > 0; i1++ {
			i1 := i1
			o.ApplicationData = append(o.ApplicationData, uint8(0))
			if err := w.ReadData(&o.ApplicationData[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ApplicationData := func(ptr interface{}) { o.ApplicationData = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.ApplicationData, _s_ApplicationData, _ptr_ApplicationData); err != nil {
		return err
	}
	return nil
}

// SystemResourceAttributeACE structure represents SYSTEM_RESOURCE_ATTRIBUTE_ACE RPC structure.
//
// The SYSTEM_RESOURCE_ATTRIBUTE_ACE structure defines an ACE for the specification
// of a resource attribute associated with an object. A SYSTEM_RESOURCE_ATTRIBUTE_ACE
// is used in conditional ACEs in specifying access or audit policy for the resource.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Header                                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Mask                                                                                                                          |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Sid (variable)                                                                                                                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Attribute Data (variable)                                                                                                     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type SystemResourceAttributeACE struct {
	// Mask (4 bytes): An ACCESS_MASK that MUST be set to zero.
	Mask uint32 `idl:"name:Mask" json:"mask"`
	// Sid (variable): The SID corresponding to the Everyone SID (S-1-1-0) in binary form.
	SID           *SID   `idl:"name:Sid" json:"sid"`
	AttributeData []byte `idl:"name:AttributeData" json:"attribute_data"`
}

func (o *SystemResourceAttributeACE) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SystemResourceAttributeACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Mask); err != nil {
		return err
	}
	if o.SID != nil {
		_ptr_Sid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SID != nil {
				if err := o.SID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SID{}).MarshalNDR(ctx, w); err != nil {
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
	if o.AttributeData != nil {
		_ptr_AttributeData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			for i1 := range o.AttributeData {
				i1 := i1
				if err := w.WriteData(o.AttributeData[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.AttributeData, _ptr_AttributeData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SystemResourceAttributeACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Mask); err != nil {
		return err
	}
	_ptr_Sid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SID == nil {
			o.SID = &SID{}
		}
		if err := o.SID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Sid := func(ptr interface{}) { o.SID = *ptr.(**SID) }
	if err := w.ReadPointer(&o.SID, _s_Sid, _ptr_Sid); err != nil {
		return err
	}
	_ptr_AttributeData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		for i1 := 0; w.Len() > 0; i1++ {
			i1 := i1
			o.AttributeData = append(o.AttributeData, uint8(0))
			if err := w.ReadData(&o.AttributeData[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_AttributeData := func(ptr interface{}) { o.AttributeData = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.AttributeData, _s_AttributeData, _ptr_AttributeData); err != nil {
		return err
	}
	return nil
}

// SystemScopedPolicyIDACE structure represents SYSTEM_SCOPED_POLICY_ID_ACE RPC structure.
//
// The SYSTEM_SCOPED_POLICY_ID_ACE structure defines an ACE for the purpose of applying
// a central access policy to the resource.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Header                                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Mask                                                                                                                          |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Sid (variable)                                                                                                                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type SystemScopedPolicyIDACE struct {
	// Mask (4 bytes): An ACCESS_MASK that MUST be set to zero.
	Mask uint32 `idl:"name:Mask" json:"mask"`
	// Sid (variable): A SID that identifies a central access policy. For a SYSTEM_SCOPED_POLICY_ID_ACE
	// to be applicable on a resource, this SID MUST match a CAPID of a CentralAccessPolicy
	// contained in the CentralAccessPoliciesList (as specified in [MS-GPCAP] section 3.2.1.1)
	// of the machine on which the access evaluation will be performed.
	SID *SID `idl:"name:Sid" json:"sid"`
}

func (o *SystemScopedPolicyIDACE) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SystemScopedPolicyIDACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Mask); err != nil {
		return err
	}
	if o.SID != nil {
		_ptr_Sid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SID != nil {
				if err := o.SID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SID{}).MarshalNDR(ctx, w); err != nil {
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
	return nil
}
func (o *SystemScopedPolicyIDACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Mask); err != nil {
		return err
	}
	_ptr_Sid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SID == nil {
			o.SID = &SID{}
		}
		if err := o.SID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Sid := func(ptr interface{}) { o.SID = *ptr.(**SID) }
	if err := w.ReadPointer(&o.SID, _s_Sid, _ptr_Sid); err != nil {
		return err
	}
	return nil
}

// RawACE structure represents RAW_ACE RPC structure.
type RawACE struct {
	RawData []byte `idl:"name:RawData" json:"raw_data"`
}

func (o *RawACE) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RawACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.RawData != nil {
		_ptr_RawData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			for i1 := range o.RawData {
				i1 := i1
				if err := w.WriteData(o.RawData[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RawData, _ptr_RawData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *RawACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_RawData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		for i1 := 0; w.Len() > 0; i1++ {
			i1 := i1
			o.RawData = append(o.RawData, uint8(0))
			if err := w.ReadData(&o.RawData[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_RawData := func(ptr interface{}) { o.RawData = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.RawData, _s_RawData, _ptr_RawData); err != nil {
		return err
	}
	return nil
}

// ACEType type represents ACE_TYPE RPC enumeration.
type ACEType uint16

var (
	ACETypeAccessAllowedACEType               ACEType = 0
	ACETypeAccessDeniedACEType                ACEType = 1
	ACETypeSystemAuditACEType                 ACEType = 2
	ACETypeSystemAlarmACEType                 ACEType = 3
	ACETypeAccessAllowedCompoundACEType       ACEType = 4
	ACETypeAccessAllowedObjectACEType         ACEType = 5
	ACETypeAccessDeniedObjectACEType          ACEType = 6
	ACETypeSystemAuditObjectACEType           ACEType = 7
	ACETypeSystemAlarmObjectACEType           ACEType = 8
	ACETypeAccessAllowedCallbackACEType       ACEType = 9
	ACETypeAccessDeniedCallbackACEType        ACEType = 10
	ACETypeAccessAllowedCallbackObjectACEType ACEType = 11
	ACETypeAccessDeniedCallbackObjectACEType  ACEType = 12
	ACETypeSystemAuditCallbackACEType         ACEType = 13
	ACETypeSystemAlarmCallbackACEType         ACEType = 14
	ACETypeSystemAuditCallbackObjectACEType   ACEType = 15
	ACETypeSystemAlarmCallbackObjectACEType   ACEType = 16
	ACETypeSystemMandatoryLabelACEType        ACEType = 17
	ACETypeSystemResourceAttributeACEType     ACEType = 18
	ACETypeSystemScopedPolicyIdaceType        ACEType = 19
)

func (o ACEType) String() string {
	switch o {
	case ACETypeAccessAllowedACEType:
		return "ACETypeAccessAllowedACEType"
	case ACETypeAccessDeniedACEType:
		return "ACETypeAccessDeniedACEType"
	case ACETypeSystemAuditACEType:
		return "ACETypeSystemAuditACEType"
	case ACETypeSystemAlarmACEType:
		return "ACETypeSystemAlarmACEType"
	case ACETypeAccessAllowedCompoundACEType:
		return "ACETypeAccessAllowedCompoundACEType"
	case ACETypeAccessAllowedObjectACEType:
		return "ACETypeAccessAllowedObjectACEType"
	case ACETypeAccessDeniedObjectACEType:
		return "ACETypeAccessDeniedObjectACEType"
	case ACETypeSystemAuditObjectACEType:
		return "ACETypeSystemAuditObjectACEType"
	case ACETypeSystemAlarmObjectACEType:
		return "ACETypeSystemAlarmObjectACEType"
	case ACETypeAccessAllowedCallbackACEType:
		return "ACETypeAccessAllowedCallbackACEType"
	case ACETypeAccessDeniedCallbackACEType:
		return "ACETypeAccessDeniedCallbackACEType"
	case ACETypeAccessAllowedCallbackObjectACEType:
		return "ACETypeAccessAllowedCallbackObjectACEType"
	case ACETypeAccessDeniedCallbackObjectACEType:
		return "ACETypeAccessDeniedCallbackObjectACEType"
	case ACETypeSystemAuditCallbackACEType:
		return "ACETypeSystemAuditCallbackACEType"
	case ACETypeSystemAlarmCallbackACEType:
		return "ACETypeSystemAlarmCallbackACEType"
	case ACETypeSystemAuditCallbackObjectACEType:
		return "ACETypeSystemAuditCallbackObjectACEType"
	case ACETypeSystemAlarmCallbackObjectACEType:
		return "ACETypeSystemAlarmCallbackObjectACEType"
	case ACETypeSystemMandatoryLabelACEType:
		return "ACETypeSystemMandatoryLabelACEType"
	case ACETypeSystemResourceAttributeACEType:
		return "ACETypeSystemResourceAttributeACEType"
	case ACETypeSystemScopedPolicyIdaceType:
		return "ACETypeSystemScopedPolicyIdaceType"
	}
	return "Invalid"
}

// ACE structure represents ACE RPC structure.
type ACE struct {
	ACEType  uint8    `idl:"name:AceType" json:"ace_type"`
	ACEFlags uint8    `idl:"name:AceFlags" json:"ace_flags"`
	ACESize  uint16   `idl:"name:AceSize" json:"ace_size"`
	Data     []byte   `idl:"name:Data;size_is:((AceSize-4))" json:"data"`
	ACEData  *ACEData `idl:"name:AceData;switch_is:AceType" json:"ace_data"`
}

func (o *ACE) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.ACESize == 0 {
		o.ACESize = uint16((len(o.Data) + 4))
	}
	if o.ACESize < 4 {
		o.ACESize = 4
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteData(o.ACEType); err != nil {
		return err
	}
	if err := w.WriteData(o.ACEFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.ACESize); err != nil {
		return err
	}
	if o.Data != nil || (o.ACESize-4) > 0 {
		_ptr_Data := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64((o.ACESize - 4))
			if o.ACESize < 4 {
				dimSize1 = uint64(0)
			}
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
	// pad 4
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	return nil
}
func (o *ACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadData(&o.ACEType); err != nil {
		return err
	}
	if err := w.ReadData(&o.ACEFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.ACESize); err != nil {
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
		if o.ACESize > 4 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64((o.ACESize - 4))
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
		_layout_AceData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_AceData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.ACEData == nil {
					o.ACEData = &ACEData{}
				}
				_swACEData := uint8(o.ACEType)
				if err := o.ACEData.UnmarshalUnionNDR(ctx, w, _swACEData); err != nil {
					return err
				}
				return nil
			})
			_s_AceData := func(ptr interface{}) { o.ACEData = *ptr.(**ACEData) }
			if err := w.ReadPointer(&o.ACEData, _s_AceData, _ptr_AceData); err != nil {
				return err
			}
			return nil
		})
		if len(o.Data) > 0 {
			if err := w.WithBytes(o.Data).Unmarshal(ctx, _layout_AceData); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Data := func(ptr interface{}) { o.Data = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Data, _s_Data, _ptr_Data); err != nil {
		return err
	}
	// pad 4
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	return nil
}

// ACEData structure represents ACE_DATA RPC union.
type ACEData struct {
	// Types that are assignable to Value
	//
	// *ACEData_AccessAllowedACE
	// *ACEData_AccessDeniedACE
	// *ACEData_SystemAuditACE
	// *ACEData_AccessAllowedObjectACE
	// *ACEData_AccessDeniedObjectACE
	// *ACEData_SystemAuditObjectACE
	// *ACEData_AccessAllowedCallbackACE
	// *ACEData_AccessDeniedCallbackACE
	// *ACEData_AccessAllowedCallbackObjectACE
	// *ACEData_AccessDeniedCallbackObjectACE
	// *ACEData_SystemAuditCallbackACE
	// *ACEData_SystemAuditCallbackObjectACE
	// *ACEData_SystemMandatoryLabelACE
	// *ACEData_SystemResourceAttributeACE
	// *ACEData_SystemScopedPolicyIDACE
	// *ACEData_RawACE
	Value is_ACEData `json:"value"`
}

func (o *ACEData) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *ACEData_AccessAllowedACE:
		if value != nil {
			return value.AccessAllowedACE
		}
	case *ACEData_AccessDeniedACE:
		if value != nil {
			return value.AccessDeniedACE
		}
	case *ACEData_SystemAuditACE:
		if value != nil {
			return value.SystemAuditACE
		}
	case *ACEData_AccessAllowedObjectACE:
		if value != nil {
			return value.AccessAllowedObjectACE
		}
	case *ACEData_AccessDeniedObjectACE:
		if value != nil {
			return value.AccessDeniedObjectACE
		}
	case *ACEData_SystemAuditObjectACE:
		if value != nil {
			return value.SystemAuditObjectACE
		}
	case *ACEData_AccessAllowedCallbackACE:
		if value != nil {
			return value.AccessAllowedCallbackACE
		}
	case *ACEData_AccessDeniedCallbackACE:
		if value != nil {
			return value.AccessDeniedCallbackACE
		}
	case *ACEData_AccessAllowedCallbackObjectACE:
		if value != nil {
			return value.AccessAllowedCallbackObjectACE
		}
	case *ACEData_AccessDeniedCallbackObjectACE:
		if value != nil {
			return value.AccessDeniedCallbackObjectACE
		}
	case *ACEData_SystemAuditCallbackACE:
		if value != nil {
			return value.SystemAuditCallbackACE
		}
	case *ACEData_SystemAuditCallbackObjectACE:
		if value != nil {
			return value.SystemAuditCallbackObjectACE
		}
	case *ACEData_SystemMandatoryLabelACE:
		if value != nil {
			return value.SystemMandatoryLabelACE
		}
	case *ACEData_SystemResourceAttributeACE:
		if value != nil {
			return value.SystemResourceAttributeACE
		}
	case *ACEData_SystemScopedPolicyIDACE:
		if value != nil {
			return value.SystemScopedPolicyIDACE
		}
	case *ACEData_RawACE:
		if value != nil {
			return value.RawACE
		}
	}
	return nil
}

type is_ACEData interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_ACEData()
}

func (o *ACEData) NDRSwitchValue(sw uint8) uint8 {
	if o == nil {
		return uint8(0)
	}
	switch (interface{})(o.Value).(type) {
	case *ACEData_AccessAllowedACE:
		return uint8(0)
	case *ACEData_AccessDeniedACE:
		return uint8(1)
	case *ACEData_SystemAuditACE:
		return uint8(2)
	case *ACEData_AccessAllowedObjectACE:
		return uint8(5)
	case *ACEData_AccessDeniedObjectACE:
		return uint8(6)
	case *ACEData_SystemAuditObjectACE:
		return uint8(7)
	case *ACEData_AccessAllowedCallbackACE:
		return uint8(9)
	case *ACEData_AccessDeniedCallbackACE:
		return uint8(10)
	case *ACEData_AccessAllowedCallbackObjectACE:
		return uint8(11)
	case *ACEData_AccessDeniedCallbackObjectACE:
		return uint8(12)
	case *ACEData_SystemAuditCallbackACE:
		return uint8(13)
	case *ACEData_SystemAuditCallbackObjectACE:
		return uint8(15)
	case *ACEData_SystemMandatoryLabelACE:
		return uint8(17)
	case *ACEData_SystemResourceAttributeACE:
		return uint8(18)
	case *ACEData_SystemScopedPolicyIDACE:
		return uint8(19)
	}
	return uint8(0)
}

func (o *ACEData) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint8) error {
	if err := w.WriteSwitch(uint8(sw)); err != nil {
		return err
	}
	switch sw {
	case uint8(0):
		_o, _ := o.Value.(*ACEData_AccessAllowedACE)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ACEData_AccessAllowedACE{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(1):
		_o, _ := o.Value.(*ACEData_AccessDeniedACE)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ACEData_AccessDeniedACE{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(2):
		_o, _ := o.Value.(*ACEData_SystemAuditACE)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ACEData_SystemAuditACE{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(5):
		_o, _ := o.Value.(*ACEData_AccessAllowedObjectACE)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ACEData_AccessAllowedObjectACE{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(6):
		_o, _ := o.Value.(*ACEData_AccessDeniedObjectACE)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ACEData_AccessDeniedObjectACE{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(7):
		_o, _ := o.Value.(*ACEData_SystemAuditObjectACE)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ACEData_SystemAuditObjectACE{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(9):
		_o, _ := o.Value.(*ACEData_AccessAllowedCallbackACE)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ACEData_AccessAllowedCallbackACE{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(10):
		_o, _ := o.Value.(*ACEData_AccessDeniedCallbackACE)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ACEData_AccessDeniedCallbackACE{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(11):
		_o, _ := o.Value.(*ACEData_AccessAllowedCallbackObjectACE)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ACEData_AccessAllowedCallbackObjectACE{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(12):
		_o, _ := o.Value.(*ACEData_AccessDeniedCallbackObjectACE)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ACEData_AccessDeniedCallbackObjectACE{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(13):
		_o, _ := o.Value.(*ACEData_SystemAuditCallbackACE)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ACEData_SystemAuditCallbackACE{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(15):
		_o, _ := o.Value.(*ACEData_SystemAuditCallbackObjectACE)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ACEData_SystemAuditCallbackObjectACE{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(17):
		_o, _ := o.Value.(*ACEData_SystemMandatoryLabelACE)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ACEData_SystemMandatoryLabelACE{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(18):
		_o, _ := o.Value.(*ACEData_SystemResourceAttributeACE)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ACEData_SystemResourceAttributeACE{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(19):
		_o, _ := o.Value.(*ACEData_SystemScopedPolicyIDACE)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ACEData_SystemScopedPolicyIDACE{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		_o, _ := o.Value.(*ACEData_RawACE)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ACEData_RawACE{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *ACEData) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint8) error {
	if err := w.ReadSwitch((*uint8)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint8(0):
		o.Value = &ACEData_AccessAllowedACE{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(1):
		o.Value = &ACEData_AccessDeniedACE{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(2):
		o.Value = &ACEData_SystemAuditACE{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(5):
		o.Value = &ACEData_AccessAllowedObjectACE{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(6):
		o.Value = &ACEData_AccessDeniedObjectACE{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(7):
		o.Value = &ACEData_SystemAuditObjectACE{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(9):
		o.Value = &ACEData_AccessAllowedCallbackACE{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(10):
		o.Value = &ACEData_AccessDeniedCallbackACE{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(11):
		o.Value = &ACEData_AccessAllowedCallbackObjectACE{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(12):
		o.Value = &ACEData_AccessDeniedCallbackObjectACE{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(13):
		o.Value = &ACEData_SystemAuditCallbackACE{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(15):
		o.Value = &ACEData_SystemAuditCallbackObjectACE{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(17):
		o.Value = &ACEData_SystemMandatoryLabelACE{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(18):
		o.Value = &ACEData_SystemResourceAttributeACE{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(19):
		o.Value = &ACEData_SystemScopedPolicyIDACE{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		o.Value = &ACEData_RawACE{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// ACEData_AccessAllowedACE structure represents ACE_DATA RPC union arm.
//
// It has following labels: 0
type ACEData_AccessAllowedACE struct {
	AccessAllowedACE *AccessAllowedACE `idl:"name:AccessAllowedAce" json:"access_allowed_ace"`
}

func (*ACEData_AccessAllowedACE) is_ACEData() {}

func (o *ACEData_AccessAllowedACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.AccessAllowedACE != nil {
		_ptr_AccessAllowedAce := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.AccessAllowedACE != nil {
				if err := o.AccessAllowedACE.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AccessAllowedACE{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.AccessAllowedACE, _ptr_AccessAllowedAce); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ACEData_AccessAllowedACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_AccessAllowedAce := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.AccessAllowedACE == nil {
			o.AccessAllowedACE = &AccessAllowedACE{}
		}
		if err := o.AccessAllowedACE.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_AccessAllowedAce := func(ptr interface{}) { o.AccessAllowedACE = *ptr.(**AccessAllowedACE) }
	if err := w.ReadPointer(&o.AccessAllowedACE, _s_AccessAllowedAce, _ptr_AccessAllowedAce); err != nil {
		return err
	}
	return nil
}

// ACEData_AccessDeniedACE structure represents ACE_DATA RPC union arm.
//
// It has following labels: 1
type ACEData_AccessDeniedACE struct {
	AccessDeniedACE *AccessDeniedACE `idl:"name:AccessDeniedAce" json:"access_denied_ace"`
}

func (*ACEData_AccessDeniedACE) is_ACEData() {}

func (o *ACEData_AccessDeniedACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.AccessDeniedACE != nil {
		_ptr_AccessDeniedAce := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.AccessDeniedACE != nil {
				if err := o.AccessDeniedACE.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AccessDeniedACE{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.AccessDeniedACE, _ptr_AccessDeniedAce); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ACEData_AccessDeniedACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_AccessDeniedAce := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.AccessDeniedACE == nil {
			o.AccessDeniedACE = &AccessDeniedACE{}
		}
		if err := o.AccessDeniedACE.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_AccessDeniedAce := func(ptr interface{}) { o.AccessDeniedACE = *ptr.(**AccessDeniedACE) }
	if err := w.ReadPointer(&o.AccessDeniedACE, _s_AccessDeniedAce, _ptr_AccessDeniedAce); err != nil {
		return err
	}
	return nil
}

// ACEData_SystemAuditACE structure represents ACE_DATA RPC union arm.
//
// It has following labels: 2
type ACEData_SystemAuditACE struct {
	SystemAuditACE *SystemAuditACE `idl:"name:SystemAuditAce" json:"system_audit_ace"`
}

func (*ACEData_SystemAuditACE) is_ACEData() {}

func (o *ACEData_SystemAuditACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.SystemAuditACE != nil {
		_ptr_SystemAuditAce := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SystemAuditACE != nil {
				if err := o.SystemAuditACE.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SystemAuditACE{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SystemAuditACE, _ptr_SystemAuditAce); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ACEData_SystemAuditACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_SystemAuditAce := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SystemAuditACE == nil {
			o.SystemAuditACE = &SystemAuditACE{}
		}
		if err := o.SystemAuditACE.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_SystemAuditAce := func(ptr interface{}) { o.SystemAuditACE = *ptr.(**SystemAuditACE) }
	if err := w.ReadPointer(&o.SystemAuditACE, _s_SystemAuditAce, _ptr_SystemAuditAce); err != nil {
		return err
	}
	return nil
}

// ACEData_AccessAllowedObjectACE structure represents ACE_DATA RPC union arm.
//
// It has following labels: 5
type ACEData_AccessAllowedObjectACE struct {
	AccessAllowedObjectACE *AccessAllowedObjectACE `idl:"name:AccessAllowedObjectAce" json:"access_allowed_object_ace"`
}

func (*ACEData_AccessAllowedObjectACE) is_ACEData() {}

func (o *ACEData_AccessAllowedObjectACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.AccessAllowedObjectACE != nil {
		_ptr_AccessAllowedObjectAce := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.AccessAllowedObjectACE != nil {
				if err := o.AccessAllowedObjectACE.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AccessAllowedObjectACE{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.AccessAllowedObjectACE, _ptr_AccessAllowedObjectAce); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ACEData_AccessAllowedObjectACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_AccessAllowedObjectAce := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.AccessAllowedObjectACE == nil {
			o.AccessAllowedObjectACE = &AccessAllowedObjectACE{}
		}
		if err := o.AccessAllowedObjectACE.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_AccessAllowedObjectAce := func(ptr interface{}) { o.AccessAllowedObjectACE = *ptr.(**AccessAllowedObjectACE) }
	if err := w.ReadPointer(&o.AccessAllowedObjectACE, _s_AccessAllowedObjectAce, _ptr_AccessAllowedObjectAce); err != nil {
		return err
	}
	return nil
}

// ACEData_AccessDeniedObjectACE structure represents ACE_DATA RPC union arm.
//
// It has following labels: 6
type ACEData_AccessDeniedObjectACE struct {
	AccessDeniedObjectACE *AccessDeniedObjectACE `idl:"name:AccessDeniedObjectAce" json:"access_denied_object_ace"`
}

func (*ACEData_AccessDeniedObjectACE) is_ACEData() {}

func (o *ACEData_AccessDeniedObjectACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.AccessDeniedObjectACE != nil {
		_ptr_AccessDeniedObjectAce := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.AccessDeniedObjectACE != nil {
				if err := o.AccessDeniedObjectACE.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AccessDeniedObjectACE{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.AccessDeniedObjectACE, _ptr_AccessDeniedObjectAce); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ACEData_AccessDeniedObjectACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_AccessDeniedObjectAce := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.AccessDeniedObjectACE == nil {
			o.AccessDeniedObjectACE = &AccessDeniedObjectACE{}
		}
		if err := o.AccessDeniedObjectACE.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_AccessDeniedObjectAce := func(ptr interface{}) { o.AccessDeniedObjectACE = *ptr.(**AccessDeniedObjectACE) }
	if err := w.ReadPointer(&o.AccessDeniedObjectACE, _s_AccessDeniedObjectAce, _ptr_AccessDeniedObjectAce); err != nil {
		return err
	}
	return nil
}

// ACEData_SystemAuditObjectACE structure represents ACE_DATA RPC union arm.
//
// It has following labels: 7
type ACEData_SystemAuditObjectACE struct {
	SystemAuditObjectACE *SystemAuditObjectACE `idl:"name:SystemAuditObjectAce" json:"system_audit_object_ace"`
}

func (*ACEData_SystemAuditObjectACE) is_ACEData() {}

func (o *ACEData_SystemAuditObjectACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.SystemAuditObjectACE != nil {
		_ptr_SystemAuditObjectAce := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SystemAuditObjectACE != nil {
				if err := o.SystemAuditObjectACE.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SystemAuditObjectACE{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SystemAuditObjectACE, _ptr_SystemAuditObjectAce); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ACEData_SystemAuditObjectACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_SystemAuditObjectAce := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SystemAuditObjectACE == nil {
			o.SystemAuditObjectACE = &SystemAuditObjectACE{}
		}
		if err := o.SystemAuditObjectACE.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_SystemAuditObjectAce := func(ptr interface{}) { o.SystemAuditObjectACE = *ptr.(**SystemAuditObjectACE) }
	if err := w.ReadPointer(&o.SystemAuditObjectACE, _s_SystemAuditObjectAce, _ptr_SystemAuditObjectAce); err != nil {
		return err
	}
	return nil
}

// ACEData_AccessAllowedCallbackACE structure represents ACE_DATA RPC union arm.
//
// It has following labels: 9
type ACEData_AccessAllowedCallbackACE struct {
	AccessAllowedCallbackACE *AccessAllowedCallbackACE `idl:"name:AccessAllowedCallbackAce" json:"access_allowed_callback_ace"`
}

func (*ACEData_AccessAllowedCallbackACE) is_ACEData() {}

func (o *ACEData_AccessAllowedCallbackACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.AccessAllowedCallbackACE != nil {
		_ptr_AccessAllowedCallbackAce := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.AccessAllowedCallbackACE != nil {
				if err := o.AccessAllowedCallbackACE.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AccessAllowedCallbackACE{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.AccessAllowedCallbackACE, _ptr_AccessAllowedCallbackAce); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ACEData_AccessAllowedCallbackACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_AccessAllowedCallbackAce := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.AccessAllowedCallbackACE == nil {
			o.AccessAllowedCallbackACE = &AccessAllowedCallbackACE{}
		}
		if err := o.AccessAllowedCallbackACE.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_AccessAllowedCallbackAce := func(ptr interface{}) { o.AccessAllowedCallbackACE = *ptr.(**AccessAllowedCallbackACE) }
	if err := w.ReadPointer(&o.AccessAllowedCallbackACE, _s_AccessAllowedCallbackAce, _ptr_AccessAllowedCallbackAce); err != nil {
		return err
	}
	return nil
}

// ACEData_AccessDeniedCallbackACE structure represents ACE_DATA RPC union arm.
//
// It has following labels: 10
type ACEData_AccessDeniedCallbackACE struct {
	AccessDeniedCallbackACE *AccessDeniedCallbackACE `idl:"name:AccessDeniedCallbackAce" json:"access_denied_callback_ace"`
}

func (*ACEData_AccessDeniedCallbackACE) is_ACEData() {}

func (o *ACEData_AccessDeniedCallbackACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.AccessDeniedCallbackACE != nil {
		_ptr_AccessDeniedCallbackAce := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.AccessDeniedCallbackACE != nil {
				if err := o.AccessDeniedCallbackACE.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AccessDeniedCallbackACE{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.AccessDeniedCallbackACE, _ptr_AccessDeniedCallbackAce); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ACEData_AccessDeniedCallbackACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_AccessDeniedCallbackAce := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.AccessDeniedCallbackACE == nil {
			o.AccessDeniedCallbackACE = &AccessDeniedCallbackACE{}
		}
		if err := o.AccessDeniedCallbackACE.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_AccessDeniedCallbackAce := func(ptr interface{}) { o.AccessDeniedCallbackACE = *ptr.(**AccessDeniedCallbackACE) }
	if err := w.ReadPointer(&o.AccessDeniedCallbackACE, _s_AccessDeniedCallbackAce, _ptr_AccessDeniedCallbackAce); err != nil {
		return err
	}
	return nil
}

// ACEData_AccessAllowedCallbackObjectACE structure represents ACE_DATA RPC union arm.
//
// It has following labels: 11
type ACEData_AccessAllowedCallbackObjectACE struct {
	AccessAllowedCallbackObjectACE *AccessAllowedCallbackObjectACE `idl:"name:AccessAllowedCallbackObjectAce" json:"access_allowed_callback_object_ace"`
}

func (*ACEData_AccessAllowedCallbackObjectACE) is_ACEData() {}

func (o *ACEData_AccessAllowedCallbackObjectACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.AccessAllowedCallbackObjectACE != nil {
		_ptr_AccessAllowedCallbackObjectAce := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.AccessAllowedCallbackObjectACE != nil {
				if err := o.AccessAllowedCallbackObjectACE.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AccessAllowedCallbackObjectACE{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.AccessAllowedCallbackObjectACE, _ptr_AccessAllowedCallbackObjectAce); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ACEData_AccessAllowedCallbackObjectACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_AccessAllowedCallbackObjectAce := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.AccessAllowedCallbackObjectACE == nil {
			o.AccessAllowedCallbackObjectACE = &AccessAllowedCallbackObjectACE{}
		}
		if err := o.AccessAllowedCallbackObjectACE.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_AccessAllowedCallbackObjectAce := func(ptr interface{}) { o.AccessAllowedCallbackObjectACE = *ptr.(**AccessAllowedCallbackObjectACE) }
	if err := w.ReadPointer(&o.AccessAllowedCallbackObjectACE, _s_AccessAllowedCallbackObjectAce, _ptr_AccessAllowedCallbackObjectAce); err != nil {
		return err
	}
	return nil
}

// ACEData_AccessDeniedCallbackObjectACE structure represents ACE_DATA RPC union arm.
//
// It has following labels: 12
type ACEData_AccessDeniedCallbackObjectACE struct {
	AccessDeniedCallbackObjectACE *AccessDeniedCallbackObjectACE `idl:"name:AccessDeniedCallbackObjectAce" json:"access_denied_callback_object_ace"`
}

func (*ACEData_AccessDeniedCallbackObjectACE) is_ACEData() {}

func (o *ACEData_AccessDeniedCallbackObjectACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.AccessDeniedCallbackObjectACE != nil {
		_ptr_AccessDeniedCallbackObjectAce := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.AccessDeniedCallbackObjectACE != nil {
				if err := o.AccessDeniedCallbackObjectACE.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AccessDeniedCallbackObjectACE{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.AccessDeniedCallbackObjectACE, _ptr_AccessDeniedCallbackObjectAce); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ACEData_AccessDeniedCallbackObjectACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_AccessDeniedCallbackObjectAce := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.AccessDeniedCallbackObjectACE == nil {
			o.AccessDeniedCallbackObjectACE = &AccessDeniedCallbackObjectACE{}
		}
		if err := o.AccessDeniedCallbackObjectACE.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_AccessDeniedCallbackObjectAce := func(ptr interface{}) { o.AccessDeniedCallbackObjectACE = *ptr.(**AccessDeniedCallbackObjectACE) }
	if err := w.ReadPointer(&o.AccessDeniedCallbackObjectACE, _s_AccessDeniedCallbackObjectAce, _ptr_AccessDeniedCallbackObjectAce); err != nil {
		return err
	}
	return nil
}

// ACEData_SystemAuditCallbackACE structure represents ACE_DATA RPC union arm.
//
// It has following labels: 13
type ACEData_SystemAuditCallbackACE struct {
	SystemAuditCallbackACE *SystemAuditCallbackACE `idl:"name:SystemAuditCallbackAce" json:"system_audit_callback_ace"`
}

func (*ACEData_SystemAuditCallbackACE) is_ACEData() {}

func (o *ACEData_SystemAuditCallbackACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.SystemAuditCallbackACE != nil {
		_ptr_SystemAuditCallbackAce := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SystemAuditCallbackACE != nil {
				if err := o.SystemAuditCallbackACE.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SystemAuditCallbackACE{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SystemAuditCallbackACE, _ptr_SystemAuditCallbackAce); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ACEData_SystemAuditCallbackACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_SystemAuditCallbackAce := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SystemAuditCallbackACE == nil {
			o.SystemAuditCallbackACE = &SystemAuditCallbackACE{}
		}
		if err := o.SystemAuditCallbackACE.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_SystemAuditCallbackAce := func(ptr interface{}) { o.SystemAuditCallbackACE = *ptr.(**SystemAuditCallbackACE) }
	if err := w.ReadPointer(&o.SystemAuditCallbackACE, _s_SystemAuditCallbackAce, _ptr_SystemAuditCallbackAce); err != nil {
		return err
	}
	return nil
}

// ACEData_SystemAuditCallbackObjectACE structure represents ACE_DATA RPC union arm.
//
// It has following labels: 15
type ACEData_SystemAuditCallbackObjectACE struct {
	SystemAuditCallbackObjectACE *SystemAuditCallbackObjectACE `idl:"name:SystemAuditCallbackObjectAce" json:"system_audit_callback_object_ace"`
}

func (*ACEData_SystemAuditCallbackObjectACE) is_ACEData() {}

func (o *ACEData_SystemAuditCallbackObjectACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.SystemAuditCallbackObjectACE != nil {
		_ptr_SystemAuditCallbackObjectAce := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SystemAuditCallbackObjectACE != nil {
				if err := o.SystemAuditCallbackObjectACE.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SystemAuditCallbackObjectACE{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SystemAuditCallbackObjectACE, _ptr_SystemAuditCallbackObjectAce); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ACEData_SystemAuditCallbackObjectACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_SystemAuditCallbackObjectAce := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SystemAuditCallbackObjectACE == nil {
			o.SystemAuditCallbackObjectACE = &SystemAuditCallbackObjectACE{}
		}
		if err := o.SystemAuditCallbackObjectACE.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_SystemAuditCallbackObjectAce := func(ptr interface{}) { o.SystemAuditCallbackObjectACE = *ptr.(**SystemAuditCallbackObjectACE) }
	if err := w.ReadPointer(&o.SystemAuditCallbackObjectACE, _s_SystemAuditCallbackObjectAce, _ptr_SystemAuditCallbackObjectAce); err != nil {
		return err
	}
	return nil
}

// ACEData_SystemMandatoryLabelACE structure represents ACE_DATA RPC union arm.
//
// It has following labels: 17
type ACEData_SystemMandatoryLabelACE struct {
	SystemMandatoryLabelACE *SystemMandatoryLabelACE `idl:"name:SystemMandatoryLabelAce" json:"system_mandatory_label_ace"`
}

func (*ACEData_SystemMandatoryLabelACE) is_ACEData() {}

func (o *ACEData_SystemMandatoryLabelACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.SystemMandatoryLabelACE != nil {
		_ptr_SystemMandatoryLabelAce := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SystemMandatoryLabelACE != nil {
				if err := o.SystemMandatoryLabelACE.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SystemMandatoryLabelACE{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SystemMandatoryLabelACE, _ptr_SystemMandatoryLabelAce); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ACEData_SystemMandatoryLabelACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_SystemMandatoryLabelAce := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SystemMandatoryLabelACE == nil {
			o.SystemMandatoryLabelACE = &SystemMandatoryLabelACE{}
		}
		if err := o.SystemMandatoryLabelACE.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_SystemMandatoryLabelAce := func(ptr interface{}) { o.SystemMandatoryLabelACE = *ptr.(**SystemMandatoryLabelACE) }
	if err := w.ReadPointer(&o.SystemMandatoryLabelACE, _s_SystemMandatoryLabelAce, _ptr_SystemMandatoryLabelAce); err != nil {
		return err
	}
	return nil
}

// ACEData_SystemResourceAttributeACE structure represents ACE_DATA RPC union arm.
//
// It has following labels: 18
type ACEData_SystemResourceAttributeACE struct {
	SystemResourceAttributeACE *SystemResourceAttributeACE `idl:"name:SystemResourceAttributeAce" json:"system_resource_attribute_ace"`
}

func (*ACEData_SystemResourceAttributeACE) is_ACEData() {}

func (o *ACEData_SystemResourceAttributeACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.SystemResourceAttributeACE != nil {
		_ptr_SystemResourceAttributeAce := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SystemResourceAttributeACE != nil {
				if err := o.SystemResourceAttributeACE.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SystemResourceAttributeACE{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SystemResourceAttributeACE, _ptr_SystemResourceAttributeAce); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ACEData_SystemResourceAttributeACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_SystemResourceAttributeAce := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SystemResourceAttributeACE == nil {
			o.SystemResourceAttributeACE = &SystemResourceAttributeACE{}
		}
		if err := o.SystemResourceAttributeACE.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_SystemResourceAttributeAce := func(ptr interface{}) { o.SystemResourceAttributeACE = *ptr.(**SystemResourceAttributeACE) }
	if err := w.ReadPointer(&o.SystemResourceAttributeACE, _s_SystemResourceAttributeAce, _ptr_SystemResourceAttributeAce); err != nil {
		return err
	}
	return nil
}

// ACEData_SystemScopedPolicyIDACE structure represents ACE_DATA RPC union arm.
//
// It has following labels: 19
type ACEData_SystemScopedPolicyIDACE struct {
	SystemScopedPolicyIDACE *SystemScopedPolicyIDACE `idl:"name:SystemScopedPolicyIdAce" json:"system_scoped_policy_id_ace"`
}

func (*ACEData_SystemScopedPolicyIDACE) is_ACEData() {}

func (o *ACEData_SystemScopedPolicyIDACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.SystemScopedPolicyIDACE != nil {
		_ptr_SystemScopedPolicyIdAce := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SystemScopedPolicyIDACE != nil {
				if err := o.SystemScopedPolicyIDACE.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SystemScopedPolicyIDACE{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SystemScopedPolicyIDACE, _ptr_SystemScopedPolicyIdAce); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ACEData_SystemScopedPolicyIDACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_SystemScopedPolicyIdAce := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SystemScopedPolicyIDACE == nil {
			o.SystemScopedPolicyIDACE = &SystemScopedPolicyIDACE{}
		}
		if err := o.SystemScopedPolicyIDACE.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_SystemScopedPolicyIdAce := func(ptr interface{}) { o.SystemScopedPolicyIDACE = *ptr.(**SystemScopedPolicyIDACE) }
	if err := w.ReadPointer(&o.SystemScopedPolicyIDACE, _s_SystemScopedPolicyIdAce, _ptr_SystemScopedPolicyIdAce); err != nil {
		return err
	}
	return nil
}

// ACEData_RawACE structure represents ACE_DATA RPC default union arm.
type ACEData_RawACE struct {
	RawACE *RawACE `idl:"name:RawAce" json:"raw_ace"`
}

func (*ACEData_RawACE) is_ACEData() {}

func (o *ACEData_RawACE) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RawACE != nil {
		_ptr_RawAce := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RawACE != nil {
				if err := o.RawACE.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RawACE{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RawACE, _ptr_RawAce); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ACEData_RawACE) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_RawAce := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RawACE == nil {
			o.RawACE = &RawACE{}
		}
		if err := o.RawACE.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_RawAce := func(ptr interface{}) { o.RawACE = *ptr.(**RawACE) }
	if err := w.ReadPointer(&o.RawACE, _s_RawAce, _ptr_RawAce); err != nil {
		return err
	}
	return nil
}

// TokenMandatoryPolicy structure represents TOKEN_MANDATORY_POLICY RPC structure.
//
// The TOKEN_MANDATORY_POLICY structure specifies the mandatory integrity policy for
// a token.
type TokenMandatoryPolicy struct {
	// Policy:  The Policy member contains a value denoting the mandatory integrity policy
	// of the token; these values are mutually exclusive.
	//
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                                   |                                                                                  |
	//	|                       VALUE                       |                                     MEANING                                      |
	//	|                                                   |                                                                                  |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	| TOKEN_MANDATORY_POLICY_OFF 0x00000000             | No mandatory integrity policy is enforced for the token.                         |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	| TOKEN_MANDATORY_POLICY_NO_WRITE_UP 0x00000001     | A process associated with the token cannot write to objects that have a greater  |
	//	|                                                   | mandatory integrity level.                                                       |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	| TOKEN_MANDATORY_POLICY_NEW_PROCESS_MIN 0x00000002 | A process created with the token has an integrity level that is the lesser of    |
	//	|                                                   | the parent-process integrity level and the executable-file integrity level.      |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	Policy uint32 `idl:"name:Policy" json:"policy"`
}

func (o *TokenMandatoryPolicy) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TokenMandatoryPolicy) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Policy); err != nil {
		return err
	}
	return nil
}
func (o *TokenMandatoryPolicy) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Policy); err != nil {
		return err
	}
	return nil
}

// MandatoryInformation structure represents MANDATORY_INFORMATION RPC structure.
//
// The MANDATORY_INFORMATION structure defines mandatory security information for a
// securable object.
type MandatoryInformation struct {
	// AllowedAccess:  The AllowedAccess member specifies the access mask that is used to
	// encode the user rights to an object.
	AllowedAccess uint32 `idl:"name:AllowedAccess" json:"allowed_access"`
	// WriteAllowed:  Specifies write properties for the object.
	WriteAllowed bool `idl:"name:WriteAllowed" json:"write_allowed"`
	// ReadAllowed:  Specifies read properties for the object.
	ReadAllowed bool `idl:"name:ReadAllowed" json:"read_allowed"`
	// ExecuteAllowed:  Specifies execution properties for the object.
	ExecuteAllowed bool `idl:"name:ExecuteAllowed" json:"execute_allowed"`
	// MandatoryPolicy:  Specifies the integrity policy for the object.
	MandatoryPolicy *TokenMandatoryPolicy `idl:"name:MandatoryPolicy" json:"mandatory_policy"`
}

func (o *MandatoryInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *MandatoryInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.AllowedAccess); err != nil {
		return err
	}
	if err := w.WriteData(o.WriteAllowed); err != nil {
		return err
	}
	if err := w.WriteData(o.ReadAllowed); err != nil {
		return err
	}
	if err := w.WriteData(o.ExecuteAllowed); err != nil {
		return err
	}
	if o.MandatoryPolicy != nil {
		if err := o.MandatoryPolicy.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TokenMandatoryPolicy{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MandatoryInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.AllowedAccess); err != nil {
		return err
	}
	if err := w.ReadData(&o.WriteAllowed); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReadAllowed); err != nil {
		return err
	}
	if err := w.ReadData(&o.ExecuteAllowed); err != nil {
		return err
	}
	if o.MandatoryPolicy == nil {
		o.MandatoryPolicy = &TokenMandatoryPolicy{}
	}
	if err := o.MandatoryPolicy.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ClaimSecurityAttributeOctetStringRelative structure represents CLAIM_SECURITY_ATTRIBUTE_OCTET_STRING_RELATIVE RPC structure.
//
// The CLAIM_SECURITY_ATTRIBUTE_OCTET_STRING_RELATIVE structure specifies an octet string.<78>
type ClaimSecurityAttributeOctetStringRelative struct {
	// Length:  The length, in bytes, of the value contained in the OctetString field.
	Length uint32 `idl:"name:Length" json:"length"`
	// OctetString:  An array of bytes containing the octet string value. The length of
	// the value is specified by the Length field.
	OctetString []byte `idl:"name:OctetString" json:"octet_string"`
}

func (o *ClaimSecurityAttributeOctetStringRelative) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *ClaimSecurityAttributeOctetStringRelative) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(len(o.OctetString))
	return []uint64{
		dimSize1,
	}
}
func (o *ClaimSecurityAttributeOctetStringRelative) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	for i1 := range o.OctetString {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.OctetString[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.OctetString); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClaimSecurityAttributeOctetStringRelative) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.OctetString", sizeInfo[0])
	}
	o.OctetString = make([]byte, sizeInfo[0])
	for i1 := range o.OctetString {
		i1 := i1
		if err := w.ReadData(&o.OctetString[i1]); err != nil {
			return err
		}
	}
	return nil
}

// ClaimSecurityAttributeRelativeV1 structure represents CLAIM_SECURITY_ATTRIBUTE_RELATIVE_V1 RPC structure.
//
// The CLAIM_SECURITY_ATTRIBUTE_RELATIVE_V1 structure defines a resource attribute that
// is defined in contiguous memory for persistence within a serialized Security Descriptor.
type ClaimSecurityAttributeRelativeV1 struct {
	// Name:  A DWORD value indicating an offset from the beginning of the CLAIM_SECURITY_ATTRIBUTE_RELATIVE_V1
	// structure to a string of Unicode characters containing the name of the claim security
	// attribute. The string MUST be at least 4 bytes in length.
	Name uint32 `idl:"name:Name" json:"name"`
	// ValueType:  A union tag value indicating the type of information referred to by the
	// Values member. The Values member MUST be an array of offsets from the beginning of
	// the CLAIM_SECURITY_ATTRIBUTE_RELATIVE_V1 structure to the specified ValueType. ValueType
	// MUST be one of the following values:
	//
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                                   |                                                                                  |
	//	|                       VALUE                       |                                     MEANING                                      |
	//	|                                                   |                                                                                  |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	| CLAIM_SECURITY_ATTRIBUTE_TYPE_INT64 0x0001        | Values member refers to an array of offsets to LONG64 value(s).                  |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	| CLAIM_SECURITY_ATTRIBUTE_TYPE_UINT64 0x0002       | Values member refers to an array of offsets to ULONG64 value(s).                 |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	| CLAIM_SECURITY_ATTRIBUTE_TYPE_STRING 0x0003       | Values member refers to an array of offsets to Unicode character string          |
	//	|                                                   | value(s).                                                                        |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	| CLAIM_SECURITY_ATTRIBUTE_TYPE_SID 0x0005          | The Values member refers to an array of offsets to                               |
	//	|                                                   | CLAIM_SECURITY_ATTRIBUTE_OCTET_STRING_RELATIVE value(s) where the OctetString    |
	//	|                                                   | value is a SID string.                                                           |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	| CLAIM_SECURITY_ATTRIBUTE_TYPE_BOOLEAN 0x0006      | The Values member refers to an array of offsets to ULONG64 values where each     |
	//	|                                                   | element indicates a Boolean value. The value 1 indicates TRUE, and the value 0   |
	//	|                                                   | indicates FALSE.                                                                 |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	| CLAIM_SECURITY_ATTRIBUTE_TYPE_OCTET_STRING 0x0010 | Values member contains an array of                                               |
	//	|                                                   | CLAIM_SECURITY_ATTRIBUTE_OCTET_STRING_RELATIVE value(s) as specified in section  |
	//	|                                                   | 2.4.10.2.                                                                        |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	ValueType uint16 `idl:"name:ValueType" json:"value_type"`
	// Reserved:  Reserved. This member MUST be set to zero when sent and MUST be ignored
	// when received.
	_ uint16 `idl:"name:Reserved"`
	// Flags:  The upper two bytes of this DWORD are available for application-specific
	// data. The two lowest-order bits in the lower of these two bytes are reserved. These
	// two bytes MAY<75> contain only one of the following values in those two bits:
	//
	//	+----------------------------------------------------+-----------------------------------------------------------------------+
	//	|                                                    |                                                                       |
	//	|                       VALUE                        |                                MEANING                                |
	//	|                                                    |                                                                       |
	//	+----------------------------------------------------+-----------------------------------------------------------------------+
	//	+----------------------------------------------------+-----------------------------------------------------------------------+
	//	| FCI_CLAIM_SECURITY_ATTRIBUTE_MANUAL 0x0001         | The CLAIM_SECURITY_ATTRIBUTE has been manually assigned               |
	//	+----------------------------------------------------+-----------------------------------------------------------------------+
	//	| FCI_CLAIM_SECURITY_ATTRIBUTE_POLICY_DERIVED 0x0002 | The CLAIM_SECURITY_ATTRIBUTE has been determined by a central policy. |
	//	+----------------------------------------------------+-----------------------------------------------------------------------+
	//
	// The lower two bytes of this DWORD MUST be zero or a bitwise combination of one or
	// more of the following values:<76>
	//
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                                      |                                                                                  |
	//	|                        VALUE                         |                                     MEANING                                      |
	//	|                                                      |                                                                                  |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| CLAIM_SECURITY_ATTRIBUTE_NON_INHERITABLE 0x0001      | This claim security attribute is not inherited across processes.<77>             |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| CLAIM_SECURITY_ATTRIBUTE_VALUE_CASE_SENSITIVE 0x0002 | The value of the claim security attribute is case sensitive. This flag is valid  |
	//	|                                                      | for values that contain string types.                                            |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| CLAIM_SECURITY_ATTRIBUTE_USE_FOR_DENY_ONLY 0x0004    | Reserved for future use.                                                         |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| CLAIM_SECURITY_ATTRIBUTE_DISABLED_BY_DEFAULT 0x0008  | The claim security attribute is disabled by default.                             |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| CLAIM_SECURITY_ATTRIBUTE_DISABLED 0x0010             | Reserved for future use.                                                         |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| CLAIM_SECURITY_ATTRIBUTE_MANDATORY 0x0020            | The claim security attribute is mandatory.                                       |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
	// ValueCount:  The number of values contained in the Values member.
	ValueCount uint32 `idl:"name:ValueCount" json:"value_count"`
	// Values:  An array of offsets from the beginning of the CLAIM_SECURITY_ATTRIBUTE_RELATIVE_V1
	// structure. Each offset indicates the location of a claim security attribute value
	// of type specified in the ValueType member.
	Values *ClaimSecurityAttributeRelativeV1_Values `idl:"name:Values" json:"values"`
}

func (o *ClaimSecurityAttributeRelativeV1) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClaimSecurityAttributeRelativeV1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Name); err != nil {
		return err
	}
	if err := w.WriteData(o.ValueType); err != nil {
		return err
	}
	// reserved Reserved
	if err := w.WriteData(uint16(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.ValueCount); err != nil {
		return err
	}
	// FIXME unknown type Values
	return nil
}
func (o *ClaimSecurityAttributeRelativeV1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Name); err != nil {
		return err
	}
	if err := w.ReadData(&o.ValueType); err != nil {
		return err
	}
	// reserved Reserved
	var _Reserved uint16
	if err := w.ReadData(&_Reserved); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.ValueCount); err != nil {
		return err
	}
	// FIXME: unknown type Values
	return nil
}

type ClaimSecurityAttributeRelativeV1_Values struct {
	Int64       []int64                                      `idl:"name:pInt64" json:"int64"`
	Uint64      []uint64                                     `idl:"name:pUint64" json:"uint64"`
	String      []string                                     `idl:"name:ppString" json:"string"`
	OctetString []*ClaimSecurityAttributeOctetStringRelative `idl:"name:pOctetString" json:"octet_string"`
}

// SID structure represents RPC_SID RPC structure.
//
// The RPC_SID structure is an IDL representation of the SID type (as specified in section
// 2.4.2) for use by RPC-based protocols.
type SID struct {
	// Revision:  An 8-bit unsigned integer that specifies the revision level of the SID.
	// This value MUST be set to 0x01.
	Revision uint8 `idl:"name:Revision" json:"revision"`
	// SubAuthorityCount:  An 8-bit unsigned integer that specifies the number of elements
	// in the SubAuthority array. The maximum number of elements allowed is 15.
	SubAuthorityCount uint8 `idl:"name:SubAuthorityCount" json:"sub_authority_count"`
	// IdentifierAuthority:  An RPC_SID_IDENTIFIER_AUTHORITY structure that indicates the
	// authority under which the SID was created. It describes the entity that created the
	// SID. The Identifier Authority value {0,0,0,0,0,5} denotes SIDs created by the NT
	// SID authority.
	IDAuthority *SIDIDAuthority `idl:"name:IdentifierAuthority" json:"id_authority"`
	// SubAuthority:  A variable length array of unsigned 32-bit integers that uniquely
	// identifies a principal relative to the IdentifierAuthority. Its length is determined
	// by SubAuthorityCount.
	SubAuthority []uint32 `idl:"name:SubAuthority;size_is:(SubAuthorityCount)" json:"sub_authority"`
}

func (o *SID) xxx_PreparePayload(ctx context.Context) error {
	if o.SubAuthority != nil && o.SubAuthorityCount == 0 {
		o.SubAuthorityCount = uint8(len(o.SubAuthority))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *SID) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.SubAuthorityCount)
	return []uint64{
		dimSize1,
	}
}
func (o *SID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.Revision); err != nil {
		return err
	}
	if err := w.WriteData(o.SubAuthorityCount); err != nil {
		return err
	}
	if o.IDAuthority != nil {
		if err := o.IDAuthority.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SIDIDAuthority{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	for i1 := range o.SubAuthority {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.SubAuthority[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.SubAuthority); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *SID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.Revision); err != nil {
		return err
	}
	if err := w.ReadData(&o.SubAuthorityCount); err != nil {
		return err
	}
	if o.IDAuthority == nil {
		o.IDAuthority = &SIDIDAuthority{}
	}
	if err := o.IDAuthority.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.SubAuthorityCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.SubAuthorityCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.SubAuthority", sizeInfo[0])
	}
	o.SubAuthority = make([]uint32, sizeInfo[0])
	for i1 := range o.SubAuthority {
		i1 := i1
		if err := w.ReadData(&o.SubAuthority[i1]); err != nil {
			return err
		}
	}
	return nil
}

// ACL structure represents ACL RPC structure.
//
// The access control list (ACL) packet is used to specify a list of individual access
// control entries (ACEs). An ACL packet and an array of ACEs comprise a complete access
// control list.
//
// The individual ACEs in an ACL are numbered from 0 to n, where n+1 is the number of
// ACEs in the ACL. When editing an ACL, an application refers to an ACE within the
// ACL by the ACE index.
//
// In the absence of implementation-specific functions to access the individual ACEs,
// access to each ACE MUST be computed by using the AclSize and AceCount fields to parse
// the wire packets following the ACL to identify each ACE_HEADER, which in turn contains
// the information needed to obtain the specific ACEs.
//
// An ACL is said to be in canonical form if:
//
// * All explicit ACEs are placed before inherited ACEs.
//
// * Within the explicit ACEs, deny ACEs come before grant ACEs.
//
// * Deny ACEs on the object come before deny ACEs on a child or property.
//
// * Grant ACEs on the object come before grant ACEs on a child or property.
//
// * Inherited ACEs are placed in the order in which they were inherited.
//
// There are two types of ACL:
//
// * A discretionary access control list (DACL) ( a66edeb1-52a0-4d64-a93b-2f5c833d7d92#gt_d727f612-7a45-48e4-9d87-71735d62b321
// ) is controlled by the owner of an object or anyone granted WRITE_DAC access to the
// object. It specifies the access particular users and groups can have to an object.
// For example, the owner of a file can use a DACL to control which users and groups
// can and cannot have access to the file.
//
// * A system access control list (SACL) ( a66edeb1-52a0-4d64-a93b-2f5c833d7d92#gt_c189801e-3752-4715-88f4-17804dad5782
// ) is similar to the DACL, except that the SACL is used to audit rather than control
// access to an object. When an audited action occurs, the operating system records
// the event in the security log. Each ACE in a SACL has a header that indicates whether
// auditing is triggered by success, failure, or both; a SID that specifies a particular
// user or security group to monitor; and an access mask that lists the operations to
// audit.
//
// The SACL also MAY contain <71> ( 11e1608c-6169-4fbc-9c33-373fc9b224f4#Appendix_A_71
// ) a label ACE that defines the integrity level of the object.
//
// The only valid ACE types for a SACL are the auditing types (SYSTEM_AUDIT_ACE_TYPE,
// SYSTEM_AUDIT_OBJECT_ACE_TYPE, SYSTEM_AUDIT_CALLBACK_ACE_TYPE, and SYSTEM_AUDIT_CALLBACK_OBJECT_ACE_TYPE),
// the label type (SYSTEM_MANDATORY_LABEL_ACE_TYPE), the system resource attribute type
// (SYSTEM_RESOURCE_ATTRIBUTE_ACE_TYPE), and the scoped policy type (SYSTEM_SCOPED_POLICY_ID_ACE_TYPE),
// as specified in section 2.4.4.1.
//
// The SACL MUST NOT contain ACEs that belong in the DACL, and the DACL MUST NOT contain
// ACE types that belong in the SACL. Doing so results in unspecified behavior.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| AclRevision                   | Sbz1                          | AclSize                                                       |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| AceCount                                                      | Sbz2                                                          |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//
// The RPC representation of the ACL data type specifies the elements needed to access
// a complete access control list, including both the ACL header structure and the array
// of ACEs. The individual members are as specified in section 2.4.5.
//
// The ACL structure MUST be aligned on a 32-bit boundary.
//
// In the absence of implementation-specific functions to access the individual ACEs,
// access to each ACE MUST be computed by using the AclSize and AceCount members to
// parse the memory following the ACL to identify each ACE_HEADER, which in turn contains
// the information needed to obtain the specific ACEs.
type ACL struct {
	// AclRevision (1 byte): An unsigned 8-bit value that specifies the revision of the
	// ACL. The only two legitimate forms of ACLs supported for on-the-wire management or
	// manipulation are type 2 and type 4. No other form is valid for manipulation on the
	// wire. Therefore this field MUST be set to one of the following values.
	//
	//	+----------------------+----------------------------------------------------------------------------------+
	//	|                      |                                                                                  |
	//	|        VALUE         |                                     MEANING                                      |
	//	|                      |                                                                                  |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| ACL_REVISION 0x02    | When set to 0x02, only AceTypes 0x00, 0x01, 0x02, 0x03, 0x11, 0x12, and 0x13 can |
	//	|                      | be present in the ACL. An AceType of 0x11 is used for SACLs but not for DACLs.   |
	//	|                      | For more information about ACE types, see section 2.4.4.1.                       |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| ACL_REVISION_DS 0x04 | When set to 0x04, AceTypes 0x05, 0x06, 0x07, 0x08, and 0x11 are allowed. ACLs    |
	//	|                      | of revision 0x04 are applicable only to directory service objects. An AceType of |
	//	|                      | 0x11 is used for SACLs but not for DACLs.                                        |
	//	+----------------------+----------------------------------------------------------------------------------+
	ACLRevision uint8 `idl:"name:AclRevision" json:"acl_revision"`
	// Sbz1 (1 byte): An unsigned 8-bit value. This field is reserved and MUST be set to
	// zero.
	SBZ1 uint8 `idl:"name:Sbz1" json:"sbz1"`
	// AclSize (2 bytes): An unsigned 16-bit integer that specifies the size, in bytes,
	// of the complete ACL, including all ACEs.
	ACLSize uint16 `idl:"name:AclSize" json:"acl_size"`
	// AceCount (2 bytes): An unsigned 16-bit integer that specifies the count of the number
	// of ACE records in the ACL.
	ACECount uint16 `idl:"name:AceCount" json:"ace_count"`
	// Sbz2 (2 bytes): An unsigned 16-bit integer. This field is reserved and MUST be set
	// to zero.
	SBZ2 uint16 `idl:"name:Sbz2" json:"sbz2"`
	Aces []*ACE `idl:"name:Aces;size_is:(AceCount)" json:"aces"`
}

func (o *ACL) xxx_PreparePayload(ctx context.Context) error {
	if o.Aces != nil && o.ACECount == 0 {
		o.ACECount = uint16(len(o.Aces))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *ACL) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.ACECount)
	return []uint64{
		dimSize1,
	}
}
func (o *ACL) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteData(o.ACLRevision); err != nil {
		return err
	}
	if err := w.WriteData(o.SBZ1); err != nil {
		return err
	}
	if err := w.WriteData(o.ACLSize); err != nil {
		return err
	}
	if err := w.WriteData(o.ACECount); err != nil {
		return err
	}
	if err := w.WriteData(o.SBZ2); err != nil {
		return err
	}
	for i1 := range o.Aces {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.Aces[i1] != nil {
			_ptr_Aces := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Aces[i1] != nil {
					if err := o.Aces[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ACE{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Aces[i1], _ptr_Aces); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Aces); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ACL) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadData(&o.ACLRevision); err != nil {
		return err
	}
	if err := w.ReadData(&o.SBZ1); err != nil {
		return err
	}
	if err := w.ReadData(&o.ACLSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.ACECount); err != nil {
		return err
	}
	if err := w.ReadData(&o.SBZ2); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.ACECount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.ACECount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Aces", sizeInfo[0])
	}
	o.Aces = make([]*ACE, sizeInfo[0])
	for i1 := range o.Aces {
		i1 := i1
		_ptr_Aces := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Aces[i1] == nil {
				o.Aces[i1] = &ACE{}
			}
			if err := o.Aces[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Aces := func(ptr interface{}) { o.Aces[i1] = *ptr.(**ACE) }
		if err := w.ReadPointer(&o.Aces[i1], _s_Aces, _ptr_Aces); err != nil {
			return err
		}
	}
	return nil
}

// SecurityDescriptor structure represents SECURITY_DESCRIPTOR RPC structure.
//
// The SECURITY_DESCRIPTOR structure defines the security attributes of an object. These
// attributes specify who owns the object; who can access the object and what they can
// do with it; what level of audit logging can be applied to the object; and what kind
// of restrictions apply to the use of the security descriptor.
//
// Security descriptors appear in one of two forms, absolute or self-relative.
//
// A security descriptor is said to be in absolute format if it stores all of its security
// information via pointer fields, as specified in the RPC representation in section
// 2.4.6.1.
//
// A security descriptor is said to be in self-relative format if it stores all of its
// security information in a contiguous block of memory and expresses all of its pointer
// fields as offsets from its beginning. The order of appearance of pointer target fields
// is not required to be in any particular order; the location of  the OwnerSid, GroupSid,
// Sacl, and/or Dacl is only based on OffsetOwner, OffsetGroup, OffsetSacl, and/or OffsetDacl
// pointers found in the fixed portion of the relative security descriptor.<72>
//
// The self-relative form of the security descriptor is required if one wants to transmit
// the SECURITY_DESCRIPTOR structure as an opaque data structure for transmission in
// communication protocols over a wire, or for storage on secondary media; the absolute
// form cannot be transmitted because it contains pointers to objects that are generally
// not accessible to the recipient.
//
// When a self-relative security descriptor is transmitted over a wire, it is sent in
// little-endian format and requires no padding.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Revision                      | Sbz1                          | Control                                                       |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| OffsetOwner                                                                                                                   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| OffsetGroup                                                                                                                   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| OffsetSacl                                                                                                                    |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| OffsetDacl                                                                                                                    |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| OwnerSid (variable)                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| GroupSid (variable)                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Sacl (variable)                                                                                                               |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Dacl (variable)                                                                                                               |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//
// The RPC representation of the SECURITY_DESCRIPTOR structure defines the in-memory
// representation of the SECURITY_DESCRIPTOR message. The individual member semantics
// for the Revision, Sbz1, Control, Owner, Group, Sacl and Dacl members are as specified
// in section 2.4.6, with the exceptions that Owner corresponds to OwnerSid, and Group
// corresponds to GroupSid, respectively.
type SecurityDescriptor struct {
	// Revision (1 byte): An unsigned 8-bit value that specifies the revision of the SECURITY_DESCRIPTOR
	// structure. This field MUST be set to one.
	//
	// Revision:  As specified in section 2.4.6.
	Revision uint8 `idl:"name:Revision" json:"revision"`
	// Sbz1 (1 byte): An unsigned 8-bit value with no meaning unless the Control RM bit
	// is set to 0x1. If the RM bit is set to 0x1, Sbz1 is interpreted as the resource manager
	// control bits that contain specific information<73> for the specific resource manager
	// that is accessing the structure. The permissible values and meanings of these bits
	// are determined by the implementation of the resource manager.
	//
	// Sbz1:  As specified in section 2.4.6.
	SBZ1 uint8 `idl:"name:Sbz1" json:"sbz1"`
	// Control (2 bytes): An unsigned 16-bit field that specifies control access bit flags.
	// The Self Relative (SR) bit MUST be set when the security descriptor is in self-relative
	// format.
	//
	//	+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+
	//	|     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |
	//	|  0  |  1  |  2  |  3  |  4  |  5  |  6  |  7  |  8  |  9  |  1  |  1  |  2  |  3  |  4  |  5  |
	//	|     |     |     |     |     |     |     |     |     |     |  0  |     |     |     |     |     |
	//	|     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |
	//	+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+
	//	+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+
	//	| S R | R M | P S | P D | S I | D I | S C | D C | S S | D T | S D | S P | D D | D P | G D | O D |
	//	+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+
	//
	// Where the bits are defined as:
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                                       |                                                                                  |
	//	|                 VALUE                 |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| SR Self-Relative                      | Set when the security descriptor is in self-relative format. Cleared when the    |
	//	|                                       | security descriptor is in absolute format.                                       |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| RM RM Control Valid                   | Set to 0x1 when the Sbz1 field is to be interpreted as resource manager control  |
	//	|                                       | bits.                                                                            |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| PS SACL Protected                     | Set when the SACL will be protected from inherit operations.                     |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| PD DACL Protected                     | Set when the DACL will be protected from inherit operations.                     |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| SI SACL Auto-Inherited                | Set when the SACL was created through inheritance.                               |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| DI DACL Auto-Inherited                | Set when the DACL was created through inheritance.                               |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| SC SACL Computed Inheritance Required | Set when the SACL is to be computed through inheritance. When both SC and SI are |
	//	|                                       | set, the resulting security descriptor sets SI; the SC setting is not preserved. |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| DC DACL Computed Inheritance Required | Set when the DACL is to be computed through inheritance. When both DC and DI are |
	//	|                                       | set, the resulting security descriptor sets DI; the DC setting is not preserved. |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| SS Server Security                    | Set when the caller wants the system to create a Server ACL based on the input   |
	//	|                                       | ACL, regardless of its source (explicit or defaulting).                          |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| DT DACL Trusted                       | Set when the ACL that is pointed to by the DACL field was provided by a trusted  |
	//	|                                       | source and does not require any editing of compound ACEs.                        |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| SD SACL Defaulted                     | Set when the SACL was established by default means.                              |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| SP SACL Present                       | Set when the SACL is present on the object.                                      |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| DD DACL Defaulted                     | Set when the DACL was established by default means.                              |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| DP DACL Present                       | Set when the DACL is present on the object.                                      |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| GD Group Defaulted                    | Set when the group was established by default means.                             |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| OD Owner Defaulted                    | Set when the owner was established by default means.                             |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//
	// Control:  As specified in section 2.4.6.
	Control uint16 `idl:"name:Control" json:"control"`
	// OffsetOwner (4 bytes): An unsigned 32-bit integer that specifies the offset to the
	// SID. This SID specifies the owner of the object to which the security descriptor
	// is associated. This must be a valid offset if the OD flag is not set. If this field
	// is set to zero, the OwnerSid field MUST not be present.
	OffsetOwner uint32 `idl:"name:OffsetOwner" json:"offset_owner"`
	// OffsetGroup (4 bytes): An unsigned 32-bit integer that specifies the offset to the
	// SID. This SID specifies the group of the object to which the security descriptor
	// is associated. This must be a valid offset if the GD flag is not set. If this field
	// is set to zero, the GroupSid field MUST not be present.
	OffsetGroup uint32 `idl:"name:OffsetGroup" json:"offset_group"`
	// OffsetSacl (4 bytes): An unsigned 32-bit integer that specifies the offset to the
	// ACL that contains system ACEs. Typically, the system ACL contains auditing ACEs (such
	// as SYSTEM_AUDIT_ACE, SYSTEM_AUDIT_CALLBACK_ACE, or SYSTEM_AUDIT_CALLBACK_OBJECT_ACE),
	// and at most one Label ACE (as specified in section 2.4.4.13). This must be a valid
	// offset if the SP flag is set; if the SP flag is not set, this field MUST be set to
	// zero. If this field is set to zero, the Sacl field MUST not be present.
	OffsetSACL uint32 `idl:"name:OffsetSacl" json:"offset_sacl"`
	// OffsetDacl (4 bytes): An unsigned 32-bit integer that specifies the offset to the
	// ACL that contains ACEs that control access. Typically, the DACL contains ACEs that
	// grant or deny access to principals or groups. This must be a valid offset if the
	// DP flag is set; if the DP flag is not set, this field MUST be set to zero. If this
	// field is set to zero, the Dacl field MUST not be present.
	OffsetDACL uint32 `idl:"name:OffsetDacl" json:"offset_dacl"`
	// Owner:  Pointer to the Owner SID (OwnerSid), as specified in section 2.4.6.
	Owner *SID `idl:"name:Owner" json:"owner"`
	// Group:  Pointer to the Group SID (GroupSid), as specified in section 2.4.6.
	Group *SID `idl:"name:Group" json:"group"`
	// Sacl (variable): The SACL of the object. The length of the SID MUST be a multiple
	// of 4. This field MUST be present if the SP flag is set.
	//
	// Sacl:  Pointer to the Sacl, as specified in section 2.4.6.
	SACL *ACL `idl:"name:Sacl" json:"sacl"`
	// Dacl (variable): The DACL of the object. The length of the SID MUST be a multiple
	// of 4. This field MUST be present if the DP flag is set.
	//
	// Dacl:  Pointer to the Dacl, as specified in section 2.4.6.
	DACL *ACL `idl:"name:Dacl" json:"dacl"`
}

func (o *SecurityDescriptor) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SecurityDescriptor) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteData(o.Revision); err != nil {
		return err
	}
	if err := w.WriteData(o.SBZ1); err != nil {
		return err
	}
	if err := w.WriteData(o.Control); err != nil {
		return err
	}
	if o.Owner != nil {
		_ptr_Owner := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Owner != nil {
				if err := o.Owner.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Owner, _ptr_Owner); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Group != nil {
		_ptr_Group := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Group != nil {
				if err := o.Group.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Group, _ptr_Group); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.SACL != nil {
		_ptr_Sacl := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SACL != nil {
				if err := o.SACL.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ACL{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SACL, _ptr_Sacl); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DACL != nil {
		_ptr_Dacl := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.DACL != nil {
				if err := o.DACL.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ACL{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.DACL, _ptr_Dacl); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SecurityDescriptor) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadData(&o.Revision); err != nil {
		return err
	}
	if err := w.ReadData(&o.SBZ1); err != nil {
		return err
	}
	if err := w.ReadData(&o.Control); err != nil {
		return err
	}
	_ptr_Owner := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Owner == nil {
			o.Owner = &SID{}
		}
		if err := o.Owner.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Owner := func(ptr interface{}) { o.Owner = *ptr.(**SID) }
	if err := w.ReadPointer(&o.Owner, _s_Owner, _ptr_Owner); err != nil {
		return err
	}
	_ptr_Group := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Group == nil {
			o.Group = &SID{}
		}
		if err := o.Group.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Group := func(ptr interface{}) { o.Group = *ptr.(**SID) }
	if err := w.ReadPointer(&o.Group, _s_Group, _ptr_Group); err != nil {
		return err
	}
	_ptr_Sacl := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SACL == nil {
			o.SACL = &ACL{}
		}
		if err := o.SACL.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Sacl := func(ptr interface{}) { o.SACL = *ptr.(**ACL) }
	if err := w.ReadPointer(&o.SACL, _s_Sacl, _ptr_Sacl); err != nil {
		return err
	}
	_ptr_Dacl := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.DACL == nil {
			o.DACL = &ACL{}
		}
		if err := o.DACL.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Dacl := func(ptr interface{}) { o.DACL = *ptr.(**ACL) }
	if err := w.ReadPointer(&o.DACL, _s_Dacl, _ptr_Dacl); err != nil {
		return err
	}
	return nil
}
