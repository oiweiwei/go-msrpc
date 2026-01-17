package ivdsdiskonline

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

// IVdsDiskOnline server interface.
type DiskOnlineServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// Online operation.
	Online(context.Context, *OnlineRequest) (*OnlineResponse, error)

	// The Offline method brings a disk to the offline state. An offline disk exposes no
	// volume devices.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT, as specified in
	// [MS-ERREF], to indicate success or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	//
	// ERROR_SUCCESS(0x00000000)
	Offline(context.Context, *OfflineRequest) (*OfflineResponse, error)
}

func RegisterDiskOnlineServer(conn dcerpc.Conn, o DiskOnlineServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewDiskOnlineServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(DiskOnlineSyntaxV0_0))...)
}

func NewDiskOnlineServerHandle(o DiskOnlineServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return DiskOnlineServerHandle(ctx, o, opNum, r)
	}
}

func DiskOnlineServerHandle(ctx context.Context, o DiskOnlineServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
	case 4: // Offline
		op := &xxx_OfflineOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OfflineRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Offline(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IVdsDiskOnline
type UnimplementedDiskOnlineServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedDiskOnlineServer) Online(context.Context, *OnlineRequest) (*OnlineResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDiskOnlineServer) Offline(context.Context, *OfflineRequest) (*OfflineResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ DiskOnlineServer = (*UnimplementedDiskOnlineServer)(nil)
