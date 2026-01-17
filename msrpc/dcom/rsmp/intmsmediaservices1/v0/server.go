package intmsmediaservices1

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

// INtmsMediaServices1 server interface.
type MediaServices1Server interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The MountNtmsMedia method mounts one or more pieces of media.
	MountNTMSMedia(context.Context, *MountNTMSMediaRequest) (*MountNTMSMediaResponse, error)

	// The DismountNtmsMedia method queues a command to move a medium in a drive to its
	// storage.
	DismountNTMSMedia(context.Context, *DismountNTMSMediaRequest) (*DismountNTMSMediaResponse, error)

	// Opnum5NotUsedOnWire operation.
	// Opnum5NotUsedOnWire

	// The AllocateNtmsMedia method allocates a piece of available media.
	AllocateNTMSMedia(context.Context, *AllocateNTMSMediaRequest) (*AllocateNTMSMediaResponse, error)

	// The DeallocateNtmsMedia method deallocates the side that is associated with a piece
	// of logical media.
	DeallocateNTMSMedia(context.Context, *DeallocateNTMSMediaRequest) (*DeallocateNTMSMediaResponse, error)

	// The SwapNtmsMedia method swaps the position of two media sides.
	SwapNTMSMedia(context.Context, *SwapNTMSMediaRequest) (*SwapNTMSMediaResponse, error)

	// The DecommissionNtmsMedia method moves media from available state to decommissioned
	// state. Media that are decommissioned by the DecommissionNtmsMedia method are recognized
	// by the server, but decommissioned media does not contain any data and is never again
	// used. Only media that are in an available state can be decommissioned.
	DecommissionNTMSMedia(context.Context, *DecommissionNTMSMediaRequest) (*DecommissionNTMSMediaResponse, error)

	// The SetNtmsMediaComplete method marks a piece of logical media as complete.
	SetNTMSMediaComplete(context.Context, *SetNTMSMediaCompleteRequest) (*SetNTMSMediaCompleteResponse, error)

	// The DeleteNtmsMedia method deletes a physical piece of offline media by removing
	// all references to it.
	DeleteNTMSMedia(context.Context, *DeleteNTMSMediaRequest) (*DeleteNTMSMediaResponse, error)

	// The CreateNtmsMediaPoolA method creates a new application media pool, with strings
	// encoded using ASCII.
	CreateNTMSMediaPoolA(context.Context, *CreateNTMSMediaPoolARequest) (*CreateNTMSMediaPoolAResponse, error)

	// The CreateNtmsMediaPoolW method creates a new application media pool whose name is
	// composed of a sequence of Unicode characters.
	CreateNTMSMediaPoolW(context.Context, *CreateNTMSMediaPoolWRequest) (*CreateNTMSMediaPoolWResponse, error)

	// The GetNtmsMediaPoolNameA method retrieves the full name hierarchy of a media pool,
	// with null-terminated strings encoded using ASCII.
	GetNTMSMediaPoolNameA(context.Context, *GetNTMSMediaPoolNameARequest) (*GetNTMSMediaPoolNameAResponse, error)

	// The GetNtmsMediaPoolNameW method retrieves the full name hierarchy of a media pool,
	// with strings encoded using Unicode.
	GetNTMSMediaPoolNameW(context.Context, *GetNTMSMediaPoolNameWRequest) (*GetNTMSMediaPoolNameWResponse, error)

	// The MoveToNtmsMediaPool method moves a medium from its current media pool to another
	// media pool.
	MoveToNTMSMediaPool(context.Context, *MoveToNTMSMediaPoolRequest) (*MoveToNTMSMediaPoolResponse, error)

	// The DeleteNtmsMediaPool method deletes an application media pool.
	DeleteNTMSMediaPool(context.Context, *DeleteNTMSMediaPoolRequest) (*DeleteNTMSMediaPoolResponse, error)

	// The AddNtmsMediaType method MUST add a media type to a library if there is not currently
	// a relation in the library. The method MUST create the system media pools (FREE, IMPORT,
	// and UNRECOGNIZED) if they do not exist.
	AddNTMSMediaType(context.Context, *AddNTMSMediaTypeRequest) (*AddNTMSMediaTypeResponse, error)

	// The DeleteNtmsMediaType method deletes a media type from a library.
	DeleteNTMSMediaType(context.Context, *DeleteNTMSMediaTypeRequest) (*DeleteNTMSMediaTypeResponse, error)

	// The ChangeNtmsMediaType method moves a physical media identifier to a new media pool
	// and sets the media type of the medium to that of the pool.
	ChangeNTMSMediaType(context.Context, *ChangeNTMSMediaTypeRequest) (*ChangeNTMSMediaTypeResponse, error)
}

func RegisterMediaServices1Server(conn dcerpc.Conn, o MediaServices1Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewMediaServices1ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(MediaServices1SyntaxV0_0))...)
}

func NewMediaServices1ServerHandle(o MediaServices1Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return MediaServices1ServerHandle(ctx, o, opNum, r)
	}
}

func MediaServices1ServerHandle(ctx context.Context, o MediaServices1Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // MountNtmsMedia
		op := &xxx_MountNTMSMediaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MountNTMSMediaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MountNTMSMedia(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // DismountNtmsMedia
		op := &xxx_DismountNTMSMediaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DismountNTMSMediaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DismountNTMSMedia(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // Opnum5NotUsedOnWire
		// Opnum5NotUsedOnWire
		return nil, nil
	case 3: // AllocateNtmsMedia
		op := &xxx_AllocateNTMSMediaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AllocateNTMSMediaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AllocateNTMSMedia(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // DeallocateNtmsMedia
		op := &xxx_DeallocateNTMSMediaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeallocateNTMSMediaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeallocateNTMSMedia(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // SwapNtmsMedia
		op := &xxx_SwapNTMSMediaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SwapNTMSMediaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SwapNTMSMedia(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // DecommissionNtmsMedia
		op := &xxx_DecommissionNTMSMediaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DecommissionNTMSMediaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DecommissionNTMSMedia(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // SetNtmsMediaComplete
		op := &xxx_SetNTMSMediaCompleteOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetNTMSMediaCompleteRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetNTMSMediaComplete(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // DeleteNtmsMedia
		op := &xxx_DeleteNTMSMediaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteNTMSMediaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteNTMSMedia(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // CreateNtmsMediaPoolA
		op := &xxx_CreateNTMSMediaPoolAOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateNTMSMediaPoolARequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateNTMSMediaPoolA(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // CreateNtmsMediaPoolW
		op := &xxx_CreateNTMSMediaPoolWOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateNTMSMediaPoolWRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateNTMSMediaPoolW(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // GetNtmsMediaPoolNameA
		op := &xxx_GetNTMSMediaPoolNameAOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNTMSMediaPoolNameARequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNTMSMediaPoolNameA(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // GetNtmsMediaPoolNameW
		op := &xxx_GetNTMSMediaPoolNameWOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNTMSMediaPoolNameWRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNTMSMediaPoolNameW(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // MoveToNtmsMediaPool
		op := &xxx_MoveToNTMSMediaPoolOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MoveToNTMSMediaPoolRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MoveToNTMSMediaPool(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // DeleteNtmsMediaPool
		op := &xxx_DeleteNTMSMediaPoolOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteNTMSMediaPoolRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteNTMSMediaPool(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // AddNtmsMediaType
		op := &xxx_AddNTMSMediaTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddNTMSMediaTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddNTMSMediaType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // DeleteNtmsMediaType
		op := &xxx_DeleteNTMSMediaTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteNTMSMediaTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteNTMSMediaType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // ChangeNtmsMediaType
		op := &xxx_ChangeNTMSMediaTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ChangeNTMSMediaTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ChangeNTMSMediaType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented INtmsMediaServices1
type UnimplementedMediaServices1Server struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedMediaServices1Server) MountNTMSMedia(context.Context, *MountNTMSMediaRequest) (*MountNTMSMediaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMediaServices1Server) DismountNTMSMedia(context.Context, *DismountNTMSMediaRequest) (*DismountNTMSMediaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMediaServices1Server) AllocateNTMSMedia(context.Context, *AllocateNTMSMediaRequest) (*AllocateNTMSMediaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMediaServices1Server) DeallocateNTMSMedia(context.Context, *DeallocateNTMSMediaRequest) (*DeallocateNTMSMediaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMediaServices1Server) SwapNTMSMedia(context.Context, *SwapNTMSMediaRequest) (*SwapNTMSMediaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMediaServices1Server) DecommissionNTMSMedia(context.Context, *DecommissionNTMSMediaRequest) (*DecommissionNTMSMediaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMediaServices1Server) SetNTMSMediaComplete(context.Context, *SetNTMSMediaCompleteRequest) (*SetNTMSMediaCompleteResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMediaServices1Server) DeleteNTMSMedia(context.Context, *DeleteNTMSMediaRequest) (*DeleteNTMSMediaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMediaServices1Server) CreateNTMSMediaPoolA(context.Context, *CreateNTMSMediaPoolARequest) (*CreateNTMSMediaPoolAResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMediaServices1Server) CreateNTMSMediaPoolW(context.Context, *CreateNTMSMediaPoolWRequest) (*CreateNTMSMediaPoolWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMediaServices1Server) GetNTMSMediaPoolNameA(context.Context, *GetNTMSMediaPoolNameARequest) (*GetNTMSMediaPoolNameAResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMediaServices1Server) GetNTMSMediaPoolNameW(context.Context, *GetNTMSMediaPoolNameWRequest) (*GetNTMSMediaPoolNameWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMediaServices1Server) MoveToNTMSMediaPool(context.Context, *MoveToNTMSMediaPoolRequest) (*MoveToNTMSMediaPoolResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMediaServices1Server) DeleteNTMSMediaPool(context.Context, *DeleteNTMSMediaPoolRequest) (*DeleteNTMSMediaPoolResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMediaServices1Server) AddNTMSMediaType(context.Context, *AddNTMSMediaTypeRequest) (*AddNTMSMediaTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMediaServices1Server) DeleteNTMSMediaType(context.Context, *DeleteNTMSMediaTypeRequest) (*DeleteNTMSMediaTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMediaServices1Server) ChangeNTMSMediaType(context.Context, *ChangeNTMSMediaTypeRequest) (*ChangeNTMSMediaTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ MediaServices1Server = (*UnimplementedMediaServices1Server)(nil)
