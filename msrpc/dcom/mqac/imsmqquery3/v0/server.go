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
type Query3Server interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The LookupQueue_v2 method is received by the server in an RPC_REQUEST packet. In
	// response, the server MUST return a pointer to an IMSMQQueueInfos4 interface pointer.
	// The returned interface allows the client to enumerate over a collection of Queues.
	// The queue query represents search criteria based on any combination of Queue.Identifier,
	// Queue.Type, Queue.Label, Queue.CreateTime, or Queue.ModifyTime  properties.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	LookupQueueV2(context.Context, *LookupQueueV2Request) (*LookupQueueV2Response, error)

	// Properties operation.
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)

	// The LookupQueue method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return a pointer to an IMSMQQueueInfos4 interface pointer. The returned
	// interface allows the client to enumerate over a collection of Queues. The queue query
	// represents search criteria based on any combination of Queue.Identifier, Queue.Type,
	// Queue.Label, Queue.CreateTime, Queue.ModifyTime, or Queue.MulticastAddress properties.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	LookupQueue(context.Context, *LookupQueueRequest) (*LookupQueueResponse, error)
}

func RegisterQuery3Server(conn dcerpc.Conn, o Query3Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewQuery3ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Query3SyntaxV0_0))...)
}

func NewQuery3ServerHandle(o Query3Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Query3ServerHandle(ctx, o, opNum, r)
	}
}

func Query3ServerHandle(ctx context.Context, o Query3Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
type UnimplementedQuery3Server struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedQuery3Server) LookupQueueV2(context.Context, *LookupQueueV2Request) (*LookupQueueV2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuery3Server) GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuery3Server) LookupQueue(context.Context, *LookupQueueRequest) (*LookupQueueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ Query3Server = (*UnimplementedQuery3Server)(nil)
