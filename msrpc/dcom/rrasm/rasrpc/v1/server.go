package rasrpc

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

// rasrpc server interface.
type RasrpcServer interface {

	// Opnum0NotUsedOnWire operation.
	// Opnum0NotUsedOnWire

	// Opnum1NotUsedOnWire operation.
	// Opnum1NotUsedOnWire

	// Opnum2NotUsedOnWire operation.
	// Opnum2NotUsedOnWire

	// Opnum3NotUsedOnWire operation.
	// Opnum3NotUsedOnWire

	// Opnum4NotUsedOnWire operation.
	// Opnum4NotUsedOnWire

	RASRPCDeleteEntry(context.Context, *RASRPCDeleteEntryRequest) (*RASRPCDeleteEntryResponse, error)

	// Opnum6NotUsedOnWire operation.
	// Opnum6NotUsedOnWire

	// Opnum7NotUsedOnWire operation.
	// Opnum7NotUsedOnWire

	// Opnum8NotUsedOnWire operation.
	// Opnum8NotUsedOnWire

	RASRPCGetUserPreferences(context.Context, *RASRPCGetUserPreferencesRequest) (*RASRPCGetUserPreferencesResponse, error)

	RASRPCSetUserPreferences(context.Context, *RASRPCSetUserPreferencesRequest) (*RASRPCSetUserPreferencesResponse, error)

	RASRPCGetSystemDirectory(context.Context, *RASRPCGetSystemDirectoryRequest) (*RASRPCGetSystemDirectoryResponse, error)

	RASRPCSubmitRequest(context.Context, *RASRPCSubmitRequestRequest) (*RASRPCSubmitRequestResponse, error)

	// Opnum13NotUsedOnWire operation.
	// Opnum13NotUsedOnWire

	RASRPCGetInstalledProtocolsEx(context.Context, *RASRPCGetInstalledProtocolsExRequest) (*RASRPCGetInstalledProtocolsExResponse, error)

	RASRPCGetVersion(context.Context, *RASRPCGetVersionRequest) (*RASRPCGetVersionResponse, error)

	// Opnum16NotUsedOnWire operation.
	// Opnum16NotUsedOnWire
}

func RegisterRasrpcServer(conn dcerpc.Conn, o RasrpcServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewRasrpcServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(RasrpcSyntaxV1_0))...)
}

func NewRasrpcServerHandle(o RasrpcServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return RasrpcServerHandle(ctx, o, opNum, r)
	}
}

func RasrpcServerHandle(ctx context.Context, o RasrpcServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // Opnum0NotUsedOnWire
		// Opnum0NotUsedOnWire
		return nil, nil
	case 1: // Opnum1NotUsedOnWire
		// Opnum1NotUsedOnWire
		return nil, nil
	case 2: // Opnum2NotUsedOnWire
		// Opnum2NotUsedOnWire
		return nil, nil
	case 3: // Opnum3NotUsedOnWire
		// Opnum3NotUsedOnWire
		return nil, nil
	case 4: // Opnum4NotUsedOnWire
		// Opnum4NotUsedOnWire
		return nil, nil
	case 5: // RasRpcDeleteEntry
		op := &xxx_RASRPCDeleteEntryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RASRPCDeleteEntryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RASRPCDeleteEntry(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // Opnum6NotUsedOnWire
		// Opnum6NotUsedOnWire
		return nil, nil
	case 7: // Opnum7NotUsedOnWire
		// Opnum7NotUsedOnWire
		return nil, nil
	case 8: // Opnum8NotUsedOnWire
		// Opnum8NotUsedOnWire
		return nil, nil
	case 9: // RasRpcGetUserPreferences
		op := &xxx_RASRPCGetUserPreferencesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RASRPCGetUserPreferencesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RASRPCGetUserPreferences(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // RasRpcSetUserPreferences
		op := &xxx_RASRPCSetUserPreferencesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RASRPCSetUserPreferencesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RASRPCSetUserPreferences(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // RasRpcGetSystemDirectory
		op := &xxx_RASRPCGetSystemDirectoryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RASRPCGetSystemDirectoryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RASRPCGetSystemDirectory(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // RasRpcSubmitRequest
		op := &xxx_RASRPCSubmitRequestOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RASRPCSubmitRequestRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RASRPCSubmitRequest(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // Opnum13NotUsedOnWire
		// Opnum13NotUsedOnWire
		return nil, nil
	case 14: // RasRpcGetInstalledProtocolsEx
		op := &xxx_RASRPCGetInstalledProtocolsExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RASRPCGetInstalledProtocolsExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RASRPCGetInstalledProtocolsEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // RasRpcGetVersion
		op := &xxx_RASRPCGetVersionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RASRPCGetVersionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RASRPCGetVersion(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // Opnum16NotUsedOnWire
		// Opnum16NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}

// Unimplemented rasrpc
type UnimplementedRasrpcServer struct {
}

func (UnimplementedRasrpcServer) RASRPCDeleteEntry(context.Context, *RASRPCDeleteEntryRequest) (*RASRPCDeleteEntryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRasrpcServer) RASRPCGetUserPreferences(context.Context, *RASRPCGetUserPreferencesRequest) (*RASRPCGetUserPreferencesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRasrpcServer) RASRPCSetUserPreferences(context.Context, *RASRPCSetUserPreferencesRequest) (*RASRPCSetUserPreferencesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRasrpcServer) RASRPCGetSystemDirectory(context.Context, *RASRPCGetSystemDirectoryRequest) (*RASRPCGetSystemDirectoryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRasrpcServer) RASRPCSubmitRequest(context.Context, *RASRPCSubmitRequestRequest) (*RASRPCSubmitRequestResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRasrpcServer) RASRPCGetInstalledProtocolsEx(context.Context, *RASRPCGetInstalledProtocolsExRequest) (*RASRPCGetInstalledProtocolsExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRasrpcServer) RASRPCGetVersion(context.Context, *RASRPCGetVersionRequest) (*RASRPCGetVersionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ RasrpcServer = (*UnimplementedRasrpcServer)(nil)
