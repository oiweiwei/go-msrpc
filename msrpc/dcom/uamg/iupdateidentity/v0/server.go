package iupdateidentity

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

// IUpdateIdentity server interface.
type UpdateIdentityServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	GetRevisionNumber(context.Context, *GetRevisionNumberRequest) (*GetRevisionNumberResponse, error)

	GetUpdateID(context.Context, *GetUpdateIDRequest) (*GetUpdateIDResponse, error)
}

func RegisterUpdateIdentityServer(conn dcerpc.Conn, o UpdateIdentityServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewUpdateIdentityServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(UpdateIdentitySyntaxV0_0))...)
}

func NewUpdateIdentityServerHandle(o UpdateIdentityServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return UpdateIdentityServerHandle(ctx, o, opNum, r)
	}
}

func UpdateIdentityServerHandle(ctx context.Context, o UpdateIdentityServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // RevisionNumber
		op := &xxx_GetRevisionNumberOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRevisionNumberRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRevisionNumber(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // UpdateID
		op := &xxx_GetUpdateIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetUpdateIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetUpdateID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IUpdateIdentity
type UnimplementedUpdateIdentityServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedUpdateIdentityServer) GetRevisionNumber(context.Context, *GetRevisionNumberRequest) (*GetRevisionNumberResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateIdentityServer) GetUpdateID(context.Context, *GetUpdateIDRequest) (*GetUpdateIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ UpdateIdentityServer = (*UnimplementedUpdateIdentityServer)(nil)
