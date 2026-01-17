package ivdsadvanceddisk

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

// IVdsAdvancedDisk server interface.
type AdvancedDiskServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The GetPartitionProperties method retrieves the properties of a partition on the
	// disk at a specified byte offset.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetPartitionProperties(context.Context, *GetPartitionPropertiesRequest) (*GetPartitionPropertiesResponse, error)

	// The QueryPartitions method enumerates a disk's partitions.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryPartitions(context.Context, *QueryPartitionsRequest) (*QueryPartitionsResponse, error)

	// The CreatePartition method creates a partition on a disk at a specified byte offset.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.<94>
	//
	// Note  Creating or deleting partitions on dynamic disks is not supported.
	CreatePartition(context.Context, *CreatePartitionRequest) (*CreatePartitionResponse, error)

	// The DeletePartition method deletes a partition from the disk at a specified byte
	// offset.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	//
	// Note  Creating or deleting partitions on dynamic disks is not supported.
	//
	// ERROR_SUCCESS (0x00000000)
	DeletePartition(context.Context, *DeletePartitionRequest) (*DeletePartitionResponse, error)

	// The ChangeAttributes method changes the attributes of the partition at byte offset
	// ullOffset on the disk.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	ChangeAttributes(context.Context, *ChangeAttributesRequest) (*ChangeAttributesResponse, error)

	// The AssignDriveLetter method assigns a drive letter to an existing OEM, ESP, or unknown
	// partition.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	AssignDriveLetter(context.Context, *AssignDriveLetterRequest) (*AssignDriveLetterResponse, error)

	// The DeleteDriveLetter method deletes a drive letter that is assigned to an OEM, ESP,
	// or unknown partition.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	DeleteDriveLetter(context.Context, *DeleteDriveLetterRequest) (*DeleteDriveLetterResponse, error)

	// The GetDriveLetter method retrieves the drive letter of an OEM, ESP, or unknown partition
	// on the disk at a specified byte offset.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	//
	// ERROR_SUCCESS (0x00000000)
	GetDriveLetter(context.Context, *GetDriveLetterRequest) (*GetDriveLetterResponse, error)

	// The FormatPartition method formats an existing OEM, ESP, or unknown partition.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	//
	// ERROR_SUCCESS (0x00000000)
	FormatPartition(context.Context, *FormatPartitionRequest) (*FormatPartitionResponse, error)

	// The Clean method cleans a disk.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	Clean(context.Context, *CleanRequest) (*CleanResponse, error)
}

func RegisterAdvancedDiskServer(conn dcerpc.Conn, o AdvancedDiskServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAdvancedDiskServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AdvancedDiskSyntaxV0_0))...)
}

func NewAdvancedDiskServerHandle(o AdvancedDiskServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AdvancedDiskServerHandle(ctx, o, opNum, r)
	}
}

func AdvancedDiskServerHandle(ctx context.Context, o AdvancedDiskServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetPartitionProperties
		op := &xxx_GetPartitionPropertiesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPartitionPropertiesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPartitionProperties(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // QueryPartitions
		op := &xxx_QueryPartitionsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryPartitionsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryPartitions(ctx, req)
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
	case 6: // DeletePartition
		op := &xxx_DeletePartitionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeletePartitionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeletePartition(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // ChangeAttributes
		op := &xxx_ChangeAttributesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ChangeAttributesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ChangeAttributes(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // AssignDriveLetter
		op := &xxx_AssignDriveLetterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AssignDriveLetterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AssignDriveLetter(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // DeleteDriveLetter
		op := &xxx_DeleteDriveLetterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteDriveLetterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteDriveLetter(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // GetDriveLetter
		op := &xxx_GetDriveLetterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDriveLetterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDriveLetter(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // FormatPartition
		op := &xxx_FormatPartitionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FormatPartitionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.FormatPartition(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // Clean
		op := &xxx_CleanOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CleanRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Clean(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IVdsAdvancedDisk
type UnimplementedAdvancedDiskServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedAdvancedDiskServer) GetPartitionProperties(context.Context, *GetPartitionPropertiesRequest) (*GetPartitionPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAdvancedDiskServer) QueryPartitions(context.Context, *QueryPartitionsRequest) (*QueryPartitionsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAdvancedDiskServer) CreatePartition(context.Context, *CreatePartitionRequest) (*CreatePartitionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAdvancedDiskServer) DeletePartition(context.Context, *DeletePartitionRequest) (*DeletePartitionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAdvancedDiskServer) ChangeAttributes(context.Context, *ChangeAttributesRequest) (*ChangeAttributesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAdvancedDiskServer) AssignDriveLetter(context.Context, *AssignDriveLetterRequest) (*AssignDriveLetterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAdvancedDiskServer) DeleteDriveLetter(context.Context, *DeleteDriveLetterRequest) (*DeleteDriveLetterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAdvancedDiskServer) GetDriveLetter(context.Context, *GetDriveLetterRequest) (*GetDriveLetterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAdvancedDiskServer) FormatPartition(context.Context, *FormatPartitionRequest) (*FormatPartitionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAdvancedDiskServer) Clean(context.Context, *CleanRequest) (*CleanResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AdvancedDiskServer = (*UnimplementedAdvancedDiskServer)(nil)
