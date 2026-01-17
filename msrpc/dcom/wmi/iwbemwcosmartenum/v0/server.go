package iwbemwcosmartenum

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

// IWbemWCOSmartEnum server interface.
type WCOSmartEnumServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// Next operation.
	Next(context.Context, *NextRequest) (*NextResponse, error)
}

func RegisterWCOSmartEnumServer(conn dcerpc.Conn, o WCOSmartEnumServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewWCOSmartEnumServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(WCOSmartEnumSyntaxV0_0))...)
}

func NewWCOSmartEnumServerHandle(o WCOSmartEnumServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return WCOSmartEnumServerHandle(ctx, o, opNum, r)
	}
}

func WCOSmartEnumServerHandle(ctx context.Context, o WCOSmartEnumServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Next
		op := &xxx_NextOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &NextRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Next(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IWbemWCOSmartEnum
type UnimplementedWCOSmartEnumServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedWCOSmartEnumServer) Next(context.Context, *NextRequest) (*NextResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ WCOSmartEnumServer = (*UnimplementedWCOSmartEnumServer)(nil)
