package iremotesstpcertcheck

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
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
	_ = iunknown.GoPackage
)

// IRemoteSstpCertCheck server interface.
type RemoteSSTPCertCheckServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// CheckIfCertificateAllowedRR operation.
	CheckInterfaceCertificateAllowedRR(context.Context, *CheckInterfaceCertificateAllowedRRRequest) (*CheckInterfaceCertificateAllowedRRResponse, error)
}

func RegisterRemoteSSTPCertCheckServer(conn dcerpc.Conn, o RemoteSSTPCertCheckServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewRemoteSSTPCertCheckServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(RemoteSSTPCertCheckSyntaxV0_0))...)
}

func NewRemoteSSTPCertCheckServerHandle(o RemoteSSTPCertCheckServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return RemoteSSTPCertCheckServerHandle(ctx, o, opNum, r)
	}
}

func RemoteSSTPCertCheckServerHandle(ctx context.Context, o RemoteSSTPCertCheckServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // CheckIfCertificateAllowedRR
		op := &xxx_CheckInterfaceCertificateAllowedRROperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CheckInterfaceCertificateAllowedRRRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CheckInterfaceCertificateAllowedRR(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IRemoteSstpCertCheck
type UnimplementedRemoteSSTPCertCheckServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedRemoteSSTPCertCheckServer) CheckInterfaceCertificateAllowedRR(context.Context, *CheckInterfaceCertificateAllowedRRRequest) (*CheckInterfaceCertificateAllowedRRResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ RemoteSSTPCertCheckServer = (*UnimplementedRemoteSSTPCertCheckServer)(nil)
