package ifsrmcollection

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
		in := &Get_NewEnumRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Get_NewEnum(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // Item
		in := &GetItemRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetItem(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // Count
		in := &GetCountRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetCount(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // State
		in := &GetStateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetState(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // Cancel
		in := &CancelRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Cancel(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // WaitForCompletion
		in := &WaitForCompletionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.WaitForCompletion(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // GetById
		in := &GetByIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetByID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
