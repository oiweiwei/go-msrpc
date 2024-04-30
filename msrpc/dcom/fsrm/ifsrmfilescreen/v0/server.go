package ifsrmfilescreen

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	ifsrmfilescreenbase "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmfilescreenbase/v0"
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
	_ = ifsrmfilescreenbase.GoPackage
)

// IFsrmFileScreen server interface.
type FileScreenServer interface {

	// IFsrmFileScreenBase base class.
	ifsrmfilescreenbase.FileScreenBaseServer

	// Path operation.
	GetPath(context.Context, *GetPathRequest) (*GetPathResponse, error)

	// SourceTemplateName operation.
	GetSourceTemplateName(context.Context, *GetSourceTemplateNameRequest) (*GetSourceTemplateNameResponse, error)

	// MatchesSourceTemplate operation.
	GetMatchesSourceTemplate(context.Context, *GetMatchesSourceTemplateRequest) (*GetMatchesSourceTemplateResponse, error)

	// UserSid operation.
	GetUserSID(context.Context, *GetUserSIDRequest) (*GetUserSIDResponse, error)

	// UserAccount operation.
	GetUserAccount(context.Context, *GetUserAccountRequest) (*GetUserAccountResponse, error)

	// ApplyTemplate operation.
	ApplyTemplate(context.Context, *ApplyTemplateRequest) (*ApplyTemplateResponse, error)
}

func RegisterFileScreenServer(conn dcerpc.Conn, o FileScreenServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewFileScreenServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(FileScreenSyntaxV0_0))...)
}

func NewFileScreenServerHandle(o FileScreenServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return FileScreenServerHandle(ctx, o, opNum, r)
	}
}

func FileScreenServerHandle(ctx context.Context, o FileScreenServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 18 {
		// IFsrmFileScreenBase base method.
		return ifsrmfilescreenbase.FileScreenBaseServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 18: // Path
		in := &GetPathRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPath(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // SourceTemplateName
		in := &GetSourceTemplateNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSourceTemplateName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // MatchesSourceTemplate
		in := &GetMatchesSourceTemplateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMatchesSourceTemplate(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // UserSid
		in := &GetUserSIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetUserSID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 22: // UserAccount
		in := &GetUserAccountRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetUserAccount(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // ApplyTemplate
		in := &ApplyTemplateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ApplyTemplate(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
