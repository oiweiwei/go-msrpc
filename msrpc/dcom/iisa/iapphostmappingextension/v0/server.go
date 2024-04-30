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
		in := &GetSiteNameFromSiteIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSiteNameFromSiteID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // GetSiteIdFromSiteName
		in := &GetSiteIDFromSiteNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSiteIDFromSiteName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // GetSiteElementFromSiteId
		in := &GetSiteElementFromSiteIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSiteElementFromSiteID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // MapPath
		in := &MapPathRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.MapPath(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
