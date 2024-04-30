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
		in := &GetNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // Type
		in := &GetTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // Type
		in := &SetTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // OverrideModeDefault
		in := &GetOverrideModeDefaultRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetOverrideModeDefault(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // OverrideModeDefault
		in := &SetOverrideModeDefaultRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetOverrideModeDefault(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // AllowDefinition
		in := &GetAllowDefinitionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetAllowDefinition(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // AllowDefinition
		in := &SetAllowDefinitionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetAllowDefinition(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // AllowLocation
		in := &GetAllowLocationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetAllowLocation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // AllowLocation
		in := &SetAllowLocationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetAllowLocation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
