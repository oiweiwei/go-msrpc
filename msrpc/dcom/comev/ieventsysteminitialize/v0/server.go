package ieventsysteminitialize

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
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
	_ = iunknown.GoPackage
)

// IEventSystemInitialize server interface.
type EventSystemInitializeServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The SetCOMCatalogBehaviour method controls the event system CatalogMode and RetainSubKeys
	// state variables.
	//
	// Return Values: The server MUST return S_OK.
	//
	// After this method is called, the event system CatalogMode state variable will be
	// true (server in catalog mode) and the RetainSubKeys variable will be set based on
	// the bRetainSubKeys parameter. If the client does not call this method, the server
	// will be in non-catalog mode. The Store, Remove, and RemoveS methods of IEventSystem
	// will have different behavior when the server is in catalog mode. See section 3.1.1.3
	// for more details.
	SetCOMCatalogBehaviour(context.Context, *SetCOMCatalogBehaviourRequest) (*SetCOMCatalogBehaviourResponse, error)
}

func RegisterEventSystemInitializeServer(conn dcerpc.Conn, o EventSystemInitializeServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewEventSystemInitializeServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(EventSystemInitializeSyntaxV0_0))...)
}

func NewEventSystemInitializeServerHandle(o EventSystemInitializeServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return EventSystemInitializeServerHandle(ctx, o, opNum, r)
	}
}

func EventSystemInitializeServerHandle(ctx context.Context, o EventSystemInitializeServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // SetCOMCatalogBehaviour
		op := &xxx_SetCOMCatalogBehaviourOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetCOMCatalogBehaviourRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetCOMCatalogBehaviour(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IEventSystemInitialize
type UnimplementedEventSystemInitializeServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedEventSystemInitializeServer) SetCOMCatalogBehaviour(context.Context, *SetCOMCatalogBehaviourRequest) (*SetCOMCatalogBehaviourResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ EventSystemInitializeServer = (*UnimplementedEventSystemInitializeServer)(nil)
