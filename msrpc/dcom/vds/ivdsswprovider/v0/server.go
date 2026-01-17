package ivdsswprovider

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

// IVdsSwProvider server interface.
type SwProviderServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The QueryPacks method retrieves the provider disk packs.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryPacks(context.Context, *QueryPacksRequest) (*QueryPacksResponse, error)

	// The CreatePack method creates a disk pack.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	CreatePack(context.Context, *CreatePackRequest) (*CreatePackResponse, error)
}

func RegisterSwProviderServer(conn dcerpc.Conn, o SwProviderServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewSwProviderServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(SwProviderSyntaxV0_0))...)
}

func NewSwProviderServerHandle(o SwProviderServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return SwProviderServerHandle(ctx, o, opNum, r)
	}
}

func SwProviderServerHandle(ctx context.Context, o SwProviderServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // QueryPacks
		op := &xxx_QueryPacksOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryPacksRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryPacks(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // CreatePack
		op := &xxx_CreatePackOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreatePackRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreatePack(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IVdsSwProvider
type UnimplementedSwProviderServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedSwProviderServer) QueryPacks(context.Context, *QueryPacksRequest) (*QueryPacksResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSwProviderServer) CreatePack(context.Context, *CreatePackRequest) (*CreatePackResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ SwProviderServer = (*UnimplementedSwProviderServer)(nil)
