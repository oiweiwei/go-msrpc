package itracedataprovidercollection

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

// ITraceDataProviderCollection server interface.
type TraceDataProviderCollectionServer interface {

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

	// The CreateTraceDataProvider method creates a trace data provider object.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	CreateTraceDataProvider(context.Context, *CreateTraceDataProviderRequest) (*CreateTraceDataProviderResponse, error)

	// The GetTraceDataProviders method populates the collection with the registered trace
	// providers.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetTraceDataProviders(context.Context, *GetTraceDataProvidersRequest) (*GetTraceDataProvidersResponse, error)

	// The GetTraceDataProvidersByProcess method populates the collection with the list
	// of providers that were registered by the given process.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetTraceDataProvidersByProcess(context.Context, *GetTraceDataProvidersByProcessRequest) (*GetTraceDataProvidersByProcessResponse, error)
}

func RegisterTraceDataProviderCollectionServer(conn dcerpc.Conn, o TraceDataProviderCollectionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewTraceDataProviderCollectionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(TraceDataProviderCollectionSyntaxV0_0))...)
}

func NewTraceDataProviderCollectionServerHandle(o TraceDataProviderCollectionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return TraceDataProviderCollectionServerHandle(ctx, o, opNum, r)
	}
}

func TraceDataProviderCollectionServerHandle(ctx context.Context, o TraceDataProviderCollectionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
	case 14: // CreateTraceDataProvider
		op := &xxx_CreateTraceDataProviderOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateTraceDataProviderRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateTraceDataProvider(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // GetTraceDataProviders
		op := &xxx_GetTraceDataProvidersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTraceDataProvidersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTraceDataProviders(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // GetTraceDataProvidersByProcess
		op := &xxx_GetTraceDataProvidersByProcessOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTraceDataProvidersByProcessRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTraceDataProvidersByProcess(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented ITraceDataProviderCollection
type UnimplementedTraceDataProviderCollectionServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedTraceDataProviderCollectionServer) GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataProviderCollectionServer) GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataProviderCollectionServer) Get_NewEnum(context.Context, *Get_NewEnumRequest) (*Get_NewEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataProviderCollectionServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataProviderCollectionServer) Remove(context.Context, *RemoveRequest) (*RemoveResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataProviderCollectionServer) Clear(context.Context, *ClearRequest) (*ClearResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataProviderCollectionServer) AddRange(context.Context, *AddRangeRequest) (*AddRangeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataProviderCollectionServer) CreateTraceDataProvider(context.Context, *CreateTraceDataProviderRequest) (*CreateTraceDataProviderResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataProviderCollectionServer) GetTraceDataProviders(context.Context, *GetTraceDataProvidersRequest) (*GetTraceDataProvidersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataProviderCollectionServer) GetTraceDataProvidersByProcess(context.Context, *GetTraceDataProvidersByProcessRequest) (*GetTraceDataProvidersByProcessResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ TraceDataProviderCollectionServer = (*UnimplementedTraceDataProviderCollectionServer)(nil)
