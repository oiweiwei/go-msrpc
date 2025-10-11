package imanagetelnetsessions

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
	_ = idispatch.GoPackage
)

// IManageTelnetSessions server interface.
type ManageTelnetSessionsServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The GetTelnetSessions method is used to query the telnet server for information about
	// all active telnet sessions.
	//
	// Return Values: The server MUST return zero if the method is successful. The server
	// MUST return 0x01 if processing fails and set output parameters to NULL. These are
	// in addition to the values that can be returned by the underlying [MS-DCOM] implementation.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// DCOM protocol [MS-DCOM].
	GetTelnetSessions(context.Context, *GetTelnetSessionsRequest) (*GetTelnetSessionsResponse, error)

	// The TerminateSession method terminates a telnet session.
	//
	// Return Values: The server MUST return zero if the method is successful. The server
	// MUST return 0x01 if processing fails. These are in addition to the values that can
	// be returned by the underlying [MS-DCOM] implementation.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// DCOM protocol [MS-DCOM].
	TerminateSession(context.Context, *TerminateSessionRequest) (*TerminateSessionResponse, error)

	// The SendMsgToASession method directs the telnet server to send a text message to
	// the telnet client that initiated the session.
	//
	// Return Values: The server MUST return zero if the method is successful. The server
	// MUST return 0x01 if processing fails. These are in addition to the values that can
	// be returned by the underlying [MS-DCOM] implementation.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// DCOM protocol [MS-DCOM].
	SendMessageToASession(context.Context, *SendMessageToASessionRequest) (*SendMessageToASessionResponse, error)
}

func RegisterManageTelnetSessionsServer(conn dcerpc.Conn, o ManageTelnetSessionsServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewManageTelnetSessionsServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ManageTelnetSessionsSyntaxV0_0))...)
}

func NewManageTelnetSessionsServerHandle(o ManageTelnetSessionsServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ManageTelnetSessionsServerHandle(ctx, o, opNum, r)
	}
}

func ManageTelnetSessionsServerHandle(ctx context.Context, o ManageTelnetSessionsServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // GetTelnetSessions
		op := &xxx_GetTelnetSessionsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTelnetSessionsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTelnetSessions(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // TerminateSession
		op := &xxx_TerminateSessionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &TerminateSessionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.TerminateSession(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // SendMsgToASession
		op := &xxx_SendMessageToASessionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SendMessageToASessionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SendMessageToASession(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IManageTelnetSessions
type UnimplementedManageTelnetSessionsServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedManageTelnetSessionsServer) GetTelnetSessions(context.Context, *GetTelnetSessionsRequest) (*GetTelnetSessionsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedManageTelnetSessionsServer) TerminateSession(context.Context, *TerminateSessionRequest) (*TerminateSessionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedManageTelnetSessionsServer) SendMessageToASession(context.Context, *SendMessageToASessionRequest) (*SendMessageToASessionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ManageTelnetSessionsServer = (*UnimplementedManageTelnetSessionsServer)(nil)
