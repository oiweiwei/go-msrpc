package imsmqquery2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = idispatch.GoPackage
)

// IMSMQQuery2 server interface.
type Query2Server interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The LookupQueue method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return a pointer to an IMSMQQueueInfos4 interface pointer. The returned
	// interface allows the client to enumerate over a collection of Queues. The queue query
	// represents search criteria based on any combination of Queue.Identifier, Queue.Type,
	// Queue.Label, Queue.CreateTime, Queue.ModifyTime, or Queue.MulticastAddress properties.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	LookupQueue(context.Context, *LookupQueueRequest) (*LookupQueueResponse, error)

	// Properties operation.
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)
}

func RegisterQuery2Server(conn dcerpc.Conn, o Query2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewQuery2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Query2SyntaxV0_0))...)
}

func NewQuery2ServerHandle(o Query2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Query2ServerHandle(ctx, o, opNum, r)
	}
}

func Query2ServerHandle(ctx context.Context, o Query2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // LookupQueue
		op := &xxx_LookupQueueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &LookupQueueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.LookupQueue(ctx, req)
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
	}
	return nil, nil
}

// Unimplemented IMSMQQuery2
type UnimplementedQuery2Server struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedQuery2Server) LookupQueue(context.Context, *LookupQueueRequest) (*LookupQueueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuery2Server) GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ Query2Server = (*UnimplementedQuery2Server)(nil)
