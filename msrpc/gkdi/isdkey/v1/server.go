package isdkey

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
)

// ISDKey server interface.
type ISDKeyServer interface {

	// The syntax for the GetKey (Opnum 0) method consists of the following.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// received from the client; otherwise, it MUST return a nonzero value.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetKey(context.Context, *GetKeyRequest) (*GetKeyResponse, error)
}

func RegisterISDKeyServer(conn dcerpc.Conn, o ISDKeyServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewISDKeyServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ISDKeySyntaxV1_0))...)
}

func NewISDKeyServerHandle(o ISDKeyServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ISDKeyServerHandle(ctx, o, opNum, r)
	}
}

func ISDKeyServerHandle(ctx context.Context, o ISDKeyServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // GetKey
		in := &GetKeyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetKey(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
