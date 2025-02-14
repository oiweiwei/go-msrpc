package ivdsopenvdisk

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

// IVdsOpenVDisk server interface.
type OpenVDiskServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The Attach method creates an operating system disk device for a virtual disk.
	//
	// Return Values: This method MUST return zero to indicate success, or return an implementation-specific
	// nonzero error code to indicate failure.
	Attach(context.Context, *AttachRequest) (*AttachResponse, error)

	// The Detach method removes the operating system disk device that represents a virtual
	// disk.
	//
	// Return Values: The method MUST return zero to indicate success, or return an implementation-specific
	// nonzero error code to indicate failure.
	Detach(context.Context, *DetachRequest) (*DetachResponse, error)

	// The DetachAndDelete method removes the operating system disk device that represents
	// a virtual disk and deletes any backing store file.
	//
	// Return Values: The method MUST return zero to indicate success, or return an implementation-specific
	// nonzero error code to indicate failure.
	DetachAndDelete(context.Context, *DetachAndDeleteRequest) (*DetachAndDeleteResponse, error)

	// The Compact method reduces the size of the virtual disk file (the backing store).
	// This requires that the virtual disk be detached. Compact is applicable only to differencing
	// type virtual disks and virtual disks created using CREATE_VIRTUAL_DISK_FLAG_NONE.
	// The Compact method does not change the size of the virtual disk.
	//
	// Return Values: The method MUST return zero to indicate success, or return an implementation-specific
	// nonzero error code to indicate failure.
	Compact(context.Context, *CompactRequest) (*CompactResponse, error)

	//	This method is applicable only to differencing type virtual disks. The Merge method moves all data blocks from a differencing virtual disk into its parent virtual disk. Merging a virtual disk requires that the virtual disk be detached during the operation. Both the virtual disk and its parent are opened READ|WRITE using the IVdsVDisk::Open method called against the virtual disk with an appropriate value for the ReadWriteDepth, as described later in this section.<145>
	//
	// For example, to merge a differencing disk that is a child of a single parent disk
	// into that parent disk, call the IVdsVDisk::Open method on the child disk with the
	// ReadWriteDepth parameter set to the value 2. This value opens both disks with the
	// READ and WRITE flags set, which is necessary for disks to be merged with subsequent
	// call to the IVdsOpenVDisk::Merge method.
	//
	// Return Values: The method MUST return zero to indicate success, or return an implementation-specific
	// nonzero error code to indicate failure.
	Merge(context.Context, *MergeRequest) (*MergeResponse, error)

	//	The Expand method increases the size of a virtual disk. Expanding a virtual disk requires that the virtual disk be detached during the operation. The virtual disk file is opened with READ|WRITE privileges using the IVdsVDisk::Open method.<147>
	//
	// Return Values: The method MUST return zero to indicate success, or return an implementation-specific
	// nonzero error code to indicate failure.
	Expand(context.Context, *ExpandRequest) (*ExpandResponse, error)
}

func RegisterOpenVDiskServer(conn dcerpc.Conn, o OpenVDiskServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewOpenVDiskServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(OpenVDiskSyntaxV0_0))...)
}

func NewOpenVDiskServerHandle(o OpenVDiskServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return OpenVDiskServerHandle(ctx, o, opNum, r)
	}
}

func OpenVDiskServerHandle(ctx context.Context, o OpenVDiskServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Attach
		op := &xxx_AttachOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AttachRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Attach(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // Detach
		op := &xxx_DetachOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DetachRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Detach(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // DetachAndDelete
		op := &xxx_DetachAndDeleteOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DetachAndDeleteRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DetachAndDelete(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // Compact
		op := &xxx_CompactOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CompactRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Compact(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // Merge
		op := &xxx_MergeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MergeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Merge(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // Expand
		op := &xxx_ExpandOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ExpandRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Expand(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IVdsOpenVDisk
type UnimplementedOpenVDiskServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedOpenVDiskServer) Attach(context.Context, *AttachRequest) (*AttachResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedOpenVDiskServer) Detach(context.Context, *DetachRequest) (*DetachResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedOpenVDiskServer) DetachAndDelete(context.Context, *DetachAndDeleteRequest) (*DetachAndDeleteResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedOpenVDiskServer) Compact(context.Context, *CompactRequest) (*CompactResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedOpenVDiskServer) Merge(context.Context, *MergeRequest) (*MergeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedOpenVDiskServer) Expand(context.Context, *ExpandRequest) (*ExpandResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ OpenVDiskServer = (*UnimplementedOpenVDiskServer)(nil)
