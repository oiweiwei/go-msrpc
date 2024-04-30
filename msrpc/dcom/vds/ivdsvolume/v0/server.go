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
		in := &GetPropertiesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetProperties(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // GetPack
		in := &GetPackRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPack(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // QueryPlexes
		in := &QueryPlexesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryPlexes(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // Extend
		in := &ExtendRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Extend(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // Shrink
		in := &ShrinkRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Shrink(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // AddPlex
		in := &AddPlexRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddPlex(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // BreakPlex
		in := &BreakPlexRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.BreakPlex(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // RemovePlex
		in := &RemovePlexRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemovePlex(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // Delete
		in := &DeleteRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Delete(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // SetFlags
		in := &SetFlagsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetFlags(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // ClearFlags
		in := &ClearFlagsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ClearFlags(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
