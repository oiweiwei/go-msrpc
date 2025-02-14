package ivdspack

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

// IVdsPack server interface.
type PackServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// GetProperties operation.
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)

	// The GetProvider method retrieves the provider that the disk pack belongs to.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetProvider(context.Context, *GetProviderRequest) (*GetProviderResponse, error)

	// The QueryVolumes method retrieves the volumes of a disk pack.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryVolumes(context.Context, *QueryVolumesRequest) (*QueryVolumesResponse, error)

	// The QueryDisks method retrieves the disks of a disk pack.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryDisks(context.Context, *QueryDisksRequest) (*QueryDisksResponse, error)

	// The CreateVolume method creates a volume in a disk pack.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	CreateVolume(context.Context, *CreateVolumeRequest) (*CreateVolumeResponse, error)

	// This method initializes a disk that has no partitioning format defined, and then
	// adds the disk to the disk pack. AddDisk cannot redefine the partitioning format on
	// a disk.<82>
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	AddDisk(context.Context, *AddDiskRequest) (*AddDiskResponse, error)

	// The MigrateDisks method migrates a set of disks from one pack to another pack.<83>
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	//
	// ERROR_SUCCESS (0x00000000)
	MigrateDisks(context.Context, *MigrateDisksRequest) (*MigrateDisksResponse, error)

	// Opnum10NotUsedOnWire operation.
	// Opnum10NotUsedOnWire

	// The IVdsPack::RemoveMissingDisk method removes the specified missing disk from a
	// disk pack. This method only applies to dynamic disks. At least one dynamic disk needs
	// to be present to enumerate missing disks.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	RemoveMissingDisk(context.Context, *RemoveMissingDiskRequest) (*RemoveMissingDiskResponse, error)

	// The Recover method restores a disk pack to a healthy state. This method is not supported
	// on basic disk packs or the INVALID dynamic disk pack (the value of VDS_PACK_PROP::pwszName
	// is INVALID for this pack). The INVALID dynamic disk pack contains dynamic disks that
	// have failed to be joined to the owning pack because there are errors or data corruption
	// has occurred.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	Recover(context.Context, *RecoverRequest) (*RecoverResponse, error)
}

func RegisterPackServer(conn dcerpc.Conn, o PackServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewPackServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(PackSyntaxV0_0))...)
}

func NewPackServerHandle(o PackServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return PackServerHandle(ctx, o, opNum, r)
	}
}

func PackServerHandle(ctx context.Context, o PackServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetProperties
		op := &xxx_GetPropertiesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPropertiesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetProperties(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // GetProvider
		op := &xxx_GetProviderOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetProviderRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetProvider(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // QueryVolumes
		op := &xxx_QueryVolumesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryVolumesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryVolumes(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // QueryDisks
		op := &xxx_QueryDisksOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryDisksRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryDisks(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // CreateVolume
		op := &xxx_CreateVolumeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateVolumeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateVolume(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // AddDisk
		op := &xxx_AddDiskOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddDiskRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddDisk(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // MigrateDisks
		op := &xxx_MigrateDisksOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MigrateDisksRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MigrateDisks(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // Opnum10NotUsedOnWire
		// Opnum10NotUsedOnWire
		return nil, nil
	case 11: // RemoveMissingDisk
		op := &xxx_RemoveMissingDiskOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveMissingDiskRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveMissingDisk(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // Recover
		op := &xxx_RecoverOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RecoverRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Recover(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IVdsPack
type UnimplementedPackServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedPackServer) GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPackServer) GetProvider(context.Context, *GetProviderRequest) (*GetProviderResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPackServer) QueryVolumes(context.Context, *QueryVolumesRequest) (*QueryVolumesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPackServer) QueryDisks(context.Context, *QueryDisksRequest) (*QueryDisksResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPackServer) CreateVolume(context.Context, *CreateVolumeRequest) (*CreateVolumeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPackServer) AddDisk(context.Context, *AddDiskRequest) (*AddDiskResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPackServer) MigrateDisks(context.Context, *MigrateDisksRequest) (*MigrateDisksResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPackServer) RemoveMissingDisk(context.Context, *RemoveMissingDiskRequest) (*RemoveMissingDiskResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPackServer) Recover(context.Context, *RecoverRequest) (*RecoverResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ PackServer = (*UnimplementedPackServer)(nil)
