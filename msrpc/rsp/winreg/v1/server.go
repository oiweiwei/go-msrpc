package winreg

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
)

// winreg server interface.
type WinregServer interface {

	// Opnum0NotImplemented operation.
	// Opnum0NotImplemented

	// Opnum1NotImplemented operation.
	// Opnum1NotImplemented

	// Opnum2NotImplemented operation.
	// Opnum2NotImplemented

	// Opnum3NotImplemented operation.
	// Opnum3NotImplemented

	// Opnum4NotImplemented operation.
	// Opnum4NotImplemented

	// Opnum5NotImplemented operation.
	// Opnum5NotImplemented

	// Opnum6NotImplemented operation.
	// Opnum6NotImplemented

	// Opnum7NotImplemented operation.
	// Opnum7NotImplemented

	// Opnum8NotImplemented operation.
	// Opnum8NotImplemented

	// Opnum9NotImplemented operation.
	// Opnum9NotImplemented

	// Opnum10NotImplemented operation.
	// Opnum10NotImplemented

	// Opnum11NotImplemented operation.
	// Opnum11NotImplemented

	// Opnum12NotImplemented operation.
	// Opnum12NotImplemented

	// Opnum13NotImplemented operation.
	// Opnum13NotImplemented

	// Opnum14NotImplemented operation.
	// Opnum14NotImplemented

	// Opnum15NotImplemented operation.
	// Opnum15NotImplemented

	// Opnum16NotImplemented operation.
	// Opnum16NotImplemented

	// Opnum17NotImplemented operation.
	// Opnum17NotImplemented

	// Opnum18NotImplemented operation.
	// Opnum18NotImplemented

	// Opnum19NotImplemented operation.
	// Opnum19NotImplemented

	// Opnum20NotImplemented operation.
	// Opnum20NotImplemented

	// Opnum21NotImplemented operation.
	// Opnum21NotImplemented

	// Opnum22NotImplemented operation.
	// Opnum22NotImplemented

	// Opnum23NotImplemented operation.
	// Opnum23NotImplemented

	// The BaseInitiateSystemShutdown method is used to initiate the shutdown of the remote
	// computer.<4>
	//
	// Return Values: The method returns ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it returns a nonzero error code.
	BaseInitiateSystemShutdown(context.Context, *BaseInitiateSystemShutdownRequest) (*BaseInitiateSystemShutdownResponse, error)

	// The BaseAbortSystemShutdown method is used to terminate the shutdown of the remote
	// computer within the waiting period.<5>
	//
	// Return Values: The method returns ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it returns a nonzero error code.
	BaseAbortSystemShutdown(context.Context, *BaseAbortSystemShutdownRequest) (*BaseAbortSystemShutdownResponse, error)

	// Opnum26NotImplemented operation.
	// Opnum26NotImplemented

	// Opnum27NotImplemented operation.
	// Opnum27NotImplemented

	// Opnum28NotImplemented operation.
	// Opnum28NotImplemented

	// Opnum29NotImplemented operation.
	// Opnum29NotImplemented

	// The BaseInitiateSystemShutdownEx method is used to initiate the shutdown of the remote
	// computer with the reason for initiating the shutdown given as a parameter to the
	// call.<6>
	//
	// Return Values: The method returns ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it returns a nonzero error code.
	BaseInitiateSystemShutdownEx(context.Context, *BaseInitiateSystemShutdownExRequest) (*BaseInitiateSystemShutdownExResponse, error)
}

func RegisterWinregServer(conn dcerpc.Conn, o WinregServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewWinregServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(WinregSyntaxV1_0))...)
}

func NewWinregServerHandle(o WinregServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return WinregServerHandle(ctx, o, opNum, r)
	}
}

func WinregServerHandle(ctx context.Context, o WinregServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // Opnum0NotImplemented
		// Opnum0NotImplemented
		return nil, nil
	case 1: // Opnum1NotImplemented
		// Opnum1NotImplemented
		return nil, nil
	case 2: // Opnum2NotImplemented
		// Opnum2NotImplemented
		return nil, nil
	case 3: // Opnum3NotImplemented
		// Opnum3NotImplemented
		return nil, nil
	case 4: // Opnum4NotImplemented
		// Opnum4NotImplemented
		return nil, nil
	case 5: // Opnum5NotImplemented
		// Opnum5NotImplemented
		return nil, nil
	case 6: // Opnum6NotImplemented
		// Opnum6NotImplemented
		return nil, nil
	case 7: // Opnum7NotImplemented
		// Opnum7NotImplemented
		return nil, nil
	case 8: // Opnum8NotImplemented
		// Opnum8NotImplemented
		return nil, nil
	case 9: // Opnum9NotImplemented
		// Opnum9NotImplemented
		return nil, nil
	case 10: // Opnum10NotImplemented
		// Opnum10NotImplemented
		return nil, nil
	case 11: // Opnum11NotImplemented
		// Opnum11NotImplemented
		return nil, nil
	case 12: // Opnum12NotImplemented
		// Opnum12NotImplemented
		return nil, nil
	case 13: // Opnum13NotImplemented
		// Opnum13NotImplemented
		return nil, nil
	case 14: // Opnum14NotImplemented
		// Opnum14NotImplemented
		return nil, nil
	case 15: // Opnum15NotImplemented
		// Opnum15NotImplemented
		return nil, nil
	case 16: // Opnum16NotImplemented
		// Opnum16NotImplemented
		return nil, nil
	case 17: // Opnum17NotImplemented
		// Opnum17NotImplemented
		return nil, nil
	case 18: // Opnum18NotImplemented
		// Opnum18NotImplemented
		return nil, nil
	case 19: // Opnum19NotImplemented
		// Opnum19NotImplemented
		return nil, nil
	case 20: // Opnum20NotImplemented
		// Opnum20NotImplemented
		return nil, nil
	case 21: // Opnum21NotImplemented
		// Opnum21NotImplemented
		return nil, nil
	case 22: // Opnum22NotImplemented
		// Opnum22NotImplemented
		return nil, nil
	case 23: // Opnum23NotImplemented
		// Opnum23NotImplemented
		return nil, nil
	case 24: // BaseInitiateSystemShutdown
		op := &xxx_BaseInitiateSystemShutdownOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BaseInitiateSystemShutdownRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BaseInitiateSystemShutdown(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // BaseAbortSystemShutdown
		op := &xxx_BaseAbortSystemShutdownOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BaseAbortSystemShutdownRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BaseAbortSystemShutdown(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 26: // Opnum26NotImplemented
		// Opnum26NotImplemented
		return nil, nil
	case 27: // Opnum27NotImplemented
		// Opnum27NotImplemented
		return nil, nil
	case 28: // Opnum28NotImplemented
		// Opnum28NotImplemented
		return nil, nil
	case 29: // Opnum29NotImplemented
		// Opnum29NotImplemented
		return nil, nil
	case 30: // BaseInitiateSystemShutdownEx
		op := &xxx_BaseInitiateSystemShutdownExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BaseInitiateSystemShutdownExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BaseInitiateSystemShutdownEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented winreg
type UnimplementedWinregServer struct {
}

func (UnimplementedWinregServer) BaseInitiateSystemShutdown(context.Context, *BaseInitiateSystemShutdownRequest) (*BaseInitiateSystemShutdownResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinregServer) BaseAbortSystemShutdown(context.Context, *BaseAbortSystemShutdownRequest) (*BaseAbortSystemShutdownResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinregServer) BaseInitiateSystemShutdownEx(context.Context, *BaseInitiateSystemShutdownExRequest) (*BaseInitiateSystemShutdownExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ WinregServer = (*UnimplementedWinregServer)(nil)
