package iapphostadminmanager

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

// IAppHostAdminManager server interface.
type AppHostAdminManagerServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// GetAdminSection operation.
	GetAdminSection(context.Context, *GetAdminSectionRequest) (*GetAdminSectionResponse, error)

	// GetMetadata operation.
	GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error)

	// SetMetadata operation.
	SetMetadata(context.Context, *SetMetadataRequest) (*SetMetadataResponse, error)

	// ConfigManager operation.
	GetConfigManager(context.Context, *GetConfigManagerRequest) (*GetConfigManagerResponse, error)
}

func RegisterAppHostAdminManagerServer(conn dcerpc.Conn, o AppHostAdminManagerServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostAdminManagerServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostAdminManagerSyntaxV0_0))...)
}

func NewAppHostAdminManagerServerHandle(o AppHostAdminManagerServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostAdminManagerServerHandle(ctx, o, opNum, r)
	}
}

func AppHostAdminManagerServerHandle(ctx context.Context, o AppHostAdminManagerServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetAdminSection
		in := &GetAdminSectionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetAdminSection(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // GetMetadata
		in := &GetMetadataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMetadata(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // SetMetadata
		in := &SetMetadataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetMetadata(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // ConfigManager
		in := &GetConfigManagerRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetConfigManager(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
