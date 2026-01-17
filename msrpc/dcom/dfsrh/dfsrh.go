// The dfsrh package implements the DFSRH client protocol.
//
// # Introduction
//
// The Distributed File System: Replication Helper (DFS-R Helper) Protocol is a set
// of DCOM interfaces for configuring and monitoring the Distributed File System.
//
// # Overview
//
// The Distributed File System: Replication Helper (DFS-R Helper) Protocol provides
// a set of DCOM interfaces for configuring and monitoring Distributed File System–Replication
// (DFS-R) on a server, as specified in [MS-FRS2]. The server end of the protocol is
// a DCOM service that implements the DCOM interface. The client end of the protocol
// is an application that invokes methods on the interface to make DFS-R configuration
// changes and monitor the status of the DFS-R service on the server.
//
// The first part of the Distributed File System: Replication Helper (DFS-R Helper)
// Protocol consists of interfaces for creating, modifying, and deleting configuration
// objects in Active Directory by using the machine account of the server.
//
// For all replication members, the configuration related to a member is stored in the
// computer object for the local machine in Active Directory. It is common for system
// components that are unrelated to DFS-R to set permissions on the computer object
// that will prevent modification of the object by some users and still permit modification
// by using the credentials for the computer. Therefore, a server implementation uses
// the credentials of the local machine account when it sends commands to update Active
// Directory objects.
//
// If a user has sufficient privileges to connect to the server that is running the
// DFS-R Helper Protocol and to invoke methods implemented by the DFS-R Helper Protocol
// interfaces, the server works as a proxy for making configuration changes on behalf
// of the client application that is running under the user's account.
//
// The client sends the server information about the Active Directory operation that
// the client is trying to accomplish. The server then attempts to execute the command
// by using the machine account and returns information about the status of the operation.
//
// The second part of the Distributed File System: Replication Helper (DFS-R Helper)
// Protocol is an interface for monitoring DFS-R on the computer and collecting various
// statistics about the DFS-R operation.
//
// The information that is collected by using the Distributed File System: Replication
// Helper (DFS-R Helper) Protocol includes, among other types of information, the following
// statistics:
//
// * Information about replication errors that are encountered by DFS-R on the server.
//
// * The count and size of replicated files on the server.
//
// * Disk use on the server.
//
// * Information about replicated folders ( acefe7bd-5912-4f18-9554-b3da50a2ee44#gt_064adaf1-86c7-43e5-a157-b0949980181e
// ) on the server.
//
// * Replication backlog—the number of files that are not yet fully replicated.
package dfsrh

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
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
	_ = oaut.GoPackage
	_ = dcom.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/dfsrh"
)

// HelperErrorsEnum type represents DfsrHelperErrorsEnum RPC enumeration.
//
// The DfsrHelperErrorsEnum enumeration defines error codes that are specific to the
// IADProxy and IADProxy2 interfaces.
//
// The UUID for this enumeration is {9009D654-250B-4e0d-9AB0-ACB63134F69F}.
type HelperErrorsEnum uint32

var (
	// dfsrHelperErrorNotLocalAdmin:  Reserved for future use.
	HelperErrorsEnumNotLocalAdmin HelperErrorsEnum = 2147753985
	// dfsrHelperErrorCreateVerifyServerControl: Cannot create LDAP_SERVER_VERIFY_NAME_OID
	// control for the LDAP command.
	HelperErrorsEnumCreateVerifyServerControl HelperErrorsEnum = 2147753986
	// dfsrHelperLdapErrorBase:  This is the base value for LDAP errors.
	HelperErrorsEnumLDAPErrorBase HelperErrorsEnum = 2147758080
)

func (o HelperErrorsEnum) String() string {
	switch o {
	case HelperErrorsEnumNotLocalAdmin:
		return "HelperErrorsEnumNotLocalAdmin"
	case HelperErrorsEnumCreateVerifyServerControl:
		return "HelperErrorsEnumCreateVerifyServerControl"
	case HelperErrorsEnumLDAPErrorBase:
		return "HelperErrorsEnumLDAPErrorBase"
	}
	return "Invalid"
}

// ReportingFlags type represents DfsrReportingFlags RPC enumeration.
//
// The DfsrReportingFlags enumeration represents the options for generating health reports,
// which are used in IServerHealthReport and IServerHealthReport2 interfaces. The UUID
// for this enumeration is {CEB5D7B4-3964-4f71-AC17-4BF57A379D87}.
//
// Any bitmask that consists of one, or a combination, of the following enumerated values
// is supported:
type ReportingFlags uint16

var (
	// REPORTING_FLAGS_NONE:  Default report options.
	ReportingFlagsNone ReportingFlags = 0
	// REPORTING_FLAGS_BACKLOG:  Return the count of backlog transactions.
	ReportingFlagsBacklog ReportingFlags = 1
	// REPORTING_FLAGS_FILES:  Return the count and cumulative size of files in the replicated
	// folders.
	ReportingFlagsFiles ReportingFlags = 2
)

func (o ReportingFlags) String() string {
	switch o {
	case ReportingFlagsNone:
		return "ReportingFlagsNone"
	case ReportingFlagsBacklog:
		return "ReportingFlagsBacklog"
	case ReportingFlagsFiles:
		return "ReportingFlagsFiles"
	}
	return "Invalid"
}

// ADAttributeData structure represents _AdAttributeData RPC structure.
type ADAttributeData struct {
	Operation      int32        `idl:"name:operation" json:"operation"`
	AttributeName  *oaut.String `idl:"name:attributeName" json:"attribute_name"`
	AttributeValue *oaut.String `idl:"name:attributeValue" json:"attribute_value"`
	IsString       int16        `idl:"name:isString" json:"is_string"`
	Length         int32        `idl:"name:length" json:"length"`
}

func (o *ADAttributeData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ADAttributeData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Operation); err != nil {
		return err
	}
	if o.AttributeName != nil {
		_ptr_attributeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.AttributeName != nil {
				if err := o.AttributeName.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.AttributeName, _ptr_attributeName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.AttributeValue != nil {
		_ptr_attributeValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.AttributeValue != nil {
				if err := o.AttributeValue.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.AttributeValue, _ptr_attributeValue); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.IsString); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *ADAttributeData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Operation); err != nil {
		return err
	}
	_ptr_attributeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.AttributeName == nil {
			o.AttributeName = &oaut.String{}
		}
		if err := o.AttributeName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_attributeName := func(ptr interface{}) { o.AttributeName = *ptr.(**oaut.String) }
	if err := w.ReadPointer(&o.AttributeName, _s_attributeName, _ptr_attributeName); err != nil {
		return err
	}
	_ptr_attributeValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.AttributeValue == nil {
			o.AttributeValue = &oaut.String{}
		}
		if err := o.AttributeValue.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_attributeValue := func(ptr interface{}) { o.AttributeValue = *ptr.(**oaut.String) }
	if err := w.ReadPointer(&o.AttributeValue, _s_attributeValue, _ptr_attributeValue); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsString); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// VersionVectorData structure represents _VersionVectorData RPC structure.
type VersionVectorData struct {
	UncompressedSize int32         `idl:"name:uncompressedSize" json:"uncompressed_size"`
	BacklogCount     int32         `idl:"name:backlogCount" json:"backlog_count"`
	ContentSetGUID   *oaut.String  `idl:"name:contentSetGuid" json:"content_set_guid"`
	VersionVector    *oaut.Variant `idl:"name:versionVector" json:"version_vector"`
}

func (o *VersionVectorData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *VersionVectorData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.UncompressedSize); err != nil {
		return err
	}
	if err := w.WriteData(o.BacklogCount); err != nil {
		return err
	}
	if o.ContentSetGUID != nil {
		_ptr_contentSetGuid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ContentSetGUID != nil {
				if err := o.ContentSetGUID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ContentSetGUID, _ptr_contentSetGuid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.VersionVector != nil {
		_ptr_versionVector := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.VersionVector != nil {
				if err := o.VersionVector.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.VersionVector, _ptr_versionVector); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *VersionVectorData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.UncompressedSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.BacklogCount); err != nil {
		return err
	}
	_ptr_contentSetGuid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ContentSetGUID == nil {
			o.ContentSetGUID = &oaut.String{}
		}
		if err := o.ContentSetGUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_contentSetGuid := func(ptr interface{}) { o.ContentSetGUID = *ptr.(**oaut.String) }
	if err := w.ReadPointer(&o.ContentSetGUID, _s_contentSetGuid, _ptr_contentSetGuid); err != nil {
		return err
	}
	_ptr_versionVector := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.VersionVector == nil {
			o.VersionVector = &oaut.Variant{}
		}
		if err := o.VersionVector.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_versionVector := func(ptr interface{}) { o.VersionVector = *ptr.(**oaut.Variant) }
	if err := w.ReadPointer(&o.VersionVector, _s_versionVector, _ptr_versionVector); err != nil {
		return err
	}
	return nil
}

// ServerHealthReport structure represents IServerHealthReport RPC structure.
type ServerHealthReport dcom.InterfacePointer

func (o *ServerHealthReport) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *ServerHealthReport) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *ServerHealthReport) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ServerHealthReport) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.DataCount); err != nil {
		return err
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
}
func (o *ServerHealthReport) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
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
}

// IADProxy2 structure represents IADProxy2 RPC structure.
type IADProxy2 dcom.InterfacePointer

func (o *IADProxy2) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *IADProxy2) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *IADProxy2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *IADProxy2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.DataCount); err != nil {
		return err
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
}
func (o *IADProxy2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
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
}

// ServerHealthReport2 structure represents IServerHealthReport2 RPC structure.
type ServerHealthReport2 dcom.InterfacePointer

func (o *ServerHealthReport2) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *ServerHealthReport2) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *ServerHealthReport2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ServerHealthReport2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.DataCount); err != nil {
		return err
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
}
func (o *ServerHealthReport2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
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
}

// IADProxy structure represents IADProxy RPC structure.
type IADProxy dcom.InterfacePointer

func (o *IADProxy) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *IADProxy) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *IADProxy) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *IADProxy) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.DataCount); err != nil {
		return err
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
}
func (o *IADProxy) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
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
}
