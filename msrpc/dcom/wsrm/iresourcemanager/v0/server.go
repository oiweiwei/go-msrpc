package iresourcemanager

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

// IResourceManager server interface.
type ResourceManagerServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	RetrieveEventList(context.Context, *RetrieveEventListRequest) (*RetrieveEventListResponse, error)

	GetSystemAffinity(context.Context, *GetSystemAffinityRequest) (*GetSystemAffinityResponse, error)

	ImportXMLFiles(context.Context, *ImportXMLFilesRequest) (*ImportXMLFilesResponse, error)

	ExportXMLFiles(context.Context, *ExportXMLFilesRequest) (*ExportXMLFilesResponse, error)

	RestoreXMLFiles(context.Context, *RestoreXMLFilesRequest) (*RestoreXMLFilesResponse, error)

	GetDependencies(context.Context, *GetDependenciesRequest) (*GetDependenciesResponse, error)

	GetServiceList(context.Context, *GetServiceListRequest) (*GetServiceListResponse, error)

	GetIISAppPoolNames(context.Context, *GetIISAppPoolNamesRequest) (*GetIISAppPoolNamesResponse, error)

	GetServerName(context.Context, *GetServerNameRequest) (*GetServerNameResponse, error)

	GetCurrentMemory(context.Context, *GetCurrentMemoryRequest) (*GetCurrentMemoryResponse, error)
}

func RegisterResourceManagerServer(conn dcerpc.Conn, o ResourceManagerServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewResourceManagerServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ResourceManagerSyntaxV0_0))...)
}

func NewResourceManagerServerHandle(o ResourceManagerServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ResourceManagerServerHandle(ctx, o, opNum, r)
	}
}

func ResourceManagerServerHandle(ctx context.Context, o ResourceManagerServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // RetrieveEventList
		op := &xxx_RetrieveEventListOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RetrieveEventListRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RetrieveEventList(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // GetSystemAffinity
		op := &xxx_GetSystemAffinityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSystemAffinityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSystemAffinity(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // ImportXMLFiles
		op := &xxx_ImportXMLFilesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ImportXMLFilesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ImportXMLFiles(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // ExportXMLFiles
		op := &xxx_ExportXMLFilesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ExportXMLFilesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ExportXMLFiles(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // RestoreXMLFiles
		op := &xxx_RestoreXMLFilesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RestoreXMLFilesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RestoreXMLFiles(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // GetDependencies
		op := &xxx_GetDependenciesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDependenciesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDependencies(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // GetServiceList
		op := &xxx_GetServiceListOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetServiceListRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetServiceList(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // GetIISAppPoolNames
		op := &xxx_GetIISAppPoolNamesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIISAppPoolNamesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIISAppPoolNames(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // GetServerName
		op := &xxx_GetServerNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetServerNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetServerName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // GetCurrentMemory
		op := &xxx_GetCurrentMemoryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCurrentMemoryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCurrentMemory(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IResourceManager
type UnimplementedResourceManagerServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedResourceManagerServer) RetrieveEventList(context.Context, *RetrieveEventListRequest) (*RetrieveEventListResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedResourceManagerServer) GetSystemAffinity(context.Context, *GetSystemAffinityRequest) (*GetSystemAffinityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedResourceManagerServer) ImportXMLFiles(context.Context, *ImportXMLFilesRequest) (*ImportXMLFilesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedResourceManagerServer) ExportXMLFiles(context.Context, *ExportXMLFilesRequest) (*ExportXMLFilesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedResourceManagerServer) RestoreXMLFiles(context.Context, *RestoreXMLFilesRequest) (*RestoreXMLFilesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedResourceManagerServer) GetDependencies(context.Context, *GetDependenciesRequest) (*GetDependenciesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedResourceManagerServer) GetServiceList(context.Context, *GetServiceListRequest) (*GetServiceListResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedResourceManagerServer) GetIISAppPoolNames(context.Context, *GetIISAppPoolNamesRequest) (*GetIISAppPoolNamesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedResourceManagerServer) GetServerName(context.Context, *GetServerNameRequest) (*GetServerNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedResourceManagerServer) GetCurrentMemory(context.Context, *GetCurrentMemoryRequest) (*GetCurrentMemoryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ResourceManagerServer = (*UnimplementedResourceManagerServer)(nil)
