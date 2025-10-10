package iwindowsdriverupdateentrycollection

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

// IWindowsDriverUpdateEntryCollection server interface.
type WindowsDriverUpdateEntryCollectionServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error)

	// _NewEnum operation.
	Get_NewEnum(context.Context, *Get_NewEnumRequest) (*Get_NewEnumResponse, error)

	GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error)
}

func RegisterWindowsDriverUpdateEntryCollectionServer(conn dcerpc.Conn, o WindowsDriverUpdateEntryCollectionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewWindowsDriverUpdateEntryCollectionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(WindowsDriverUpdateEntryCollectionSyntaxV0_0))...)
}

func NewWindowsDriverUpdateEntryCollectionServerHandle(o WindowsDriverUpdateEntryCollectionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return WindowsDriverUpdateEntryCollectionServerHandle(ctx, o, opNum, r)
	}
}

func WindowsDriverUpdateEntryCollectionServerHandle(ctx context.Context, o WindowsDriverUpdateEntryCollectionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
	case 8: // _NewEnum
		op := &xxx_Get_NewEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &Get_NewEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Get_NewEnum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // Count
		op := &xxx_GetCountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IWindowsDriverUpdateEntryCollection
type UnimplementedWindowsDriverUpdateEntryCollectionServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedWindowsDriverUpdateEntryCollectionServer) GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWindowsDriverUpdateEntryCollectionServer) Get_NewEnum(context.Context, *Get_NewEnumRequest) (*Get_NewEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWindowsDriverUpdateEntryCollectionServer) GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ WindowsDriverUpdateEntryCollectionServer = (*UnimplementedWindowsDriverUpdateEntryCollectionServer)(nil)
