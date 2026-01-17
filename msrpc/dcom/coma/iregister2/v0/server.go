package iregister2

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

// IRegister2 server interface.
type Register2Server interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// This method is called by a client to create a component full configuration for an
	// existing component in an existing conglomeration in the global partition.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	CreateFullConfiguration(context.Context, *CreateFullConfigurationRequest) (*CreateFullConfigurationResponse, error)

	// This method is called by a client to create a component legacy configuration for
	// an existing component in an existing conglomeration in the global partition.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	CreateLegacyConfiguration(context.Context, *CreateLegacyConfigurationRequest) (*CreateLegacyConfigurationResponse, error)

	// This method is called by a client to convert an existing component legacy configuration
	// for a component into a component full configuration for that component.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result (as specified in [MS-ERREF] section 2.1) on failure. All failure results MUST
	// be treated identically.
	PromoteLegacyConfiguration(context.Context, *PromoteLegacyConfigurationRequest) (*PromoteLegacyConfigurationResponse, error)

	// Opnum6NotUsedOnWire operation.
	// Opnum6NotUsedOnWire

	// Opnum7NotUsedOnWire operation.
	// Opnum7NotUsedOnWire

	// This method is called by a client to register the components in one or more modules
	// and to create component full configurations for those modules in an existing conglomeration.
	// This method supports conglomerations in any partition.
	//
	// Alternatively, this method can be called to verify modules without actually registering
	// the components. As a side effect, this method returns implementation-specific detailed
	// results of the registration or verification operation for informational purposes.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1,  on failure. All failure results
	// MUST be treated identically.
	RegisterModule2(context.Context, *RegisterModule2Request) (*RegisterModule2Response, error)

	// Opnum9NotUsedOnWire operation.
	// Opnum9NotUsedOnWire
}

func RegisterRegister2Server(conn dcerpc.Conn, o Register2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewRegister2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Register2SyntaxV0_0))...)
}

func NewRegister2ServerHandle(o Register2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Register2ServerHandle(ctx, o, opNum, r)
	}
}

func Register2ServerHandle(ctx context.Context, o Register2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // CreateFullConfiguration
		op := &xxx_CreateFullConfigurationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateFullConfigurationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateFullConfiguration(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // CreateLegacyConfiguration
		op := &xxx_CreateLegacyConfigurationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateLegacyConfigurationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateLegacyConfiguration(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // PromoteLegacyConfiguration
		op := &xxx_PromoteLegacyConfigurationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PromoteLegacyConfigurationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PromoteLegacyConfiguration(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // Opnum6NotUsedOnWire
		// Opnum6NotUsedOnWire
		return nil, nil
	case 7: // Opnum7NotUsedOnWire
		// Opnum7NotUsedOnWire
		return nil, nil
	case 8: // RegisterModule2
		op := &xxx_RegisterModule2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RegisterModule2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RegisterModule2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // Opnum9NotUsedOnWire
		// Opnum9NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}

// Unimplemented IRegister2
type UnimplementedRegister2Server struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedRegister2Server) CreateFullConfiguration(context.Context, *CreateFullConfigurationRequest) (*CreateFullConfigurationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRegister2Server) CreateLegacyConfiguration(context.Context, *CreateLegacyConfigurationRequest) (*CreateLegacyConfigurationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRegister2Server) PromoteLegacyConfiguration(context.Context, *PromoteLegacyConfigurationRequest) (*PromoteLegacyConfigurationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRegister2Server) RegisterModule2(context.Context, *RegisterModule2Request) (*RegisterModule2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ Register2Server = (*UnimplementedRegister2Server)(nil)
