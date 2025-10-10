package intmsmediaservices1

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

// INtmsMediaServices1 server interface.
type MediaServices1Server interface {

	// IUnknown base class.
	iunknown.UnknownServer

	MountNTMSMedia(context.Context, *MountNTMSMediaRequest) (*MountNTMSMediaResponse, error)

	DismountNTMSMedia(context.Context, *DismountNTMSMediaRequest) (*DismountNTMSMediaResponse, error)

	// Opnum5NotUsedOnWire operation.
	// Opnum5NotUsedOnWire

	AllocateNTMSMedia(context.Context, *AllocateNTMSMediaRequest) (*AllocateNTMSMediaResponse, error)

	DeallocateNTMSMedia(context.Context, *DeallocateNTMSMediaRequest) (*DeallocateNTMSMediaResponse, error)

	SwapNTMSMedia(context.Context, *SwapNTMSMediaRequest) (*SwapNTMSMediaResponse, error)

	DecommissionNTMSMedia(context.Context, *DecommissionNTMSMediaRequest) (*DecommissionNTMSMediaResponse, error)

	SetNTMSMediaComplete(context.Context, *SetNTMSMediaCompleteRequest) (*SetNTMSMediaCompleteResponse, error)

	DeleteNTMSMedia(context.Context, *DeleteNTMSMediaRequest) (*DeleteNTMSMediaResponse, error)

	CreateNTMSMediaPoolA(context.Context, *CreateNTMSMediaPoolARequest) (*CreateNTMSMediaPoolAResponse, error)

	CreateNTMSMediaPoolW(context.Context, *CreateNTMSMediaPoolWRequest) (*CreateNTMSMediaPoolWResponse, error)

	GetNTMSMediaPoolNameA(context.Context, *GetNTMSMediaPoolNameARequest) (*GetNTMSMediaPoolNameAResponse, error)

	GetNTMSMediaPoolNameW(context.Context, *GetNTMSMediaPoolNameWRequest) (*GetNTMSMediaPoolNameWResponse, error)

	MoveToNTMSMediaPool(context.Context, *MoveToNTMSMediaPoolRequest) (*MoveToNTMSMediaPoolResponse, error)

	DeleteNTMSMediaPool(context.Context, *DeleteNTMSMediaPoolRequest) (*DeleteNTMSMediaPoolResponse, error)

	AddNTMSMediaType(context.Context, *AddNTMSMediaTypeRequest) (*AddNTMSMediaTypeResponse, error)

	DeleteNTMSMediaType(context.Context, *DeleteNTMSMediaTypeRequest) (*DeleteNTMSMediaTypeResponse, error)

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
