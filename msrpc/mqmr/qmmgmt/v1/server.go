package qmmgmt

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

// qmmgmt server interface.
type QmmgmtServer interface {

	// The R_QMMgmtGetInfo method requests information on an MSMQ installation on a server
	// or on a specific queue.
	//
	// Return Values: On success, this method MUST return MQ_OK (0x00000000).
	//
	//	+---------------------+----------------------------------------------------------------------------------+
	//	|       RETURN        |                                                                                  |
	//	|     VALUE/CODE      |                                   DESCRIPTION                                    |
	//	|                     |                                                                                  |
	//	+---------------------+----------------------------------------------------------------------------------+
	//	+---------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 MQ_OK    |                                                                                  |
	//	+---------------------+----------------------------------------------------------------------------------+
	//	| 0xC00E0001 MQ_ERROR | Generic error code. This error code is also the first of several error codes     |
	//	|                     | beginning with the string "MQ_ERR". A list of the errors prefaced with "MQ-ERR"  |
	//	|                     | is specified in 2.4.                                                             |
	//	+---------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// If an error occurs, the server MUST return a failure HRESULT and MUST NOT set any
	// [out] parameter values.
	//
	// The opnum field value for this method MUST be 0 and is received at a dynamically
	// assigned endpoint supplied by the RPC endpoint mapper, as specified in [MS-RPCE].
	//
	// If the pObjectFormat parameter specifies an MgmtObjectType of MGMT_MACHINE, the server
	// MUST return only those properties that pertain to the MSMQ installation. If pObjectFormat
	// specifies an MgmtObjectType of MGMT_QUEUE, the server MUST return only those properties
	// that pertain to a queue. If pObjectFormat specifies an MgmtObjectType of MGMT_SESSION,
	// the call MUST fail, and the error message MAY be MQ_ERROR_INVALID_PARAMETER (0xC00E0006).<15>
	//
	// If the pObjectFormat parameter specifies a computer, and one or more of the properties
	// specified in aProp are different than those specified in section 2.2.3.1, the call
	// MAY fail with MQ_ERROR_ILLEGAL_PROPID (0xC00E0039). If the pObjectFormat parameter
	// specifies a queue, and one or more of the properties specified in aProp are different
	// than those specified in section 2.2.3.2, the call MAY fail with MQ_ERROR_ILLEGAL_PROPID
	// (0xC00E0039).<16>
	//
	// MSMQ properties are specified in [MS-MQMQ] section 2.
	//
	// For MSMQ error codes, see [MSDN-MQEIC]. The structure and sequence of data on the
	// wire are specified in [C706] Transfer Syntax NDR.
	ManagementGetInfo(context.Context, *ManagementGetInfoRequest) (*ManagementGetInfoResponse, error)

	// The R_QMMgmtAction method requests the server to perform a management function on
	// a specific queue or MSMQ installation.
	//
	// Return Values: On success, this method MUST return MQ_OK (0x00000000).
	//
	// MQ_OK (0x00000000)
	//
	// MQ_ERROR (0xC00E0001)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// If pObjectFormat specifies an MgmtObjectType of MGMT_SESSION or an lpwszAction has
	// different value than those in the table above, the call MUST fail and the error message
	// MAY be MQ_ERROR_INVALID_PARAMETER (0xC00E0006).<17>
	//
	// If an error occurs, the server MUST return a failure HRESULT.
	//
	// The opnum field value for this method MUST be 1 and is received at a dynamically
	// assigned endpoint supplied by the RPC endpoint mapper, as specified in [MS-RPCE].
	//
	// For MSMQ error codes, see [MSDN-MQEIC]. The structure and sequence of data on the
	// wire are specified in the Transfer Syntax NDR section in [C706].
	ManagementAction(context.Context, *ManagementActionRequest) (*ManagementActionResponse, error)
}

func RegisterQmmgmtServer(conn dcerpc.Conn, o QmmgmtServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewQmmgmtServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(QmmgmtSyntaxV1_0))...)
}

func NewQmmgmtServerHandle(o QmmgmtServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return QmmgmtServerHandle(ctx, o, opNum, r)
	}
}

func QmmgmtServerHandle(ctx context.Context, o QmmgmtServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // R_QMMgmtGetInfo
		op := &xxx_ManagementGetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ManagementGetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ManagementGetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // R_QMMgmtAction
		op := &xxx_ManagementActionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ManagementActionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ManagementAction(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented qmmgmt
type UnimplementedQmmgmtServer struct {
}

func (UnimplementedQmmgmtServer) ManagementGetInfo(context.Context, *ManagementGetInfoRequest) (*ManagementGetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQmmgmtServer) ManagementAction(context.Context, *ManagementActionRequest) (*ManagementActionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ QmmgmtServer = (*UnimplementedQmmgmtServer)(nil)
