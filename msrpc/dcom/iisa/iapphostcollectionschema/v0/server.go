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
		in := &GetAddElementNamesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetAddElementNames(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // GetAddElementSchema
		in := &GetAddElementSchemaRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetAddElementSchema(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // RemoveElementSchema
		in := &GetRemoveElementSchemaRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetRemoveElementSchema(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // ClearElementSchema
		in := &GetClearElementSchemaRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetClearElementSchema(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // IsMergeAppend
		in := &GetIsMergeAppendRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetIsMergeAppend(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // GetMetadata
		in := &GetMetadataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMetadata(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // DoesAllowDuplicates
		in := &GetDoesAllowDuplicatesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDoesAllowDuplicates(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
