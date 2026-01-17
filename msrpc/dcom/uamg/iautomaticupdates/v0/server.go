package iautomaticupdates

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
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
	_ = idispatch.GoPackage
)

// IAutomaticUpdates server interface.
type AutomaticUpdatesServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The server SHOULD trigger the automatic update agent to perform an update search.
	DetectNow(context.Context, *DetectNowRequest) (*DetectNowResponse, error)

	// Opnum9NotUsedOnWire operation.
	// Opnum9NotUsedOnWire

	// Opnum10NotUsedOnWire operation.
	// Opnum10NotUsedOnWire

	// Opnum11NotUsedOnWire operation.
	// Opnum11NotUsedOnWire

	// Opnum12NotUsedOnWire operation.
	// Opnum12NotUsedOnWire

	// Opnum13NotUsedOnWire operation.
	// Opnum13NotUsedOnWire

	// Opnum14NotUsedOnWire operation.
	// Opnum14NotUsedOnWire
}

func RegisterAutomaticUpdatesServer(conn dcerpc.Conn, o AutomaticUpdatesServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAutomaticUpdatesServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AutomaticUpdatesSyntaxV0_0))...)
}

func NewAutomaticUpdatesServerHandle(o AutomaticUpdatesServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AutomaticUpdatesServerHandle(ctx, o, opNum, r)
	}
}

func AutomaticUpdatesServerHandle(ctx context.Context, o AutomaticUpdatesServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // DetectNow
		op := &xxx_DetectNowOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DetectNowRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DetectNow(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // Opnum9NotUsedOnWire
		// Opnum9NotUsedOnWire
		return nil, nil
	case 9: // Opnum10NotUsedOnWire
		// Opnum10NotUsedOnWire
		return nil, nil
	case 10: // Opnum11NotUsedOnWire
		// Opnum11NotUsedOnWire
		return nil, nil
	case 11: // Opnum12NotUsedOnWire
		// Opnum12NotUsedOnWire
		return nil, nil
	case 12: // Opnum13NotUsedOnWire
		// Opnum13NotUsedOnWire
		return nil, nil
	case 13: // Opnum14NotUsedOnWire
		// Opnum14NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}

// Unimplemented IAutomaticUpdates
type UnimplementedAutomaticUpdatesServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedAutomaticUpdatesServer) DetectNow(context.Context, *DetectNowRequest) (*DetectNowResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AutomaticUpdatesServer = (*UnimplementedAutomaticUpdatesServer)(nil)
