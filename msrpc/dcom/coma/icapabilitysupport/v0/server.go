package icapabilitysupport

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
		op := &xxx_StartOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StartRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Start(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // Stop
		op := &xxx_StopOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StopRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Stop(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // Opnum5NotUsedOnWire
		// Opnum5NotUsedOnWire
		return nil, nil
	case 6: // Opnum6NotUsedOnWire
		// Opnum6NotUsedOnWire
		return nil, nil
	case 7: // IsInstalled
		op := &xxx_IsInstalledOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IsInstalledRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.IsInstalled(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // IsRunning
		op := &xxx_IsRunningOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IsRunningRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.IsRunning(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // Opnum9NotUsedOnWire
		// Opnum9NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}

// Unimplemented ICapabilitySupport
type UnimplementedCapabilitySupportServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedCapabilitySupportServer) Start(context.Context, *StartRequest) (*StartResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCapabilitySupportServer) Stop(context.Context, *StopRequest) (*StopResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCapabilitySupportServer) IsInstalled(context.Context, *IsInstalledRequest) (*IsInstalledResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCapabilitySupportServer) IsRunning(context.Context, *IsRunningRequest) (*IsRunningResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ CapabilitySupportServer = (*UnimplementedCapabilitySupportServer)(nil)
