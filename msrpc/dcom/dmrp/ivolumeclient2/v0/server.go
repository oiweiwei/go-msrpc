package ivolumeclient2

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

// IVolumeClient2 server interface.
type VolumeClient2Server interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// GetMaxAdjustedFreeSpace operation.
	GetMaxAdjustedFreeSpace(context.Context, *GetMaxAdjustedFreeSpaceRequest) (*GetMaxAdjustedFreeSpaceResponse, error)
}

func RegisterVolumeClient2Server(conn dcerpc.Conn, o VolumeClient2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewVolumeClient2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(VolumeClient2SyntaxV0_0))...)
}

func NewVolumeClient2ServerHandle(o VolumeClient2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return VolumeClient2ServerHandle(ctx, o, opNum, r)
	}
}

func VolumeClient2ServerHandle(ctx context.Context, o VolumeClient2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetMaxAdjustedFreeSpace
		op := &xxx_GetMaxAdjustedFreeSpaceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMaxAdjustedFreeSpaceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMaxAdjustedFreeSpace(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IVolumeClient2
type UnimplementedVolumeClient2Server struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedVolumeClient2Server) GetMaxAdjustedFreeSpace(context.Context, *GetMaxAdjustedFreeSpaceRequest) (*GetMaxAdjustedFreeSpaceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ VolumeClient2Server = (*UnimplementedVolumeClient2Server)(nil)
