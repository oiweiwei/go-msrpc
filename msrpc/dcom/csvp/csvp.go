// The csvp package implements the CSVP client protocol.
//
// # Introduction
//
// The Failover Cluster: Setup and Validation Protocol (ClusPrep) consists of DCOM interfaces,
// as specified in [MS-DCOM], that are used for remotely configuring cluster nodes cleaning
// up cluster nodes and validating that hardware and software settings are compatible
// with use in a failover cluster.
//
// # Overview
//
// The Failover Cluster: Setup and Validation Protocol (ClusPrep) is a COM/DCOM protocol
// to setup, cleanup, and validate a set of machines for failover cluster capability.
// The setup of state includes pushing to a new server the state required to participate
// in the failover cluster. Validation of state includes retrieving failover cluster
// state from a server to validate the state and directing the server to perform specific
// tasks and report status to validate it is functioning correctly. Cleanup of state
// includes removing restrictions on accessing shared disks and deletion of cluster
// state from the server.
//
// The ClusPrep Protocol provides DCOM interfaces that enable a client to:
//
// * Validate the server configuration so as to make it eligible to become a node (
// b30bcd1c-8ff6-46d2-a27e-61bd85ace9f4#gt_762051d8-4fdc-437e-af9d-3f4da77c3c7d ) in
// a failover cluster ( b30bcd1c-8ff6-46d2-a27e-61bd85ace9f4#gt_acd8b49c-8762-48fd-9272-26a03643322b
// ).
//
// * Configure a server to no longer be a node in a failover cluster.
//
// * Retrieve log information from a node in a failover cluster.
//
// * Determine whether the hardware/software settings of a server meet the requirements
// to be part of a failover cluster.
package csvp

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
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
	_ = oaut.GoPackage
	_ = dcom.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/csvp"
)

// ClusterDiskIDEnum type represents CPREP_DISKID_ENUM RPC enumeration.
//
// The CPREP_DISKID_ENUM enumeration defines possible kinds of disk identifiers.
type ClusterDiskIDEnum uint16

var (
	// CprepIdSignature:  A small computer system interface (SCSI) signature that is 4
	// bytes in length. Used to identify master boot record (MBR) disks.
	ClusterDiskIDEnumSignature ClusterDiskIDEnum = 0
	// CprepIdGuid:  A signature of a GUID partitioning table (GPT) disk, which is a GUID.
	// A GUID, also known as a UUID, is a 16-byte structure, intended to serve as a unique
	// identifier for an object.
	ClusterDiskIDEnumGUID ClusterDiskIDEnum = 1
	// CprepIdNumber:  The number by which the disk is identified.
	ClusterDiskIDEnumNumber ClusterDiskIDEnum = 4000
	// CprepIdUnknown:  Used for disks that are not identified via one of the previously
	// described ways.
	ClusterDiskIDEnumUnknown ClusterDiskIDEnum = 5000
)

func (o ClusterDiskIDEnum) String() string {
	switch o {
	case ClusterDiskIDEnumSignature:
		return "ClusterDiskIDEnumSignature"
	case ClusterDiskIDEnumGUID:
		return "ClusterDiskIDEnumGUID"
	case ClusterDiskIDEnumNumber:
		return "ClusterDiskIDEnumNumber"
	case ClusterDiskIDEnumUnknown:
		return "ClusterDiskIDEnumUnknown"
	}
	return "Invalid"
}

// ClusterDiskID structure represents CPREP_DISKID RPC structure.
//
// The CPREP_DISKID structure identifies an operating system disk and typically corresponds
// to a LUN. This structure holds either the operating system disk number (not the BIOS
// disk number) or the disk signature .
type ClusterDiskID struct {
	// DiskIdType:  This MUST be one of the valid CPREP_DISKID_ENUM values.
	DiskIDType    ClusterDiskIDEnum            `idl:"name:DiskIdType" json:"disk_id_type"`
	ClusterDiskID *ClusterDiskID_ClusterDiskID `idl:"name:ClusterDiskID;switch_is:DiskIdType" json:"cluster_disk_id"`
}

func (o *ClusterDiskID) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClusterDiskID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.DiskIDType)); err != nil {
		return err
	}
	_swClusterDiskID := uint16(o.DiskIDType)
	if o.ClusterDiskID != nil {
		if err := o.ClusterDiskID.MarshalUnionNDR(ctx, w, _swClusterDiskID); err != nil {
			return err
		}
	} else {
		if err := (&ClusterDiskID_ClusterDiskID{}).MarshalUnionNDR(ctx, w, _swClusterDiskID); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClusterDiskID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.DiskIDType)); err != nil {
		return err
	}
	if o.ClusterDiskID == nil {
		o.ClusterDiskID = &ClusterDiskID_ClusterDiskID{}
	}
	_swClusterDiskID := uint16(o.DiskIDType)
	if err := o.ClusterDiskID.UnmarshalUnionNDR(ctx, w, _swClusterDiskID); err != nil {
		return err
	}
	return nil
}

// ClusterDiskID_ClusterDiskID structure represents CPREP_DISKID union anonymous member.
//
// The CPREP_DISKID structure identifies an operating system disk and typically corresponds
// to a LUN. This structure holds either the operating system disk number (not the BIOS
// disk number) or the disk signature .
type ClusterDiskID_ClusterDiskID struct {
	// Types that are assignable to Value
	//
	// *ClusterDiskID_DiskSignature
	// *ClusterDiskID_DiskGUID
	// *ClusterDiskID_DeviceNumber
	// *ClusterDiskID_Junk
	Value is_ClusterDiskID_ClusterDiskID `json:"value"`
}

func (o *ClusterDiskID_ClusterDiskID) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *ClusterDiskID_DiskSignature:
		if value != nil {
			return value.DiskSignature
		}
	case *ClusterDiskID_DiskGUID:
		if value != nil {
			return value.DiskGUID
		}
	case *ClusterDiskID_DeviceNumber:
		if value != nil {
			return value.DeviceNumber
		}
	case *ClusterDiskID_Junk:
		if value != nil {
			return value.Junk
		}
	}
	return nil
}

type is_ClusterDiskID_ClusterDiskID interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_ClusterDiskID_ClusterDiskID()
}

func (o *ClusterDiskID_ClusterDiskID) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *ClusterDiskID_DiskSignature:
		return uint16(0)
	case *ClusterDiskID_DiskGUID:
		return uint16(1)
	case *ClusterDiskID_DeviceNumber:
		return uint16(4000)
	case *ClusterDiskID_Junk:
		return uint16(5000)
	}
	return uint16(0)
}

func (o *ClusterDiskID_ClusterDiskID) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(0):
		_o, _ := o.Value.(*ClusterDiskID_DiskSignature)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ClusterDiskID_DiskSignature{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(1):
		_o, _ := o.Value.(*ClusterDiskID_DiskGUID)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ClusterDiskID_DiskGUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(4000):
		_o, _ := o.Value.(*ClusterDiskID_DeviceNumber)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ClusterDiskID_DeviceNumber{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(5000):
		_o, _ := o.Value.(*ClusterDiskID_Junk)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ClusterDiskID_Junk{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *ClusterDiskID_ClusterDiskID) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch((*uint16)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(0):
		o.Value = &ClusterDiskID_DiskSignature{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(1):
		o.Value = &ClusterDiskID_DiskGUID{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(4000):
		o.Value = &ClusterDiskID_DeviceNumber{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(5000):
		o.Value = &ClusterDiskID_Junk{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// ClusterDiskID_DiskSignature structure represents ClusterDiskID_ClusterDiskID RPC union arm.
//
// It has following labels: 0
type ClusterDiskID_DiskSignature struct {
	// DiskSignature:  This field is valid only if DiskIdType is CprepIdSignature. It MUST
	// contain the 4-byte signature of the disk. How the disk signature is assigned is implementation-specific.
	DiskSignature uint32 `idl:"name:DiskSignature" json:"disk_signature"`
}

func (*ClusterDiskID_DiskSignature) is_ClusterDiskID_ClusterDiskID() {}

func (o *ClusterDiskID_DiskSignature) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.DiskSignature); err != nil {
		return err
	}
	return nil
}
func (o *ClusterDiskID_DiskSignature) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.DiskSignature); err != nil {
		return err
	}
	return nil
}

// ClusterDiskID_DiskGUID structure represents ClusterDiskID_ClusterDiskID RPC union arm.
//
// It has following labels: 1
type ClusterDiskID_DiskGUID struct {
	// DiskGuid:   This field is valid only if DiskIdType is CprepIdGuid. It MUST contain
	// the GUID that identifies the disk. How the disk GUID is assigned is implementation-specific.
	DiskGUID *dtyp.GUID `idl:"name:DiskGuid" json:"disk_guid"`
}

func (*ClusterDiskID_DiskGUID) is_ClusterDiskID_ClusterDiskID() {}

func (o *ClusterDiskID_DiskGUID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.DiskGUID != nil {
		if err := o.DiskGUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClusterDiskID_DiskGUID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.DiskGUID == nil {
		o.DiskGUID = &dtyp.GUID{}
	}
	if err := o.DiskGUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ClusterDiskID_DeviceNumber structure represents ClusterDiskID_ClusterDiskID RPC union arm.
//
// It has following labels: 4000
type ClusterDiskID_DeviceNumber struct {
	// DeviceNumber:  This field is valid only if DiskIdType is CprepIdNumber. It MUST contain
	// the operating system disk number, not the BIOS disk number. The device number ranges
	// from zero to the number of disks accessible by the server minus one. How the device
	// number is assigned is implementation-specific.
	DeviceNumber uint32 `idl:"name:DeviceNumber" json:"device_number"`
}

func (*ClusterDiskID_DeviceNumber) is_ClusterDiskID_ClusterDiskID() {}

func (o *ClusterDiskID_DeviceNumber) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.DeviceNumber); err != nil {
		return err
	}
	return nil
}
func (o *ClusterDiskID_DeviceNumber) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.DeviceNumber); err != nil {
		return err
	}
	return nil
}

// ClusterDiskID_Junk structure represents ClusterDiskID_ClusterDiskID RPC union arm.
//
// It has following labels: 5000
type ClusterDiskID_Junk struct {
	// Junk:   This field is valid only if DiskIdType is CprepIdUnknown. The value of this
	// field is not used.
	Junk uint32 `idl:"name:Junk" json:"junk"`
}

func (*ClusterDiskID_Junk) is_ClusterDiskID_ClusterDiskID() {}

func (o *ClusterDiskID_Junk) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Junk); err != nil {
		return err
	}
	return nil
}
func (o *ClusterDiskID_Junk) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Junk); err != nil {
		return err
	}
	return nil
}

// DiskStackType type represents DiskStackType RPC enumeration.
//
// The DiskStackType enumeration defines valid driver types that a disk driver is implemented
// as.
type DiskStackType uint16

var (
	// DiskStackScsiPort:  The driver is a SCSIPort driver.
	DiskStackTypeSCSIPort DiskStackType = 0
	// DiskStackStorPort:  The driver is a StorPort driver.
	DiskStackTypeStorPort DiskStackType = 1
	// DiskStackFullPort:  The driver is a monolithic driver and does not conform to any
	// storage driver submodel.
	DiskStackTypeFullPort DiskStackType = 2
)

func (o DiskStackType) String() string {
	switch o {
	case DiskStackTypeSCSIPort:
		return "DiskStackTypeSCSIPort"
	case DiskStackTypeStorPort:
		return "DiskStackTypeStorPort"
	case DiskStackTypeFullPort:
		return "DiskStackTypeFullPort"
	}
	return "Invalid"
}

// ClusterSCSIAddress structure represents CPREP_SCSI_ADDRESS RPC structure.
//
// The CPREP_SCSI_ADDRESS structure holds information to identify a disk via the SCSI
// protocol. The structure is included in this document because it is referenced by
// the DISK_PROPS structure; however, the values in this structure are never read by
// the client.
type ClusterSCSIAddress struct {
	// Length:  Contains the length of this structure in bytes.
	Length uint32 `idl:"name:Length" json:"length"`
	// PortNumber:  Contains the number of the SCSI adapter.
	PortNumber uint8 `idl:"name:PortNumber" json:"port_number"`
	// PathId:  Contains the number of the bus.
	PathID uint8 `idl:"name:PathId" json:"path_id"`
	// TargetId:  Contains the number of the target device.
	TargetID uint8 `idl:"name:TargetId" json:"target_id"`
	// Lun:  Contains the logical unit number.
	LUN uint8 `idl:"name:Lun" json:"lun"`
}

func (o *ClusterSCSIAddress) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClusterSCSIAddress) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if err := w.WriteData(o.PortNumber); err != nil {
		return err
	}
	if err := w.WriteData(o.PathID); err != nil {
		return err
	}
	if err := w.WriteData(o.TargetID); err != nil {
		return err
	}
	if err := w.WriteData(o.LUN); err != nil {
		return err
	}
	return nil
}
func (o *ClusterSCSIAddress) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	if err := w.ReadData(&o.PortNumber); err != nil {
		return err
	}
	if err := w.ReadData(&o.PathID); err != nil {
		return err
	}
	if err := w.ReadData(&o.TargetID); err != nil {
		return err
	}
	if err := w.ReadData(&o.LUN); err != nil {
		return err
	}
	return nil
}

// DiskProperties structure represents DISK_PROPS RPC structure.
//
// The DISK_PROPS structure holds information about a single disk.<3>
type DiskProperties struct {
	// DiskNumber:  The zero-based device number assigned to the disk by the operating system.
	DiskNumber uint32 `idl:"name:DiskNumber" json:"disk_number"`
	// DiskId:  A valid CPREP_DISKID structure with the correct identifier for the disk.
	DiskID *ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	// DiskBusType:  The type of bus to which the disk is attached. It MAY<4> be one of
	// the following values.
	//
	//	+-------------------------------------+--------------------------------------------------------------+
	//	|                                     |                                                              |
	//	|                VALUE                |                           MEANING                            |
	//	|                                     |                                                              |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusTypeUnknown 0x00000000           | The bus type is not one of those that follows.               |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusTypeScsi 0x00000001              | The bus type is SCSI.                                        |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusTypeAtapi 0x00000002             | The bus type is AT attachment packet interface (ATAPI).      |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusTypeAta 0x00000003               | The bus type is advanced technology attachment (ATA).        |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusType1394 0x00000004              | The bus type is IEEE 1394.                                   |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusTypeSsa 0x00000005               | The bus type is serial storage architecture (SSA).           |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusTypeFibre 0x00000006             | The bus type is Fibre Channel.                               |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusTypeUsb 0x00000007               | The bus type is universal serial bus (USB).                  |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusTypeRAID 0x00000008              | The bus type is redundant array of independent disks (RAID). |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusTypeiScsi 0x00000009             | The bus type is iSCSI.                                       |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusTypeSas 0x0000000A               | The bus type is Serial Attached SCSI (SAS).                  |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusTypeSata 0x0000000B              | The bus type is Serial ATA (SATA).                           |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusTypeSd 0x0000000C                | The bus type is Sd.                                          |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusTypeMmc 0x0000000D               | The bus type is Mmc.                                         |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusTypeVirtual 0x0000000E           | The bus type is Virtual.                                     |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusTypeFileBackedVirtual 0x0000000F | The bus type is File Backed Virtual.                         |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusTypeSpaces 0x00000010            | The bus is type Spaces.                                      |
	//	+-------------------------------------+--------------------------------------------------------------+
	DiskBusType uint32 `idl:"name:DiskBusType" json:"disk_bus_type"`
	// StackType:  The driver subtype of the device driver. It MUST be one of the valid
	// values for DiskStackType.
	StackType DiskStackType `idl:"name:StackType" json:"stack_type"`
	// ScsiAddress:  The SCSI address of the disk. It MUST be a valid CPREP_SCSI_ADDRESS.
	SCSIAddress *ClusterSCSIAddress `idl:"name:ScsiAddress" json:"scsi_address"`
	// DiskIsClusterable:  A Boolean flag that indicates whether the disk can be represented
	// by a storage class resource in a failover cluster, as specified in [MS-CMRP]. A value
	// of TRUE or 1 indicates that the disk can be represented by a storage class resource.
	// A value of FALSE or 0 indicates that the disk cannot be represented by a storage
	// class resource. The value of the DiskIsClusterable member can be determined in an
	// implementation-specific way.
	DiskIsClusterable int32 `idl:"name:DiskIsClusterable" json:"disk_is_clusterable"`
	// AdapterDesc:  A user-friendly description of the adapter to which the disk is connected.
	AdapterDesc []uint16 `idl:"name:AdapterDesc" json:"adapter_desc"`
	// NumPaths:  The number of IO paths to the disk. A Multipath I/O (MPIO) disk has a
	// number greater than 1.
	PathsLength uint32 `idl:"name:NumPaths" json:"paths_length"`
	// Flags:  Information about the disk. It MAY<5> be one or more of the following values
	// bitwise OR'd together.
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                                         |                                                                                  |
	//	|                  VALUE                  |                                     MEANING                                      |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_BOOT 0x00000001                    | The disk is the boot device.                                                     |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_SYSTEM 0x00000002                  | The disk contains the operating system.                                          |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_PAGEFILE 0x00000004                | The disk contains an operating system pagefile.                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_HIBERNATE 0x00000008               | The disk will be used to store system hibernation data.                          |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_CRASHDUMP 0x00000010               | The disk will be used to store system crash dump data.                           |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_REMOVABLE 0x00000020               | The disk is on removable media.                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_CLUSTERNOSUPP 0x00000040           | The disk is not supported by the cluster implementation. The criteria for        |
	//	|                                         | support are implementation-specific.                                             |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_BUSNOSUPP 0x00000100               | The disk is on a bus not supported by the cluster implementation. The criteria   |
	//	|                                         | for support are implementation-specific.                                         |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_SYSTEMBUS 0x00000200               | The disk is on the same bus as the disk containing the operating system.         |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_ALREADY_CLUSTERED 0x00000400       | The disk is already controlled by the cluster.                                   |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_SYTLE_MBR 0x00001000               | The disk is MBR.                                                                 |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_STYLE_GPT 0x00002000               | The disk is GPT.                                                                 |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_STYLE_RAW 0x00004000               | The disk is neither MBR nor GPT.                                                 |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_PART_BASIC 0x00008000              | The disk is configured with basic volumes.                                       |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_PART_DYNAMIC 0x00010000            | The disk is configured with dynamic volumes.                                     |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_CLUSTERED_ONLINE 0x00020000        | The disk is controlled by the cluster and is online.                             |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_UNREADABLE 0x00040000              | The disk cannot be read.                                                         |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_MPIO 0x00080000                    | The disk is controlled by MPIO.                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_CLUSTERED_OTHER 0x00100000         | The disk is controlled by cluster software other than the failover cluster       |
	//	|                                         | implementation.                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_MISSING 0x00200000                 | The disk could not be found.                                                     |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_REDUNDANT 0x00400000               | The disk is exposed to the operating system multiple times through redundant     |
	//	|                                         | paths.                                                                           |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_SNAPSHOT 0x00800000                | The disk is a snapshot disk.                                                     |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_FAILING_IO 0x02000000              | The disk is unable to gather disk information.                                   |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_NO_PAGE83 0x04000000               | The disk does not have a Device Identification VPD page (see [SPC-3] section     |
	//	|                                         | 7.6.3) with PAGE CODE (see [SPC-3] table 294) set to 83h, a device ASSOCIATION   |
	//	|                                         | (see [SPC-3] table 297), and IDENTIFIER TYPE (see [SPC-3] table 298) of Type 8,  |
	//	|                                         | Type 3, or Type 2.                                                               |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_COLLISION 0x08000000               | The disk's signature collides with the signature on another disk visible to this |
	//	|                                         | server, and disk signature collision resolution is disabled.                     |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_OUTOFSPACE 0x10000000              | The disk is a thin-provisioned LUN that has no free space.                       |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_POOL_DRIVE 0x20000000              | The disk is a member of a storage pool.                                          |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_POOL_DRIVE_NOT_TESTABLE 0x40000000 | The disk is a member of a storage pool and cannot be tested because the storage  |
	//	|                                         | pool is in use.                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_POOL_CLUSTERED 0x80000000          | The disk is member of a storage pool and the storage pool to which it belongs is |
	//	|                                         | a cluster resource.                                                              |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
}

func (o *DiskProperties) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DiskProperties) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DiskNumber); err != nil {
		return err
	}
	if o.DiskID != nil {
		if err := o.DiskID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ClusterDiskID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.DiskBusType); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.StackType)); err != nil {
		return err
	}
	if o.SCSIAddress != nil {
		if err := o.SCSIAddress.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ClusterSCSIAddress{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.DiskIsClusterable); err != nil {
		return err
	}
	for i1 := range o.AdapterDesc {
		i1 := i1
		if uint64(i1) >= 260 {
			break
		}
		if err := w.WriteData(o.AdapterDesc[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.AdapterDesc); uint64(i1) < 260; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PathsLength); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	return nil
}
func (o *DiskProperties) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DiskNumber); err != nil {
		return err
	}
	if o.DiskID == nil {
		o.DiskID = &ClusterDiskID{}
	}
	if err := o.DiskID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.DiskBusType); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.StackType)); err != nil {
		return err
	}
	if o.SCSIAddress == nil {
		o.SCSIAddress = &ClusterSCSIAddress{}
	}
	if err := o.SCSIAddress.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.DiskIsClusterable); err != nil {
		return err
	}
	o.AdapterDesc = make([]uint16, 260)
	for i1 := range o.AdapterDesc {
		i1 := i1
		if err := w.ReadData(&o.AdapterDesc[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.PathsLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	return nil
}

// DiskPropertiesEx structure represents DISK_PROPS_EX RPC structure.
//
// The DISK_PROPS_EX structure holds information about a single disk. This structure
// SHOULD<6> be supported and is required for the IClusterStorage3 interface.
type DiskPropertiesEx struct {
	// DiskNumber:  The zero-based device number assigned to the disk by the operating system.
	DiskNumber uint32 `idl:"name:DiskNumber" json:"disk_number"`
	// DiskId:  A valid CPREP_DISKID structure with the correct identifier for the disk.
	DiskID *ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	// DiskBusType:  The type of bus to which the disk is attached. It contains one of the
	// following values.
	//
	//	+-------------------------------------+--------------------------------------------------------------+
	//	|                                     |                                                              |
	//	|                VALUE                |                           MEANING                            |
	//	|                                     |                                                              |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusTypeUnknown 0x00000000           | The bus type is not one of those that follow.                |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusTypeScsi 0x00000001              | The bus type is SCSI.                                        |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusTypeAtapi 0x00000002             | The bus type is AT attachment packet interface (ATAPI).      |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusTypeAta 0x00000003               | The bus type is advanced technology attachment (ATA).        |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusType1394 0x00000004              | The bus type is IEEE 1394.                                   |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusTypeSsa 0x00000005               | The bus type is serial storage architecture (SSA).           |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusTypeFibre 0x00000006             | The bus type is Fibre Channel.                               |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusTypeUsb 0x00000007               | The bus type is universal serial bus (USB).                  |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusTypeRAID 0x00000008              | The bus type is redundant array of independent disks (RAID). |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusTypeiScsi 0x00000009             | The bus type is iSCSI.                                       |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusTypeSas 0x0000000A               | The bus type is Serial Attached SCSI (SAS).                  |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusTypeSata 0x0000000B              | The bus type is Serial ATA (SATA).                           |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusTypeSd 0x0000000C                | The bus type is Sd.                                          |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusTypeMmc 0x0000000D               | The bus type is Mmc.                                         |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusTypeVirtual 0x0000000E           | The bus type is Virtual.                                     |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusTypeFileBackedVirtual 0x0000000F | The bus type is File Backed Virtual.                         |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| BusTypeSpaces 0x00000010            | The bus type is Spaces.                                      |
	//	+-------------------------------------+--------------------------------------------------------------+
	DiskBusType uint32 `idl:"name:DiskBusType" json:"disk_bus_type"`
	// StackType:  The driver subtype of the device driver. It MUST be one of the valid
	// values for DiskStackType.
	StackType DiskStackType `idl:"name:StackType" json:"stack_type"`
	// ScsiAddress:  The SCSI address of the disk. It MUST be a valid CPREP_SCSI_ADDRESS.
	SCSIAddress *ClusterSCSIAddress `idl:"name:ScsiAddress" json:"scsi_address"`
	// DiskIsClusterable:  A Boolean flag that indicates whether the disk can be clustered.
	// A value of TRUE or 1 indicates that the disk can be clustered. A value of FALSE or
	// 0 indicates that the disk cannot be clustered. The value of the DiskIsClusterable
	// member can be determined in an implementation-specific way.
	DiskIsClusterable bool `idl:"name:DiskIsClusterable" json:"disk_is_clusterable"`
	// AdapterDesc:  A user-friendly description of the adapter to which the disk is connected.
	AdapterDesc []uint16 `idl:"name:AdapterDesc" json:"adapter_desc"`
	// pwszFriendlyName:  A null-terminated string containing a user-friendly description
	// of the disk. Memory for this string is allocated by the server and MUST be freed
	// by the client.
	FriendlyName string `idl:"name:pwszFriendlyName;string" json:"friendly_name"`
	// NumPaths:  The number of IO paths to the disk. A Multipath I/O (MPIO) disk has a
	// number greater than 1.
	PathsLength uint32 `idl:"name:NumPaths" json:"paths_length"`
	// Flags:  Information about the disk. It contains one or more of the following values
	// bitwise OR'd together.
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                                         |                                                                                  |
	//	|                  VALUE                  |                                     MEANING                                      |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_BOOT 0x00000001                    | The disk is the boot device.                                                     |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_SYSTEM 0x00000002                  | The disk contains the operating system.                                          |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_PAGEFILE 0x00000004                | The disk contains an operating system pagefile.                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_HIBERNATE 0x00000008               | The disk will be used to store system hibernation data.                          |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_CRASHDUMP 0x00000010               | The disk will be used to store system crash dump data.                           |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_REMOVABLE 0x00000020               | The disk is on removable media.                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_CLUSTERNOSUPP 0x00000040           | The disk is not supported by the cluster implementation. The criteria for        |
	//	|                                         | support are implementation-specific.                                             |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_BUSNOSUPP 0x00000100               | The disk is on a bus not supported by the cluster implementation. The criteria   |
	//	|                                         | for support are implementation-specific.                                         |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_SYSTEMBUS 0x00000200               | The disk is on the same bus as the disk containing the operating system.         |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_ALREADY_CLUSTERED 0x00000400       | The disk is already controlled by the cluster.                                   |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_SYTLE_MBR 0x00001000               | The disk is MBR.                                                                 |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_STYLE_GPT 0x00002000               | The disk is GPT.                                                                 |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_STYLE_RAW 0x00004000               | The disk is neither MBR nor GPT.                                                 |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_PART_BASIC 0x00008000              | The disk is configured with basic volumes.                                       |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_PART_DYNAMIC 0x00010000            | The disk is configured with dynamic volumes.                                     |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_CLUSTERED_ONLINE 0x00020000        | The disk is controlled by the cluster and is online.                             |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_UNREADABLE 0x00040000              | The disk cannot be read.                                                         |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_MPIO 0x00080000                    | The disk is controlled by MPIO.                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_CLUSTERED_OTHER 0x00100000         | The disk is controlled by cluster software other than the failover cluster       |
	//	|                                         | implementation.                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_MISSING 0x00200000                 | The disk could not be found.                                                     |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_REDUNDANT 0x00400000               | The disk is exposed to the operating system more than once through redundant     |
	//	|                                         | paths.                                                                           |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_SNAPSHOT 0x00800000                | The disk is a snapshot disk.                                                     |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_FAILING_IO 0x02000000              | The disk is unable to gather disk information.                                   |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_NO_PAGE83 0x04000000               | The disk does not have a Device Identification VPD page (see [SPC-3] section     |
	//	|                                         | 7.6.3) with PAGE CODE (see [SPC-3] table 294) set to 83h, a device ASSOCIATION   |
	//	|                                         | (see [SPC-3] table 297), and IDENTIFIER TYPE (see [SPC-3] table 298) of Type 8,  |
	//	|                                         | Type 3, or Type 2.                                                               |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_COLLISION 0x08000000               | The disk's signature collides with the signature of another disk visible to this |
	//	|                                         | server, and disk signature collision resolution is disabled.                     |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_OUTOFSPACE 0x10000000              | The disk is a thin-provisioned LUN that has no free space.                       |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_POOL_DRIVE 0x20000000              | The disk is a member of a storage pool.                                          |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_POOL_DRIVE_NOT_TESTABLE 0x40000000 | The disk is a member of a storage pool but does not meet implementation-specific |
	//	|                                         | criteria for testing.                                                            |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_POOL_CLUSTERED 0x80000000          | The disk is a member of a storage pool, and the storage pool to which it belongs |
	//	|                                         | is a cluster resource.                                                           |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
	// ExtendedFlags:  Additional information about the disk. It contains one or more of
	// the following values bitwise OR'd together.
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                                         |                                                                                  |
	//	|                  VALUE                  |                                     MEANING                                      |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_EX_SPLITPOOLCONFIG 0x00000001      | The storage pool drive is configured for both pool and non-pool data.            |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DISK_EX_POOL_NOT_CLUSTERABLE 0x00000002 | The storage pool drive is part of a pool that is not suitable for failover       |
	//	|                                         | clustering.                                                                      |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	ExtendedFlags uint32 `idl:"name:ExtendedFlags" json:"extended_flags"`
	// pwszPoolName:  A null-terminated string indicating the name of the storage pool that
	// the disk is a member of. If the disk is not a member of a storage pool, this field
	// MUST be initialized to NULL.
	PoolName string `idl:"name:pwszPoolName;string" json:"pool_name"`
	// pwszPage83Id:  A null-terminated string containing a VPD 83h identifier (see [SPC-3]
	// section 7.6.3) associated with the addressed logical unit number. The VPD 83h ASSOCIATION
	// field (see [SPC-3] table 297) has the value 00bh, and IDENTIFIER TYPE (see [SPC-3]
	// table 298) equal to SCSI name string (8h), NAA (3h), or EUI-64 based (2h).
	//
	// The order of precedence when choosing a VPD 83h identifier to return is: SCSI name
	// string type has precedence over NAA or EUI-64 based, and NAA has precedence over
	// EUI-64 based.
	Page83ID string `idl:"name:pwszPage83Id;string" json:"page83_id"`
	// pwszSerialNumber:  A null-terminated string containing the VPD page 80h (Unit Serial
	// Number see [SPC-3]section 7.6.10). This field is optional, as defined in [SPC-3]
	// (it can be all spaces). Memory for this string is allocated by the server and MUST
	// be freed by the client.
	SerialNumber string `idl:"name:pwszSerialNumber;string" json:"serial_number"`
	// guidPoolId:  The identifier of the storage pool that the disk is a member of. If
	// the disk is not a member of a storage pool, this field MUST be initialized to zero.
	PoolID *dtyp.GUID `idl:"name:guidPoolId" json:"pool_id"`
}

func (o *DiskPropertiesEx) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DiskPropertiesEx) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.DiskNumber); err != nil {
		return err
	}
	if o.DiskID != nil {
		if err := o.DiskID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ClusterDiskID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.DiskBusType); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.StackType)); err != nil {
		return err
	}
	if o.SCSIAddress != nil {
		if err := o.SCSIAddress.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ClusterSCSIAddress{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if !o.DiskIsClusterable {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	for i1 := range o.AdapterDesc {
		i1 := i1
		if uint64(i1) >= 260 {
			break
		}
		if err := w.WriteData(o.AdapterDesc[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.AdapterDesc); uint64(i1) < 260; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if o.FriendlyName != "" {
		_ptr_pwszFriendlyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.FriendlyName); err != nil {
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
	if err := w.WriteData(o.PathsLength); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.ExtendedFlags); err != nil {
		return err
	}
	if o.PoolName != "" {
		_ptr_pwszPoolName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.PoolName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.PoolName, _ptr_pwszPoolName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Page83ID != "" {
		_ptr_pwszPage83Id := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Page83ID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Page83ID, _ptr_pwszPage83Id); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.SerialNumber != "" {
		_ptr_pwszSerialNumber := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.SerialNumber); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SerialNumber, _ptr_pwszSerialNumber); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.PoolID != nil {
		if err := o.PoolID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DiskPropertiesEx) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.DiskNumber); err != nil {
		return err
	}
	if o.DiskID == nil {
		o.DiskID = &ClusterDiskID{}
	}
	if err := o.DiskID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.DiskBusType); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.StackType)); err != nil {
		return err
	}
	if o.SCSIAddress == nil {
		o.SCSIAddress = &ClusterSCSIAddress{}
	}
	if err := o.SCSIAddress.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	var _bDiskIsClusterable int32
	if err := w.ReadData(&_bDiskIsClusterable); err != nil {
		return err
	}
	o.DiskIsClusterable = _bDiskIsClusterable != 0
	o.AdapterDesc = make([]uint16, 260)
	for i1 := range o.AdapterDesc {
		i1 := i1
		if err := w.ReadData(&o.AdapterDesc[i1]); err != nil {
			return err
		}
	}
	_ptr_pwszFriendlyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.FriendlyName); err != nil {
			return err
		}
		return nil
	})
	_s_pwszFriendlyName := func(ptr interface{}) { o.FriendlyName = *ptr.(*string) }
	if err := w.ReadPointer(&o.FriendlyName, _s_pwszFriendlyName, _ptr_pwszFriendlyName); err != nil {
		return err
	}
	if err := w.ReadData(&o.PathsLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.ExtendedFlags); err != nil {
		return err
	}
	_ptr_pwszPoolName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.PoolName); err != nil {
			return err
		}
		return nil
	})
	_s_pwszPoolName := func(ptr interface{}) { o.PoolName = *ptr.(*string) }
	if err := w.ReadPointer(&o.PoolName, _s_pwszPoolName, _ptr_pwszPoolName); err != nil {
		return err
	}
	_ptr_pwszPage83Id := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Page83ID); err != nil {
			return err
		}
		return nil
	})
	_s_pwszPage83Id := func(ptr interface{}) { o.Page83ID = *ptr.(*string) }
	if err := w.ReadPointer(&o.Page83ID, _s_pwszPage83Id, _ptr_pwszPage83Id); err != nil {
		return err
	}
	_ptr_pwszSerialNumber := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.SerialNumber); err != nil {
			return err
		}
		return nil
	})
	_s_pwszSerialNumber := func(ptr interface{}) { o.SerialNumber = *ptr.(*string) }
	if err := w.ReadPointer(&o.SerialNumber, _s_pwszSerialNumber, _ptr_pwszSerialNumber); err != nil {
		return err
	}
	if o.PoolID == nil {
		o.PoolID = &dtyp.GUID{}
	}
	if err := o.PoolID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ClusterNetworkProfile type represents CLUSTER_NETWORK_PROFILE RPC enumeration.
//
// The CLUSTER_NETWORK_PROFILE enumeration defines the valid values for network adapter
// firewall profiles. When the server firewall enforces policies specified in [MS-FASP],
// the server SHOULD determine the network adapter firewall profile by querying the
// server firewall for the network adapter profile and mapping that value as specified
// below.
type ClusterNetworkProfile uint16

var (
	// ClusterNetworkProfilePublic:  See FW_PROFILE_TYPE_PUBLIC in [MS-FASP] section 2.2.2.
	ClusterNetworkProfilePublic ClusterNetworkProfile = 0
	// ClusterNetworkProfilePrivate:  See FW_PROFILE_TYPE_PRIVATE in [MS-FASP] section
	// 2.2.2.
	ClusterNetworkProfilePrivate ClusterNetworkProfile = 1
	// ClusterNetworkProfileDomainAuthenticated:  See FW_PROFILE_TYPE_DOMAIN in [MS-FASP]
	// section 2.2.2.
	ClusterNetworkProfileDomainAuthenticated ClusterNetworkProfile = 2
)

func (o ClusterNetworkProfile) String() string {
	switch o {
	case ClusterNetworkProfilePublic:
		return "ClusterNetworkProfilePublic"
	case ClusterNetworkProfilePrivate:
		return "ClusterNetworkProfilePrivate"
	case ClusterNetworkProfileDomainAuthenticated:
		return "ClusterNetworkProfileDomainAuthenticated"
	}
	return "Invalid"
}

// NodeRouteInfo structure represents NODE_ROUTE_INFO RPC structure.
//
// A client uses a NODE_ROUTE_INFO structure<7> to add routes that share the same remoteVirtualIP
// IP address field.
//
// The IP addresses in the remoteVirtualIP field and the elements of the localUnicastIPs
// array and the remoteUnicastIPs array can be either IPv4 or IPv6 and are contained
// in Unicode strings. IPv4 addresses MUST be represented in dotted decimal notation.
// IPv6 addresses MUST be represented in the form specified by [RFC1924].
type NodeRouteInfo struct {
	// remoteVirtualIP:  An IP address that is common to all routes designated by the NODE_ROUTE_INFO
	// data structure. A client uses a remoteVirtualIP as the common identifier for all
	// communication routes to a remote host.
	RemoteVirtualIP *oaut.String `idl:"name:remoteVirtualIP" json:"remote_virtual_ip"`
	// localUnicastIPs:  An array of IP addresses. A client sets the elements of localUnicastIPs
	// to the IP addresses from which the server can send network traffic to a remote host.
	LocalUnicastIPs *oaut.SafeArray `idl:"name:localUnicastIPs" json:"local_unicast_i_ps"`
	// remoteUnicastIPs:  An array of IP addresses. A client sets the elements of remoteUnicastIPs
	// to the IP address to which network traffic can be sent on a remote host.
	RemoteUnicastIPs *oaut.SafeArray `idl:"name:remoteUnicastIPs" json:"remote_unicast_i_ps"`
	// indices:  An array of unsigned integers that MUST be unique among all indices specified
	// in all NODE_ROUTE_INFO structures contained in an ADD_ROUTES_REQUEST structure.
	Indices *oaut.SafeArray `idl:"name:indices" json:"indices"`
}

func (o *NodeRouteInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *NodeRouteInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.RemoteVirtualIP != nil {
		_ptr_remoteVirtualIP := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RemoteVirtualIP != nil {
				if err := o.RemoteVirtualIP.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RemoteVirtualIP, _ptr_remoteVirtualIP); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LocalUnicastIPs != nil {
		_ptr_localUnicastIPs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.LocalUnicastIPs != nil {
				if err := o.LocalUnicastIPs.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalUnicastIPs, _ptr_localUnicastIPs); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.RemoteUnicastIPs != nil {
		_ptr_remoteUnicastIPs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RemoteUnicastIPs != nil {
				if err := o.RemoteUnicastIPs.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RemoteUnicastIPs, _ptr_remoteUnicastIPs); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Indices != nil {
		_ptr_indices := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Indices != nil {
				if err := o.Indices.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Indices, _ptr_indices); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *NodeRouteInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_remoteVirtualIP := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RemoteVirtualIP == nil {
			o.RemoteVirtualIP = &oaut.String{}
		}
		if err := o.RemoteVirtualIP.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_remoteVirtualIP := func(ptr interface{}) { o.RemoteVirtualIP = *ptr.(**oaut.String) }
	if err := w.ReadPointer(&o.RemoteVirtualIP, _s_remoteVirtualIP, _ptr_remoteVirtualIP); err != nil {
		return err
	}
	_ptr_localUnicastIPs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.LocalUnicastIPs == nil {
			o.LocalUnicastIPs = &oaut.SafeArray{}
		}
		if err := o.LocalUnicastIPs.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_localUnicastIPs := func(ptr interface{}) { o.LocalUnicastIPs = *ptr.(**oaut.SafeArray) }
	if err := w.ReadPointer(&o.LocalUnicastIPs, _s_localUnicastIPs, _ptr_localUnicastIPs); err != nil {
		return err
	}
	_ptr_remoteUnicastIPs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RemoteUnicastIPs == nil {
			o.RemoteUnicastIPs = &oaut.SafeArray{}
		}
		if err := o.RemoteUnicastIPs.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_remoteUnicastIPs := func(ptr interface{}) { o.RemoteUnicastIPs = *ptr.(**oaut.SafeArray) }
	if err := w.ReadPointer(&o.RemoteUnicastIPs, _s_remoteUnicastIPs, _ptr_remoteUnicastIPs); err != nil {
		return err
	}
	_ptr_indices := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Indices == nil {
			o.Indices = &oaut.SafeArray{}
		}
		if err := o.Indices.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_indices := func(ptr interface{}) { o.Indices = *ptr.(**oaut.SafeArray) }
	if err := w.ReadPointer(&o.Indices, _s_indices, _ptr_indices); err != nil {
		return err
	}
	return nil
}

// AddRoutesRequest structure represents ADD_ROUTES_REQUEST RPC structure.
//
// The ADD_ROUTES_REQUEST structure<8> designates a collection of communication routes
// to monitor for status and packet loss. The manifestation of a communication route
// is implementation-specific. A communication route includes network endpoints, identified
// by IP addresses, between which packets can be sent and received.
type AddRoutesRequest struct {
	// localVirtualIP:  The IP address that is common to all routes initiated from a server.
	// Typically a client uses an arbitrary localVirtualIP as the common identifier for
	// all communication routes from the server to any remote host. The IP address is represented
	// as a Unicode string and can be either IPv4 or IPv6. IPv4 addresses MUST be represented
	// in dotted decimal notation. IPv6 addresses MUST be represented in the form specified
	// by [RFC1924].
	LocalVirtualIP *oaut.String `idl:"name:localVirtualIP" json:"local_virtual_ip"`
	// nodeRouteInfos:  An array of NODE_ROUTE_INFO objects.
	NodeRouteInfos *oaut.SafeArray `idl:"name:nodeRouteInfos" json:"node_route_infos"`
}

func (o *AddRoutesRequest) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AddRoutesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.LocalVirtualIP != nil {
		_ptr_localVirtualIP := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.LocalVirtualIP != nil {
				if err := o.LocalVirtualIP.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalVirtualIP, _ptr_localVirtualIP); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.NodeRouteInfos != nil {
		_ptr_nodeRouteInfos := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.NodeRouteInfos != nil {
				if err := o.NodeRouteInfos.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.NodeRouteInfos, _ptr_nodeRouteInfos); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *AddRoutesRequest) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_localVirtualIP := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.LocalVirtualIP == nil {
			o.LocalVirtualIP = &oaut.String{}
		}
		if err := o.LocalVirtualIP.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_localVirtualIP := func(ptr interface{}) { o.LocalVirtualIP = *ptr.(**oaut.String) }
	if err := w.ReadPointer(&o.LocalVirtualIP, _s_localVirtualIP, _ptr_localVirtualIP); err != nil {
		return err
	}
	_ptr_nodeRouteInfos := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.NodeRouteInfos == nil {
			o.NodeRouteInfos = &oaut.SafeArray{}
		}
		if err := o.NodeRouteInfos.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_nodeRouteInfos := func(ptr interface{}) { o.NodeRouteInfos = *ptr.(**oaut.SafeArray) }
	if err := w.ReadPointer(&o.NodeRouteInfos, _s_nodeRouteInfos, _ptr_nodeRouteInfos); err != nil {
		return err
	}
	return nil
}

// RouteStatus type represents ROUTE_STATUS RPC enumeration.
//
// The ROUTE_STATUS enumeration<9> defines the possible states of a communication route.
type RouteStatus uint16

var (
	// DOWN:  Using implementation-specific mechanisms, the server deemed the route unsuitable
	// for communication to the remote host.
	RouteStatusDown RouteStatus = 0
	// UP:  Using implementation-specific mechanisms, the server deemed the route suitable
	// for communication to the remote host.
	RouteStatusUp RouteStatus = 1
	// UP_DOWN:  Using implementation-specific mechanisms, the server deemed the route
	// neither consistently suitable nor consistently unsuitable for communication to the
	// remote host.
	RouteStatusUpDown RouteStatus = 2
)

func (o RouteStatus) String() string {
	switch o {
	case RouteStatusDown:
		return "RouteStatusDown"
	case RouteStatusUp:
		return "RouteStatusUp"
	case RouteStatusUpDown:
		return "RouteStatusUpDown"
	}
	return "Invalid"
}

// RouteLossAndState structure represents ROUTE_LOSS_AND_STATE RPC structure.
//
// The ROUTE_LOSS_AND_STATE structure<10> contains information about a route’s packet
// loss and its status.
type RouteLossAndState struct {
	// packetLoss:  A value between 0x00000000 and 0x00000064. Designates the reliability
	// of communication on the route measured by the server using implementation-specific
	// mechanisms. A value of 0x00000000 represents most reliable, and 0x00000064 designates
	// least reliable. A server sends a collection of network packets to the remote host
	// and measures the number of packets that are successfully delivered.
	PacketLoss uint32 `idl:"name:packetLoss" json:"packet_loss"`
	// status:  The status of the communication route.
	Status RouteStatus `idl:"name:status" json:"status"`
}

func (o *RouteLossAndState) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RouteLossAndState) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.PacketLoss); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Status)); err != nil {
		return err
	}
	return nil
}
func (o *RouteLossAndState) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.PacketLoss); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Status)); err != nil {
		return err
	}
	return nil
}

// AddRoutesReply structure represents ADD_ROUTES_REPLY RPC structure.
//
// The ADD_ROUTES_REPLY structure<11> contains information about packet loss and route
// status for routes previously added by the client.
type AddRoutesReply struct {
	// indices:  An array of unsigned integers matching the indices previously designated
	// by the client in an ADD_ROUTES_REQUEST data structure.
	Indices *oaut.SafeArray `idl:"name:indices" json:"indices"`
	// replies:  An array of ROUTE_LOSS_AND_STATE (section 2.2.21) objects representing
	// the communication data collected by the server using implementation-specific mechanisms.
	Replies *oaut.SafeArray `idl:"name:replies" json:"replies"`
	// routeUnavailable:  A value of TRUE indicates that the server was not in the correct
	// state to set the remaining fields of the ROUTE_LOSS_AND_STATE data structure. In
	// this case, the indices and replies fields MUST be ignored.
	RouteUnavailable bool `idl:"name:routeUnavailable" json:"route_unavailable"`
}

func (o *AddRoutesReply) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AddRoutesReply) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.Indices != nil {
		_ptr_indices := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Indices != nil {
				if err := o.Indices.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Indices, _ptr_indices); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Replies != nil {
		_ptr_replies := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Replies != nil {
				if err := o.Replies.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Replies, _ptr_replies); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.RouteUnavailable); err != nil {
		return err
	}
	return nil
}
func (o *AddRoutesReply) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_indices := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Indices == nil {
			o.Indices = &oaut.SafeArray{}
		}
		if err := o.Indices.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_indices := func(ptr interface{}) { o.Indices = *ptr.(**oaut.SafeArray) }
	if err := w.ReadPointer(&o.Indices, _s_indices, _ptr_indices); err != nil {
		return err
	}
	_ptr_replies := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Replies == nil {
			o.Replies = &oaut.SafeArray{}
		}
		if err := o.Replies.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_replies := func(ptr interface{}) { o.Replies = *ptr.(**oaut.SafeArray) }
	if err := w.ReadPointer(&o.Replies, _s_replies, _ptr_replies); err != nil {
		return err
	}
	if err := w.ReadData(&o.RouteUnavailable); err != nil {
		return err
	}
	return nil
}

// DiskMediaType type represents DiskMediaType RPC enumeration.
type DiskMediaType uint16

var (
	DiskMediaTypeUnknown DiskMediaType = 0
	DiskMediaTypeHDD     DiskMediaType = 1
	DiskMediaTypeSSD     DiskMediaType = 2
	DiskMediaTypeSCM     DiskMediaType = 3
)

func (o DiskMediaType) String() string {
	switch o {
	case DiskMediaTypeUnknown:
		return "DiskMediaTypeUnknown"
	case DiskMediaTypeHDD:
		return "DiskMediaTypeHDD"
	case DiskMediaTypeSSD:
		return "DiskMediaTypeSSD"
	case DiskMediaTypeSCM:
		return "DiskMediaTypeSCM"
	}
	return "Invalid"
}

// ClusterLogExFlag type represents ClusterLogExFlag RPC enumeration.
type ClusterLogExFlag uint16

var (
	ClusterLogExFlagNone             ClusterLogExFlag = 0
	ClusterLogExFlagLocalTime        ClusterLogExFlag = 1
	ClusterLogExFlagSkipClusterState ClusterLogExFlag = 2
)

func (o ClusterLogExFlag) String() string {
	switch o {
	case ClusterLogExFlagNone:
		return "ClusterLogExFlagNone"
	case ClusterLogExFlagLocalTime:
		return "ClusterLogExFlagLocalTime"
	case ClusterLogExFlagSkipClusterState:
		return "ClusterLogExFlagSkipClusterState"
	}
	return "Invalid"
}

// ClusterCertType type represents CLUSTER_CERTTYPE RPC enumeration.
type ClusterCertType uint16

var (
	ClusterCertTypeSChannel    ClusterCertType = 0
	ClusterCertTypeSetSChannel ClusterCertType = 1
	ClusterCertTypePku2u       ClusterCertType = 2
	ClusterCertTypeSetPku2u    ClusterCertType = 3
)

func (o ClusterCertType) String() string {
	switch o {
	case ClusterCertTypeSChannel:
		return "ClusterCertTypeSChannel"
	case ClusterCertTypeSetSChannel:
		return "ClusterCertTypeSetSChannel"
	case ClusterCertTypePku2u:
		return "ClusterCertTypePku2u"
	case ClusterCertTypeSetPku2u:
		return "ClusterCertTypeSetPku2u"
	}
	return "Invalid"
}

// ClusterCert structure represents CLUSTER_CERT RPC structure.
//
// The CLUSTER_CERT structure contains certificate information and the cluster secret
// that is distributed by the client to all nodes in the cluster.
type ClusterCert struct {
	// CbCertData: Length of the CertData field.
	CertDataLength uint32 `idl:"name:CbCertData" json:"cert_data_length"`
	// CbKeyData: Length of the KeyData field.
	KeyDataLength uint32 `idl:"name:CbKeyData" json:"key_data_length"`
	// CertData: Exported certificate blob from the certificate store.
	CertData []byte `idl:"name:CertData" json:"cert_data"`
	// KeyData: Exported private key blob from the crypto container that matches the certificate.
	KeyData []byte `idl:"name:KeyData" json:"key_data"`
	// ClusterSecret: Cluster secret data as defined in section 3.10.1.
	ClusterSecret []uint16 `idl:"name:ClusterSecret" json:"cluster_secret"`
}

func (o *ClusterCert) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClusterCert) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.CertDataLength); err != nil {
		return err
	}
	if err := w.WriteData(o.KeyDataLength); err != nil {
		return err
	}
	for i1 := range o.CertData {
		i1 := i1
		if uint64(i1) >= 5120 {
			break
		}
		if err := w.WriteData(o.CertData[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.CertData); uint64(i1) < 5120; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.KeyData {
		i1 := i1
		if uint64(i1) >= 10240 {
			break
		}
		if err := w.WriteData(o.KeyData[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.KeyData); uint64(i1) < 10240; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.ClusterSecret {
		i1 := i1
		if uint64(i1) >= 33 {
			break
		}
		if err := w.WriteData(o.ClusterSecret[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.ClusterSecret); uint64(i1) < 33; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClusterCert) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.CertDataLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.KeyDataLength); err != nil {
		return err
	}
	o.CertData = make([]byte, 5120)
	for i1 := range o.CertData {
		i1 := i1
		if err := w.ReadData(&o.CertData[i1]); err != nil {
			return err
		}
	}
	o.KeyData = make([]byte, 10240)
	for i1 := range o.KeyData {
		i1 := i1
		if err := w.ReadData(&o.KeyData[i1]); err != nil {
			return err
		}
	}
	o.ClusterSecret = make([]uint16, 33)
	for i1 := range o.ClusterSecret {
		i1 := i1
		if err := w.ReadData(&o.ClusterSecret[i1]); err != nil {
			return err
		}
	}
	return nil
}

// ClusterLogEx structure represents IClusterLogEx RPC structure.
type ClusterLogEx dcom.InterfacePointer

func (o *ClusterLogEx) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *ClusterLogEx) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClusterLogEx) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ClusterLogEx) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ClusterLogEx) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ClusterStorage3 structure represents IClusterStorage3 RPC structure.
type ClusterStorage3 dcom.InterfacePointer

func (o *ClusterStorage3) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *ClusterStorage3) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClusterStorage3) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ClusterStorage3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ClusterStorage3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ClusterLog structure represents IClusterLog RPC structure.
type ClusterLog dcom.InterfacePointer

func (o *ClusterLog) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *ClusterLog) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClusterLog) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ClusterLog) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ClusterLog) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ClusterStorage2 structure represents IClusterStorage2 RPC structure.
type ClusterStorage2 dcom.InterfacePointer

func (o *ClusterStorage2) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *ClusterStorage2) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClusterStorage2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ClusterStorage2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ClusterStorage2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ClusterSetup structure represents IClusterSetup RPC structure.
type ClusterSetup dcom.InterfacePointer

func (o *ClusterSetup) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *ClusterSetup) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClusterSetup) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ClusterSetup) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ClusterSetup) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ClusterFirewall structure represents IClusterFirewall RPC structure.
type ClusterFirewall dcom.InterfacePointer

func (o *ClusterFirewall) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *ClusterFirewall) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClusterFirewall) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ClusterFirewall) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ClusterFirewall) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ClusterCleanup structure represents IClusterCleanup RPC structure.
type ClusterCleanup dcom.InterfacePointer

func (o *ClusterCleanup) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *ClusterCleanup) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClusterCleanup) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ClusterCleanup) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ClusterCleanup) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ClusterNetwork2 structure represents IClusterNetwork2 RPC structure.
type ClusterNetwork2 dcom.InterfacePointer

func (o *ClusterNetwork2) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *ClusterNetwork2) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClusterNetwork2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ClusterNetwork2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ClusterNetwork2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ClusterUpdate structure represents IClusterUpdate RPC structure.
type ClusterUpdate dcom.InterfacePointer

func (o *ClusterUpdate) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *ClusterUpdate) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClusterUpdate) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ClusterUpdate) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ClusterUpdate) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ClusterStorage2 class identifier c72b09db-4d53-4f41-8dcc-2d752ab56f7c
var ClusterStorage2ClassID = &dcom.ClassID{Data1: 0xc72b09db, Data2: 0x4d53, Data3: 0x4f41, Data4: []byte{0x8d, 0xcc, 0x2d, 0x75, 0x2a, 0xb5, 0x6f, 0x7c}}

// ClusterNetwork2 class identifier e1568352-586d-43e4-933f-8e6dc4de317a
var ClusterNetwork2ClassID = &dcom.ClassID{Data1: 0xe1568352, Data2: 0x586d, Data3: 0x43e4, Data4: []byte{0x93, 0x3f, 0x8e, 0x6d, 0xc4, 0xde, 0x31, 0x7a}}

// ClusterCleanup class identifier a6d3e32b-9814-4409-8de3-cfa673e6d3de
var ClusterCleanupClassID = &dcom.ClassID{Data1: 0xa6d3e32b, Data2: 0x9814, Data3: 0x4409, Data4: []byte{0x8d, 0xe3, 0xcf, 0xa6, 0x73, 0xe6, 0xd3, 0xde}}

// ClusterSetup class identifier 04d55210-b6ac-4248-9e69-2a569d1d2ab6
var ClusterSetupClassID = &dcom.ClassID{Data1: 0x04d55210, Data2: 0xb6ac, Data3: 0x4248, Data4: []byte{0x9e, 0x69, 0x2a, 0x56, 0x9d, 0x1d, 0x2a, 0xb6}}

// ClusterLog class identifier 88e7ac6d-c561-4f03-9a60-39dd768f867d
var ClusterLogClassID = &dcom.ClassID{Data1: 0x88e7ac6d, Data2: 0xc561, Data3: 0x4f03, Data4: []byte{0x9a, 0x60, 0x39, 0xdd, 0x76, 0x8f, 0x86, 0x7d}}

// ClusterFirewall class identifier 3cfee98c-fb4b-44c6-bd98-a1db14abca3f
var ClusterFirewallClassID = &dcom.ClassID{Data1: 0x3cfee98c, Data2: 0xfb4b, Data3: 0x44c6, Data4: []byte{0xbd, 0x98, 0xa1, 0xdb, 0x14, 0xab, 0xca, 0x3f}}

// ClusterUpdate class identifier 4142dd5d-3472-4370-8641-de7856431fb0
var ClusterUpdateClassID = &dcom.ClassID{Data1: 0x4142dd5d, Data2: 0x3472, Data3: 0x4370, Data4: []byte{0x86, 0x41, 0xde, 0x78, 0x56, 0x43, 0x1f, 0xb0}}
