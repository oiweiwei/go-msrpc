package ipchservice

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

// IPCHService server interface.
type PCHServiceServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// Opnum7NotUsedByProtocol operation.
	// Opnum7NotUsedByProtocol

	// Opnum8NotUsedByProtocol operation.
	// Opnum8NotUsedByProtocol

	// Opnum9NotUsedByProtocol operation.
	// Opnum9NotUsedByProtocol

	// Opnum10NotUsedByProtocol operation.
	// Opnum10NotUsedByProtocol

	// Opnum11NotUsedByProtocol operation.
	// Opnum11NotUsedByProtocol

	// Opnum12NotUsedByProtocol operation.
	// Opnum12NotUsedByProtocol

	// Opnum13NotUsedByProtocol operation.
	// Opnum13NotUsedByProtocol

	// Opnum14NotUsedByProtocol operation.
	// Opnum14NotUsedByProtocol

	// Opnum15NotUsedByProtocol operation.
	// Opnum15NotUsedByProtocol

	// Opnum16NotUsedByProtocol operation.
	// Opnum16NotUsedByProtocol

	// Opnum17NotUsedByProtocol operation.
	// Opnum17NotUsedByProtocol

	// Opnum18NotUsedByProtocol operation.
	// Opnum18NotUsedByProtocol

	// The RemoteConnectionParms method gets the Remote Assistance connection parameters
	// for a specific UserName, DomainName, and SessionID triple.
	//
	// Return Values: A signed 32-bit value indicating return status. This method MUST return
	// zero to indicate success, or an HRESULT error value (as specified in [MS-ERREF] section
	// 2.1.1) to indicate failure. If the UserName and DomainName are valid BSTRs, the return
	// code is one listed in the following table. If the UserName and DomainName are invalid
	// BSTRs, the HRESULT value returned is the corresponding HRESULT to the system error
	// code ERROR_NONE_MAPPED.
	//
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	|                   RETURN                   |                                                                                  |
	//	|                 VALUE/CODE                 |                                   DESCRIPTION                                    |
	//	|                                            |                                                                                  |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                            | The call was successful.                                                         |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED                  | General access denied error. <8>                                                 |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8007000E E_OUTOFMEMORY                   | Out of memory.                                                                   |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800704EC ERROR_ACCESS_DISABLED_BY_POLICY | The program cannot be opened because of a software restriction policy. For more  |
	//	|                                            | information, contact the system administrator.                                   |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown:
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol [MS-RPCE].
	RemoteConnectionParameters(context.Context, *RemoteConnectionParametersRequest) (*RemoteConnectionParametersResponse, error)

	// The RemoteUserSessionInfo method returns the collection of the terminal services
	// sessions on the remote novice machine. All the terminal services session information
	// is returned as a standard IPCHCollection interface. The members of this collection
	// are objects of type ISAFSession. ISAFSession includes the DomainName, SessionID,
	// SessionState, and UserName for each session.
	//
	// Return Values: A signed 32-bit value indicating return status. This method MUST return
	// zero to indicate success, or an HRESULT error value (as specified in [MS-ERREF] section
	// 2.1.1) to indicate failure.
	//
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	|                   RETURN                   |                                                                                  |
	//	|                 VALUE/CODE                 |                                   DESCRIPTION                                    |
	//	|                                            |                                                                                  |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                            | The call was successful.                                                         |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED                  | General access denied error. <9>                                                 |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8007000E E_OUTOFMEMORY                   | Out of memory.                                                                   |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800704EC ERROR_ACCESS_DISABLED_BY_POLICY | The program cannot be opened because of a software restriction policy. For more  |
	//	|                                            | information, contact the system administrator.                                   |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown:
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol [MS-RPCE].
	RemoteUserSessionInfo(context.Context, *RemoteUserSessionInfoRequest) (*RemoteUserSessionInfoResponse, error)

	// Opnum21NotUsedByProtocol operation.
	// Opnum21NotUsedByProtocol
}

func RegisterPCHServiceServer(conn dcerpc.Conn, o PCHServiceServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewPCHServiceServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(PCHServiceSyntaxV0_0))...)
}

func NewPCHServiceServerHandle(o PCHServiceServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return PCHServiceServerHandle(ctx, o, opNum, r)
	}
}

func PCHServiceServerHandle(ctx context.Context, o PCHServiceServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Opnum7NotUsedByProtocol
		// Opnum7NotUsedByProtocol
		return nil, nil
	case 8: // Opnum8NotUsedByProtocol
		// Opnum8NotUsedByProtocol
		return nil, nil
	case 9: // Opnum9NotUsedByProtocol
		// Opnum9NotUsedByProtocol
		return nil, nil
	case 10: // Opnum10NotUsedByProtocol
		// Opnum10NotUsedByProtocol
		return nil, nil
	case 11: // Opnum11NotUsedByProtocol
		// Opnum11NotUsedByProtocol
		return nil, nil
	case 12: // Opnum12NotUsedByProtocol
		// Opnum12NotUsedByProtocol
		return nil, nil
	case 13: // Opnum13NotUsedByProtocol
		// Opnum13NotUsedByProtocol
		return nil, nil
	case 14: // Opnum14NotUsedByProtocol
		// Opnum14NotUsedByProtocol
		return nil, nil
	case 15: // Opnum15NotUsedByProtocol
		// Opnum15NotUsedByProtocol
		return nil, nil
	case 16: // Opnum16NotUsedByProtocol
		// Opnum16NotUsedByProtocol
		return nil, nil
	case 17: // Opnum17NotUsedByProtocol
		// Opnum17NotUsedByProtocol
		return nil, nil
	case 18: // Opnum18NotUsedByProtocol
		// Opnum18NotUsedByProtocol
		return nil, nil
	case 19: // RemoteConnectionParms
		op := &xxx_RemoteConnectionParametersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoteConnectionParametersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoteConnectionParameters(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // RemoteUserSessionInfo
		op := &xxx_RemoteUserSessionInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoteUserSessionInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoteUserSessionInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // Opnum21NotUsedByProtocol
		// Opnum21NotUsedByProtocol
		return nil, nil
	}
	return nil, nil
}

// Unimplemented IPCHService
type UnimplementedPCHServiceServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedPCHServiceServer) RemoteConnectionParameters(context.Context, *RemoteConnectionParametersRequest) (*RemoteConnectionParametersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPCHServiceServer) RemoteUserSessionInfo(context.Context, *RemoteUserSessionInfoRequest) (*RemoteUserSessionInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ PCHServiceServer = (*UnimplementedPCHServiceServer)(nil)
