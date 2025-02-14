package iapphostconfiglocationcollection

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

// IAppHostConfigLocationCollection server interface.
type AppHostConfigLocationCollectionServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// Count operation.
	GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error)

	// Item operation.
	GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error)

	// AddLocation operation.
	AddLocation(context.Context, *AddLocationRequest) (*AddLocationResponse, error)

	// DeleteLocation operation.
	DeleteLocation(context.Context, *DeleteLocationRequest) (*DeleteLocationResponse, error)
}

func RegisterAppHostConfigLocationCollectionServer(conn dcerpc.Conn, o AppHostConfigLocationCollectionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostConfigLocationCollectionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostConfigLocationCollectionSyntaxV0_0))...)
}

func NewAppHostConfigLocationCollectionServerHandle(o AppHostConfigLocationCollectionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostConfigLocationCollectionServerHandle(ctx, o, opNum, r)
	}
}

func AppHostConfigLocationCollectionServerHandle(ctx context.Context, o AppHostConfigLocationCollectionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Count
		op := &xxx_GetCountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // Item
		op := &xxx_GetItemOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetItemRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetItem(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // AddLocation
		op := &xxx_AddLocationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddLocationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddLocation(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // DeleteLocation
		op := &xxx_DeleteLocationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteLocationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteLocation(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IAppHostConfigLocationCollection
type UnimplementedAppHostConfigLocationCollectionServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedAppHostConfigLocationCollectionServer) GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostConfigLocationCollectionServer) GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostConfigLocationCollectionServer) AddLocation(context.Context, *AddLocationRequest) (*AddLocationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostConfigLocationCollectionServer) DeleteLocation(context.Context, *DeleteLocationRequest) (*DeleteLocationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AppHostConfigLocationCollectionServer = (*UnimplementedAppHostConfigLocationCollectionServer)(nil)
