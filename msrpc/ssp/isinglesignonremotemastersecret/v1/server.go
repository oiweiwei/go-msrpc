package isinglesignonremotemastersecret

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
)

// ISingleSignonRemoteMasterSecret server interface.
type SSORemoteMasterSecretServer interface {

	// RemoteGetMasterSecret operation.
	RemoteGetMasterSecret(context.Context, *RemoteGetMasterSecretRequest) (*RemoteGetMasterSecretResponse, error)
}

func RegisterSSORemoteMasterSecretServer(conn dcerpc.Conn, o SSORemoteMasterSecretServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewSSORemoteMasterSecretServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(SSORemoteMasterSecretSyntaxV1_0))...)
}

func NewSSORemoteMasterSecretServerHandle(o SSORemoteMasterSecretServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return SSORemoteMasterSecretServerHandle(ctx, o, opNum, r)
	}
}

func SSORemoteMasterSecretServerHandle(ctx context.Context, o SSORemoteMasterSecretServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // RemoteGetMasterSecret
		op := &xxx_RemoteGetMasterSecretOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoteGetMasterSecretRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoteGetMasterSecret(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented ISingleSignonRemoteMasterSecret
type UnimplementedSSORemoteMasterSecretServer struct {
}

func (UnimplementedSSORemoteMasterSecretServer) RemoteGetMasterSecret(context.Context, *RemoteGetMasterSecretRequest) (*RemoteGetMasterSecretResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ SSORemoteMasterSecretServer = (*UnimplementedSSORemoteMasterSecretServer)(nil)
