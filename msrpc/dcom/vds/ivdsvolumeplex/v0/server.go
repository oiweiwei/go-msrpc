package ivdsvolumeplex

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

// IVdsVolumePlex server interface.
type VolumePlexServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// GetProperties operation.
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)

	// The GetVolume method retrieves the volume that the volume plex belongs to.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetVolume(context.Context, *GetVolumeRequest) (*GetVolumeResponse, error)

	// QueryExtents operation.
	QueryExtents(context.Context, *QueryExtentsRequest) (*QueryExtentsResponse, error)

	// The Repair method repairs a fault-tolerant volume plex by moving defective members
	// to good disks. Only plexes that are RAID-5, striped with parity, can be repaired
	// with this method.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	//
	// ERROR_SUCCESS (0x00000000)
	Repair(context.Context, *RepairRequest) (*RepairResponse, error)
}

func RegisterVolumePlexServer(conn dcerpc.Conn, o VolumePlexServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewVolumePlexServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(VolumePlexSyntaxV0_0))...)
}

func NewVolumePlexServerHandle(o VolumePlexServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return VolumePlexServerHandle(ctx, o, opNum, r)
	}
}

func VolumePlexServerHandle(ctx context.Context, o VolumePlexServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetProperties
		op := &xxx_GetPropertiesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPropertiesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetProperties(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // GetVolume
		op := &xxx_GetVolumeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetVolumeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetVolume(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // QueryExtents
		op := &xxx_QueryExtentsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryExtentsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryExtents(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // Repair
		op := &xxx_RepairOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RepairRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Repair(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IVdsVolumePlex
type UnimplementedVolumePlexServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedVolumePlexServer) GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumePlexServer) GetVolume(context.Context, *GetVolumeRequest) (*GetVolumeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumePlexServer) QueryExtents(context.Context, *QueryExtentsRequest) (*QueryExtentsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumePlexServer) Repair(context.Context, *RepairRequest) (*RepairResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ VolumePlexServer = (*UnimplementedVolumePlexServer)(nil)
