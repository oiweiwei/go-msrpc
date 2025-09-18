package ipersistmoniker

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

// IPersistMoniker server interface.
type PersistMonikerServer interface {

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

	// SaveCompleted operation.
	SaveCompleted(context.Context, *SaveCompletedRequest) (*SaveCompletedResponse, error)

	// GetCurMoniker operation.
	GetCurMoniker(context.Context, *GetCurMonikerRequest) (*GetCurMonikerResponse, error)
}

func RegisterPersistMonikerServer(conn dcerpc.Conn, o PersistMonikerServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewPersistMonikerServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(PersistMonikerSyntaxV0_0))...)
}

func NewPersistMonikerServerHandle(o PersistMonikerServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return PersistMonikerServerHandle(ctx, o, opNum, r)
	}
}

func PersistMonikerServerHandle(ctx context.Context, o PersistMonikerServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
	case 7: // SaveCompleted
		op := &xxx_SaveCompletedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SaveCompletedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SaveCompleted(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // GetCurMoniker
		op := &xxx_GetCurMonikerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCurMonikerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCurMoniker(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IPersistMoniker
type UnimplementedPersistMonikerServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedPersistMonikerServer) GetClassID(context.Context, *GetClassIDRequest) (*GetClassIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPersistMonikerServer) IsDirty(context.Context, *IsDirtyRequest) (*IsDirtyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPersistMonikerServer) Load(context.Context, *LoadRequest) (*LoadResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPersistMonikerServer) Save(context.Context, *SaveRequest) (*SaveResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPersistMonikerServer) SaveCompleted(context.Context, *SaveCompletedRequest) (*SaveCompletedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPersistMonikerServer) GetCurMoniker(context.Context, *GetCurMonikerRequest) (*GetCurMonikerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ PersistMonikerServer = (*UnimplementedPersistMonikerServer)(nil)
