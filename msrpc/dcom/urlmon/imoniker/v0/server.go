package imoniker

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

// IMoniker server interface.
type MonikerServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// GetClassID operation.
	GetClassID(context.Context, *GetClassIDRequest) (*GetClassIDResponse, error)

	// IsDirty operation.
	IsDirty(context.Context, *IsDirtyRequest) (*IsDirtyResponse, error)

	// Load operation.
	Load(context.Context, *LoadRequest) (*LoadResponse, error)

	// Save operation.
	Save(context.Context, *SaveRequest) (*SaveResponse, error)

	// GetSizeMax operation.
	GetSizeMax(context.Context, *GetSizeMaxRequest) (*GetSizeMaxResponse, error)

	// BindToObject operation.
	BindToObject(context.Context, *BindToObjectRequest) (*BindToObjectResponse, error)

	// BindToStorage operation.
	BindToStorage(context.Context, *BindToStorageRequest) (*BindToStorageResponse, error)

	// Reduce operation.
	Reduce(context.Context, *ReduceRequest) (*ReduceResponse, error)

	// ComposeWith operation.
	ComposeWith(context.Context, *ComposeWithRequest) (*ComposeWithResponse, error)

	// Enum operation.
	Enum(context.Context, *EnumRequest) (*EnumResponse, error)

	// IsEqual operation.
	IsEqual(context.Context, *IsEqualRequest) (*IsEqualResponse, error)

	// Hash operation.
	Hash(context.Context, *HashRequest) (*HashResponse, error)

	// IsRunning operation.
	IsRunning(context.Context, *IsRunningRequest) (*IsRunningResponse, error)

	// GetTimeOfLastChange operation.
	GetTimeOfLastChange(context.Context, *GetTimeOfLastChangeRequest) (*GetTimeOfLastChangeResponse, error)

	// Inverse operation.
	Inverse(context.Context, *InverseRequest) (*InverseResponse, error)

	// CommonPrefixWith operation.
	CommonPrefixWith(context.Context, *CommonPrefixWithRequest) (*CommonPrefixWithResponse, error)

	// RelativePathTo operation.
	RelativePathTo(context.Context, *RelativePathToRequest) (*RelativePathToResponse, error)

	// GetDisplayName operation.
	GetDisplayName(context.Context, *GetDisplayNameRequest) (*GetDisplayNameResponse, error)

	// ParseDisplayName operation.
	ParseDisplayName(context.Context, *ParseDisplayNameRequest) (*ParseDisplayNameResponse, error)

	// IsSystemMoniker operation.
	IsSystemMoniker(context.Context, *IsSystemMonikerRequest) (*IsSystemMonikerResponse, error)
}

func RegisterMonikerServer(conn dcerpc.Conn, o MonikerServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewMonikerServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(MonikerSyntaxV0_0))...)
}

func NewMonikerServerHandle(o MonikerServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return MonikerServerHandle(ctx, o, opNum, r)
	}
}

func MonikerServerHandle(ctx context.Context, o MonikerServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetClassID
		op := &xxx_GetClassIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetClassIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetClassID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // IsDirty
		op := &xxx_IsDirtyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IsDirtyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.IsDirty(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // Load
		op := &xxx_LoadOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &LoadRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Load(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // Save
		op := &xxx_SaveOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SaveRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Save(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // GetSizeMax
		op := &xxx_GetSizeMaxOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSizeMaxRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSizeMax(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // BindToObject
		op := &xxx_BindToObjectOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BindToObjectRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BindToObject(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // BindToStorage
		op := &xxx_BindToStorageOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BindToStorageRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BindToStorage(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // Reduce
		op := &xxx_ReduceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReduceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Reduce(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // ComposeWith
		op := &xxx_ComposeWithOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ComposeWithRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ComposeWith(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // Enum
		op := &xxx_EnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Enum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // IsEqual
		op := &xxx_IsEqualOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IsEqualRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.IsEqual(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // Hash
		op := &xxx_HashOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &HashRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Hash(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // IsRunning
		op := &xxx_IsRunningOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IsRunningRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.IsRunning(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // GetTimeOfLastChange
		op := &xxx_GetTimeOfLastChangeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTimeOfLastChangeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTimeOfLastChange(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // Inverse
		op := &xxx_InverseOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &InverseRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Inverse(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // CommonPrefixWith
		op := &xxx_CommonPrefixWithOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CommonPrefixWithRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CommonPrefixWith(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // RelativePathTo
		op := &xxx_RelativePathToOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RelativePathToRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RelativePathTo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // GetDisplayName
		op := &xxx_GetDisplayNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDisplayNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDisplayName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // ParseDisplayName
		op := &xxx_ParseDisplayNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ParseDisplayNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ParseDisplayName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // IsSystemMoniker
		op := &xxx_IsSystemMonikerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IsSystemMonikerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.IsSystemMoniker(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMoniker
type UnimplementedMonikerServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedMonikerServer) GetClassID(context.Context, *GetClassIDRequest) (*GetClassIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMonikerServer) IsDirty(context.Context, *IsDirtyRequest) (*IsDirtyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMonikerServer) Load(context.Context, *LoadRequest) (*LoadResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMonikerServer) Save(context.Context, *SaveRequest) (*SaveResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMonikerServer) GetSizeMax(context.Context, *GetSizeMaxRequest) (*GetSizeMaxResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMonikerServer) BindToObject(context.Context, *BindToObjectRequest) (*BindToObjectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMonikerServer) BindToStorage(context.Context, *BindToStorageRequest) (*BindToStorageResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMonikerServer) Reduce(context.Context, *ReduceRequest) (*ReduceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMonikerServer) ComposeWith(context.Context, *ComposeWithRequest) (*ComposeWithResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMonikerServer) Enum(context.Context, *EnumRequest) (*EnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMonikerServer) IsEqual(context.Context, *IsEqualRequest) (*IsEqualResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMonikerServer) Hash(context.Context, *HashRequest) (*HashResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMonikerServer) IsRunning(context.Context, *IsRunningRequest) (*IsRunningResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMonikerServer) GetTimeOfLastChange(context.Context, *GetTimeOfLastChangeRequest) (*GetTimeOfLastChangeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMonikerServer) Inverse(context.Context, *InverseRequest) (*InverseResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMonikerServer) CommonPrefixWith(context.Context, *CommonPrefixWithRequest) (*CommonPrefixWithResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMonikerServer) RelativePathTo(context.Context, *RelativePathToRequest) (*RelativePathToResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMonikerServer) GetDisplayName(context.Context, *GetDisplayNameRequest) (*GetDisplayNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMonikerServer) ParseDisplayName(context.Context, *ParseDisplayNameRequest) (*ParseDisplayNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMonikerServer) IsSystemMoniker(context.Context, *IsSystemMonikerRequest) (*IsSystemMonikerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ MonikerServer = (*UnimplementedMonikerServer)(nil)
