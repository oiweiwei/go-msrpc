package isearchresult

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

// ISearchResult server interface.
type SearchResultServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	GetResultCode(context.Context, *GetResultCodeRequest) (*GetResultCodeResponse, error)

	GetRootCategories(context.Context, *GetRootCategoriesRequest) (*GetRootCategoriesResponse, error)

	GetUpdates(context.Context, *GetUpdatesRequest) (*GetUpdatesResponse, error)

	GetWarnings(context.Context, *GetWarningsRequest) (*GetWarningsResponse, error)
}

func RegisterSearchResultServer(conn dcerpc.Conn, o SearchResultServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewSearchResultServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(SearchResultSyntaxV0_0))...)
}

func NewSearchResultServerHandle(o SearchResultServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return SearchResultServerHandle(ctx, o, opNum, r)
	}
}

func SearchResultServerHandle(ctx context.Context, o SearchResultServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // ResultCode
		op := &xxx_GetResultCodeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetResultCodeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetResultCode(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // RootCategories
		op := &xxx_GetRootCategoriesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRootCategoriesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRootCategories(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // Updates
		op := &xxx_GetUpdatesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetUpdatesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetUpdates(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // Warnings
		op := &xxx_GetWarningsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetWarningsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetWarnings(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented ISearchResult
type UnimplementedSearchResultServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedSearchResultServer) GetResultCode(context.Context, *GetResultCodeRequest) (*GetResultCodeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSearchResultServer) GetRootCategories(context.Context, *GetRootCategoriesRequest) (*GetRootCategoriesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSearchResultServer) GetUpdates(context.Context, *GetUpdatesRequest) (*GetUpdatesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSearchResultServer) GetWarnings(context.Context, *GetWarningsRequest) (*GetWarningsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ SearchResultServer = (*UnimplementedSearchResultServer)(nil)
