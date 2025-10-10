package istringcollection

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

// IStringCollection server interface.
type StringCollectionServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error)

	SetItem(context.Context, *SetItemRequest) (*SetItemResponse, error)

	// _NewEnum operation.
	Get_NewEnum(context.Context, *Get_NewEnumRequest) (*Get_NewEnumResponse, error)

	GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error)

	GetReadOnly(context.Context, *GetReadOnlyRequest) (*GetReadOnlyResponse, error)

	Add(context.Context, *AddRequest) (*AddResponse, error)

	Clear(context.Context, *ClearRequest) (*ClearResponse, error)

	Copy(context.Context, *CopyRequest) (*CopyResponse, error)

	Insert(context.Context, *InsertRequest) (*InsertResponse, error)

	RemoveAT(context.Context, *RemoveATRequest) (*RemoveATResponse, error)
}

func RegisterStringCollectionServer(conn dcerpc.Conn, o StringCollectionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewStringCollectionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(StringCollectionSyntaxV0_0))...)
}

func NewStringCollectionServerHandle(o StringCollectionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return StringCollectionServerHandle(ctx, o, opNum, r)
	}
}

func StringCollectionServerHandle(ctx context.Context, o StringCollectionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Item
		op := &xxx_GetItemOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetItemRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetItem(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // Item
		op := &xxx_SetItemOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetItemRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetItem(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // _NewEnum
		op := &xxx_Get_NewEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &Get_NewEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Get_NewEnum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // Count
		op := &xxx_GetCountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // ReadOnly
		op := &xxx_GetReadOnlyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetReadOnlyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetReadOnly(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // Add
		op := &xxx_AddOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Add(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // Clear
		op := &xxx_ClearOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ClearRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Clear(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // Copy
		op := &xxx_CopyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CopyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Copy(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // Insert
		op := &xxx_InsertOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &InsertRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Insert(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // RemoveAt
		op := &xxx_RemoveATOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveATRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveAT(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IStringCollection
type UnimplementedStringCollectionServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedStringCollectionServer) GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedStringCollectionServer) SetItem(context.Context, *SetItemRequest) (*SetItemResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedStringCollectionServer) Get_NewEnum(context.Context, *Get_NewEnumRequest) (*Get_NewEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedStringCollectionServer) GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedStringCollectionServer) GetReadOnly(context.Context, *GetReadOnlyRequest) (*GetReadOnlyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedStringCollectionServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedStringCollectionServer) Clear(context.Context, *ClearRequest) (*ClearResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedStringCollectionServer) Copy(context.Context, *CopyRequest) (*CopyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedStringCollectionServer) Insert(context.Context, *InsertRequest) (*InsertResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedStringCollectionServer) RemoveAT(context.Context, *RemoveATRequest) (*RemoveATResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ StringCollectionServer = (*UnimplementedStringCollectionServer)(nil)
