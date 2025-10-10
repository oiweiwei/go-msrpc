package imsmqcollection

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
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
	_ = idispatch.GoPackage
)

// IMSMQCollection server interface.
type ImsmqCollectionServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// Item operation.
	Item(context.Context, *ItemRequest) (*ItemResponse, error)

	// Count operation.
	GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error)

	// _NewEnum operation.
	_NewEnum(context.Context, *_NewEnumRequest) (*_NewEnumResponse, error)
}

func RegisterImsmqCollectionServer(conn dcerpc.Conn, o ImsmqCollectionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewImsmqCollectionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ImsmqCollectionSyntaxV0_0))...)
}

func NewImsmqCollectionServerHandle(o ImsmqCollectionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ImsmqCollectionServerHandle(ctx, o, opNum, r)
	}
}

func ImsmqCollectionServerHandle(ctx context.Context, o ImsmqCollectionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Item
		op := &xxx_ItemOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ItemRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Item(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // Count
		op := &xxx_GetCountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // _NewEnum
		op := &xxx__NewEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &_NewEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o._NewEnum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMSMQCollection
type UnimplementedImsmqCollectionServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedImsmqCollectionServer) Item(context.Context, *ItemRequest) (*ItemResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqCollectionServer) GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqCollectionServer) _NewEnum(context.Context, *_NewEnumRequest) (*_NewEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ImsmqCollectionServer = (*UnimplementedImsmqCollectionServer)(nil)
