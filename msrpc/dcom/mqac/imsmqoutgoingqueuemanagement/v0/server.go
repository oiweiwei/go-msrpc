package imsmqoutgoingqueuemanagement

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	imsmqmanagement "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqmanagement/v0"
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
	_ = imsmqmanagement.GoPackage
)

// IMSMQOutgoingQueueManagement server interface.
type OutgoingQueueManagementServer interface {

	// IMSMQManagement base class.
	imsmqmanagement.ManagementServer

	// The State method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return the represented OutgoingQueue.State.
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
	// * *iPropID* = PROPID_MGMT_QUEUE_STATE
	//
	// * If the *rStatus* return value is not equal to MQ_OK (0x00000000), the server MUST
	// return *rStatus* and MUST take no further action.
	//
	// * Else:
	//
	// * If the value of the returned *rPropVar* was "LOCAL CONNECTION":
	//
	// * The server MUST set the plState output variable to MQ_QUEUE_STATE_LOCAL_CONNECTION,
	// and MUST take no further action.
	//
	// * If the value of the returned *rPropVar* was "DISCONNECTED":
	//
	// * The server MUST set the plState output variable to MQ_QUEUE_STATE_DISCONNECTED,
	// and MUST take no further action.
	//
	// * If the value of the returned *rPropVar* was "LOCKED":
	//
	// * The server MUST set the plState output variable to MQ_QUEUE_STATE_LOCKED, and MUST
	// take no further action.
	//
	// * If the value of the returned *rPropVar* was "WAITING":
	//
	// * The server MUST set the plState output variable to MQ_QUEUE_STATE_WAITING, and
	// MUST take no further action.
	//
	// * If the value of the returned *rPropVar* was "NEED VALIDATION":
	//
	// * The server MUST set the plState output variable to MQ_QUEUE_STATE_NEEDVALIDATE,
	// and MUST take no further action.
	//
	// * If the value of the returned *rPropVar* was "ONHOLD":
	//
	// * The server MUST set the plState output variable to MQ_QUEUE_STATE_ONHOLD, and MUST
	// take no further action.
	//
	// * If the value of the returned *rPropVar* was "INACTIVE":
	//
	// * The server MUST set the plState output variable to MQ_QUEUE_STATE_NONACTIVE, and
	// MUST take no further action.
	//
	// * If the value of the returned *rPropVar* was "CONNECTED":
	//
	// * The server MUST set the plState output variable to MQ_QUEUE_STATE_CONNECTED, and
	// MUST take no further action.
	//
	// * If the value of the returned *rPropVar* was "DISCONNECTING":
	//
	// * The server MUST set the plState output variable to MQ_QUEUE_STATE_DISCONNECTING,
	// and MUST take no further action.
	GetState(context.Context, *GetStateRequest) (*GetStateResponse, error)

	// The NextHops method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return the represented OutgoingQueue.NextHops.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	GetNextHops(context.Context, *GetNextHopsRequest) (*GetNextHopsResponse, error)

	// The EodGetSendInfo method is received by the server in an RPC_REQUEST packet. In
	// response, the server MUST return the represented OutgoingQueue.OutgoingTransferInfoReference.
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
	// * The server MUST create a local variable tempCollection and initialize it as an
	// empty collection.
	//
	// * The server MUST generate a QMMgmt Get Info event with the following input:
	//
	// * PROPID_MGMT_QUEUE_EOD_LAST_ACK
	//
	// * If the rStatus return value is not equal to MQ_OK (0x00000000), the server MUST
	// return rStatus and MUST take no further action.
	//
	// * The value of the returned rPropVa MUST be added to tempCollection.
	//
	// * The server MUST generate a QMMgmt Get Info event with the following input:
	//
	// * PROPID_MGMT_QUEUE_EOD_LAST_ACK_TIME
	//
	// * If the rStatus return value is not equal to MQ_OK (0x00000000), the server MUST
	// return rStatus and MUST take no further action.
	//
	// * The value of the returned rPropVa MUST be added to tempCollection.
	//
	// * The server MUST generate a QMMgmt Get Info event with the following input:
	//
	// * PROPID_MGMT_QUEUE_EOD_LAST_ACK_COUNT
	//
	// * If the rStatus return value is not equal to MQ_OK (0x00000000), the server MUST
	// return rStatus and MUST take no further action.
	//
	// * The value of the returned rPropVa MUST be added to tempCollection.
	//
	// * The server MUST generate a QMMgmt Get Info event with the following input:
	//
	// * PROPID_MGMT_QUEUE_EOD_FIRST_NON_ACK
	//
	// * If the rStatus return value is not equal to MQ_OK (0x00000000), the server MUST
	// return rStatus and MUST take no further action.
	//
	// * The value of the returned rPropVa MUST be added to tempCollection.
	//
	// * The server MUST generate a QMMgmt Get Info event with the following input:
	//
	// * PROPID_MGMT_QUEUE_EOD_LAST_NON_ACK
	//
	// * If the rStatus return value is not equal to MQ_OK (0x00000000), the server MUST
	// return rStatus and MUST take no further action.
	//
	// * The value of the returned rPropVa MUST be added to tempCollection.
	//
	// * The server MUST generate a QMMgmt Get Info event with the following input:
	//
	// * PROPID_MGMT_QUEUE_EOD_NEXT_SEQ
	//
	// * If the rStatus return value is not equal to MQ_OK (0x00000000), the server MUST
	// return rStatus and MUST take no further action.
	//
	// * The value of the returned rPropVa MUST be added to tempCollection.
	//
	// * The server MUST generate a QMMgmt Get Info event with the following input:
	//
	// * PROPID_MGMT_QUEUE_EOD_NO_READ_COUNT
	//
	// * If the rStatus return value is not equal to MQ_OK (0x00000000), the server MUST
	// return rStatus and MUST take no further action.
	//
	// * The value of the returned rPropVa MUST be added to tempCollection.
	//
	// * The server MUST generate a QMMgmt Get Info event with the following input:
	//
	// * PROPID_MGMT_QUEUE_EOD_NO_ACK_COUNT
	//
	// * If the rStatus return value is not equal to MQ_OK (0x00000000), the server MUST
	// return rStatus and MUST take no further action.
	//
	// * The value of the returned rPropVa MUST be added to tempCollection.
	//
	// * The server MUST generate a QMMgmt Get Info event with the following input:
	//
	// * PROPID_MGMT_QUEUE_EOD_RESEND_TIME
	//
	// * If the rStatus return value is not equal to MQ_OK (0x00000000), the server MUST
	// return rStatus and MUST take no further action.
	//
	// * The value of the returned rPropVa MUST be added to tempCollection.
	//
	// * The server MUST generate a QMMgmt Get Info event with the following input:
	//
	// * PROPID_MGMT_QUEUE_EOD_RESEND_INTERVAL
	//
	// * If the rStatus return value is not equal to MQ_OK (0x00000000), the server MUST
	// return rStatus and MUST take no further action.
	//
	// * The value of the returned rPropVa MUST be added to tempCollection.
	//
	// * The server MUST generate a QMMgmt Get Info event with the following input:
	//
	// * PROPID_MGMT_QUEUE_EOD_RESEND_COUNT
	//
	// * If the rStatus return value is not equal to MQ_OK (0x00000000), the server MUST
	// return rStatus and MUST take no further action.
	//
	// * The value of the returned rPropVa MUST be added to tempCollection.
	//
	// * The server MUST copy tempCollection to the ppCollection output variable and return
	// S_OK (0x00000000).
	EODGetSendInfo(context.Context, *EODGetSendInfoRequest) (*EODGetSendInfoResponse, error)

	// The Resume method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST resume the transfer of messages from the represented outgoing Queue.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.<20>
	//
	// When the server processes this call, it MUST follow these guidelines:
	//
	// * If the *ObjectIsInitialized* instance variable is False:
	//
	// * The server MUST return MQ_ERROR_UNINITIALIZED_OBJECT (0xC00E0094), and MUST take
	// no further action.
	//
	// * The server MUST generate a QMMgmtAction event with the following inputs:
	//
	// * *iAction* = "RESUME"
	//
	// * The server MUST return rStatus, and MUST take no further action.
	Resume(context.Context, *ResumeRequest) (*ResumeResponse, error)

	// The Pause method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST pause the transmission of messages from the referenced OutgoingQueue.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.<21>
	Pause(context.Context, *PauseRequest) (*PauseResponse, error)

	// The EodResend method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST resend the pending sequence of transactional messages in the represented
	// OutgoingQueue.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.<22>
	//
	// * If the *ObjectIsInitialized* instance variable is False:
	//
	// * The server ( 3b7be3f7-651c-4f9c-930b-a9a7c4355ad8#gt_434b0234-e970-4e8c-bdfa-e16a30d96703
	// ) MUST return MQ_ERROR_UNINITIALIZED_OBJECT (0xC00E0094) and take no further action.
	//
	// * The server MUST generate a QMMgmtAction event with the following inputs:
	//
	// * iAction = "Pause"
	//
	// * The server MUST return rStatus, and take no further action.
	EODResend(context.Context, *EODResendRequest) (*EODResendResponse, error)
}

func RegisterOutgoingQueueManagementServer(conn dcerpc.Conn, o OutgoingQueueManagementServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewOutgoingQueueManagementServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(OutgoingQueueManagementSyntaxV0_0))...)
}

func NewOutgoingQueueManagementServerHandle(o OutgoingQueueManagementServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return OutgoingQueueManagementServerHandle(ctx, o, opNum, r)
	}
}

func OutgoingQueueManagementServerHandle(ctx context.Context, o OutgoingQueueManagementServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 16 {
		// IMSMQManagement base method.
		return imsmqmanagement.ManagementServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 16: // State
		op := &xxx_GetStateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetStateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetState(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // NextHops
		op := &xxx_GetNextHopsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNextHopsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNextHops(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // EodGetSendInfo
		op := &xxx_EODGetSendInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EODGetSendInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EODGetSendInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // Resume
		op := &xxx_ResumeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ResumeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Resume(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // Pause
		op := &xxx_PauseOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PauseRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Pause(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // EodResend
		op := &xxx_EODResendOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EODResendRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EODResend(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMSMQOutgoingQueueManagement
type UnimplementedOutgoingQueueManagementServer struct {
	imsmqmanagement.UnimplementedManagementServer
}

func (UnimplementedOutgoingQueueManagementServer) GetState(context.Context, *GetStateRequest) (*GetStateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedOutgoingQueueManagementServer) GetNextHops(context.Context, *GetNextHopsRequest) (*GetNextHopsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedOutgoingQueueManagementServer) EODGetSendInfo(context.Context, *EODGetSendInfoRequest) (*EODGetSendInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedOutgoingQueueManagementServer) Resume(context.Context, *ResumeRequest) (*ResumeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedOutgoingQueueManagementServer) Pause(context.Context, *PauseRequest) (*PauseResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedOutgoingQueueManagementServer) EODResend(context.Context, *EODResendRequest) (*EODResendResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ OutgoingQueueManagementServer = (*UnimplementedOutgoingQueueManagementServer)(nil)
