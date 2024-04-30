package qmcomm2

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

// qmcomm2 server interface.
type Qmcomm2Server interface {

	// A client invokes QMSendMessageInternalEx if the server returns STATUS_RETRY (0xc000022d)
	// from a prior call to rpc_ACSendMessageEx. Implementations of this protocol SHOULD
	// NOT return STATUS_RETRY from rpc_ACSendMessageEx, rendering this method unnecessary.
	// Such implementations MUST take no action when QMSendMessageInternalEx is invoked
	// and return MQ_ERROR_ILLEGAL_OPERATION (0xc00e0064).
	//
	// Return Values:  On success, this method MUST return MQ_OK (0x00000000); otherwise,
	// the server MUST return a failure HRESULT,<67> and the client MUST treat all failure
	// HRESULTs identically. Additionally, if a failure HRESULT is returned, the client
	// MUST disregard all out-parameter values.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// This method is invoked at the dynamically assigned endpoint returned by the R_QMGetRTQMServerPort
	// method when IP_HANDSHAKE (0x00000000) or IPX_HANDSHAKE (0x00000002) is the interface
	// specified by the fIP parameter.
	QMSendMessageInternalEx(context.Context, *QMSendMessageInternalExRequest) (*QMSendMessageInternalExResponse, error)

	// A client calls the rpc_ACSendMessageEx method to place a message into a message queue
	// for delivery.
	//
	// Return Values: On success, this method MUST return MQ_OK (0x00000000); otherwise,
	// the server MUST return a failure HRESULT,<72> and the client MUST treat all failure
	// HRESULTs identically. Additionally, if a failure HRESULT is returned, the client
	// MUST disregard all out-parameter values.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// This method is invoked at the dynamically assigned endpoint returned by the R_QMGetRTQMServerPort
	// (section 3.1.4.24) method when IP_HANDSHAKE (0x00000000) or IPX_HANDSHAKE (0x00000002)
	// is the interface specified by the fIP parameter.
	//
	// Security Considerations: The caller can request that the server perform security
	// related operations such as signing and encrypting the message. These operations are
	// requested by setting members of the ptb.CACTransferBufferOld structure.
	SendMessageEx(context.Context, *SendMessageExRequest) (*SendMessageExResponse, error)

	// A client calls rpc_ACReceiveMessageEx to peek or receive a message from a message
	// queue.
	//
	// Return Values:  On success, this method MUST return MQ_OK (0x00000000); otherwise,
	// the server MUST return a failure HRESULT,<73><74> and the client MUST treat all failure
	// HRESULTs identically. Additionally, if a failure HRESULT is returned, the client
	// MUST disregard all out-parameter values.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// This method is invoked at the dynamically assigned endpoint returned by the R_QMGetRTQMServerPort
	// method when IP_HANDSHAKE (0x00000000) or IPX_HANDSHAKE (0x00000002) is the interface
	// specified by the fIP parameter.
	ReceiveMessageEx(context.Context, *ReceiveMessageExRequest) (*ReceiveMessageExResponse, error)

	// A client calls rpc_ACCreateCursorEx to create a cursor for use when peeking and receiving
	// from a message queue.
	//
	// Return Values:  On success, this method MUST return MQ_OK (0x00000000); otherwise,
	// the server MUST return a failure HRESULT,<75> and the client MUST treat all failure
	// HRESULTs identically. Additionally, if a failure HRESULT is returned, the client
	// MUST disregard all out-parameter values.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// This method is invoked at the dynamically assigned endpoint returned by the R_QMGetRTQMServerPort
	// (section 3.1.4.24) method when IP_HANDSHAKE (0x00000000) or IPX_HANDSHAKE (0x00000002)
	// is the interface specified by the fIP parameter.
	CreateCursorEx(context.Context, *CreateCursorExRequest) (*CreateCursorExResponse, error)
}

func RegisterQmcomm2Server(conn dcerpc.Conn, o Qmcomm2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewQmcomm2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Qmcomm2SyntaxV1_0))...)
}

func NewQmcomm2ServerHandle(o Qmcomm2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Qmcomm2ServerHandle(ctx, o, opNum, r)
	}
}

func Qmcomm2ServerHandle(ctx context.Context, o Qmcomm2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // QMSendMessageInternalEx
		in := &QMSendMessageInternalExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QMSendMessageInternalEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 1: // rpc_ACSendMessageEx
		in := &SendMessageExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SendMessageEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 2: // rpc_ACReceiveMessageEx
		in := &ReceiveMessageExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ReceiveMessageEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 3: // rpc_ACCreateCursorEx
		in := &CreateCursorExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateCursorEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
