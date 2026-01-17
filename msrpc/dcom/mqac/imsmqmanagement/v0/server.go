package imsmqmanagement

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

// IMSMQManagement server interface.
type ManagementServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The Init method is received by the server in an RPC_REQUEST packet. In response,
	// the server initializes the object to represent the state of a Queue. The represented
	// Queue MUST have Queue.Active set to True. If the represented Queue.Active is False,
	// or the Pathname and FormatName input parameters cannot be resolved, this method MUST
	// return an error without setting the ObjectIsInitialized instance variable to True.
	// This method MUST be called prior to calling any other operation on MSMQManagement.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	Init(context.Context, *InitRequest) (*InitResponse, error)

	// FormatName operation.
	GetFormatName(context.Context, *GetFormatNameRequest) (*GetFormatNameResponse, error)

	// Machine operation.
	GetMachine(context.Context, *GetMachineRequest) (*GetMachineResponse, error)

	// The MessageCount method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return the number of messages in the represented Queue.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	GetMessageCount(context.Context, *GetMessageCountRequest) (*GetMessageCountResponse, error)

	// The ForeignStatus method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return an enumerated value to indicate whether a Queue is a foreign
	// queue or an OutgoingQueue that transfers messages to a foreign queue.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	//
	// When the server processes this call, it MUST follow these guidelines:
	//
	// * If the *ObjectIsInitialized* instance variable is False:
	//
	// * The server MUST return MQ_ERROR_UNINITIALIZED_OBJECT (0xC00E0094) and MUST take
	// no further action.
	//
	// * The server MUST generate a *QMMgmt Get Info* event with the following inputs:
	//
	// * iPropID = PROPID_MGMT_QUEUE_FOREIGN
	//
	// * If the *rStatus* return value is not equal to MQ_OK (0x00000000), the server MUST
	// return *rStatus* and MUST take no further action.
	//
	// * Else:
	//
	// * If the value of the returned *rPropVar* was "Yes":
	//
	// * The *pIForeignStatus* output variable MUST be set to MQ_STATUS_FOREIGN, and the
	// server MUST take no further action.
	//
	// * If the value of the returned *rPropVar* was "No":
	//
	// * The *pIForeignStatus* output variable MUST be set to MQ_STATUS_NOT_FOREIGN, and
	// the server MUST take no further action.
	//
	// * If the value of the returned *rPropVar* was "Unknown":
	//
	// * The *pIForeignStatus* output variable MUST be set to MQ_STATUS_UNKNOWN, and the
	// server MUST take no further action.
	GetForeignStatus(context.Context, *GetForeignStatusRequest) (*GetForeignStatusResponse, error)

	// The QueueType method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return the type of the referenced Queue.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	//
	// When the server processes this call, it MUST follow these guidelines:
	//
	// * If the *ObjectIsInitialized* instance variable is False:
	//
	// * The server MUST return MQ_ERROR_UNINITIALIZED_OBJECT (0xC00E0094), and MUST take
	// no further action.
	//
	// * Else:
	//
	// * The server MUST generate a *QMMgmt Get Info* event with the following inputs:
	//
	// * iPropID = PROPID_MGMT_QUEUE_TYPE
	//
	// * If the rStatus return value is not equal to MQ_OK (0x00000000), the server MUST
	// return rStatus and MUST take no further action.
	//
	// * If the value of the returned *rPropVar* was "PUBLIC":
	//
	// * The plQueueType output variable MUST be set to MQ_TYPE_PUBLIC, and the server MUST
	// return S_OK (0x00000000) and MUST take no further action.
	//
	// * If the value of the returned was "PRIVATE":
	//
	// * The plQueueType output variable MUST be set to MQ_TYPE_PRIVATE, and the server
	// MUST return S_OK (0x00000000) and MUST take no further action.
	//
	// * If the value of the returned was "MACHINE":
	//
	// * The plQueueType output variable MUST be set to MQ_TYPE_MACHINE, and the server
	// MUST return S_OK (0x00000000) and MUST take no further action.
	//
	// * If the value of the returned was "CONNECTOR":
	//
	// * The plQueueType output variable MUST be set to MGMT_TYPE_CONNECTOR, and the server
	// MUST return S_OK (0x00000000) and MUST take no further action.
	//
	// * If the value of the returned was "MULTICAST":
	//
	// * The plQueueType output variable MUST be set to MQ_TYPE_MULTICAST, and the server
	// MUST return S_OK (0x00000000) and MUST take no further action.
	GetQueueType(context.Context, *GetQueueTypeRequest) (*GetQueueTypeResponse, error)

	// The IsLocal method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return a BOOLEAN value that indicates whether the represented Queue
	// is an OutgoingQueue (False) or not (True).
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	//
	// When the server processes this call, it MUST follow these guidelines:
	//
	// * If the *ObjectIsInitialized* instance variable is False:
	//
	// * The server MUST return MQ_ERROR_UNINITIALIZED_OBJECT (0xC00E0094) and MUST take
	// no further action.
	//
	// * The server MUST generate a QMMgmt Get Info ( a0b1e28b-0f93-415d-8753-e3e1133678db
	// ) event with the following inputs:
	//
	// * *iPropID* = PROPID_MGMT_QUEUE_LOCATION
	//
	// * If the *rStatus* return value is not equal to MQ_OK (0x00000000), the server MUST
	// return *rStatus* and MUST take no further action.
	//
	// * Else:
	//
	// * If the value of the returned *rPropVar* was "LOCAL":
	//
	// * The pfIsLocal output variable MUST be set to True, and the server MUST take no
	// further action.
	//
	// * If the value of the returned *rPropVar* was "REMOTE":
	//
	// * The pfIsLocal output variable MUST be set to False, and the server MUST take no
	// further action.
	//
	// * If the value of the returned *rPropVar* was "LOCAL":
	//
	// * The pfIsLocal output variable MUST be set to True, and the server MUST take no
	// further action.
	GetIsLocal(context.Context, *GetIsLocalRequest) (*GetIsLocalResponse, error)

	// The TransactionalStatus method is received by the server in an RPC_REQUEST packet.
	// In response, the server MUST return the represented Queue.Transactional value.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	GetTransactionalStatus(context.Context, *GetTransactionalStatusRequest) (*GetTransactionalStatusResponse, error)

	// The BytesInQueue method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return the value of the represented Queue.TotalBytes property.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	//
	// When the server processes this call, it MUST follow these guidelines:
	//
	// * If the *ObjectIsInitialized* instance variable is False:
	//
	// * The server MUST return MQ_ERROR_UNINITIALIZED_OBJECT (0xC00E0094), and MUST take
	// no further action.
	//
	// * The server MUST generate a *QMMgmt Get Info* event with the following inputs:
	//
	// * *iPropID* = PROPID_MGMT_QUEUE_BYTES_IN_ QUEUE
	//
	// * If the *rStatus* return value is not equal to MQ_OK (0x00000000), the server MUST
	// return *rStatus* and MUST take no further action.
	//
	// * Else:
	//
	// * The pvBytesInQueue output variable MUST be set to the value of the returned *rPropVar*.
	GetBytesInQueue(context.Context, *GetBytesInQueueRequest) (*GetBytesInQueueResponse, error)
}

func RegisterManagementServer(conn dcerpc.Conn, o ManagementServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewManagementServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ManagementSyntaxV0_0))...)
}

func NewManagementServerHandle(o ManagementServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ManagementServerHandle(ctx, o, opNum, r)
	}
}

func ManagementServerHandle(ctx context.Context, o ManagementServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Init
		op := &xxx_InitOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &InitRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Init(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // FormatName
		op := &xxx_GetFormatNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFormatNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFormatName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // Machine
		op := &xxx_GetMachineOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMachineRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMachine(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // MessageCount
		op := &xxx_GetMessageCountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMessageCountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMessageCount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // ForeignStatus
		op := &xxx_GetForeignStatusOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetForeignStatusRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetForeignStatus(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // QueueType
		op := &xxx_GetQueueTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetQueueTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetQueueType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // IsLocal
		op := &xxx_GetIsLocalOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsLocalRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsLocal(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // TransactionalStatus
		op := &xxx_GetTransactionalStatusOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTransactionalStatusRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTransactionalStatus(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // BytesInQueue
		op := &xxx_GetBytesInQueueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetBytesInQueueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetBytesInQueue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMSMQManagement
type UnimplementedManagementServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedManagementServer) Init(context.Context, *InitRequest) (*InitResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedManagementServer) GetFormatName(context.Context, *GetFormatNameRequest) (*GetFormatNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedManagementServer) GetMachine(context.Context, *GetMachineRequest) (*GetMachineResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedManagementServer) GetMessageCount(context.Context, *GetMessageCountRequest) (*GetMessageCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedManagementServer) GetForeignStatus(context.Context, *GetForeignStatusRequest) (*GetForeignStatusResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedManagementServer) GetQueueType(context.Context, *GetQueueTypeRequest) (*GetQueueTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedManagementServer) GetIsLocal(context.Context, *GetIsLocalRequest) (*GetIsLocalResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedManagementServer) GetTransactionalStatus(context.Context, *GetTransactionalStatusRequest) (*GetTransactionalStatusResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedManagementServer) GetBytesInQueue(context.Context, *GetBytesInQueueRequest) (*GetBytesInQueueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ManagementServer = (*UnimplementedManagementServer)(nil)
