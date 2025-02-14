package iapphostwritableadminmanager

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iapphostadminmanager "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostadminmanager/v0"
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
	_ = iapphostadminmanager.GoPackage
)

// IAppHostWritableAdminManager server interface.
type AppHostWritableAdminManagerServer interface {

	// IAppHostAdminManager base class.
	iapphostadminmanager.AppHostAdminManagerServer

	// CommitChanges operation.
	CommitChanges(context.Context, *CommitChangesRequest) (*CommitChangesResponse, error)

	// CommitPath operation.
	GetCommitPath(context.Context, *GetCommitPathRequest) (*GetCommitPathResponse, error)

	// CommitPath operation.
	SetCommitPath(context.Context, *SetCommitPathRequest) (*SetCommitPathResponse, error)
}

func RegisterAppHostWritableAdminManagerServer(conn dcerpc.Conn, o AppHostWritableAdminManagerServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostWritableAdminManagerServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostWritableAdminManagerSyntaxV0_0))...)
}

func NewAppHostWritableAdminManagerServerHandle(o AppHostWritableAdminManagerServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostWritableAdminManagerServerHandle(ctx, o, opNum, r)
	}
}

func AppHostWritableAdminManagerServerHandle(ctx context.Context, o AppHostWritableAdminManagerServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IAppHostAdminManager base method.
		return iapphostadminmanager.AppHostAdminManagerServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // CommitChanges
		op := &xxx_CommitChangesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CommitChangesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CommitChanges(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // CommitPath
		op := &xxx_GetCommitPathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCommitPathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCommitPath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // CommitPath
		op := &xxx_SetCommitPathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetCommitPathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetCommitPath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IAppHostWritableAdminManager
type UnimplementedAppHostWritableAdminManagerServer struct {
	iapphostadminmanager.UnimplementedAppHostAdminManagerServer
}

func (UnimplementedAppHostWritableAdminManagerServer) CommitChanges(context.Context, *CommitChangesRequest) (*CommitChangesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostWritableAdminManagerServer) GetCommitPath(context.Context, *GetCommitPathRequest) (*GetCommitPathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostWritableAdminManagerServer) SetCommitPath(context.Context, *SetCommitPathRequest) (*SetCommitPathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AppHostWritableAdminManagerServer = (*UnimplementedAppHostWritableAdminManagerServer)(nil)
