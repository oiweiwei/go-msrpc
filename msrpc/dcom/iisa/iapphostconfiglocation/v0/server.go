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
		op := &xxx_GetPathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // Count
		op := &xxx_GetCountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // Item
		op := &xxx_GetItemOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetItemRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetItem(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // AddConfigSection
		op := &xxx_AddConfigSectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddConfigSectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddConfigSection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // DeleteConfigSection
		op := &xxx_DeleteConfigSectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteConfigSectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteConfigSection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IAppHostConfigLocation
type UnimplementedAppHostConfigLocationServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedAppHostConfigLocationServer) GetPath(context.Context, *GetPathRequest) (*GetPathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostConfigLocationServer) GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostConfigLocationServer) GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostConfigLocationServer) AddConfigSection(context.Context, *AddConfigSectionRequest) (*AddConfigSectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostConfigLocationServer) DeleteConfigSection(context.Context, *DeleteConfigSectionRequest) (*DeleteConfigSectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AppHostConfigLocationServer = (*UnimplementedAppHostConfigLocationServer)(nil)
