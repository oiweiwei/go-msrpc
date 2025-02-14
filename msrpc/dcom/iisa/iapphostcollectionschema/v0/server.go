package iapphostcollectionschema

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

// IAppHostCollectionSchema server interface.
type AppHostCollectionSchemaServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// AddElementNames operation.
	GetAddElementNames(context.Context, *GetAddElementNamesRequest) (*GetAddElementNamesResponse, error)

	// GetAddElementSchema operation.
	GetAddElementSchema(context.Context, *GetAddElementSchemaRequest) (*GetAddElementSchemaResponse, error)

	// RemoveElementSchema operation.
	GetRemoveElementSchema(context.Context, *GetRemoveElementSchemaRequest) (*GetRemoveElementSchemaResponse, error)

	// ClearElementSchema operation.
	GetClearElementSchema(context.Context, *GetClearElementSchemaRequest) (*GetClearElementSchemaResponse, error)

	// IsMergeAppend operation.
	GetIsMergeAppend(context.Context, *GetIsMergeAppendRequest) (*GetIsMergeAppendResponse, error)

	// GetMetadata operation.
	GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error)

	// DoesAllowDuplicates operation.
	GetDoesAllowDuplicates(context.Context, *GetDoesAllowDuplicatesRequest) (*GetDoesAllowDuplicatesResponse, error)
}

func RegisterAppHostCollectionSchemaServer(conn dcerpc.Conn, o AppHostCollectionSchemaServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostCollectionSchemaServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostCollectionSchemaSyntaxV0_0))...)
}

func NewAppHostCollectionSchemaServerHandle(o AppHostCollectionSchemaServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostCollectionSchemaServerHandle(ctx, o, opNum, r)
	}
}

func AppHostCollectionSchemaServerHandle(ctx context.Context, o AppHostCollectionSchemaServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // AddElementNames
		op := &xxx_GetAddElementNamesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAddElementNamesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAddElementNames(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // GetAddElementSchema
		op := &xxx_GetAddElementSchemaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAddElementSchemaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAddElementSchema(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // RemoveElementSchema
		op := &xxx_GetRemoveElementSchemaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRemoveElementSchemaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRemoveElementSchema(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // ClearElementSchema
		op := &xxx_GetClearElementSchemaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetClearElementSchemaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetClearElementSchema(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // IsMergeAppend
		op := &xxx_GetIsMergeAppendOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsMergeAppendRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsMergeAppend(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // GetMetadata
		op := &xxx_GetMetadataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMetadataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMetadata(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // DoesAllowDuplicates
		op := &xxx_GetDoesAllowDuplicatesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDoesAllowDuplicatesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDoesAllowDuplicates(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IAppHostCollectionSchema
type UnimplementedAppHostCollectionSchemaServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedAppHostCollectionSchemaServer) GetAddElementNames(context.Context, *GetAddElementNamesRequest) (*GetAddElementNamesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostCollectionSchemaServer) GetAddElementSchema(context.Context, *GetAddElementSchemaRequest) (*GetAddElementSchemaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostCollectionSchemaServer) GetRemoveElementSchema(context.Context, *GetRemoveElementSchemaRequest) (*GetRemoveElementSchemaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostCollectionSchemaServer) GetClearElementSchema(context.Context, *GetClearElementSchemaRequest) (*GetClearElementSchemaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostCollectionSchemaServer) GetIsMergeAppend(context.Context, *GetIsMergeAppendRequest) (*GetIsMergeAppendResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostCollectionSchemaServer) GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostCollectionSchemaServer) GetDoesAllowDuplicates(context.Context, *GetDoesAllowDuplicatesRequest) (*GetDoesAllowDuplicatesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AppHostCollectionSchemaServer = (*UnimplementedAppHostCollectionSchemaServer)(nil)
