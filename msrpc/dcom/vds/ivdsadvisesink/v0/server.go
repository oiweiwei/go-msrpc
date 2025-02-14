package ivdsadvisesink

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

// IVdsAdviseSink server interface.
type AdviseSinkServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The OnNotify method passes notifications from VDS to applications.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	OnNotify(context.Context, *OnNotifyRequest) (*OnNotifyResponse, error)
}

func RegisterAdviseSinkServer(conn dcerpc.Conn, o AdviseSinkServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAdviseSinkServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AdviseSinkSyntaxV0_0))...)
}

func NewAdviseSinkServerHandle(o AdviseSinkServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AdviseSinkServerHandle(ctx, o, opNum, r)
	}
}

func AdviseSinkServerHandle(ctx context.Context, o AdviseSinkServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // OnNotify
		op := &xxx_OnNotifyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OnNotifyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OnNotify(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IVdsAdviseSink
type UnimplementedAdviseSinkServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedAdviseSinkServer) OnNotify(context.Context, *OnNotifyRequest) (*OnNotifyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AdviseSinkServer = (*UnimplementedAdviseSinkServer)(nil)
