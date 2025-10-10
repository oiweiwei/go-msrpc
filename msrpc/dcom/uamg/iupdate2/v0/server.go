package iupdate2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iupdate "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdate/v0"
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
	_ = iupdate.GoPackage
)

// IUpdate2 server interface.
type Update2Server interface {

	// IUpdate base class.
	iupdate.UpdateServer

	GetRebootRequired(context.Context, *GetRebootRequiredRequest) (*GetRebootRequiredResponse, error)

	GetIsPresent(context.Context, *GetIsPresentRequest) (*GetIsPresentResponse, error)

	GetCveIDs(context.Context, *GetCveIDsRequest) (*GetCveIDsResponse, error)

	// Opnum56NotUsedOnWire operation.
	// Opnum56NotUsedOnWire
}

func RegisterUpdate2Server(conn dcerpc.Conn, o Update2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewUpdate2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Update2SyntaxV0_0))...)
}

func NewUpdate2ServerHandle(o Update2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Update2ServerHandle(ctx, o, opNum, r)
	}
}

func Update2ServerHandle(ctx context.Context, o Update2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 52 {
		// IUpdate base method.
		return iupdate.UpdateServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 52: // RebootRequired
		op := &xxx_GetRebootRequiredOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRebootRequiredRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRebootRequired(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 53: // IsPresent
		op := &xxx_GetIsPresentOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsPresentRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsPresent(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 54: // CveIDs
		op := &xxx_GetCveIDsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCveIDsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCveIDs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 55: // Opnum56NotUsedOnWire
		// Opnum56NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}

// Unimplemented IUpdate2
type UnimplementedUpdate2Server struct {
	iupdate.UnimplementedUpdateServer
}

func (UnimplementedUpdate2Server) GetRebootRequired(context.Context, *GetRebootRequiredRequest) (*GetRebootRequiredResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdate2Server) GetIsPresent(context.Context, *GetIsPresentRequest) (*GetIsPresentResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdate2Server) GetCveIDs(context.Context, *GetCveIDsRequest) (*GetCveIDsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ Update2Server = (*UnimplementedUpdate2Server)(nil)
