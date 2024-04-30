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
		in := &GetPropertiesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetProperties(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // GetProvider
		in := &GetProviderRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetProvider(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // QueryVolumes
		in := &QueryVolumesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryVolumes(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // QueryDisks
		in := &QueryDisksRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryDisks(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // CreateVolume
		in := &CreateVolumeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateVolume(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // AddDisk
		in := &AddDiskRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddDisk(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // MigrateDisks
		in := &MigrateDisksRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.MigrateDisks(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // Opnum10NotUsedOnWire
		// Opnum10NotUsedOnWire
		return nil, nil
	case 11: // RemoveMissingDisk
		in := &RemoveMissingDiskRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoveMissingDisk(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // Recover
		in := &RecoverRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Recover(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
