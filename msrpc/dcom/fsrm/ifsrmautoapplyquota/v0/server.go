package ifsrmautoapplyquota

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	ifsrmquotaobject "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmquotaobject/v0"
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
	_ = ifsrmquotaobject.GoPackage
)

// IFsrmAutoApplyQuota server interface.
type AutoApplyQuotaServer interface {

	// IFsrmQuotaObject base class.
	ifsrmquotaobject.QuotaObjectServer

	// ExcludeFolders operation.
	GetExcludeFolders(context.Context, *GetExcludeFoldersRequest) (*GetExcludeFoldersResponse, error)

	// ExcludeFolders operation.
	SetExcludeFolders(context.Context, *SetExcludeFoldersRequest) (*SetExcludeFoldersResponse, error)

	// CommitAndUpdateDerived operation.
	CommitAndUpdateDerived(context.Context, *CommitAndUpdateDerivedRequest) (*CommitAndUpdateDerivedResponse, error)
}

func RegisterAutoApplyQuotaServer(conn dcerpc.Conn, o AutoApplyQuotaServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAutoApplyQuotaServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AutoApplyQuotaSyntaxV0_0))...)
}

func NewAutoApplyQuotaServerHandle(o AutoApplyQuotaServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AutoApplyQuotaServerHandle(ctx, o, opNum, r)
	}
}

func AutoApplyQuotaServerHandle(ctx context.Context, o AutoApplyQuotaServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 28 {
		// IFsrmQuotaObject base method.
		return ifsrmquotaobject.QuotaObjectServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 28: // ExcludeFolders
		op := &xxx_GetExcludeFoldersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetExcludeFoldersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetExcludeFolders(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 29: // ExcludeFolders
		op := &xxx_SetExcludeFoldersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetExcludeFoldersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetExcludeFolders(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // CommitAndUpdateDerived
		op := &xxx_CommitAndUpdateDerivedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CommitAndUpdateDerivedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CommitAndUpdateDerived(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmAutoApplyQuota
type UnimplementedAutoApplyQuotaServer struct {
	ifsrmquotaobject.UnimplementedQuotaObjectServer
}

func (UnimplementedAutoApplyQuotaServer) GetExcludeFolders(context.Context, *GetExcludeFoldersRequest) (*GetExcludeFoldersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAutoApplyQuotaServer) SetExcludeFolders(context.Context, *SetExcludeFoldersRequest) (*SetExcludeFoldersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAutoApplyQuotaServer) CommitAndUpdateDerived(context.Context, *CommitAndUpdateDerivedRequest) (*CommitAndUpdateDerivedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AutoApplyQuotaServer = (*UnimplementedAutoApplyQuotaServer)(nil)
