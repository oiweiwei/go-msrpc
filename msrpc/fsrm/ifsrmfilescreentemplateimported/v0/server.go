package ifsrmfilescreentemplateimported

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	ifsrmfilescreentemplate "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmfilescreentemplate/v0"
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
	_ = ifsrmfilescreentemplate.GoPackage
)

// IFsrmFileScreenTemplateImported server interface.
type FileScreenTemplateImportedServer interface {

	// IFsrmFileScreenTemplate base class.
	ifsrmfilescreentemplate.FileScreenTemplateServer

	// OverwriteOnCommit operation.
	GetOverwriteOnCommit(context.Context, *GetOverwriteOnCommitRequest) (*GetOverwriteOnCommitResponse, error)

	// OverwriteOnCommit operation.
	SetOverwriteOnCommit(context.Context, *SetOverwriteOnCommitRequest) (*SetOverwriteOnCommitResponse, error)
}

func RegisterFileScreenTemplateImportedServer(conn dcerpc.Conn, o FileScreenTemplateImportedServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewFileScreenTemplateImportedServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(FileScreenTemplateImportedSyntaxV0_0))...)
}

func NewFileScreenTemplateImportedServerHandle(o FileScreenTemplateImportedServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return FileScreenTemplateImportedServerHandle(ctx, o, opNum, r)
	}
}

func FileScreenTemplateImportedServerHandle(ctx context.Context, o FileScreenTemplateImportedServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 22 {
		// IFsrmFileScreenTemplate base method.
		return ifsrmfilescreentemplate.FileScreenTemplateServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 22: // OverwriteOnCommit
		op := &xxx_GetOverwriteOnCommitOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetOverwriteOnCommitRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetOverwriteOnCommit(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // OverwriteOnCommit
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

// Unimplemented IFsrmFileScreenTemplateImported
type UnimplementedFileScreenTemplateImportedServer struct {
	ifsrmfilescreentemplate.UnimplementedFileScreenTemplateServer
}

func (UnimplementedFileScreenTemplateImportedServer) GetOverwriteOnCommit(context.Context, *GetOverwriteOnCommitRequest) (*GetOverwriteOnCommitResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileScreenTemplateImportedServer) SetOverwriteOnCommit(context.Context, *SetOverwriteOnCommitRequest) (*SetOverwriteOnCommitResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ FileScreenTemplateImportedServer = (*UnimplementedFileScreenTemplateImportedServer)(nil)
