package iremoteicficsconfig

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

// IRemoteICFICSConfig server interface.
type RemoteIcficsConfigServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// GetIcfEnabled operation.
	GetIcfEnabled(context.Context, *GetIcfEnabledRequest) (*GetIcfEnabledResponse, error)

	// GetIcsEnabled operation.
	GetIcsEnabled(context.Context, *GetIcsEnabledRequest) (*GetIcsEnabledResponse, error)
}

func RegisterRemoteIcficsConfigServer(conn dcerpc.Conn, o RemoteIcficsConfigServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewRemoteIcficsConfigServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(RemoteIcficsConfigSyntaxV0_0))...)
}

func NewRemoteIcficsConfigServerHandle(o RemoteIcficsConfigServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return RemoteIcficsConfigServerHandle(ctx, o, opNum, r)
	}
}

func RemoteIcficsConfigServerHandle(ctx context.Context, o RemoteIcficsConfigServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetIcfEnabled
		op := &xxx_GetIcfEnabledOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIcfEnabledRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIcfEnabled(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // GetIcsEnabled
		op := &xxx_GetIcsEnabledOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIcsEnabledRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIcsEnabled(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IRemoteICFICSConfig
type UnimplementedRemoteIcficsConfigServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedRemoteIcficsConfigServer) GetIcfEnabled(context.Context, *GetIcfEnabledRequest) (*GetIcfEnabledResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteIcficsConfigServer) GetIcsEnabled(context.Context, *GetIcsEnabledRequest) (*GetIcsEnabledResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ RemoteIcficsConfigServer = (*UnimplementedRemoteIcficsConfigServer)(nil)
