package ivolumeclient

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
	_ = iunknown.GoPackage
)

// IVolumeClient server interface.
type VolumeClientServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The EnumDisks method enumerates the server's mass storage devices.
	//
	// Return Values: The method MUST return 0 or a nonerror HRESULT on success, or an implementation-specific
	// nonzero error code on failure (as specified in [MS-ERREF] section 2.1; see also section
	// 2.2.1 for HRESULT values predefined by the Disk Management Remote Protocol).
	EnumDisks(context.Context, *EnumDisksRequest) (*EnumDisksResponse, error)

	// The EnumDiskRegions method enumerates all used and free regions of a specified disk.
	//
	// Return Values: The method MUST return 0 or a nonerror HRESULT on success, or an implementation-specific
	// nonzero error code on failure (as specified in [MS-ERREF] section 2.1; see also section
	// 2.2.1 for HRESULT values predefined by the Disk Management Remote Protocol).
	EnumDiskRegions(context.Context, *EnumDiskRegionsRequest) (*EnumDiskRegionsResponse, error)

	// CreatePartition operation.
	CreatePartition(context.Context, *CreatePartitionRequest) (*CreatePartitionResponse, error)

	// CreatePartitionAssignAndFormat operation.
	CreatePartitionAssignAndFormat(context.Context, *CreatePartitionAssignAndFormatRequest) (*CreatePartitionAssignAndFormatResponse, error)

	// CreatePartitionAssignAndFormatEx operation.
	CreatePartitionAssignAndFormatEx(context.Context, *CreatePartitionAssignAndFormatExRequest) (*CreatePartitionAssignAndFormatExResponse, error)

	// DeletePartition operation.
	DeletePartition(context.Context, *DeletePartitionRequest) (*DeletePartitionResponse, error)

	// The WriteSignature method writes a disk signature to a specified disk. This is a
	// synchronous task.
	//
	// Return Values: The method MUST return 0 or a nonerror HRESULT on success, or an implementation-specific
	// nonzero error code on failure (as specified in [MS-ERREF] section 2.1; see also section
	// 2.2.1 for HRESULT values predefined by the Disk Management Remote Protocol).
	WriteSignature(context.Context, *WriteSignatureRequest) (*WriteSignatureResponse, error)

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

	// Reserved27 operation.
	// _

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

	// AddMirror operation.
	AddMirror(context.Context, *AddMirrorRequest) (*AddMirrorResponse, error)

	// RemoveMirror operation.
	RemoveMirror(context.Context, *RemoveMirrorRequest) (*RemoveMirrorResponse, error)

	// SplitMirror operation.
	SplitMirror(context.Context, *SplitMirrorRequest) (*SplitMirrorResponse, error)

	// The InitializeDisk method converts an uninitialized disk into a dynamic disk. This
	// is a synchronous task.
	//
	// Return Values: The method MUST return 0 or a nonerror HRESULT on success, or an implementation-specific
	// nonzero error code on failure (as specified in [MS-ERREF]; see also section 2.2.1
	// for HRESULT values predefined by the Disk Management Remote Protocol).
	InitializeDisk(context.Context, *InitializeDiskRequest) (*InitializeDiskResponse, error)

	// UninitializeDisk operation.
	UninitializeDisk(context.Context, *UninitializeDiskRequest) (*UninitializeDiskResponse, error)

	// ReConnectDisk operation.
	ReConnectDisk(context.Context, *ReConnectDiskRequest) (*ReConnectDiskResponse, error)

	// Reserved_Opnum42 operation.
	// ReservedOpnum42

	// ImportDiskGroup operation.
	ImportDiskGroup(context.Context, *ImportDiskGroupRequest) (*ImportDiskGroupResponse, error)

	// DiskMergeQuery operation.
	DiskMergeQuery(context.Context, *DiskMergeQueryRequest) (*DiskMergeQueryResponse, error)

	// DiskMerge operation.
	DiskMerge(context.Context, *DiskMergeRequest) (*DiskMergeResponse, error)

	// Reserved_Opnum46 operation.
	// ReservedOpnum46

	// ReAttachDisk operation.
	ReAttachDisk(context.Context, *ReAttachDiskRequest) (*ReAttachDiskResponse, error)

	// Reserved_Opnum48 operation.
	// ReservedOpnum48

	// Reserved_Opnum49 operation.
	// ReservedOpnum49

	// Reserved_Opnum50 operation.
	// ReservedOpnum50

	// ReplaceRaid5Column operation.
	ReplaceRAID5Column(context.Context, *ReplaceRAID5ColumnRequest) (*ReplaceRAID5ColumnResponse, error)

	// RestartVolume operation.
	RestartVolume(context.Context, *RestartVolumeRequest) (*RestartVolumeResponse, error)

	// The GetEncapsulateDiskInfo method gathers the information needed to convert the specified
	// basic disks to dynamic disks. This is a synchronous task.
	//
	// Return Values: The method MUST return 0 or a nonerror HRESULT on success, or an implementation-specific
	// nonzero error code on failure (as specified in [MS-ERREF]; see also section 2.2.1
	// for HRESULT values predefined by the Disk Management Remote Protocol).
	GetEncapsulateDiskInfo(context.Context, *GetEncapsulateDiskInfoRequest) (*GetEncapsulateDiskInfoResponse, error)

	// The EncapsulateDisk method converts the specified basic disks to dynamic disks. This
	// is a synchronous task.
	//
	// Return Values: The method MUST return 0 or a nonerror HRESULT on success, or an implementation-specific
	// nonzero error code on failure (as specified in [MS-ERREF]; see also section 2.2.1
	// for HRESULT values predefined by the Disk Management Remote Protocol).
	EncapsulateDisk(context.Context, *EncapsulateDiskRequest) (*EncapsulateDiskResponse, error)

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

func RegisterVolumeClientServer(conn dcerpc.Conn, o VolumeClientServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewVolumeClientServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(VolumeClientSyntaxV0_0))...)
}

func NewVolumeClientServerHandle(o VolumeClientServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return VolumeClientServerHandle(ctx, o, opNum, r)
	}
}

func VolumeClientServerHandle(ctx context.Context, o VolumeClientServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // EnumDisks
		in := &EnumDisksRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumDisks(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // EnumDiskRegions
		in := &EnumDiskRegionsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumDiskRegions(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // CreatePartition
		in := &CreatePartitionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreatePartition(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // CreatePartitionAssignAndFormat
		in := &CreatePartitionAssignAndFormatRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreatePartitionAssignAndFormat(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // CreatePartitionAssignAndFormatEx
		in := &CreatePartitionAssignAndFormatExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreatePartitionAssignAndFormatEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // DeletePartition
		in := &DeletePartitionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeletePartition(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // WriteSignature
		in := &WriteSignatureRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.WriteSignature(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // MarkActivePartition
		in := &MarkActivePartitionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.MarkActivePartition(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // Eject
		in := &EjectRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Eject(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // Reserved_Opnum12
		// Reserved_Opnum12
		return nil, nil
	case 13: // FTEnumVolumes
		in := &FTEnumVolumesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FTEnumVolumes(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // FTEnumLogicalDiskMembers
		in := &FTEnumLogicalDiskMembersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FTEnumLogicalDiskMembers(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // FTDeleteVolume
		in := &FTDeleteVolumeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FTDeleteVolume(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // FTBreakMirror
		in := &FTBreakMirrorRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FTBreakMirror(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // FTResyncMirror
		in := &FTResyncMirrorRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FTResyncMirror(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // FTRegenerateParityStripe
		in := &FTRegenerateParityStripeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FTRegenerateParityStripe(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // FTReplaceMirrorPartition
		in := &FTReplaceMirrorPartitionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FTReplaceMirrorPartition(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // FTReplaceParityStripePartition
		in := &FTReplaceParityStripePartitionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FTReplaceParityStripePartition(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // EnumDriveLetters
		in := &EnumDriveLettersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumDriveLetters(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 22: // AssignDriveLetter
		in := &AssignDriveLetterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AssignDriveLetter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // FreeDriveLetter
		in := &FreeDriveLetterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FreeDriveLetter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // EnumLocalFileSystems
		in := &EnumLocalFileSystemsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumLocalFileSystems(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // GetInstalledFileSystems
		in := &GetInstalledFileSystemsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetInstalledFileSystems(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 26: // Format
		in := &FormatRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Format(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 27: // Reserved27
		// Reserved27
		return nil, nil
	case 28: // EnumVolumes
		in := &EnumVolumesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumVolumes(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 29: // EnumVolumeMembers
		in := &EnumVolumeMembersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumVolumeMembers(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 30: // CreateVolume
		in := &CreateVolumeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateVolume(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 31: // CreateVolumeAssignAndFormat
		in := &CreateVolumeAssignAndFormatRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateVolumeAssignAndFormat(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 32: // CreateVolumeAssignAndFormatEx
		in := &CreateVolumeAssignAndFormatExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateVolumeAssignAndFormatEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 33: // GetVolumeMountName
		in := &GetVolumeMountNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetVolumeMountName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 34: // GrowVolume
		in := &GrowVolumeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GrowVolume(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 35: // DeleteVolume
		in := &DeleteVolumeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteVolume(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 36: // AddMirror
		in := &AddMirrorRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddMirror(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 37: // RemoveMirror
		in := &RemoveMirrorRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoveMirror(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 38: // SplitMirror
		in := &SplitMirrorRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SplitMirror(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 39: // InitializeDisk
		in := &InitializeDiskRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.InitializeDisk(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 40: // UninitializeDisk
		in := &UninitializeDiskRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.UninitializeDisk(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 41: // ReConnectDisk
		in := &ReConnectDiskRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ReConnectDisk(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 42: // Reserved_Opnum42
		// Reserved_Opnum42
		return nil, nil
	case 43: // ImportDiskGroup
		in := &ImportDiskGroupRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ImportDiskGroup(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 44: // DiskMergeQuery
		in := &DiskMergeQueryRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DiskMergeQuery(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 45: // DiskMerge
		in := &DiskMergeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DiskMerge(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 46: // Reserved_Opnum46
		// Reserved_Opnum46
		return nil, nil
	case 47: // ReAttachDisk
		in := &ReAttachDiskRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ReAttachDisk(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 48: // Reserved_Opnum48
		// Reserved_Opnum48
		return nil, nil
	case 49: // Reserved_Opnum49
		// Reserved_Opnum49
		return nil, nil
	case 50: // Reserved_Opnum50
		// Reserved_Opnum50
		return nil, nil
	case 51: // ReplaceRaid5Column
		in := &ReplaceRAID5ColumnRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ReplaceRAID5Column(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 52: // RestartVolume
		in := &RestartVolumeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RestartVolume(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 53: // GetEncapsulateDiskInfo
		in := &GetEncapsulateDiskInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetEncapsulateDiskInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 54: // EncapsulateDisk
		in := &EncapsulateDiskRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EncapsulateDisk(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 55: // QueryChangePartitionNumbers
		in := &QueryChangePartitionNumbersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryChangePartitionNumbers(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 56: // DeletePartitionNumberInfoFromRegistry
		in := &DeletePartitionNumberInfoFromRegistryRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeletePartitionNumberInfoFromRegistry(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 57: // SetDontShow
		in := &SetDontShowRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetDontShow(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 58: // GetDontShow
		in := &GetDontShowRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDontShow(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 59: // Reserved0
		// Reserved0
		return nil, nil
	case 60: // Reserved1
		// Reserved1
		return nil, nil
	case 61: // Reserved2
		// Reserved2
		return nil, nil
	case 62: // Reserved3
		// Reserved3
		return nil, nil
	case 63: // Reserved4
		// Reserved4
		return nil, nil
	case 64: // Reserved5
		// Reserved5
		return nil, nil
	case 65: // Reserved6
		// Reserved6
		return nil, nil
	case 66: // Reserved7
		// Reserved7
		return nil, nil
	case 67: // EnumTasks
		in := &EnumTasksRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumTasks(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 68: // GetTaskDetail
		in := &GetTaskDetailRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTaskDetail(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 69: // AbortTask
		in := &AbortTaskRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AbortTask(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 70: // HrGetErrorData
		in := &HResultGetErrorDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.HResultGetErrorData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 71: // Initialize
		in := &InitializeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Initialize(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 72: // Uninitialize
		in := &UninitializeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Uninitialize(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 73: // Refresh
		in := &RefreshRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Refresh(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 74: // RescanDisks
		in := &RescanDisksRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RescanDisks(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 75: // RefreshFileSys
		in := &RefreshFileSystemRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RefreshFileSystem(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 76: // SecureSystemPartition
		in := &SecureSystemPartitionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SecureSystemPartition(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 77: // ShutDownSystem
		in := &ShutDownSystemRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ShutDownSystem(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 78: // EnumAccessPath
		in := &EnumAccessPathRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumAccessPath(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 79: // EnumAccessPathForVolume
		in := &EnumAccessPathForVolumeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumAccessPathForVolume(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 80: // AddAccessPath
		in := &AddAccessPathRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddAccessPath(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 81: // DeleteAccessPath
		in := &DeleteAccessPathRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteAccessPath(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
