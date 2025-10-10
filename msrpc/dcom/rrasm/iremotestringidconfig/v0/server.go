package iremotestringidconfig

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

// IRemoteStringIdConfig server interface.
type RemoteStringIDConfigServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// GetStringFromId operation.
	GetStringFromID(context.Context, *GetStringFromIDRequest) (*GetStringFromIDResponse, error)
}

func RegisterRemoteStringIDConfigServer(conn dcerpc.Conn, o RemoteStringIDConfigServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewRemoteStringIDConfigServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(RemoteStringIDConfigSyntaxV0_0))...)
}

func NewRemoteStringIDConfigServerHandle(o RemoteStringIDConfigServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return RemoteStringIDConfigServerHandle(ctx, o, opNum, r)
	}
}

func RemoteStringIDConfigServerHandle(ctx context.Context, o RemoteStringIDConfigServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetStringFromId
		op := &xxx_GetStringFromIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetStringFromIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetStringFromID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IRemoteStringIdConfig
type UnimplementedRemoteStringIDConfigServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedRemoteStringIDConfigServer) GetStringFromID(context.Context, *GetStringFromIDRequest) (*GetStringFromIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ RemoteStringIDConfigServer = (*UnimplementedRemoteStringIDConfigServer)(nil)
