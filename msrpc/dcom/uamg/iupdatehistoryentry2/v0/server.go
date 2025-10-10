package iupdatehistoryentry2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iupdatehistoryentry "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdatehistoryentry/v0"
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
	_ = iupdatehistoryentry.GoPackage
)

// IUpdateHistoryEntry2 server interface.
type UpdateHistoryEntry2Server interface {

	// IUpdateHistoryEntry base class.
	iupdatehistoryentry.UpdateHistoryEntryServer

	GetCategories(context.Context, *GetCategoriesRequest) (*GetCategoriesResponse, error)
}

func RegisterUpdateHistoryEntry2Server(conn dcerpc.Conn, o UpdateHistoryEntry2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewUpdateHistoryEntry2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(UpdateHistoryEntry2SyntaxV0_0))...)
}

func NewUpdateHistoryEntry2ServerHandle(o UpdateHistoryEntry2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return UpdateHistoryEntry2ServerHandle(ctx, o, opNum, r)
	}
}

func UpdateHistoryEntry2ServerHandle(ctx context.Context, o UpdateHistoryEntry2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 21 {
		// IUpdateHistoryEntry base method.
		return iupdatehistoryentry.UpdateHistoryEntryServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 21: // Categories
		op := &xxx_GetCategoriesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCategoriesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCategories(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IUpdateHistoryEntry2
type UnimplementedUpdateHistoryEntry2Server struct {
	iupdatehistoryentry.UnimplementedUpdateHistoryEntryServer
}

func (UnimplementedUpdateHistoryEntry2Server) GetCategories(context.Context, *GetCategoriesRequest) (*GetCategoriesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ UpdateHistoryEntry2Server = (*UnimplementedUpdateHistoryEntry2Server)(nil)
