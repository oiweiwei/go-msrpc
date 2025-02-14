package ifsrmquotatemplate

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	ifsrmquotabase "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmquotabase/v0"
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
	_ = ifsrmquotabase.GoPackage
)

// IFsrmQuotaTemplate server interface.
type QuotaTemplateServer interface {

	// IFsrmQuotaBase base class.
	ifsrmquotabase.QuotaBaseServer

	// Name operation.
	GetName(context.Context, *GetNameRequest) (*GetNameResponse, error)

	// Name operation.
	SetName(context.Context, *SetNameRequest) (*SetNameResponse, error)

	// CopyTemplate operation.
	CopyTemplate(context.Context, *CopyTemplateRequest) (*CopyTemplateResponse, error)

	// CommitAndUpdateDerived operation.
	CommitAndUpdateDerived(context.Context, *CommitAndUpdateDerivedRequest) (*CommitAndUpdateDerivedResponse, error)
}

func RegisterQuotaTemplateServer(conn dcerpc.Conn, o QuotaTemplateServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewQuotaTemplateServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(QuotaTemplateSyntaxV0_0))...)
}

func NewQuotaTemplateServerHandle(o QuotaTemplateServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return QuotaTemplateServerHandle(ctx, o, opNum, r)
	}
}

func QuotaTemplateServerHandle(ctx context.Context, o QuotaTemplateServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 22 {
		// IFsrmQuotaBase base method.
		return ifsrmquotabase.QuotaBaseServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 22: // Name
		op := &xxx_GetNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // Name
		op := &xxx_SetNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // CopyTemplate
		op := &xxx_CopyTemplateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CopyTemplateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CopyTemplate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // CommitAndUpdateDerived
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

// Unimplemented IFsrmQuotaTemplate
type UnimplementedQuotaTemplateServer struct {
	ifsrmquotabase.UnimplementedQuotaBaseServer
}

func (UnimplementedQuotaTemplateServer) GetName(context.Context, *GetNameRequest) (*GetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaTemplateServer) SetName(context.Context, *SetNameRequest) (*SetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaTemplateServer) CopyTemplate(context.Context, *CopyTemplateRequest) (*CopyTemplateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaTemplateServer) CommitAndUpdateDerived(context.Context, *CommitAndUpdateDerivedRequest) (*CommitAndUpdateDerivedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ QuotaTemplateServer = (*UnimplementedQuotaTemplateServer)(nil)
