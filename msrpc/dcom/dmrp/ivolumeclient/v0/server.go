package ivolumeclient

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
		op := &xxx_EnumDisksOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumDisksRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumDisks(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // EnumDiskRegions
		op := &xxx_EnumDiskRegionsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumDiskRegionsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumDiskRegions(ctx, req)
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
	case 9: // WriteSignature
		op := &xxx_WriteSignatureOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &WriteSignatureRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.WriteSignature(ctx, req)
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
	case 27: // Reserved27
		// Reserved27
		return nil, nil
	case 28: // EnumVolumes
		op := &xxx_EnumVolumesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumVolumesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumVolumes(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 29: // EnumVolumeMembers
		op := &xxx_EnumVolumeMembersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumVolumeMembersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumVolumeMembers(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // CreateVolume
		op := &xxx_CreateVolumeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateVolumeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateVolume(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 31: // CreateVolumeAssignAndFormat
		op := &xxx_CreateVolumeAssignAndFormatOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateVolumeAssignAndFormatRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateVolumeAssignAndFormat(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 32: // CreateVolumeAssignAndFormatEx
		op := &xxx_CreateVolumeAssignAndFormatExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateVolumeAssignAndFormatExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateVolumeAssignAndFormatEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 33: // GetVolumeMountName
		op := &xxx_GetVolumeMountNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetVolumeMountNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetVolumeMountName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 34: // GrowVolume
		op := &xxx_GrowVolumeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GrowVolumeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GrowVolume(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 35: // DeleteVolume
		op := &xxx_DeleteVolumeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteVolumeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteVolume(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 36: // AddMirror
		op := &xxx_AddMirrorOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddMirrorRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddMirror(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 37: // RemoveMirror
		op := &xxx_RemoveMirrorOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveMirrorRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveMirror(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 38: // SplitMirror
		op := &xxx_SplitMirrorOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SplitMirrorRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SplitMirror(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 39: // InitializeDisk
		op := &xxx_InitializeDiskOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &InitializeDiskRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.InitializeDisk(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 40: // UninitializeDisk
		op := &xxx_UninitializeDiskOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UninitializeDiskRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.UninitializeDisk(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 41: // ReConnectDisk
		op := &xxx_ReConnectDiskOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReConnectDiskRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReConnectDisk(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 42: // Reserved_Opnum42
		// Reserved_Opnum42
		return nil, nil
	case 43: // ImportDiskGroup
		op := &xxx_ImportDiskGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ImportDiskGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ImportDiskGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 44: // DiskMergeQuery
		op := &xxx_DiskMergeQueryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DiskMergeQueryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DiskMergeQuery(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 45: // DiskMerge
		op := &xxx_DiskMergeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DiskMergeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DiskMerge(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 46: // Reserved_Opnum46
		// Reserved_Opnum46
		return nil, nil
	case 47: // ReAttachDisk
		op := &xxx_ReAttachDiskOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReAttachDiskRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReAttachDisk(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
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
		op := &xxx_ReplaceRAID5ColumnOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReplaceRAID5ColumnRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReplaceRAID5Column(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 52: // RestartVolume
		op := &xxx_RestartVolumeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RestartVolumeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RestartVolume(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 53: // GetEncapsulateDiskInfo
		op := &xxx_GetEncapsulateDiskInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetEncapsulateDiskInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetEncapsulateDiskInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 54: // EncapsulateDisk
		op := &xxx_EncapsulateDiskOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EncapsulateDiskRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EncapsulateDisk(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 55: // QueryChangePartitionNumbers
		op := &xxx_QueryChangePartitionNumbersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryChangePartitionNumbersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryChangePartitionNumbers(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 56: // DeletePartitionNumberInfoFromRegistry
		op := &xxx_DeletePartitionNumberInfoFromRegistryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeletePartitionNumberInfoFromRegistryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeletePartitionNumberInfoFromRegistry(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 57: // SetDontShow
		op := &xxx_SetDontShowOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetDontShowRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetDontShow(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 58: // GetDontShow
		op := &xxx_GetDontShowOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDontShowRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDontShow(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
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
		op := &xxx_EnumTasksOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumTasksRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumTasks(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 68: // GetTaskDetail
		op := &xxx_GetTaskDetailOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTaskDetailRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTaskDetail(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 69: // AbortTask
		op := &xxx_AbortTaskOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AbortTaskRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AbortTask(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 70: // HrGetErrorData
		op := &xxx_HResultGetErrorDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &HResultGetErrorDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.HResultGetErrorData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 71: // Initialize
		op := &xxx_InitializeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &InitializeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Initialize(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 72: // Uninitialize
		op := &xxx_UninitializeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UninitializeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Uninitialize(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 73: // Refresh
		op := &xxx_RefreshOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RefreshRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Refresh(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 74: // RescanDisks
		op := &xxx_RescanDisksOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RescanDisksRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RescanDisks(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 75: // RefreshFileSys
		op := &xxx_RefreshFileSystemOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RefreshFileSystemRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RefreshFileSystem(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 76: // SecureSystemPartition
		op := &xxx_SecureSystemPartitionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SecureSystemPartitionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SecureSystemPartition(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 77: // ShutDownSystem
		op := &xxx_ShutDownSystemOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ShutDownSystemRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ShutDownSystem(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 78: // EnumAccessPath
		op := &xxx_EnumAccessPathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumAccessPathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumAccessPath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 79: // EnumAccessPathForVolume
		op := &xxx_EnumAccessPathForVolumeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumAccessPathForVolumeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumAccessPathForVolume(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 80: // AddAccessPath
		op := &xxx_AddAccessPathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddAccessPathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddAccessPath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 81: // DeleteAccessPath
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

// Unimplemented IVolumeClient
type UnimplementedVolumeClientServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedVolumeClientServer) EnumDisks(context.Context, *EnumDisksRequest) (*EnumDisksResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) EnumDiskRegions(context.Context, *EnumDiskRegionsRequest) (*EnumDiskRegionsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) CreatePartition(context.Context, *CreatePartitionRequest) (*CreatePartitionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) CreatePartitionAssignAndFormat(context.Context, *CreatePartitionAssignAndFormatRequest) (*CreatePartitionAssignAndFormatResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) CreatePartitionAssignAndFormatEx(context.Context, *CreatePartitionAssignAndFormatExRequest) (*CreatePartitionAssignAndFormatExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) DeletePartition(context.Context, *DeletePartitionRequest) (*DeletePartitionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) WriteSignature(context.Context, *WriteSignatureRequest) (*WriteSignatureResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) MarkActivePartition(context.Context, *MarkActivePartitionRequest) (*MarkActivePartitionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) Eject(context.Context, *EjectRequest) (*EjectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) FTEnumVolumes(context.Context, *FTEnumVolumesRequest) (*FTEnumVolumesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) FTEnumLogicalDiskMembers(context.Context, *FTEnumLogicalDiskMembersRequest) (*FTEnumLogicalDiskMembersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) FTDeleteVolume(context.Context, *FTDeleteVolumeRequest) (*FTDeleteVolumeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) FTBreakMirror(context.Context, *FTBreakMirrorRequest) (*FTBreakMirrorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) FTResyncMirror(context.Context, *FTResyncMirrorRequest) (*FTResyncMirrorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) FTRegenerateParityStripe(context.Context, *FTRegenerateParityStripeRequest) (*FTRegenerateParityStripeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) FTReplaceMirrorPartition(context.Context, *FTReplaceMirrorPartitionRequest) (*FTReplaceMirrorPartitionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) FTReplaceParityStripePartition(context.Context, *FTReplaceParityStripePartitionRequest) (*FTReplaceParityStripePartitionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) EnumDriveLetters(context.Context, *EnumDriveLettersRequest) (*EnumDriveLettersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) AssignDriveLetter(context.Context, *AssignDriveLetterRequest) (*AssignDriveLetterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) FreeDriveLetter(context.Context, *FreeDriveLetterRequest) (*FreeDriveLetterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) EnumLocalFileSystems(context.Context, *EnumLocalFileSystemsRequest) (*EnumLocalFileSystemsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) GetInstalledFileSystems(context.Context, *GetInstalledFileSystemsRequest) (*GetInstalledFileSystemsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) Format(context.Context, *FormatRequest) (*FormatResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) EnumVolumes(context.Context, *EnumVolumesRequest) (*EnumVolumesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) EnumVolumeMembers(context.Context, *EnumVolumeMembersRequest) (*EnumVolumeMembersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) CreateVolume(context.Context, *CreateVolumeRequest) (*CreateVolumeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) CreateVolumeAssignAndFormat(context.Context, *CreateVolumeAssignAndFormatRequest) (*CreateVolumeAssignAndFormatResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) CreateVolumeAssignAndFormatEx(context.Context, *CreateVolumeAssignAndFormatExRequest) (*CreateVolumeAssignAndFormatExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) GetVolumeMountName(context.Context, *GetVolumeMountNameRequest) (*GetVolumeMountNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) GrowVolume(context.Context, *GrowVolumeRequest) (*GrowVolumeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) DeleteVolume(context.Context, *DeleteVolumeRequest) (*DeleteVolumeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) AddMirror(context.Context, *AddMirrorRequest) (*AddMirrorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) RemoveMirror(context.Context, *RemoveMirrorRequest) (*RemoveMirrorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) SplitMirror(context.Context, *SplitMirrorRequest) (*SplitMirrorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) InitializeDisk(context.Context, *InitializeDiskRequest) (*InitializeDiskResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) UninitializeDisk(context.Context, *UninitializeDiskRequest) (*UninitializeDiskResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) ReConnectDisk(context.Context, *ReConnectDiskRequest) (*ReConnectDiskResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) ImportDiskGroup(context.Context, *ImportDiskGroupRequest) (*ImportDiskGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) DiskMergeQuery(context.Context, *DiskMergeQueryRequest) (*DiskMergeQueryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) DiskMerge(context.Context, *DiskMergeRequest) (*DiskMergeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) ReAttachDisk(context.Context, *ReAttachDiskRequest) (*ReAttachDiskResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) ReplaceRAID5Column(context.Context, *ReplaceRAID5ColumnRequest) (*ReplaceRAID5ColumnResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) RestartVolume(context.Context, *RestartVolumeRequest) (*RestartVolumeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) GetEncapsulateDiskInfo(context.Context, *GetEncapsulateDiskInfoRequest) (*GetEncapsulateDiskInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) EncapsulateDisk(context.Context, *EncapsulateDiskRequest) (*EncapsulateDiskResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) QueryChangePartitionNumbers(context.Context, *QueryChangePartitionNumbersRequest) (*QueryChangePartitionNumbersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) DeletePartitionNumberInfoFromRegistry(context.Context, *DeletePartitionNumberInfoFromRegistryRequest) (*DeletePartitionNumberInfoFromRegistryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) SetDontShow(context.Context, *SetDontShowRequest) (*SetDontShowResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) GetDontShow(context.Context, *GetDontShowRequest) (*GetDontShowResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) EnumTasks(context.Context, *EnumTasksRequest) (*EnumTasksResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) GetTaskDetail(context.Context, *GetTaskDetailRequest) (*GetTaskDetailResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) AbortTask(context.Context, *AbortTaskRequest) (*AbortTaskResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) HResultGetErrorData(context.Context, *HResultGetErrorDataRequest) (*HResultGetErrorDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) Initialize(context.Context, *InitializeRequest) (*InitializeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) Uninitialize(context.Context, *UninitializeRequest) (*UninitializeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) RescanDisks(context.Context, *RescanDisksRequest) (*RescanDisksResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) RefreshFileSystem(context.Context, *RefreshFileSystemRequest) (*RefreshFileSystemResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) SecureSystemPartition(context.Context, *SecureSystemPartitionRequest) (*SecureSystemPartitionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) ShutDownSystem(context.Context, *ShutDownSystemRequest) (*ShutDownSystemResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) EnumAccessPath(context.Context, *EnumAccessPathRequest) (*EnumAccessPathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) EnumAccessPathForVolume(context.Context, *EnumAccessPathForVolumeRequest) (*EnumAccessPathForVolumeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) AddAccessPath(context.Context, *AddAccessPathRequest) (*AddAccessPathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClientServer) DeleteAccessPath(context.Context, *DeleteAccessPathRequest) (*DeleteAccessPathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ VolumeClientServer = (*UnimplementedVolumeClientServer)(nil)
