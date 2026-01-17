package imsmqapplication3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = imsmqapplication2.GoPackage
)

// IMSMQApplication3 server interface.
type Application3Server interface {

	// IMSMQApplication2 base class.
	imsmqapplication2.Application2Server

	// The ActiveQueues method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return an array of strings that contain the format names of all the
	// represented QueueManager.QueueCollection.Queues, where Queue.Active is equal to True.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	//
	// When the server processes this call, it MUST follow these guidelines:
	//
	// * If the ComputerName instance variable is NULL:
	//
	// * Identify all the local QueueManager.QueueCollection.Queues, where Queue.Active
	// is equal to True.
	//
	// * Set the pvActiveQueues output variable to an array that contains the format names
	// that specify all the identified Queues.
	//
	// * Else:
	//
	// * The server MUST generate a QMMgmt Get Info ( a0b1e28b-0f93-415d-8753-e3e1133678db
	// ) event with the following inputs:
	//
	// * iPropID = PROPID_MGMT_MSMQ_ACTIVEQUEUES
	//
	// * If the rStatus return value is not equal to MQ_OK (0x00000000), the server MUST
	// return rStatus and MUST take no further action. Otherwise, the pvActiveQueues output
	// variable MUST be set to the value of the returned rPropVar.
	ActiveQueues(context.Context, *ActiveQueuesRequest) (*ActiveQueuesResponse, error)

	// The PrivateQueues method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return an array of strings that contain the path names of all the
	// represented QueueManager.QueueCollection.Queues, where Queue.QueueType is equal to
	// Private.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	GetPrivateQueues(context.Context, *GetPrivateQueuesRequest) (*GetPrivateQueuesResponse, error)

	// The DirectoryServiceServer method is received by the server in an RPC_REQUEST packet.
	// In response, the server MUST return a string that contains the name of the current
	// directory computer.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	//
	// When the server processes this call, it MUST follow these guidelines:
	//
	// * If the ComputerName instance variable is NULL:
	//
	// * Set the pbstrDirectoryServiceServer output variable to the DNS or NetBIOS name
	// of the directory computer, <16> ( 71c359c3-e9ec-4fe6-a101-aab1eabecdcf#Appendix_A_16
	// ) prefixed by "\\".
	//
	// * Else:
	//
	// * The server MUST generate a QMMgmt Get Info event with the following inputs:
	//
	// * iPropID = PROPID_MGMT_MSMQ_DSSERVER
	//
	// * If the rStatus return value is not equal to MQ_OK (0x00000000), the server MUST
	// return rStatus and MUST take no further action. Otherwise, the pbstrDirectoryServiceServer
	// output variable MUST be set to the value of the returned rPropVar.
	GetDirectoryServiceServer(context.Context, *GetDirectoryServiceServerRequest) (*GetDirectoryServiceServerResponse, error)

	// The IsConnected method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return a BOOLEAN value that indicates the connection status of the
	// represented QueueManager.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	//
	// When the server processes this call, it MUST follow these guidelines:
	//
	// * If the ComputerName instance variable is NULL:
	//
	// * Set the pfIsConnected output variable to local QueueManager.ConnectionActive.
	//
	// * Else:
	//
	// * The server MUST generate a QMMgmt Get Info event with the following inputs:
	//
	// * iPropID = PROPID_MGMT_MSMQ_CONNECTED
	//
	// * If the rStatus return value is not equal to MQ_OK (0x00000000), the server MUST
	// return rStatus and MUST take no further action. Otherwise, the pfIsConnected output
	// variable MUST be set to the value of the returned rPropVar.
	GetIsConnected(context.Context, *GetIsConnectedRequest) (*GetIsConnectedResponse, error)

	// The BytesInAllQueues method is received by the server in an RPC_REQUEST packet. In
	// response, the server MUST return the number of message bytes that are currently stored
	// in all Queues of the represented QueueManager.QueueCollection.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	GetBytesInAllQueues(context.Context, *GetBytesInAllQueuesRequest) (*GetBytesInAllQueuesResponse, error)

	// Machine operation.
	SetMachine(context.Context, *SetMachineRequest) (*SetMachineResponse, error)

	// Machine operation.
	GetMachine(context.Context, *GetMachineRequest) (*GetMachineResponse, error)

	// The Connect method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST connect the represented QueueManager to the network and to the directory.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.<17>
	//
	// When the server processes this call, it MUST follow these guidelines:
	//
	// * If the ComputerName instance variable is NULL:
	//
	// * Send a Bring Online ( ../ms-mqdmpr/f8539502-ed84-4cdb-97e7-a8927c97fbbf ) event,
	// as defined in [MS-MQDMPR] ( ../ms-mqdmpr/5eafe0a6-a22f-436b-a0d9-4cbc25c52b47 ) section
	// 3.1.4.13, to the local QueueManager.
	//
	// * Else:
	//
	// * The server MUST generate a QMMgmt Action event with the following inputs:
	//
	// * iAction = "CONNECT"
	//
	// * The server MUST return rStatus , and MUST take no further action.
	Connect(context.Context, *ConnectRequest) (*ConnectResponse, error)

	// The Disconnect method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST disconnect the represented QueueManager from the network and the
	// directory.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.<18>
	//
	// When the server processes this call, it MUST follow these guidelines:
	//
	// * If the ComputerName instance variable is NULL:
	//
	// * Send a Take Offline ( ../ms-mqdmpr/3f4f55d3-aa90-41fb-985d-288fd76b2703 ) event
	// as defined in [MS-MQDMPR] ( ../ms-mqdmpr/5eafe0a6-a22f-436b-a0d9-4cbc25c52b47 ) section
	// 3.1.4.12 to the local QueueManager.
	//
	// * Else:
	//
	// * The server MUST generate a QMMgmt Action event with the following inputs:
	//
	// * iAction = "DISCONNECT"
	//
	// * The server MUST return rStatus , and MUST take no further action.
	Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error)

	// The Tidy method is received by the server in an RPC_REQUEST packet. In response,
	// the server SHOULD perform implementation-specific tasks to release unused resources
	// of the represented QueueManager.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.<19>
	//
	// When the server processes this call, it MUST follow these guidelines:
	//
	// * If the ComputerName instance variable is NULL:
	//
	// * The local QueueManager SHOULD perform implementation-specific tasks to release
	// unused resources.
	//
	// * Else:
	//
	// * The server MUST generate a QMMgmt Action event with the following inputs:
	//
	// * iAction = "TIDY"
	//
	// * The server MUST return rStatus , and MUST take no further action.
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
