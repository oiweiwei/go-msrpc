package ntfrsapi

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

// NtFrsApi server interface.
type NTFrsAPIServer interface {

	// Opnum0NotUsedOnWire operation.
	// Opnum0NotUsedOnWire

	// Opnum1NotUsedOnWire operation.
	// Opnum1NotUsedOnWire

	// Opnum2NotUsedOnWire operation.
	// Opnum2NotUsedOnWire

	// Opnum3NotUsedOnWire operation.
	// Opnum3NotUsedOnWire

	NTFrsAPIRPCSetDSPollingIntervalW(context.Context, *NTFrsAPIRPCSetDSPollingIntervalWRequest) (*NTFrsAPIRPCSetDSPollingIntervalWResponse, error)

	NTFrsAPIRPCGetDSPollingIntervalW(context.Context, *NTFrsAPIRPCGetDSPollingIntervalWRequest) (*NTFrsAPIRPCGetDSPollingIntervalWResponse, error)

	// Opnum6NotUsedOnWire operation.
	// Opnum6NotUsedOnWire

	NTFrsAPIRPCInfoW(context.Context, *NTFrsAPIRPCInfoWRequest) (*NTFrsAPIRPCInfoWResponse, error)

	NTFrsAPIRPCIsPathReplicated(context.Context, *NTFrsAPIRPCIsPathReplicatedRequest) (*NTFrsAPIRPCIsPathReplicatedResponse, error)

	NTFrsAPIRPCWriterCommand(context.Context, *NTFrsAPIRPCWriterCommandRequest) (*NTFrsAPIRPCWriterCommandResponse, error)

	NTFrsAPIRPCForceReplication(context.Context, *NTFrsAPIRPCForceReplicationRequest) (*NTFrsAPIRPCForceReplicationResponse, error)
}

func RegisterNTFrsAPIServer(conn dcerpc.Conn, o NTFrsAPIServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewNTFrsAPIServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(NTFrsAPISyntaxV1_1))...)
}

func NewNTFrsAPIServerHandle(o NTFrsAPIServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return NTFrsAPIServerHandle(ctx, o, opNum, r)
	}
}

func NTFrsAPIServerHandle(ctx context.Context, o NTFrsAPIServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
	case 4: // NtFrsApi_Rpc_Set_DsPollingIntervalW
		op := &xxx_NTFrsAPIRPCSetDSPollingIntervalWOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &NTFrsAPIRPCSetDSPollingIntervalWRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.NTFrsAPIRPCSetDSPollingIntervalW(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // NtFrsApi_Rpc_Get_DsPollingIntervalW
		op := &xxx_NTFrsAPIRPCGetDSPollingIntervalWOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &NTFrsAPIRPCGetDSPollingIntervalWRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.NTFrsAPIRPCGetDSPollingIntervalW(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // Opnum6NotUsedOnWire
		// Opnum6NotUsedOnWire
		return nil, nil
	case 7: // NtFrsApi_Rpc_InfoW
		op := &xxx_NTFrsAPIRPCInfoWOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &NTFrsAPIRPCInfoWRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.NTFrsAPIRPCInfoW(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // NtFrsApi_Rpc_IsPathReplicated
		op := &xxx_NTFrsAPIRPCIsPathReplicatedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &NTFrsAPIRPCIsPathReplicatedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.NTFrsAPIRPCIsPathReplicated(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // NtFrsApi_Rpc_WriterCommand
		op := &xxx_NTFrsAPIRPCWriterCommandOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &NTFrsAPIRPCWriterCommandRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.NTFrsAPIRPCWriterCommand(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // NtFrsApi_Rpc_ForceReplication
		op := &xxx_NTFrsAPIRPCForceReplicationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &NTFrsAPIRPCForceReplicationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.NTFrsAPIRPCForceReplication(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented NtFrsApi
type UnimplementedNTFrsAPIServer struct {
}

func (UnimplementedNTFrsAPIServer) NTFrsAPIRPCSetDSPollingIntervalW(context.Context, *NTFrsAPIRPCSetDSPollingIntervalWRequest) (*NTFrsAPIRPCSetDSPollingIntervalWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTFrsAPIServer) NTFrsAPIRPCGetDSPollingIntervalW(context.Context, *NTFrsAPIRPCGetDSPollingIntervalWRequest) (*NTFrsAPIRPCGetDSPollingIntervalWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTFrsAPIServer) NTFrsAPIRPCInfoW(context.Context, *NTFrsAPIRPCInfoWRequest) (*NTFrsAPIRPCInfoWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTFrsAPIServer) NTFrsAPIRPCIsPathReplicated(context.Context, *NTFrsAPIRPCIsPathReplicatedRequest) (*NTFrsAPIRPCIsPathReplicatedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTFrsAPIServer) NTFrsAPIRPCWriterCommand(context.Context, *NTFrsAPIRPCWriterCommandRequest) (*NTFrsAPIRPCWriterCommandResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTFrsAPIServer) NTFrsAPIRPCForceReplication(context.Context, *NTFrsAPIRPCForceReplicationRequest) (*NTFrsAPIRPCForceReplicationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ NTFrsAPIServer = (*UnimplementedNTFrsAPIServer)(nil)
