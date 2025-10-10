package iupdate4

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iupdate3 "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdate3/v0"
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
	_ = iupdate3.GoPackage
)

// IUpdate4 server interface.
type Update4Server interface {

	// IUpdate3 base class.
	iupdate3.Update3Server

	GetPerUser(context.Context, *GetPerUserRequest) (*GetPerUserResponse, error)
}

func RegisterUpdate4Server(conn dcerpc.Conn, o Update4Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewUpdate4ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Update4SyntaxV0_0))...)
}

func NewUpdate4ServerHandle(o Update4Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Update4ServerHandle(ctx, o, opNum, r)
	}
}

func Update4ServerHandle(ctx context.Context, o Update4Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 57 {
		// IUpdate3 base method.
		return iupdate3.Update3ServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 57: // PerUser
		op := &xxx_GetPerUserOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPerUserRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPerUser(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IUpdate4
type UnimplementedUpdate4Server struct {
	iupdate3.UnimplementedUpdate3Server
}

func (UnimplementedUpdate4Server) GetPerUser(context.Context, *GetPerUserRequest) (*GetPerUserResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ Update4Server = (*UnimplementedUpdate4Server)(nil)
