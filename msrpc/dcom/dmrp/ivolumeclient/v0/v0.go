package ivolumeclient

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	dmrp "github.com/oiweiwei/go-msrpc/msrpc/dcom/dmrp"
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
	_ = dcom.GoPackage
	_ = iunknown.GoPackage
	_ = dmrp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/dmrp"
)

var (
	// IVolumeClient interface identifier d2d79df5-3400-11d0-b40b-00aa005ff586
	VolumeClientIID = &dcom.IID{Data1: 0xd2d79df5, Data2: 0x3400, Data3: 0x11d0, Data4: []byte{0xb4, 0x0b, 0x00, 0xaa, 0x00, 0x5f, 0xf5, 0x86}}
	// Syntax UUID
	VolumeClientSyntaxUUID = &uuid.UUID{TimeLow: 0xd2d79df5, TimeMid: 0x3400, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xb4, ClockSeqLow: 0xb, Node: [6]uint8{0x0, 0xaa, 0x0, 0x5f, 0xf5, 0x86}}
	// Syntax ID
	VolumeClientSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: VolumeClientSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVolumeClient interface.
type VolumeClientClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The EnumDisks method enumerates the server's mass storage devices.
	//
	// Return Values: The method MUST return 0 or a nonerror HRESULT on success, or an implementation-specific
	// nonzero error code on failure (as specified in [MS-ERREF] section 2.1; see also section
	// 2.2.1 for HRESULT values predefined by the Disk Management Remote Protocol).
	EnumDisks(context.Context, *EnumDisksRequest, ...dcerpc.CallOption) (*EnumDisksResponse, error)

	// The EnumDiskRegions method enumerates all used and free regions of a specified disk.
	//
	// Return Values: The method MUST return 0 or a nonerror HRESULT on success, or an implementation-specific
	// nonzero error code on failure (as specified in [MS-ERREF] section 2.1; see also section
	// 2.2.1 for HRESULT values predefined by the Disk Management Remote Protocol).
	EnumDiskRegions(context.Context, *EnumDiskRegionsRequest, ...dcerpc.CallOption) (*EnumDiskRegionsResponse, error)

	// CreatePartition operation.
	CreatePartition(context.Context, *CreatePartitionRequest, ...dcerpc.CallOption) (*CreatePartitionResponse, error)

	// CreatePartitionAssignAndFormat operation.
	CreatePartitionAssignAndFormat(context.Context, *CreatePartitionAssignAndFormatRequest, ...dcerpc.CallOption) (*CreatePartitionAssignAndFormatResponse, error)

	// CreatePartitionAssignAndFormatEx operation.
	CreatePartitionAssignAndFormatEx(context.Context, *CreatePartitionAssignAndFormatExRequest, ...dcerpc.CallOption) (*CreatePartitionAssignAndFormatExResponse, error)

	// DeletePartition operation.
	DeletePartition(context.Context, *DeletePartitionRequest, ...dcerpc.CallOption) (*DeletePartitionResponse, error)

	// The WriteSignature method writes a disk signature to a specified disk. This is a
	// synchronous task.
	//
	// Return Values: The method MUST return 0 or a nonerror HRESULT on success, or an implementation-specific
	// nonzero error code on failure (as specified in [MS-ERREF] section 2.1; see also section
	// 2.2.1 for HRESULT values predefined by the Disk Management Remote Protocol).
	WriteSignature(context.Context, *WriteSignatureRequest, ...dcerpc.CallOption) (*WriteSignatureResponse, error)

	// MarkActivePartition operation.
	MarkActivePartition(context.Context, *MarkActivePartitionRequest, ...dcerpc.CallOption) (*MarkActivePartitionResponse, error)

	// Eject operation.
	Eject(context.Context, *EjectRequest, ...dcerpc.CallOption) (*EjectResponse, error)

	// Reserved_Opnum12 operation.
	// ReservedOpnum12

	// FTEnumVolumes operation.
	FTEnumVolumes(context.Context, *FTEnumVolumesRequest, ...dcerpc.CallOption) (*FTEnumVolumesResponse, error)

	// FTEnumLogicalDiskMembers operation.
	FTEnumLogicalDiskMembers(context.Context, *FTEnumLogicalDiskMembersRequest, ...dcerpc.CallOption) (*FTEnumLogicalDiskMembersResponse, error)

	// FTDeleteVolume operation.
	FTDeleteVolume(context.Context, *FTDeleteVolumeRequest, ...dcerpc.CallOption) (*FTDeleteVolumeResponse, error)

	// FTBreakMirror operation.
	FTBreakMirror(context.Context, *FTBreakMirrorRequest, ...dcerpc.CallOption) (*FTBreakMirrorResponse, error)

	// FTResyncMirror operation.
	FTResyncMirror(context.Context, *FTResyncMirrorRequest, ...dcerpc.CallOption) (*FTResyncMirrorResponse, error)

	// FTRegenerateParityStripe operation.
	FTRegenerateParityStripe(context.Context, *FTRegenerateParityStripeRequest, ...dcerpc.CallOption) (*FTRegenerateParityStripeResponse, error)

	// FTReplaceMirrorPartition operation.
	FTReplaceMirrorPartition(context.Context, *FTReplaceMirrorPartitionRequest, ...dcerpc.CallOption) (*FTReplaceMirrorPartitionResponse, error)

	// FTReplaceParityStripePartition operation.
	FTReplaceParityStripePartition(context.Context, *FTReplaceParityStripePartitionRequest, ...dcerpc.CallOption) (*FTReplaceParityStripePartitionResponse, error)

	// EnumDriveLetters operation.
	EnumDriveLetters(context.Context, *EnumDriveLettersRequest, ...dcerpc.CallOption) (*EnumDriveLettersResponse, error)

	// AssignDriveLetter operation.
	AssignDriveLetter(context.Context, *AssignDriveLetterRequest, ...dcerpc.CallOption) (*AssignDriveLetterResponse, error)

	// FreeDriveLetter operation.
	FreeDriveLetter(context.Context, *FreeDriveLetterRequest, ...dcerpc.CallOption) (*FreeDriveLetterResponse, error)

	// EnumLocalFileSystems operation.
	EnumLocalFileSystems(context.Context, *EnumLocalFileSystemsRequest, ...dcerpc.CallOption) (*EnumLocalFileSystemsResponse, error)

	// GetInstalledFileSystems operation.
	GetInstalledFileSystems(context.Context, *GetInstalledFileSystemsRequest, ...dcerpc.CallOption) (*GetInstalledFileSystemsResponse, error)

	// Format operation.
	Format(context.Context, *FormatRequest, ...dcerpc.CallOption) (*FormatResponse, error)

	// Reserved27 operation.
	// _

	// EnumVolumes operation.
	EnumVolumes(context.Context, *EnumVolumesRequest, ...dcerpc.CallOption) (*EnumVolumesResponse, error)

	// EnumVolumeMembers operation.
	EnumVolumeMembers(context.Context, *EnumVolumeMembersRequest, ...dcerpc.CallOption) (*EnumVolumeMembersResponse, error)

	// CreateVolume operation.
	CreateVolume(context.Context, *CreateVolumeRequest, ...dcerpc.CallOption) (*CreateVolumeResponse, error)

	// CreateVolumeAssignAndFormat operation.
	CreateVolumeAssignAndFormat(context.Context, *CreateVolumeAssignAndFormatRequest, ...dcerpc.CallOption) (*CreateVolumeAssignAndFormatResponse, error)

	// CreateVolumeAssignAndFormatEx operation.
	CreateVolumeAssignAndFormatEx(context.Context, *CreateVolumeAssignAndFormatExRequest, ...dcerpc.CallOption) (*CreateVolumeAssignAndFormatExResponse, error)

	// GetVolumeMountName operation.
	GetVolumeMountName(context.Context, *GetVolumeMountNameRequest, ...dcerpc.CallOption) (*GetVolumeMountNameResponse, error)

	// GrowVolume operation.
	GrowVolume(context.Context, *GrowVolumeRequest, ...dcerpc.CallOption) (*GrowVolumeResponse, error)

	// DeleteVolume operation.
	DeleteVolume(context.Context, *DeleteVolumeRequest, ...dcerpc.CallOption) (*DeleteVolumeResponse, error)

	// AddMirror operation.
	AddMirror(context.Context, *AddMirrorRequest, ...dcerpc.CallOption) (*AddMirrorResponse, error)

	// RemoveMirror operation.
	RemoveMirror(context.Context, *RemoveMirrorRequest, ...dcerpc.CallOption) (*RemoveMirrorResponse, error)

	// SplitMirror operation.
	SplitMirror(context.Context, *SplitMirrorRequest, ...dcerpc.CallOption) (*SplitMirrorResponse, error)

	// The InitializeDisk method converts an uninitialized disk into a dynamic disk. This
	// is a synchronous task.
	//
	// Return Values: The method MUST return 0 or a nonerror HRESULT on success, or an implementation-specific
	// nonzero error code on failure (as specified in [MS-ERREF]; see also section 2.2.1
	// for HRESULT values predefined by the Disk Management Remote Protocol).
	InitializeDisk(context.Context, *InitializeDiskRequest, ...dcerpc.CallOption) (*InitializeDiskResponse, error)

	// UninitializeDisk operation.
	UninitializeDisk(context.Context, *UninitializeDiskRequest, ...dcerpc.CallOption) (*UninitializeDiskResponse, error)

	// ReConnectDisk operation.
	ReConnectDisk(context.Context, *ReConnectDiskRequest, ...dcerpc.CallOption) (*ReConnectDiskResponse, error)

	// Reserved_Opnum42 operation.
	// ReservedOpnum42

	// ImportDiskGroup operation.
	ImportDiskGroup(context.Context, *ImportDiskGroupRequest, ...dcerpc.CallOption) (*ImportDiskGroupResponse, error)

	// DiskMergeQuery operation.
	DiskMergeQuery(context.Context, *DiskMergeQueryRequest, ...dcerpc.CallOption) (*DiskMergeQueryResponse, error)

	// DiskMerge operation.
	DiskMerge(context.Context, *DiskMergeRequest, ...dcerpc.CallOption) (*DiskMergeResponse, error)

	// Reserved_Opnum46 operation.
	// ReservedOpnum46

	// ReAttachDisk operation.
	ReAttachDisk(context.Context, *ReAttachDiskRequest, ...dcerpc.CallOption) (*ReAttachDiskResponse, error)

	// Reserved_Opnum48 operation.
	// ReservedOpnum48

	// Reserved_Opnum49 operation.
	// ReservedOpnum49

	// Reserved_Opnum50 operation.
	// ReservedOpnum50

	// ReplaceRaid5Column operation.
	ReplaceRAID5Column(context.Context, *ReplaceRAID5ColumnRequest, ...dcerpc.CallOption) (*ReplaceRAID5ColumnResponse, error)

	// RestartVolume operation.
	RestartVolume(context.Context, *RestartVolumeRequest, ...dcerpc.CallOption) (*RestartVolumeResponse, error)

	// The GetEncapsulateDiskInfo method gathers the information needed to convert the specified
	// basic disks to dynamic disks. This is a synchronous task.
	//
	// Return Values: The method MUST return 0 or a nonerror HRESULT on success, or an implementation-specific
	// nonzero error code on failure (as specified in [MS-ERREF]; see also section 2.2.1
	// for HRESULT values predefined by the Disk Management Remote Protocol).
	GetEncapsulateDiskInfo(context.Context, *GetEncapsulateDiskInfoRequest, ...dcerpc.CallOption) (*GetEncapsulateDiskInfoResponse, error)

	// The EncapsulateDisk method converts the specified basic disks to dynamic disks. This
	// is a synchronous task.
	//
	// Return Values: The method MUST return 0 or a nonerror HRESULT on success, or an implementation-specific
	// nonzero error code on failure (as specified in [MS-ERREF]; see also section 2.2.1
	// for HRESULT values predefined by the Disk Management Remote Protocol).
	EncapsulateDisk(context.Context, *EncapsulateDiskRequest, ...dcerpc.CallOption) (*EncapsulateDiskResponse, error)

	// QueryChangePartitionNumbers operation.
	QueryChangePartitionNumbers(context.Context, *QueryChangePartitionNumbersRequest, ...dcerpc.CallOption) (*QueryChangePartitionNumbersResponse, error)

	// DeletePartitionNumberInfoFromRegistry operation.
	DeletePartitionNumberInfoFromRegistry(context.Context, *DeletePartitionNumberInfoFromRegistryRequest, ...dcerpc.CallOption) (*DeletePartitionNumberInfoFromRegistryResponse, error)

	// SetDontShow operation.
	SetDontShow(context.Context, *SetDontShowRequest, ...dcerpc.CallOption) (*SetDontShowResponse, error)

	// GetDontShow operation.
	GetDontShow(context.Context, *GetDontShowRequest, ...dcerpc.CallOption) (*GetDontShowResponse, error)

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
	EnumTasks(context.Context, *EnumTasksRequest, ...dcerpc.CallOption) (*EnumTasksResponse, error)

	// GetTaskDetail operation.
	GetTaskDetail(context.Context, *GetTaskDetailRequest, ...dcerpc.CallOption) (*GetTaskDetailResponse, error)

	// AbortTask operation.
	AbortTask(context.Context, *AbortTaskRequest, ...dcerpc.CallOption) (*AbortTaskResponse, error)

	// HrGetErrorData operation.
	HResultGetErrorData(context.Context, *HResultGetErrorDataRequest, ...dcerpc.CallOption) (*HResultGetErrorDataResponse, error)

	// Initialize operation.
	Initialize(context.Context, *InitializeRequest, ...dcerpc.CallOption) (*InitializeResponse, error)

	// Uninitialize operation.
	Uninitialize(context.Context, *UninitializeRequest, ...dcerpc.CallOption) (*UninitializeResponse, error)

	// Refresh operation.
	Refresh(context.Context, *RefreshRequest, ...dcerpc.CallOption) (*RefreshResponse, error)

	// RescanDisks operation.
	RescanDisks(context.Context, *RescanDisksRequest, ...dcerpc.CallOption) (*RescanDisksResponse, error)

	// RefreshFileSys operation.
	RefreshFileSystem(context.Context, *RefreshFileSystemRequest, ...dcerpc.CallOption) (*RefreshFileSystemResponse, error)

	// SecureSystemPartition operation.
	SecureSystemPartition(context.Context, *SecureSystemPartitionRequest, ...dcerpc.CallOption) (*SecureSystemPartitionResponse, error)

	// ShutDownSystem operation.
	ShutDownSystem(context.Context, *ShutDownSystemRequest, ...dcerpc.CallOption) (*ShutDownSystemResponse, error)

	// EnumAccessPath operation.
	EnumAccessPath(context.Context, *EnumAccessPathRequest, ...dcerpc.CallOption) (*EnumAccessPathResponse, error)

	// EnumAccessPathForVolume operation.
	EnumAccessPathForVolume(context.Context, *EnumAccessPathForVolumeRequest, ...dcerpc.CallOption) (*EnumAccessPathForVolumeResponse, error)

	// AddAccessPath operation.
	AddAccessPath(context.Context, *AddAccessPathRequest, ...dcerpc.CallOption) (*AddAccessPathResponse, error)

	// DeleteAccessPath operation.
	DeleteAccessPath(context.Context, *DeleteAccessPathRequest, ...dcerpc.CallOption) (*DeleteAccessPathResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) VolumeClientClient
}

type xxx_DefaultVolumeClientClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultVolumeClientClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultVolumeClientClient) EnumDisks(ctx context.Context, in *EnumDisksRequest, opts ...dcerpc.CallOption) (*EnumDisksResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumDisksResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) EnumDiskRegions(ctx context.Context, in *EnumDiskRegionsRequest, opts ...dcerpc.CallOption) (*EnumDiskRegionsResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumDiskRegionsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) CreatePartition(ctx context.Context, in *CreatePartitionRequest, opts ...dcerpc.CallOption) (*CreatePartitionResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreatePartitionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) CreatePartitionAssignAndFormat(ctx context.Context, in *CreatePartitionAssignAndFormatRequest, opts ...dcerpc.CallOption) (*CreatePartitionAssignAndFormatResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreatePartitionAssignAndFormatResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) CreatePartitionAssignAndFormatEx(ctx context.Context, in *CreatePartitionAssignAndFormatExRequest, opts ...dcerpc.CallOption) (*CreatePartitionAssignAndFormatExResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreatePartitionAssignAndFormatExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) DeletePartition(ctx context.Context, in *DeletePartitionRequest, opts ...dcerpc.CallOption) (*DeletePartitionResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DeletePartitionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) WriteSignature(ctx context.Context, in *WriteSignatureRequest, opts ...dcerpc.CallOption) (*WriteSignatureResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WriteSignatureResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) MarkActivePartition(ctx context.Context, in *MarkActivePartitionRequest, opts ...dcerpc.CallOption) (*MarkActivePartitionResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MarkActivePartitionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) Eject(ctx context.Context, in *EjectRequest, opts ...dcerpc.CallOption) (*EjectResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EjectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) FTEnumVolumes(ctx context.Context, in *FTEnumVolumesRequest, opts ...dcerpc.CallOption) (*FTEnumVolumesResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &FTEnumVolumesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) FTEnumLogicalDiskMembers(ctx context.Context, in *FTEnumLogicalDiskMembersRequest, opts ...dcerpc.CallOption) (*FTEnumLogicalDiskMembersResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &FTEnumLogicalDiskMembersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) FTDeleteVolume(ctx context.Context, in *FTDeleteVolumeRequest, opts ...dcerpc.CallOption) (*FTDeleteVolumeResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &FTDeleteVolumeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) FTBreakMirror(ctx context.Context, in *FTBreakMirrorRequest, opts ...dcerpc.CallOption) (*FTBreakMirrorResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &FTBreakMirrorResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) FTResyncMirror(ctx context.Context, in *FTResyncMirrorRequest, opts ...dcerpc.CallOption) (*FTResyncMirrorResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &FTResyncMirrorResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) FTRegenerateParityStripe(ctx context.Context, in *FTRegenerateParityStripeRequest, opts ...dcerpc.CallOption) (*FTRegenerateParityStripeResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &FTRegenerateParityStripeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) FTReplaceMirrorPartition(ctx context.Context, in *FTReplaceMirrorPartitionRequest, opts ...dcerpc.CallOption) (*FTReplaceMirrorPartitionResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &FTReplaceMirrorPartitionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) FTReplaceParityStripePartition(ctx context.Context, in *FTReplaceParityStripePartitionRequest, opts ...dcerpc.CallOption) (*FTReplaceParityStripePartitionResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &FTReplaceParityStripePartitionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) EnumDriveLetters(ctx context.Context, in *EnumDriveLettersRequest, opts ...dcerpc.CallOption) (*EnumDriveLettersResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumDriveLettersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) AssignDriveLetter(ctx context.Context, in *AssignDriveLetterRequest, opts ...dcerpc.CallOption) (*AssignDriveLetterResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AssignDriveLetterResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) FreeDriveLetter(ctx context.Context, in *FreeDriveLetterRequest, opts ...dcerpc.CallOption) (*FreeDriveLetterResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &FreeDriveLetterResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) EnumLocalFileSystems(ctx context.Context, in *EnumLocalFileSystemsRequest, opts ...dcerpc.CallOption) (*EnumLocalFileSystemsResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumLocalFileSystemsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) GetInstalledFileSystems(ctx context.Context, in *GetInstalledFileSystemsRequest, opts ...dcerpc.CallOption) (*GetInstalledFileSystemsResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetInstalledFileSystemsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) Format(ctx context.Context, in *FormatRequest, opts ...dcerpc.CallOption) (*FormatResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &FormatResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) EnumVolumes(ctx context.Context, in *EnumVolumesRequest, opts ...dcerpc.CallOption) (*EnumVolumesResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumVolumesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) EnumVolumeMembers(ctx context.Context, in *EnumVolumeMembersRequest, opts ...dcerpc.CallOption) (*EnumVolumeMembersResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumVolumeMembersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) CreateVolume(ctx context.Context, in *CreateVolumeRequest, opts ...dcerpc.CallOption) (*CreateVolumeResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateVolumeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) CreateVolumeAssignAndFormat(ctx context.Context, in *CreateVolumeAssignAndFormatRequest, opts ...dcerpc.CallOption) (*CreateVolumeAssignAndFormatResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateVolumeAssignAndFormatResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) CreateVolumeAssignAndFormatEx(ctx context.Context, in *CreateVolumeAssignAndFormatExRequest, opts ...dcerpc.CallOption) (*CreateVolumeAssignAndFormatExResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateVolumeAssignAndFormatExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) GetVolumeMountName(ctx context.Context, in *GetVolumeMountNameRequest, opts ...dcerpc.CallOption) (*GetVolumeMountNameResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetVolumeMountNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) GrowVolume(ctx context.Context, in *GrowVolumeRequest, opts ...dcerpc.CallOption) (*GrowVolumeResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GrowVolumeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) DeleteVolume(ctx context.Context, in *DeleteVolumeRequest, opts ...dcerpc.CallOption) (*DeleteVolumeResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DeleteVolumeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) AddMirror(ctx context.Context, in *AddMirrorRequest, opts ...dcerpc.CallOption) (*AddMirrorResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AddMirrorResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) RemoveMirror(ctx context.Context, in *RemoveMirrorRequest, opts ...dcerpc.CallOption) (*RemoveMirrorResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemoveMirrorResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) SplitMirror(ctx context.Context, in *SplitMirrorRequest, opts ...dcerpc.CallOption) (*SplitMirrorResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SplitMirrorResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) InitializeDisk(ctx context.Context, in *InitializeDiskRequest, opts ...dcerpc.CallOption) (*InitializeDiskResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &InitializeDiskResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) UninitializeDisk(ctx context.Context, in *UninitializeDiskRequest, opts ...dcerpc.CallOption) (*UninitializeDiskResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &UninitializeDiskResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) ReConnectDisk(ctx context.Context, in *ReConnectDiskRequest, opts ...dcerpc.CallOption) (*ReConnectDiskResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ReConnectDiskResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) ImportDiskGroup(ctx context.Context, in *ImportDiskGroupRequest, opts ...dcerpc.CallOption) (*ImportDiskGroupResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ImportDiskGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) DiskMergeQuery(ctx context.Context, in *DiskMergeQueryRequest, opts ...dcerpc.CallOption) (*DiskMergeQueryResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DiskMergeQueryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) DiskMerge(ctx context.Context, in *DiskMergeRequest, opts ...dcerpc.CallOption) (*DiskMergeResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DiskMergeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) ReAttachDisk(ctx context.Context, in *ReAttachDiskRequest, opts ...dcerpc.CallOption) (*ReAttachDiskResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ReAttachDiskResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) ReplaceRAID5Column(ctx context.Context, in *ReplaceRAID5ColumnRequest, opts ...dcerpc.CallOption) (*ReplaceRAID5ColumnResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ReplaceRAID5ColumnResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) RestartVolume(ctx context.Context, in *RestartVolumeRequest, opts ...dcerpc.CallOption) (*RestartVolumeResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RestartVolumeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) GetEncapsulateDiskInfo(ctx context.Context, in *GetEncapsulateDiskInfoRequest, opts ...dcerpc.CallOption) (*GetEncapsulateDiskInfoResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetEncapsulateDiskInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) EncapsulateDisk(ctx context.Context, in *EncapsulateDiskRequest, opts ...dcerpc.CallOption) (*EncapsulateDiskResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EncapsulateDiskResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) QueryChangePartitionNumbers(ctx context.Context, in *QueryChangePartitionNumbersRequest, opts ...dcerpc.CallOption) (*QueryChangePartitionNumbersResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryChangePartitionNumbersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) DeletePartitionNumberInfoFromRegistry(ctx context.Context, in *DeletePartitionNumberInfoFromRegistryRequest, opts ...dcerpc.CallOption) (*DeletePartitionNumberInfoFromRegistryResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DeletePartitionNumberInfoFromRegistryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) SetDontShow(ctx context.Context, in *SetDontShowRequest, opts ...dcerpc.CallOption) (*SetDontShowResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetDontShowResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) GetDontShow(ctx context.Context, in *GetDontShowRequest, opts ...dcerpc.CallOption) (*GetDontShowResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetDontShowResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) EnumTasks(ctx context.Context, in *EnumTasksRequest, opts ...dcerpc.CallOption) (*EnumTasksResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumTasksResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) GetTaskDetail(ctx context.Context, in *GetTaskDetailRequest, opts ...dcerpc.CallOption) (*GetTaskDetailResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetTaskDetailResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) AbortTask(ctx context.Context, in *AbortTaskRequest, opts ...dcerpc.CallOption) (*AbortTaskResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AbortTaskResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) HResultGetErrorData(ctx context.Context, in *HResultGetErrorDataRequest, opts ...dcerpc.CallOption) (*HResultGetErrorDataResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &HResultGetErrorDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) Initialize(ctx context.Context, in *InitializeRequest, opts ...dcerpc.CallOption) (*InitializeResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &InitializeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) Uninitialize(ctx context.Context, in *UninitializeRequest, opts ...dcerpc.CallOption) (*UninitializeResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &UninitializeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) Refresh(ctx context.Context, in *RefreshRequest, opts ...dcerpc.CallOption) (*RefreshResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RefreshResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) RescanDisks(ctx context.Context, in *RescanDisksRequest, opts ...dcerpc.CallOption) (*RescanDisksResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RescanDisksResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) RefreshFileSystem(ctx context.Context, in *RefreshFileSystemRequest, opts ...dcerpc.CallOption) (*RefreshFileSystemResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RefreshFileSystemResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) SecureSystemPartition(ctx context.Context, in *SecureSystemPartitionRequest, opts ...dcerpc.CallOption) (*SecureSystemPartitionResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SecureSystemPartitionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) ShutDownSystem(ctx context.Context, in *ShutDownSystemRequest, opts ...dcerpc.CallOption) (*ShutDownSystemResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ShutDownSystemResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) EnumAccessPath(ctx context.Context, in *EnumAccessPathRequest, opts ...dcerpc.CallOption) (*EnumAccessPathResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumAccessPathResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) EnumAccessPathForVolume(ctx context.Context, in *EnumAccessPathForVolumeRequest, opts ...dcerpc.CallOption) (*EnumAccessPathForVolumeResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumAccessPathForVolumeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) AddAccessPath(ctx context.Context, in *AddAccessPathRequest, opts ...dcerpc.CallOption) (*AddAccessPathResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AddAccessPathResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) DeleteAccessPath(ctx context.Context, in *DeleteAccessPathRequest, opts ...dcerpc.CallOption) (*DeleteAccessPathResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DeleteAccessPathResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClientClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultVolumeClientClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultVolumeClientClient) IPID(ctx context.Context, ipid *dcom.IPID) VolumeClientClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultVolumeClientClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewVolumeClientClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (VolumeClientClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(VolumeClientSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iunknown.NewUnknownClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultVolumeClientClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_EnumDisksOperation structure represents the EnumDisks operation
type xxx_EnumDisksOperation struct {
	This      *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat   `idl:"name:That" json:"that"`
	DiskCount uint32           `idl:"name:diskCount" json:"disk_count"`
	DiskList  []*dmrp.DiskInfo `idl:"name:diskList;size_is:(, diskCount)" json:"disk_list"`
	Return    int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumDisksOperation) OpNum() int { return 3 }

func (o *xxx_EnumDisksOperation) OpName() string { return "/IVolumeClient/v0/EnumDisks" }

func (o *xxx_EnumDisksOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumDisksOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumDisksOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumDisksOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.DiskList != nil && o.DiskCount == 0 {
		o.DiskCount = uint32(len(o.DiskList))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumDisksOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// diskCount {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.DiskCount); err != nil {
			return err
		}
	}
	// diskList {out} (1:{pointer=ref}*(2)*(1))(2:{alias=DISK_INFO}[dim:0,size_is=diskCount](struct))
	{
		if o.DiskList != nil || o.DiskCount > 0 {
			_ptr_diskList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.DiskCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.DiskList {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.DiskList[i1] != nil {
						if err := o.DiskList[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dmrp.DiskInfo{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.DiskList); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&dmrp.DiskInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DiskList, _ptr_diskList); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumDisksOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// diskCount {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.DiskCount); err != nil {
			return err
		}
	}
	// diskList {out} (1:{pointer=ref}*(2)*(1))(2:{alias=DISK_INFO}[dim:0,size_is=diskCount](struct))
	{
		_ptr_diskList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.DiskList", sizeInfo[0])
			}
			o.DiskList = make([]*dmrp.DiskInfo, sizeInfo[0])
			for i1 := range o.DiskList {
				i1 := i1
				if o.DiskList[i1] == nil {
					o.DiskList[i1] = &dmrp.DiskInfo{}
				}
				if err := o.DiskList[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_diskList := func(ptr interface{}) { o.DiskList = *ptr.(*[]*dmrp.DiskInfo) }
		if err := w.ReadPointer(&o.DiskList, _s_diskList, _ptr_diskList); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumDisksRequest structure represents the EnumDisks operation request
type EnumDisksRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *EnumDisksRequest) xxx_ToOp(ctx context.Context) *xxx_EnumDisksOperation {
	if o == nil {
		return &xxx_EnumDisksOperation{}
	}
	return &xxx_EnumDisksOperation{
		This: o.This,
	}
}

func (o *EnumDisksRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumDisksOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *EnumDisksRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumDisksRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumDisksOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumDisksResponse structure represents the EnumDisks operation response
type EnumDisksResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// diskCount: Number of pointers in diskList.
	DiskCount uint32 `idl:"name:diskCount" json:"disk_count"`
	// diskList: Pointer to an array of DISK_INFO structures.
	DiskList []*dmrp.DiskInfo `idl:"name:diskList;size_is:(, diskCount)" json:"disk_list"`
	// Return: The EnumDisks return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumDisksResponse) xxx_ToOp(ctx context.Context) *xxx_EnumDisksOperation {
	if o == nil {
		return &xxx_EnumDisksOperation{}
	}
	return &xxx_EnumDisksOperation{
		That:      o.That,
		DiskCount: o.DiskCount,
		DiskList:  o.DiskList,
		Return:    o.Return,
	}
}

func (o *EnumDisksResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumDisksOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DiskCount = op.DiskCount
	o.DiskList = op.DiskList
	o.Return = op.Return
}
func (o *EnumDisksResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumDisksResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumDisksOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumDiskRegionsOperation structure represents the EnumDiskRegions operation
type xxx_EnumDiskRegionsOperation struct {
	This          *dcom.ORPCThis     `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat     `idl:"name:That" json:"that"`
	DiskID        int64              `idl:"name:diskId" json:"disk_id"`
	RegionsLength uint32             `idl:"name:numRegions" json:"regions_length"`
	RegionList    []*dmrp.RegionInfo `idl:"name:regionList;size_is:(, numRegions)" json:"region_list"`
	Return        int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumDiskRegionsOperation) OpNum() int { return 4 }

func (o *xxx_EnumDiskRegionsOperation) OpName() string { return "/IVolumeClient/v0/EnumDiskRegions" }

func (o *xxx_EnumDiskRegionsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumDiskRegionsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// diskId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.DiskID); err != nil {
			return err
		}
	}
	// numRegions {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.RegionsLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumDiskRegionsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// diskId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.DiskID); err != nil {
			return err
		}
	}
	// numRegions {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.RegionsLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumDiskRegionsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.RegionList != nil && o.RegionsLength == 0 {
		o.RegionsLength = uint32(len(o.RegionList))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumDiskRegionsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// numRegions {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.RegionsLength); err != nil {
			return err
		}
	}
	// regionList {out} (1:{pointer=ref}*(2)*(1))(2:{alias=REGION_INFO}[dim:0,size_is=numRegions](struct))
	{
		if o.RegionList != nil || o.RegionsLength > 0 {
			_ptr_regionList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.RegionsLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.RegionList {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.RegionList[i1] != nil {
						if err := o.RegionList[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dmrp.RegionInfo{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.RegionList); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&dmrp.RegionInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RegionList, _ptr_regionList); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumDiskRegionsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// numRegions {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.RegionsLength); err != nil {
			return err
		}
	}
	// regionList {out} (1:{pointer=ref}*(2)*(1))(2:{alias=REGION_INFO}[dim:0,size_is=numRegions](struct))
	{
		_ptr_regionList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.RegionList", sizeInfo[0])
			}
			o.RegionList = make([]*dmrp.RegionInfo, sizeInfo[0])
			for i1 := range o.RegionList {
				i1 := i1
				if o.RegionList[i1] == nil {
					o.RegionList[i1] = &dmrp.RegionInfo{}
				}
				if err := o.RegionList[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_regionList := func(ptr interface{}) { o.RegionList = *ptr.(*[]*dmrp.RegionInfo) }
		if err := w.ReadPointer(&o.RegionList, _s_regionList, _ptr_regionList); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumDiskRegionsRequest structure represents the EnumDiskRegions operation request
type EnumDiskRegionsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// diskId: Specifies the OID of the disk for which regions are being enumerated.
	DiskID int64 `idl:"name:diskId" json:"disk_id"`
	// numRegions: Pointer to the number of regions in regionList.
	RegionsLength uint32 `idl:"name:numRegions" json:"regions_length"`
}

func (o *EnumDiskRegionsRequest) xxx_ToOp(ctx context.Context) *xxx_EnumDiskRegionsOperation {
	if o == nil {
		return &xxx_EnumDiskRegionsOperation{}
	}
	return &xxx_EnumDiskRegionsOperation{
		This:          o.This,
		DiskID:        o.DiskID,
		RegionsLength: o.RegionsLength,
	}
}

func (o *EnumDiskRegionsRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumDiskRegionsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
	o.RegionsLength = op.RegionsLength
}
func (o *EnumDiskRegionsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumDiskRegionsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumDiskRegionsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumDiskRegionsResponse structure represents the EnumDiskRegions operation response
type EnumDiskRegionsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// numRegions: Pointer to the number of regions in regionList.
	RegionsLength uint32 `idl:"name:numRegions" json:"regions_length"`
	// regionList: Pointer to an array of REGION_INFO structures.
	RegionList []*dmrp.RegionInfo `idl:"name:regionList;size_is:(, numRegions)" json:"region_list"`
	// Return: The EnumDiskRegions return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumDiskRegionsResponse) xxx_ToOp(ctx context.Context) *xxx_EnumDiskRegionsOperation {
	if o == nil {
		return &xxx_EnumDiskRegionsOperation{}
	}
	return &xxx_EnumDiskRegionsOperation{
		That:          o.That,
		RegionsLength: o.RegionsLength,
		RegionList:    o.RegionList,
		Return:        o.Return,
	}
}

func (o *EnumDiskRegionsResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumDiskRegionsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.RegionsLength = op.RegionsLength
	o.RegionList = op.RegionList
	o.Return = op.Return
}
func (o *EnumDiskRegionsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumDiskRegionsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumDiskRegionsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreatePartitionOperation structure represents the CreatePartition operation
type xxx_CreatePartitionOperation struct {
	This          *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat   `idl:"name:That" json:"that"`
	PartitionSpec *dmrp.RegionSpec `idl:"name:partitionSpec" json:"partition_spec"`
	TaskInfo      *dmrp.TaskInfo   `idl:"name:tinfo" json:"task_info"`
	Return        int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_CreatePartitionOperation) OpNum() int { return 5 }

func (o *xxx_CreatePartitionOperation) OpName() string { return "/IVolumeClient/v0/CreatePartition" }

func (o *xxx_CreatePartitionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreatePartitionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// partitionSpec {in} (1:{alias=REGION_SPEC}(struct))
	{
		if o.PartitionSpec != nil {
			if err := o.PartitionSpec.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.RegionSpec{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CreatePartitionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// partitionSpec {in} (1:{alias=REGION_SPEC}(struct))
	{
		if o.PartitionSpec == nil {
			o.PartitionSpec = &dmrp.RegionSpec{}
		}
		if err := o.PartitionSpec.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreatePartitionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreatePartitionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreatePartitionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreatePartitionRequest structure represents the CreatePartition operation request
type CreatePartitionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This          *dcom.ORPCThis   `idl:"name:This" json:"this"`
	PartitionSpec *dmrp.RegionSpec `idl:"name:partitionSpec" json:"partition_spec"`
}

func (o *CreatePartitionRequest) xxx_ToOp(ctx context.Context) *xxx_CreatePartitionOperation {
	if o == nil {
		return &xxx_CreatePartitionOperation{}
	}
	return &xxx_CreatePartitionOperation{
		This:          o.This,
		PartitionSpec: o.PartitionSpec,
	}
}

func (o *CreatePartitionRequest) xxx_FromOp(ctx context.Context, op *xxx_CreatePartitionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PartitionSpec = op.PartitionSpec
}
func (o *CreatePartitionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreatePartitionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreatePartitionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreatePartitionResponse structure represents the CreatePartition operation response
type CreatePartitionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The CreatePartition return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreatePartitionResponse) xxx_ToOp(ctx context.Context) *xxx_CreatePartitionOperation {
	if o == nil {
		return &xxx_CreatePartitionOperation{}
	}
	return &xxx_CreatePartitionOperation{
		That:     o.That,
		TaskInfo: o.TaskInfo,
		Return:   o.Return,
	}
}

func (o *CreatePartitionResponse) xxx_FromOp(ctx context.Context, op *xxx_CreatePartitionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *CreatePartitionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreatePartitionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreatePartitionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreatePartitionAssignAndFormatOperation structure represents the CreatePartitionAssignAndFormat operation
type xxx_CreatePartitionAssignAndFormatOperation struct {
	This                 *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat       `idl:"name:That" json:"that"`
	PartitionSpec        *dmrp.RegionSpec     `idl:"name:partitionSpec" json:"partition_spec"`
	Letter               uint16               `idl:"name:letter" json:"letter"`
	LetterLastKnownState int64                `idl:"name:letterLastKnownState" json:"letter_last_known_state"`
	FSSpec               *dmrp.FileSystemInfo `idl:"name:fsSpec" json:"fs_spec"`
	QuickFormat          bool                 `idl:"name:quickFormat" json:"quick_format"`
	TaskInfo             *dmrp.TaskInfo       `idl:"name:tinfo" json:"task_info"`
	Return               int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_CreatePartitionAssignAndFormatOperation) OpNum() int { return 6 }

func (o *xxx_CreatePartitionAssignAndFormatOperation) OpName() string {
	return "/IVolumeClient/v0/CreatePartitionAssignAndFormat"
}

func (o *xxx_CreatePartitionAssignAndFormatOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreatePartitionAssignAndFormatOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// partitionSpec {in} (1:{alias=REGION_SPEC}(struct))
	{
		if o.PartitionSpec != nil {
			if err := o.PartitionSpec.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.RegionSpec{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// letter {in} (1:(wchar))
	{
		if err := w.WriteData(o.Letter); err != nil {
			return err
		}
	}
	// letterLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.LetterLastKnownState); err != nil {
			return err
		}
	}
	// fsSpec {in} (1:{alias=FILE_SYSTEM_INFO}(struct))
	{
		if o.FSSpec != nil {
			if err := o.FSSpec.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.FileSystemInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// quickFormat {in} (1:(bool))
	{
		if err := w.WriteData(o.QuickFormat); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreatePartitionAssignAndFormatOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// partitionSpec {in} (1:{alias=REGION_SPEC}(struct))
	{
		if o.PartitionSpec == nil {
			o.PartitionSpec = &dmrp.RegionSpec{}
		}
		if err := o.PartitionSpec.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// letter {in} (1:(wchar))
	{
		if err := w.ReadData(&o.Letter); err != nil {
			return err
		}
	}
	// letterLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.LetterLastKnownState); err != nil {
			return err
		}
	}
	// fsSpec {in} (1:{alias=FILE_SYSTEM_INFO}(struct))
	{
		if o.FSSpec == nil {
			o.FSSpec = &dmrp.FileSystemInfo{}
		}
		if err := o.FSSpec.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// quickFormat {in} (1:(bool))
	{
		if err := w.ReadData(&o.QuickFormat); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreatePartitionAssignAndFormatOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreatePartitionAssignAndFormatOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreatePartitionAssignAndFormatOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreatePartitionAssignAndFormatRequest structure represents the CreatePartitionAssignAndFormat operation request
type CreatePartitionAssignAndFormatRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis       `idl:"name:This" json:"this"`
	PartitionSpec        *dmrp.RegionSpec     `idl:"name:partitionSpec" json:"partition_spec"`
	Letter               uint16               `idl:"name:letter" json:"letter"`
	LetterLastKnownState int64                `idl:"name:letterLastKnownState" json:"letter_last_known_state"`
	FSSpec               *dmrp.FileSystemInfo `idl:"name:fsSpec" json:"fs_spec"`
	QuickFormat          bool                 `idl:"name:quickFormat" json:"quick_format"`
}

func (o *CreatePartitionAssignAndFormatRequest) xxx_ToOp(ctx context.Context) *xxx_CreatePartitionAssignAndFormatOperation {
	if o == nil {
		return &xxx_CreatePartitionAssignAndFormatOperation{}
	}
	return &xxx_CreatePartitionAssignAndFormatOperation{
		This:                 o.This,
		PartitionSpec:        o.PartitionSpec,
		Letter:               o.Letter,
		LetterLastKnownState: o.LetterLastKnownState,
		FSSpec:               o.FSSpec,
		QuickFormat:          o.QuickFormat,
	}
}

func (o *CreatePartitionAssignAndFormatRequest) xxx_FromOp(ctx context.Context, op *xxx_CreatePartitionAssignAndFormatOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PartitionSpec = op.PartitionSpec
	o.Letter = op.Letter
	o.LetterLastKnownState = op.LetterLastKnownState
	o.FSSpec = op.FSSpec
	o.QuickFormat = op.QuickFormat
}
func (o *CreatePartitionAssignAndFormatRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreatePartitionAssignAndFormatRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreatePartitionAssignAndFormatOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreatePartitionAssignAndFormatResponse structure represents the CreatePartitionAssignAndFormat operation response
type CreatePartitionAssignAndFormatResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The CreatePartitionAssignAndFormat return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreatePartitionAssignAndFormatResponse) xxx_ToOp(ctx context.Context) *xxx_CreatePartitionAssignAndFormatOperation {
	if o == nil {
		return &xxx_CreatePartitionAssignAndFormatOperation{}
	}
	return &xxx_CreatePartitionAssignAndFormatOperation{
		That:     o.That,
		TaskInfo: o.TaskInfo,
		Return:   o.Return,
	}
}

func (o *CreatePartitionAssignAndFormatResponse) xxx_FromOp(ctx context.Context, op *xxx_CreatePartitionAssignAndFormatOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *CreatePartitionAssignAndFormatResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreatePartitionAssignAndFormatResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreatePartitionAssignAndFormatOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreatePartitionAssignAndFormatExOperation structure represents the CreatePartitionAssignAndFormatEx operation
type xxx_CreatePartitionAssignAndFormatExOperation struct {
	This                 *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat       `idl:"name:That" json:"that"`
	PartitionSpec        *dmrp.RegionSpec     `idl:"name:partitionSpec" json:"partition_spec"`
	Letter               uint16               `idl:"name:letter" json:"letter"`
	LetterLastKnownState int64                `idl:"name:letterLastKnownState" json:"letter_last_known_state"`
	AccessPathLength     int32                `idl:"name:cchAccessPath" json:"access_path_length"`
	AccessPath           string               `idl:"name:AccessPath;size_is:(cchAccessPath)" json:"access_path"`
	FSSpec               *dmrp.FileSystemInfo `idl:"name:fsSpec" json:"fs_spec"`
	QuickFormat          bool                 `idl:"name:quickFormat" json:"quick_format"`
	Flags                uint32               `idl:"name:dwFlags" json:"flags"`
	TaskInfo             *dmrp.TaskInfo       `idl:"name:tinfo" json:"task_info"`
	Return               int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_CreatePartitionAssignAndFormatExOperation) OpNum() int { return 7 }

func (o *xxx_CreatePartitionAssignAndFormatExOperation) OpName() string {
	return "/IVolumeClient/v0/CreatePartitionAssignAndFormatEx"
}

func (o *xxx_CreatePartitionAssignAndFormatExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.AccessPath != "" && o.AccessPathLength == 0 {
		o.AccessPathLength = int32(len(o.AccessPath))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreatePartitionAssignAndFormatExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// partitionSpec {in} (1:{alias=REGION_SPEC}(struct))
	{
		if o.PartitionSpec != nil {
			if err := o.PartitionSpec.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.RegionSpec{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// letter {in} (1:(wchar))
	{
		if err := w.WriteData(o.Letter); err != nil {
			return err
		}
	}
	// letterLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.LetterLastKnownState); err != nil {
			return err
		}
	}
	// cchAccessPath {in} (1:(int32))
	{
		if err := w.WriteData(o.AccessPathLength); err != nil {
			return err
		}
	}
	// AccessPath {in} (1:{pointer=ref}*(1)[dim:0,size_is=cchAccessPath,string](wchar))
	{
		dimSize1 := uint64(o.AccessPathLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		_AccessPath_buf := utf16.Encode([]rune(o.AccessPath))
		if uint64(len(_AccessPath_buf)) > sizeInfo[0] {
			_AccessPath_buf = _AccessPath_buf[:sizeInfo[0]]
		}
		for i1 := range _AccessPath_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_AccessPath_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_AccessPath_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	// fsSpec {in} (1:{alias=FILE_SYSTEM_INFO}(struct))
	{
		if o.FSSpec != nil {
			if err := o.FSSpec.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.FileSystemInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// quickFormat {in} (1:(bool))
	{
		if err := w.WriteData(o.QuickFormat); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreatePartitionAssignAndFormatExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// partitionSpec {in} (1:{alias=REGION_SPEC}(struct))
	{
		if o.PartitionSpec == nil {
			o.PartitionSpec = &dmrp.RegionSpec{}
		}
		if err := o.PartitionSpec.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// letter {in} (1:(wchar))
	{
		if err := w.ReadData(&o.Letter); err != nil {
			return err
		}
	}
	// letterLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.LetterLastKnownState); err != nil {
			return err
		}
	}
	// cchAccessPath {in} (1:(int32))
	{
		if err := w.ReadData(&o.AccessPathLength); err != nil {
			return err
		}
	}
	// AccessPath {in} (1:{pointer=ref}*(1)[dim:0,size_is=cchAccessPath,string](wchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _AccessPath_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _AccessPath_buf", sizeInfo[0])
		}
		_AccessPath_buf = make([]uint16, sizeInfo[0])
		for i1 := range _AccessPath_buf {
			i1 := i1
			if err := w.ReadData(&_AccessPath_buf[i1]); err != nil {
				return err
			}
		}
		o.AccessPath = strings.TrimRight(string(utf16.Decode(_AccessPath_buf)), ndr.ZeroString)
	}
	// fsSpec {in} (1:{alias=FILE_SYSTEM_INFO}(struct))
	{
		if o.FSSpec == nil {
			o.FSSpec = &dmrp.FileSystemInfo{}
		}
		if err := o.FSSpec.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// quickFormat {in} (1:(bool))
	{
		if err := w.ReadData(&o.QuickFormat); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreatePartitionAssignAndFormatExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreatePartitionAssignAndFormatExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreatePartitionAssignAndFormatExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreatePartitionAssignAndFormatExRequest structure represents the CreatePartitionAssignAndFormatEx operation request
type CreatePartitionAssignAndFormatExRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis       `idl:"name:This" json:"this"`
	PartitionSpec        *dmrp.RegionSpec     `idl:"name:partitionSpec" json:"partition_spec"`
	Letter               uint16               `idl:"name:letter" json:"letter"`
	LetterLastKnownState int64                `idl:"name:letterLastKnownState" json:"letter_last_known_state"`
	AccessPathLength     int32                `idl:"name:cchAccessPath" json:"access_path_length"`
	AccessPath           string               `idl:"name:AccessPath;size_is:(cchAccessPath)" json:"access_path"`
	FSSpec               *dmrp.FileSystemInfo `idl:"name:fsSpec" json:"fs_spec"`
	QuickFormat          bool                 `idl:"name:quickFormat" json:"quick_format"`
	Flags                uint32               `idl:"name:dwFlags" json:"flags"`
}

func (o *CreatePartitionAssignAndFormatExRequest) xxx_ToOp(ctx context.Context) *xxx_CreatePartitionAssignAndFormatExOperation {
	if o == nil {
		return &xxx_CreatePartitionAssignAndFormatExOperation{}
	}
	return &xxx_CreatePartitionAssignAndFormatExOperation{
		This:                 o.This,
		PartitionSpec:        o.PartitionSpec,
		Letter:               o.Letter,
		LetterLastKnownState: o.LetterLastKnownState,
		AccessPathLength:     o.AccessPathLength,
		AccessPath:           o.AccessPath,
		FSSpec:               o.FSSpec,
		QuickFormat:          o.QuickFormat,
		Flags:                o.Flags,
	}
}

func (o *CreatePartitionAssignAndFormatExRequest) xxx_FromOp(ctx context.Context, op *xxx_CreatePartitionAssignAndFormatExOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PartitionSpec = op.PartitionSpec
	o.Letter = op.Letter
	o.LetterLastKnownState = op.LetterLastKnownState
	o.AccessPathLength = op.AccessPathLength
	o.AccessPath = op.AccessPath
	o.FSSpec = op.FSSpec
	o.QuickFormat = op.QuickFormat
	o.Flags = op.Flags
}
func (o *CreatePartitionAssignAndFormatExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreatePartitionAssignAndFormatExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreatePartitionAssignAndFormatExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreatePartitionAssignAndFormatExResponse structure represents the CreatePartitionAssignAndFormatEx operation response
type CreatePartitionAssignAndFormatExResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The CreatePartitionAssignAndFormatEx return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreatePartitionAssignAndFormatExResponse) xxx_ToOp(ctx context.Context) *xxx_CreatePartitionAssignAndFormatExOperation {
	if o == nil {
		return &xxx_CreatePartitionAssignAndFormatExOperation{}
	}
	return &xxx_CreatePartitionAssignAndFormatExOperation{
		That:     o.That,
		TaskInfo: o.TaskInfo,
		Return:   o.Return,
	}
}

func (o *CreatePartitionAssignAndFormatExResponse) xxx_FromOp(ctx context.Context, op *xxx_CreatePartitionAssignAndFormatExOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *CreatePartitionAssignAndFormatExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreatePartitionAssignAndFormatExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreatePartitionAssignAndFormatExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeletePartitionOperation structure represents the DeletePartition operation
type xxx_DeletePartitionOperation struct {
	This          *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat   `idl:"name:That" json:"that"`
	PartitionSpec *dmrp.RegionSpec `idl:"name:partitionSpec" json:"partition_spec"`
	Force         bool             `idl:"name:force" json:"force"`
	TaskInfo      *dmrp.TaskInfo   `idl:"name:tinfo" json:"task_info"`
	Return        int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_DeletePartitionOperation) OpNum() int { return 8 }

func (o *xxx_DeletePartitionOperation) OpName() string { return "/IVolumeClient/v0/DeletePartition" }

func (o *xxx_DeletePartitionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeletePartitionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// partitionSpec {in} (1:{alias=REGION_SPEC}(struct))
	{
		if o.PartitionSpec != nil {
			if err := o.PartitionSpec.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.RegionSpec{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// force {in} (1:(bool))
	{
		if err := w.WriteData(o.Force); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeletePartitionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// partitionSpec {in} (1:{alias=REGION_SPEC}(struct))
	{
		if o.PartitionSpec == nil {
			o.PartitionSpec = &dmrp.RegionSpec{}
		}
		if err := o.PartitionSpec.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// force {in} (1:(bool))
	{
		if err := w.ReadData(&o.Force); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeletePartitionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeletePartitionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeletePartitionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DeletePartitionRequest structure represents the DeletePartition operation request
type DeletePartitionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This          *dcom.ORPCThis   `idl:"name:This" json:"this"`
	PartitionSpec *dmrp.RegionSpec `idl:"name:partitionSpec" json:"partition_spec"`
	Force         bool             `idl:"name:force" json:"force"`
}

func (o *DeletePartitionRequest) xxx_ToOp(ctx context.Context) *xxx_DeletePartitionOperation {
	if o == nil {
		return &xxx_DeletePartitionOperation{}
	}
	return &xxx_DeletePartitionOperation{
		This:          o.This,
		PartitionSpec: o.PartitionSpec,
		Force:         o.Force,
	}
}

func (o *DeletePartitionRequest) xxx_FromOp(ctx context.Context, op *xxx_DeletePartitionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PartitionSpec = op.PartitionSpec
	o.Force = op.Force
}
func (o *DeletePartitionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeletePartitionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeletePartitionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeletePartitionResponse structure represents the DeletePartition operation response
type DeletePartitionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The DeletePartition return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeletePartitionResponse) xxx_ToOp(ctx context.Context) *xxx_DeletePartitionOperation {
	if o == nil {
		return &xxx_DeletePartitionOperation{}
	}
	return &xxx_DeletePartitionOperation{
		That:     o.That,
		TaskInfo: o.TaskInfo,
		Return:   o.Return,
	}
}

func (o *DeletePartitionResponse) xxx_FromOp(ctx context.Context, op *xxx_DeletePartitionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *DeletePartitionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeletePartitionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeletePartitionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WriteSignatureOperation structure represents the WriteSignature operation
type xxx_WriteSignatureOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	DiskID             int64          `idl:"name:diskId" json:"disk_id"`
	DiskLastKnownState int64          `idl:"name:diskLastKnownState" json:"disk_last_known_state"`
	TaskInfo           *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_WriteSignatureOperation) OpNum() int { return 9 }

func (o *xxx_WriteSignatureOperation) OpName() string { return "/IVolumeClient/v0/WriteSignature" }

func (o *xxx_WriteSignatureOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WriteSignatureOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// diskId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.DiskID); err != nil {
			return err
		}
	}
	// diskLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.DiskLastKnownState); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WriteSignatureOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// diskId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.DiskID); err != nil {
			return err
		}
	}
	// diskLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.DiskLastKnownState); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WriteSignatureOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WriteSignatureOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WriteSignatureOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WriteSignatureRequest structure represents the WriteSignature operation request
type WriteSignatureRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// diskId: Specifies the object identifier of the target disk for the signature.
	DiskID int64 `idl:"name:diskId" json:"disk_id"`
	// diskLastKnownState: Disk's last known modification sequence number.
	DiskLastKnownState int64 `idl:"name:diskLastKnownState" json:"disk_last_known_state"`
}

func (o *WriteSignatureRequest) xxx_ToOp(ctx context.Context) *xxx_WriteSignatureOperation {
	if o == nil {
		return &xxx_WriteSignatureOperation{}
	}
	return &xxx_WriteSignatureOperation{
		This:               o.This,
		DiskID:             o.DiskID,
		DiskLastKnownState: o.DiskLastKnownState,
	}
}

func (o *WriteSignatureRequest) xxx_FromOp(ctx context.Context, op *xxx_WriteSignatureOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
	o.DiskLastKnownState = op.DiskLastKnownState
}
func (o *WriteSignatureRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *WriteSignatureRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WriteSignatureOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WriteSignatureResponse structure represents the WriteSignature operation response
type WriteSignatureResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// tinfo: Pointer to a TASK_INFO structure that the client can use to track the request's
	// progress.
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The WriteSignature return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *WriteSignatureResponse) xxx_ToOp(ctx context.Context) *xxx_WriteSignatureOperation {
	if o == nil {
		return &xxx_WriteSignatureOperation{}
	}
	return &xxx_WriteSignatureOperation{
		That:     o.That,
		TaskInfo: o.TaskInfo,
		Return:   o.Return,
	}
}

func (o *WriteSignatureResponse) xxx_FromOp(ctx context.Context, op *xxx_WriteSignatureOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *WriteSignatureResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *WriteSignatureResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WriteSignatureOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MarkActivePartitionOperation structure represents the MarkActivePartition operation
type xxx_MarkActivePartitionOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	RegionID             int64          `idl:"name:regionId" json:"region_id"`
	RegionLastKnownState int64          `idl:"name:regionLastKnownState" json:"region_last_known_state"`
	TaskInfo             *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_MarkActivePartitionOperation) OpNum() int { return 10 }

func (o *xxx_MarkActivePartitionOperation) OpName() string {
	return "/IVolumeClient/v0/MarkActivePartition"
}

func (o *xxx_MarkActivePartitionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MarkActivePartitionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// regionId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.RegionID); err != nil {
			return err
		}
	}
	// regionLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.RegionLastKnownState); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MarkActivePartitionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// regionId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.RegionID); err != nil {
			return err
		}
	}
	// regionLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.RegionLastKnownState); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MarkActivePartitionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MarkActivePartitionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MarkActivePartitionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// MarkActivePartitionRequest structure represents the MarkActivePartition operation request
type MarkActivePartitionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	RegionID             int64          `idl:"name:regionId" json:"region_id"`
	RegionLastKnownState int64          `idl:"name:regionLastKnownState" json:"region_last_known_state"`
}

func (o *MarkActivePartitionRequest) xxx_ToOp(ctx context.Context) *xxx_MarkActivePartitionOperation {
	if o == nil {
		return &xxx_MarkActivePartitionOperation{}
	}
	return &xxx_MarkActivePartitionOperation{
		This:                 o.This,
		RegionID:             o.RegionID,
		RegionLastKnownState: o.RegionLastKnownState,
	}
}

func (o *MarkActivePartitionRequest) xxx_FromOp(ctx context.Context, op *xxx_MarkActivePartitionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.RegionID = op.RegionID
	o.RegionLastKnownState = op.RegionLastKnownState
}
func (o *MarkActivePartitionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *MarkActivePartitionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MarkActivePartitionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MarkActivePartitionResponse structure represents the MarkActivePartition operation response
type MarkActivePartitionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The MarkActivePartition return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *MarkActivePartitionResponse) xxx_ToOp(ctx context.Context) *xxx_MarkActivePartitionOperation {
	if o == nil {
		return &xxx_MarkActivePartitionOperation{}
	}
	return &xxx_MarkActivePartitionOperation{
		That:     o.That,
		TaskInfo: o.TaskInfo,
		Return:   o.Return,
	}
}

func (o *MarkActivePartitionResponse) xxx_FromOp(ctx context.Context, op *xxx_MarkActivePartitionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *MarkActivePartitionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *MarkActivePartitionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MarkActivePartitionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EjectOperation structure represents the Eject operation
type xxx_EjectOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	DiskID             int64          `idl:"name:diskId" json:"disk_id"`
	DiskLastKnownState int64          `idl:"name:diskLastKnownState" json:"disk_last_known_state"`
	TaskInfo           *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_EjectOperation) OpNum() int { return 11 }

func (o *xxx_EjectOperation) OpName() string { return "/IVolumeClient/v0/Eject" }

func (o *xxx_EjectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EjectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// diskId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.DiskID); err != nil {
			return err
		}
	}
	// diskLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.DiskLastKnownState); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EjectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// diskId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.DiskID); err != nil {
			return err
		}
	}
	// diskLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.DiskLastKnownState); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EjectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EjectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EjectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EjectRequest structure represents the Eject operation request
type EjectRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	DiskID             int64          `idl:"name:diskId" json:"disk_id"`
	DiskLastKnownState int64          `idl:"name:diskLastKnownState" json:"disk_last_known_state"`
}

func (o *EjectRequest) xxx_ToOp(ctx context.Context) *xxx_EjectOperation {
	if o == nil {
		return &xxx_EjectOperation{}
	}
	return &xxx_EjectOperation{
		This:               o.This,
		DiskID:             o.DiskID,
		DiskLastKnownState: o.DiskLastKnownState,
	}
}

func (o *EjectRequest) xxx_FromOp(ctx context.Context, op *xxx_EjectOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
	o.DiskLastKnownState = op.DiskLastKnownState
}
func (o *EjectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EjectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EjectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EjectResponse structure represents the Eject operation response
type EjectResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The Eject return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EjectResponse) xxx_ToOp(ctx context.Context) *xxx_EjectOperation {
	if o == nil {
		return &xxx_EjectOperation{}
	}
	return &xxx_EjectOperation{
		That:     o.That,
		TaskInfo: o.TaskInfo,
		Return:   o.Return,
	}
}

func (o *EjectResponse) xxx_FromOp(ctx context.Context, op *xxx_EjectOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *EjectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EjectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EjectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FTEnumVolumesOperation structure represents the FTEnumVolumes operation
type xxx_FTEnumVolumesOperation struct {
	This        *dcom.ORPCThis     `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat     `idl:"name:That" json:"that"`
	VolumeCount uint32             `idl:"name:volumeCount" json:"volume_count"`
	VolumeList  []*dmrp.VolumeInfo `idl:"name:ftVolumeList;size_is:(, volumeCount)" json:"volume_list"`
	Return      int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_FTEnumVolumesOperation) OpNum() int { return 13 }

func (o *xxx_FTEnumVolumesOperation) OpName() string { return "/IVolumeClient/v0/FTEnumVolumes" }

func (o *xxx_FTEnumVolumesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTEnumVolumesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// volumeCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.VolumeCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTEnumVolumesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// volumeCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.VolumeCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTEnumVolumesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.VolumeList != nil && o.VolumeCount == 0 {
		o.VolumeCount = uint32(len(o.VolumeList))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTEnumVolumesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// volumeCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.VolumeCount); err != nil {
			return err
		}
	}
	// ftVolumeList {out} (1:{pointer=ref}*(2)*(1))(2:{alias=VOLUME_INFO}[dim:0,size_is=volumeCount](struct))
	{
		if o.VolumeList != nil || o.VolumeCount > 0 {
			_ptr_ftVolumeList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.VolumeCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.VolumeList {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.VolumeList[i1] != nil {
						if err := o.VolumeList[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dmrp.VolumeInfo{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.VolumeList); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&dmrp.VolumeInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.VolumeList, _ptr_ftVolumeList); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTEnumVolumesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// volumeCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.VolumeCount); err != nil {
			return err
		}
	}
	// ftVolumeList {out} (1:{pointer=ref}*(2)*(1))(2:{alias=VOLUME_INFO}[dim:0,size_is=volumeCount](struct))
	{
		_ptr_ftVolumeList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.VolumeList", sizeInfo[0])
			}
			o.VolumeList = make([]*dmrp.VolumeInfo, sizeInfo[0])
			for i1 := range o.VolumeList {
				i1 := i1
				if o.VolumeList[i1] == nil {
					o.VolumeList[i1] = &dmrp.VolumeInfo{}
				}
				if err := o.VolumeList[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ftVolumeList := func(ptr interface{}) { o.VolumeList = *ptr.(*[]*dmrp.VolumeInfo) }
		if err := w.ReadPointer(&o.VolumeList, _s_ftVolumeList, _ptr_ftVolumeList); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// FTEnumVolumesRequest structure represents the FTEnumVolumes operation request
type FTEnumVolumesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	VolumeCount uint32         `idl:"name:volumeCount" json:"volume_count"`
}

func (o *FTEnumVolumesRequest) xxx_ToOp(ctx context.Context) *xxx_FTEnumVolumesOperation {
	if o == nil {
		return &xxx_FTEnumVolumesOperation{}
	}
	return &xxx_FTEnumVolumesOperation{
		This:        o.This,
		VolumeCount: o.VolumeCount,
	}
}

func (o *FTEnumVolumesRequest) xxx_FromOp(ctx context.Context, op *xxx_FTEnumVolumesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VolumeCount = op.VolumeCount
}
func (o *FTEnumVolumesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *FTEnumVolumesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FTEnumVolumesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FTEnumVolumesResponse structure represents the FTEnumVolumes operation response
type FTEnumVolumesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat     `idl:"name:That" json:"that"`
	VolumeCount uint32             `idl:"name:volumeCount" json:"volume_count"`
	VolumeList  []*dmrp.VolumeInfo `idl:"name:ftVolumeList;size_is:(, volumeCount)" json:"volume_list"`
	// Return: The FTEnumVolumes return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *FTEnumVolumesResponse) xxx_ToOp(ctx context.Context) *xxx_FTEnumVolumesOperation {
	if o == nil {
		return &xxx_FTEnumVolumesOperation{}
	}
	return &xxx_FTEnumVolumesOperation{
		That:        o.That,
		VolumeCount: o.VolumeCount,
		VolumeList:  o.VolumeList,
		Return:      o.Return,
	}
}

func (o *FTEnumVolumesResponse) xxx_FromOp(ctx context.Context, op *xxx_FTEnumVolumesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.VolumeCount = op.VolumeCount
	o.VolumeList = op.VolumeList
	o.Return = op.Return
}
func (o *FTEnumVolumesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *FTEnumVolumesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FTEnumVolumesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FTEnumLogicalDiskMembersOperation structure represents the FTEnumLogicalDiskMembers operation
type xxx_FTEnumLogicalDiskMembersOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	VolumeID    int64          `idl:"name:volumeId" json:"volume_id"`
	MemberCount uint32         `idl:"name:memberCount" json:"member_count"`
	MemberList  []int64        `idl:"name:memberList;size_is:(, memberCount)" json:"member_list"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_FTEnumLogicalDiskMembersOperation) OpNum() int { return 14 }

func (o *xxx_FTEnumLogicalDiskMembersOperation) OpName() string {
	return "/IVolumeClient/v0/FTEnumLogicalDiskMembers"
}

func (o *xxx_FTEnumLogicalDiskMembersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTEnumLogicalDiskMembersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.VolumeID); err != nil {
			return err
		}
	}
	// memberCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.MemberCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTEnumLogicalDiskMembersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.VolumeID); err != nil {
			return err
		}
	}
	// memberCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.MemberCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTEnumLogicalDiskMembersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.MemberList != nil && o.MemberCount == 0 {
		o.MemberCount = uint32(len(o.MemberList))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTEnumLogicalDiskMembersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// memberCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.MemberCount); err != nil {
			return err
		}
	}
	// memberList {out} (1:{pointer=ref}*(2)*(1))(2:{alias=LdmObjectId, names=LONGLONG}[dim:0,size_is=memberCount](int64))
	{
		if o.MemberList != nil || o.MemberCount > 0 {
			_ptr_memberList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.MemberCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.MemberList {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.MemberList[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.MemberList); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(int64(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MemberList, _ptr_memberList); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTEnumLogicalDiskMembersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// memberCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.MemberCount); err != nil {
			return err
		}
	}
	// memberList {out} (1:{pointer=ref}*(2)*(1))(2:{alias=LdmObjectId, names=LONGLONG}[dim:0,size_is=memberCount](int64))
	{
		_ptr_memberList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.MemberList", sizeInfo[0])
			}
			o.MemberList = make([]int64, sizeInfo[0])
			for i1 := range o.MemberList {
				i1 := i1
				if err := w.ReadData(&o.MemberList[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_memberList := func(ptr interface{}) { o.MemberList = *ptr.(*[]int64) }
		if err := w.ReadPointer(&o.MemberList, _s_memberList, _ptr_memberList); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// FTEnumLogicalDiskMembersRequest structure represents the FTEnumLogicalDiskMembers operation request
type FTEnumLogicalDiskMembersRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	VolumeID    int64          `idl:"name:volumeId" json:"volume_id"`
	MemberCount uint32         `idl:"name:memberCount" json:"member_count"`
}

func (o *FTEnumLogicalDiskMembersRequest) xxx_ToOp(ctx context.Context) *xxx_FTEnumLogicalDiskMembersOperation {
	if o == nil {
		return &xxx_FTEnumLogicalDiskMembersOperation{}
	}
	return &xxx_FTEnumLogicalDiskMembersOperation{
		This:        o.This,
		VolumeID:    o.VolumeID,
		MemberCount: o.MemberCount,
	}
}

func (o *FTEnumLogicalDiskMembersRequest) xxx_FromOp(ctx context.Context, op *xxx_FTEnumLogicalDiskMembersOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VolumeID = op.VolumeID
	o.MemberCount = op.MemberCount
}
func (o *FTEnumLogicalDiskMembersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *FTEnumLogicalDiskMembersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FTEnumLogicalDiskMembersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FTEnumLogicalDiskMembersResponse structure represents the FTEnumLogicalDiskMembers operation response
type FTEnumLogicalDiskMembersResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	MemberCount uint32         `idl:"name:memberCount" json:"member_count"`
	MemberList  []int64        `idl:"name:memberList;size_is:(, memberCount)" json:"member_list"`
	// Return: The FTEnumLogicalDiskMembers return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *FTEnumLogicalDiskMembersResponse) xxx_ToOp(ctx context.Context) *xxx_FTEnumLogicalDiskMembersOperation {
	if o == nil {
		return &xxx_FTEnumLogicalDiskMembersOperation{}
	}
	return &xxx_FTEnumLogicalDiskMembersOperation{
		That:        o.That,
		MemberCount: o.MemberCount,
		MemberList:  o.MemberList,
		Return:      o.Return,
	}
}

func (o *FTEnumLogicalDiskMembersResponse) xxx_FromOp(ctx context.Context, op *xxx_FTEnumLogicalDiskMembersOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MemberCount = op.MemberCount
	o.MemberList = op.MemberList
	o.Return = op.Return
}
func (o *FTEnumLogicalDiskMembersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *FTEnumLogicalDiskMembersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FTEnumLogicalDiskMembersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FTDeleteVolumeOperation structure represents the FTDeleteVolume operation
type xxx_FTDeleteVolumeOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	VolumeID             int64          `idl:"name:volumeId" json:"volume_id"`
	Force                bool           `idl:"name:force" json:"force"`
	VolumeLastKnownState int64          `idl:"name:volumeLastKnownState" json:"volume_last_known_state"`
	TaskInfo             *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_FTDeleteVolumeOperation) OpNum() int { return 15 }

func (o *xxx_FTDeleteVolumeOperation) OpName() string { return "/IVolumeClient/v0/FTDeleteVolume" }

func (o *xxx_FTDeleteVolumeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTDeleteVolumeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.VolumeID); err != nil {
			return err
		}
	}
	// force {in} (1:(bool))
	{
		if err := w.WriteData(o.Force); err != nil {
			return err
		}
	}
	// volumeLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.VolumeLastKnownState); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTDeleteVolumeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.VolumeID); err != nil {
			return err
		}
	}
	// force {in} (1:(bool))
	{
		if err := w.ReadData(&o.Force); err != nil {
			return err
		}
	}
	// volumeLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.VolumeLastKnownState); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTDeleteVolumeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTDeleteVolumeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTDeleteVolumeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// FTDeleteVolumeRequest structure represents the FTDeleteVolume operation request
type FTDeleteVolumeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	VolumeID             int64          `idl:"name:volumeId" json:"volume_id"`
	Force                bool           `idl:"name:force" json:"force"`
	VolumeLastKnownState int64          `idl:"name:volumeLastKnownState" json:"volume_last_known_state"`
}

func (o *FTDeleteVolumeRequest) xxx_ToOp(ctx context.Context) *xxx_FTDeleteVolumeOperation {
	if o == nil {
		return &xxx_FTDeleteVolumeOperation{}
	}
	return &xxx_FTDeleteVolumeOperation{
		This:                 o.This,
		VolumeID:             o.VolumeID,
		Force:                o.Force,
		VolumeLastKnownState: o.VolumeLastKnownState,
	}
}

func (o *FTDeleteVolumeRequest) xxx_FromOp(ctx context.Context, op *xxx_FTDeleteVolumeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VolumeID = op.VolumeID
	o.Force = op.Force
	o.VolumeLastKnownState = op.VolumeLastKnownState
}
func (o *FTDeleteVolumeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *FTDeleteVolumeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FTDeleteVolumeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FTDeleteVolumeResponse structure represents the FTDeleteVolume operation response
type FTDeleteVolumeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The FTDeleteVolume return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *FTDeleteVolumeResponse) xxx_ToOp(ctx context.Context) *xxx_FTDeleteVolumeOperation {
	if o == nil {
		return &xxx_FTDeleteVolumeOperation{}
	}
	return &xxx_FTDeleteVolumeOperation{
		That:     o.That,
		TaskInfo: o.TaskInfo,
		Return:   o.Return,
	}
}

func (o *FTDeleteVolumeResponse) xxx_FromOp(ctx context.Context, op *xxx_FTDeleteVolumeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *FTDeleteVolumeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *FTDeleteVolumeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FTDeleteVolumeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FTBreakMirrorOperation structure represents the FTBreakMirror operation
type xxx_FTBreakMirrorOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	VolumeID             int64          `idl:"name:volumeId" json:"volume_id"`
	VolumeLastKnownState int64          `idl:"name:volumeLastKnownState" json:"volume_last_known_state"`
	Force                bool           `idl:"name:bForce" json:"force"`
	TaskInfo             *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_FTBreakMirrorOperation) OpNum() int { return 16 }

func (o *xxx_FTBreakMirrorOperation) OpName() string { return "/IVolumeClient/v0/FTBreakMirror" }

func (o *xxx_FTBreakMirrorOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTBreakMirrorOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.VolumeID); err != nil {
			return err
		}
	}
	// volumeLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.VolumeLastKnownState); err != nil {
			return err
		}
	}
	// bForce {in} (1:(bool))
	{
		if err := w.WriteData(o.Force); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTBreakMirrorOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.VolumeID); err != nil {
			return err
		}
	}
	// volumeLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.VolumeLastKnownState); err != nil {
			return err
		}
	}
	// bForce {in} (1:(bool))
	{
		if err := w.ReadData(&o.Force); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTBreakMirrorOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTBreakMirrorOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTBreakMirrorOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// FTBreakMirrorRequest structure represents the FTBreakMirror operation request
type FTBreakMirrorRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	VolumeID             int64          `idl:"name:volumeId" json:"volume_id"`
	VolumeLastKnownState int64          `idl:"name:volumeLastKnownState" json:"volume_last_known_state"`
	Force                bool           `idl:"name:bForce" json:"force"`
}

func (o *FTBreakMirrorRequest) xxx_ToOp(ctx context.Context) *xxx_FTBreakMirrorOperation {
	if o == nil {
		return &xxx_FTBreakMirrorOperation{}
	}
	return &xxx_FTBreakMirrorOperation{
		This:                 o.This,
		VolumeID:             o.VolumeID,
		VolumeLastKnownState: o.VolumeLastKnownState,
		Force:                o.Force,
	}
}

func (o *FTBreakMirrorRequest) xxx_FromOp(ctx context.Context, op *xxx_FTBreakMirrorOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VolumeID = op.VolumeID
	o.VolumeLastKnownState = op.VolumeLastKnownState
	o.Force = op.Force
}
func (o *FTBreakMirrorRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *FTBreakMirrorRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FTBreakMirrorOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FTBreakMirrorResponse structure represents the FTBreakMirror operation response
type FTBreakMirrorResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The FTBreakMirror return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *FTBreakMirrorResponse) xxx_ToOp(ctx context.Context) *xxx_FTBreakMirrorOperation {
	if o == nil {
		return &xxx_FTBreakMirrorOperation{}
	}
	return &xxx_FTBreakMirrorOperation{
		That:     o.That,
		TaskInfo: o.TaskInfo,
		Return:   o.Return,
	}
}

func (o *FTBreakMirrorResponse) xxx_FromOp(ctx context.Context, op *xxx_FTBreakMirrorOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *FTBreakMirrorResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *FTBreakMirrorResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FTBreakMirrorOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FTResyncMirrorOperation structure represents the FTResyncMirror operation
type xxx_FTResyncMirrorOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	VolumeID             int64          `idl:"name:volumeId" json:"volume_id"`
	VolumeLastKnownState int64          `idl:"name:volumeLastKnownState" json:"volume_last_known_state"`
	TaskInfo             *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_FTResyncMirrorOperation) OpNum() int { return 17 }

func (o *xxx_FTResyncMirrorOperation) OpName() string { return "/IVolumeClient/v0/FTResyncMirror" }

func (o *xxx_FTResyncMirrorOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTResyncMirrorOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.VolumeID); err != nil {
			return err
		}
	}
	// volumeLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.VolumeLastKnownState); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTResyncMirrorOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.VolumeID); err != nil {
			return err
		}
	}
	// volumeLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.VolumeLastKnownState); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTResyncMirrorOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTResyncMirrorOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTResyncMirrorOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// FTResyncMirrorRequest structure represents the FTResyncMirror operation request
type FTResyncMirrorRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	VolumeID             int64          `idl:"name:volumeId" json:"volume_id"`
	VolumeLastKnownState int64          `idl:"name:volumeLastKnownState" json:"volume_last_known_state"`
}

func (o *FTResyncMirrorRequest) xxx_ToOp(ctx context.Context) *xxx_FTResyncMirrorOperation {
	if o == nil {
		return &xxx_FTResyncMirrorOperation{}
	}
	return &xxx_FTResyncMirrorOperation{
		This:                 o.This,
		VolumeID:             o.VolumeID,
		VolumeLastKnownState: o.VolumeLastKnownState,
	}
}

func (o *FTResyncMirrorRequest) xxx_FromOp(ctx context.Context, op *xxx_FTResyncMirrorOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VolumeID = op.VolumeID
	o.VolumeLastKnownState = op.VolumeLastKnownState
}
func (o *FTResyncMirrorRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *FTResyncMirrorRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FTResyncMirrorOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FTResyncMirrorResponse structure represents the FTResyncMirror operation response
type FTResyncMirrorResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The FTResyncMirror return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *FTResyncMirrorResponse) xxx_ToOp(ctx context.Context) *xxx_FTResyncMirrorOperation {
	if o == nil {
		return &xxx_FTResyncMirrorOperation{}
	}
	return &xxx_FTResyncMirrorOperation{
		That:     o.That,
		TaskInfo: o.TaskInfo,
		Return:   o.Return,
	}
}

func (o *FTResyncMirrorResponse) xxx_FromOp(ctx context.Context, op *xxx_FTResyncMirrorOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *FTResyncMirrorResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *FTResyncMirrorResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FTResyncMirrorOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FTRegenerateParityStripeOperation structure represents the FTRegenerateParityStripe operation
type xxx_FTRegenerateParityStripeOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	VolumeID             int64          `idl:"name:volumeId" json:"volume_id"`
	VolumeLastKnownState int64          `idl:"name:volumeLastKnownState" json:"volume_last_known_state"`
	TaskInfo             *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_FTRegenerateParityStripeOperation) OpNum() int { return 18 }

func (o *xxx_FTRegenerateParityStripeOperation) OpName() string {
	return "/IVolumeClient/v0/FTRegenerateParityStripe"
}

func (o *xxx_FTRegenerateParityStripeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTRegenerateParityStripeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.VolumeID); err != nil {
			return err
		}
	}
	// volumeLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.VolumeLastKnownState); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTRegenerateParityStripeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.VolumeID); err != nil {
			return err
		}
	}
	// volumeLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.VolumeLastKnownState); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTRegenerateParityStripeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTRegenerateParityStripeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTRegenerateParityStripeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// FTRegenerateParityStripeRequest structure represents the FTRegenerateParityStripe operation request
type FTRegenerateParityStripeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	VolumeID             int64          `idl:"name:volumeId" json:"volume_id"`
	VolumeLastKnownState int64          `idl:"name:volumeLastKnownState" json:"volume_last_known_state"`
}

func (o *FTRegenerateParityStripeRequest) xxx_ToOp(ctx context.Context) *xxx_FTRegenerateParityStripeOperation {
	if o == nil {
		return &xxx_FTRegenerateParityStripeOperation{}
	}
	return &xxx_FTRegenerateParityStripeOperation{
		This:                 o.This,
		VolumeID:             o.VolumeID,
		VolumeLastKnownState: o.VolumeLastKnownState,
	}
}

func (o *FTRegenerateParityStripeRequest) xxx_FromOp(ctx context.Context, op *xxx_FTRegenerateParityStripeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VolumeID = op.VolumeID
	o.VolumeLastKnownState = op.VolumeLastKnownState
}
func (o *FTRegenerateParityStripeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *FTRegenerateParityStripeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FTRegenerateParityStripeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FTRegenerateParityStripeResponse structure represents the FTRegenerateParityStripe operation response
type FTRegenerateParityStripeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The FTRegenerateParityStripe return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *FTRegenerateParityStripeResponse) xxx_ToOp(ctx context.Context) *xxx_FTRegenerateParityStripeOperation {
	if o == nil {
		return &xxx_FTRegenerateParityStripeOperation{}
	}
	return &xxx_FTRegenerateParityStripeOperation{
		That:     o.That,
		TaskInfo: o.TaskInfo,
		Return:   o.Return,
	}
}

func (o *FTRegenerateParityStripeResponse) xxx_FromOp(ctx context.Context, op *xxx_FTRegenerateParityStripeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *FTRegenerateParityStripeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *FTRegenerateParityStripeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FTRegenerateParityStripeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FTReplaceMirrorPartitionOperation structure represents the FTReplaceMirrorPartition operation
type xxx_FTReplaceMirrorPartitionOperation struct {
	This                    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                    *dcom.ORPCThat `idl:"name:That" json:"that"`
	VolumeID                int64          `idl:"name:volumeId" json:"volume_id"`
	VolumeLastKnownState    int64          `idl:"name:volumeLastKnownState" json:"volume_last_known_state"`
	OldMemberID             int64          `idl:"name:oldMemberId" json:"old_member_id"`
	OldMemberLastKnownState int64          `idl:"name:oldMemberLastKnownState" json:"old_member_last_known_state"`
	NewRegionID             int64          `idl:"name:newRegionId" json:"new_region_id"`
	NewRegionLastKnownState int64          `idl:"name:newRegionLastKnownState" json:"new_region_last_known_state"`
	Flags                   uint32         `idl:"name:flags" json:"flags"`
	TaskInfo                *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	Return                  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_FTReplaceMirrorPartitionOperation) OpNum() int { return 19 }

func (o *xxx_FTReplaceMirrorPartitionOperation) OpName() string {
	return "/IVolumeClient/v0/FTReplaceMirrorPartition"
}

func (o *xxx_FTReplaceMirrorPartitionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTReplaceMirrorPartitionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.VolumeID); err != nil {
			return err
		}
	}
	// volumeLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.VolumeLastKnownState); err != nil {
			return err
		}
	}
	// oldMemberId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.OldMemberID); err != nil {
			return err
		}
	}
	// oldMemberLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.OldMemberLastKnownState); err != nil {
			return err
		}
	}
	// newRegionId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.NewRegionID); err != nil {
			return err
		}
	}
	// newRegionLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.NewRegionLastKnownState); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTReplaceMirrorPartitionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.VolumeID); err != nil {
			return err
		}
	}
	// volumeLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.VolumeLastKnownState); err != nil {
			return err
		}
	}
	// oldMemberId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.OldMemberID); err != nil {
			return err
		}
	}
	// oldMemberLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.OldMemberLastKnownState); err != nil {
			return err
		}
	}
	// newRegionId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.NewRegionID); err != nil {
			return err
		}
	}
	// newRegionLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.NewRegionLastKnownState); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTReplaceMirrorPartitionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTReplaceMirrorPartitionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTReplaceMirrorPartitionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// FTReplaceMirrorPartitionRequest structure represents the FTReplaceMirrorPartition operation request
type FTReplaceMirrorPartitionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                    *dcom.ORPCThis `idl:"name:This" json:"this"`
	VolumeID                int64          `idl:"name:volumeId" json:"volume_id"`
	VolumeLastKnownState    int64          `idl:"name:volumeLastKnownState" json:"volume_last_known_state"`
	OldMemberID             int64          `idl:"name:oldMemberId" json:"old_member_id"`
	OldMemberLastKnownState int64          `idl:"name:oldMemberLastKnownState" json:"old_member_last_known_state"`
	NewRegionID             int64          `idl:"name:newRegionId" json:"new_region_id"`
	NewRegionLastKnownState int64          `idl:"name:newRegionLastKnownState" json:"new_region_last_known_state"`
	Flags                   uint32         `idl:"name:flags" json:"flags"`
}

func (o *FTReplaceMirrorPartitionRequest) xxx_ToOp(ctx context.Context) *xxx_FTReplaceMirrorPartitionOperation {
	if o == nil {
		return &xxx_FTReplaceMirrorPartitionOperation{}
	}
	return &xxx_FTReplaceMirrorPartitionOperation{
		This:                    o.This,
		VolumeID:                o.VolumeID,
		VolumeLastKnownState:    o.VolumeLastKnownState,
		OldMemberID:             o.OldMemberID,
		OldMemberLastKnownState: o.OldMemberLastKnownState,
		NewRegionID:             o.NewRegionID,
		NewRegionLastKnownState: o.NewRegionLastKnownState,
		Flags:                   o.Flags,
	}
}

func (o *FTReplaceMirrorPartitionRequest) xxx_FromOp(ctx context.Context, op *xxx_FTReplaceMirrorPartitionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VolumeID = op.VolumeID
	o.VolumeLastKnownState = op.VolumeLastKnownState
	o.OldMemberID = op.OldMemberID
	o.OldMemberLastKnownState = op.OldMemberLastKnownState
	o.NewRegionID = op.NewRegionID
	o.NewRegionLastKnownState = op.NewRegionLastKnownState
	o.Flags = op.Flags
}
func (o *FTReplaceMirrorPartitionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *FTReplaceMirrorPartitionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FTReplaceMirrorPartitionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FTReplaceMirrorPartitionResponse structure represents the FTReplaceMirrorPartition operation response
type FTReplaceMirrorPartitionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The FTReplaceMirrorPartition return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *FTReplaceMirrorPartitionResponse) xxx_ToOp(ctx context.Context) *xxx_FTReplaceMirrorPartitionOperation {
	if o == nil {
		return &xxx_FTReplaceMirrorPartitionOperation{}
	}
	return &xxx_FTReplaceMirrorPartitionOperation{
		That:     o.That,
		TaskInfo: o.TaskInfo,
		Return:   o.Return,
	}
}

func (o *FTReplaceMirrorPartitionResponse) xxx_FromOp(ctx context.Context, op *xxx_FTReplaceMirrorPartitionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *FTReplaceMirrorPartitionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *FTReplaceMirrorPartitionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FTReplaceMirrorPartitionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FTReplaceParityStripePartitionOperation structure represents the FTReplaceParityStripePartition operation
type xxx_FTReplaceParityStripePartitionOperation struct {
	This                    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                    *dcom.ORPCThat `idl:"name:That" json:"that"`
	VolumeID                int64          `idl:"name:volumeId" json:"volume_id"`
	VolumeLastKnownState    int64          `idl:"name:volumeLastKnownState" json:"volume_last_known_state"`
	OldMemberID             int64          `idl:"name:oldMemberId" json:"old_member_id"`
	OldMemberLastKnownState int64          `idl:"name:oldMemberLastKnownState" json:"old_member_last_known_state"`
	NewRegionID             int64          `idl:"name:newRegionId" json:"new_region_id"`
	NewRegionLastKnownState int64          `idl:"name:newRegionLastKnownState" json:"new_region_last_known_state"`
	Flags                   uint32         `idl:"name:flags" json:"flags"`
	TaskInfo                *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	Return                  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_FTReplaceParityStripePartitionOperation) OpNum() int { return 20 }

func (o *xxx_FTReplaceParityStripePartitionOperation) OpName() string {
	return "/IVolumeClient/v0/FTReplaceParityStripePartition"
}

func (o *xxx_FTReplaceParityStripePartitionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTReplaceParityStripePartitionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.VolumeID); err != nil {
			return err
		}
	}
	// volumeLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.VolumeLastKnownState); err != nil {
			return err
		}
	}
	// oldMemberId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.OldMemberID); err != nil {
			return err
		}
	}
	// oldMemberLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.OldMemberLastKnownState); err != nil {
			return err
		}
	}
	// newRegionId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.NewRegionID); err != nil {
			return err
		}
	}
	// newRegionLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.NewRegionLastKnownState); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTReplaceParityStripePartitionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.VolumeID); err != nil {
			return err
		}
	}
	// volumeLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.VolumeLastKnownState); err != nil {
			return err
		}
	}
	// oldMemberId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.OldMemberID); err != nil {
			return err
		}
	}
	// oldMemberLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.OldMemberLastKnownState); err != nil {
			return err
		}
	}
	// newRegionId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.NewRegionID); err != nil {
			return err
		}
	}
	// newRegionLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.NewRegionLastKnownState); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTReplaceParityStripePartitionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTReplaceParityStripePartitionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTReplaceParityStripePartitionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// FTReplaceParityStripePartitionRequest structure represents the FTReplaceParityStripePartition operation request
type FTReplaceParityStripePartitionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                    *dcom.ORPCThis `idl:"name:This" json:"this"`
	VolumeID                int64          `idl:"name:volumeId" json:"volume_id"`
	VolumeLastKnownState    int64          `idl:"name:volumeLastKnownState" json:"volume_last_known_state"`
	OldMemberID             int64          `idl:"name:oldMemberId" json:"old_member_id"`
	OldMemberLastKnownState int64          `idl:"name:oldMemberLastKnownState" json:"old_member_last_known_state"`
	NewRegionID             int64          `idl:"name:newRegionId" json:"new_region_id"`
	NewRegionLastKnownState int64          `idl:"name:newRegionLastKnownState" json:"new_region_last_known_state"`
	Flags                   uint32         `idl:"name:flags" json:"flags"`
}

func (o *FTReplaceParityStripePartitionRequest) xxx_ToOp(ctx context.Context) *xxx_FTReplaceParityStripePartitionOperation {
	if o == nil {
		return &xxx_FTReplaceParityStripePartitionOperation{}
	}
	return &xxx_FTReplaceParityStripePartitionOperation{
		This:                    o.This,
		VolumeID:                o.VolumeID,
		VolumeLastKnownState:    o.VolumeLastKnownState,
		OldMemberID:             o.OldMemberID,
		OldMemberLastKnownState: o.OldMemberLastKnownState,
		NewRegionID:             o.NewRegionID,
		NewRegionLastKnownState: o.NewRegionLastKnownState,
		Flags:                   o.Flags,
	}
}

func (o *FTReplaceParityStripePartitionRequest) xxx_FromOp(ctx context.Context, op *xxx_FTReplaceParityStripePartitionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VolumeID = op.VolumeID
	o.VolumeLastKnownState = op.VolumeLastKnownState
	o.OldMemberID = op.OldMemberID
	o.OldMemberLastKnownState = op.OldMemberLastKnownState
	o.NewRegionID = op.NewRegionID
	o.NewRegionLastKnownState = op.NewRegionLastKnownState
	o.Flags = op.Flags
}
func (o *FTReplaceParityStripePartitionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *FTReplaceParityStripePartitionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FTReplaceParityStripePartitionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FTReplaceParityStripePartitionResponse structure represents the FTReplaceParityStripePartition operation response
type FTReplaceParityStripePartitionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The FTReplaceParityStripePartition return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *FTReplaceParityStripePartitionResponse) xxx_ToOp(ctx context.Context) *xxx_FTReplaceParityStripePartitionOperation {
	if o == nil {
		return &xxx_FTReplaceParityStripePartitionOperation{}
	}
	return &xxx_FTReplaceParityStripePartitionOperation{
		That:     o.That,
		TaskInfo: o.TaskInfo,
		Return:   o.Return,
	}
}

func (o *FTReplaceParityStripePartitionResponse) xxx_FromOp(ctx context.Context, op *xxx_FTReplaceParityStripePartitionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *FTReplaceParityStripePartitionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *FTReplaceParityStripePartitionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FTReplaceParityStripePartitionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumDriveLettersOperation structure represents the EnumDriveLetters operation
type xxx_EnumDriveLettersOperation struct {
	This             *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat          `idl:"name:That" json:"that"`
	DriveLetterCount uint32                  `idl:"name:driveLetterCount" json:"drive_letter_count"`
	DriveLetterList  []*dmrp.DriveLetterInfo `idl:"name:driveLetterList;size_is:(, driveLetterCount)" json:"drive_letter_list"`
	Return           int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumDriveLettersOperation) OpNum() int { return 21 }

func (o *xxx_EnumDriveLettersOperation) OpName() string { return "/IVolumeClient/v0/EnumDriveLetters" }

func (o *xxx_EnumDriveLettersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumDriveLettersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// driveLetterCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.DriveLetterCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumDriveLettersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// driveLetterCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.DriveLetterCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumDriveLettersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.DriveLetterList != nil && o.DriveLetterCount == 0 {
		o.DriveLetterCount = uint32(len(o.DriveLetterList))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumDriveLettersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// driveLetterCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.DriveLetterCount); err != nil {
			return err
		}
	}
	// driveLetterList {out} (1:{pointer=ref}*(2)*(1))(2:{alias=DRIVE_LETTER_INFO}[dim:0,size_is=driveLetterCount](struct))
	{
		if o.DriveLetterList != nil || o.DriveLetterCount > 0 {
			_ptr_driveLetterList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.DriveLetterCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.DriveLetterList {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.DriveLetterList[i1] != nil {
						if err := o.DriveLetterList[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dmrp.DriveLetterInfo{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.DriveLetterList); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&dmrp.DriveLetterInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DriveLetterList, _ptr_driveLetterList); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumDriveLettersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// driveLetterCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.DriveLetterCount); err != nil {
			return err
		}
	}
	// driveLetterList {out} (1:{pointer=ref}*(2)*(1))(2:{alias=DRIVE_LETTER_INFO}[dim:0,size_is=driveLetterCount](struct))
	{
		_ptr_driveLetterList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.DriveLetterList", sizeInfo[0])
			}
			o.DriveLetterList = make([]*dmrp.DriveLetterInfo, sizeInfo[0])
			for i1 := range o.DriveLetterList {
				i1 := i1
				if o.DriveLetterList[i1] == nil {
					o.DriveLetterList[i1] = &dmrp.DriveLetterInfo{}
				}
				if err := o.DriveLetterList[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_driveLetterList := func(ptr interface{}) { o.DriveLetterList = *ptr.(*[]*dmrp.DriveLetterInfo) }
		if err := w.ReadPointer(&o.DriveLetterList, _s_driveLetterList, _ptr_driveLetterList); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumDriveLettersRequest structure represents the EnumDriveLetters operation request
type EnumDriveLettersRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	DriveLetterCount uint32         `idl:"name:driveLetterCount" json:"drive_letter_count"`
}

func (o *EnumDriveLettersRequest) xxx_ToOp(ctx context.Context) *xxx_EnumDriveLettersOperation {
	if o == nil {
		return &xxx_EnumDriveLettersOperation{}
	}
	return &xxx_EnumDriveLettersOperation{
		This:             o.This,
		DriveLetterCount: o.DriveLetterCount,
	}
}

func (o *EnumDriveLettersRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumDriveLettersOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DriveLetterCount = op.DriveLetterCount
}
func (o *EnumDriveLettersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumDriveLettersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumDriveLettersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumDriveLettersResponse structure represents the EnumDriveLetters operation response
type EnumDriveLettersResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That             *dcom.ORPCThat          `idl:"name:That" json:"that"`
	DriveLetterCount uint32                  `idl:"name:driveLetterCount" json:"drive_letter_count"`
	DriveLetterList  []*dmrp.DriveLetterInfo `idl:"name:driveLetterList;size_is:(, driveLetterCount)" json:"drive_letter_list"`
	// Return: The EnumDriveLetters return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumDriveLettersResponse) xxx_ToOp(ctx context.Context) *xxx_EnumDriveLettersOperation {
	if o == nil {
		return &xxx_EnumDriveLettersOperation{}
	}
	return &xxx_EnumDriveLettersOperation{
		That:             o.That,
		DriveLetterCount: o.DriveLetterCount,
		DriveLetterList:  o.DriveLetterList,
		Return:           o.Return,
	}
}

func (o *EnumDriveLettersResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumDriveLettersOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DriveLetterCount = op.DriveLetterCount
	o.DriveLetterList = op.DriveLetterList
	o.Return = op.Return
}
func (o *EnumDriveLettersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumDriveLettersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumDriveLettersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AssignDriveLetterOperation structure represents the AssignDriveLetter operation
type xxx_AssignDriveLetterOperation struct {
	This                  *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                  *dcom.ORPCThat `idl:"name:That" json:"that"`
	Letter                uint16         `idl:"name:letter" json:"letter"`
	ForceOption           uint32         `idl:"name:forceOption" json:"force_option"`
	LetterLastKnownState  int64          `idl:"name:letterLastKnownState" json:"letter_last_known_state"`
	StorageID             int64          `idl:"name:storageId" json:"storage_id"`
	StorageLastKnownState int64          `idl:"name:storageLastKnownState" json:"storage_last_known_state"`
	TaskInfo              *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	Return                int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_AssignDriveLetterOperation) OpNum() int { return 22 }

func (o *xxx_AssignDriveLetterOperation) OpName() string {
	return "/IVolumeClient/v0/AssignDriveLetter"
}

func (o *xxx_AssignDriveLetterOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AssignDriveLetterOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// letter {in} (1:(wchar))
	{
		if err := w.WriteData(o.Letter); err != nil {
			return err
		}
	}
	// forceOption {in} (1:(uint32))
	{
		if err := w.WriteData(o.ForceOption); err != nil {
			return err
		}
	}
	// letterLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.LetterLastKnownState); err != nil {
			return err
		}
	}
	// storageId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.StorageID); err != nil {
			return err
		}
	}
	// storageLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.StorageLastKnownState); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AssignDriveLetterOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// letter {in} (1:(wchar))
	{
		if err := w.ReadData(&o.Letter); err != nil {
			return err
		}
	}
	// forceOption {in} (1:(uint32))
	{
		if err := w.ReadData(&o.ForceOption); err != nil {
			return err
		}
	}
	// letterLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.LetterLastKnownState); err != nil {
			return err
		}
	}
	// storageId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.StorageID); err != nil {
			return err
		}
	}
	// storageLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.StorageLastKnownState); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AssignDriveLetterOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AssignDriveLetterOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AssignDriveLetterOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AssignDriveLetterRequest structure represents the AssignDriveLetter operation request
type AssignDriveLetterRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                  *dcom.ORPCThis `idl:"name:This" json:"this"`
	Letter                uint16         `idl:"name:letter" json:"letter"`
	ForceOption           uint32         `idl:"name:forceOption" json:"force_option"`
	LetterLastKnownState  int64          `idl:"name:letterLastKnownState" json:"letter_last_known_state"`
	StorageID             int64          `idl:"name:storageId" json:"storage_id"`
	StorageLastKnownState int64          `idl:"name:storageLastKnownState" json:"storage_last_known_state"`
}

func (o *AssignDriveLetterRequest) xxx_ToOp(ctx context.Context) *xxx_AssignDriveLetterOperation {
	if o == nil {
		return &xxx_AssignDriveLetterOperation{}
	}
	return &xxx_AssignDriveLetterOperation{
		This:                  o.This,
		Letter:                o.Letter,
		ForceOption:           o.ForceOption,
		LetterLastKnownState:  o.LetterLastKnownState,
		StorageID:             o.StorageID,
		StorageLastKnownState: o.StorageLastKnownState,
	}
}

func (o *AssignDriveLetterRequest) xxx_FromOp(ctx context.Context, op *xxx_AssignDriveLetterOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Letter = op.Letter
	o.ForceOption = op.ForceOption
	o.LetterLastKnownState = op.LetterLastKnownState
	o.StorageID = op.StorageID
	o.StorageLastKnownState = op.StorageLastKnownState
}
func (o *AssignDriveLetterRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AssignDriveLetterRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AssignDriveLetterOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AssignDriveLetterResponse structure represents the AssignDriveLetter operation response
type AssignDriveLetterResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The AssignDriveLetter return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AssignDriveLetterResponse) xxx_ToOp(ctx context.Context) *xxx_AssignDriveLetterOperation {
	if o == nil {
		return &xxx_AssignDriveLetterOperation{}
	}
	return &xxx_AssignDriveLetterOperation{
		That:     o.That,
		TaskInfo: o.TaskInfo,
		Return:   o.Return,
	}
}

func (o *AssignDriveLetterResponse) xxx_FromOp(ctx context.Context, op *xxx_AssignDriveLetterOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *AssignDriveLetterResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AssignDriveLetterResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AssignDriveLetterOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FreeDriveLetterOperation structure represents the FreeDriveLetter operation
type xxx_FreeDriveLetterOperation struct {
	This                  *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                  *dcom.ORPCThat `idl:"name:That" json:"that"`
	Letter                uint16         `idl:"name:letter" json:"letter"`
	ForceOption           uint32         `idl:"name:forceOption" json:"force_option"`
	LetterLastKnownState  int64          `idl:"name:letterLastKnownState" json:"letter_last_known_state"`
	StorageID             int64          `idl:"name:storageId" json:"storage_id"`
	StorageLastKnownState int64          `idl:"name:storageLastKnownState" json:"storage_last_known_state"`
	TaskInfo              *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	Return                int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_FreeDriveLetterOperation) OpNum() int { return 23 }

func (o *xxx_FreeDriveLetterOperation) OpName() string { return "/IVolumeClient/v0/FreeDriveLetter" }

func (o *xxx_FreeDriveLetterOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FreeDriveLetterOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// letter {in} (1:(wchar))
	{
		if err := w.WriteData(o.Letter); err != nil {
			return err
		}
	}
	// forceOption {in} (1:(uint32))
	{
		if err := w.WriteData(o.ForceOption); err != nil {
			return err
		}
	}
	// letterLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.LetterLastKnownState); err != nil {
			return err
		}
	}
	// storageId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.StorageID); err != nil {
			return err
		}
	}
	// storageLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.StorageLastKnownState); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FreeDriveLetterOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// letter {in} (1:(wchar))
	{
		if err := w.ReadData(&o.Letter); err != nil {
			return err
		}
	}
	// forceOption {in} (1:(uint32))
	{
		if err := w.ReadData(&o.ForceOption); err != nil {
			return err
		}
	}
	// letterLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.LetterLastKnownState); err != nil {
			return err
		}
	}
	// storageId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.StorageID); err != nil {
			return err
		}
	}
	// storageLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.StorageLastKnownState); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FreeDriveLetterOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FreeDriveLetterOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FreeDriveLetterOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// FreeDriveLetterRequest structure represents the FreeDriveLetter operation request
type FreeDriveLetterRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                  *dcom.ORPCThis `idl:"name:This" json:"this"`
	Letter                uint16         `idl:"name:letter" json:"letter"`
	ForceOption           uint32         `idl:"name:forceOption" json:"force_option"`
	LetterLastKnownState  int64          `idl:"name:letterLastKnownState" json:"letter_last_known_state"`
	StorageID             int64          `idl:"name:storageId" json:"storage_id"`
	StorageLastKnownState int64          `idl:"name:storageLastKnownState" json:"storage_last_known_state"`
}

func (o *FreeDriveLetterRequest) xxx_ToOp(ctx context.Context) *xxx_FreeDriveLetterOperation {
	if o == nil {
		return &xxx_FreeDriveLetterOperation{}
	}
	return &xxx_FreeDriveLetterOperation{
		This:                  o.This,
		Letter:                o.Letter,
		ForceOption:           o.ForceOption,
		LetterLastKnownState:  o.LetterLastKnownState,
		StorageID:             o.StorageID,
		StorageLastKnownState: o.StorageLastKnownState,
	}
}

func (o *FreeDriveLetterRequest) xxx_FromOp(ctx context.Context, op *xxx_FreeDriveLetterOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Letter = op.Letter
	o.ForceOption = op.ForceOption
	o.LetterLastKnownState = op.LetterLastKnownState
	o.StorageID = op.StorageID
	o.StorageLastKnownState = op.StorageLastKnownState
}
func (o *FreeDriveLetterRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *FreeDriveLetterRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FreeDriveLetterOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FreeDriveLetterResponse structure represents the FreeDriveLetter operation response
type FreeDriveLetterResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The FreeDriveLetter return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *FreeDriveLetterResponse) xxx_ToOp(ctx context.Context) *xxx_FreeDriveLetterOperation {
	if o == nil {
		return &xxx_FreeDriveLetterOperation{}
	}
	return &xxx_FreeDriveLetterOperation{
		That:     o.That,
		TaskInfo: o.TaskInfo,
		Return:   o.Return,
	}
}

func (o *FreeDriveLetterResponse) xxx_FromOp(ctx context.Context, op *xxx_FreeDriveLetterOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *FreeDriveLetterResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *FreeDriveLetterResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FreeDriveLetterOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumLocalFileSystemsOperation structure represents the EnumLocalFileSystems operation
type xxx_EnumLocalFileSystemsOperation struct {
	This            *dcom.ORPCThis         `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat         `idl:"name:That" json:"that"`
	FileSystemCount uint32                 `idl:"name:fileSystemCount" json:"file_system_count"`
	FileSystemList  []*dmrp.FileSystemInfo `idl:"name:fileSystemList;size_is:(, fileSystemCount)" json:"file_system_list"`
	Return          int32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumLocalFileSystemsOperation) OpNum() int { return 24 }

func (o *xxx_EnumLocalFileSystemsOperation) OpName() string {
	return "/IVolumeClient/v0/EnumLocalFileSystems"
}

func (o *xxx_EnumLocalFileSystemsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumLocalFileSystemsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumLocalFileSystemsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumLocalFileSystemsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.FileSystemList != nil && o.FileSystemCount == 0 {
		o.FileSystemCount = uint32(len(o.FileSystemList))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumLocalFileSystemsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// fileSystemCount {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.FileSystemCount); err != nil {
			return err
		}
	}
	// fileSystemList {out} (1:{pointer=ref}*(2)*(1))(2:{alias=FILE_SYSTEM_INFO}[dim:0,size_is=fileSystemCount](struct))
	{
		if o.FileSystemList != nil || o.FileSystemCount > 0 {
			_ptr_fileSystemList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.FileSystemCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.FileSystemList {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.FileSystemList[i1] != nil {
						if err := o.FileSystemList[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dmrp.FileSystemInfo{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.FileSystemList); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&dmrp.FileSystemInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FileSystemList, _ptr_fileSystemList); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumLocalFileSystemsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// fileSystemCount {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.FileSystemCount); err != nil {
			return err
		}
	}
	// fileSystemList {out} (1:{pointer=ref}*(2)*(1))(2:{alias=FILE_SYSTEM_INFO}[dim:0,size_is=fileSystemCount](struct))
	{
		_ptr_fileSystemList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.FileSystemList", sizeInfo[0])
			}
			o.FileSystemList = make([]*dmrp.FileSystemInfo, sizeInfo[0])
			for i1 := range o.FileSystemList {
				i1 := i1
				if o.FileSystemList[i1] == nil {
					o.FileSystemList[i1] = &dmrp.FileSystemInfo{}
				}
				if err := o.FileSystemList[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_fileSystemList := func(ptr interface{}) { o.FileSystemList = *ptr.(*[]*dmrp.FileSystemInfo) }
		if err := w.ReadPointer(&o.FileSystemList, _s_fileSystemList, _ptr_fileSystemList); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumLocalFileSystemsRequest structure represents the EnumLocalFileSystems operation request
type EnumLocalFileSystemsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *EnumLocalFileSystemsRequest) xxx_ToOp(ctx context.Context) *xxx_EnumLocalFileSystemsOperation {
	if o == nil {
		return &xxx_EnumLocalFileSystemsOperation{}
	}
	return &xxx_EnumLocalFileSystemsOperation{
		This: o.This,
	}
}

func (o *EnumLocalFileSystemsRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumLocalFileSystemsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *EnumLocalFileSystemsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumLocalFileSystemsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumLocalFileSystemsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumLocalFileSystemsResponse structure represents the EnumLocalFileSystems operation response
type EnumLocalFileSystemsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That            *dcom.ORPCThat         `idl:"name:That" json:"that"`
	FileSystemCount uint32                 `idl:"name:fileSystemCount" json:"file_system_count"`
	FileSystemList  []*dmrp.FileSystemInfo `idl:"name:fileSystemList;size_is:(, fileSystemCount)" json:"file_system_list"`
	// Return: The EnumLocalFileSystems return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumLocalFileSystemsResponse) xxx_ToOp(ctx context.Context) *xxx_EnumLocalFileSystemsOperation {
	if o == nil {
		return &xxx_EnumLocalFileSystemsOperation{}
	}
	return &xxx_EnumLocalFileSystemsOperation{
		That:            o.That,
		FileSystemCount: o.FileSystemCount,
		FileSystemList:  o.FileSystemList,
		Return:          o.Return,
	}
}

func (o *EnumLocalFileSystemsResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumLocalFileSystemsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FileSystemCount = op.FileSystemCount
	o.FileSystemList = op.FileSystemList
	o.Return = op.Return
}
func (o *EnumLocalFileSystemsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumLocalFileSystemsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumLocalFileSystemsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetInstalledFileSystemsOperation structure represents the GetInstalledFileSystems operation
type xxx_GetInstalledFileSystemsOperation struct {
	This    *dcom.ORPCThis                  `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat                  `idl:"name:That" json:"that"`
	FSCount uint32                          `idl:"name:fsCount" json:"fs_count"`
	FSList  []*dmrp.InstalledFileSystemInfo `idl:"name:fsList;size_is:(, fsCount)" json:"fs_list"`
	Return  int32                           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetInstalledFileSystemsOperation) OpNum() int { return 25 }

func (o *xxx_GetInstalledFileSystemsOperation) OpName() string {
	return "/IVolumeClient/v0/GetInstalledFileSystems"
}

func (o *xxx_GetInstalledFileSystemsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInstalledFileSystemsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInstalledFileSystemsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInstalledFileSystemsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.FSList != nil && o.FSCount == 0 {
		o.FSCount = uint32(len(o.FSList))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInstalledFileSystemsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// fsCount {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.FSCount); err != nil {
			return err
		}
	}
	// fsList {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IFILE_SYSTEM_INFO}[dim:0,size_is=fsCount](struct))
	{
		if o.FSList != nil || o.FSCount > 0 {
			_ptr_fsList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.FSCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.FSList {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.FSList[i1] != nil {
						if err := o.FSList[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dmrp.InstalledFileSystemInfo{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.FSList); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&dmrp.InstalledFileSystemInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FSList, _ptr_fsList); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInstalledFileSystemsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// fsCount {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.FSCount); err != nil {
			return err
		}
	}
	// fsList {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IFILE_SYSTEM_INFO}[dim:0,size_is=fsCount](struct))
	{
		_ptr_fsList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.FSList", sizeInfo[0])
			}
			o.FSList = make([]*dmrp.InstalledFileSystemInfo, sizeInfo[0])
			for i1 := range o.FSList {
				i1 := i1
				if o.FSList[i1] == nil {
					o.FSList[i1] = &dmrp.InstalledFileSystemInfo{}
				}
				if err := o.FSList[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_fsList := func(ptr interface{}) { o.FSList = *ptr.(*[]*dmrp.InstalledFileSystemInfo) }
		if err := w.ReadPointer(&o.FSList, _s_fsList, _ptr_fsList); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetInstalledFileSystemsRequest structure represents the GetInstalledFileSystems operation request
type GetInstalledFileSystemsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetInstalledFileSystemsRequest) xxx_ToOp(ctx context.Context) *xxx_GetInstalledFileSystemsOperation {
	if o == nil {
		return &xxx_GetInstalledFileSystemsOperation{}
	}
	return &xxx_GetInstalledFileSystemsOperation{
		This: o.This,
	}
}

func (o *GetInstalledFileSystemsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetInstalledFileSystemsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetInstalledFileSystemsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetInstalledFileSystemsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetInstalledFileSystemsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetInstalledFileSystemsResponse structure represents the GetInstalledFileSystems operation response
type GetInstalledFileSystemsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat                  `idl:"name:That" json:"that"`
	FSCount uint32                          `idl:"name:fsCount" json:"fs_count"`
	FSList  []*dmrp.InstalledFileSystemInfo `idl:"name:fsList;size_is:(, fsCount)" json:"fs_list"`
	// Return: The GetInstalledFileSystems return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetInstalledFileSystemsResponse) xxx_ToOp(ctx context.Context) *xxx_GetInstalledFileSystemsOperation {
	if o == nil {
		return &xxx_GetInstalledFileSystemsOperation{}
	}
	return &xxx_GetInstalledFileSystemsOperation{
		That:    o.That,
		FSCount: o.FSCount,
		FSList:  o.FSList,
		Return:  o.Return,
	}
}

func (o *GetInstalledFileSystemsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetInstalledFileSystemsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FSCount = op.FSCount
	o.FSList = op.FSList
	o.Return = op.Return
}
func (o *GetInstalledFileSystemsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetInstalledFileSystemsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetInstalledFileSystemsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FormatOperation structure represents the Format operation
type xxx_FormatOperation struct {
	This                  *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That                  *dcom.ORPCThat       `idl:"name:That" json:"that"`
	StorageID             int64                `idl:"name:storageId" json:"storage_id"`
	FSSpec                *dmrp.FileSystemInfo `idl:"name:fsSpec" json:"fs_spec"`
	QuickFormat           bool                 `idl:"name:quickFormat" json:"quick_format"`
	Force                 bool                 `idl:"name:force" json:"force"`
	StorageLastKnownState int64                `idl:"name:storageLastKnownState" json:"storage_last_known_state"`
	TaskInfo              *dmrp.TaskInfo       `idl:"name:tinfo" json:"task_info"`
	Return                int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_FormatOperation) OpNum() int { return 26 }

func (o *xxx_FormatOperation) OpName() string { return "/IVolumeClient/v0/Format" }

func (o *xxx_FormatOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FormatOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// storageId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.StorageID); err != nil {
			return err
		}
	}
	// fsSpec {in} (1:{alias=FILE_SYSTEM_INFO}(struct))
	{
		if o.FSSpec != nil {
			if err := o.FSSpec.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.FileSystemInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// quickFormat {in} (1:(bool))
	{
		if err := w.WriteData(o.QuickFormat); err != nil {
			return err
		}
	}
	// force {in} (1:(bool))
	{
		if err := w.WriteData(o.Force); err != nil {
			return err
		}
	}
	// storageLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.StorageLastKnownState); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FormatOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// storageId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.StorageID); err != nil {
			return err
		}
	}
	// fsSpec {in} (1:{alias=FILE_SYSTEM_INFO}(struct))
	{
		if o.FSSpec == nil {
			o.FSSpec = &dmrp.FileSystemInfo{}
		}
		if err := o.FSSpec.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// quickFormat {in} (1:(bool))
	{
		if err := w.ReadData(&o.QuickFormat); err != nil {
			return err
		}
	}
	// force {in} (1:(bool))
	{
		if err := w.ReadData(&o.Force); err != nil {
			return err
		}
	}
	// storageLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.StorageLastKnownState); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FormatOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FormatOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FormatOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// FormatRequest structure represents the Format operation request
type FormatRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                  *dcom.ORPCThis       `idl:"name:This" json:"this"`
	StorageID             int64                `idl:"name:storageId" json:"storage_id"`
	FSSpec                *dmrp.FileSystemInfo `idl:"name:fsSpec" json:"fs_spec"`
	QuickFormat           bool                 `idl:"name:quickFormat" json:"quick_format"`
	Force                 bool                 `idl:"name:force" json:"force"`
	StorageLastKnownState int64                `idl:"name:storageLastKnownState" json:"storage_last_known_state"`
}

func (o *FormatRequest) xxx_ToOp(ctx context.Context) *xxx_FormatOperation {
	if o == nil {
		return &xxx_FormatOperation{}
	}
	return &xxx_FormatOperation{
		This:                  o.This,
		StorageID:             o.StorageID,
		FSSpec:                o.FSSpec,
		QuickFormat:           o.QuickFormat,
		Force:                 o.Force,
		StorageLastKnownState: o.StorageLastKnownState,
	}
}

func (o *FormatRequest) xxx_FromOp(ctx context.Context, op *xxx_FormatOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.StorageID = op.StorageID
	o.FSSpec = op.FSSpec
	o.QuickFormat = op.QuickFormat
	o.Force = op.Force
	o.StorageLastKnownState = op.StorageLastKnownState
}
func (o *FormatRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *FormatRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FormatOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FormatResponse structure represents the Format operation response
type FormatResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The Format return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *FormatResponse) xxx_ToOp(ctx context.Context) *xxx_FormatOperation {
	if o == nil {
		return &xxx_FormatOperation{}
	}
	return &xxx_FormatOperation{
		That:     o.That,
		TaskInfo: o.TaskInfo,
		Return:   o.Return,
	}
}

func (o *FormatResponse) xxx_FromOp(ctx context.Context, op *xxx_FormatOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *FormatResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *FormatResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FormatOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumVolumesOperation structure represents the EnumVolumes operation
type xxx_EnumVolumesOperation struct {
	This          *dcom.ORPCThis     `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat     `idl:"name:That" json:"that"`
	VolumeCount   uint32             `idl:"name:volumeCount" json:"volume_count"`
	LDMVolumeList []*dmrp.VolumeInfo `idl:"name:LdmVolumeList;size_is:(, volumeCount)" json:"ldm_volume_list"`
	Return        int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumVolumesOperation) OpNum() int { return 28 }

func (o *xxx_EnumVolumesOperation) OpName() string { return "/IVolumeClient/v0/EnumVolumes" }

func (o *xxx_EnumVolumesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumVolumesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// volumeCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.VolumeCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumVolumesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// volumeCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.VolumeCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumVolumesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.LDMVolumeList != nil && o.VolumeCount == 0 {
		o.VolumeCount = uint32(len(o.LDMVolumeList))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumVolumesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// volumeCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.VolumeCount); err != nil {
			return err
		}
	}
	// LdmVolumeList {out} (1:{pointer=ref}*(2)*(1))(2:{alias=VOLUME_INFO}[dim:0,size_is=volumeCount](struct))
	{
		if o.LDMVolumeList != nil || o.VolumeCount > 0 {
			_ptr_LdmVolumeList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.VolumeCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.LDMVolumeList {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.LDMVolumeList[i1] != nil {
						if err := o.LDMVolumeList[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dmrp.VolumeInfo{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.LDMVolumeList); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&dmrp.VolumeInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.LDMVolumeList, _ptr_LdmVolumeList); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumVolumesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// volumeCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.VolumeCount); err != nil {
			return err
		}
	}
	// LdmVolumeList {out} (1:{pointer=ref}*(2)*(1))(2:{alias=VOLUME_INFO}[dim:0,size_is=volumeCount](struct))
	{
		_ptr_LdmVolumeList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.LDMVolumeList", sizeInfo[0])
			}
			o.LDMVolumeList = make([]*dmrp.VolumeInfo, sizeInfo[0])
			for i1 := range o.LDMVolumeList {
				i1 := i1
				if o.LDMVolumeList[i1] == nil {
					o.LDMVolumeList[i1] = &dmrp.VolumeInfo{}
				}
				if err := o.LDMVolumeList[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_LdmVolumeList := func(ptr interface{}) { o.LDMVolumeList = *ptr.(*[]*dmrp.VolumeInfo) }
		if err := w.ReadPointer(&o.LDMVolumeList, _s_LdmVolumeList, _ptr_LdmVolumeList); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumVolumesRequest structure represents the EnumVolumes operation request
type EnumVolumesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	VolumeCount uint32         `idl:"name:volumeCount" json:"volume_count"`
}

func (o *EnumVolumesRequest) xxx_ToOp(ctx context.Context) *xxx_EnumVolumesOperation {
	if o == nil {
		return &xxx_EnumVolumesOperation{}
	}
	return &xxx_EnumVolumesOperation{
		This:        o.This,
		VolumeCount: o.VolumeCount,
	}
}

func (o *EnumVolumesRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumVolumesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VolumeCount = op.VolumeCount
}
func (o *EnumVolumesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumVolumesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumVolumesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumVolumesResponse structure represents the EnumVolumes operation response
type EnumVolumesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That          *dcom.ORPCThat     `idl:"name:That" json:"that"`
	VolumeCount   uint32             `idl:"name:volumeCount" json:"volume_count"`
	LDMVolumeList []*dmrp.VolumeInfo `idl:"name:LdmVolumeList;size_is:(, volumeCount)" json:"ldm_volume_list"`
	// Return: The EnumVolumes return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumVolumesResponse) xxx_ToOp(ctx context.Context) *xxx_EnumVolumesOperation {
	if o == nil {
		return &xxx_EnumVolumesOperation{}
	}
	return &xxx_EnumVolumesOperation{
		That:          o.That,
		VolumeCount:   o.VolumeCount,
		LDMVolumeList: o.LDMVolumeList,
		Return:        o.Return,
	}
}

func (o *EnumVolumesResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumVolumesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.VolumeCount = op.VolumeCount
	o.LDMVolumeList = op.LDMVolumeList
	o.Return = op.Return
}
func (o *EnumVolumesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumVolumesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumVolumesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumVolumeMembersOperation structure represents the EnumVolumeMembers operation
type xxx_EnumVolumeMembersOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	VolumeID    int64          `idl:"name:volumeId" json:"volume_id"`
	MemberCount uint32         `idl:"name:memberCount" json:"member_count"`
	MemberList  []int64        `idl:"name:memberList;size_is:(, memberCount)" json:"member_list"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumVolumeMembersOperation) OpNum() int { return 29 }

func (o *xxx_EnumVolumeMembersOperation) OpName() string {
	return "/IVolumeClient/v0/EnumVolumeMembers"
}

func (o *xxx_EnumVolumeMembersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumVolumeMembersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.VolumeID); err != nil {
			return err
		}
	}
	// memberCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.MemberCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumVolumeMembersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.VolumeID); err != nil {
			return err
		}
	}
	// memberCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.MemberCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumVolumeMembersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.MemberList != nil && o.MemberCount == 0 {
		o.MemberCount = uint32(len(o.MemberList))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumVolumeMembersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// memberCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.MemberCount); err != nil {
			return err
		}
	}
	// memberList {out} (1:{pointer=ref}*(2)*(1))(2:{alias=LdmObjectId, names=LONGLONG}[dim:0,size_is=memberCount](int64))
	{
		if o.MemberList != nil || o.MemberCount > 0 {
			_ptr_memberList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.MemberCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.MemberList {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.MemberList[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.MemberList); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(int64(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MemberList, _ptr_memberList); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumVolumeMembersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// memberCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.MemberCount); err != nil {
			return err
		}
	}
	// memberList {out} (1:{pointer=ref}*(2)*(1))(2:{alias=LdmObjectId, names=LONGLONG}[dim:0,size_is=memberCount](int64))
	{
		_ptr_memberList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.MemberList", sizeInfo[0])
			}
			o.MemberList = make([]int64, sizeInfo[0])
			for i1 := range o.MemberList {
				i1 := i1
				if err := w.ReadData(&o.MemberList[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_memberList := func(ptr interface{}) { o.MemberList = *ptr.(*[]int64) }
		if err := w.ReadPointer(&o.MemberList, _s_memberList, _ptr_memberList); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumVolumeMembersRequest structure represents the EnumVolumeMembers operation request
type EnumVolumeMembersRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	VolumeID    int64          `idl:"name:volumeId" json:"volume_id"`
	MemberCount uint32         `idl:"name:memberCount" json:"member_count"`
}

func (o *EnumVolumeMembersRequest) xxx_ToOp(ctx context.Context) *xxx_EnumVolumeMembersOperation {
	if o == nil {
		return &xxx_EnumVolumeMembersOperation{}
	}
	return &xxx_EnumVolumeMembersOperation{
		This:        o.This,
		VolumeID:    o.VolumeID,
		MemberCount: o.MemberCount,
	}
}

func (o *EnumVolumeMembersRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumVolumeMembersOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VolumeID = op.VolumeID
	o.MemberCount = op.MemberCount
}
func (o *EnumVolumeMembersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumVolumeMembersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumVolumeMembersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumVolumeMembersResponse structure represents the EnumVolumeMembers operation response
type EnumVolumeMembersResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	MemberCount uint32         `idl:"name:memberCount" json:"member_count"`
	MemberList  []int64        `idl:"name:memberList;size_is:(, memberCount)" json:"member_list"`
	// Return: The EnumVolumeMembers return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumVolumeMembersResponse) xxx_ToOp(ctx context.Context) *xxx_EnumVolumeMembersOperation {
	if o == nil {
		return &xxx_EnumVolumeMembersOperation{}
	}
	return &xxx_EnumVolumeMembersOperation{
		That:        o.That,
		MemberCount: o.MemberCount,
		MemberList:  o.MemberList,
		Return:      o.Return,
	}
}

func (o *EnumVolumeMembersResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumVolumeMembersOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MemberCount = op.MemberCount
	o.MemberList = op.MemberList
	o.Return = op.Return
}
func (o *EnumVolumeMembersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumVolumeMembersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumVolumeMembersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateVolumeOperation structure represents the CreateVolume operation
type xxx_CreateVolumeOperation struct {
	This       *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat   `idl:"name:That" json:"that"`
	VolumeSpec *dmrp.VolumeSpec `idl:"name:volumeSpec" json:"volume_spec"`
	DiskCount  uint32           `idl:"name:diskCount" json:"disk_count"`
	DiskList   []*dmrp.DiskSpec `idl:"name:diskList;size_is:(diskCount)" json:"disk_list"`
	TaskInfo   *dmrp.TaskInfo   `idl:"name:tinfo" json:"task_info"`
	Return     int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateVolumeOperation) OpNum() int { return 30 }

func (o *xxx_CreateVolumeOperation) OpName() string { return "/IVolumeClient/v0/CreateVolume" }

func (o *xxx_CreateVolumeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.DiskList != nil && o.DiskCount == 0 {
		o.DiskCount = uint32(len(o.DiskList))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateVolumeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// volumeSpec {in} (1:{alias=VOLUME_SPEC}(struct))
	{
		if o.VolumeSpec != nil {
			if err := o.VolumeSpec.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.VolumeSpec{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// diskCount {in} (1:(uint32))
	{
		if err := w.WriteData(o.DiskCount); err != nil {
			return err
		}
	}
	// diskList {in} (1:{pointer=ref}*(1))(2:{alias=DISK_SPEC}[dim:0,size_is=diskCount](struct))
	{
		dimSize1 := uint64(o.DiskCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.DiskList {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.DiskList[i1] != nil {
				if err := o.DiskList[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dmrp.DiskSpec{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.DiskList); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&dmrp.DiskSpec{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CreateVolumeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// volumeSpec {in} (1:{alias=VOLUME_SPEC}(struct))
	{
		if o.VolumeSpec == nil {
			o.VolumeSpec = &dmrp.VolumeSpec{}
		}
		if err := o.VolumeSpec.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// diskCount {in} (1:(uint32))
	{
		if err := w.ReadData(&o.DiskCount); err != nil {
			return err
		}
	}
	// diskList {in} (1:{pointer=ref}*(1))(2:{alias=DISK_SPEC}[dim:0,size_is=diskCount](struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.DiskList", sizeInfo[0])
		}
		o.DiskList = make([]*dmrp.DiskSpec, sizeInfo[0])
		for i1 := range o.DiskList {
			i1 := i1
			if o.DiskList[i1] == nil {
				o.DiskList[i1] = &dmrp.DiskSpec{}
			}
			if err := o.DiskList[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CreateVolumeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateVolumeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateVolumeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateVolumeRequest structure represents the CreateVolume operation request
type CreateVolumeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis   `idl:"name:This" json:"this"`
	VolumeSpec *dmrp.VolumeSpec `idl:"name:volumeSpec" json:"volume_spec"`
	DiskCount  uint32           `idl:"name:diskCount" json:"disk_count"`
	DiskList   []*dmrp.DiskSpec `idl:"name:diskList;size_is:(diskCount)" json:"disk_list"`
}

func (o *CreateVolumeRequest) xxx_ToOp(ctx context.Context) *xxx_CreateVolumeOperation {
	if o == nil {
		return &xxx_CreateVolumeOperation{}
	}
	return &xxx_CreateVolumeOperation{
		This:       o.This,
		VolumeSpec: o.VolumeSpec,
		DiskCount:  o.DiskCount,
		DiskList:   o.DiskList,
	}
}

func (o *CreateVolumeRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateVolumeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VolumeSpec = op.VolumeSpec
	o.DiskCount = op.DiskCount
	o.DiskList = op.DiskList
}
func (o *CreateVolumeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateVolumeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateVolumeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateVolumeResponse structure represents the CreateVolume operation response
type CreateVolumeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The CreateVolume return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateVolumeResponse) xxx_ToOp(ctx context.Context) *xxx_CreateVolumeOperation {
	if o == nil {
		return &xxx_CreateVolumeOperation{}
	}
	return &xxx_CreateVolumeOperation{
		That:     o.That,
		TaskInfo: o.TaskInfo,
		Return:   o.Return,
	}
}

func (o *CreateVolumeResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateVolumeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *CreateVolumeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateVolumeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateVolumeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateVolumeAssignAndFormatOperation structure represents the CreateVolumeAssignAndFormat operation
type xxx_CreateVolumeAssignAndFormatOperation struct {
	This                 *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat       `idl:"name:That" json:"that"`
	VolumeSpec           *dmrp.VolumeSpec     `idl:"name:volumeSpec" json:"volume_spec"`
	DiskCount            uint32               `idl:"name:diskCount" json:"disk_count"`
	DiskList             []*dmrp.DiskSpec     `idl:"name:diskList;size_is:(diskCount)" json:"disk_list"`
	Letter               uint16               `idl:"name:letter" json:"letter"`
	LetterLastKnownState int64                `idl:"name:letterLastKnownState" json:"letter_last_known_state"`
	FSSpec               *dmrp.FileSystemInfo `idl:"name:fsSpec" json:"fs_spec"`
	QuickFormat          bool                 `idl:"name:quickFormat" json:"quick_format"`
	TaskInfo             *dmrp.TaskInfo       `idl:"name:tinfo" json:"task_info"`
	Return               int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateVolumeAssignAndFormatOperation) OpNum() int { return 31 }

func (o *xxx_CreateVolumeAssignAndFormatOperation) OpName() string {
	return "/IVolumeClient/v0/CreateVolumeAssignAndFormat"
}

func (o *xxx_CreateVolumeAssignAndFormatOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.DiskList != nil && o.DiskCount == 0 {
		o.DiskCount = uint32(len(o.DiskList))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateVolumeAssignAndFormatOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// volumeSpec {in} (1:{alias=VOLUME_SPEC}(struct))
	{
		if o.VolumeSpec != nil {
			if err := o.VolumeSpec.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.VolumeSpec{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// diskCount {in} (1:(uint32))
	{
		if err := w.WriteData(o.DiskCount); err != nil {
			return err
		}
	}
	// diskList {in} (1:{pointer=ref}*(1))(2:{alias=DISK_SPEC}[dim:0,size_is=diskCount](struct))
	{
		dimSize1 := uint64(o.DiskCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.DiskList {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.DiskList[i1] != nil {
				if err := o.DiskList[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dmrp.DiskSpec{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.DiskList); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&dmrp.DiskSpec{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// letter {in} (1:(wchar))
	{
		if err := w.WriteData(o.Letter); err != nil {
			return err
		}
	}
	// letterLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.LetterLastKnownState); err != nil {
			return err
		}
	}
	// fsSpec {in} (1:{alias=FILE_SYSTEM_INFO}(struct))
	{
		if o.FSSpec != nil {
			if err := o.FSSpec.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.FileSystemInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// quickFormat {in} (1:(bool))
	{
		if err := w.WriteData(o.QuickFormat); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateVolumeAssignAndFormatOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// volumeSpec {in} (1:{alias=VOLUME_SPEC}(struct))
	{
		if o.VolumeSpec == nil {
			o.VolumeSpec = &dmrp.VolumeSpec{}
		}
		if err := o.VolumeSpec.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// diskCount {in} (1:(uint32))
	{
		if err := w.ReadData(&o.DiskCount); err != nil {
			return err
		}
	}
	// diskList {in} (1:{pointer=ref}*(1))(2:{alias=DISK_SPEC}[dim:0,size_is=diskCount](struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.DiskList", sizeInfo[0])
		}
		o.DiskList = make([]*dmrp.DiskSpec, sizeInfo[0])
		for i1 := range o.DiskList {
			i1 := i1
			if o.DiskList[i1] == nil {
				o.DiskList[i1] = &dmrp.DiskSpec{}
			}
			if err := o.DiskList[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// letter {in} (1:(wchar))
	{
		if err := w.ReadData(&o.Letter); err != nil {
			return err
		}
	}
	// letterLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.LetterLastKnownState); err != nil {
			return err
		}
	}
	// fsSpec {in} (1:{alias=FILE_SYSTEM_INFO}(struct))
	{
		if o.FSSpec == nil {
			o.FSSpec = &dmrp.FileSystemInfo{}
		}
		if err := o.FSSpec.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// quickFormat {in} (1:(bool))
	{
		if err := w.ReadData(&o.QuickFormat); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateVolumeAssignAndFormatOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateVolumeAssignAndFormatOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateVolumeAssignAndFormatOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateVolumeAssignAndFormatRequest structure represents the CreateVolumeAssignAndFormat operation request
type CreateVolumeAssignAndFormatRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis       `idl:"name:This" json:"this"`
	VolumeSpec           *dmrp.VolumeSpec     `idl:"name:volumeSpec" json:"volume_spec"`
	DiskCount            uint32               `idl:"name:diskCount" json:"disk_count"`
	DiskList             []*dmrp.DiskSpec     `idl:"name:diskList;size_is:(diskCount)" json:"disk_list"`
	Letter               uint16               `idl:"name:letter" json:"letter"`
	LetterLastKnownState int64                `idl:"name:letterLastKnownState" json:"letter_last_known_state"`
	FSSpec               *dmrp.FileSystemInfo `idl:"name:fsSpec" json:"fs_spec"`
	QuickFormat          bool                 `idl:"name:quickFormat" json:"quick_format"`
}

func (o *CreateVolumeAssignAndFormatRequest) xxx_ToOp(ctx context.Context) *xxx_CreateVolumeAssignAndFormatOperation {
	if o == nil {
		return &xxx_CreateVolumeAssignAndFormatOperation{}
	}
	return &xxx_CreateVolumeAssignAndFormatOperation{
		This:                 o.This,
		VolumeSpec:           o.VolumeSpec,
		DiskCount:            o.DiskCount,
		DiskList:             o.DiskList,
		Letter:               o.Letter,
		LetterLastKnownState: o.LetterLastKnownState,
		FSSpec:               o.FSSpec,
		QuickFormat:          o.QuickFormat,
	}
}

func (o *CreateVolumeAssignAndFormatRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateVolumeAssignAndFormatOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VolumeSpec = op.VolumeSpec
	o.DiskCount = op.DiskCount
	o.DiskList = op.DiskList
	o.Letter = op.Letter
	o.LetterLastKnownState = op.LetterLastKnownState
	o.FSSpec = op.FSSpec
	o.QuickFormat = op.QuickFormat
}
func (o *CreateVolumeAssignAndFormatRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateVolumeAssignAndFormatRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateVolumeAssignAndFormatOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateVolumeAssignAndFormatResponse structure represents the CreateVolumeAssignAndFormat operation response
type CreateVolumeAssignAndFormatResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The CreateVolumeAssignAndFormat return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateVolumeAssignAndFormatResponse) xxx_ToOp(ctx context.Context) *xxx_CreateVolumeAssignAndFormatOperation {
	if o == nil {
		return &xxx_CreateVolumeAssignAndFormatOperation{}
	}
	return &xxx_CreateVolumeAssignAndFormatOperation{
		That:     o.That,
		TaskInfo: o.TaskInfo,
		Return:   o.Return,
	}
}

func (o *CreateVolumeAssignAndFormatResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateVolumeAssignAndFormatOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *CreateVolumeAssignAndFormatResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateVolumeAssignAndFormatResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateVolumeAssignAndFormatOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateVolumeAssignAndFormatExOperation structure represents the CreateVolumeAssignAndFormatEx operation
type xxx_CreateVolumeAssignAndFormatExOperation struct {
	This                 *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat       `idl:"name:That" json:"that"`
	VolumeSpec           *dmrp.VolumeSpec     `idl:"name:volumeSpec" json:"volume_spec"`
	DiskCount            uint32               `idl:"name:diskCount" json:"disk_count"`
	DiskList             []*dmrp.DiskSpec     `idl:"name:diskList;size_is:(diskCount)" json:"disk_list"`
	Letter               uint16               `idl:"name:letter" json:"letter"`
	LetterLastKnownState int64                `idl:"name:letterLastKnownState" json:"letter_last_known_state"`
	AccessPathLength     int32                `idl:"name:cchAccessPath" json:"access_path_length"`
	AccessPath           string               `idl:"name:AccessPath;size_is:(cchAccessPath)" json:"access_path"`
	FSSpec               *dmrp.FileSystemInfo `idl:"name:fsSpec" json:"fs_spec"`
	QuickFormat          bool                 `idl:"name:quickFormat" json:"quick_format"`
	Flags                uint32               `idl:"name:dwFlags" json:"flags"`
	TaskInfo             *dmrp.TaskInfo       `idl:"name:tinfo" json:"task_info"`
	Return               int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateVolumeAssignAndFormatExOperation) OpNum() int { return 32 }

func (o *xxx_CreateVolumeAssignAndFormatExOperation) OpName() string {
	return "/IVolumeClient/v0/CreateVolumeAssignAndFormatEx"
}

func (o *xxx_CreateVolumeAssignAndFormatExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.DiskList != nil && o.DiskCount == 0 {
		o.DiskCount = uint32(len(o.DiskList))
	}
	if o.AccessPath != "" && o.AccessPathLength == 0 {
		o.AccessPathLength = int32(len(o.AccessPath))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateVolumeAssignAndFormatExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// volumeSpec {in} (1:{alias=VOLUME_SPEC}(struct))
	{
		if o.VolumeSpec != nil {
			if err := o.VolumeSpec.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.VolumeSpec{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// diskCount {in} (1:(uint32))
	{
		if err := w.WriteData(o.DiskCount); err != nil {
			return err
		}
	}
	// diskList {in} (1:{pointer=ref}*(1))(2:{alias=DISK_SPEC}[dim:0,size_is=diskCount](struct))
	{
		dimSize1 := uint64(o.DiskCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.DiskList {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.DiskList[i1] != nil {
				if err := o.DiskList[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dmrp.DiskSpec{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.DiskList); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&dmrp.DiskSpec{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// letter {in} (1:(wchar))
	{
		if err := w.WriteData(o.Letter); err != nil {
			return err
		}
	}
	// letterLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.LetterLastKnownState); err != nil {
			return err
		}
	}
	// cchAccessPath {in} (1:(int32))
	{
		if err := w.WriteData(o.AccessPathLength); err != nil {
			return err
		}
	}
	// AccessPath {in} (1:{pointer=ref}*(1)[dim:0,size_is=cchAccessPath,string](wchar))
	{
		dimSize1 := uint64(o.AccessPathLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		_AccessPath_buf := utf16.Encode([]rune(o.AccessPath))
		if uint64(len(_AccessPath_buf)) > sizeInfo[0] {
			_AccessPath_buf = _AccessPath_buf[:sizeInfo[0]]
		}
		for i1 := range _AccessPath_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_AccessPath_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_AccessPath_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	// fsSpec {in} (1:{alias=FILE_SYSTEM_INFO}(struct))
	{
		if o.FSSpec != nil {
			if err := o.FSSpec.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.FileSystemInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// quickFormat {in} (1:(bool))
	{
		if err := w.WriteData(o.QuickFormat); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateVolumeAssignAndFormatExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// volumeSpec {in} (1:{alias=VOLUME_SPEC}(struct))
	{
		if o.VolumeSpec == nil {
			o.VolumeSpec = &dmrp.VolumeSpec{}
		}
		if err := o.VolumeSpec.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// diskCount {in} (1:(uint32))
	{
		if err := w.ReadData(&o.DiskCount); err != nil {
			return err
		}
	}
	// diskList {in} (1:{pointer=ref}*(1))(2:{alias=DISK_SPEC}[dim:0,size_is=diskCount](struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.DiskList", sizeInfo[0])
		}
		o.DiskList = make([]*dmrp.DiskSpec, sizeInfo[0])
		for i1 := range o.DiskList {
			i1 := i1
			if o.DiskList[i1] == nil {
				o.DiskList[i1] = &dmrp.DiskSpec{}
			}
			if err := o.DiskList[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// letter {in} (1:(wchar))
	{
		if err := w.ReadData(&o.Letter); err != nil {
			return err
		}
	}
	// letterLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.LetterLastKnownState); err != nil {
			return err
		}
	}
	// cchAccessPath {in} (1:(int32))
	{
		if err := w.ReadData(&o.AccessPathLength); err != nil {
			return err
		}
	}
	// AccessPath {in} (1:{pointer=ref}*(1)[dim:0,size_is=cchAccessPath,string](wchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _AccessPath_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _AccessPath_buf", sizeInfo[0])
		}
		_AccessPath_buf = make([]uint16, sizeInfo[0])
		for i1 := range _AccessPath_buf {
			i1 := i1
			if err := w.ReadData(&_AccessPath_buf[i1]); err != nil {
				return err
			}
		}
		o.AccessPath = strings.TrimRight(string(utf16.Decode(_AccessPath_buf)), ndr.ZeroString)
	}
	// fsSpec {in} (1:{alias=FILE_SYSTEM_INFO}(struct))
	{
		if o.FSSpec == nil {
			o.FSSpec = &dmrp.FileSystemInfo{}
		}
		if err := o.FSSpec.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// quickFormat {in} (1:(bool))
	{
		if err := w.ReadData(&o.QuickFormat); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateVolumeAssignAndFormatExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateVolumeAssignAndFormatExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateVolumeAssignAndFormatExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateVolumeAssignAndFormatExRequest structure represents the CreateVolumeAssignAndFormatEx operation request
type CreateVolumeAssignAndFormatExRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis       `idl:"name:This" json:"this"`
	VolumeSpec           *dmrp.VolumeSpec     `idl:"name:volumeSpec" json:"volume_spec"`
	DiskCount            uint32               `idl:"name:diskCount" json:"disk_count"`
	DiskList             []*dmrp.DiskSpec     `idl:"name:diskList;size_is:(diskCount)" json:"disk_list"`
	Letter               uint16               `idl:"name:letter" json:"letter"`
	LetterLastKnownState int64                `idl:"name:letterLastKnownState" json:"letter_last_known_state"`
	AccessPathLength     int32                `idl:"name:cchAccessPath" json:"access_path_length"`
	AccessPath           string               `idl:"name:AccessPath;size_is:(cchAccessPath)" json:"access_path"`
	FSSpec               *dmrp.FileSystemInfo `idl:"name:fsSpec" json:"fs_spec"`
	QuickFormat          bool                 `idl:"name:quickFormat" json:"quick_format"`
	Flags                uint32               `idl:"name:dwFlags" json:"flags"`
}

func (o *CreateVolumeAssignAndFormatExRequest) xxx_ToOp(ctx context.Context) *xxx_CreateVolumeAssignAndFormatExOperation {
	if o == nil {
		return &xxx_CreateVolumeAssignAndFormatExOperation{}
	}
	return &xxx_CreateVolumeAssignAndFormatExOperation{
		This:                 o.This,
		VolumeSpec:           o.VolumeSpec,
		DiskCount:            o.DiskCount,
		DiskList:             o.DiskList,
		Letter:               o.Letter,
		LetterLastKnownState: o.LetterLastKnownState,
		AccessPathLength:     o.AccessPathLength,
		AccessPath:           o.AccessPath,
		FSSpec:               o.FSSpec,
		QuickFormat:          o.QuickFormat,
		Flags:                o.Flags,
	}
}

func (o *CreateVolumeAssignAndFormatExRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateVolumeAssignAndFormatExOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VolumeSpec = op.VolumeSpec
	o.DiskCount = op.DiskCount
	o.DiskList = op.DiskList
	o.Letter = op.Letter
	o.LetterLastKnownState = op.LetterLastKnownState
	o.AccessPathLength = op.AccessPathLength
	o.AccessPath = op.AccessPath
	o.FSSpec = op.FSSpec
	o.QuickFormat = op.QuickFormat
	o.Flags = op.Flags
}
func (o *CreateVolumeAssignAndFormatExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateVolumeAssignAndFormatExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateVolumeAssignAndFormatExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateVolumeAssignAndFormatExResponse structure represents the CreateVolumeAssignAndFormatEx operation response
type CreateVolumeAssignAndFormatExResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The CreateVolumeAssignAndFormatEx return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateVolumeAssignAndFormatExResponse) xxx_ToOp(ctx context.Context) *xxx_CreateVolumeAssignAndFormatExOperation {
	if o == nil {
		return &xxx_CreateVolumeAssignAndFormatExOperation{}
	}
	return &xxx_CreateVolumeAssignAndFormatExOperation{
		That:     o.That,
		TaskInfo: o.TaskInfo,
		Return:   o.Return,
	}
}

func (o *CreateVolumeAssignAndFormatExResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateVolumeAssignAndFormatExOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *CreateVolumeAssignAndFormatExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateVolumeAssignAndFormatExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateVolumeAssignAndFormatExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetVolumeMountNameOperation structure represents the GetVolumeMountName operation
type xxx_GetVolumeMountNameOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	VolumeID        int64          `idl:"name:volumeId" json:"volume_id"`
	MountNameLength uint32         `idl:"name:cchMountName" json:"mount_name_length"`
	MountName       string         `idl:"name:mountName;size_is:(, cchMountName)" json:"mount_name"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetVolumeMountNameOperation) OpNum() int { return 33 }

func (o *xxx_GetVolumeMountNameOperation) OpName() string {
	return "/IVolumeClient/v0/GetVolumeMountName"
}

func (o *xxx_GetVolumeMountNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVolumeMountNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.VolumeID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVolumeMountNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.VolumeID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVolumeMountNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.MountName != "" && o.MountNameLength == 0 {
		o.MountNameLength = uint32(len(o.MountName))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVolumeMountNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// cchMountName {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.MountNameLength); err != nil {
			return err
		}
	}
	// mountName {out} (1:{pointer=ref}*(2)*(1))(2:{alias=WCHAR}[dim:0,size_is=cchMountName,string](wchar))
	{
		if o.MountName != "" || o.MountNameLength > 0 {
			_ptr_mountName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.MountNameLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				_MountName_buf := utf16.Encode([]rune(o.MountName))
				if uint64(len(_MountName_buf)) > sizeInfo[0] {
					_MountName_buf = _MountName_buf[:sizeInfo[0]]
				}
				for i1 := range _MountName_buf {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(_MountName_buf[i1]); err != nil {
						return err
					}
				}
				for i1 := len(_MountName_buf); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint16(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MountName, _ptr_mountName); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVolumeMountNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cchMountName {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.MountNameLength); err != nil {
			return err
		}
	}
	// mountName {out} (1:{pointer=ref}*(2)*(1))(2:{alias=WCHAR}[dim:0,size_is=cchMountName,string](wchar))
	{
		_ptr_mountName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			var _MountName_buf []uint16
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array _MountName_buf", sizeInfo[0])
			}
			_MountName_buf = make([]uint16, sizeInfo[0])
			for i1 := range _MountName_buf {
				i1 := i1
				if err := w.ReadData(&_MountName_buf[i1]); err != nil {
					return err
				}
			}
			o.MountName = strings.TrimRight(string(utf16.Decode(_MountName_buf)), ndr.ZeroString)
			return nil
		})
		_s_mountName := func(ptr interface{}) { o.MountName = *ptr.(*string) }
		if err := w.ReadPointer(&o.MountName, _s_mountName, _ptr_mountName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetVolumeMountNameRequest structure represents the GetVolumeMountName operation request
type GetVolumeMountNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	VolumeID int64          `idl:"name:volumeId" json:"volume_id"`
}

func (o *GetVolumeMountNameRequest) xxx_ToOp(ctx context.Context) *xxx_GetVolumeMountNameOperation {
	if o == nil {
		return &xxx_GetVolumeMountNameOperation{}
	}
	return &xxx_GetVolumeMountNameOperation{
		This:     o.This,
		VolumeID: o.VolumeID,
	}
}

func (o *GetVolumeMountNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetVolumeMountNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VolumeID = op.VolumeID
}
func (o *GetVolumeMountNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetVolumeMountNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetVolumeMountNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetVolumeMountNameResponse structure represents the GetVolumeMountName operation response
type GetVolumeMountNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	MountNameLength uint32         `idl:"name:cchMountName" json:"mount_name_length"`
	MountName       string         `idl:"name:mountName;size_is:(, cchMountName)" json:"mount_name"`
	// Return: The GetVolumeMountName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetVolumeMountNameResponse) xxx_ToOp(ctx context.Context) *xxx_GetVolumeMountNameOperation {
	if o == nil {
		return &xxx_GetVolumeMountNameOperation{}
	}
	return &xxx_GetVolumeMountNameOperation{
		That:            o.That,
		MountNameLength: o.MountNameLength,
		MountName:       o.MountName,
		Return:          o.Return,
	}
}

func (o *GetVolumeMountNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetVolumeMountNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MountNameLength = op.MountNameLength
	o.MountName = op.MountName
	o.Return = op.Return
}
func (o *GetVolumeMountNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetVolumeMountNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetVolumeMountNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GrowVolumeOperation structure represents the GrowVolume operation
type xxx_GrowVolumeOperation struct {
	This       *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat   `idl:"name:That" json:"that"`
	VolumeID   int64            `idl:"name:volumeId" json:"volume_id"`
	VolumeSpec *dmrp.VolumeSpec `idl:"name:volumeSpec" json:"volume_spec"`
	DiskCount  uint32           `idl:"name:diskCount" json:"disk_count"`
	DiskList   []*dmrp.DiskSpec `idl:"name:diskList;size_is:(diskCount)" json:"disk_list"`
	Force      bool             `idl:"name:force" json:"force"`
	TaskInfo   *dmrp.TaskInfo   `idl:"name:tinfo" json:"task_info"`
	Return     int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GrowVolumeOperation) OpNum() int { return 34 }

func (o *xxx_GrowVolumeOperation) OpName() string { return "/IVolumeClient/v0/GrowVolume" }

func (o *xxx_GrowVolumeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.DiskList != nil && o.DiskCount == 0 {
		o.DiskCount = uint32(len(o.DiskList))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GrowVolumeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.VolumeID); err != nil {
			return err
		}
	}
	// volumeSpec {in} (1:{alias=VOLUME_SPEC}(struct))
	{
		if o.VolumeSpec != nil {
			if err := o.VolumeSpec.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.VolumeSpec{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// diskCount {in} (1:(uint32))
	{
		if err := w.WriteData(o.DiskCount); err != nil {
			return err
		}
	}
	// diskList {in} (1:{pointer=ref}*(1))(2:{alias=DISK_SPEC}[dim:0,size_is=diskCount](struct))
	{
		dimSize1 := uint64(o.DiskCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.DiskList {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.DiskList[i1] != nil {
				if err := o.DiskList[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dmrp.DiskSpec{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.DiskList); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&dmrp.DiskSpec{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// force {in} (1:(bool))
	{
		if err := w.WriteData(o.Force); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GrowVolumeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.VolumeID); err != nil {
			return err
		}
	}
	// volumeSpec {in} (1:{alias=VOLUME_SPEC}(struct))
	{
		if o.VolumeSpec == nil {
			o.VolumeSpec = &dmrp.VolumeSpec{}
		}
		if err := o.VolumeSpec.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// diskCount {in} (1:(uint32))
	{
		if err := w.ReadData(&o.DiskCount); err != nil {
			return err
		}
	}
	// diskList {in} (1:{pointer=ref}*(1))(2:{alias=DISK_SPEC}[dim:0,size_is=diskCount](struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.DiskList", sizeInfo[0])
		}
		o.DiskList = make([]*dmrp.DiskSpec, sizeInfo[0])
		for i1 := range o.DiskList {
			i1 := i1
			if o.DiskList[i1] == nil {
				o.DiskList[i1] = &dmrp.DiskSpec{}
			}
			if err := o.DiskList[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// force {in} (1:(bool))
	{
		if err := w.ReadData(&o.Force); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GrowVolumeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GrowVolumeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GrowVolumeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GrowVolumeRequest structure represents the GrowVolume operation request
type GrowVolumeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis   `idl:"name:This" json:"this"`
	VolumeID   int64            `idl:"name:volumeId" json:"volume_id"`
	VolumeSpec *dmrp.VolumeSpec `idl:"name:volumeSpec" json:"volume_spec"`
	DiskCount  uint32           `idl:"name:diskCount" json:"disk_count"`
	DiskList   []*dmrp.DiskSpec `idl:"name:diskList;size_is:(diskCount)" json:"disk_list"`
	Force      bool             `idl:"name:force" json:"force"`
}

func (o *GrowVolumeRequest) xxx_ToOp(ctx context.Context) *xxx_GrowVolumeOperation {
	if o == nil {
		return &xxx_GrowVolumeOperation{}
	}
	return &xxx_GrowVolumeOperation{
		This:       o.This,
		VolumeID:   o.VolumeID,
		VolumeSpec: o.VolumeSpec,
		DiskCount:  o.DiskCount,
		DiskList:   o.DiskList,
		Force:      o.Force,
	}
}

func (o *GrowVolumeRequest) xxx_FromOp(ctx context.Context, op *xxx_GrowVolumeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VolumeID = op.VolumeID
	o.VolumeSpec = op.VolumeSpec
	o.DiskCount = op.DiskCount
	o.DiskList = op.DiskList
	o.Force = op.Force
}
func (o *GrowVolumeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GrowVolumeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GrowVolumeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GrowVolumeResponse structure represents the GrowVolume operation response
type GrowVolumeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The GrowVolume return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GrowVolumeResponse) xxx_ToOp(ctx context.Context) *xxx_GrowVolumeOperation {
	if o == nil {
		return &xxx_GrowVolumeOperation{}
	}
	return &xxx_GrowVolumeOperation{
		That:     o.That,
		TaskInfo: o.TaskInfo,
		Return:   o.Return,
	}
}

func (o *GrowVolumeResponse) xxx_FromOp(ctx context.Context, op *xxx_GrowVolumeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *GrowVolumeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GrowVolumeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GrowVolumeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteVolumeOperation structure represents the DeleteVolume operation
type xxx_DeleteVolumeOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	VolumeID             int64          `idl:"name:volumeId" json:"volume_id"`
	Force                bool           `idl:"name:force" json:"force"`
	VolumeLastKnownState int64          `idl:"name:volumeLastKnownState" json:"volume_last_known_state"`
	TaskInfo             *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteVolumeOperation) OpNum() int { return 35 }

func (o *xxx_DeleteVolumeOperation) OpName() string { return "/IVolumeClient/v0/DeleteVolume" }

func (o *xxx_DeleteVolumeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteVolumeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.VolumeID); err != nil {
			return err
		}
	}
	// force {in} (1:(bool))
	{
		if err := w.WriteData(o.Force); err != nil {
			return err
		}
	}
	// volumeLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.VolumeLastKnownState); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteVolumeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.VolumeID); err != nil {
			return err
		}
	}
	// force {in} (1:(bool))
	{
		if err := w.ReadData(&o.Force); err != nil {
			return err
		}
	}
	// volumeLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.VolumeLastKnownState); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteVolumeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteVolumeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteVolumeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DeleteVolumeRequest structure represents the DeleteVolume operation request
type DeleteVolumeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	VolumeID             int64          `idl:"name:volumeId" json:"volume_id"`
	Force                bool           `idl:"name:force" json:"force"`
	VolumeLastKnownState int64          `idl:"name:volumeLastKnownState" json:"volume_last_known_state"`
}

func (o *DeleteVolumeRequest) xxx_ToOp(ctx context.Context) *xxx_DeleteVolumeOperation {
	if o == nil {
		return &xxx_DeleteVolumeOperation{}
	}
	return &xxx_DeleteVolumeOperation{
		This:                 o.This,
		VolumeID:             o.VolumeID,
		Force:                o.Force,
		VolumeLastKnownState: o.VolumeLastKnownState,
	}
}

func (o *DeleteVolumeRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteVolumeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VolumeID = op.VolumeID
	o.Force = op.Force
	o.VolumeLastKnownState = op.VolumeLastKnownState
}
func (o *DeleteVolumeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteVolumeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteVolumeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteVolumeResponse structure represents the DeleteVolume operation response
type DeleteVolumeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The DeleteVolume return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteVolumeResponse) xxx_ToOp(ctx context.Context) *xxx_DeleteVolumeOperation {
	if o == nil {
		return &xxx_DeleteVolumeOperation{}
	}
	return &xxx_DeleteVolumeOperation{
		That:     o.That,
		TaskInfo: o.TaskInfo,
		Return:   o.Return,
	}
}

func (o *DeleteVolumeResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteVolumeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *DeleteVolumeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteVolumeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteVolumeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddMirrorOperation structure represents the AddMirror operation
type xxx_AddMirrorOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	VolumeID             int64          `idl:"name:volumeId" json:"volume_id"`
	VolumeLastKnownState int64          `idl:"name:volumeLastKnownState" json:"volume_last_known_state"`
	DiskSpec             *dmrp.DiskSpec `idl:"name:diskSpec" json:"disk_spec"`
	DiskNumber           int32          `idl:"name:diskNumber" json:"disk_number"`
	PartitionNumber      int32          `idl:"name:partitionNumber" json:"partition_number"`
	TaskInfo             *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_AddMirrorOperation) OpNum() int { return 36 }

func (o *xxx_AddMirrorOperation) OpName() string { return "/IVolumeClient/v0/AddMirror" }

func (o *xxx_AddMirrorOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddMirrorOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.VolumeID); err != nil {
			return err
		}
	}
	// volumeLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.VolumeLastKnownState); err != nil {
			return err
		}
	}
	// diskSpec {in} (1:{alias=DISK_SPEC}(struct))
	{
		if o.DiskSpec != nil {
			if err := o.DiskSpec.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.DiskSpec{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// diskNumber {in, out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.DiskNumber); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddMirrorOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.VolumeID); err != nil {
			return err
		}
	}
	// volumeLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.VolumeLastKnownState); err != nil {
			return err
		}
	}
	// diskSpec {in} (1:{alias=DISK_SPEC}(struct))
	{
		if o.DiskSpec == nil {
			o.DiskSpec = &dmrp.DiskSpec{}
		}
		if err := o.DiskSpec.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// diskNumber {in, out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.DiskNumber); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddMirrorOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddMirrorOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// diskNumber {in, out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.DiskNumber); err != nil {
			return err
		}
	}
	// partitionNumber {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.PartitionNumber); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddMirrorOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// diskNumber {in, out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.DiskNumber); err != nil {
			return err
		}
	}
	// partitionNumber {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.PartitionNumber); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AddMirrorRequest structure represents the AddMirror operation request
type AddMirrorRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	VolumeID             int64          `idl:"name:volumeId" json:"volume_id"`
	VolumeLastKnownState int64          `idl:"name:volumeLastKnownState" json:"volume_last_known_state"`
	DiskSpec             *dmrp.DiskSpec `idl:"name:diskSpec" json:"disk_spec"`
	DiskNumber           int32          `idl:"name:diskNumber" json:"disk_number"`
}

func (o *AddMirrorRequest) xxx_ToOp(ctx context.Context) *xxx_AddMirrorOperation {
	if o == nil {
		return &xxx_AddMirrorOperation{}
	}
	return &xxx_AddMirrorOperation{
		This:                 o.This,
		VolumeID:             o.VolumeID,
		VolumeLastKnownState: o.VolumeLastKnownState,
		DiskSpec:             o.DiskSpec,
		DiskNumber:           o.DiskNumber,
	}
}

func (o *AddMirrorRequest) xxx_FromOp(ctx context.Context, op *xxx_AddMirrorOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VolumeID = op.VolumeID
	o.VolumeLastKnownState = op.VolumeLastKnownState
	o.DiskSpec = op.DiskSpec
	o.DiskNumber = op.DiskNumber
}
func (o *AddMirrorRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AddMirrorRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddMirrorOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddMirrorResponse structure represents the AddMirror operation response
type AddMirrorResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	DiskNumber      int32          `idl:"name:diskNumber" json:"disk_number"`
	PartitionNumber int32          `idl:"name:partitionNumber" json:"partition_number"`
	TaskInfo        *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The AddMirror return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AddMirrorResponse) xxx_ToOp(ctx context.Context) *xxx_AddMirrorOperation {
	if o == nil {
		return &xxx_AddMirrorOperation{}
	}
	return &xxx_AddMirrorOperation{
		That:            o.That,
		DiskNumber:      o.DiskNumber,
		PartitionNumber: o.PartitionNumber,
		TaskInfo:        o.TaskInfo,
		Return:          o.Return,
	}
}

func (o *AddMirrorResponse) xxx_FromOp(ctx context.Context, op *xxx_AddMirrorOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DiskNumber = op.DiskNumber
	o.PartitionNumber = op.PartitionNumber
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *AddMirrorResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AddMirrorResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddMirrorOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoveMirrorOperation structure represents the RemoveMirror operation
type xxx_RemoveMirrorOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	VolumeID             int64          `idl:"name:volumeId" json:"volume_id"`
	VolumeLastKnownState int64          `idl:"name:volumeLastKnownState" json:"volume_last_known_state"`
	DiskID               int64          `idl:"name:diskId" json:"disk_id"`
	DiskLastKnownState   int64          `idl:"name:diskLastKnownState" json:"disk_last_known_state"`
	TaskInfo             *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoveMirrorOperation) OpNum() int { return 37 }

func (o *xxx_RemoveMirrorOperation) OpName() string { return "/IVolumeClient/v0/RemoveMirror" }

func (o *xxx_RemoveMirrorOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveMirrorOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.VolumeID); err != nil {
			return err
		}
	}
	// volumeLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.VolumeLastKnownState); err != nil {
			return err
		}
	}
	// diskId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.DiskID); err != nil {
			return err
		}
	}
	// diskLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.DiskLastKnownState); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveMirrorOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.VolumeID); err != nil {
			return err
		}
	}
	// volumeLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.VolumeLastKnownState); err != nil {
			return err
		}
	}
	// diskId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.DiskID); err != nil {
			return err
		}
	}
	// diskLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.DiskLastKnownState); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveMirrorOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveMirrorOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveMirrorOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RemoveMirrorRequest structure represents the RemoveMirror operation request
type RemoveMirrorRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	VolumeID             int64          `idl:"name:volumeId" json:"volume_id"`
	VolumeLastKnownState int64          `idl:"name:volumeLastKnownState" json:"volume_last_known_state"`
	DiskID               int64          `idl:"name:diskId" json:"disk_id"`
	DiskLastKnownState   int64          `idl:"name:diskLastKnownState" json:"disk_last_known_state"`
}

func (o *RemoveMirrorRequest) xxx_ToOp(ctx context.Context) *xxx_RemoveMirrorOperation {
	if o == nil {
		return &xxx_RemoveMirrorOperation{}
	}
	return &xxx_RemoveMirrorOperation{
		This:                 o.This,
		VolumeID:             o.VolumeID,
		VolumeLastKnownState: o.VolumeLastKnownState,
		DiskID:               o.DiskID,
		DiskLastKnownState:   o.DiskLastKnownState,
	}
}

func (o *RemoveMirrorRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoveMirrorOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VolumeID = op.VolumeID
	o.VolumeLastKnownState = op.VolumeLastKnownState
	o.DiskID = op.DiskID
	o.DiskLastKnownState = op.DiskLastKnownState
}
func (o *RemoveMirrorRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RemoveMirrorRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveMirrorOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoveMirrorResponse structure represents the RemoveMirror operation response
type RemoveMirrorResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The RemoveMirror return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RemoveMirrorResponse) xxx_ToOp(ctx context.Context) *xxx_RemoveMirrorOperation {
	if o == nil {
		return &xxx_RemoveMirrorOperation{}
	}
	return &xxx_RemoveMirrorOperation{
		That:     o.That,
		TaskInfo: o.TaskInfo,
		Return:   o.Return,
	}
}

func (o *RemoveMirrorResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoveMirrorOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *RemoveMirrorResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RemoveMirrorResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveMirrorOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SplitMirrorOperation structure represents the SplitMirror operation
type xxx_SplitMirrorOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	VolumeID             int64          `idl:"name:volumeId" json:"volume_id"`
	VolumeLastKnownState int64          `idl:"name:volumeLastKnownState" json:"volume_last_known_state"`
	DiskID               int64          `idl:"name:diskId" json:"disk_id"`
	DiskLastKnownState   int64          `idl:"name:diskLastKnownState" json:"disk_last_known_state"`
	Letter               uint16         `idl:"name:letter" json:"letter"`
	LetterLastKnownState int64          `idl:"name:letterLastKnownState" json:"letter_last_known_state"`
	TaskInfo             *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SplitMirrorOperation) OpNum() int { return 38 }

func (o *xxx_SplitMirrorOperation) OpName() string { return "/IVolumeClient/v0/SplitMirror" }

func (o *xxx_SplitMirrorOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SplitMirrorOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.VolumeID); err != nil {
			return err
		}
	}
	// volumeLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.VolumeLastKnownState); err != nil {
			return err
		}
	}
	// diskId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.DiskID); err != nil {
			return err
		}
	}
	// diskLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.DiskLastKnownState); err != nil {
			return err
		}
	}
	// letter {in} (1:(wchar))
	{
		if err := w.WriteData(o.Letter); err != nil {
			return err
		}
	}
	// letterLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.LetterLastKnownState); err != nil {
			return err
		}
	}
	// tinfo {in, out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_SplitMirrorOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.VolumeID); err != nil {
			return err
		}
	}
	// volumeLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.VolumeLastKnownState); err != nil {
			return err
		}
	}
	// diskId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.DiskID); err != nil {
			return err
		}
	}
	// diskLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.DiskLastKnownState); err != nil {
			return err
		}
	}
	// letter {in} (1:(wchar))
	{
		if err := w.ReadData(&o.Letter); err != nil {
			return err
		}
	}
	// letterLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.LetterLastKnownState); err != nil {
			return err
		}
	}
	// tinfo {in, out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SplitMirrorOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SplitMirrorOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {in, out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SplitMirrorOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {in, out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SplitMirrorRequest structure represents the SplitMirror operation request
type SplitMirrorRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	VolumeID             int64          `idl:"name:volumeId" json:"volume_id"`
	VolumeLastKnownState int64          `idl:"name:volumeLastKnownState" json:"volume_last_known_state"`
	DiskID               int64          `idl:"name:diskId" json:"disk_id"`
	DiskLastKnownState   int64          `idl:"name:diskLastKnownState" json:"disk_last_known_state"`
	Letter               uint16         `idl:"name:letter" json:"letter"`
	LetterLastKnownState int64          `idl:"name:letterLastKnownState" json:"letter_last_known_state"`
	TaskInfo             *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
}

func (o *SplitMirrorRequest) xxx_ToOp(ctx context.Context) *xxx_SplitMirrorOperation {
	if o == nil {
		return &xxx_SplitMirrorOperation{}
	}
	return &xxx_SplitMirrorOperation{
		This:                 o.This,
		VolumeID:             o.VolumeID,
		VolumeLastKnownState: o.VolumeLastKnownState,
		DiskID:               o.DiskID,
		DiskLastKnownState:   o.DiskLastKnownState,
		Letter:               o.Letter,
		LetterLastKnownState: o.LetterLastKnownState,
		TaskInfo:             o.TaskInfo,
	}
}

func (o *SplitMirrorRequest) xxx_FromOp(ctx context.Context, op *xxx_SplitMirrorOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VolumeID = op.VolumeID
	o.VolumeLastKnownState = op.VolumeLastKnownState
	o.DiskID = op.DiskID
	o.DiskLastKnownState = op.DiskLastKnownState
	o.Letter = op.Letter
	o.LetterLastKnownState = op.LetterLastKnownState
	o.TaskInfo = op.TaskInfo
}
func (o *SplitMirrorRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SplitMirrorRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SplitMirrorOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SplitMirrorResponse structure represents the SplitMirror operation response
type SplitMirrorResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The SplitMirror return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SplitMirrorResponse) xxx_ToOp(ctx context.Context) *xxx_SplitMirrorOperation {
	if o == nil {
		return &xxx_SplitMirrorOperation{}
	}
	return &xxx_SplitMirrorOperation{
		That:     o.That,
		TaskInfo: o.TaskInfo,
		Return:   o.Return,
	}
}

func (o *SplitMirrorResponse) xxx_FromOp(ctx context.Context, op *xxx_SplitMirrorOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *SplitMirrorResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SplitMirrorResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SplitMirrorOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_InitializeDiskOperation structure represents the InitializeDisk operation
type xxx_InitializeDiskOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	DiskID             int64          `idl:"name:diskId" json:"disk_id"`
	DiskLastKnownState int64          `idl:"name:diskLastKnownState" json:"disk_last_known_state"`
	TaskInfo           *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_InitializeDiskOperation) OpNum() int { return 39 }

func (o *xxx_InitializeDiskOperation) OpName() string { return "/IVolumeClient/v0/InitializeDisk" }

func (o *xxx_InitializeDiskOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitializeDiskOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// diskId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.DiskID); err != nil {
			return err
		}
	}
	// diskLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.DiskLastKnownState); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitializeDiskOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// diskId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.DiskID); err != nil {
			return err
		}
	}
	// diskLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.DiskLastKnownState); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitializeDiskOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitializeDiskOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitializeDiskOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// InitializeDiskRequest structure represents the InitializeDisk operation request
type InitializeDiskRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// diskId: Specifies the OID of the disk to initialize.
	DiskID int64 `idl:"name:diskId" json:"disk_id"`
	// diskLastKnownState: Disk's last known modification sequence number.
	DiskLastKnownState int64 `idl:"name:diskLastKnownState" json:"disk_last_known_state"`
}

func (o *InitializeDiskRequest) xxx_ToOp(ctx context.Context) *xxx_InitializeDiskOperation {
	if o == nil {
		return &xxx_InitializeDiskOperation{}
	}
	return &xxx_InitializeDiskOperation{
		This:               o.This,
		DiskID:             o.DiskID,
		DiskLastKnownState: o.DiskLastKnownState,
	}
}

func (o *InitializeDiskRequest) xxx_FromOp(ctx context.Context, op *xxx_InitializeDiskOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
	o.DiskLastKnownState = op.DiskLastKnownState
}
func (o *InitializeDiskRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *InitializeDiskRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InitializeDiskOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// InitializeDiskResponse structure represents the InitializeDisk operation response
type InitializeDiskResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// tinfo: Pointer to a TASK_INFO structure that the client can use to track the request's
	// progress.
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The InitializeDisk return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *InitializeDiskResponse) xxx_ToOp(ctx context.Context) *xxx_InitializeDiskOperation {
	if o == nil {
		return &xxx_InitializeDiskOperation{}
	}
	return &xxx_InitializeDiskOperation{
		That:     o.That,
		TaskInfo: o.TaskInfo,
		Return:   o.Return,
	}
}

func (o *InitializeDiskResponse) xxx_FromOp(ctx context.Context, op *xxx_InitializeDiskOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *InitializeDiskResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *InitializeDiskResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InitializeDiskOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_UninitializeDiskOperation structure represents the UninitializeDisk operation
type xxx_UninitializeDiskOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	DiskID             int64          `idl:"name:diskId" json:"disk_id"`
	DiskLastKnownState int64          `idl:"name:diskLastKnownState" json:"disk_last_known_state"`
	TaskInfo           *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_UninitializeDiskOperation) OpNum() int { return 40 }

func (o *xxx_UninitializeDiskOperation) OpName() string { return "/IVolumeClient/v0/UninitializeDisk" }

func (o *xxx_UninitializeDiskOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UninitializeDiskOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// diskId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.DiskID); err != nil {
			return err
		}
	}
	// diskLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.DiskLastKnownState); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UninitializeDiskOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// diskId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.DiskID); err != nil {
			return err
		}
	}
	// diskLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.DiskLastKnownState); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UninitializeDiskOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UninitializeDiskOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UninitializeDiskOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// UninitializeDiskRequest structure represents the UninitializeDisk operation request
type UninitializeDiskRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	DiskID             int64          `idl:"name:diskId" json:"disk_id"`
	DiskLastKnownState int64          `idl:"name:diskLastKnownState" json:"disk_last_known_state"`
}

func (o *UninitializeDiskRequest) xxx_ToOp(ctx context.Context) *xxx_UninitializeDiskOperation {
	if o == nil {
		return &xxx_UninitializeDiskOperation{}
	}
	return &xxx_UninitializeDiskOperation{
		This:               o.This,
		DiskID:             o.DiskID,
		DiskLastKnownState: o.DiskLastKnownState,
	}
}

func (o *UninitializeDiskRequest) xxx_FromOp(ctx context.Context, op *xxx_UninitializeDiskOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
	o.DiskLastKnownState = op.DiskLastKnownState
}
func (o *UninitializeDiskRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *UninitializeDiskRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UninitializeDiskOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// UninitializeDiskResponse structure represents the UninitializeDisk operation response
type UninitializeDiskResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The UninitializeDisk return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *UninitializeDiskResponse) xxx_ToOp(ctx context.Context) *xxx_UninitializeDiskOperation {
	if o == nil {
		return &xxx_UninitializeDiskOperation{}
	}
	return &xxx_UninitializeDiskOperation{
		That:     o.That,
		TaskInfo: o.TaskInfo,
		Return:   o.Return,
	}
}

func (o *UninitializeDiskResponse) xxx_FromOp(ctx context.Context, op *xxx_UninitializeDiskOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *UninitializeDiskResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *UninitializeDiskResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UninitializeDiskOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReConnectDiskOperation structure represents the ReConnectDisk operation
type xxx_ReConnectDiskOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	DiskID   int64          `idl:"name:diskId" json:"disk_id"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ReConnectDiskOperation) OpNum() int { return 41 }

func (o *xxx_ReConnectDiskOperation) OpName() string { return "/IVolumeClient/v0/ReConnectDisk" }

func (o *xxx_ReConnectDiskOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReConnectDiskOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// diskId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.DiskID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReConnectDiskOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// diskId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.DiskID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReConnectDiskOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReConnectDiskOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReConnectDiskOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ReConnectDiskRequest structure represents the ReConnectDisk operation request
type ReConnectDiskRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	DiskID int64          `idl:"name:diskId" json:"disk_id"`
}

func (o *ReConnectDiskRequest) xxx_ToOp(ctx context.Context) *xxx_ReConnectDiskOperation {
	if o == nil {
		return &xxx_ReConnectDiskOperation{}
	}
	return &xxx_ReConnectDiskOperation{
		This:   o.This,
		DiskID: o.DiskID,
	}
}

func (o *ReConnectDiskRequest) xxx_FromOp(ctx context.Context, op *xxx_ReConnectDiskOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
}
func (o *ReConnectDiskRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ReConnectDiskRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReConnectDiskOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReConnectDiskResponse structure represents the ReConnectDisk operation response
type ReConnectDiskResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The ReConnectDisk return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReConnectDiskResponse) xxx_ToOp(ctx context.Context) *xxx_ReConnectDiskOperation {
	if o == nil {
		return &xxx_ReConnectDiskOperation{}
	}
	return &xxx_ReConnectDiskOperation{
		That:     o.That,
		TaskInfo: o.TaskInfo,
		Return:   o.Return,
	}
}

func (o *ReConnectDiskResponse) xxx_FromOp(ctx context.Context, op *xxx_ReConnectDiskOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *ReConnectDiskResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ReConnectDiskResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReConnectDiskOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ImportDiskGroupOperation structure represents the ImportDiskGroup operation
type xxx_ImportDiskGroupOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	DiskGroupIDLength int32          `idl:"name:cchDgid" json:"disk_group_id_length"`
	DiskGroupID       []byte         `idl:"name:dgid;size_is:(cchDgid)" json:"disk_group_id"`
	TaskInfo          *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ImportDiskGroupOperation) OpNum() int { return 43 }

func (o *xxx_ImportDiskGroupOperation) OpName() string { return "/IVolumeClient/v0/ImportDiskGroup" }

func (o *xxx_ImportDiskGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.DiskGroupID != nil && o.DiskGroupIDLength == 0 {
		o.DiskGroupIDLength = int32(len(o.DiskGroupID))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportDiskGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// cchDgid {in} (1:(int32))
	{
		if err := w.WriteData(o.DiskGroupIDLength); err != nil {
			return err
		}
	}
	// dgid {in} (1:{pointer=ref}*(1)[dim:0,size_is=cchDgid](uint8))
	{
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
	}
	return nil
}

func (o *xxx_ImportDiskGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cchDgid {in} (1:(int32))
	{
		if err := w.ReadData(&o.DiskGroupIDLength); err != nil {
			return err
		}
	}
	// dgid {in} (1:{pointer=ref}*(1)[dim:0,size_is=cchDgid](uint8))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
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
	}
	return nil
}

func (o *xxx_ImportDiskGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportDiskGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportDiskGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ImportDiskGroupRequest structure represents the ImportDiskGroup operation request
type ImportDiskGroupRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	DiskGroupIDLength int32          `idl:"name:cchDgid" json:"disk_group_id_length"`
	DiskGroupID       []byte         `idl:"name:dgid;size_is:(cchDgid)" json:"disk_group_id"`
}

func (o *ImportDiskGroupRequest) xxx_ToOp(ctx context.Context) *xxx_ImportDiskGroupOperation {
	if o == nil {
		return &xxx_ImportDiskGroupOperation{}
	}
	return &xxx_ImportDiskGroupOperation{
		This:              o.This,
		DiskGroupIDLength: o.DiskGroupIDLength,
		DiskGroupID:       o.DiskGroupID,
	}
}

func (o *ImportDiskGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_ImportDiskGroupOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskGroupIDLength = op.DiskGroupIDLength
	o.DiskGroupID = op.DiskGroupID
}
func (o *ImportDiskGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ImportDiskGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ImportDiskGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ImportDiskGroupResponse structure represents the ImportDiskGroup operation response
type ImportDiskGroupResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The ImportDiskGroup return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ImportDiskGroupResponse) xxx_ToOp(ctx context.Context) *xxx_ImportDiskGroupOperation {
	if o == nil {
		return &xxx_ImportDiskGroupOperation{}
	}
	return &xxx_ImportDiskGroupOperation{
		That:     o.That,
		TaskInfo: o.TaskInfo,
		Return:   o.Return,
	}
}

func (o *ImportDiskGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_ImportDiskGroupOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *ImportDiskGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ImportDiskGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ImportDiskGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DiskMergeQueryOperation structure represents the DiskMergeQuery operation
type xxx_DiskMergeQueryOperation struct {
	This              *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat          `idl:"name:That" json:"that"`
	DiskGroupIDLength int32                   `idl:"name:cchDgid" json:"disk_group_id_length"`
	DiskGroupID       []byte                  `idl:"name:dgid;size_is:(cchDgid)" json:"disk_group_id"`
	DisksLength       int32                   `idl:"name:numDisks" json:"disks_length"`
	DiskList          []int64                 `idl:"name:diskList;size_is:(numDisks)" json:"disk_list"`
	MergeConfigTID    int64                   `idl:"name:merge_config_tid" json:"merge_config_tid"`
	RIDsLength        int32                   `idl:"name:numRids" json:"rids_length"`
	MergeDMRIDs       []int64                 `idl:"name:merge_dm_rids;size_is:(, numRids)" json:"merge_dm_rids"`
	ObjectsLength     int32                   `idl:"name:numObjects" json:"objects_length"`
	MergeObjectInfo   []*dmrp.MergeObjectInfo `idl:"name:mergeObjectInfo;size_is:(, numObjects)" json:"merge_object_info"`
	Flags             uint32                  `idl:"name:flags" json:"flags"`
	TaskInfo          *dmrp.TaskInfo          `idl:"name:tinfo" json:"task_info"`
	Return            int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_DiskMergeQueryOperation) OpNum() int { return 44 }

func (o *xxx_DiskMergeQueryOperation) OpName() string { return "/IVolumeClient/v0/DiskMergeQuery" }

func (o *xxx_DiskMergeQueryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.DiskGroupID != nil && o.DiskGroupIDLength == 0 {
		o.DiskGroupIDLength = int32(len(o.DiskGroupID))
	}
	if o.DiskList != nil && o.DisksLength == 0 {
		o.DisksLength = int32(len(o.DiskList))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DiskMergeQueryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// cchDgid {in} (1:(int32))
	{
		if err := w.WriteData(o.DiskGroupIDLength); err != nil {
			return err
		}
	}
	// dgid {in} (1:{pointer=ref}*(1)[dim:0,size_is=cchDgid](uint8))
	{
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
	}
	// numDisks {in} (1:(int32))
	{
		if err := w.WriteData(o.DisksLength); err != nil {
			return err
		}
	}
	// diskList {in} (1:{pointer=ref}*(1))(2:{alias=LdmObjectId, names=LONGLONG}[dim:0,size_is=numDisks](int64))
	{
		dimSize1 := uint64(o.DisksLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.DiskList {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.DiskList[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.DiskList); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(int64(0)); err != nil {
				return err
			}
		}
	}
	// flags {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DiskMergeQueryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cchDgid {in} (1:(int32))
	{
		if err := w.ReadData(&o.DiskGroupIDLength); err != nil {
			return err
		}
	}
	// dgid {in} (1:{pointer=ref}*(1)[dim:0,size_is=cchDgid](uint8))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
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
	}
	// numDisks {in} (1:(int32))
	{
		if err := w.ReadData(&o.DisksLength); err != nil {
			return err
		}
	}
	// diskList {in} (1:{pointer=ref}*(1))(2:{alias=LdmObjectId, names=LONGLONG}[dim:0,size_is=numDisks](int64))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.DiskList", sizeInfo[0])
		}
		o.DiskList = make([]int64, sizeInfo[0])
		for i1 := range o.DiskList {
			i1 := i1
			if err := w.ReadData(&o.DiskList[i1]); err != nil {
				return err
			}
		}
	}
	// flags {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DiskMergeQueryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.MergeDMRIDs != nil && o.RIDsLength == 0 {
		o.RIDsLength = int32(len(o.MergeDMRIDs))
	}
	if o.MergeObjectInfo != nil && o.ObjectsLength == 0 {
		o.ObjectsLength = int32(len(o.MergeObjectInfo))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DiskMergeQueryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// merge_config_tid {out} (1:{pointer=ref}*(1)(int64))
	{
		if err := w.WriteData(o.MergeConfigTID); err != nil {
			return err
		}
	}
	// numRids {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.RIDsLength); err != nil {
			return err
		}
	}
	// merge_dm_rids {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=numRids](int64))
	{
		if o.MergeDMRIDs != nil || o.RIDsLength > 0 {
			_ptr_merge_dm_rids := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.RIDsLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.MergeDMRIDs {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.MergeDMRIDs[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.MergeDMRIDs); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(int64(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MergeDMRIDs, _ptr_merge_dm_rids); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// numObjects {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.ObjectsLength); err != nil {
			return err
		}
	}
	// mergeObjectInfo {out} (1:{pointer=ref}*(2)*(1))(2:{alias=MERGE_OBJECT_INFO}[dim:0,size_is=numObjects](struct))
	{
		if o.MergeObjectInfo != nil || o.ObjectsLength > 0 {
			_ptr_mergeObjectInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ObjectsLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.MergeObjectInfo {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.MergeObjectInfo[i1] != nil {
						if err := o.MergeObjectInfo[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dmrp.MergeObjectInfo{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.MergeObjectInfo); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&dmrp.MergeObjectInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MergeObjectInfo, _ptr_mergeObjectInfo); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// flags {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DiskMergeQueryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// merge_config_tid {out} (1:{pointer=ref}*(1)(int64))
	{
		if err := w.ReadData(&o.MergeConfigTID); err != nil {
			return err
		}
	}
	// numRids {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.RIDsLength); err != nil {
			return err
		}
	}
	// merge_dm_rids {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=numRids](int64))
	{
		_ptr_merge_dm_rids := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.MergeDMRIDs", sizeInfo[0])
			}
			o.MergeDMRIDs = make([]int64, sizeInfo[0])
			for i1 := range o.MergeDMRIDs {
				i1 := i1
				if err := w.ReadData(&o.MergeDMRIDs[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_merge_dm_rids := func(ptr interface{}) { o.MergeDMRIDs = *ptr.(*[]int64) }
		if err := w.ReadPointer(&o.MergeDMRIDs, _s_merge_dm_rids, _ptr_merge_dm_rids); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// numObjects {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.ObjectsLength); err != nil {
			return err
		}
	}
	// mergeObjectInfo {out} (1:{pointer=ref}*(2)*(1))(2:{alias=MERGE_OBJECT_INFO}[dim:0,size_is=numObjects](struct))
	{
		_ptr_mergeObjectInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.MergeObjectInfo", sizeInfo[0])
			}
			o.MergeObjectInfo = make([]*dmrp.MergeObjectInfo, sizeInfo[0])
			for i1 := range o.MergeObjectInfo {
				i1 := i1
				if o.MergeObjectInfo[i1] == nil {
					o.MergeObjectInfo[i1] = &dmrp.MergeObjectInfo{}
				}
				if err := o.MergeObjectInfo[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_mergeObjectInfo := func(ptr interface{}) { o.MergeObjectInfo = *ptr.(*[]*dmrp.MergeObjectInfo) }
		if err := w.ReadPointer(&o.MergeObjectInfo, _s_mergeObjectInfo, _ptr_mergeObjectInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// flags {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DiskMergeQueryRequest structure represents the DiskMergeQuery operation request
type DiskMergeQueryRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	DiskGroupIDLength int32          `idl:"name:cchDgid" json:"disk_group_id_length"`
	DiskGroupID       []byte         `idl:"name:dgid;size_is:(cchDgid)" json:"disk_group_id"`
	DisksLength       int32          `idl:"name:numDisks" json:"disks_length"`
	DiskList          []int64        `idl:"name:diskList;size_is:(numDisks)" json:"disk_list"`
	Flags             uint32         `idl:"name:flags" json:"flags"`
}

func (o *DiskMergeQueryRequest) xxx_ToOp(ctx context.Context) *xxx_DiskMergeQueryOperation {
	if o == nil {
		return &xxx_DiskMergeQueryOperation{}
	}
	return &xxx_DiskMergeQueryOperation{
		This:              o.This,
		DiskGroupIDLength: o.DiskGroupIDLength,
		DiskGroupID:       o.DiskGroupID,
		DisksLength:       o.DisksLength,
		DiskList:          o.DiskList,
		Flags:             o.Flags,
	}
}

func (o *DiskMergeQueryRequest) xxx_FromOp(ctx context.Context, op *xxx_DiskMergeQueryOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskGroupIDLength = op.DiskGroupIDLength
	o.DiskGroupID = op.DiskGroupID
	o.DisksLength = op.DisksLength
	o.DiskList = op.DiskList
	o.Flags = op.Flags
}
func (o *DiskMergeQueryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DiskMergeQueryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DiskMergeQueryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DiskMergeQueryResponse structure represents the DiskMergeQuery operation response
type DiskMergeQueryResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That            *dcom.ORPCThat          `idl:"name:That" json:"that"`
	MergeConfigTID  int64                   `idl:"name:merge_config_tid" json:"merge_config_tid"`
	RIDsLength      int32                   `idl:"name:numRids" json:"rids_length"`
	MergeDMRIDs     []int64                 `idl:"name:merge_dm_rids;size_is:(, numRids)" json:"merge_dm_rids"`
	ObjectsLength   int32                   `idl:"name:numObjects" json:"objects_length"`
	MergeObjectInfo []*dmrp.MergeObjectInfo `idl:"name:mergeObjectInfo;size_is:(, numObjects)" json:"merge_object_info"`
	Flags           uint32                  `idl:"name:flags" json:"flags"`
	TaskInfo        *dmrp.TaskInfo          `idl:"name:tinfo" json:"task_info"`
	// Return: The DiskMergeQuery return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DiskMergeQueryResponse) xxx_ToOp(ctx context.Context) *xxx_DiskMergeQueryOperation {
	if o == nil {
		return &xxx_DiskMergeQueryOperation{}
	}
	return &xxx_DiskMergeQueryOperation{
		That:            o.That,
		MergeConfigTID:  o.MergeConfigTID,
		RIDsLength:      o.RIDsLength,
		MergeDMRIDs:     o.MergeDMRIDs,
		ObjectsLength:   o.ObjectsLength,
		MergeObjectInfo: o.MergeObjectInfo,
		Flags:           o.Flags,
		TaskInfo:        o.TaskInfo,
		Return:          o.Return,
	}
}

func (o *DiskMergeQueryResponse) xxx_FromOp(ctx context.Context, op *xxx_DiskMergeQueryOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MergeConfigTID = op.MergeConfigTID
	o.RIDsLength = op.RIDsLength
	o.MergeDMRIDs = op.MergeDMRIDs
	o.ObjectsLength = op.ObjectsLength
	o.MergeObjectInfo = op.MergeObjectInfo
	o.Flags = op.Flags
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *DiskMergeQueryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DiskMergeQueryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DiskMergeQueryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DiskMergeOperation structure represents the DiskMerge operation
type xxx_DiskMergeOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	DiskGroupIDLength int32          `idl:"name:cchDgid" json:"disk_group_id_length"`
	DiskGroupID       []byte         `idl:"name:dgid;size_is:(cchDgid)" json:"disk_group_id"`
	DisksLength       int32          `idl:"name:numDisks" json:"disks_length"`
	DiskList          []int64        `idl:"name:diskList;size_is:(numDisks)" json:"disk_list"`
	MergeConfigTID    int64          `idl:"name:merge_config_tid" json:"merge_config_tid"`
	RIDsLength        int32          `idl:"name:numRids" json:"rids_length"`
	MergeDMRIDs       []int64        `idl:"name:merge_dm_rids;size_is:(numRids)" json:"merge_dm_rids"`
	TaskInfo          *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DiskMergeOperation) OpNum() int { return 45 }

func (o *xxx_DiskMergeOperation) OpName() string { return "/IVolumeClient/v0/DiskMerge" }

func (o *xxx_DiskMergeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.DiskGroupID != nil && o.DiskGroupIDLength == 0 {
		o.DiskGroupIDLength = int32(len(o.DiskGroupID))
	}
	if o.DiskList != nil && o.DisksLength == 0 {
		o.DisksLength = int32(len(o.DiskList))
	}
	if o.MergeDMRIDs != nil && o.RIDsLength == 0 {
		o.RIDsLength = int32(len(o.MergeDMRIDs))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DiskMergeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// cchDgid {in} (1:(int32))
	{
		if err := w.WriteData(o.DiskGroupIDLength); err != nil {
			return err
		}
	}
	// dgid {in} (1:{pointer=ref}*(1)[dim:0,size_is=cchDgid](uint8))
	{
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
	}
	// numDisks {in} (1:(int32))
	{
		if err := w.WriteData(o.DisksLength); err != nil {
			return err
		}
	}
	// diskList {in} (1:{pointer=ref}*(1))(2:{alias=LdmObjectId, names=LONGLONG}[dim:0,size_is=numDisks](int64))
	{
		dimSize1 := uint64(o.DisksLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.DiskList {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.DiskList[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.DiskList); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(int64(0)); err != nil {
				return err
			}
		}
	}
	// merge_config_tid {in} (1:(int64))
	{
		if err := w.WriteData(o.MergeConfigTID); err != nil {
			return err
		}
	}
	// numRids {in} (1:(int32))
	{
		if err := w.WriteData(o.RIDsLength); err != nil {
			return err
		}
	}
	// merge_dm_rids {in} (1:{pointer=ref}*(1)[dim:0,size_is=numRids](int64))
	{
		dimSize1 := uint64(o.RIDsLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.MergeDMRIDs {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.MergeDMRIDs[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.MergeDMRIDs); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(int64(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_DiskMergeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cchDgid {in} (1:(int32))
	{
		if err := w.ReadData(&o.DiskGroupIDLength); err != nil {
			return err
		}
	}
	// dgid {in} (1:{pointer=ref}*(1)[dim:0,size_is=cchDgid](uint8))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
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
	}
	// numDisks {in} (1:(int32))
	{
		if err := w.ReadData(&o.DisksLength); err != nil {
			return err
		}
	}
	// diskList {in} (1:{pointer=ref}*(1))(2:{alias=LdmObjectId, names=LONGLONG}[dim:0,size_is=numDisks](int64))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.DiskList", sizeInfo[0])
		}
		o.DiskList = make([]int64, sizeInfo[0])
		for i1 := range o.DiskList {
			i1 := i1
			if err := w.ReadData(&o.DiskList[i1]); err != nil {
				return err
			}
		}
	}
	// merge_config_tid {in} (1:(int64))
	{
		if err := w.ReadData(&o.MergeConfigTID); err != nil {
			return err
		}
	}
	// numRids {in} (1:(int32))
	{
		if err := w.ReadData(&o.RIDsLength); err != nil {
			return err
		}
	}
	// merge_dm_rids {in} (1:{pointer=ref}*(1)[dim:0,size_is=numRids](int64))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.MergeDMRIDs", sizeInfo[0])
		}
		o.MergeDMRIDs = make([]int64, sizeInfo[0])
		for i1 := range o.MergeDMRIDs {
			i1 := i1
			if err := w.ReadData(&o.MergeDMRIDs[i1]); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_DiskMergeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DiskMergeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DiskMergeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DiskMergeRequest structure represents the DiskMerge operation request
type DiskMergeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	DiskGroupIDLength int32          `idl:"name:cchDgid" json:"disk_group_id_length"`
	DiskGroupID       []byte         `idl:"name:dgid;size_is:(cchDgid)" json:"disk_group_id"`
	DisksLength       int32          `idl:"name:numDisks" json:"disks_length"`
	DiskList          []int64        `idl:"name:diskList;size_is:(numDisks)" json:"disk_list"`
	MergeConfigTID    int64          `idl:"name:merge_config_tid" json:"merge_config_tid"`
	RIDsLength        int32          `idl:"name:numRids" json:"rids_length"`
	MergeDMRIDs       []int64        `idl:"name:merge_dm_rids;size_is:(numRids)" json:"merge_dm_rids"`
}

func (o *DiskMergeRequest) xxx_ToOp(ctx context.Context) *xxx_DiskMergeOperation {
	if o == nil {
		return &xxx_DiskMergeOperation{}
	}
	return &xxx_DiskMergeOperation{
		This:              o.This,
		DiskGroupIDLength: o.DiskGroupIDLength,
		DiskGroupID:       o.DiskGroupID,
		DisksLength:       o.DisksLength,
		DiskList:          o.DiskList,
		MergeConfigTID:    o.MergeConfigTID,
		RIDsLength:        o.RIDsLength,
		MergeDMRIDs:       o.MergeDMRIDs,
	}
}

func (o *DiskMergeRequest) xxx_FromOp(ctx context.Context, op *xxx_DiskMergeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskGroupIDLength = op.DiskGroupIDLength
	o.DiskGroupID = op.DiskGroupID
	o.DisksLength = op.DisksLength
	o.DiskList = op.DiskList
	o.MergeConfigTID = op.MergeConfigTID
	o.RIDsLength = op.RIDsLength
	o.MergeDMRIDs = op.MergeDMRIDs
}
func (o *DiskMergeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DiskMergeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DiskMergeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DiskMergeResponse structure represents the DiskMerge operation response
type DiskMergeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The DiskMerge return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DiskMergeResponse) xxx_ToOp(ctx context.Context) *xxx_DiskMergeOperation {
	if o == nil {
		return &xxx_DiskMergeOperation{}
	}
	return &xxx_DiskMergeOperation{
		That:     o.That,
		TaskInfo: o.TaskInfo,
		Return:   o.Return,
	}
}

func (o *DiskMergeResponse) xxx_FromOp(ctx context.Context, op *xxx_DiskMergeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *DiskMergeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DiskMergeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DiskMergeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReAttachDiskOperation structure represents the ReAttachDisk operation
type xxx_ReAttachDiskOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	DiskID             int64          `idl:"name:diskId" json:"disk_id"`
	DiskLastKnownState int64          `idl:"name:diskLastKnownState" json:"disk_last_known_state"`
	TaskInfo           *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ReAttachDiskOperation) OpNum() int { return 47 }

func (o *xxx_ReAttachDiskOperation) OpName() string { return "/IVolumeClient/v0/ReAttachDisk" }

func (o *xxx_ReAttachDiskOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReAttachDiskOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// diskId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.DiskID); err != nil {
			return err
		}
	}
	// diskLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.DiskLastKnownState); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReAttachDiskOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// diskId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.DiskID); err != nil {
			return err
		}
	}
	// diskLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.DiskLastKnownState); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReAttachDiskOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReAttachDiskOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReAttachDiskOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ReAttachDiskRequest structure represents the ReAttachDisk operation request
type ReAttachDiskRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	DiskID             int64          `idl:"name:diskId" json:"disk_id"`
	DiskLastKnownState int64          `idl:"name:diskLastKnownState" json:"disk_last_known_state"`
}

func (o *ReAttachDiskRequest) xxx_ToOp(ctx context.Context) *xxx_ReAttachDiskOperation {
	if o == nil {
		return &xxx_ReAttachDiskOperation{}
	}
	return &xxx_ReAttachDiskOperation{
		This:               o.This,
		DiskID:             o.DiskID,
		DiskLastKnownState: o.DiskLastKnownState,
	}
}

func (o *ReAttachDiskRequest) xxx_FromOp(ctx context.Context, op *xxx_ReAttachDiskOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
	o.DiskLastKnownState = op.DiskLastKnownState
}
func (o *ReAttachDiskRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ReAttachDiskRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReAttachDiskOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReAttachDiskResponse structure represents the ReAttachDisk operation response
type ReAttachDiskResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The ReAttachDisk return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReAttachDiskResponse) xxx_ToOp(ctx context.Context) *xxx_ReAttachDiskOperation {
	if o == nil {
		return &xxx_ReAttachDiskOperation{}
	}
	return &xxx_ReAttachDiskOperation{
		That:     o.That,
		TaskInfo: o.TaskInfo,
		Return:   o.Return,
	}
}

func (o *ReAttachDiskResponse) xxx_FromOp(ctx context.Context, op *xxx_ReAttachDiskOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *ReAttachDiskResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ReAttachDiskResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReAttachDiskOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReplaceRAID5ColumnOperation structure represents the ReplaceRaid5Column operation
type xxx_ReplaceRAID5ColumnOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	VolumeID             int64          `idl:"name:volumeId" json:"volume_id"`
	VolumeLastKnownState int64          `idl:"name:volumeLastKnownState" json:"volume_last_known_state"`
	NewDiskID            int64          `idl:"name:newDiskId" json:"new_disk_id"`
	DiskLastKnownState   int64          `idl:"name:diskLastKnownState" json:"disk_last_known_state"`
	TaskInfo             *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ReplaceRAID5ColumnOperation) OpNum() int { return 51 }

func (o *xxx_ReplaceRAID5ColumnOperation) OpName() string {
	return "/IVolumeClient/v0/ReplaceRaid5Column"
}

func (o *xxx_ReplaceRAID5ColumnOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReplaceRAID5ColumnOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.VolumeID); err != nil {
			return err
		}
	}
	// volumeLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.VolumeLastKnownState); err != nil {
			return err
		}
	}
	// newDiskId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.NewDiskID); err != nil {
			return err
		}
	}
	// diskLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.DiskLastKnownState); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReplaceRAID5ColumnOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.VolumeID); err != nil {
			return err
		}
	}
	// volumeLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.VolumeLastKnownState); err != nil {
			return err
		}
	}
	// newDiskId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.NewDiskID); err != nil {
			return err
		}
	}
	// diskLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.DiskLastKnownState); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReplaceRAID5ColumnOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReplaceRAID5ColumnOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReplaceRAID5ColumnOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ReplaceRAID5ColumnRequest structure represents the ReplaceRaid5Column operation request
type ReplaceRAID5ColumnRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	VolumeID             int64          `idl:"name:volumeId" json:"volume_id"`
	VolumeLastKnownState int64          `idl:"name:volumeLastKnownState" json:"volume_last_known_state"`
	NewDiskID            int64          `idl:"name:newDiskId" json:"new_disk_id"`
	DiskLastKnownState   int64          `idl:"name:diskLastKnownState" json:"disk_last_known_state"`
}

func (o *ReplaceRAID5ColumnRequest) xxx_ToOp(ctx context.Context) *xxx_ReplaceRAID5ColumnOperation {
	if o == nil {
		return &xxx_ReplaceRAID5ColumnOperation{}
	}
	return &xxx_ReplaceRAID5ColumnOperation{
		This:                 o.This,
		VolumeID:             o.VolumeID,
		VolumeLastKnownState: o.VolumeLastKnownState,
		NewDiskID:            o.NewDiskID,
		DiskLastKnownState:   o.DiskLastKnownState,
	}
}

func (o *ReplaceRAID5ColumnRequest) xxx_FromOp(ctx context.Context, op *xxx_ReplaceRAID5ColumnOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VolumeID = op.VolumeID
	o.VolumeLastKnownState = op.VolumeLastKnownState
	o.NewDiskID = op.NewDiskID
	o.DiskLastKnownState = op.DiskLastKnownState
}
func (o *ReplaceRAID5ColumnRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ReplaceRAID5ColumnRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReplaceRAID5ColumnOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReplaceRAID5ColumnResponse structure represents the ReplaceRaid5Column operation response
type ReplaceRAID5ColumnResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The ReplaceRaid5Column return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReplaceRAID5ColumnResponse) xxx_ToOp(ctx context.Context) *xxx_ReplaceRAID5ColumnOperation {
	if o == nil {
		return &xxx_ReplaceRAID5ColumnOperation{}
	}
	return &xxx_ReplaceRAID5ColumnOperation{
		That:     o.That,
		TaskInfo: o.TaskInfo,
		Return:   o.Return,
	}
}

func (o *ReplaceRAID5ColumnResponse) xxx_FromOp(ctx context.Context, op *xxx_ReplaceRAID5ColumnOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *ReplaceRAID5ColumnResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ReplaceRAID5ColumnResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReplaceRAID5ColumnOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RestartVolumeOperation structure represents the RestartVolume operation
type xxx_RestartVolumeOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	VolumeID             int64          `idl:"name:volumeId" json:"volume_id"`
	VolumeLastKnownState int64          `idl:"name:volumeLastKnownState" json:"volume_last_known_state"`
	TaskInfo             *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RestartVolumeOperation) OpNum() int { return 52 }

func (o *xxx_RestartVolumeOperation) OpName() string { return "/IVolumeClient/v0/RestartVolume" }

func (o *xxx_RestartVolumeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RestartVolumeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.VolumeID); err != nil {
			return err
		}
	}
	// volumeLastKnownState {in} (1:(int64))
	{
		if err := w.WriteData(o.VolumeLastKnownState); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RestartVolumeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.VolumeID); err != nil {
			return err
		}
	}
	// volumeLastKnownState {in} (1:(int64))
	{
		if err := w.ReadData(&o.VolumeLastKnownState); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RestartVolumeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RestartVolumeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RestartVolumeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RestartVolumeRequest structure represents the RestartVolume operation request
type RestartVolumeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	VolumeID             int64          `idl:"name:volumeId" json:"volume_id"`
	VolumeLastKnownState int64          `idl:"name:volumeLastKnownState" json:"volume_last_known_state"`
}

func (o *RestartVolumeRequest) xxx_ToOp(ctx context.Context) *xxx_RestartVolumeOperation {
	if o == nil {
		return &xxx_RestartVolumeOperation{}
	}
	return &xxx_RestartVolumeOperation{
		This:                 o.This,
		VolumeID:             o.VolumeID,
		VolumeLastKnownState: o.VolumeLastKnownState,
	}
}

func (o *RestartVolumeRequest) xxx_FromOp(ctx context.Context, op *xxx_RestartVolumeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VolumeID = op.VolumeID
	o.VolumeLastKnownState = op.VolumeLastKnownState
}
func (o *RestartVolumeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RestartVolumeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RestartVolumeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RestartVolumeResponse structure represents the RestartVolume operation response
type RestartVolumeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The RestartVolume return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RestartVolumeResponse) xxx_ToOp(ctx context.Context) *xxx_RestartVolumeOperation {
	if o == nil {
		return &xxx_RestartVolumeOperation{}
	}
	return &xxx_RestartVolumeOperation{
		That:     o.That,
		TaskInfo: o.TaskInfo,
		Return:   o.Return,
	}
}

func (o *RestartVolumeResponse) xxx_FromOp(ctx context.Context, op *xxx_RestartVolumeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *RestartVolumeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RestartVolumeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RestartVolumeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetEncapsulateDiskInfoOperation structure represents the GetEncapsulateDiskInfo operation
type xxx_GetEncapsulateDiskInfoOperation struct {
	This                *dcom.ORPCThis     `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat     `idl:"name:That" json:"that"`
	DiskCount           uint32             `idl:"name:diskCount" json:"disk_count"`
	DiskSpecList        []*dmrp.DiskSpec   `idl:"name:diskSpecList;size_is:(diskCount)" json:"disk_spec_list"`
	EncapInfoFlags      uint32             `idl:"name:encapInfoFlags" json:"encap_info_flags"`
	AffectedDiskCount   uint32             `idl:"name:affectedDiskCount" json:"affected_disk_count"`
	AffectedDiskList    []*dmrp.DiskInfo   `idl:"name:affectedDiskList;size_is:(, affectedDiskCount)" json:"affected_disk_list"`
	AffectedDiskFlags   []uint32           `idl:"name:affectedDiskFlags;size_is:(, affectedDiskCount)" json:"affected_disk_flags"`
	AffectedVolumeCount uint32             `idl:"name:affectedVolumeCount" json:"affected_volume_count"`
	AffectedVolumeList  []*dmrp.VolumeInfo `idl:"name:affectedVolumeList;size_is:(, affectedVolumeCount)" json:"affected_volume_list"`
	AffectedRegionCount uint32             `idl:"name:affectedRegionCount" json:"affected_region_count"`
	AffectedRegionList  []*dmrp.RegionInfo `idl:"name:affectedRegionList;size_is:(, affectedRegionCount)" json:"affected_region_list"`
	TaskInfo            *dmrp.TaskInfo     `idl:"name:tinfo" json:"task_info"`
	Return              int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_GetEncapsulateDiskInfoOperation) OpNum() int { return 53 }

func (o *xxx_GetEncapsulateDiskInfoOperation) OpName() string {
	return "/IVolumeClient/v0/GetEncapsulateDiskInfo"
}

func (o *xxx_GetEncapsulateDiskInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.DiskSpecList != nil && o.DiskCount == 0 {
		o.DiskCount = uint32(len(o.DiskSpecList))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEncapsulateDiskInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// diskCount {in} (1:(uint32))
	{
		if err := w.WriteData(o.DiskCount); err != nil {
			return err
		}
	}
	// diskSpecList {in} (1:{pointer=ref}*(1))(2:{alias=DISK_SPEC}[dim:0,size_is=diskCount](struct))
	{
		dimSize1 := uint64(o.DiskCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.DiskSpecList {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.DiskSpecList[i1] != nil {
				if err := o.DiskSpecList[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dmrp.DiskSpec{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.DiskSpecList); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&dmrp.DiskSpec{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetEncapsulateDiskInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// diskCount {in} (1:(uint32))
	{
		if err := w.ReadData(&o.DiskCount); err != nil {
			return err
		}
	}
	// diskSpecList {in} (1:{pointer=ref}*(1))(2:{alias=DISK_SPEC}[dim:0,size_is=diskCount](struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.DiskSpecList", sizeInfo[0])
		}
		o.DiskSpecList = make([]*dmrp.DiskSpec, sizeInfo[0])
		for i1 := range o.DiskSpecList {
			i1 := i1
			if o.DiskSpecList[i1] == nil {
				o.DiskSpecList[i1] = &dmrp.DiskSpec{}
			}
			if err := o.DiskSpecList[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetEncapsulateDiskInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.AffectedDiskList != nil && o.AffectedDiskCount == 0 {
		o.AffectedDiskCount = uint32(len(o.AffectedDiskList))
	}
	if o.AffectedDiskFlags != nil && o.AffectedDiskCount == 0 {
		o.AffectedDiskCount = uint32(len(o.AffectedDiskFlags))
	}
	if o.AffectedVolumeList != nil && o.AffectedVolumeCount == 0 {
		o.AffectedVolumeCount = uint32(len(o.AffectedVolumeList))
	}
	if o.AffectedRegionList != nil && o.AffectedRegionCount == 0 {
		o.AffectedRegionCount = uint32(len(o.AffectedRegionList))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEncapsulateDiskInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// encapInfoFlags {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.EncapInfoFlags); err != nil {
			return err
		}
	}
	// affectedDiskCount {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.AffectedDiskCount); err != nil {
			return err
		}
	}
	// affectedDiskList {out} (1:{pointer=ref}*(2)*(1))(2:{alias=DISK_INFO}[dim:0,size_is=affectedDiskCount](struct))
	{
		if o.AffectedDiskList != nil || o.AffectedDiskCount > 0 {
			_ptr_affectedDiskList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.AffectedDiskCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.AffectedDiskList {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.AffectedDiskList[i1] != nil {
						if err := o.AffectedDiskList[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dmrp.DiskInfo{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.AffectedDiskList); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&dmrp.DiskInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.AffectedDiskList, _ptr_affectedDiskList); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// affectedDiskFlags {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=affectedDiskCount](uint32))
	{
		if o.AffectedDiskFlags != nil || o.AffectedDiskCount > 0 {
			_ptr_affectedDiskFlags := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.AffectedDiskCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.AffectedDiskFlags {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.AffectedDiskFlags[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.AffectedDiskFlags); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint32(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.AffectedDiskFlags, _ptr_affectedDiskFlags); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// affectedVolumeCount {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.AffectedVolumeCount); err != nil {
			return err
		}
	}
	// affectedVolumeList {out} (1:{pointer=ref}*(2)*(1))(2:{alias=VOLUME_INFO}[dim:0,size_is=affectedVolumeCount](struct))
	{
		if o.AffectedVolumeList != nil || o.AffectedVolumeCount > 0 {
			_ptr_affectedVolumeList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.AffectedVolumeCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.AffectedVolumeList {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.AffectedVolumeList[i1] != nil {
						if err := o.AffectedVolumeList[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dmrp.VolumeInfo{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.AffectedVolumeList); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&dmrp.VolumeInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.AffectedVolumeList, _ptr_affectedVolumeList); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// affectedRegionCount {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.AffectedRegionCount); err != nil {
			return err
		}
	}
	// affectedRegionList {out} (1:{pointer=ref}*(2)*(1))(2:{alias=REGION_INFO}[dim:0,size_is=affectedRegionCount](struct))
	{
		if o.AffectedRegionList != nil || o.AffectedRegionCount > 0 {
			_ptr_affectedRegionList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.AffectedRegionCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.AffectedRegionList {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.AffectedRegionList[i1] != nil {
						if err := o.AffectedRegionList[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dmrp.RegionInfo{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.AffectedRegionList); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&dmrp.RegionInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.AffectedRegionList, _ptr_affectedRegionList); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEncapsulateDiskInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// encapInfoFlags {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.EncapInfoFlags); err != nil {
			return err
		}
	}
	// affectedDiskCount {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.AffectedDiskCount); err != nil {
			return err
		}
	}
	// affectedDiskList {out} (1:{pointer=ref}*(2)*(1))(2:{alias=DISK_INFO}[dim:0,size_is=affectedDiskCount](struct))
	{
		_ptr_affectedDiskList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.AffectedDiskList", sizeInfo[0])
			}
			o.AffectedDiskList = make([]*dmrp.DiskInfo, sizeInfo[0])
			for i1 := range o.AffectedDiskList {
				i1 := i1
				if o.AffectedDiskList[i1] == nil {
					o.AffectedDiskList[i1] = &dmrp.DiskInfo{}
				}
				if err := o.AffectedDiskList[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_affectedDiskList := func(ptr interface{}) { o.AffectedDiskList = *ptr.(*[]*dmrp.DiskInfo) }
		if err := w.ReadPointer(&o.AffectedDiskList, _s_affectedDiskList, _ptr_affectedDiskList); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// affectedDiskFlags {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=affectedDiskCount](uint32))
	{
		_ptr_affectedDiskFlags := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.AffectedDiskFlags", sizeInfo[0])
			}
			o.AffectedDiskFlags = make([]uint32, sizeInfo[0])
			for i1 := range o.AffectedDiskFlags {
				i1 := i1
				if err := w.ReadData(&o.AffectedDiskFlags[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_affectedDiskFlags := func(ptr interface{}) { o.AffectedDiskFlags = *ptr.(*[]uint32) }
		if err := w.ReadPointer(&o.AffectedDiskFlags, _s_affectedDiskFlags, _ptr_affectedDiskFlags); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// affectedVolumeCount {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.AffectedVolumeCount); err != nil {
			return err
		}
	}
	// affectedVolumeList {out} (1:{pointer=ref}*(2)*(1))(2:{alias=VOLUME_INFO}[dim:0,size_is=affectedVolumeCount](struct))
	{
		_ptr_affectedVolumeList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.AffectedVolumeList", sizeInfo[0])
			}
			o.AffectedVolumeList = make([]*dmrp.VolumeInfo, sizeInfo[0])
			for i1 := range o.AffectedVolumeList {
				i1 := i1
				if o.AffectedVolumeList[i1] == nil {
					o.AffectedVolumeList[i1] = &dmrp.VolumeInfo{}
				}
				if err := o.AffectedVolumeList[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_affectedVolumeList := func(ptr interface{}) { o.AffectedVolumeList = *ptr.(*[]*dmrp.VolumeInfo) }
		if err := w.ReadPointer(&o.AffectedVolumeList, _s_affectedVolumeList, _ptr_affectedVolumeList); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// affectedRegionCount {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.AffectedRegionCount); err != nil {
			return err
		}
	}
	// affectedRegionList {out} (1:{pointer=ref}*(2)*(1))(2:{alias=REGION_INFO}[dim:0,size_is=affectedRegionCount](struct))
	{
		_ptr_affectedRegionList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.AffectedRegionList", sizeInfo[0])
			}
			o.AffectedRegionList = make([]*dmrp.RegionInfo, sizeInfo[0])
			for i1 := range o.AffectedRegionList {
				i1 := i1
				if o.AffectedRegionList[i1] == nil {
					o.AffectedRegionList[i1] = &dmrp.RegionInfo{}
				}
				if err := o.AffectedRegionList[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_affectedRegionList := func(ptr interface{}) { o.AffectedRegionList = *ptr.(*[]*dmrp.RegionInfo) }
		if err := w.ReadPointer(&o.AffectedRegionList, _s_affectedRegionList, _ptr_affectedRegionList); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetEncapsulateDiskInfoRequest structure represents the GetEncapsulateDiskInfo operation request
type GetEncapsulateDiskInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// diskCount: Number of elements passed in the diskSpecList array.
	DiskCount uint32 `idl:"name:diskCount" json:"disk_count"`
	// diskSpecList: Array of DISK_SPEC structures that specifies the disks to be encapsulated.
	DiskSpecList []*dmrp.DiskSpec `idl:"name:diskSpecList;size_is:(diskCount)" json:"disk_spec_list"`
}

func (o *GetEncapsulateDiskInfoRequest) xxx_ToOp(ctx context.Context) *xxx_GetEncapsulateDiskInfoOperation {
	if o == nil {
		return &xxx_GetEncapsulateDiskInfoOperation{}
	}
	return &xxx_GetEncapsulateDiskInfoOperation{
		This:         o.This,
		DiskCount:    o.DiskCount,
		DiskSpecList: o.DiskSpecList,
	}
}

func (o *GetEncapsulateDiskInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetEncapsulateDiskInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskCount = op.DiskCount
	o.DiskSpecList = op.DiskSpecList
}
func (o *GetEncapsulateDiskInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetEncapsulateDiskInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEncapsulateDiskInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetEncapsulateDiskInfoResponse structure represents the GetEncapsulateDiskInfo operation response
type GetEncapsulateDiskInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// encapInfoFlags: Bitmap of flags that returns information about encapsulating the
	// disks specified in diskSpecList. The value of this field is generated by combining
	// zero or more of the applicable flags defined as follows with a logical OR operation.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                                       |                                                                                  |
	//	|                 VALUE                 |                                     MEANING                                      |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ENCAP_INFO_CANT_PROCEED 0x00000001    | Encapsulation for disk will not succeed. The other flags specify the reason.     |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ENCAP_INFO_NO_FREE_SPACE 0x00000002   | Volume manager could not find sufficient free space on the disk for              |
	//	|                                       | encapsulation.                                                                   |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ENCAP_INFO_BAD_ACTIVE 0x00000004      | Disk contains an active partition from which the current operating system was    |
	//	|                                       | started.                                                                         |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ENCAP_INFO_UNKNOWN_PART 0x00000008    | Volume manager was unable to determine the type of a partition on the disk.      |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ENCAP_INFO_FT_UNHEALTHY 0x00000010    | Disk contains an FT set volume that is not functioning properly.                 |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ENCAP_INFO_FT_QUERY_FAILED 0x00000020 | Volume manager was unable to obtain information about an FT set volume on the    |
	//	|                                       | disk.                                                                            |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ENCAP_INFO_FT_HAS_RAID5 0x00000040    | Disk is part of an FT RAID-5 set, which this interface does not support for      |
	//	|                                       | encapsulation.                                                                   |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ENCAP_INFO_FT_ON_BOOT 0x00000080      | Disk is both part of an FT set volume and bootable, which this interface does    |
	//	|                                       | not support for encapsulation.                                                   |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ENCAP_INFO_REBOOT_REQD 0x00000100     | Encapsulation of the disk requires a restart of the computer.                    |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ENCAP_INFO_CONTAINS_FT 0x00000200     | Disk is part of an FT set volume.                                                |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ENCAP_INFO_VOLUME_BUSY 0x00000400     | Disk is currently in use.                                                        |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ENCAP_INFO_PART_NR_CHANGE 0x00000800  | Encapsulation of the disk requires modification of the boot configuration.       |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	EncapInfoFlags uint32 `idl:"name:encapInfoFlags" json:"encap_info_flags"`
	// affectedDiskCount: Pointer to the number of disks that will be affected by the encapsulation.
	AffectedDiskCount uint32 `idl:"name:affectedDiskCount" json:"affected_disk_count"`
	// affectedDiskList: Pointer to an array of new DISK_INFO structures that represents
	// the disks that will be affected by the encapsulation. Memory for the array is allocated
	// by the server and freed by the client.
	AffectedDiskList []*dmrp.DiskInfo `idl:"name:affectedDiskList;size_is:(, affectedDiskCount)" json:"affected_disk_list"`
	// affectedDiskFlags: Pointer to an array of bitmaps of flags that provides information
	// about the disks that will be affected by the encapsulation. Memory for the array
	// is allocated by the server and freed by the client. The value of this field is a
	// logical 'OR' of 0 or more of the following applicable flags.
	//
	//	+----------------------------------------+---------------------------------------------------+
	//	|                                        |                                                   |
	//	|                 VALUE                  |                      MEANING                      |
	//	|                                        |                                                   |
	//	+----------------------------------------+---------------------------------------------------+
	//	+----------------------------------------+---------------------------------------------------+
	//	| CONTAINS_FT 0x00000001                 | Disk contains an FT set volume.                   |
	//	+----------------------------------------+---------------------------------------------------+
	//	| CONTAINS_RAID5 0x00000002              | Disk contains part of an FT RAID-5 set.           |
	//	+----------------------------------------+---------------------------------------------------+
	//	| CONTAINS_REDISTRIBUTION 0x00000004     | Disk contains an unknown volume type.             |
	//	+----------------------------------------+---------------------------------------------------+
	//	| CONTAINS_BOOTABLE_PARTITION 0x00000008 | Disk contains a bootable partition.               |
	//	+----------------------------------------+---------------------------------------------------+
	//	| CONTAINS_LOCKED_PARTITION 0x00000010   | Disk contains a locked partition.                 |
	//	+----------------------------------------+---------------------------------------------------+
	//	| CONTAINS_NO_FREE_SPACE 0x00000020      | Disk is full.                                     |
	//	+----------------------------------------+---------------------------------------------------+
	//	| CONTAINS_EXTENDED_PARTITION 0x00000040 | Disk contains an empty partition.                 |
	//	+----------------------------------------+---------------------------------------------------+
	//	| PARTITION_NUMBER_CHANGE 0x00000080     | A partition number on the disk has changed.       |
	//	+----------------------------------------+---------------------------------------------------+
	//	| CONTAINS_BOOTINDICATOR 0x00000100      | Disk contains the active partition.               |
	//	+----------------------------------------+---------------------------------------------------+
	//	| CONTAINS_BOOTLOADER 0x00000200         | Disk contains the boot loader.                    |
	//	+----------------------------------------+---------------------------------------------------+
	//	| CONTAINS_SYSTEMDIR 0x00000400          | Partition contains the system directory.          |
	//	+----------------------------------------+---------------------------------------------------+
	//	| CONTAINS_MIXED_PARTITIONS 0x00000800   | Partition contains different types of partitions. |
	//	+----------------------------------------+---------------------------------------------------+
	AffectedDiskFlags []uint32 `idl:"name:affectedDiskFlags;size_is:(, affectedDiskCount)" json:"affected_disk_flags"`
	// affectedVolumeCount: Pointer to the number of volumes that will be affected by the
	// encapsulation.
	AffectedVolumeCount uint32 `idl:"name:affectedVolumeCount" json:"affected_volume_count"`
	// affectedVolumeList: Pointer to an array of VOLUME_INFO structures that represents
	// the volumes that will be affected by the encapsulation. Memory for the array is allocated
	// by the server and freed by the client.
	AffectedVolumeList []*dmrp.VolumeInfo `idl:"name:affectedVolumeList;size_is:(, affectedVolumeCount)" json:"affected_volume_list"`
	// affectedRegionCount: Pointer to the number of regions that will be affected by the
	// encapsulation.
	AffectedRegionCount uint32 `idl:"name:affectedRegionCount" json:"affected_region_count"`
	// affectedRegionList: Pointer to an array of REGION_INFO structures that represents
	// the regions that will be affected by the encapsulation. Memory for the array is allocated
	// by the server and freed by the client.
	AffectedRegionList []*dmrp.RegionInfo `idl:"name:affectedRegionList;size_is:(, affectedRegionCount)" json:"affected_region_list"`
	// tinfo: Pointer to a TASK_INFO structure that the client can use to track the request's
	// progress.
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The GetEncapsulateDiskInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetEncapsulateDiskInfoResponse) xxx_ToOp(ctx context.Context) *xxx_GetEncapsulateDiskInfoOperation {
	if o == nil {
		return &xxx_GetEncapsulateDiskInfoOperation{}
	}
	return &xxx_GetEncapsulateDiskInfoOperation{
		That:                o.That,
		EncapInfoFlags:      o.EncapInfoFlags,
		AffectedDiskCount:   o.AffectedDiskCount,
		AffectedDiskList:    o.AffectedDiskList,
		AffectedDiskFlags:   o.AffectedDiskFlags,
		AffectedVolumeCount: o.AffectedVolumeCount,
		AffectedVolumeList:  o.AffectedVolumeList,
		AffectedRegionCount: o.AffectedRegionCount,
		AffectedRegionList:  o.AffectedRegionList,
		TaskInfo:            o.TaskInfo,
		Return:              o.Return,
	}
}

func (o *GetEncapsulateDiskInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetEncapsulateDiskInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.EncapInfoFlags = op.EncapInfoFlags
	o.AffectedDiskCount = op.AffectedDiskCount
	o.AffectedDiskList = op.AffectedDiskList
	o.AffectedDiskFlags = op.AffectedDiskFlags
	o.AffectedVolumeCount = op.AffectedVolumeCount
	o.AffectedVolumeList = op.AffectedVolumeList
	o.AffectedRegionCount = op.AffectedRegionCount
	o.AffectedRegionList = op.AffectedRegionList
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *GetEncapsulateDiskInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetEncapsulateDiskInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEncapsulateDiskInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EncapsulateDiskOperation structure represents the EncapsulateDisk operation
type xxx_EncapsulateDiskOperation struct {
	This                *dcom.ORPCThis     `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat     `idl:"name:That" json:"that"`
	AffectedDiskCount   uint32             `idl:"name:affectedDiskCount" json:"affected_disk_count"`
	AffectedDiskList    []*dmrp.DiskInfo   `idl:"name:affectedDiskList;size_is:(affectedDiskCount)" json:"affected_disk_list"`
	AffectedVolumeCount uint32             `idl:"name:affectedVolumeCount" json:"affected_volume_count"`
	AffectedVolumeList  []*dmrp.VolumeInfo `idl:"name:affectedVolumeList;size_is:(affectedVolumeCount)" json:"affected_volume_list"`
	AffectedRegionCount uint32             `idl:"name:affectedRegionCount" json:"affected_region_count"`
	AffectedRegionList  []*dmrp.RegionInfo `idl:"name:affectedRegionList;size_is:(affectedRegionCount)" json:"affected_region_list"`
	EncapInfoFlags      uint32             `idl:"name:encapInfoFlags" json:"encap_info_flags"`
	TaskInfo            *dmrp.TaskInfo     `idl:"name:tinfo" json:"task_info"`
	Return              int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_EncapsulateDiskOperation) OpNum() int { return 54 }

func (o *xxx_EncapsulateDiskOperation) OpName() string { return "/IVolumeClient/v0/EncapsulateDisk" }

func (o *xxx_EncapsulateDiskOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.AffectedDiskList != nil && o.AffectedDiskCount == 0 {
		o.AffectedDiskCount = uint32(len(o.AffectedDiskList))
	}
	if o.AffectedVolumeList != nil && o.AffectedVolumeCount == 0 {
		o.AffectedVolumeCount = uint32(len(o.AffectedVolumeList))
	}
	if o.AffectedRegionList != nil && o.AffectedRegionCount == 0 {
		o.AffectedRegionCount = uint32(len(o.AffectedRegionList))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EncapsulateDiskOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// affectedDiskCount {in} (1:(uint32))
	{
		if err := w.WriteData(o.AffectedDiskCount); err != nil {
			return err
		}
	}
	// affectedDiskList {in} (1:{pointer=ref}*(1))(2:{alias=DISK_INFO}[dim:0,size_is=affectedDiskCount](struct))
	{
		dimSize1 := uint64(o.AffectedDiskCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.AffectedDiskList {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.AffectedDiskList[i1] != nil {
				if err := o.AffectedDiskList[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dmrp.DiskInfo{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.AffectedDiskList); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&dmrp.DiskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// affectedVolumeCount {in} (1:(uint32))
	{
		if err := w.WriteData(o.AffectedVolumeCount); err != nil {
			return err
		}
	}
	// affectedVolumeList {in} (1:{pointer=ref}*(1))(2:{alias=VOLUME_INFO}[dim:0,size_is=affectedVolumeCount](struct))
	{
		dimSize1 := uint64(o.AffectedVolumeCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.AffectedVolumeList {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.AffectedVolumeList[i1] != nil {
				if err := o.AffectedVolumeList[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dmrp.VolumeInfo{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.AffectedVolumeList); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&dmrp.VolumeInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// affectedRegionCount {in} (1:(uint32))
	{
		if err := w.WriteData(o.AffectedRegionCount); err != nil {
			return err
		}
	}
	// affectedRegionList {in} (1:{pointer=ref}*(1))(2:{alias=REGION_INFO}[dim:0,size_is=affectedRegionCount](struct))
	{
		dimSize1 := uint64(o.AffectedRegionCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.AffectedRegionList {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.AffectedRegionList[i1] != nil {
				if err := o.AffectedRegionList[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dmrp.RegionInfo{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.AffectedRegionList); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&dmrp.RegionInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_EncapsulateDiskOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// affectedDiskCount {in} (1:(uint32))
	{
		if err := w.ReadData(&o.AffectedDiskCount); err != nil {
			return err
		}
	}
	// affectedDiskList {in} (1:{pointer=ref}*(1))(2:{alias=DISK_INFO}[dim:0,size_is=affectedDiskCount](struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.AffectedDiskList", sizeInfo[0])
		}
		o.AffectedDiskList = make([]*dmrp.DiskInfo, sizeInfo[0])
		for i1 := range o.AffectedDiskList {
			i1 := i1
			if o.AffectedDiskList[i1] == nil {
				o.AffectedDiskList[i1] = &dmrp.DiskInfo{}
			}
			if err := o.AffectedDiskList[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// affectedVolumeCount {in} (1:(uint32))
	{
		if err := w.ReadData(&o.AffectedVolumeCount); err != nil {
			return err
		}
	}
	// affectedVolumeList {in} (1:{pointer=ref}*(1))(2:{alias=VOLUME_INFO}[dim:0,size_is=affectedVolumeCount](struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.AffectedVolumeList", sizeInfo[0])
		}
		o.AffectedVolumeList = make([]*dmrp.VolumeInfo, sizeInfo[0])
		for i1 := range o.AffectedVolumeList {
			i1 := i1
			if o.AffectedVolumeList[i1] == nil {
				o.AffectedVolumeList[i1] = &dmrp.VolumeInfo{}
			}
			if err := o.AffectedVolumeList[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// affectedRegionCount {in} (1:(uint32))
	{
		if err := w.ReadData(&o.AffectedRegionCount); err != nil {
			return err
		}
	}
	// affectedRegionList {in} (1:{pointer=ref}*(1))(2:{alias=REGION_INFO}[dim:0,size_is=affectedRegionCount](struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.AffectedRegionList", sizeInfo[0])
		}
		o.AffectedRegionList = make([]*dmrp.RegionInfo, sizeInfo[0])
		for i1 := range o.AffectedRegionList {
			i1 := i1
			if o.AffectedRegionList[i1] == nil {
				o.AffectedRegionList[i1] = &dmrp.RegionInfo{}
			}
			if err := o.AffectedRegionList[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_EncapsulateDiskOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EncapsulateDiskOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// encapInfoFlags {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.EncapInfoFlags); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EncapsulateDiskOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// encapInfoFlags {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.EncapInfoFlags); err != nil {
			return err
		}
	}
	// tinfo {out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EncapsulateDiskRequest structure represents the EncapsulateDisk operation request
type EncapsulateDiskRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// affectedDiskCount: The number of elements passed in the affectedDiskList array.
	AffectedDiskCount uint32 `idl:"name:affectedDiskCount" json:"affected_disk_count"`
	// affectedDiskList: An array of DISK_INFO structures that specifies the disks to be
	// encapsulated.
	AffectedDiskList []*dmrp.DiskInfo `idl:"name:affectedDiskList;size_is:(affectedDiskCount)" json:"affected_disk_list"`
	// affectedVolumeCount: The number of elements passed in the affectedVolumeList array.
	AffectedVolumeCount uint32 `idl:"name:affectedVolumeCount" json:"affected_volume_count"`
	// affectedVolumeList: An array of VOLUME_INFO structures that represents the volumes
	// affected by the encapsulation. If the number of affect volumes is zero, a pointer
	// to a zero length array MUST be passed. This pointer MUST NOT be input as NULL.
	AffectedVolumeList []*dmrp.VolumeInfo `idl:"name:affectedVolumeList;size_is:(affectedVolumeCount)" json:"affected_volume_list"`
	// affectedRegionCount: The number of elements passed in the affectedRegionList array.
	AffectedRegionCount uint32 `idl:"name:affectedRegionCount" json:"affected_region_count"`
	// affectedRegionList: An array of REGION_INFO structures that represents the regions
	// affected by the encapsulation. If the number of affect regions is zero, a pointer
	// to a zero length array MUST be passed. This pointer MUST NOT be input as NULL.
	AffectedRegionList []*dmrp.RegionInfo `idl:"name:affectedRegionList;size_is:(affectedRegionCount)" json:"affected_region_list"`
}

func (o *EncapsulateDiskRequest) xxx_ToOp(ctx context.Context) *xxx_EncapsulateDiskOperation {
	if o == nil {
		return &xxx_EncapsulateDiskOperation{}
	}
	return &xxx_EncapsulateDiskOperation{
		This:                o.This,
		AffectedDiskCount:   o.AffectedDiskCount,
		AffectedDiskList:    o.AffectedDiskList,
		AffectedVolumeCount: o.AffectedVolumeCount,
		AffectedVolumeList:  o.AffectedVolumeList,
		AffectedRegionCount: o.AffectedRegionCount,
		AffectedRegionList:  o.AffectedRegionList,
	}
}

func (o *EncapsulateDiskRequest) xxx_FromOp(ctx context.Context, op *xxx_EncapsulateDiskOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.AffectedDiskCount = op.AffectedDiskCount
	o.AffectedDiskList = op.AffectedDiskList
	o.AffectedVolumeCount = op.AffectedVolumeCount
	o.AffectedVolumeList = op.AffectedVolumeList
	o.AffectedRegionCount = op.AffectedRegionCount
	o.AffectedRegionList = op.AffectedRegionList
}
func (o *EncapsulateDiskRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EncapsulateDiskRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EncapsulateDiskOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EncapsulateDiskResponse structure represents the EncapsulateDisk operation response
type EncapsulateDiskResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// encapInfoFlags: Bitmap of flags that provide information about the encapsulation.
	// The value of this field is a logical 'OR' of zero or more of the following applicable
	// flags.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                                       |                                                                                  |
	//	|                 VALUE                 |                                     MEANING                                      |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ENCAP_INFO_CANT_PROCEED 0x00000001    | Encapsulation for disk did not succeed. The other flags specify the reason.      |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ENCAP_INFO_NO_FREE_SPACE 0x00000002   | The volume manager could not find sufficient free space on the disk for          |
	//	|                                       | encapsulation.                                                                   |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ENCAP_INFO_BAD_ACTIVE 0x00000004      | The disk contains an active partition from which the current operating system    |
	//	|                                       | was not started.                                                                 |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ENCAP_INFO_UNKNOWN_PART 0x00000008    | The volume manager was unable to determine the type of a partition on the disk.  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ENCAP_INFO_FT_UNHEALTHY 0x00000010    | The disk contains an unhealthy FT set volume.                                    |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ENCAP_INFO_FT_QUERY_FAILED 0x00000020 | The volume manager was unable to obtain information about an FT set volume on    |
	//	|                                       | the disk.                                                                        |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ENCAP_INFO_FT_HAS_RAID5 0x00000040    | The disk is part of an FT RAID-5 set, which this interface does not support for  |
	//	|                                       | encapsulation.                                                                   |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ENCAP_INFO_FT_ON_BOOT 0x00000080      | The disk is part of an FT set volume and bootable, which this interface does not |
	//	|                                       | support for encapsulation.                                                       |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ENCAP_INFO_REBOOT_REQD 0x00000100     | Encapsulation of the disk requires a restart of the computer.                    |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ENCAP_INFO_CONTAINS_FT 0x00000200     | The disk is part of an FT set volume.                                            |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ENCAP_INFO_VOLUME_BUSY 0x00000400     | The disk is currently in use.                                                    |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ENCAP_INFO_PART_NR_CHANGE 0x00000800  | Encapsulation of the disk requires modification of the boot configuration.       |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	EncapInfoFlags uint32 `idl:"name:encapInfoFlags" json:"encap_info_flags"`
	// tinfo: Pointer to a TASK_INFO structure that the client can use to track the request's
	// progress.
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The EncapsulateDisk return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EncapsulateDiskResponse) xxx_ToOp(ctx context.Context) *xxx_EncapsulateDiskOperation {
	if o == nil {
		return &xxx_EncapsulateDiskOperation{}
	}
	return &xxx_EncapsulateDiskOperation{
		That:           o.That,
		EncapInfoFlags: o.EncapInfoFlags,
		TaskInfo:       o.TaskInfo,
		Return:         o.Return,
	}
}

func (o *EncapsulateDiskResponse) xxx_FromOp(ctx context.Context, op *xxx_EncapsulateDiskOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.EncapInfoFlags = op.EncapInfoFlags
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *EncapsulateDiskResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EncapsulateDiskResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EncapsulateDiskOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryChangePartitionNumbersOperation structure represents the QueryChangePartitionNumbers operation
type xxx_QueryChangePartitionNumbersOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	OldPartitionNumber int32          `idl:"name:oldPartitionNumber" json:"old_partition_number"`
	NewPartitionNumber int32          `idl:"name:newPartitionNumber" json:"new_partition_number"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryChangePartitionNumbersOperation) OpNum() int { return 55 }

func (o *xxx_QueryChangePartitionNumbersOperation) OpName() string {
	return "/IVolumeClient/v0/QueryChangePartitionNumbers"
}

func (o *xxx_QueryChangePartitionNumbersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryChangePartitionNumbersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryChangePartitionNumbersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryChangePartitionNumbersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryChangePartitionNumbersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// oldPartitionNumber {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.OldPartitionNumber); err != nil {
			return err
		}
	}
	// newPartitionNumber {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.NewPartitionNumber); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryChangePartitionNumbersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// oldPartitionNumber {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.OldPartitionNumber); err != nil {
			return err
		}
	}
	// newPartitionNumber {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.NewPartitionNumber); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryChangePartitionNumbersRequest structure represents the QueryChangePartitionNumbers operation request
type QueryChangePartitionNumbersRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *QueryChangePartitionNumbersRequest) xxx_ToOp(ctx context.Context) *xxx_QueryChangePartitionNumbersOperation {
	if o == nil {
		return &xxx_QueryChangePartitionNumbersOperation{}
	}
	return &xxx_QueryChangePartitionNumbersOperation{
		This: o.This,
	}
}

func (o *QueryChangePartitionNumbersRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryChangePartitionNumbersOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *QueryChangePartitionNumbersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryChangePartitionNumbersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryChangePartitionNumbersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryChangePartitionNumbersResponse structure represents the QueryChangePartitionNumbers operation response
type QueryChangePartitionNumbersResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	OldPartitionNumber int32          `idl:"name:oldPartitionNumber" json:"old_partition_number"`
	NewPartitionNumber int32          `idl:"name:newPartitionNumber" json:"new_partition_number"`
	// Return: The QueryChangePartitionNumbers return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryChangePartitionNumbersResponse) xxx_ToOp(ctx context.Context) *xxx_QueryChangePartitionNumbersOperation {
	if o == nil {
		return &xxx_QueryChangePartitionNumbersOperation{}
	}
	return &xxx_QueryChangePartitionNumbersOperation{
		That:               o.That,
		OldPartitionNumber: o.OldPartitionNumber,
		NewPartitionNumber: o.NewPartitionNumber,
		Return:             o.Return,
	}
}

func (o *QueryChangePartitionNumbersResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryChangePartitionNumbersOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.OldPartitionNumber = op.OldPartitionNumber
	o.NewPartitionNumber = op.NewPartitionNumber
	o.Return = op.Return
}
func (o *QueryChangePartitionNumbersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryChangePartitionNumbersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryChangePartitionNumbersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeletePartitionNumberInfoFromRegistryOperation structure represents the DeletePartitionNumberInfoFromRegistry operation
type xxx_DeletePartitionNumberInfoFromRegistryOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeletePartitionNumberInfoFromRegistryOperation) OpNum() int { return 56 }

func (o *xxx_DeletePartitionNumberInfoFromRegistryOperation) OpName() string {
	return "/IVolumeClient/v0/DeletePartitionNumberInfoFromRegistry"
}

func (o *xxx_DeletePartitionNumberInfoFromRegistryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeletePartitionNumberInfoFromRegistryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeletePartitionNumberInfoFromRegistryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeletePartitionNumberInfoFromRegistryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeletePartitionNumberInfoFromRegistryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeletePartitionNumberInfoFromRegistryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DeletePartitionNumberInfoFromRegistryRequest structure represents the DeletePartitionNumberInfoFromRegistry operation request
type DeletePartitionNumberInfoFromRegistryRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *DeletePartitionNumberInfoFromRegistryRequest) xxx_ToOp(ctx context.Context) *xxx_DeletePartitionNumberInfoFromRegistryOperation {
	if o == nil {
		return &xxx_DeletePartitionNumberInfoFromRegistryOperation{}
	}
	return &xxx_DeletePartitionNumberInfoFromRegistryOperation{
		This: o.This,
	}
}

func (o *DeletePartitionNumberInfoFromRegistryRequest) xxx_FromOp(ctx context.Context, op *xxx_DeletePartitionNumberInfoFromRegistryOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *DeletePartitionNumberInfoFromRegistryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeletePartitionNumberInfoFromRegistryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeletePartitionNumberInfoFromRegistryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeletePartitionNumberInfoFromRegistryResponse structure represents the DeletePartitionNumberInfoFromRegistry operation response
type DeletePartitionNumberInfoFromRegistryResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeletePartitionNumberInfoFromRegistry return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeletePartitionNumberInfoFromRegistryResponse) xxx_ToOp(ctx context.Context) *xxx_DeletePartitionNumberInfoFromRegistryOperation {
	if o == nil {
		return &xxx_DeletePartitionNumberInfoFromRegistryOperation{}
	}
	return &xxx_DeletePartitionNumberInfoFromRegistryOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *DeletePartitionNumberInfoFromRegistryResponse) xxx_FromOp(ctx context.Context, op *xxx_DeletePartitionNumberInfoFromRegistryOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeletePartitionNumberInfoFromRegistryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeletePartitionNumberInfoFromRegistryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeletePartitionNumberInfoFromRegistryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetDontShowOperation structure represents the SetDontShow operation
type xxx_SetDontShowOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	SetNoShow bool           `idl:"name:bSetNoShow" json:"set_no_show"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetDontShowOperation) OpNum() int { return 57 }

func (o *xxx_SetDontShowOperation) OpName() string { return "/IVolumeClient/v0/SetDontShow" }

func (o *xxx_SetDontShowOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDontShowOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// bSetNoShow {in} (1:(bool))
	{
		if err := w.WriteData(o.SetNoShow); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDontShowOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bSetNoShow {in} (1:(bool))
	{
		if err := w.ReadData(&o.SetNoShow); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDontShowOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDontShowOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDontShowOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetDontShowRequest structure represents the SetDontShow operation request
type SetDontShowRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	SetNoShow bool           `idl:"name:bSetNoShow" json:"set_no_show"`
}

func (o *SetDontShowRequest) xxx_ToOp(ctx context.Context) *xxx_SetDontShowOperation {
	if o == nil {
		return &xxx_SetDontShowOperation{}
	}
	return &xxx_SetDontShowOperation{
		This:      o.This,
		SetNoShow: o.SetNoShow,
	}
}

func (o *SetDontShowRequest) xxx_FromOp(ctx context.Context, op *xxx_SetDontShowOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SetNoShow = op.SetNoShow
}
func (o *SetDontShowRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetDontShowRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDontShowOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetDontShowResponse structure represents the SetDontShow operation response
type SetDontShowResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetDontShow return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetDontShowResponse) xxx_ToOp(ctx context.Context) *xxx_SetDontShowOperation {
	if o == nil {
		return &xxx_SetDontShowOperation{}
	}
	return &xxx_SetDontShowOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetDontShowResponse) xxx_FromOp(ctx context.Context, op *xxx_SetDontShowOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetDontShowResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetDontShowResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDontShowOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDontShowOperation structure represents the GetDontShow operation
type xxx_GetDontShowOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	GetNoShow bool           `idl:"name:bGetNoShow" json:"get_no_show"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDontShowOperation) OpNum() int { return 58 }

func (o *xxx_GetDontShowOperation) OpName() string { return "/IVolumeClient/v0/GetDontShow" }

func (o *xxx_GetDontShowOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDontShowOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDontShowOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDontShowOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDontShowOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// bGetNoShow {out} (1:{pointer=ref}*(1)(bool))
	{
		if err := w.WriteData(o.GetNoShow); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDontShowOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bGetNoShow {out} (1:{pointer=ref}*(1)(bool))
	{
		if err := w.ReadData(&o.GetNoShow); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetDontShowRequest structure represents the GetDontShow operation request
type GetDontShowRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetDontShowRequest) xxx_ToOp(ctx context.Context) *xxx_GetDontShowOperation {
	if o == nil {
		return &xxx_GetDontShowOperation{}
	}
	return &xxx_GetDontShowOperation{
		This: o.This,
	}
}

func (o *GetDontShowRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDontShowOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetDontShowRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetDontShowRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDontShowOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDontShowResponse structure represents the GetDontShow operation response
type GetDontShowResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	GetNoShow bool           `idl:"name:bGetNoShow" json:"get_no_show"`
	// Return: The GetDontShow return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDontShowResponse) xxx_ToOp(ctx context.Context) *xxx_GetDontShowOperation {
	if o == nil {
		return &xxx_GetDontShowOperation{}
	}
	return &xxx_GetDontShowOperation{
		That:      o.That,
		GetNoShow: o.GetNoShow,
		Return:    o.Return,
	}
}

func (o *GetDontShowResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDontShowOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.GetNoShow = op.GetNoShow
	o.Return = op.Return
}
func (o *GetDontShowResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetDontShowResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDontShowOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumTasksOperation structure represents the EnumTasks operation
type xxx_EnumTasksOperation struct {
	This      *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat   `idl:"name:That" json:"that"`
	TaskCount uint32           `idl:"name:taskCount" json:"task_count"`
	TaskList  []*dmrp.TaskInfo `idl:"name:taskList;size_is:(, taskCount)" json:"task_list"`
	Return    int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumTasksOperation) OpNum() int { return 67 }

func (o *xxx_EnumTasksOperation) OpName() string { return "/IVolumeClient/v0/EnumTasks" }

func (o *xxx_EnumTasksOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumTasksOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// taskCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.TaskCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumTasksOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// taskCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.TaskCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumTasksOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.TaskList != nil && o.TaskCount == 0 {
		o.TaskCount = uint32(len(o.TaskList))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumTasksOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// taskCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.TaskCount); err != nil {
			return err
		}
	}
	// taskList {out} (1:{pointer=ref}*(2)*(1))(2:{alias=TASK_INFO}[dim:0,size_is=taskCount](struct))
	{
		if o.TaskList != nil || o.TaskCount > 0 {
			_ptr_taskList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.TaskCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.TaskList {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.TaskList[i1] != nil {
						if err := o.TaskList[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.TaskList); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.TaskList, _ptr_taskList); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumTasksOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// taskCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.TaskCount); err != nil {
			return err
		}
	}
	// taskList {out} (1:{pointer=ref}*(2)*(1))(2:{alias=TASK_INFO}[dim:0,size_is=taskCount](struct))
	{
		_ptr_taskList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.TaskList", sizeInfo[0])
			}
			o.TaskList = make([]*dmrp.TaskInfo, sizeInfo[0])
			for i1 := range o.TaskList {
				i1 := i1
				if o.TaskList[i1] == nil {
					o.TaskList[i1] = &dmrp.TaskInfo{}
				}
				if err := o.TaskList[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_taskList := func(ptr interface{}) { o.TaskList = *ptr.(*[]*dmrp.TaskInfo) }
		if err := w.ReadPointer(&o.TaskList, _s_taskList, _ptr_taskList); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumTasksRequest structure represents the EnumTasks operation request
type EnumTasksRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	TaskCount uint32         `idl:"name:taskCount" json:"task_count"`
}

func (o *EnumTasksRequest) xxx_ToOp(ctx context.Context) *xxx_EnumTasksOperation {
	if o == nil {
		return &xxx_EnumTasksOperation{}
	}
	return &xxx_EnumTasksOperation{
		This:      o.This,
		TaskCount: o.TaskCount,
	}
}

func (o *EnumTasksRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumTasksOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.TaskCount = op.TaskCount
}
func (o *EnumTasksRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumTasksRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumTasksOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumTasksResponse structure represents the EnumTasks operation response
type EnumTasksResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That      *dcom.ORPCThat   `idl:"name:That" json:"that"`
	TaskCount uint32           `idl:"name:taskCount" json:"task_count"`
	TaskList  []*dmrp.TaskInfo `idl:"name:taskList;size_is:(, taskCount)" json:"task_list"`
	// Return: The EnumTasks return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumTasksResponse) xxx_ToOp(ctx context.Context) *xxx_EnumTasksOperation {
	if o == nil {
		return &xxx_EnumTasksOperation{}
	}
	return &xxx_EnumTasksOperation{
		That:      o.That,
		TaskCount: o.TaskCount,
		TaskList:  o.TaskList,
		Return:    o.Return,
	}
}

func (o *EnumTasksResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumTasksOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskCount = op.TaskCount
	o.TaskList = op.TaskList
	o.Return = op.Return
}
func (o *EnumTasksResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumTasksResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumTasksOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetTaskDetailOperation structure represents the GetTaskDetail operation
type xxx_GetTaskDetailOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	ID       int64          `idl:"name:id" json:"id"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTaskDetailOperation) OpNum() int { return 68 }

func (o *xxx_GetTaskDetailOperation) OpName() string { return "/IVolumeClient/v0/GetTaskDetail" }

func (o *xxx_GetTaskDetailOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTaskDetailOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// id {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.ID); err != nil {
			return err
		}
	}
	// tinfo {in, out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetTaskDetailOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// id {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.ID); err != nil {
			return err
		}
	}
	// tinfo {in, out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTaskDetailOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTaskDetailOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// tinfo {in, out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo != nil {
			if err := o.TaskInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dmrp.TaskInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTaskDetailOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tinfo {in, out} (1:{pointer=ref}*(1))(2:{alias=TASK_INFO}(struct))
	{
		if o.TaskInfo == nil {
			o.TaskInfo = &dmrp.TaskInfo{}
		}
		if err := o.TaskInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetTaskDetailRequest structure represents the GetTaskDetail operation request
type GetTaskDetailRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	ID       int64          `idl:"name:id" json:"id"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
}

func (o *GetTaskDetailRequest) xxx_ToOp(ctx context.Context) *xxx_GetTaskDetailOperation {
	if o == nil {
		return &xxx_GetTaskDetailOperation{}
	}
	return &xxx_GetTaskDetailOperation{
		This:     o.This,
		ID:       o.ID,
		TaskInfo: o.TaskInfo,
	}
}

func (o *GetTaskDetailRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTaskDetailOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ID = op.ID
	o.TaskInfo = op.TaskInfo
}
func (o *GetTaskDetailRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetTaskDetailRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTaskDetailOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetTaskDetailResponse structure represents the GetTaskDetail operation response
type GetTaskDetailResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TaskInfo *dmrp.TaskInfo `idl:"name:tinfo" json:"task_info"`
	// Return: The GetTaskDetail return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTaskDetailResponse) xxx_ToOp(ctx context.Context) *xxx_GetTaskDetailOperation {
	if o == nil {
		return &xxx_GetTaskDetailOperation{}
	}
	return &xxx_GetTaskDetailOperation{
		That:     o.That,
		TaskInfo: o.TaskInfo,
		Return:   o.Return,
	}
}

func (o *GetTaskDetailResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTaskDetailOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TaskInfo = op.TaskInfo
	o.Return = op.Return
}
func (o *GetTaskDetailResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetTaskDetailResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTaskDetailOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AbortTaskOperation structure represents the AbortTask operation
type xxx_AbortTaskOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	ID     int64          `idl:"name:id" json:"id"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_AbortTaskOperation) OpNum() int { return 69 }

func (o *xxx_AbortTaskOperation) OpName() string { return "/IVolumeClient/v0/AbortTask" }

func (o *xxx_AbortTaskOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AbortTaskOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// id {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.ID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AbortTaskOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// id {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.ID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AbortTaskOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AbortTaskOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AbortTaskOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AbortTaskRequest structure represents the AbortTask operation request
type AbortTaskRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	ID   int64          `idl:"name:id" json:"id"`
}

func (o *AbortTaskRequest) xxx_ToOp(ctx context.Context) *xxx_AbortTaskOperation {
	if o == nil {
		return &xxx_AbortTaskOperation{}
	}
	return &xxx_AbortTaskOperation{
		This: o.This,
		ID:   o.ID,
	}
}

func (o *AbortTaskRequest) xxx_FromOp(ctx context.Context, op *xxx_AbortTaskOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ID = op.ID
}
func (o *AbortTaskRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AbortTaskRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AbortTaskOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AbortTaskResponse structure represents the AbortTask operation response
type AbortTaskResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AbortTask return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AbortTaskResponse) xxx_ToOp(ctx context.Context) *xxx_AbortTaskOperation {
	if o == nil {
		return &xxx_AbortTaskOperation{}
	}
	return &xxx_AbortTaskOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *AbortTaskResponse) xxx_FromOp(ctx context.Context, op *xxx_AbortTaskOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *AbortTaskResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AbortTaskResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AbortTaskOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_HResultGetErrorDataOperation structure represents the HrGetErrorData operation
type xxx_HResultGetErrorDataOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	HResult     int32          `idl:"name:hr" json:"hresult"`
	Flags       uint32         `idl:"name:dwFlags" json:"flags"`
	StoredFlags uint32         `idl:"name:pdwStoredFlags" json:"stored_flags"`
	Size        int32          `idl:"name:pcszw" json:"size"`
	Data        []string       `idl:"name:prgszw;size_is:(, pcszw, );string" json:"data"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_HResultGetErrorDataOperation) OpNum() int { return 70 }

func (o *xxx_HResultGetErrorDataOperation) OpName() string { return "/IVolumeClient/v0/HrGetErrorData" }

func (o *xxx_HResultGetErrorDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_HResultGetErrorDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// hr {in} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.HResult); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_HResultGetErrorDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// hr {in} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.HResult); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_HResultGetErrorDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Data != nil && o.Size == 0 {
		o.Size = int32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_HResultGetErrorDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pdwStoredFlags {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.StoredFlags); err != nil {
			return err
		}
	}
	// pcszw {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.Size); err != nil {
			return err
		}
	}
	// prgszw {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,size_is=pcszw]*(1)[dim:0,string,null](wchar))
	{
		if o.Data != nil || o.Size > 0 {
			_ptr_prgszw := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.Size)
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
					if o.Data[i1] != "" {
						_ptr_prgszw := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
							if err := ndr.WriteUTF16NString(ctx, w, o.Data[i1]); err != nil {
								return err
							}
							return nil
						})
						if err := w.WritePointer(&o.Data[i1], _ptr_prgszw); err != nil {
							return err
						}
					} else {
						if err := w.WritePointer(nil); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Data, _ptr_prgszw); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_HResultGetErrorDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pdwStoredFlags {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.StoredFlags); err != nil {
			return err
		}
	}
	// pcszw {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.Size); err != nil {
			return err
		}
	}
	// prgszw {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,size_is=pcszw]*(1)[dim:0,string,null](wchar))
	{
		_ptr_prgszw := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
			}
			o.Data = make([]string, sizeInfo[0])
			for i1 := range o.Data {
				i1 := i1
				_ptr_prgszw := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
					if err := ndr.ReadUTF16NString(ctx, w, &o.Data[i1]); err != nil {
						return err
					}
					return nil
				})
				_s_prgszw := func(ptr interface{}) { o.Data[i1] = *ptr.(*string) }
				if err := w.ReadPointer(&o.Data[i1], _s_prgszw, _ptr_prgszw); err != nil {
					return err
				}
			}
			return nil
		})
		_s_prgszw := func(ptr interface{}) { o.Data = *ptr.(*[]string) }
		if err := w.ReadPointer(&o.Data, _s_prgszw, _ptr_prgszw); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// HResultGetErrorDataRequest structure represents the HrGetErrorData operation request
type HResultGetErrorDataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	HResult int32          `idl:"name:hr" json:"hresult"`
	Flags   uint32         `idl:"name:dwFlags" json:"flags"`
}

func (o *HResultGetErrorDataRequest) xxx_ToOp(ctx context.Context) *xxx_HResultGetErrorDataOperation {
	if o == nil {
		return &xxx_HResultGetErrorDataOperation{}
	}
	return &xxx_HResultGetErrorDataOperation{
		This:    o.This,
		HResult: o.HResult,
		Flags:   o.Flags,
	}
}

func (o *HResultGetErrorDataRequest) xxx_FromOp(ctx context.Context, op *xxx_HResultGetErrorDataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.HResult = op.HResult
	o.Flags = op.Flags
}
func (o *HResultGetErrorDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *HResultGetErrorDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_HResultGetErrorDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// HResultGetErrorDataResponse structure represents the HrGetErrorData operation response
type HResultGetErrorDataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	StoredFlags uint32         `idl:"name:pdwStoredFlags" json:"stored_flags"`
	Size        int32          `idl:"name:pcszw" json:"size"`
	Data        []string       `idl:"name:prgszw;size_is:(, pcszw, );string" json:"data"`
	// Return: The HrGetErrorData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *HResultGetErrorDataResponse) xxx_ToOp(ctx context.Context) *xxx_HResultGetErrorDataOperation {
	if o == nil {
		return &xxx_HResultGetErrorDataOperation{}
	}
	return &xxx_HResultGetErrorDataOperation{
		That:        o.That,
		StoredFlags: o.StoredFlags,
		Size:        o.Size,
		Data:        o.Data,
		Return:      o.Return,
	}
}

func (o *HResultGetErrorDataResponse) xxx_FromOp(ctx context.Context, op *xxx_HResultGetErrorDataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.StoredFlags = op.StoredFlags
	o.Size = op.Size
	o.Data = op.Data
	o.Return = op.Return
}
func (o *HResultGetErrorDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *HResultGetErrorDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_HResultGetErrorDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_InitializeOperation structure represents the Initialize operation
type xxx_InitializeOperation struct {
	This                  *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                  *dcom.ORPCThat `idl:"name:That" json:"that"`
	NotificationInterface *dcom.Unknown  `idl:"name:notificationInterface" json:"notification_interface"`
	IDLVersion            uint32         `idl:"name:ulIDLVersion" json:"idl_version"`
	Flags                 uint32         `idl:"name:pdwFlags" json:"flags"`
	ClientID              int64          `idl:"name:clientId" json:"client_id"`
	RemoteCount           uint32         `idl:"name:cRemote" json:"remote_count"`
	Return                int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_InitializeOperation) OpNum() int { return 71 }

func (o *xxx_InitializeOperation) OpName() string { return "/IVolumeClient/v0/Initialize" }

func (o *xxx_InitializeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitializeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// notificationInterface {in} (1:{pointer=ref}*(1))(2:{alias=IUnknown}(interface))
	{
		if o.NotificationInterface != nil {
			_ptr_notificationInterface := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.NotificationInterface != nil {
					if err := o.NotificationInterface.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcom.Unknown{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NotificationInterface, _ptr_notificationInterface); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// cRemote {in} (1:(uint32))
	{
		if err := w.WriteData(o.RemoteCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitializeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// notificationInterface {in} (1:{pointer=ref}*(1))(2:{alias=IUnknown}(interface))
	{
		_ptr_notificationInterface := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.NotificationInterface == nil {
				o.NotificationInterface = &dcom.Unknown{}
			}
			if err := o.NotificationInterface.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_notificationInterface := func(ptr interface{}) { o.NotificationInterface = *ptr.(**dcom.Unknown) }
		if err := w.ReadPointer(&o.NotificationInterface, _s_notificationInterface, _ptr_notificationInterface); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cRemote {in} (1:(uint32))
	{
		if err := w.ReadData(&o.RemoteCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitializeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitializeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ulIDLVersion {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.IDLVersion); err != nil {
			return err
		}
	}
	// pdwFlags {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// clientId {out} (1:{pointer=ref}*(1))(2:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.ClientID); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitializeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ulIDLVersion {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.IDLVersion); err != nil {
			return err
		}
	}
	// pdwFlags {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// clientId {out} (1:{pointer=ref}*(1))(2:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.ClientID); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// InitializeRequest structure represents the Initialize operation request
type InitializeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                  *dcom.ORPCThis `idl:"name:This" json:"this"`
	NotificationInterface *dcom.Unknown  `idl:"name:notificationInterface" json:"notification_interface"`
	RemoteCount           uint32         `idl:"name:cRemote" json:"remote_count"`
}

func (o *InitializeRequest) xxx_ToOp(ctx context.Context) *xxx_InitializeOperation {
	if o == nil {
		return &xxx_InitializeOperation{}
	}
	return &xxx_InitializeOperation{
		This:                  o.This,
		NotificationInterface: o.NotificationInterface,
		RemoteCount:           o.RemoteCount,
	}
}

func (o *InitializeRequest) xxx_FromOp(ctx context.Context, op *xxx_InitializeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.NotificationInterface = op.NotificationInterface
	o.RemoteCount = op.RemoteCount
}
func (o *InitializeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *InitializeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InitializeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// InitializeResponse structure represents the Initialize operation response
type InitializeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	IDLVersion uint32         `idl:"name:ulIDLVersion" json:"idl_version"`
	Flags      uint32         `idl:"name:pdwFlags" json:"flags"`
	ClientID   int64          `idl:"name:clientId" json:"client_id"`
	// Return: The Initialize return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *InitializeResponse) xxx_ToOp(ctx context.Context) *xxx_InitializeOperation {
	if o == nil {
		return &xxx_InitializeOperation{}
	}
	return &xxx_InitializeOperation{
		That:       o.That,
		IDLVersion: o.IDLVersion,
		Flags:      o.Flags,
		ClientID:   o.ClientID,
		Return:     o.Return,
	}
}

func (o *InitializeResponse) xxx_FromOp(ctx context.Context, op *xxx_InitializeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.IDLVersion = op.IDLVersion
	o.Flags = op.Flags
	o.ClientID = op.ClientID
	o.Return = op.Return
}
func (o *InitializeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *InitializeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InitializeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_UninitializeOperation structure represents the Uninitialize operation
type xxx_UninitializeOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_UninitializeOperation) OpNum() int { return 72 }

func (o *xxx_UninitializeOperation) OpName() string { return "/IVolumeClient/v0/Uninitialize" }

func (o *xxx_UninitializeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UninitializeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UninitializeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UninitializeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UninitializeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UninitializeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// UninitializeRequest structure represents the Uninitialize operation request
type UninitializeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *UninitializeRequest) xxx_ToOp(ctx context.Context) *xxx_UninitializeOperation {
	if o == nil {
		return &xxx_UninitializeOperation{}
	}
	return &xxx_UninitializeOperation{
		This: o.This,
	}
}

func (o *UninitializeRequest) xxx_FromOp(ctx context.Context, op *xxx_UninitializeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *UninitializeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *UninitializeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UninitializeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// UninitializeResponse structure represents the Uninitialize operation response
type UninitializeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Uninitialize return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *UninitializeResponse) xxx_ToOp(ctx context.Context) *xxx_UninitializeOperation {
	if o == nil {
		return &xxx_UninitializeOperation{}
	}
	return &xxx_UninitializeOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *UninitializeResponse) xxx_FromOp(ctx context.Context, op *xxx_UninitializeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *UninitializeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *UninitializeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UninitializeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RefreshOperation structure represents the Refresh operation
type xxx_RefreshOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RefreshOperation) OpNum() int { return 73 }

func (o *xxx_RefreshOperation) OpName() string { return "/IVolumeClient/v0/Refresh" }

func (o *xxx_RefreshOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RefreshOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RefreshOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RefreshOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RefreshOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RefreshOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RefreshRequest structure represents the Refresh operation request
type RefreshRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *RefreshRequest) xxx_ToOp(ctx context.Context) *xxx_RefreshOperation {
	if o == nil {
		return &xxx_RefreshOperation{}
	}
	return &xxx_RefreshOperation{
		This: o.This,
	}
}

func (o *RefreshRequest) xxx_FromOp(ctx context.Context, op *xxx_RefreshOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *RefreshRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RefreshRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RefreshOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RefreshResponse structure represents the Refresh operation response
type RefreshResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Refresh return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RefreshResponse) xxx_ToOp(ctx context.Context) *xxx_RefreshOperation {
	if o == nil {
		return &xxx_RefreshOperation{}
	}
	return &xxx_RefreshOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *RefreshResponse) xxx_FromOp(ctx context.Context, op *xxx_RefreshOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RefreshResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RefreshResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RefreshOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RescanDisksOperation structure represents the RescanDisks operation
type xxx_RescanDisksOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RescanDisksOperation) OpNum() int { return 74 }

func (o *xxx_RescanDisksOperation) OpName() string { return "/IVolumeClient/v0/RescanDisks" }

func (o *xxx_RescanDisksOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RescanDisksOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RescanDisksOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RescanDisksOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RescanDisksOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RescanDisksOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RescanDisksRequest structure represents the RescanDisks operation request
type RescanDisksRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *RescanDisksRequest) xxx_ToOp(ctx context.Context) *xxx_RescanDisksOperation {
	if o == nil {
		return &xxx_RescanDisksOperation{}
	}
	return &xxx_RescanDisksOperation{
		This: o.This,
	}
}

func (o *RescanDisksRequest) xxx_FromOp(ctx context.Context, op *xxx_RescanDisksOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *RescanDisksRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RescanDisksRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RescanDisksOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RescanDisksResponse structure represents the RescanDisks operation response
type RescanDisksResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RescanDisks return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RescanDisksResponse) xxx_ToOp(ctx context.Context) *xxx_RescanDisksOperation {
	if o == nil {
		return &xxx_RescanDisksOperation{}
	}
	return &xxx_RescanDisksOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *RescanDisksResponse) xxx_FromOp(ctx context.Context, op *xxx_RescanDisksOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RescanDisksResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RescanDisksResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RescanDisksOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RefreshFileSystemOperation structure represents the RefreshFileSys operation
type xxx_RefreshFileSystemOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RefreshFileSystemOperation) OpNum() int { return 75 }

func (o *xxx_RefreshFileSystemOperation) OpName() string { return "/IVolumeClient/v0/RefreshFileSys" }

func (o *xxx_RefreshFileSystemOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RefreshFileSystemOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RefreshFileSystemOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RefreshFileSystemOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RefreshFileSystemOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RefreshFileSystemOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RefreshFileSystemRequest structure represents the RefreshFileSys operation request
type RefreshFileSystemRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *RefreshFileSystemRequest) xxx_ToOp(ctx context.Context) *xxx_RefreshFileSystemOperation {
	if o == nil {
		return &xxx_RefreshFileSystemOperation{}
	}
	return &xxx_RefreshFileSystemOperation{
		This: o.This,
	}
}

func (o *RefreshFileSystemRequest) xxx_FromOp(ctx context.Context, op *xxx_RefreshFileSystemOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *RefreshFileSystemRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RefreshFileSystemRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RefreshFileSystemOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RefreshFileSystemResponse structure represents the RefreshFileSys operation response
type RefreshFileSystemResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RefreshFileSys return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RefreshFileSystemResponse) xxx_ToOp(ctx context.Context) *xxx_RefreshFileSystemOperation {
	if o == nil {
		return &xxx_RefreshFileSystemOperation{}
	}
	return &xxx_RefreshFileSystemOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *RefreshFileSystemResponse) xxx_FromOp(ctx context.Context, op *xxx_RefreshFileSystemOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RefreshFileSystemResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RefreshFileSystemResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RefreshFileSystemOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SecureSystemPartitionOperation structure represents the SecureSystemPartition operation
type xxx_SecureSystemPartitionOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SecureSystemPartitionOperation) OpNum() int { return 76 }

func (o *xxx_SecureSystemPartitionOperation) OpName() string {
	return "/IVolumeClient/v0/SecureSystemPartition"
}

func (o *xxx_SecureSystemPartitionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SecureSystemPartitionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SecureSystemPartitionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SecureSystemPartitionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SecureSystemPartitionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SecureSystemPartitionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SecureSystemPartitionRequest structure represents the SecureSystemPartition operation request
type SecureSystemPartitionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *SecureSystemPartitionRequest) xxx_ToOp(ctx context.Context) *xxx_SecureSystemPartitionOperation {
	if o == nil {
		return &xxx_SecureSystemPartitionOperation{}
	}
	return &xxx_SecureSystemPartitionOperation{
		This: o.This,
	}
}

func (o *SecureSystemPartitionRequest) xxx_FromOp(ctx context.Context, op *xxx_SecureSystemPartitionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *SecureSystemPartitionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SecureSystemPartitionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SecureSystemPartitionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SecureSystemPartitionResponse structure represents the SecureSystemPartition operation response
type SecureSystemPartitionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SecureSystemPartition return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SecureSystemPartitionResponse) xxx_ToOp(ctx context.Context) *xxx_SecureSystemPartitionOperation {
	if o == nil {
		return &xxx_SecureSystemPartitionOperation{}
	}
	return &xxx_SecureSystemPartitionOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SecureSystemPartitionResponse) xxx_FromOp(ctx context.Context, op *xxx_SecureSystemPartitionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SecureSystemPartitionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SecureSystemPartitionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SecureSystemPartitionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ShutDownSystemOperation structure represents the ShutDownSystem operation
type xxx_ShutDownSystemOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ShutDownSystemOperation) OpNum() int { return 77 }

func (o *xxx_ShutDownSystemOperation) OpName() string { return "/IVolumeClient/v0/ShutDownSystem" }

func (o *xxx_ShutDownSystemOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ShutDownSystemOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ShutDownSystemOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ShutDownSystemOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ShutDownSystemOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ShutDownSystemOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ShutDownSystemRequest structure represents the ShutDownSystem operation request
type ShutDownSystemRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *ShutDownSystemRequest) xxx_ToOp(ctx context.Context) *xxx_ShutDownSystemOperation {
	if o == nil {
		return &xxx_ShutDownSystemOperation{}
	}
	return &xxx_ShutDownSystemOperation{
		This: o.This,
	}
}

func (o *ShutDownSystemRequest) xxx_FromOp(ctx context.Context, op *xxx_ShutDownSystemOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *ShutDownSystemRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ShutDownSystemRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ShutDownSystemOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ShutDownSystemResponse structure represents the ShutDownSystem operation response
type ShutDownSystemResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ShutDownSystem return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ShutDownSystemResponse) xxx_ToOp(ctx context.Context) *xxx_ShutDownSystemOperation {
	if o == nil {
		return &xxx_ShutDownSystemOperation{}
	}
	return &xxx_ShutDownSystemOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *ShutDownSystemResponse) xxx_FromOp(ctx context.Context, op *xxx_ShutDownSystemOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ShutDownSystemResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ShutDownSystemResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ShutDownSystemOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumAccessPathOperation structure represents the EnumAccessPath operation
type xxx_EnumAccessPathOperation struct {
	This   *dcom.ORPCThis        `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat        `idl:"name:That" json:"that"`
	Count  int32                 `idl:"name:lCount" json:"count"`
	Paths  []*dmrp.CountedString `idl:"name:paths;size_is:(, lCount)" json:"paths"`
	Return int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumAccessPathOperation) OpNum() int { return 78 }

func (o *xxx_EnumAccessPathOperation) OpName() string { return "/IVolumeClient/v0/EnumAccessPath" }

func (o *xxx_EnumAccessPathOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumAccessPathOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lCount {in, out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumAccessPathOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lCount {in, out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumAccessPathOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Paths != nil && o.Count == 0 {
		o.Count = int32(len(o.Paths))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumAccessPathOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lCount {in, out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	// paths {out} (1:{pointer=ref}*(2)*(1))(2:{alias=COUNTED_STRING}[dim:0,size_is=lCount](struct))
	{
		if o.Paths != nil || o.Count > 0 {
			_ptr_paths := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.Count)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.Paths {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.Paths[i1] != nil {
						if err := o.Paths[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dmrp.CountedString{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.Paths); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&dmrp.CountedString{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Paths, _ptr_paths); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumAccessPathOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lCount {in, out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	// paths {out} (1:{pointer=ref}*(2)*(1))(2:{alias=COUNTED_STRING}[dim:0,size_is=lCount](struct))
	{
		_ptr_paths := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Paths", sizeInfo[0])
			}
			o.Paths = make([]*dmrp.CountedString, sizeInfo[0])
			for i1 := range o.Paths {
				i1 := i1
				if o.Paths[i1] == nil {
					o.Paths[i1] = &dmrp.CountedString{}
				}
				if err := o.Paths[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_paths := func(ptr interface{}) { o.Paths = *ptr.(*[]*dmrp.CountedString) }
		if err := w.ReadPointer(&o.Paths, _s_paths, _ptr_paths); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumAccessPathRequest structure represents the EnumAccessPath operation request
type EnumAccessPathRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This  *dcom.ORPCThis `idl:"name:This" json:"this"`
	Count int32          `idl:"name:lCount" json:"count"`
}

func (o *EnumAccessPathRequest) xxx_ToOp(ctx context.Context) *xxx_EnumAccessPathOperation {
	if o == nil {
		return &xxx_EnumAccessPathOperation{}
	}
	return &xxx_EnumAccessPathOperation{
		This:  o.This,
		Count: o.Count,
	}
}

func (o *EnumAccessPathRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumAccessPathOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Count = op.Count
}
func (o *EnumAccessPathRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumAccessPathRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumAccessPathOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumAccessPathResponse structure represents the EnumAccessPath operation response
type EnumAccessPathResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat        `idl:"name:That" json:"that"`
	Count int32                 `idl:"name:lCount" json:"count"`
	Paths []*dmrp.CountedString `idl:"name:paths;size_is:(, lCount)" json:"paths"`
	// Return: The EnumAccessPath return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumAccessPathResponse) xxx_ToOp(ctx context.Context) *xxx_EnumAccessPathOperation {
	if o == nil {
		return &xxx_EnumAccessPathOperation{}
	}
	return &xxx_EnumAccessPathOperation{
		That:   o.That,
		Count:  o.Count,
		Paths:  o.Paths,
		Return: o.Return,
	}
}

func (o *EnumAccessPathResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumAccessPathOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Count = op.Count
	o.Paths = op.Paths
	o.Return = op.Return
}
func (o *EnumAccessPathResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumAccessPathResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumAccessPathOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumAccessPathForVolumeOperation structure represents the EnumAccessPathForVolume operation
type xxx_EnumAccessPathForVolumeOperation struct {
	This     *dcom.ORPCThis        `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat        `idl:"name:That" json:"that"`
	VolumeID int64                 `idl:"name:VolumeId" json:"volume_id"`
	Count    int32                 `idl:"name:lCount" json:"count"`
	Paths    []*dmrp.CountedString `idl:"name:paths;size_is:(, lCount)" json:"paths"`
	Return   int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumAccessPathForVolumeOperation) OpNum() int { return 79 }

func (o *xxx_EnumAccessPathForVolumeOperation) OpName() string {
	return "/IVolumeClient/v0/EnumAccessPathForVolume"
}

func (o *xxx_EnumAccessPathForVolumeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumAccessPathForVolumeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// VolumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.VolumeID); err != nil {
			return err
		}
	}
	// lCount {in, out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumAccessPathForVolumeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// VolumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.VolumeID); err != nil {
			return err
		}
	}
	// lCount {in, out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumAccessPathForVolumeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Paths != nil && o.Count == 0 {
		o.Count = int32(len(o.Paths))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumAccessPathForVolumeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lCount {in, out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	// paths {out} (1:{pointer=ref}*(2)*(1))(2:{alias=COUNTED_STRING}[dim:0,size_is=lCount](struct))
	{
		if o.Paths != nil || o.Count > 0 {
			_ptr_paths := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.Count)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.Paths {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.Paths[i1] != nil {
						if err := o.Paths[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dmrp.CountedString{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.Paths); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&dmrp.CountedString{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Paths, _ptr_paths); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumAccessPathForVolumeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lCount {in, out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	// paths {out} (1:{pointer=ref}*(2)*(1))(2:{alias=COUNTED_STRING}[dim:0,size_is=lCount](struct))
	{
		_ptr_paths := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Paths", sizeInfo[0])
			}
			o.Paths = make([]*dmrp.CountedString, sizeInfo[0])
			for i1 := range o.Paths {
				i1 := i1
				if o.Paths[i1] == nil {
					o.Paths[i1] = &dmrp.CountedString{}
				}
				if err := o.Paths[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_paths := func(ptr interface{}) { o.Paths = *ptr.(*[]*dmrp.CountedString) }
		if err := w.ReadPointer(&o.Paths, _s_paths, _ptr_paths); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumAccessPathForVolumeRequest structure represents the EnumAccessPathForVolume operation request
type EnumAccessPathForVolumeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	VolumeID int64          `idl:"name:VolumeId" json:"volume_id"`
	Count    int32          `idl:"name:lCount" json:"count"`
}

func (o *EnumAccessPathForVolumeRequest) xxx_ToOp(ctx context.Context) *xxx_EnumAccessPathForVolumeOperation {
	if o == nil {
		return &xxx_EnumAccessPathForVolumeOperation{}
	}
	return &xxx_EnumAccessPathForVolumeOperation{
		This:     o.This,
		VolumeID: o.VolumeID,
		Count:    o.Count,
	}
}

func (o *EnumAccessPathForVolumeRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumAccessPathForVolumeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VolumeID = op.VolumeID
	o.Count = op.Count
}
func (o *EnumAccessPathForVolumeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumAccessPathForVolumeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumAccessPathForVolumeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumAccessPathForVolumeResponse structure represents the EnumAccessPathForVolume operation response
type EnumAccessPathForVolumeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat        `idl:"name:That" json:"that"`
	Count int32                 `idl:"name:lCount" json:"count"`
	Paths []*dmrp.CountedString `idl:"name:paths;size_is:(, lCount)" json:"paths"`
	// Return: The EnumAccessPathForVolume return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumAccessPathForVolumeResponse) xxx_ToOp(ctx context.Context) *xxx_EnumAccessPathForVolumeOperation {
	if o == nil {
		return &xxx_EnumAccessPathForVolumeOperation{}
	}
	return &xxx_EnumAccessPathForVolumeOperation{
		That:   o.That,
		Count:  o.Count,
		Paths:  o.Paths,
		Return: o.Return,
	}
}

func (o *EnumAccessPathForVolumeResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumAccessPathForVolumeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Count = op.Count
	o.Paths = op.Paths
	o.Return = op.Return
}
func (o *EnumAccessPathForVolumeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumAccessPathForVolumeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumAccessPathForVolumeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddAccessPathOperation structure represents the AddAccessPath operation
type xxx_AddAccessPathOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	PathLength int32          `idl:"name:cch_path" json:"path_length"`
	Path       string         `idl:"name:path;size_is:(cch_path)" json:"path"`
	TargetID   int64          `idl:"name:targetId" json:"target_id"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_AddAccessPathOperation) OpNum() int { return 80 }

func (o *xxx_AddAccessPathOperation) OpName() string { return "/IVolumeClient/v0/AddAccessPath" }

func (o *xxx_AddAccessPathOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Path != "" && o.PathLength == 0 {
		o.PathLength = int32(len(o.Path))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddAccessPathOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// cch_path {in} (1:(int32))
	{
		if err := w.WriteData(o.PathLength); err != nil {
			return err
		}
	}
	// path {in} (1:{pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,size_is=cch_path,string](wchar))
	{
		dimSize1 := uint64(o.PathLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		_Path_buf := utf16.Encode([]rune(o.Path))
		if uint64(len(_Path_buf)) > sizeInfo[0] {
			_Path_buf = _Path_buf[:sizeInfo[0]]
		}
		for i1 := range _Path_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_Path_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_Path_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	// targetId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.TargetID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddAccessPathOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cch_path {in} (1:(int32))
	{
		if err := w.ReadData(&o.PathLength); err != nil {
			return err
		}
	}
	// path {in} (1:{pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,size_is=cch_path,string](wchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _Path_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Path_buf", sizeInfo[0])
		}
		_Path_buf = make([]uint16, sizeInfo[0])
		for i1 := range _Path_buf {
			i1 := i1
			if err := w.ReadData(&_Path_buf[i1]); err != nil {
				return err
			}
		}
		o.Path = strings.TrimRight(string(utf16.Decode(_Path_buf)), ndr.ZeroString)
	}
	// targetId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.TargetID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddAccessPathOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddAccessPathOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddAccessPathOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AddAccessPathRequest structure represents the AddAccessPath operation request
type AddAccessPathRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	PathLength int32          `idl:"name:cch_path" json:"path_length"`
	Path       string         `idl:"name:path;size_is:(cch_path)" json:"path"`
	TargetID   int64          `idl:"name:targetId" json:"target_id"`
}

func (o *AddAccessPathRequest) xxx_ToOp(ctx context.Context) *xxx_AddAccessPathOperation {
	if o == nil {
		return &xxx_AddAccessPathOperation{}
	}
	return &xxx_AddAccessPathOperation{
		This:       o.This,
		PathLength: o.PathLength,
		Path:       o.Path,
		TargetID:   o.TargetID,
	}
}

func (o *AddAccessPathRequest) xxx_FromOp(ctx context.Context, op *xxx_AddAccessPathOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PathLength = op.PathLength
	o.Path = op.Path
	o.TargetID = op.TargetID
}
func (o *AddAccessPathRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AddAccessPathRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddAccessPathOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddAccessPathResponse structure represents the AddAccessPath operation response
type AddAccessPathResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AddAccessPath return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AddAccessPathResponse) xxx_ToOp(ctx context.Context) *xxx_AddAccessPathOperation {
	if o == nil {
		return &xxx_AddAccessPathOperation{}
	}
	return &xxx_AddAccessPathOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *AddAccessPathResponse) xxx_FromOp(ctx context.Context, op *xxx_AddAccessPathOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *AddAccessPathResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AddAccessPathResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddAccessPathOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteAccessPathOperation structure represents the DeleteAccessPath operation
type xxx_DeleteAccessPathOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	VolumeID   int64          `idl:"name:volumeId" json:"volume_id"`
	PathLength int32          `idl:"name:cch_path" json:"path_length"`
	Path       string         `idl:"name:path;size_is:(cch_path)" json:"path"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteAccessPathOperation) OpNum() int { return 81 }

func (o *xxx_DeleteAccessPathOperation) OpName() string { return "/IVolumeClient/v0/DeleteAccessPath" }

func (o *xxx_DeleteAccessPathOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Path != "" && o.PathLength == 0 {
		o.PathLength = int32(len(o.Path))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteAccessPathOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.VolumeID); err != nil {
			return err
		}
	}
	// cch_path {in} (1:(int32))
	{
		if err := w.WriteData(o.PathLength); err != nil {
			return err
		}
	}
	// path {in} (1:{pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,size_is=cch_path,string](wchar))
	{
		dimSize1 := uint64(o.PathLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		_Path_buf := utf16.Encode([]rune(o.Path))
		if uint64(len(_Path_buf)) > sizeInfo[0] {
			_Path_buf = _Path_buf[:sizeInfo[0]]
		}
		for i1 := range _Path_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_Path_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_Path_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_DeleteAccessPathOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.VolumeID); err != nil {
			return err
		}
	}
	// cch_path {in} (1:(int32))
	{
		if err := w.ReadData(&o.PathLength); err != nil {
			return err
		}
	}
	// path {in} (1:{pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,size_is=cch_path,string](wchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _Path_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Path_buf", sizeInfo[0])
		}
		_Path_buf = make([]uint16, sizeInfo[0])
		for i1 := range _Path_buf {
			i1 := i1
			if err := w.ReadData(&_Path_buf[i1]); err != nil {
				return err
			}
		}
		o.Path = strings.TrimRight(string(utf16.Decode(_Path_buf)), ndr.ZeroString)
	}
	return nil
}

func (o *xxx_DeleteAccessPathOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteAccessPathOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteAccessPathOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DeleteAccessPathRequest structure represents the DeleteAccessPath operation request
type DeleteAccessPathRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	VolumeID   int64          `idl:"name:volumeId" json:"volume_id"`
	PathLength int32          `idl:"name:cch_path" json:"path_length"`
	Path       string         `idl:"name:path;size_is:(cch_path)" json:"path"`
}

func (o *DeleteAccessPathRequest) xxx_ToOp(ctx context.Context) *xxx_DeleteAccessPathOperation {
	if o == nil {
		return &xxx_DeleteAccessPathOperation{}
	}
	return &xxx_DeleteAccessPathOperation{
		This:       o.This,
		VolumeID:   o.VolumeID,
		PathLength: o.PathLength,
		Path:       o.Path,
	}
}

func (o *DeleteAccessPathRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteAccessPathOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VolumeID = op.VolumeID
	o.PathLength = op.PathLength
	o.Path = op.Path
}
func (o *DeleteAccessPathRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteAccessPathRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteAccessPathOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteAccessPathResponse structure represents the DeleteAccessPath operation response
type DeleteAccessPathResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteAccessPath return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteAccessPathResponse) xxx_ToOp(ctx context.Context) *xxx_DeleteAccessPathOperation {
	if o == nil {
		return &xxx_DeleteAccessPathOperation{}
	}
	return &xxx_DeleteAccessPathOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *DeleteAccessPathResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteAccessPathOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteAccessPathResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteAccessPathResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteAccessPathOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
