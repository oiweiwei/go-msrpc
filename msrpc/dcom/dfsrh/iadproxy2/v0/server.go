package iadproxy2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iadproxy "github.com/oiweiwei/go-msrpc/msrpc/dcom/dfsrh/iadproxy/v0"
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
	_ = iadproxy.GoPackage
)

// IADProxy2 server interface.
type IADProxy2Server interface {

	// IADProxy base class.
	iadproxy.IADProxyServer

	// CreateObject2 operation.
	CreateObject2(context.Context, *CreateObject2Request) (*CreateObject2Response, error)

	// DeleteObject2 operation.
	DeleteObject2(context.Context, *DeleteObject2Request) (*DeleteObject2Response, error)

	// ModifyObject2 operation.
	ModifyObject2(context.Context, *ModifyObject2Request) (*ModifyObject2Response, error)
}

func RegisterIADProxy2Server(conn dcerpc.Conn, o IADProxy2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewIADProxy2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(IADProxy2SyntaxV0_0))...)
}

func NewIADProxy2ServerHandle(o IADProxy2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return IADProxy2ServerHandle(ctx, o, opNum, r)
	}
}

func IADProxy2ServerHandle(ctx context.Context, o IADProxy2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 6 {
		// IADProxy base method.
		return iadproxy.IADProxyServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 6: // CreateObject2
		in := &CreateObject2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateObject2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // DeleteObject2
		in := &DeleteObject2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteObject2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // ModifyObject2
		in := &ModifyObject2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ModifyObject2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
