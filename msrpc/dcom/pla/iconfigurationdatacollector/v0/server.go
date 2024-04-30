package iconfigurationdatacollector

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	idatacollector "github.com/oiweiwei/go-msrpc/msrpc/dcom/pla/idatacollector/v0"
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
	_ = idatacollector.GoPackage
)

// IConfigurationDataCollector server interface.
type ConfigurationDataCollectorServer interface {

	// IDataCollector base class.
	idatacollector.DataCollectorServer

	// FileMaxCount operation.
	GetFileMaxCount(context.Context, *GetFileMaxCountRequest) (*GetFileMaxCountResponse, error)

	// FileMaxCount operation.
	SetFileMaxCount(context.Context, *SetFileMaxCountRequest) (*SetFileMaxCountResponse, error)

	// FileMaxRecursiveDepth operation.
	GetFileMaxRecursiveDepth(context.Context, *GetFileMaxRecursiveDepthRequest) (*GetFileMaxRecursiveDepthResponse, error)

	// FileMaxRecursiveDepth operation.
	SetFileMaxRecursiveDepth(context.Context, *SetFileMaxRecursiveDepthRequest) (*SetFileMaxRecursiveDepthResponse, error)

	// FileMaxTotalSize operation.
	GetFileMaxTotalSize(context.Context, *GetFileMaxTotalSizeRequest) (*GetFileMaxTotalSizeResponse, error)

	// FileMaxTotalSize operation.
	SetFileMaxTotalSize(context.Context, *SetFileMaxTotalSizeRequest) (*SetFileMaxTotalSizeResponse, error)

	// Files operation.
	GetFiles(context.Context, *GetFilesRequest) (*GetFilesResponse, error)

	// Files operation.
	SetFiles(context.Context, *SetFilesRequest) (*SetFilesResponse, error)

	// ManagementQueries operation.
	GetManagementQueries(context.Context, *GetManagementQueriesRequest) (*GetManagementQueriesResponse, error)

	// ManagementQueries operation.
	SetManagementQueries(context.Context, *SetManagementQueriesRequest) (*SetManagementQueriesResponse, error)

	// QueryNetworkAdapters operation.
	GetQueryNetworkAdapters(context.Context, *GetQueryNetworkAdaptersRequest) (*GetQueryNetworkAdaptersResponse, error)

	// QueryNetworkAdapters operation.
	SetQueryNetworkAdapters(context.Context, *SetQueryNetworkAdaptersRequest) (*SetQueryNetworkAdaptersResponse, error)

	// RegistryKeys operation.
	GetRegistryKeys(context.Context, *GetRegistryKeysRequest) (*GetRegistryKeysResponse, error)

	// RegistryKeys operation.
	SetRegistryKeys(context.Context, *SetRegistryKeysRequest) (*SetRegistryKeysResponse, error)

	// RegistryMaxRecursiveDepth operation.
	GetRegistryMaxRecursiveDepth(context.Context, *GetRegistryMaxRecursiveDepthRequest) (*GetRegistryMaxRecursiveDepthResponse, error)

	// RegistryMaxRecursiveDepth operation.
	SetRegistryMaxRecursiveDepth(context.Context, *SetRegistryMaxRecursiveDepthRequest) (*SetRegistryMaxRecursiveDepthResponse, error)

	// SystemStateFile operation.
	GetSystemStateFile(context.Context, *GetSystemStateFileRequest) (*GetSystemStateFileResponse, error)

	// SystemStateFile operation.
	SetSystemStateFile(context.Context, *SetSystemStateFileRequest) (*SetSystemStateFileResponse, error)
}

func RegisterConfigurationDataCollectorServer(conn dcerpc.Conn, o ConfigurationDataCollectorServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewConfigurationDataCollectorServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ConfigurationDataCollectorSyntaxV0_0))...)
}

func NewConfigurationDataCollectorServerHandle(o ConfigurationDataCollectorServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ConfigurationDataCollectorServerHandle(ctx, o, opNum, r)
	}
}

func ConfigurationDataCollectorServerHandle(ctx context.Context, o ConfigurationDataCollectorServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 32 {
		// IDataCollector base method.
		return idatacollector.DataCollectorServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 32: // FileMaxCount
		in := &GetFileMaxCountRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFileMaxCount(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 33: // FileMaxCount
		in := &SetFileMaxCountRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetFileMaxCount(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 34: // FileMaxRecursiveDepth
		in := &GetFileMaxRecursiveDepthRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFileMaxRecursiveDepth(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 35: // FileMaxRecursiveDepth
		in := &SetFileMaxRecursiveDepthRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetFileMaxRecursiveDepth(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 36: // FileMaxTotalSize
		in := &GetFileMaxTotalSizeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFileMaxTotalSize(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 37: // FileMaxTotalSize
		in := &SetFileMaxTotalSizeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetFileMaxTotalSize(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 38: // Files
		in := &GetFilesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFiles(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 39: // Files
		in := &SetFilesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetFiles(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 40: // ManagementQueries
		in := &GetManagementQueriesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetManagementQueries(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 41: // ManagementQueries
		in := &SetManagementQueriesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetManagementQueries(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 42: // QueryNetworkAdapters
		in := &GetQueryNetworkAdaptersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetQueryNetworkAdapters(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 43: // QueryNetworkAdapters
		in := &SetQueryNetworkAdaptersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetQueryNetworkAdapters(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 44: // RegistryKeys
		in := &GetRegistryKeysRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetRegistryKeys(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 45: // RegistryKeys
		in := &SetRegistryKeysRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetRegistryKeys(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 46: // RegistryMaxRecursiveDepth
		in := &GetRegistryMaxRecursiveDepthRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetRegistryMaxRecursiveDepth(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 47: // RegistryMaxRecursiveDepth
		in := &SetRegistryMaxRecursiveDepthRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetRegistryMaxRecursiveDepth(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 48: // SystemStateFile
		in := &GetSystemStateFileRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSystemStateFile(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 49: // SystemStateFile
		in := &SetSystemStateFileRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSystemStateFile(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
