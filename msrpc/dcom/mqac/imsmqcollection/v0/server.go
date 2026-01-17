package imsmqcollection

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = idispatch.GoPackage
)

// IMSMQCollection server interface.
type CollectionServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The Item method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return a VARIANT from the VariantCollection instance variable where
	// the key matches the Index input parameter.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	Item(context.Context, *ItemRequest) (*ItemResponse, error)

	// The Count method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return the total number of elements in the VariantCollection instance
	// variable.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error)

	// _NewEnum operation.
	_NewEnum(context.Context, *_NewEnumRequest) (*_NewEnumResponse, error)
}

func RegisterCollectionServer(conn dcerpc.Conn, o CollectionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewCollectionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(CollectionSyntaxV0_0))...)
}

func NewCollectionServerHandle(o CollectionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return CollectionServerHandle(ctx, o, opNum, r)
	}
}

func CollectionServerHandle(ctx context.Context, o CollectionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
type UnimplementedCollectionServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedCollectionServer) Item(context.Context, *ItemRequest) (*ItemResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCollectionServer) GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCollectionServer) _NewEnum(context.Context, *_NewEnumRequest) (*_NewEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ CollectionServer = (*UnimplementedCollectionServer)(nil)
