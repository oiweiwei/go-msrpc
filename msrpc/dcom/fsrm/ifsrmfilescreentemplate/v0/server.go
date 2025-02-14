package ifsrmfilescreentemplate

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

// IFsrmFileScreenTemplate server interface.
type FileScreenTemplateServer interface {

	// IFsrmFileScreenBase base class.
	ifsrmfilescreenbase.FileScreenBaseServer

	// Name operation.
	GetName(context.Context, *GetNameRequest) (*GetNameResponse, error)

	// Name operation.
	SetName(context.Context, *SetNameRequest) (*SetNameResponse, error)

	// CopyTemplate operation.
	CopyTemplate(context.Context, *CopyTemplateRequest) (*CopyTemplateResponse, error)

	// CommitAndUpdateDerived operation.
	CommitAndUpdateDerived(context.Context, *CommitAndUpdateDerivedRequest) (*CommitAndUpdateDerivedResponse, error)
}

func RegisterFileScreenTemplateServer(conn dcerpc.Conn, o FileScreenTemplateServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewFileScreenTemplateServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(FileScreenTemplateSyntaxV0_0))...)
}

func NewFileScreenTemplateServerHandle(o FileScreenTemplateServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return FileScreenTemplateServerHandle(ctx, o, opNum, r)
	}
}

func FileScreenTemplateServerHandle(ctx context.Context, o FileScreenTemplateServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 18 {
		// IFsrmFileScreenBase base method.
		return ifsrmfilescreenbase.FileScreenBaseServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 18: // Name
		op := &xxx_GetNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // Name
		op := &xxx_SetNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // CopyTemplate
		op := &xxx_CopyTemplateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CopyTemplateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CopyTemplate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // CommitAndUpdateDerived
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

// Unimplemented IFsrmFileScreenTemplate
type UnimplementedFileScreenTemplateServer struct {
	ifsrmfilescreenbase.UnimplementedFileScreenBaseServer
}

func (UnimplementedFileScreenTemplateServer) GetName(context.Context, *GetNameRequest) (*GetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileScreenTemplateServer) SetName(context.Context, *SetNameRequest) (*SetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileScreenTemplateServer) CopyTemplate(context.Context, *CopyTemplateRequest) (*CopyTemplateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileScreenTemplateServer) CommitAndUpdateDerived(context.Context, *CommitAndUpdateDerivedRequest) (*CommitAndUpdateDerivedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ FileScreenTemplateServer = (*UnimplementedFileScreenTemplateServer)(nil)
