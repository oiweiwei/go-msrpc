// The vds package implements the VDS client protocol.
//
// # Introduction
//
// The Virtual Disk Service (VDS) Remote Protocol is a set of Distributed Component
// Object Model (DCOM) interfaces for managing the configuration of disk storage on
// a computer. The Virtual Disk Service Remote Protocol deals with detailed low-level
// operating system and storage concepts.
//
// Although this specification outlines the basic concepts that you need to know, this
// specification assumes that you are familiar with these technologies. For information
// about storage, disk, and volume concepts, see [MSDN-STC] and [MSDN-PARTITIONINFO];
// for information on disk management, see [MSDN-DISKMAN].
//
// The Virtual Disk Service Remote Protocol is used to programmatically enumerate and
// configure disks, volumes, host bus adapter (HBA) ports, and iSCSI initiators on remote
// computers. This protocol supersedes the Disk Management Remote Protocol, as specified
// in [MS-DMRP].
//
// # Overview
//
// The Virtual Disk Service Remote Protocol provides a mechanism for remote configuration
// of disks, partitions, volumes, and iSCSI initiators on a server. Through the Virtual
// Disk Service Remote Protocol, a client can change the configuration of disks into
// partitions, partitions into volumes, and volumes into file systems. The protocol
// also enables clients to obtain notifications of changes to these storage objects.
//
// In the Virtual Disk Service Remote Protocol, two entities are involved: the server,
// whose storage is configured, and the client, which accesses and requests changes
// to the server storage configuration.
//
// The Virtual Disk Service Remote Protocol is expressed as a set of DCOM interfaces.
// For a server, this protocol implements support for the DCOM interface in order to
// manage storage. For a client, this protocol invokes method calls on the interface
// in order to perform disk and volume configuration tasks on the server.<1>
package vds

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
	GoPackage = "dcom/vds"
)

// MaxPath represents the MAX_PATH RPC constant
var MaxPath = 260

// MaxFSNameSize represents the MAX_FS_NAME_SIZE RPC constant
var MaxFSNameSize = 8

// VerVDSLUNInformation represents the VER_VDS_LUN_INFORMATION RPC constant
var VerVDSLUNInformation = 1

// NFPackArrive represents the VDS_NF_PACK_ARRIVE RPC constant
var NFPackArrive = 1

// NFPackDepart represents the VDS_NF_PACK_DEPART RPC constant
var NFPackDepart = 2

// NFPackModify represents the VDS_NF_PACK_MODIFY RPC constant
var NFPackModify = 3

// NFVolumeArrive represents the VDS_NF_VOLUME_ARRIVE RPC constant
var NFVolumeArrive = 4

// NFVolumeDepart represents the VDS_NF_VOLUME_DEPART RPC constant
var NFVolumeDepart = 5

// NFVolumeModify represents the VDS_NF_VOLUME_MODIFY RPC constant
var NFVolumeModify = 6

// NFVolumeRebuildingProgress represents the VDS_NF_VOLUME_REBUILDING_PROGRESS RPC constant
var NFVolumeRebuildingProgress = 7

// NFDiskArrive represents the VDS_NF_DISK_ARRIVE RPC constant
var NFDiskArrive = 8

// NFDiskDepart represents the VDS_NF_DISK_DEPART RPC constant
var NFDiskDepart = 9

// NFDiskModify represents the VDS_NF_DISK_MODIFY RPC constant
var NFDiskModify = 10

// NFPartitionArrive represents the VDS_NF_PARTITION_ARRIVE RPC constant
var NFPartitionArrive = 11

// NFPartitionDepart represents the VDS_NF_PARTITION_DEPART RPC constant
var NFPartitionDepart = 12

// NFPartitionModify represents the VDS_NF_PARTITION_MODIFY RPC constant
var NFPartitionModify = 13

// NFDriveLetterFree represents the VDS_NF_DRIVE_LETTER_FREE RPC constant
var NFDriveLetterFree = 201

// NFDriveLetterAssign represents the VDS_NF_DRIVE_LETTER_ASSIGN RPC constant
var NFDriveLetterAssign = 202

// NFFileSystemModify represents the VDS_NF_FILE_SYSTEM_MODIFY RPC constant
var NFFileSystemModify = 203

// NFFileSystemFormatProgress represents the VDS_NF_FILE_SYSTEM_FORMAT_PROGRESS RPC constant
var NFFileSystemFormatProgress = 204

// NFMountPointsChange represents the VDS_NF_MOUNT_POINTS_CHANGE RPC constant
var NFMountPointsChange = 205

// NFServiceOutOfSync represents the VDS_NF_SERVICE_OUT_OF_SYNC RPC constant
var NFServiceOutOfSync = 301

// ObjectID structure represents VDS_OBJECT_ID RPC structure.
type ObjectID dtyp.GUID

func (o *ObjectID) GUID() *dtyp.GUID { return (*dtyp.GUID)(o) }

func (o *ObjectID) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ObjectID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Health type represents VDS_HEALTH RPC enumeration.
//
// The VDS_HEALTH enumeration defines the possible health states of the storage objects
// in the Virtual Disk Service Remote Protocol. The storage objects are packs, volumes,
// volume plexes and disks.
type Health uint16

var (
	// VDS_H_UNKNOWN:  The health of the object cannot be determined.
	HealthUnknown Health = 0
	// VDS_H_HEALTHY:  The object indicates online status. If the object is a disk, the
	// disk is not missing, dynamic disk log and configuration files are synchronized, and
	// the disk is free of input/output errors. If the object is a LUN or volume, all plexes
	// (mirrored, simple, spanned, and striped) and columns (RAID-5) are active. For a volume,
	// VDS_H_HEALTHY implies no disks containing volume extents have IO errors. For a LUN,
	// VDS_H_HEALTHY implies no drives containing LUN extents have IO errors.
	HealthHealthy Health = 1
	// VDS_H_REBUILDING:  The volume is resynchronizing all plexes, or a striped with parity
	// (RAID-5) plex is regenerating the parity.
	HealthRebuilding Health = 2
	// VDS_H_STALE:  The object configuration is stale.
	HealthStale Health = 3
	// VDS_H_FAILING:  The object is failing but still working. For example, a volume with
	// failing health might produce occasional input/output errors from which it can still
	// recover.
	HealthFailing Health = 4
	// VDS_H_FAILING_REDUNDANCY:  One or more plexes have errors, but the object is working
	// and all plexes are online.
	HealthFailingRedundancy Health = 5
	// VDS_H_FAILED_REDUNDANCY:  One or more plexes have failed, but at least one plex
	// is working.
	HealthFailedRedundancy Health = 6
	// VDS_H_FAILED_REDUNDANCY_FAILING:  The last working plex is failing.
	HealthFailedRedundancyFailing Health = 7
	// VDS_H_FAILED:  The object has failed. Any object with a failed health status also
	// has a failed object status.
	HealthFailed Health = 8
)

func (o Health) String() string {
	switch o {
	case HealthUnknown:
		return "HealthUnknown"
	case HealthHealthy:
		return "HealthHealthy"
	case HealthRebuilding:
		return "HealthRebuilding"
	case HealthStale:
		return "HealthStale"
	case HealthFailing:
		return "HealthFailing"
	case HealthFailingRedundancy:
		return "HealthFailingRedundancy"
	case HealthFailedRedundancy:
		return "HealthFailedRedundancy"
	case HealthFailedRedundancyFailing:
		return "HealthFailedRedundancyFailing"
	case HealthFailed:
		return "HealthFailed"
	}
	return "Invalid"
}

// NotificationTargetType type represents VDS_NOTIFICATION_TARGET_TYPE RPC enumeration.
//
// The VDS_NOTIFICATION_TARGET_TYPE enumeration defines the set of valid target types
// (subjects) of a Virtual Disk Service Remote Protocol notification.
type NotificationTargetType uint16

var (
	// VDS_NTT_UNKNOWN:  Notification is of an unknown type.
	NotificationTargetTypeUnknown NotificationTargetType = 0
	// VDS_NTT_PACK:  Notification refers to a pack.
	NotificationTargetTypePack NotificationTargetType = 10
	// VDS_NTT_VOLUME:  Notification refers to a volume.
	NotificationTargetTypeVolume NotificationTargetType = 11
	// VDS_NTT_DISK:  Notification refers to a disk.
	NotificationTargetTypeDisk NotificationTargetType = 13
	// VDS_NTT_PARTITION:  Notification refers to a partition.
	NotificationTargetTypePartition NotificationTargetType = 60
	// VDS_NTT_DRIVE_LETTER:  Notification refers to a drive letter.
	NotificationTargetTypeDriveLetter NotificationTargetType = 61
	// VDS_NTT_FILE_SYSTEM:  Notification refers to a file system.
	NotificationTargetTypeFileSystem NotificationTargetType = 62
	// VDS_NTT_MOUNT_POINT:  Notification refers to a mount point.
	NotificationTargetTypeMountPoint NotificationTargetType = 63
	// VDS_NTT_SERVICE:  Notification refers to the Virtual Disk Service.<6>
	NotificationTargetTypeService NotificationTargetType = 200
)

func (o NotificationTargetType) String() string {
	switch o {
	case NotificationTargetTypeUnknown:
		return "NotificationTargetTypeUnknown"
	case NotificationTargetTypePack:
		return "NotificationTargetTypePack"
	case NotificationTargetTypeVolume:
		return "NotificationTargetTypeVolume"
	case NotificationTargetTypeDisk:
		return "NotificationTargetTypeDisk"
	case NotificationTargetTypePartition:
		return "NotificationTargetTypePartition"
	case NotificationTargetTypeDriveLetter:
		return "NotificationTargetTypeDriveLetter"
	case NotificationTargetTypeFileSystem:
		return "NotificationTargetTypeFileSystem"
	case NotificationTargetTypeMountPoint:
		return "NotificationTargetTypeMountPoint"
	case NotificationTargetTypeService:
		return "NotificationTargetTypeService"
	}
	return "Invalid"
}

// AsyncOutputType type represents VDS_ASYNC_OUTPUT_TYPE RPC enumeration.
//
// The VDS_ASYNC_OUTPUT_TYPE enumeration defines the types of operation information
// that the VDS_ASYNC_OUTPUT structure returns.
type AsyncOutputType uint16

var (
	// VDS_ASYNCOUT_UNKNOWN:  Information is about an unknown type of operation.
	AsyncOutputTypeUnknown AsyncOutputType = 0
	// VDS_ASYNCOUT_CREATEVOLUME:  Information is about creating a volume.
	AsyncOutputTypeCreateVolume AsyncOutputType = 1
	// VDS_ASYNCOUT_EXTENDVOLUME:  Information is about extending the size of a volume.
	AsyncOutputTypeExtendVolume AsyncOutputType = 2
	// VDS_ASYNCOUT_SHRINKVOLUME:  Information is about shrinking the size of a volume.
	AsyncOutputTypeShrinkVolume AsyncOutputType = 3
	// VDS_ASYNCOUT_ADDVOLUMEPLEX:  Information is about adding a volume plex.
	AsyncOutputTypeAddVolumePlex AsyncOutputType = 4
	// VDS_ASYNCOUT_BREAKVOLUMEPLEX:  Information is about breaking a volume plex.
	AsyncOutputTypeBreakVolumePlex AsyncOutputType = 5
	// VDS_ASYNCOUT_REMOVEVOLUMEPLEX:  Information is about removing a volume plex.
	AsyncOutputTypeRemoveVolumePlex AsyncOutputType = 6
	// VDS_ASYNCOUT_REPAIRVOLUMEPLEX:  Information is about repairing a volume plex.
	AsyncOutputTypeRepairVolumePlex AsyncOutputType = 7
	// VDS_ASYNCOUT_RECOVERPACK:  Information is about recovering a pack.
	AsyncOutputTypeRecoverPack AsyncOutputType = 8
	// VDS_ASYNCOUT_REPLACEDISK:  Information is about replacing a disk.
	AsyncOutputTypeReplaceDisk AsyncOutputType = 9
	// VDS_ASYNCOUT_CREATEPARTITION:  Information is about creating a partition.
	AsyncOutputTypeCreatePartition AsyncOutputType = 10
	// VDS_ASYNCOUT_CLEAN:  Information is about cleaning a disk.
	AsyncOutputTypeClean AsyncOutputType = 11
	// VDS_ASYNCOUT_CREATELUN:  Information is about creating a LUN.
	AsyncOutputTypeCreateLUN AsyncOutputType = 50
	// VDS_ASYNCOUT_FORMAT:  Information is about formatting a file system.
	AsyncOutputTypeFormat AsyncOutputType = 101
	// VDS_ASYNCOUT_CREATE_VDISK:  Information is about creating a virtual disk.
	AsyncOutputTypeCreateVDisk AsyncOutputType = 200
	// VDS_ASYNCOUT_SURFACE_VDISK:  Information is about attaching a virtual disk.
	AsyncOutputTypeSurfaceVDisk AsyncOutputType = 201
	// VDS_ASYNCOUT_COMPACT_VDISK:  Information is about compacting a virtual disk.
	AsyncOutputTypeCompactVDisk AsyncOutputType = 202
	// VDS_ASYNCOUT_MERGE_VDISK:  Information is about merging a virtual disk.
	AsyncOutputTypeMergeVDisk AsyncOutputType = 203
	// VDS_ASYNCOUT_EXPAND_VDISK:  Information is about expanding a virtual disk.
	AsyncOutputTypeExpandVDisk AsyncOutputType = 204
)

func (o AsyncOutputType) String() string {
	switch o {
	case AsyncOutputTypeUnknown:
		return "AsyncOutputTypeUnknown"
	case AsyncOutputTypeCreateVolume:
		return "AsyncOutputTypeCreateVolume"
	case AsyncOutputTypeExtendVolume:
		return "AsyncOutputTypeExtendVolume"
	case AsyncOutputTypeShrinkVolume:
		return "AsyncOutputTypeShrinkVolume"
	case AsyncOutputTypeAddVolumePlex:
		return "AsyncOutputTypeAddVolumePlex"
	case AsyncOutputTypeBreakVolumePlex:
		return "AsyncOutputTypeBreakVolumePlex"
	case AsyncOutputTypeRemoveVolumePlex:
		return "AsyncOutputTypeRemoveVolumePlex"
	case AsyncOutputTypeRepairVolumePlex:
		return "AsyncOutputTypeRepairVolumePlex"
	case AsyncOutputTypeRecoverPack:
		return "AsyncOutputTypeRecoverPack"
	case AsyncOutputTypeReplaceDisk:
		return "AsyncOutputTypeReplaceDisk"
	case AsyncOutputTypeCreatePartition:
		return "AsyncOutputTypeCreatePartition"
	case AsyncOutputTypeClean:
		return "AsyncOutputTypeClean"
	case AsyncOutputTypeCreateLUN:
		return "AsyncOutputTypeCreateLUN"
	case AsyncOutputTypeFormat:
		return "AsyncOutputTypeFormat"
	case AsyncOutputTypeCreateVDisk:
		return "AsyncOutputTypeCreateVDisk"
	case AsyncOutputTypeSurfaceVDisk:
		return "AsyncOutputTypeSurfaceVDisk"
	case AsyncOutputTypeCompactVDisk:
		return "AsyncOutputTypeCompactVDisk"
	case AsyncOutputTypeMergeVDisk:
		return "AsyncOutputTypeMergeVDisk"
	case AsyncOutputTypeExpandVDisk:
		return "AsyncOutputTypeExpandVDisk"
	}
	return "Invalid"
}

// StorageBusType type represents VDS_STORAGE_BUS_TYPE RPC enumeration.
//
// The VDS_STORAGE_BUS_TYPE enumeration defines the type of bus on which a disk resides.
type StorageBusType uint16

var (
	// VDSBusTypeUnknown:  Bus type is unknown.
	StorageBusTypeUnknown StorageBusType = 0
	// VDSBusTypeScsi:  Disk resides on a SCSI bus.
	StorageBusTypeSCSI StorageBusType = 1
	// VDSBusTypeAtapi:  Disk resides on an AT Attachment Packet Interface (ATAPI) bus.
	// For more information on this bus type, see [ANSI/INCITS-397-2005].
	StorageBusTypeATAPI StorageBusType = 2
	// VDSBusTypeAta:  Disk resides on an AT Attached (ATA) bus. For more information on
	// this bus type, see [ANSI/INCITS-451-2008].
	StorageBusTypeATA StorageBusType = 3
	// VDSBusType1394:  Disk resides on an IEEE 1394 bus. For more information, see [IEEE1394-2008].
	StorageBusTypeType1394 StorageBusType = 4
	// VDSBusTypeSsa:  Disk resides on a serial storage architecture (SSA) bus. For more
	// information on this bus type, see [IEEE-SSA].
	StorageBusTypeSSA StorageBusType = 5
	// VDSBusTypeFibre:  Disk resides on a fiber channel bus.
	StorageBusTypeFibre StorageBusType = 6
	// VDSBusTypeUsb:  Disk resides on a universal serial bus (USB).
	StorageBusTypeUSB StorageBusType = 7
	// VDSBusTypeRAID:  Disk resides on a RAID bus.
	StorageBusTypeRAID StorageBusType = 8
	// VDSBusTypeiScsi:  Disk resides on an iSCSI bus.
	StorageBusTypeISCSI StorageBusType = 9
	// VDSBusTypeSas:  Disk resides on a Serial Attached SCSI (SAS) bus. For more information
	// on this bus type, see [ANSI/INCITS-457-2010].
	StorageBusTypeSAS StorageBusType = 10
	// VDSBusTypeSata:  Disk resides on a Serial ATA (SATA) bus. For more information on
	// this bus type, see [SATA-3.0].
	StorageBusTypeSATA StorageBusType = 11
	// VDSBusTypeSd:  Disk resides on a secure digital (SD) bus.
	StorageBusTypeSD StorageBusType = 12
	// VDSBusTypeMmc:  Indicates a multimedia card (MMC) bus type. For information on multimedia
	// cards, which are a flash memory card standard, see [JEDEC-MO227-A].
	StorageBusTypeMMC StorageBusType = 13
	// VDSBusTypeMax:  Maximum bus type value. Note that this value does not identify a
	// particular bus type; rather, it serves as an end value of the enumeration.<7>
	StorageBusTypeMax StorageBusType = 14
	// VDSBusTypeVirtual:  The disk SHOULD<8> reside on a virtual bus
	StorageBusTypeVirtual StorageBusType = 14
	// VDSBusTypeFileBackedVirtual:  Disk is backed by a file.
	StorageBusTypeFileBackedVirtual StorageBusType = 15
	// VDSBusTypeSpaces:  The disk SHOULD<9> be backed by Storage Spaces.
	StorageBusTypeSpaces StorageBusType = 16
	// VDSBusTypeMaxReserved:  Maximum reserved bus type value. Bus type values below this
	// range are reserved.
	StorageBusTypeMaxReserved StorageBusType = 127
)

func (o StorageBusType) String() string {
	switch o {
	case StorageBusTypeUnknown:
		return "StorageBusTypeUnknown"
	case StorageBusTypeSCSI:
		return "StorageBusTypeSCSI"
	case StorageBusTypeATAPI:
		return "StorageBusTypeATAPI"
	case StorageBusTypeATA:
		return "StorageBusTypeATA"
	case StorageBusTypeType1394:
		return "StorageBusTypeType1394"
	case StorageBusTypeSSA:
		return "StorageBusTypeSSA"
	case StorageBusTypeFibre:
		return "StorageBusTypeFibre"
	case StorageBusTypeUSB:
		return "StorageBusTypeUSB"
	case StorageBusTypeRAID:
		return "StorageBusTypeRAID"
	case StorageBusTypeISCSI:
		return "StorageBusTypeISCSI"
	case StorageBusTypeSAS:
		return "StorageBusTypeSAS"
	case StorageBusTypeSATA:
		return "StorageBusTypeSATA"
	case StorageBusTypeSD:
		return "StorageBusTypeSD"
	case StorageBusTypeMMC:
		return "StorageBusTypeMMC"
	case StorageBusTypeMax:
		return "StorageBusTypeMax"
	case StorageBusTypeVirtual:
		return "StorageBusTypeVirtual"
	case StorageBusTypeFileBackedVirtual:
		return "StorageBusTypeFileBackedVirtual"
	case StorageBusTypeSpaces:
		return "StorageBusTypeSpaces"
	case StorageBusTypeMaxReserved:
		return "StorageBusTypeMaxReserved"
	}
	return "Invalid"
}

// StorageIDCodeSet type represents VDS_STORAGE_IDENTIFIER_CODE_SET RPC enumeration.
//
// The VDS_STORAGE_IDENTIFIER_CODE_SET enumeration defines the code set that is used
// by the storage device identifier, as specified in [SPC-3]
type StorageIDCodeSet uint16

var (
	// VDSStorageIdCodeSetReserved:  This value is reserved by the SPC-3 standard and is
	// not used.
	StorageIDCodeSetReserved StorageIDCodeSet = 0
	// VDSStorageIdCodeSetBinary:  The identifier contains binary values.
	StorageIDCodeSetBinary StorageIDCodeSet = 1
	// VDSStorageIdCodeSetAscii:  The identifier contains ASCII values.
	StorageIDCodeSetASCII StorageIDCodeSet = 2
	// VDSStorageIdCodeSetUtf8:  The identifier contains UTF-8 values.
	StorageIDCodeSetUTF8 StorageIDCodeSet = 3
)

func (o StorageIDCodeSet) String() string {
	switch o {
	case StorageIDCodeSetReserved:
		return "StorageIDCodeSetReserved"
	case StorageIDCodeSetBinary:
		return "StorageIDCodeSetBinary"
	case StorageIDCodeSetASCII:
		return "StorageIDCodeSetASCII"
	case StorageIDCodeSetUTF8:
		return "StorageIDCodeSetUTF8"
	}
	return "Invalid"
}

// StorageIDType type represents VDS_STORAGE_IDENTIFIER_TYPE RPC enumeration.
//
// The VDS_STORAGE_IDENTIFIER_TYPE enumeration defines the types of storage device identifiers,
// as specified in [SPC-3].
type StorageIDType uint16

var (
	// VDSStorageIdTypeVendorSpecific:  Storage identifier is vendor-specific.
	StorageIDTypeVendorSpecific StorageIDType = 0
	// VDSStorageIdTypeVendorId:  Storage identifier is a vendor identifier.
	StorageIDTypeVendorID StorageIDType = 1
	// VDSStorageIdTypeEUI64:  Storage identifier is a 64-bit extended unique identifier
	// (EUI-64).
	StorageIDTypeEUI64 StorageIDType = 2
	// VDSStorageIdTypeFCPHName:  Storage identifier is a Fibre Channel Physical and Signaling
	// Interface (FC-PH) identifier.
	StorageIDTypeFCPHName StorageIDType = 3
	// VDSStorageIdTypePortRelative:  Storage identifier is a relative target port identifier.
	StorageIDTypePortRelative StorageIDType = 4
	// VDSStorageIdTypeTargetPortGroup:  Storage identifier is a target port group number.
	StorageIDTypeTargetPortGroup StorageIDType = 5
	// VDSStorageIdTypeLogicalUnitGroup:  Storage identifier is a logical unit group number.
	StorageIDTypeLogicalUnitGroup StorageIDType = 6
	// VDSStorageIdTypeMD5LogicalUnitIdentifier:  Storage identifier is an MD5 logical
	// unit number (LUN).
	StorageIDTypeMD5LogicalUnitID StorageIDType = 7
	// VDSStorageIdTypeScsiNameString:  Storage identifier is an SCSI name string identifier.
	StorageIDTypeSCSINameString StorageIDType = 8
)

func (o StorageIDType) String() string {
	switch o {
	case StorageIDTypeVendorSpecific:
		return "StorageIDTypeVendorSpecific"
	case StorageIDTypeVendorID:
		return "StorageIDTypeVendorID"
	case StorageIDTypeEUI64:
		return "StorageIDTypeEUI64"
	case StorageIDTypeFCPHName:
		return "StorageIDTypeFCPHName"
	case StorageIDTypePortRelative:
		return "StorageIDTypePortRelative"
	case StorageIDTypeTargetPortGroup:
		return "StorageIDTypeTargetPortGroup"
	case StorageIDTypeLogicalUnitGroup:
		return "StorageIDTypeLogicalUnitGroup"
	case StorageIDTypeMD5LogicalUnitID:
		return "StorageIDTypeMD5LogicalUnitID"
	case StorageIDTypeSCSINameString:
		return "StorageIDTypeSCSINameString"
	}
	return "Invalid"
}

// InterconnectAddressType type represents VDS_INTERCONNECT_ADDRESS_TYPE RPC enumeration.
//
// The VDS_INTERCONNECT_ADDRESS_TYPE enumeration defines the set of valid address types
// of a physical interconnect.
type InterconnectAddressType uint16

var (
	// VDS_IA_UNKNOWN:  This value is reserved.
	InterconnectAddressTypeUnknown InterconnectAddressType = 0
	// VDS_IA_FCFS:  Address type is first come, first served.
	InterconnectAddressTypeFCFS InterconnectAddressType = 1
	// VDS_IA_FCPH:  Address type is FC-PH. For more information, see [ANSI-289-1996].
	InterconnectAddressTypeFCPH InterconnectAddressType = 2
	// VDS_IA_FCPH3:  Address type is FC-PH-3. For more information, see [ANSI-289-1996].
	InterconnectAddressTypeFCPH3 InterconnectAddressType = 3
	// VDS_IA_MAC:  Address type is media access control (MAC).
	InterconnectAddressTypeMAC InterconnectAddressType = 4
	// VDS_IA_SCSI:  Address type is SCSI.
	InterconnectAddressTypeSCSI InterconnectAddressType = 5
)

func (o InterconnectAddressType) String() string {
	switch o {
	case InterconnectAddressTypeUnknown:
		return "InterconnectAddressTypeUnknown"
	case InterconnectAddressTypeFCFS:
		return "InterconnectAddressTypeFCFS"
	case InterconnectAddressTypeFCPH:
		return "InterconnectAddressTypeFCPH"
	case InterconnectAddressTypeFCPH3:
		return "InterconnectAddressTypeFCPH3"
	case InterconnectAddressTypeMAC:
		return "InterconnectAddressTypeMAC"
	case InterconnectAddressTypeSCSI:
		return "InterconnectAddressTypeSCSI"
	}
	return "Invalid"
}

// FileSystemType type represents VDS_FILE_SYSTEM_TYPE RPC enumeration.
//
// The VDS_FILE_SYSTEM_TYPE enumeration defines the set of valid types for a file system.
type FileSystemType uint16

var (
	// VDS_FST_UNKNOWN:  The file system is unknown.
	FileSystemTypeUnknown FileSystemType = 0
	// VDS_FST_RAW:  The file system is raw.
	FileSystemTypeRaw FileSystemType = 1
	// VDS_FST_FAT:  The file system is a FAT file system.
	FileSystemTypeFAT FileSystemType = 2
	// VDS_FST_FAT32:  The file system is FAT32.
	FileSystemTypeFAT32 FileSystemType = 3
	// VDS_FST_NTFS:  The file system is the NTFS file system.
	FileSystemTypeNTFS FileSystemType = 4
	// VDS_FST_CDFS:  The file system is the Compact Disc File System (CDFS).
	FileSystemTypeCDFS FileSystemType = 5
	// VDS_FST_UDF:  The file system is Universal Disk Format (UDF).
	FileSystemTypeUDF FileSystemType = 6
	// VDS_FST_EXFAT:  The file system is Extended File Allocation Table (ExFAT). For more
	// information, see [MSDN-EFFS].
	FileSystemTypeExFAT FileSystemType = 7
	// VDS_FST_CSVFS<10>:  The file system is Cluster Shared Volume File System (CSVFS).
	FileSystemTypeCSVFS FileSystemType = 8
	// VDS_FST_REFS<11>:  The file system is Resilient File System (ReFS).
	FileSystemTypeReFS FileSystemType = 9
)

func (o FileSystemType) String() string {
	switch o {
	case FileSystemTypeUnknown:
		return "FileSystemTypeUnknown"
	case FileSystemTypeRaw:
		return "FileSystemTypeRaw"
	case FileSystemTypeFAT:
		return "FileSystemTypeFAT"
	case FileSystemTypeFAT32:
		return "FileSystemTypeFAT32"
	case FileSystemTypeNTFS:
		return "FileSystemTypeNTFS"
	case FileSystemTypeCDFS:
		return "FileSystemTypeCDFS"
	case FileSystemTypeUDF:
		return "FileSystemTypeUDF"
	case FileSystemTypeExFAT:
		return "FileSystemTypeExFAT"
	case FileSystemTypeCSVFS:
		return "FileSystemTypeCSVFS"
	case FileSystemTypeReFS:
		return "FileSystemTypeReFS"
	}
	return "Invalid"
}

// FileSystemFlag type represents VDS_FILE_SYSTEM_FLAG RPC enumeration.
//
// The VDS_FILE_SYSTEM_FLAG enumeration defines the set of valid flags for a file system
// format type.
//
// If more than one flag is specified, the file system type supports all the file system
// allocation sizes that are specified. However, a specific file system on a volume
// does not have multiple allocation sizes at the same time.
type FileSystemFlag uint32

var (
	// VDS_FSF_SUPPORT_FORMAT:  If set, the file system format type supports format.
	FileSystemFlagSupportFormat FileSystemFlag = 1
	// VDS_FSF_SUPPORT_QUICK_FORMAT:  If set, the file system format type supports quick
	// format.
	FileSystemFlagSupportQuickFormat FileSystemFlag = 2
	// VDS_FSF_SUPPORT_COMPRESS:  If set, the file system format type supports file compression.
	FileSystemFlagSupportCompress FileSystemFlag = 4
	// VDS_FSF_SUPPORT_SPECIFY_LABEL:  If set, the file system format type supports file
	// system labels.
	FileSystemFlagSupportSpecifyLabel FileSystemFlag = 8
	// VDS_FSF_SUPPORT_MOUNT_POINT:  If set, the file system format type supports mount
	// points.
	FileSystemFlagSupportMountPoint FileSystemFlag = 16
	// VDS_FSF_SUPPORT_REMOVABLE_MEDIA:  If set, the file system format type supports removable
	// media.
	FileSystemFlagSupportRemovableMedia FileSystemFlag = 32
	// VDS_FSF_SUPPORT_EXTEND:  If set, the file system format type supports extending
	// volumes.
	FileSystemFlagSupportExtend FileSystemFlag = 64
	// VDS_FSF_ALLOCATION_UNIT_512:  If set, the file system format supports allocation
	// units of 512 bytes.
	FileSystemFlagAllocationUnit512 FileSystemFlag = 65536
	// VDS_FSF_ALLOCATION_UNIT_1K:  If set, the file system format type supports allocation
	// units of 1 kilobyte.
	FileSystemFlagAllocationUnit1k FileSystemFlag = 131072
	// VDS_FSF_ALLOCATION_UNIT_2K:  If set, the file system format type supports allocation
	// units of 2 kilobytes.
	FileSystemFlagAllocationUnit2k FileSystemFlag = 262144
	// VDS_FSF_ALLOCATION_UNIT_4K:  If set, the file system format type supports allocation
	// units of 4 kilobytes.
	FileSystemFlagAllocationUnit4k FileSystemFlag = 524288
	// VDS_FSF_ALLOCATION_UNIT_8K:  If set, the file system format type supports allocation
	// units of 8 kilobytes.
	FileSystemFlagAllocationUnit8k FileSystemFlag = 1048576
	// VDS_FSF_ALLOCATION_UNIT_16K:  If set, the file system format type supports allocation
	// units of 16 kilobytes.
	FileSystemFlagAllocationUnit16k FileSystemFlag = 2097152
	// VDS_FSF_ALLOCATION_UNIT_32K:  If set, the file system format type supports allocation
	// units of 32 kilobytes.
	FileSystemFlagAllocationUnit32k FileSystemFlag = 4194304
	// VDS_FSF_ALLOCATION_UNIT_64K:  If set, the file system format type supports allocation
	// units of 64 kilobytes.
	FileSystemFlagAllocationUnit64k FileSystemFlag = 8388608
	// VDS_FSF_ALLOCATION_UNIT_128K:  If set, the file system format type supports allocation
	// units of 128 kilobytes.
	FileSystemFlagAllocationUnit128k FileSystemFlag = 16777216
	// VDS_FSF_ALLOCATION_UNIT_256K:  If set, the file system format type supports allocation
	// units of 256 kilobytes.
	FileSystemFlagAllocationUnit256k FileSystemFlag = 33554432
)

func (o FileSystemFlag) String() string {
	switch o {
	case FileSystemFlagSupportFormat:
		return "FileSystemFlagSupportFormat"
	case FileSystemFlagSupportQuickFormat:
		return "FileSystemFlagSupportQuickFormat"
	case FileSystemFlagSupportCompress:
		return "FileSystemFlagSupportCompress"
	case FileSystemFlagSupportSpecifyLabel:
		return "FileSystemFlagSupportSpecifyLabel"
	case FileSystemFlagSupportMountPoint:
		return "FileSystemFlagSupportMountPoint"
	case FileSystemFlagSupportRemovableMedia:
		return "FileSystemFlagSupportRemovableMedia"
	case FileSystemFlagSupportExtend:
		return "FileSystemFlagSupportExtend"
	case FileSystemFlagAllocationUnit512:
		return "FileSystemFlagAllocationUnit512"
	case FileSystemFlagAllocationUnit1k:
		return "FileSystemFlagAllocationUnit1k"
	case FileSystemFlagAllocationUnit2k:
		return "FileSystemFlagAllocationUnit2k"
	case FileSystemFlagAllocationUnit4k:
		return "FileSystemFlagAllocationUnit4k"
	case FileSystemFlagAllocationUnit8k:
		return "FileSystemFlagAllocationUnit8k"
	case FileSystemFlagAllocationUnit16k:
		return "FileSystemFlagAllocationUnit16k"
	case FileSystemFlagAllocationUnit32k:
		return "FileSystemFlagAllocationUnit32k"
	case FileSystemFlagAllocationUnit64k:
		return "FileSystemFlagAllocationUnit64k"
	case FileSystemFlagAllocationUnit128k:
		return "FileSystemFlagAllocationUnit128k"
	case FileSystemFlagAllocationUnit256k:
		return "FileSystemFlagAllocationUnit256k"
	}
	return "Invalid"
}

// FileSystemPropertyFlag type represents VDS_FILE_SYSTEM_PROP_FLAG RPC enumeration.
//
// The VDS_FILE_SYSTEM_PROP_FLAG enumeration defines the set of fields for a file system.
// A value that accepts these flags SHOULD have the following flag set.
type FileSystemPropertyFlag uint16

var (
	// VDS_FPF_COMPRESSED:  If set, the file system supports file compression.
	FileSystemPropertyFlagCompressed FileSystemPropertyFlag = 1
)

func (o FileSystemPropertyFlag) String() string {
	switch o {
	case FileSystemPropertyFlagCompressed:
		return "FileSystemPropertyFlagCompressed"
	}
	return "Invalid"
}

// FileSystemFormatSupportFlag type represents VDS_FILE_SYSTEM_FORMAT_SUPPORT_FLAG RPC enumeration.
//
// The VDS_FILE_SYSTEM_FORMAT_SUPPORT_FLAG enumeration defines the properties of file
// systems that are supported for formatting volumes.<12>
type FileSystemFormatSupportFlag uint16

var (
	// VDS_FSS_DEFAULT:  The file system is the default file system for formatting the
	// volume.
	FileSystemFormatSupportFlagDefault FileSystemFormatSupportFlag = 1
	// VDS_FSS_PREVIOUS_REVISION:  The revision of the file system is not the latest revision
	// that is supported for formatting the volume.
	FileSystemFormatSupportFlagPreviousRevision FileSystemFormatSupportFlag = 2
	// VDS_FSS_RECOMMENDED:  The file system is the recommended file system for formatting
	// the volume.
	FileSystemFormatSupportFlagRecommended FileSystemFormatSupportFlag = 4
)

func (o FileSystemFormatSupportFlag) String() string {
	switch o {
	case FileSystemFormatSupportFlagDefault:
		return "FileSystemFormatSupportFlagDefault"
	case FileSystemFormatSupportFlagPreviousRevision:
		return "FileSystemFormatSupportFlagPreviousRevision"
	case FileSystemFormatSupportFlagRecommended:
		return "FileSystemFormatSupportFlagRecommended"
	}
	return "Invalid"
}

// DiskExtentType type represents VDS_DISK_EXTENT_TYPE RPC enumeration.
//
// The VDS_DISK_EXTENT_TYPE enumeration defines the set of valid types for a disk extent.
type DiskExtentType uint16

var (
	// VDS_DET_UNKNOWN:  The extent belongs to an unknown partition type.
	DiskExtentTypeUnknown DiskExtentType = 0
	// VDS_DET_FREE:  The extent belongs to an area of free space.
	DiskExtentTypeFree DiskExtentType = 1
	// VDS_DET_DATA:  The extent belongs to a volume.
	DiskExtentTypeData DiskExtentType = 2
	// VDS_DET_OEM:  The extent belongs to an OEM partition.
	DiskExtentTypeOEM DiskExtentType = 3
	// VDS_DET_ESP:  The extent belongs to an Extensible Firmware Interface (EFI) system
	// partition.
	DiskExtentTypeESP DiskExtentType = 4
	// VDS_DET_MSR:  The extent belongs to a Microsoft Reserved (MSR) partition.
	DiskExtentTypeMSR DiskExtentType = 5
	// VDS_DET_LDM:  The extent belongs to a disk management metadata partition.
	DiskExtentTypeLDM DiskExtentType = 6
	// VDS_DET_UNUSABLE:  The extent belongs to an area of unusable space.
	DiskExtentTypeUnusable DiskExtentType = 32767
)

func (o DiskExtentType) String() string {
	switch o {
	case DiskExtentTypeUnknown:
		return "DiskExtentTypeUnknown"
	case DiskExtentTypeFree:
		return "DiskExtentTypeFree"
	case DiskExtentTypeData:
		return "DiskExtentTypeData"
	case DiskExtentTypeOEM:
		return "DiskExtentTypeOEM"
	case DiskExtentTypeESP:
		return "DiskExtentTypeESP"
	case DiskExtentTypeMSR:
		return "DiskExtentTypeMSR"
	case DiskExtentTypeLDM:
		return "DiskExtentTypeLDM"
	case DiskExtentTypeUnusable:
		return "DiskExtentTypeUnusable"
	}
	return "Invalid"
}

// PartitionStyle type represents VDS_PARTITION_STYLE RPC enumeration.
//
// The VDS_PARTITION_STYLE enumeration defines the styles of partitions.
type PartitionStyle uint16

var (
	// VDS_PST_UNKNOWN:  The partition format is unknown.
	PartitionStyleUnknown PartitionStyle = 0
	// VDS_PST_MBR:  The partition format is master boot record (MBR).
	PartitionStyleMBR PartitionStyle = 1
	// VDS_PST_GPT:  The partition format is GUID partitioning table (GPT).
	PartitionStyleGPT PartitionStyle = 2
)

func (o PartitionStyle) String() string {
	switch o {
	case PartitionStyleUnknown:
		return "PartitionStyleUnknown"
	case PartitionStyleMBR:
		return "PartitionStyleMBR"
	case PartitionStyleGPT:
		return "PartitionStyleGPT"
	}
	return "Invalid"
}

// PartitionFlag type represents VDS_PARTITION_FLAG RPC enumeration.
//
// The VDS_PARTITION_FLAG enumeration defines flags that describe partitions. A value
// that accepts these flags MUST have the following flag set or MUST have the value
// set to 0 if the flag is not applicable to a given partition.
type PartitionFlag uint16

var (
	// VDS_PTF_SYSTEM:  Value that indicates that the partition is a system partition.
	PartitionFlagSystem PartitionFlag = 1
)

func (o PartitionFlag) String() string {
	switch o {
	case PartitionFlagSystem:
		return "PartitionFlagSystem"
	}
	return "Invalid"
}

// VolumeType type represents VDS_VOLUME_TYPE RPC enumeration.
//
// The VDS_VOLUME_TYPE enumeration defines the set of valid types for a volume object.
type VolumeType uint16

var (
	// VDS_VT_UNKNOWN:  The status of the volume is unknown.
	VolumeTypeUnknown VolumeType = 0
	// VDS_VT_SIMPLE:  The volume type is simple: it is composed of extents from exactly
	// one disk.
	VolumeTypeSimple VolumeType = 10
	// VDS_VT_SPAN:  The volume type is spanned: it is composed of extents from more than
	// one disk.
	VolumeTypeSpan VolumeType = 11
	// VDS_VT_STRIPE:  The volume type is striped, which is equivalent to RAID-0.
	VolumeTypeStripe VolumeType = 12
	// VDS_VT_MIRROR:  The volume type is mirrored, which is equivalent to RAID-1.
	VolumeTypeMirror VolumeType = 13
	// VDS_VT_PARITY:  The volume type is striped with parity, which accounts for RAID
	// levels 3, 4, 5, and 6.
	VolumeTypeParity VolumeType = 14
)

func (o VolumeType) String() string {
	switch o {
	case VolumeTypeUnknown:
		return "VolumeTypeUnknown"
	case VolumeTypeSimple:
		return "VolumeTypeSimple"
	case VolumeTypeSpan:
		return "VolumeTypeSpan"
	case VolumeTypeStripe:
		return "VolumeTypeStripe"
	case VolumeTypeMirror:
		return "VolumeTypeMirror"
	case VolumeTypeParity:
		return "VolumeTypeParity"
	}
	return "Invalid"
}

// TransitionState type represents VDS_TRANSITION_STATE RPC enumeration.
//
// The VDS_TRANSITION_STATE enumeration defines the set of valid transition state values
// for a VDS object.
type TransitionState uint16

var (
	// VDS_TS_UNKNOWN:  The transition state of the object cannot be determined.
	TransitionStateUnknown TransitionState = 0
	// VDS_TS_STABLE:  The object is stable. No configuration activity is currently in
	// progress.
	TransitionStateStable TransitionState = 1
	// VDS_TS_EXTENDING:  The object is being extended.
	TransitionStateExtending TransitionState = 2
	// VDS_TS_SHRINKING:  The object is being shrunk.
	TransitionStateShrinking TransitionState = 3
	// VDS_TS_RECONFIGING:  The object is being automatically reconfigured.
	TransitionStateReconfiging TransitionState = 4
)

func (o TransitionState) String() string {
	switch o {
	case TransitionStateUnknown:
		return "TransitionStateUnknown"
	case TransitionStateStable:
		return "TransitionStateStable"
	case TransitionStateExtending:
		return "TransitionStateExtending"
	case TransitionStateShrinking:
		return "TransitionStateShrinking"
	case TransitionStateReconfiging:
		return "TransitionStateReconfiging"
	}
	return "Invalid"
}

// FormatOptionFlags type represents VDS_FORMAT_OPTION_FLAGS RPC enumeration.
//
// The VDS_FORMAT_OPTION_FLAGS enumeration defines the set of valid format option values.
type FormatOptionFlags uint16

var (
	FormatOptionFlagNone              FormatOptionFlags = 0
	FormatOptionFlagForce             FormatOptionFlags = 1
	FormatOptionFlagQuick             FormatOptionFlags = 2
	FormatOptionFlagCompression       FormatOptionFlags = 4
	FormatOptionFlagDuplicateMetadata FormatOptionFlags = 8
)

func (o FormatOptionFlags) String() string {
	switch o {
	case FormatOptionFlagNone:
		return "FormatOptionFlagNone"
	case FormatOptionFlagForce:
		return "FormatOptionFlagForce"
	case FormatOptionFlagQuick:
		return "FormatOptionFlagQuick"
	case FormatOptionFlagCompression:
		return "FormatOptionFlagCompression"
	case FormatOptionFlagDuplicateMetadata:
		return "FormatOptionFlagDuplicateMetadata"
	}
	return "Invalid"
}

// DiskFlag type represents VDS_DISK_FLAG RPC enumeration.
//
// The VDS_DISK_FLAG enumeration defines the properties of a disk.
type DiskFlag uint16

var (
	// VDS_DF_AUDIO_CD:  The disk is an audio CD, as specified in [IEC60908].
	DiskFlagAudioCD DiskFlag = 1
	// VDS_DF_HOTSPARE:  The disk is a hot spare.
	DiskFlagHotSpare DiskFlag = 2
	// VDS_DF_RESERVE_CAPABLE:  The disk can be reserved for a host.
	DiskFlagReserveCapable DiskFlag = 4
	// VDS_DF_MASKED:  The disk is currently hidden from the host.
	DiskFlagMasked DiskFlag = 8
	// VDS_DF_STYLE_CONVERTIBLE:  The disk is convertible between the MBR partition format
	// and the GPT partition format.
	DiskFlagStyleConvertible DiskFlag = 16
	// VDS_DF_CLUSTERED:  The disk is clustered.
	DiskFlagClustered DiskFlag = 32
	// VDS_DF_READ_ONLY:  The disk read-only attribute is set.
	DiskFlagReadOnly DiskFlag = 64
	// VDS_DF_SYSTEM_DISK:  The disk contains the system volume.
	DiskFlagSystemDisk DiskFlag = 128
	// VDS_DF_BOOT_DISK:  The disk contains the boot volume.
	DiskFlagBootDisk DiskFlag = 256
	// VDS_DF_PAGEFILE_DISK:  The disk contains the paging file on one of its volumes.
	DiskFlagPageFileDisk DiskFlag = 512
	// VDS_DF_HIBERNATIONFILE_DISK:  The disk contains the hibernation file on one of its
	// volumes.
	DiskFlagHibernationFileDisk DiskFlag = 1024
	// VDS_DF_CRASHDUMP_DISK:  The disk is configured to contain a crash-dump file on one
	// of its volumes.
	DiskFlagCrashDumpDisk DiskFlag = 2048
	// VDS_DF_HAS_ARC_PATH:  The disk has an Advanced RISC Computing (ARC) path specified
	// in the BIOS. For information on ARC paths, see [KB102873].
	DiskFlagHasARCPath DiskFlag = 4096
	// VDS_DF_DYNAMIC:  The disk is a logical disk manager dynamic disk.
	DiskFlagDynamic DiskFlag = 8192
	// VDS_DF_BOOT_FROM_DISK:  Indicates the disk from which the machine will boot. Note
	// that this is BIOS disk 0 on the MBR, not the current system volume disk. For example,
	// if the machine boots to Windows PE, this flag is set on BIOS disk 0. For EFI machines,
	// this flag is set on a disk containing the EFI system partition used to boot the machine.
	DiskFlagBootFromDisk DiskFlag = 16384
	// VDS_DF_CURRENT_READ_ONLY:  Indicates that the disk is in a read-only state. If this
	// flag is not set, the disk is read/write. Unlike the VDS_DF_READ_ONLY flag, which
	// is used to indicate the disk's read-only attribute maintained by the operating system,
	// this flag reflects the actual disk state. This flag cannot be set by using the IVdsDisk::SetFlags
	// method or cleared by using the IVdsDisk::ClearFlags method.
	DiskFlagCurrentReadOnly DiskFlag = 32768
)

func (o DiskFlag) String() string {
	switch o {
	case DiskFlagAudioCD:
		return "DiskFlagAudioCD"
	case DiskFlagHotSpare:
		return "DiskFlagHotSpare"
	case DiskFlagReserveCapable:
		return "DiskFlagReserveCapable"
	case DiskFlagMasked:
		return "DiskFlagMasked"
	case DiskFlagStyleConvertible:
		return "DiskFlagStyleConvertible"
	case DiskFlagClustered:
		return "DiskFlagClustered"
	case DiskFlagReadOnly:
		return "DiskFlagReadOnly"
	case DiskFlagSystemDisk:
		return "DiskFlagSystemDisk"
	case DiskFlagBootDisk:
		return "DiskFlagBootDisk"
	case DiskFlagPageFileDisk:
		return "DiskFlagPageFileDisk"
	case DiskFlagHibernationFileDisk:
		return "DiskFlagHibernationFileDisk"
	case DiskFlagCrashDumpDisk:
		return "DiskFlagCrashDumpDisk"
	case DiskFlagHasARCPath:
		return "DiskFlagHasARCPath"
	case DiskFlagDynamic:
		return "DiskFlagDynamic"
	case DiskFlagBootFromDisk:
		return "DiskFlagBootFromDisk"
	case DiskFlagCurrentReadOnly:
		return "DiskFlagCurrentReadOnly"
	}
	return "Invalid"
}

// DiskStatus type represents VDS_DISK_STATUS RPC enumeration.
//
// The VDS_DISK_STATUS enumeration defines the status of a disk.
type DiskStatus uint16

var (
	// VDS_DS_UNKNOWN:  The disk status is unknown.
	DiskStatusUnknown DiskStatus = 0
	// VDS_DS_ONLINE:  The disk is online.
	DiskStatusOnline DiskStatus = 1
	// VDS_DS_NOT_READY:  The disk is not ready.
	DiskStatusNotReady DiskStatus = 2
	// VDS_DS_NO_MEDIA:  The disk has no media.
	DiskStatusNoMedia DiskStatus = 3
	// VDS_DS_OFFLINE:  The disk is offline. Offline disks have no volume devices exposed.
	DiskStatusOffline DiskStatus = 4
	// VDS_DS_FAILED:  The disk failed.
	DiskStatusFailed DiskStatus = 5
	// VDS_DS_MISSING:  The disk is missing; it is no longer available to the operating
	// system.
	DiskStatusMissing DiskStatus = 6
)

func (o DiskStatus) String() string {
	switch o {
	case DiskStatusUnknown:
		return "DiskStatusUnknown"
	case DiskStatusOnline:
		return "DiskStatusOnline"
	case DiskStatusNotReady:
		return "DiskStatusNotReady"
	case DiskStatusNoMedia:
		return "DiskStatusNoMedia"
	case DiskStatusOffline:
		return "DiskStatusOffline"
	case DiskStatusFailed:
		return "DiskStatusFailed"
	case DiskStatusMissing:
		return "DiskStatusMissing"
	}
	return "Invalid"
}

// LUNReserveMode type represents VDS_LUN_RESERVE_MODE RPC enumeration.
//
// The VDS_LUN_RESERVE_MODE enumeration defines the sharing mode of a disk.
type LUNReserveMode uint16

var (
	// VDS_LRM_NONE:  The disk has no assigned sharing mode.
	LUNReserveModeNone LUNReserveMode = 0
	// VDS_LRM_EXCLUSIVE_RW:  The disk is reserved for exclusive access.
	LUNReserveModeExclusiveRW LUNReserveMode = 1
	// VDS_LRM_EXCLUSIVE_RO:  The disk is available for read access.
	LUNReserveModeExclusiveReadOnly LUNReserveMode = 2
	// VDS_LRM_SHARED_RO:  The disk is available for shared read access.
	LUNReserveModeSharedReadOnly LUNReserveMode = 3
	// VDS_LRM_SHARED_RW:  The disk is available for shared read/write access.
	LUNReserveModeSharedRW LUNReserveMode = 4
)

func (o LUNReserveMode) String() string {
	switch o {
	case LUNReserveModeNone:
		return "LUNReserveModeNone"
	case LUNReserveModeExclusiveRW:
		return "LUNReserveModeExclusiveRW"
	case LUNReserveModeExclusiveReadOnly:
		return "LUNReserveModeExclusiveReadOnly"
	case LUNReserveModeSharedReadOnly:
		return "LUNReserveModeSharedReadOnly"
	case LUNReserveModeSharedRW:
		return "LUNReserveModeSharedRW"
	}
	return "Invalid"
}

// VolumeStatus type represents VDS_VOLUME_STATUS RPC enumeration.
//
// The VDS_VOLUME_STATUS enumeration defines the set of object status values for a volume.
type VolumeStatus uint16

var (
	// VDS_VS_UNKNOWN:  The status of the volume is unknown.
	VolumeStatusUnknown VolumeStatus = 0
	// VDS_VS_ONLINE:  The volume is available.
	VolumeStatusOnline VolumeStatus = 1
	// VDS_VS_NO_MEDIA:  The volume belongs to a removable media device, such as a CD-ROM
	// or DVD-ROM drive, but the device does not currently have media in the drive.
	VolumeStatusNoMedia VolumeStatus = 3
	// VDS_VS_OFFLINE:  When this status is set, it (1) indicates that no path names for the volume are available for use by applications, and (2) prevents READ and READ|WRITE handles to the volume device being opened. When a volume transitions to this state, calls to open a new handle against the volume device fail, but any in-progress I/O against the volume will complete before all I/O to the volume is stopped.<13>
	VolumeStatusOffline VolumeStatus = 4
	// VDS_VS_FAILED:  The volume is unavailable.
	VolumeStatusFailed VolumeStatus = 5
)

func (o VolumeStatus) String() string {
	switch o {
	case VolumeStatusUnknown:
		return "VolumeStatusUnknown"
	case VolumeStatusOnline:
		return "VolumeStatusOnline"
	case VolumeStatusNoMedia:
		return "VolumeStatusNoMedia"
	case VolumeStatusOffline:
		return "VolumeStatusOffline"
	case VolumeStatusFailed:
		return "VolumeStatusFailed"
	}
	return "Invalid"
}

// VolumeFlag type represents VDS_VOLUME_FLAG RPC enumeration.
//
// The VDS_VOLUME_FLAG enumeration defines the set of valid flags for a volume object.
type VolumeFlag uint32

var (
	// VDS_VF_SYSTEM_VOLUME:  If set, the volume is a system volume. It contains the boot
	// loader that is used to invoke the operating system on the boot volume.
	VolumeFlagSystemVolume VolumeFlag = 1
	// VDS_VF_BOOT_VOLUME:  If set, the volume is a boot volume that contains the operating
	// system.
	VolumeFlagBootVolume VolumeFlag = 2
	// VDS_VF_ACTIVE:  If set, the volume is an active volume. It can become the system
	// volume at system startup if the BIOS is configured to select that disk for startup.
	VolumeFlagActive VolumeFlag = 4
	// VDS_VF_READONLY:  If set, the volume can be read from but not written to.
	VolumeFlagReadonly VolumeFlag = 8
	// VDS_VF_HIDDEN:  If set, the volume does not automatically get assigned mount points
	// or drive letters that can be used to access the volume.
	VolumeFlagHidden VolumeFlag = 16
	// VDS_VF_CAN_EXTEND:  If set, the volume size can be extended.
	VolumeFlagCanExtend VolumeFlag = 32
	// VDS_VF_CAN_SHRINK:  If set, the volume size can be reduced.
	VolumeFlagCanShrink VolumeFlag = 64
	// VDS_VF_PAGEFILE:  If this flag is set, the volume contains a page file.
	VolumeFlagPageFile VolumeFlag = 128
	// VDS_VF_HIBERNATION:  If set, the volume holds the files that are used when the system
	// hibernates.
	VolumeFlagHibernation VolumeFlag = 256
	// VDS_VF_CRASHDUMP:  If set, the volume acts as a crash-dump device.
	VolumeFlagCrashDump VolumeFlag = 512
	// VDS_VF_INSTALLABLE:  If set, callers can use the volume to install an operating
	// system.
	VolumeFlagInstallable VolumeFlag = 1024
	// VDS_VF_LBN_REMAP_ENABLED:  If set, VDS can dynamically change the position of the
	// volume on the disk. This flag is not valid for basic and dynamic volumes and is only
	// supported by some third-party volume managers.
	VolumeFlagLbnRemapEnabled VolumeFlag = 2048
	// VDS_VF_FORMATTING:  If set, the volume is being formatted.
	VolumeFlagFormatting VolumeFlag = 4096
	// VDS_VF_NOT_FORMATTABLE:  If set, the volume cannot be formatted.
	VolumeFlagNotFormattable VolumeFlag = 8192
	// VDS_VF_NTFS_NOT_SUPPORTED:  If set, the volume does not support the NTFS file system
	// but can support other file systems.
	VolumeFlagNTFSNotSupported VolumeFlag = 16384
	// VDS_VF_FAT32_NOT_SUPPORTED:  If set, the volume does not support FAT32.
	VolumeFlagFAT32NotSupported VolumeFlag = 32768
	// VDS_VF_FAT_NOT_SUPPORTED:  If set, the volume does not support FAT.
	VolumeFlagFATNotSupported VolumeFlag = 65536
	// VDS_VF_NO_DEFAULT_DRIVE_LETTER:  If set, the operating system does not automatically
	// assign a drive letter when the volume is created or a disk containing existing volumes
	// is connected to the operating system. When cleared, the operating system assigns
	// a drive letter to the volume. Callers can set and clear this flag. For basic GPT
	// volumes and dynamic disk volumes, assigning or removing a drive letter toggles this
	// flag.<14>
	VolumeFlagNoDefaultDriveLetter VolumeFlag = 131072
	// VDS_VF_PERMANENTLY_DISMOUNTED:  If set, the volume is unavailable and requires a
	// mount-point assignment. VDS sets this flag after the caller invokes the IVdsVolumeMF::Dismount
	// method, setting the bForce and bPermanent parameters to TRUE.
	VolumeFlagPermanentlyDismounted VolumeFlag = 262144
	// VDS_VF_PERMANENT_DISMOUNT_SUPPORTED:  If set, the volume supports bPermanent for
	// the IVdsVolumeMF::Dismount method. This flag cannot be set or cleared by the client.
	// This flag is set by the server if it applies.
	VolumeFlagPermanentDismountSupported VolumeFlag = 524288
	// VDS_VF_SHADOW_COPY:  If set, the volume is a shadow copy of another volume. This
	// flag is set when the snapshot is taken, and it is cleared when the snapshot is broken
	// from the original volume. The VDS_VF_SHADOW_COPY flag is an indication for software-like
	// file system filter drivers (for example, antivirus) to avoid attaching to the volume.
	// Applications can use the attribute to differentiate snapshots from production volumes.
	// Applications that create a Fast Recovery, in which a shadow copy LUN is made into
	// a non-snapshot by clearing the read-only and hidden bit, will need to clear this
	// bit as well.
	VolumeFlagShadowCopy VolumeFlag = 1048576
	// VDS_VF_FVE_ENABLED:  The volume is encrypted with full-volume encryption.<15>
	VolumeFlagFveEnabled VolumeFlag = 2097152
	// VDS_VF_DIRTY<16>:  The volume's dirty bit is set.
	VolumeFlagDirty VolumeFlag = 4194304
	// VDS_VF_REFS_NOT_SUPPORTED<17>:  The volume does not support ReFS.
	//
	// * *Dynamic disk volumes* - The flag is per volume. *VDS_VF_NO_DEFAULT_DRIVE_LETTER*
	// is set at volume creation. <18> ( 6b2552ee-ba27-409f-99a8-18841573327a#Appendix_A_18
	// ) The flag toggles when drive letters are assigned or removed, and can also be set
	// or cleared using any of the Set/ClearFlags methods.
	//
	// * *MBR basic disk volumes* - The flag is applied to all volumes created on the disk
	// after the flag is set. It is set per disk for basic MBR ( 8eef5d42-22d7-4302-aed9-12face91505a#gt_b251c771-5ccf-40f2-b98d-0119db210b4b
	// ) disks, not per volume. The flag is only set or cleared if an explicit call is made
	// to the IVdsVolume::SetFlags (section 3.4.5.2.32.10) ( b0ed8172-e096-4685-9e68-f07fb6fddc4d
	// ) and IVdsVolume::ClearFlag (section 3.4.5.2.32.11) ( 002b6121-2e6d-4334-864a-be59ecbb8a07
	// ) methods, respectively. For example, the *VDS_VF_NO_DEFAULT_DRIVE_LETTER* flag is
	// not toggled as drive letters are assigned to or removed from specific volumes.
	//
	// * *GPT basic disk volumes* - The flag is per volume, data partitions ( 8eef5d42-22d7-4302-aed9-12face91505a#gt_2f24f458-7d39-47a2-93f7-de433ea85c75
	// ) only. *VDS_VF_NO_DEFAULT_DRIVE_LETTER* is set at volume creation and toggled when
	// drive letters are assigned or removed (by VDS).
	VolumeFlagReFSNotSupported VolumeFlag = 8388608
)

func (o VolumeFlag) String() string {
	switch o {
	case VolumeFlagSystemVolume:
		return "VolumeFlagSystemVolume"
	case VolumeFlagBootVolume:
		return "VolumeFlagBootVolume"
	case VolumeFlagActive:
		return "VolumeFlagActive"
	case VolumeFlagReadonly:
		return "VolumeFlagReadonly"
	case VolumeFlagHidden:
		return "VolumeFlagHidden"
	case VolumeFlagCanExtend:
		return "VolumeFlagCanExtend"
	case VolumeFlagCanShrink:
		return "VolumeFlagCanShrink"
	case VolumeFlagPageFile:
		return "VolumeFlagPageFile"
	case VolumeFlagHibernation:
		return "VolumeFlagHibernation"
	case VolumeFlagCrashDump:
		return "VolumeFlagCrashDump"
	case VolumeFlagInstallable:
		return "VolumeFlagInstallable"
	case VolumeFlagLbnRemapEnabled:
		return "VolumeFlagLbnRemapEnabled"
	case VolumeFlagFormatting:
		return "VolumeFlagFormatting"
	case VolumeFlagNotFormattable:
		return "VolumeFlagNotFormattable"
	case VolumeFlagNTFSNotSupported:
		return "VolumeFlagNTFSNotSupported"
	case VolumeFlagFAT32NotSupported:
		return "VolumeFlagFAT32NotSupported"
	case VolumeFlagFATNotSupported:
		return "VolumeFlagFATNotSupported"
	case VolumeFlagNoDefaultDriveLetter:
		return "VolumeFlagNoDefaultDriveLetter"
	case VolumeFlagPermanentlyDismounted:
		return "VolumeFlagPermanentlyDismounted"
	case VolumeFlagPermanentDismountSupported:
		return "VolumeFlagPermanentDismountSupported"
	case VolumeFlagShadowCopy:
		return "VolumeFlagShadowCopy"
	case VolumeFlagFveEnabled:
		return "VolumeFlagFveEnabled"
	case VolumeFlagDirty:
		return "VolumeFlagDirty"
	case VolumeFlagReFSNotSupported:
		return "VolumeFlagReFSNotSupported"
	}
	return "Invalid"
}

// PackNotification structure represents VDS_PACK_NOTIFICATION RPC structure.
//
// The VDS_PACK_NOTIFICATION structure provides information about a pack notification.
type PackNotification struct {
	// ulEvent:  The type of pack notification; it MUST be one of the following values.
	//
	//	+-------------------------------+-----------------------------+
	//	|                               |                             |
	//	|             VALUE             |           MEANING           |
	//	|                               |                             |
	//	+-------------------------------+-----------------------------+
	//	+-------------------------------+-----------------------------+
	//	| VDS_NF_PACK_ARRIVE 0x00000001 | The pack was newly created. |
	//	+-------------------------------+-----------------------------+
	//	| VDS_NF_PACK_DEPART 0x00000002 | The pack was deleted.       |
	//	+-------------------------------+-----------------------------+
	//	| VDS_NF_PACK_MODIFY 0x00000003 | The pack was modified.      |
	//	+-------------------------------+-----------------------------+
	Event uint32 `idl:"name:ulEvent" json:"event"`
	// packId:  The VDS object ID of the pack object to which the notification refers.
	PackID *ObjectID `idl:"name:packId" json:"pack_id"`
}

func (o *PackNotification) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PackNotification) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Event); err != nil {
		return err
	}
	if o.PackID != nil {
		if err := o.PackID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PackNotification) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Event); err != nil {
		return err
	}
	if o.PackID == nil {
		o.PackID = &ObjectID{}
	}
	if err := o.PackID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DiskNotification structure represents VDS_DISK_NOTIFICATION RPC structure.
//
// The VDS_DISK_NOTIFICATION structure provides information about a disk notification.
type DiskNotification struct {
	// ulEvent:  The type of disk notification; it MUST be one of the following values.
	//
	//	+-------------------------------+--------------------------------------------------------+
	//	|                               |                                                        |
	//	|             VALUE             |                        MEANING                         |
	//	|                               |                                                        |
	//	+-------------------------------+--------------------------------------------------------+
	//	+-------------------------------+--------------------------------------------------------+
	//	| VDS_NF_DISK_ARRIVE 0x00000008 | The disk has become visible to the operating system.   |
	//	+-------------------------------+--------------------------------------------------------+
	//	| VDS_NF_DISK_DEPART 0x00000009 | The disk is no longer visible to the operating system. |
	//	+-------------------------------+--------------------------------------------------------+
	//	| VDS_NF_DISK_MODIFY 0x0000000A | The disk or its properties were modified.              |
	//	+-------------------------------+--------------------------------------------------------+
	Event uint32 `idl:"name:ulEvent" json:"event"`
	// diskId:  The VDS object ID of the disk object to which the notification refers.
	DiskID *ObjectID `idl:"name:diskId" json:"disk_id"`
}

func (o *DiskNotification) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DiskNotification) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Event); err != nil {
		return err
	}
	if o.DiskID != nil {
		if err := o.DiskID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DiskNotification) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Event); err != nil {
		return err
	}
	if o.DiskID == nil {
		o.DiskID = &ObjectID{}
	}
	if err := o.DiskID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// VolumeNotification structure represents VDS_VOLUME_NOTIFICATION RPC structure.
//
// The VDS_VOLUME_NOTIFICATION structure provides information about a volume change
// notification.
type VolumeNotification struct {
	// ulEvent:  Determines the volume event for which an application will be notified;
	// it MUST be one of the following values.
	//
	//	+----------------------------------------------+------------------------------------------------------------------+
	//	|                                              |                                                                  |
	//	|                    VALUE                     |                             MEANING                              |
	//	|                                              |                                                                  |
	//	+----------------------------------------------+------------------------------------------------------------------+
	//	+----------------------------------------------+------------------------------------------------------------------+
	//	| VDS_NF_VOLUME_ARRIVE 0x00000004              | A new volume is visible to the operating system.                 |
	//	+----------------------------------------------+------------------------------------------------------------------+
	//	| VDS_NF_VOLUME_DEPART 0x00000005              | An existing volume is no longer visible to the operating system. |
	//	+----------------------------------------------+------------------------------------------------------------------+
	//	| VDS_NF_VOLUME_MODIFY 0x00000006              | The volume was modified.                                         |
	//	+----------------------------------------------+------------------------------------------------------------------+
	//	| VDS_NF_VOLUME_REBUILDING_PROGRESS 0x00000007 | A fault tolerant volume is being regenerated or resynchronized.  |
	//	+----------------------------------------------+------------------------------------------------------------------+
	Event uint32 `idl:"name:ulEvent" json:"event"`
	// volumeId:  The VDS object ID of the volume object to which the notification refers.
	VolumeID *ObjectID `idl:"name:volumeId" json:"volume_id"`
	// plexId:  The VDS object ID of a volume plex object to which the notification refers,
	// if any. VDS applies this identifier during the rebuild operation, which can execute
	// on multiple plexes at different rates.
	PlexID *ObjectID `idl:"name:plexId" json:"plex_id"`
	// ulPercentCompleted:  The percentage of completion for the operation. Valid values
	// range from 0-100.
	PercentCompleted uint32 `idl:"name:ulPercentCompleted" json:"percent_completed"`
}

func (o *VolumeNotification) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *VolumeNotification) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Event); err != nil {
		return err
	}
	if o.VolumeID != nil {
		if err := o.VolumeID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.PlexID != nil {
		if err := o.PlexID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PercentCompleted); err != nil {
		return err
	}
	return nil
}
func (o *VolumeNotification) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Event); err != nil {
		return err
	}
	if o.VolumeID == nil {
		o.VolumeID = &ObjectID{}
	}
	if err := o.VolumeID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.PlexID == nil {
		o.PlexID = &ObjectID{}
	}
	if err := o.PlexID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.PercentCompleted); err != nil {
		return err
	}
	return nil
}

// PartitionNotification structure represents VDS_PARTITION_NOTIFICATION RPC structure.
//
// The VDS_PARTITION_NOTIFICATION structure provides information about a partition notification.
type PartitionNotification struct {
	// ulEvent:  Determines the partition event for which an application will be notified;
	// it MUST be one of the following values.
	//
	//	+------------------------------------+---------------------------------------------------------------------+
	//	|                                    |                                                                     |
	//	|               VALUE                |                               MEANING                               |
	//	|                                    |                                                                     |
	//	+------------------------------------+---------------------------------------------------------------------+
	//	+------------------------------------+---------------------------------------------------------------------+
	//	| VDS_NF_PARTITION_ARRIVE 0x0000000B | A new partition is visible to the operating system.                 |
	//	+------------------------------------+---------------------------------------------------------------------+
	//	| VDS_NF_PARTITION_DEPART 0x0000000C | An existing partition is no longer visible to the operating system. |
	//	+------------------------------------+---------------------------------------------------------------------+
	//	| VDS_NF_PARTITION_MODIFY 0x0000000D | An existing partition changed.                                      |
	//	+------------------------------------+---------------------------------------------------------------------+
	Event uint32 `idl:"name:ulEvent" json:"event"`
	// diskId:  The VDS object ID of the disk object containing the partition that triggered
	// the event.
	DiskID *ObjectID `idl:"name:diskId" json:"disk_id"`
	// ullOffset:  The byte offset of the partition from the beginning of the disk.
	Offset uint64 `idl:"name:ullOffset" json:"offset"`
}

func (o *PartitionNotification) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PartitionNotification) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.Event); err != nil {
		return err
	}
	if o.DiskID != nil {
		if err := o.DiskID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Offset); err != nil {
		return err
	}
	return nil
}
func (o *PartitionNotification) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.Event); err != nil {
		return err
	}
	if o.DiskID == nil {
		o.DiskID = &ObjectID{}
	}
	if err := o.DiskID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Offset); err != nil {
		return err
	}
	return nil
}

// DriveLetterNotification structure represents VDS_DRIVE_LETTER_NOTIFICATION RPC structure.
//
// The VDS_DRIVE_LETTER_NOTIFICATION structure provides information about a drive letter
// notification.
type DriveLetterNotification struct {
	// ulEvent:  Determines the drive-letter event for which an application will be notified;
	// it MUST be one of the following values.
	//
	//	+---------------------------------------+-------------------------------------------------+
	//	|                                       |                                                 |
	//	|                 VALUE                 |                     MEANING                     |
	//	|                                       |                                                 |
	//	+---------------------------------------+-------------------------------------------------+
	//	+---------------------------------------+-------------------------------------------------+
	//	| VDS_NF_DRIVE_LETTER_FREE 0x000000C9   | The drive letter is no longer in use.           |
	//	+---------------------------------------+-------------------------------------------------+
	//	| VDS_NF_DRIVE_LETTER_ASSIGN 0x000000CA | The drive letter has been assigned to a volume. |
	//	+---------------------------------------+-------------------------------------------------+
	Event uint32 `idl:"name:ulEvent" json:"event"`
	// wcLetter:  The drive letter that triggered the event, as a single uppercase or lowercase
	// alphabetical (A-Z) Unicode character.
	Letter uint16 `idl:"name:wcLetter" json:"letter"`
	// volumeId:  The VDS object ID of the volume object to which the drive letter is assigned.
	// If the drive letter is freed, the volume identifier is GUID_NULL.
	VolumeID *ObjectID `idl:"name:volumeId" json:"volume_id"`
}

func (o *DriveLetterNotification) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DriveLetterNotification) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Event); err != nil {
		return err
	}
	if err := w.WriteData(o.Letter); err != nil {
		return err
	}
	if o.VolumeID != nil {
		if err := o.VolumeID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DriveLetterNotification) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Event); err != nil {
		return err
	}
	if err := w.ReadData(&o.Letter); err != nil {
		return err
	}
	if o.VolumeID == nil {
		o.VolumeID = &ObjectID{}
	}
	if err := o.VolumeID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// FileSystemNotification structure represents VDS_FILE_SYSTEM_NOTIFICATION RPC structure.
//
// The VDS_FILE_SYSTEM_NOTIFICATION structure provides information about a file system
// notification.
type FileSystemNotification struct {
	// ulEvent:  Determines the file system event for which an application will be notified;
	// it MUST be one of the following values.
	//
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                               |                                                                                  |
	//	|                     VALUE                     |                                     MEANING                                      |
	//	|                                               |                                                                                  |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| VDS_NF_FILE_SYSTEM_MODIFY 0x000000CB          | A volume received a new label, or a file system was extended or shrunk; does not |
	//	|                                               | include a change to the file system compression flags.                           |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| VDS_NF_FILE_SYSTEM_FORMAT_PROGRESS 0x000000CC | A file system is being formatted.                                                |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	Event uint32 `idl:"name:ulEvent" json:"event"`
	// volumeId:  The VDS object ID of the volume object containing the file system that
	// triggered the event.
	VolumeID *ObjectID `idl:"name:volumeId" json:"volume_id"`
	// dwPercentCompleted:  The completed format progress as a percentage of the whole.
	PercentCompleted uint32 `idl:"name:dwPercentCompleted" json:"percent_completed"`
}

func (o *FileSystemNotification) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *FileSystemNotification) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Event); err != nil {
		return err
	}
	if o.VolumeID != nil {
		if err := o.VolumeID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PercentCompleted); err != nil {
		return err
	}
	return nil
}
func (o *FileSystemNotification) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Event); err != nil {
		return err
	}
	if o.VolumeID == nil {
		o.VolumeID = &ObjectID{}
	}
	if err := o.VolumeID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.PercentCompleted); err != nil {
		return err
	}
	return nil
}

// MountPointNotification structure represents VDS_MOUNT_POINT_NOTIFICATION RPC structure.
//
// The VDS_MOUNT_POINT_NOTIFICATION structure provides information about a mount point
// change notification.
type MountPointNotification struct {
	// ulEvent:  Determines the mount point event for which an application will be notified;
	// it MUST be the following value.
	//
	//	+---------------------------------------+--------------------------+
	//	|                                       |                          |
	//	|                 VALUE                 |         MEANING          |
	//	|                                       |                          |
	//	+---------------------------------------+--------------------------+
	//	+---------------------------------------+--------------------------+
	//	| VDS_NF_MOUNT_POINTS_CHANGE 0x000000CD | The mount point changed. |
	//	+---------------------------------------+--------------------------+
	Event uint32 `idl:"name:ulEvent" json:"event"`
	// volumeId:  The VDS object ID of the volume object containing the mount point that
	// triggered the event.
	VolumeID *ObjectID `idl:"name:volumeId" json:"volume_id"`
}

func (o *MountPointNotification) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *MountPointNotification) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Event); err != nil {
		return err
	}
	if o.VolumeID != nil {
		if err := o.VolumeID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MountPointNotification) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Event); err != nil {
		return err
	}
	if o.VolumeID == nil {
		o.VolumeID = &ObjectID{}
	}
	if err := o.VolumeID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RecoverAction type represents VDS_RECOVER_ACTION RPC enumeration.
//
// The VDS_RECOVER_ACTION enumeration defines the set of valid client actions to be
// taken in response to a notification with target type VDS_NTT_SERVICE.
type RecoverAction uint16

var (
	// VDS_RA_UNKNOWN:  Client action to be taken is unknown.
	RecoverActionUnknown RecoverAction = 0
	// VDS_RA_REFRESH:  Client action to be taken is to call the IVdsService::Refresh (section
	// 3.4.5.2.4.10) method.
	RecoverActionRefresh RecoverAction = 1
	// VDS_RA_RESTART:  Client action to be taken is to restart the service.
	RecoverActionRestart RecoverAction = 2
)

func (o RecoverAction) String() string {
	switch o {
	case RecoverActionUnknown:
		return "RecoverActionUnknown"
	case RecoverActionRefresh:
		return "RecoverActionRefresh"
	case RecoverActionRestart:
		return "RecoverActionRestart"
	}
	return "Invalid"
}

// ServiceNotification structure represents VDS_SERVICE_NOTIFICATION RPC structure.
//
// The VDS_ SERVICE _NOTIFICATION structure provides information about state changes
// to the service object.<19>
type ServiceNotification struct {
	// ulEvent:  The type of service notification; it MUST be set to the following value.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                                       |                                                                                  |
	//	|                 VALUE                 |                                     MEANING                                      |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| VDS_NF_SERVICE_OUT_OF_SYNC 0x0000012D | The service's cache has become inconsistent or the service has encountered an    |
	//	|                                       | error requiring client action.                                                   |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	Event uint32 `idl:"name:ulEvent" json:"event"`
	// action:  The type of action required by the client to return the service to a functioning,
	// consistent state; it MUST be one of the following values.
	//
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	|                           |                                                                                  |
	//	|           VALUE           |                                     MEANING                                      |
	//	|                           |                                                                                  |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| VDS_RA_UNKNOWN 0x00000000 | The client corrective action is unknown.                                         |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| VDS_RA_REFRESH 0x00000001 | The client corrective action is to call the IVdsService::Refresh (section        |
	//	|                           | 3.4.5.2.4.10) method.                                                            |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| VDS_RA_RESTART 0x00000002 | The client corrective action is to restart the service.                          |
	//	+---------------------------+----------------------------------------------------------------------------------+
	Action RecoverAction `idl:"name:action" json:"action"`
}

func (o *ServiceNotification) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ServiceNotification) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Event); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Action)); err != nil {
		return err
	}
	return nil
}
func (o *ServiceNotification) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Event); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Action)); err != nil {
		return err
	}
	return nil
}

// Notification structure represents VDS_NOTIFICATION RPC structure.
//
// The VDS_NOTIFICATION structure provides information about a notification.
type Notification struct {
	// objectType:  A value defined in the VDS_NOTIFICATION_TARGET_TYPE enumeration that
	// describes the type of notification.
	ObjectType   NotificationTargetType     `idl:"name:objectType" json:"object_type"`
	Notification *Notification_Notification `idl:"name:Notification;switch_is:objectType" json:"notification"`
}

func (o *Notification) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Notification) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.ObjectType)); err != nil {
		return err
	}
	_swNotification := uint16(o.ObjectType)
	if o.Notification != nil {
		if err := o.Notification.MarshalUnionNDR(ctx, w, _swNotification); err != nil {
			return err
		}
	} else {
		if err := (&Notification_Notification{}).MarshalUnionNDR(ctx, w, _swNotification); err != nil {
			return err
		}
	}
	return nil
}
func (o *Notification) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.ObjectType)); err != nil {
		return err
	}
	if o.Notification == nil {
		o.Notification = &Notification_Notification{}
	}
	_swNotification := uint16(o.ObjectType)
	if err := o.Notification.UnmarshalUnionNDR(ctx, w, _swNotification); err != nil {
		return err
	}
	return nil
}

// Notification_Notification structure represents VDS_NOTIFICATION union anonymous member.
//
// The VDS_NOTIFICATION structure provides information about a notification.
type Notification_Notification struct {
	// Types that are assignable to Value
	//
	// *Notification_Pack
	// *Notification_Disk
	// *Notification_Volume
	// *Notification_Partition
	// *Notification_Letter
	// *Notification_FileSystem
	// *Notification_MountPoint
	// *Notification_Service
	Value is_Notification_Notification `json:"value"`
}

func (o *Notification_Notification) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *Notification_Pack:
		if value != nil {
			return value.Pack
		}
	case *Notification_Disk:
		if value != nil {
			return value.Disk
		}
	case *Notification_Volume:
		if value != nil {
			return value.Volume
		}
	case *Notification_Partition:
		if value != nil {
			return value.Partition
		}
	case *Notification_Letter:
		if value != nil {
			return value.Letter
		}
	case *Notification_FileSystem:
		if value != nil {
			return value.FileSystem
		}
	case *Notification_MountPoint:
		if value != nil {
			return value.MountPoint
		}
	case *Notification_Service:
		if value != nil {
			return value.Service
		}
	}
	return nil
}

type is_Notification_Notification interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_Notification_Notification()
}

func (o *Notification_Notification) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *Notification_Pack:
		return uint16(10)
	case *Notification_Disk:
		return uint16(13)
	case *Notification_Volume:
		return uint16(11)
	case *Notification_Partition:
		return uint16(60)
	case *Notification_Letter:
		return uint16(61)
	case *Notification_FileSystem:
		return uint16(62)
	case *Notification_MountPoint:
		return uint16(63)
	case *Notification_Service:
		return uint16(200)
	}
	return uint16(0)
}

func (o *Notification_Notification) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(10):
		_o, _ := o.Value.(*Notification_Pack)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Notification_Pack{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(13):
		_o, _ := o.Value.(*Notification_Disk)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Notification_Disk{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(11):
		_o, _ := o.Value.(*Notification_Volume)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Notification_Volume{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(60):
		_o, _ := o.Value.(*Notification_Partition)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Notification_Partition{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(61):
		_o, _ := o.Value.(*Notification_Letter)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Notification_Letter{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(62):
		_o, _ := o.Value.(*Notification_FileSystem)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Notification_FileSystem{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(63):
		_o, _ := o.Value.(*Notification_MountPoint)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Notification_MountPoint{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(200):
		_o, _ := o.Value.(*Notification_Service)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Notification_Service{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *Notification_Notification) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch((*uint16)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(10):
		o.Value = &Notification_Pack{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(13):
		o.Value = &Notification_Disk{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(11):
		o.Value = &Notification_Volume{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(60):
		o.Value = &Notification_Partition{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(61):
		o.Value = &Notification_Letter{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(62):
		o.Value = &Notification_FileSystem{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(63):
		o.Value = &Notification_MountPoint{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(200):
		o.Value = &Notification_Service{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// Notification_Pack structure represents Notification_Notification RPC union arm.
//
// It has following labels: 10
type Notification_Pack struct {
	// Pack:  A VDS_PACK_NOTIFICATION structure that describes a pack change.
	Pack *PackNotification `idl:"name:Pack" json:"pack"`
}

func (*Notification_Pack) is_Notification_Notification() {}

func (o *Notification_Pack) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Pack != nil {
		if err := o.Pack.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PackNotification{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Notification_Pack) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Pack == nil {
		o.Pack = &PackNotification{}
	}
	if err := o.Pack.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Notification_Disk structure represents Notification_Notification RPC union arm.
//
// It has following labels: 13
type Notification_Disk struct {
	// Disk:  A VDS_DISK_NOTIFICATION structure that describes a disk change.
	Disk *DiskNotification `idl:"name:Disk" json:"disk"`
}

func (*Notification_Disk) is_Notification_Notification() {}

func (o *Notification_Disk) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Disk != nil {
		if err := o.Disk.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DiskNotification{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Notification_Disk) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Disk == nil {
		o.Disk = &DiskNotification{}
	}
	if err := o.Disk.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Notification_Volume structure represents Notification_Notification RPC union arm.
//
// It has following labels: 11
type Notification_Volume struct {
	// Volume:  A VDS_VOLUME_NOTIFICATION structure that describes a volume change.
	Volume *VolumeNotification `idl:"name:Volume" json:"volume"`
}

func (*Notification_Volume) is_Notification_Notification() {}

func (o *Notification_Volume) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Volume != nil {
		if err := o.Volume.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&VolumeNotification{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Notification_Volume) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Volume == nil {
		o.Volume = &VolumeNotification{}
	}
	if err := o.Volume.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Notification_Partition structure represents Notification_Notification RPC union arm.
//
// It has following labels: 60
type Notification_Partition struct {
	// Partition:  A VDS_PARTITION_NOTIFICATION structure that describes a partition change.
	Partition *PartitionNotification `idl:"name:Partition" json:"partition"`
}

func (*Notification_Partition) is_Notification_Notification() {}

func (o *Notification_Partition) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Partition != nil {
		if err := o.Partition.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PartitionNotification{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Notification_Partition) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Partition == nil {
		o.Partition = &PartitionNotification{}
	}
	if err := o.Partition.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Notification_Letter structure represents Notification_Notification RPC union arm.
//
// It has following labels: 61
type Notification_Letter struct {
	// Letter:  A VDS_DRIVE_LETTER_NOTIFICATION structure that describes a drive letter
	// change.
	Letter *DriveLetterNotification `idl:"name:Letter" json:"letter"`
}

func (*Notification_Letter) is_Notification_Notification() {}

func (o *Notification_Letter) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Letter != nil {
		if err := o.Letter.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DriveLetterNotification{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Notification_Letter) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Letter == nil {
		o.Letter = &DriveLetterNotification{}
	}
	if err := o.Letter.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Notification_FileSystem structure represents Notification_Notification RPC union arm.
//
// It has following labels: 62
type Notification_FileSystem struct {
	// FileSystem:  A VDS_FILE_SYSTEM_NOTIFICATION structure that describes a file system
	// change.
	FileSystem *FileSystemNotification `idl:"name:FileSystem" json:"file_system"`
}

func (*Notification_FileSystem) is_Notification_Notification() {}

func (o *Notification_FileSystem) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.FileSystem != nil {
		if err := o.FileSystem.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&FileSystemNotification{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Notification_FileSystem) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.FileSystem == nil {
		o.FileSystem = &FileSystemNotification{}
	}
	if err := o.FileSystem.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Notification_MountPoint structure represents Notification_Notification RPC union arm.
//
// It has following labels: 63
type Notification_MountPoint struct {
	// MountPoint:  A VDS_MOUNT_POINT_NOTIFICATION structure that describes a mount point
	// change.
	MountPoint *MountPointNotification `idl:"name:MountPoint" json:"mount_point"`
}

func (*Notification_MountPoint) is_Notification_Notification() {}

func (o *Notification_MountPoint) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.MountPoint != nil {
		if err := o.MountPoint.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&MountPointNotification{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Notification_MountPoint) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.MountPoint == nil {
		o.MountPoint = &MountPointNotification{}
	}
	if err := o.MountPoint.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Notification_Service structure represents Notification_Notification RPC union arm.
//
// It has following labels: 200
type Notification_Service struct {
	// Service:  A VDS_SERVICE_NOTIFICATION structure that provides information about a
	// state change to the service object.
	Service *ServiceNotification `idl:"name:Service" json:"service"`
}

func (*Notification_Service) is_Notification_Notification() {}

func (o *Notification_Service) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Service != nil {
		if err := o.Service.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ServiceNotification{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Notification_Service) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Service == nil {
		o.Service = &ServiceNotification{}
	}
	if err := o.Service.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// AsyncOutput structure represents VDS_ASYNC_OUTPUT RPC structure.
//
// The VDS_ASYNC_OUTPUT structure provides information from a completed asynchronous
// operation.
type AsyncOutput struct {
	// type:  A value from the VDS_ASYNC_OUTPUT_TYPE enumeration that indicates the type
	// of operation information.
	Type        AsyncOutputType          `idl:"name:type" json:"type"`
	AsyncOutput *AsyncOutput_AsyncOutput `idl:"name:async_output;switch_is:type" json:"async_output"`
}

func (o *AsyncOutput) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AsyncOutput) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Type)); err != nil {
		return err
	}
	_swAsyncOutput := uint16(o.Type)
	if o.AsyncOutput != nil {
		if err := o.AsyncOutput.MarshalUnionNDR(ctx, w, _swAsyncOutput); err != nil {
			return err
		}
	} else {
		if err := (&AsyncOutput_AsyncOutput{}).MarshalUnionNDR(ctx, w, _swAsyncOutput); err != nil {
			return err
		}
	}
	return nil
}
func (o *AsyncOutput) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Type)); err != nil {
		return err
	}
	if o.AsyncOutput == nil {
		o.AsyncOutput = &AsyncOutput_AsyncOutput{}
	}
	_swAsyncOutput := uint16(o.Type)
	if err := o.AsyncOutput.UnmarshalUnionNDR(ctx, w, _swAsyncOutput); err != nil {
		return err
	}
	return nil
}

// AsyncOutput_AsyncOutput structure represents VDS_ASYNC_OUTPUT union anonymous member.
//
// The VDS_ASYNC_OUTPUT structure provides information from a completed asynchronous
// operation.
type AsyncOutput_AsyncOutput struct {
	// Types that are assignable to Value
	//
	// *AsyncOutput_CreatePartition
	// *AsyncOutput_CreateVolume
	// *AsyncOutput_BreakVolumePlex
	// *AsyncOutput_ShrinkVolume
	// *AsyncOutput_CreateVDisk
	// *AsyncOutput_DefaultAsyncOutput
	Value is_AsyncOutput_AsyncOutput `json:"value"`
}

func (o *AsyncOutput_AsyncOutput) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *AsyncOutput_CreatePartition:
		if value != nil {
			return value.CreatePartition
		}
	case *AsyncOutput_CreateVolume:
		if value != nil {
			return value.CreateVolume
		}
	case *AsyncOutput_BreakVolumePlex:
		if value != nil {
			return value.BreakVolumePlex
		}
	case *AsyncOutput_ShrinkVolume:
		if value != nil {
			return value.ShrinkVolume
		}
	case *AsyncOutput_CreateVDisk:
		if value != nil {
			return value.CreateVDisk
		}
	}
	return nil
}

type is_AsyncOutput_AsyncOutput interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_AsyncOutput_AsyncOutput()
}

func (o *AsyncOutput_AsyncOutput) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *AsyncOutput_CreatePartition:
		return uint16(10)
	case *AsyncOutput_CreateVolume:
		return uint16(1)
	case *AsyncOutput_BreakVolumePlex:
		return uint16(5)
	case *AsyncOutput_ShrinkVolume:
		return uint16(3)
	case *AsyncOutput_CreateVDisk:
		return uint16(200)
	}
	return uint16(0)
}

func (o *AsyncOutput_AsyncOutput) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(10):
		_o, _ := o.Value.(*AsyncOutput_CreatePartition)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&AsyncOutput_CreatePartition{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(1):
		_o, _ := o.Value.(*AsyncOutput_CreateVolume)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&AsyncOutput_CreateVolume{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(5):
		_o, _ := o.Value.(*AsyncOutput_BreakVolumePlex)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&AsyncOutput_BreakVolumePlex{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(3):
		_o, _ := o.Value.(*AsyncOutput_ShrinkVolume)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&AsyncOutput_ShrinkVolume{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(200):
		_o, _ := o.Value.(*AsyncOutput_CreateVDisk)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&AsyncOutput_CreateVDisk{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		// no-op
	}
	return nil
}

func (o *AsyncOutput_AsyncOutput) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch((*uint16)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(10):
		o.Value = &AsyncOutput_CreatePartition{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(1):
		o.Value = &AsyncOutput_CreateVolume{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(5):
		o.Value = &AsyncOutput_BreakVolumePlex{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(3):
		o.Value = &AsyncOutput_ShrinkVolume{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(200):
		o.Value = &AsyncOutput_CreateVDisk{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		o.Value = &AsyncOutput_DefaultAsyncOutput{}
	}
	return nil
}

// AsyncOutput_CreatePartition structure represents AsyncOutput_AsyncOutput RPC union arm.
//
// It has following labels: 10
type AsyncOutput_CreatePartition struct {
	// cp:  The cp structure provides information about a newly created partition.
	CreatePartition *AsyncOutput_AsyncOutput_CreatePartition `idl:"name:cp" json:"create_partition"`
}

func (*AsyncOutput_CreatePartition) is_AsyncOutput_AsyncOutput() {}

func (o *AsyncOutput_CreatePartition) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.CreatePartition != nil {
		if err := o.CreatePartition.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&AsyncOutput_AsyncOutput_CreatePartition{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *AsyncOutput_CreatePartition) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.CreatePartition == nil {
		o.CreatePartition = &AsyncOutput_AsyncOutput_CreatePartition{}
	}
	if err := o.CreatePartition.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// AsyncOutput_AsyncOutput_CreatePartition structure represents VDS_ASYNC_OUTPUT structure anonymous member.
//
// The VDS_ASYNC_OUTPUT structure provides information from a completed asynchronous
// operation.
type AsyncOutput_AsyncOutput_CreatePartition struct {
	// ullOffset:  The byte offset of the partition from the beginning of the disk.
	Offset uint64 `idl:"name:ullOffset" json:"offset"`
	// volumeId:  The VDS object ID of the associated volume object, if the partition is
	// a volume.
	VolumeID *ObjectID `idl:"name:volumeId" json:"volume_id"`
}

func (o *AsyncOutput_AsyncOutput_CreatePartition) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AsyncOutput_AsyncOutput_CreatePartition) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.Offset); err != nil {
		return err
	}
	if o.VolumeID != nil {
		if err := o.VolumeID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *AsyncOutput_AsyncOutput_CreatePartition) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.Offset); err != nil {
		return err
	}
	if o.VolumeID == nil {
		o.VolumeID = &ObjectID{}
	}
	if err := o.VolumeID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// AsyncOutput_CreateVolume structure represents AsyncOutput_AsyncOutput RPC union arm.
//
// It has following labels: 1
type AsyncOutput_CreateVolume struct {
	// cv:  The cv structure provides information about a newly created volume.
	CreateVolume *AsyncOutput_AsyncOutput_CreateVolume `idl:"name:cv" json:"create_volume"`
}

func (*AsyncOutput_CreateVolume) is_AsyncOutput_AsyncOutput() {}

func (o *AsyncOutput_CreateVolume) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.CreateVolume != nil {
		if err := o.CreateVolume.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&AsyncOutput_AsyncOutput_CreateVolume{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *AsyncOutput_CreateVolume) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.CreateVolume == nil {
		o.CreateVolume = &AsyncOutput_AsyncOutput_CreateVolume{}
	}
	if err := o.CreateVolume.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// AsyncOutput_AsyncOutput_CreateVolume structure represents VDS_ASYNC_OUTPUT structure anonymous member.
//
// The VDS_ASYNC_OUTPUT structure provides information from a completed asynchronous
// operation.
type AsyncOutput_AsyncOutput_CreateVolume struct {
	// pVolumeUnk:  A pointer to the IUnknown interface of the newly created volume.
	VolumeUnknown *dcom.Unknown `idl:"name:pVolumeUnk" json:"volume_unknown"`
}

func (o *AsyncOutput_AsyncOutput_CreateVolume) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AsyncOutput_AsyncOutput_CreateVolume) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.VolumeUnknown != nil {
		_ptr_pVolumeUnk := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.VolumeUnknown != nil {
				if err := o.VolumeUnknown.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dcom.Unknown{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.VolumeUnknown, _ptr_pVolumeUnk); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *AsyncOutput_AsyncOutput_CreateVolume) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_pVolumeUnk := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.VolumeUnknown == nil {
			o.VolumeUnknown = &dcom.Unknown{}
		}
		if err := o.VolumeUnknown.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pVolumeUnk := func(ptr interface{}) { o.VolumeUnknown = *ptr.(**dcom.Unknown) }
	if err := w.ReadPointer(&o.VolumeUnknown, _s_pVolumeUnk, _ptr_pVolumeUnk); err != nil {
		return err
	}
	return nil
}

// AsyncOutput_BreakVolumePlex structure represents AsyncOutput_AsyncOutput RPC union arm.
//
// It has following labels: 5
type AsyncOutput_BreakVolumePlex struct {
	// bvp:  The bvp structure provides information about a volume after a plex is broken.
	BreakVolumePlex *AsyncOutput_AsyncOutput_BreakVolumePlex `idl:"name:bvp" json:"break_volume_plex"`
}

func (*AsyncOutput_BreakVolumePlex) is_AsyncOutput_AsyncOutput() {}

func (o *AsyncOutput_BreakVolumePlex) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.BreakVolumePlex != nil {
		if err := o.BreakVolumePlex.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&AsyncOutput_AsyncOutput_BreakVolumePlex{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *AsyncOutput_BreakVolumePlex) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.BreakVolumePlex == nil {
		o.BreakVolumePlex = &AsyncOutput_AsyncOutput_BreakVolumePlex{}
	}
	if err := o.BreakVolumePlex.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// AsyncOutput_AsyncOutput_BreakVolumePlex structure represents VDS_ASYNC_OUTPUT structure anonymous member.
//
// The VDS_ASYNC_OUTPUT structure provides information from a completed asynchronous
// operation.
type AsyncOutput_AsyncOutput_BreakVolumePlex struct {
	// pVolumeUnk:  A pointer to the IUnknown interface of the newly created volume.
	VolumeUnknown *dcom.Unknown `idl:"name:pVolumeUnk" json:"volume_unknown"`
}

func (o *AsyncOutput_AsyncOutput_BreakVolumePlex) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AsyncOutput_AsyncOutput_BreakVolumePlex) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.VolumeUnknown != nil {
		_ptr_pVolumeUnk := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.VolumeUnknown != nil {
				if err := o.VolumeUnknown.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dcom.Unknown{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.VolumeUnknown, _ptr_pVolumeUnk); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *AsyncOutput_AsyncOutput_BreakVolumePlex) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_pVolumeUnk := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.VolumeUnknown == nil {
			o.VolumeUnknown = &dcom.Unknown{}
		}
		if err := o.VolumeUnknown.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pVolumeUnk := func(ptr interface{}) { o.VolumeUnknown = *ptr.(**dcom.Unknown) }
	if err := w.ReadPointer(&o.VolumeUnknown, _s_pVolumeUnk, _ptr_pVolumeUnk); err != nil {
		return err
	}
	return nil
}

// AsyncOutput_ShrinkVolume structure represents AsyncOutput_AsyncOutput RPC union arm.
//
// It has following labels: 3
type AsyncOutput_ShrinkVolume struct {
	// sv:  The sv structure provides information about a volume shrink operation.
	ShrinkVolume *AsyncOutput_AsyncOutput_ShrinkVolume `idl:"name:sv" json:"shrink_volume"`
}

func (*AsyncOutput_ShrinkVolume) is_AsyncOutput_AsyncOutput() {}

func (o *AsyncOutput_ShrinkVolume) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ShrinkVolume != nil {
		if err := o.ShrinkVolume.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&AsyncOutput_AsyncOutput_ShrinkVolume{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *AsyncOutput_ShrinkVolume) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ShrinkVolume == nil {
		o.ShrinkVolume = &AsyncOutput_AsyncOutput_ShrinkVolume{}
	}
	if err := o.ShrinkVolume.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// AsyncOutput_AsyncOutput_ShrinkVolume structure represents VDS_ASYNC_OUTPUT structure anonymous member.
//
// The VDS_ASYNC_OUTPUT structure provides information from a completed asynchronous
// operation.
type AsyncOutput_AsyncOutput_ShrinkVolume struct {
	// ullReclaimedBytes:  The number of bytes that the volume shrink operation reclaimed.
	ReclaimedBytes uint64 `idl:"name:ullReclaimedBytes" json:"reclaimed_bytes"`
}

func (o *AsyncOutput_AsyncOutput_ShrinkVolume) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AsyncOutput_AsyncOutput_ShrinkVolume) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.ReclaimedBytes); err != nil {
		return err
	}
	return nil
}
func (o *AsyncOutput_AsyncOutput_ShrinkVolume) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReclaimedBytes); err != nil {
		return err
	}
	return nil
}

// AsyncOutput_CreateVDisk structure represents AsyncOutput_AsyncOutput RPC union arm.
//
// It has following labels: 200
type AsyncOutput_CreateVDisk struct {
	// cvd:  The cvd structure provides information about a newly created virtual disk.
	CreateVDisk *AsyncOutput_AsyncOutput_CreateVDisk `idl:"name:cvd" json:"create_v_disk"`
}

func (*AsyncOutput_CreateVDisk) is_AsyncOutput_AsyncOutput() {}

func (o *AsyncOutput_CreateVDisk) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.CreateVDisk != nil {
		if err := o.CreateVDisk.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&AsyncOutput_AsyncOutput_CreateVDisk{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *AsyncOutput_CreateVDisk) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.CreateVDisk == nil {
		o.CreateVDisk = &AsyncOutput_AsyncOutput_CreateVDisk{}
	}
	if err := o.CreateVDisk.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// AsyncOutput_AsyncOutput_CreateVDisk structure represents VDS_ASYNC_OUTPUT structure anonymous member.
//
// The VDS_ASYNC_OUTPUT structure provides information from a completed asynchronous
// operation.
type AsyncOutput_AsyncOutput_CreateVDisk struct {
	// pVDiskUnk:  A pointer to the IUnknown interface of the newly created virtual disk.
	VDiskUnknown *dcom.Unknown `idl:"name:pVDiskUnk" json:"v_disk_unknown"`
}

func (o *AsyncOutput_AsyncOutput_CreateVDisk) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AsyncOutput_AsyncOutput_CreateVDisk) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.VDiskUnknown != nil {
		_ptr_pVDiskUnk := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.VDiskUnknown != nil {
				if err := o.VDiskUnknown.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dcom.Unknown{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.VDiskUnknown, _ptr_pVDiskUnk); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *AsyncOutput_AsyncOutput_CreateVDisk) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_pVDiskUnk := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.VDiskUnknown == nil {
			o.VDiskUnknown = &dcom.Unknown{}
		}
		if err := o.VDiskUnknown.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pVDiskUnk := func(ptr interface{}) { o.VDiskUnknown = *ptr.(**dcom.Unknown) }
	if err := w.ReadPointer(&o.VDiskUnknown, _s_pVDiskUnk, _ptr_pVDiskUnk); err != nil {
		return err
	}
	return nil
}

// AsyncOutput_DefaultAsyncOutput structure represents AsyncOutput_AsyncOutput RPC default union arm.
type AsyncOutput_DefaultAsyncOutput struct {
}

func (*AsyncOutput_DefaultAsyncOutput) is_AsyncOutput_AsyncOutput() {}

func (o *AsyncOutput_DefaultAsyncOutput) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return nil
}
func (o *AsyncOutput_DefaultAsyncOutput) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	return nil
}

// PartitionInfoMBR structure represents VDS_PARTITION_INFO_MBR RPC structure.
//
// The VDS_PARTITION_INFO_MBR structure provides information about an MBR partition.
type PartitionInfoMBR struct {
	// partitionType:  The byte value indicating the partition type.<20>
	PartitionType uint8 `idl:"name:partitionType" json:"partition_type"`
	// bootIndicator:  A Boolean value that indicates whether the partition is bootable.
	BootIndicator bool `idl:"name:bootIndicator" json:"boot_indicator"`
	// recognizedPartition:  A Boolean value that indicates whether the partition will be
	// exposed as a volume.
	RecognizedPartition bool `idl:"name:recognizedPartition" json:"recognized_partition"`
	// hiddenSectors:  The number of sectors between the start of the partition and the
	// partition's first usable area.
	HiddenSectors uint32 `idl:"name:hiddenSectors" json:"hidden_sectors"`
}

func (o *PartitionInfoMBR) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PartitionInfoMBR) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.PartitionType); err != nil {
		return err
	}
	if err := w.WriteData(o.BootIndicator); err != nil {
		return err
	}
	if err := w.WriteData(o.RecognizedPartition); err != nil {
		return err
	}
	if err := w.WriteData(o.HiddenSectors); err != nil {
		return err
	}
	return nil
}
func (o *PartitionInfoMBR) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.PartitionType); err != nil {
		return err
	}
	if err := w.ReadData(&o.BootIndicator); err != nil {
		return err
	}
	if err := w.ReadData(&o.RecognizedPartition); err != nil {
		return err
	}
	if err := w.ReadData(&o.HiddenSectors); err != nil {
		return err
	}
	return nil
}

// PartitionInfoGPT structure represents VDS_PARTITION_INFO_GPT RPC structure.
//
// The VDS_PARTITION_INFO_GPT structure provides information about a partition in a
// GPT.
type PartitionInfoGPT struct {
	// partitionType:  A GUID indicating the partition type.<21>
	PartitionType *dtyp.GUID `idl:"name:partitionType" json:"partition_type"`
	// partitionId:  The GUID of the partition.
	PartitionID *dtyp.GUID `idl:"name:partitionId" json:"partition_id"`
	// attributes:  The attributes of the partition; they can have a combination of the
	// following values.
	//
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                                             |                                                                                  |
	//	|                            VALUE                            |                                     MEANING                                      |
	//	|                                                             |                                                                                  |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| GPT_ATTRIBUTE_PLATFORM_REQUIRED 0x0000000000000001          | Partition is required for the platform to function properly.<22>                 |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| GPT_BASIC_DATA_ATTRIBUTE_READ_ONLY 0x1000000000000000       | Partition cannot be written to but can be read from. Used only with the basic    |
	//	|                                                             | data partition type.                                                             |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| GPT_BASIC_DATA_ATTRIBUTE_SHADOW_COPY 0x2000000000000000     | Partition is a shadow copy. Used only with the basic data partition type.        |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| GPT_BASIC_DATA_ATTRIBUTE_HIDDEN 0x4000000000000000          | Partition is hidden and will not be mounted. Used only with the basic data       |
	//	|                                                             | partition type.                                                                  |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| GPT_BASIC_DATA_ATTRIBUTE_NO_DRIVE_LETTER 0x8000000000000000 | Partition does not receive a drive letter by default when moving the disk to     |
	//	|                                                             | another machine. Used only with the basic data partition type.                   |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	Attributes uint64 `idl:"name:attributes" json:"attributes"`
	// name:  Null-terminated Unicode name of the partition.
	Name []uint16 `idl:"name:name" json:"name"`
}

func (o *PartitionInfoGPT) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PartitionInfoGPT) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.PartitionType != nil {
		if err := o.PartitionType.MarshalNDR(ctx, w); err != nil {
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
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	for i1 := range o.Name {
		i1 := i1
		if uint64(i1) >= 36 {
			break
		}
		if err := w.WriteData(o.Name[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Name); uint64(i1) < 36; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *PartitionInfoGPT) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.PartitionType == nil {
		o.PartitionType = &dtyp.GUID{}
	}
	if err := o.PartitionType.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.PartitionID == nil {
		o.PartitionID = &dtyp.GUID{}
	}
	if err := o.PartitionID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	o.Name = make([]uint16, 36)
	for i1 := range o.Name {
		i1 := i1
		if err := w.ReadData(&o.Name[i1]); err != nil {
			return err
		}
	}
	return nil
}

// StorageID structure represents VDS_STORAGE_IDENTIFIER RPC structure.
//
// The VDS_STORAGE_IDENTIFIER structure provides information about a storage identifier.
type StorageID struct {
	// m_CodeSet:  Value from the VDS_STORAGE_IDENTIFIER_CODE_SET enumeration that defines
	// the code set of the storage identifier.
	CodeSet StorageIDCodeSet `idl:"name:m_CodeSet" json:"code_set"`
	// m_Type:  Value from the VDS_STORAGE_IDENTIFIER_TYPE enumeration that defines the
	// type of the storage identifier.
	Type StorageIDType `idl:"name:m_Type" json:"type"`
	// m_cbIdentifier:  Length of the m_rgbIdentifier identifier in bytes.
	IDLength uint32 `idl:"name:m_cbIdentifier" json:"id_length"`
	// m_rgbIdentifier:  Value of the storage identifier. These identifiers depend on both
	// the code set and the type.
	ID []byte `idl:"name:m_rgbIdentifier;size_is:(m_cbIdentifier)" json:"id"`
}

func (o *StorageID) xxx_PreparePayload(ctx context.Context) error {
	if o.ID != nil && o.IDLength == 0 {
		o.IDLength = uint32(len(o.ID))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *StorageID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.CodeSet)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Type)); err != nil {
		return err
	}
	if err := w.WriteData(o.IDLength); err != nil {
		return err
	}
	if o.ID != nil || o.IDLength > 0 {
		_ptr_m_rgbIdentifier := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.IDLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ID {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.ID[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.ID); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ID, _ptr_m_rgbIdentifier); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *StorageID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.CodeSet)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Type)); err != nil {
		return err
	}
	if err := w.ReadData(&o.IDLength); err != nil {
		return err
	}
	_ptr_m_rgbIdentifier := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.IDLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.IDLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ID", sizeInfo[0])
		}
		o.ID = make([]byte, sizeInfo[0])
		for i1 := range o.ID {
			i1 := i1
			if err := w.ReadData(&o.ID[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_m_rgbIdentifier := func(ptr interface{}) { o.ID = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.ID, _s_m_rgbIdentifier, _ptr_m_rgbIdentifier); err != nil {
		return err
	}
	return nil
}

// StorageDeviceIDDescriptor structure represents VDS_STORAGE_DEVICE_ID_DESCRIPTOR RPC structure.
//
// The VDS_STORAGE_DEVICE_ID_DESCRIPTOR structure provides information about a device
// identification descriptor.
type StorageDeviceIDDescriptor struct {
	// m_version:  The version number of the VDS_STORAGE_DEVICE_ID_DESCRIPTOR structure
	// as specified by the device manufacturer and in [SPC-3].
	Version uint32 `idl:"name:m_version" json:"version"`
	// m_cIdentifiers:  The number of elements in the m_rgIdentifiers array.
	IdentifiersCount uint32 `idl:"name:m_cIdentifiers" json:"identifiers_count"`
	// m_rgIdentifiers:  The array of VDS_STORAGE_IDENTIFIER structures that contain the
	// storage identifier information.
	Identifiers []*StorageID `idl:"name:m_rgIdentifiers;size_is:(m_cIdentifiers)" json:"identifiers"`
}

func (o *StorageDeviceIDDescriptor) xxx_PreparePayload(ctx context.Context) error {
	if o.Identifiers != nil && o.IdentifiersCount == 0 {
		o.IdentifiersCount = uint32(len(o.Identifiers))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *StorageDeviceIDDescriptor) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	if err := w.WriteData(o.IdentifiersCount); err != nil {
		return err
	}
	if o.Identifiers != nil || o.IdentifiersCount > 0 {
		_ptr_m_rgIdentifiers := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.IdentifiersCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Identifiers {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Identifiers[i1] != nil {
					if err := o.Identifiers[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&StorageID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Identifiers); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&StorageID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Identifiers, _ptr_m_rgIdentifiers); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *StorageDeviceIDDescriptor) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	if err := w.ReadData(&o.IdentifiersCount); err != nil {
		return err
	}
	_ptr_m_rgIdentifiers := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.IdentifiersCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.IdentifiersCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Identifiers", sizeInfo[0])
		}
		o.Identifiers = make([]*StorageID, sizeInfo[0])
		for i1 := range o.Identifiers {
			i1 := i1
			if o.Identifiers[i1] == nil {
				o.Identifiers[i1] = &StorageID{}
			}
			if err := o.Identifiers[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_m_rgIdentifiers := func(ptr interface{}) { o.Identifiers = *ptr.(*[]*StorageID) }
	if err := w.ReadPointer(&o.Identifiers, _s_m_rgIdentifiers, _ptr_m_rgIdentifiers); err != nil {
		return err
	}
	return nil
}

// Interconnect structure represents VDS_INTERCONNECT RPC structure.
//
// The VDS_INTERCONNECT structure defines the address data of a physical interconnect,
// as specified in [SPC-3].
type Interconnect struct {
	// m_addressType:  A VDS_INTERCONNECT_ADDRESS_TYPE structure that stores the address
	// type of the interconnect.
	AddressType InterconnectAddressType `idl:"name:m_addressType" json:"address_type"`
	// m_cbPort:  The size, in bytes, of the interconnect address data for the LUN port
	// to which m_pbPort refers.
	PortLength uint32 `idl:"name:m_cbPort" json:"port_length"`
	// m_pbPort:  A pointer to the interconnect address data for the LUN port.
	Port []byte `idl:"name:m_pbPort;size_is:(m_cbPort)" json:"port"`
	// m_cbAddress:  The size, in bytes, of the interconnect address data for the LUN to
	// which m_pbAddress refers.
	AddressLength uint32 `idl:"name:m_cbAddress" json:"address_length"`
	// m_pbAddress:  A pointer to the interconnect address data for the LUN.
	Address []byte `idl:"name:m_pbAddress;size_is:(m_cbAddress)" json:"address"`
}

func (o *Interconnect) xxx_PreparePayload(ctx context.Context) error {
	if o.Port != nil && o.PortLength == 0 {
		o.PortLength = uint32(len(o.Port))
	}
	if o.Address != nil && o.AddressLength == 0 {
		o.AddressLength = uint32(len(o.Address))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Interconnect) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.AddressType)); err != nil {
		return err
	}
	if err := w.WriteData(o.PortLength); err != nil {
		return err
	}
	if o.Port != nil || o.PortLength > 0 {
		_ptr_m_pbPort := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.PortLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Port {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Port[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Port); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Port, _ptr_m_pbPort); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.AddressLength); err != nil {
		return err
	}
	if o.Address != nil || o.AddressLength > 0 {
		_ptr_m_pbAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.AddressLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Address {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Address[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Address); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Address, _ptr_m_pbAddress); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Interconnect) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.AddressType)); err != nil {
		return err
	}
	if err := w.ReadData(&o.PortLength); err != nil {
		return err
	}
	_ptr_m_pbPort := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.PortLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.PortLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Port", sizeInfo[0])
		}
		o.Port = make([]byte, sizeInfo[0])
		for i1 := range o.Port {
			i1 := i1
			if err := w.ReadData(&o.Port[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_m_pbPort := func(ptr interface{}) { o.Port = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Port, _s_m_pbPort, _ptr_m_pbPort); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddressLength); err != nil {
		return err
	}
	_ptr_m_pbAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.AddressLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.AddressLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Address", sizeInfo[0])
		}
		o.Address = make([]byte, sizeInfo[0])
		for i1 := range o.Address {
			i1 := i1
			if err := w.ReadData(&o.Address[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_m_pbAddress := func(ptr interface{}) { o.Address = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Address, _s_m_pbAddress, _ptr_m_pbAddress); err != nil {
		return err
	}
	return nil
}

// LUNInformation structure represents VDS_LUN_INFORMATION RPC structure.
//
// The VDS_LUN_INFORMATION structure provides information about a SCSI-2 device. For
// information about a SCSI-2 device, see [ANSI-131-1994].
type LUNInformation struct {
	// m_version:  The version number of the VDS_LUN_INFORMATION structure. As of the current
	// version of this protocol, this value is always VER_VDS_LUN_INFORMATION (0x00000001).
	Version uint32 `idl:"name:m_version" json:"version"`
	// m_DeviceType:  The SCSI-2 device type of the device, as specified in [SPC-3].
	DeviceType uint8 `idl:"name:m_DeviceType" json:"device_type"`
	// m_DeviceTypeModifier:  The SCSI-2 device type modifier, if any, as specified in [SPC-3].
	DeviceTypeModifier uint8 `idl:"name:m_DeviceTypeModifier" json:"device_type_modifier"`
	// m_bCommandQueuing:  A Boolean value that indicates whether the device supports multiple
	// outstanding commands.
	CommandQueuing int32 `idl:"name:m_bCommandQueuing" json:"command_queuing"`
	// m_BusType:  A value from the VDS_STORAGE_BUS_TYPE enumeration that indicates the
	// bus type of the device.
	BusType StorageBusType `idl:"name:m_BusType" json:"bus_type"`
	// m_szVendorId:  The null-terminated vendor identification Unicode string of the device.
	// This value is NULL if no vendor ID exists.
	VendorID string `idl:"name:m_szVendorId;string" json:"vendor_id"`
	// m_szProductId:  The null-terminated product identification Unicode string of the
	// device. This value is NULL if no product ID exists.
	ProductID string `idl:"name:m_szProductId;string" json:"product_id"`
	// m_szProductRevision:  The null-terminated product revision Unicode string of the
	// device. This value is NULL if no product revision information exists.
	ProductRevision string `idl:"name:m_szProductRevision;string" json:"product_revision"`
	// m_szSerialNumber:  The null-terminated serial number of the device. This value is
	// NULL if no serial number exists.
	SerialNumber string `idl:"name:m_szSerialNumber;string" json:"serial_number"`
	// m_diskSignature:  The disk signature of the disk.
	DiskSignature *dtyp.GUID `idl:"name:m_diskSignature" json:"disk_signature"`
	// m_deviceIdDescriptor:  A VDS_STORAGE_DEVICE_ID_DESCRIPTOR structure that contains
	// the identification descriptor of the device.
	DeviceIDDescriptor *StorageDeviceIDDescriptor `idl:"name:m_deviceIdDescriptor" json:"device_id_descriptor"`
	// m_cInterconnects:  The number of elements in the m_rgInterconnects array.
	InterconnectsCount uint32 `idl:"name:m_cInterconnects" json:"interconnects_count"`
	// m_rgInterconnects:  Any array of VDS_INTERCONNECT structures that describe the physical
	// interconnects to the device.
	Interconnects []*Interconnect `idl:"name:m_rgInterconnects;size_is:(m_cInterconnects)" json:"interconnects"`
}

func (o *LUNInformation) xxx_PreparePayload(ctx context.Context) error {
	if o.Interconnects != nil && o.InterconnectsCount == 0 {
		o.InterconnectsCount = uint32(len(o.Interconnects))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *LUNInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	if err := w.WriteData(o.DeviceType); err != nil {
		return err
	}
	if err := w.WriteData(o.DeviceTypeModifier); err != nil {
		return err
	}
	if err := w.WriteData(o.CommandQueuing); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.BusType)); err != nil {
		return err
	}
	if o.VendorID != "" {
		_ptr_m_szVendorId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.VendorID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.VendorID, _ptr_m_szVendorId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ProductID != "" {
		_ptr_m_szProductId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.ProductID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ProductID, _ptr_m_szProductId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ProductRevision != "" {
		_ptr_m_szProductRevision := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.ProductRevision); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ProductRevision, _ptr_m_szProductRevision); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.SerialNumber != "" {
		_ptr_m_szSerialNumber := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.SerialNumber); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SerialNumber, _ptr_m_szSerialNumber); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DiskSignature != nil {
		if err := o.DiskSignature.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.DeviceIDDescriptor != nil {
		if err := o.DeviceIDDescriptor.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&StorageDeviceIDDescriptor{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.InterconnectsCount); err != nil {
		return err
	}
	if o.Interconnects != nil || o.InterconnectsCount > 0 {
		_ptr_m_rgInterconnects := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.InterconnectsCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Interconnects {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Interconnects[i1] != nil {
					if err := o.Interconnects[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&Interconnect{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Interconnects); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&Interconnect{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Interconnects, _ptr_m_rgInterconnects); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *LUNInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	if err := w.ReadData(&o.DeviceType); err != nil {
		return err
	}
	if err := w.ReadData(&o.DeviceTypeModifier); err != nil {
		return err
	}
	if err := w.ReadData(&o.CommandQueuing); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.BusType)); err != nil {
		return err
	}
	_ptr_m_szVendorId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.VendorID); err != nil {
			return err
		}
		return nil
	})
	_s_m_szVendorId := func(ptr interface{}) { o.VendorID = *ptr.(*string) }
	if err := w.ReadPointer(&o.VendorID, _s_m_szVendorId, _ptr_m_szVendorId); err != nil {
		return err
	}
	_ptr_m_szProductId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.ProductID); err != nil {
			return err
		}
		return nil
	})
	_s_m_szProductId := func(ptr interface{}) { o.ProductID = *ptr.(*string) }
	if err := w.ReadPointer(&o.ProductID, _s_m_szProductId, _ptr_m_szProductId); err != nil {
		return err
	}
	_ptr_m_szProductRevision := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.ProductRevision); err != nil {
			return err
		}
		return nil
	})
	_s_m_szProductRevision := func(ptr interface{}) { o.ProductRevision = *ptr.(*string) }
	if err := w.ReadPointer(&o.ProductRevision, _s_m_szProductRevision, _ptr_m_szProductRevision); err != nil {
		return err
	}
	_ptr_m_szSerialNumber := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.SerialNumber); err != nil {
			return err
		}
		return nil
	})
	_s_m_szSerialNumber := func(ptr interface{}) { o.SerialNumber = *ptr.(*string) }
	if err := w.ReadPointer(&o.SerialNumber, _s_m_szSerialNumber, _ptr_m_szSerialNumber); err != nil {
		return err
	}
	if o.DiskSignature == nil {
		o.DiskSignature = &dtyp.GUID{}
	}
	if err := o.DiskSignature.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.DeviceIDDescriptor == nil {
		o.DeviceIDDescriptor = &StorageDeviceIDDescriptor{}
	}
	if err := o.DeviceIDDescriptor.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterconnectsCount); err != nil {
		return err
	}
	_ptr_m_rgInterconnects := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.InterconnectsCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.InterconnectsCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Interconnects", sizeInfo[0])
		}
		o.Interconnects = make([]*Interconnect, sizeInfo[0])
		for i1 := range o.Interconnects {
			i1 := i1
			if o.Interconnects[i1] == nil {
				o.Interconnects[i1] = &Interconnect{}
			}
			if err := o.Interconnects[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_m_rgInterconnects := func(ptr interface{}) { o.Interconnects = *ptr.(*[]*Interconnect) }
	if err := w.ReadPointer(&o.Interconnects, _s_m_rgInterconnects, _ptr_m_rgInterconnects); err != nil {
		return err
	}
	return nil
}

// FileSystemProperty structure represents VDS_FILE_SYSTEM_PROP RPC structure.
//
// The VDS_FILE_SYSTEM_PROP structure provides information about the properties of a
// file system.
type FileSystemProperty struct {
	// type:  A VDS_FILE_SYSTEM_TYPE value that provides information about the type of the
	// file system.
	Type FileSystemType `idl:"name:type" json:"type"`
	// volumeId:  The VDS object ID of the volume object on which the file system resides.
	VolumeID *ObjectID `idl:"name:volumeId" json:"volume_id"`
	// ulFlags:  The combination of any values, by using the bitwise OR operator, that are
	// defined in the VDS_FILE_SYSTEM_PROP_FLAG enumeration.
	Flags uint32 `idl:"name:ulFlags" json:"flags"`
	// ullTotalAllocationUnits:  The total number of allocation units on the file system.
	TotalAllocationUnits uint64 `idl:"name:ullTotalAllocationUnits" json:"total_allocation_units"`
	// ullAvailableAllocationUnits:  The number of allocation units available on the file
	// system.
	AvailableAllocationUnits uint64 `idl:"name:ullAvailableAllocationUnits" json:"available_allocation_units"`
	// ulAllocationUnitSize:  The size of the allocation units in use by the file system.
	AllocationUnitSize uint32 `idl:"name:ulAllocationUnitSize" json:"allocation_unit_size"`
	// pwszLabel:  A null-terminated Unicode label of the file system.
	Label string `idl:"name:pwszLabel;string" json:"label"`
}

func (o *FileSystemProperty) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *FileSystemProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Type)); err != nil {
		return err
	}
	if o.VolumeID != nil {
		if err := o.VolumeID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalAllocationUnits); err != nil {
		return err
	}
	if err := w.WriteData(o.AvailableAllocationUnits); err != nil {
		return err
	}
	if err := w.WriteData(o.AllocationUnitSize); err != nil {
		return err
	}
	if o.Label != "" {
		_ptr_pwszLabel := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Label); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Label, _ptr_pwszLabel); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *FileSystemProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Type)); err != nil {
		return err
	}
	if o.VolumeID == nil {
		o.VolumeID = &ObjectID{}
	}
	if err := o.VolumeID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalAllocationUnits); err != nil {
		return err
	}
	if err := w.ReadData(&o.AvailableAllocationUnits); err != nil {
		return err
	}
	if err := w.ReadData(&o.AllocationUnitSize); err != nil {
		return err
	}
	_ptr_pwszLabel := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Label); err != nil {
			return err
		}
		return nil
	})
	_s_pwszLabel := func(ptr interface{}) { o.Label = *ptr.(*string) }
	if err := w.ReadPointer(&o.Label, _s_pwszLabel, _ptr_pwszLabel); err != nil {
		return err
	}
	return nil
}

// FileSystemFormatSupportProperty structure represents VDS_FILE_SYSTEM_FORMAT_SUPPORT_PROP RPC structure.
//
// The VDS_FILE_SYSTEM_FORMAT_SUPPORT_PROP structure provides information about file
// systems that are supported for formatting volumes.<23>
type FileSystemFormatSupportProperty struct {
	// ulFlags:  The combination of any values, by using the bitwise OR operator, that are
	// defined in the VDS_FILE_SYSTEM_FORMAT_SUPPORT_FLAG enumeration.
	Flags uint32 `idl:"name:ulFlags" json:"flags"`
	// usRevision:  A 16-bit, binary-coded decimal number that indicates the file system
	// version, if any. The first two (most significant) digits (8-bits) indicate the major
	// version while the last two (least significant) digits (8-bits) indicate the minor
	// version. For example, a value that has a bit pattern of 00000010 01010000 (0x0250
	// in hexadecimal) represents version 2.50; 0x1195 represents version 11.95, and so
	// on.
	Revision uint16 `idl:"name:usRevision" json:"revision"`
	// ulDefaultUnitAllocationSize:  The default allocation unit size, in bytes, that the
	// file system uses for formatting the volume. This value MUST be a power of 2 and MUST
	// also appear in rgulAllowedUnitAllocationSizes.
	DefaultUnitAllocationSize uint32 `idl:"name:ulDefaultUnitAllocationSize" json:"default_unit_allocation_size"`
	// rgulAllowedUnitAllocationSizes:  A zero-terminated array of allocation unit sizes,
	// in bytes, that the file system supports for formatting the volume. An array is not
	// zero-terminated if the array contains 32 elements. Each of the values in the array
	// MUST be a power of 2.
	AllowedUnitAllocationSizes []uint32 `idl:"name:rgulAllowedUnitAllocationSizes" json:"allowed_unit_allocation_sizes"`
	// wszName:  A null-terminated Unicode wide-character string that indicates the name
	// of the file system.
	Name []uint16 `idl:"name:wszName" json:"name"`
}

func (o *FileSystemFormatSupportProperty) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *FileSystemFormatSupportProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.Revision); err != nil {
		return err
	}
	if err := w.WriteData(o.DefaultUnitAllocationSize); err != nil {
		return err
	}
	for i1 := range o.AllowedUnitAllocationSizes {
		i1 := i1
		if uint64(i1) >= 32 {
			break
		}
		if err := w.WriteData(o.AllowedUnitAllocationSizes[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.AllowedUnitAllocationSizes); uint64(i1) < 32; i1++ {
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	for i1 := range o.Name {
		i1 := i1
		if uint64(i1) >= 32 {
			break
		}
		if err := w.WriteData(o.Name[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Name); uint64(i1) < 32; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *FileSystemFormatSupportProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.Revision); err != nil {
		return err
	}
	if err := w.ReadData(&o.DefaultUnitAllocationSize); err != nil {
		return err
	}
	o.AllowedUnitAllocationSizes = make([]uint32, 32)
	for i1 := range o.AllowedUnitAllocationSizes {
		i1 := i1
		if err := w.ReadData(&o.AllowedUnitAllocationSizes[i1]); err != nil {
			return err
		}
	}
	o.Name = make([]uint16, 32)
	for i1 := range o.Name {
		i1 := i1
		if err := w.ReadData(&o.Name[i1]); err != nil {
			return err
		}
	}
	return nil
}

// DiskExtent structure represents VDS_DISK_EXTENT RPC structure.
//
// The VDS_DISK_EXTENT structure provides information about a disk extent.
type DiskExtent struct {
	// diskId:  The VDS object ID of the disk object on which the extent resides.
	DiskID *ObjectID `idl:"name:diskId" json:"disk_id"`
	// type:  The value from the VDS_DISK_EXTENT_TYPE enumeration that indicates the type
	// of the extent.
	Type DiskExtentType `idl:"name:type" json:"type"`
	// ullOffset:  The byte offset of the disk extent from the beginning of the disk.
	Offset uint64 `idl:"name:ullOffset" json:"offset"`
	// ullSize:  The size, in bytes, of the extent.
	Size uint64 `idl:"name:ullSize" json:"size"`
	// volumeId:  The VDS object ID of the volume object to which the extent belongs, if
	// any.
	VolumeID *ObjectID `idl:"name:volumeId" json:"volume_id"`
	// plexId:  The VDS object ID of the volume plex object to which the extent belongs,
	// if it belongs to a volume.
	PlexID *ObjectID `idl:"name:plexId" json:"plex_id"`
	// memberIdx:  The zero-based index of the volume plex member to which the extent belongs,
	// if it belongs to a volume plex.
	MemberIndex uint32 `idl:"name:memberIdx" json:"member_index"`
}

func (o *DiskExtent) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DiskExtent) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.DiskID != nil {
		if err := o.DiskID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint16(o.Type)); err != nil {
		return err
	}
	if err := w.WriteData(o.Offset); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if o.VolumeID != nil {
		if err := o.VolumeID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.PlexID != nil {
		if err := o.PlexID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.MemberIndex); err != nil {
		return err
	}
	return nil
}
func (o *DiskExtent) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.DiskID == nil {
		o.DiskID = &ObjectID{}
	}
	if err := o.DiskID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Type)); err != nil {
		return err
	}
	if err := w.ReadData(&o.Offset); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if o.VolumeID == nil {
		o.VolumeID = &ObjectID{}
	}
	if err := o.VolumeID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.PlexID == nil {
		o.PlexID = &ObjectID{}
	}
	if err := o.PlexID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.MemberIndex); err != nil {
		return err
	}
	return nil
}

// DiskFreeExtent structure represents VDS_DISK_FREE_EXTENT RPC structure.
//
// The VDS_DISK_FREE_EXTENT structure provides information about a disk extent associated
// with free space on the disk.
type DiskFreeExtent struct {
	// diskId:  The VDS object ID of the disk object on which the extent resides.
	DiskID *ObjectID `idl:"name:diskId" json:"disk_id"`
	// ullOffset:  The byte offset of the disk extent from the beginning of the disk.
	Offset uint64 `idl:"name:ullOffset" json:"offset"`
	// ullSize:  The size, in bytes, of the extent.
	Size uint64 `idl:"name:ullSize" json:"size"`
}

func (o *DiskFreeExtent) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DiskFreeExtent) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.DiskID != nil {
		if err := o.DiskID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Offset); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	return nil
}
func (o *DiskFreeExtent) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.DiskID == nil {
		o.DiskID = &ObjectID{}
	}
	if err := o.DiskID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Offset); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	return nil
}

// PartitionProperty structure represents VDS_PARTITION_PROP RPC structure.
//
// The VDS_PARTITION_PROP structure provides information about partition properties.
type PartitionProperty struct {
	// PartitionStyle:  The value from the VDS_PARTITION_STYLE enumeration that describes
	// the partition format of the disk where the partition resides.
	PartitionStyle PartitionStyle `idl:"name:PartitionStyle" json:"partition_style"`
	// ulFlags:  The combination of any values, by using the bitwise OR operator, from the
	// VDS_PARTITION_FLAG enumeration describing the partition.
	Flags uint32 `idl:"name:ulFlags" json:"flags"`
	// ulPartitionNumber:  The one-based index number of the partition that the operating
	// system assigns.
	PartitionNumber uint32 `idl:"name:ulPartitionNumber" json:"partition_number"`
	// ullOffset:  The byte offset of the partition from the beginning of the disk.
	Offset uint64 `idl:"name:ullOffset" json:"offset"`
	// ullSize:  The size of the partition, in bytes.
	Size              uint64                               `idl:"name:ullSize" json:"size"`
	PartitionProperty *PartitionProperty_PartitionProperty `idl:"name:PartitionProperty;switch_is:PartitionStyle" json:"partition_property"`
}

func (o *PartitionProperty) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PartitionProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.PartitionStyle)); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.PartitionNumber); err != nil {
		return err
	}
	if err := w.WriteData(o.Offset); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	_swPartitionProperty := uint16(o.PartitionStyle)
	if o.PartitionProperty != nil {
		if err := o.PartitionProperty.MarshalUnionNDR(ctx, w, _swPartitionProperty); err != nil {
			return err
		}
	} else {
		if err := (&PartitionProperty_PartitionProperty{}).MarshalUnionNDR(ctx, w, _swPartitionProperty); err != nil {
			return err
		}
	}
	return nil
}
func (o *PartitionProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.PartitionStyle)); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.PartitionNumber); err != nil {
		return err
	}
	if err := w.ReadData(&o.Offset); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if o.PartitionProperty == nil {
		o.PartitionProperty = &PartitionProperty_PartitionProperty{}
	}
	_swPartitionProperty := uint16(o.PartitionStyle)
	if err := o.PartitionProperty.UnmarshalUnionNDR(ctx, w, _swPartitionProperty); err != nil {
		return err
	}
	return nil
}

// PartitionProperty_PartitionProperty structure represents VDS_PARTITION_PROP union anonymous member.
//
// The VDS_PARTITION_PROP structure provides information about partition properties.
type PartitionProperty_PartitionProperty struct {
	// Types that are assignable to Value
	//
	// *PartitionProperty_MBR
	// *PartitionProperty_GPT
	// *PartitionProperty_DefaultPartitionProperty
	Value is_PartitionProperty_PartitionProperty `json:"value"`
}

func (o *PartitionProperty_PartitionProperty) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *PartitionProperty_MBR:
		if value != nil {
			return value.MBR
		}
	case *PartitionProperty_GPT:
		if value != nil {
			return value.GPT
		}
	}
	return nil
}

type is_PartitionProperty_PartitionProperty interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_PartitionProperty_PartitionProperty()
}

func (o *PartitionProperty_PartitionProperty) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *PartitionProperty_MBR:
		return uint16(1)
	case *PartitionProperty_GPT:
		return uint16(2)
	}
	return uint16(0)
}

func (o *PartitionProperty_PartitionProperty) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		_o, _ := o.Value.(*PartitionProperty_MBR)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PartitionProperty_MBR{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*PartitionProperty_GPT)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PartitionProperty_GPT{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		// no-op
	}
	return nil
}

func (o *PartitionProperty_PartitionProperty) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch((*uint16)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		o.Value = &PartitionProperty_MBR{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &PartitionProperty_GPT{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		o.Value = &PartitionProperty_DefaultPartitionProperty{}
	}
	return nil
}

// PartitionProperty_MBR structure represents PartitionProperty_PartitionProperty RPC union arm.
//
// It has following labels: 1
type PartitionProperty_MBR struct {
	// Mbr:  A VDS_PARTITION_INFO_MBR structure that describes the MBR partition.
	MBR *PartitionInfoMBR `idl:"name:Mbr" json:"mbr"`
}

func (*PartitionProperty_MBR) is_PartitionProperty_PartitionProperty() {}

func (o *PartitionProperty_MBR) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.MBR != nil {
		if err := o.MBR.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PartitionInfoMBR{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PartitionProperty_MBR) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.MBR == nil {
		o.MBR = &PartitionInfoMBR{}
	}
	if err := o.MBR.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PartitionProperty_GPT structure represents PartitionProperty_PartitionProperty RPC union arm.
//
// It has following labels: 2
type PartitionProperty_GPT struct {
	// Gpt:  A VDS_PARTITION_INFO_GPT structure that describes the GPT partition.
	GPT *PartitionInfoGPT `idl:"name:Gpt" json:"gpt"`
}

func (*PartitionProperty_GPT) is_PartitionProperty_PartitionProperty() {}

func (o *PartitionProperty_GPT) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.GPT != nil {
		if err := o.GPT.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PartitionInfoGPT{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PartitionProperty_GPT) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.GPT == nil {
		o.GPT = &PartitionInfoGPT{}
	}
	if err := o.GPT.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PartitionProperty_DefaultPartitionProperty structure represents PartitionProperty_PartitionProperty RPC default union arm.
type PartitionProperty_DefaultPartitionProperty struct {
}

func (*PartitionProperty_DefaultPartitionProperty) is_PartitionProperty_PartitionProperty() {}

func (o *PartitionProperty_DefaultPartitionProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return nil
}
func (o *PartitionProperty_DefaultPartitionProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	return nil
}

// InputDisk structure represents VDS_INPUT_DISK RPC structure.
//
// The VDS_INPUT_DISK structure provides information about a disk for volume creation
// and modification.
type InputDisk struct {
	// diskId:  The VDS object ID of the disk object.
	DiskID *ObjectID `idl:"name:diskId" json:"disk_id"`
	// ullSize:  The size of the disk to use, in bytes.
	Size uint64 `idl:"name:ullSize" json:"size"`
	// plexId:  When extending a volume, the VDS object ID of the plex object to which the
	// disk will be added. A volume can only be extended by extending all members of all
	// plexes in the same operation. This member is used when extending any volume and ignored
	// when creating a volume or repairing a RAID-5 volume.
	PlexID *ObjectID `idl:"name:plexId" json:"plex_id"`
	// memberIdx:  The zero-based member index of the disk to which the extent belongs.
	// Either specify a memberIdx for all disks or specify it for none. VDS uses disks with
	// the same memberIdx in the order they appear in the array. For example, the first
	// disk in the array is always used first, even if it does not have the lowest index.
	// This member is ignored when repairing a RAID-5 volume.
	MemberIndex uint32 `idl:"name:memberIdx" json:"member_index"`
}

func (o *InputDisk) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *InputDisk) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.DiskID != nil {
		if err := o.DiskID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if o.PlexID != nil {
		if err := o.PlexID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.MemberIndex); err != nil {
		return err
	}
	return nil
}
func (o *InputDisk) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.DiskID == nil {
		o.DiskID = &ObjectID{}
	}
	if err := o.DiskID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if o.PlexID == nil {
		o.PlexID = &ObjectID{}
	}
	if err := o.PlexID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.MemberIndex); err != nil {
		return err
	}
	return nil
}

// CreatePartitionParameters structure represents CREATE_PARTITION_PARAMETERS RPC structure.
//
// The CREATE_PARTITION_PARAMETERS structure provides information about partition properties.
type CreatePartitionParameters struct {
	// style:  A value from the VDS_PARTITION_STYLE enumeration that describes the disk
	// partition format.
	Style                     PartitionStyle                                       `idl:"name:style" json:"style"`
	CreatePartitionParameters *CreatePartitionParameters_CreatePartitionParameters `idl:"name:CreatePartitionParameters;switch_is:style" json:"create_partition_parameters"`
}

func (o *CreatePartitionParameters) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CreatePartitionParameters) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Style)); err != nil {
		return err
	}
	_swCreatePartitionParameters := uint16(o.Style)
	if o.CreatePartitionParameters != nil {
		if err := o.CreatePartitionParameters.MarshalUnionNDR(ctx, w, _swCreatePartitionParameters); err != nil {
			return err
		}
	} else {
		if err := (&CreatePartitionParameters_CreatePartitionParameters{}).MarshalUnionNDR(ctx, w, _swCreatePartitionParameters); err != nil {
			return err
		}
	}
	return nil
}
func (o *CreatePartitionParameters) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Style)); err != nil {
		return err
	}
	if o.CreatePartitionParameters == nil {
		o.CreatePartitionParameters = &CreatePartitionParameters_CreatePartitionParameters{}
	}
	_swCreatePartitionParameters := uint16(o.Style)
	if err := o.CreatePartitionParameters.UnmarshalUnionNDR(ctx, w, _swCreatePartitionParameters); err != nil {
		return err
	}
	return nil
}

// CreatePartitionParameters_CreatePartitionParameters structure represents CREATE_PARTITION_PARAMETERS union anonymous member.
//
// The CREATE_PARTITION_PARAMETERS structure provides information about partition properties.
type CreatePartitionParameters_CreatePartitionParameters struct {
	// Types that are assignable to Value
	//
	// *CreatePartitionParameters_MBRPartitionInfo
	// *CreatePartitionParameters_GPTPartitionInfo
	// *CreatePartitionParameters_DefaultCreatePartitionParameters
	Value is_CreatePartitionParameters_CreatePartitionParameters `json:"value"`
}

func (o *CreatePartitionParameters_CreatePartitionParameters) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *CreatePartitionParameters_MBRPartitionInfo:
		if value != nil {
			return value.MBRPartitionInfo
		}
	case *CreatePartitionParameters_GPTPartitionInfo:
		if value != nil {
			return value.GPTPartitionInfo
		}
	}
	return nil
}

type is_CreatePartitionParameters_CreatePartitionParameters interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_CreatePartitionParameters_CreatePartitionParameters()
}

func (o *CreatePartitionParameters_CreatePartitionParameters) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *CreatePartitionParameters_MBRPartitionInfo:
		return uint16(1)
	case *CreatePartitionParameters_GPTPartitionInfo:
		return uint16(2)
	}
	return uint16(0)
}

func (o *CreatePartitionParameters_CreatePartitionParameters) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		_o, _ := o.Value.(*CreatePartitionParameters_MBRPartitionInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&CreatePartitionParameters_MBRPartitionInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*CreatePartitionParameters_GPTPartitionInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&CreatePartitionParameters_GPTPartitionInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		// no-op
	}
	return nil
}

func (o *CreatePartitionParameters_CreatePartitionParameters) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch((*uint16)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		o.Value = &CreatePartitionParameters_MBRPartitionInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &CreatePartitionParameters_GPTPartitionInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		o.Value = &CreatePartitionParameters_DefaultCreatePartitionParameters{}
	}
	return nil
}

// CreatePartitionParameters_MBRPartitionInfo structure represents CreatePartitionParameters_CreatePartitionParameters RPC union arm.
//
// It has following labels: 1
type CreatePartitionParameters_MBRPartitionInfo struct {
	// MbrPartInfo:  Contains information for an MBR partition.
	MBRPartitionInfo *CreatePartitionParameters_CreatePartitionParameters_MBRPartitionInfo `idl:"name:MbrPartInfo" json:"mbr_partition_info"`
}

func (*CreatePartitionParameters_MBRPartitionInfo) is_CreatePartitionParameters_CreatePartitionParameters() {
}

func (o *CreatePartitionParameters_MBRPartitionInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.MBRPartitionInfo != nil {
		if err := o.MBRPartitionInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&CreatePartitionParameters_CreatePartitionParameters_MBRPartitionInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *CreatePartitionParameters_MBRPartitionInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.MBRPartitionInfo == nil {
		o.MBRPartitionInfo = &CreatePartitionParameters_CreatePartitionParameters_MBRPartitionInfo{}
	}
	if err := o.MBRPartitionInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// CreatePartitionParameters_CreatePartitionParameters_MBRPartitionInfo structure represents CREATE_PARTITION_PARAMETERS structure anonymous member.
//
// The CREATE_PARTITION_PARAMETERS structure provides information about partition properties.
type CreatePartitionParameters_CreatePartitionParameters_MBRPartitionInfo struct {
	// partitionType:  The byte value that indicates the partition type to create.
	PartitionType uint8 `idl:"name:partitionType" json:"partition_type"`
	// bootIndicator:  A Boolean value that indicates whether the partition is bootable.
	BootIndicator bool `idl:"name:bootIndicator" json:"boot_indicator"`
}

func (o *CreatePartitionParameters_CreatePartitionParameters_MBRPartitionInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CreatePartitionParameters_CreatePartitionParameters_MBRPartitionInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteData(o.PartitionType); err != nil {
		return err
	}
	if err := w.WriteData(o.BootIndicator); err != nil {
		return err
	}
	return nil
}
func (o *CreatePartitionParameters_CreatePartitionParameters_MBRPartitionInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.PartitionType); err != nil {
		return err
	}
	if err := w.ReadData(&o.BootIndicator); err != nil {
		return err
	}
	return nil
}

// CreatePartitionParameters_GPTPartitionInfo structure represents CreatePartitionParameters_CreatePartitionParameters RPC union arm.
//
// It has following labels: 2
type CreatePartitionParameters_GPTPartitionInfo struct {
	// GptPartInfo:  Contains information about a GPT partition.
	GPTPartitionInfo *CreatePartitionParameters_CreatePartitionParameters_GPTPartitionInfo `idl:"name:GptPartInfo" json:"gpt_partition_info"`
}

func (*CreatePartitionParameters_GPTPartitionInfo) is_CreatePartitionParameters_CreatePartitionParameters() {
}

func (o *CreatePartitionParameters_GPTPartitionInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.GPTPartitionInfo != nil {
		if err := o.GPTPartitionInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&CreatePartitionParameters_CreatePartitionParameters_GPTPartitionInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *CreatePartitionParameters_GPTPartitionInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.GPTPartitionInfo == nil {
		o.GPTPartitionInfo = &CreatePartitionParameters_CreatePartitionParameters_GPTPartitionInfo{}
	}
	if err := o.GPTPartitionInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// CreatePartitionParameters_CreatePartitionParameters_GPTPartitionInfo structure represents CREATE_PARTITION_PARAMETERS structure anonymous member.
//
// The CREATE_PARTITION_PARAMETERS structure provides information about partition properties.
type CreatePartitionParameters_CreatePartitionParameters_GPTPartitionInfo struct {
	// partitionType:  The byte value that indicates the partition type to create.
	PartitionType *dtyp.GUID `idl:"name:partitionType" json:"partition_type"`
	// partitionId:  The GUID of the partition.
	PartitionID *dtyp.GUID `idl:"name:partitionId" json:"partition_id"`
	// attributes:  A bitwise OR operator of attributes that is used to create the partition;
	// it can have a combination of the following values.
	//
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                                             |                                                                                  |
	//	|                            VALUE                            |                                     MEANING                                      |
	//	|                                                             |                                                                                  |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| GPT_ATTRIBUTE_PLATFORM_REQUIRED 0x0000000000000001          | A partition is required for the platform to function properly.<25>               |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| GPT_BASIC_DATA_ATTRIBUTE_READ_ONLY 0x1000000000000000       | The partition can be read from, but not written to. Used only with the basic     |
	//	|                                                             | data partition type.                                                             |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| GPT_BASIC_DATA_ATTRIBUTE_HIDDEN 0x4000000000000000          | The partition is hidden and is not mounted. Used only with the basic data        |
	//	|                                                             | partition type.                                                                  |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| GPT_BASIC_DATA_ATTRIBUTE_NO_DRIVE_LETTER 0x8000000000000000 | The partition does not receive a drive letter by default when moving the disk to |
	//	|                                                             | another computer. Used only with the basic data partition type.                  |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	Attributes uint64 `idl:"name:attributes" json:"attributes"`
	// name:  The null-terminated Unicode name of the partition.
	Name []uint16 `idl:"name:name" json:"name"`
}

func (o *CreatePartitionParameters_CreatePartitionParameters_GPTPartitionInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CreatePartitionParameters_CreatePartitionParameters_GPTPartitionInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.PartitionType != nil {
		if err := o.PartitionType.MarshalNDR(ctx, w); err != nil {
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
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	for i1 := range o.Name {
		i1 := i1
		if uint64(i1) >= 24 {
			break
		}
		if err := w.WriteData(o.Name[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Name); uint64(i1) < 24; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *CreatePartitionParameters_CreatePartitionParameters_GPTPartitionInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.PartitionType == nil {
		o.PartitionType = &dtyp.GUID{}
	}
	if err := o.PartitionType.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.PartitionID == nil {
		o.PartitionID = &dtyp.GUID{}
	}
	if err := o.PartitionID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	o.Name = make([]uint16, 24)
	for i1 := range o.Name {
		i1 := i1
		if err := w.ReadData(&o.Name[i1]); err != nil {
			return err
		}
	}
	return nil
}

// CreatePartitionParameters_DefaultCreatePartitionParameters structure represents CreatePartitionParameters_CreatePartitionParameters RPC default union arm.
type CreatePartitionParameters_DefaultCreatePartitionParameters struct {
}

func (*CreatePartitionParameters_DefaultCreatePartitionParameters) is_CreatePartitionParameters_CreatePartitionParameters() {
}

func (o *CreatePartitionParameters_DefaultCreatePartitionParameters) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return nil
}
func (o *CreatePartitionParameters_DefaultCreatePartitionParameters) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	return nil
}

// VirtualStorageType structure represents VIRTUAL_STORAGE_TYPE RPC structure.
//
// The VIRTUAL_STORAGE_TYPE structure specifies the device and vendor of the virtual
// disk.<26>
type VirtualStorageType struct {
	// DeviceId:  The virtual disk type. It can have one of the following values.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                                       |                                                                                  |
	//	|                 VALUE                 |                                     MEANING                                      |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| VIRTUAL_STORAGE_TYPE_DEVICE_UNKNOWN 0 | The virtual disk type is unknown.                                                |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| VIRTUAL_STORAGE_TYPE_DEVICE_ISO 1     | The virtual disk is an ISO image (.iso) file. For more information, see          |
	//	|                                       | [ECMA-119] and [OSTA-UDFS].                                                      |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| VIRTUAL_STORAGE_TYPE_DEVICE_VHD 2     | The virtual disk is a virtual hard disk (.vhd) file.                             |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	DeviceID uint32 `idl:"name:DeviceId" json:"device_id"`
	// VendorId:  A GUID that uniquely identifies the virtual disk vendor.
	VendorID *dtyp.GUID `idl:"name:VendorId" json:"vendor_id"`
}

func (o *VirtualStorageType) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *VirtualStorageType) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DeviceID); err != nil {
		return err
	}
	if o.VendorID != nil {
		if err := o.VendorID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *VirtualStorageType) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DeviceID); err != nil {
		return err
	}
	if o.VendorID == nil {
		o.VendorID = &dtyp.GUID{}
	}
	if err := o.VendorID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PartitionStyle2 type represents __VDS_PARTITION_STYLE2 RPC enumeration.
type PartitionStyle2 uint16

var (
	PartitionStyle2MBR PartitionStyle2 = 0
	PartitionStyle2GPT PartitionStyle2 = 1
	PartitionStyle2Raw PartitionStyle2 = 2
)

func (o PartitionStyle2) String() string {
	switch o {
	case PartitionStyle2MBR:
		return "PartitionStyle2MBR"
	case PartitionStyle2GPT:
		return "PartitionStyle2GPT"
	case PartitionStyle2Raw:
		return "PartitionStyle2Raw"
	}
	return "Invalid"
}

// ObjectType type represents VDS_OBJECT_TYPE RPC enumeration.
//
// The VDS_OBJECT_TYPE enumeration defines the set of valid VDS object types.
type ObjectType uint16

var (
	// VDS_OT_UNKNOWN:  The object has an unknown type.
	ObjectTypeUnknown ObjectType = 0
	// VDS_OT_PROVIDER:  The object is a provider.
	ObjectTypeProvider ObjectType = 1
	// VDS_OT_PACK:  The object is a pack (a disk group).
	ObjectTypePack ObjectType = 10
	// VDS_OT_VOLUME:  The object is a volume.
	ObjectTypeVolume ObjectType = 11
	// VDS_OT_VOLUME_PLEX:  The object is a plex of a volume.
	ObjectTypeVolumePlex ObjectType = 12
	// VDS_OT_DISK:  The object is a disk.
	ObjectTypeDisk ObjectType = 13
	// VDS_OT_HBAPORT:  The object is an HBA port.
	ObjectTypeHBAPort ObjectType = 90
	// VDS_OT_INIT_ADAPTER:  The object is an iSCSI initiator adapter.
	ObjectTypeInitiatorAdapter ObjectType = 91
	// VDS_OT_INIT_PORTAL:  The object is an iSCSI initiator portal.
	ObjectTypeInitiatorPortal ObjectType = 92
	// VDS_OT_ASYNC:  The object maintains the status of an asynchronous VDS operation.
	ObjectTypeAsync ObjectType = 100
	// VDS_OT_ENUM:  The object is an enumerator that contains an enumeration of other
	// VDS objects.
	ObjectTypeEnum ObjectType = 101
	// VDS_OT_VDISK:  The object is a virtual disk.
	ObjectTypeVDisk ObjectType = 200
	// VDS_OT_OPEN_VDISK:  The object represents an open virtual disk (an OpenVirtualDisk
	// object).
	ObjectTypeOpenVDisk ObjectType = 201
)

func (o ObjectType) String() string {
	switch o {
	case ObjectTypeUnknown:
		return "ObjectTypeUnknown"
	case ObjectTypeProvider:
		return "ObjectTypeProvider"
	case ObjectTypePack:
		return "ObjectTypePack"
	case ObjectTypeVolume:
		return "ObjectTypeVolume"
	case ObjectTypeVolumePlex:
		return "ObjectTypeVolumePlex"
	case ObjectTypeDisk:
		return "ObjectTypeDisk"
	case ObjectTypeHBAPort:
		return "ObjectTypeHBAPort"
	case ObjectTypeInitiatorAdapter:
		return "ObjectTypeInitiatorAdapter"
	case ObjectTypeInitiatorPortal:
		return "ObjectTypeInitiatorPortal"
	case ObjectTypeAsync:
		return "ObjectTypeAsync"
	case ObjectTypeEnum:
		return "ObjectTypeEnum"
	case ObjectTypeVDisk:
		return "ObjectTypeVDisk"
	case ObjectTypeOpenVDisk:
		return "ObjectTypeOpenVDisk"
	}
	return "Invalid"
}

// ServiceFlag type represents VDS_SERVICE_FLAG RPC enumeration.
//
// The VDS_SERVICE_FLAG enumeration defines the properties of the service.
type ServiceFlag uint16

var (
	// VDS_SVF_SUPPORT_DYNAMIC:  The server supports dynamic disks.
	ServiceFlagSupportDynamic ServiceFlag = 1
	// VDS_SVF_SUPPORT_FAULT_TOLERANT:  The server supports fault-tolerant disks.
	ServiceFlagSupportFaultTolerant ServiceFlag = 2
	// VDS_SVF_SUPPORT_GPT:  The server supports the GPT partition format.
	ServiceFlagSupportGPT ServiceFlag = 4
	// VDS_SVF_SUPPORT_DYNAMIC_1394:  The server supports dynamic disks that use the IEEE
	// 1394 interface for the host bus adapter connection. For more information on IEEE
	// 1394, see [IEEE1394-2008].
	ServiceFlagSupportDynamic1394 ServiceFlag = 8
	// VDS_SVF_CLUSTER_SERVICE_CONFIGURED:  The server is running on a cluster.
	ServiceFlagClusterServiceConfigured ServiceFlag = 16
	// VDS_SVF_AUTO_MOUNT_OFF:  The server will not automatically mount disks.
	ServiceFlagAutoMountOff ServiceFlag = 32
	// VDS_SVF_OS_UNINSTALL_VALID:  The server has an uninstall image to which it can roll
	// back.
	ServiceFlagOSUninstallValid ServiceFlag = 64
	// VDS_SVF_EFI:  The computer starts an EFI from a GPT partition.
	ServiceFlagEFI ServiceFlag = 128
	// VDS_SVF_SUPPORT_MIRROR:  The server supports mirrored volumes (RAID-1).
	ServiceFlagSupportMirror ServiceFlag = 256
	// VDS_SVF_SUPPORT_RAIDS:  The server supports striped with parity volumes (RAID-5).
	ServiceFlagSupportRAID5 ServiceFlag = 512
	// VDS_SVF_SUPPORT_REFS<27>:  The server supports the ReFS.
	ServiceFlagReFS ServiceFlag = 1024
)

func (o ServiceFlag) String() string {
	switch o {
	case ServiceFlagSupportDynamic:
		return "ServiceFlagSupportDynamic"
	case ServiceFlagSupportFaultTolerant:
		return "ServiceFlagSupportFaultTolerant"
	case ServiceFlagSupportGPT:
		return "ServiceFlagSupportGPT"
	case ServiceFlagSupportDynamic1394:
		return "ServiceFlagSupportDynamic1394"
	case ServiceFlagClusterServiceConfigured:
		return "ServiceFlagClusterServiceConfigured"
	case ServiceFlagAutoMountOff:
		return "ServiceFlagAutoMountOff"
	case ServiceFlagOSUninstallValid:
		return "ServiceFlagOSUninstallValid"
	case ServiceFlagEFI:
		return "ServiceFlagEFI"
	case ServiceFlagSupportMirror:
		return "ServiceFlagSupportMirror"
	case ServiceFlagSupportRAID5:
		return "ServiceFlagSupportRAID5"
	case ServiceFlagReFS:
		return "ServiceFlagReFS"
	}
	return "Invalid"
}

// ProviderType type represents VDS_PROVIDER_TYPE RPC enumeration.
//
// The VDS_PROVIDER_TYPE enumeration defines the set of valid types for a provider.
type ProviderType uint16

var (
	// VDS_PT_UNKNOWN:  The type is neither a software nor a hardware provider.
	ProviderTypeUnknown ProviderType = 0
	// VDS_PT_SOFTWARE:  The type indicates a program that is responsible for volume management.
	ProviderTypeSoftware ProviderType = 1
	// VDS_PT_HARDWARE:  The type indicates a program that is responsible for aspects of
	// hardware storage management.
	ProviderTypeHardware ProviderType = 2
	// VDS_PT_VIRTUALDISK:  The type indicates a program that is responsible for aspects
	// of hardware storage management.
	ProviderTypeVirtualDisk ProviderType = 3
	// VDS_PT_MAX:  Denotes the maximum acceptable value for this type. VDS_PT_MAX - 1,
	// ('3'), is the maximum acceptable value.
	ProviderTypeMax ProviderType = 4
)

func (o ProviderType) String() string {
	switch o {
	case ProviderTypeUnknown:
		return "ProviderTypeUnknown"
	case ProviderTypeSoftware:
		return "ProviderTypeSoftware"
	case ProviderTypeHardware:
		return "ProviderTypeHardware"
	case ProviderTypeVirtualDisk:
		return "ProviderTypeVirtualDisk"
	case ProviderTypeMax:
		return "ProviderTypeMax"
	}
	return "Invalid"
}

// ProviderFlag type represents VDS_PROVIDER_FLAG RPC enumeration.
//
// The VDS_PROVIDER_FLAG enumeration defines the set of valid flags for a provider object.
type ProviderFlag uint32

var (
	// VDS_PF_DYNAMIC:  If set, all disks that the current provider manages are dynamic.
	// This flag MUST be set only by a dynamic provider. By definition, dynamic providers
	// manage only dynamic disks.
	ProviderFlagDynamic ProviderFlag = 1
	// VDS_PF_INTERNAL_HARDWARE_PROVIDER:  Reserved for internal use.
	ProviderFlagInternalHardwareProvider ProviderFlag = 2
	// VDS_PF_ONE_DISK_ONLY_PER_PACK:  If set, the provider supports single disk packs
	// only. Typically, the basic provider sets this flag to simulate a disk pack that has
	// one disk.
	ProviderFlagOneDiskOnlyPerPack ProviderFlag = 4
	// VDS_PF_ONE_PACK_ONLINE_ONLY:  If set, the dynamic provider supports online status
	// for only one pack at a time.
	ProviderFlagOnePackOnlineOnly ProviderFlag = 8
	// VDS_PF_VOLUME_SPACE_MUST_BE_CONTIGUOUS:  If set, all volumes that this provider
	// manages MUST have contiguous space. This flag applies to the basic provider only.
	ProviderFlagVolumeSpaceMustBeContiguous ProviderFlag = 16
	// VDS_PF_SUPPORT_MIRROR:  If set, the provider supports mirrored volumes (RAID-1).
	ProviderFlagSupportMirror ProviderFlag = 32
	// VDS_PF_SUPPORT_RAID5:  If set, the provider supports striped with parity volumes
	// (RAID-5).
	ProviderFlagSupportRAID5 ProviderFlag = 64
	// VDS_PF_SUPPORT_DYNAMIC_1394:  If set, the provider supports IEEE 1394 dynamic disks.
	// This flag MUST be set only by the dynamic provider on systems that support IEEE 1394
	// dynamic disks.
	ProviderFlagSupportDynamic1394 ProviderFlag = 536870912
	// VDS_PF_SUPPORT_FAULT_TOLERANT:  If set, the provider supports fault-tolerant disks.
	// This flag MUST be set only by the dynamic provider on systems that support fault-tolerant
	// volumes.
	ProviderFlagSupportFaultTolerant ProviderFlag = 1073741824
	// VDS_PF_SUPPORT_DYNAMIC:  If set, the provider supports managing dynamic disks. This
	// flag MUST be set only by the dynamic provider on systems that support dynamic disks.
	ProviderFlagSupportDynamic ProviderFlag = 2147483648
)

func (o ProviderFlag) String() string {
	switch o {
	case ProviderFlagDynamic:
		return "ProviderFlagDynamic"
	case ProviderFlagInternalHardwareProvider:
		return "ProviderFlagInternalHardwareProvider"
	case ProviderFlagOneDiskOnlyPerPack:
		return "ProviderFlagOneDiskOnlyPerPack"
	case ProviderFlagOnePackOnlineOnly:
		return "ProviderFlagOnePackOnlineOnly"
	case ProviderFlagVolumeSpaceMustBeContiguous:
		return "ProviderFlagVolumeSpaceMustBeContiguous"
	case ProviderFlagSupportMirror:
		return "ProviderFlagSupportMirror"
	case ProviderFlagSupportRAID5:
		return "ProviderFlagSupportRAID5"
	case ProviderFlagSupportDynamic1394:
		return "ProviderFlagSupportDynamic1394"
	case ProviderFlagSupportFaultTolerant:
		return "ProviderFlagSupportFaultTolerant"
	case ProviderFlagSupportDynamic:
		return "ProviderFlagSupportDynamic"
	}
	return "Invalid"
}

// QueryProviderFlag type represents VDS_QUERY_PROVIDER_FLAG RPC enumeration.
//
// The VDS_QUERY_PROVIDER_FLAG enumeration defines the set of valid flags for provider
// query operations. Callers can query for hardware providers, software providers, or
// both.<28>
type QueryProviderFlag uint16

var (
	// VDS_QUERY_SOFTWARE_PROVIDERS:  If set, the operation queries for software providers.
	QueryProviderFlagSoftwareProviders QueryProviderFlag = 1
	// VDS_QUERY_HARDWARE_PROVIDERS:  If set, the operation queries for hardware providers.
	QueryProviderFlagHardwareProviders QueryProviderFlag = 2
	// VDS_QUERY_VIRTUALDISK_PROVIDERS:  If set, the operation queries for virtual disk
	// providers.
	QueryProviderFlagVirtualDiskProviders QueryProviderFlag = 4
)

func (o QueryProviderFlag) String() string {
	switch o {
	case QueryProviderFlagSoftwareProviders:
		return "QueryProviderFlagSoftwareProviders"
	case QueryProviderFlagHardwareProviders:
		return "QueryProviderFlagHardwareProviders"
	case QueryProviderFlagVirtualDiskProviders:
		return "QueryProviderFlagVirtualDiskProviders"
	}
	return "Invalid"
}

// DriveLetterFlag type represents VDS_DRIVE_LETTER_FLAG RPC enumeration.
//
// The VDS_DRIVE_LETTER_FLAG enumeration defines the set of valid flags for a drive
// letter.
type DriveLetterFlag uint16

var (
	// VDS_DLF_NON_PERSISTENT:  If set, the drive letter no longer appears after the computer
	// is restarted.
	DriveLetterFlagNonPersistent DriveLetterFlag = 1
)

func (o DriveLetterFlag) String() string {
	switch o {
	case DriveLetterFlagNonPersistent:
		return "DriveLetterFlagNonPersistent"
	}
	return "Invalid"
}

// PackStatus type represents VDS_PACK_STATUS RPC enumeration.
//
// The VDS_PACK_STATUS enumeration defines the set of object status values for a disk
// pack.
type PackStatus uint16

var (
	// VDS_PS_UNKNOWN:  The status of the disk pack cannot be determined.
	PackStatusUnknown PackStatus = 0
	// VDS_PS_ONLINE:  The disk pack is available.
	PackStatusOnline PackStatus = 1
	// VDS_PS_OFFLINE:  The disk pack is unavailable; the disks are not accessible.
	PackStatusOffline PackStatus = 4
)

func (o PackStatus) String() string {
	switch o {
	case PackStatusUnknown:
		return "PackStatusUnknown"
	case PackStatusOnline:
		return "PackStatusOnline"
	case PackStatusOffline:
		return "PackStatusOffline"
	}
	return "Invalid"
}

// PackFlag type represents VDS_PACK_FLAG RPC enumeration.
//
// The VDS_PACK_FLAG enumeration defines the set of valid flags for a disk pack object.
type PackFlag uint16

var (
	// VDS_PKF_FOREIGN:  If set, an external disk pack is eligible for online status.
	PackFlagForeign PackFlag = 1
	// VDS_PKF_NOQUORUM:  If set, a dynamic disk pack lacks the required disk quorum.
	PackFlagNoQuorum PackFlag = 2
	// VDS_PKF_POLICY:  If set, management policy forbids the disk pack from gaining online
	// status.
	PackFlagPolicy PackFlag = 4
	// VDS_PKF_CORRUPTED:  If set, a disk pack contains a disk that has a corrupted LDM
	// database.
	PackFlagCorrupted PackFlag = 8
	// VDS_PKF_ONLINE_ERROR:  If set, a disk pack with sufficient disk quorum failed to
	// achieve online status due to an error.
	PackFlagOnlineError PackFlag = 16
)

func (o PackFlag) String() string {
	switch o {
	case PackFlagForeign:
		return "PackFlagForeign"
	case PackFlagNoQuorum:
		return "PackFlagNoQuorum"
	case PackFlagPolicy:
		return "PackFlagPolicy"
	case PackFlagCorrupted:
		return "PackFlagCorrupted"
	case PackFlagOnlineError:
		return "PackFlagOnlineError"
	}
	return "Invalid"
}

// DiskOfflineReason type represents VDS_DISK_OFFLINE_REASON RPC enumeration.
//
// The VDS_DISK_OFFLINE_REASON enumeration defines the reason for the disk to be kept
// offline.
type DiskOfflineReason uint16

var (
	// VDSDiskOfflineReasonNone:  The reason is unknown.
	DiskOfflineReasonNone DiskOfflineReason = 0
	// VDSDiskOfflineReasonPolicy:  The disk is offline because of the SAN policy.
	DiskOfflineReasonPolicy DiskOfflineReason = 1
	// VDSDiskOfflineReasonRedundantPath:  The disk is offline because it was determined
	// that the disk is a redundant path to another disk that is online.
	DiskOfflineReasonRedundantPath DiskOfflineReason = 2
	// VDSDiskOfflineReasonSnapshot:  The disk is offline because it is a snapshot disk.
	DiskOfflineReasonSnapshot DiskOfflineReason = 3
	// VDSDiskOfflineReasonCollision:  The disk is offline because its disk signature is
	// the same as the disk signature of another disk that is online.
	DiskOfflineReasonCollision DiskOfflineReason = 4
	// VDSDiskOfflineReasonResourceExhaustion<43>:  The disk is offline because of lack
	// of capacity.
	DiskOfflineReasonResourceExhaustion DiskOfflineReason = 5
	// VDSDiskOfflineReasonWriteFailure<44>:  The disk is offline because of critical write
	// failures.
	DiskOfflineReasonWriteFailure DiskOfflineReason = 6
	// VDSDiskOfflineReasonDIScan<45>:  The disk is offline because a data integrity scan
	// is required.
	DiskOfflineReasonDIScan DiskOfflineReason = 7
)

func (o DiskOfflineReason) String() string {
	switch o {
	case DiskOfflineReasonNone:
		return "DiskOfflineReasonNone"
	case DiskOfflineReasonPolicy:
		return "DiskOfflineReasonPolicy"
	case DiskOfflineReasonRedundantPath:
		return "DiskOfflineReasonRedundantPath"
	case DiskOfflineReasonSnapshot:
		return "DiskOfflineReasonSnapshot"
	case DiskOfflineReasonCollision:
		return "DiskOfflineReasonCollision"
	case DiskOfflineReasonResourceExhaustion:
		return "DiskOfflineReasonResourceExhaustion"
	case DiskOfflineReasonWriteFailure:
		return "DiskOfflineReasonWriteFailure"
	case DiskOfflineReasonDIScan:
		return "DiskOfflineReasonDIScan"
	}
	return "Invalid"
}

// VolumePlexType type represents VDS_VOLUME_PLEX_TYPE RPC enumeration.
//
// The VDS_VOLUME_PLEX_TYPE enumeration defines the set of valid types for a volume
// plex.
type VolumePlexType uint16

var (
	// VDS_VPT_UNKNOWN:  The volume plex type is unknown.
	VolumePlexTypeUnknown VolumePlexType = 0
	// VDS_VPT_SIMPLE:  The plex type is simple; it is composed of extents from exactly
	// one disk.
	VolumePlexTypeSimple VolumePlexType = 10
	// VDS_VPT_SPAN:  The plex type is spanned; it is composed of extents from more than
	// one disk.
	VolumePlexTypeSpan VolumePlexType = 11
	// VDS_VPT_STRIPE:  The plex type is striped, which is equivalent to RAID-0.
	VolumePlexTypeStripe VolumePlexType = 12
	// VDS_VPT_PARITY:  The plex type is striped with parity, which accounts for RAID levels
	// 3, 4, 5, and 6.
	VolumePlexTypeParity VolumePlexType = 14
)

func (o VolumePlexType) String() string {
	switch o {
	case VolumePlexTypeUnknown:
		return "VolumePlexTypeUnknown"
	case VolumePlexTypeSimple:
		return "VolumePlexTypeSimple"
	case VolumePlexTypeSpan:
		return "VolumePlexTypeSpan"
	case VolumePlexTypeStripe:
		return "VolumePlexTypeStripe"
	case VolumePlexTypeParity:
		return "VolumePlexTypeParity"
	}
	return "Invalid"
}

// VolumePlexStatus type represents VDS_VOLUME_PLEX_STATUS RPC enumeration.
//
// The VDS_VOLUME_PLEX_STATUS enumeration defines the set of object status values for
// a volume plex.
type VolumePlexStatus uint16

var (
	// VDS_VPS_UNKNOWN:  The status of the volume plex is unknown.
	VolumePlexStatusUnknown VolumePlexStatus = 0
	// VDS_VPS_ONLINE:  The volume plex is available.
	VolumePlexStatusOnline VolumePlexStatus = 1
	// VDS_VPS_NO_MEDIA:  The volume plex has no media.
	VolumePlexStatusNoMedia VolumePlexStatus = 3
	// VDS_VPS_FAILED:  The volume plex is unavailable.
	VolumePlexStatusFailed VolumePlexStatus = 5
)

func (o VolumePlexStatus) String() string {
	switch o {
	case VolumePlexStatusUnknown:
		return "VolumePlexStatusUnknown"
	case VolumePlexStatusOnline:
		return "VolumePlexStatusOnline"
	case VolumePlexStatusNoMedia:
		return "VolumePlexStatusNoMedia"
	case VolumePlexStatusFailed:
		return "VolumePlexStatusFailed"
	}
	return "Invalid"
}

// IPAddressType type represents VDS_IPADDRESS_TYPE RPC enumeration.
//
// The VDS_IPADDRESS_TYPE enumeration defines the set of valid types for an IP address.
// These type values are used in the type member of the VDS_IPADDRESS structure.<40>
type IPAddressType uint16

var (
	// VDS_IPT_TEXT:  The IP address is a text string.
	IPAddressTypeText IPAddressType = 0
	// VDS_IPT_IPV4:  The IP address is an IPv4 address.
	IPAddressTypeIPv4 IPAddressType = 1
	// VDS_IPT_IPV6:  The IP address is an IPv6 address.
	IPAddressTypeIPv6 IPAddressType = 2
	// VDS_IPT_EMPTY:  An IP address is not specified.
	IPAddressTypeEmpty IPAddressType = 3
)

func (o IPAddressType) String() string {
	switch o {
	case IPAddressTypeText:
		return "IPAddressTypeText"
	case IPAddressTypeIPv4:
		return "IPAddressTypeIPv4"
	case IPAddressTypeIPv6:
		return "IPAddressTypeIPv6"
	case IPAddressTypeEmpty:
		return "IPAddressTypeEmpty"
	}
	return "Invalid"
}

// HBAPortType type represents VDS_HBAPORT_TYPE RPC enumeration.
//
// The VDS_HBAPORT_TYPE enumeration defines the set of valid types for an HBA port.
// These types correspond to the HBA_PORTTYPE values, as specified in [HBAAPI]. These
// values are used in the type member of the VDS_HBAPORT_PROP structure.<34>
type HBAPortType uint16

var (
	// VDS_HPT_UNKNOWN:  The port type is unknown.
	HBAPortTypeUnknown HBAPortType = 1
	// VDS_HPT_OTHER:  The port type is another (undefined) type.
	HBAPortTypeOther HBAPortType = 2
	// VDS_HPT_NOTPRESENT:  The port type is not present.
	HBAPortTypeNotPresent HBAPortType = 3
	// VDS_HPT_NPORT:  The port type is a fabric.
	HBAPortTypeN HBAPortType = 5
	// VDS_HPT_NLPORT:  The port type is a public loop.
	HBAPortTypeNL HBAPortType = 6
	// VDS_HPT_FLPORT:  The port type is a fabric on a loop.
	HBAPortTypeFL HBAPortType = 7
	// VDS_HPT_FPORT:  The port type is a fabric port.
	HBAPortTypeF HBAPortType = 8
	// VDS_HPT_EPORT:  The port type is a fabric expansion port.
	HBAPortTypeE HBAPortType = 9
	// VDS_HPT_GPORT:  The port type is a generic fabric port.
	HBAPortTypeG HBAPortType = 10
	// VDS_HPT_LPORT:  The port type is a private loop.
	HBAPortTypeL HBAPortType = 20
	// VDS_HPT_PTP:  The port type is point-to-point.
	HBAPortTypePTP HBAPortType = 21
)

func (o HBAPortType) String() string {
	switch o {
	case HBAPortTypeUnknown:
		return "HBAPortTypeUnknown"
	case HBAPortTypeOther:
		return "HBAPortTypeOther"
	case HBAPortTypeNotPresent:
		return "HBAPortTypeNotPresent"
	case HBAPortTypeN:
		return "HBAPortTypeN"
	case HBAPortTypeNL:
		return "HBAPortTypeNL"
	case HBAPortTypeFL:
		return "HBAPortTypeFL"
	case HBAPortTypeF:
		return "HBAPortTypeF"
	case HBAPortTypeE:
		return "HBAPortTypeE"
	case HBAPortTypeG:
		return "HBAPortTypeG"
	case HBAPortTypeL:
		return "HBAPortTypeL"
	case HBAPortTypePTP:
		return "HBAPortTypePTP"
	}
	return "Invalid"
}

// HBAPortStatus type represents VDS_HBAPORT_STATUS RPC enumeration.
//
// The VDS_HBAPORT_STATUS enumeration defines the set of valid statuses for an HBA port.
// These values are used in the status member of the VDS_HBAPORT_PROP structure. These
// states correspond to the HBA_PORTSTATE values, as specified in [HBAAPI].<35>
type HBAPortStatus uint16

var (
	// VDS_HPS_UNKNOWN:  The HBA port status is unknown.
	HBAPortStatusUnknown HBAPortStatus = 1
	// VDS_HPS_ONLINE:  The HBA port is operational.
	HBAPortStatusOnline HBAPortStatus = 2
	// VDS_HPS_OFFLINE:  The HBA port was set offline by a user.
	HBAPortStatusOffline HBAPortStatus = 3
	// VDS_HPS_BYPASSED:  The HBA port is bypassed.
	HBAPortStatusBypassed HBAPortStatus = 4
	// VDS_HPS_DIAGNOSTICS:  The HBA port is in diagnostics mode.
	HBAPortStatusDiagnostics HBAPortStatus = 5
	// VDS_HPS_LINKDOWN:  The HBA port link is down.
	HBAPortStatusLinkDown HBAPortStatus = 6
	// VDS_HPS_ERROR:  The HBA port has an error.
	HBAPortStatusError HBAPortStatus = 7
	// VDS_HPS_LOOPBACK:  The HBA port is loopback.
	HBAPortStatusLoopback HBAPortStatus = 8
)

func (o HBAPortStatus) String() string {
	switch o {
	case HBAPortStatusUnknown:
		return "HBAPortStatusUnknown"
	case HBAPortStatusOnline:
		return "HBAPortStatusOnline"
	case HBAPortStatusOffline:
		return "HBAPortStatusOffline"
	case HBAPortStatusBypassed:
		return "HBAPortStatusBypassed"
	case HBAPortStatusDiagnostics:
		return "HBAPortStatusDiagnostics"
	case HBAPortStatusLinkDown:
		return "HBAPortStatusLinkDown"
	case HBAPortStatusError:
		return "HBAPortStatusError"
	case HBAPortStatusLoopback:
		return "HBAPortStatusLoopback"
	}
	return "Invalid"
}

// HBAPortSpeedFlag type represents VDS_HBAPORT_SPEED_FLAG RPC enumeration.
//
// The VDS_HBAPORT_SPEED_FLAG enumeration type defines the set of valid flags for determining
// the speeds that an HBA port supports. These values are used in the ulPortSpeed member
// of the VDS_HBAPORT_PROP structure. These flags correspond to the HBA_PORTSPEED flags,
// as specified in [HBAAPI].<36>
type HBAPortSpeedFlag uint16

var (
	// VDS_HSF_UNKNOWN:  The HBA port speed is unknown.
	HBAPortSpeedFlagUnknown HBAPortSpeedFlag = 0
	// VDS_HSF_1GBIT:  The HBA port supports a transfer rate of 1 gigabit per second.
	HBAPortSpeedFlag1Gbit HBAPortSpeedFlag = 1
	// VDS_HSF_2GBIT:  The HBA port supports a transfer rate of 2 gigabits per second.
	HBAPortSpeedFlag2Gbit HBAPortSpeedFlag = 2
	// VDS_HSF_10GBIT:  The HBA port supports a transfer rate of 10 gigabits per second.
	HBAPortSpeedFlag10Gbit HBAPortSpeedFlag = 4
	// VDS_HSF_4GBIT:  The HBA port supports a transfer rate of 4 gigabits per second.
	HBAPortSpeedFlag4Gbit HBAPortSpeedFlag = 8
	// VDS_HSF_NOT_NEGOTIATED:  The HBA port speed has not been established.
	HBAPortSpeedFlagNotNegotiated HBAPortSpeedFlag = 32768
)

func (o HBAPortSpeedFlag) String() string {
	switch o {
	case HBAPortSpeedFlagUnknown:
		return "HBAPortSpeedFlagUnknown"
	case HBAPortSpeedFlag1Gbit:
		return "HBAPortSpeedFlag1Gbit"
	case HBAPortSpeedFlag2Gbit:
		return "HBAPortSpeedFlag2Gbit"
	case HBAPortSpeedFlag10Gbit:
		return "HBAPortSpeedFlag10Gbit"
	case HBAPortSpeedFlag4Gbit:
		return "HBAPortSpeedFlag4Gbit"
	case HBAPortSpeedFlagNotNegotiated:
		return "HBAPortSpeedFlagNotNegotiated"
	}
	return "Invalid"
}

// PathStatus type represents VDS_PATH_STATUS RPC enumeration.
//
// The VDS_PATH_STATUS enumeration defines the set of status values for a path to a
// storage device.
type PathStatus uint16

var (
	// VDS_MPS_UNKNOWN:  The status of the path is unknown.
	PathStatusUnknown PathStatus = 0
	// VDS_MPS_ONLINE:  The path is available.
	PathStatusOnline PathStatus = 1
	// VDS_MPS_FAILED:  The path is unavailable.
	PathStatusFailed PathStatus = 5
	// VDS_MPS_STANDBY:  The path is on standby; it is available but will not be used unless
	// other paths fail.
	PathStatusStandby PathStatus = 7
)

func (o PathStatus) String() string {
	switch o {
	case PathStatusUnknown:
		return "PathStatusUnknown"
	case PathStatusOnline:
		return "PathStatusOnline"
	case PathStatusFailed:
		return "PathStatusFailed"
	case PathStatusStandby:
		return "PathStatusStandby"
	}
	return "Invalid"
}

// ReparsePointProperty structure represents VDS_REPARSE_POINT_PROP RPC structure.
//
// The VDS_REPARSE_POINT_PROP structure defines the reparse point properties of the
// mount point to a volume object.
type ReparsePointProperty struct {
	// SourceVolumeId:  The VDS object ID of the volume object that the reparse point refers
	// to.
	SourceVolumeID *ObjectID `idl:"name:SourceVolumeId" json:"source_volume_id"`
	// pwszPath:  The null-terminated Unicode path of the reparse point. The path does not
	// contain a drive letter; for example, "\mount".
	Path string `idl:"name:pwszPath;string" json:"path"`
}

func (o *ReparsePointProperty) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReparsePointProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.SourceVolumeID != nil {
		if err := o.SourceVolumeID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Path != "" {
		_ptr_pwszPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Path, _ptr_pwszPath); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReparsePointProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.SourceVolumeID == nil {
		o.SourceVolumeID = &ObjectID{}
	}
	if err := o.SourceVolumeID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_pwszPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
			return err
		}
		return nil
	})
	_s_pwszPath := func(ptr interface{}) { o.Path = *ptr.(*string) }
	if err := w.ReadPointer(&o.Path, _s_pwszPath, _ptr_pwszPath); err != nil {
		return err
	}
	return nil
}

// DriveLetterProperty structure represents VDS_DRIVE_LETTER_PROP RPC structure.
//
// The VDS_DRIVE_LETTER_PROP structure provides information about a drive letter .
type DriveLetterProperty struct {
	// wcLetter:  The drive letter as a single uppercase or lowercase alphabetical (A-Z)
	// Unicode character.
	Letter uint16 `idl:"name:wcLetter" json:"letter"`
	// volumeId:  The VDS object ID of the volume object to which the drive letter is assigned.
	// If the drive letter is not assigned to any volume, the value MUST be GUID_NULL.
	VolumeID *ObjectID `idl:"name:volumeId" json:"volume_id"`
	// ulFlags:  The combination of any values, by using a bitwise OR operator, that is
	// defined in the VDS_DRIVE_LETTER_FLAG enumeration.
	Flags uint32 `idl:"name:ulFlags" json:"flags"`
	// bUsed:  A Boolean value that indicates whether the drive letter is already in use.
	Used int32 `idl:"name:bUsed" json:"used"`
}

func (o *DriveLetterProperty) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DriveLetterProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Letter); err != nil {
		return err
	}
	if o.VolumeID != nil {
		if err := o.VolumeID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.Used); err != nil {
		return err
	}
	return nil
}
func (o *DriveLetterProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Letter); err != nil {
		return err
	}
	if o.VolumeID == nil {
		o.VolumeID = &ObjectID{}
	}
	if err := o.VolumeID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.Used); err != nil {
		return err
	}
	return nil
}

// SANPolicy type represents VDS_SAN_POLICY RPC enumeration.
//
// The VDS_SAN_POLICY enumeration defines the set of valid SAN policy values.
type SANPolicy uint16

var (
	// VDS_SP_UNKNOWN:  The SAN policy is unknown.
	SANPolicyUnknown SANPolicy = 0
	// VDS_SP_ONLINE:  All newly discovered disks are brought online and made WRITABLE.
	// If the disk is offline, no volume devices are exposed for the disk. If the disk is
	// online, the volume devices for the disk are exposed. WRITABLE is the normal state
	// for a disk. A disk can also be made READONLY. If the disk is READONLY, disk data
	// and metadata can be read, but writes to the disk will fail.
	SANPolicyOnline SANPolicy = 1
	// VDS_SP_OFFLINE_SHARED:  All newly discovered disks not residing on a shared bus
	// are brought online and made WRITABLE. If the disk is offline, no volume devices are
	// exposed for the disk. If the disk is online, the volume devices for the disk are
	// exposed. WRITABLE is the normal state for a disk. A disk can also be made READONLY.
	// If the disk is READONLY, disk data and metadata can be read, but writes to the disk
	// will fail.
	SANPolicyOfflineShared SANPolicy = 2
	// VDS_SP_OFFLINE:  All newly discovered disks remain offline and READONLY. If the
	// disk is offline, no volume devices are exposed for the disk. If the disk is online,
	// the volume devices for the disk are exposed. WRITABLE is the normal state for a disk.
	// A disk can also be made READONLY. If the disk is READONLY, disk data and metadata
	// can be read, but writes to the disk will fail.
	SANPolicyOffline SANPolicy = 3
	// VDS_SP_OFFLINE_INTERNAL<31>:  All newly discovered internal disks remain offline
	// and READONLY. If the disk is offline, no volume devices are exposed for the disk.
	// If the disk is online, the volume devices for the disk are exposed. WRITABLE is the
	// normal state for a disk. A disk can also be made READONLY. If the disk is READONLY,
	// disk data and metadata can be read, but writes to the disk will fail.
	SANPolicyOfflineInternal SANPolicy = 4
	// VDS_SP_MAX<32>:  Denotes the maximum acceptable value for this type. VDS_SP_MAX
	// - 1, ('4'), is the maximum acceptable value.
	SANPolicyMax SANPolicy = 5
)

func (o SANPolicy) String() string {
	switch o {
	case SANPolicyUnknown:
		return "SANPolicyUnknown"
	case SANPolicyOnline:
		return "SANPolicyOnline"
	case SANPolicyOfflineShared:
		return "SANPolicyOfflineShared"
	case SANPolicyOffline:
		return "SANPolicyOffline"
	case SANPolicyOfflineInternal:
		return "SANPolicyOfflineInternal"
	case SANPolicyMax:
		return "SANPolicyMax"
	}
	return "Invalid"
}

// FileSystemTypeProperty structure represents VDS_FILE_SYSTEM_TYPE_PROP RPC structure.
//
// The VDS_FILE_SYSTEM_TYPE_PROP structure provides information about a file system
// format.<30>
type FileSystemTypeProperty struct {
	// type:  A value from the VDS_FILE_SYSTEM_TYPE enumeration that indicates the file
	// system format type.
	Type FileSystemType `idl:"name:type" json:"type"`
	// wszName:  A null-terminated Unicode name of the file system format, for example,
	// NTFS or FAT32.
	Name []uint16 `idl:"name:wszName" json:"name"`
	// ulFlags:  A combination of any values, by using a bitwise OR operator, that are defined
	// in the VDS_FILE_SYSTEM_FLAG enumeration.
	Flags uint32 `idl:"name:ulFlags" json:"flags"`
	// ulCompressionFlags:  A bitwise OR operator of any allocation units that are defined
	// in the VDS_FILE_SYSTEM_PROP_FLAG enumeration.
	CompressionFlags uint32 `idl:"name:ulCompressionFlags" json:"compression_flags"`
	// ulMaxLabelLength:  The maximum allowable length of a label for the file system format.
	MaxLabelLength uint32 `idl:"name:ulMaxLabelLength" json:"max_label_length"`
	// pwszIllegalLabelCharSet:  A null-terminated sequence of Unicode characters that are
	// not allowed in the label of the file system format.
	IllegalLabelCharSet string `idl:"name:pwszIllegalLabelCharSet;string" json:"illegal_label_char_set"`
}

func (o *FileSystemTypeProperty) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *FileSystemTypeProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Type)); err != nil {
		return err
	}
	for i1 := range o.Name {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.Name[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Name); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.CompressionFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxLabelLength); err != nil {
		return err
	}
	if o.IllegalLabelCharSet != "" {
		_ptr_pwszIllegalLabelCharSet := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.IllegalLabelCharSet); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.IllegalLabelCharSet, _ptr_pwszIllegalLabelCharSet); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *FileSystemTypeProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Type)); err != nil {
		return err
	}
	o.Name = make([]uint16, 8)
	for i1 := range o.Name {
		i1 := i1
		if err := w.ReadData(&o.Name[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.CompressionFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxLabelLength); err != nil {
		return err
	}
	_ptr_pwszIllegalLabelCharSet := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.IllegalLabelCharSet); err != nil {
			return err
		}
		return nil
	})
	_s_pwszIllegalLabelCharSet := func(ptr interface{}) { o.IllegalLabelCharSet = *ptr.(*string) }
	if err := w.ReadPointer(&o.IllegalLabelCharSet, _s_pwszIllegalLabelCharSet, _ptr_pwszIllegalLabelCharSet); err != nil {
		return err
	}
	return nil
}

// ChangeAttributesParameters structure represents CHANGE_ATTRIBUTES_PARAMETERS RPC structure.
//
// The CHANGE_ATTRIBUTES_PARAMETERS structure describes the attributes to change on
// a partition.
type ChangeAttributesParameters struct {
	// style:  The value from the VDS_PARTITION_STYLE enumeration that describes the partition
	// format of the disk. If the disk partitioning format is MBR, the only value that can
	// be changed is the bootIndicator. If the disk partitioning format is GPT, the only
	// value that can be changed is the GPT attribute.
	Style                      PartitionStyle                                         `idl:"name:style" json:"style"`
	ChangeAttributesParameters *ChangeAttributesParameters_ChangeAttributesParameters `idl:"name:change_attributes_parameters;switch_is:style" json:"change_attributes_parameters"`
}

func (o *ChangeAttributesParameters) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ChangeAttributesParameters) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Style)); err != nil {
		return err
	}
	_swChangeAttributesParameters := uint16(o.Style)
	if o.ChangeAttributesParameters != nil {
		if err := o.ChangeAttributesParameters.MarshalUnionNDR(ctx, w, _swChangeAttributesParameters); err != nil {
			return err
		}
	} else {
		if err := (&ChangeAttributesParameters_ChangeAttributesParameters{}).MarshalUnionNDR(ctx, w, _swChangeAttributesParameters); err != nil {
			return err
		}
	}
	return nil
}
func (o *ChangeAttributesParameters) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Style)); err != nil {
		return err
	}
	if o.ChangeAttributesParameters == nil {
		o.ChangeAttributesParameters = &ChangeAttributesParameters_ChangeAttributesParameters{}
	}
	_swChangeAttributesParameters := uint16(o.Style)
	if err := o.ChangeAttributesParameters.UnmarshalUnionNDR(ctx, w, _swChangeAttributesParameters); err != nil {
		return err
	}
	return nil
}

// ChangeAttributesParameters_ChangeAttributesParameters structure represents CHANGE_ATTRIBUTES_PARAMETERS union anonymous member.
//
// The CHANGE_ATTRIBUTES_PARAMETERS structure describes the attributes to change on
// a partition.
type ChangeAttributesParameters_ChangeAttributesParameters struct {
	// Types that are assignable to Value
	//
	// *ChangeAttributesParameters_MBRPartitionInfo
	// *ChangeAttributesParameters_GPTPartitionInfo
	// *ChangeAttributesParameters_DefaultChangeAttributesParameters
	Value is_ChangeAttributesParameters_ChangeAttributesParameters `json:"value"`
}

func (o *ChangeAttributesParameters_ChangeAttributesParameters) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *ChangeAttributesParameters_MBRPartitionInfo:
		if value != nil {
			return value.MBRPartitionInfo
		}
	case *ChangeAttributesParameters_GPTPartitionInfo:
		if value != nil {
			return value.GPTPartitionInfo
		}
	}
	return nil
}

type is_ChangeAttributesParameters_ChangeAttributesParameters interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_ChangeAttributesParameters_ChangeAttributesParameters()
}

func (o *ChangeAttributesParameters_ChangeAttributesParameters) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *ChangeAttributesParameters_MBRPartitionInfo:
		return uint16(1)
	case *ChangeAttributesParameters_GPTPartitionInfo:
		return uint16(2)
	}
	return uint16(0)
}

func (o *ChangeAttributesParameters_ChangeAttributesParameters) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		_o, _ := o.Value.(*ChangeAttributesParameters_MBRPartitionInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ChangeAttributesParameters_MBRPartitionInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*ChangeAttributesParameters_GPTPartitionInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ChangeAttributesParameters_GPTPartitionInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		// no-op
	}
	return nil
}

func (o *ChangeAttributesParameters_ChangeAttributesParameters) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch((*uint16)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		o.Value = &ChangeAttributesParameters_MBRPartitionInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &ChangeAttributesParameters_GPTPartitionInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		o.Value = &ChangeAttributesParameters_DefaultChangeAttributesParameters{}
	}
	return nil
}

// ChangeAttributesParameters_MBRPartitionInfo structure represents ChangeAttributesParameters_ChangeAttributesParameters RPC union arm.
//
// It has following labels: 1
type ChangeAttributesParameters_MBRPartitionInfo struct {
	// MbrPartInfo:  Contains information for an MBR partition.
	MBRPartitionInfo *ChangeAttributesParameters_ChangeAttributesParameters_MBRPartitionInfo `idl:"name:MbrPartInfo" json:"mbr_partition_info"`
}

func (*ChangeAttributesParameters_MBRPartitionInfo) is_ChangeAttributesParameters_ChangeAttributesParameters() {
}

func (o *ChangeAttributesParameters_MBRPartitionInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.MBRPartitionInfo != nil {
		if err := o.MBRPartitionInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ChangeAttributesParameters_ChangeAttributesParameters_MBRPartitionInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ChangeAttributesParameters_MBRPartitionInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.MBRPartitionInfo == nil {
		o.MBRPartitionInfo = &ChangeAttributesParameters_ChangeAttributesParameters_MBRPartitionInfo{}
	}
	if err := o.MBRPartitionInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ChangeAttributesParameters_ChangeAttributesParameters_MBRPartitionInfo structure represents CHANGE_ATTRIBUTES_PARAMETERS structure anonymous member.
//
// The CHANGE_ATTRIBUTES_PARAMETERS structure describes the attributes to change on
// a partition.
type ChangeAttributesParameters_ChangeAttributesParameters_MBRPartitionInfo struct {
	// bootIndicator:  The Boolean value that indicates whether the partition is bootable.
	BootIndicator bool `idl:"name:bootIndicator" json:"boot_indicator"`
}

func (o *ChangeAttributesParameters_ChangeAttributesParameters_MBRPartitionInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ChangeAttributesParameters_ChangeAttributesParameters_MBRPartitionInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteData(o.BootIndicator); err != nil {
		return err
	}
	return nil
}
func (o *ChangeAttributesParameters_ChangeAttributesParameters_MBRPartitionInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.BootIndicator); err != nil {
		return err
	}
	return nil
}

// ChangeAttributesParameters_GPTPartitionInfo structure represents ChangeAttributesParameters_ChangeAttributesParameters RPC union arm.
//
// It has following labels: 2
type ChangeAttributesParameters_GPTPartitionInfo struct {
	// GptPartInfo:  Contains information for a partition in a GPT.
	GPTPartitionInfo *ChangeAttributesParameters_ChangeAttributesParameters_GPTPartitionInfo `idl:"name:GptPartInfo" json:"gpt_partition_info"`
}

func (*ChangeAttributesParameters_GPTPartitionInfo) is_ChangeAttributesParameters_ChangeAttributesParameters() {
}

func (o *ChangeAttributesParameters_GPTPartitionInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.GPTPartitionInfo != nil {
		if err := o.GPTPartitionInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ChangeAttributesParameters_ChangeAttributesParameters_GPTPartitionInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ChangeAttributesParameters_GPTPartitionInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.GPTPartitionInfo == nil {
		o.GPTPartitionInfo = &ChangeAttributesParameters_ChangeAttributesParameters_GPTPartitionInfo{}
	}
	if err := o.GPTPartitionInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ChangeAttributesParameters_ChangeAttributesParameters_GPTPartitionInfo structure represents CHANGE_ATTRIBUTES_PARAMETERS structure anonymous member.
//
// The CHANGE_ATTRIBUTES_PARAMETERS structure describes the attributes to change on
// a partition.
type ChangeAttributesParameters_ChangeAttributesParameters_GPTPartitionInfo struct {
	// attributes:  The bitwise OR operator of attributes to change; it can have a combination
	// of the following values.
	//
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                                             |                                                                                  |
	//	|                            VALUE                            |                                     MEANING                                      |
	//	|                                                             |                                                                                  |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| GPT_ATTRIBUTE_PLATFORM_REQUIRED 0x0000000000000001          | Partition is required for the platform to function properly.<46>                 |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| GPT_BASIC_DATA_ATTRIBUTE_READ_ONLY 0x1000000000000000       | The partition can be read from but not written to. Used only with the basic data |
	//	|                                                             | partition type.                                                                  |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| GPT_BASIC_DATA_ATTRIBUTE_HIDDEN 0x4000000000000000          | The partition is hidden and is not mounted. Used only with the basic data        |
	//	|                                                             | partition type.                                                                  |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| GPT_BASIC_DATA_ATTRIBUTE_NO_DRIVE_LETTER 0x8000000000000000 | The partition does not receive a drive letter by default when moving the disk to |
	//	|                                                             | another machine. Used only with the basic data partition type.                   |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	Attributes uint64 `idl:"name:attributes" json:"attributes"`
}

func (o *ChangeAttributesParameters_ChangeAttributesParameters_GPTPartitionInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ChangeAttributesParameters_ChangeAttributesParameters_GPTPartitionInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	return nil
}
func (o *ChangeAttributesParameters_ChangeAttributesParameters_GPTPartitionInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	return nil
}

// ChangeAttributesParameters_DefaultChangeAttributesParameters structure represents ChangeAttributesParameters_ChangeAttributesParameters RPC default union arm.
type ChangeAttributesParameters_DefaultChangeAttributesParameters struct {
}

func (*ChangeAttributesParameters_DefaultChangeAttributesParameters) is_ChangeAttributesParameters_ChangeAttributesParameters() {
}

func (o *ChangeAttributesParameters_DefaultChangeAttributesParameters) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return nil
}
func (o *ChangeAttributesParameters_DefaultChangeAttributesParameters) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	return nil
}

// ChangePartitionTypeParameters structure represents CHANGE_PARTITION_TYPE_PARAMETERS RPC structure.
//
// The CHANGE_PARTITION_TYPE_PARAMETERS structure describes parameters to use when changing
// a partition type.<47>
type ChangePartitionTypeParameters struct {
	// style:  A value from the VDS_PARITION_STYLE enumeration that describes the disk partition
	// format.
	Style                         PartitionStyle                                               `idl:"name:style" json:"style"`
	ChangePartitionTypeParameters *ChangePartitionTypeParameters_ChangePartitionTypeParameters `idl:"name:change_partition_type_parameters;switch_is:style" json:"change_partition_type_parameters"`
}

func (o *ChangePartitionTypeParameters) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ChangePartitionTypeParameters) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Style)); err != nil {
		return err
	}
	_swChangePartitionTypeParameters := uint16(o.Style)
	if o.ChangePartitionTypeParameters != nil {
		if err := o.ChangePartitionTypeParameters.MarshalUnionNDR(ctx, w, _swChangePartitionTypeParameters); err != nil {
			return err
		}
	} else {
		if err := (&ChangePartitionTypeParameters_ChangePartitionTypeParameters{}).MarshalUnionNDR(ctx, w, _swChangePartitionTypeParameters); err != nil {
			return err
		}
	}
	return nil
}
func (o *ChangePartitionTypeParameters) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Style)); err != nil {
		return err
	}
	if o.ChangePartitionTypeParameters == nil {
		o.ChangePartitionTypeParameters = &ChangePartitionTypeParameters_ChangePartitionTypeParameters{}
	}
	_swChangePartitionTypeParameters := uint16(o.Style)
	if err := o.ChangePartitionTypeParameters.UnmarshalUnionNDR(ctx, w, _swChangePartitionTypeParameters); err != nil {
		return err
	}
	return nil
}

// ChangePartitionTypeParameters_ChangePartitionTypeParameters structure represents CHANGE_PARTITION_TYPE_PARAMETERS union anonymous member.
//
// The CHANGE_PARTITION_TYPE_PARAMETERS structure describes parameters to use when changing
// a partition type.<47>
type ChangePartitionTypeParameters_ChangePartitionTypeParameters struct {
	// Types that are assignable to Value
	//
	// *ChangePartitionTypeParameters_MBRPartitionInfo
	// *ChangePartitionTypeParameters_GPTPartitionInfo
	// *ChangePartitionTypeParameters_DefaultChangePartitionTypeParameters
	Value is_ChangePartitionTypeParameters_ChangePartitionTypeParameters `json:"value"`
}

func (o *ChangePartitionTypeParameters_ChangePartitionTypeParameters) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *ChangePartitionTypeParameters_MBRPartitionInfo:
		if value != nil {
			return value.MBRPartitionInfo
		}
	case *ChangePartitionTypeParameters_GPTPartitionInfo:
		if value != nil {
			return value.GPTPartitionInfo
		}
	}
	return nil
}

type is_ChangePartitionTypeParameters_ChangePartitionTypeParameters interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_ChangePartitionTypeParameters_ChangePartitionTypeParameters()
}

func (o *ChangePartitionTypeParameters_ChangePartitionTypeParameters) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *ChangePartitionTypeParameters_MBRPartitionInfo:
		return uint16(1)
	case *ChangePartitionTypeParameters_GPTPartitionInfo:
		return uint16(2)
	}
	return uint16(0)
}

func (o *ChangePartitionTypeParameters_ChangePartitionTypeParameters) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		_o, _ := o.Value.(*ChangePartitionTypeParameters_MBRPartitionInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ChangePartitionTypeParameters_MBRPartitionInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*ChangePartitionTypeParameters_GPTPartitionInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ChangePartitionTypeParameters_GPTPartitionInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		// no-op
	}
	return nil
}

func (o *ChangePartitionTypeParameters_ChangePartitionTypeParameters) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch((*uint16)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		o.Value = &ChangePartitionTypeParameters_MBRPartitionInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &ChangePartitionTypeParameters_GPTPartitionInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		o.Value = &ChangePartitionTypeParameters_DefaultChangePartitionTypeParameters{}
	}
	return nil
}

// ChangePartitionTypeParameters_MBRPartitionInfo structure represents ChangePartitionTypeParameters_ChangePartitionTypeParameters RPC union arm.
//
// It has following labels: 1
type ChangePartitionTypeParameters_MBRPartitionInfo struct {
	// MbrPartInfo:  Contains information for an MBR partition.
	MBRPartitionInfo *ChangePartitionTypeParameters_ChangePartitionTypeParameters_MBRPartitionInfo `idl:"name:MbrPartInfo" json:"mbr_partition_info"`
}

func (*ChangePartitionTypeParameters_MBRPartitionInfo) is_ChangePartitionTypeParameters_ChangePartitionTypeParameters() {
}

func (o *ChangePartitionTypeParameters_MBRPartitionInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.MBRPartitionInfo != nil {
		if err := o.MBRPartitionInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ChangePartitionTypeParameters_ChangePartitionTypeParameters_MBRPartitionInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ChangePartitionTypeParameters_MBRPartitionInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.MBRPartitionInfo == nil {
		o.MBRPartitionInfo = &ChangePartitionTypeParameters_ChangePartitionTypeParameters_MBRPartitionInfo{}
	}
	if err := o.MBRPartitionInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ChangePartitionTypeParameters_ChangePartitionTypeParameters_MBRPartitionInfo structure represents CHANGE_PARTITION_TYPE_PARAMETERS structure anonymous member.
//
// The CHANGE_PARTITION_TYPE_PARAMETERS structure describes parameters to use when changing
// a partition type.<47>
type ChangePartitionTypeParameters_ChangePartitionTypeParameters_MBRPartitionInfo struct {
	// partitionType:  The byte value indicating the partition type to change the partition
	// to.
	PartitionType uint8 `idl:"name:partitionType" json:"partition_type"`
}

func (o *ChangePartitionTypeParameters_ChangePartitionTypeParameters_MBRPartitionInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ChangePartitionTypeParameters_ChangePartitionTypeParameters_MBRPartitionInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteData(o.PartitionType); err != nil {
		return err
	}
	return nil
}
func (o *ChangePartitionTypeParameters_ChangePartitionTypeParameters_MBRPartitionInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.PartitionType); err != nil {
		return err
	}
	return nil
}

// ChangePartitionTypeParameters_GPTPartitionInfo structure represents ChangePartitionTypeParameters_ChangePartitionTypeParameters RPC union arm.
//
// It has following labels: 2
type ChangePartitionTypeParameters_GPTPartitionInfo struct {
	// GptPartInfo:  Contains information for the partition of a GPT.
	GPTPartitionInfo *ChangePartitionTypeParameters_ChangePartitionTypeParameters_GPTPartitionInfo `idl:"name:GptPartInfo" json:"gpt_partition_info"`
}

func (*ChangePartitionTypeParameters_GPTPartitionInfo) is_ChangePartitionTypeParameters_ChangePartitionTypeParameters() {
}

func (o *ChangePartitionTypeParameters_GPTPartitionInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.GPTPartitionInfo != nil {
		if err := o.GPTPartitionInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ChangePartitionTypeParameters_ChangePartitionTypeParameters_GPTPartitionInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ChangePartitionTypeParameters_GPTPartitionInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.GPTPartitionInfo == nil {
		o.GPTPartitionInfo = &ChangePartitionTypeParameters_ChangePartitionTypeParameters_GPTPartitionInfo{}
	}
	if err := o.GPTPartitionInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ChangePartitionTypeParameters_ChangePartitionTypeParameters_GPTPartitionInfo structure represents CHANGE_PARTITION_TYPE_PARAMETERS structure anonymous member.
//
// The CHANGE_PARTITION_TYPE_PARAMETERS structure describes parameters to use when changing
// a partition type.<47>
type ChangePartitionTypeParameters_ChangePartitionTypeParameters_GPTPartitionInfo struct {
	// partitionType:  The byte value indicating the partition type to change the partition
	// to.
	PartitionType *dtyp.GUID `idl:"name:partitionType" json:"partition_type"`
}

func (o *ChangePartitionTypeParameters_ChangePartitionTypeParameters_GPTPartitionInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ChangePartitionTypeParameters_ChangePartitionTypeParameters_GPTPartitionInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.PartitionType != nil {
		if err := o.PartitionType.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ChangePartitionTypeParameters_ChangePartitionTypeParameters_GPTPartitionInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.PartitionType == nil {
		o.PartitionType = &dtyp.GUID{}
	}
	if err := o.PartitionType.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ChangePartitionTypeParameters_DefaultChangePartitionTypeParameters structure represents ChangePartitionTypeParameters_ChangePartitionTypeParameters RPC default union arm.
type ChangePartitionTypeParameters_DefaultChangePartitionTypeParameters struct {
}

func (*ChangePartitionTypeParameters_DefaultChangePartitionTypeParameters) is_ChangePartitionTypeParameters_ChangePartitionTypeParameters() {
}

func (o *ChangePartitionTypeParameters_DefaultChangePartitionTypeParameters) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return nil
}
func (o *ChangePartitionTypeParameters_DefaultChangePartitionTypeParameters) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	return nil
}

// WWN structure represents VDS_WWN RPC structure.
//
// The VDS_WWN structure defines a worldwide name (WWN). This structure corresponds
// to the HBA_WWN structure, as specified in [HBAAPI], which also defines the WWN term.<37>
type WWN struct {
	// rguchWwn:  An array of 8 bytes that specifies the 64-bit WWN value. The first element
	// of the array is the most significant byte of the WWN, and the most significant bit
	// of that byte is the most significant bit of the WWN.
	WWN []byte `idl:"name:rguchWwn" json:"wwn"`
}

func (o *WWN) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *WWN) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	for i1 := range o.WWN {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.WWN[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.WWN); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *WWN) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	o.WWN = make([]byte, 8)
	for i1 := range o.WWN {
		i1 := i1
		if err := w.ReadData(&o.WWN[i1]); err != nil {
			return err
		}
	}
	return nil
}

// IPAddress structure represents VDS_IPADDRESS RPC structure.
//
// The VDS_IPADDRESS structure defines an IP address and port.<41>
type IPAddress struct {
	// type:  The type of address as enumerated by VDS_IPADDRESS_TYPE.
	Type IPAddressType `idl:"name:type" json:"type"`
	// ipv4Address:  If the type member is VDS_IPT_IPV4, this member contains the binary
	// IPv4 address in network byte order. The field 3 (last octet) byte value is contained
	// in bits 0 through 7. The byte value for field 2 is contained in bits 8 through 15.
	// The byte value for field 1 is contained in bits 16 through 23. The byte value for
	// field 0 is contained in bits 24 through 31. Otherwise, this value is ignored.
	IPv4Address uint32 `idl:"name:ipv4Address" json:"ipv4_address"`
	// ipv6Address:  If the type member is VDS_IPT_IPV6, this member contains the binary
	// IPv6 address in network byte order. Otherwise, this value is ignored.
	IPv6Address []byte `idl:"name:ipv6Address" json:"ipv6_address"`
	// ulIpv6FlowInfo:  If the type member is VDS_IPT_IPV6, this member contains the flow
	// information as defined in IPv6. Otherwise, this value is ignored.
	IPv6FlowInfo uint32 `idl:"name:ulIpv6FlowInfo" json:"ipv6_flow_info"`
	// ulIpv6ScopeId:  If the type member is VDS_IPT_IPV6, this member contains the scope
	// ID as defined in IPv6. Otherwise, this value is ignored.
	IPv6ScopeID uint32 `idl:"name:ulIpv6ScopeId" json:"ipv6_scope_id"`
	// wszTextAddress:  If the type member is VDS_IPT_TEXT, this member contains the null-terminated
	// Unicode text address, which is either a DNS address, an IPv4 dotted address, or an
	// IPv6 hexadecimal address. Otherwise, this value is ignored.
	TextAddress []uint16 `idl:"name:wszTextAddress" json:"text_address"`
	// ulPort:  If the type member is VDS_IPT_IPV4, VDS_IPT_IPV6, or VDS_IPT_TEXT, this
	// member contains the TCP port number. Otherwise, this value is ignored.
	Port uint32 `idl:"name:ulPort" json:"port"`
}

func (o *IPAddress) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPAddress) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Type)); err != nil {
		return err
	}
	if err := w.WriteData(o.IPv4Address); err != nil {
		return err
	}
	for i1 := range o.IPv6Address {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.IPv6Address[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.IPv6Address); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.IPv6FlowInfo); err != nil {
		return err
	}
	if err := w.WriteData(o.IPv6ScopeID); err != nil {
		return err
	}
	for i1 := range o.TextAddress {
		i1 := i1
		if uint64(i1) >= 257 {
			break
		}
		if err := w.WriteData(o.TextAddress[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.TextAddress); uint64(i1) < 257; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Port); err != nil {
		return err
	}
	return nil
}
func (o *IPAddress) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Type)); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPv4Address); err != nil {
		return err
	}
	o.IPv6Address = make([]byte, 16)
	for i1 := range o.IPv6Address {
		i1 := i1
		if err := w.ReadData(&o.IPv6Address[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.IPv6FlowInfo); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPv6ScopeID); err != nil {
		return err
	}
	o.TextAddress = make([]uint16, 257)
	for i1 := range o.TextAddress {
		i1 := i1
		if err := w.ReadData(&o.TextAddress[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.Port); err != nil {
		return err
	}
	return nil
}

// ISCSISharedSecret structure represents VDS_ISCSI_SHARED_SECRET RPC structure.
//
// The VDS_ISCSI_SHARED_SECRET structure defines the Challenge-Handshake Authentication
// Protocol (CHAP), as specified in [MS-CHAP], shared secret for an iSCSI initiator.
type ISCSISharedSecret struct {
	// pSharedSecret:  A pointer to an array of bytes that contains the secret.
	SharedSecret []byte `idl:"name:pSharedSecret;size_is:(ulSharedSecretSize)" json:"shared_secret"`
	// ulSharedSecretSize:  The number of bytes contained in the array that pSharedSecret
	// references. Bytes MUST be at least 12 and less than or equal to 16.<33> If a shared
	// secret of size less than 12 bytes is used, the server does not return an error. However,
	// the operation will not complete.
	SharedSecretSize uint32 `idl:"name:ulSharedSecretSize" json:"shared_secret_size"`
}

func (o *ISCSISharedSecret) xxx_PreparePayload(ctx context.Context) error {
	if o.SharedSecret != nil && o.SharedSecretSize == 0 {
		o.SharedSecretSize = uint32(len(o.SharedSecret))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ISCSISharedSecret) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.SharedSecret != nil || o.SharedSecretSize > 0 {
		_ptr_pSharedSecret := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.SharedSecretSize)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.SharedSecret {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.SharedSecret[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.SharedSecret); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
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
	if err := w.WriteData(o.SharedSecretSize); err != nil {
		return err
	}
	return nil
}
func (o *ISCSISharedSecret) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_pSharedSecret := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.SharedSecretSize > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.SharedSecretSize)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.SharedSecret", sizeInfo[0])
		}
		o.SharedSecret = make([]byte, sizeInfo[0])
		for i1 := range o.SharedSecret {
			i1 := i1
			if err := w.ReadData(&o.SharedSecret[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pSharedSecret := func(ptr interface{}) { o.SharedSecret = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.SharedSecret, _s_pSharedSecret, _ptr_pSharedSecret); err != nil {
		return err
	}
	if err := w.ReadData(&o.SharedSecretSize); err != nil {
		return err
	}
	return nil
}

// ServiceProperty structure represents VDS_SERVICE_PROP RPC structure.
//
// The VDS_SERVICE_PROP structure provides information about the properties of a service.
type ServiceProperty struct {
	// pwszVersion:  The version of VDS; a human-readable, null-terminated Unicode string.
	// This string can be any human-readable, null-terminated Unicode value.<29>
	Version string `idl:"name:pwszVersion;string" json:"version"`
	// ulFlags:  A combination of any values, by using the bitwise OR operator, that is
	// defined in the VDS_SERVICE_FLAG enumeration.
	Flags uint32 `idl:"name:ulFlags" json:"flags"`
}

func (o *ServiceProperty) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ServiceProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Version != "" {
		_ptr_pwszVersion := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Version); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Version, _ptr_pwszVersion); err != nil {
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
	return nil
}
func (o *ServiceProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_pwszVersion := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Version); err != nil {
			return err
		}
		return nil
	})
	_s_pwszVersion := func(ptr interface{}) { o.Version = *ptr.(*string) }
	if err := w.ReadPointer(&o.Version, _s_pwszVersion, _ptr_pwszVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	return nil
}

// HBAPortProperty structure represents VDS_HBAPORT_PROP RPC structure.
//
// The VDS_HBAPORT_PROP structure defines the properties of an HBA port.<38>
type HBAPortProperty struct {
	// id:  The VDS object ID of the HBA port object.
	ID *ObjectID `idl:"name:id" json:"id"`
	// wwnNode:  The node WWN for the HBA port.
	Node *WWN `idl:"name:wwnNode" json:"node"`
	// wwnPort:  The port WWN of the HBA port.
	Port *WWN `idl:"name:wwnPort" json:"port"`
	// type:  The type of the HBA port that VDS_HBAPORT_TYPE enumerates.
	Type HBAPortType `idl:"name:type" json:"type"`
	// status:  The status of the HBA port that VDS_HBAPORT_STATUS enumerates.
	Status HBAPortStatus `idl:"name:status" json:"status"`
	// ulPortSpeed:  The speed of the HBA port that VDS_HBAPORT_SPEED_FLAG enumerates. Only
	// one bit can be set in this bitmask.
	PortSpeed uint32 `idl:"name:ulPortSpeed" json:"port_speed"`
	// ulSupportedPortSpeed:  The combination of values, by using a bitwise OR operator,
	// from the VDS_HBAPORT_SPEED_FLAG enumeration that describes the set of supported speeds
	// of the HBA port.
	SupportedPortSpeed uint32 `idl:"name:ulSupportedPortSpeed" json:"supported_port_speed"`
}

func (o *HBAPortProperty) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *HBAPortProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.ID != nil {
		if err := o.ID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Node != nil {
		if err := o.Node.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&WWN{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Port != nil {
		if err := o.Port.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&WWN{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint16(o.Type)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Status)); err != nil {
		return err
	}
	if err := w.WriteData(o.PortSpeed); err != nil {
		return err
	}
	if err := w.WriteData(o.SupportedPortSpeed); err != nil {
		return err
	}
	return nil
}
func (o *HBAPortProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.ID == nil {
		o.ID = &ObjectID{}
	}
	if err := o.ID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Node == nil {
		o.Node = &WWN{}
	}
	if err := o.Node.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Port == nil {
		o.Port = &WWN{}
	}
	if err := o.Port.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Type)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Status)); err != nil {
		return err
	}
	if err := w.ReadData(&o.PortSpeed); err != nil {
		return err
	}
	if err := w.ReadData(&o.SupportedPortSpeed); err != nil {
		return err
	}
	return nil
}

// ISCSIInitiatorAdapterProperty structure represents VDS_ISCSI_INITIATOR_ADAPTER_PROP RPC structure.
//
// The VDS_ISCSI_INITIATOR_ADAPTER_PROP structure defines the properties of an iSCSI
// initiator adapter.<39>
type ISCSIInitiatorAdapterProperty struct {
	// id:  The VDS object ID of the initiator adapter object.
	ID *ObjectID `idl:"name:id" json:"id"`
	// pwszName:  A human-readable, null-terminated Unicode string that is the name of the
	// initiator adapter.
	Name string `idl:"name:pwszName;string" json:"name"`
}

func (o *ISCSIInitiatorAdapterProperty) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ISCSIInitiatorAdapterProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.ID != nil {
		if err := o.ID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Name != "" {
		_ptr_pwszName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_pwszName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ISCSIInitiatorAdapterProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.ID == nil {
		o.ID = &ObjectID{}
	}
	if err := o.ID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_pwszName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_pwszName := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_pwszName, _ptr_pwszName); err != nil {
		return err
	}
	return nil
}

// ISCSIInitiatorPortalProperty structure represents VDS_ISCSI_INITIATOR_PORTAL_PROP RPC structure.
//
// The VDS_ISCSI_INITIATOR_PORTAL_PROP structure defines the properties of an iSCSI
// initiator portal.<42>
type ISCSIInitiatorPortalProperty struct {
	// id:  The VDS object ID of the initiator portal object.
	ID *ObjectID `idl:"name:id" json:"id"`
	// address:  The IP address and port of the portal.
	Address *IPAddress `idl:"name:address" json:"address"`
	// ulPortIndex:  The port index that the iSCSI initiators service assigned to the portal.
	PortIndex uint32 `idl:"name:ulPortIndex" json:"port_index"`
}

func (o *ISCSIInitiatorPortalProperty) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ISCSIInitiatorPortalProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.ID != nil {
		if err := o.ID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Address != nil {
		if err := o.Address.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IPAddress{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PortIndex); err != nil {
		return err
	}
	return nil
}
func (o *ISCSIInitiatorPortalProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.ID == nil {
		o.ID = &ObjectID{}
	}
	if err := o.ID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Address == nil {
		o.Address = &IPAddress{}
	}
	if err := o.Address.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.PortIndex); err != nil {
		return err
	}
	return nil
}

// ProviderProperty structure represents VDS_PROVIDER_PROP RPC structure.
//
// The VDS_PROVIDER_PROP structure provides information about provider properties.
type ProviderProperty struct {
	// id:  The VDS object ID of the provider object.
	ID *ObjectID `idl:"name:id" json:"id"`
	// pwszName:  The null-terminated Unicode name of the provider.
	Name string `idl:"name:pwszName;string" json:"name"`
	// guidVersionId:  The version GUID of the provider. This GUID MUST be unique to each
	// version of the provider.
	VersionID *dtyp.GUID `idl:"name:guidVersionId" json:"version_id"`
	// pwszVersion:  The null-terminated Unicode version string of the provider. The convention
	// for this string is <major version number>.<minor version number>.
	Version string `idl:"name:pwszVersion;string" json:"version"`
	// type:  A value from the VDS_PROVIDER_TYPE enumeration that indicates the provider
	// type.
	Type ProviderType `idl:"name:type" json:"type"`
	// ulFlags:  A combination of any values, by using a bitwise OR operator, from the VDS_PROVIDER_FLAG
	// enumeration.
	Flags uint32 `idl:"name:ulFlags" json:"flags"`
	// ulStripeSizeFlags:  Stripe size that the provider supports, which MUST be a power
	// of 2. Each bit in the 32-bit integer represents a size that the provider supports.
	// For example, if the nth bit is set, the provider supports a stripe size of 2^n. This
	// parameter is used only for software providers. The basic provider sets this value
	// to zero and the dynamic provider sets this value to 64K.
	StripeSizeFlags uint32 `idl:"name:ulStripeSizeFlags" json:"stripe_size_flags"`
	// sRebuildPriority:  The rebuild priority of all volumes that the provider manages.
	// It specifies the regeneration order when a mirrored or RAID-5 volume requires rebuilding.
	// Priority levels MUST be from 0 through 15. A higher value indicates a higher priority.
	// This parameter is used only for software providers and does not apply to the basic
	// provider.
	RebuildPriority int16 `idl:"name:sRebuildPriority" json:"rebuild_priority"`
}

func (o *ProviderProperty) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ProviderProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.ID != nil {
		if err := o.ID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Name != "" {
		_ptr_pwszName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_pwszName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.VersionID != nil {
		if err := o.VersionID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Version != "" {
		_ptr_pwszVersion := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Version); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Version, _ptr_pwszVersion); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint16(o.Type)); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.StripeSizeFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.RebuildPriority); err != nil {
		return err
	}
	return nil
}
func (o *ProviderProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.ID == nil {
		o.ID = &ObjectID{}
	}
	if err := o.ID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_pwszName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_pwszName := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_pwszName, _ptr_pwszName); err != nil {
		return err
	}
	if o.VersionID == nil {
		o.VersionID = &dtyp.GUID{}
	}
	if err := o.VersionID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_pwszVersion := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Version); err != nil {
			return err
		}
		return nil
	})
	_s_pwszVersion := func(ptr interface{}) { o.Version = *ptr.(*string) }
	if err := w.ReadPointer(&o.Version, _s_pwszVersion, _ptr_pwszVersion); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Type)); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.StripeSizeFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.RebuildPriority); err != nil {
		return err
	}
	return nil
}

// PackProperty structure represents VDS_PACK_PROP RPC structure.
//
// The VDS_PACK_PROP structure provides information about the properties of a disk pack.
type PackProperty struct {
	// id:  The VDS object ID of the disk pack object.
	ID *ObjectID `idl:"name:id" json:"id"`
	// pwszName:  The null-terminated Unicode name of the disk pack. If the pack has no
	// name, this pointer is set to NULL.
	Name string `idl:"name:pwszName;string" json:"name"`
	// status:  The value from the VDS_PACK_STATUS enumeration that indicates the status
	// of the disk pack.
	Status PackStatus `idl:"name:status" json:"status"`
	// ulFlags:  A combination of any values, by using a bitwise OR operator, of the disk
	// pack flags that are defined in the VDS_PACK_FLAG enumeration. ulFlags can be 0 if
	// none of the VDS_PACK_FLAG values apply.
	Flags uint32 `idl:"name:ulFlags" json:"flags"`
}

func (o *PackProperty) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PackProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.ID != nil {
		if err := o.ID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Name != "" {
		_ptr_pwszName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_pwszName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint16(o.Status)); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	return nil
}
func (o *PackProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.ID == nil {
		o.ID = &ObjectID{}
	}
	if err := o.ID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_pwszName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_pwszName := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_pwszName, _ptr_pwszName); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Status)); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	return nil
}

// DiskProperty structure represents VDS_DISK_PROP RPC structure.
//
// The VDS_DISK_PROP structure provides the properties of a disk.
type DiskProperty struct {
	// id:  The VDS object ID of the disk object.
	ID *ObjectID `idl:"name:id" json:"id"`
	// status:  The value from the VDS_DISK_STATUS enumeration that indicates the disk status.
	Status DiskStatus `idl:"name:status" json:"status"`
	// ReserveMode:  The value from the VDS_LUN_RESERVE_MODE enumeration that indicates
	// the sharing mode of the disk.
	ReserveMode LUNReserveMode `idl:"name:ReserveMode" json:"reserve_mode"`
	// health:  The value from the VDS_HEALTH enumeration that indicates the health of the
	// disk.
	Health Health `idl:"name:health" json:"health"`
	// dwDeviceType:  The device type of the disk. Note that this value refers to the disk
	// type and not the drive type. Thus, if there is CD media in a DVD/CD drive, it is
	// identified as FILE_DEVICE_CD_ROM; however, DVD media in the same drive is identified
	// as FILE_DEVICE_DVD. This field can have the following values.
	//
	//	+-------------------------------+------------------------------------------------+
	//	|                               |                                                |
	//	|             VALUE             |                    MEANING                     |
	//	|                               |                                                |
	//	+-------------------------------+------------------------------------------------+
	//	+-------------------------------+------------------------------------------------+
	//	| FILE_DEVICE_CD_ROM 0x00000002 | The device is a CD-ROM.                        |
	//	+-------------------------------+------------------------------------------------+
	//	| FILE_DEVICE_DISK 0x00000007   | The device is a hard disk or removable device. |
	//	+-------------------------------+------------------------------------------------+
	//	| FILE_DEVICE_DVD 0x00000033    | The device is a DVD.                           |
	//	+-------------------------------+------------------------------------------------+
	DeviceType uint32 `idl:"name:dwDeviceType" json:"device_type"`
	// dwMediaType:  The media type of the disk; it can have the following values.
	//
	//	+---------------------------+---------------------------------+
	//	|                           |                                 |
	//	|           VALUE           |             MEANING             |
	//	|                           |                                 |
	//	+---------------------------+---------------------------------+
	//	+---------------------------+---------------------------------+
	//	| Unknown 0x00000000        | The disk media type is unknown. |
	//	+---------------------------+---------------------------------+
	//	| RemovableMedia 0x0000000B | The disk media is removable.    |
	//	+---------------------------+---------------------------------+
	//	| FixedMedia 0x0000000C     | The disk media is fixed.        |
	//	+---------------------------+---------------------------------+
	MediaType uint32 `idl:"name:dwMediaType" json:"media_type"`
	// ullSize:  The size of the disk, in bytes.
	Size uint64 `idl:"name:ullSize" json:"size"`
	// ulBytesPerSector:  The size of the sectors for the disk, in bytes.
	BytesPerSector uint32 `idl:"name:ulBytesPerSector" json:"bytes_per_sector"`
	// ulSectorsPerTrack:  The number of sectors per track on the disk.
	SectorsPerTrack uint32 `idl:"name:ulSectorsPerTrack" json:"sectors_per_track"`
	// ulTracksPerCylinder:  The number of tracks per cylinder on the disk.
	TracksPerCylinder uint32 `idl:"name:ulTracksPerCylinder" json:"tracks_per_cylinder"`
	// ulFlags:  The combination of any values, by using a bitwise OR operator, that are
	// defined in the VDS_DISK_FLAG enumeration. This field can be zero if none of the VDS_DISK_FLAG
	// values apply.
	Flags uint32 `idl:"name:ulFlags" json:"flags"`
	// BusType:  The value from the VDS_STORAGE_BUS_TYPE enumeration that indicates the
	// type of bus where the disk resides.
	BusType StorageBusType `idl:"name:BusType" json:"bus_type"`
	// PartitionStyle:  The value from the VDS_PARTITION_STYLE enumeration that indicates
	// the partitioning format of the disk.
	PartitionStyle PartitionStyle             `idl:"name:PartitionStyle" json:"partition_style"`
	DiskProperty   *DiskProperty_DiskProperty `idl:"name:DiskProperty;switch_is:PartitionStyle" json:"disk_property"`
	// pwszDiskAddress:  The null-terminated Unicode address of the disk, if the disk uses
	// a SCSI-like address; otherwise, NULL. If present, a client can use this property
	// to determine the port number, bus, target number, and LUN of the disk.
	DiskAddress string `idl:"name:pwszDiskAddress;string" json:"disk_address"`
	// pwszName:  The null-terminated Unicode name that the operating system uses to identify
	// the disk. If present, a client can use this property to determine the disk's PNP
	// device number. This is the number obtained from the DeviceNumber member of STORAGE_DEVICE_NUMBER
	// (see [MSDN-STRGEDEVNUM]). For a hard disk, this name has the format \\?\PhysicalDriveN,
	// where N signifies the device number of the disk. For a DVD/CD drive, this name has
	// the format \\?\CdRomN, where N signifies the device number of the DVD/CD drive. A
	// client can use this property to identify the disk.
	Name string `idl:"name:pwszName;string" json:"name"`
	// pwszFriendlyName:  The null-terminated Unicode friendly (human-readable) name of
	// the disk as assigned by the operating system. This property MAY be NULL. If present,
	// a client can use this property to display a human-readable name of the disk.
	FriendlyName string `idl:"name:pwszFriendlyName;string" json:"friendly_name"`
	// pwszAdaptorName:  The null-terminated Unicode name that the operating system assigns
	// to the adapter to which the disk is attached. This property MAY be NULL. If present,
	// a client can use this property to display the adapter name of the disk.
	AdaptorName string `idl:"name:pwszAdaptorName;string" json:"adaptor_name"`
	// pwszDevicePath:  The null-terminated Unicode device path that the operating system
	// uses to identify the device for the disk. This property MAY be NULL. If present,
	// a client can use this property to display the device path of the disk. This string
	// is used to load the property page information for a disk.
	DevicePath string `idl:"name:pwszDevicePath;string" json:"device_path"`
}

func (o *DiskProperty) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DiskProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.ID != nil {
		if err := o.ID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint16(o.Status)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.ReserveMode)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Health)); err != nil {
		return err
	}
	if err := w.WriteData(o.DeviceType); err != nil {
		return err
	}
	if err := w.WriteData(o.MediaType); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if err := w.WriteData(o.BytesPerSector); err != nil {
		return err
	}
	if err := w.WriteData(o.SectorsPerTrack); err != nil {
		return err
	}
	if err := w.WriteData(o.TracksPerCylinder); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.BusType)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.PartitionStyle)); err != nil {
		return err
	}
	_swDiskProperty := uint16(o.PartitionStyle)
	if o.DiskProperty != nil {
		if err := o.DiskProperty.MarshalUnionNDR(ctx, w, _swDiskProperty); err != nil {
			return err
		}
	} else {
		if err := (&DiskProperty_DiskProperty{}).MarshalUnionNDR(ctx, w, _swDiskProperty); err != nil {
			return err
		}
	}
	if o.DiskAddress != "" {
		_ptr_pwszDiskAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DiskAddress); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DiskAddress, _ptr_pwszDiskAddress); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Name != "" {
		_ptr_pwszName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_pwszName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
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
	if o.AdaptorName != "" {
		_ptr_pwszAdaptorName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.AdaptorName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.AdaptorName, _ptr_pwszAdaptorName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DevicePath != "" {
		_ptr_pwszDevicePath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DevicePath); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DevicePath, _ptr_pwszDevicePath); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *DiskProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.ID == nil {
		o.ID = &ObjectID{}
	}
	if err := o.ID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Status)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.ReserveMode)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Health)); err != nil {
		return err
	}
	if err := w.ReadData(&o.DeviceType); err != nil {
		return err
	}
	if err := w.ReadData(&o.MediaType); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if err := w.ReadData(&o.BytesPerSector); err != nil {
		return err
	}
	if err := w.ReadData(&o.SectorsPerTrack); err != nil {
		return err
	}
	if err := w.ReadData(&o.TracksPerCylinder); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.BusType)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.PartitionStyle)); err != nil {
		return err
	}
	if o.DiskProperty == nil {
		o.DiskProperty = &DiskProperty_DiskProperty{}
	}
	_swDiskProperty := uint16(o.PartitionStyle)
	if err := o.DiskProperty.UnmarshalUnionNDR(ctx, w, _swDiskProperty); err != nil {
		return err
	}
	_ptr_pwszDiskAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DiskAddress); err != nil {
			return err
		}
		return nil
	})
	_s_pwszDiskAddress := func(ptr interface{}) { o.DiskAddress = *ptr.(*string) }
	if err := w.ReadPointer(&o.DiskAddress, _s_pwszDiskAddress, _ptr_pwszDiskAddress); err != nil {
		return err
	}
	_ptr_pwszName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_pwszName := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_pwszName, _ptr_pwszName); err != nil {
		return err
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
	_ptr_pwszAdaptorName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.AdaptorName); err != nil {
			return err
		}
		return nil
	})
	_s_pwszAdaptorName := func(ptr interface{}) { o.AdaptorName = *ptr.(*string) }
	if err := w.ReadPointer(&o.AdaptorName, _s_pwszAdaptorName, _ptr_pwszAdaptorName); err != nil {
		return err
	}
	_ptr_pwszDevicePath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DevicePath); err != nil {
			return err
		}
		return nil
	})
	_s_pwszDevicePath := func(ptr interface{}) { o.DevicePath = *ptr.(*string) }
	if err := w.ReadPointer(&o.DevicePath, _s_pwszDevicePath, _ptr_pwszDevicePath); err != nil {
		return err
	}
	return nil
}

// DiskProperty_DiskProperty structure represents VDS_DISK_PROP union anonymous member.
//
// The VDS_DISK_PROP structure provides the properties of a disk.
type DiskProperty_DiskProperty struct {
	// Types that are assignable to Value
	//
	// *DiskProperty_Signature
	// *DiskProperty_DiskGUID
	// *DiskProperty_DefaultDiskProperty
	Value is_DiskProperty_DiskProperty `json:"value"`
}

func (o *DiskProperty_DiskProperty) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *DiskProperty_Signature:
		if value != nil {
			return value.Signature
		}
	case *DiskProperty_DiskGUID:
		if value != nil {
			return value.DiskGUID
		}
	}
	return nil
}

type is_DiskProperty_DiskProperty interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_DiskProperty_DiskProperty()
}

func (o *DiskProperty_DiskProperty) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *DiskProperty_Signature:
		return uint16(1)
	case *DiskProperty_DiskGUID:
		return uint16(2)
	}
	return uint16(0)
}

func (o *DiskProperty_DiskProperty) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		_o, _ := o.Value.(*DiskProperty_Signature)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&DiskProperty_Signature{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*DiskProperty_DiskGUID)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&DiskProperty_DiskGUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		// no-op
	}
	return nil
}

func (o *DiskProperty_DiskProperty) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch((*uint16)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		o.Value = &DiskProperty_Signature{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &DiskProperty_DiskGUID{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		o.Value = &DiskProperty_DefaultDiskProperty{}
	}
	return nil
}

// DiskProperty_Signature structure represents DiskProperty_DiskProperty RPC union arm.
//
// It has following labels: 1
type DiskProperty_Signature struct {
	// dwSignature:  The MBR disk signature of the disk.
	Signature uint32 `idl:"name:dwSignature" json:"signature"`
}

func (*DiskProperty_Signature) is_DiskProperty_DiskProperty() {}

func (o *DiskProperty_Signature) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Signature); err != nil {
		return err
	}
	return nil
}
func (o *DiskProperty_Signature) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Signature); err != nil {
		return err
	}
	return nil
}

// DiskProperty_DiskGUID structure represents DiskProperty_DiskProperty RPC union arm.
//
// It has following labels: 2
type DiskProperty_DiskGUID struct {
	// DiskGuid:  The GUID in the GPT that identifies the disk.
	DiskGUID *dtyp.GUID `idl:"name:DiskGuid" json:"disk_guid"`
}

func (*DiskProperty_DiskGUID) is_DiskProperty_DiskProperty() {}

func (o *DiskProperty_DiskGUID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *DiskProperty_DiskGUID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.DiskGUID == nil {
		o.DiskGUID = &dtyp.GUID{}
	}
	if err := o.DiskGUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DiskProperty_DefaultDiskProperty structure represents DiskProperty_DiskProperty RPC default union arm.
type DiskProperty_DefaultDiskProperty struct {
}

func (*DiskProperty_DefaultDiskProperty) is_DiskProperty_DiskProperty() {}

func (o *DiskProperty_DefaultDiskProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return nil
}
func (o *DiskProperty_DefaultDiskProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	return nil
}

// DiskProperty2 structure represents VDS_DISK_PROP2 RPC structure.
//
// The VDS_DISK_PROP2 structure provides the properties of a disk.
type DiskProperty2 struct {
	// id:  The VDS object ID of the disk object.
	ID *ObjectID `idl:"name:id" json:"id"`
	// status:  The value from the VDS_DISK_STATUS enumeration that indicates the disk status.
	Status DiskStatus `idl:"name:status" json:"status"`
	// ReserveMode:  The value from the VDS_LUN_RESERVE_MODE enumeration that includes the
	// sharing mode of the disk.
	ReserveMode LUNReserveMode `idl:"name:ReserveMode" json:"reserve_mode"`
	// health:  The value from the VDS_HEALTH enumeration that indicates the health of the
	// disk.
	Health Health `idl:"name:health" json:"health"`
	// dwDeviceType:  The device type of the disk. Note that this volume refers to the disk
	// type and not the drive type. Thus, if there is CD media in a DVD/CD drive, it is
	// identified as FILE_DEVICE_CD_ROM; however, DVD media in the same drive is identified
	// as FILE_DEVICE_DVD. This field can have the following values.
	//
	//	+-------------------------------+------------------------------------------------+
	//	|                               |                                                |
	//	|             VALUE             |                    MEANING                     |
	//	|                               |                                                |
	//	+-------------------------------+------------------------------------------------+
	//	+-------------------------------+------------------------------------------------+
	//	| FILE_DEVICE_CD_ROM 0x00000002 | The device is a CD-ROM.                        |
	//	+-------------------------------+------------------------------------------------+
	//	| FILE_DEVICE_DISK 0x00000007   | The device is a hard disk or removable device. |
	//	+-------------------------------+------------------------------------------------+
	//	| FILE_DEVICE_DVD 0x00000033    | The device is a DVD.                           |
	//	+-------------------------------+------------------------------------------------+
	DeviceType uint32 `idl:"name:dwDeviceType" json:"device_type"`
	// dwMediaType:  The media type of the disk. It can have the following values.
	//
	//	+---------------------------+---------------------------------+
	//	|                           |                                 |
	//	|           VALUE           |             MEANING             |
	//	|                           |                                 |
	//	+---------------------------+---------------------------------+
	//	+---------------------------+---------------------------------+
	//	| Unknown 0x00000000        | The disk media type is unknown. |
	//	+---------------------------+---------------------------------+
	//	| RemovableMedia 0x0000000B | The disk is removable media.    |
	//	+---------------------------+---------------------------------+
	//	| FixedMedia 0x0000000C     | The disk media is fixed.        |
	//	+---------------------------+---------------------------------+
	MediaType uint32 `idl:"name:dwMediaType" json:"media_type"`
	// ullSize:  The size of the disk, in bytes.
	Size uint64 `idl:"name:ullSize" json:"size"`
	// ulBytesPerSector:  The size of the sectors for the disk, in bytes.
	BytesPerSector uint32 `idl:"name:ulBytesPerSector" json:"bytes_per_sector"`
	// ulSectorsPerTrack:  The number of sectors per track on the disk.
	SectorsPerTrack uint32 `idl:"name:ulSectorsPerTrack" json:"sectors_per_track"`
	// ulTracksPerCylinder:  The number of tracks per cylinder on the disk.
	TracksPerCylinder uint32 `idl:"name:ulTracksPerCylinder" json:"tracks_per_cylinder"`
	// ulFlags:  The combination of any values, by using a bitwise OR operator, that are
	// defined in the VDS_DISK_FLAG enumeration.
	Flags uint32 `idl:"name:ulFlags" json:"flags"`
	// BusType:  The value from the VDS_STORAGE_BUS_TYPE enumeration that indicates the
	// type of bus where the disk resides.
	BusType StorageBusType `idl:"name:BusType" json:"bus_type"`
	// PartitionStyle:  The value from the VDS_PARTITION_STYLE enumeration that indicates
	// the partitioning format of the disk.
	PartitionStyle PartitionStyle               `idl:"name:PartitionStyle" json:"partition_style"`
	DiskProperty2  *DiskProperty2_DiskProperty2 `idl:"name:DiskProperty2;switch_is:PartitionStyle" json:"disk_property2"`
	// pwszDiskAddress:  The null-terminated Unicode address of the disk, if the disk uses
	// a SCSI-like address. Otherwise, NULL. If present, a client can use this property
	// to determine the port number, bus, target number, and LUN of the disk.
	DiskAddress string `idl:"name:pwszDiskAddress;string" json:"disk_address"`
	// pwszName:  The null-terminated Unicode name that the operating system uses to identify
	// the disk. If present, a client can use this property to determine the disk's PNP
	// device number. For a hard disk, this name has the format \\?\PhysicalDriveN; where
	// N signifies the device number of the disk. For a DVD/CD drive, this name has the
	// format \\?\CdRomN; where N signifies the device number of the DVD/CD drive. A client
	// can use this property to identify the disk.
	Name string `idl:"name:pwszName;string" json:"name"`
	// pwszFriendlyName:  The null-terminated Unicode friendly (human-readable) name of
	// the disk as assigned by the operating system. This property MAY be NULL. If present,
	// a client can use this property to display a human-readable name of the disk.
	FriendlyName string `idl:"name:pwszFriendlyName;string" json:"friendly_name"`
	// pwszAdaptorName:  The null-terminated Unicode name that the operating system assigns
	// to the adapter to which the disk is attached. This property MAY be NULL. If present,
	// a client can use this property to display the adapter name of the disk.
	AdaptorName string `idl:"name:pwszAdaptorName;string" json:"adaptor_name"`
	// pwszDevicePath:  The null-terminated Unicode device path that the operating system
	// uses to identify the device for the disk. This property MAY be NULL. If present,
	// a client can use this property to display the device path of the disk. This string
	// is used to load the property page information for a disk.
	DevicePath string `idl:"name:pwszDevicePath;string" json:"device_path"`
	// pwszLocationPath:  This string is built from a combination of the DEVPKEY_Device_LocationPaths
	// value for the disk's adapter, the bus type, and the SCSI address. The DEVPKEY_Device_LocationPaths
	// property represents the location of a device instance in the device tree.
	//
	// The following table shows examples of location paths built for various bus/disk types.
	//
	//	+-------------------+----------------------------------------------------------+
	//	|     BUS/DISK      |                         LOCATION                         |
	//	|       TYPE        |                           PATH                           |
	//	+-------------------+----------------------------------------------------------+
	//	+-------------------+----------------------------------------------------------+
	//	| IDE\ATA\PATA\SATA | PCIROOT(0)#PCI(0100)#ATA(C01T03L00)                      |
	//	+-------------------+----------------------------------------------------------+
	//	| SCSI              | PCIROOT(0)#PCI(1C00)#PCI(0000)#SCSI(P00T01L01)           |
	//	+-------------------+----------------------------------------------------------+
	//	| SAS               | PCIROOT(1)#PCI(0300)#SAS(P00T03L00)                      |
	//	+-------------------+----------------------------------------------------------+
	//	| PCI RAID          | PCIROOT(0)#PCI(0200)#PCI(0003)#PCI(0100)#RAID(P02T00L00) |
	//	+-------------------+----------------------------------------------------------+
	LocationPath string `idl:"name:pwszLocationPath;string" json:"location_path"`
}

func (o *DiskProperty2) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DiskProperty2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.ID != nil {
		if err := o.ID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint16(o.Status)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.ReserveMode)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Health)); err != nil {
		return err
	}
	if err := w.WriteData(o.DeviceType); err != nil {
		return err
	}
	if err := w.WriteData(o.MediaType); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if err := w.WriteData(o.BytesPerSector); err != nil {
		return err
	}
	if err := w.WriteData(o.SectorsPerTrack); err != nil {
		return err
	}
	if err := w.WriteData(o.TracksPerCylinder); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.BusType)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.PartitionStyle)); err != nil {
		return err
	}
	_swDiskProperty2 := uint16(o.PartitionStyle)
	if o.DiskProperty2 != nil {
		if err := o.DiskProperty2.MarshalUnionNDR(ctx, w, _swDiskProperty2); err != nil {
			return err
		}
	} else {
		if err := (&DiskProperty2_DiskProperty2{}).MarshalUnionNDR(ctx, w, _swDiskProperty2); err != nil {
			return err
		}
	}
	if o.DiskAddress != "" {
		_ptr_pwszDiskAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DiskAddress); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DiskAddress, _ptr_pwszDiskAddress); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Name != "" {
		_ptr_pwszName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_pwszName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
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
	if o.AdaptorName != "" {
		_ptr_pwszAdaptorName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.AdaptorName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.AdaptorName, _ptr_pwszAdaptorName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DevicePath != "" {
		_ptr_pwszDevicePath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DevicePath); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DevicePath, _ptr_pwszDevicePath); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LocationPath != "" {
		_ptr_pwszLocationPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LocationPath); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LocationPath, _ptr_pwszLocationPath); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *DiskProperty2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.ID == nil {
		o.ID = &ObjectID{}
	}
	if err := o.ID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Status)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.ReserveMode)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Health)); err != nil {
		return err
	}
	if err := w.ReadData(&o.DeviceType); err != nil {
		return err
	}
	if err := w.ReadData(&o.MediaType); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if err := w.ReadData(&o.BytesPerSector); err != nil {
		return err
	}
	if err := w.ReadData(&o.SectorsPerTrack); err != nil {
		return err
	}
	if err := w.ReadData(&o.TracksPerCylinder); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.BusType)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.PartitionStyle)); err != nil {
		return err
	}
	if o.DiskProperty2 == nil {
		o.DiskProperty2 = &DiskProperty2_DiskProperty2{}
	}
	_swDiskProperty2 := uint16(o.PartitionStyle)
	if err := o.DiskProperty2.UnmarshalUnionNDR(ctx, w, _swDiskProperty2); err != nil {
		return err
	}
	_ptr_pwszDiskAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DiskAddress); err != nil {
			return err
		}
		return nil
	})
	_s_pwszDiskAddress := func(ptr interface{}) { o.DiskAddress = *ptr.(*string) }
	if err := w.ReadPointer(&o.DiskAddress, _s_pwszDiskAddress, _ptr_pwszDiskAddress); err != nil {
		return err
	}
	_ptr_pwszName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_pwszName := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_pwszName, _ptr_pwszName); err != nil {
		return err
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
	_ptr_pwszAdaptorName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.AdaptorName); err != nil {
			return err
		}
		return nil
	})
	_s_pwszAdaptorName := func(ptr interface{}) { o.AdaptorName = *ptr.(*string) }
	if err := w.ReadPointer(&o.AdaptorName, _s_pwszAdaptorName, _ptr_pwszAdaptorName); err != nil {
		return err
	}
	_ptr_pwszDevicePath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DevicePath); err != nil {
			return err
		}
		return nil
	})
	_s_pwszDevicePath := func(ptr interface{}) { o.DevicePath = *ptr.(*string) }
	if err := w.ReadPointer(&o.DevicePath, _s_pwszDevicePath, _ptr_pwszDevicePath); err != nil {
		return err
	}
	_ptr_pwszLocationPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LocationPath); err != nil {
			return err
		}
		return nil
	})
	_s_pwszLocationPath := func(ptr interface{}) { o.LocationPath = *ptr.(*string) }
	if err := w.ReadPointer(&o.LocationPath, _s_pwszLocationPath, _ptr_pwszLocationPath); err != nil {
		return err
	}
	return nil
}

// DiskProperty2_DiskProperty2 structure represents VDS_DISK_PROP2 union anonymous member.
//
// The VDS_DISK_PROP2 structure provides the properties of a disk.
type DiskProperty2_DiskProperty2 struct {
	// Types that are assignable to Value
	//
	// *DiskProperty2_Signature
	// *DiskProperty2_DiskGUID
	// *DiskProperty2_DefaultDiskProperty2
	Value is_DiskProperty2_DiskProperty2 `json:"value"`
}

func (o *DiskProperty2_DiskProperty2) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *DiskProperty2_Signature:
		if value != nil {
			return value.Signature
		}
	case *DiskProperty2_DiskGUID:
		if value != nil {
			return value.DiskGUID
		}
	}
	return nil
}

type is_DiskProperty2_DiskProperty2 interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_DiskProperty2_DiskProperty2()
}

func (o *DiskProperty2_DiskProperty2) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *DiskProperty2_Signature:
		return uint16(1)
	case *DiskProperty2_DiskGUID:
		return uint16(2)
	}
	return uint16(0)
}

func (o *DiskProperty2_DiskProperty2) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		_o, _ := o.Value.(*DiskProperty2_Signature)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&DiskProperty2_Signature{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*DiskProperty2_DiskGUID)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&DiskProperty2_DiskGUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		// no-op
	}
	return nil
}

func (o *DiskProperty2_DiskProperty2) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch((*uint16)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		o.Value = &DiskProperty2_Signature{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &DiskProperty2_DiskGUID{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		o.Value = &DiskProperty2_DefaultDiskProperty2{}
	}
	return nil
}

// DiskProperty2_Signature structure represents DiskProperty2_DiskProperty2 RPC union arm.
//
// It has following labels: 1
type DiskProperty2_Signature struct {
	// dwSignature:  The MBR disk signature of the disk.
	Signature uint32 `idl:"name:dwSignature" json:"signature"`
}

func (*DiskProperty2_Signature) is_DiskProperty2_DiskProperty2() {}

func (o *DiskProperty2_Signature) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Signature); err != nil {
		return err
	}
	return nil
}
func (o *DiskProperty2_Signature) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Signature); err != nil {
		return err
	}
	return nil
}

// DiskProperty2_DiskGUID structure represents DiskProperty2_DiskProperty2 RPC union arm.
//
// It has following labels: 2
type DiskProperty2_DiskGUID struct {
	// DiskGuid:  The GUID in the GPT that identifies the disk.
	DiskGUID *dtyp.GUID `idl:"name:DiskGuid" json:"disk_guid"`
}

func (*DiskProperty2_DiskGUID) is_DiskProperty2_DiskProperty2() {}

func (o *DiskProperty2_DiskGUID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *DiskProperty2_DiskGUID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.DiskGUID == nil {
		o.DiskGUID = &dtyp.GUID{}
	}
	if err := o.DiskGUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DiskProperty2_DefaultDiskProperty2 structure represents DiskProperty2_DiskProperty2 RPC default union arm.
type DiskProperty2_DefaultDiskProperty2 struct {
}

func (*DiskProperty2_DefaultDiskProperty2) is_DiskProperty2_DiskProperty2() {}

func (o *DiskProperty2_DefaultDiskProperty2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return nil
}
func (o *DiskProperty2_DefaultDiskProperty2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	return nil
}

// AdvancedDiskProperty structure represents VDS_ADVANCEDDISK_PROP RPC structure.
type AdvancedDiskProperty struct {
	ID                   string                                     `idl:"name:pwszId;string" json:"id"`
	PathName             string                                     `idl:"name:pwszPathname;string" json:"path_name"`
	Location             string                                     `idl:"name:pwszLocation;string" json:"location"`
	FriendlyName         string                                     `idl:"name:pwszFriendlyName;string" json:"friendly_name"`
	VPDID                string                                     `idl:"name:pswzIdentifier;string" json:"vpd_id"`
	IDFormat             uint16                                     `idl:"name:usIdentifierFormat" json:"id_format"`
	Number               uint32                                     `idl:"name:ulNumber" json:"number"`
	SerialNumber         string                                     `idl:"name:pwszSerialNumber;string" json:"serial_number"`
	FirmwareVersion      string                                     `idl:"name:pwszFirmwareVersion;string" json:"firmware_version"`
	Manufacturer         string                                     `idl:"name:pwszManufacturer;string" json:"manufacturer"`
	Model                string                                     `idl:"name:pwszModel;string" json:"model"`
	TotalSize            uint64                                     `idl:"name:ullTotalSize" json:"total_size"`
	AllocatedSize        uint64                                     `idl:"name:ullAllocatedSize" json:"allocated_size"`
	LogicalSectorSize    uint32                                     `idl:"name:ulLogicalSectorSize" json:"logical_sector_size"`
	PhysicalSectorSize   uint32                                     `idl:"name:ulPhysicalSectorSize" json:"physical_sector_size"`
	PartitionCount       uint32                                     `idl:"name:ulPartitionCount" json:"partition_count"`
	Status               DiskStatus                                 `idl:"name:status" json:"status"`
	Health               Health                                     `idl:"name:health" json:"health"`
	BusType              StorageBusType                             `idl:"name:BusType" json:"bus_type"`
	PartitionStyle       PartitionStyle                             `idl:"name:PartitionStyle" json:"partition_style"`
	AdvancedDiskProperty *AdvancedDiskProperty_AdvancedDiskProperty `idl:"name:AdvancedDiskProperty;switch_is:PartitionStyle" json:"advanced_disk_property"`
	Flags                uint32                                     `idl:"name:ulFlags" json:"flags"`
	DeviceType           uint32                                     `idl:"name:dwDeviceType" json:"device_type"`
}

func (o *AdvancedDiskProperty) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AdvancedDiskProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.ID != "" {
		_ptr_pwszId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ID, _ptr_pwszId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.PathName != "" {
		_ptr_pwszPathname := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.PathName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.PathName, _ptr_pwszPathname); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Location != "" {
		_ptr_pwszLocation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Location); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Location, _ptr_pwszLocation); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
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
	if o.VPDID != "" {
		_ptr_pswzIdentifier := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.VPDID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.VPDID, _ptr_pswzIdentifier); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.IDFormat); err != nil {
		return err
	}
	if err := w.WriteData(o.Number); err != nil {
		return err
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
	if o.FirmwareVersion != "" {
		_ptr_pwszFirmwareVersion := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.FirmwareVersion); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.FirmwareVersion, _ptr_pwszFirmwareVersion); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Manufacturer != "" {
		_ptr_pwszManufacturer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Manufacturer); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Manufacturer, _ptr_pwszManufacturer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Model != "" {
		_ptr_pwszModel := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Model); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Model, _ptr_pwszModel); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.TotalSize); err != nil {
		return err
	}
	if err := w.WriteData(o.AllocatedSize); err != nil {
		return err
	}
	if err := w.WriteData(o.LogicalSectorSize); err != nil {
		return err
	}
	if err := w.WriteData(o.PhysicalSectorSize); err != nil {
		return err
	}
	if err := w.WriteData(o.PartitionCount); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Status)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Health)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.BusType)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.PartitionStyle)); err != nil {
		return err
	}
	_swAdvancedDiskProperty := uint16(o.PartitionStyle)
	if o.AdvancedDiskProperty != nil {
		if err := o.AdvancedDiskProperty.MarshalUnionNDR(ctx, w, _swAdvancedDiskProperty); err != nil {
			return err
		}
	} else {
		if err := (&AdvancedDiskProperty_AdvancedDiskProperty{}).MarshalUnionNDR(ctx, w, _swAdvancedDiskProperty); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.DeviceType); err != nil {
		return err
	}
	return nil
}
func (o *AdvancedDiskProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	_ptr_pwszId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ID); err != nil {
			return err
		}
		return nil
	})
	_s_pwszId := func(ptr interface{}) { o.ID = *ptr.(*string) }
	if err := w.ReadPointer(&o.ID, _s_pwszId, _ptr_pwszId); err != nil {
		return err
	}
	_ptr_pwszPathname := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.PathName); err != nil {
			return err
		}
		return nil
	})
	_s_pwszPathname := func(ptr interface{}) { o.PathName = *ptr.(*string) }
	if err := w.ReadPointer(&o.PathName, _s_pwszPathname, _ptr_pwszPathname); err != nil {
		return err
	}
	_ptr_pwszLocation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Location); err != nil {
			return err
		}
		return nil
	})
	_s_pwszLocation := func(ptr interface{}) { o.Location = *ptr.(*string) }
	if err := w.ReadPointer(&o.Location, _s_pwszLocation, _ptr_pwszLocation); err != nil {
		return err
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
	_ptr_pswzIdentifier := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.VPDID); err != nil {
			return err
		}
		return nil
	})
	_s_pswzIdentifier := func(ptr interface{}) { o.VPDID = *ptr.(*string) }
	if err := w.ReadPointer(&o.VPDID, _s_pswzIdentifier, _ptr_pswzIdentifier); err != nil {
		return err
	}
	if err := w.ReadData(&o.IDFormat); err != nil {
		return err
	}
	if err := w.ReadData(&o.Number); err != nil {
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
	_ptr_pwszFirmwareVersion := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.FirmwareVersion); err != nil {
			return err
		}
		return nil
	})
	_s_pwszFirmwareVersion := func(ptr interface{}) { o.FirmwareVersion = *ptr.(*string) }
	if err := w.ReadPointer(&o.FirmwareVersion, _s_pwszFirmwareVersion, _ptr_pwszFirmwareVersion); err != nil {
		return err
	}
	_ptr_pwszManufacturer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Manufacturer); err != nil {
			return err
		}
		return nil
	})
	_s_pwszManufacturer := func(ptr interface{}) { o.Manufacturer = *ptr.(*string) }
	if err := w.ReadPointer(&o.Manufacturer, _s_pwszManufacturer, _ptr_pwszManufacturer); err != nil {
		return err
	}
	_ptr_pwszModel := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Model); err != nil {
			return err
		}
		return nil
	})
	_s_pwszModel := func(ptr interface{}) { o.Model = *ptr.(*string) }
	if err := w.ReadPointer(&o.Model, _s_pwszModel, _ptr_pwszModel); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.AllocatedSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.LogicalSectorSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.PhysicalSectorSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.PartitionCount); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Status)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Health)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.BusType)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.PartitionStyle)); err != nil {
		return err
	}
	if o.AdvancedDiskProperty == nil {
		o.AdvancedDiskProperty = &AdvancedDiskProperty_AdvancedDiskProperty{}
	}
	_swAdvancedDiskProperty := uint16(o.PartitionStyle)
	if err := o.AdvancedDiskProperty.UnmarshalUnionNDR(ctx, w, _swAdvancedDiskProperty); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.DeviceType); err != nil {
		return err
	}
	return nil
}

// AdvancedDiskProperty_AdvancedDiskProperty structure represents VDS_ADVANCEDDISK_PROP union anonymous member.
type AdvancedDiskProperty_AdvancedDiskProperty struct {
	// Types that are assignable to Value
	//
	// *AdvancedDiskProperty_Signature
	// *AdvancedDiskProperty_DiskGUID
	// *AdvancedDiskProperty_DefaultAdvancedDiskProperty
	Value is_AdvancedDiskProperty_AdvancedDiskProperty `json:"value"`
}

func (o *AdvancedDiskProperty_AdvancedDiskProperty) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *AdvancedDiskProperty_Signature:
		if value != nil {
			return value.Signature
		}
	case *AdvancedDiskProperty_DiskGUID:
		if value != nil {
			return value.DiskGUID
		}
	}
	return nil
}

type is_AdvancedDiskProperty_AdvancedDiskProperty interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_AdvancedDiskProperty_AdvancedDiskProperty()
}

func (o *AdvancedDiskProperty_AdvancedDiskProperty) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *AdvancedDiskProperty_Signature:
		return uint16(1)
	case *AdvancedDiskProperty_DiskGUID:
		return uint16(2)
	}
	return uint16(0)
}

func (o *AdvancedDiskProperty_AdvancedDiskProperty) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		_o, _ := o.Value.(*AdvancedDiskProperty_Signature)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&AdvancedDiskProperty_Signature{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*AdvancedDiskProperty_DiskGUID)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&AdvancedDiskProperty_DiskGUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		// no-op
	}
	return nil
}

func (o *AdvancedDiskProperty_AdvancedDiskProperty) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch((*uint16)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		o.Value = &AdvancedDiskProperty_Signature{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &AdvancedDiskProperty_DiskGUID{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		o.Value = &AdvancedDiskProperty_DefaultAdvancedDiskProperty{}
	}
	return nil
}

// AdvancedDiskProperty_Signature structure represents AdvancedDiskProperty_AdvancedDiskProperty RPC union arm.
//
// It has following labels: 1
type AdvancedDiskProperty_Signature struct {
	Signature uint32 `idl:"name:dwSignature" json:"signature"`
}

func (*AdvancedDiskProperty_Signature) is_AdvancedDiskProperty_AdvancedDiskProperty() {}

func (o *AdvancedDiskProperty_Signature) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Signature); err != nil {
		return err
	}
	return nil
}
func (o *AdvancedDiskProperty_Signature) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Signature); err != nil {
		return err
	}
	return nil
}

// AdvancedDiskProperty_DiskGUID structure represents AdvancedDiskProperty_AdvancedDiskProperty RPC union arm.
//
// It has following labels: 2
type AdvancedDiskProperty_DiskGUID struct {
	DiskGUID *dtyp.GUID `idl:"name:DiskGuid" json:"disk_guid"`
}

func (*AdvancedDiskProperty_DiskGUID) is_AdvancedDiskProperty_AdvancedDiskProperty() {}

func (o *AdvancedDiskProperty_DiskGUID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *AdvancedDiskProperty_DiskGUID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.DiskGUID == nil {
		o.DiskGUID = &dtyp.GUID{}
	}
	if err := o.DiskGUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// AdvancedDiskProperty_DefaultAdvancedDiskProperty structure represents AdvancedDiskProperty_AdvancedDiskProperty RPC default union arm.
type AdvancedDiskProperty_DefaultAdvancedDiskProperty struct {
}

func (*AdvancedDiskProperty_DefaultAdvancedDiskProperty) is_AdvancedDiskProperty_AdvancedDiskProperty() {
}

func (o *AdvancedDiskProperty_DefaultAdvancedDiskProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return nil
}
func (o *AdvancedDiskProperty_DefaultAdvancedDiskProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	return nil
}

// VolumeProperty structure represents VDS_VOLUME_PROP RPC structure.
//
// The VDS_VOLUME_PROP structure that provides the properties of a volume.
type VolumeProperty struct {
	// id:  The VDS object ID of the volume object.
	ID *ObjectID `idl:"name:id" json:"id"`
	// type:  The value from the VDS_VOLUME_TYPE enumeration that defines the type of the
	// volume.
	Type VolumeType `idl:"name:type" json:"type"`
	// status:  The value from the VDS_VOLUME_STATUS enumeration that defines the status
	// of the volume.
	Status VolumeStatus `idl:"name:status" json:"status"`
	// health:  The value from the VDS_HEALTH enumeration that defines the health of the
	// volume.
	Health Health `idl:"name:health" json:"health"`
	// TransitionState:  The value from the VDS_TRANSITION_STATE enumeration that defines
	// the configuration stability of the volume.
	TransitionState TransitionState `idl:"name:TransitionState" json:"transition_state"`
	// ullSize:  The size of the volume, in bytes.
	Size uint64 `idl:"name:ullSize" json:"size"`
	// ulFlags:  The combination of any values by using the bitwise OR operator of volume
	// flags from the VDS_VOLUME_FLAG enumeration.
	Flags uint32 `idl:"name:ulFlags" json:"flags"`
	// RecommendedFileSystemType:  The value from the VDS_FILE_SYSTEM_TYPE enumeration that
	// defines the recommended file system type for the volume.
	RecommendedFileSystemType FileSystemType `idl:"name:RecommendedFileSystemType" json:"recommended_file_system_type"`
	// pwszName:  The null-terminated Unicode name that the operating system uses to identify
	// the volume.
	Name string `idl:"name:pwszName;string" json:"name"`
}

func (o *VolumeProperty) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *VolumeProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.ID != nil {
		if err := o.ID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint16(o.Type)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Status)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Health)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.TransitionState)); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.RecommendedFileSystemType)); err != nil {
		return err
	}
	if o.Name != "" {
		_ptr_pwszName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_pwszName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *VolumeProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.ID == nil {
		o.ID = &ObjectID{}
	}
	if err := o.ID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Type)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Status)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Health)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.TransitionState)); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.RecommendedFileSystemType)); err != nil {
		return err
	}
	_ptr_pwszName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_pwszName := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_pwszName, _ptr_pwszName); err != nil {
		return err
	}
	return nil
}

// VolumeProperty2 structure represents VDS_VOLUME_PROP2 RPC structure.
//
// The VDS_VOLUME_PROP2 structure provides the properties of a volume.
type VolumeProperty2 struct {
	// id:  The VDS object ID of the volume object.
	ID *ObjectID `idl:"name:id" json:"id"`
	// type:  The value from the VDS_VOLUME_TYPE enumeration that defines the type of the
	// volume.
	Type VolumeType `idl:"name:type" json:"type"`
	// status:  The value from the VDS_VOLUME_STATUS enumeration that defines the status
	// of the volume.
	Status VolumeStatus `idl:"name:status" json:"status"`
	// health:  The value from the VDS_HEALTH enumeration that defines the health of the
	// volume.
	Health Health `idl:"name:health" json:"health"`
	// TransitionState:  The value from the VDS_TRANSITION_STATE enumeration that defines
	// the configuration stability of the volume.
	TransitionState TransitionState `idl:"name:TransitionState" json:"transition_state"`
	// ullSize:  The size of the volume, in bytes.
	Size uint64 `idl:"name:ullSize" json:"size"`
	// ulFlags:  The combination of any values, by using the bitwise OR operator, of volume
	// flags from the VDS_VOLUME_FLAG enumeration.
	Flags uint32 `idl:"name:ulFlags" json:"flags"`
	// RecommendedFileSystemType:  The value from the VDS_FILE_SYSTEM_TYPE enumeration that
	// defines the recommended file system type for the volume.
	RecommendedFileSystemType FileSystemType `idl:"name:RecommendedFileSystemType" json:"recommended_file_system_type"`
	// cbUniqueId:  Count of bytes for pUniqueId.
	UniqueIDLength uint32 `idl:"name:cbUniqueId" json:"unique_id_length"`
	// pwszName:  The null-terminated Unicode name that the operating system uses to identify
	// the volume.
	Name string `idl:"name:pwszName;string" json:"name"`
	// pUniqueId:  A byte array containing the volume's unique id.
	UniqueID []byte `idl:"name:pUniqueId;size_is:(cbUniqueId)" json:"unique_id"`
}

func (o *VolumeProperty2) xxx_PreparePayload(ctx context.Context) error {
	if o.UniqueID != nil && o.UniqueIDLength == 0 {
		o.UniqueIDLength = uint32(len(o.UniqueID))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *VolumeProperty2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.ID != nil {
		if err := o.ID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint16(o.Type)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Status)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Health)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.TransitionState)); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.RecommendedFileSystemType)); err != nil {
		return err
	}
	if err := w.WriteData(o.UniqueIDLength); err != nil {
		return err
	}
	if o.Name != "" {
		_ptr_pwszName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_pwszName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.UniqueID != nil || o.UniqueIDLength > 0 {
		_ptr_pUniqueId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.UniqueIDLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.UniqueID {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.UniqueID[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.UniqueID); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.UniqueID, _ptr_pUniqueId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *VolumeProperty2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.ID == nil {
		o.ID = &ObjectID{}
	}
	if err := o.ID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Type)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Status)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Health)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.TransitionState)); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.RecommendedFileSystemType)); err != nil {
		return err
	}
	if err := w.ReadData(&o.UniqueIDLength); err != nil {
		return err
	}
	_ptr_pwszName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_pwszName := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_pwszName, _ptr_pwszName); err != nil {
		return err
	}
	_ptr_pUniqueId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.UniqueIDLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.UniqueIDLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.UniqueID", sizeInfo[0])
		}
		o.UniqueID = make([]byte, sizeInfo[0])
		for i1 := range o.UniqueID {
			i1 := i1
			if err := w.ReadData(&o.UniqueID[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pUniqueId := func(ptr interface{}) { o.UniqueID = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.UniqueID, _s_pUniqueId, _ptr_pUniqueId); err != nil {
		return err
	}
	return nil
}

// VolumePlexProperty structure represents VDS_VOLUME_PLEX_PROP RPC structure.
//
// The VDS_VOLUME_PLEX_PROP structure provides information about the properties of a
// volume plex.
type VolumePlexProperty struct {
	// id:  The GUID of the plex object.
	ID *ObjectID `idl:"name:id" json:"id"`
	// type:  The plex type that is enumerated by VDS_VOLUME_PLEX_TYPE. The type of the
	// plex need not match that of the volume to which it belongs. For example, a mirrored
	// RAID-1 volume can be composed of plexes that are simple (composed of extents from
	// exactly one disk).
	Type VolumePlexType `idl:"name:type" json:"type"`
	// status:  The status of the plex object that is enumerated by VDS_VOLUME_PLEX_STATUS.
	// The status of the plex need not match that of the volume to which it belongs. For
	// example, a volume plex can have a failed status (VDS_VPS_FAILED), but if the volume
	// is fault-tolerant and its other plexes are online (VDS_VPS_ONLINE), the volume will
	// still be online (VDS_VS_ONLINE).
	Status VolumePlexStatus `idl:"name:status" json:"status"`
	// health:  Value from the VDS_HEALTH enumeration that defines the health of the volume.
	// The health of the plex need not match that of the volume to which it belongs. For
	// instance, a volume's plex can have failed health (VDS_H_FAILED), but if the volume
	// is a mirror volume (RAID-1) and its other plexes are healthy (VDS_H_HEALTHY), the
	// volume will have failed redundancy health (VDS_H_FAILED_REDUNDANCY).
	Health Health `idl:"name:health" json:"health"`
	// TransitionState:  Value from the VDS_TRANSITION_STATE enumeration that defines the
	// configuration stability of the plex. The TransitionState of the plex matches the
	// TransitionState of the volume to which it belongs.
	TransitionState TransitionState `idl:"name:TransitionState" json:"transition_state"`
	// ullSize:  The size of the plex, in bytes. The size can be equal to, or greater than,
	// that of the volume to which it belongs. The plex cannot be smaller than the volume.
	Size uint64 `idl:"name:ullSize" json:"size"`
	// ulStripeSize:  The stripe interleave size, in bytes. This member applies only for
	// plexes of type VDS_VPT_STRIPE (striped) and VDS_VPT_PARITY (striped with parity).
	StripeSize uint32 `idl:"name:ulStripeSize" json:"stripe_size"`
	// ulNumberOfMembers:  The number of members (RAID columns) in the volume plex.
	NumberOfMembers uint32 `idl:"name:ulNumberOfMembers" json:"number_of_members"`
}

func (o *VolumePlexProperty) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *VolumePlexProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.ID != nil {
		if err := o.ID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint16(o.Type)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Status)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Health)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.TransitionState)); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if err := w.WriteData(o.StripeSize); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfMembers); err != nil {
		return err
	}
	return nil
}
func (o *VolumePlexProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.ID == nil {
		o.ID = &ObjectID{}
	}
	if err := o.ID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Type)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Status)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Health)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.TransitionState)); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if err := w.ReadData(&o.StripeSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfMembers); err != nil {
		return err
	}
	return nil
}

// CreateVirtualDiskFlag type represents CREATE_VIRTUAL_DISK_FLAG RPC enumeration.
//
// The CREATE_VIRTUAL_DISK_FLAG enumeration defines the properties of a virtual disk
// that is being created.
type CreateVirtualDiskFlag uint16

var (
	// CREATE_VIRTUAL_DISK_FLAG_NONE:  Indicates to the server that no flags have been
	// specified. CREATE_VIRTUAL_DISK_FLAG_NONE specifies that a virtual disk file will
	// be allocated as the backing store for the virtual disk and that this file will automatically
	// increase in size to accommodate the allocated data.
	//
	// A virtual disk created using the CREATE_ VIRTUAL_DISK_FLAG_NONE flag has a virtual
	// disk file backing store that at any given time is as large as the actual data written
	// to it, plus the size of the header and footer. As more data is written, the virtual
	// disk file automatically increases in size.
	CreateVirtualDiskFlagNone CreateVirtualDiskFlag = 0
	// CREATE_VIRTUAL_DISK_FLAG_FULL_PHYSICAL_ALLOCATION: Specifies that the server preallocates
	// all physical space necessary for the virtual size of the virtual disk. A fixed size
	// virtual disk file will be allocated as the backing store for the virtual disk. For
	// example, creating a fixed size virtual disk that is 2 gigabytes in size using this
	// flag will result in a virtual disk file that is approximately 2 gigabytes in size.
	CreateVirtualDiskFlagFullPhysicalAllocation CreateVirtualDiskFlag = 1
)

func (o CreateVirtualDiskFlag) String() string {
	switch o {
	case CreateVirtualDiskFlagNone:
		return "CreateVirtualDiskFlagNone"
	case CreateVirtualDiskFlagFullPhysicalAllocation:
		return "CreateVirtualDiskFlagFullPhysicalAllocation"
	}
	return "Invalid"
}

// OpenVirtualDiskFlag type represents OPEN_VIRTUAL_DISK_FLAG RPC enumeration.
//
// The OPEN_VIRTUAL_DISK_FLAG enumeration defines flags that are used to open a virtual
// disk object.
type OpenVirtualDiskFlag uint16

var (
	// OPEN_VIRTUAL_DISK_FLAG_NONE:  Indicates that no flag has been specified.
	OpenVirtualDiskFlagNone OpenVirtualDiskFlag = 0
	// OPEN_VIRTUAL_DISK_FLAG_NO_PARENTS:  Applicable only to differencing type virtual
	// disks. Opens the backing store without opening the backing store for any differencing
	// chain parents.
	OpenVirtualDiskFlagNoParents OpenVirtualDiskFlag = 1
	// OPEN_VIRTUAL_DISK_FLAG_BLANK_FILE:  Opens the backing store as an empty file without
	// performing virtual disk verification.
	OpenVirtualDiskFlagBlankFile OpenVirtualDiskFlag = 2
	// OPEN_VIRTUAL_DISK_FLAG_BOOT_DRIVE:  This flag MUST not be used by VDS virtual disk
	// providers or their clients.<51>
	OpenVirtualDiskFlagBootDrive OpenVirtualDiskFlag = 4
)

func (o OpenVirtualDiskFlag) String() string {
	switch o {
	case OpenVirtualDiskFlagNone:
		return "OpenVirtualDiskFlagNone"
	case OpenVirtualDiskFlagNoParents:
		return "OpenVirtualDiskFlagNoParents"
	case OpenVirtualDiskFlagBlankFile:
		return "OpenVirtualDiskFlagBlankFile"
	case OpenVirtualDiskFlagBootDrive:
		return "OpenVirtualDiskFlagBootDrive"
	}
	return "Invalid"
}

// CreateVDiskParameters structure represents VDS_CREATE_VDISK_PARAMETERS RPC structure.
//
// The VDS_CREATE_VDISK_PARAMETERS structure contains the parameters to be used when
// a virtual disk is created.
type CreateVDiskParameters struct {
	// UniqueId:  A unique and non-zero GUID value to be assigned to the virtual disk.
	UniqueID *dtyp.GUID `idl:"name:UniqueId" json:"unique_id"`
	// MaximumSize:  The maximum virtual size, in bytes, of the virtual disk object.
	MaximumSize uint64 `idl:"name:MaximumSize" json:"maximum_size"`
	// BlockSizeInBytes:  The internal block size, in bytes, of the virtual disk object.
	// If the virtual disk object being created is a differencing disk, this value MUST
	// be 0. If the virtual disk object being created is not a differencing disk, setting
	// this value to 0 causes the virtual disk object being created to use the default block
	// size.<50>
	BlockSizeInBytes uint32 `idl:"name:BlockSizeInBytes" json:"block_size_in_bytes"`
	// SectorSizeInBytes:  Internal sector size, in bytes, of the virtual disk object. This
	// value MUST be set to 512 (CREATE_VIRTUAL_DISK_PARAMETERS_DEFAULT_SECTOR_SIZE).
	SectorSizeInBytes uint32 `idl:"name:SectorSizeInBytes" json:"sector_size_in_bytes"`
	// pParentPath:  A null-terminated wide-character string containing an optional path
	// to a parent virtual disk object. This member associates the new virtual hard disk
	// with an existing virtual hard disk. Used when creating a differencing disk. The differencing
	// disk gets its size from its parent.
	ParentPath string `idl:"name:pParentPath;string" json:"parent_path"`
	// pSourcePath:  A null-terminated wide-character string containing an optional path
	// to a source of data to be copied to the new virtual hard disk. When pSourcePath is
	// specified, data from the input virtual disk file is copied block for block from the
	// input virtual disk file to the created virtual disk file. There is no parent-child
	// relationship established.
	SourcePath string `idl:"name:pSourcePath;string" json:"source_path"`
}

func (o *CreateVDiskParameters) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CreateVDiskParameters) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.UniqueID != nil {
		if err := o.UniqueID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.MaximumSize); err != nil {
		return err
	}
	if err := w.WriteData(o.BlockSizeInBytes); err != nil {
		return err
	}
	if err := w.WriteData(o.SectorSizeInBytes); err != nil {
		return err
	}
	if o.ParentPath != "" {
		_ptr_pParentPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ParentPath); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ParentPath, _ptr_pParentPath); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.SourcePath != "" {
		_ptr_pSourcePath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.SourcePath); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SourcePath, _ptr_pSourcePath); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *CreateVDiskParameters) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.UniqueID == nil {
		o.UniqueID = &dtyp.GUID{}
	}
	if err := o.UniqueID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaximumSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.BlockSizeInBytes); err != nil {
		return err
	}
	if err := w.ReadData(&o.SectorSizeInBytes); err != nil {
		return err
	}
	_ptr_pParentPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ParentPath); err != nil {
			return err
		}
		return nil
	})
	_s_pParentPath := func(ptr interface{}) { o.ParentPath = *ptr.(*string) }
	if err := w.ReadPointer(&o.ParentPath, _s_pParentPath, _ptr_pParentPath); err != nil {
		return err
	}
	_ptr_pSourcePath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.SourcePath); err != nil {
			return err
		}
		return nil
	})
	_s_pSourcePath := func(ptr interface{}) { o.SourcePath = *ptr.(*string) }
	if err := w.ReadPointer(&o.SourcePath, _s_pSourcePath, _ptr_pSourcePath); err != nil {
		return err
	}
	return nil
}

// VDiskState type represents VDS_VDISK_STATE RPC enumeration.
//
// The VDS_VDISK_STATE enumeration describes the state of a virtual disk.
type VDiskState uint16

var (
	// VDS_VST_UNKNOWN:  VDS was not able to identify the disk's current status.
	VDiskStateUnknown VDiskState = 0
	// VDS_VST_ADDED:  The virtual disk is added to the service's list of objects.
	VDiskStateAdded VDiskState = 1
	// VDS_VST_OPEN:  The virtual disk has been added to the service's list of objects,
	// and the virtual disk file has been opened using IVdsVDisk::Open.
	VDiskStateOpen VDiskState = 2
	// VDS_VST_ATTACH_PENDING:  The virtual disk has been added to the service's list of
	// objects, the virtual disk file has been opened using IVdsVDisk::Open, and the virtual
	// disk is in the process of being attached.
	VDiskStateAttachPending VDiskState = 3
	// VDS_VST_ATTACHED_NOT_OPEN:  The virtual disk has been added to the service's list
	// of objects and the virtual disk is attached, but the virtual disk file is not open.
	VDiskStateAttachedNotOpen VDiskState = 4
	// VDS_VST_ATTACHED:  The virtual disk has been added to the service's list of objects,
	// the virtual disk file has been opened using IVdsVDisk::Open, and the virtual disk
	// is attached.
	VDiskStateAttached VDiskState = 5
	// VDS_VST_DETACH_PENDING:  The virtual disk has been added to the service's list of
	// objects, the virtual disk file has been opened using IVdsVDisk::Open, and the virtual
	// disk is in the process of being detached.
	VDiskStateDetachPending VDiskState = 6
	// VDS_VST_COMPACTING:  The virtual disk has been added to the service's list of objects,
	// the virtual disk file has been opened using IVdsVDisk::Open, and the virtual disk
	// is being compacted.
	VDiskStateCompacting VDiskState = 7
	// VDS_VST_MERGING:  The virtual disk has been added to the service's list of objects,
	// the virtual disk file has been opened using IVdsVDisk::Open, and the virtual disk
	// is being merged.
	VDiskStateMerging VDiskState = 8
	// VDS_VST_EXPANDING:  The virtual disk has been added to the service's list of objects,
	// the virtual disk file has been opened using IVdsVDisk::Open, and the virtual disk
	// is being expanded.
	VDiskStateExpanding VDiskState = 9
	// VDS_VST_DELETED:  The virtual disk has been deleted.
	VDiskStateDeleted VDiskState = 10
	// VDS_VST_MAX:  Denotes the maximum acceptable value for this type. VDS_VST_MAX -
	// 1 is the maximum acceptable value.
	//
	// When the service has been made aware of a virtual disk, the state is set to VDS_VST_ADDED.
	// In order to perform any operations on the virtual disk such as attaching it, detaching
	// it, merging, compacting, or expanding, the virtual disk file is opened using IVdsVDisk::Open.
	// Once the virtual disk is opened, its state transitions to VDS_VST_OPEN.
	//
	// Attach: To attach a virtual disk, the virtual disk object is first added to the service's
	// list of objects and its state is set to VDS_VS_ADDED. Next IVdsVDisk::Open MUST be
	// called against the virtual disk, and the state transitions to VDS_VST_OPEN. When
	// the attach operation is initiated against the virtual disk, the state of the virtual
	// disk transitions to VDS_VST_ATTACH_PENDING. The virtual disk remains in this state
	// until the operating system disk object corresponding to the virtual disk has been
	// instantiated. Once this object is instantiated, the virtual disk object's state transitions
	// to VDS_VST_ATTACHED. The IVdsOpenVDisk interface is then released, the OpenVirtualDisk
	// object is removed, and the state transitions to VDS_VST_ATTACHED_NOT_OPEN.
	//
	// Detach: To detach a virtual disk, the virtual disk object is first added to the service's
	// list of objects and its state is set to VDS_VST_ADDED. Next IVdsVDisk::Open MUST
	// be called against the virtual disk, and the state transitions to VDS_VST_OPEN. When
	// the detach operation is initiated against the virtual disk, the state of the virtual
	// disk transitions to VDS_VST_DETACH_PENDING. The virtual disk remains in this state
	// until the operating system disk object corresponding to the virtual disk has been
	// removed. Once this object is removed, the virtual disk object's state transitions
	// to VDS_VST_OPEN. The IVdsOpenVDisk interface is then released, the OpenVirtualDisk
	// object is removed, and the state transitions to VDS_VST_ADDED.
	VDiskStateMax VDiskState = 11
)

func (o VDiskState) String() string {
	switch o {
	case VDiskStateUnknown:
		return "VDiskStateUnknown"
	case VDiskStateAdded:
		return "VDiskStateAdded"
	case VDiskStateOpen:
		return "VDiskStateOpen"
	case VDiskStateAttachPending:
		return "VDiskStateAttachPending"
	case VDiskStateAttachedNotOpen:
		return "VDiskStateAttachedNotOpen"
	case VDiskStateAttached:
		return "VDiskStateAttached"
	case VDiskStateDetachPending:
		return "VDiskStateDetachPending"
	case VDiskStateCompacting:
		return "VDiskStateCompacting"
	case VDiskStateMerging:
		return "VDiskStateMerging"
	case VDiskStateExpanding:
		return "VDiskStateExpanding"
	case VDiskStateDeleted:
		return "VDiskStateDeleted"
	case VDiskStateMax:
		return "VDiskStateMax"
	}
	return "Invalid"
}

// AttachVirtualDiskFlag type represents ATTACH_VIRTUAL_DISK_FLAG RPC enumeration.
//
// The ATTACH_VIRTUAL_DISK_FLAG enumeration defines options for attaching a virtual
// disk.
type AttachVirtualDiskFlag uint16

var (
	// ATTACH_VIRTUAL_DISK_FLAG_NONE:  Indicates that no flag has been specified. This
	// flag implies that the operating system disk device created when the virtual disk
	// is attached will be read\write.
	AttachVirtualDiskFlagNone AttachVirtualDiskFlag = 0
	// ATTACH_VIRTUAL_DISK_FLAG_READ_ONLY:  Attaches the operating system disk device created
	// when the virtual disk is attached as read-only.
	AttachVirtualDiskFlagReadOnly AttachVirtualDiskFlag = 1
	// ATTACH_VIRTUAL_DISK_FLAG_NO_DRIVE_LETTER:  If this flag is set, no drive letters
	// are assigned to the disk's volumes.
	AttachVirtualDiskFlagNoDriveLetter AttachVirtualDiskFlag = 2
	// ATTACH_VIRTUAL_DISK_FLAG_PERMANENT_LIFETIME:  MUST NOT be used by virtual disk providers
	// or their clients.<52>
	AttachVirtualDiskFlagPermanentLifetime AttachVirtualDiskFlag = 4
	// ATTACH_VIRTUAL_DISK_FLAG_NO_LOCAL_HOST:  Specifies that the operating system disk
	// device created when the virtual disk is attached will not be exposed to the local
	// system, but rather to a virtual machine running on the local system.
	AttachVirtualDiskFlagNoLocalHost AttachVirtualDiskFlag = 8
)

func (o AttachVirtualDiskFlag) String() string {
	switch o {
	case AttachVirtualDiskFlagNone:
		return "AttachVirtualDiskFlagNone"
	case AttachVirtualDiskFlagReadOnly:
		return "AttachVirtualDiskFlagReadOnly"
	case AttachVirtualDiskFlagNoDriveLetter:
		return "AttachVirtualDiskFlagNoDriveLetter"
	case AttachVirtualDiskFlagPermanentLifetime:
		return "AttachVirtualDiskFlagPermanentLifetime"
	case AttachVirtualDiskFlagNoLocalHost:
		return "AttachVirtualDiskFlagNoLocalHost"
	}
	return "Invalid"
}

// DetachVirtualDiskFlag type represents DETACH_VIRTUAL_DISK_FLAG RPC enumeration.
//
// The DETACH_VIRTUAL_DISK_FLAG enumeration defines options for detaching a virtual
// disk.
type DetachVirtualDiskFlag uint16

var (
	// DETACH_VIRTUAL_DISK_FLAG_NONE:  Indicates that no flag has been specified. Currently,
	// this is the only flag defined.
	DetachVirtualDiskFlagNone DetachVirtualDiskFlag = 0
)

func (o DetachVirtualDiskFlag) String() string {
	switch o {
	case DetachVirtualDiskFlagNone:
		return "DetachVirtualDiskFlagNone"
	}
	return "Invalid"
}

// CompactVirtualDiskFlag type represents COMPACT_VIRTUAL_DISK_FLAG RPC enumeration.
//
// The COMPACT_VIRTUAL_DISK_FLAG enumeration defines options for compacting a virtual
// disk.
type CompactVirtualDiskFlag uint16

var (
	// COMPACT_VIRTUAL_DISK_FLAG_NONE:  Indicates that no flag has been specified. Currently,
	// this is the only flag defined.
	CompactVirtualDiskFlagNone CompactVirtualDiskFlag = 0
)

func (o CompactVirtualDiskFlag) String() string {
	switch o {
	case CompactVirtualDiskFlagNone:
		return "CompactVirtualDiskFlagNone"
	}
	return "Invalid"
}

// MergeVirtualDiskFlag type represents MERGE_VIRTUAL_DISK_FLAG RPC enumeration.
//
// The MERGE_VIRTUAL_DISK_FLAG enumeration defines options for merging a virtual disk.
type MergeVirtualDiskFlag uint16

var (
	// MERGE_VIRTUAL_DISK_FLAG_NONE:  Indicates that no flag has been specified. Currently,
	// this is the only flag defined.
	MergeVirtualDiskFlagNone MergeVirtualDiskFlag = 0
)

func (o MergeVirtualDiskFlag) String() string {
	switch o {
	case MergeVirtualDiskFlagNone:
		return "MergeVirtualDiskFlagNone"
	}
	return "Invalid"
}

// ExpandVirtualDiskFlag type represents EXPAND_VIRTUAL_DISK_FLAG RPC enumeration.
//
// The EXPAND_VIRTUAL_DISK_FLAG enumeration defines options for expanding a virtual
// disk.
type ExpandVirtualDiskFlag uint16

var (
	// EXPAND_VIRTUAL_DISK_FLAG_NONE:  Indicates that no flag has been specified. Currently,
	// this is the only flag defined.
	ExpandVirtualDiskFlagNone ExpandVirtualDiskFlag = 0
)

func (o ExpandVirtualDiskFlag) String() string {
	switch o {
	case ExpandVirtualDiskFlagNone:
		return "ExpandVirtualDiskFlagNone"
	}
	return "Invalid"
}

// DependentDiskFlag type represents DEPENDENT_DISK_FLAG RPC enumeration.
//
// The DEPENDENT_DISK_FLAG enumeration contains virtual disk dependency information
// flags.
type DependentDiskFlag uint16

var (
	// DEPENDENT_DISK_FLAG_NONE:  No flags specified. Use system defaults.
	DependentDiskFlagNone DependentDiskFlag = 0
	// DEPENDENT_DISK_FLAG_MULT_BACKING_FILES:  Multiple files backing the virtual disk.
	DependentDiskFlagMultBackingFiles DependentDiskFlag = 1
	// DEPENDENT_DISK_FLAG_FULLY_ALLOCATED:  Fully allocated virtual disk.
	DependentDiskFlagFullyAllocated DependentDiskFlag = 2
	// DEPENDENT_DISK_FLAG_READ_ONLY:  Read-only virtual disk.
	DependentDiskFlagReadOnly DependentDiskFlag = 4
	// DEPENDENT_DISK_FLAG_REMOTE:  The backing file of the virtual disk is not on a local
	// physical disk.
	DependentDiskFlagRemote DependentDiskFlag = 8
	// DEPENDENT_DISK_FLAG_SYSTEM_VOLUME:  Reserved.
	DependentDiskFlagSystemVolume DependentDiskFlag = 16
	// DEPENDENT_DISK_FLAG_SYSTEM_VOLUME_PARENT:  The backing file of the virtual disk
	// is on the system volume.
	DependentDiskFlagSystemVolumeParent DependentDiskFlag = 32
	// DEPENDENT_DISK_FLAG_REMOVABLE:  The backing file of the virtual disk is on a removable
	// physical disk.
	DependentDiskFlagRemovable DependentDiskFlag = 64
	// DEPENDENT_DISK_FLAG_NO_DRIVE_LETTER:  Drive letters are not automatically assigned
	// to the volumes on the virtual disk.
	DependentDiskFlagNoDriveLetter DependentDiskFlag = 128
	// DEPENDENT_DISK_FLAG_PARENT:  The virtual disk is a parent in a differencing chain.
	DependentDiskFlagParent DependentDiskFlag = 256
	// DEPENDENT_DISK_FLAG_NO_HOST_DISK:  The virtual disk is not surfaced on (attached
	// to) the local host. For example, it is attached to a guest virtual machine.
	DependentDiskFlagNoHostDisk DependentDiskFlag = 512
	// DEPENDENT_DISK_FLAG_PERMANENT_LIFETIME:  The lifetime of the virtual disk is not
	// tied to any application or process.
	DependentDiskFlagPermanentLifetime DependentDiskFlag = 1024
)

func (o DependentDiskFlag) String() string {
	switch o {
	case DependentDiskFlagNone:
		return "DependentDiskFlagNone"
	case DependentDiskFlagMultBackingFiles:
		return "DependentDiskFlagMultBackingFiles"
	case DependentDiskFlagFullyAllocated:
		return "DependentDiskFlagFullyAllocated"
	case DependentDiskFlagReadOnly:
		return "DependentDiskFlagReadOnly"
	case DependentDiskFlagRemote:
		return "DependentDiskFlagRemote"
	case DependentDiskFlagSystemVolume:
		return "DependentDiskFlagSystemVolume"
	case DependentDiskFlagSystemVolumeParent:
		return "DependentDiskFlagSystemVolumeParent"
	case DependentDiskFlagRemovable:
		return "DependentDiskFlagRemovable"
	case DependentDiskFlagNoDriveLetter:
		return "DependentDiskFlagNoDriveLetter"
	case DependentDiskFlagParent:
		return "DependentDiskFlagParent"
	case DependentDiskFlagNoHostDisk:
		return "DependentDiskFlagNoHostDisk"
	case DependentDiskFlagPermanentLifetime:
		return "DependentDiskFlagPermanentLifetime"
	}
	return "Invalid"
}

// VDiskProperties structure represents VDS_VDISK_PROPERTIES RPC structure.
//
// The VDS_VDISK_PROPERTIES structure defines the properties of a virtual disk.
type VDiskProperties struct {
	// Id:  A unique VDS-specific session identifier of the disk.
	ID *ObjectID `idl:"name:Id" json:"id"`
	// State:  A VDS_VDISK_STATE enumeration value that specifies the virtual disk state.
	State VDiskState `idl:"name:State" json:"state"`
	// VirtualDeviceType:  A pointer to a VIRTUAL_STORAGE_TYPE structure that specifies
	// the storage device type of the virtual disk.
	VirtualDeviceType *VirtualStorageType `idl:"name:VirtualDeviceType" json:"virtual_device_type"`
	// VirtualSize:  The size, in bytes, of the virtual disk.
	VirtualSize uint64 `idl:"name:VirtualSize" json:"virtual_size"`
	// PhysicalSize:  The on-disk size, in bytes, of the virtual hard disk backing file.
	PhysicalSize uint64 `idl:"name:PhysicalSize" json:"physical_size"`
	// pPath:  A null-terminated wide-character string containing the name and directory
	// path of the backing file for the virtual hard disk.
	Path string `idl:"name:pPath;string" json:"path"`
	// pDeviceName:  A null-terminated wide-character string containing the name and device
	// path of the disk device object for the volume where the virtual hard disk resides.
	DeviceName string `idl:"name:pDeviceName;string" json:"device_name"`
	// DiskFlag:  Type of virtual disk that uses values from the DEPENDENT_DISK_FLAG (section
	// 2.2.2.19.1.3) enumeration.
	DiskFlag DependentDiskFlag `idl:"name:DiskFlag" json:"disk_flag"`
	// bIsChild:  A Boolean value that specifies if the virtual disk is a child virtual
	// disk.
	IsChild bool `idl:"name:bIsChild" json:"is_child"`
	// pParentPath:  A null-terminated wide-character string containing an optional path
	// to the parent virtual disk.
	ParentPath string `idl:"name:pParentPath;string" json:"parent_path"`
}

func (o *VDiskProperties) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *VDiskProperties) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.ID != nil {
		if err := o.ID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint16(o.State)); err != nil {
		return err
	}
	if o.VirtualDeviceType != nil {
		if err := o.VirtualDeviceType.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&VirtualStorageType{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.VirtualSize); err != nil {
		return err
	}
	if err := w.WriteData(o.PhysicalSize); err != nil {
		return err
	}
	if o.Path != "" {
		_ptr_pPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Path, _ptr_pPath); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DeviceName != "" {
		_ptr_pDeviceName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DeviceName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DeviceName, _ptr_pDeviceName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint16(o.DiskFlag)); err != nil {
		return err
	}
	if !o.IsChild {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if o.ParentPath != "" {
		_ptr_pParentPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ParentPath); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ParentPath, _ptr_pParentPath); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *VDiskProperties) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.ID == nil {
		o.ID = &ObjectID{}
	}
	if err := o.ID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.State)); err != nil {
		return err
	}
	if o.VirtualDeviceType == nil {
		o.VirtualDeviceType = &VirtualStorageType{}
	}
	if err := o.VirtualDeviceType.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.VirtualSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.PhysicalSize); err != nil {
		return err
	}
	_ptr_pPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
			return err
		}
		return nil
	})
	_s_pPath := func(ptr interface{}) { o.Path = *ptr.(*string) }
	if err := w.ReadPointer(&o.Path, _s_pPath, _ptr_pPath); err != nil {
		return err
	}
	_ptr_pDeviceName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DeviceName); err != nil {
			return err
		}
		return nil
	})
	_s_pDeviceName := func(ptr interface{}) { o.DeviceName = *ptr.(*string) }
	if err := w.ReadPointer(&o.DeviceName, _s_pDeviceName, _ptr_pDeviceName); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.DiskFlag)); err != nil {
		return err
	}
	var _bIsChild int32
	if err := w.ReadData(&_bIsChild); err != nil {
		return err
	}
	o.IsChild = _bIsChild != 0
	_ptr_pParentPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ParentPath); err != nil {
			return err
		}
		return nil
	})
	_s_pParentPath := func(ptr interface{}) { o.ParentPath = *ptr.(*string) }
	if err := w.ReadPointer(&o.ParentPath, _s_pParentPath, _ptr_pParentPath); err != nil {
		return err
	}
	return nil
}

// VirtualDiskAccessMask type represents VIRTUAL_DISK_ACCESS_MASK RPC enumeration.
//
// The VIRTUAL_DISK_ACCESS_MASK enumeration contains the bit mask for specifying access
// rights to a virtual hard disk (VHD).
type VirtualDiskAccessMask uint32

var (
	// VIRTUAL_DISK_ACCESS_SURFACE_RO:  Open the VHD for read-only surfacing (attaching)
	// access. The caller MUST have READ access to the virtual disk image file. If used
	// in a request to open a VHD that is already open, the other handles are limited to
	// either VIRTUAL_DISK_ACCESS_UNSURFACE or VIRTUAL_DISK_ACCESS_GET_INFO access; otherwise,
	// the open request with this flag will fail.
	VirtualDiskAccessMaskSurfaceReadOnly VirtualDiskAccessMask = 65536
	// VIRTUAL_DISK_ACCESS_SURFACE_RW:  Open the VHD for read-write surfacing (attaching) access. The caller MUST have (READ | WRITE) access to the virtual disk image file. If used in a request to open a VHD that is already open, the other handles are limited to either VIRTUAL_DISK_ACCESS_UNSURFACE or VIRTUAL_DISK_ACCESS_GET_INFO access; otherwise, the open request with this flag will fail. If the VHD is part of a differencing chain, the disk number for this request cannot be less than the ReadWriteDepth specified during the prior open request for that differencing chain.
	VirtualDiskAccessMaskSurfaceRW VirtualDiskAccessMask = 131072
	// VIRTUAL_DISK_ACCESS_UNSURFACE:  Open the VHD to allow unsurfacing (detaching) of a surfaced (attached) VHD. The caller MUST have (FILE_READ_ATTRIBUTES | FILE_READ_DATA) access to the virtual disk image file.
	VirtualDiskAccessMaskUnsurface VirtualDiskAccessMask = 262144
	// VIRTUAL_DISK_ACCESS_GET_INFO:  Open the VHD for retrieval of information. The caller
	// MUST have READ access to the virtual disk image file.
	VirtualDiskAccessMaskGetInfo VirtualDiskAccessMask = 524288
	// VIRTUAL_DISK_ACCESS_CREATE:  Open the VHD for creation.
	VirtualDiskAccessMaskCreate VirtualDiskAccessMask = 1048576
	// VIRTUAL_DISK_ACCESS_METAOPS:  Open the VHD to perform offline metaoperations. For information on the offline metaoperations, see [MSDN-CompactVirtualDisk], [MSDN-ExpandVirtualDisk], [MSDN-MergeVirtualDisk], [MSDN-SetVirtualDiskInfo], and [MSDN-VIRTDSKACCMSK]. The caller MUST have (READ | WRITE) access to the virtual disk image file, up to ReadWriteDepth if working with a differencing chain. If the VHD is part of a differencing chain, the backing store (host volume) is opened in read/write exclusive mode up to ReadWriteDepth.
	VirtualDiskAccessMaskMetaops VirtualDiskAccessMask = 2097152
	// VIRTUAL_DISK_ACCESS_READ:  Reserved.
	VirtualDiskAccessMaskRead VirtualDiskAccessMask = 851968
	// VIRTUAL_DISK_ACCESS_ALL:  Allows unrestricted access to the VHD. The caller MUST
	// have unrestricted access rights to the virtual disk image file.
	VirtualDiskAccessMaskAll VirtualDiskAccessMask = 4128768
	// VIRTUAL_DISK_ACCESS_WRITABLE:  Reserved.
	VirtualDiskAccessMaskWritable VirtualDiskAccessMask = 3276800
)

func (o VirtualDiskAccessMask) String() string {
	switch o {
	case VirtualDiskAccessMaskSurfaceReadOnly:
		return "VirtualDiskAccessMaskSurfaceReadOnly"
	case VirtualDiskAccessMaskSurfaceRW:
		return "VirtualDiskAccessMaskSurfaceRW"
	case VirtualDiskAccessMaskUnsurface:
		return "VirtualDiskAccessMaskUnsurface"
	case VirtualDiskAccessMaskGetInfo:
		return "VirtualDiskAccessMaskGetInfo"
	case VirtualDiskAccessMaskCreate:
		return "VirtualDiskAccessMaskCreate"
	case VirtualDiskAccessMaskMetaops:
		return "VirtualDiskAccessMaskMetaops"
	case VirtualDiskAccessMaskRead:
		return "VirtualDiskAccessMaskRead"
	case VirtualDiskAccessMaskAll:
		return "VirtualDiskAccessMaskAll"
	case VirtualDiskAccessMaskWritable:
		return "VirtualDiskAccessMaskWritable"
	}
	return "Invalid"
}

// VolumePlex structure represents IVdsVolumePlex RPC structure.
type VolumePlex dcom.InterfacePointer

func (o *VolumePlex) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *VolumePlex) xxx_PreparePayload(ctx context.Context) error {
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
func (o *VolumePlex) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *VolumePlex) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *VolumePlex) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ServiceHBA structure represents IVdsServiceHba RPC structure.
type ServiceHBA dcom.InterfacePointer

func (o *ServiceHBA) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *ServiceHBA) xxx_PreparePayload(ctx context.Context) error {
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
func (o *ServiceHBA) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ServiceHBA) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ServiceHBA) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ServiceSAN structure represents IVdsServiceSAN RPC structure.
type ServiceSAN dcom.InterfacePointer

func (o *ServiceSAN) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *ServiceSAN) xxx_PreparePayload(ctx context.Context) error {
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
func (o *ServiceSAN) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ServiceSAN) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ServiceSAN) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// SwProvider structure represents IVdsSwProvider RPC structure.
type SwProvider dcom.InterfacePointer

func (o *SwProvider) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *SwProvider) xxx_PreparePayload(ctx context.Context) error {
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
func (o *SwProvider) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *SwProvider) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *SwProvider) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// VolumeOnline structure represents IVdsVolumeOnline RPC structure.
type VolumeOnline dcom.InterfacePointer

func (o *VolumeOnline) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *VolumeOnline) xxx_PreparePayload(ctx context.Context) error {
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
func (o *VolumeOnline) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *VolumeOnline) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *VolumeOnline) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// AdvancedDisk structure represents IVdsAdvancedDisk RPC structure.
type AdvancedDisk dcom.InterfacePointer

func (o *AdvancedDisk) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *AdvancedDisk) xxx_PreparePayload(ctx context.Context) error {
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
func (o *AdvancedDisk) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AdvancedDisk) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *AdvancedDisk) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// DiskOnline structure represents IVdsDiskOnline RPC structure.
type DiskOnline dcom.InterfacePointer

func (o *DiskOnline) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *DiskOnline) xxx_PreparePayload(ctx context.Context) error {
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
func (o *DiskOnline) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *DiskOnline) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *DiskOnline) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ServiceLoader structure represents IVdsServiceLoader RPC structure.
type ServiceLoader dcom.InterfacePointer

func (o *ServiceLoader) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *ServiceLoader) xxx_PreparePayload(ctx context.Context) error {
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
func (o *ServiceLoader) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ServiceLoader) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ServiceLoader) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// HBAPort structure represents IVdsHbaPort RPC structure.
type HBAPort dcom.InterfacePointer

func (o *HBAPort) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *HBAPort) xxx_PreparePayload(ctx context.Context) error {
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
func (o *HBAPort) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *HBAPort) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *HBAPort) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Volume structure represents IVdsVolume RPC structure.
type Volume dcom.InterfacePointer

func (o *Volume) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Volume) xxx_PreparePayload(ctx context.Context) error {
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
func (o *Volume) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Volume) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Volume) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ServiceSw structure represents IVdsServiceSw RPC structure.
type ServiceSw dcom.InterfacePointer

func (o *ServiceSw) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *ServiceSw) xxx_PreparePayload(ctx context.Context) error {
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
func (o *ServiceSw) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ServiceSw) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ServiceSw) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ServiceISCSI structure represents IVdsServiceIscsi RPC structure.
type ServiceISCSI dcom.InterfacePointer

func (o *ServiceISCSI) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *ServiceISCSI) xxx_PreparePayload(ctx context.Context) error {
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
func (o *ServiceISCSI) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ServiceISCSI) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ServiceISCSI) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Disk structure represents IVdsDisk RPC structure.
type Disk dcom.InterfacePointer

func (o *Disk) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Disk) xxx_PreparePayload(ctx context.Context) error {
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
func (o *Disk) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Disk) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Disk) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Async structure represents IVdsAsync RPC structure.
type Async dcom.InterfacePointer

func (o *Async) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Async) xxx_PreparePayload(ctx context.Context) error {
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
func (o *Async) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Async) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Async) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// CreatePartitionEx structure represents IVdsCreatePartitionEx RPC structure.
type CreatePartitionEx dcom.InterfacePointer

func (o *CreatePartitionEx) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *CreatePartitionEx) xxx_PreparePayload(ctx context.Context) error {
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
func (o *CreatePartitionEx) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *CreatePartitionEx) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *CreatePartitionEx) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ServiceUninstallDisk structure represents IVdsServiceUninstallDisk RPC structure.
type ServiceUninstallDisk dcom.InterfacePointer

func (o *ServiceUninstallDisk) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *ServiceUninstallDisk) xxx_PreparePayload(ctx context.Context) error {
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
func (o *ServiceUninstallDisk) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ServiceUninstallDisk) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ServiceUninstallDisk) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ISCSIInitiatorAdapter structure represents IVdsIscsiInitiatorAdapter RPC structure.
type ISCSIInitiatorAdapter dcom.InterfacePointer

func (o *ISCSIInitiatorAdapter) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *ISCSIInitiatorAdapter) xxx_PreparePayload(ctx context.Context) error {
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
func (o *ISCSIInitiatorAdapter) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ISCSIInitiatorAdapter) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ISCSIInitiatorAdapter) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// VDiskProvider structure represents IVdsVdProvider RPC structure.
type VDiskProvider dcom.InterfacePointer

func (o *VDiskProvider) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *VDiskProvider) xxx_PreparePayload(ctx context.Context) error {
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
func (o *VDiskProvider) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *VDiskProvider) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *VDiskProvider) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// SubSystemImportTarget structure represents IVdsSubSystemImportTarget RPC structure.
type SubSystemImportTarget dcom.InterfacePointer

func (o *SubSystemImportTarget) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *SubSystemImportTarget) xxx_PreparePayload(ctx context.Context) error {
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
func (o *SubSystemImportTarget) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *SubSystemImportTarget) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *SubSystemImportTarget) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// DiskPartitionMF2 structure represents IVdsDiskPartitionMF2 RPC structure.
type DiskPartitionMF2 dcom.InterfacePointer

func (o *DiskPartitionMF2) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *DiskPartitionMF2) xxx_PreparePayload(ctx context.Context) error {
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
func (o *DiskPartitionMF2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *DiskPartitionMF2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *DiskPartitionMF2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// VDisk structure represents IVdsVDisk RPC structure.
type VDisk dcom.InterfacePointer

func (o *VDisk) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *VDisk) xxx_PreparePayload(ctx context.Context) error {
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
func (o *VDisk) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *VDisk) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *VDisk) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Service structure represents IVdsService RPC structure.
type Service dcom.InterfacePointer

func (o *Service) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Service) xxx_PreparePayload(ctx context.Context) error {
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
func (o *Service) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Service) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Service) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ServiceInitialization structure represents IVdsServiceInitialization RPC structure.
type ServiceInitialization dcom.InterfacePointer

func (o *ServiceInitialization) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *ServiceInitialization) xxx_PreparePayload(ctx context.Context) error {
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
func (o *ServiceInitialization) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ServiceInitialization) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ServiceInitialization) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// EnumObject structure represents IEnumVdsObject RPC structure.
type EnumObject dcom.InterfacePointer

func (o *EnumObject) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *EnumObject) xxx_PreparePayload(ctx context.Context) error {
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
func (o *EnumObject) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *EnumObject) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *EnumObject) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// AdviseSink structure represents IVdsAdviseSink RPC structure.
type AdviseSink dcom.InterfacePointer

func (o *AdviseSink) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *AdviseSink) xxx_PreparePayload(ctx context.Context) error {
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
func (o *AdviseSink) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AdviseSink) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *AdviseSink) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ISCSIInitiatorPortal structure represents IVdsIscsiInitiatorPortal RPC structure.
type ISCSIInitiatorPortal dcom.InterfacePointer

func (o *ISCSIInitiatorPortal) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *ISCSIInitiatorPortal) xxx_PreparePayload(ctx context.Context) error {
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
func (o *ISCSIInitiatorPortal) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ISCSIInitiatorPortal) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ISCSIInitiatorPortal) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// DiskPartitionMF structure represents IVdsDiskPartitionMF RPC structure.
type DiskPartitionMF dcom.InterfacePointer

func (o *DiskPartitionMF) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *DiskPartitionMF) xxx_PreparePayload(ctx context.Context) error {
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
func (o *DiskPartitionMF) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *DiskPartitionMF) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *DiskPartitionMF) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Provider structure represents IVdsProvider RPC structure.
type Provider dcom.InterfacePointer

func (o *Provider) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Provider) xxx_PreparePayload(ctx context.Context) error {
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
func (o *Provider) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Provider) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Provider) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// OpenVDisk structure represents IVdsOpenVDisk RPC structure.
type OpenVDisk dcom.InterfacePointer

func (o *OpenVDisk) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *OpenVDisk) xxx_PreparePayload(ctx context.Context) error {
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
func (o *OpenVDisk) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *OpenVDisk) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *OpenVDisk) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// VolumeShrink structure represents IVdsVolumeShrink RPC structure.
type VolumeShrink dcom.InterfacePointer

func (o *VolumeShrink) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *VolumeShrink) xxx_PreparePayload(ctx context.Context) error {
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
func (o *VolumeShrink) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *VolumeShrink) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *VolumeShrink) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Removable structure represents IVdsRemovable RPC structure.
type Removable dcom.InterfacePointer

func (o *Removable) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Removable) xxx_PreparePayload(ctx context.Context) error {
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
func (o *Removable) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Removable) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Removable) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// AdvancedDisk2 structure represents IVdsAdvancedDisk2 RPC structure.
type AdvancedDisk2 dcom.InterfacePointer

func (o *AdvancedDisk2) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *AdvancedDisk2) xxx_PreparePayload(ctx context.Context) error {
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
func (o *AdvancedDisk2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AdvancedDisk2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *AdvancedDisk2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// AdvancedDisk3 structure represents IVdsAdvancedDisk3 RPC structure.
type AdvancedDisk3 dcom.InterfacePointer

func (o *AdvancedDisk3) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *AdvancedDisk3) xxx_PreparePayload(ctx context.Context) error {
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
func (o *AdvancedDisk3) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AdvancedDisk3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *AdvancedDisk3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Volume2 structure represents IVdsVolume2 RPC structure.
type Volume2 dcom.InterfacePointer

func (o *Volume2) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Volume2) xxx_PreparePayload(ctx context.Context) error {
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
func (o *Volume2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Volume2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Volume2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// HwProvider structure represents IVdsHwProvider RPC structure.
type HwProvider dcom.InterfacePointer

func (o *HwProvider) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *HwProvider) xxx_PreparePayload(ctx context.Context) error {
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
func (o *HwProvider) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *HwProvider) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *HwProvider) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// VolumeMF structure represents IVdsVolumeMF RPC structure.
type VolumeMF dcom.InterfacePointer

func (o *VolumeMF) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *VolumeMF) xxx_PreparePayload(ctx context.Context) error {
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
func (o *VolumeMF) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *VolumeMF) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *VolumeMF) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Pack structure represents IVdsPack RPC structure.
type Pack dcom.InterfacePointer

func (o *Pack) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Pack) xxx_PreparePayload(ctx context.Context) error {
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
func (o *Pack) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Pack) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Pack) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// VolumeMF2 structure represents IVdsVolumeMF2 RPC structure.
type VolumeMF2 dcom.InterfacePointer

func (o *VolumeMF2) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *VolumeMF2) xxx_PreparePayload(ctx context.Context) error {
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
func (o *VolumeMF2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *VolumeMF2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *VolumeMF2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// VolumeMF3 structure represents IVdsVolumeMF3 RPC structure.
type VolumeMF3 dcom.InterfacePointer

func (o *VolumeMF3) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *VolumeMF3) xxx_PreparePayload(ctx context.Context) error {
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
func (o *VolumeMF3) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *VolumeMF3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *VolumeMF3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Pack2 structure represents IVdsPack2 RPC structure.
type Pack2 dcom.InterfacePointer

func (o *Pack2) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Pack2) xxx_PreparePayload(ctx context.Context) error {
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
func (o *Pack2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Pack2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Pack2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Disk2 structure represents IVdsDisk2 RPC structure.
type Disk2 dcom.InterfacePointer

func (o *Disk2) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Disk2) xxx_PreparePayload(ctx context.Context) error {
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
func (o *Disk2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Disk2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Disk2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Disk3 structure represents IVdsDisk3 RPC structure.
type Disk3 dcom.InterfacePointer

func (o *Disk3) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Disk3) xxx_PreparePayload(ctx context.Context) error {
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
func (o *Disk3) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Disk3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Disk3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
