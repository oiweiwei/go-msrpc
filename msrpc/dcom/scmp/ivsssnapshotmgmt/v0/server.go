package ivsssnapshotmgmt

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

// IVssSnapshotMgmt server interface.
type SnapshotManagementServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	GetProviderManagementInterface(context.Context, *GetProviderManagementInterfaceRequest) (*GetProviderManagementInterfaceResponse, error)

	QueryVolumesSupportedForSnapshots(context.Context, *QueryVolumesSupportedForSnapshotsRequest) (*QueryVolumesSupportedForSnapshotsResponse, error)

	QuerySnapshotsByVolume(context.Context, *QuerySnapshotsByVolumeRequest) (*QuerySnapshotsByVolumeResponse, error)
}

func RegisterSnapshotManagementServer(conn dcerpc.Conn, o SnapshotManagementServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewSnapshotManagementServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(SnapshotManagementSyntaxV0_0))...)
}

func NewSnapshotManagementServerHandle(o SnapshotManagementServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return SnapshotManagementServerHandle(ctx, o, opNum, r)
	}
}

func SnapshotManagementServerHandle(ctx context.Context, o SnapshotManagementServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetProviderMgmtInterface
		op := &xxx_GetProviderManagementInterfaceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetProviderManagementInterfaceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetProviderManagementInterface(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // QueryVolumesSupportedForSnapshots
		op := &xxx_QueryVolumesSupportedForSnapshotsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryVolumesSupportedForSnapshotsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryVolumesSupportedForSnapshots(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // QuerySnapshotsByVolume
		op := &xxx_QuerySnapshotsByVolumeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QuerySnapshotsByVolumeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QuerySnapshotsByVolume(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IVssSnapshotMgmt
type UnimplementedSnapshotManagementServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedSnapshotManagementServer) GetProviderManagementInterface(context.Context, *GetProviderManagementInterfaceRequest) (*GetProviderManagementInterfaceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSnapshotManagementServer) QueryVolumesSupportedForSnapshots(context.Context, *QueryVolumesSupportedForSnapshotsRequest) (*QueryVolumesSupportedForSnapshotsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSnapshotManagementServer) QuerySnapshotsByVolume(context.Context, *QuerySnapshotsByVolumeRequest) (*QuerySnapshotsByVolumeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ SnapshotManagementServer = (*UnimplementedSnapshotManagementServer)(nil)
