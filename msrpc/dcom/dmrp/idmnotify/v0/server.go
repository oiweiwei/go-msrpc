package idmnotify

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

// IDMNotify server interface.
type IDMNotifyServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The ObjectsChanged method notifies the client of object changes.
	//
	// Return Values: The method MUST return 0 or a nonerror HRESULT on success, or an implementation-specific
	// nonzero error code on failure (as specified in [MS-ERREF] section 2.1).
	ObjectsChanged(context.Context, *ObjectsChangedRequest) (*ObjectsChangedResponse, error)
}

func RegisterIDMNotifyServer(conn dcerpc.Conn, o IDMNotifyServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewIDMNotifyServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(IDMNotifySyntaxV0_0))...)
}

func NewIDMNotifyServerHandle(o IDMNotifyServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return IDMNotifyServerHandle(ctx, o, opNum, r)
	}
}

func IDMNotifyServerHandle(ctx context.Context, o IDMNotifyServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // ObjectsChanged
		op := &xxx_ObjectsChangedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ObjectsChangedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ObjectsChanged(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IDMNotify
type UnimplementedIDMNotifyServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedIDMNotifyServer) ObjectsChanged(context.Context, *ObjectsChangedRequest) (*ObjectsChangedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ IDMNotifyServer = (*UnimplementedIDMNotifyServer)(nil)
