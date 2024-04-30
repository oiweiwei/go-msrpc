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
		in := &GetNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // Schema
		in := &GetSchemaRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSchema(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // CreateInstance
		in := &CreateInstanceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateInstance(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
