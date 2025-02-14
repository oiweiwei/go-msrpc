package bitspeerauth

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

// BitsPeerAuth server interface.
type BitsPeerAuthServer interface {

	// Return Values: An HRESULT indicating return status. See [MS-ERREF] for details of
	// the HRESULT type.
	//
	// ERROR_SUCCESS (0x00000000)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	ExchangePublicKeys(context.Context, *ExchangePublicKeysRequest) (*ExchangePublicKeysResponse, error)
}

func RegisterBitsPeerAuthServer(conn dcerpc.Conn, o BitsPeerAuthServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewBitsPeerAuthServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(BitsPeerAuthSyntaxV1_0))...)
}

func NewBitsPeerAuthServerHandle(o BitsPeerAuthServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return BitsPeerAuthServerHandle(ctx, o, opNum, r)
	}
}

func BitsPeerAuthServerHandle(ctx context.Context, o BitsPeerAuthServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // ExchangePublicKeys
		op := &xxx_ExchangePublicKeysOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ExchangePublicKeysRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ExchangePublicKeys(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented BitsPeerAuth
type UnimplementedBitsPeerAuthServer struct {
}

func (UnimplementedBitsPeerAuthServer) ExchangePublicKeys(context.Context, *ExchangePublicKeysRequest) (*ExchangePublicKeysResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ BitsPeerAuthServer = (*UnimplementedBitsPeerAuthServer)(nil)
