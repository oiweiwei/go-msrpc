package idatacollectorcollection

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

// IDataCollectorCollection server interface.
type DataCollectorCollectionServer interface {

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

	// The CreateDataCollectorFromXml method creates a data collector using XML.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	CreateDataCollectorFromXML(context.Context, *CreateDataCollectorFromXMLRequest) (*CreateDataCollectorFromXMLResponse, error)

	// The CreateDataCollector method creates a data collector of the specified type.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	CreateDataCollector(context.Context, *CreateDataCollectorRequest) (*CreateDataCollectorResponse, error)
}

func RegisterDataCollectorCollectionServer(conn dcerpc.Conn, o DataCollectorCollectionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewDataCollectorCollectionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(DataCollectorCollectionSyntaxV0_0))...)
}

func NewDataCollectorCollectionServerHandle(o DataCollectorCollectionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return DataCollectorCollectionServerHandle(ctx, o, opNum, r)
	}
}

func DataCollectorCollectionServerHandle(ctx context.Context, o DataCollectorCollectionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
	case 14: // CreateDataCollectorFromXml
		op := &xxx_CreateDataCollectorFromXMLOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateDataCollectorFromXMLRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateDataCollectorFromXML(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // CreateDataCollector
		op := &xxx_CreateDataCollectorOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateDataCollectorRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateDataCollector(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IDataCollectorCollection
type UnimplementedDataCollectorCollectionServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedDataCollectorCollectionServer) GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorCollectionServer) GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorCollectionServer) Get_NewEnum(context.Context, *Get_NewEnumRequest) (*Get_NewEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorCollectionServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorCollectionServer) Remove(context.Context, *RemoveRequest) (*RemoveResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorCollectionServer) Clear(context.Context, *ClearRequest) (*ClearResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorCollectionServer) AddRange(context.Context, *AddRangeRequest) (*AddRangeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorCollectionServer) CreateDataCollectorFromXML(context.Context, *CreateDataCollectorFromXMLRequest) (*CreateDataCollectorFromXMLResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorCollectionServer) CreateDataCollector(context.Context, *CreateDataCollectorRequest) (*CreateDataCollectorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ DataCollectorCollectionServer = (*UnimplementedDataCollectorCollectionServer)(nil)
