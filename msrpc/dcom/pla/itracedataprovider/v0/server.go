package itracedataprovider

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

// ITraceDataProvider server interface.
type TraceDataProviderServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// DisplayName operation.
	GetDisplayName(context.Context, *GetDisplayNameRequest) (*GetDisplayNameResponse, error)

	// DisplayName operation.
	SetDisplayName(context.Context, *SetDisplayNameRequest) (*SetDisplayNameResponse, error)

	// Guid operation.
	GetGUID(context.Context, *GetGUIDRequest) (*GetGUIDResponse, error)

	// Guid operation.
	SetGUID(context.Context, *SetGUIDRequest) (*SetGUIDResponse, error)

	// The Level (Get) method retrieves the Level property, as specified in the property
	// table in section 3.2.4.11.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetLevel(context.Context, *GetLevelRequest) (*GetLevelResponse, error)

	// The KeywordsAny (Get) method retrieves the KeywordsAny property, as specified in
	// the property table in section 3.2.4.11.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetKeywordsAny(context.Context, *GetKeywordsAnyRequest) (*GetKeywordsAnyResponse, error)

	// The KeywordsAll (Get) method retrieves the KeywordsAll property, as specified in
	// the property table in section 3.2.4.11.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetKeywordsAll(context.Context, *GetKeywordsAllRequest) (*GetKeywordsAllResponse, error)

	// The Properties (Get) method retrieves the Properties property, as specified in the
	// property table in section 3.2.4.11.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)

	// FilterEnabled operation.
	GetFilterEnabled(context.Context, *GetFilterEnabledRequest) (*GetFilterEnabledResponse, error)

	// FilterEnabled operation.
	SetFilterEnabled(context.Context, *SetFilterEnabledRequest) (*SetFilterEnabledResponse, error)

	// FilterType operation.
	GetFilterType(context.Context, *GetFilterTypeRequest) (*GetFilterTypeResponse, error)

	// FilterType operation.
	SetFilterType(context.Context, *SetFilterTypeRequest) (*SetFilterTypeResponse, error)

	// FilterData operation.
	GetFilterData(context.Context, *GetFilterDataRequest) (*GetFilterDataResponse, error)

	// FilterData operation.
	SetFilterData(context.Context, *SetFilterDataRequest) (*SetFilterDataResponse, error)

	// Query operation.
	Query(context.Context, *QueryRequest) (*QueryResponse, error)

	// The Resolve method, given another event trace data provider or a collection of event
	// trace data providers, updates the properties of the original provider with information
	// from one of the passed-in provider(s). If only one provider is passed-in, information
	// from that provider is used. If multiple providers are passed-in, information is used
	// from the first provider in the collection that has the same value for the ITraceDataProvider::Guid
	// property as the original ITraceDataProvider. The Guid in this context refers to the
	// COM interface property which is used as a PLA-UID and does not comply with the semantics
	// of GUID specified in [C706].
	//
	// If no provider from the passed-in collection has the same PLA-UID, the original provider
	// is not updated. If the original provider is updated, the DisplayName property of
	// the original provider is overwritten by the DisplayName of the passed-in provider,
	// and the ValueMapItems in the Level, KeywordsAny, and KeywordsAll properties of the
	// passed-in provider replace the ValueMapItems for the existing Level, KeywordsAny,
	// and KeywordsAll properties of the original provider. However, actual value of the
	// Level, KeywordsAny, and KeywordsAll properties are not overwritten. Consequently,
	// the existing settings are not lost; that is, if the level is 5, the symbolic names
	// of all the levels (which are stored as ValueMapItems) are added, but the value of
	// the level remains as 5. For ValueMaps, see section 3.2.4.18.
	//
	// Because only one provider is used to update the original provider, there is no possibility
	// for conflicting or duplicate properties.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	Resolve(context.Context, *ResolveRequest) (*ResolveResponse, error)

	// The SetSecurity method updates the system-wide security descriptor of the provider.
	// Because the security descriptor is system-wide, the update will impact the ability
	// of all users (local or remote) to view, modify, enable, or delete the provider.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	SetSecurity(context.Context, *SetSecurityRequest) (*SetSecurityResponse, error)

	// The GetSecurity method retrieves the system-wide security descriptor of the provider.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetSecurity(context.Context, *GetSecurityRequest) (*GetSecurityResponse, error)

	// The GetRegisteredProcesses method retrieves a list of processes that have registered
	// as an event trace data provider.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetRegisteredProcesses(context.Context, *GetRegisteredProcessesRequest) (*GetRegisteredProcessesResponse, error)
}

func RegisterTraceDataProviderServer(conn dcerpc.Conn, o TraceDataProviderServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewTraceDataProviderServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(TraceDataProviderSyntaxV0_0))...)
}

func NewTraceDataProviderServerHandle(o TraceDataProviderServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return TraceDataProviderServerHandle(ctx, o, opNum, r)
	}
}

func TraceDataProviderServerHandle(ctx context.Context, o TraceDataProviderServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // DisplayName
		op := &xxx_GetDisplayNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDisplayNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDisplayName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // DisplayName
		op := &xxx_SetDisplayNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetDisplayNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetDisplayName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // Guid
		op := &xxx_GetGUIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetGUIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetGUID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // Guid
		op := &xxx_SetGUIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetGUIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetGUID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // Level
		op := &xxx_GetLevelOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLevelRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLevel(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // KeywordsAny
		op := &xxx_GetKeywordsAnyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetKeywordsAnyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetKeywordsAny(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // KeywordsAll
		op := &xxx_GetKeywordsAllOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetKeywordsAllRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetKeywordsAll(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // Properties
		op := &xxx_GetPropertiesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPropertiesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetProperties(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // FilterEnabled
		op := &xxx_GetFilterEnabledOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFilterEnabledRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFilterEnabled(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // FilterEnabled
		op := &xxx_SetFilterEnabledOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFilterEnabledRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFilterEnabled(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // FilterType
		op := &xxx_GetFilterTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFilterTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFilterType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // FilterType
		op := &xxx_SetFilterTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFilterTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFilterType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // FilterData
		op := &xxx_GetFilterDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFilterDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFilterData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // FilterData
		op := &xxx_SetFilterDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFilterDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFilterData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // Query
		op := &xxx_QueryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Query(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // Resolve
		op := &xxx_ResolveOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ResolveRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Resolve(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // SetSecurity
		op := &xxx_SetSecurityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSecurityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSecurity(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // GetSecurity
		op := &xxx_GetSecurityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSecurityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSecurity(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // GetRegisteredProcesses
		op := &xxx_GetRegisteredProcessesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRegisteredProcessesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRegisteredProcesses(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented ITraceDataProvider
type UnimplementedTraceDataProviderServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedTraceDataProviderServer) GetDisplayName(context.Context, *GetDisplayNameRequest) (*GetDisplayNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataProviderServer) SetDisplayName(context.Context, *SetDisplayNameRequest) (*SetDisplayNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataProviderServer) GetGUID(context.Context, *GetGUIDRequest) (*GetGUIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataProviderServer) SetGUID(context.Context, *SetGUIDRequest) (*SetGUIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataProviderServer) GetLevel(context.Context, *GetLevelRequest) (*GetLevelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataProviderServer) GetKeywordsAny(context.Context, *GetKeywordsAnyRequest) (*GetKeywordsAnyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataProviderServer) GetKeywordsAll(context.Context, *GetKeywordsAllRequest) (*GetKeywordsAllResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataProviderServer) GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataProviderServer) GetFilterEnabled(context.Context, *GetFilterEnabledRequest) (*GetFilterEnabledResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataProviderServer) SetFilterEnabled(context.Context, *SetFilterEnabledRequest) (*SetFilterEnabledResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataProviderServer) GetFilterType(context.Context, *GetFilterTypeRequest) (*GetFilterTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataProviderServer) SetFilterType(context.Context, *SetFilterTypeRequest) (*SetFilterTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataProviderServer) GetFilterData(context.Context, *GetFilterDataRequest) (*GetFilterDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataProviderServer) SetFilterData(context.Context, *SetFilterDataRequest) (*SetFilterDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataProviderServer) Query(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataProviderServer) Resolve(context.Context, *ResolveRequest) (*ResolveResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataProviderServer) SetSecurity(context.Context, *SetSecurityRequest) (*SetSecurityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataProviderServer) GetSecurity(context.Context, *GetSecurityRequest) (*GetSecurityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTraceDataProviderServer) GetRegisteredProcesses(context.Context, *GetRegisteredProcessesRequest) (*GetRegisteredProcessesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ TraceDataProviderServer = (*UnimplementedTraceDataProviderServer)(nil)
