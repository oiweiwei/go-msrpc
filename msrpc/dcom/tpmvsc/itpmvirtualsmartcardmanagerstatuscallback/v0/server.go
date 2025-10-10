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
type VirtualSmartCardManagerStatusCallbackServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	ReportProgress(context.Context, *ReportProgressRequest) (*ReportProgressResponse, error)

	ReportError(context.Context, *ReportErrorRequest) (*ReportErrorResponse, error)
}

func RegisterVirtualSmartCardManagerStatusCallbackServer(conn dcerpc.Conn, o VirtualSmartCardManagerStatusCallbackServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewVirtualSmartCardManagerStatusCallbackServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(VirtualSmartCardManagerStatusCallbackSyntaxV0_0))...)
}

func NewVirtualSmartCardManagerStatusCallbackServerHandle(o VirtualSmartCardManagerStatusCallbackServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return VirtualSmartCardManagerStatusCallbackServerHandle(ctx, o, opNum, r)
	}
}

func VirtualSmartCardManagerStatusCallbackServerHandle(ctx context.Context, o VirtualSmartCardManagerStatusCallbackServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
type UnimplementedVirtualSmartCardManagerStatusCallbackServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedVirtualSmartCardManagerStatusCallbackServer) ReportProgress(context.Context, *ReportProgressRequest) (*ReportProgressResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedVirtualSmartCardManagerStatusCallbackServer) ReportError(context.Context, *ReportErrorRequest) (*ReportErrorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ VirtualSmartCardManagerStatusCallbackServer = (*UnimplementedVirtualSmartCardManagerStatusCallbackServer)(nil)
