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
		in := &GetPathRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPath(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // UserSid
		in := &GetUserSIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetUserSID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // UserAccount
		in := &GetUserAccountRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetUserAccount(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // SourceTemplateName
		in := &GetSourceTemplateNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSourceTemplateName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 26: // MatchesSourceTemplate
		in := &GetMatchesSourceTemplateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMatchesSourceTemplate(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 27: // ApplyTemplate
		in := &ApplyTemplateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ApplyTemplate(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
