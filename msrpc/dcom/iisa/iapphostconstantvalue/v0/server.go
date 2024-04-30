package iapphostconstantvalue

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

// IAppHostConstantValue server interface.
type AppHostConstantValueServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// Name operation.
	GetName(context.Context, *GetNameRequest) (*GetNameResponse, error)

	// Value operation.
	GetValue(context.Context, *GetValueRequest) (*GetValueResponse, error)
}

func RegisterAppHostConstantValueServer(conn dcerpc.Conn, o AppHostConstantValueServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostConstantValueServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostConstantValueSyntaxV0_0))...)
}

func NewAppHostConstantValueServerHandle(o AppHostConstantValueServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostConstantValueServerHandle(ctx, o, opNum, r)
	}
}

func AppHostConstantValueServerHandle(ctx context.Context, o AppHostConstantValueServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
	case 4: // Value
		in := &GetValueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetValue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
