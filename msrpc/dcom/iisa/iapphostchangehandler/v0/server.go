package iapphostchangehandler

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

// IAppHostChangeHandler server interface.
type AppHostChangeHandlerServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// OnSectionChanges operation.
	OnSectionChanges(context.Context, *OnSectionChangesRequest) (*OnSectionChangesResponse, error)
}

func RegisterAppHostChangeHandlerServer(conn dcerpc.Conn, o AppHostChangeHandlerServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostChangeHandlerServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostChangeHandlerSyntaxV0_0))...)
}

func NewAppHostChangeHandlerServerHandle(o AppHostChangeHandlerServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostChangeHandlerServerHandle(ctx, o, opNum, r)
	}
}

func AppHostChangeHandlerServerHandle(ctx context.Context, o AppHostChangeHandlerServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // OnSectionChanges
		op := &xxx_OnSectionChangesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OnSectionChangesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OnSectionChanges(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IAppHostChangeHandler
type UnimplementedAppHostChangeHandlerServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedAppHostChangeHandlerServer) OnSectionChanges(context.Context, *OnSectionChangesRequest) (*OnSectionChangesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AppHostChangeHandlerServer = (*UnimplementedAppHostChangeHandlerServer)(nil)
