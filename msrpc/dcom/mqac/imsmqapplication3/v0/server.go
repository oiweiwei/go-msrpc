package imsmqapplication3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	imsmqapplication2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqapplication2/v0"
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
	_ = imsmqapplication2.GoPackage
)

// IMSMQApplication3 server interface.
type Application3Server interface {

	// IMSMQApplication2 base class.
	imsmqapplication2.Application2Server

	// ActiveQueues operation.
	ActiveQueues(context.Context, *ActiveQueuesRequest) (*ActiveQueuesResponse, error)

	// PrivateQueues operation.
	GetPrivateQueues(context.Context, *GetPrivateQueuesRequest) (*GetPrivateQueuesResponse, error)

	// DirectoryServiceServer operation.
	GetDirectoryServiceServer(context.Context, *GetDirectoryServiceServerRequest) (*GetDirectoryServiceServerResponse, error)

	// IsConnected operation.
	GetIsConnected(context.Context, *GetIsConnectedRequest) (*GetIsConnectedResponse, error)

	// BytesInAllQueues operation.
	GetBytesInAllQueues(context.Context, *GetBytesInAllQueuesRequest) (*GetBytesInAllQueuesResponse, error)

	// Machine operation.
	SetMachine(context.Context, *SetMachineRequest) (*SetMachineResponse, error)

	// Machine operation.
	GetMachine(context.Context, *GetMachineRequest) (*GetMachineResponse, error)

	// Connect operation.
	Connect(context.Context, *ConnectRequest) (*ConnectResponse, error)

	// Disconnect operation.
	Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error)

	// Tidy operation.
	Tidy(context.Context, *TidyRequest) (*TidyResponse, error)
}

func RegisterApplication3Server(conn dcerpc.Conn, o Application3Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewApplication3ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Application3SyntaxV0_0))...)
}

func NewApplication3ServerHandle(o Application3Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Application3ServerHandle(ctx, o, opNum, r)
	}
}

func Application3ServerHandle(ctx context.Context, o Application3Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 15 {
		// IMSMQApplication2 base method.
		return imsmqapplication2.Application2ServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 15: // ActiveQueues
		op := &xxx_ActiveQueuesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ActiveQueuesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ActiveQueues(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // PrivateQueues
		op := &xxx_GetPrivateQueuesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPrivateQueuesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPrivateQueues(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // DirectoryServiceServer
		op := &xxx_GetDirectoryServiceServerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDirectoryServiceServerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDirectoryServiceServer(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // IsConnected
		op := &xxx_GetIsConnectedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsConnectedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsConnected(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // BytesInAllQueues
		op := &xxx_GetBytesInAllQueuesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetBytesInAllQueuesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetBytesInAllQueues(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // Machine
		op := &xxx_SetMachineOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMachineRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMachine(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // Machine
		op := &xxx_GetMachineOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMachineRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMachine(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // Connect
		op := &xxx_ConnectOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ConnectRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Connect(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // Disconnect
		op := &xxx_DisconnectOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DisconnectRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Disconnect(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // Tidy
		op := &xxx_TidyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &TidyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Tidy(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMSMQApplication3
type UnimplementedApplication3Server struct {
	imsmqapplication2.UnimplementedApplication2Server
}

func (UnimplementedApplication3Server) ActiveQueues(context.Context, *ActiveQueuesRequest) (*ActiveQueuesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedApplication3Server) GetPrivateQueues(context.Context, *GetPrivateQueuesRequest) (*GetPrivateQueuesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedApplication3Server) GetDirectoryServiceServer(context.Context, *GetDirectoryServiceServerRequest) (*GetDirectoryServiceServerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedApplication3Server) GetIsConnected(context.Context, *GetIsConnectedRequest) (*GetIsConnectedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedApplication3Server) GetBytesInAllQueues(context.Context, *GetBytesInAllQueuesRequest) (*GetBytesInAllQueuesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedApplication3Server) SetMachine(context.Context, *SetMachineRequest) (*SetMachineResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedApplication3Server) GetMachine(context.Context, *GetMachineRequest) (*GetMachineResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedApplication3Server) Connect(context.Context, *ConnectRequest) (*ConnectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedApplication3Server) Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedApplication3Server) Tidy(context.Context, *TidyRequest) (*TidyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ Application3Server = (*UnimplementedApplication3Server)(nil)
