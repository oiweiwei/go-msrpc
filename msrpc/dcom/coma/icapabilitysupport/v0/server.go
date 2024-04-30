package icapabilitysupport

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

// ICapabilitySupport server interface.
type CapabilitySupportServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// This method is called by a client to start instance load balancing.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	Start(context.Context, *StartRequest) (*StartResponse, error)

	// This method is called by a client to stop instance load balancing.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	Stop(context.Context, *StopRequest) (*StopResponse, error)

	// Opnum5NotUsedOnWire operation.
	// Opnum5NotUsedOnWire

	// Opnum6NotUsedOnWire operation.
	// Opnum6NotUsedOnWire

	// This method is called by a client to determine whether instance load balancing is
	// installed.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	IsInstalled(context.Context, *IsInstalledRequest) (*IsInstalledResponse, error)

	// This method is called by a client to determine whether instance load balancing is
	// running.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	IsRunning(context.Context, *IsRunningRequest) (*IsRunningResponse, error)

	// Opnum9NotUsedOnWire operation.
	// Opnum9NotUsedOnWire
}

func RegisterCapabilitySupportServer(conn dcerpc.Conn, o CapabilitySupportServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewCapabilitySupportServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(CapabilitySupportSyntaxV0_0))...)
}

func NewCapabilitySupportServerHandle(o CapabilitySupportServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return CapabilitySupportServerHandle(ctx, o, opNum, r)
	}
}

func CapabilitySupportServerHandle(ctx context.Context, o CapabilitySupportServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Start
		in := &StartRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Start(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // Stop
		in := &StopRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Stop(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // Opnum5NotUsedOnWire
		// Opnum5NotUsedOnWire
		return nil, nil
	case 6: // Opnum6NotUsedOnWire
		// Opnum6NotUsedOnWire
		return nil, nil
	case 7: // IsInstalled
		in := &IsInstalledRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.IsInstalled(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // IsRunning
		in := &IsRunningRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.IsRunning(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // Opnum9NotUsedOnWire
		// Opnum9NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}
