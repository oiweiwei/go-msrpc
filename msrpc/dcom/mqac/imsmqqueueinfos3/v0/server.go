package imsmqqueueinfos3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
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
	_ = idispatch.GoPackage
)

// IMSMQQueueInfos3 server interface.
type ImsmqQueueInfos3Server interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// Reset operation.
	Reset(context.Context, *ResetRequest) (*ResetResponse, error)

	// Next operation.
	Next(context.Context, *NextRequest) (*NextResponse, error)

	// Properties operation.
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)
}

func RegisterImsmqQueueInfos3Server(conn dcerpc.Conn, o ImsmqQueueInfos3Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewImsmqQueueInfos3ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ImsmqQueueInfos3SyntaxV0_0))...)
}

func NewImsmqQueueInfos3ServerHandle(o ImsmqQueueInfos3Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ImsmqQueueInfos3ServerHandle(ctx, o, opNum, r)
	}
}

func ImsmqQueueInfos3ServerHandle(ctx context.Context, o ImsmqQueueInfos3Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Reset
		op := &xxx_ResetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ResetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Reset(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // Next
		op := &xxx_NextOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &NextRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Next(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // Properties
		op := &xxx_GetPropertiesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPropertiesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetProperties(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMSMQQueueInfos3
type UnimplementedImsmqQueueInfos3Server struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedImsmqQueueInfos3Server) Reset(context.Context, *ResetRequest) (*ResetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfos3Server) Next(context.Context, *NextRequest) (*NextResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfos3Server) GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ImsmqQueueInfos3Server = (*UnimplementedImsmqQueueInfos3Server)(nil)
