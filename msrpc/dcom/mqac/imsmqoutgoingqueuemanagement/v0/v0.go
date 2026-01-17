package imsmqoutgoingqueuemanagement

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	mqac "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac"
	imsmqmanagement "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqmanagement/v0"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
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
	_ = dcom.GoPackage
	_ = imsmqmanagement.GoPackage
	_ = oaut.GoPackage
	_ = mqac.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/mqac"
)

var (
	// IMSMQOutgoingQueueManagement interface identifier 64c478fb-f9b0-4695-8a7f-439ac94326d3
	OutgoingQueueManagementIID = &dcom.IID{Data1: 0x64c478fb, Data2: 0xf9b0, Data3: 0x4695, Data4: []byte{0x8a, 0x7f, 0x43, 0x9a, 0xc9, 0x43, 0x26, 0xd3}}
	// Syntax UUID
	OutgoingQueueManagementSyntaxUUID = &uuid.UUID{TimeLow: 0x64c478fb, TimeMid: 0xf9b0, TimeHiAndVersion: 0x4695, ClockSeqHiAndReserved: 0x8a, ClockSeqLow: 0x7f, Node: [6]uint8{0x43, 0x9a, 0xc9, 0x43, 0x26, 0xd3}}
	// Syntax ID
	OutgoingQueueManagementSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: OutgoingQueueManagementSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IMSMQOutgoingQueueManagement interface.
type OutgoingQueueManagementClient interface {

	// IMSMQManagement retrieval method.
	Management() imsmqmanagement.ManagementClient

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
	GetState(context.Context, *GetStateRequest, ...dcerpc.CallOption) (*GetStateResponse, error)

	// The NextHops method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return the represented OutgoingQueue.NextHops.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	GetNextHops(context.Context, *GetNextHopsRequest, ...dcerpc.CallOption) (*GetNextHopsResponse, error)

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
	EODGetSendInfo(context.Context, *EODGetSendInfoRequest, ...dcerpc.CallOption) (*EODGetSendInfoResponse, error)

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
	Resume(context.Context, *ResumeRequest, ...dcerpc.CallOption) (*ResumeResponse, error)

	// The Pause method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST pause the transmission of messages from the referenced OutgoingQueue.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.<21>
	Pause(context.Context, *PauseRequest, ...dcerpc.CallOption) (*PauseResponse, error)

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
	EODResend(context.Context, *EODResendRequest, ...dcerpc.CallOption) (*EODResendResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) OutgoingQueueManagementClient
}

type xxx_DefaultOutgoingQueueManagementClient struct {
	imsmqmanagement.ManagementClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultOutgoingQueueManagementClient) Management() imsmqmanagement.ManagementClient {
	return o.ManagementClient
}

func (o *xxx_DefaultOutgoingQueueManagementClient) GetState(ctx context.Context, in *GetStateRequest, opts ...dcerpc.CallOption) (*GetStateResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetStateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultOutgoingQueueManagementClient) GetNextHops(ctx context.Context, in *GetNextHopsRequest, opts ...dcerpc.CallOption) (*GetNextHopsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetNextHopsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultOutgoingQueueManagementClient) EODGetSendInfo(ctx context.Context, in *EODGetSendInfoRequest, opts ...dcerpc.CallOption) (*EODGetSendInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EODGetSendInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultOutgoingQueueManagementClient) Resume(ctx context.Context, in *ResumeRequest, opts ...dcerpc.CallOption) (*ResumeResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ResumeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultOutgoingQueueManagementClient) Pause(ctx context.Context, in *PauseRequest, opts ...dcerpc.CallOption) (*PauseResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &PauseResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultOutgoingQueueManagementClient) EODResend(ctx context.Context, in *EODResendRequest, opts ...dcerpc.CallOption) (*EODResendResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EODResendResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultOutgoingQueueManagementClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultOutgoingQueueManagementClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultOutgoingQueueManagementClient) IPID(ctx context.Context, ipid *dcom.IPID) OutgoingQueueManagementClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultOutgoingQueueManagementClient{
		ManagementClient: o.ManagementClient.IPID(ctx, ipid),
		cc:               o.cc,
		ipid:             ipid,
	}
}

func NewOutgoingQueueManagementClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (OutgoingQueueManagementClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(OutgoingQueueManagementSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := imsmqmanagement.NewManagementClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultOutgoingQueueManagementClient{
		ManagementClient: base,
		cc:               cc,
		ipid:             ipid,
	}, nil
}

// xxx_GetStateOperation structure represents the State operation
type xxx_GetStateOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	State  int32          `idl:"name:plState" json:"state"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetStateOperation) OpNum() int { return 16 }

func (o *xxx_GetStateOperation) OpName() string { return "/IMSMQOutgoingQueueManagement/v0/State" }

func (o *xxx_GetStateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetStateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetStateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetStateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetStateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// plState {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.State); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetStateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// plState {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.State); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetStateRequest structure represents the State operation request
type GetStateRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetStateRequest) xxx_ToOp(ctx context.Context, op *xxx_GetStateOperation) *xxx_GetStateOperation {
	if op == nil {
		op = &xxx_GetStateOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetStateRequest) xxx_FromOp(ctx context.Context, op *xxx_GetStateOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetStateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetStateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetStateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetStateResponse structure represents the State operation response
type GetStateResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// plState: A pointer to a long that corresponds to the QUEUE_STATE (section 2.2.2.19)
	// enumeration.
	State int32 `idl:"name:plState" json:"state"`
	// Return: The State return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetStateResponse) xxx_ToOp(ctx context.Context, op *xxx_GetStateOperation) *xxx_GetStateOperation {
	if op == nil {
		op = &xxx_GetStateOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.State = o.State
	op.Return = o.Return
	return op
}

func (o *GetStateResponse) xxx_FromOp(ctx context.Context, op *xxx_GetStateOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.State = op.State
	o.Return = op.Return
}
func (o *GetStateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetStateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetStateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNextHopsOperation structure represents the NextHops operation
type xxx_GetNextHopsOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	NextHops *oaut.Variant  `idl:"name:pvNextHops" json:"next_hops"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNextHopsOperation) OpNum() int { return 17 }

func (o *xxx_GetNextHopsOperation) OpName() string {
	return "/IMSMQOutgoingQueueManagement/v0/NextHops"
}

func (o *xxx_GetNextHopsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNextHopsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNextHopsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNextHopsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNextHopsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pvNextHops {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.NextHops != nil {
			_ptr_pvNextHops := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.NextHops != nil {
					if err := o.NextHops.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NextHops, _ptr_pvNextHops); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNextHopsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pvNextHops {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pvNextHops := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.NextHops == nil {
				o.NextHops = &oaut.Variant{}
			}
			if err := o.NextHops.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pvNextHops := func(ptr interface{}) { o.NextHops = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.NextHops, _s_pvNextHops, _ptr_pvNextHops); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetNextHopsRequest structure represents the NextHops operation request
type GetNextHopsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetNextHopsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNextHopsOperation) *xxx_GetNextHopsOperation {
	if op == nil {
		op = &xxx_GetNextHopsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetNextHopsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNextHopsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetNextHopsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNextHopsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNextHopsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNextHopsResponse structure represents the NextHops operation response
type GetNextHopsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pvNextHops: A pointer to a VARIANT that contains an array of zero or more strings (VT_ARRAY | VT_BSTR) that specify the routing addresses.
	NextHops *oaut.Variant `idl:"name:pvNextHops" json:"next_hops"`
	// Return: The NextHops return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNextHopsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNextHopsOperation) *xxx_GetNextHopsOperation {
	if op == nil {
		op = &xxx_GetNextHopsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.NextHops = o.NextHops
	op.Return = o.Return
	return op
}

func (o *GetNextHopsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNextHopsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.NextHops = op.NextHops
	o.Return = op.Return
}
func (o *GetNextHopsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNextHopsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNextHopsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EODGetSendInfoOperation structure represents the EodGetSendInfo operation
type xxx_EODGetSendInfoOperation struct {
	This       *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Collection *mqac.Collection `idl:"name:ppCollection" json:"collection"`
	Return     int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_EODGetSendInfoOperation) OpNum() int { return 18 }

func (o *xxx_EODGetSendInfoOperation) OpName() string {
	return "/IMSMQOutgoingQueueManagement/v0/EodGetSendInfo"
}

func (o *xxx_EODGetSendInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EODGetSendInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EODGetSendInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EODGetSendInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EODGetSendInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ppCollection {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQCollection}(interface))
	{
		if o.Collection != nil {
			_ptr_ppCollection := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Collection != nil {
					if err := o.Collection.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Collection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Collection, _ptr_ppCollection); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EODGetSendInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppCollection {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQCollection}(interface))
	{
		_ptr_ppCollection := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Collection == nil {
				o.Collection = &mqac.Collection{}
			}
			if err := o.Collection.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppCollection := func(ptr interface{}) { o.Collection = *ptr.(**mqac.Collection) }
		if err := w.ReadPointer(&o.Collection, _s_ppCollection, _ptr_ppCollection); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EODGetSendInfoRequest structure represents the EodGetSendInfo operation request
type EODGetSendInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *EODGetSendInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_EODGetSendInfoOperation) *xxx_EODGetSendInfoOperation {
	if op == nil {
		op = &xxx_EODGetSendInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *EODGetSendInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_EODGetSendInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *EODGetSendInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EODGetSendInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EODGetSendInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EODGetSendInfoResponse structure represents the EodGetSendInfo operation response
type EODGetSendInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppCollection: A pointer to an EODSourceInfo (section 2.2.4.2) collection.
	Collection *mqac.Collection `idl:"name:ppCollection" json:"collection"`
	// Return: The EodGetSendInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EODGetSendInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_EODGetSendInfoOperation) *xxx_EODGetSendInfoOperation {
	if op == nil {
		op = &xxx_EODGetSendInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Collection = o.Collection
	op.Return = o.Return
	return op
}

func (o *EODGetSendInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_EODGetSendInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Collection = op.Collection
	o.Return = op.Return
}
func (o *EODGetSendInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EODGetSendInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EODGetSendInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ResumeOperation structure represents the Resume operation
type xxx_ResumeOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ResumeOperation) OpNum() int { return 19 }

func (o *xxx_ResumeOperation) OpName() string { return "/IMSMQOutgoingQueueManagement/v0/Resume" }

func (o *xxx_ResumeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResumeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResumeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResumeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResumeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResumeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ResumeRequest structure represents the Resume operation request
type ResumeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *ResumeRequest) xxx_ToOp(ctx context.Context, op *xxx_ResumeOperation) *xxx_ResumeOperation {
	if op == nil {
		op = &xxx_ResumeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *ResumeRequest) xxx_FromOp(ctx context.Context, op *xxx_ResumeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *ResumeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ResumeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResumeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ResumeResponse structure represents the Resume operation response
type ResumeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Resume return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ResumeResponse) xxx_ToOp(ctx context.Context, op *xxx_ResumeOperation) *xxx_ResumeOperation {
	if op == nil {
		op = &xxx_ResumeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ResumeResponse) xxx_FromOp(ctx context.Context, op *xxx_ResumeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ResumeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ResumeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResumeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PauseOperation structure represents the Pause operation
type xxx_PauseOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_PauseOperation) OpNum() int { return 20 }

func (o *xxx_PauseOperation) OpName() string { return "/IMSMQOutgoingQueueManagement/v0/Pause" }

func (o *xxx_PauseOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PauseOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PauseOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PauseOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PauseOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PauseOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PauseRequest structure represents the Pause operation request
type PauseRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *PauseRequest) xxx_ToOp(ctx context.Context, op *xxx_PauseOperation) *xxx_PauseOperation {
	if op == nil {
		op = &xxx_PauseOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *PauseRequest) xxx_FromOp(ctx context.Context, op *xxx_PauseOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *PauseRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PauseRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PauseOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PauseResponse structure represents the Pause operation response
type PauseResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Pause return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PauseResponse) xxx_ToOp(ctx context.Context, op *xxx_PauseOperation) *xxx_PauseOperation {
	if op == nil {
		op = &xxx_PauseOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *PauseResponse) xxx_FromOp(ctx context.Context, op *xxx_PauseOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *PauseResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PauseResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PauseOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EODResendOperation structure represents the EodResend operation
type xxx_EODResendOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_EODResendOperation) OpNum() int { return 21 }

func (o *xxx_EODResendOperation) OpName() string { return "/IMSMQOutgoingQueueManagement/v0/EodResend" }

func (o *xxx_EODResendOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EODResendOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EODResendOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EODResendOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EODResendOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EODResendOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EODResendRequest structure represents the EodResend operation request
type EODResendRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *EODResendRequest) xxx_ToOp(ctx context.Context, op *xxx_EODResendOperation) *xxx_EODResendOperation {
	if op == nil {
		op = &xxx_EODResendOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *EODResendRequest) xxx_FromOp(ctx context.Context, op *xxx_EODResendOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *EODResendRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EODResendRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EODResendOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EODResendResponse structure represents the EodResend operation response
type EODResendResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The EodResend return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EODResendResponse) xxx_ToOp(ctx context.Context, op *xxx_EODResendOperation) *xxx_EODResendOperation {
	if op == nil {
		op = &xxx_EODResendOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *EODResendResponse) xxx_FromOp(ctx context.Context, op *xxx_EODResendOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *EODResendResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EODResendResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EODResendOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
