package iserverhealthreport2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iserverhealthreport "github.com/oiweiwei/go-msrpc/msrpc/dcom/dfsrh/iserverhealthreport/v0"
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
	_ = iserverhealthreport.GoPackage
)

// IServerHealthReport2 server interface.
type ServerHealthReport2Server interface {

	// IServerHealthReport base class.
	iserverhealthreport.ServerHealthReportServer

	// GetReport2 operation.
	GetReport2(context.Context, *GetReport2Request) (*GetReport2Response, error)

	// GetCompressedReport2 operation.
	GetCompressedReport2(context.Context, *GetCompressedReport2Request) (*GetCompressedReport2Response, error)
}

func RegisterServerHealthReport2Server(conn dcerpc.Conn, o ServerHealthReport2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewServerHealthReport2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ServerHealthReport2SyntaxV0_0))...)
}

func NewServerHealthReport2ServerHandle(o ServerHealthReport2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ServerHealthReport2ServerHandle(ctx, o, opNum, r)
	}
}

func ServerHealthReport2ServerHandle(ctx context.Context, o ServerHealthReport2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 9 {
		// IServerHealthReport base method.
		return iserverhealthreport.ServerHealthReportServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 9: // GetReport2
		op := &xxx_GetReport2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetReport2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetReport2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // GetCompressedReport2
		op := &xxx_GetCompressedReport2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCompressedReport2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCompressedReport2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IServerHealthReport2
type UnimplementedServerHealthReport2Server struct {
	iserverhealthreport.UnimplementedServerHealthReportServer
}

func (UnimplementedServerHealthReport2Server) GetReport2(context.Context, *GetReport2Request) (*GetReport2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServerHealthReport2Server) GetCompressedReport2(context.Context, *GetCompressedReport2Request) (*GetCompressedReport2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ServerHealthReport2Server = (*UnimplementedServerHealthReport2Server)(nil)
