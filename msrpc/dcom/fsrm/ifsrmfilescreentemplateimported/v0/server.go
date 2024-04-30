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
		in := &GetOverwriteOnCommitRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetOverwriteOnCommit(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // OverwriteOnCommit
		in := &SetOverwriteOnCommitRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetOverwriteOnCommit(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
