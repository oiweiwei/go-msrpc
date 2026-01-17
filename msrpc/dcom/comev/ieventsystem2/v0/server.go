package ieventsystem2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	ieventsystem "github.com/oiweiwei/go-msrpc/msrpc/dcom/comev/ieventsystem/v0"
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
	_ = ieventsystem.GoPackage
)

// IEventSystem2 server interface.
type EventSystem2Server interface {

	// IEventSystem base class.
	ieventsystem.EventSystemServer

	// The GetVersion method retrieves the version of the server implementation of the protocol.
	//
	// Return Values: An HRESULT specifying success or failure. All success codes MUST be
	// treated the same, and all failure codes MUST be treated the same.
	GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error)

	// The VerifyTransientSubscribers method verifies the state of the transient subscribers.
	//
	// This method has no parameters.
	//
	// Return Values: This function MUST return S_OK.
	VerifyTransientSubscribers(context.Context, *VerifyTransientSubscribersRequest) (*VerifyTransientSubscribersResponse, error)
}

func RegisterEventSystem2Server(conn dcerpc.Conn, o EventSystem2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewEventSystem2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(EventSystem2SyntaxV0_0))...)
}

func NewEventSystem2ServerHandle(o EventSystem2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return EventSystem2ServerHandle(ctx, o, opNum, r)
	}
}

func EventSystem2ServerHandle(ctx context.Context, o EventSystem2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 13 {
		// IEventSystem base method.
		return ieventsystem.EventSystemServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 13: // GetVersion
		op := &xxx_GetVersionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetVersionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetVersion(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // VerifyTransientSubscribers
		op := &xxx_VerifyTransientSubscribersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &VerifyTransientSubscribersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.VerifyTransientSubscribers(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IEventSystem2
type UnimplementedEventSystem2Server struct {
	ieventsystem.UnimplementedEventSystemServer
}

func (UnimplementedEventSystem2Server) GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSystem2Server) VerifyTransientSubscribers(context.Context, *VerifyTransientSubscribersRequest) (*VerifyTransientSubscribersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ EventSystem2Server = (*UnimplementedEventSystem2Server)(nil)
