// The coma package implements the COMA client protocol.
//
// # Introduction
//
// This document specifies the Component Object Model Plus (COM+) Remote Administration
// Protocol (COMA), which allows clients to manage the configuration of software components
// and to control running instances of these components.
//
// # Overview
//
// The COM+ Remote Administration Protocol (COMA) enables remote clients to register,
// import, remove, configure, control, and monitor components and conglomerations for
// an Object Request Broker (ORB). The server end of the protocol is a conceptual service
// that maintains a catalog of configurations for an ORB. A COMA server exposes interfaces
// that enable a client to manage the catalog and control component instances and instance
// containers.
package coma

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = dtyp.GoPackage
	_ = dcom.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/coma"
)

// PropertyMeta structure represents PropertyMeta RPC structure.
//
// The PropertyMeta structure represents the type, size, and meta-properties (specified
// in this section) of a property in a table.
type PropertyMeta struct {
	// dataType:   The eDataType (section 2.2.1.2) value that represents the data type of
	// the property.
	DataType uint32 `idl:"name:dataType" json:"data_type"`
	// cbSize:   A size, in bytes, associated with the property. The meaning of this value
	// depends on the value of the dataType field and whether the fPROPERTY_FIXEDLENGTH
	// flag is set in the flags field.
	//
	//	+-------------------+----------------------------+----------------------------------------------------------------------------------+
	//	|     VALUE OF      |   FPROPERTY FIXEDLENGTH    |                                                                                  |
	//	|     DATATYPE      |            SET?            |                                     MEANING                                      |
	//	|                   |                            |                                                                                  |
	//	+-------------------+----------------------------+----------------------------------------------------------------------------------+
	//	+-------------------+----------------------------+----------------------------------------------------------------------------------+
	//	| eDT_ULONG         | -                          | The fixed size of the property. MUST be set to 0x00000004.                       |
	//	+-------------------+----------------------------+----------------------------------------------------------------------------------+
	//	| eDT_GUID          | -                          | The fixed size of the property. MUST be set to 0x00000010 (decimal 16).          |
	//	+-------------------+----------------------------+----------------------------------------------------------------------------------+
	//	| eDT_BYTES         | No                         | The maximum size of the property. A value of 0xFFFFFFFF indicates the property's |
	//	|                   |                            | size is unconstrained.                                                           |
	//	+-------------------+----------------------------+----------------------------------------------------------------------------------+
	//	| eDT_BYTES         | Yes                        | The fixed size of the property.                                                  |
	//	+-------------------+----------------------------+----------------------------------------------------------------------------------+
	//	| eDT_LPWSTR        | No                         | The maximum size of the property. A value of 0xFFFFFFFF indicates the property's |
	//	|                   |                            | size is unconstrained.                                                           |
	//	+-------------------+----------------------------+----------------------------------------------------------------------------------+
	//	| eDT_LPWSTR        | Yes                        | The fixed size of the property. MUST be set to a multiple of 2.                  |
	//	+-------------------+----------------------------+----------------------------------------------------------------------------------+
	Length uint32 `idl:"name:cbSize" json:"length"`
	// flags:   A bit field specifying the meta-properties of the property. MUST be a combination
	// of zero or more of the following flags.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                                      |                                                                                  |
	//	|                VALUE                 |                                     MEANING                                      |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| fPROPERTY_PRIMARYKEY 0x00000001      | This property is part of the primary key for its table. MUST be set if           |
	//	|                                      | fPROPERTY_NOTNULLABLE is set.                                                    |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| fPROPERTY_NOTNULLABLE 0x00000002     | This property cannot be null.                                                    |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| fPROPERTY_FIXEDLENGTH 0x00000004     | This eDT_BYTES or eDT_LPWSTR property has a fixed size. MUST NOT be set for      |
	//	|                                      | properties of type eDT_ULONG or eDT_GUID.                                        |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| fPROPERTY_NOTPERSISTABLE 0x00000008  | This property contains sensitive data such as passwords that MUST NOT be written |
	//	|                                      | in plaintext to persistent storage.                                              |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| fPROPERTY_CASEINSENSITIVE 0x00000020 | This eDT_LPWSTR property MUST be treated as case-insensitive for purposes of     |
	//	|                                      | comparison. MUST NOT be set for properties of type eDT_ULONG, eDT_GUID, or       |
	//	|                                      | eDT_BYTES.                                                                       |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:flags" json:"flags"`
}

func (o *PropertyMeta) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PropertyMeta) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataType); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	return nil
}
func (o *PropertyMeta) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataType); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	return nil
}

// ComponentType type represents eComponentType RPC enumeration.
//
// The eComponentType enumeration is used to select a component bitness when multiple
// bitnesses might exist for the same component.
type ComponentType uint16

var (
	// eCT_UNKNOWN:   The component bitness is unknown to the client. The server MUST
	// select the native bitness of the component if it exists; otherwise, the server MUST
	// select the non-native bitness of the component.
	ComponentTypeUnknown ComponentType = 0
	// eCT_32BIT:   The server MUST select the 32-bit bitness of the component.
	ComponentTypeCt32Bit ComponentType = 1
	// eCT_64BIT:   The server MUST select the 64-bit bitness of the component.
	ComponentTypeCt64Bit ComponentType = 2
	// eCT_NATIVE:   The server MUST select the native bitness (see section 3.1.4.4) of
	// the component.
	ComponentTypeNative ComponentType = 4096
)

func (o ComponentType) String() string {
	switch o {
	case ComponentTypeUnknown:
		return "ComponentTypeUnknown"
	case ComponentTypeCt32Bit:
		return "ComponentTypeCt32Bit"
	case ComponentTypeCt64Bit:
		return "ComponentTypeCt64Bit"
	case ComponentTypeNative:
		return "ComponentTypeNative"
	}
	return "Invalid"
}

// SRPLevelInfo structure represents SRPLevelInfo RPC structure.
//
// The SRPLevelInfo structure defines a software restriction policy trust level, as
// specified in section 3.1.1.1.9, supported by the server.
type SRPLevelInfo struct {
	// dwSRPLevel:   The numerical identifier of the software restriction policy level.
	// This MUST be between 0x00000000 and 0x00040000.
	SRPLevel uint32 `idl:"name:dwSRPLevel" json:"srp_level"`
	// wszFriendlyName:   A user-friendly display name for the software restriction policy
	// level.
	FriendlyName string `idl:"name:wszFriendlyName;string" json:"friendly_name"`
}

func (o *SRPLevelInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SRPLevelInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.SRPLevel); err != nil {
		return err
	}
	if o.FriendlyName != "" {
		_ptr_wszFriendlyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.FriendlyName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.FriendlyName, _ptr_wszFriendlyName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SRPLevelInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.SRPLevel); err != nil {
		return err
	}
	_ptr_wszFriendlyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.FriendlyName); err != nil {
			return err
		}
		return nil
	})
	_s_wszFriendlyName := func(ptr interface{}) { o.FriendlyName = *ptr.(*string) }
	if err := w.ReadPointer(&o.FriendlyName, _s_wszFriendlyName, _ptr_wszFriendlyName); err != nil {
		return err
	}
	return nil
}

// CatServerServices type represents CatSrvServices RPC enumeration.
//
// The CatSrvServices enumeration identifies the optional catalog-related capabilities
// of a COMA server that can be controlled dynamically by the ICapabilitySupport (section
// 3.1.4.19) interface. The current version of COMA defines one such capability, instance
// load balancing.
type CatServerServices uint16

var (
	// css_lb:   Identifies the instance load balancing capability.
	CatServerServicesLoadBalancer CatServerServices = 1
)

func (o CatServerServices) String() string {
	switch o {
	case CatServerServicesLoadBalancer:
		return "CatServerServicesLoadBalancer"
	}
	return "Invalid"
}

// CatServerServiceState type represents CatSrvServiceState RPC enumeration.
//
// The CatSrvServiceState enumeration identifies possible run-time states for instance
// load balancing.
type CatServerServiceState uint16

var (
	// css_serviceStopped:  Instance load balancing is not running.
	CatServerServiceStateStopped CatServerServiceState = 0
	// css_serviceStartPending:  Instance load balancing is not yet running, but it is
	// in the process of starting.
	CatServerServiceStateStartPending CatServerServiceState = 1
	// css_serviceStopPending:  Instance load balancing is running, but it is in the process
	// of stopping.
	CatServerServiceStateStopPending CatServerServiceState = 2
	// css_serviceRunning:  Instance load balancing is running.
	CatServerServiceStateRunning CatServerServiceState = 3
	// css_serviceContinuePending:  Instance load balancing is running, has been paused,
	// and is in the process of resuming.
	CatServerServiceStateContinuePending CatServerServiceState = 4
	// css_servicePausePending:  Instance load balancing is running, but it is in the process
	// of pausing.
	CatServerServiceStatePausePending CatServerServiceState = 5
	// css_servicePaused:  Instance load balancing is running, but it is paused.
	CatServerServiceStatePaused CatServerServiceState = 6
	// css_serviceUnknownState:   The server was unable to determine the state of instance
	// load balancing.
	CatServerServiceStateUnknownState CatServerServiceState = 7
)

func (o CatServerServiceState) String() string {
	switch o {
	case CatServerServiceStateStopped:
		return "CatServerServiceStateStopped"
	case CatServerServiceStateStartPending:
		return "CatServerServiceStateStartPending"
	case CatServerServiceStateStopPending:
		return "CatServerServiceStateStopPending"
	case CatServerServiceStateRunning:
		return "CatServerServiceStateRunning"
	case CatServerServiceStateContinuePending:
		return "CatServerServiceStateContinuePending"
	case CatServerServiceStatePausePending:
		return "CatServerServiceStatePausePending"
	case CatServerServiceStatePaused:
		return "CatServerServiceStatePaused"
	case CatServerServiceStateUnknownState:
		return "CatServerServiceStateUnknownState"
	}
	return "Invalid"
}

// InstanceContainer structure represents InstanceContainer RPC structure.
//
// The InstanceContainer structure represents an instance container.
type InstanceContainer struct {
	// ConglomerationID:   The conglomeration identifier of the conglomeration associated
	// with the instance container.
	ConglomerationID *dtyp.GUID `idl:"name:ConglomerationID" json:"conglomeration_id"`
	// PartitionID:   The partition identifier of the partition that contains the conglomeration
	// associated with the instance container.
	PartitionID *dtyp.GUID `idl:"name:PartitionID" json:"partition_id"`
	// ContainerID:   The activation of the instance container.
	ContainerID *dtyp.GUID `idl:"name:ContainerID" json:"container_id"`
	// dwProcessID:   The value of the instance container's ProcessIdentifier property,
	// as described in section 3.1.1.3.21.
	ProcessID uint32 `idl:"name:dwProcessID" json:"process_id"`
	// bPaused:   A flag that indicates whether or not the instance container is paused.
	Paused bool `idl:"name:bPaused" json:"paused"`
	// bRecycled:   A flag that indicates whether or not the instance container has been
	// recycled.
	Recycled bool `idl:"name:bRecycled" json:"recycled"`
}

func (o *InstanceContainer) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *InstanceContainer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.ConglomerationID != nil {
		if err := o.ConglomerationID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.PartitionID != nil {
		if err := o.PartitionID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ContainerID != nil {
		if err := o.ContainerID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ProcessID); err != nil {
		return err
	}
	if !o.Paused {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.Recycled {
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
func (o *InstanceContainer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.ConglomerationID == nil {
		o.ConglomerationID = &dtyp.GUID{}
	}
	if err := o.ConglomerationID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.PartitionID == nil {
		o.PartitionID = &dtyp.GUID{}
	}
	if err := o.PartitionID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ContainerID == nil {
		o.ContainerID = &dtyp.GUID{}
	}
	if err := o.ContainerID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.ProcessID); err != nil {
		return err
	}
	var _bPaused int32
	if err := w.ReadData(&_bPaused); err != nil {
		return err
	}
	o.Paused = _bPaused != 0
	var _bRecycled int32
	if err := w.ReadData(&_bRecycled); err != nil {
		return err
	}
	o.Recycled = _bRecycled != 0
	return nil
}

// Import structure represents IImport RPC structure.
//
// The IImport interface provides methods for importing, as specified in Export and
// Import (section 1.3.7), conglomerations and partitions from installer package files
// and returning information about installer package files. This interface inherits
// from IUnknown, as specified in [MS-DCOM] section 3.1.1.5.8.
//
// To receive incoming remote calls for this interface, the server MUST implement a
// DCOM Object Class with the CLSID CLSID_COMAServer (see section 1.9) using the UUID
// {C2BE6970-DF9E-11D1-8B87-00C04FD7A924} for this interface.
//
// This interface includes the following methods beyond those of IUnknown.
//
// Methods in RPC Opnum Order
//
//	+---------------------+---------------------------------------------------------------+
//	|                     |                                                               |
//	|       METHOD        |                          DESCRIPTION                          |
//	|                     |                                                               |
//	+---------------------+---------------------------------------------------------------+
//	+---------------------+---------------------------------------------------------------+
//	| ImportFromFile      | Imports a conglomeration from a file. Opnum: 3                |
//	+---------------------+---------------------------------------------------------------+
//	| QueryFile           | Returns information about an installer package file. Opnum: 4 |
//	+---------------------+---------------------------------------------------------------+
//	| Opnum5NotUsedOnWire | Reserved for local use. Opnum: 5                              |
//	+---------------------+---------------------------------------------------------------+
//	| Opnum6NotUsedOnWire | Reserved for local use. Opnum: 6                              |
//	+---------------------+---------------------------------------------------------------+
//
// In the previous table, the phrase "Reserved for local use" means that the client
// MUST NOT send the opnum and the server behavior is undefined since it does not affect
// interoperability.<334>
//
// All methods MUST NOT throw exceptions.
type Import dcom.InterfacePointer

func (o *Import) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Import) xxx_PreparePayload(ctx context.Context) error {
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

func (o *Import) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Import) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Import) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// CapabilitySupport structure represents ICapabilitySupport RPC structure.
//
// The ICapabilitySupport interface provides methods for starting and stopping optional,
// dynamically controllable, catalog-related capabilities of a COMA server. This version
// of COMA defines one such capability, instance load balancing (section 1.3.9). This
// interface inherits from IUnknown, as specified in [MS-DCOM] section 3.1.1.5.8.
//
// To receive incoming remote calls for this interface, the server MUST implement a
// DCOM Object Class with the CLSID CLSID_COMAServer, specified in section 1.9, using
// the UUID {47CDE9A1-0BF6-11D2-8016-00C04FB9988E} for this interface.
//
// This interface includes the following methods beyond those of IUnknown.
//
// Methods in RPC Opnum Order
//
//	+---------------------+---------------------------------------------------------------------------+
//	|                     |                                                                           |
//	|       METHOD        |                                DESCRIPTION                                |
//	|                     |                                                                           |
//	+---------------------+---------------------------------------------------------------------------+
//	+---------------------+---------------------------------------------------------------------------+
//	| Start               | Starts instance load balancing. Opnum: 3                                  |
//	+---------------------+---------------------------------------------------------------------------+
//	| Stop                | Stops instance load balancing. Opnum: 4                                   |
//	+---------------------+---------------------------------------------------------------------------+
//	| Opnum5NotUsedOnWire | Reserved for local use. Opnum: 5                                          |
//	+---------------------+---------------------------------------------------------------------------+
//	| Opnum6NotUsedOnWire | Reserved for local use. Opnum: 6                                          |
//	+---------------------+---------------------------------------------------------------------------+
//	| IsInstalled         | Determines whether instance load balancing support is installed. Opnum: 7 |
//	+---------------------+---------------------------------------------------------------------------+
//	| IsRunning           | Determines whether instance load balancing is running. Opnum: 8           |
//	+---------------------+---------------------------------------------------------------------------+
//	| Opnum9NotUsedOnWire | Reserved for local use. Opnum: 9                                          |
//	+---------------------+---------------------------------------------------------------------------+
//
// In the previous table, the phrase "Reserved for local use" means that the client
// MUST NOT send the opnum, and the server behavior is undefined since it does not affect
// interoperability.<348>
//
// All methods MUST NOT throw exceptions.
type CapabilitySupport dcom.InterfacePointer

func (o *CapabilitySupport) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *CapabilitySupport) xxx_PreparePayload(ctx context.Context) error {
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

func (o *CapabilitySupport) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *CapabilitySupport) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *CapabilitySupport) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Export2 structure represents IExport2 RPC structure.
//
// The IExport2 interface provides a method for exporting, as specified in Export and
// Import (section 1.3.7), a partition to an installer package file. This interface
// inherits from IUnknown, as specified in [MS-DCOM] section 3.1.1.5.8.
//
// To receive incoming remote calls for this interface, the server MUST implement a
// DCOM Object Class with the CLSID CLSID_COMAServer, as specified in section 1.9, using
// the UUID {F131EA3E-B7BE-480E-A60D-51CB2785779E} for this interface.
//
// This interface includes the following methods beyond those of IUnknown.
//
// Methods in RPC Opnum Order
//
//	+-----------------+------------------------------------------------------------+
//	|                 |                                                            |
//	|     METHOD      |                        DESCRIPTION                         |
//	|                 |                                                            |
//	+-----------------+------------------------------------------------------------+
//	+-----------------+------------------------------------------------------------+
//	| ExportPartition | Exports a partition to an installer package file. Opnum: 3 |
//	+-----------------+------------------------------------------------------------+
//
// All methods MUST NOT throw exceptions.
type Export2 dcom.InterfacePointer

func (o *Export2) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Export2) xxx_PreparePayload(ctx context.Context) error {
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

func (o *Export2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Export2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Export2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// CatalogUtils2 structure represents ICatalogUtils2 RPC structure.
//
// The ICatalogUtils2 interface is a miscellaneous utility interface. This interface
// inherits from IUnknown, as specified in [MS-DCOM] section 3.1.1.5.8.
//
// To receive incoming remote calls for this interface, the server MUST implement a
// DCOM Object Class with the CLSID CLSID_COMAServer, as specified in section 1.9, using
// the UUID {C726744E-5735-4F08-8286-C510EE638FB6} for this interface.
//
// This interface includes the following methods beyond those of IUnknown.
//
// Methods in RPC Opnum Order
//
//	+----------------------------+----------------------------------------------------------------------------------+
//	|                            |                                                                                  |
//	|           METHOD           |                                   DESCRIPTION                                    |
//	|                            |                                                                                  |
//	+----------------------------+----------------------------------------------------------------------------------+
//	+----------------------------+----------------------------------------------------------------------------------+
//	| CopyConglomerations        | Copies one or more conglomerations from one partition into another partition.    |
//	|                            | Opnum: 3                                                                         |
//	+----------------------------+----------------------------------------------------------------------------------+
//	| CopyComponentConfiguration | Copies a component configuration from one conglomeration into another            |
//	|                            | conglomeration. Opnum: 4                                                         |
//	+----------------------------+----------------------------------------------------------------------------------+
//	| AliasComponent             | Creates an alias component configuration. Opnum: 5                               |
//	+----------------------------+----------------------------------------------------------------------------------+
//	| MoveComponentConfiguration | Moves a component configuration from one conglomeration into another             |
//	|                            | conglomeration. Opnum: 6                                                         |
//	+----------------------------+----------------------------------------------------------------------------------+
//	| GetEventClassesForIID2     | Returns the event classes associated with an interface identifier (IID). Opnum:  |
//	|                            |                                                                                7 |
//	+----------------------------+----------------------------------------------------------------------------------+
//	| IsSafeToDelete             | Determines whether it is safe to delete a file. Opnum: 8                         |
//	+----------------------------+----------------------------------------------------------------------------------+
//	| FlushPartitionCache        | Flushes a server's local cache of partition user information. Opnum: 9           |
//	+----------------------------+----------------------------------------------------------------------------------+
//	| EnumerateSRPLevels         | Returns a list of software restriction policy levels supported by the server.    |
//	|                            | Opnum: 10                                                                        |
//	+----------------------------+----------------------------------------------------------------------------------+
//	| GetComponentVersions       | Returns a list of component configurations for a component. Opnum: 11            |
//	+----------------------------+----------------------------------------------------------------------------------+
//
// All methods MUST NOT throw exceptions.
type CatalogUtils2 dcom.InterfacePointer

func (o *CatalogUtils2) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *CatalogUtils2) xxx_PreparePayload(ctx context.Context) error {
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

func (o *CatalogUtils2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *CatalogUtils2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *CatalogUtils2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// CatalogSession structure represents ICatalogSession RPC structure.
//
// The ICatalogSession interface provides methods for Catalog Version Negotiation (section
// 3.1.4.1) and for Multiple-partition Support Capability Negotiation (section 3.1.4.3).
// This interface inherits from IUnknown, as specified in [MS-DCOM] section 3.1.1.5.8.
//
// To receive incoming remote calls for this interface, the server MUST implement a
// DCOM Object Class with the CLSID CLSID_COMAServer (see section 1.9) using the UUID
// {182C40FA-32E4-11D0-818B-00A0C9231C29} for this interface.
//
// Besides the methods of IUnknown, this interface includes the following methods.
//
// Methods in RPC Opnum Order
//
//	+----------------------+----------------------------------------------------------------------------------+
//	|                      |                                                                                  |
//	|        METHOD        |                                   DESCRIPTION                                    |
//	|                      |                                                                                  |
//	+----------------------+----------------------------------------------------------------------------------+
//	+----------------------+----------------------------------------------------------------------------------+
//	| Opnum3NotUsedOnWire  | Reserved for local use. Opnum: 3                                                 |
//	+----------------------+----------------------------------------------------------------------------------+
//	| Opnum4NotUsedOnWire  | Reserved for local use. Opnum: 4                                                 |
//	+----------------------+----------------------------------------------------------------------------------+
//	| Opnum5NotUsedOnWire  | Reserved for local use. Opnum: 5                                                 |
//	+----------------------+----------------------------------------------------------------------------------+
//	| Opnum6NotUsedOnWire  | Reserved for local use. Opnum: 6                                                 |
//	+----------------------+----------------------------------------------------------------------------------+
//	| InitializeSession    | Performs catalog version negotiation. Opnum: 7                                   |
//	+----------------------+----------------------------------------------------------------------------------+
//	| GetServerInformation | Performs capability negotiation for the multiple-partition support capability.   |
//	|                      | Opnum: 8                                                                         |
//	+----------------------+----------------------------------------------------------------------------------+
//
// In the previous table, the phrase "Reserved for local use" means that the client
// MUST NOT send the opnum, and the server behavior is undefined because it does not
// affect interoperability.<284>
//
// All methods MUST NOT throw exceptions.
type CatalogSession dcom.InterfacePointer

func (o *CatalogSession) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *CatalogSession) xxx_PreparePayload(ctx context.Context) error {
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

func (o *CatalogSession) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *CatalogSession) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *CatalogSession) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ReplicationUtil structure represents IReplicationUtil RPC structure.
//
// The IReplicationUtil interface provides methods for miscellaneous tasks specific
// to replication scenarios. This interface inherits from IUnknown, as specified in
// [MS-DCOM] section 3.1.1.5.8.
//
// To receive incoming remote calls for this interface, the server MUST implement a
// DCOM object class with the CLSID CLSID_COMAServer, as specified in section 1.9, using
// the UUID {98315903-7BE5-11D2-ADC1-00A02463D6E7} for this interface.
//
// This interface includes the following methods beyond those of IUnknown:
//
// Methods in RPC Opnum Order
//
//	+-----------------------------+----------------------------------------------------------------------------------+
//	|                             |                                                                                  |
//	|           METHOD            |                                   DESCRIPTION                                    |
//	|                             |                                                                                  |
//	+-----------------------------+----------------------------------------------------------------------------------+
//	+-----------------------------+----------------------------------------------------------------------------------+
//	| CreateShare                 | Creates a file share for copying installer package files. Opnum: 3               |
//	+-----------------------------+----------------------------------------------------------------------------------+
//	| CreateEmptyDir              | Creates an empty directory. Opnum: 4                                             |
//	+-----------------------------+----------------------------------------------------------------------------------+
//	| RemoveShare                 | Removes a previously created file share. Opnum: 5                                |
//	+-----------------------------+----------------------------------------------------------------------------------+
//	| BeginReplicationAsTarget    | Creates the file share for a new replication operation, optionally managing      |
//	|                             | replication history information and/or backup state. Opnum: 6                    |
//	+-----------------------------+----------------------------------------------------------------------------------+
//	| QueryConglomerationPassword | Returns the Password property of a conglomeration. Opnum: 7                      |
//	+-----------------------------+----------------------------------------------------------------------------------+
//	| CreateReplicationDir        | Ensures that the base replication directory exists, and returns its path. Opnum: |
//	|                             |                                                                                8 |
//	+-----------------------------+----------------------------------------------------------------------------------+
//
// All methods MUST NOT throw exceptions.
type ReplicationUtil dcom.InterfacePointer

func (o *ReplicationUtil) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *ReplicationUtil) xxx_PreparePayload(ctx context.Context) error {
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

func (o *ReplicationUtil) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ReplicationUtil) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ReplicationUtil) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Import2 structure represents IImport2 RPC structure.
//
// The IImport2 interface provides a method for setting the import target partition,
// as specified in Per-Session State (section 3.1.1.5). This interface inherits from
// IUnknown, as specified in [MS-DCOM] section 3.1.1.5.8.
//
// To receive incoming remote calls for this interface, the server MUST implement a
// DCOM Object Class with the CLSID CLSID_COMAServer, as specified in section 1.9, using
// the UUID {1F7B1697-ECB2-4CBB-8A0E-75C427F4A6F0} for this interface.
//
// This interface includes the following methods beyond those of IUnknown.
//
// Methods in RPC Opnum Order
//
//	+---------------------+--------------------------------------------------------------------------+
//	|                     |                                                                          |
//	|       METHOD        |                               DESCRIPTION                                |
//	|                     |                                                                          |
//	+---------------------+--------------------------------------------------------------------------+
//	+---------------------+--------------------------------------------------------------------------+
//	| SetPartition        | Sets the partition into which conglomerations will be imported. Opnum: 3 |
//	+---------------------+--------------------------------------------------------------------------+
//	| Opnum4NotUsedOnWire | Reserved for local use. Opnum: 4                                         |
//	+---------------------+--------------------------------------------------------------------------+
//	| Opnum5NotUsedOnWire | Reserved for local use. Opnum: 5                                         |
//	+---------------------+--------------------------------------------------------------------------+
//
// In the previous table, the phrase "Reserved for local use" means that the client
// MUST NOT send the opnum, and the server behavior is undefined because it does not
// affect interoperability.<339>
//
// All methods MUST NOT throw exceptions.
type Import2 dcom.InterfacePointer

func (o *Import2) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Import2) xxx_PreparePayload(ctx context.Context) error {
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

func (o *Import2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Import2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Import2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// CatalogUtils structure represents ICatalogUtils RPC structure.
//
// The ICatalogUtils interface is a miscellaneous utility interface. This interface
// inherits from IUnknown, as specified in [MS-DCOM] section 3.1.1.5.8.
//
// To receive incoming remote calls for this interface, the server MUST implement a
// DCOM Object Class with the CLSID CLSID_COMAServer, as specified in section 1.9, using
// the UUID {456129E2-1078-11D2-B0F9-00805FC73204} for this interface.
//
// This interface includes the following methods beyond those of IUnknown.
//
// Methods in RPC Opnum Order
//
//	+-----------------------+----------------------------------------------------------------------------------+
//	|                       |                                                                                  |
//	|        METHOD         |                                   DESCRIPTION                                    |
//	|                       |                                                                                  |
//	+-----------------------+----------------------------------------------------------------------------------+
//	+-----------------------+----------------------------------------------------------------------------------+
//	| ValidateUser          | Verifies a user account name and password. Opnum: 3                              |
//	+-----------------------+----------------------------------------------------------------------------------+
//	| WaitForEndWrites      | Waits for writes to the catalog to complete. Opnum: 4                            |
//	+-----------------------+----------------------------------------------------------------------------------+
//	| GetEventClassesForIID | Returns the event classes associated with an interface identifier (IID). Opnum:  |
//	|                       |                                                                                5 |
//	+-----------------------+----------------------------------------------------------------------------------+
//
// All methods MUST NOT throw exceptions.
type CatalogUtils dcom.InterfacePointer

func (o *CatalogUtils) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *CatalogUtils) xxx_PreparePayload(ctx context.Context) error {
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

func (o *CatalogUtils) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *CatalogUtils) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *CatalogUtils) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Catalog64BitSupport structure represents ICatalog64BitSupport RPC structure.
//
// The ICatalog64BitSupport interface provides methods for capability negotiation for
// the multiple-bitness capability, as specified in section 3.1.4.4, and the 64-bit
// QueryCell marshaling format capability, as specified in section 3.1.4.2. This interface
// inherits from IUnknown, as specified in [MS-DCOM] section 3.1.1.5.8.
//
// To receive incoming remote calls for this interface, the server MUST implement a
// DCOM Object Class with the CLSID CLSID_COMAServer, as specified in section 1.9, using
// the UUID {1D118904-94B3-4A64-9FA6-ED432666A7B9} for this interface.
//
// This interface includes the following methods beyond those of IUnknown.
//
// Methods in RPC Opnum Order
//
//	+---------------------------------+----------------------------------------------------------------------------------+
//	|                                 |                                                                                  |
//	|             METHOD              |                                   DESCRIPTION                                    |
//	|                                 |                                                                                  |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| SupportsMultipleBitness         | Performs capability negotiation for the multiple-bitness capability. Opnum: 3    |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| Initialize64BitQueryCellSupport | Performs capability negotiation for the 64-bit QueryCell marshaling format       |
//	|                                 | capability. Opnum: 4                                                             |
//	+---------------------------------+----------------------------------------------------------------------------------+
//
// All methods MUST NOT throw exceptions.
type Catalog64BitSupport dcom.InterfacePointer

func (o *Catalog64BitSupport) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *Catalog64BitSupport) xxx_PreparePayload(ctx context.Context) error {
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

func (o *Catalog64BitSupport) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Catalog64BitSupport) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Catalog64BitSupport) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// CatalogTableInfo structure represents ICatalogTableInfo RPC structure.
//
// The ICatalogTableInfo interface provides a method to retrieve table metadata, as
// specified in section 3.1.1.2.1, for a catalog table. This interface inherits from
// IUnknown, as specified in [MS-DCOM] section 3.1.1.5.8.
//
// To receive incoming remote calls for this interface, the server MUST implement a
// DCOM Object Class with the CLSID CLSID_COMAServer, as specified in section 1.9, using
// the UUID {A8927A41-D3CE-11D1-8472-006008B0E5CA} for this interface.
//
// This interface includes the following methods beyond those of IUnknown.
//
// Methods in RPC Opnum Order
//
//	+--------------------+----------------------------------------+
//	|                    |                                        |
//	|       METHOD       |              DESCRIPTION               |
//	|                    |                                        |
//	+--------------------+----------------------------------------+
//	+--------------------+----------------------------------------+
//	| GetClientTableInfo | Returns metadata for a table. Opnum: 3 |
//	+--------------------+----------------------------------------+
//
// All methods MUST NOT throw exceptions.
type CatalogTableInfo dcom.InterfacePointer

func (o *CatalogTableInfo) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *CatalogTableInfo) xxx_PreparePayload(ctx context.Context) error {
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

func (o *CatalogTableInfo) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *CatalogTableInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *CatalogTableInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Register structure represents IRegister RPC structure.
//
// The IRegister interface provides a method for registration, as specified in section
// 1.3.6. This interface inherits from IUnknown, as specified in [MS-DCOM] section 3.1.1.5.8.
//
// To receive incoming remote calls for this interface, the server MUST implement a
// DCOM Object Class with the CLSID CLSID_COMAServer, as specified in section 1.9, using
// the UUID {8DB2180E-BD29-11D1-8B7E-00C04FD7A924} for this interface.
//
// This interface includes the following methods beyond those of IUnknown.
//
// Methods in RPC Opnum Order
//
//	+---------------------+-----------------------------------------------------------+
//	|                     |                                                           |
//	|       METHOD        |                        DESCRIPTION                        |
//	|                     |                                                           |
//	+---------------------+-----------------------------------------------------------+
//	+---------------------+-----------------------------------------------------------+
//	| RegisterModule      | Registers the components in one or more modules. Opnum: 3 |
//	+---------------------+-----------------------------------------------------------+
//	| Opnum4NotUsedOnWire | Reserved for local use. Opnum: 4                          |
//	+---------------------+-----------------------------------------------------------+
//
// In the previous table, the phrase "Reserved for local use" means that the client
// MUST NOT send the opnum and the server behavior is undefined because it does not
// affect interoperability.<316>
type Register dcom.InterfacePointer

func (o *Register) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Register) xxx_PreparePayload(ctx context.Context) error {
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

func (o *Register) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Register) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Register) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// AlternateLaunch structure represents IAlternateLaunch RPC structure.
//
// The IAlternateLaunch interface provides methods for creating and deleting Alternate
// launch configurations (section 3.1.1.4). This interface inherits from IUnknown, as
// specified in [MS-DCOM] section 3.1.1.5.8.
//
// To receive incoming remote calls for this interface, the server MUST implement a
// DCOM Object Class with the CLSID CLSID_COMAServer, as specified in section 1.9, using
// the UUID {7F43B400-1A0E-4D57-BBC9-6B0C65F7A889} for this interface.
//
// Methods in RPC Opnum Order
//
//	+---------------------+----------------------------------------------------------------------------+
//	|                     |                                                                            |
//	|       METHOD        |                                DESCRIPTION                                 |
//	|                     |                                                                            |
//	+---------------------+----------------------------------------------------------------------------+
//	+---------------------+----------------------------------------------------------------------------+
//	| CreateConfiguration | Creates an alternate launch configuration for a conglomeration. Opnum: 3   |
//	+---------------------+----------------------------------------------------------------------------+
//	| DeleteConfiguration | Removes an alternative launch configuration for a conglomeration. Opnum: 4 |
//	+---------------------+----------------------------------------------------------------------------+
//
// All methods MUST NOT throw exceptions.
type AlternateLaunch dcom.InterfacePointer

func (o *AlternateLaunch) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *AlternateLaunch) xxx_PreparePayload(ctx context.Context) error {
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

func (o *AlternateLaunch) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AlternateLaunch) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *AlternateLaunch) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ContainerControl structure represents IContainerControl RPC structure.
//
// The IContainerControl interface provides methods for creating and controlling an
// InstanceContainer (section 2.2.9). This interface inherits from IUnknown, as specified
// in [MS-DCOM] section 3.1.1.5.8.
//
// To receive incoming remote calls for this interface, the server MUST implement a
// DCOM object class with the CLSID CLSID_COMAServer, as specified in section 1.9, using
// the UUID {3F3B1B86-DBBE-11D1-9DA6-00805F85CFE3} for this interface.
//
// This interface includes the following methods beyond those of IUnknown.
//
// Methods in RPC Opnum Order
//
//	+--------------------+-------------------------------------------------------------------+
//	|                    |                                                                   |
//	|       METHOD       |                            DESCRIPTION                            |
//	|                    |                                                                   |
//	+--------------------+-------------------------------------------------------------------+
//	+--------------------+-------------------------------------------------------------------+
//	| CreateContainer    | Creates an instance container for a conglomeration. Opnum: 3      |
//	+--------------------+-------------------------------------------------------------------+
//	| ShutdownContainers | Shuts down all instance containers for a conglomeration. Opnum: 4 |
//	+--------------------+-------------------------------------------------------------------+
//	| RefreshComponents  | Performs implementation-specific repairs to the catalog. Opnum: 5 |
//	+--------------------+-------------------------------------------------------------------+
//
// All methods MUST NOT throw exceptions.
type ContainerControl dcom.InterfacePointer

func (o *ContainerControl) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *ContainerControl) xxx_PreparePayload(ctx context.Context) error {
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

func (o *ContainerControl) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ContainerControl) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ContainerControl) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ContainerControl2 structure represents IContainerControl2 RPC structure.
//
// The IContainerControl2 interface provides methods for controlling InstanceContainers
// (section 2.2.9). This interface inherits from IUnknown, as specified in [MS-DCOM]
// section 3.1.1.5.8.
//
// To receive incoming remote calls for this interface, the server MUST implement a
// DCOM Object Class with the CLSID CLSID_COMAServer, specified in section 1.9, using
// the UUID {6C935649-30A6-4211-8687-C4C83E5FE1C7} for this interface.
//
// This interface includes the following methods beyond those of IUnknown.
//
// Methods in RPC Opnum Order
//
//	+------------------------------------+----------------------------------------------------------------------------------+
//	|                                    |                                                                                  |
//	|               METHOD               |                                   DESCRIPTION                                    |
//	|                                    |                                                                                  |
//	+------------------------------------+----------------------------------------------------------------------------------+
//	+------------------------------------+----------------------------------------------------------------------------------+
//	| ShutdownContainer                  | Shuts down an instance container. Opnum: 3                                       |
//	+------------------------------------+----------------------------------------------------------------------------------+
//	| PauseContainer                     | Pauses an instance container. Opnum: 4                                           |
//	+------------------------------------+----------------------------------------------------------------------------------+
//	| ResumeContainer                    | Resumes a paused instance container. Opnum: 5                                    |
//	+------------------------------------+----------------------------------------------------------------------------------+
//	| IsContainerPaused                  | Determines whether an instance container is paused. Opnum: 6                     |
//	+------------------------------------+----------------------------------------------------------------------------------+
//	| GetRunningContainers               | Returns a list of running instance containers for a conglomeration. Opnum: 7     |
//	+------------------------------------+----------------------------------------------------------------------------------+
//	| GetContainerIDFromProcessID        | Returns the instance container identifier corresponding to a process identifier. |
//	|                                    | Opnum: 8                                                                         |
//	+------------------------------------+----------------------------------------------------------------------------------+
//	| RecycleContainer                   | Forces an instance container to be recycled. Opnum: 9                            |
//	+------------------------------------+----------------------------------------------------------------------------------+
//	| GetContainerIDFromConglomerationID | Returns the instance container identifier of an instance container for a         |
//	|                                    | conglomeration. Opnum: 10                                                        |
//	+------------------------------------+----------------------------------------------------------------------------------+
//
// All methods MUST NOT throw exceptions.
type ContainerControl2 dcom.InterfacePointer

func (o *ContainerControl2) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *ContainerControl2) xxx_PreparePayload(ctx context.Context) error {
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

func (o *ContainerControl2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ContainerControl2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ContainerControl2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Register2 structure represents IRegister2 RPC structure.
//
// The IRegister2 interface provides methods for registration, as specified in section
// 1.3.6, and creating component configurations. This interface inherits from IUnknown,
// as specified in [MS-DCOM] section 3.1.1.5.8.
//
// To receive incoming remote calls for this interface, the server MUST implement a
// DCOM Object Class with the CLSID CLSID_COMAServer, as specified in section 1.9, using
// the UUID {971668DC-C3FE-4EA1-9643-0C7230F494A1} for this interface.
//
// This interface includes the following methods beyond those of IUnknown.
//
// Methods in RPC Opnum Order
//
//	+----------------------------+----------------------------------------------------------------------------------+
//	|                            |                                                                                  |
//	|           METHOD           |                                   DESCRIPTION                                    |
//	|                            |                                                                                  |
//	+----------------------------+----------------------------------------------------------------------------------+
//	+----------------------------+----------------------------------------------------------------------------------+
//	| CreateFullConfiguration    | Creates a component full configuration for an existing component. Opnum: 3       |
//	+----------------------------+----------------------------------------------------------------------------------+
//	| CreateLegacyConfiguration  | Creates a component legacy configuration for an existing component. Opnum: 4     |
//	+----------------------------+----------------------------------------------------------------------------------+
//	| PromoteLegacyConfiguration | Converts a component legacy configuration into a component full configuration.   |
//	|                            | Opnum: 5                                                                         |
//	+----------------------------+----------------------------------------------------------------------------------+
//	| Opnum6NotUsedOnWire        | Reserved for local use. Opnum: 6                                                 |
//	+----------------------------+----------------------------------------------------------------------------------+
//	| Opnum7NotUsedOnWire        | Reserved for local use. Opnum: 7                                                 |
//	+----------------------------+----------------------------------------------------------------------------------+
//	| RegisterModule2            | Registers the components in one or more modules. Opnum: 8                        |
//	+----------------------------+----------------------------------------------------------------------------------+
//	| Opnum9NotUsedOnWire        | Reserved for local use. Opnum: 9                                                 |
//	+----------------------------+----------------------------------------------------------------------------------+
//
// In the previous table, the phrase "Reserved for local use" means that the client
// MUST NOT send the opnum, and the server behavior is undefined since it does not affect
// interoperability.<324>
//
// All methods MUST NOT throw exceptions.
type Register2 dcom.InterfacePointer

func (o *Register2) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Register2) xxx_PreparePayload(ctx context.Context) error {
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

func (o *Register2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Register2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Register2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Export structure represents IExport RPC structure.
//
// The IExport interface provides a method for exporting, as specified in Export and
// Import section 1.3.7, a conglomeration to an installer package file. This interface
// inherits from IUnknown, as specified in [MS-DCOM] section 3.1.1.5.8.
//
// To receive incoming remote calls for this interface, the server MUST implement a
// DCOM Object Class with the CLSID CLSID_COMAServer, as specified in section 1.9, using
// the UUID {CFADAC84-E12C-11D1-B34C-00C04F990D54} for this interface.
//
// This interface includes the following methods beyond those of IUnknown.
//
// Methods in RPC Opnum Order
//
//	+----------------------+-----------------------------------------------------------------+
//	|                      |                                                                 |
//	|        METHOD        |                           DESCRIPTION                           |
//	|                      |                                                                 |
//	+----------------------+-----------------------------------------------------------------+
//	+----------------------+-----------------------------------------------------------------+
//	| ExportConglomeration | Exports a conglomeration to an installer package file. Opnum: 3 |
//	+----------------------+-----------------------------------------------------------------+
//	| Opnum4NotUsedOnWire  | Reserved for local use. Opnum: 4                                |
//	+----------------------+-----------------------------------------------------------------+
//	| Opnum5NotUsedOnWire  | Reserved for local use. Opnum: 5                                |
//	+----------------------+-----------------------------------------------------------------+
//	| Opnum6NotUsedOnWire  | Reserved for local use. Opnum: 6                                |
//	+----------------------+-----------------------------------------------------------------+
//
// In the previous table, the phrase "Reserved for local use" means that the client
// MUST NOT send the opnum and the server behavior is undefined because it does not
// affect interoperability.<340>
//
// All methods MUST NOT throw exceptions.
type Export dcom.InterfacePointer

func (o *Export) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Export) xxx_PreparePayload(ctx context.Context) error {
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

func (o *Export) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Export) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Export) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// CatalogTableRead structure represents ICatalogTableRead RPC structure.
//
// The ICatalogTableRead interface provides a method to read entries from a catalog
// table. This interface inherits from IUnknown, as specified in [MS-DCOM] section 3.1.1.5.8.
//
// To receive incoming remote calls for this interface, the server MUST implement a
// DCOM Object Class with the CLSID CLSID_COMAServer, as specified in section 1.9, using
// the UUID {0E3D6630-B46B-11D1-9D2D-006008B0E5CA} for this interface.
//
// This interface includes the following methods beyond those of IUnknown.
//
// Methods in RPC Opnum Order
//
//	+-----------+--------------------------------------+
//	|           |                                      |
//	|  METHOD   |             DESCRIPTION              |
//	|           |                                      |
//	+-----------+--------------------------------------+
//	+-----------+--------------------------------------+
//	| ReadTable | Reads entries from a table. Opnum: 3 |
//	+-----------+--------------------------------------+
//
// All methods MUST NOT throw exceptions.
type CatalogTableRead dcom.InterfacePointer

func (o *CatalogTableRead) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *CatalogTableRead) xxx_PreparePayload(ctx context.Context) error {
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

func (o *CatalogTableRead) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *CatalogTableRead) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *CatalogTableRead) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// CatalogTableWrite structure represents ICatalogTableWrite RPC structure.
//
// The ICatalogTableWrite interface provides a method to write entries to a catalog
// table. This interface inherits from IUnknown, as specified in [MS-DCOM] section 3.1.1.5.8.
//
// To receive incoming remote calls for this interface, the server MUST implement a
// DCOM Object Class with the CLSID CLSID_COMAServer, as specified in section 1.9, using
// the UUID {0E3D6631-B46B-11D1-9D2D-006008B0E5CA} for this interface.
//
// This interface includes the following methods beyond those of IUnknown.
//
// Methods in RPC Opnum Order
//
//	+------------+-------------------------------------+
//	|            |                                     |
//	|   METHOD   |             DESCRIPTION             |
//	|            |                                     |
//	+------------+-------------------------------------+
//	+------------+-------------------------------------+
//	| WriteTable | Writes entries to a table. Opnum: 3 |
//	+------------+-------------------------------------+
//
// All methods MUST NOT throw exceptions.
type CatalogTableWrite dcom.InterfacePointer

func (o *CatalogTableWrite) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *CatalogTableWrite) xxx_PreparePayload(ctx context.Context) error {
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

func (o *CatalogTableWrite) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *CatalogTableWrite) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *CatalogTableWrite) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
