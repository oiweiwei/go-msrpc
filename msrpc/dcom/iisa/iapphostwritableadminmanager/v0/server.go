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
		in := &CommitChangesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CommitChanges(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // CommitPath
		in := &GetCommitPathRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetCommitPath(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // CommitPath
		in := &SetCommitPathRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetCommitPath(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
