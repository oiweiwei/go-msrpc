package itracedataprovider

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
		in := &GetDisplayNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDisplayName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // DisplayName
		in := &SetDisplayNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetDisplayName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // Guid
		in := &GetGUIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetGUID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // Guid
		in := &SetGUIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetGUID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // Level
		in := &GetLevelRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetLevel(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // KeywordsAny
		in := &GetKeywordsAnyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetKeywordsAny(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // KeywordsAll
		in := &GetKeywordsAllRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetKeywordsAll(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // Properties
		in := &GetPropertiesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetProperties(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // FilterEnabled
		in := &GetFilterEnabledRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFilterEnabled(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // FilterEnabled
		in := &SetFilterEnabledRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetFilterEnabled(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // FilterType
		in := &GetFilterTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFilterType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // FilterType
		in := &SetFilterTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetFilterType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // FilterData
		in := &GetFilterDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFilterData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // FilterData
		in := &SetFilterDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetFilterData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // Query
		in := &QueryRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Query(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 22: // Resolve
		in := &ResolveRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Resolve(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // SetSecurity
		in := &SetSecurityRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSecurity(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // GetSecurity
		in := &GetSecurityRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSecurity(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // GetRegisteredProcesses
		in := &GetRegisteredProcessesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetRegisteredProcesses(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
