package ivdsserviceuninstalldisk

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

// IVdsServiceUninstallDisk server interface.
type ServiceUninstallDiskServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The GetDiskIdFromLunInfo method retrieves the VDS object ID of a disk that corresponds
	// to a specified LUN information structure.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetDiskIDFromLUNInfo(context.Context, *GetDiskIDFromLUNInfoRequest) (*GetDiskIDFromLUNInfoResponse, error)

	// The UninstallDisks method uninstalls a specific set of disks when it is given a list
	// of the VDS object IDs for the disks. All volumes that are contained wholly or partially
	// on the disks are also uninstalled, and the obsolete mount points are removed.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	UninstallDisks(context.Context, *UninstallDisksRequest) (*UninstallDisksResponse, error)
}

func RegisterServiceUninstallDiskServer(conn dcerpc.Conn, o ServiceUninstallDiskServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewServiceUninstallDiskServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ServiceUninstallDiskSyntaxV0_0))...)
}

func NewServiceUninstallDiskServerHandle(o ServiceUninstallDiskServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ServiceUninstallDiskServerHandle(ctx, o, opNum, r)
	}
}

func ServiceUninstallDiskServerHandle(ctx context.Context, o ServiceUninstallDiskServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetDiskIdFromLunInfo
		op := &xxx_GetDiskIDFromLUNInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDiskIDFromLUNInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDiskIDFromLUNInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // UninstallDisks
		op := &xxx_UninstallDisksOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UninstallDisksRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.UninstallDisks(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IVdsServiceUninstallDisk
type UnimplementedServiceUninstallDiskServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedServiceUninstallDiskServer) GetDiskIDFromLUNInfo(context.Context, *GetDiskIDFromLUNInfoRequest) (*GetDiskIDFromLUNInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServiceUninstallDiskServer) UninstallDisks(context.Context, *UninstallDisksRequest) (*UninstallDisksResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ServiceUninstallDiskServer = (*UnimplementedServiceUninstallDiskServer)(nil)
