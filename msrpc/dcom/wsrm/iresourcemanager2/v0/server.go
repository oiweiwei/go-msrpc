package iresourcemanager2

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

// IResourceManager2 server interface.
type ResourceManager2Server interface {

	// IDispatch base class.
	idispatch.DispatchServer

	ExportObjects(context.Context, *ExportObjectsRequest) (*ExportObjectsResponse, error)

	GetImportConflicts(context.Context, *GetImportConflictsRequest) (*GetImportConflictsResponse, error)

	ImportXML(context.Context, *ImportXMLRequest) (*ImportXMLResponse, error)

	ExportXML(context.Context, *ExportXMLRequest) (*ExportXMLResponse, error)
}

func RegisterResourceManager2Server(conn dcerpc.Conn, o ResourceManager2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewResourceManager2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ResourceManager2SyntaxV0_0))...)
}

func NewResourceManager2ServerHandle(o ResourceManager2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ResourceManager2ServerHandle(ctx, o, opNum, r)
	}
}

func ResourceManager2ServerHandle(ctx context.Context, o ResourceManager2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // ExportObjects
		op := &xxx_ExportObjectsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ExportObjectsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ExportObjects(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // GetImportConflicts
		op := &xxx_GetImportConflictsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetImportConflictsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetImportConflicts(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // ImportXml
		op := &xxx_ImportXMLOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ImportXMLRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ImportXML(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // ExportXml
		op := &xxx_ExportXMLOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ExportXMLRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ExportXML(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IResourceManager2
type UnimplementedResourceManager2Server struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedResourceManager2Server) ExportObjects(context.Context, *ExportObjectsRequest) (*ExportObjectsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedResourceManager2Server) GetImportConflicts(context.Context, *GetImportConflictsRequest) (*GetImportConflictsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedResourceManager2Server) ImportXML(context.Context, *ImportXMLRequest) (*ImportXMLResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedResourceManager2Server) ExportXML(context.Context, *ExportXMLRequest) (*ExportXMLResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ResourceManager2Server = (*UnimplementedResourceManager2Server)(nil)
