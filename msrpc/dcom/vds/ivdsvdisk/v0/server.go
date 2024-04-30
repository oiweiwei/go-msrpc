package ivdsvdisk

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

// IVdsVDisk server interface.
type VDiskServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The Open method opens a handle to the specified virtual disk file and returns an
	// IVdsOpenVDisk (section 3.1.15.2) interface pointer to an object representing the
	// open virtual disk (an OpenVirtualDisk object). Release the IVdsOpenVDisk interface
	// to close the handle to the virtual disk.
	//
	// Return Values: The method MUST return zero to indicate success, or return an implementation-specific
	// nonzero error code to indicate failure.
	Open(context.Context, *OpenRequest) (*OpenResponse, error)

	// GetProperties operation.
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)

	// The GetHostVolume method returns an interface pointer to the volume object for the
	// volume on which the virtual disk backing store file resides.<139>
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetHostVolume(context.Context, *GetHostVolumeRequest) (*GetHostVolumeResponse, error)

	// The GetDeviceName method returns the device name of the disk.
	//
	// Return Values: The method MUST return zero to indicate success, or return an implementation-specific
	// nonzero error code to indicate failure.
	GetDeviceName(context.Context, *GetDeviceNameRequest) (*GetDeviceNameResponse, error)
}

func RegisterVDiskServer(conn dcerpc.Conn, o VDiskServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewVDiskServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(VDiskSyntaxV0_0))...)
}

func NewVDiskServerHandle(o VDiskServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return VDiskServerHandle(ctx, o, opNum, r)
	}
}

func VDiskServerHandle(ctx context.Context, o VDiskServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Open
		in := &OpenRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Open(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // GetProperties
		in := &GetPropertiesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetProperties(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // GetHostVolume
		in := &GetHostVolumeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetHostVolume(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // GetDeviceName
		in := &GetDeviceNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDeviceName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
