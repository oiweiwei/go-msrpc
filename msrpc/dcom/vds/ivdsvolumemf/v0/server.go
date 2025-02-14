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
		op := &xxx_GetFileSystemPropertiesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFileSystemPropertiesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFileSystemProperties(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // Format
		op := &xxx_FormatOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FormatRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Format(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // AddAccessPath
		op := &xxx_AddAccessPathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddAccessPathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddAccessPath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // QueryAccessPaths
		op := &xxx_QueryAccessPathsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryAccessPathsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryAccessPaths(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // QueryReparsePoints
		op := &xxx_QueryReparsePointsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryReparsePointsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryReparsePoints(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // DeleteAccessPath
		op := &xxx_DeleteAccessPathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteAccessPathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteAccessPath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // Mount
		op := &xxx_MountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Mount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // Dismount
		op := &xxx_DismountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DismountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Dismount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // SetFileSystemFlags
		op := &xxx_SetFileSystemFlagsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFileSystemFlagsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFileSystemFlags(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // ClearFileSystemFlags
		op := &xxx_ClearFileSystemFlagsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ClearFileSystemFlagsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ClearFileSystemFlags(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IVdsVolumeMF
type UnimplementedVolumeMFServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedVolumeMFServer) GetFileSystemProperties(context.Context, *GetFileSystemPropertiesRequest) (*GetFileSystemPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeMFServer) Format(context.Context, *FormatRequest) (*FormatResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeMFServer) AddAccessPath(context.Context, *AddAccessPathRequest) (*AddAccessPathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeMFServer) QueryAccessPaths(context.Context, *QueryAccessPathsRequest) (*QueryAccessPathsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeMFServer) QueryReparsePoints(context.Context, *QueryReparsePointsRequest) (*QueryReparsePointsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeMFServer) DeleteAccessPath(context.Context, *DeleteAccessPathRequest) (*DeleteAccessPathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeMFServer) Mount(context.Context, *MountRequest) (*MountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeMFServer) Dismount(context.Context, *DismountRequest) (*DismountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeMFServer) SetFileSystemFlags(context.Context, *SetFileSystemFlagsRequest) (*SetFileSystemFlagsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeMFServer) ClearFileSystemFlags(context.Context, *ClearFileSystemFlagsRequest) (*ClearFileSystemFlagsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ VolumeMFServer = (*UnimplementedVolumeMFServer)(nil)
