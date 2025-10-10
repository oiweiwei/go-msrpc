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
type VSSSnapshotManagementServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	GetProviderManagementInterface(context.Context, *GetProviderManagementInterfaceRequest) (*GetProviderManagementInterfaceResponse, error)

	QueryVolumesSupportedForSnapshots(context.Context, *QueryVolumesSupportedForSnapshotsRequest) (*QueryVolumesSupportedForSnapshotsResponse, error)

	QuerySnapshotsByVolume(context.Context, *QuerySnapshotsByVolumeRequest) (*QuerySnapshotsByVolumeResponse, error)
}

func RegisterVSSSnapshotManagementServer(conn dcerpc.Conn, o VSSSnapshotManagementServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewVSSSnapshotManagementServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(VSSSnapshotManagementSyntaxV0_0))...)
}

func NewVSSSnapshotManagementServerHandle(o VSSSnapshotManagementServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return VSSSnapshotManagementServerHandle(ctx, o, opNum, r)
	}
}

func VSSSnapshotManagementServerHandle(ctx context.Context, o VSSSnapshotManagementServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
type UnimplementedVSSSnapshotManagementServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedVSSSnapshotManagementServer) GetProviderManagementInterface(context.Context, *GetProviderManagementInterfaceRequest) (*GetProviderManagementInterfaceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVSSSnapshotManagementServer) QueryVolumesSupportedForSnapshots(context.Context, *QueryVolumesSupportedForSnapshotsRequest) (*QueryVolumesSupportedForSnapshotsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVSSSnapshotManagementServer) QuerySnapshotsByVolume(context.Context, *QuerySnapshotsByVolumeRequest) (*QuerySnapshotsByVolumeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ VSSSnapshotManagementServer = (*UnimplementedVSSSnapshotManagementServer)(nil)
