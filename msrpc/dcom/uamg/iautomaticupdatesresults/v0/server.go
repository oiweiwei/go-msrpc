package iautomaticupdatesresults

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

// IAutomaticUpdatesResults server interface.
type AutomaticUpdatesResultsServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	GetLastSearchSuccessDate(context.Context, *GetLastSearchSuccessDateRequest) (*GetLastSearchSuccessDateResponse, error)

	GetLastInstallationSuccessDate(context.Context, *GetLastInstallationSuccessDateRequest) (*GetLastInstallationSuccessDateResponse, error)
}

func RegisterAutomaticUpdatesResultsServer(conn dcerpc.Conn, o AutomaticUpdatesResultsServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAutomaticUpdatesResultsServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AutomaticUpdatesResultsSyntaxV0_0))...)
}

func NewAutomaticUpdatesResultsServerHandle(o AutomaticUpdatesResultsServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AutomaticUpdatesResultsServerHandle(ctx, o, opNum, r)
	}
}

func AutomaticUpdatesResultsServerHandle(ctx context.Context, o AutomaticUpdatesResultsServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // LastSearchSuccessDate
		op := &xxx_GetLastSearchSuccessDateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLastSearchSuccessDateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLastSearchSuccessDate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // LastInstallationSuccessDate
		op := &xxx_GetLastInstallationSuccessDateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLastInstallationSuccessDateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLastInstallationSuccessDate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IAutomaticUpdatesResults
type UnimplementedAutomaticUpdatesResultsServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedAutomaticUpdatesResultsServer) GetLastSearchSuccessDate(context.Context, *GetLastSearchSuccessDateRequest) (*GetLastSearchSuccessDateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAutomaticUpdatesResultsServer) GetLastInstallationSuccessDate(context.Context, *GetLastInstallationSuccessDateRequest) (*GetLastInstallationSuccessDateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AutomaticUpdatesResultsServer = (*UnimplementedAutomaticUpdatesResultsServer)(nil)
