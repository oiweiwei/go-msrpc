package iapphostconfiglocation

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

// IAppHostConfigLocation server interface.
type AppHostConfigLocationServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// Path operation.
	GetPath(context.Context, *GetPathRequest) (*GetPathResponse, error)

	// Count operation.
	GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error)

	// Item operation.
	GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error)

	// AddConfigSection operation.
	AddConfigSection(context.Context, *AddConfigSectionRequest) (*AddConfigSectionResponse, error)

	// DeleteConfigSection operation.
	DeleteConfigSection(context.Context, *DeleteConfigSectionRequest) (*DeleteConfigSectionResponse, error)
}

func RegisterAppHostConfigLocationServer(conn dcerpc.Conn, o AppHostConfigLocationServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostConfigLocationServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostConfigLocationSyntaxV0_0))...)
}

func NewAppHostConfigLocationServerHandle(o AppHostConfigLocationServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostConfigLocationServerHandle(ctx, o, opNum, r)
	}
}

func AppHostConfigLocationServerHandle(ctx context.Context, o AppHostConfigLocationServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Path
		in := &GetPathRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPath(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // Count
		in := &GetCountRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetCount(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // Item
		in := &GetItemRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetItem(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // AddConfigSection
		in := &AddConfigSectionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddConfigSection(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // DeleteConfigSection
		in := &DeleteConfigSectionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteConfigSection(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
