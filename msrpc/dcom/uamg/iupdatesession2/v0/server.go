package iupdatesession2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iupdatesession "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdatesession/v0"
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
	_ = iupdatesession.GoPackage
)

// IUpdateSession2 server interface.
type UpdateSession2Server interface {

	// IUpdateSession base class.
	iupdatesession.UpdateSessionServer

	// The IUpdateSession2::UserLocale (opnum 14) method gets the language for update results.
	//
	// The IUpdateSession2::UserLocale (opnum 15) method sets the language for update results.
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
	// This method SHOULD return the value of the UserLocale ADM element.
	GetUserLocale(context.Context, *GetUserLocaleRequest) (*GetUserLocaleResponse, error)

	// The IUpdateSession2::UserLocale (opnum 14) method gets the language for update results.
	//
	// The IUpdateSession2::UserLocale (opnum 15) method sets the language for update results.
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
	// This method SHOULD return the value of the UserLocale ADM element.
	SetUserLocale(context.Context, *SetUserLocaleRequest) (*SetUserLocaleResponse, error)
}

func RegisterUpdateSession2Server(conn dcerpc.Conn, o UpdateSession2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewUpdateSession2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(UpdateSession2SyntaxV0_0))...)
}

func NewUpdateSession2ServerHandle(o UpdateSession2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return UpdateSession2ServerHandle(ctx, o, opNum, r)
	}
}

func UpdateSession2ServerHandle(ctx context.Context, o UpdateSession2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 15 {
		// IUpdateSession base method.
		return iupdatesession.UpdateSessionServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 15: // UserLocale
		op := &xxx_GetUserLocaleOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetUserLocaleRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetUserLocale(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // UserLocale
		op := &xxx_SetUserLocaleOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetUserLocaleRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetUserLocale(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IUpdateSession2
type UnimplementedUpdateSession2Server struct {
	iupdatesession.UnimplementedUpdateSessionServer
}

func (UnimplementedUpdateSession2Server) GetUserLocale(context.Context, *GetUserLocaleRequest) (*GetUserLocaleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSession2Server) SetUserLocale(context.Context, *SetUserLocaleRequest) (*SetUserLocaleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ UpdateSession2Server = (*UnimplementedUpdateSession2Server)(nil)
