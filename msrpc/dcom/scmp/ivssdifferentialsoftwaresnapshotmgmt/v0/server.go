package ivssdifferentialsoftwaresnapshotmgmt

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

// IVssDifferentialSoftwareSnapshotMgmt server interface.
type DifferentialSoftwareSnapshotManagementServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	AddDiffArea(context.Context, *AddDiffAreaRequest) (*AddDiffAreaResponse, error)

	ChangeDiffAreaMaximumSize(context.Context, *ChangeDiffAreaMaximumSizeRequest) (*ChangeDiffAreaMaximumSizeResponse, error)

	QueryVolumesSupportedForDiffAreas(context.Context, *QueryVolumesSupportedForDiffAreasRequest) (*QueryVolumesSupportedForDiffAreasResponse, error)

	QueryDiffAreasForVolume(context.Context, *QueryDiffAreasForVolumeRequest) (*QueryDiffAreasForVolumeResponse, error)

	QueryDiffAreasOnVolume(context.Context, *QueryDiffAreasOnVolumeRequest) (*QueryDiffAreasOnVolumeResponse, error)

	// Opnum08NotUsedOnWire operation.
	// Opnum08NotUsedOnWire
}

func RegisterDifferentialSoftwareSnapshotManagementServer(conn dcerpc.Conn, o DifferentialSoftwareSnapshotManagementServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewDifferentialSoftwareSnapshotManagementServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(DifferentialSoftwareSnapshotManagementSyntaxV0_0))...)
}

func NewDifferentialSoftwareSnapshotManagementServerHandle(o DifferentialSoftwareSnapshotManagementServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return DifferentialSoftwareSnapshotManagementServerHandle(ctx, o, opNum, r)
	}
}

func DifferentialSoftwareSnapshotManagementServerHandle(ctx context.Context, o DifferentialSoftwareSnapshotManagementServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // AddDiffArea
		op := &xxx_AddDiffAreaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddDiffAreaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddDiffArea(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // ChangeDiffAreaMaximumSize
		op := &xxx_ChangeDiffAreaMaximumSizeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ChangeDiffAreaMaximumSizeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ChangeDiffAreaMaximumSize(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // QueryVolumesSupportedForDiffAreas
		op := &xxx_QueryVolumesSupportedForDiffAreasOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryVolumesSupportedForDiffAreasRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryVolumesSupportedForDiffAreas(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // QueryDiffAreasForVolume
		op := &xxx_QueryDiffAreasForVolumeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryDiffAreasForVolumeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryDiffAreasForVolume(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // QueryDiffAreasOnVolume
		op := &xxx_QueryDiffAreasOnVolumeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryDiffAreasOnVolumeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryDiffAreasOnVolume(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // Opnum08NotUsedOnWire
		// Opnum08NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}

// Unimplemented IVssDifferentialSoftwareSnapshotMgmt
type UnimplementedDifferentialSoftwareSnapshotManagementServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedDifferentialSoftwareSnapshotManagementServer) AddDiffArea(context.Context, *AddDiffAreaRequest) (*AddDiffAreaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDifferentialSoftwareSnapshotManagementServer) ChangeDiffAreaMaximumSize(context.Context, *ChangeDiffAreaMaximumSizeRequest) (*ChangeDiffAreaMaximumSizeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDifferentialSoftwareSnapshotManagementServer) QueryVolumesSupportedForDiffAreas(context.Context, *QueryVolumesSupportedForDiffAreasRequest) (*QueryVolumesSupportedForDiffAreasResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDifferentialSoftwareSnapshotManagementServer) QueryDiffAreasForVolume(context.Context, *QueryDiffAreasForVolumeRequest) (*QueryDiffAreasForVolumeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDifferentialSoftwareSnapshotManagementServer) QueryDiffAreasOnVolume(context.Context, *QueryDiffAreasOnVolumeRequest) (*QueryDiffAreasOnVolumeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ DifferentialSoftwareSnapshotManagementServer = (*UnimplementedDifferentialSoftwareSnapshotManagementServer)(nil)
