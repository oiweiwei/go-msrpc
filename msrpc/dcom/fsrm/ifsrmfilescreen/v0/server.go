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
		op := &xxx_GetPathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // SourceTemplateName
		op := &xxx_GetSourceTemplateNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSourceTemplateNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSourceTemplateName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // MatchesSourceTemplate
		op := &xxx_GetMatchesSourceTemplateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMatchesSourceTemplateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMatchesSourceTemplate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // UserSid
		op := &xxx_GetUserSIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetUserSIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetUserSID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // UserAccount
		op := &xxx_GetUserAccountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetUserAccountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetUserAccount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // ApplyTemplate
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

// Unimplemented IFsrmFileScreen
type UnimplementedFileScreenServer struct {
	ifsrmfilescreenbase.UnimplementedFileScreenBaseServer
}

func (UnimplementedFileScreenServer) GetPath(context.Context, *GetPathRequest) (*GetPathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileScreenServer) GetSourceTemplateName(context.Context, *GetSourceTemplateNameRequest) (*GetSourceTemplateNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileScreenServer) GetMatchesSourceTemplate(context.Context, *GetMatchesSourceTemplateRequest) (*GetMatchesSourceTemplateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileScreenServer) GetUserSID(context.Context, *GetUserSIDRequest) (*GetUserSIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileScreenServer) GetUserAccount(context.Context, *GetUserAccountRequest) (*GetUserAccountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileScreenServer) ApplyTemplate(context.Context, *ApplyTemplateRequest) (*ApplyTemplateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ FileScreenServer = (*UnimplementedFileScreenServer)(nil)
