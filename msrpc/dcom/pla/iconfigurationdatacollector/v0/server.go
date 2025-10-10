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
		op := &xxx_GetFileMaxCountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFileMaxCountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFileMaxCount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 33: // FileMaxCount
		op := &xxx_SetFileMaxCountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFileMaxCountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFileMaxCount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 34: // FileMaxRecursiveDepth
		op := &xxx_GetFileMaxRecursiveDepthOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFileMaxRecursiveDepthRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFileMaxRecursiveDepth(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 35: // FileMaxRecursiveDepth
		op := &xxx_SetFileMaxRecursiveDepthOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFileMaxRecursiveDepthRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFileMaxRecursiveDepth(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 36: // FileMaxTotalSize
		op := &xxx_GetFileMaxTotalSizeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFileMaxTotalSizeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFileMaxTotalSize(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 37: // FileMaxTotalSize
		op := &xxx_SetFileMaxTotalSizeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFileMaxTotalSizeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFileMaxTotalSize(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 38: // Files
		op := &xxx_GetFilesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFilesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFiles(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 39: // Files
		op := &xxx_SetFilesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFilesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFiles(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 40: // ManagementQueries
		op := &xxx_GetManagementQueriesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetManagementQueriesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetManagementQueries(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 41: // ManagementQueries
		op := &xxx_SetManagementQueriesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetManagementQueriesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetManagementQueries(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 42: // QueryNetworkAdapters
		op := &xxx_GetQueryNetworkAdaptersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetQueryNetworkAdaptersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetQueryNetworkAdapters(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 43: // QueryNetworkAdapters
		op := &xxx_SetQueryNetworkAdaptersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetQueryNetworkAdaptersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetQueryNetworkAdapters(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 44: // RegistryKeys
		op := &xxx_GetRegistryKeysOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRegistryKeysRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRegistryKeys(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 45: // RegistryKeys
		op := &xxx_SetRegistryKeysOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetRegistryKeysRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetRegistryKeys(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 46: // RegistryMaxRecursiveDepth
		op := &xxx_GetRegistryMaxRecursiveDepthOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRegistryMaxRecursiveDepthRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRegistryMaxRecursiveDepth(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 47: // RegistryMaxRecursiveDepth
		op := &xxx_SetRegistryMaxRecursiveDepthOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetRegistryMaxRecursiveDepthRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetRegistryMaxRecursiveDepth(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 48: // SystemStateFile
		op := &xxx_GetSystemStateFileOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSystemStateFileRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSystemStateFile(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 49: // SystemStateFile
		op := &xxx_SetSystemStateFileOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSystemStateFileRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSystemStateFile(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IConfigurationDataCollector
type UnimplementedConfigurationDataCollectorServer struct {
	idatacollector.UnimplementedDataCollectorServer
}

func (UnimplementedConfigurationDataCollectorServer) GetFileMaxCount(context.Context, *GetFileMaxCountRequest) (*GetFileMaxCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigurationDataCollectorServer) SetFileMaxCount(context.Context, *SetFileMaxCountRequest) (*SetFileMaxCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigurationDataCollectorServer) GetFileMaxRecursiveDepth(context.Context, *GetFileMaxRecursiveDepthRequest) (*GetFileMaxRecursiveDepthResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigurationDataCollectorServer) SetFileMaxRecursiveDepth(context.Context, *SetFileMaxRecursiveDepthRequest) (*SetFileMaxRecursiveDepthResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigurationDataCollectorServer) GetFileMaxTotalSize(context.Context, *GetFileMaxTotalSizeRequest) (*GetFileMaxTotalSizeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigurationDataCollectorServer) SetFileMaxTotalSize(context.Context, *SetFileMaxTotalSizeRequest) (*SetFileMaxTotalSizeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigurationDataCollectorServer) GetFiles(context.Context, *GetFilesRequest) (*GetFilesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigurationDataCollectorServer) SetFiles(context.Context, *SetFilesRequest) (*SetFilesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigurationDataCollectorServer) GetManagementQueries(context.Context, *GetManagementQueriesRequest) (*GetManagementQueriesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigurationDataCollectorServer) SetManagementQueries(context.Context, *SetManagementQueriesRequest) (*SetManagementQueriesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigurationDataCollectorServer) GetQueryNetworkAdapters(context.Context, *GetQueryNetworkAdaptersRequest) (*GetQueryNetworkAdaptersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigurationDataCollectorServer) SetQueryNetworkAdapters(context.Context, *SetQueryNetworkAdaptersRequest) (*SetQueryNetworkAdaptersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigurationDataCollectorServer) GetRegistryKeys(context.Context, *GetRegistryKeysRequest) (*GetRegistryKeysResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigurationDataCollectorServer) SetRegistryKeys(context.Context, *SetRegistryKeysRequest) (*SetRegistryKeysResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigurationDataCollectorServer) GetRegistryMaxRecursiveDepth(context.Context, *GetRegistryMaxRecursiveDepthRequest) (*GetRegistryMaxRecursiveDepthResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigurationDataCollectorServer) SetRegistryMaxRecursiveDepth(context.Context, *SetRegistryMaxRecursiveDepthRequest) (*SetRegistryMaxRecursiveDepthResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigurationDataCollectorServer) GetSystemStateFile(context.Context, *GetSystemStateFileRequest) (*GetSystemStateFileResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConfigurationDataCollectorServer) SetSystemStateFile(context.Context, *SetSystemStateFileRequest) (*SetSystemStateFileResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ConfigurationDataCollectorServer = (*UnimplementedConfigurationDataCollectorServer)(nil)
