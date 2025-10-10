package iupdateservicemanager2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iupdateservicemanager "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdateservicemanager/v0"
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
	_ = iupdateservicemanager.GoPackage
)

// IUpdateServiceManager2 server interface.
type UpdateServiceManager2Server interface {

	// IUpdateServiceManager base class.
	iupdateservicemanager.UpdateServiceManagerServer

	GetClientApplicationID(context.Context, *GetClientApplicationIDRequest) (*GetClientApplicationIDResponse, error)

	SetClientApplicationID(context.Context, *SetClientApplicationIDRequest) (*SetClientApplicationIDResponse, error)

	QueryServiceRegistration(context.Context, *QueryServiceRegistrationRequest) (*QueryServiceRegistrationResponse, error)

	AddService2(context.Context, *AddService2Request) (*AddService2Response, error)
}

func RegisterUpdateServiceManager2Server(conn dcerpc.Conn, o UpdateServiceManager2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewUpdateServiceManager2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(UpdateServiceManager2SyntaxV0_0))...)
}

func NewUpdateServiceManager2ServerHandle(o UpdateServiceManager2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return UpdateServiceManager2ServerHandle(ctx, o, opNum, r)
	}
}

func UpdateServiceManager2ServerHandle(ctx context.Context, o UpdateServiceManager2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 14 {
		// IUpdateServiceManager base method.
		return iupdateservicemanager.UpdateServiceManagerServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 14: // ClientApplicationID
		op := &xxx_GetClientApplicationIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetClientApplicationIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetClientApplicationID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // ClientApplicationID
		op := &xxx_SetClientApplicationIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetClientApplicationIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetClientApplicationID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // QueryServiceRegistration
		op := &xxx_QueryServiceRegistrationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryServiceRegistrationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryServiceRegistration(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // AddService2
		op := &xxx_AddService2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddService2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddService2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IUpdateServiceManager2
type UnimplementedUpdateServiceManager2Server struct {
	iupdateservicemanager.UnimplementedUpdateServiceManagerServer
}

func (UnimplementedUpdateServiceManager2Server) GetClientApplicationID(context.Context, *GetClientApplicationIDRequest) (*GetClientApplicationIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServiceManager2Server) SetClientApplicationID(context.Context, *SetClientApplicationIDRequest) (*SetClientApplicationIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServiceManager2Server) QueryServiceRegistration(context.Context, *QueryServiceRegistrationRequest) (*QueryServiceRegistrationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServiceManager2Server) AddService2(context.Context, *AddService2Request) (*AddService2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ UpdateServiceManager2Server = (*UnimplementedUpdateServiceManager2Server)(nil)
