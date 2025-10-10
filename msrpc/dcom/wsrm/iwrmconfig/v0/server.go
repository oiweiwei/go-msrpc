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
type ConfigServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error)

	SetConfig(context.Context, *SetConfigRequest) (*SetConfigResponse, error)

	IsEnabled(context.Context, *IsEnabledRequest) (*IsEnabledResponse, error)

	EnableDisable(context.Context, *EnableDisableRequest) (*EnableDisableResponse, error)

	GetExclusionList(context.Context, *GetExclusionListRequest) (*GetExclusionListResponse, error)

	SetExclusionList(context.Context, *SetExclusionListRequest) (*SetExclusionListResponse, error)

	WSRMActivate(context.Context, *WSRMActivateRequest) (*WSRMActivateResponse, error)

	IsWSRMActivated(context.Context, *IsWSRMActivatedRequest) (*IsWSRMActivatedResponse, error)

	RestoreExclusionList(context.Context, *RestoreExclusionListRequest) (*RestoreExclusionListResponse, error)
}

func RegisterConfigServer(conn dcerpc.Conn, o ConfigServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewConfigServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ConfigSyntaxV0_0))...)
}

func NewConfigServerHandle(o ConfigServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ConfigServerHandle(ctx, o, opNum, r)
	}
}

func ConfigServerHandle(ctx context.Context, o ConfigServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
		op := &xxx_WSRMActivateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &WSRMActivateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.WSRMActivate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // IsWSRMActivated
		op := &xxx_IsWSRMActivatedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IsWSRMActivatedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.IsWSRMActivated(ctx, req)
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
type UnimplementedConfigServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedConfigServer) GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigServer) SetConfig(context.Context, *SetConfigRequest) (*SetConfigResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigServer) IsEnabled(context.Context, *IsEnabledRequest) (*IsEnabledResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigServer) EnableDisable(context.Context, *EnableDisableRequest) (*EnableDisableResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigServer) GetExclusionList(context.Context, *GetExclusionListRequest) (*GetExclusionListResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigServer) SetExclusionList(context.Context, *SetExclusionListRequest) (*SetExclusionListResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigServer) WSRMActivate(context.Context, *WSRMActivateRequest) (*WSRMActivateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigServer) IsWSRMActivated(context.Context, *IsWSRMActivatedRequest) (*IsWSRMActivatedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigServer) RestoreExclusionList(context.Context, *RestoreExclusionListRequest) (*RestoreExclusionListResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ConfigServer = (*UnimplementedConfigServer)(nil)
