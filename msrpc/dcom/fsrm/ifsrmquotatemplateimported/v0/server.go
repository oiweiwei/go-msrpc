package ifsrmquotatemplateimported

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	ifsrmquotatemplate "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmquotatemplate/v0"
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
	_ = ifsrmquotatemplate.GoPackage
)

// IFsrmQuotaTemplateImported server interface.
type QuotaTemplateImportedServer interface {

	// IFsrmQuotaTemplate base class.
	ifsrmquotatemplate.QuotaTemplateServer

	// OverwriteOnCommit operation.
	GetOverwriteOnCommit(context.Context, *GetOverwriteOnCommitRequest) (*GetOverwriteOnCommitResponse, error)

	// OverwriteOnCommit operation.
	SetOverwriteOnCommit(context.Context, *SetOverwriteOnCommitRequest) (*SetOverwriteOnCommitResponse, error)
}

func RegisterQuotaTemplateImportedServer(conn dcerpc.Conn, o QuotaTemplateImportedServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewQuotaTemplateImportedServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(QuotaTemplateImportedSyntaxV0_0))...)
}

func NewQuotaTemplateImportedServerHandle(o QuotaTemplateImportedServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return QuotaTemplateImportedServerHandle(ctx, o, opNum, r)
	}
}

func QuotaTemplateImportedServerHandle(ctx context.Context, o QuotaTemplateImportedServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 26 {
		// IFsrmQuotaTemplate base method.
		return ifsrmquotatemplate.QuotaTemplateServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 26: // OverwriteOnCommit
		op := &xxx_GetOverwriteOnCommitOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetOverwriteOnCommitRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetOverwriteOnCommit(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // OverwriteOnCommit
		op := &xxx_SetOverwriteOnCommitOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetOverwriteOnCommitRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetOverwriteOnCommit(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmQuotaTemplateImported
type UnimplementedQuotaTemplateImportedServer struct {
	ifsrmquotatemplate.UnimplementedQuotaTemplateServer
}

func (UnimplementedQuotaTemplateImportedServer) GetOverwriteOnCommit(context.Context, *GetOverwriteOnCommitRequest) (*GetOverwriteOnCommitResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaTemplateImportedServer) SetOverwriteOnCommit(context.Context, *SetOverwriteOnCommitRequest) (*SetOverwriteOnCommitResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ QuotaTemplateImportedServer = (*UnimplementedQuotaTemplateImportedServer)(nil)
