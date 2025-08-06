// The comt package implements the COMT client protocol.
//
// # Introduction
//
// This document specifies the Component Object Model Plus (COM+) Tracker Service Protocol
// (COMT), which allows clients to monitor running instances of components.
//
// # Overview
//
// The COM+ Tracker Service Protocol enables remote clients to monitor instances of
// components running on a server. The server end of the protocol tracks the status
// of component instances and instance containers on the server and implements an interface
// that clients can use to poll for this status. It also optionally includes an event-driven
// notification system in which the client can supply (via another protocol) a callback
// interface for receiving tracker events. The server then calls the client's callback
// interface whenever new tracking data is available, for example, as a result of local
// events on the server.
package comt

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
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
	_ = dcom.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/comt"
)

// ContainerStatistics structure represents ContainerStatistics RPC structure.
//
// The ContainerStatistics type represents activity statistics for an instance container.
type ContainerStatistics struct {
	// cCalls:  The number of method calls that the component instances perform in an instance
	// container.
	CallsCount uint32 `idl:"name:cCalls" json:"calls_count"`
	// cComponentInstances:  The number of component instances in an instance container.
	ComponentInstancesCount uint32 `idl:"name:cComponentInstances" json:"component_instances_count"`
	// cComponents:  The number of distinct components currently instantiated in an instance
	// container.
	ComponentsCount uint32 `idl:"name:cComponents" json:"components_count"`
	// cCallsPerSecond:  This SHOULD be set to a running average, over an implementation-specific
	// time period,<1> of the number of method calls per second received by an instance
	// container. Alternatively, an implementation MAY instead simply set this field to
	// zero.
	CallsPerSecondCount uint32 `idl:"name:cCallsPerSecond" json:"calls_per_second_count"`
}

func (o *ContainerStatistics) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ContainerStatistics) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.CallsCount); err != nil {
		return err
	}
	if err := w.WriteData(o.ComponentInstancesCount); err != nil {
		return err
	}
	if err := w.WriteData(o.ComponentsCount); err != nil {
		return err
	}
	if err := w.WriteData(o.CallsPerSecondCount); err != nil {
		return err
	}
	return nil
}
func (o *ContainerStatistics) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.CallsCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.ComponentInstancesCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.ComponentsCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.CallsPerSecondCount); err != nil {
		return err
	}
	return nil
}

// ContainerData structure represents ContainerData RPC structure.
//
// The ContainerData type represents run-time information for a conglomeration that
// has one or more instance containers on the server. The meanings of the fields in
// this structure depend on the number of instance containers that exist on the server
// for the conglomeration represented, as specified in the following section.
type ContainerData struct {
	// dwLegacyId:  The container legacy identifier of one of the instance containers, arbitrarily
	// selected by the server, that exist for the conglomeration represented.
	LegacyID uint32 `idl:"name:dwLegacyId" json:"legacy_id"`
	// wszApplicationIdentifier:  A null-terminated Unicode string that MUST contain the
	// CurlyBraceGuidString (section 2.2.1) representation of a conglomeration identifier.
	// Note that a null-terminated CurlyBraceGuidString is 39 Unicode characters, including
	// the null character, and this field is 40 characters long. The final element in this
	// array is unused. It SHOULD be set to 0 and MUST be ignored upon receipt.
	ApplicationID []uint16 `idl:"name:wszApplicationIdentifier" json:"application_id"`
	// dwProcessId:  The process identifier of the process that contains one of the instance
	// containers, arbitrarily selected by the server, that exist for the conglomeration
	// represented.
	ProcessID uint32 `idl:"name:dwProcessId" json:"process_id"`
	// statistics:  A ContainerStatistics (section 2.2.2) structure with fields that contain
	// statistics averaged across all instance containers that exist for the conglomeration
	// represented.
	Statistics *ContainerStatistics `idl:"name:statistics" json:"statistics"`
}

func (o *ContainerData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ContainerData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.LegacyID); err != nil {
		return err
	}
	for i1 := range o.ApplicationID {
		i1 := i1
		if uint64(i1) >= 40 {
			break
		}
		if err := w.WriteData(o.ApplicationID[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.ApplicationID); uint64(i1) < 40; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ProcessID); err != nil {
		return err
	}
	if o.Statistics != nil {
		if err := o.Statistics.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ContainerStatistics{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ContainerData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.LegacyID); err != nil {
		return err
	}
	o.ApplicationID = make([]uint16, 40)
	for i1 := range o.ApplicationID {
		i1 := i1
		if err := w.ReadData(&o.ApplicationID[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.ProcessID); err != nil {
		return err
	}
	if o.Statistics == nil {
		o.Statistics = &ContainerStatistics{}
	}
	if err := o.Statistics.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ComponentData structure represents ComponentData RPC structure.
//
// This type represents activity statistics for a component that has one or more component
// instances in an instance container.
type ComponentData struct {
	// clsid:  The CLSID of the component.
	ClassID *dtyp.GUID `idl:"name:clsid" json:"class_id"`
	// cTotalReferences:  An implementation-specific<2> count of the number of references
	// to all component instances of the component. This MUST be set to 0xffffffff if the
	// server does not track this information.<3>
	TotalReferencesCount uint32 `idl:"name:cTotalReferences" json:"total_references_count"`
	// cBoundReferences:  The number of references to all active (not pooled) component
	// instances of the component. This MUST be set to 0xffffffff if the server does not
	// track this information.<4>
	BoundReferencesCount uint32 `idl:"name:cBoundReferences" json:"bound_references_count"`
	// cPooledInstances:  The number of pooled component instances of the component, if
	// the server enables instance pooling. This MUST be set to 0xffffffff if the server
	// does not track this information.<5>
	PooledInstancesCount uint32 `idl:"name:cPooledInstances" json:"pooled_instances_count"`
	// cInstancesInCall:  The number of component instances of the component that are currently
	// performing a method call. This MUST be set to 0xffffffff if the server does not track
	// this information.<6>
	InstancesInCallCount uint32 `idl:"name:cInstancesInCall" json:"instances_in_call_count"`
	// dwResponseTime:  A value that indicates the average time, in milliseconds, it takes
	// to complete method calls to component instances of the component. Calculation of
	// this value is implementation-specific.<7> This MUST be set to 0xffffffff if the server
	// does not track this information.<8>
	ResponseTime uint32 `idl:"name:dwResponseTime" json:"response_time"`
	// cCallsCompleted:  The number of method calls to component instances of the component
	// that were successfully completed in an implementation-specific<9> time period. Whether
	// a server considers a method call successfully completed is implementation-specific.<10>
	// This MUST be set to 0xffffffff if the server does not track this information.<11>
	CallsCompletedCount uint32 `idl:"name:cCallsCompleted" json:"calls_completed_count"`
	// cCallsFailed:  The number of method calls to component instances of the component
	// that failed in an implementation-specific<12> time period. Whether a server considers
	// a method call to have failed is implementation-specific.<13> This MUST be set to
	// 0xffffffff if the server does not track this information.<14>
	CallsFailedCount uint32 `idl:"name:cCallsFailed" json:"calls_failed_count"`
}

func (o *ComponentData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ComponentData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.ClassID != nil {
		if err := o.ClassID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.TotalReferencesCount); err != nil {
		return err
	}
	if err := w.WriteData(o.BoundReferencesCount); err != nil {
		return err
	}
	if err := w.WriteData(o.PooledInstancesCount); err != nil {
		return err
	}
	if err := w.WriteData(o.InstancesInCallCount); err != nil {
		return err
	}
	if err := w.WriteData(o.ResponseTime); err != nil {
		return err
	}
	if err := w.WriteData(o.CallsCompletedCount); err != nil {
		return err
	}
	if err := w.WriteData(o.CallsFailedCount); err != nil {
		return err
	}
	return nil
}
func (o *ComponentData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.ClassID == nil {
		o.ClassID = &dtyp.GUID{}
	}
	if err := o.ClassID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalReferencesCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.BoundReferencesCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.PooledInstancesCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.InstancesInCallCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.ResponseTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.CallsCompletedCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.CallsFailedCount); err != nil {
		return err
	}
	return nil
}

// GetTrackingData structure represents IGetTrackingData RPC structure.
//
// The IGetTrackingData interface provides methods for a client to poll for tracking
// information. This interface inherits from IUnknown, as specified in [MS-DCOM] section
// 3.1.1.5.8. The version for this interface is 0.0.
//
// To receive incoming remote calls for this interface, the server MUST implement a
// DCOM object class with the CLSID CLSID_TrackerService (as specified in section 1.9)
// by using the UUID {B60040E0-BCF3-11D1-861D-0080C729264D} for this interface.
//
// The IGetTrackingData interface includes the following methods beyond those of IUnknown.
//
// Methods in RPC Opnum Order
//
//	+-------------------------------------+----------------------------------------------------------------------------------+
//	|                                     |                                                                                  |
//	|               METHOD                |                                   DESCRIPTION                                    |
//	|                                     |                                                                                  |
//	+-------------------------------------+----------------------------------------------------------------------------------+
//	+-------------------------------------+----------------------------------------------------------------------------------+
//	| Opnum3NotUsedOnWire                 | Reserved for local use. Opnum: 3                                                 |
//	+-------------------------------------+----------------------------------------------------------------------------------+
//	| GetContainerData                    | Returns tracking information for instance containers. Opnum: 4                   |
//	+-------------------------------------+----------------------------------------------------------------------------------+
//	| GetComponentDataByContainer         | Returns tracking information for components by instance container. Opnum: 5      |
//	+-------------------------------------+----------------------------------------------------------------------------------+
//	| GetComponentDataByContainerAndCLSID | Returns tracking information for a component by instance container and CLSID.    |
//	|                                     | Opnum: 6                                                                         |
//	+-------------------------------------+----------------------------------------------------------------------------------+
//	| Opnum7NotUsedOnWire                 | Reserved for local use. Opnum: 7                                                 |
//	+-------------------------------------+----------------------------------------------------------------------------------+
//
// In the preceding table, the phrase "Reserved for local use" means that the client
// MUST NOT send the opnum and that the server behavior is undefined<19> because it
// does not affect interoperability.
//
// All methods MUST NOT throw exceptions.
type GetTrackingData dcom.InterfacePointer

func (o *GetTrackingData) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *GetTrackingData) xxx_PreparePayload(ctx context.Context) error {
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

func (o *GetTrackingData) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *GetTrackingData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *GetTrackingData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// COMTrackingInfoEvents structure represents IComTrackingInfoEvents RPC structure.
//
// The IComTrackingInfoEvents interface provides a method for a server to send the client
// tracker events. This interface inherits from IUnknown, as specified in [MS-DCOM]
// section 3.1.1.5.8. The version for this interface is 0.0.
//
// This interface includes the following method beyond those of IUnknown.
//
// Methods in RPC Opnum Order
//
//	+-------------------+-------------------------------------+
//	|                   |                                     |
//	|      METHOD       |             DESCRIPTION             |
//	|                   |                                     |
//	+-------------------+-------------------------------------+
//	+-------------------+-------------------------------------+
//	| OnNewTrackingInfo | Handles new tracking info. Opnum: 3 |
//	+-------------------+-------------------------------------+
//
// This method MUST NOT throw exceptions.
type COMTrackingInfoEvents dcom.InterfacePointer

func (o *COMTrackingInfoEvents) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *COMTrackingInfoEvents) xxx_PreparePayload(ctx context.Context) error {
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

func (o *COMTrackingInfoEvents) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *COMTrackingInfoEvents) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *COMTrackingInfoEvents) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ProcessDump structure represents IProcessDump RPC structure.
//
// The IProcessDump interface provides methods for a client to request a process dump
// of a process containing an instance container on the COMT server. This interface
// inherits from IDispatch, as specified in [MS-OAUT] section 3.1.4. The version for
// this interface is 0.0.
//
// To receive incoming remote calls for this interface, the server MUST implement a
// DCOM object class with the CLSID CLSID_ProcessDump (as specified in section 1.9)
// by using the UUID {23C9DD26-2355-4FE2-84DE-F779A238ADBD} for this interface.
//
// This interface includes the following methods beyond those of IDispatch.
//
// Methods in RPC Opnum Order
//
//	+-------------+--------------------------------------------------------------------------------+
//	|             |                                                                                |
//	|   METHOD    |                                  DESCRIPTION                                   |
//	|             |                                                                                |
//	+-------------+--------------------------------------------------------------------------------+
//	+-------------+--------------------------------------------------------------------------------+
//	| IsSupported | Returns a result indicating whether or not process dump is supported. Opnum: 7 |
//	+-------------+--------------------------------------------------------------------------------+
//	| DumpProcess | Requests a process dump. Opnum: 8                                              |
//	+-------------+--------------------------------------------------------------------------------+
//
// All methods MUST NOT throw exceptions.
type ProcessDump dcom.InterfacePointer

func (o *ProcessDump) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *ProcessDump) xxx_PreparePayload(ctx context.Context) error {
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

func (o *ProcessDump) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ProcessDump) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ProcessDump) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
