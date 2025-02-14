package ifsrmquotaobject

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

// IFsrmQuotaObject server interface.
type QuotaObjectServer interface {

	// IFsrmQuotaBase base class.
	ifsrmquotabase.QuotaBaseServer

	// Path operation.
	GetPath(context.Context, *GetPathRequest) (*GetPathResponse, error)

	// UserSid operation.
	GetUserSID(context.Context, *GetUserSIDRequest) (*GetUserSIDResponse, error)

	// UserAccount operation.
	GetUserAccount(context.Context, *GetUserAccountRequest) (*GetUserAccountResponse, error)

	// SourceTemplateName operation.
	GetSourceTemplateName(context.Context, *GetSourceTemplateNameRequest) (*GetSourceTemplateNameResponse, error)

	// MatchesSourceTemplate operation.
	GetMatchesSourceTemplate(context.Context, *GetMatchesSourceTemplateRequest) (*GetMatchesSourceTemplateResponse, error)

	// ApplyTemplate operation.
	ApplyTemplate(context.Context, *ApplyTemplateRequest) (*ApplyTemplateResponse, error)
}

func RegisterQuotaObjectServer(conn dcerpc.Conn, o QuotaObjectServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewQuotaObjectServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(QuotaObjectSyntaxV0_0))...)
}

func NewQuotaObjectServerHandle(o QuotaObjectServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return QuotaObjectServerHandle(ctx, o, opNum, r)
	}
}

func QuotaObjectServerHandle(ctx context.Context, o QuotaObjectServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 22 {
		// IFsrmQuotaBase base method.
		return ifsrmquotabase.QuotaBaseServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 22: // Path
		op := &xxx_GetPathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // UserSid
		op := &xxx_GetUserSIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetUserSIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetUserSID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // UserAccount
		op := &xxx_GetUserAccountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetUserAccountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetUserAccount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // SourceTemplateName
		op := &xxx_GetSourceTemplateNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSourceTemplateNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSourceTemplateName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 26: // MatchesSourceTemplate
		op := &xxx_GetMatchesSourceTemplateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMatchesSourceTemplateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMatchesSourceTemplate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // ApplyTemplate
		op := &xxx_ApplyTemplateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ApplyTemplateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ApplyTemplate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmQuotaObject
type UnimplementedQuotaObjectServer struct {
	ifsrmquotabase.UnimplementedQuotaBaseServer
}

func (UnimplementedQuotaObjectServer) GetPath(context.Context, *GetPathRequest) (*GetPathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaObjectServer) GetUserSID(context.Context, *GetUserSIDRequest) (*GetUserSIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaObjectServer) GetUserAccount(context.Context, *GetUserAccountRequest) (*GetUserAccountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaObjectServer) GetSourceTemplateName(context.Context, *GetSourceTemplateNameRequest) (*GetSourceTemplateNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaObjectServer) GetMatchesSourceTemplate(context.Context, *GetMatchesSourceTemplateRequest) (*GetMatchesSourceTemplateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaObjectServer) ApplyTemplate(context.Context, *ApplyTemplateRequest) (*ApplyTemplateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ QuotaObjectServer = (*UnimplementedQuotaObjectServer)(nil)
