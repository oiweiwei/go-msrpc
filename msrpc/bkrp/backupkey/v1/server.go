package backupkey

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
)

// BackupKey server interface.
type BackupKeyServer interface {

	// This section specifies the BackuprKey method.
	//
	// Return Values: The server MUST return 0 if it successfully processes the message
	// received from the client, and a nonzero value otherwise.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	BackupKey(context.Context, *BackupKeyRequest) (*BackupKeyResponse, error)
}

func RegisterBackupKeyServer(conn dcerpc.Conn, o BackupKeyServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewBackupKeyServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(BackupKeySyntaxV1_0))...)
}

func NewBackupKeyServerHandle(o BackupKeyServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return BackupKeyServerHandle(ctx, o, opNum, r)
	}
}

func BackupKeyServerHandle(ctx context.Context, o BackupKeyServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // BackuprKey
		op := &xxx_BackupKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BackupKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BackupKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented BackupKey
type UnimplementedBackupKeyServer struct {
}

func (UnimplementedBackupKeyServer) BackupKey(context.Context, *BackupKeyRequest) (*BackupKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ BackupKeyServer = (*UnimplementedBackupKeyServer)(nil)
