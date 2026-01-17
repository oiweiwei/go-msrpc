package ntfrsapi

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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

	// The NtFrsApi_Rpc_Set_DsPollingIntervalW method adjusts the interval at which Active
	// Directory is polled (see section 3.1.5.1) for updates unless both LongInterval and
	// ShortInterval are 0, and then MUST initiate a polling cycle.
	//
	// Return Values: The method MUST return 0 on success or a nonzero error code on failure.
	// All nonzero values MUST be treated as equivalent failures unless otherwise specified.
	//
	//	+--------------------------+------------------------------------+
	//	|          RETURN          |                                    |
	//	|        VALUE/CODE        |            DESCRIPTION             |
	//	|                          |                                    |
	//	+--------------------------+------------------------------------+
	//	+--------------------------+------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The method completed successfully. |
	//	+--------------------------+------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetDSPollingIntervalW(context.Context, *SetDSPollingIntervalWRequest) (*SetDSPollingIntervalWResponse, error)

	// The NtFrsApi_Rpc_Get_DsPollingIntervalW method MUST return the current Active Directory
	// polling intervals.
	//
	// Provides the current, long, and short polling intervals. All polling intervals are
	// in minutes.
	//
	// Return Values: The method MUST return 0 on success or a nonzero error code on failure.
	// All nonzero values MUST be treated as equivalent failures unless otherwise specified.
	//
	//	+--------------------------+------------------------------------+
	//	|          RETURN          |                                    |
	//	|        VALUE/CODE        |            DESCRIPTION             |
	//	|                          |                                    |
	//	+--------------------------+------------------------------------+
	//	+--------------------------+------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The method completed successfully. |
	//	+--------------------------+------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetDSPollingIntervalW(context.Context, *GetDSPollingIntervalWRequest) (*GetDSPollingIntervalWResponse, error)

	// Opnum6NotUsedOnWire operation.
	// Opnum6NotUsedOnWire

	// The NtFrsApi_Rpc_InfoW method MUST return internal information. The method is not
	// used in server-to-server interoperation. The vendor MUST fill the BLOB with implementation-dependent
	// data structures.
	//
	// Return Values: The method MUST return 0 on success or a nonzero error code on failure.
	// All nonzero values MUST be treated as equivalent failures unless otherwise specified.
	//
	//	+--------------------------+------------------------------------+
	//	|          RETURN          |                                    |
	//	|        VALUE/CODE        |            DESCRIPTION             |
	//	|                          |                                    |
	//	+--------------------------+------------------------------------+
	//	+--------------------------+------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The method completed successfully. |
	//	+--------------------------+------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	InfoW(context.Context, *InfoWRequest) (*InfoWResponse, error)

	// This method is not used by FRS.
	//
	// Return Values: The method MUST return 0 on success or a nonzero error code on failure.
	// All nonzero values MUST be treated as equivalent failures unless otherwise specified.
	//
	//	+--------------------------+------------------------------------+
	//	|          RETURN          |                                    |
	//	|        VALUE/CODE        |            DESCRIPTION             |
	//	|                          |                                    |
	//	+--------------------------+------------------------------------+
	//	+--------------------------+------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The method completed successfully. |
	//	+--------------------------+------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	IsPathReplicated(context.Context, *IsPathReplicatedRequest) (*IsPathReplicatedResponse, error)

	// The NtFrsApi_Rpc_WriterCommand method MUST deactivate or reactivate the replication
	// of the specified Replica Set.
	//
	// Return Values: The method MUST return 0 on success or a nonzero error code on failure.
	// All nonzero values MUST be treated as equivalent failures unless otherwise specified.
	//
	//	+--------------------------+------------------------------------+
	//	|          RETURN          |                                    |
	//	|        VALUE/CODE        |            DESCRIPTION             |
	//	|                          |                                    |
	//	+--------------------------+------------------------------------+
	//	+--------------------------+------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The method completed successfully. |
	//	+--------------------------+------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	WriterCommand(context.Context, *WriterCommandRequest) (*WriterCommandResponse, error)

	// This call MUST trigger replication on the connection even if the schedule is off.
	//
	// Return Values: The method MUST return 0 on success or a nonzero error code on failure.
	// All nonzero values MUST be treated as equivalent failures unless otherwise specified.
	//
	//	+--------------------------+------------------------------------+
	//	|          RETURN          |                                    |
	//	|        VALUE/CODE        |            DESCRIPTION             |
	//	|                          |                                    |
	//	+--------------------------+------------------------------------+
	//	+--------------------------+------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The method completed successfully. |
	//	+--------------------------+------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	ForceReplication(context.Context, *ForceReplicationRequest) (*ForceReplicationResponse, error)
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
		op := &xxx_SetDSPollingIntervalWOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetDSPollingIntervalWRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetDSPollingIntervalW(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // NtFrsApi_Rpc_Get_DsPollingIntervalW
		op := &xxx_GetDSPollingIntervalWOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDSPollingIntervalWRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDSPollingIntervalW(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // Opnum6NotUsedOnWire
		// Opnum6NotUsedOnWire
		return nil, nil
	case 7: // NtFrsApi_Rpc_InfoW
		op := &xxx_InfoWOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &InfoWRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.InfoW(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // NtFrsApi_Rpc_IsPathReplicated
		op := &xxx_IsPathReplicatedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IsPathReplicatedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.IsPathReplicated(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // NtFrsApi_Rpc_WriterCommand
		op := &xxx_WriterCommandOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &WriterCommandRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.WriterCommand(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // NtFrsApi_Rpc_ForceReplication
		op := &xxx_ForceReplicationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ForceReplicationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ForceReplication(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented NtFrsApi
type UnimplementedNTFrsAPIServer struct {
}

func (UnimplementedNTFrsAPIServer) SetDSPollingIntervalW(context.Context, *SetDSPollingIntervalWRequest) (*SetDSPollingIntervalWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTFrsAPIServer) GetDSPollingIntervalW(context.Context, *GetDSPollingIntervalWRequest) (*GetDSPollingIntervalWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTFrsAPIServer) InfoW(context.Context, *InfoWRequest) (*InfoWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTFrsAPIServer) IsPathReplicated(context.Context, *IsPathReplicatedRequest) (*IsPathReplicatedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTFrsAPIServer) WriterCommand(context.Context, *WriterCommandRequest) (*WriterCommandResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTFrsAPIServer) ForceReplication(context.Context, *ForceReplicationRequest) (*ForceReplicationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ NTFrsAPIServer = (*UnimplementedNTFrsAPIServer)(nil)
