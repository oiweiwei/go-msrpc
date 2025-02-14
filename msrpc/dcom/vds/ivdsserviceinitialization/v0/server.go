package ivdsserviceinitialization

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

// IVdsServiceInitialization server interface.
type ServiceInitializationServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The Initialize method starts the initialization of the server.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	Initialize(context.Context, *InitializeRequest) (*InitializeResponse, error)
}

func RegisterServiceInitializationServer(conn dcerpc.Conn, o ServiceInitializationServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewServiceInitializationServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ServiceInitializationSyntaxV0_0))...)
}

func NewServiceInitializationServerHandle(o ServiceInitializationServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ServiceInitializationServerHandle(ctx, o, opNum, r)
	}
}

func ServiceInitializationServerHandle(ctx context.Context, o ServiceInitializationServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Initialize
		op := &xxx_InitializeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &InitializeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Initialize(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IVdsServiceInitialization
type UnimplementedServiceInitializationServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedServiceInitializationServer) Initialize(context.Context, *InitializeRequest) (*InitializeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ServiceInitializationServer = (*UnimplementedServiceInitializationServer)(nil)
