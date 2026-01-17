package ivolumeclient3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
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
	_ = iunknown.GoPackage
)

// IVolumeClient3 server interface.
type VolumeClient3Server interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The EnumDisksEx method enumerates the server's mass storage devices.
	//
	// Return Values: The method MUST return 0 or a nonerror HRESULT on success, or an implementation-specific
	// nonzero error code on failure (as specified in [MS-ERREF]; see also section 2.2.1
	// for HRESULT values predefined by the Disk Management Remote Protocol).
	EnumDisksEx(context.Context, *EnumDisksExRequest) (*EnumDisksExResponse, error)

	// The EnumDiskRegionsEx method enumerates all used and free regions of a specified
	// disk.
	//
	// Return Values: The method MUST return 0 or a nonerror HRESULT on success, or an implementation-specific
	// nonzero error code on failure (as specified in [MS-ERREF]; see also section 2.2.1
	// for HRESULT values predefined by the Disk Management Remote Protocol).
	EnumDiskRegionsEx(context.Context, *EnumDiskRegionsExRequest) (*EnumDiskRegionsExResponse, error)

	// CreatePartition operation.
	CreatePartition(context.Context, *CreatePartitionRequest) (*CreatePartitionResponse, error)

	// CreatePartitionAssignAndFormat operation.
	CreatePartitionAssignAndFormat(context.Context, *CreatePartitionAssignAndFormatRequest) (*CreatePartitionAssignAndFormatResponse, error)

	// CreatePartitionAssignAndFormatEx operation.
	CreatePartitionAssignAndFormatEx(context.Context, *CreatePartitionAssignAndFormatExRequest) (*CreatePartitionAssignAndFormatExResponse, error)

	// DeletePartition operation.
	DeletePartition(context.Context, *DeletePartitionRequest) (*DeletePartitionResponse, error)

	// The InitializeDiskStyle method sets the partition style and writes a signature to
	// a disk. This is a synchronous task.
	//
	// Return Values: The method MUST return 0 or a nonerror HRESULT on success, or an implementation-specific
	// nonzero error code on failure (as specified in [MS-ERREF]; see also section 2.2.1
	// for HRESULT values predefined by the Disk Management Remote Protocol).
	InitializeDiskStyle(context.Context, *InitializeDiskStyleRequest) (*InitializeDiskStyleResponse, error)

	// MarkActivePartition operation.
	MarkActivePartition(context.Context, *MarkActivePartitionRequest) (*MarkActivePartitionResponse, error)

	// Eject operation.
	Eject(context.Context, *EjectRequest) (*EjectResponse, error)

	// Reserved_Opnum12 operation.
	// ReservedOpnum12

	// FTEnumVolumes operation.
	FTEnumVolumes(context.Context, *FTEnumVolumesRequest) (*FTEnumVolumesResponse, error)

	// FTEnumLogicalDiskMembers operation.
	FTEnumLogicalDiskMembers(context.Context, *FTEnumLogicalDiskMembersRequest) (*FTEnumLogicalDiskMembersResponse, error)

	// FTDeleteVolume operation.
	FTDeleteVolume(context.Context, *FTDeleteVolumeRequest) (*FTDeleteVolumeResponse, error)

	// FTBreakMirror operation.
	FTBreakMirror(context.Context, *FTBreakMirrorRequest) (*FTBreakMirrorResponse, error)

	// FTResyncMirror operation.
	FTResyncMirror(context.Context, *FTResyncMirrorRequest) (*FTResyncMirrorResponse, error)

	// FTRegenerateParityStripe operation.
	FTRegenerateParityStripe(context.Context, *FTRegenerateParityStripeRequest) (*FTRegenerateParityStripeResponse, error)

	// FTReplaceMirrorPartition operation.
	FTReplaceMirrorPartition(context.Context, *FTReplaceMirrorPartitionRequest) (*FTReplaceMirrorPartitionResponse, error)

	// FTReplaceParityStripePartition operation.
	FTReplaceParityStripePartition(context.Context, *FTReplaceParityStripePartitionRequest) (*FTReplaceParityStripePartitionResponse, error)

	// EnumDriveLetters operation.
	EnumDriveLetters(context.Context, *EnumDriveLettersRequest) (*EnumDriveLettersResponse, error)

	// AssignDriveLetter operation.
	AssignDriveLetter(context.Context, *AssignDriveLetterRequest) (*AssignDriveLetterResponse, error)

	// FreeDriveLetter operation.
	FreeDriveLetter(context.Context, *FreeDriveLetterRequest) (*FreeDriveLetterResponse, error)

	// EnumLocalFileSystems operation.
	EnumLocalFileSystems(context.Context, *EnumLocalFileSystemsRequest) (*EnumLocalFileSystemsResponse, error)

	// GetInstalledFileSystems operation.
	GetInstalledFileSystems(context.Context, *GetInstalledFileSystemsRequest) (*GetInstalledFileSystemsResponse, error)

	// Format operation.
	Format(context.Context, *FormatRequest) (*FormatResponse, error)

	// EnumVolumes operation.
	EnumVolumes(context.Context, *EnumVolumesRequest) (*EnumVolumesResponse, error)

	// EnumVolumeMembers operation.
	EnumVolumeMembers(context.Context, *EnumVolumeMembersRequest) (*EnumVolumeMembersResponse, error)

	// CreateVolume operation.
	CreateVolume(context.Context, *CreateVolumeRequest) (*CreateVolumeResponse, error)

	// CreateVolumeAssignAndFormat operation.
	CreateVolumeAssignAndFormat(context.Context, *CreateVolumeAssignAndFormatRequest) (*CreateVolumeAssignAndFormatResponse, error)

	// CreateVolumeAssignAndFormatEx operation.
	CreateVolumeAssignAndFormatEx(context.Context, *CreateVolumeAssignAndFormatExRequest) (*CreateVolumeAssignAndFormatExResponse, error)

	// GetVolumeMountName operation.
	GetVolumeMountName(context.Context, *GetVolumeMountNameRequest) (*GetVolumeMountNameResponse, error)

	// GrowVolume operation.
	GrowVolume(context.Context, *GrowVolumeRequest) (*GrowVolumeResponse, error)

	// DeleteVolume operation.
	DeleteVolume(context.Context, *DeleteVolumeRequest) (*DeleteVolumeResponse, error)

	// The CreatePartitionsForVolume method creates a partition underneath a volume. This
	// is a synchronous task.
	//
	// Return Values: The method MUST return 0 or a nonerror HRESULT on success, or an implementation-specific
	// nonzero error code on failure (as specified in [MS-ERREF]; see also section 2.2.1
	// for HRESULT values predefined by the Disk Management Remote Protocol).
	CreatePartitionsForVolume(context.Context, *CreatePartitionsForVolumeRequest) (*CreatePartitionsForVolumeResponse, error)

	// The DeletePartitionsForVolume method deletes the partitions underneath a dynamic
	// disk volume. This is a synchronous task.
	//
	// Return Values: The method MUST return 0 or a nonerror HRESULT on success, or an implementation-specific
	// nonzero error code on failure (as specified in [MS-ERREF]; see also section 2.2.1
	// for HRESULT values predefined by the Disk Management Remote Protocol).
	DeletePartitionsForVolume(context.Context, *DeletePartitionsForVolumeRequest) (*DeletePartitionsForVolumeResponse, error)

	// GetMaxAdjustedFreeSpace operation.
	GetMaxAdjustedFreeSpace(context.Context, *GetMaxAdjustedFreeSpaceRequest) (*GetMaxAdjustedFreeSpaceResponse, error)

	// AddMirror operation.
	AddMirror(context.Context, *AddMirrorRequest) (*AddMirrorResponse, error)

	// RemoveMirror operation.
	RemoveMirror(context.Context, *RemoveMirrorRequest) (*RemoveMirrorResponse, error)

	// SplitMirror operation.
	SplitMirror(context.Context, *SplitMirrorRequest) (*SplitMirrorResponse, error)

	// The InitializeDiskEx method initializes a disk for control by the volume manager.
	// This is a synchronous task.
	//
	// Return Values: The method MUST return 0 or a nonerror HRESULT on success, or an implementation-specific
	// nonzero error code on failure (as specified in [MS-ERREF]; see also section 2.2.1
	// for HRESULT values predefined by the Disk Management Remote Protocol).
	InitializeDiskEx(context.Context, *InitializeDiskExRequest) (*InitializeDiskExResponse, error)

	// UninitializeDisk operation.
	UninitializeDisk(context.Context, *UninitializeDiskRequest) (*UninitializeDiskResponse, error)

	// ReConnectDisk operation.
	ReConnectDisk(context.Context, *ReConnectDiskRequest) (*ReConnectDiskResponse, error)

	// ImportDiskGroup operation.
	ImportDiskGroup(context.Context, *ImportDiskGroupRequest) (*ImportDiskGroupResponse, error)

	// DiskMergeQuery operation.
	DiskMergeQuery(context.Context, *DiskMergeQueryRequest) (*DiskMergeQueryResponse, error)

	// DiskMerge operation.
	DiskMerge(context.Context, *DiskMergeRequest) (*DiskMergeResponse, error)

	// ReAttachDisk operation.
	ReAttachDisk(context.Context, *ReAttachDiskRequest) (*ReAttachDiskResponse, error)

	// ReplaceRaid5Column operation.
	ReplaceRAID5Column(context.Context, *ReplaceRAID5ColumnRequest) (*ReplaceRAID5ColumnResponse, error)

	// RestartVolume operation.
	RestartVolume(context.Context, *RestartVolumeRequest) (*RestartVolumeResponse, error)

	// The GetEncapsulateDiskInfoEx method gathers the information needed to convert the
	// specified basic disks to dynamic disks.
	//
	// Return Values: The method MUST return 0 or a nonerror HRESULT on success, or an implementation-specific
	// nonzero error code on failure (as specified in [MS-ERREF]; see also section 2.2.1
	// for HRESULT values predefined by the Disk Management Remote Protocol).
	GetEncapsulateDiskInfoEx(context.Context, *GetEncapsulateDiskInfoExRequest) (*GetEncapsulateDiskInfoExResponse, error)

	// The EncapsulateDiskEx method converts the specified basic disks to dynamic disks.
	// This is a synchronous task.
	//
	// Return Values: The method MUST return 0 or a nonerror HRESULT on success, or an implementation-specific
	// nonzero error code on failure (as specified in [MS-ERREF]; see also section 2.2.1
	// for HRESULT values predefined by the Disk Management Remote Protocol).
	EncapsulateDiskEx(context.Context, *EncapsulateDiskExRequest) (*EncapsulateDiskExResponse, error)

	// QueryChangePartitionNumbers operation.
	QueryChangePartitionNumbers(context.Context, *QueryChangePartitionNumbersRequest) (*QueryChangePartitionNumbersResponse, error)

	// DeletePartitionNumberInfoFromRegistry operation.
	DeletePartitionNumberInfoFromRegistry(context.Context, *DeletePartitionNumberInfoFromRegistryRequest) (*DeletePartitionNumberInfoFromRegistryResponse, error)

	// SetDontShow operation.
	SetDontShow(context.Context, *SetDontShowRequest) (*SetDontShowResponse, error)

	// GetDontShow operation.
	GetDontShow(context.Context, *GetDontShowRequest) (*GetDontShowResponse, error)

	// Reserved0 operation.
	// _

	// Reserved1 operation.
	// _

	// Reserved2 operation.
	// _

	// Reserved3 operation.
	// _

	// Reserved4 operation.
	// _

	// Reserved5 operation.
	// _

	// Reserved6 operation.
	// _

	// Reserved7 operation.
	// _

	// EnumTasks operation.
	EnumTasks(context.Context, *EnumTasksRequest) (*EnumTasksResponse, error)

	// GetTaskDetail operation.
	GetTaskDetail(context.Context, *GetTaskDetailRequest) (*GetTaskDetailResponse, error)

	// AbortTask operation.
	AbortTask(context.Context, *AbortTaskRequest) (*AbortTaskResponse, error)

	// HrGetErrorData operation.
	HResultGetErrorData(context.Context, *HResultGetErrorDataRequest) (*HResultGetErrorDataResponse, error)

	// Initialize operation.
	Initialize(context.Context, *InitializeRequest) (*InitializeResponse, error)

	// Uninitialize operation.
	Uninitialize(context.Context, *UninitializeRequest) (*UninitializeResponse, error)

	// Refresh operation.
	Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error)

	// RescanDisks operation.
	RescanDisks(context.Context, *RescanDisksRequest) (*RescanDisksResponse, error)

	// RefreshFileSys operation.
	RefreshFileSystem(context.Context, *RefreshFileSystemRequest) (*RefreshFileSystemResponse, error)

	// SecureSystemPartition operation.
	SecureSystemPartition(context.Context, *SecureSystemPartitionRequest) (*SecureSystemPartitionResponse, error)

	// ShutDownSystem operation.
	ShutDownSystem(context.Context, *ShutDownSystemRequest) (*ShutDownSystemResponse, error)

	// EnumAccessPath operation.
	EnumAccessPath(context.Context, *EnumAccessPathRequest) (*EnumAccessPathResponse, error)

	// EnumAccessPathForVolume operation.
	EnumAccessPathForVolume(context.Context, *EnumAccessPathForVolumeRequest) (*EnumAccessPathForVolumeResponse, error)

	// AddAccessPath operation.
	AddAccessPath(context.Context, *AddAccessPathRequest) (*AddAccessPathResponse, error)

	// DeleteAccessPath operation.
	DeleteAccessPath(context.Context, *DeleteAccessPathRequest) (*DeleteAccessPathResponse, error)
}

func RegisterVolumeClient3Server(conn dcerpc.Conn, o VolumeClient3Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewVolumeClient3ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(VolumeClient3SyntaxV0_0))...)
}

func NewVolumeClient3ServerHandle(o VolumeClient3Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return VolumeClient3ServerHandle(ctx, o, opNum, r)
	}
}

func VolumeClient3ServerHandle(ctx context.Context, o VolumeClient3Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // EnumDisksEx
		op := &xxx_EnumDisksExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumDisksExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumDisksEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // EnumDiskRegionsEx
		op := &xxx_EnumDiskRegionsExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumDiskRegionsExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumDiskRegionsEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // CreatePartition
		op := &xxx_CreatePartitionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreatePartitionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreatePartition(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // CreatePartitionAssignAndFormat
		op := &xxx_CreatePartitionAssignAndFormatOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreatePartitionAssignAndFormatRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreatePartitionAssignAndFormat(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // CreatePartitionAssignAndFormatEx
		op := &xxx_CreatePartitionAssignAndFormatExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreatePartitionAssignAndFormatExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreatePartitionAssignAndFormatEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // DeletePartition
		op := &xxx_DeletePartitionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeletePartitionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeletePartition(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // InitializeDiskStyle
		op := &xxx_InitializeDiskStyleOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &InitializeDiskStyleRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.InitializeDiskStyle(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // MarkActivePartition
		op := &xxx_MarkActivePartitionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MarkActivePartitionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MarkActivePartition(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // Eject
		op := &xxx_EjectOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EjectRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Eject(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // Reserved_Opnum12
		// Reserved_Opnum12
		return nil, nil
	case 13: // FTEnumVolumes
		op := &xxx_FTEnumVolumesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FTEnumVolumesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.FTEnumVolumes(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // FTEnumLogicalDiskMembers
		op := &xxx_FTEnumLogicalDiskMembersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FTEnumLogicalDiskMembersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.FTEnumLogicalDiskMembers(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // FTDeleteVolume
		op := &xxx_FTDeleteVolumeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FTDeleteVolumeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.FTDeleteVolume(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // FTBreakMirror
		op := &xxx_FTBreakMirrorOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FTBreakMirrorRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.FTBreakMirror(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // FTResyncMirror
		op := &xxx_FTResyncMirrorOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FTResyncMirrorRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.FTResyncMirror(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // FTRegenerateParityStripe
		op := &xxx_FTRegenerateParityStripeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FTRegenerateParityStripeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.FTRegenerateParityStripe(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // FTReplaceMirrorPartition
		op := &xxx_FTReplaceMirrorPartitionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FTReplaceMirrorPartitionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.FTReplaceMirrorPartition(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // FTReplaceParityStripePartition
		op := &xxx_FTReplaceParityStripePartitionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FTReplaceParityStripePartitionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.FTReplaceParityStripePartition(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // EnumDriveLetters
		op := &xxx_EnumDriveLettersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumDriveLettersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumDriveLetters(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // AssignDriveLetter
		op := &xxx_AssignDriveLetterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AssignDriveLetterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AssignDriveLetter(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // FreeDriveLetter
		op := &xxx_FreeDriveLetterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FreeDriveLetterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.FreeDriveLetter(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // EnumLocalFileSystems
		op := &xxx_EnumLocalFileSystemsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumLocalFileSystemsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumLocalFileSystems(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // GetInstalledFileSystems
		op := &xxx_GetInstalledFileSystemsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetInstalledFileSystemsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetInstalledFileSystems(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 26: // Format
		op := &xxx_FormatOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FormatRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Format(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // EnumVolumes
		op := &xxx_EnumVolumesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumVolumesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumVolumes(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 28: // EnumVolumeMembers
		op := &xxx_EnumVolumeMembersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumVolumeMembersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumVolumeMembers(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 29: // CreateVolume
		op := &xxx_CreateVolumeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateVolumeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateVolume(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // CreateVolumeAssignAndFormat
		op := &xxx_CreateVolumeAssignAndFormatOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateVolumeAssignAndFormatRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateVolumeAssignAndFormat(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 31: // CreateVolumeAssignAndFormatEx
		op := &xxx_CreateVolumeAssignAndFormatExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateVolumeAssignAndFormatExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateVolumeAssignAndFormatEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 32: // GetVolumeMountName
		op := &xxx_GetVolumeMountNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetVolumeMountNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetVolumeMountName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 33: // GrowVolume
		op := &xxx_GrowVolumeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GrowVolumeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GrowVolume(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 34: // DeleteVolume
		op := &xxx_DeleteVolumeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteVolumeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteVolume(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 35: // CreatePartitionsForVolume
		op := &xxx_CreatePartitionsForVolumeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreatePartitionsForVolumeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreatePartitionsForVolume(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 36: // DeletePartitionsForVolume
		op := &xxx_DeletePartitionsForVolumeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeletePartitionsForVolumeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeletePartitionsForVolume(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 37: // GetMaxAdjustedFreeSpace
		op := &xxx_GetMaxAdjustedFreeSpaceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMaxAdjustedFreeSpaceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMaxAdjustedFreeSpace(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 38: // AddMirror
		op := &xxx_AddMirrorOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddMirrorRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddMirror(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 39: // RemoveMirror
		op := &xxx_RemoveMirrorOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveMirrorRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveMirror(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 40: // SplitMirror
		op := &xxx_SplitMirrorOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SplitMirrorRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SplitMirror(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 41: // InitializeDiskEx
		op := &xxx_InitializeDiskExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &InitializeDiskExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.InitializeDiskEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 42: // UninitializeDisk
		op := &xxx_UninitializeDiskOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UninitializeDiskRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.UninitializeDisk(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 43: // ReConnectDisk
		op := &xxx_ReConnectDiskOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReConnectDiskRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReConnectDisk(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 44: // ImportDiskGroup
		op := &xxx_ImportDiskGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ImportDiskGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ImportDiskGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 45: // DiskMergeQuery
		op := &xxx_DiskMergeQueryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DiskMergeQueryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DiskMergeQuery(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 46: // DiskMerge
		op := &xxx_DiskMergeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DiskMergeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DiskMerge(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 47: // ReAttachDisk
		op := &xxx_ReAttachDiskOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReAttachDiskRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReAttachDisk(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 48: // ReplaceRaid5Column
		op := &xxx_ReplaceRAID5ColumnOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReplaceRAID5ColumnRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReplaceRAID5Column(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 49: // RestartVolume
		op := &xxx_RestartVolumeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RestartVolumeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RestartVolume(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 50: // GetEncapsulateDiskInfoEx
		op := &xxx_GetEncapsulateDiskInfoExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetEncapsulateDiskInfoExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetEncapsulateDiskInfoEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 51: // EncapsulateDiskEx
		op := &xxx_EncapsulateDiskExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EncapsulateDiskExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EncapsulateDiskEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 52: // QueryChangePartitionNumbers
		op := &xxx_QueryChangePartitionNumbersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryChangePartitionNumbersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryChangePartitionNumbers(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 53: // DeletePartitionNumberInfoFromRegistry
		op := &xxx_DeletePartitionNumberInfoFromRegistryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeletePartitionNumberInfoFromRegistryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeletePartitionNumberInfoFromRegistry(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 54: // SetDontShow
		op := &xxx_SetDontShowOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetDontShowRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetDontShow(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 55: // GetDontShow
		op := &xxx_GetDontShowOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDontShowRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDontShow(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 56: // Reserved0
		// Reserved0
		return nil, nil
	case 57: // Reserved1
		// Reserved1
		return nil, nil
	case 58: // Reserved2
		// Reserved2
		return nil, nil
	case 59: // Reserved3
		// Reserved3
		return nil, nil
	case 60: // Reserved4
		// Reserved4
		return nil, nil
	case 61: // Reserved5
		// Reserved5
		return nil, nil
	case 62: // Reserved6
		// Reserved6
		return nil, nil
	case 63: // Reserved7
		// Reserved7
		return nil, nil
	case 64: // EnumTasks
		op := &xxx_EnumTasksOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumTasksRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumTasks(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 65: // GetTaskDetail
		op := &xxx_GetTaskDetailOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTaskDetailRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTaskDetail(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 66: // AbortTask
		op := &xxx_AbortTaskOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AbortTaskRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AbortTask(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 67: // HrGetErrorData
		op := &xxx_HResultGetErrorDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &HResultGetErrorDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.HResultGetErrorData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 68: // Initialize
		op := &xxx_InitializeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &InitializeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Initialize(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 69: // Uninitialize
		op := &xxx_UninitializeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UninitializeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Uninitialize(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 70: // Refresh
		op := &xxx_RefreshOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RefreshRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Refresh(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 71: // RescanDisks
		op := &xxx_RescanDisksOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RescanDisksRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RescanDisks(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 72: // RefreshFileSys
		op := &xxx_RefreshFileSystemOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RefreshFileSystemRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RefreshFileSystem(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 73: // SecureSystemPartition
		op := &xxx_SecureSystemPartitionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SecureSystemPartitionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SecureSystemPartition(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 74: // ShutDownSystem
		op := &xxx_ShutDownSystemOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ShutDownSystemRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ShutDownSystem(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 75: // EnumAccessPath
		op := &xxx_EnumAccessPathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumAccessPathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumAccessPath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 76: // EnumAccessPathForVolume
		op := &xxx_EnumAccessPathForVolumeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumAccessPathForVolumeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumAccessPathForVolume(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 77: // AddAccessPath
		op := &xxx_AddAccessPathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddAccessPathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddAccessPath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 78: // DeleteAccessPath
		op := &xxx_DeleteAccessPathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteAccessPathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteAccessPath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IVolumeClient3
type UnimplementedVolumeClient3Server struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedVolumeClient3Server) EnumDisksEx(context.Context, *EnumDisksExRequest) (*EnumDisksExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) EnumDiskRegionsEx(context.Context, *EnumDiskRegionsExRequest) (*EnumDiskRegionsExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) CreatePartition(context.Context, *CreatePartitionRequest) (*CreatePartitionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) CreatePartitionAssignAndFormat(context.Context, *CreatePartitionAssignAndFormatRequest) (*CreatePartitionAssignAndFormatResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) CreatePartitionAssignAndFormatEx(context.Context, *CreatePartitionAssignAndFormatExRequest) (*CreatePartitionAssignAndFormatExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) DeletePartition(context.Context, *DeletePartitionRequest) (*DeletePartitionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) InitializeDiskStyle(context.Context, *InitializeDiskStyleRequest) (*InitializeDiskStyleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) MarkActivePartition(context.Context, *MarkActivePartitionRequest) (*MarkActivePartitionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) Eject(context.Context, *EjectRequest) (*EjectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) FTEnumVolumes(context.Context, *FTEnumVolumesRequest) (*FTEnumVolumesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) FTEnumLogicalDiskMembers(context.Context, *FTEnumLogicalDiskMembersRequest) (*FTEnumLogicalDiskMembersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) FTDeleteVolume(context.Context, *FTDeleteVolumeRequest) (*FTDeleteVolumeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) FTBreakMirror(context.Context, *FTBreakMirrorRequest) (*FTBreakMirrorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) FTResyncMirror(context.Context, *FTResyncMirrorRequest) (*FTResyncMirrorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) FTRegenerateParityStripe(context.Context, *FTRegenerateParityStripeRequest) (*FTRegenerateParityStripeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) FTReplaceMirrorPartition(context.Context, *FTReplaceMirrorPartitionRequest) (*FTReplaceMirrorPartitionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) FTReplaceParityStripePartition(context.Context, *FTReplaceParityStripePartitionRequest) (*FTReplaceParityStripePartitionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) EnumDriveLetters(context.Context, *EnumDriveLettersRequest) (*EnumDriveLettersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) AssignDriveLetter(context.Context, *AssignDriveLetterRequest) (*AssignDriveLetterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) FreeDriveLetter(context.Context, *FreeDriveLetterRequest) (*FreeDriveLetterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) EnumLocalFileSystems(context.Context, *EnumLocalFileSystemsRequest) (*EnumLocalFileSystemsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) GetInstalledFileSystems(context.Context, *GetInstalledFileSystemsRequest) (*GetInstalledFileSystemsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) Format(context.Context, *FormatRequest) (*FormatResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) EnumVolumes(context.Context, *EnumVolumesRequest) (*EnumVolumesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) EnumVolumeMembers(context.Context, *EnumVolumeMembersRequest) (*EnumVolumeMembersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) CreateVolume(context.Context, *CreateVolumeRequest) (*CreateVolumeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) CreateVolumeAssignAndFormat(context.Context, *CreateVolumeAssignAndFormatRequest) (*CreateVolumeAssignAndFormatResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) CreateVolumeAssignAndFormatEx(context.Context, *CreateVolumeAssignAndFormatExRequest) (*CreateVolumeAssignAndFormatExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) GetVolumeMountName(context.Context, *GetVolumeMountNameRequest) (*GetVolumeMountNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) GrowVolume(context.Context, *GrowVolumeRequest) (*GrowVolumeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) DeleteVolume(context.Context, *DeleteVolumeRequest) (*DeleteVolumeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) CreatePartitionsForVolume(context.Context, *CreatePartitionsForVolumeRequest) (*CreatePartitionsForVolumeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) DeletePartitionsForVolume(context.Context, *DeletePartitionsForVolumeRequest) (*DeletePartitionsForVolumeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) GetMaxAdjustedFreeSpace(context.Context, *GetMaxAdjustedFreeSpaceRequest) (*GetMaxAdjustedFreeSpaceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) AddMirror(context.Context, *AddMirrorRequest) (*AddMirrorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) RemoveMirror(context.Context, *RemoveMirrorRequest) (*RemoveMirrorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) SplitMirror(context.Context, *SplitMirrorRequest) (*SplitMirrorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) InitializeDiskEx(context.Context, *InitializeDiskExRequest) (*InitializeDiskExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) UninitializeDisk(context.Context, *UninitializeDiskRequest) (*UninitializeDiskResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) ReConnectDisk(context.Context, *ReConnectDiskRequest) (*ReConnectDiskResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) ImportDiskGroup(context.Context, *ImportDiskGroupRequest) (*ImportDiskGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) DiskMergeQuery(context.Context, *DiskMergeQueryRequest) (*DiskMergeQueryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) DiskMerge(context.Context, *DiskMergeRequest) (*DiskMergeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) ReAttachDisk(context.Context, *ReAttachDiskRequest) (*ReAttachDiskResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) ReplaceRAID5Column(context.Context, *ReplaceRAID5ColumnRequest) (*ReplaceRAID5ColumnResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) RestartVolume(context.Context, *RestartVolumeRequest) (*RestartVolumeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) GetEncapsulateDiskInfoEx(context.Context, *GetEncapsulateDiskInfoExRequest) (*GetEncapsulateDiskInfoExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) EncapsulateDiskEx(context.Context, *EncapsulateDiskExRequest) (*EncapsulateDiskExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) QueryChangePartitionNumbers(context.Context, *QueryChangePartitionNumbersRequest) (*QueryChangePartitionNumbersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) DeletePartitionNumberInfoFromRegistry(context.Context, *DeletePartitionNumberInfoFromRegistryRequest) (*DeletePartitionNumberInfoFromRegistryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) SetDontShow(context.Context, *SetDontShowRequest) (*SetDontShowResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) GetDontShow(context.Context, *GetDontShowRequest) (*GetDontShowResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) EnumTasks(context.Context, *EnumTasksRequest) (*EnumTasksResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) GetTaskDetail(context.Context, *GetTaskDetailRequest) (*GetTaskDetailResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) AbortTask(context.Context, *AbortTaskRequest) (*AbortTaskResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) HResultGetErrorData(context.Context, *HResultGetErrorDataRequest) (*HResultGetErrorDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) Initialize(context.Context, *InitializeRequest) (*InitializeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) Uninitialize(context.Context, *UninitializeRequest) (*UninitializeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) RescanDisks(context.Context, *RescanDisksRequest) (*RescanDisksResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) RefreshFileSystem(context.Context, *RefreshFileSystemRequest) (*RefreshFileSystemResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) SecureSystemPartition(context.Context, *SecureSystemPartitionRequest) (*SecureSystemPartitionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) ShutDownSystem(context.Context, *ShutDownSystemRequest) (*ShutDownSystemResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) EnumAccessPath(context.Context, *EnumAccessPathRequest) (*EnumAccessPathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) EnumAccessPathForVolume(context.Context, *EnumAccessPathForVolumeRequest) (*EnumAccessPathForVolumeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) AddAccessPath(context.Context, *AddAccessPathRequest) (*AddAccessPathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient3Server) DeleteAccessPath(context.Context, *DeleteAccessPathRequest) (*DeleteAccessPathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ VolumeClient3Server = (*UnimplementedVolumeClient3Server)(nil)
