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

	// The IUpdate4::PerUser (opnum 58) method retrieves whether the update is per user.
	//
	// The IWindowsDriverUpdate4::PerUser (opnum 67) method retrieves whether the update
	// is per user.
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
	// This method SHOULD return the value of the PerUser ADM element.
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
