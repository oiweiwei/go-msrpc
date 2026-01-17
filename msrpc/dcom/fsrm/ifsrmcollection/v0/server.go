package ifsrmcollection

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

// IFsrmCollection server interface.
type CollectionServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// _NewEnum operation.
	Get_NewEnum(context.Context, *Get_NewEnumRequest) (*Get_NewEnumResponse, error)

	// The Item method returns a pointer to the object at the requested position in the
	// collection of Objects Being Enumerated.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN                |                                                                                  |
	//	|             VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80004003 E_POINTER                | The item parameter is NULL.                                                      |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80131502 COR_E_ARGUMENTOUTOFRANGE | The value of the index parameter is greater than the number of Objects Being     |
	//	|                                     | Enumerated in the collection.                                                    |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error)

	// The Count method returns the number of objects in the collection of Objects Being
	// Enumerated.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+----------------------+------------------------------+
	//	|        RETURN        |                              |
	//	|      VALUE/CODE      |         DESCRIPTION          |
	//	|                      |                              |
	//	+----------------------+------------------------------+
	//	+----------------------+------------------------------+
	//	| 0x80004003 E_POINTER | The count parameter is NULL. |
	//	+----------------------+------------------------------+
	GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error)

	// The State method returns the state FsrmCollectionState_Complete.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+----------------------+------------------------------+
	//	|        RETURN        |                              |
	//	|      VALUE/CODE      |         DESCRIPTION          |
	//	|                      |                              |
	//	+----------------------+------------------------------+
	//	+----------------------+------------------------------+
	//	| 0x80004003 E_POINTER | The state parameter is NULL. |
	//	+----------------------+------------------------------+
	GetState(context.Context, *GetStateRequest) (*GetStateResponse, error)

	// Cancel operation.
	Cancel(context.Context, *CancelRequest) (*CancelResponse, error)

	// WaitForCompletion operation.
	WaitForCompletion(context.Context, *WaitForCompletionRequest) (*WaitForCompletionResponse, error)

	// The GetById method returns the object from the collection of Objects Being Enumerated
	// (section 3.2.1.11) whose ID matches the specified id.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	|           RETURN            |                                                                                  |
	//	|         VALUE/CODE          |                                   DESCRIPTION                                    |
	//	|                             |                                                                                  |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND | An object with the specified ID was not found in the collection.                 |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG     | This code is returned for the following reasons: The entry parameter is          |
	//	|                             | NULL. The objects in the collection of Objects Being Enumerated are not          |
	//	|                             | one of the following interfaces: IFsrmFileScreen, IFsrmFileScreenException,      |
	//	|                             | IFsrmFileScreenTemplate, IFsrmFileGroup, IFsrmQuota, IFsrmQuotaTemplate,         |
	//	|                             | IFsrmAction, IFsrmReportJob, IFsrmReport, IFsrmClassifcationRule,                |
	//	|                             | IFsrmPropertyDefinition, IFsrmPipelineModuleDefinition or                        |
	//	|                             | IFsrmFileManagementJob.                                                          |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	GetByID(context.Context, *GetByIDRequest) (*GetByIDResponse, error)
}

func RegisterCollectionServer(conn dcerpc.Conn, o CollectionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewCollectionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(CollectionSyntaxV0_0))...)
}

func NewCollectionServerHandle(o CollectionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return CollectionServerHandle(ctx, o, opNum, r)
	}
}

func CollectionServerHandle(ctx context.Context, o CollectionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // _NewEnum
		op := &xxx_Get_NewEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &Get_NewEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Get_NewEnum(ctx, req)
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
	case 9: // Count
		op := &xxx_GetCountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // State
		op := &xxx_GetStateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetStateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetState(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // Cancel
		op := &xxx_CancelOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CancelRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Cancel(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // WaitForCompletion
		op := &xxx_WaitForCompletionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &WaitForCompletionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.WaitForCompletion(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // GetById
		op := &xxx_GetByIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetByIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetByID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmCollection
type UnimplementedCollectionServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedCollectionServer) Get_NewEnum(context.Context, *Get_NewEnumRequest) (*Get_NewEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCollectionServer) GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCollectionServer) GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCollectionServer) GetState(context.Context, *GetStateRequest) (*GetStateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCollectionServer) Cancel(context.Context, *CancelRequest) (*CancelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCollectionServer) WaitForCompletion(context.Context, *WaitForCompletionRequest) (*WaitForCompletionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCollectionServer) GetByID(context.Context, *GetByIDRequest) (*GetByIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ CollectionServer = (*UnimplementedCollectionServer)(nil)
