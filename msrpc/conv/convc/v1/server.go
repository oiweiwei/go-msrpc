package convc

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
)

// convc server interface.
type ConvcServer interface {

	// convc_indy operation.
	Indy(context.Context, *IndyRequest) (*IndyResponse, error)
}

func RegisterConvcServer(conn dcerpc.Conn, o ConvcServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewConvcServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ConvcSyntaxV1_0))...)
}

func NewConvcServerHandle(o ConvcServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ConvcServerHandle(ctx, o, opNum, r)
	}
}

func ConvcServerHandle(ctx context.Context, o ConvcServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // convc_indy
		op := &xxx_IndyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IndyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Indy(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented convc
type UnimplementedConvcServer struct {
}

func (UnimplementedConvcServer) Indy(context.Context, *IndyRequest) (*IndyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ConvcServer = (*UnimplementedConvcServer)(nil)
