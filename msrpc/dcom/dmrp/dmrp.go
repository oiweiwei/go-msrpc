// The dmrp package implements the DMRP client protocol.
//
// # Introduction
//
// The Disk Management Remote Protocol is a set of Distributed Component Object Model
// (DCOM) interfaces, as specified in [MS-DCOM], built for managing storage objects
// on a machine. The Disk Management Remote Protocol relies on detailed, low-level operating
// system and storage concepts. While the basic concepts are outlined in this specification,
// it is assumed that the reader has familiarity with these technologies.
//
// For background information on storage, disk, and volume concepts, see [MSDN-DISKMAN]
// and [MSDN-VOLMAN]. For the IDL specification, see sections 6.1 and 6.2.
//
// # Overview
//
// The Disk Management Remote Protocol provides a set of DCOM interfaces for managing
// storage objects, such as disks and volumes. The protocol also enables clients to
// obtain notifications of changes to storage objects. The server end of the protocol
// implements supports that let the DCOM interface handle requests for storage management
// services for a server system over the network. The client end of the protocol is
// an application that invokes method calls on the interface to perform various disk
// and volume configuration tasks.
//
// This protocol includes the following six DCOM interfaces:
//
// * IVolumeClient ( e31bc7db-aca1-4ca1-9d32-77b1415f8584 )
//
// * IVolumeClient2 ( 575cc030-9a49-40c8-9a6c-0c3de32ce514 )
//
// * IVolumeClient3 ( be8a9c4c-f7b2-496c-a0e1-cb0655cea66f )
//
// * IVolumeClient4 ( 7993e053-3bd7-4a09-ab55-209d2d31be6f )
//
// * IDMRemoteServer ( d7504302-df62-40a1-b3c1-7a44302559e7 )
//
// * IDMNotify ( 65b511db-f735-4225-856b-022bfd7a1a50 )
//
// The IVolumeClient and IVolumeClient2 interfaces provide methods for managing storage
// objects, such as disks and volumes.
//
// IVolumeClient3 supersedes IVolumeClient and IVolumeClient2, and contains new functionality
// related to the GUID partition table (GPT) disk-partitioning style. The IVolumeClient
// and IVolumeClient2 interfaces do not support the GPT disk-partitioning style and
// cannot be used with GPT partitioned disks.
//
// IVolumeClient4 includes additional functionality to augment what is provided by IVolumeClient3;
// IVolumeClient4 verifies that disk access and disk media record information is valid
// when the cache is refreshed, and it queries the device path for a volume.
//
// IDMRemoteServer includes functionality to create an instance of the Disk Management
// server on a remote machine.
//
// IDMNotify is the interface implemented by the client to receive notifications from
// the Disk Management server.<1>
package dmrp

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
	GoPackage = "dcom/dmrp"
)

// DiskAudioCD represents the DISK_AUDIO_CD RPC constant
const DiskAudioCD = 0x00000001

// DiskNec98 represents the DISK_NEC98 RPC constant
const DiskNec98 = 0x00000002

// DeviceTypeUnknown represents the DEVICETYPE_UNKNOWN RPC constant
var DeviceTypeUnknown = 0

// DeviceTypeVm represents the DEVICETYPE_VM RPC constant
var DeviceTypeVm = 1

// DeviceTypeRemovable represents the DEVICETYPE_REMOVABLE RPC constant
var DeviceTypeRemovable = 2

// DeviceTypeCDROM represents the DEVICETYPE_CDROM RPC constant
var DeviceTypeCDROM = 3

// DeviceTypeFdisk represents the DEVICETYPE_FDISK RPC constant
var DeviceTypeFdisk = 4

// DeviceTypeDVD represents the DEVICETYPE_DVD RPC constant
var DeviceTypeDVD = 5

// DevicestateUnknown represents the DEVICESTATE_UNKNOWN RPC constant
var DevicestateUnknown = 0

// DevicestateHealthy represents the DEVICESTATE_HEALTHY RPC constant
var DevicestateHealthy = 1

// DevicestateNoMedia represents the DEVICESTATE_NO_MEDIA RPC constant
var DevicestateNoMedia = 2

// DevicestateNoSig represents the DEVICESTATE_NOSIG RPC constant
var DevicestateNoSig = 4

// DevicestateBad represents the DEVICESTATE_BAD RPC constant
var DevicestateBad = 8

// DevicestateNotReady represents the DEVICESTATE_NOT_READY RPC constant
var DevicestateNotReady = 16

// DevicestateMissing represents the DEVICESTATE_MISSING RPC constant
var DevicestateMissing = 32

// DevicestateOffline represents the DEVICESTATE_OFFLINE RPC constant
var DevicestateOffline = 64

// DevicestateFailing represents the DEVICESTATE_FAILING RPC constant
var DevicestateFailing = 128

// DevicestateImportFailed represents the DEVICESTATE_IMPORT_FAILED RPC constant
var DevicestateImportFailed = 256

// DevicestateUnclaimed represents the DEVICESTATE_UNCLAIMED RPC constant
var DevicestateUnclaimed = 512

// BusTypeUnknown represents the BUSTYPE_UNKNOWN RPC constant
var BusTypeUnknown = 0

// BusTypeIDE represents the BUSTYPE_IDE RPC constant
var BusTypeIDE = 1

// BusTypeSCSI represents the BUSTYPE_SCSI RPC constant
var BusTypeSCSI = 2

// BusTypeFibre represents the BUSTYPE_FIBRE RPC constant
var BusTypeFibre = 3

// BusTypeUSB represents the BUSTYPE_USB RPC constant
var BusTypeUSB = 4

// BusTypeSSA represents the BUSTYPE_SSA RPC constant
var BusTypeSSA = 5

// BusType1394 represents the BUSTYPE_1394 RPC constant
var BusType1394 = 6

// DeviceAttributeNone represents the DEVICEATTR_NONE RPC constant
var DeviceAttributeNone = 0

// DeviceAttributeReadOnly represents the DEVICEATTR_RDONLY RPC constant
var DeviceAttributeReadOnly = 1

// DeviceAttributeNTMS represents the DEVICEATTR_NTMS RPC constant
var DeviceAttributeNTMS = 2

// ContainsFT represents the CONTAINS_FT RPC constant
var ContainsFT = 1

// ContainsRAID5 represents the CONTAINS_RAID5 RPC constant
var ContainsRAID5 = 2

// ContainsRedistribution represents the CONTAINS_REDISTRIBUTION RPC constant
var ContainsRedistribution = 4

// ContainsBootablePartition represents the CONTAINS_BOOTABLE_PARTITION RPC constant
var ContainsBootablePartition = 8

// ContainsLockedPartition represents the CONTAINS_LOCKED_PARTITION RPC constant
var ContainsLockedPartition = 16

// ContainsNoFreeSpace represents the CONTAINS_NO_FREE_SPACE RPC constant
var ContainsNoFreeSpace = 32

// ContainsExtendedPartition represents the CONTAINS_EXTENDED_PARTITION RPC constant
var ContainsExtendedPartition = 64

// PartitionNumberChange represents the PARTITION_NUMBER_CHANGE RPC constant
var PartitionNumberChange = 128

// ContainsBootIndicator represents the CONTAINS_BOOTINDICATOR RPC constant
var ContainsBootIndicator = 256

// ContainsBootLoader represents the CONTAINS_BOOTLOADER RPC constant
var ContainsBootLoader = 512

// ContainsSystemDir represents the CONTAINS_SYSTEMDIR RPC constant
var ContainsSystemDir = 1024

// ContainsMixedPartitions represents the CONTAINS_MIXED_PARTITIONS RPC constant
var ContainsMixedPartitions = 2048

// PartitionOS2Boot represents the PARTITION_OS2_BOOT RPC constant
const PartitionOS2Boot = 0x0000000A

// PartitionEISA represents the PARTITION_EISA RPC constant
const PartitionEISA = 0x00000012

// PartitionHibernation represents the PARTITION_HIBERNATION RPC constant
const PartitionHibernation = 0x00000084

// PartitionDiagnostic represents the PARTITION_DIAGNOSTIC RPC constant
const PartitionDiagnostic = 0x000000A0

// PartitionDell represents the PARTITION_DELL RPC constant
const PartitionDell = 0x000000DE

// PartitionIBM represents the PARTITION_IBM RPC constant
const PartitionIBM = 0x000000FE

// RegionFormatInProgress represents the REGION_FORMAT_IN_PROGRESS RPC constant
const RegionFormatInProgress = 0x00000001

// VolumeFormatInProgress represents the VOLUME_FORMAT_IN_PROGRESS RPC constant
const VolumeFormatInProgress = 0x00000001

// RegionIsSystemPartition represents the REGION_IS_SYSTEM_PARTITION RPC constant
const RegionIsSystemPartition = 0x00000002

// RegionHasPageFile represents the REGION_HAS_PAGEFILE RPC constant
const RegionHasPageFile = 0x00000004

// VolumeHasPageFile represents the VOLUME_HAS_PAGEFILE RPC constant
const VolumeHasPageFile = 0x00000004

// RegionHadBootINI represents the REGION_HAD_BOOT_INI RPC constant
const RegionHadBootINI = 0x00000040

// VolumeIsBootVolume represents the VOLUME_IS_BOOT_VOLUME RPC constant
const VolumeIsBootVolume = 0x00000100

// VolumeIsRestartable represents the VOLUME_IS_RESTARTABLE RPC constant
const VolumeIsRestartable = 0x00000400

// VolumeIsSystemVolume represents the VOLUME_IS_SYSTEM_VOLUME RPC constant
const VolumeIsSystemVolume = 0x00000800

// VolumeHasRetainPartition represents the VOLUME_HAS_RETAIN_PARTITION RPC constant
const VolumeHasRetainPartition = 0x00001000

// VolumeHadBootINI represents the VOLUME_HAD_BOOT_INI RPC constant
const VolumeHadBootINI = 0x00002000

// VolumeCorrupt represents the VOLUME_CORRUPT RPC constant
const VolumeCorrupt = 0x00004000

// VolumeHasCrashDump represents the VOLUME_HAS_CRASHDUMP RPC constant
const VolumeHasCrashDump = 0x00008000

// VolumeIsCurrBootVolume represents the VOLUME_IS_CURR_BOOT_VOLUME RPC constant
const VolumeIsCurrBootVolume = 0x00010000

// VolumeHasHibernation represents the VOLUME_HAS_HIBERNATION RPC constant
const VolumeHasHibernation = 0x00020000

// NoForceOperation represents the NO_FORCE_OPERATION RPC constant
const NoForceOperation = 0x00000000

// ForceOperation represents the FORCE_OPERATION RPC constant
const ForceOperation = 0x00000001

// DLPendingRemoval represents the DL_PENDING_REMOVAL RPC constant
const DLPendingRemoval = 0x00000001

// SystemFlagServer represents the SYSFLAG_SERVER RPC constant
const SystemFlagServer = 0x00000001

// SystemFlagAlpha represents the SYSFLAG_ALPHA RPC constant
const SystemFlagAlpha = 0x00000002

// SystemFlagSyspartSecure represents the SYSFLAG_SYSPART_SECURE RPC constant
const SystemFlagSyspartSecure = 0x00000004

// SystemFlagNec98 represents the SYSFLAG_NEC_98 RPC constant
const SystemFlagNec98 = 0x00000008

// SystemFlagLaptop represents the SYSFLAG_LAPTOP RPC constant
const SystemFlagLaptop = 0x00000010

// SystemFlagWolfpack represents the SYSFLAG_WOLFPACK RPC constant
const SystemFlagWolfpack = 0x00000020

// DskmergeDelete represents the DSKMERGE_DELETE RPC constant
const DskmergeDelete = 0x00000001

// DskmergeDeleteRedundancy represents the DSKMERGE_DELETE_REDUNDANCY RPC constant
const DskmergeDeleteRedundancy = 0x00000002

// DskmergeStaleData represents the DSKMERGE_STALE_DATA RPC constant
const DskmergeStaleData = 0x00000004

// DskmergeRelated represents the DSKMERGE_RELATED RPC constant
const DskmergeRelated = 0x00000008

// DskmergeInNoUnrelated represents the DSKMERGE_IN_NO_UNRELATED RPC constant
const DskmergeInNoUnrelated = 0x00000001

// DskmergeOutNoPrimaryDiskGroup represents the DSKMERGE_OUT_NO_PRIMARY_DG RPC constant
const DskmergeOutNoPrimaryDiskGroup = 0x00000001

// FtreplaceForce represents the FTREPLACE_FORCE RPC constant
const FtreplaceForce = 0x00000001

// FtreplaceDeleteOnFail represents the FTREPLACE_DELETE_ON_FAIL RPC constant
const FtreplaceDeleteOnFail = 0x00000002

// CreateAssignAccessPath represents the CREATE_ASSIGN_ACCESS_PATH RPC constant
const CreateAssignAccessPath = 0x00000001

// EnableVolumeCompression represents the ENABLE_VOLUME_COMPRESSION RPC constant
const EnableVolumeCompression = 0x00000001

// MaxFSNameSize represents the MAX_FS_NAME_SIZE RPC constant
const MaxFSNameSize = 0x00000008

// FormatOptionCompress represents the FSF_FMT_OPTION_COMPRESS RPC constant
const FormatOptionCompress = 0x00000001

// FormatOptionLabel represents the FSF_FMT_OPTION_LABEL RPC constant
const FormatOptionLabel = 0x00000002

// MountPointSupport represents the FSF_MNT_POINT_SUPPORT RPC constant
const MountPointSupport = 0x00000004

// RemovableMediaSupport represents the FSF_REMOVABLE_MEDIA_SUPPORT RPC constant
const RemovableMediaSupport = 0x00000008

// FSGrowSupport represents the FSF_FS_GROW_SUPPORT RPC constant
const FSGrowSupport = 0x00000010

// FSQuickFormatEnable represents the FSF_FS_QUICK_FORMAT_ENABLE RPC constant
const FSQuickFormatEnable = 0x00000020

// FSAllocSize512 represents the FSF_FS_ALLOC_SZ_512 RPC constant
const FSAllocSize512 = 0x00000040

// FSAllocSize1K represents the FSF_FS_ALLOC_SZ_1K RPC constant
const FSAllocSize1K = 0x00000080

// FSAllocSize2K represents the FSF_FS_ALLOC_SZ_2K RPC constant
const FSAllocSize2K = 0x00000100

// FSAllocSize4K represents the FSF_FS_ALLOC_SZ_4K RPC constant
const FSAllocSize4K = 0x00000200

// FSAllocSize8K represents the FSF_FS_ALLOC_SZ_8K RPC constant
const FSAllocSize8K = 0x00000400

// FSAllocSize16K represents the FSF_FS_ALLOC_SZ_16K RPC constant
const FSAllocSize16K = 0x00000800

// FSAllocSize32K represents the FSF_FS_ALLOC_SZ_32K RPC constant
const FSAllocSize32K = 0x00001000

// FSAllocSize64K represents the FSF_FS_ALLOC_SZ_64K RPC constant
const FSAllocSize64K = 0x00002000

// FSAllocSize128K represents the FSF_FS_ALLOC_SZ_128K RPC constant
const FSAllocSize128K = 0x00004000

// FSAllocSize256K represents the FSF_FS_ALLOC_SZ_256K RPC constant
const FSAllocSize256K = 0x00008000

// FSAllocSizeOther represents the FSF_FS_ALLOC_SZ_OTHER RPC constant
const FSAllocSizeOther = 0x00010000

// FSFormatSupported represents the FSF_FS_FORMAT_SUPPORTED RPC constant
const FSFormatSupported = 0x00020000

// FSValidBits represents the FSF_FS_VALID_BITS RPC constant
const FSValidBits = 0x0003FFFF

// FSTypeUnknown represents the FSTYPE_UNKNOWN RPC constant
var FSTypeUnknown = 0

// FSTypeNTFS represents the FSTYPE_NTFS RPC constant
var FSTypeNTFS = 1

// FSTypeFAT represents the FSTYPE_FAT RPC constant
var FSTypeFAT = 2

// FSTypeFAT32 represents the FSTYPE_FAT32 RPC constant
var FSTypeFAT32 = 3

// FSTypeCDFS represents the FSTYPE_CDFS RPC constant
var FSTypeCDFS = 4

// FSTypeUDF represents the FSTYPE_UDF RPC constant
var FSTypeUDF = 5

// FSTypeOther represents the FSTYPE_OTHER RPC constant
var FSTypeOther = 2147483648

// EncapInfoCantProceed represents the ENCAP_INFO_CANT_PROCEED RPC constant
const EncapInfoCantProceed = 0x00000001

// EncapInfoNoFreeSpace represents the ENCAP_INFO_NO_FREE_SPACE RPC constant
const EncapInfoNoFreeSpace = 0x00000002

// EncapInfoBadActive represents the ENCAP_INFO_BAD_ACTIVE RPC constant
const EncapInfoBadActive = 0x00000004

// EncapInfoUnknownPart represents the ENCAP_INFO_UNKNOWN_PART RPC constant
const EncapInfoUnknownPart = 0x00000008

// EncapInfoFTUnhealthy represents the ENCAP_INFO_FT_UNHEALTHY RPC constant
const EncapInfoFTUnhealthy = 0x00000010

// EncapInfoFTQueryFailed represents the ENCAP_INFO_FT_QUERY_FAILED RPC constant
const EncapInfoFTQueryFailed = 0x00000020

// EncapInfoFTHasRAID5 represents the ENCAP_INFO_FT_HAS_RAID5 RPC constant
const EncapInfoFTHasRAID5 = 0x00000040

// EncapInfoFTOnBoot represents the ENCAP_INFO_FT_ON_BOOT RPC constant
const EncapInfoFTOnBoot = 0x00000080

// EncapInfoRebootRequired represents the ENCAP_INFO_REBOOT_REQD RPC constant
const EncapInfoRebootRequired = 0x00000100

// EncapInfoContainsFT represents the ENCAP_INFO_CONTAINS_FT RPC constant
const EncapInfoContainsFT = 0x00000200

// EncapInfoVolumeBusy represents the ENCAP_INFO_VOLUME_BUSY RPC constant
const EncapInfoVolumeBusy = 0x00000400

// EncapInfoPartNRChange represents the ENCAP_INFO_PART_NR_CHANGE RPC constant
const EncapInfoPartNRChange = 0x00000800

// SystemFlagNoDynamic represents the SYSFLAG_NO_DYNAMIC RPC constant
const SystemFlagNoDynamic = 0x00000010

// SystemFlagIA64 represents the SYSFLAG_IA64 RPC constant
const SystemFlagIA64 = 0x00000040

// SystemFlagUninstallValid represents the SYSFLAG_UNINSTALL_VALID RPC constant
const SystemFlagUninstallValid = 0x00000080

// SystemFlagDynamic1394 represents the SYSFLAG_DYNAMIC_1394 RPC constant
const SystemFlagDynamic1394 = 0x00000100

// DiskFormattableDVD represents the DISK_FORMATTABLE_DVD RPC constant
const DiskFormattableDVD = 0x00000004

// DiskMemoryStick represents the DISK_MEMORY_STICK RPC constant
const DiskMemoryStick = 0x00000008

// DiskNTFSNotSupported represents the DISK_NTFS_NOT_SUPPORTED RPC constant
const DiskNTFSNotSupported = 0x00000010

// RegionHidden represents the REGION_HIDDEN RPC constant
const RegionHidden = 0x00040000

// EncapInfoMixedPartitions represents the ENCAP_INFO_MIXED_PARTITIONS RPC constant
const EncapInfoMixedPartitions = 0x00001000

// EncapInfoOpenFailed represents the ENCAP_INFO_OPEN_FAILED RPC constant
const EncapInfoOpenFailed = 0x00002000

// RegionType type represents REGIONTYPE RPC enumeration.
//
// The REGIONTYPE enumeration defines values for region types.
type RegionType uint16

var (
	// REGION_UNKNOWN:  Region type is unknown.
	RegionTypeUnknown RegionType = 0
	// REGION_FREE:  Region resides in free space.
	RegionTypeFree RegionType = 1
	// REGION_EXTENDED_FREE:  Region resides in the free space of an extended partition.
	RegionTypeExtendedFree RegionType = 2
	// REGION_PRIMARY:  Region resides in a primary partition.
	RegionTypePrimary RegionType = 3
	// REGION_LOGICAL:  Region resides in a logical partition.
	RegionTypeLogical RegionType = 4
	// REGION_EXTENDED:  Region resides in an extended partition.
	RegionTypeExtended RegionType = 5
	// REGION_SUBDISK:  Region resides on a subdisk.
	RegionTypeSubdisk RegionType = 6
	// REGION_CDROM:  Region resides on a CD-ROM device.
	RegionTypeCDROM RegionType = 7
	// REGION_REMOVABLE:  Region resides on a device with removable media.
	RegionTypeRemovable RegionType = 8
)

func (o RegionType) String() string {
	switch o {
	case RegionTypeUnknown:
		return "RegionTypeUnknown"
	case RegionTypeFree:
		return "RegionTypeFree"
	case RegionTypeExtendedFree:
		return "RegionTypeExtendedFree"
	case RegionTypePrimary:
		return "RegionTypePrimary"
	case RegionTypeLogical:
		return "RegionTypeLogical"
	case RegionTypeExtended:
		return "RegionTypeExtended"
	case RegionTypeSubdisk:
		return "RegionTypeSubdisk"
	case RegionTypeCDROM:
		return "RegionTypeCDROM"
	case RegionTypeRemovable:
		return "RegionTypeRemovable"
	}
	return "Invalid"
}

// VolumeType type represents VOLUMETYPE RPC enumeration.
//
// The VOLUMETYPE enumeration defines values for types of volumes.
type VolumeType uint16

var (
	// VOLUMETYPE_UNKNOWN:  Volume is of an unknown type.
	VolumeTypeUnknown VolumeType = 0
	// VOLUMETYPE_PRIMARY_PARTITION:  Volume is a primary partition.
	VolumeTypePrimaryPartition VolumeType = 1
	// VOLUMETYPE_LOGICAL_DRIVE:  Volume is a logical drive.
	VolumeTypeLogicalDrive VolumeType = 2
	// VOLUMETYPE_FT:  Volume is a Windows NT 4.0 operating system–style fault-tolerant
	// volume.
	VolumeTypeFT VolumeType = 3
	// VOLUMETYPE_VM:  Volume is controlled by the volume manager.
	VolumeTypeVm VolumeType = 4
	// VOLUMETYPE_CDROM:  Volume resides on a CD-ROM device.
	VolumeTypeCDROM VolumeType = 5
	// VOLUMETYPE_REMOVABLE:  Volume resides on a device with removable media.
	VolumeTypeRemovable VolumeType = 6
)

func (o VolumeType) String() string {
	switch o {
	case VolumeTypeUnknown:
		return "VolumeTypeUnknown"
	case VolumeTypePrimaryPartition:
		return "VolumeTypePrimaryPartition"
	case VolumeTypeLogicalDrive:
		return "VolumeTypeLogicalDrive"
	case VolumeTypeFT:
		return "VolumeTypeFT"
	case VolumeTypeVm:
		return "VolumeTypeVm"
	case VolumeTypeCDROM:
		return "VolumeTypeCDROM"
	case VolumeTypeRemovable:
		return "VolumeTypeRemovable"
	}
	return "Invalid"
}

// VolumeLayout type represents VOLUMELAYOUT RPC enumeration.
//
// The VOLUMELAYOUT enumeration defines values for volume layouts.
type VolumeLayout uint16

var (
	// VOLUMELAYOUT_UNKNOWN:  Volume has an unknown layout.
	VolumeLayoutUnknown VolumeLayout = 0
	// VOLUMELAYOUT_PARTITION:  Volume is a partition.
	VolumeLayoutPartition VolumeLayout = 1
	// VOLUMELAYOUT_SIMPLE:  Volume is a basic disk.
	VolumeLayoutSimple VolumeLayout = 2
	// VOLUMELAYOUT_SPANNED:  Volume spans multiple disks.
	VolumeLayoutSpanned VolumeLayout = 3
	// VOLUMELAYOUT_MIRROR:  Volume is a mirror.
	VolumeLayoutMirror VolumeLayout = 4
	// VOLUMELAYOUT_STRIPE:  Volume is a striped set.
	VolumeLayoutStripe VolumeLayout = 5
	// VOLUMELAYOUT_RAID5:  Volume is a RAID-5 set.
	VolumeLayoutRAID5 VolumeLayout = 6
)

func (o VolumeLayout) String() string {
	switch o {
	case VolumeLayoutUnknown:
		return "VolumeLayoutUnknown"
	case VolumeLayoutPartition:
		return "VolumeLayoutPartition"
	case VolumeLayoutSimple:
		return "VolumeLayoutSimple"
	case VolumeLayoutSpanned:
		return "VolumeLayoutSpanned"
	case VolumeLayoutMirror:
		return "VolumeLayoutMirror"
	case VolumeLayoutStripe:
		return "VolumeLayoutStripe"
	case VolumeLayoutRAID5:
		return "VolumeLayoutRAID5"
	}
	return "Invalid"
}

// RequestStatus type represents REQSTATUS RPC enumeration.
//
// The REQSTATUS enumeration defines values for the status of a request.
type RequestStatus uint16

var (
	// REQ_UNKNOWN:  Request state is unknown.
	RequestStatusUnknown RequestStatus = 0
	// REQ_STARTED:  Request has started.
	RequestStatusStarted RequestStatus = 1
	// REQ_IN_PROGRESS:  Request is in progress.
	RequestStatusInProgress RequestStatus = 2
	// REQ_COMPLETED:  Request has finished.
	RequestStatusCompleted RequestStatus = 3
	// REQ_ABORTED:  Request has terminated.
	RequestStatusAborted RequestStatus = 4
	// REQ_FAILED:  Request has failed.
	RequestStatusFailed RequestStatus = 5
)

func (o RequestStatus) String() string {
	switch o {
	case RequestStatusUnknown:
		return "RequestStatusUnknown"
	case RequestStatusStarted:
		return "RequestStatusStarted"
	case RequestStatusInProgress:
		return "RequestStatusInProgress"
	case RequestStatusCompleted:
		return "RequestStatusCompleted"
	case RequestStatusAborted:
		return "RequestStatusAborted"
	case RequestStatusFailed:
		return "RequestStatusFailed"
	}
	return "Invalid"
}

// RegionStatus type represents REGIONSTATUS RPC enumeration.
//
// The REGIONSTATUS enumeration defines values for a region's status.
type RegionStatus uint16

var (
	// REGIONSTATUS_UNKNOWN:  Region's status is unknown.
	RegionStatusUnknown RegionStatus = 0
	// REGIONSTATUS_OK:  Region is intact.
	RegionStatusOK RegionStatus = 1
	// REGIONSTATUS_FAILED:  Region failed.
	RegionStatusFailed RegionStatus = 2
	// REGIONSTATUS_FAILING:  Region is in the process of failing.
	RegionStatusFailing RegionStatus = 3
	// REGIONSTATUS_REGENERATING:  Region is regenerating data from the fault-tolerant
	// check information.
	RegionStatusRegenerating RegionStatus = 4
	// REGIONSTATUS_NEEDSRESYNC:  Region needs resynchronization.
	RegionStatusNeedsResync RegionStatus = 5
)

func (o RegionStatus) String() string {
	switch o {
	case RegionStatusUnknown:
		return "RegionStatusUnknown"
	case RegionStatusOK:
		return "RegionStatusOK"
	case RegionStatusFailed:
		return "RegionStatusFailed"
	case RegionStatusFailing:
		return "RegionStatusFailing"
	case RegionStatusRegenerating:
		return "RegionStatusRegenerating"
	case RegionStatusNeedsResync:
		return "RegionStatusNeedsResync"
	}
	return "Invalid"
}

// VolumeStatus type represents VOLUMESTATUS RPC enumeration.
//
// The VOLUMESTATUS enumeration defines values for a volume's status. For more information
// about redundant data and fault-tolerant volumes, see [MSDN-DISKMAN].
type VolumeStatus uint16

var (
	// VOLUME_STATUS_UNKNOWN:  Volume has an unknown status.
	VolumeStatusUnknown VolumeStatus = 0
	// VOLUME_STATUS_HEALTHY:  Volume is fully functional.
	VolumeStatusHealthy VolumeStatus = 1
	// VOLUME_STATUS_FAILED:  Volume is in a failed state.
	VolumeStatusFailed VolumeStatus = 2
	// VOLUME_STATUS_FAILED_REDUNDANCY:  Volume's redundant data in a fault-tolerant volume
	// has failed.
	VolumeStatusFailedRedundancy VolumeStatus = 3
	// VOLUME_STATUS_FAILING:  Volume has encountered I/O errors.
	VolumeStatusFailing VolumeStatus = 4
	// VOLUME_STATUS_FAILING_REDUNDANCY:  Volume is fault-tolerant, and it encountered
	// I/O errors.
	VolumeStatusFailingRedundancy VolumeStatus = 5
	// VOLUME_STATUS_FAILED_REDUNDANCY_FAILING: Redundant data in a fault-tolerant volume
	// has failed, and the volume encountered I/O errors in the last remaining copy of the
	// data.
	VolumeStatusFailedRedundancyFailing VolumeStatus = 6
	// VOLUME_STATUS_SYNCHING:  Volume is resynchronizing fault-tolerant data for a mirrored
	// volume.
	VolumeStatusSynching VolumeStatus = 7
	// VOLUME_STATUS_REGENERATING:  Volume is regenerating fault-tolerant data for a RAID-5
	// volume.
	VolumeStatusRegenerating VolumeStatus = 8
	// VOLUME_STATUS_INITIALIZING:  Volume is initializing to volume manager control.
	VolumeStatusInitializing VolumeStatus = 9
	// VOLUME_STATUS_FORMATTING:  Volume is currently being formatted.
	VolumeStatusFormatting VolumeStatus = 10
)

func (o VolumeStatus) String() string {
	switch o {
	case VolumeStatusUnknown:
		return "VolumeStatusUnknown"
	case VolumeStatusHealthy:
		return "VolumeStatusHealthy"
	case VolumeStatusFailed:
		return "VolumeStatusFailed"
	case VolumeStatusFailedRedundancy:
		return "VolumeStatusFailedRedundancy"
	case VolumeStatusFailing:
		return "VolumeStatusFailing"
	case VolumeStatusFailingRedundancy:
		return "VolumeStatusFailingRedundancy"
	case VolumeStatusFailedRedundancyFailing:
		return "VolumeStatusFailedRedundancyFailing"
	case VolumeStatusSynching:
		return "VolumeStatusSynching"
	case VolumeStatusRegenerating:
		return "VolumeStatusRegenerating"
	case VolumeStatusInitializing:
		return "VolumeStatusInitializing"
	case VolumeStatusFormatting:
		return "VolumeStatusFormatting"
	}
	return "Invalid"
}

// Action type represents LDMACTION RPC enumeration.
//
// The LDMACTION enumeration defines the type of action described by an ObjectsChanged
// call.
type Action uint16

var (
	// LDMACTION_UNKNOWN:  Object underwent an unknown type of change.
	ActionUnknown Action = 0
	// LDMACTION_CREATED:  Object was created.
	ActionCreated Action = 1
	// LDMACTION_DELETED:  Object was deleted.
	ActionDeleted Action = 2
	// LDMACTION_MODIFIED:  Object was modified.
	ActionModified Action = 3
	// LDMACTION_FAILED:  Object failed.
	ActionFailed Action = 4
)

func (o Action) String() string {
	switch o {
	case ActionUnknown:
		return "ActionUnknown"
	case ActionCreated:
		return "ActionCreated"
	case ActionDeleted:
		return "ActionDeleted"
	case ActionModified:
		return "ActionModified"
	case ActionFailed:
		return "ActionFailed"
	}
	return "Invalid"
}

// NotifyInfoType type represents DMNOTIFY_INFO_TYPE RPC enumeration.
//
// The DMNOTIFY_INFO_TYPE enumeration defines the type of object described by an ObjectsChanged
// call.
type NotifyInfoType uint16

var (
	// DMNOTIFY_UNKNOWN_INFO:  Object is of an unknown type.
	NotifyInfoTypeUnknownInfo NotifyInfoType = 0
	// DMNOTIFY_DISK_INFO:  Object is a disk.
	NotifyInfoTypeDiskInfo NotifyInfoType = 1
	// DMNOTIFY_VOLUME_INFO:  Object is a volume.
	NotifyInfoTypeVolumeInfo NotifyInfoType = 2
	// DMNOTIFY_REGION_INFO:  Object is a region.
	NotifyInfoTypeRegionInfo NotifyInfoType = 3
	// DMNOTIFY_TASK_INFO:  Object is a task.
	NotifyInfoTypeTaskInfo NotifyInfoType = 4
	// DMNOTIFY_DL_INFO:  Object is a drive letter.
	NotifyInfoTypeDLInfo NotifyInfoType = 5
	// DMNOTIFY_FS_INFO:  Object is a file system.
	NotifyInfoTypeFSInfo NotifyInfoType = 6
	// DMNOTIFY_SYSTEM_INFO:  Object is the Disk Management system.
	NotifyInfoTypeSystemInfo NotifyInfoType = 7
)

func (o NotifyInfoType) String() string {
	switch o {
	case NotifyInfoTypeUnknownInfo:
		return "NotifyInfoTypeUnknownInfo"
	case NotifyInfoTypeDiskInfo:
		return "NotifyInfoTypeDiskInfo"
	case NotifyInfoTypeVolumeInfo:
		return "NotifyInfoTypeVolumeInfo"
	case NotifyInfoTypeRegionInfo:
		return "NotifyInfoTypeRegionInfo"
	case NotifyInfoTypeTaskInfo:
		return "NotifyInfoTypeTaskInfo"
	case NotifyInfoTypeDLInfo:
		return "NotifyInfoTypeDLInfo"
	case NotifyInfoTypeFSInfo:
		return "NotifyInfoTypeFSInfo"
	case NotifyInfoTypeSystemInfo:
		return "NotifyInfoTypeSystemInfo"
	}
	return "Invalid"
}

// ProgressType type represents DMPROGRESS_TYPE RPC enumeration.
type ProgressType uint16

var (
	ProgressTypeUnknown  ProgressType = 0
	ProgressTypeFormat   ProgressType = 1
	ProgressTypeSynching ProgressType = 2
)

func (o ProgressType) String() string {
	switch o {
	case ProgressTypeUnknown:
		return "ProgressTypeUnknown"
	case ProgressTypeFormat:
		return "ProgressTypeFormat"
	case ProgressTypeSynching:
		return "ProgressTypeSynching"
	}
	return "Invalid"
}

// VolumeSpec structure represents VOLUME_SPEC RPC structure.
//
// The VOLUME_SPEC structure specifies a new volume to create. VOLUME_SPEC is a typedef
// of this structure.
type VolumeSpec struct {
	// type:  Specifies the volume type.
	Type VolumeType `idl:"name:type" json:"type"`
	// layout:  Specifies the volume layout.
	Layout VolumeLayout `idl:"name:layout" json:"layout"`
	// partitionType:  Specifies the type of the underlying region, if this volume will
	// be a partition.
	PartitionType RegionType `idl:"name:partitionType" json:"partition_type"`
	// length:  Specifies the length of the volume in bytes. The volume length MUST always
	// be a multiple of the disk sector size.
	Length int64 `idl:"name:length" json:"length"`
	// lastKnownState:  Specifies the volume's last known modification sequence number.
	LastKnownState int64 `idl:"name:lastKnownState" json:"last_known_state"`
}

func (o *VolumeSpec) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *VolumeSpec) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Type)); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Layout)); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.PartitionType)); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if err := w.WriteData(o.LastKnownState); err != nil {
		return err
	}
	return nil
}
func (o *VolumeSpec) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Type)); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Layout)); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.PartitionType)); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastKnownState); err != nil {
		return err
	}
	return nil
}

// VolumeInfo structure represents VOLUME_INFO RPC structure.
//
// The VOLUME_INFO structure provides information about a volume.
type VolumeInfo struct {
	// id:  Specifies the object identifier (OID) for the volume.
	ID int64 `idl:"name:id" json:"id"`
	// type:   Specifies the volume type.
	Type VolumeType `idl:"name:type" json:"type"`
	// layout:  Specifies the volume layout.
	Layout VolumeLayout `idl:"name:layout" json:"layout"`
	// length:  Specifies the length of the volume in bytes.
	Length int64 `idl:"name:length" json:"length"`
	// fsId:  Specifies the object identifier for the volume's file system, which defaults
	// to 0 if no file system is present on the volume.
	FSID int64 `idl:"name:fsId" json:"fs_id"`
	// memberCount:  Specifies the number of regions that compose the volume.
	MemberCount uint32 `idl:"name:memberCount" json:"member_count"`
	// status:  Specifies the volume status.
	Status VolumeStatus `idl:"name:status" json:"status"`
	// lastKnownState:  Specifies the volume's modification sequence number.
	LastKnownState int64 `idl:"name:lastKnownState" json:"last_known_state"`
	// taskId:  Specifies the task identifier of the associated user request. If no request
	// is made, the value is 0. For more information, see section 2.2.17.
	TaskID int64 `idl:"name:taskId" json:"task_id"`
	// vflags:  Specifies the bitmap of volume flags. The value of this field is generated
	// by combining zero or more of the following applicable flags with a logical OR operation.
	//
	// This field MUST be one of the following values.
	//
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	|                                        |                                                                                  |
	//	|                 VALUE                  |                                     MEANING                                      |
	//	|                                        |                                                                                  |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| VOLUME_FORMAT_IN_PROGRESS 0x00000001   | Volume is currently being formatted.                                             |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| VOLUME_HAS_PAGEFILE 0x00000004         | Volume contains the paging file.                                                 |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| VOLUME_IS_BOOT_VOLUME 0x00000100       | Volume contains the boot partition.                                              |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| VOLUME_IS_RESTARTABLE 0x00000400       | The RestartVolume method can be successfully called on this volume.              |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| VOLUME_IS_SYSTEM_VOLUME 0x00000800     | Volume contains the system directory.                                            |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| VOLUME_HAS_RETAIN_PARTITION 0x00001000 | Volume has an underlying partition.                                              |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| VOLUME_HAD_BOOT_INI 0x00002000         | Volume contained the Boot.ini file used when the operating system was last       |
	//	|                                        | started.                                                                         |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| VOLUME_CORRUPT 0x00004000              | Volume is corrupt.                                                               |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| VOLUME_HAS_CRASHDUMP 0x00008000        | Volume contains a crash dump file.                                               |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| VOLUME_IS_CURR_BOOT_VOLUME 0x00010000  | Volume is the current boot volume.                                               |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| VOLUME_HAS_HIBERNATION 0x00020000      | Volume contains a hibernation image.                                             |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	VolumeFlags uint32 `idl:"name:vflags" json:"volume_flags"`
}

func (o *VolumeInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *VolumeInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.ID); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Type)); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Layout)); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if err := w.WriteData(o.FSID); err != nil {
		return err
	}
	if err := w.WriteData(o.MemberCount); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Status)); err != nil {
		return err
	}
	if err := w.WriteData(o.LastKnownState); err != nil {
		return err
	}
	if err := w.WriteData(o.TaskID); err != nil {
		return err
	}
	if err := w.WriteData(o.VolumeFlags); err != nil {
		return err
	}
	return nil
}
func (o *VolumeInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.ID); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Type)); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Layout)); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	if err := w.ReadData(&o.FSID); err != nil {
		return err
	}
	if err := w.ReadData(&o.MemberCount); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Status)); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastKnownState); err != nil {
		return err
	}
	if err := w.ReadData(&o.TaskID); err != nil {
		return err
	}
	if err := w.ReadData(&o.VolumeFlags); err != nil {
		return err
	}
	return nil
}

// DiskSpec structure represents DISK_SPEC RPC structure.
//
// The DISK_SPEC structure specifies a disk for a volume modification or a creation
// request.
type DiskSpec struct {
	// diskId:  Specifies the OID for the disk.
	DiskID int64 `idl:"name:diskId" json:"disk_id"`
	// length:  Specifies the byte length to use.
	Length int64 `idl:"name:length" json:"length"`
	// needContiguous:  Boolean value that specifies if contiguous space is needed on the
	// disk.
	//
	//	+---------+---------------------------------------------+
	//	|         |                                             |
	//	|  VALUE  |                   MEANING                   |
	//	|         |                                             |
	//	+---------+---------------------------------------------+
	//	+---------+---------------------------------------------+
	//	| FALSE 0 | Contiguous space is not needed on the disk. |
	//	+---------+---------------------------------------------+
	//	| TRUE 1  | Contiguous space is needed on the disk.     |
	//	+---------+---------------------------------------------+
	NeedContiguous bool `idl:"name:needContiguous" json:"need_contiguous"`
	// lastKnownState:  Last known modification sequence number of the disk.
	LastKnownState int64 `idl:"name:lastKnownState" json:"last_known_state"`
}

func (o *DiskSpec) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DiskSpec) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.DiskID); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if err := w.WriteData(o.NeedContiguous); err != nil {
		return err
	}
	if err := w.WriteData(o.LastKnownState); err != nil {
		return err
	}
	return nil
}
func (o *DiskSpec) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.DiskID); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	if err := w.ReadData(&o.NeedContiguous); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastKnownState); err != nil {
		return err
	}
	return nil
}

// DiskInfo structure represents DISK_INFO RPC structure.
//
// The DISK_INFO structure provides information about a disk.
type DiskInfo struct {
	// id:  Specifies the OID of the disk.
	ID int64 `idl:"name:id" json:"id"`
	// length:   Size of the disk, in bytes.
	Length int64 `idl:"name:length" json:"length"`
	// freeBytes:  Number of unallocated bytes on the disk.
	FreeBytes int64 `idl:"name:freeBytes" json:"free_bytes"`
	// bytesPerTrack:  Size of a disk track, in bytes.
	BytesPerTrack uint32 `idl:"name:bytesPerTrack" json:"bytes_per_track"`
	// bytesPerCylinder:  Size of a disk cylinder, in bytes.
	BytesPerCylinder uint32 `idl:"name:bytesPerCylinder" json:"bytes_per_cylinder"`
	// bytesPerSector:  Size of a disk sector, in bytes.
	BytesPerSector uint32 `idl:"name:bytesPerSector" json:"bytes_per_sector"`
	// regionCount:  Total number of regions on the disk.
	RegionCount uint32 `idl:"name:regionCount" json:"region_count"`
	// dflags:  Disk type of the disk.
	//
	//	+--------------------------+----------------------------------------------+
	//	|                          |                                              |
	//	|          VALUE           |                   MEANING                    |
	//	|                          |                                              |
	//	+--------------------------+----------------------------------------------+
	//	+--------------------------+----------------------------------------------+
	//	| DISK_AUDIO_CD 0x00000001 | Disk is an audio CD.                         |
	//	+--------------------------+----------------------------------------------+
	//	| DISK_NEC98 0x00000002    | This value is obsolete and MUST NOT be used. |
	//	+--------------------------+----------------------------------------------+
	DiskFlags uint32 `idl:"name:dflags" json:"disk_flags"`
	// deviceType:  Device type of the disk. This field contains one of the following values.
	//
	//	+---------------------------------+-------------------------------+
	//	|                                 |                               |
	//	|              VALUE              |            MEANING            |
	//	|                                 |                               |
	//	+---------------------------------+-------------------------------+
	//	+---------------------------------+-------------------------------+
	//	| DEVICETYPE_UNKNOWN 0x00000000   | Device is of an unknown type. |
	//	+---------------------------------+-------------------------------+
	//	| DEVICETYPE_VM 0x00000001        | Device is a dynamic disk.     |
	//	+---------------------------------+-------------------------------+
	//	| DEVICETYPE_REMOVABLE 0x00000002 | Device uses removable media.  |
	//	+---------------------------------+-------------------------------+
	//	| DEVICETYPE_CDROM 0x00000003     | Device is a CD-ROM.           |
	//	+---------------------------------+-------------------------------+
	//	| DEVICETYPE_FDISK 0x00000004     | Device is a basic disk.       |
	//	+---------------------------------+-------------------------------+
	//	| DEVICETYPE_DVD 0x00000005       | Device is a DVD.              |
	//	+---------------------------------+-------------------------------+
	DeviceType uint32 `idl:"name:deviceType" json:"device_type"`
	// deviceState:   Device state of the disk. The value of this field is generated by
	// combining zero or more of the applicable flags with a logical OR operation. Valid
	// combinations are device-type dependent.
	//
	//	+--------------------------------------+-----------------------------------------------------------------+
	//	|                                      |                                                                 |
	//	|                VALUE                 |                             MEANING                             |
	//	|                                      |                                                                 |
	//	+--------------------------------------+-----------------------------------------------------------------+
	//	+--------------------------------------+-----------------------------------------------------------------+
	//	| DEVICESTATE_UNKNOWN 0x00000000       | Disk is in an unknown state.                                    |
	//	+--------------------------------------+-----------------------------------------------------------------+
	//	| DEVICESTATE_HEALTHY 0x00000001       | Disk is fully functional.                                       |
	//	+--------------------------------------+-----------------------------------------------------------------+
	//	| DEVICESTATE_NO_MEDIA 0x00000002      | Disk has no media.                                              |
	//	+--------------------------------------+-----------------------------------------------------------------+
	//	| DEVICESTATE_NOSIG 0x00000004         | Disk has an invalid signature.                                  |
	//	+--------------------------------------+-----------------------------------------------------------------+
	//	| DEVICESTATE_BAD 0x00000008           | Disk was deleted or experienced an install or hardware problem. |
	//	+--------------------------------------+-----------------------------------------------------------------+
	//	| DEVICESTATE_NOT_READY 0x00000010     | Disk is not ready yet.                                          |
	//	+--------------------------------------+-----------------------------------------------------------------+
	//	| DEVICESTATE_MISSING 0x00000020       | Disk is no longer available.                                    |
	//	+--------------------------------------+-----------------------------------------------------------------+
	//	| DEVICESTATE_OFFLINE 0x00000040       | Disk is offline.                                                |
	//	+--------------------------------------+-----------------------------------------------------------------+
	//	| DEVICESTATE_FAILING 0x00000080       | Disk experienced a physical I/O error.                          |
	//	+--------------------------------------+-----------------------------------------------------------------+
	//	| DEVICESTATE_IMPORT_FAILED 0x00000100 | Disk belongs to a group whose import failed.                    |
	//	+--------------------------------------+-----------------------------------------------------------------+
	//	| DEVICESTATE_UNCLAIMED 0x00000200     | Disk belongs to a foreign disk group.                           |
	//	+--------------------------------------+-----------------------------------------------------------------+
	DeviceState uint32 `idl:"name:deviceState" json:"device_state"`
	// busType:  Type of bus on which the disk resides. This field contains one of the following
	// values.
	//
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	|                            |                                                                                  |
	//	|           VALUE            |                                     MEANING                                      |
	//	|                            |                                                                                  |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| BUSTYPE_UNKNOWN 0x00000000 | Bus type is unknown.                                                             |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| BUSTYPE_IDE 0x00000001     | Disk resides on an Integrated Drive Electronics (IDE) bus.                       |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| BUSTYPE_SCSI 0x00000002    | Disk resides on a SCSI bus.                                                      |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| BUSTYPE_FIBRE 0x00000003   | Disk resides on a fiber channel bus.                                             |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| BUSTYPE_USB 0x00000004     | Disk resides on a universal serial bus (USB).                                    |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| BUSTYPE_SSA 0x00000005     | Disk resides on a serial storage architecture (SSA) Bus.                         |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| BUSTYPE_1394 0x00000006    | Disk resides on an Institute of Electronics and Electrical Engineers (IEEE) 1394 |
	//	|                            | bus.                                                                             |
	//	+----------------------------+----------------------------------------------------------------------------------+
	BusType uint32 `idl:"name:busType" json:"bus_type"`
	// attributes:  Bitmap of disk attributes. The value of this field is generated by combining
	// zero or more of the applicable flags defined in the following table with a logical
	// OR operation.
	//
	//	+------------------------------+-------------------------+
	//	|                              |                         |
	//	|            VALUE             |         MEANING         |
	//	|                              |                         |
	//	+------------------------------+-------------------------+
	//	+------------------------------+-------------------------+
	//	| DEVICEATTR_NONE 0x00000000   | Disk has no attributes. |
	//	+------------------------------+-------------------------+
	//	| DEVICEATTR_RDONLY 0x00000001 | Disk is read-only.      |
	//	+------------------------------+-------------------------+
	//	| DEVICEATTR_NTMS 0x00000002   | This value is obsolete. |
	//	+------------------------------+-------------------------+
	Attributes uint32 `idl:"name:attributes" json:"attributes"`
	// isUpgradeable:  Boolean value that indicates whether the disk can be converted to
	// a dynamic disk. Will be true if the disk is basic, healthy, and has 512 byte sectors.
	//
	//	+---------+----------------------------------------------------------+
	//	|         |                                                          |
	//	|  VALUE  |                         MEANING                          |
	//	|         |                                                          |
	//	+---------+----------------------------------------------------------+
	//	+---------+----------------------------------------------------------+
	//	| FALSE 0 | Disk cannot be converted to a dynamic disk.              |
	//	+---------+----------------------------------------------------------+
	//	| TRUE 1  | Disk can be encapsulated or converted to a dynamic disk. |
	//	+---------+----------------------------------------------------------+
	IsUpgradeable bool `idl:"name:isUpgradeable" json:"is_upgradeable"`
	// portNumber:  SCSI port number of the disk, if the bus reports this information.
	PortNumber int32 `idl:"name:portNumber" json:"port_number"`
	// targetNumber:  SCSI target identifier of the disk, if the bus reports this information.
	TargetNumber int32 `idl:"name:targetNumber" json:"target_number"`
	// lunNumber:  SCSI logical unit number (LUN) of the disk, if the bus reports this information.
	LUNNumber int32 `idl:"name:lunNumber" json:"lun_number"`
	// lastKnownState:  Modification sequence number of the disk.
	LastKnownState int64 `idl:"name:lastKnownState" json:"last_known_state"`
	// taskId:  The task identifier of the associated user request. If no request is made,
	// the value is 0. For more information about this task identifier, see section 2.2.17.
	TaskID int64 `idl:"name:taskId" json:"task_id"`
	// cchName:  Length of the hard disk physical name, in Unicode characters, including
	// the terminating null character.
	NameLength int32 `idl:"name:cchName" json:"name_length"`
	// cchVendor:  Length of the disk's vendor name, in Unicode characters, including the
	// terminating null character.
	VendorLength int32 `idl:"name:cchVendor" json:"vendor_length"`
	// cchDgid:  Length of the disk's group identification handle, in ASCII characters,
	// including the terminating null character.
	DiskGroupIDLength int32 `idl:"name:cchDgid" json:"disk_group_id_length"`
	// cchAdapterName:  Length of the disk's adapter name, in Unicode characters, including
	// the terminating null character.
	AdapterNameLength int32 `idl:"name:cchAdapterName" json:"adapter_name_length"`
	// cchDgName:  Length of the disk's group name, in Unicode characters, including the
	// terminating null character.
	DiskGroupNameLength int32 `idl:"name:cchDgName" json:"disk_group_name_length"`
	// name:  Null-terminated physical device name of the hard disk, in the format '\device\Harddisk1'.
	// This is Unicode.
	Name string `idl:"name:name;size_is:(cchName)" json:"name"`
	// vendor:  Null-terminated name of the hard disk vendor. This is the disk vendor's
	// disk model name. This is Unicode.
	Vendor string `idl:"name:vendor;size_is:(cchVendor)" json:"vendor"`
	// dgid:  Specifies the object identifier of the disk's disk group. This is ASCII.
	DiskGroupID []byte `idl:"name:dgid;size_is:(cchDgid)" json:"disk_group_id"`
	// adapterName:  Null-terminated name of the disk adapter as returned by the disk adapter
	// firmware; for example, 'Adaptec AHA-2940U2W - Ultra2 SCSI'. This is Unicode.
	AdapterName string `idl:"name:adapterName;size_is:(cchAdapterName)" json:"adapter_name"`
	// dgName:  Null-terminated name for the disk's disk group, if the disk is dynamic.
	// Only dynamic disks have an associated disk group. Basic disks do not. This is Unicode.
	DiskGroupName string `idl:"name:dgName;size_is:(cchDgName)" json:"disk_group_name"`
}

func (o *DiskInfo) xxx_PreparePayload(ctx context.Context) error {
	if o.Name != "" && o.NameLength == 0 {
		o.NameLength = int32(len(o.Name))
	}
	if o.Vendor != "" && o.VendorLength == 0 {
		o.VendorLength = int32(len(o.Vendor))
	}
	if o.DiskGroupID != nil && o.DiskGroupIDLength == 0 {
		o.DiskGroupIDLength = int32(len(o.DiskGroupID))
	}
	if o.AdapterName != "" && o.AdapterNameLength == 0 {
		o.AdapterNameLength = int32(len(o.AdapterName))
	}
	if o.DiskGroupName != "" && o.DiskGroupNameLength == 0 {
		o.DiskGroupNameLength = int32(len(o.DiskGroupName))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DiskInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.ID); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if err := w.WriteData(o.FreeBytes); err != nil {
		return err
	}
	if err := w.WriteData(o.BytesPerTrack); err != nil {
		return err
	}
	if err := w.WriteData(o.BytesPerCylinder); err != nil {
		return err
	}
	if err := w.WriteData(o.BytesPerSector); err != nil {
		return err
	}
	if err := w.WriteData(o.RegionCount); err != nil {
		return err
	}
	if err := w.WriteData(o.DiskFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.DeviceType); err != nil {
		return err
	}
	if err := w.WriteData(o.DeviceState); err != nil {
		return err
	}
	if err := w.WriteData(o.BusType); err != nil {
		return err
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	if err := w.WriteData(o.IsUpgradeable); err != nil {
		return err
	}
	if err := w.WriteData(o.PortNumber); err != nil {
		return err
	}
	if err := w.WriteData(o.TargetNumber); err != nil {
		return err
	}
	if err := w.WriteData(o.LUNNumber); err != nil {
		return err
	}
	if err := w.WriteData(o.LastKnownState); err != nil {
		return err
	}
	if err := w.WriteData(o.TaskID); err != nil {
		return err
	}
	if err := w.WriteData(o.NameLength); err != nil {
		return err
	}
	if err := w.WriteData(o.VendorLength); err != nil {
		return err
	}
	if err := w.WriteData(o.DiskGroupIDLength); err != nil {
		return err
	}
	if err := w.WriteData(o.AdapterNameLength); err != nil {
		return err
	}
	if err := w.WriteData(o.DiskGroupNameLength); err != nil {
		return err
	}
	if o.Name != "" || o.NameLength > 0 {
		_ptr_name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.NameLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			_Name_buf := utf16.Encode([]rune(o.Name))
			if uint64(len(_Name_buf)) > sizeInfo[0] {
				_Name_buf = _Name_buf[:sizeInfo[0]]
			}
			for i1 := range _Name_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_Name_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_Name_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_name); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Vendor != "" || o.VendorLength > 0 {
		_ptr_vendor := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.VendorLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			_Vendor_buf := utf16.Encode([]rune(o.Vendor))
			if uint64(len(_Vendor_buf)) > sizeInfo[0] {
				_Vendor_buf = _Vendor_buf[:sizeInfo[0]]
			}
			for i1 := range _Vendor_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_Vendor_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_Vendor_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Vendor, _ptr_vendor); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DiskGroupID != nil || o.DiskGroupIDLength > 0 {
		_ptr_dgid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.DiskGroupIDLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.DiskGroupID {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.DiskGroupID[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.DiskGroupID); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.DiskGroupID, _ptr_dgid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.AdapterName != "" || o.AdapterNameLength > 0 {
		_ptr_adapterName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.AdapterNameLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			_AdapterName_buf := utf16.Encode([]rune(o.AdapterName))
			if uint64(len(_AdapterName_buf)) > sizeInfo[0] {
				_AdapterName_buf = _AdapterName_buf[:sizeInfo[0]]
			}
			for i1 := range _AdapterName_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_AdapterName_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_AdapterName_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.AdapterName, _ptr_adapterName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DiskGroupName != "" || o.DiskGroupNameLength > 0 {
		_ptr_dgName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.DiskGroupNameLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			_DiskGroupName_buf := utf16.Encode([]rune(o.DiskGroupName))
			if uint64(len(_DiskGroupName_buf)) > sizeInfo[0] {
				_DiskGroupName_buf = _DiskGroupName_buf[:sizeInfo[0]]
			}
			for i1 := range _DiskGroupName_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_DiskGroupName_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_DiskGroupName_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.DiskGroupName, _ptr_dgName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *DiskInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.ID); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	if err := w.ReadData(&o.FreeBytes); err != nil {
		return err
	}
	if err := w.ReadData(&o.BytesPerTrack); err != nil {
		return err
	}
	if err := w.ReadData(&o.BytesPerCylinder); err != nil {
		return err
	}
	if err := w.ReadData(&o.BytesPerSector); err != nil {
		return err
	}
	if err := w.ReadData(&o.RegionCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.DiskFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.DeviceType); err != nil {
		return err
	}
	if err := w.ReadData(&o.DeviceState); err != nil {
		return err
	}
	if err := w.ReadData(&o.BusType); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsUpgradeable); err != nil {
		return err
	}
	if err := w.ReadData(&o.PortNumber); err != nil {
		return err
	}
	if err := w.ReadData(&o.TargetNumber); err != nil {
		return err
	}
	if err := w.ReadData(&o.LUNNumber); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastKnownState); err != nil {
		return err
	}
	if err := w.ReadData(&o.TaskID); err != nil {
		return err
	}
	if err := w.ReadData(&o.NameLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.VendorLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.DiskGroupIDLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.AdapterNameLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.DiskGroupNameLength); err != nil {
		return err
	}
	_ptr_name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
		var _Name_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Name_buf", sizeInfo[0])
		}
		_Name_buf = make([]uint16, sizeInfo[0])
		for i1 := range _Name_buf {
			i1 := i1
			if err := w.ReadData(&_Name_buf[i1]); err != nil {
				return err
			}
		}
		o.Name = strings.TrimRight(string(utf16.Decode(_Name_buf)), ndr.ZeroString)
		return nil
	})
	_s_name := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_name, _ptr_name); err != nil {
		return err
	}
	_ptr_vendor := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.VendorLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.VendorLength)
		}
		var _Vendor_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Vendor_buf", sizeInfo[0])
		}
		_Vendor_buf = make([]uint16, sizeInfo[0])
		for i1 := range _Vendor_buf {
			i1 := i1
			if err := w.ReadData(&_Vendor_buf[i1]); err != nil {
				return err
			}
		}
		o.Vendor = strings.TrimRight(string(utf16.Decode(_Vendor_buf)), ndr.ZeroString)
		return nil
	})
	_s_vendor := func(ptr interface{}) { o.Vendor = *ptr.(*string) }
	if err := w.ReadPointer(&o.Vendor, _s_vendor, _ptr_vendor); err != nil {
		return err
	}
	_ptr_dgid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.DiskGroupIDLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.DiskGroupIDLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.DiskGroupID", sizeInfo[0])
		}
		o.DiskGroupID = make([]byte, sizeInfo[0])
		for i1 := range o.DiskGroupID {
			i1 := i1
			if err := w.ReadData(&o.DiskGroupID[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_dgid := func(ptr interface{}) { o.DiskGroupID = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.DiskGroupID, _s_dgid, _ptr_dgid); err != nil {
		return err
	}
	_ptr_adapterName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.AdapterNameLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.AdapterNameLength)
		}
		var _AdapterName_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _AdapterName_buf", sizeInfo[0])
		}
		_AdapterName_buf = make([]uint16, sizeInfo[0])
		for i1 := range _AdapterName_buf {
			i1 := i1
			if err := w.ReadData(&_AdapterName_buf[i1]); err != nil {
				return err
			}
		}
		o.AdapterName = strings.TrimRight(string(utf16.Decode(_AdapterName_buf)), ndr.ZeroString)
		return nil
	})
	_s_adapterName := func(ptr interface{}) { o.AdapterName = *ptr.(*string) }
	if err := w.ReadPointer(&o.AdapterName, _s_adapterName, _ptr_adapterName); err != nil {
		return err
	}
	_ptr_dgName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.DiskGroupNameLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.DiskGroupNameLength)
		}
		var _DiskGroupName_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _DiskGroupName_buf", sizeInfo[0])
		}
		_DiskGroupName_buf = make([]uint16, sizeInfo[0])
		for i1 := range _DiskGroupName_buf {
			i1 := i1
			if err := w.ReadData(&_DiskGroupName_buf[i1]); err != nil {
				return err
			}
		}
		o.DiskGroupName = strings.TrimRight(string(utf16.Decode(_DiskGroupName_buf)), ndr.ZeroString)
		return nil
	})
	_s_dgName := func(ptr interface{}) { o.DiskGroupName = *ptr.(*string) }
	if err := w.ReadPointer(&o.DiskGroupName, _s_dgName, _ptr_dgName); err != nil {
		return err
	}
	return nil
}

// RegionSpec structure represents REGION_SPEC RPC structure.
//
// The REGION_SPEC structure specifies a region for partition creation and deletion.
type RegionSpec struct {
	// regionId:  Specifies the OID for the region.
	RegionID int64 `idl:"name:regionId" json:"region_id"`
	// regionType:  Specifies the region type.
	RegionType RegionType `idl:"name:regionType" json:"region_type"`
	// diskId:  Specifies the OID for the disk on which the region resides.
	DiskID int64 `idl:"name:diskId" json:"disk_id"`
	// start:  Specifies the byte offset of the region on disk.
	Start int64 `idl:"name:start" json:"start"`
	// length:  Specifies the length of the region in bytes.
	Length int64 `idl:"name:length" json:"length"`
	// lastKnownState:  Specifies the region's last known modification sequence number.
	LastKnownState int64 `idl:"name:lastKnownState" json:"last_known_state"`
}

func (o *RegionSpec) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RegionSpec) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.RegionID); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.RegionType)); err != nil {
		return err
	}
	if err := w.WriteData(o.DiskID); err != nil {
		return err
	}
	if err := w.WriteData(o.Start); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if err := w.WriteData(o.LastKnownState); err != nil {
		return err
	}
	return nil
}
func (o *RegionSpec) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.RegionID); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.RegionType)); err != nil {
		return err
	}
	if err := w.ReadData(&o.DiskID); err != nil {
		return err
	}
	if err := w.ReadData(&o.Start); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastKnownState); err != nil {
		return err
	}
	return nil
}

// RegionInfo structure represents REGION_INFO RPC structure.
//
// The REGION_INFO structure provides information about a region.
type RegionInfo struct {
	// id:  Specifies the region's OID.
	ID int64 `idl:"name:id" json:"id"`
	// diskId:  Specifies the OID of the disk on which the region resides.
	DiskID int64 `idl:"name:diskId" json:"disk_id"`
	// volId:  Specifies the OID of the volume on the region, if any. The value of this
	// field is nonzero if it is valid.
	VolID int64 `idl:"name:volId" json:"vol_id"`
	// fsId:  Specifies the OID of the file system on the region, if any. The value of this
	// field is nonzero if it is valid.
	FSID int64 `idl:"name:fsId" json:"fs_id"`
	// start:  Byte offset of the region on the disk.
	Start int64 `idl:"name:start" json:"start"`
	// length:  Length of the region in bytes.
	Length int64 `idl:"name:length" json:"length"`
	// regionType:  Value from the REGIONTYPE enumeration that indicates the region type.
	RegionType RegionType `idl:"name:regionType" json:"region_type"`
	// partitionType:  Type of the partition on the region. This field contains one of the
	// following values.
	//
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	|                             |                                                                                  |
	//	|            VALUE            |                                     MEANING                                      |
	//	|                             |                                                                                  |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| PARTITION_ENTRY_UNUSED 0x00 | An unused entry partition.                                                       |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| PARTITION_EXTENDED 0x05     | An extended partition.                                                           |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| PARTITION_FAT_12 0x01       | A FAT12 file system partition.                                                   |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| PARTITION_FAT_16 0x04       | A FAT16 file system partition.                                                   |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| PARTITION_FAT32 0x0B        | A FAT32 file system partition.                                                   |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| PARTITION_IFS 0x07          | An installable file system (IFS) partition.                                      |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| PARTITION_LDM 0x42          | An LDM partition.                                                                |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| PARTITION_NTFT 0x80         | A Windows NT operating system fault-tolerant (FT) partition.                     |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| VALID_NTFT 0xC0             | A valid Windows NT FT partition. The high bit of a partition type code indicates |
	//	|                             | that a partition is part of an NT FT mirror or striped array.                    |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	PartitionType uint32 `idl:"name:partitionType" json:"partition_type"`
	// isActive:  Boolean value that indicates whether the region is an active partition.
	//
	//	+---------+----------------------------------+
	//	|         |                                  |
	//	|  VALUE  |             MEANING              |
	//	|         |                                  |
	//	+---------+----------------------------------+
	//	+---------+----------------------------------+
	//	| FALSE 0 | Region is an inactive partition. |
	//	+---------+----------------------------------+
	//	| TRUE 1  | Region is an active partition.   |
	//	+---------+----------------------------------+
	IsActive bool `idl:"name:isActive" json:"is_active"`
	// status:   Value from the REGIONSTATUS enumeration that indicates the region's status.
	Status RegionStatus `idl:"name:status" json:"status"`
	// lastKnownState:  Modification sequence number of the region.
	LastKnownState int64 `idl:"name:lastKnownState" json:"last_known_state"`
	// taskId:   This LdmObjectId is the task identifier of the associated user request.
	// If no request is made, the value is 0. For more information about this task identifier,
	// see section 2.2.17.
	TaskID int64 `idl:"name:taskId" json:"task_id"`
	// rflags:  Bitmap of region flags. The value of this field is generated by combining
	// zero or more of the applicable flags with a logical OR operation.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                                       |                                                                                  |
	//	|                 VALUE                 |                                     MEANING                                      |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| REGION_FORMAT_IN_PROGRESS 0x00000001  | Region is currently being formatted.                                             |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| REGION_IS_SYSTEM_PARTITION 0x00000002 | Region contains the system directory. The system directory has the operating     |
	//	|                                       | system installed to it. This is not necessarily the "active" partition that      |
	//	|                                       | contains the boot loader file.                                                   |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| REGION_HAS_PAGEFILE 0x00000004        | Region contains the paging file.                                                 |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| REGION_HAD_BOOT_INI 0x00000040        | Boot file was located in this region when the operating system was last started. |
	//	|                                       | This is the "active" partition that contains the boot loader file.               |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	RegionFlags uint32 `idl:"name:rflags" json:"region_flags"`
	// currentPartitionNumber:   Number of the partition on the region, if any.
	CurrentPartitionNumber uint32 `idl:"name:currentPartitionNumber" json:"current_partition_number"`
}

func (o *RegionInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RegionInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.ID); err != nil {
		return err
	}
	if err := w.WriteData(o.DiskID); err != nil {
		return err
	}
	if err := w.WriteData(o.VolID); err != nil {
		return err
	}
	if err := w.WriteData(o.FSID); err != nil {
		return err
	}
	if err := w.WriteData(o.Start); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.RegionType)); err != nil {
		return err
	}
	if err := w.WriteData(o.PartitionType); err != nil {
		return err
	}
	if err := w.WriteData(o.IsActive); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Status)); err != nil {
		return err
	}
	if err := w.WriteData(o.LastKnownState); err != nil {
		return err
	}
	if err := w.WriteData(o.TaskID); err != nil {
		return err
	}
	if err := w.WriteData(o.RegionFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.CurrentPartitionNumber); err != nil {
		return err
	}
	return nil
}
func (o *RegionInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.ID); err != nil {
		return err
	}
	if err := w.ReadData(&o.DiskID); err != nil {
		return err
	}
	if err := w.ReadData(&o.VolID); err != nil {
		return err
	}
	if err := w.ReadData(&o.FSID); err != nil {
		return err
	}
	if err := w.ReadData(&o.Start); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.RegionType)); err != nil {
		return err
	}
	if err := w.ReadData(&o.PartitionType); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsActive); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Status)); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastKnownState); err != nil {
		return err
	}
	if err := w.ReadData(&o.TaskID); err != nil {
		return err
	}
	if err := w.ReadData(&o.RegionFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.CurrentPartitionNumber); err != nil {
		return err
	}
	return nil
}

// DriveLetterInfo structure represents DRIVE_LETTER_INFO RPC structure.
//
// The DRIVE_LETTER_INFO structure provides information about a drive letter. It is
// used for drive letter assignment and free requests, for notification of drive letter
// changes, and for enumeration.
type DriveLetterInfo struct {
	// letter:  Drive letter as a single case-insensitive alphabetical Unicode character.
	Letter uint16 `idl:"name:letter" json:"letter"`
	// storageId:  Specifies the OID of the volume, partition, or logical drive to which
	// the drive letter is assigned, if any.
	StorageID int64 `idl:"name:storageId" json:"storage_id"`
	// isUsed:   Boolean value that specifies if the drive letter is in use.
	//
	//	+---------+-------------------------+
	//	|         |                         |
	//	|  VALUE  |         MEANING         |
	//	|         |                         |
	//	+---------+-------------------------+
	//	+---------+-------------------------+
	//	| FALSE 0 | Drive letter is free.   |
	//	+---------+-------------------------+
	//	| TRUE 1  | Drive letter is in use. |
	//	+---------+-------------------------+
	IsUsed bool `idl:"name:isUsed" json:"is_used"`
	// lastKnownState:  Modification sequence number of the drive letter.
	LastKnownState int64 `idl:"name:lastKnownState" json:"last_known_state"`
	// taskId:  Specifies the task identifier of the associated user request. If no request
	// is made, the value is 0. For more information about this task identifier, see section
	// 2.2.17.
	TaskID int64 `idl:"name:taskId" json:"task_id"`
	// dlflags:  Bitmap of drive letter flags. The value of this field is generated by combining
	// zero or more of the applicable flags defined as follows with a logical OR operation.
	//
	//	+-------------------------------+-----------------------------------------------+
	//	|                               |                                               |
	//	|             VALUE             |                    MEANING                    |
	//	|                               |                                               |
	//	+-------------------------------+-----------------------------------------------+
	//	+-------------------------------+-----------------------------------------------+
	//	| DL_PENDING_REMOVAL 0x00000001 | Drive letter has a removal operation pending. |
	//	+-------------------------------+-----------------------------------------------+
	DriveLetterFlags uint32 `idl:"name:dlflags" json:"drive_letter_flags"`
}

func (o *DriveLetterInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DriveLetterInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.Letter); err != nil {
		return err
	}
	if err := w.WriteData(o.StorageID); err != nil {
		return err
	}
	if err := w.WriteData(o.IsUsed); err != nil {
		return err
	}
	if err := w.WriteData(o.LastKnownState); err != nil {
		return err
	}
	if err := w.WriteData(o.TaskID); err != nil {
		return err
	}
	if err := w.WriteData(o.DriveLetterFlags); err != nil {
		return err
	}
	return nil
}
func (o *DriveLetterInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.Letter); err != nil {
		return err
	}
	if err := w.ReadData(&o.StorageID); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsUsed); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastKnownState); err != nil {
		return err
	}
	if err := w.ReadData(&o.TaskID); err != nil {
		return err
	}
	if err := w.ReadData(&o.DriveLetterFlags); err != nil {
		return err
	}
	return nil
}

// FileSystemInfo structure represents FILE_SYSTEM_INFO RPC structure.
//
// The FILE_SYSTEM_INFO structure provides information about a file system. This structure
// is used for file system enumeration, file system operations, and notification of
// file system changes in the configuration database. For more information about the
// parameters, see [MSDN-STC].
type FileSystemInfo struct {
	// id:  Specifies the OID for the file system.
	ID int64 `idl:"name:id" json:"id"`
	// storageId:   Specifies the OID for the volume, partition, or logical drive associated
	// with the file system.
	StorageID int64 `idl:"name:storageId" json:"storage_id"`
	// totalAllocationUnits:  Total number of file allocation units in the file system.
	TotalAllocationUnits int64 `idl:"name:totalAllocationUnits" json:"total_allocation_units"`
	// availableAllocationUnits:  Number of available file allocation units in the file
	// system.
	AvailableAllocationUnits int64 `idl:"name:availableAllocationUnits" json:"available_allocation_units"`
	// allocationUnitSize:  Size of a file allocation unit in bytes.
	AllocationUnitSize uint32 `idl:"name:allocationUnitSize" json:"allocation_unit_size"`
	// fsflags:  Bitmap of file system flags. The value of this field is generated by combining
	// zero or more of the applicable flags with a logical OR operation.
	//
	//	+--------------------------------------+---------------------------------------------------------+
	//	|                                      |                                                         |
	//	|                VALUE                 |                         MEANING                         |
	//	|                                      |                                                         |
	//	+--------------------------------------+---------------------------------------------------------+
	//	+--------------------------------------+---------------------------------------------------------+
	//	| ENABLE_VOLUME_COMPRESSION 0x00000001 | File system supports NT file system (NTFS) compression. |
	//	+--------------------------------------+---------------------------------------------------------+
	Fsflags uint32 `idl:"name:fsflags" json:"fsflags"`
	// lastKnownState:  File system's last known modification sequence number.
	LastKnownState int64 `idl:"name:lastKnownState" json:"last_known_state"`
	// taskId:  Specifies the task identifier of the associated user request. If no request
	// is made, the value is 0. For more information about this task identifier, see section
	// 2.2.17.
	TaskID int64 `idl:"name:taskId" json:"task_id"`
	// fsType:  Type of the file system.
	//
	//	+---------------------------+------------------------------------------------------+
	//	|                           |                                                      |
	//	|           VALUE           |                       MEANING                        |
	//	|                           |                                                      |
	//	+---------------------------+------------------------------------------------------+
	//	+---------------------------+------------------------------------------------------+
	//	| FSTYPE_UNKNOWN 0x00000000 | File system type is unknown.                         |
	//	+---------------------------+------------------------------------------------------+
	//	| FSTYPE_NTFS 0x00000001    | File system type is NTFS.                            |
	//	+---------------------------+------------------------------------------------------+
	//	| FSTYPE_FAT 0x00000002     | File system type is file allocation table (FAT).     |
	//	+---------------------------+------------------------------------------------------+
	//	| FSTYPE_FAT32 0x00000003   | File system type is a FAT32 file system.             |
	//	+---------------------------+------------------------------------------------------+
	//	| FSTYPE_CDFS 0x00000004    | File system type is Compact Disc File System (CDFS). |
	//	+---------------------------+------------------------------------------------------+
	//	| FSTYPE_UDF 0x00000005     | File system type is Universal Disk Format (UDF).     |
	//	+---------------------------+------------------------------------------------------+
	//	| FSTYPE_OTHER 0x80000000   | File system type is not listed.                      |
	//	+---------------------------+------------------------------------------------------+
	FSType int32 `idl:"name:fsType" json:"fs_type"`
	// cchLabel:   Length of the label of the file system, in Unicode characters, including
	// the terminating null character.
	LabelLength int32 `idl:"name:cchLabel" json:"label_length"`
	// label:  Null-terminated label of the file system. This is Unicode.
	Label string `idl:"name:label;size_is:(cchLabel)" json:"label"`
}

func (o *FileSystemInfo) xxx_PreparePayload(ctx context.Context) error {
	if o.Label != "" && o.LabelLength == 0 {
		o.LabelLength = int32(len(o.Label))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *FileSystemInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.ID); err != nil {
		return err
	}
	if err := w.WriteData(o.StorageID); err != nil {
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
	if err := w.WriteData(o.Fsflags); err != nil {
		return err
	}
	if err := w.WriteData(o.LastKnownState); err != nil {
		return err
	}
	if err := w.WriteData(o.TaskID); err != nil {
		return err
	}
	if err := w.WriteData(o.FSType); err != nil {
		return err
	}
	if err := w.WriteData(o.LabelLength); err != nil {
		return err
	}
	if o.Label != "" || o.LabelLength > 0 {
		_ptr_label := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.LabelLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			_Label_buf := utf16.Encode([]rune(o.Label))
			if uint64(len(_Label_buf)) > sizeInfo[0] {
				_Label_buf = _Label_buf[:sizeInfo[0]]
			}
			for i1 := range _Label_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_Label_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_Label_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Label, _ptr_label); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *FileSystemInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.ID); err != nil {
		return err
	}
	if err := w.ReadData(&o.StorageID); err != nil {
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
	if err := w.ReadData(&o.Fsflags); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastKnownState); err != nil {
		return err
	}
	if err := w.ReadData(&o.TaskID); err != nil {
		return err
	}
	if err := w.ReadData(&o.FSType); err != nil {
		return err
	}
	if err := w.ReadData(&o.LabelLength); err != nil {
		return err
	}
	_ptr_label := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.LabelLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.LabelLength)
		}
		var _Label_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Label_buf", sizeInfo[0])
		}
		_Label_buf = make([]uint16, sizeInfo[0])
		for i1 := range _Label_buf {
			i1 := i1
			if err := w.ReadData(&_Label_buf[i1]); err != nil {
				return err
			}
		}
		o.Label = strings.TrimRight(string(utf16.Decode(_Label_buf)), ndr.ZeroString)
		return nil
	})
	_s_label := func(ptr interface{}) { o.Label = *ptr.(*string) }
	if err := w.ReadPointer(&o.Label, _s_label, _ptr_label); err != nil {
		return err
	}
	return nil
}

// InstalledFileSystemInfo structure represents IFILE_SYSTEM_INFO RPC structure.
//
// The IFILE_SYSTEM_INFO structure provides information about an installed file system.
// For more information, see [MSDN-STC].
type InstalledFileSystemInfo struct {
	// fsType:   Type of the file system. This field contains one of the following values.
	//
	//	+---------------------------+----------------------------------------+
	//	|                           |                                        |
	//	|           VALUE           |                MEANING                 |
	//	|                           |                                        |
	//	+---------------------------+----------------------------------------+
	//	+---------------------------+----------------------------------------+
	//	| FSTYPE_UNKNOWN 0x00000000 | File system type is unknown.           |
	//	+---------------------------+----------------------------------------+
	//	| FSTYPE_NTFS 0x00000001    | File system type is NTFS.              |
	//	+---------------------------+----------------------------------------+
	//	| FSTYPE_FAT 0x00000002     | File system type is FAT.               |
	//	+---------------------------+----------------------------------------+
	//	| FSTYPE_FAT32 0x00000003   | File system type is FAT32 file system. |
	//	+---------------------------+----------------------------------------+
	//	| FSTYPE_CDFS 0x00000004    | File system type is CDFS.              |
	//	+---------------------------+----------------------------------------+
	//	| FSTYPE_UDF 0x00000005     | File system type is UDF.               |
	//	+---------------------------+----------------------------------------+
	//	| FSTYPE_OTHER 0x80000000   | File system type is not listed.        |
	//	+---------------------------+----------------------------------------+
	FSType int32 `idl:"name:fsType" json:"fs_type"`
	// fsName:  Null-terminated Unicode file system name.
	FSName []uint16 `idl:"name:fsName" json:"fs_name"`
	// fsFlags:  Bitmap of file system flags. The value of this field is a logical OR of
	// zero or more of the applicable flags.
	//
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	|                                        |                                                                                  |
	//	|                 VALUE                  |                                     MEANING                                      |
	//	|                                        |                                                                                  |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| FSF_FMT_OPTION_COMPRESS 0x00000001     | File system supports compression.                                                |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| FSF_FMT_OPTION_LABEL 0x00000002        | File system supports label specification.                                        |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| FSF_MNT_POINT_SUPPORT 0x00000004       | File system supports creation of mount points.                                   |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| FSF_REMOVABLE_MEDIA_SUPPORT 0x00000008 | File system supports creation of removable media.                                |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| FSF_FS_GROW_SUPPORT 0x00000010         | File system supports the extend operation.                                       |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| FSF_FS_QUICK_FORMAT_ENABLE 0x00000020  | File system supports quick formatting.                                           |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| FSF_FS_ALLOC_SZ_512 0x00000040         | File system supports an allocation unit size of 512 bytes.                       |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| FSF_FS_ALLOC_SZ_1K 0x00000080          | File system supports an allocation unit size of 1 kilobyte.                      |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| FSF_FS_ALLOC_SZ_2K 0x00000100          | File system supports an allocation unit size of 2 kilobytes.                     |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| FSF_FS_ALLOC_SZ_4K 0x00000200          | File system supports an allocation unit size of 4 kilobytes.                     |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| FSF_FS_ALLOC_SZ_8K 0x00000400          | File system supports an allocation unit size of 8 kilobytes.                     |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| FSF_FS_ALLOC_SZ_16K 0x00000800         | File system supports an allocation unit size of 16 kilobytes.                    |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| FSF_FS_ALLOC_SZ_32K 0x00001000         | File system supports an allocation unit size of 32 kilobytes.                    |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| FSF_FS_ALLOC_SZ_64K 0x00002000         | File system supports an allocation unit size of 64 kilobytes.                    |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| FSF_FS_ALLOC_SZ_128K 0x00004000        | File system supports an allocation unit size of 128 kilobytes.                   |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| FSF_FS_ALLOC_SZ_256K 0x00008000        | File system supports an allocation unit size of 256 kilobytes.                   |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| FSF_FS_ALLOC_SZ_OTHER 0x00010000       | File system supports any allocation unit size that the user provides.            |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| FSF_FS_FORMAT_SUPPORTED 0x00020000     | File system supports formatting.                                                 |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| FSF_FS_VALID_BITS 0x0003FFFF           | All other bits in the bitmap MUST be ignored. The server does a bitwise AND      |
	//	|                                        | operation with this value to clear upper-level bits that may be present but are  |
	//	|                                        | not supported.                                                                   |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	FSFlags uint32 `idl:"name:fsFlags" json:"fs_flags"`
	// fsCompressionFlags:  Bitmap of allocation unit sizes that are valid for compression.
	// The value of this field is a logical 'OR' of zero or more of the applicable flags.
	//
	//	+----------------------------------+-----------------------------------------------------------------------+
	//	|                                  |                                                                       |
	//	|              VALUE               |                                MEANING                                |
	//	|                                  |                                                                       |
	//	+----------------------------------+-----------------------------------------------------------------------+
	//	+----------------------------------+-----------------------------------------------------------------------+
	//	| FSF_FS_ALLOC_SZ_1K 0x00000080    | File system supports an allocation unit size of 1 kilobyte.           |
	//	+----------------------------------+-----------------------------------------------------------------------+
	//	| FSF_FS_ALLOC_SZ_2K 0x00000100    | File system supports an allocation unit size of 2 kilobytes.          |
	//	+----------------------------------+-----------------------------------------------------------------------+
	//	| FSF_FS_ALLOC_SZ_4K 0x00000200    | File system supports an allocation unit size of 4 kilobytes.          |
	//	+----------------------------------+-----------------------------------------------------------------------+
	//	| FSF_FS_ALLOC_SZ_8K 0x00000400    | File system supports an allocation unit size of 8 kilobytes.          |
	//	+----------------------------------+-----------------------------------------------------------------------+
	//	| FSF_FS_ALLOC_SZ_16K 0x00000800   | File system supports an allocation unit size of 16 kilobytes.         |
	//	+----------------------------------+-----------------------------------------------------------------------+
	//	| FSF_FS_ALLOC_SZ_32K 0x00001000   | File system supports an allocation unit size of 32 kilobytes.         |
	//	+----------------------------------+-----------------------------------------------------------------------+
	//	| FSF_FS_ALLOC_SZ_64K 0x00002000   | File system supports an allocation unit size of 64 kilobytes.         |
	//	+----------------------------------+-----------------------------------------------------------------------+
	//	| FSF_FS_ALLOC_SZ_128K 0x00004000  | File system supports an allocation unit size of 128 kilobytes.        |
	//	+----------------------------------+-----------------------------------------------------------------------+
	//	| FSF_FS_ALLOC_SZ_256K 0x00008000  | File system supports an allocation unit size of 256 kilobytes.        |
	//	+----------------------------------+-----------------------------------------------------------------------+
	//	| FSF_FS_ALLOC_SZ_OTHER 0x00010000 | File system supports any allocation unit size that the user provides. |
	//	+----------------------------------+-----------------------------------------------------------------------+
	FSCompressionFlags uint32 `idl:"name:fsCompressionFlags" json:"fs_compression_flags"`
	// cchLabelLimit:  Maximum number of characters allowed in the file system's label.
	LabelLimitLength int32 `idl:"name:cchLabelLimit" json:"label_limit_length"`
	// cchLabel:  Length of the iLabelChSet member in bytes.
	LabelLength int32 `idl:"name:cchLabel" json:"label_length"`
	// iLabelChSet:  Array of characters that are not allowed in the file system's label.
	LabelCharSet string `idl:"name:iLabelChSet;size_is:(cchLabel)" json:"label_char_set"`
}

func (o *InstalledFileSystemInfo) xxx_PreparePayload(ctx context.Context) error {
	if o.LabelCharSet != "" && o.LabelLength == 0 {
		o.LabelLength = int32(len(o.LabelCharSet))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *InstalledFileSystemInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.FSType); err != nil {
		return err
	}
	for i1 := range o.FSName {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.FSName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.FSName); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.FSFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.FSCompressionFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.LabelLimitLength); err != nil {
		return err
	}
	if err := w.WriteData(o.LabelLength); err != nil {
		return err
	}
	if o.LabelCharSet != "" || o.LabelLength > 0 {
		_ptr_iLabelChSet := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.LabelLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			_LabelCharSet_buf := utf16.Encode([]rune(o.LabelCharSet))
			if uint64(len(_LabelCharSet_buf)) > sizeInfo[0] {
				_LabelCharSet_buf = _LabelCharSet_buf[:sizeInfo[0]]
			}
			for i1 := range _LabelCharSet_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_LabelCharSet_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_LabelCharSet_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.LabelCharSet, _ptr_iLabelChSet); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *InstalledFileSystemInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.FSType); err != nil {
		return err
	}
	o.FSName = make([]uint16, 8)
	for i1 := range o.FSName {
		i1 := i1
		if err := w.ReadData(&o.FSName[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.FSFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.FSCompressionFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.LabelLimitLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.LabelLength); err != nil {
		return err
	}
	_ptr_iLabelChSet := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.LabelLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.LabelLength)
		}
		var _LabelCharSet_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _LabelCharSet_buf", sizeInfo[0])
		}
		_LabelCharSet_buf = make([]uint16, sizeInfo[0])
		for i1 := range _LabelCharSet_buf {
			i1 := i1
			if err := w.ReadData(&_LabelCharSet_buf[i1]); err != nil {
				return err
			}
		}
		o.LabelCharSet = strings.TrimRight(string(utf16.Decode(_LabelCharSet_buf)), ndr.ZeroString)
		return nil
	})
	_s_iLabelChSet := func(ptr interface{}) { o.LabelCharSet = *ptr.(*string) }
	if err := w.ReadPointer(&o.LabelCharSet, _s_iLabelChSet, _ptr_iLabelChSet); err != nil {
		return err
	}
	return nil
}

// TaskInfo structure represents TASK_INFO RPC structure.
//
// The TASK_INFO structure provides information about a task on the server.
type TaskInfo struct {
	// id:  Specifies the OID for the task.
	ID int64 `idl:"name:id" json:"id"`
	// storageId:  Specifies the OID of the object associated with the task.
	StorageID int64 `idl:"name:storageId" json:"storage_id"`
	// createTime:  Unused. This field MUST be set to 0 by servers and ignored by clients.
	CreateTime int64 `idl:"name:createTime" json:"create_time"`
	// clientID:  Specifies the OID of the client that requested the task.
	ClientID int64 `idl:"name:clientID" json:"client_id"`
	// percentComplete:  Percentage of the task that is complete. This field MUST be between
	// 0 and 100, inclusive.
	PercentComplete uint32 `idl:"name:percentComplete" json:"percent_complete"`
	// status:  Specifies the status of the request.
	Status RequestStatus `idl:"name:status" json:"status"`
	// type:  Specifies the kind of operation referred to by the percentComplete member.
	// For more information, see section 2.2.18.
	Type ProgressType `idl:"name:type" json:"type"`
	// error:  The HRESULT error if the value of the status member is REQ_FAILED.
	Error int32 `idl:"name:error" json:"error"`
	// tflag:  Unused. This field MUST be set to 0 by servers and ignored by clients.
	//
	// A TASK_INFO structure is returned by all Disk Management methods that perform configuration
	// operations. The TASK_INFO structure provides information about the task that is being
	// performed by the server in response to the request. The id member of this structure
	// identifies this task from all other tasks being performed by the server. Notifications
	// received by the client as a task progresses can be associated with the original request
	// by comparing the taskId member of the notification structure with the id member of
	// this structure.
	TaskFlag uint32 `idl:"name:tflag" json:"task_flag"`
}

func (o *TaskInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TaskInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.ID); err != nil {
		return err
	}
	if err := w.WriteData(o.StorageID); err != nil {
		return err
	}
	if err := w.WriteData(o.CreateTime); err != nil {
		return err
	}
	if err := w.WriteData(o.ClientID); err != nil {
		return err
	}
	if err := w.WriteData(o.PercentComplete); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Status)); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Type)); err != nil {
		return err
	}
	if err := w.WriteData(o.Error); err != nil {
		return err
	}
	if err := w.WriteData(o.TaskFlag); err != nil {
		return err
	}
	return nil
}
func (o *TaskInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.ID); err != nil {
		return err
	}
	if err := w.ReadData(&o.StorageID); err != nil {
		return err
	}
	if err := w.ReadData(&o.CreateTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClientID); err != nil {
		return err
	}
	if err := w.ReadData(&o.PercentComplete); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Status)); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Type)); err != nil {
		return err
	}
	if err := w.ReadData(&o.Error); err != nil {
		return err
	}
	if err := w.ReadData(&o.TaskFlag); err != nil {
		return err
	}
	return nil
}

// CountedString structure represents COUNTED_STRING RPC structure.
//
// The COUNTED_STRING structure provides information about a mounted folder.
type CountedString struct {
	// sourceId:  Specifies the OID of the source volume. The source volume has a folder
	// to which the target volume will be mounted.
	SourceID int64 `idl:"name:sourceId" json:"source_id"`
	// targetId:  Specifies the OID of the target volume.
	TargetID int64 `idl:"name:targetId" json:"target_id"`
	// cchString:  Specifies the length of the mount path, including the terminating null
	// character.
	StringLength int32 `idl:"name:cchString" json:"string_length"`
	// sstring:  Null-terminated Unicode string that contains the mount path of the source.
	String string `idl:"name:sstring;size_is:(cchString)" json:"string"`
}

func (o *CountedString) xxx_PreparePayload(ctx context.Context) error {
	if o.String != "" && o.StringLength == 0 {
		o.StringLength = int32(len(o.String))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CountedString) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.SourceID); err != nil {
		return err
	}
	if err := w.WriteData(o.TargetID); err != nil {
		return err
	}
	if err := w.WriteData(o.StringLength); err != nil {
		return err
	}
	if o.String != "" || o.StringLength > 0 {
		_ptr_sstring := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.StringLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			_String_buf := utf16.Encode([]rune(o.String))
			if uint64(len(_String_buf)) > sizeInfo[0] {
				_String_buf = _String_buf[:sizeInfo[0]]
			}
			for i1 := range _String_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_String_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_String_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.String, _ptr_sstring); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *CountedString) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.SourceID); err != nil {
		return err
	}
	if err := w.ReadData(&o.TargetID); err != nil {
		return err
	}
	if err := w.ReadData(&o.StringLength); err != nil {
		return err
	}
	_ptr_sstring := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.StringLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.StringLength)
		}
		var _String_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _String_buf", sizeInfo[0])
		}
		_String_buf = make([]uint16, sizeInfo[0])
		for i1 := range _String_buf {
			i1 := i1
			if err := w.ReadData(&_String_buf[i1]); err != nil {
				return err
			}
		}
		o.String = strings.TrimRight(string(utf16.Decode(_String_buf)), ndr.ZeroString)
		return nil
	})
	_s_sstring := func(ptr interface{}) { o.String = *ptr.(*string) }
	if err := w.ReadPointer(&o.String, _s_sstring, _ptr_sstring); err != nil {
		return err
	}
	return nil
}

// MergeObjectInfo structure represents MERGE_OBJECT_INFO RPC structure.
//
// The MERGE_OBJECT_INFO structure provides change information for a merge operation.
type MergeObjectInfo struct {
	// type:  This parameter MUST be set to 0x00000001.
	Type uint32 `idl:"name:type" json:"type"`
	// flags:   Bitmap of merge flags. The value of this field is generated by combining
	// zero or more of the applicable flags with a logical OR operation.
	//
	//	+---------------------------------------+------------------------------------------------------------+
	//	|                                       |                                                            |
	//	|                 VALUE                 |                          MEANING                           |
	//	|                                       |                                                            |
	//	+---------------------------------------+------------------------------------------------------------+
	//	+---------------------------------------+------------------------------------------------------------+
	//	| DSKMERGE_DELETE 0x00000001            | Volume will be deleted.                                    |
	//	+---------------------------------------+------------------------------------------------------------+
	//	| DSKMERGE_DELETE_REDUNDANCY 0x00000002 | Redundant data in a fault-tolerant volume will be deleted. |
	//	+---------------------------------------+------------------------------------------------------------+
	//	| DSKMERGE_STALE_DATA 0x00000004        | Volume contents will be stale.                             |
	//	+---------------------------------------+------------------------------------------------------------+
	//	| DSKMERGE_RELATED 0x00000008           | Volume has subdisks on merged disks.                       |
	//	+---------------------------------------+------------------------------------------------------------+
	Flags uint32 `idl:"name:flags" json:"flags"`
	// layout:  Value from the VOLUMELAYOUT enumeration that indicates the volume's new
	// layout.
	Layout VolumeLayout `idl:"name:layout" json:"layout"`
	// length:  Volume's new size in bytes.
	Length int64 `idl:"name:length" json:"length"`
}

func (o *MergeObjectInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *MergeObjectInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Layout)); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	return nil
}
func (o *MergeObjectInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Layout)); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	return nil
}

// PartitionStyle type represents PARTITIONSTYLE RPC enumeration.
//
// The PARTITIONSTYLE enumeration defines the style of a partition.
type PartitionStyle uint16

var (
	// PARTITIONSTYLE_UNKNOWN:  Partition is of an unknown style.
	PartitionStyleUnknown PartitionStyle = 0
	// PARTITIONSTYLE_MBR:  Partition is of the MBR style.
	PartitionStyleMBR PartitionStyle = 1
	// PARTITIONSTYLE_GPT:  Partition is of the GPT style.
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

// DiskInfoEx structure represents DISK_INFO_EX RPC structure.
//
// The DISK_INFO_EX structure provides information about a disk.
type DiskInfoEx struct {
	// id:  Specifies the OID of the disk.
	ID int64 `idl:"name:id" json:"id"`
	// length:   Size of the disk in bytes.
	Length int64 `idl:"name:length" json:"length"`
	// freeBytes:  Number of unallocated bytes on the disk.
	FreeBytes int64 `idl:"name:freeBytes" json:"free_bytes"`
	// bytesPerTrack:  Size of a disk track in bytes.
	BytesPerTrack uint32 `idl:"name:bytesPerTrack" json:"bytes_per_track"`
	// bytesPerCylinder:  Size of a disk cylinder in bytes.
	BytesPerCylinder uint32 `idl:"name:bytesPerCylinder" json:"bytes_per_cylinder"`
	// bytesPerSector:  Size of a disk sector in bytes.
	BytesPerSector uint32 `idl:"name:bytesPerSector" json:"bytes_per_sector"`
	// regionCount:  Total number of regions on the disk.
	RegionCount uint32 `idl:"name:regionCount" json:"region_count"`
	// dflags:  Disk type of the disk. The value of this field is generated by combining
	// zero or more of the applicable flags with a logical OR operation.
	//
	//	+-------------------------------------+--------------------------------------------------+
	//	|                                     |                                                  |
	//	|                VALUE                |                     MEANING                      |
	//	|                                     |                                                  |
	//	+-------------------------------------+--------------------------------------------------+
	//	+-------------------------------------+--------------------------------------------------+
	//	| DISK_AUDIO_CD 0x00000001            | Disk is an audio CD.                             |
	//	+-------------------------------------+--------------------------------------------------+
	//	| DISK_NEC98 0x00000002               | This value is obsolete and MUST NOT be returned. |
	//	+-------------------------------------+--------------------------------------------------+
	//	| DISK_FORMATTABLE_DVD 0x00000004     | Disk is a DVD that can be formatted.             |
	//	+-------------------------------------+--------------------------------------------------+
	//	| DISK_MEMORY_STICK 0x00000008        | Disk is a memory stick.                          |
	//	+-------------------------------------+--------------------------------------------------+
	//	| DISK_NTFS_NOT_SUPPORTED 0x000000010 | Disk does not support being formatted as NTFS.   |
	//	+-------------------------------------+--------------------------------------------------+
	DiskFlags uint32 `idl:"name:dflags" json:"disk_flags"`
	// deviceType:  Device type of the disk.
	//
	//	+---------------------------------+-------------------------------+
	//	|                                 |                               |
	//	|              VALUE              |            MEANING            |
	//	|                                 |                               |
	//	+---------------------------------+-------------------------------+
	//	+---------------------------------+-------------------------------+
	//	| DEVICETYPE_UNKNOWN 0x00000000   | Device is of an unknown type. |
	//	+---------------------------------+-------------------------------+
	//	| DEVICETYPE_VM 0x00000001        | Device is a dynamic disk.     |
	//	+---------------------------------+-------------------------------+
	//	| DEVICETYPE_REMOVABLE 0x00000002 | Device uses removable media.  |
	//	+---------------------------------+-------------------------------+
	//	| DEVICETYPE_CDROM 0x00000003     | Device is a CD-ROM.           |
	//	+---------------------------------+-------------------------------+
	//	| DEVICETYPE_FDISK 0x00000004     | Device is a basic disk.       |
	//	+---------------------------------+-------------------------------+
	//	| DEVICETYPE_DVD 0x00000005       | Device is a DVD.              |
	//	+---------------------------------+-------------------------------+
	DeviceType uint32 `idl:"name:deviceType" json:"device_type"`
	// deviceState:   Device state of the disk.
	//
	//	+--------------------------------------+---------------------------------------------------------------------+
	//	|                                      |                                                                     |
	//	|                VALUE                 |                               MEANING                               |
	//	|                                      |                                                                     |
	//	+--------------------------------------+---------------------------------------------------------------------+
	//	+--------------------------------------+---------------------------------------------------------------------+
	//	| DEVICESTATE_UNKNOWN 0x00000000       | Disk is in an unknown state.                                        |
	//	+--------------------------------------+---------------------------------------------------------------------+
	//	| DEVICESTATE_HEALTHY 0x00000001       | Disk is fully functional.                                           |
	//	+--------------------------------------+---------------------------------------------------------------------+
	//	| DEVICESTATE_NO_MEDIA 0x00000002      | Disk has no media.                                                  |
	//	+--------------------------------------+---------------------------------------------------------------------+
	//	| DEVICESTATE_NOSIG 0x00000004         | Disk has an invalid signature.                                      |
	//	+--------------------------------------+---------------------------------------------------------------------+
	//	| DEVICESTATE_BAD 0x00000008           | Disk experienced a geometry failure.                                |
	//	+--------------------------------------+---------------------------------------------------------------------+
	//	| DEVICESTATE_NOT_READY 0x00000010     | Disk is not ready yet.                                              |
	//	+--------------------------------------+---------------------------------------------------------------------+
	//	| DEVICESTATE_MISSING 0x00000020       | Disk is no longer available.                                        |
	//	+--------------------------------------+---------------------------------------------------------------------+
	//	| DEVICESTATE_OFFLINE 0x00000040       | Disk is offline.                                                    |
	//	+--------------------------------------+---------------------------------------------------------------------+
	//	| DEVICESTATE_FAILING 0x00000080       | Disk experienced a physical I/O error.                              |
	//	+--------------------------------------+---------------------------------------------------------------------+
	//	| DEVICESTATE_IMPORT_FAILED 0x00000100 | Disk belongs to a group whose import failed. See disk group import. |
	//	+--------------------------------------+---------------------------------------------------------------------+
	//	| DEVICESTATE_UNCLAIMED 0x00000200     | Disk belongs to a foreign disk group.                               |
	//	+--------------------------------------+---------------------------------------------------------------------+
	DeviceState uint32 `idl:"name:deviceState" json:"device_state"`
	// busType:   Type of bus on which the disk resides.
	//
	//	+----------------------------+--------------------------------------+
	//	|                            |                                      |
	//	|           VALUE            |               MEANING                |
	//	|                            |                                      |
	//	+----------------------------+--------------------------------------+
	//	+----------------------------+--------------------------------------+
	//	| BUSTYPE_UNKNOWN 0x00000000 | Bus type is unknown.                 |
	//	+----------------------------+--------------------------------------+
	//	| BUSTYPE_IDE 0x00000001     | Disk resides on an IDE bus.          |
	//	+----------------------------+--------------------------------------+
	//	| BUSTYPE_SCSI 0x00000002    | Disk resides on an SCSI bus.         |
	//	+----------------------------+--------------------------------------+
	//	| BUSTYPE_FIBRE 0x00000003   | Disk resides on a fiber channel bus. |
	//	+----------------------------+--------------------------------------+
	//	| BUSTYPE_USB 0x00000004     | Disk resides on a USB.               |
	//	+----------------------------+--------------------------------------+
	//	| BUSTYPE_SSA 0x00000005     | Disk resides on an SSA bus.          |
	//	+----------------------------+--------------------------------------+
	//	| BUSTYPE_1394 0x00000006    | Disk resides on an IEEE 1394 bus.    |
	//	+----------------------------+--------------------------------------+
	BusType uint32 `idl:"name:busType" json:"bus_type"`
	// attributes:  Bitmap of disk attributes.
	//
	//	+------------------------------+-------------------------+
	//	|                              |                         |
	//	|            VALUE             |         MEANING         |
	//	|                              |                         |
	//	+------------------------------+-------------------------+
	//	+------------------------------+-------------------------+
	//	| DEVICEATTR_NONE 0x00000000   | Disk has no attributes. |
	//	+------------------------------+-------------------------+
	//	| DEVICEATTR_RDONLY 0x00000001 | Disk is read-only.      |
	//	+------------------------------+-------------------------+
	//	| DEVICEATTR_NTMS 0x00000002   | This value is obsolete. |
	//	+------------------------------+-------------------------+
	Attributes uint32 `idl:"name:attributes" json:"attributes"`
	// maxPartitionCount:   Maximum number of partitions on the disk.
	MaxPartitionCount uint32 `idl:"name:maxPartitionCount" json:"max_partition_count"`
	// isUpgradeable:   Boolean value that indicates if the disk can be converted to a dynamic
	// disk. True if the disk is basic, healthy, and has 512-byte sectors.
	//
	//	+---------+-------------------------------------------------------------+
	//	|         |                                                             |
	//	|  VALUE  |                           MEANING                           |
	//	|         |                                                             |
	//	+---------+-------------------------------------------------------------+
	//	+---------+-------------------------------------------------------------+
	//	| FALSE 0 | Disk cannot be encapsulated or converted to a dynamic disk. |
	//	+---------+-------------------------------------------------------------+
	//	| TRUE 1  | Disk can be encapsulated or converted to a dynamic disk.    |
	//	+---------+-------------------------------------------------------------+
	IsUpgradeable bool `idl:"name:isUpgradeable" json:"is_upgradeable"`
	// maySwitchStyle:   Boolean value that indicates if the disk's partition style can
	// be changed from MBR to GPT, or changed from GPT to MBR.
	//
	//	+---------+-----------------------------------------------------------------+
	//	|         |                                                                 |
	//	|  VALUE  |                             MEANING                             |
	//	|         |                                                                 |
	//	+---------+-----------------------------------------------------------------+
	//	+---------+-----------------------------------------------------------------+
	//	| FALSE 0 | Partition style of the disk cannot be changed.                  |
	//	+---------+-----------------------------------------------------------------+
	//	| TRUE 1  | Partition style of the disk can be changed between MBR and GPT. |
	//	+---------+-----------------------------------------------------------------+
	MaySwitchStyle bool `idl:"name:maySwitchStyle" json:"may_switch_style"`
	// partitionStyle:   Value from the PARTITIONSTYLE enumeration that indicates the disk's
	// partitioning style.
	PartitionStyle PartitionStyle         `idl:"name:partitionStyle" json:"partition_style"`
	DiskInfoEx     *DiskInfoEx_DiskInfoEx `idl:"name:DiskInfoEx;switch_is:partitionStyle" json:"disk_info_ex"`
	// portNumber:  SCSI port number of the disk.
	PortNumber int32 `idl:"name:portNumber" json:"port_number"`
	// targetNumber:   SCSI target identifier of the disk.
	TargetNumber int32 `idl:"name:targetNumber" json:"target_number"`
	// lunNumber:  SCSI LUN of the disk.
	LUNNumber int32 `idl:"name:lunNumber" json:"lun_number"`
	// lastKnownState:  Modification sequence number of the disk.
	LastKnownState int64 `idl:"name:lastKnownState" json:"last_known_state"`
	// taskId:   The task identifier of the associated user request. If no request is made,
	// the value is 0.
	TaskID int64 `idl:"name:taskId" json:"task_id"`
	// cchName:   Length of the hard disk's physical name, including the terminating null
	// character.
	NameLength int32 `idl:"name:cchName" json:"name_length"`
	// cchVendor:   Length of the disk's vendor name, including the terminating null character.
	VendorLength int32 `idl:"name:cchVendor" json:"vendor_length"`
	// cchDgid:   Length of the disk's group identification handle, including the terminating
	// null character.
	DiskGroupIDLength int32 `idl:"name:cchDgid" json:"disk_group_id_length"`
	// cchAdapterName:   Length of the disk's adapter name, including the terminating null
	// character.
	AdapterNameLength int32 `idl:"name:cchAdapterName" json:"adapter_name_length"`
	// cchDgName:   Length of the disk's group name, including the terminating null character.
	DiskGroupNameLength int32 `idl:"name:cchDgName" json:"disk_group_name_length"`
	// cchDevInstId:   Length of the disk's device instance path, including the terminating
	// null character.
	DevInstanceIDLength int32 `idl:"name:cchDevInstId" json:"dev_instance_id_length"`
	// name:  Null-terminated physical name of the hard disk. For example: '\device\Harddisk1'.
	Name string `idl:"name:name;size_is:(cchName)" json:"name"`
	// vendor:   Null-terminated name of the hard disk vendor. This is the disk vendor's
	// disk model name. For example: "SEAGATE ST34573N SCSI Disk Device".
	Vendor string `idl:"name:vendor;size_is:(cchVendor)" json:"vendor"`
	// dgid:   Specifies the object identifier of the disk's disk group.
	DiskGroupID []byte `idl:"name:dgid;size_is:(cchDgid)" json:"disk_group_id"`
	// adapterName:   Null-terminated name of the disk adapter. For example: "Adaptec AHA-2940U2W
	// - Ultra2 SCSI".
	AdapterName string `idl:"name:adapterName;size_is:(cchAdapterName)" json:"adapter_name"`
	// dgName:  Null-terminated name for the disk's disk group, if the disk is dynamic.
	DiskGroupName string `idl:"name:dgName;size_is:(cchDgName)" json:"disk_group_name"`
	// devInstId:  Null-terminated device instance path of the disk with the backslashes
	// replaced by "#", "\\?\" prepended to the beginning, and the Pnp disk class GUID,
	// as specified in [MS-DTYP] section 2.3.4.3, appended to the end. For example: "\\?\ide#diskwdc_wd1600jd-75hbb0_____________________08.02d08#5&15c8d966&0&0.0.0#{53f56307-b6bf-11d0-94f2-00a0c91efb8b}".
	DevInstanceID string `idl:"name:devInstId;size_is:(cchDevInstId)" json:"dev_instance_id"`
}

func (o *DiskInfoEx) xxx_PreparePayload(ctx context.Context) error {
	if o.Name != "" && o.NameLength == 0 {
		o.NameLength = int32(len(o.Name))
	}
	if o.Vendor != "" && o.VendorLength == 0 {
		o.VendorLength = int32(len(o.Vendor))
	}
	if o.DiskGroupID != nil && o.DiskGroupIDLength == 0 {
		o.DiskGroupIDLength = int32(len(o.DiskGroupID))
	}
	if o.AdapterName != "" && o.AdapterNameLength == 0 {
		o.AdapterNameLength = int32(len(o.AdapterName))
	}
	if o.DiskGroupName != "" && o.DiskGroupNameLength == 0 {
		o.DiskGroupNameLength = int32(len(o.DiskGroupName))
	}
	if o.DevInstanceID != "" && o.DevInstanceIDLength == 0 {
		o.DevInstanceIDLength = int32(len(o.DevInstanceID))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DiskInfoEx) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.ID); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if err := w.WriteData(o.FreeBytes); err != nil {
		return err
	}
	if err := w.WriteData(o.BytesPerTrack); err != nil {
		return err
	}
	if err := w.WriteData(o.BytesPerCylinder); err != nil {
		return err
	}
	if err := w.WriteData(o.BytesPerSector); err != nil {
		return err
	}
	if err := w.WriteData(o.RegionCount); err != nil {
		return err
	}
	if err := w.WriteData(o.DiskFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.DeviceType); err != nil {
		return err
	}
	if err := w.WriteData(o.DeviceState); err != nil {
		return err
	}
	if err := w.WriteData(o.BusType); err != nil {
		return err
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxPartitionCount); err != nil {
		return err
	}
	if err := w.WriteData(o.IsUpgradeable); err != nil {
		return err
	}
	if err := w.WriteData(o.MaySwitchStyle); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.PartitionStyle)); err != nil {
		return err
	}
	_swDiskInfoEx := uint16(o.PartitionStyle)
	if o.DiskInfoEx != nil {
		if err := o.DiskInfoEx.MarshalUnionNDR(ctx, w, _swDiskInfoEx); err != nil {
			return err
		}
	} else {
		if err := (&DiskInfoEx_DiskInfoEx{}).MarshalUnionNDR(ctx, w, _swDiskInfoEx); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PortNumber); err != nil {
		return err
	}
	if err := w.WriteData(o.TargetNumber); err != nil {
		return err
	}
	if err := w.WriteData(o.LUNNumber); err != nil {
		return err
	}
	if err := w.WriteData(o.LastKnownState); err != nil {
		return err
	}
	if err := w.WriteData(o.TaskID); err != nil {
		return err
	}
	if err := w.WriteData(o.NameLength); err != nil {
		return err
	}
	if err := w.WriteData(o.VendorLength); err != nil {
		return err
	}
	if err := w.WriteData(o.DiskGroupIDLength); err != nil {
		return err
	}
	if err := w.WriteData(o.AdapterNameLength); err != nil {
		return err
	}
	if err := w.WriteData(o.DiskGroupNameLength); err != nil {
		return err
	}
	if err := w.WriteData(o.DevInstanceIDLength); err != nil {
		return err
	}
	if o.Name != "" || o.NameLength > 0 {
		_ptr_name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.NameLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			_Name_buf := utf16.Encode([]rune(o.Name))
			if uint64(len(_Name_buf)) > sizeInfo[0] {
				_Name_buf = _Name_buf[:sizeInfo[0]]
			}
			for i1 := range _Name_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_Name_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_Name_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_name); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Vendor != "" || o.VendorLength > 0 {
		_ptr_vendor := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.VendorLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			_Vendor_buf := utf16.Encode([]rune(o.Vendor))
			if uint64(len(_Vendor_buf)) > sizeInfo[0] {
				_Vendor_buf = _Vendor_buf[:sizeInfo[0]]
			}
			for i1 := range _Vendor_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_Vendor_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_Vendor_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Vendor, _ptr_vendor); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DiskGroupID != nil || o.DiskGroupIDLength > 0 {
		_ptr_dgid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.DiskGroupIDLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.DiskGroupID {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.DiskGroupID[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.DiskGroupID); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.DiskGroupID, _ptr_dgid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.AdapterName != "" || o.AdapterNameLength > 0 {
		_ptr_adapterName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.AdapterNameLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			_AdapterName_buf := utf16.Encode([]rune(o.AdapterName))
			if uint64(len(_AdapterName_buf)) > sizeInfo[0] {
				_AdapterName_buf = _AdapterName_buf[:sizeInfo[0]]
			}
			for i1 := range _AdapterName_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_AdapterName_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_AdapterName_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.AdapterName, _ptr_adapterName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DiskGroupName != "" || o.DiskGroupNameLength > 0 {
		_ptr_dgName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.DiskGroupNameLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			_DiskGroupName_buf := utf16.Encode([]rune(o.DiskGroupName))
			if uint64(len(_DiskGroupName_buf)) > sizeInfo[0] {
				_DiskGroupName_buf = _DiskGroupName_buf[:sizeInfo[0]]
			}
			for i1 := range _DiskGroupName_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_DiskGroupName_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_DiskGroupName_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.DiskGroupName, _ptr_dgName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DevInstanceID != "" || o.DevInstanceIDLength > 0 {
		_ptr_devInstId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.DevInstanceIDLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			_DevInstanceID_buf := utf16.Encode([]rune(o.DevInstanceID))
			if uint64(len(_DevInstanceID_buf)) > sizeInfo[0] {
				_DevInstanceID_buf = _DevInstanceID_buf[:sizeInfo[0]]
			}
			for i1 := range _DevInstanceID_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_DevInstanceID_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_DevInstanceID_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.DevInstanceID, _ptr_devInstId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *DiskInfoEx) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.ID); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	if err := w.ReadData(&o.FreeBytes); err != nil {
		return err
	}
	if err := w.ReadData(&o.BytesPerTrack); err != nil {
		return err
	}
	if err := w.ReadData(&o.BytesPerCylinder); err != nil {
		return err
	}
	if err := w.ReadData(&o.BytesPerSector); err != nil {
		return err
	}
	if err := w.ReadData(&o.RegionCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.DiskFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.DeviceType); err != nil {
		return err
	}
	if err := w.ReadData(&o.DeviceState); err != nil {
		return err
	}
	if err := w.ReadData(&o.BusType); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxPartitionCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsUpgradeable); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaySwitchStyle); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.PartitionStyle)); err != nil {
		return err
	}
	if o.DiskInfoEx == nil {
		o.DiskInfoEx = &DiskInfoEx_DiskInfoEx{}
	}
	_swDiskInfoEx := uint16(o.PartitionStyle)
	if err := o.DiskInfoEx.UnmarshalUnionNDR(ctx, w, _swDiskInfoEx); err != nil {
		return err
	}
	if err := w.ReadData(&o.PortNumber); err != nil {
		return err
	}
	if err := w.ReadData(&o.TargetNumber); err != nil {
		return err
	}
	if err := w.ReadData(&o.LUNNumber); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastKnownState); err != nil {
		return err
	}
	if err := w.ReadData(&o.TaskID); err != nil {
		return err
	}
	if err := w.ReadData(&o.NameLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.VendorLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.DiskGroupIDLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.AdapterNameLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.DiskGroupNameLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.DevInstanceIDLength); err != nil {
		return err
	}
	_ptr_name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
		var _Name_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Name_buf", sizeInfo[0])
		}
		_Name_buf = make([]uint16, sizeInfo[0])
		for i1 := range _Name_buf {
			i1 := i1
			if err := w.ReadData(&_Name_buf[i1]); err != nil {
				return err
			}
		}
		o.Name = strings.TrimRight(string(utf16.Decode(_Name_buf)), ndr.ZeroString)
		return nil
	})
	_s_name := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_name, _ptr_name); err != nil {
		return err
	}
	_ptr_vendor := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.VendorLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.VendorLength)
		}
		var _Vendor_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Vendor_buf", sizeInfo[0])
		}
		_Vendor_buf = make([]uint16, sizeInfo[0])
		for i1 := range _Vendor_buf {
			i1 := i1
			if err := w.ReadData(&_Vendor_buf[i1]); err != nil {
				return err
			}
		}
		o.Vendor = strings.TrimRight(string(utf16.Decode(_Vendor_buf)), ndr.ZeroString)
		return nil
	})
	_s_vendor := func(ptr interface{}) { o.Vendor = *ptr.(*string) }
	if err := w.ReadPointer(&o.Vendor, _s_vendor, _ptr_vendor); err != nil {
		return err
	}
	_ptr_dgid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.DiskGroupIDLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.DiskGroupIDLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.DiskGroupID", sizeInfo[0])
		}
		o.DiskGroupID = make([]byte, sizeInfo[0])
		for i1 := range o.DiskGroupID {
			i1 := i1
			if err := w.ReadData(&o.DiskGroupID[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_dgid := func(ptr interface{}) { o.DiskGroupID = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.DiskGroupID, _s_dgid, _ptr_dgid); err != nil {
		return err
	}
	_ptr_adapterName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.AdapterNameLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.AdapterNameLength)
		}
		var _AdapterName_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _AdapterName_buf", sizeInfo[0])
		}
		_AdapterName_buf = make([]uint16, sizeInfo[0])
		for i1 := range _AdapterName_buf {
			i1 := i1
			if err := w.ReadData(&_AdapterName_buf[i1]); err != nil {
				return err
			}
		}
		o.AdapterName = strings.TrimRight(string(utf16.Decode(_AdapterName_buf)), ndr.ZeroString)
		return nil
	})
	_s_adapterName := func(ptr interface{}) { o.AdapterName = *ptr.(*string) }
	if err := w.ReadPointer(&o.AdapterName, _s_adapterName, _ptr_adapterName); err != nil {
		return err
	}
	_ptr_dgName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.DiskGroupNameLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.DiskGroupNameLength)
		}
		var _DiskGroupName_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _DiskGroupName_buf", sizeInfo[0])
		}
		_DiskGroupName_buf = make([]uint16, sizeInfo[0])
		for i1 := range _DiskGroupName_buf {
			i1 := i1
			if err := w.ReadData(&_DiskGroupName_buf[i1]); err != nil {
				return err
			}
		}
		o.DiskGroupName = strings.TrimRight(string(utf16.Decode(_DiskGroupName_buf)), ndr.ZeroString)
		return nil
	})
	_s_dgName := func(ptr interface{}) { o.DiskGroupName = *ptr.(*string) }
	if err := w.ReadPointer(&o.DiskGroupName, _s_dgName, _ptr_dgName); err != nil {
		return err
	}
	_ptr_devInstId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.DevInstanceIDLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.DevInstanceIDLength)
		}
		var _DevInstanceID_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _DevInstanceID_buf", sizeInfo[0])
		}
		_DevInstanceID_buf = make([]uint16, sizeInfo[0])
		for i1 := range _DevInstanceID_buf {
			i1 := i1
			if err := w.ReadData(&_DevInstanceID_buf[i1]); err != nil {
				return err
			}
		}
		o.DevInstanceID = strings.TrimRight(string(utf16.Decode(_DevInstanceID_buf)), ndr.ZeroString)
		return nil
	})
	_s_devInstId := func(ptr interface{}) { o.DevInstanceID = *ptr.(*string) }
	if err := w.ReadPointer(&o.DevInstanceID, _s_devInstId, _ptr_devInstId); err != nil {
		return err
	}
	return nil
}

// DiskInfoEx_DiskInfoEx structure represents DISK_INFO_EX union anonymous member.
//
// The DISK_INFO_EX structure provides information about a disk.
type DiskInfoEx_DiskInfoEx struct {
	// Types that are assignable to Value
	//
	// *DiskInfoEx_MBR
	// *DiskInfoEx_GPT
	Value is_DiskInfoEx_DiskInfoEx `json:"value"`
}

func (o *DiskInfoEx_DiskInfoEx) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *DiskInfoEx_MBR:
		if value != nil {
			return value.MBR
		}
	case *DiskInfoEx_GPT:
		if value != nil {
			return value.GPT
		}
	}
	return nil
}

type is_DiskInfoEx_DiskInfoEx interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_DiskInfoEx_DiskInfoEx()
}

func (o *DiskInfoEx_DiskInfoEx) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *DiskInfoEx_MBR:
		return uint16(1)
	case *DiskInfoEx_GPT:
		return uint16(2)
	}
	return uint16(0)
}

func (o *DiskInfoEx_DiskInfoEx) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(4); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		_o, _ := o.Value.(*DiskInfoEx_MBR)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&DiskInfoEx_MBR{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*DiskInfoEx_GPT)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&DiskInfoEx_GPT{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *DiskInfoEx_DiskInfoEx) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch(ndr.Enum((*uint16)(&sw))); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(4); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		o.Value = &DiskInfoEx_MBR{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &DiskInfoEx_GPT{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// DiskInfoEx_MBR structure represents DiskInfoEx_DiskInfoEx RPC union arm.
//
// It has following labels: 1
type DiskInfoEx_MBR struct {
	MBR *DiskInfoEx_DiskInfoEx_MBR `idl:"name:mbr" json:"mbr"`
}

func (*DiskInfoEx_MBR) is_DiskInfoEx_DiskInfoEx() {}

func (o *DiskInfoEx_MBR) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.MBR != nil {
		if err := o.MBR.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DiskInfoEx_DiskInfoEx_MBR{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DiskInfoEx_MBR) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.MBR == nil {
		o.MBR = &DiskInfoEx_DiskInfoEx_MBR{}
	}
	if err := o.MBR.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DiskInfoEx_DiskInfoEx_MBR structure represents DISK_INFO_EX structure anonymous member.
//
// The DISK_INFO_EX structure provides information about a disk.
type DiskInfoEx_DiskInfoEx_MBR struct {
	// signature:  Signature of the disk. The disk signature is not guaranteed to be unique
	// across machines.<8>
	Signature uint32 `idl:"name:signature" json:"signature"`
}

func (o *DiskInfoEx_DiskInfoEx_MBR) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DiskInfoEx_DiskInfoEx_MBR) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Signature); err != nil {
		return err
	}
	return nil
}
func (o *DiskInfoEx_DiskInfoEx_MBR) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Signature); err != nil {
		return err
	}
	return nil
}

// DiskInfoEx_GPT structure represents DiskInfoEx_DiskInfoEx RPC union arm.
//
// It has following labels: 2
type DiskInfoEx_GPT struct {
	GPT *DiskInfoEx_DiskInfoEx_GPT `idl:"name:gpt" json:"gpt"`
}

func (*DiskInfoEx_GPT) is_DiskInfoEx_DiskInfoEx() {}

func (o *DiskInfoEx_GPT) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.GPT != nil {
		if err := o.GPT.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DiskInfoEx_DiskInfoEx_GPT{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DiskInfoEx_GPT) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.GPT == nil {
		o.GPT = &DiskInfoEx_DiskInfoEx_GPT{}
	}
	if err := o.GPT.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DiskInfoEx_DiskInfoEx_GPT structure represents DISK_INFO_EX structure anonymous member.
//
// The DISK_INFO_EX structure provides information about a disk.
type DiskInfoEx_DiskInfoEx_GPT struct {
	// diskId:  GUID, as specified in [MS-DTYP], section 2.3.4.1, of the disk.<9>
	DiskID *dtyp.GUID `idl:"name:diskId" json:"disk_id"`
}

func (o *DiskInfoEx_DiskInfoEx_GPT) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DiskInfoEx_DiskInfoEx_GPT) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.DiskID != nil {
		if err := o.DiskID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DiskInfoEx_DiskInfoEx_GPT) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.DiskID == nil {
		o.DiskID = &dtyp.GUID{}
	}
	if err := o.DiskID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RegionInfoEx structure represents REGION_INFO_EX RPC structure.
//
// The REGION_INFO_EX structure provides information about a region.
type RegionInfoEx struct {
	// id:  Specifies the region's OID.
	ID int64 `idl:"name:id" json:"id"`
	// diskId:   Specifies the OID of the disk on which the region resides.
	DiskID int64 `idl:"name:diskId" json:"disk_id"`
	// volId:  Specifies the OID of the volume on the region, if any.
	VolID int64 `idl:"name:volId" json:"vol_id"`
	// fsId:  Specifies the OID of the file system on the region, if any.
	FSID int64 `idl:"name:fsId" json:"fs_id"`
	// start:  Byte offset of the region on the disk.
	Start int64 `idl:"name:start" json:"start"`
	// length:  Length of the region in bytes.
	Length int64 `idl:"name:length" json:"length"`
	// regionType:  Value from the REGIONTYPE enumeration that indicates the region type.
	RegionType RegionType `idl:"name:regionType" json:"region_type"`
	// partitionStyle:   Value from the PARTITIONSTYLE enumeration that indicates the region's
	// partitioning style.
	PartitionStyle PartitionStyle             `idl:"name:partitionStyle" json:"partition_style"`
	RegionInfoEx   *RegionInfoEx_RegionInfoEx `idl:"name:RegionInfoEx;switch_is:partitionStyle" json:"region_info_ex"`
	// status:  Value from the REGIONSTATUS enumeration that indicates the region's status.
	Status RegionStatus `idl:"name:status" json:"status"`
	// lastKnownState:  Modification sequence number of the region.
	LastKnownState int64 `idl:"name:lastKnownState" json:"last_known_state"`
	// taskId:   This LdmObjectId is the task identifier of the associated user request.
	// If no request is made, the value MUST be 0.
	TaskID int64 `idl:"name:taskId" json:"task_id"`
	// rflags:  Bitmap of region flags. The value of this field is generated by combining
	// zero or more of the applicable flags with a logical OR operation.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                                       |                                                                                  |
	//	|                 VALUE                 |                                     MEANING                                      |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| REGION_FORMAT_IN_PROGRESS 0x00000001  | Region is currently being formatted.                                             |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| REGION_IS_SYSTEM_PARTITION 0x00000002 | Region contains the system directory. The system directory has the operating     |
	//	|                                       | system installed on it. This is not necessarily the "active" partition that      |
	//	|                                       | contains the boot loader file.                                                   |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| REGION_HAS_PAGEFILE 0x00000004        | Region contains the paging file.                                                 |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| REGION_HAD_BOOT_INI 0x00000040        | Boot.ini file was located in this region when the operating system was last      |
	//	|                                       | started. This is the "active" partition that contains the boot loader file.      |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| REGION_HIDDEN 0x00040000              | This region is part of a volume that is not accessible through any               |
	//	|                                       | user-available path names.<11>                                                   |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	RegionFlags uint32 `idl:"name:rflags" json:"region_flags"`
	// currentPartitionNumber:  Number of the partition on the region, if any.
	CurrentPartitionNumber uint32 `idl:"name:currentPartitionNumber" json:"current_partition_number"`
	// cchName:   Length of the region's name, including the terminating null character.
	NameLength int32 `idl:"name:cchName" json:"name_length"`
	// name:   Null-terminated name of the region.
	Name string `idl:"name:name;size_is:(cchName)" json:"name"`
}

func (o *RegionInfoEx) xxx_PreparePayload(ctx context.Context) error {
	if o.Name != "" && o.NameLength == 0 {
		o.NameLength = int32(len(o.Name))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RegionInfoEx) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.ID); err != nil {
		return err
	}
	if err := w.WriteData(o.DiskID); err != nil {
		return err
	}
	if err := w.WriteData(o.VolID); err != nil {
		return err
	}
	if err := w.WriteData(o.FSID); err != nil {
		return err
	}
	if err := w.WriteData(o.Start); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.RegionType)); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.PartitionStyle)); err != nil {
		return err
	}
	_swRegionInfoEx := uint16(o.PartitionStyle)
	if o.RegionInfoEx != nil {
		if err := o.RegionInfoEx.MarshalUnionNDR(ctx, w, _swRegionInfoEx); err != nil {
			return err
		}
	} else {
		if err := (&RegionInfoEx_RegionInfoEx{}).MarshalUnionNDR(ctx, w, _swRegionInfoEx); err != nil {
			return err
		}
	}
	if err := w.WriteEnum(uint16(o.Status)); err != nil {
		return err
	}
	if err := w.WriteData(o.LastKnownState); err != nil {
		return err
	}
	if err := w.WriteData(o.TaskID); err != nil {
		return err
	}
	if err := w.WriteData(o.RegionFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.CurrentPartitionNumber); err != nil {
		return err
	}
	if err := w.WriteData(o.NameLength); err != nil {
		return err
	}
	if o.Name != "" || o.NameLength > 0 {
		_ptr_name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.NameLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			_Name_buf := utf16.Encode([]rune(o.Name))
			if uint64(len(_Name_buf)) > sizeInfo[0] {
				_Name_buf = _Name_buf[:sizeInfo[0]]
			}
			for i1 := range _Name_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_Name_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_Name_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_name); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *RegionInfoEx) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.ID); err != nil {
		return err
	}
	if err := w.ReadData(&o.DiskID); err != nil {
		return err
	}
	if err := w.ReadData(&o.VolID); err != nil {
		return err
	}
	if err := w.ReadData(&o.FSID); err != nil {
		return err
	}
	if err := w.ReadData(&o.Start); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.RegionType)); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.PartitionStyle)); err != nil {
		return err
	}
	if o.RegionInfoEx == nil {
		o.RegionInfoEx = &RegionInfoEx_RegionInfoEx{}
	}
	_swRegionInfoEx := uint16(o.PartitionStyle)
	if err := o.RegionInfoEx.UnmarshalUnionNDR(ctx, w, _swRegionInfoEx); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Status)); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastKnownState); err != nil {
		return err
	}
	if err := w.ReadData(&o.TaskID); err != nil {
		return err
	}
	if err := w.ReadData(&o.RegionFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.CurrentPartitionNumber); err != nil {
		return err
	}
	if err := w.ReadData(&o.NameLength); err != nil {
		return err
	}
	_ptr_name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
		var _Name_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Name_buf", sizeInfo[0])
		}
		_Name_buf = make([]uint16, sizeInfo[0])
		for i1 := range _Name_buf {
			i1 := i1
			if err := w.ReadData(&_Name_buf[i1]); err != nil {
				return err
			}
		}
		o.Name = strings.TrimRight(string(utf16.Decode(_Name_buf)), ndr.ZeroString)
		return nil
	})
	_s_name := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_name, _ptr_name); err != nil {
		return err
	}
	return nil
}

// RegionInfoEx_RegionInfoEx structure represents REGION_INFO_EX union anonymous member.
//
// The REGION_INFO_EX structure provides information about a region.
type RegionInfoEx_RegionInfoEx struct {
	// Types that are assignable to Value
	//
	// *RegionInfoEx_MBR
	// *RegionInfoEx_GPT
	Value is_RegionInfoEx_RegionInfoEx `json:"value"`
}

func (o *RegionInfoEx_RegionInfoEx) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *RegionInfoEx_MBR:
		if value != nil {
			return value.MBR
		}
	case *RegionInfoEx_GPT:
		if value != nil {
			return value.GPT
		}
	}
	return nil
}

type is_RegionInfoEx_RegionInfoEx interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_RegionInfoEx_RegionInfoEx()
}

func (o *RegionInfoEx_RegionInfoEx) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *RegionInfoEx_MBR:
		return uint16(1)
	case *RegionInfoEx_GPT:
		return uint16(2)
	}
	return uint16(0)
}

func (o *RegionInfoEx_RegionInfoEx) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(8); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		_o, _ := o.Value.(*RegionInfoEx_MBR)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RegionInfoEx_MBR{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*RegionInfoEx_GPT)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RegionInfoEx_GPT{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *RegionInfoEx_RegionInfoEx) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch(ndr.Enum((*uint16)(&sw))); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(8); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		o.Value = &RegionInfoEx_MBR{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &RegionInfoEx_GPT{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// RegionInfoEx_MBR structure represents RegionInfoEx_RegionInfoEx RPC union arm.
//
// It has following labels: 1
type RegionInfoEx_MBR struct {
	MBR *RegionInfoEx_RegionInfoEx_MBR `idl:"name:mbr" json:"mbr"`
}

func (*RegionInfoEx_MBR) is_RegionInfoEx_RegionInfoEx() {}

func (o *RegionInfoEx_MBR) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.MBR != nil {
		if err := o.MBR.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&RegionInfoEx_RegionInfoEx_MBR{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RegionInfoEx_MBR) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.MBR == nil {
		o.MBR = &RegionInfoEx_RegionInfoEx_MBR{}
	}
	if err := o.MBR.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RegionInfoEx_RegionInfoEx_MBR structure represents REGION_INFO_EX structure anonymous member.
//
// The REGION_INFO_EX structure provides information about a region.
type RegionInfoEx_RegionInfoEx_MBR struct {
	// partitionType:  Windows NT 3.1 operating system, Windows NT 3.5 operating system,
	// Windows NT 3.51 operating system, and Windows NT 4.0 operating system partition style
	// for the region. This field contains one of the following values.
	//
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	|                             |                                                                                  |
	//	|            VALUE            |                                     MEANING                                      |
	//	|                             |                                                                                  |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| PARTITION_ENTRY_UNUSED 0x00 | An unused entry partition.                                                       |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| PARTITION_EXTENDED 0x05     | An extended partition.                                                           |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| PARTITION_FAT_12 0x01       | A FAT12 file system partition.                                                   |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| PARTITION_FAT_16 0x04       | A FAT16 file system partition.                                                   |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| PARTITION_FAT32 0x0B        | A FAT32 file system partition.                                                   |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| PARTITION_IFS 0x07          | An IFS partition.                                                                |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| PARTITION_LDM 0x42          | An LDM partition.                                                                |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| PARTITION_NTFT 0x80         | A Windows NT operating system fault-tolerant (FT) partition.                     |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| VALID_NTFT 0xC0             | A valid Windows NT FT partition. The high bit of a partition type code indicates |
	//	|                             | that a partition is part of an NTFT mirror or striped array.                     |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	PartitionType uint32 `idl:"name:partitionType" json:"partition_type"`
	// isActive:  Boolean value that indicates whether the partition is active. The partition
	// MUST be marked as active in order for the BIOS to start from the partition on x86
	// and x64 platforms.
	//
	//	+---------+--------------------------+
	//	|         |                          |
	//	|  VALUE  |         MEANING          |
	//	|         |                          |
	//	+---------+--------------------------+
	//	+---------+--------------------------+
	//	| FALSE 0 | Partition is not active. |
	//	+---------+--------------------------+
	//	| TRUE 1  | Partition is active.     |
	//	+---------+--------------------------+
	IsActive bool `idl:"name:isActive" json:"is_active"`
}

func (o *RegionInfoEx_RegionInfoEx_MBR) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RegionInfoEx_RegionInfoEx_MBR) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.PartitionType); err != nil {
		return err
	}
	if err := w.WriteData(o.IsActive); err != nil {
		return err
	}
	return nil
}
func (o *RegionInfoEx_RegionInfoEx_MBR) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.PartitionType); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsActive); err != nil {
		return err
	}
	return nil
}

// RegionInfoEx_GPT structure represents RegionInfoEx_RegionInfoEx RPC union arm.
//
// It has following labels: 2
type RegionInfoEx_GPT struct {
	GPT *RegionInfoEx_RegionInfoEx_GPT `idl:"name:gpt" json:"gpt"`
}

func (*RegionInfoEx_GPT) is_RegionInfoEx_RegionInfoEx() {}

func (o *RegionInfoEx_GPT) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.GPT != nil {
		if err := o.GPT.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&RegionInfoEx_RegionInfoEx_GPT{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RegionInfoEx_GPT) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.GPT == nil {
		o.GPT = &RegionInfoEx_RegionInfoEx_GPT{}
	}
	if err := o.GPT.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RegionInfoEx_RegionInfoEx_GPT structure represents REGION_INFO_EX structure anonymous member.
//
// The REGION_INFO_EX structure provides information about a region.
type RegionInfoEx_RegionInfoEx_GPT struct {
	// partitionType:  Windows NT 3.1 operating system, Windows NT 3.5 operating system,
	// Windows NT 3.51 operating system, and Windows NT 4.0 operating system partition style
	// for the region. This field contains one of the following values.
	//
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	|                             |                                                                                  |
	//	|            VALUE            |                                     MEANING                                      |
	//	|                             |                                                                                  |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| PARTITION_ENTRY_UNUSED 0x00 | An unused entry partition.                                                       |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| PARTITION_EXTENDED 0x05     | An extended partition.                                                           |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| PARTITION_FAT_12 0x01       | A FAT12 file system partition.                                                   |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| PARTITION_FAT_16 0x04       | A FAT16 file system partition.                                                   |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| PARTITION_FAT32 0x0B        | A FAT32 file system partition.                                                   |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| PARTITION_IFS 0x07          | An IFS partition.                                                                |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| PARTITION_LDM 0x42          | An LDM partition.                                                                |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| PARTITION_NTFT 0x80         | A Windows NT operating system fault-tolerant (FT) partition.                     |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| VALID_NTFT 0xC0             | A valid Windows NT FT partition. The high bit of a partition type code indicates |
	//	|                             | that a partition is part of an NTFT mirror or striped array.                     |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	PartitionType *dtyp.GUID `idl:"name:partitionType" json:"partition_type"`
	// partitionId:  A GUID that uniquely identifies a partition on a disk.
	PartitionID *dtyp.GUID `idl:"name:partitionId" json:"partition_id"`
	// attributes:  Bitmap of partition flags.<10>
	Attributes uint64 `idl:"name:attributes" json:"attributes"`
}

func (o *RegionInfoEx_RegionInfoEx_GPT) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RegionInfoEx_RegionInfoEx_GPT) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	return nil
}
func (o *RegionInfoEx_RegionInfoEx_GPT) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	return nil
}

// VolumeClient4 structure represents IVolumeClient4 RPC structure.
type VolumeClient4 dcom.InterfacePointer

func (o *VolumeClient4) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *VolumeClient4) xxx_PreparePayload(ctx context.Context) error {
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

func (o *VolumeClient4) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *VolumeClient4) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *VolumeClient4) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// VolumeClient3 structure represents IVolumeClient3 RPC structure.
type VolumeClient3 dcom.InterfacePointer

func (o *VolumeClient3) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *VolumeClient3) xxx_PreparePayload(ctx context.Context) error {
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

func (o *VolumeClient3) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *VolumeClient3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *VolumeClient3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// VolumeClient2 structure represents IVolumeClient2 RPC structure.
type VolumeClient2 dcom.InterfacePointer

func (o *VolumeClient2) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *VolumeClient2) xxx_PreparePayload(ctx context.Context) error {
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

func (o *VolumeClient2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *VolumeClient2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *VolumeClient2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// IDMRemoteServer structure represents IDMRemoteServer RPC structure.
type IDMRemoteServer dcom.InterfacePointer

func (o *IDMRemoteServer) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *IDMRemoteServer) xxx_PreparePayload(ctx context.Context) error {
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

func (o *IDMRemoteServer) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *IDMRemoteServer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *IDMRemoteServer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// VolumeClient structure represents IVolumeClient RPC structure.
type VolumeClient dcom.InterfacePointer

func (o *VolumeClient) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *VolumeClient) xxx_PreparePayload(ctx context.Context) error {
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

func (o *VolumeClient) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *VolumeClient) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *VolumeClient) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// IDMNotify structure represents IDMNotify RPC structure.
type IDMNotify dcom.InterfacePointer

func (o *IDMNotify) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *IDMNotify) xxx_PreparePayload(ctx context.Context) error {
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

func (o *IDMNotify) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *IDMNotify) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *IDMNotify) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
