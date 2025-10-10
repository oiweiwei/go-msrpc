package iupdatesearcher3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iupdatesearcher2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdatesearcher2/v0"
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
	_ = iupdatesearcher2.GoPackage
)

// IUpdateSearcher3 server interface.
type UpdateSearcher3Server interface {

	// IUpdateSearcher2 base class.
	iupdatesearcher2.UpdateSearcher2Server

	GetSearchScope(context.Context, *GetSearchScopeRequest) (*GetSearchScopeResponse, error)

	SetSearchScope(context.Context, *SetSearchScopeRequest) (*SetSearchScopeResponse, error)
}

func RegisterUpdateSearcher3Server(conn dcerpc.Conn, o UpdateSearcher3Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewUpdateSearcher3ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(UpdateSearcher3SyntaxV0_0))...)
}

func NewUpdateSearcher3ServerHandle(o UpdateSearcher3Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return UpdateSearcher3ServerHandle(ctx, o, opNum, r)
	}
}

func UpdateSearcher3ServerHandle(ctx context.Context, o UpdateSearcher3Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 27 {
		// IUpdateSearcher2 base method.
		return iupdatesearcher2.UpdateSearcher2ServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 27: // SearchScope
		op := &xxx_GetSearchScopeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSearchScopeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSearchScope(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 28: // SearchScope
		op := &xxx_SetSearchScopeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSearchScopeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSearchScope(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IUpdateSearcher3
type UnimplementedUpdateSearcher3Server struct {
	iupdatesearcher2.UnimplementedUpdateSearcher2Server
}

func (UnimplementedUpdateSearcher3Server) GetSearchScope(context.Context, *GetSearchScopeRequest) (*GetSearchScopeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSearcher3Server) SetSearchScope(context.Context, *SetSearchScopeRequest) (*SetSearchScopeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ UpdateSearcher3Server = (*UnimplementedUpdateSearcher3Server)(nil)
