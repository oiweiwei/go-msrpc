package iupdatesearcher

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

// IUpdateSearcher server interface.
type UpdateSearcherServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	GetCanAutomaticallyUpgradeService(context.Context, *GetCanAutomaticallyUpgradeServiceRequest) (*GetCanAutomaticallyUpgradeServiceResponse, error)

	SetCanAutomaticallyUpgradeService(context.Context, *SetCanAutomaticallyUpgradeServiceRequest) (*SetCanAutomaticallyUpgradeServiceResponse, error)

	GetClientApplicationID(context.Context, *GetClientApplicationIDRequest) (*GetClientApplicationIDResponse, error)

	SetClientApplicationID(context.Context, *SetClientApplicationIDRequest) (*SetClientApplicationIDResponse, error)

	GetIncludePotentiallySupersededUpdates(context.Context, *GetIncludePotentiallySupersededUpdatesRequest) (*GetIncludePotentiallySupersededUpdatesResponse, error)

	SetIncludePotentiallySupersededUpdates(context.Context, *SetIncludePotentiallySupersededUpdatesRequest) (*SetIncludePotentiallySupersededUpdatesResponse, error)

	GetServerSelection(context.Context, *GetServerSelectionRequest) (*GetServerSelectionResponse, error)

	SetServerSelection(context.Context, *SetServerSelectionRequest) (*SetServerSelectionResponse, error)

	// Opnum16NotUsedOnWire operation.
	// Opnum16NotUsedOnWire

	// Opnum17NotUsedOnWire operation.
	// Opnum17NotUsedOnWire

	EscapeString(context.Context, *EscapeStringRequest) (*EscapeStringResponse, error)

	QueryHistory(context.Context, *QueryHistoryRequest) (*QueryHistoryResponse, error)

	Search(context.Context, *SearchRequest) (*SearchResponse, error)

	GetOnline(context.Context, *GetOnlineRequest) (*GetOnlineResponse, error)

	SetOnline(context.Context, *SetOnlineRequest) (*SetOnlineResponse, error)

	GetTotalHistoryCount(context.Context, *GetTotalHistoryCountRequest) (*GetTotalHistoryCountResponse, error)

	GetServiceID(context.Context, *GetServiceIDRequest) (*GetServiceIDResponse, error)

	SetServiceID(context.Context, *SetServiceIDRequest) (*SetServiceIDResponse, error)
}

func RegisterUpdateSearcherServer(conn dcerpc.Conn, o UpdateSearcherServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewUpdateSearcherServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(UpdateSearcherSyntaxV0_0))...)
}

func NewUpdateSearcherServerHandle(o UpdateSearcherServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return UpdateSearcherServerHandle(ctx, o, opNum, r)
	}
}

func UpdateSearcherServerHandle(ctx context.Context, o UpdateSearcherServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // CanAutomaticallyUpgradeService
		op := &xxx_GetCanAutomaticallyUpgradeServiceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCanAutomaticallyUpgradeServiceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCanAutomaticallyUpgradeService(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // CanAutomaticallyUpgradeService
		op := &xxx_SetCanAutomaticallyUpgradeServiceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetCanAutomaticallyUpgradeServiceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetCanAutomaticallyUpgradeService(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // ClientApplicationID
		op := &xxx_GetClientApplicationIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetClientApplicationIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetClientApplicationID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // ClientApplicationID
		op := &xxx_SetClientApplicationIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetClientApplicationIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetClientApplicationID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // IncludePotentiallySupersededUpdates
		op := &xxx_GetIncludePotentiallySupersededUpdatesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIncludePotentiallySupersededUpdatesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIncludePotentiallySupersededUpdates(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // IncludePotentiallySupersededUpdates
		op := &xxx_SetIncludePotentiallySupersededUpdatesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetIncludePotentiallySupersededUpdatesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetIncludePotentiallySupersededUpdates(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // ServerSelection
		op := &xxx_GetServerSelectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetServerSelectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetServerSelection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // ServerSelection
		op := &xxx_SetServerSelectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetServerSelectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetServerSelection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // Opnum16NotUsedOnWire
		// Opnum16NotUsedOnWire
		return nil, nil
	case 16: // Opnum17NotUsedOnWire
		// Opnum17NotUsedOnWire
		return nil, nil
	case 17: // EscapeString
		op := &xxx_EscapeStringOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EscapeStringRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EscapeString(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // QueryHistory
		op := &xxx_QueryHistoryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryHistoryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryHistory(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // Search
		op := &xxx_SearchOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SearchRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Search(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // Online
		op := &xxx_GetOnlineOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetOnlineRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetOnline(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // Online
		op := &xxx_SetOnlineOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetOnlineRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetOnline(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // GetTotalHistoryCount
		op := &xxx_GetTotalHistoryCountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTotalHistoryCountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTotalHistoryCount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // ServiceID
		op := &xxx_GetServiceIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetServiceIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetServiceID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // ServiceID
		op := &xxx_SetServiceIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetServiceIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetServiceID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IUpdateSearcher
type UnimplementedUpdateSearcherServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedUpdateSearcherServer) GetCanAutomaticallyUpgradeService(context.Context, *GetCanAutomaticallyUpgradeServiceRequest) (*GetCanAutomaticallyUpgradeServiceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSearcherServer) SetCanAutomaticallyUpgradeService(context.Context, *SetCanAutomaticallyUpgradeServiceRequest) (*SetCanAutomaticallyUpgradeServiceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSearcherServer) GetClientApplicationID(context.Context, *GetClientApplicationIDRequest) (*GetClientApplicationIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSearcherServer) SetClientApplicationID(context.Context, *SetClientApplicationIDRequest) (*SetClientApplicationIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSearcherServer) GetIncludePotentiallySupersededUpdates(context.Context, *GetIncludePotentiallySupersededUpdatesRequest) (*GetIncludePotentiallySupersededUpdatesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSearcherServer) SetIncludePotentiallySupersededUpdates(context.Context, *SetIncludePotentiallySupersededUpdatesRequest) (*SetIncludePotentiallySupersededUpdatesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSearcherServer) GetServerSelection(context.Context, *GetServerSelectionRequest) (*GetServerSelectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSearcherServer) SetServerSelection(context.Context, *SetServerSelectionRequest) (*SetServerSelectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSearcherServer) EscapeString(context.Context, *EscapeStringRequest) (*EscapeStringResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSearcherServer) QueryHistory(context.Context, *QueryHistoryRequest) (*QueryHistoryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSearcherServer) Search(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSearcherServer) GetOnline(context.Context, *GetOnlineRequest) (*GetOnlineResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSearcherServer) SetOnline(context.Context, *SetOnlineRequest) (*SetOnlineResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSearcherServer) GetTotalHistoryCount(context.Context, *GetTotalHistoryCountRequest) (*GetTotalHistoryCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSearcherServer) GetServiceID(context.Context, *GetServiceIDRequest) (*GetServiceIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSearcherServer) SetServiceID(context.Context, *SetServiceIDRequest) (*SetServiceIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ UpdateSearcherServer = (*UnimplementedUpdateSearcherServer)(nil)
