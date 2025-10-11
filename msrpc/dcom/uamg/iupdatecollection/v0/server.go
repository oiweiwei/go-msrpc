package iupdatecollection

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

// IUpdateCollection server interface.
type UpdateCollectionServer interface {

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
	SetItem(context.Context, *SetItemRequest) (*SetItemResponse, error)

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

	// The IUpdateSession::ReadOnly (opnum 10) method returns whether the session is read-only.
	//
	// The IUpdateCollection::ReadOnly (opnum 12) method returns whether the collection
	// is read-only.
	//
	// The IStringCollection::ReadOnly (opnum 12) method retrieves whether the collection
	// is read-only.
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
	// This method SHOULD return the value of the ReadOnly ADM element.
	GetReadOnly(context.Context, *GetReadOnlyRequest) (*GetReadOnlyResponse, error)

	// The IStringCollection::Add (opnum 13) method adds an item to the collection.
	//
	// The IUpdateCollection::Add (opnum 13) method adds an item to the collection.
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
	// This method SHOULD append the given value to the List ADM element and increment the
	// value of the Count ADM element.
	Add(context.Context, *AddRequest) (*AddResponse, error)

	// The IStringCollection::Clear (opnum 14) method removes all items from the collection.
	//
	// This method has no parameters.
	//
	// The IUpdateCollection::Clear (opnum 14) method empties the collection.
	//
	// This method has no parameters.
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
	// This method SHOULD empty the List ADM element and set the value of the Count ADM
	// element to 0.
	Clear(context.Context, *ClearRequest) (*ClearResponse, error)

	// Opnum15NotUsedOnWire operation.
	// Opnum15NotUsedOnWire

	// The IStringCollection::Insert (opnum 16) method inserts an item into the collection
	// at a given position.
	//
	// The IUpdateCollection::Insert (opnum 16) method inserts an item into the collection.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// If the collection is read-only, the server SHOULD return WU_E_NOT_SUPPORTED (0x80240037).
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD insert the given value into its List ADM element at the given
	// index, and increment the value of its Count ADM element.
	Insert(context.Context, *InsertRequest) (*InsertResponse, error)

	// The IStringCollection::RemoveAt (opnum 17) method removes an item from the collection.
	//
	// The IUpdateCollection::RemoveAt (opnum 17) method removes an item from the collection.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// If the collection is read-only, the server SHOULD return WU_E_NOT_SUPPORTED (0x80240037).
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD remove the item in its List ADM element at the given index and
	// decrement the value of its Count ADM element.
	RemoveAT(context.Context, *RemoveATRequest) (*RemoveATResponse, error)
}

func RegisterUpdateCollectionServer(conn dcerpc.Conn, o UpdateCollectionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewUpdateCollectionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(UpdateCollectionSyntaxV0_0))...)
}

func NewUpdateCollectionServerHandle(o UpdateCollectionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return UpdateCollectionServerHandle(ctx, o, opNum, r)
	}
}

func UpdateCollectionServerHandle(ctx context.Context, o UpdateCollectionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
	case 14: // Opnum15NotUsedOnWire
		// Opnum15NotUsedOnWire
		return nil, nil
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

// Unimplemented IUpdateCollection
type UnimplementedUpdateCollectionServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedUpdateCollectionServer) GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateCollectionServer) SetItem(context.Context, *SetItemRequest) (*SetItemResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateCollectionServer) Get_NewEnum(context.Context, *Get_NewEnumRequest) (*Get_NewEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateCollectionServer) GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateCollectionServer) GetReadOnly(context.Context, *GetReadOnlyRequest) (*GetReadOnlyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateCollectionServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateCollectionServer) Clear(context.Context, *ClearRequest) (*ClearResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateCollectionServer) Insert(context.Context, *InsertRequest) (*InsertResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateCollectionServer) RemoveAT(context.Context, *RemoveATRequest) (*RemoveATResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ UpdateCollectionServer = (*UnimplementedUpdateCollectionServer)(nil)
