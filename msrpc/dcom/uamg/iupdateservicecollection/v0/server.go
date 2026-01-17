package iupdateservicecollection

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

// IUpdateServiceCollection server interface.
type UpdateServiceCollectionServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The IStringCollection::Item (opnum 9) method sets an item in the collection.
	//
	// The IWindowsDriverUpdateEntryCollection::Item (opnum 8) method retrieves an item
	// by index.
	//
	// );
	//
	// The IUpdateDownloadContentCollection::Item (opnum 8) method retrieves an item from
	// the collection by index.
	//
	// The ICategoryCollection::Item (opnum 8) method retrieves an item by index.
	//
	// The IUpdateExceptionCollection::Item (opnum 8) method retrieves an item from the
	// collection by index.
	//
	// The IUpdateHistoryEntryCollection::Item (opnum 8) method retrieves an item from the
	// collection by index.
	//
	// The IUpdateCollection::Item (opnum 9) method replaces an item in the collection by
	// index.
	//
	// The IStringCollection::Item (opnum 8) method gets an item by index.
	//
	// The IUpdateCollection::Item (opnum 8) method retrieves an item from the collection
	// by index.
	//
	// The IUpdateServiceCollection::Item (opnum 8) method retrieves an item from the collection
	// by index.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// If the index is less than 0x00000000 or greater than or equal to the number of items
	// in the collection, the server SHOULD return WU_E_INVALIDINDEX (0x80240007).
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD set the item at the given index in the List ADM element list to
	// the given value.
	GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error)

	// _NewEnum operation.
	Get_NewEnum(context.Context, *Get_NewEnumRequest) (*Get_NewEnumResponse, error)

	// The IUpdateExceptionCollection::Count (opnum 10) method retrieves the number of items
	// in the collection.
	//
	// The IUpdateHistoryEntryCollection::Count (opnum 10) method retrieves the number of
	// items in the collection.
	//
	// The IUpdateDownloadContentCollection::Count (opnum 10) method retrieves the number
	// of items in the collection.
	//
	// The IUpdateCollection::Count (opnum 11) method returns the number of items in the
	// collection.
	//
	// The ICategoryCollection::Count (opnum 10) method retrieves a count of the number
	// of items in the collection of ICategory instances.
	//
	// The IWindowsDriverUpdateEntryCollection::Count (opnum 10) method retrieves a count
	// of the number of items in the collection of IWindowsDriverUpdateEntry instances.
	//
	// The IUpdateServiceCollection::Count (opnum 10) method retrieves the number of items
	// in the collection.
	//
	// The IStringCollection::Count (opnum 11) method retrieves a count of the items in
	// the collection.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the Count ADM element.
	GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error)
}

func RegisterUpdateServiceCollectionServer(conn dcerpc.Conn, o UpdateServiceCollectionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewUpdateServiceCollectionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(UpdateServiceCollectionSyntaxV0_0))...)
}

func NewUpdateServiceCollectionServerHandle(o UpdateServiceCollectionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return UpdateServiceCollectionServerHandle(ctx, o, opNum, r)
	}
}

func UpdateServiceCollectionServerHandle(ctx context.Context, o UpdateServiceCollectionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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

// Unimplemented IUpdateServiceCollection
type UnimplementedUpdateServiceCollectionServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedUpdateServiceCollectionServer) GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServiceCollectionServer) Get_NewEnum(context.Context, *Get_NewEnumRequest) (*Get_NewEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServiceCollectionServer) GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ UpdateServiceCollectionServer = (*UnimplementedUpdateServiceCollectionServer)(nil)
