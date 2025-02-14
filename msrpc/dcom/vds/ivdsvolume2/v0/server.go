package ivdsvolume2

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

// IVdsVolume2 server interface.
type Volume2Server interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// GetProperties2 operation.
	GetProperties2(context.Context, *GetProperties2Request) (*GetProperties2Response, error)
}

func RegisterVolume2Server(conn dcerpc.Conn, o Volume2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewVolume2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Volume2SyntaxV0_0))...)
}

func NewVolume2ServerHandle(o Volume2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Volume2ServerHandle(ctx, o, opNum, r)
	}
}

func Volume2ServerHandle(ctx context.Context, o Volume2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetProperties2
		op := &xxx_GetProperties2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetProperties2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetProperties2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IVdsVolume2
type UnimplementedVolume2Server struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedVolume2Server) GetProperties2(context.Context, *GetProperties2Request) (*GetProperties2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ Volume2Server = (*UnimplementedVolume2Server)(nil)
