package ifsrmquotatemplatemanager

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
		op := &xxx_CreateTemplateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateTemplateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateTemplate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // GetTemplate
		op := &xxx_GetTemplateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTemplateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTemplate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // EnumTemplates
		op := &xxx_EnumTemplatesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumTemplatesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumTemplates(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // ExportTemplates
		op := &xxx_ExportTemplatesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ExportTemplatesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ExportTemplates(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // ImportTemplates
		op := &xxx_ImportTemplatesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ImportTemplatesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ImportTemplates(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmQuotaTemplateManager
type UnimplementedQuotaTemplateManagerServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedQuotaTemplateManagerServer) CreateTemplate(context.Context, *CreateTemplateRequest) (*CreateTemplateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaTemplateManagerServer) GetTemplate(context.Context, *GetTemplateRequest) (*GetTemplateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaTemplateManagerServer) EnumTemplates(context.Context, *EnumTemplatesRequest) (*EnumTemplatesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaTemplateManagerServer) ExportTemplates(context.Context, *ExportTemplatesRequest) (*ExportTemplatesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaTemplateManagerServer) ImportTemplates(context.Context, *ImportTemplatesRequest) (*ImportTemplatesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ QuotaTemplateManagerServer = (*UnimplementedQuotaTemplateManagerServer)(nil)
