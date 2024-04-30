package ivdsvolumemf

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

// IVdsVolumeMF server interface.
type VolumeMFServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The GetFileSystemProperties method returns property details about the file system
	// on the current volume.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetFileSystemProperties(context.Context, *GetFileSystemPropertiesRequest) (*GetFileSystemPropertiesResponse, error)

	// The Format method formats a file system on the current volume.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	Format(context.Context, *FormatRequest) (*FormatResponse, error)

	// The AddAccessPath method adds an access path to the current volume.<129>
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	AddAccessPath(context.Context, *AddAccessPathRequest) (*AddAccessPathResponse, error)

	// The QueryAccessPaths method returns a list of access paths and a drive letter as
	// a single case-insensitive Unicode character, if one exists, for the current volume.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryAccessPaths(context.Context, *QueryAccessPathsRequest) (*QueryAccessPathsResponse, error)

	// The QueryReparsePoints method returns all reparse points for the current volume.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryReparsePoints(context.Context, *QueryReparsePointsRequest) (*QueryReparsePointsResponse, error)

	// The DeleteAccessPath method removes the access path from the current volume.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	DeleteAccessPath(context.Context, *DeleteAccessPathRequest) (*DeleteAccessPathResponse, error)

	// The Mount method mounts a volume.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	Mount(context.Context, *MountRequest) (*MountResponse, error)

	// The Dismount method dismounts a mounted volume.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	Dismount(context.Context, *DismountRequest) (*DismountResponse, error)

	// The SetFileSystemFlags method sets the file system flags.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	SetFileSystemFlags(context.Context, *SetFileSystemFlagsRequest) (*SetFileSystemFlagsResponse, error)

	// The ClearFileSystemFlags method clears the file system flags.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	ClearFileSystemFlags(context.Context, *ClearFileSystemFlagsRequest) (*ClearFileSystemFlagsResponse, error)
}

func RegisterVolumeMFServer(conn dcerpc.Conn, o VolumeMFServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewVolumeMFServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(VolumeMFSyntaxV0_0))...)
}

func NewVolumeMFServerHandle(o VolumeMFServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return VolumeMFServerHandle(ctx, o, opNum, r)
	}
}

func VolumeMFServerHandle(ctx context.Context, o VolumeMFServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetFileSystemProperties
		in := &GetFileSystemPropertiesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFileSystemProperties(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // Format
		in := &FormatRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Format(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // AddAccessPath
		in := &AddAccessPathRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddAccessPath(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // QueryAccessPaths
		in := &QueryAccessPathsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryAccessPaths(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // QueryReparsePoints
		in := &QueryReparsePointsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryReparsePoints(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // DeleteAccessPath
		in := &DeleteAccessPathRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteAccessPath(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // Mount
		in := &MountRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Mount(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // Dismount
		in := &DismountRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Dismount(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // SetFileSystemFlags
		in := &SetFileSystemFlagsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetFileSystemFlags(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // ClearFileSystemFlags
		in := &ClearFileSystemFlagsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ClearFileSystemFlags(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
