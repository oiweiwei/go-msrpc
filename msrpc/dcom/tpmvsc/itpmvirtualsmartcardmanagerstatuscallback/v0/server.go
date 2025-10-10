package itpmvirtualsmartcardmanagerstatuscallback

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

// ITpmVirtualSmartCardManagerStatusCallback server interface.
type TpmVirtualSmartCardManagerStatusCallbackServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	ReportProgress(context.Context, *ReportProgressRequest) (*ReportProgressResponse, error)

	ReportError(context.Context, *ReportErrorRequest) (*ReportErrorResponse, error)
}

func RegisterTpmVirtualSmartCardManagerStatusCallbackServer(conn dcerpc.Conn, o TpmVirtualSmartCardManagerStatusCallbackServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewTpmVirtualSmartCardManagerStatusCallbackServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(TpmVirtualSmartCardManagerStatusCallbackSyntaxV0_0))...)
}

func NewTpmVirtualSmartCardManagerStatusCallbackServerHandle(o TpmVirtualSmartCardManagerStatusCallbackServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return TpmVirtualSmartCardManagerStatusCallbackServerHandle(ctx, o, opNum, r)
	}
}

func TpmVirtualSmartCardManagerStatusCallbackServerHandle(ctx context.Context, o TpmVirtualSmartCardManagerStatusCallbackServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // ReportProgress
		op := &xxx_ReportProgressOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReportProgressRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReportProgress(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // ReportError
		op := &xxx_ReportErrorOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReportErrorRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReportError(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented ITpmVirtualSmartCardManagerStatusCallback
type UnimplementedTpmVirtualSmartCardManagerStatusCallbackServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedTpmVirtualSmartCardManagerStatusCallbackServer) ReportProgress(context.Context, *ReportProgressRequest) (*ReportProgressResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTpmVirtualSmartCardManagerStatusCallbackServer) ReportError(context.Context, *ReportErrorRequest) (*ReportErrorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ TpmVirtualSmartCardManagerStatusCallbackServer = (*UnimplementedTpmVirtualSmartCardManagerStatusCallbackServer)(nil)
