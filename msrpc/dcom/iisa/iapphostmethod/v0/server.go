package iapphostmethod

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

// IAppHostMethod server interface.
type AppHostMethodServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// Name operation.
	GetName(context.Context, *GetNameRequest) (*GetNameResponse, error)

	// Schema operation.
	GetSchema(context.Context, *GetSchemaRequest) (*GetSchemaResponse, error)

	// CreateInstance operation.
	CreateInstance(context.Context, *CreateInstanceRequest) (*CreateInstanceResponse, error)
}

func RegisterAppHostMethodServer(conn dcerpc.Conn, o AppHostMethodServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostMethodServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostMethodSyntaxV0_0))...)
}

func NewAppHostMethodServerHandle(o AppHostMethodServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostMethodServerHandle(ctx, o, opNum, r)
	}
}

func AppHostMethodServerHandle(ctx context.Context, o AppHostMethodServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Name
		op := &xxx_GetNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // Schema
		op := &xxx_GetSchemaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSchemaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSchema(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // CreateInstance
		op := &xxx_CreateInstanceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateInstanceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateInstance(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IAppHostMethod
type UnimplementedAppHostMethodServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedAppHostMethodServer) GetName(context.Context, *GetNameRequest) (*GetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostMethodServer) GetSchema(context.Context, *GetSchemaRequest) (*GetSchemaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostMethodServer) CreateInstance(context.Context, *CreateInstanceRequest) (*CreateInstanceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AppHostMethodServer = (*UnimplementedAppHostMethodServer)(nil)
