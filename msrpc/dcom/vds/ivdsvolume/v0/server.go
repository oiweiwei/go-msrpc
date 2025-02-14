package ivdsvolume

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

// IVdsVolume server interface.
type VolumeServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// GetProperties operation.
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)

	// GetPack operation.
	GetPack(context.Context, *GetPackRequest) (*GetPackResponse, error)

	// The QueryPlexes method enumerates the plexes of a volume.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryPlexes(context.Context, *QueryPlexesRequest) (*QueryPlexesResponse, error)

	// The Extend method expands the size of the current volume by adding disk extents to
	// each member of each plex.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	Extend(context.Context, *ExtendRequest) (*ExtendResponse, error)

	// Shrink operation.
	Shrink(context.Context, *ShrinkRequest) (*ShrinkResponse, error)

	// The AddPlex method adds a volume as a plex to the current volume.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	AddPlex(context.Context, *AddPlexRequest) (*AddPlexResponse, error)

	// The BreakPlex method removes a specified plex from the current volume. The interface
	// pointer for the new volume object can be retrieved by calling IVdsAsync::Wait through
	// the ppAsync parameter. The VDS_ASYNC_OUTPUT structure that is returned contains the
	// volume object interface pointer in the bvp.pVolumeUnk member.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	BreakPlex(context.Context, *BreakPlexRequest) (*BreakPlexResponse, error)

	// The RemovePlex method removes a specified plex from a volume. The last plex of a
	// volume cannot be removed.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	RemovePlex(context.Context, *RemovePlexRequest) (*RemovePlexResponse, error)

	// The Delete method deletes all plexes in a volume.<118>
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	//
	// ERROR_SUCCESS (0x00000000)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)

	// SetFlags operation.
	SetFlags(context.Context, *SetFlagsRequest) (*SetFlagsResponse, error)

	// ClearFlags operation.
	ClearFlags(context.Context, *ClearFlagsRequest) (*ClearFlagsResponse, error)
}

func RegisterVolumeServer(conn dcerpc.Conn, o VolumeServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewVolumeServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(VolumeSyntaxV0_0))...)
}

func NewVolumeServerHandle(o VolumeServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return VolumeServerHandle(ctx, o, opNum, r)
	}
}

func VolumeServerHandle(ctx context.Context, o VolumeServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
	case 4: // GetPack
		op := &xxx_GetPackOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPackRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPack(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // QueryPlexes
		op := &xxx_QueryPlexesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryPlexesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryPlexes(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // Extend
		op := &xxx_ExtendOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ExtendRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Extend(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // Shrink
		op := &xxx_ShrinkOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ShrinkRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Shrink(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // AddPlex
		op := &xxx_AddPlexOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddPlexRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddPlex(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // BreakPlex
		op := &xxx_BreakPlexOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BreakPlexRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BreakPlex(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // RemovePlex
		op := &xxx_RemovePlexOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemovePlexRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemovePlex(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // Delete
		op := &xxx_DeleteOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Delete(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // SetFlags
		op := &xxx_SetFlagsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFlagsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFlags(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // ClearFlags
		op := &xxx_ClearFlagsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ClearFlagsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ClearFlags(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IVdsVolume
type UnimplementedVolumeServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedVolumeServer) GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeServer) GetPack(context.Context, *GetPackRequest) (*GetPackResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeServer) QueryPlexes(context.Context, *QueryPlexesRequest) (*QueryPlexesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeServer) Extend(context.Context, *ExtendRequest) (*ExtendResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeServer) Shrink(context.Context, *ShrinkRequest) (*ShrinkResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeServer) AddPlex(context.Context, *AddPlexRequest) (*AddPlexResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeServer) BreakPlex(context.Context, *BreakPlexRequest) (*BreakPlexResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeServer) RemovePlex(context.Context, *RemovePlexRequest) (*RemovePlexResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeServer) SetFlags(context.Context, *SetFlagsRequest) (*SetFlagsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVolumeServer) ClearFlags(context.Context, *ClearFlagsRequest) (*ClearFlagsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ VolumeServer = (*UnimplementedVolumeServer)(nil)
