package faxclient

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

// faxclient server interface.
type FaxclientServer interface {

	// The FAX_OpenConnection (Opnum 0) method returns the context handle that is supplied
	// by the FAX_StartServerNotification family of calls. This is done to provide a security
	// layer, by verifying that the notifications are coming from an expected source.
	//
	// Exceptions Thrown: No exceptions are thrown except those thrown by the underlying
	// RPC protocol, [MS-RPCE].
	//
	// The returned PRPC_FAX_HANDLE is the Context parameter cast to a HANDLE.
	//
	// The FAX_OpenConnection method returns the context handle supplied by the FAX_StartServerNotification
	// family of calls. This is done to provide a security layer, by verifying that the
	// notifications are coming from an expected source.
	OpenConnection(context.Context, *OpenConnectionRequest) (*OpenConnectionResponse, error)

	// The FAX_ClientEventQueue (Opnum 1) method is called by the fax server (acting as
	// an RPC client for this call) when it needs to deliver a legacy fax event to the fax
	// client (acting as an RPC server for this call). The fax client registers for notifications
	// of legacy events with the fax server by calling FAX_StartServerNotification (section
	// 3.1.4.1.100) or FAX_StartServerNotificationEx (section 3.1.4.1.101). In this call,
	// the fax client MUST pass a fax client notification context, which the fax server
	// MUST pass back to the fax client when it sends an event. This is done to provide
	// a security layer, by verifying that the notifications are coming from an expected
	// source.
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	ClientEventQueue(context.Context, *ClientEventQueueRequest) (*ClientEventQueueResponse, error)

	// The FAX_CloseConnection (Opnum 2) method is called by the fax server (acting as an
	// RPC client for this call) when it needs to release the connection to the fax client
	// (acting as an RPC server for this call). When the fax client calls FAX_EndServerNotification
	// (section 3.1.4.1.17), the fax server MUST release the RPC connection to the fax client
	// through this call.
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	CloseConnection(context.Context, *CloseConnectionRequest) (*CloseConnectionResponse, error)

	// The FAX_ClientEventQueueEx (Opnum 3) method is called by the fax server (acting as
	// an RPC client for this call) when it needs to deliver an extended fax event to the
	// fax client (acting as an RPC server for this call). The fax client registers for
	// notifications with the fax server by calling either FAX_StartServerNotificationEx
	// (section 3.1.4.1.101) or FAX_StartServerNotificationEx2 (section 3.1.4.1.102). In
	// this call, the fax client MUST pass a fax client notification context, which the
	// fax server MUST pass back to the fax client when it sends an event. This is done
	// to provide a security layer, by verifying that the notifications are coming from
	// an expected source.
	//
	// Data in FAX_ClientEventQueueEx is serialized. Pointers to variable-size data (such
	// as strings) are replaced with offsets from the beginning of the buffer.
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	//
	// Data in FAX_ClientEventQueueEx is serialized. Pointers to variable size data (such
	// as strings) are replaced with offsets from the beginning of the buffer.
	ClientEventQueueEx(context.Context, *ClientEventQueueExRequest) (*ClientEventQueueExResponse, error)
}

func RegisterFaxclientServer(conn dcerpc.Conn, o FaxclientServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewFaxclientServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(FaxclientSyntaxV3_0))...)
}

func NewFaxclientServerHandle(o FaxclientServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return FaxclientServerHandle(ctx, o, opNum, r)
	}
}

func FaxclientServerHandle(ctx context.Context, o FaxclientServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // FAX_OpenConnection
		op := &xxx_OpenConnectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenConnectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenConnection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // FAX_ClientEventQueue
		op := &xxx_ClientEventQueueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ClientEventQueueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ClientEventQueue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // FAX_CloseConnection
		op := &xxx_CloseConnectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloseConnectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CloseConnection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // FAX_ClientEventQueueEx
		op := &xxx_ClientEventQueueExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ClientEventQueueExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ClientEventQueueEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented faxclient
type UnimplementedFaxclientServer struct {
}

func (UnimplementedFaxclientServer) OpenConnection(context.Context, *OpenConnectionRequest) (*OpenConnectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxclientServer) ClientEventQueue(context.Context, *ClientEventQueueRequest) (*ClientEventQueueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxclientServer) CloseConnection(context.Context, *CloseConnectionRequest) (*CloseConnectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxclientServer) ClientEventQueueEx(context.Context, *ClientEventQueueExRequest) (*ClientEventQueueExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ FaxclientServer = (*UnimplementedFaxclientServer)(nil)
