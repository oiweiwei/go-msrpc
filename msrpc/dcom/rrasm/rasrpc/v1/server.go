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

	DeleteEntry(context.Context, *DeleteEntryRequest) (*DeleteEntryResponse, error)

	// Opnum6NotUsedOnWire operation.
	// Opnum6NotUsedOnWire

	// Opnum7NotUsedOnWire operation.
	// Opnum7NotUsedOnWire

	// Opnum8NotUsedOnWire operation.
	// Opnum8NotUsedOnWire

	GetUserPreferences(context.Context, *GetUserPreferencesRequest) (*GetUserPreferencesResponse, error)

	SetUserPreferences(context.Context, *SetUserPreferencesRequest) (*SetUserPreferencesResponse, error)

	GetSystemDirectory(context.Context, *GetSystemDirectoryRequest) (*GetSystemDirectoryResponse, error)

	SubmitRequest(context.Context, *SubmitRequestRequest) (*SubmitRequestResponse, error)

	// Opnum13NotUsedOnWire operation.
	// Opnum13NotUsedOnWire

	GetInstalledProtocolsEx(context.Context, *GetInstalledProtocolsExRequest) (*GetInstalledProtocolsExResponse, error)

	GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error)

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
		op := &xxx_DeleteEntryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteEntryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteEntry(ctx, req)
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
		op := &xxx_GetUserPreferencesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetUserPreferencesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetUserPreferences(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // RasRpcSetUserPreferences
		op := &xxx_SetUserPreferencesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetUserPreferencesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetUserPreferences(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // RasRpcGetSystemDirectory
		op := &xxx_GetSystemDirectoryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSystemDirectoryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSystemDirectory(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // RasRpcSubmitRequest
		op := &xxx_SubmitRequestOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SubmitRequestRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SubmitRequest(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // Opnum13NotUsedOnWire
		// Opnum13NotUsedOnWire
		return nil, nil
	case 14: // RasRpcGetInstalledProtocolsEx
		op := &xxx_GetInstalledProtocolsExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetInstalledProtocolsExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetInstalledProtocolsEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // RasRpcGetVersion
		op := &xxx_GetVersionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetVersionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetVersion(ctx, req)
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

func (UnimplementedRasrpcServer) DeleteEntry(context.Context, *DeleteEntryRequest) (*DeleteEntryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRasrpcServer) GetUserPreferences(context.Context, *GetUserPreferencesRequest) (*GetUserPreferencesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRasrpcServer) SetUserPreferences(context.Context, *SetUserPreferencesRequest) (*SetUserPreferencesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRasrpcServer) GetSystemDirectory(context.Context, *GetSystemDirectoryRequest) (*GetSystemDirectoryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRasrpcServer) SubmitRequest(context.Context, *SubmitRequestRequest) (*SubmitRequestResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRasrpcServer) GetInstalledProtocolsEx(context.Context, *GetInstalledProtocolsExRequest) (*GetInstalledProtocolsExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRasrpcServer) GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ RasrpcServer = (*UnimplementedRasrpcServer)(nil)
