package ifsrmquotatemplatemanager

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

// IFsrmQuotaTemplateManager server interface.
type QuotaTemplateManagerServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// CreateTemplate operation.
	CreateTemplate(context.Context, *CreateTemplateRequest) (*CreateTemplateResponse, error)

	// GetTemplate operation.
	GetTemplate(context.Context, *GetTemplateRequest) (*GetTemplateResponse, error)

	// EnumTemplates operation.
	EnumTemplates(context.Context, *EnumTemplatesRequest) (*EnumTemplatesResponse, error)

	// ExportTemplates operation.
	ExportTemplates(context.Context, *ExportTemplatesRequest) (*ExportTemplatesResponse, error)

	// ImportTemplates operation.
	ImportTemplates(context.Context, *ImportTemplatesRequest) (*ImportTemplatesResponse, error)
}

func RegisterQuotaTemplateManagerServer(conn dcerpc.Conn, o QuotaTemplateManagerServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewQuotaTemplateManagerServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(QuotaTemplateManagerSyntaxV0_0))...)
}

func NewQuotaTemplateManagerServerHandle(o QuotaTemplateManagerServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return QuotaTemplateManagerServerHandle(ctx, o, opNum, r)
	}
}

func QuotaTemplateManagerServerHandle(ctx context.Context, o QuotaTemplateManagerServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // CreateTemplate
		in := &CreateTemplateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateTemplate(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // GetTemplate
		in := &GetTemplateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTemplate(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // EnumTemplates
		in := &EnumTemplatesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumTemplates(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // ExportTemplates
		in := &ExportTemplatesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ExportTemplates(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // ImportTemplates
		in := &ImportTemplatesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ImportTemplates(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
