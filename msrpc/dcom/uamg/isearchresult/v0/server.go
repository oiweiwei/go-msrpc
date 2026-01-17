package isearchresult

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = idispatch.GoPackage
)

// ISearchResult server interface.
type SearchResultServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The ISearchResult::ResultCode (opnum 8) method retrieves the result code for the
	// search.
	//
	// The IUpdateHistoryEntry::ResultCode (opnum 9) method describes the result of the
	// operation.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the ResultCode ADM element.
	GetResultCode(context.Context, *GetResultCodeRequest) (*GetResultCodeResponse, error)

	// The ISearchResult::RootCategories (opnum 9) method retrieves the root categories
	// found during the search.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetRootCategories(context.Context, *GetRootCategoriesRequest) (*GetRootCategoriesResponse, error)

	// The ICategory::Updates (opnum 16) method retrieves an IUpdateCollection interface
	// containing the updates that belong to the update category.
	//
	// The ISearchResult::Updates (opnum 10) method retrieves the updates found during the
	// search.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the Updates ADM element.
	GetUpdates(context.Context, *GetUpdatesRequest) (*GetUpdatesResponse, error)

	// The ISearchResult::Warnings (opnum 11) method retrieves the warnings generated during
	// the search.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the Warnings ADM element.
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
