package iapphostsectiondefinition

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

// IAppHostSectionDefinition server interface.
type AppHostSectionDefinitionServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// Name operation.
	GetName(context.Context, *GetNameRequest) (*GetNameResponse, error)

	// Type operation.
	GetType(context.Context, *GetTypeRequest) (*GetTypeResponse, error)

	// Type operation.
	SetType(context.Context, *SetTypeRequest) (*SetTypeResponse, error)

	// OverrideModeDefault operation.
	GetOverrideModeDefault(context.Context, *GetOverrideModeDefaultRequest) (*GetOverrideModeDefaultResponse, error)

	// OverrideModeDefault operation.
	SetOverrideModeDefault(context.Context, *SetOverrideModeDefaultRequest) (*SetOverrideModeDefaultResponse, error)

	// AllowDefinition operation.
	GetAllowDefinition(context.Context, *GetAllowDefinitionRequest) (*GetAllowDefinitionResponse, error)

	// AllowDefinition operation.
	SetAllowDefinition(context.Context, *SetAllowDefinitionRequest) (*SetAllowDefinitionResponse, error)

	// AllowLocation operation.
	GetAllowLocation(context.Context, *GetAllowLocationRequest) (*GetAllowLocationResponse, error)

	// AllowLocation operation.
	SetAllowLocation(context.Context, *SetAllowLocationRequest) (*SetAllowLocationResponse, error)
}

func RegisterAppHostSectionDefinitionServer(conn dcerpc.Conn, o AppHostSectionDefinitionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostSectionDefinitionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostSectionDefinitionSyntaxV0_0))...)
}

func NewAppHostSectionDefinitionServerHandle(o AppHostSectionDefinitionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostSectionDefinitionServerHandle(ctx, o, opNum, r)
	}
}

func AppHostSectionDefinitionServerHandle(ctx context.Context, o AppHostSectionDefinitionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Name
		op := &xxx_GetNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // Type
		op := &xxx_GetTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // Type
		op := &xxx_SetTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // OverrideModeDefault
		op := &xxx_GetOverrideModeDefaultOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetOverrideModeDefaultRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetOverrideModeDefault(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // OverrideModeDefault
		op := &xxx_SetOverrideModeDefaultOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetOverrideModeDefaultRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetOverrideModeDefault(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // AllowDefinition
		op := &xxx_GetAllowDefinitionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAllowDefinitionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAllowDefinition(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // AllowDefinition
		op := &xxx_SetAllowDefinitionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetAllowDefinitionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetAllowDefinition(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // AllowLocation
		op := &xxx_GetAllowLocationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAllowLocationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAllowLocation(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // AllowLocation
		op := &xxx_SetAllowLocationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetAllowLocationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetAllowLocation(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IAppHostSectionDefinition
type UnimplementedAppHostSectionDefinitionServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedAppHostSectionDefinitionServer) GetName(context.Context, *GetNameRequest) (*GetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostSectionDefinitionServer) GetType(context.Context, *GetTypeRequest) (*GetTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostSectionDefinitionServer) SetType(context.Context, *SetTypeRequest) (*SetTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostSectionDefinitionServer) GetOverrideModeDefault(context.Context, *GetOverrideModeDefaultRequest) (*GetOverrideModeDefaultResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostSectionDefinitionServer) SetOverrideModeDefault(context.Context, *SetOverrideModeDefaultRequest) (*SetOverrideModeDefaultResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostSectionDefinitionServer) GetAllowDefinition(context.Context, *GetAllowDefinitionRequest) (*GetAllowDefinitionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostSectionDefinitionServer) SetAllowDefinition(context.Context, *SetAllowDefinitionRequest) (*SetAllowDefinitionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostSectionDefinitionServer) GetAllowLocation(context.Context, *GetAllowLocationRequest) (*GetAllowLocationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostSectionDefinitionServer) SetAllowLocation(context.Context, *SetAllowLocationRequest) (*SetAllowLocationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AppHostSectionDefinitionServer = (*UnimplementedAppHostSectionDefinitionServer)(nil)
