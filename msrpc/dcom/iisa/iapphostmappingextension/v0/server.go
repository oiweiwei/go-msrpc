package iapphostmappingextension

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

// IAppHostMappingExtension server interface.
type AppHostMappingExtensionServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// GetSiteNameFromSiteId operation.
	GetSiteNameFromSiteID(context.Context, *GetSiteNameFromSiteIDRequest) (*GetSiteNameFromSiteIDResponse, error)

	// GetSiteIdFromSiteName operation.
	GetSiteIDFromSiteName(context.Context, *GetSiteIDFromSiteNameRequest) (*GetSiteIDFromSiteNameResponse, error)

	// GetSiteElementFromSiteId operation.
	GetSiteElementFromSiteID(context.Context, *GetSiteElementFromSiteIDRequest) (*GetSiteElementFromSiteIDResponse, error)

	// MapPath operation.
	MapPath(context.Context, *MapPathRequest) (*MapPathResponse, error)
}

func RegisterAppHostMappingExtensionServer(conn dcerpc.Conn, o AppHostMappingExtensionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostMappingExtensionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostMappingExtensionSyntaxV0_0))...)
}

func NewAppHostMappingExtensionServerHandle(o AppHostMappingExtensionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostMappingExtensionServerHandle(ctx, o, opNum, r)
	}
}

func AppHostMappingExtensionServerHandle(ctx context.Context, o AppHostMappingExtensionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetSiteNameFromSiteId
		op := &xxx_GetSiteNameFromSiteIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSiteNameFromSiteIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSiteNameFromSiteID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // GetSiteIdFromSiteName
		op := &xxx_GetSiteIDFromSiteNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSiteIDFromSiteNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSiteIDFromSiteName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // GetSiteElementFromSiteId
		op := &xxx_GetSiteElementFromSiteIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSiteElementFromSiteIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSiteElementFromSiteID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // MapPath
		op := &xxx_MapPathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MapPathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MapPath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IAppHostMappingExtension
type UnimplementedAppHostMappingExtensionServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedAppHostMappingExtensionServer) GetSiteNameFromSiteID(context.Context, *GetSiteNameFromSiteIDRequest) (*GetSiteNameFromSiteIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostMappingExtensionServer) GetSiteIDFromSiteName(context.Context, *GetSiteIDFromSiteNameRequest) (*GetSiteIDFromSiteNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostMappingExtensionServer) GetSiteElementFromSiteID(context.Context, *GetSiteElementFromSiteIDRequest) (*GetSiteElementFromSiteIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostMappingExtensionServer) MapPath(context.Context, *MapPathRequest) (*MapPathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AppHostMappingExtensionServer = (*UnimplementedAppHostMappingExtensionServer)(nil)
