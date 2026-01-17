package ivolumeclient4

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

// IVolumeClient4 server interface.
type VolumeClient4Server interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The RefreshEx method refreshes the server's cache of storage objects, including regions,
	// removable media and CD-ROM drive media, file systems, and drive letters.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return 0 or a nonerror HRESULT on success, or an implementation-specific
	// nonzero error code on failure (as specified in [MS-ERREF]; see also section 2.2.1
	// for HRESULT values predefined by the Disk Management Remote Protocol).
	//
	// The handling of this message is identical to the handling of IVolumeClient::Refresh
	// and IVolumeClient3::Refresh except that the server MUST perform an extra low-level
	// refresh of the list of storage objects by looking for missing dynamic disks or dynamic
	// disks that were missing and are now present. This verification updates the status
	// for missing disks, volumes that reside on missing disks, or disk regions that reside
	// on missing disks.<228>
	//
	// In addition to the preceding actions, the server MUST check whether the lengths of
	// the disks have changed and make appropriate changes to the disk objects in the list
	// of storage objects.
	RefreshEx(context.Context, *RefreshExRequest) (*RefreshExResponse, error)

	// The GetVolumeDeviceName method retrieves the Windows NT operating system device name
	// of a dynamic volume on the server.
	//
	// Return Values: The method MUST return 0 or a nonerror HRESULT on success, or an implementation-specific
	// nonzero error code on failure (as specified in [MS-ERREF]; see also section 2.2.1
	// for HRESULT values predefined by the Disk Management Remote Protocol).
	GetVolumeDeviceName(context.Context, *GetVolumeDeviceNameRequest) (*GetVolumeDeviceNameResponse, error)
}

func RegisterVolumeClient4Server(conn dcerpc.Conn, o VolumeClient4Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewVolumeClient4ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(VolumeClient4SyntaxV0_0))...)
}

func NewVolumeClient4ServerHandle(o VolumeClient4Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return VolumeClient4ServerHandle(ctx, o, opNum, r)
	}
}

func VolumeClient4ServerHandle(ctx context.Context, o VolumeClient4Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // RefreshEx
		op := &xxx_RefreshExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RefreshExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RefreshEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // GetVolumeDeviceName
		op := &xxx_GetVolumeDeviceNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetVolumeDeviceNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetVolumeDeviceName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IVolumeClient4
type UnimplementedVolumeClient4Server struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedVolumeClient4Server) RefreshEx(context.Context, *RefreshExRequest) (*RefreshExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeClient4Server) GetVolumeDeviceName(context.Context, *GetVolumeDeviceNameRequest) (*GetVolumeDeviceNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ VolumeClient4Server = (*UnimplementedVolumeClient4Server)(nil)
