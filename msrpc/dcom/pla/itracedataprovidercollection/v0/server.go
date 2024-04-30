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
		in := &GetCountRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetCount(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // Item
		in := &GetItemRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetItem(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // _NewEnum
		in := &Get_NewEnumRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Get_NewEnum(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // Add
		in := &AddRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Add(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // Remove
		in := &RemoveRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Remove(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // Clear
		in := &ClearRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Clear(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // AddRange
		in := &AddRangeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddRange(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // CreateTraceDataProvider
		in := &CreateTraceDataProviderRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateTraceDataProvider(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // GetTraceDataProviders
		in := &GetTraceDataProvidersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTraceDataProviders(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // GetTraceDataProvidersByProcess
		in := &GetTraceDataProvidersByProcessRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTraceDataProvidersByProcess(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
