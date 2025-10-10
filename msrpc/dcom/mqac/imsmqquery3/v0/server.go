package imsmqquery3

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

// IMSMQQuery3 server interface.
type ImsmqQuery3Server interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// LookupQueue_v2 operation.
	LookupQueueV2(context.Context, *LookupQueueV2Request) (*LookupQueueV2Response, error)

	// Properties operation.
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)

	// LookupQueue operation.
	LookupQueue(context.Context, *LookupQueueRequest) (*LookupQueueResponse, error)
}

func RegisterImsmqQuery3Server(conn dcerpc.Conn, o ImsmqQuery3Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewImsmqQuery3ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ImsmqQuery3SyntaxV0_0))...)
}

func NewImsmqQuery3ServerHandle(o ImsmqQuery3Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ImsmqQuery3ServerHandle(ctx, o, opNum, r)
	}
}

func ImsmqQuery3ServerHandle(ctx context.Context, o ImsmqQuery3Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // LookupQueue_v2
		op := &xxx_LookupQueueV2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &LookupQueueV2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.LookupQueueV2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // Properties
		op := &xxx_GetPropertiesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPropertiesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetProperties(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // LookupQueue
		op := &xxx_LookupQueueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &LookupQueueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.LookupQueue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMSMQQuery3
type UnimplementedImsmqQuery3Server struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedImsmqQuery3Server) LookupQueueV2(context.Context, *LookupQueueV2Request) (*LookupQueueV2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQuery3Server) GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQuery3Server) LookupQueue(context.Context, *LookupQueueRequest) (*LookupQueueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ImsmqQuery3Server = (*UnimplementedImsmqQuery3Server)(nil)
