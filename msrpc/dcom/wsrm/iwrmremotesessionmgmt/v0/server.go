package iwrmremotesessionmgmt

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

// IWRMRemoteSessionMgmt server interface.
type RemoteSessionManagementServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The GetRemoteUserCategories method retrieves user categories information from the
	// WSRM server.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------+-----------------------+
	//	|      RETURN       |                       |
	//	|    VALUE/CODE     |      DESCRIPTION      |
	//	|                   |                       |
	//	+-------------------+-----------------------+
	//	+-------------------+-----------------------+
	//	| 0x00000000 S_OK   | Operation successful. |
	//	+-------------------+-----------------------+
	//
	// Additional IWRMRemoteSessionMgmt interface methods are specified in section 3.2.4.9.
	GetRemoteUserCategories(context.Context, *GetRemoteUserCategoriesRequest) (*GetRemoteUserCategoriesResponse, error)

	// The SetRemoteUserCategories method sets user categories information on the WSRM server.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                      | Operation successful.                                                            |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG              | One or more arguments are invalid.                                               |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0070 WRM_ERR_TAGS_NOT_IN_ORDER | The XML data that is maintained by the management service is invalid or cannot   |
	//	|                                      | be processed.<114>                                                               |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMRemoteSessionMgmt interface methods are specified in section 3.2.4.9.
	SetRemoteUserCategories(context.Context, *SetRemoteUserCategoriesRequest) (*SetRemoteUserCategoriesResponse, error)

	// The RefreshRemoteSessionWeights method forces reallocation of CPU quotas for the
	// sessions run by users according to the category type specified in bstrTargetUserSessions.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0070 WRM_ERR_TAGS_NOT_IN_ORDER | The XML data is invalid or cannot be processed.                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF006F WRM_ERR_ID_VALUE          | The specified name contains characters that are invalid. The name cannot         |
	//	|                                      | start with a hyphen (-), cannot contain spaces, and cannot contain any of the    |
	//	|                                      | following characters: \ / ? * | : < > " , ;                                      |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                      | Operation successful.                                                            |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMRemoteSessionMgmt interface methods are specified in section 3.2.4.9.
	RefreshRemoteSessionWeights(context.Context, *RefreshRemoteSessionWeightsRequest) (*RefreshRemoteSessionWeightsResponse, error)
}

func RegisterRemoteSessionManagementServer(conn dcerpc.Conn, o RemoteSessionManagementServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewRemoteSessionManagementServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(RemoteSessionManagementSyntaxV0_0))...)
}

func NewRemoteSessionManagementServerHandle(o RemoteSessionManagementServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return RemoteSessionManagementServerHandle(ctx, o, opNum, r)
	}
}

func RemoteSessionManagementServerHandle(ctx context.Context, o RemoteSessionManagementServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // GetRemoteUserCategories
		op := &xxx_GetRemoteUserCategoriesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRemoteUserCategoriesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRemoteUserCategories(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // SetRemoteUserCategories
		op := &xxx_SetRemoteUserCategoriesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetRemoteUserCategoriesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetRemoteUserCategories(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // RefreshRemoteSessionWeights
		op := &xxx_RefreshRemoteSessionWeightsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RefreshRemoteSessionWeightsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RefreshRemoteSessionWeights(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IWRMRemoteSessionMgmt
type UnimplementedRemoteSessionManagementServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedRemoteSessionManagementServer) GetRemoteUserCategories(context.Context, *GetRemoteUserCategoriesRequest) (*GetRemoteUserCategoriesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteSessionManagementServer) SetRemoteUserCategories(context.Context, *SetRemoteUserCategoriesRequest) (*SetRemoteUserCategoriesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteSessionManagementServer) RefreshRemoteSessionWeights(context.Context, *RefreshRemoteSessionWeightsRequest) (*RefreshRemoteSessionWeightsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ RemoteSessionManagementServer = (*UnimplementedRemoteSessionManagementServer)(nil)
