package iwbemremoterefresher

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

// IWbemRemoteRefresher server interface.
type RemoteRefresherServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The IWbemRemoteRefresher::RemoteRefresh method MUST return the updated collection
	// of CIM instances and enumerations previously configured by the IWbemRefreshingServices
	// interface pointer.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call.
	//
	// The server MUST return WBEM_S_NO_ERROR (specified in section 2.2.11) to indicate
	// the successful completion of the method.
	//
	// The IWbemRemoteRefresher::RemoteRefresh method MUST be called on the IWbemRemoteRefresher
	// interface pointer returned as a member of the _WBEM_REFRESH_INFO structure from IWbemRefreshingServices
	// methods or on the interface returned by IWbemRefreshingServices::GetRemoteRefresher
	// method invocation.
	RemoteRefresh(context.Context, *RemoteRefreshRequest) (*RemoteRefreshResponse, error)

	// The IWbemRemoteRefresher::StopRefreshing method MUST remove a set of CIM instances
	// or enumerations from the collection previously configured by the IWbemRefreshingServices
	// interface pointer.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. In case of success, the server MUST return WBEM_S_NO_ERROR (as
	// specified in section 2.2.11) to indicate the successful completion of the method.
	//
	// The IWbemRemoteRefresher::StopRefreshing method MUST be called on the IWbemRemoteRefresher
	// interface pointer that is returned as a member of the _WBEM_REFRESH_INFO structure
	// from the methods of the IWbemRefreshingServices interface or on the interface that
	// is returned by the IWbemRefreshingServices::GetRemoteRefresher method invocation.
	StopRefreshing(context.Context, *StopRefreshingRequest) (*StopRefreshingResponse, error)

	// The IWbemRemoteRefresher::Opnum5NotUsedOnWire method MUST return a random GUID that
	// identifies the server object that receives the call.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	//
	// Opnum5NotUsedOnWire
}

func RegisterRemoteRefresherServer(conn dcerpc.Conn, o RemoteRefresherServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewRemoteRefresherServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(RemoteRefresherSyntaxV0_0))...)
}

func NewRemoteRefresherServerHandle(o RemoteRefresherServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return RemoteRefresherServerHandle(ctx, o, opNum, r)
	}
}

func RemoteRefresherServerHandle(ctx context.Context, o RemoteRefresherServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // RemoteRefresh
		op := &xxx_RemoteRefreshOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoteRefreshRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoteRefresh(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // StopRefreshing
		op := &xxx_StopRefreshingOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StopRefreshingRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.StopRefreshing(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // Opnum5NotUsedOnWire
		// Opnum5NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}

// Unimplemented IWbemRemoteRefresher
type UnimplementedRemoteRefresherServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedRemoteRefresherServer) RemoteRefresh(context.Context, *RemoteRefreshRequest) (*RemoteRefreshResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteRefresherServer) StopRefreshing(context.Context, *StopRefreshingRequest) (*StopRefreshingResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ RemoteRefresherServer = (*UnimplementedRemoteRefresherServer)(nil)
