package iregister

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

// IRegister server interface.
type RegisterServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// This method is called by a client to register the components in one or more modules
	// and to create component full configurations for those modules in an existing conglomeration.
	// This method supports conglomerations in the global partition only.
	//
	// Alternatively, this method can be called to verify modules without actually registering
	// the components. As a side effect, this method returns implementation-specific detailed
	// results of the registration or verification operation for informational purposes.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	RegisterModule(context.Context, *RegisterModuleRequest) (*RegisterModuleResponse, error)

	// Opnum4NotUsedOnWire operation.
	// Opnum4NotUsedOnWire
}

func RegisterRegisterServer(conn dcerpc.Conn, o RegisterServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewRegisterServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(RegisterSyntaxV0_0))...)
}

func NewRegisterServerHandle(o RegisterServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return RegisterServerHandle(ctx, o, opNum, r)
	}
}

func RegisterServerHandle(ctx context.Context, o RegisterServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // RegisterModule
		op := &xxx_RegisterModuleOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RegisterModuleRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RegisterModule(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // Opnum4NotUsedOnWire
		// Opnum4NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}

// Unimplemented IRegister
type UnimplementedRegisterServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedRegisterServer) RegisterModule(context.Context, *RegisterModuleRequest) (*RegisterModuleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ RegisterServer = (*UnimplementedRegisterServer)(nil)
