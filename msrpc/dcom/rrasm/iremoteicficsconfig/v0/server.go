package iremoteicficsconfig

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = iunknown.GoPackage
)

// IRemoteICFICSConfig server interface.
type RemoteICFICSConfigServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// GetIcfEnabled operation.
	GetICFEnabled(context.Context, *GetICFEnabledRequest) (*GetICFEnabledResponse, error)

	// GetIcsEnabled operation.
	GetICSEnabled(context.Context, *GetICSEnabledRequest) (*GetICSEnabledResponse, error)
}

func RegisterRemoteICFICSConfigServer(conn dcerpc.Conn, o RemoteICFICSConfigServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewRemoteICFICSConfigServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(RemoteICFICSConfigSyntaxV0_0))...)
}

func NewRemoteICFICSConfigServerHandle(o RemoteICFICSConfigServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return RemoteICFICSConfigServerHandle(ctx, o, opNum, r)
	}
}

func RemoteICFICSConfigServerHandle(ctx context.Context, o RemoteICFICSConfigServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetIcfEnabled
		op := &xxx_GetICFEnabledOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetICFEnabledRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetICFEnabled(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // GetIcsEnabled
		op := &xxx_GetICSEnabledOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetICSEnabledRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetICSEnabled(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IRemoteICFICSConfig
type UnimplementedRemoteICFICSConfigServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedRemoteICFICSConfigServer) GetICFEnabled(context.Context, *GetICFEnabledRequest) (*GetICFEnabledResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteICFICSConfigServer) GetICSEnabled(context.Context, *GetICSEnabledRequest) (*GetICSEnabledResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ RemoteICFICSConfigServer = (*UnimplementedRemoteICFICSConfigServer)(nil)
