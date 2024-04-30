package iregister

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
		in := &RegisterModuleRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RegisterModule(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // Opnum4NotUsedOnWire
		// Opnum4NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}
