package iwbemrefreshingservices

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
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
	_ = iunknown.GoPackage
)

// IWbemRefreshingServices server interface.
type RefreshingServicesServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The IWbemRefreshingServices::AddObjectToRefresher method MUST add a CIM instance,
	// which is identified by its CIM path, to the list of CIM instances that can be refreshed.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	AddObjectToRefresher(context.Context, *AddObjectToRefresherRequest) (*AddObjectToRefresherResponse, error)

	// The IWbemRefreshingServices::AddObjectToRefresherByTemplate method MUST add a CIM
	// instance, which is identified by its CIM object instance, to the list of CIM instances
	// to be refreshed.
	//
	// The AddObjectToRefresherByTemplate method opnum equals 4.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	AddObjectToRefresherByTemplate(context.Context, *AddObjectToRefresherByTemplateRequest) (*AddObjectToRefresherByTemplateResponse, error)

	// The IWbemRefreshingServices::AddEnumToRefresher method MUST add all CIM instances
	// that are identified by the CIM class name to the list of CIM instances to be refreshed.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	AddEnumToRefresher(context.Context, *AddEnumToRefresherRequest) (*AddEnumToRefresherResponse, error)

	// The IWbemRefreshingServices::RemoveObjectFromRefresher method MUST remove a CIM instance,
	// which is identified by its CIM path, from the list of CIM instances that can be refreshed.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. If there are no failures, the server MUST always return WBEM_E_NOT_AVAILABLE.<68>
	RemoveObjectFromRefresher(context.Context, *RemoveObjectFromRefresherRequest) (*RemoveObjectFromRefresherResponse, error)

	// The IWbemRefreshingServices::GetRemoteRefresher method MUST return an IWbemRemoteRefresher
	// interface pointer. This pointer is needed by the client to refresh objects and enumerations.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	//
	// In case of failure, the server MUST return an HRESULT whose S (severity) bit is set
	// as specified in [MS-ERREF] section 2.1. The actual HRESULT value is implementation
	// dependent.
	GetRemoteRefresher(context.Context, *GetRemoteRefresherRequest) (*GetRemoteRefresherResponse, error)

	// The IWbemRefreshingServices::ReconnectRemoteRefresher method MUST restore a set of
	// CIM instances and enumerations that are passed in apReconnectInfo to a refresher.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR, as specified in section
	// 2.2.11, to indicate the successful completion of the method.
	ReconnectRemoteRefresher(context.Context, *ReconnectRemoteRefresherRequest) (*ReconnectRemoteRefresherResponse, error)
}

func RegisterRefreshingServicesServer(conn dcerpc.Conn, o RefreshingServicesServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewRefreshingServicesServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(RefreshingServicesSyntaxV0_0))...)
}

func NewRefreshingServicesServerHandle(o RefreshingServicesServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return RefreshingServicesServerHandle(ctx, o, opNum, r)
	}
}

func RefreshingServicesServerHandle(ctx context.Context, o RefreshingServicesServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // AddObjectToRefresher
		in := &AddObjectToRefresherRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddObjectToRefresher(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // AddObjectToRefresherByTemplate
		in := &AddObjectToRefresherByTemplateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddObjectToRefresherByTemplate(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // AddEnumToRefresher
		in := &AddEnumToRefresherRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddEnumToRefresher(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // RemoveObjectFromRefresher
		in := &RemoveObjectFromRefresherRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoveObjectFromRefresher(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // GetRemoteRefresher
		in := &GetRemoteRefresherRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetRemoteRefresher(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // ReconnectRemoteRefresher
		in := &ReconnectRemoteRefresherRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ReconnectRemoteRefresher(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
