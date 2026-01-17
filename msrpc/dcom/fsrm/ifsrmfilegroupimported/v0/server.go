package ifsrmfilegroupimported

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	ifsrmfilegroup "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmfilegroup/v0"
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
	_ = ifsrmfilegroup.GoPackage
)

// IFsrmFileGroupImported server interface.
type FileGroupImportedServer interface {

	// IFsrmFileGroup base class.
	ifsrmfilegroup.FileGroupServer

	// OverwriteOnCommit operation.
	GetOverwriteOnCommit(context.Context, *GetOverwriteOnCommitRequest) (*GetOverwriteOnCommitResponse, error)

	// OverwriteOnCommit operation.
	SetOverwriteOnCommit(context.Context, *SetOverwriteOnCommitRequest) (*SetOverwriteOnCommitResponse, error)
}

func RegisterFileGroupImportedServer(conn dcerpc.Conn, o FileGroupImportedServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewFileGroupImportedServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(FileGroupImportedSyntaxV0_0))...)
}

func NewFileGroupImportedServerHandle(o FileGroupImportedServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return FileGroupImportedServerHandle(ctx, o, opNum, r)
	}
}

func FileGroupImportedServerHandle(ctx context.Context, o FileGroupImportedServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 18 {
		// IFsrmFileGroup base method.
		return ifsrmfilegroup.FileGroupServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 18: // OverwriteOnCommit
		op := &xxx_GetOverwriteOnCommitOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetOverwriteOnCommitRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetOverwriteOnCommit(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // OverwriteOnCommit
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

// Unimplemented IFsrmFileGroupImported
type UnimplementedFileGroupImportedServer struct {
	ifsrmfilegroup.UnimplementedFileGroupServer
}

func (UnimplementedFileGroupImportedServer) GetOverwriteOnCommit(context.Context, *GetOverwriteOnCommitRequest) (*GetOverwriteOnCommitResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileGroupImportedServer) SetOverwriteOnCommit(context.Context, *SetOverwriteOnCommitRequest) (*SetOverwriteOnCommitResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ FileGroupImportedServer = (*UnimplementedFileGroupImportedServer)(nil)
