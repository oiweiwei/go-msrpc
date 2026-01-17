package ivdshwprovider

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

// IVdsHwProvider server interface.
type HwProviderServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The QuerySubSystems method retrieves the subsystems that are managed by the provider.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QuerySubSystems(context.Context, *QuerySubSystemsRequest) (*QuerySubSystemsResponse, error)

	// Opnum04NotUsedOnWire operation.
	// Opnum04NotUsedOnWire

	// Opnum05NotUsedOnWire operation.
	// Opnum05NotUsedOnWire
}

func RegisterHwProviderServer(conn dcerpc.Conn, o HwProviderServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewHwProviderServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(HwProviderSyntaxV0_0))...)
}

func NewHwProviderServerHandle(o HwProviderServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return HwProviderServerHandle(ctx, o, opNum, r)
	}
}

func HwProviderServerHandle(ctx context.Context, o HwProviderServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // QuerySubSystems
		op := &xxx_QuerySubSystemsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QuerySubSystemsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QuerySubSystems(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // Opnum04NotUsedOnWire
		// Opnum04NotUsedOnWire
		return nil, nil
	case 5: // Opnum05NotUsedOnWire
		// Opnum05NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}

// Unimplemented IVdsHwProvider
type UnimplementedHwProviderServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedHwProviderServer) QuerySubSystems(context.Context, *QuerySubSystemsRequest) (*QuerySubSystemsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ HwProviderServer = (*UnimplementedHwProviderServer)(nil)
