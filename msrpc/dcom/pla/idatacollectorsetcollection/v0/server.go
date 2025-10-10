package idatacollectorsetcollection

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

// IDataCollectorSetCollection server interface.
type DataCollectorSetCollectionServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// Count operation.
	GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error)

	// Item operation.
	GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error)

	// _NewEnum operation.
	Get_NewEnum(context.Context, *Get_NewEnumRequest) (*Get_NewEnumResponse, error)

	// Add operation.
	Add(context.Context, *AddRequest) (*AddResponse, error)

	// Remove operation.
	Remove(context.Context, *RemoveRequest) (*RemoveResponse, error)

	// Clear operation.
	Clear(context.Context, *ClearRequest) (*ClearResponse, error)

	// AddRange operation.
	AddRange(context.Context, *AddRangeRequest) (*AddRangeResponse, error)

	// The GetDataCollectorSets method populates data collector set collection with the
	// persisted data collector sets. A data collector set is considered to be persisted
	// if either of the following conditions is met:
	//
	// *
	//
	// The data collector set is in a running state. A data collector set is considered
	// to be in a running state if a call to IDataCollectorSet::getState ( 695c56b5-8762-419d-a047-aa9762d188c4
	// ) returns plaRunning. A data collector set enters a running state by calling IDataCollectorSet::Start
	// ( d2ebf91b-9e67-440d-90e9-3134ee1613a0 ). A data collector set can be removed from
	// a running state by calling IDataCollectorSet::Stop ( 2617595b-63fc-4c23-b674-4cb3e062eb6f
	// ).
	//
	// *
	//
	// The data collector set is committed. A data collector set is committed after a successful
	// call to IDataCollectorSet::Commit ( b7b0e0f8-327c-46de-b034-3d598306890e ) where
	// the CommitMode has any value other than plaValidateOnly. A data collector set can
	// be removed from a committed state by calling IDataCollectorSet::Delete ( 305c3442-6a96-43b6-8144-c0cfb1ebbec1
	// ).
	//
	// Semantically, a data collector set is persisted if it has been committed to a permanent
	// store, such as the filesystem.<35>
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetDataCollectorSets(context.Context, *GetDataCollectorSetsRequest) (*GetDataCollectorSetsResponse, error)
}

func RegisterDataCollectorSetCollectionServer(conn dcerpc.Conn, o DataCollectorSetCollectionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewDataCollectorSetCollectionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(DataCollectorSetCollectionSyntaxV0_0))...)
}

func NewDataCollectorSetCollectionServerHandle(o DataCollectorSetCollectionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return DataCollectorSetCollectionServerHandle(ctx, o, opNum, r)
	}
}

func DataCollectorSetCollectionServerHandle(ctx context.Context, o DataCollectorSetCollectionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Count
		op := &xxx_GetCountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // Item
		op := &xxx_GetItemOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetItemRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetItem(ctx, req)
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
	case 10: // Add
		op := &xxx_AddOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Add(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // Remove
		op := &xxx_RemoveOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Remove(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // Clear
		op := &xxx_ClearOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ClearRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Clear(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // AddRange
		op := &xxx_AddRangeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddRangeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddRange(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // GetDataCollectorSets
		op := &xxx_GetDataCollectorSetsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDataCollectorSetsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDataCollectorSets(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IDataCollectorSetCollection
type UnimplementedDataCollectorSetCollectionServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedDataCollectorSetCollectionServer) GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetCollectionServer) GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetCollectionServer) Get_NewEnum(context.Context, *Get_NewEnumRequest) (*Get_NewEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetCollectionServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetCollectionServer) Remove(context.Context, *RemoveRequest) (*RemoveResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetCollectionServer) Clear(context.Context, *ClearRequest) (*ClearResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetCollectionServer) AddRange(context.Context, *AddRangeRequest) (*AddRangeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorSetCollectionServer) GetDataCollectorSets(context.Context, *GetDataCollectorSetsRequest) (*GetDataCollectorSetsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ DataCollectorSetCollectionServer = (*UnimplementedDataCollectorSetCollectionServer)(nil)
