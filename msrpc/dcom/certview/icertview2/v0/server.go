package icertview2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	icertview "github.com/oiweiwei/go-msrpc/msrpc/dcom/certview/icertview/v0"
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
	_ = icertview.GoPackage
)

// ICertView2 server interface.
type CertView2Server interface {

	// ICertView base class.
	icertview.CertViewServer

	// SetTable operation.
	SetTable(context.Context, *SetTableRequest) (*SetTableResponse, error)
}

func RegisterCertView2Server(conn dcerpc.Conn, o CertView2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewCertView2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(CertView2SyntaxV0_0))...)
}

func NewCertView2ServerHandle(o CertView2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return CertView2ServerHandle(ctx, o, opNum, r)
	}
}

func CertView2ServerHandle(ctx context.Context, o CertView2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 15 {
		// ICertView base method.
		return icertview.CertViewServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 15: // SetTable
		op := &xxx_SetTableOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetTableRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetTable(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented ICertView2
type UnimplementedCertView2Server struct {
	icertview.UnimplementedCertViewServer
}

func (UnimplementedCertView2Server) SetTable(context.Context, *SetTableRequest) (*SetTableResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ CertView2Server = (*UnimplementedCertView2Server)(nil)
