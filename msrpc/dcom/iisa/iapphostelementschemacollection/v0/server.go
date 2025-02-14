package iapphostelementschemacollection

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

// IAppHostElementSchemaCollection server interface.
type AppHostElementSchemaCollectionServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// Count operation.
	GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error)

	// Item operation.
	GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error)
}

func RegisterAppHostElementSchemaCollectionServer(conn dcerpc.Conn, o AppHostElementSchemaCollectionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostElementSchemaCollectionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostElementSchemaCollectionSyntaxV0_0))...)
}

func NewAppHostElementSchemaCollectionServerHandle(o AppHostElementSchemaCollectionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostElementSchemaCollectionServerHandle(ctx, o, opNum, r)
	}
}

func AppHostElementSchemaCollectionServerHandle(ctx context.Context, o AppHostElementSchemaCollectionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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

// Unimplemented IAppHostElementSchemaCollection
type UnimplementedAppHostElementSchemaCollectionServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedAppHostElementSchemaCollectionServer) GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostElementSchemaCollectionServer) GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AppHostElementSchemaCollectionServer = (*UnimplementedAppHostElementSchemaCollectionServer)(nil)
