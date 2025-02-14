package ivdsvolumeonline

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

// IVdsVolumeOnline server interface.
type VolumeOnlineServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// Online operation.
	Online(context.Context, *OnlineRequest) (*OnlineResponse, error)
}

func RegisterVolumeOnlineServer(conn dcerpc.Conn, o VolumeOnlineServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewVolumeOnlineServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(VolumeOnlineSyntaxV0_0))...)
}

func NewVolumeOnlineServerHandle(o VolumeOnlineServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return VolumeOnlineServerHandle(ctx, o, opNum, r)
	}
}

func VolumeOnlineServerHandle(ctx context.Context, o VolumeOnlineServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Online
		op := &xxx_OnlineOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OnlineRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Online(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IVdsVolumeOnline
type UnimplementedVolumeOnlineServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedVolumeOnlineServer) Online(context.Context, *OnlineRequest) (*OnlineResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ VolumeOnlineServer = (*UnimplementedVolumeOnlineServer)(nil)
