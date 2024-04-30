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
		in := &GetNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // Name
		in := &SetNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // CopyTemplate
		in := &CopyTemplateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CopyTemplate(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // CommitAndUpdateDerived
		in := &CommitAndUpdateDerivedRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CommitAndUpdateDerived(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
