package iapphostconstantvaluecollection

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

// IAppHostConstantValueCollection server interface.
type AppHostConstantValueCollectionServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// Count operation.
	GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error)

	// Item operation.
	GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error)
}

func RegisterAppHostConstantValueCollectionServer(conn dcerpc.Conn, o AppHostConstantValueCollectionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostConstantValueCollectionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostConstantValueCollectionSyntaxV0_0))...)
}

func NewAppHostConstantValueCollectionServerHandle(o AppHostConstantValueCollectionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostConstantValueCollectionServerHandle(ctx, o, opNum, r)
	}
}

func AppHostConstantValueCollectionServerHandle(ctx context.Context, o AppHostConstantValueCollectionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
	}
	return nil, nil
}

// Unimplemented IAppHostConstantValueCollection
type UnimplementedAppHostConstantValueCollectionServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedAppHostConstantValueCollectionServer) GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostConstantValueCollectionServer) GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AppHostConstantValueCollectionServer = (*UnimplementedAppHostConstantValueCollectionServer)(nil)
