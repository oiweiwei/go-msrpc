package iwrmconfig

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

// IWRMConfig server interface.
type IwrmConfigServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error)

	SetConfig(context.Context, *SetConfigRequest) (*SetConfigResponse, error)

	IsEnabled(context.Context, *IsEnabledRequest) (*IsEnabledResponse, error)

	EnableDisable(context.Context, *EnableDisableRequest) (*EnableDisableResponse, error)

	GetExclusionList(context.Context, *GetExclusionListRequest) (*GetExclusionListResponse, error)

	SetExclusionList(context.Context, *SetExclusionListRequest) (*SetExclusionListResponse, error)

	WsrmActivate(context.Context, *WsrmActivateRequest) (*WsrmActivateResponse, error)

	IsWsrmActivated(context.Context, *IsWsrmActivatedRequest) (*IsWsrmActivatedResponse, error)

	RestoreExclusionList(context.Context, *RestoreExclusionListRequest) (*RestoreExclusionListResponse, error)
}

func RegisterIwrmConfigServer(conn dcerpc.Conn, o IwrmConfigServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewIwrmConfigServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(IwrmConfigSyntaxV0_0))...)
}

func NewIwrmConfigServerHandle(o IwrmConfigServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return IwrmConfigServerHandle(ctx, o, opNum, r)
	}
}

func IwrmConfigServerHandle(ctx context.Context, o IwrmConfigServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // GetConfig
		op := &xxx_GetConfigOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetConfigRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetConfig(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // SetConfig
		op := &xxx_SetConfigOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetConfigRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetConfig(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // IsEnabled
		op := &xxx_IsEnabledOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IsEnabledRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.IsEnabled(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // EnableDisable
		op := &xxx_EnableDisableOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnableDisableRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnableDisable(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // GetExclusionList
		op := &xxx_GetExclusionListOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetExclusionListRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetExclusionList(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // SetExclusionList
		op := &xxx_SetExclusionListOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetExclusionListRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetExclusionList(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // WSRMActivate
		op := &xxx_WsrmActivateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &WsrmActivateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.WsrmActivate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // IsWSRMActivated
		op := &xxx_IsWsrmActivatedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IsWsrmActivatedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.IsWsrmActivated(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // RestoreExclusionList
		op := &xxx_RestoreExclusionListOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RestoreExclusionListRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RestoreExclusionList(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IWRMConfig
type UnimplementedIwrmConfigServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedIwrmConfigServer) GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIwrmConfigServer) SetConfig(context.Context, *SetConfigRequest) (*SetConfigResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIwrmConfigServer) IsEnabled(context.Context, *IsEnabledRequest) (*IsEnabledResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIwrmConfigServer) EnableDisable(context.Context, *EnableDisableRequest) (*EnableDisableResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIwrmConfigServer) GetExclusionList(context.Context, *GetExclusionListRequest) (*GetExclusionListResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIwrmConfigServer) SetExclusionList(context.Context, *SetExclusionListRequest) (*SetExclusionListResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIwrmConfigServer) WsrmActivate(context.Context, *WsrmActivateRequest) (*WsrmActivateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIwrmConfigServer) IsWsrmActivated(context.Context, *IsWsrmActivatedRequest) (*IsWsrmActivatedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIwrmConfigServer) RestoreExclusionList(context.Context, *RestoreExclusionListRequest) (*RestoreExclusionListResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ IwrmConfigServer = (*UnimplementedIwrmConfigServer)(nil)
