package imsmqquery4

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

// IMSMQQuery4 server interface.
type Query4Server interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// LookupQueue_v2 operation.
	LookupQueueV2(context.Context, *LookupQueueV2Request) (*LookupQueueV2Response, error)

	// Properties operation.
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)

	// LookupQueue operation.
	LookupQueue(context.Context, *LookupQueueRequest) (*LookupQueueResponse, error)
}

func RegisterQuery4Server(conn dcerpc.Conn, o Query4Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewQuery4ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Query4SyntaxV0_0))...)
}

func NewQuery4ServerHandle(o Query4Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Query4ServerHandle(ctx, o, opNum, r)
	}
}

func Query4ServerHandle(ctx context.Context, o Query4Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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

// Unimplemented IMSMQQuery4
type UnimplementedQuery4Server struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedQuery4Server) LookupQueueV2(context.Context, *LookupQueueV2Request) (*LookupQueueV2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuery4Server) GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuery4Server) LookupQueue(context.Context, *LookupQueueRequest) (*LookupQueueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ Query4Server = (*UnimplementedQuery4Server)(nil)
