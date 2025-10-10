package iupdateservice2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iupdateservice "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdateservice/v0"
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
	_ = iupdateservice.GoPackage
)

// IUpdateService2 server interface.
type UpdateService2Server interface {

	// IUpdateService base class.
	iupdateservice.UpdateServiceServer

	GetIsDefaultAuService(context.Context, *GetIsDefaultAuServiceRequest) (*GetIsDefaultAuServiceResponse, error)
}

func RegisterUpdateService2Server(conn dcerpc.Conn, o UpdateService2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewUpdateService2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(UpdateService2SyntaxV0_0))...)
}

func NewUpdateService2ServerHandle(o UpdateService2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return UpdateService2ServerHandle(ctx, o, opNum, r)
	}
}

func UpdateService2ServerHandle(ctx context.Context, o UpdateService2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 20 {
		// IUpdateService base method.
		return iupdateservice.UpdateServiceServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 20: // IsDefaultAUService
		op := &xxx_GetIsDefaultAuServiceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsDefaultAuServiceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsDefaultAuService(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IUpdateService2
type UnimplementedUpdateService2Server struct {
	iupdateservice.UnimplementedUpdateServiceServer
}

func (UnimplementedUpdateService2Server) GetIsDefaultAuService(context.Context, *GetIsDefaultAuServiceRequest) (*GetIsDefaultAuServiceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ UpdateService2Server = (*UnimplementedUpdateService2Server)(nil)
