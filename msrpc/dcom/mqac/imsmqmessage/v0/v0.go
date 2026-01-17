package imsmqmessage

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	mqac "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
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
	_ = dcom.GoPackage
	_ = idispatch.GoPackage
	_ = mqac.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/mqac"
)

var (
	// IMSMQMessage interface identifier d7d6e074-dccd-11d0-aa4b-0060970debae
	MessageIID = &dcom.IID{Data1: 0xd7d6e074, Data2: 0xdccd, Data3: 0x11d0, Data4: []byte{0xaa, 0x4b, 0x00, 0x60, 0x97, 0x0d, 0xeb, 0xae}}
	// Syntax UUID
	MessageSyntaxUUID = &uuid.UUID{TimeLow: 0xd7d6e074, TimeMid: 0xdccd, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0x4b, Node: [6]uint8{0x0, 0x60, 0x97, 0xd, 0xeb, 0xae}}
	// Syntax ID
	MessageSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: MessageSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IMSMQMessage interface.
type MessageClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// The Class method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return the represented Message.Class.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetClass(context.Context, *GetClassRequest, ...dcerpc.CallOption) (*GetClassResponse, error)

	// PrivLevel operation.
	GetPrivacyLevel(context.Context, *GetPrivacyLevelRequest, ...dcerpc.CallOption) (*GetPrivacyLevelResponse, error)

	// PrivLevel operation.
	SetPrivacyLevel(context.Context, *SetPrivacyLevelRequest, ...dcerpc.CallOption) (*SetPrivacyLevelResponse, error)

	// AuthLevel operation.
	GetAuthLevel(context.Context, *GetAuthLevelRequest, ...dcerpc.CallOption) (*GetAuthLevelResponse, error)

	// AuthLevel operation.
	SetAuthLevel(context.Context, *SetAuthLevelRequest, ...dcerpc.CallOption) (*SetAuthLevelResponse, error)

	// The IsAuthenticated method is received by the server in an RPC_REQUEST packet. In
	// response, the server MUST return a BOOLEAN flag indicating whether the message was
	// authenticated by the Queue Manager that received the message.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetIsAuthenticated(context.Context, *GetIsAuthenticatedRequest, ...dcerpc.CallOption) (*GetIsAuthenticatedResponse, error)

	// Delivery operation.
	GetDelivery(context.Context, *GetDeliveryRequest, ...dcerpc.CallOption) (*GetDeliveryResponse, error)

	// Delivery operation.
	SetDelivery(context.Context, *SetDeliveryRequest, ...dcerpc.CallOption) (*SetDeliveryResponse, error)

	// Trace operation.
	GetTrace(context.Context, *GetTraceRequest, ...dcerpc.CallOption) (*GetTraceResponse, error)

	// Trace operation.
	SetTrace(context.Context, *SetTraceRequest, ...dcerpc.CallOption) (*SetTraceResponse, error)

	// Priority operation.
	GetPriority(context.Context, *GetPriorityRequest, ...dcerpc.CallOption) (*GetPriorityResponse, error)

	// Priority operation.
	SetPriority(context.Context, *SetPriorityRequest, ...dcerpc.CallOption) (*SetPriorityResponse, error)

	// Journal operation.
	GetJournal(context.Context, *GetJournalRequest, ...dcerpc.CallOption) (*GetJournalResponse, error)

	// Journal operation.
	SetJournal(context.Context, *SetJournalRequest, ...dcerpc.CallOption) (*SetJournalResponse, error)

	// ResponseQueueInfo operation.
	GetResponseQueueInfo(context.Context, *GetResponseQueueInfoRequest, ...dcerpc.CallOption) (*GetResponseQueueInfoResponse, error)

	// ResponseQueueInfo operation.
	SetByRefResponseQueueInfo(context.Context, *SetByRefResponseQueueInfoRequest, ...dcerpc.CallOption) (*SetByRefResponseQueueInfoResponse, error)

	// AppSpecific operation.
	GetAppSpecific(context.Context, *GetAppSpecificRequest, ...dcerpc.CallOption) (*GetAppSpecificResponse, error)

	// AppSpecific operation.
	SetAppSpecific(context.Context, *SetAppSpecificRequest, ...dcerpc.CallOption) (*SetAppSpecificResponse, error)

	// The SourceMachineGuid method is received by the server in an RPC_REQUEST packet.
	// In response, the server MUST return the represented Message.SourceMachineIdentifier.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetSourceMachineGUID(context.Context, *GetSourceMachineGUIDRequest, ...dcerpc.CallOption) (*GetSourceMachineGUIDResponse, error)

	// The BodyLength method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return the number of bytes in the represented Message.Body.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetBodyLength(context.Context, *GetBodyLengthRequest, ...dcerpc.CallOption) (*GetBodyLengthResponse, error)

	// Body operation.
	GetBody(context.Context, *GetBodyRequest, ...dcerpc.CallOption) (*GetBodyResponse, error)

	// Body operation.
	SetBody(context.Context, *SetBodyRequest, ...dcerpc.CallOption) (*SetBodyResponse, error)

	// AdminQueueInfo operation.
	GetAdminQueueInfo(context.Context, *GetAdminQueueInfoRequest, ...dcerpc.CallOption) (*GetAdminQueueInfoResponse, error)

	// AdminQueueInfo operation.
	SetByRefAdminQueueInfo(context.Context, *SetByRefAdminQueueInfoRequest, ...dcerpc.CallOption) (*SetByRefAdminQueueInfoResponse, error)

	// The Id method is received by the server in an RPC_REQUEST packet. In response, the
	// server MUST return the represented Message.Identifier.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetID(context.Context, *GetIDRequest, ...dcerpc.CallOption) (*GetIDResponse, error)

	// CorrelationId operation.
	GetCorrelationID(context.Context, *GetCorrelationIDRequest, ...dcerpc.CallOption) (*GetCorrelationIDResponse, error)

	// CorrelationId operation.
	SetCorrelationID(context.Context, *SetCorrelationIDRequest, ...dcerpc.CallOption) (*SetCorrelationIDResponse, error)

	// Ack operation.
	GetAck(context.Context, *GetAckRequest, ...dcerpc.CallOption) (*GetAckResponse, error)

	// Ack operation.
	SetAck(context.Context, *SetAckRequest, ...dcerpc.CallOption) (*SetAckResponse, error)

	// Label operation.
	GetLabel(context.Context, *GetLabelRequest, ...dcerpc.CallOption) (*GetLabelResponse, error)

	// Label operation.
	SetLabel(context.Context, *SetLabelRequest, ...dcerpc.CallOption) (*SetLabelResponse, error)

	// MaxTimeToReachQueue operation.
	GetMaxTimeToReachQueue(context.Context, *GetMaxTimeToReachQueueRequest, ...dcerpc.CallOption) (*GetMaxTimeToReachQueueResponse, error)

	// MaxTimeToReachQueue operation.
	SetMaxTimeToReachQueue(context.Context, *SetMaxTimeToReachQueueRequest, ...dcerpc.CallOption) (*SetMaxTimeToReachQueueResponse, error)

	// MaxTimeToReceive operation.
	GetMaxTimeToReceive(context.Context, *GetMaxTimeToReceiveRequest, ...dcerpc.CallOption) (*GetMaxTimeToReceiveResponse, error)

	// MaxTimeToReceive operation.
	SetMaxTimeToReceive(context.Context, *SetMaxTimeToReceiveRequest, ...dcerpc.CallOption) (*SetMaxTimeToReceiveResponse, error)

	// HashAlgorithm operation.
	GetHashAlgorithm(context.Context, *GetHashAlgorithmRequest, ...dcerpc.CallOption) (*GetHashAlgorithmResponse, error)

	// HashAlgorithm operation.
	SetHashAlgorithm(context.Context, *SetHashAlgorithmRequest, ...dcerpc.CallOption) (*SetHashAlgorithmResponse, error)

	// EncryptAlgorithm operation.
	GetEncryptAlgorithm(context.Context, *GetEncryptAlgorithmRequest, ...dcerpc.CallOption) (*GetEncryptAlgorithmResponse, error)

	// EncryptAlgorithm operation.
	SetEncryptAlgorithm(context.Context, *SetEncryptAlgorithmRequest, ...dcerpc.CallOption) (*SetEncryptAlgorithmResponse, error)

	// The SentTime method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return the represented Message.SentTime.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetSentTime(context.Context, *GetSentTimeRequest, ...dcerpc.CallOption) (*GetSentTimeResponse, error)

	// The ArrivedTime method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return the represented Message.ArrivalTime.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetArrivedTime(context.Context, *GetArrivedTimeRequest, ...dcerpc.CallOption) (*GetArrivedTimeResponse, error)

	// The DestinationQueueInfo method is received by the server in an RPC_REQUEST packet.
	// In response, the server MUST return an IMSMQQueueInfo interface pointer to an MSMQQueueInfo
	// object that represents the Queue identified by the represented Message.DestinationQueueFormatName.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetDestinationQueueInfo(context.Context, *GetDestinationQueueInfoRequest, ...dcerpc.CallOption) (*GetDestinationQueueInfoResponse, error)

	// SenderCertificate operation.
	GetSenderCertificate(context.Context, *GetSenderCertificateRequest, ...dcerpc.CallOption) (*GetSenderCertificateResponse, error)

	// SenderCertificate operation.
	SetSenderCertificate(context.Context, *SetSenderCertificateRequest, ...dcerpc.CallOption) (*SetSenderCertificateResponse, error)

	// SenderId operation.
	GetSenderID(context.Context, *GetSenderIDRequest, ...dcerpc.CallOption) (*GetSenderIDResponse, error)

	// SenderIdType operation.
	GetSenderIDType(context.Context, *GetSenderIDTypeRequest, ...dcerpc.CallOption) (*GetSenderIDTypeResponse, error)

	// SenderIdType operation.
	SetSenderIDType(context.Context, *SetSenderIDTypeRequest, ...dcerpc.CallOption) (*SetSenderIDTypeResponse, error)

	// The Send method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST send a message.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	Send(context.Context, *SendRequest, ...dcerpc.CallOption) (*SendResponse, error)

	// The AttachCurrentSecurityContext method is received by the server in an RPC_REQUEST
	// packet. In response, the server MUST cache the relevant information required to sign
	// a message on behalf of the client, including the Message.SenderIdentifier and Message.SenderCertificate.
	// This method is provided purely as an optimization to allow the client to reduce lookups
	// of the security information about the calling client each time the message is sent.
	// The represented Message.SenderIdentifier and Message.SenderCertificate property values
	// MUST NOT be updated as a result of calling this method. This method is superseded
	// by IMSMQMessage4::AttachSecurityContext2.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	AttachCurrentSecurityContext(context.Context, *AttachCurrentSecurityContextRequest, ...dcerpc.CallOption) (*AttachCurrentSecurityContextResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) MessageClient
}

type xxx_DefaultMessageClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultMessageClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultMessageClient) GetClass(ctx context.Context, in *GetClassRequest, opts ...dcerpc.CallOption) (*GetClassResponse, error) {
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
	out := &GetClassResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) GetPrivacyLevel(ctx context.Context, in *GetPrivacyLevelRequest, opts ...dcerpc.CallOption) (*GetPrivacyLevelResponse, error) {
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
	out := &GetPrivacyLevelResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) SetPrivacyLevel(ctx context.Context, in *SetPrivacyLevelRequest, opts ...dcerpc.CallOption) (*SetPrivacyLevelResponse, error) {
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
	out := &SetPrivacyLevelResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) GetAuthLevel(ctx context.Context, in *GetAuthLevelRequest, opts ...dcerpc.CallOption) (*GetAuthLevelResponse, error) {
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
	out := &GetAuthLevelResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) SetAuthLevel(ctx context.Context, in *SetAuthLevelRequest, opts ...dcerpc.CallOption) (*SetAuthLevelResponse, error) {
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
	out := &SetAuthLevelResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) GetIsAuthenticated(ctx context.Context, in *GetIsAuthenticatedRequest, opts ...dcerpc.CallOption) (*GetIsAuthenticatedResponse, error) {
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
	out := &GetIsAuthenticatedResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) GetDelivery(ctx context.Context, in *GetDeliveryRequest, opts ...dcerpc.CallOption) (*GetDeliveryResponse, error) {
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
	out := &GetDeliveryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) SetDelivery(ctx context.Context, in *SetDeliveryRequest, opts ...dcerpc.CallOption) (*SetDeliveryResponse, error) {
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
	out := &SetDeliveryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) GetTrace(ctx context.Context, in *GetTraceRequest, opts ...dcerpc.CallOption) (*GetTraceResponse, error) {
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
	out := &GetTraceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) SetTrace(ctx context.Context, in *SetTraceRequest, opts ...dcerpc.CallOption) (*SetTraceResponse, error) {
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
	out := &SetTraceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) GetPriority(ctx context.Context, in *GetPriorityRequest, opts ...dcerpc.CallOption) (*GetPriorityResponse, error) {
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
	out := &GetPriorityResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) SetPriority(ctx context.Context, in *SetPriorityRequest, opts ...dcerpc.CallOption) (*SetPriorityResponse, error) {
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
	out := &SetPriorityResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) GetJournal(ctx context.Context, in *GetJournalRequest, opts ...dcerpc.CallOption) (*GetJournalResponse, error) {
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
	out := &GetJournalResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) SetJournal(ctx context.Context, in *SetJournalRequest, opts ...dcerpc.CallOption) (*SetJournalResponse, error) {
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
	out := &SetJournalResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) GetResponseQueueInfo(ctx context.Context, in *GetResponseQueueInfoRequest, opts ...dcerpc.CallOption) (*GetResponseQueueInfoResponse, error) {
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
	out := &GetResponseQueueInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) SetByRefResponseQueueInfo(ctx context.Context, in *SetByRefResponseQueueInfoRequest, opts ...dcerpc.CallOption) (*SetByRefResponseQueueInfoResponse, error) {
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
	out := &SetByRefResponseQueueInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) GetAppSpecific(ctx context.Context, in *GetAppSpecificRequest, opts ...dcerpc.CallOption) (*GetAppSpecificResponse, error) {
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
	out := &GetAppSpecificResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) SetAppSpecific(ctx context.Context, in *SetAppSpecificRequest, opts ...dcerpc.CallOption) (*SetAppSpecificResponse, error) {
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
	out := &SetAppSpecificResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) GetSourceMachineGUID(ctx context.Context, in *GetSourceMachineGUIDRequest, opts ...dcerpc.CallOption) (*GetSourceMachineGUIDResponse, error) {
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
	out := &GetSourceMachineGUIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) GetBodyLength(ctx context.Context, in *GetBodyLengthRequest, opts ...dcerpc.CallOption) (*GetBodyLengthResponse, error) {
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
	out := &GetBodyLengthResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) GetBody(ctx context.Context, in *GetBodyRequest, opts ...dcerpc.CallOption) (*GetBodyResponse, error) {
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
	out := &GetBodyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) SetBody(ctx context.Context, in *SetBodyRequest, opts ...dcerpc.CallOption) (*SetBodyResponse, error) {
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
	out := &SetBodyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) GetAdminQueueInfo(ctx context.Context, in *GetAdminQueueInfoRequest, opts ...dcerpc.CallOption) (*GetAdminQueueInfoResponse, error) {
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
	out := &GetAdminQueueInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) SetByRefAdminQueueInfo(ctx context.Context, in *SetByRefAdminQueueInfoRequest, opts ...dcerpc.CallOption) (*SetByRefAdminQueueInfoResponse, error) {
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
	out := &SetByRefAdminQueueInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) GetID(ctx context.Context, in *GetIDRequest, opts ...dcerpc.CallOption) (*GetIDResponse, error) {
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
	out := &GetIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) GetCorrelationID(ctx context.Context, in *GetCorrelationIDRequest, opts ...dcerpc.CallOption) (*GetCorrelationIDResponse, error) {
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
	out := &GetCorrelationIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) SetCorrelationID(ctx context.Context, in *SetCorrelationIDRequest, opts ...dcerpc.CallOption) (*SetCorrelationIDResponse, error) {
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
	out := &SetCorrelationIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) GetAck(ctx context.Context, in *GetAckRequest, opts ...dcerpc.CallOption) (*GetAckResponse, error) {
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
	out := &GetAckResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) SetAck(ctx context.Context, in *SetAckRequest, opts ...dcerpc.CallOption) (*SetAckResponse, error) {
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
	out := &SetAckResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) GetLabel(ctx context.Context, in *GetLabelRequest, opts ...dcerpc.CallOption) (*GetLabelResponse, error) {
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
	out := &GetLabelResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) SetLabel(ctx context.Context, in *SetLabelRequest, opts ...dcerpc.CallOption) (*SetLabelResponse, error) {
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
	out := &SetLabelResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) GetMaxTimeToReachQueue(ctx context.Context, in *GetMaxTimeToReachQueueRequest, opts ...dcerpc.CallOption) (*GetMaxTimeToReachQueueResponse, error) {
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
	out := &GetMaxTimeToReachQueueResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) SetMaxTimeToReachQueue(ctx context.Context, in *SetMaxTimeToReachQueueRequest, opts ...dcerpc.CallOption) (*SetMaxTimeToReachQueueResponse, error) {
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
	out := &SetMaxTimeToReachQueueResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) GetMaxTimeToReceive(ctx context.Context, in *GetMaxTimeToReceiveRequest, opts ...dcerpc.CallOption) (*GetMaxTimeToReceiveResponse, error) {
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
	out := &GetMaxTimeToReceiveResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) SetMaxTimeToReceive(ctx context.Context, in *SetMaxTimeToReceiveRequest, opts ...dcerpc.CallOption) (*SetMaxTimeToReceiveResponse, error) {
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
	out := &SetMaxTimeToReceiveResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) GetHashAlgorithm(ctx context.Context, in *GetHashAlgorithmRequest, opts ...dcerpc.CallOption) (*GetHashAlgorithmResponse, error) {
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
	out := &GetHashAlgorithmResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) SetHashAlgorithm(ctx context.Context, in *SetHashAlgorithmRequest, opts ...dcerpc.CallOption) (*SetHashAlgorithmResponse, error) {
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
	out := &SetHashAlgorithmResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) GetEncryptAlgorithm(ctx context.Context, in *GetEncryptAlgorithmRequest, opts ...dcerpc.CallOption) (*GetEncryptAlgorithmResponse, error) {
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
	out := &GetEncryptAlgorithmResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) SetEncryptAlgorithm(ctx context.Context, in *SetEncryptAlgorithmRequest, opts ...dcerpc.CallOption) (*SetEncryptAlgorithmResponse, error) {
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
	out := &SetEncryptAlgorithmResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) GetSentTime(ctx context.Context, in *GetSentTimeRequest, opts ...dcerpc.CallOption) (*GetSentTimeResponse, error) {
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
	out := &GetSentTimeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) GetArrivedTime(ctx context.Context, in *GetArrivedTimeRequest, opts ...dcerpc.CallOption) (*GetArrivedTimeResponse, error) {
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
	out := &GetArrivedTimeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) GetDestinationQueueInfo(ctx context.Context, in *GetDestinationQueueInfoRequest, opts ...dcerpc.CallOption) (*GetDestinationQueueInfoResponse, error) {
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
	out := &GetDestinationQueueInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) GetSenderCertificate(ctx context.Context, in *GetSenderCertificateRequest, opts ...dcerpc.CallOption) (*GetSenderCertificateResponse, error) {
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
	out := &GetSenderCertificateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) SetSenderCertificate(ctx context.Context, in *SetSenderCertificateRequest, opts ...dcerpc.CallOption) (*SetSenderCertificateResponse, error) {
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
	out := &SetSenderCertificateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) GetSenderID(ctx context.Context, in *GetSenderIDRequest, opts ...dcerpc.CallOption) (*GetSenderIDResponse, error) {
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
	out := &GetSenderIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) GetSenderIDType(ctx context.Context, in *GetSenderIDTypeRequest, opts ...dcerpc.CallOption) (*GetSenderIDTypeResponse, error) {
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
	out := &GetSenderIDTypeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) SetSenderIDType(ctx context.Context, in *SetSenderIDTypeRequest, opts ...dcerpc.CallOption) (*SetSenderIDTypeResponse, error) {
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
	out := &SetSenderIDTypeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) Send(ctx context.Context, in *SendRequest, opts ...dcerpc.CallOption) (*SendResponse, error) {
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
	out := &SendResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) AttachCurrentSecurityContext(ctx context.Context, in *AttachCurrentSecurityContextRequest, opts ...dcerpc.CallOption) (*AttachCurrentSecurityContextResponse, error) {
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
	out := &AttachCurrentSecurityContextResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessageClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultMessageClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultMessageClient) IPID(ctx context.Context, ipid *dcom.IPID) MessageClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultMessageClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewMessageClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (MessageClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(MessageSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := idispatch.NewDispatchClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultMessageClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_GetClassOperation structure represents the Class operation
type xxx_GetClassOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Class  int32          `idl:"name:plClass" json:"class"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetClassOperation) OpNum() int { return 7 }

func (o *xxx_GetClassOperation) OpName() string { return "/IMSMQMessage/v0/Class" }

func (o *xxx_GetClassOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetClassOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetClassOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plClass {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.Class); err != nil {
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

func (o *xxx_GetClassOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plClass {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.Class); err != nil {
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

// GetClassRequest structure represents the Class operation request
type GetClassRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetClassRequest) xxx_ToOp(ctx context.Context, op *xxx_GetClassOperation) *xxx_GetClassOperation {
	if op == nil {
		op = &xxx_GetClassOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetClassRequest) xxx_FromOp(ctx context.Context, op *xxx_GetClassOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetClassRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetClassRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClassOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetClassResponse structure represents the Class operation response
type GetClassResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// plClass: A pointer to a long that identifies the message type. This parameter corresponds
	// to the MQMSGCLASS (section 2.2.2.9) enum.
	Class int32 `idl:"name:plClass" json:"class"`
	// Return: The Class return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetClassResponse) xxx_ToOp(ctx context.Context, op *xxx_GetClassOperation) *xxx_GetClassOperation {
	if op == nil {
		op = &xxx_GetClassOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Class = o.Class
	op.Return = o.Return
	return op
}

func (o *GetClassResponse) xxx_FromOp(ctx context.Context, op *xxx_GetClassOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Class = op.Class
	o.Return = op.Return
}
func (o *GetClassResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetClassResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClassOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPrivacyLevelOperation structure represents the PrivLevel operation
type xxx_GetPrivacyLevelOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	PrivacyLevel int32          `idl:"name:plPrivLevel" json:"privacy_level"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPrivacyLevelOperation) OpNum() int { return 8 }

func (o *xxx_GetPrivacyLevelOperation) OpName() string { return "/IMSMQMessage/v0/PrivLevel" }

func (o *xxx_GetPrivacyLevelOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPrivacyLevelOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPrivacyLevelOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetPrivacyLevelOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPrivacyLevelOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plPrivLevel {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.PrivacyLevel); err != nil {
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

func (o *xxx_GetPrivacyLevelOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plPrivLevel {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.PrivacyLevel); err != nil {
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

// GetPrivacyLevelRequest structure represents the PrivLevel operation request
type GetPrivacyLevelRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPrivacyLevelRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPrivacyLevelOperation) *xxx_GetPrivacyLevelOperation {
	if op == nil {
		op = &xxx_GetPrivacyLevelOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetPrivacyLevelRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPrivacyLevelOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPrivacyLevelRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPrivacyLevelRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPrivacyLevelOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPrivacyLevelResponse structure represents the PrivLevel operation response
type GetPrivacyLevelResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	PrivacyLevel int32          `idl:"name:plPrivLevel" json:"privacy_level"`
	// Return: The PrivLevel return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPrivacyLevelResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPrivacyLevelOperation) *xxx_GetPrivacyLevelOperation {
	if op == nil {
		op = &xxx_GetPrivacyLevelOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.PrivacyLevel = o.PrivacyLevel
	op.Return = o.Return
	return op
}

func (o *GetPrivacyLevelResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPrivacyLevelOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PrivacyLevel = op.PrivacyLevel
	o.Return = op.Return
}
func (o *GetPrivacyLevelResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPrivacyLevelResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPrivacyLevelOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetPrivacyLevelOperation structure represents the PrivLevel operation
type xxx_SetPrivacyLevelOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	PrivacyLevel int32          `idl:"name:lPrivLevel" json:"privacy_level"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetPrivacyLevelOperation) OpNum() int { return 9 }

func (o *xxx_SetPrivacyLevelOperation) OpName() string { return "/IMSMQMessage/v0/PrivLevel" }

func (o *xxx_SetPrivacyLevelOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPrivacyLevelOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lPrivLevel {in} (1:(int32))
	{
		if err := w.WriteData(o.PrivacyLevel); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPrivacyLevelOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lPrivLevel {in} (1:(int32))
	{
		if err := w.ReadData(&o.PrivacyLevel); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPrivacyLevelOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPrivacyLevelOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetPrivacyLevelOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetPrivacyLevelRequest structure represents the PrivLevel operation request
type SetPrivacyLevelRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	PrivacyLevel int32          `idl:"name:lPrivLevel" json:"privacy_level"`
}

func (o *SetPrivacyLevelRequest) xxx_ToOp(ctx context.Context, op *xxx_SetPrivacyLevelOperation) *xxx_SetPrivacyLevelOperation {
	if op == nil {
		op = &xxx_SetPrivacyLevelOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PrivacyLevel = o.PrivacyLevel
	return op
}

func (o *SetPrivacyLevelRequest) xxx_FromOp(ctx context.Context, op *xxx_SetPrivacyLevelOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PrivacyLevel = op.PrivacyLevel
}
func (o *SetPrivacyLevelRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetPrivacyLevelRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPrivacyLevelOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetPrivacyLevelResponse structure represents the PrivLevel operation response
type SetPrivacyLevelResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The PrivLevel return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetPrivacyLevelResponse) xxx_ToOp(ctx context.Context, op *xxx_SetPrivacyLevelOperation) *xxx_SetPrivacyLevelOperation {
	if op == nil {
		op = &xxx_SetPrivacyLevelOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetPrivacyLevelResponse) xxx_FromOp(ctx context.Context, op *xxx_SetPrivacyLevelOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetPrivacyLevelResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetPrivacyLevelResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPrivacyLevelOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAuthLevelOperation structure represents the AuthLevel operation
type xxx_GetAuthLevelOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	AuthLevel int32          `idl:"name:plAuthLevel" json:"auth_level"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAuthLevelOperation) OpNum() int { return 10 }

func (o *xxx_GetAuthLevelOperation) OpName() string { return "/IMSMQMessage/v0/AuthLevel" }

func (o *xxx_GetAuthLevelOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAuthLevelOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetAuthLevelOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetAuthLevelOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAuthLevelOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plAuthLevel {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.AuthLevel); err != nil {
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

func (o *xxx_GetAuthLevelOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plAuthLevel {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.AuthLevel); err != nil {
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

// GetAuthLevelRequest structure represents the AuthLevel operation request
type GetAuthLevelRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetAuthLevelRequest) xxx_ToOp(ctx context.Context, op *xxx_GetAuthLevelOperation) *xxx_GetAuthLevelOperation {
	if op == nil {
		op = &xxx_GetAuthLevelOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetAuthLevelRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAuthLevelOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetAuthLevelRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetAuthLevelRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAuthLevelOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAuthLevelResponse structure represents the AuthLevel operation response
type GetAuthLevelResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	AuthLevel int32          `idl:"name:plAuthLevel" json:"auth_level"`
	// Return: The AuthLevel return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAuthLevelResponse) xxx_ToOp(ctx context.Context, op *xxx_GetAuthLevelOperation) *xxx_GetAuthLevelOperation {
	if op == nil {
		op = &xxx_GetAuthLevelOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.AuthLevel = o.AuthLevel
	op.Return = o.Return
	return op
}

func (o *GetAuthLevelResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAuthLevelOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.AuthLevel = op.AuthLevel
	o.Return = op.Return
}
func (o *GetAuthLevelResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetAuthLevelResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAuthLevelOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetAuthLevelOperation structure represents the AuthLevel operation
type xxx_SetAuthLevelOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	AuthLevel int32          `idl:"name:lAuthLevel" json:"auth_level"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetAuthLevelOperation) OpNum() int { return 11 }

func (o *xxx_SetAuthLevelOperation) OpName() string { return "/IMSMQMessage/v0/AuthLevel" }

func (o *xxx_SetAuthLevelOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAuthLevelOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lAuthLevel {in} (1:(int32))
	{
		if err := w.WriteData(o.AuthLevel); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAuthLevelOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lAuthLevel {in} (1:(int32))
	{
		if err := w.ReadData(&o.AuthLevel); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAuthLevelOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAuthLevelOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetAuthLevelOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetAuthLevelRequest structure represents the AuthLevel operation request
type SetAuthLevelRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	AuthLevel int32          `idl:"name:lAuthLevel" json:"auth_level"`
}

func (o *SetAuthLevelRequest) xxx_ToOp(ctx context.Context, op *xxx_SetAuthLevelOperation) *xxx_SetAuthLevelOperation {
	if op == nil {
		op = &xxx_SetAuthLevelOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.AuthLevel = o.AuthLevel
	return op
}

func (o *SetAuthLevelRequest) xxx_FromOp(ctx context.Context, op *xxx_SetAuthLevelOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.AuthLevel = op.AuthLevel
}
func (o *SetAuthLevelRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetAuthLevelRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAuthLevelOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetAuthLevelResponse structure represents the AuthLevel operation response
type SetAuthLevelResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AuthLevel return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetAuthLevelResponse) xxx_ToOp(ctx context.Context, op *xxx_SetAuthLevelOperation) *xxx_SetAuthLevelOperation {
	if op == nil {
		op = &xxx_SetAuthLevelOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetAuthLevelResponse) xxx_FromOp(ctx context.Context, op *xxx_SetAuthLevelOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetAuthLevelResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetAuthLevelResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAuthLevelOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetIsAuthenticatedOperation structure represents the IsAuthenticated operation
type xxx_GetIsAuthenticatedOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	IsAuthenticated int16          `idl:"name:pisAuthenticated" json:"is_authenticated"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetIsAuthenticatedOperation) OpNum() int { return 12 }

func (o *xxx_GetIsAuthenticatedOperation) OpName() string { return "/IMSMQMessage/v0/IsAuthenticated" }

func (o *xxx_GetIsAuthenticatedOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsAuthenticatedOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetIsAuthenticatedOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetIsAuthenticatedOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsAuthenticatedOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pisAuthenticated {out, retval} (1:{pointer=ref}*(1)(int16))
	{
		if err := w.WriteData(o.IsAuthenticated); err != nil {
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

func (o *xxx_GetIsAuthenticatedOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pisAuthenticated {out, retval} (1:{pointer=ref}*(1)(int16))
	{
		if err := w.ReadData(&o.IsAuthenticated); err != nil {
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

// GetIsAuthenticatedRequest structure represents the IsAuthenticated operation request
type GetIsAuthenticatedRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetIsAuthenticatedRequest) xxx_ToOp(ctx context.Context, op *xxx_GetIsAuthenticatedOperation) *xxx_GetIsAuthenticatedOperation {
	if op == nil {
		op = &xxx_GetIsAuthenticatedOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetIsAuthenticatedRequest) xxx_FromOp(ctx context.Context, op *xxx_GetIsAuthenticatedOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetIsAuthenticatedRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetIsAuthenticatedRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsAuthenticatedOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetIsAuthenticatedResponse structure represents the IsAuthenticated operation response
type GetIsAuthenticatedResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pisAuthenticated: A pointer to a short that specifies whether the message was authenticated.
	IsAuthenticated int16 `idl:"name:pisAuthenticated" json:"is_authenticated"`
	// Return: The IsAuthenticated return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetIsAuthenticatedResponse) xxx_ToOp(ctx context.Context, op *xxx_GetIsAuthenticatedOperation) *xxx_GetIsAuthenticatedOperation {
	if op == nil {
		op = &xxx_GetIsAuthenticatedOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.IsAuthenticated = o.IsAuthenticated
	op.Return = o.Return
	return op
}

func (o *GetIsAuthenticatedResponse) xxx_FromOp(ctx context.Context, op *xxx_GetIsAuthenticatedOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.IsAuthenticated = op.IsAuthenticated
	o.Return = op.Return
}
func (o *GetIsAuthenticatedResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetIsAuthenticatedResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsAuthenticatedOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDeliveryOperation structure represents the Delivery operation
type xxx_GetDeliveryOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Delivery int32          `idl:"name:plDelivery" json:"delivery"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDeliveryOperation) OpNum() int { return 13 }

func (o *xxx_GetDeliveryOperation) OpName() string { return "/IMSMQMessage/v0/Delivery" }

func (o *xxx_GetDeliveryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDeliveryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetDeliveryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetDeliveryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDeliveryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plDelivery {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.Delivery); err != nil {
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

func (o *xxx_GetDeliveryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plDelivery {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.Delivery); err != nil {
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

// GetDeliveryRequest structure represents the Delivery operation request
type GetDeliveryRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetDeliveryRequest) xxx_ToOp(ctx context.Context, op *xxx_GetDeliveryOperation) *xxx_GetDeliveryOperation {
	if op == nil {
		op = &xxx_GetDeliveryOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetDeliveryRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDeliveryOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetDeliveryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetDeliveryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDeliveryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDeliveryResponse structure represents the Delivery operation response
type GetDeliveryResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Delivery int32          `idl:"name:plDelivery" json:"delivery"`
	// Return: The Delivery return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDeliveryResponse) xxx_ToOp(ctx context.Context, op *xxx_GetDeliveryOperation) *xxx_GetDeliveryOperation {
	if op == nil {
		op = &xxx_GetDeliveryOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Delivery = o.Delivery
	op.Return = o.Return
	return op
}

func (o *GetDeliveryResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDeliveryOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Delivery = op.Delivery
	o.Return = op.Return
}
func (o *GetDeliveryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetDeliveryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDeliveryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetDeliveryOperation structure represents the Delivery operation
type xxx_SetDeliveryOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Delivery int32          `idl:"name:lDelivery" json:"delivery"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetDeliveryOperation) OpNum() int { return 14 }

func (o *xxx_SetDeliveryOperation) OpName() string { return "/IMSMQMessage/v0/Delivery" }

func (o *xxx_SetDeliveryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDeliveryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lDelivery {in} (1:(int32))
	{
		if err := w.WriteData(o.Delivery); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDeliveryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lDelivery {in} (1:(int32))
	{
		if err := w.ReadData(&o.Delivery); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDeliveryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDeliveryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetDeliveryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetDeliveryRequest structure represents the Delivery operation request
type SetDeliveryRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	Delivery int32          `idl:"name:lDelivery" json:"delivery"`
}

func (o *SetDeliveryRequest) xxx_ToOp(ctx context.Context, op *xxx_SetDeliveryOperation) *xxx_SetDeliveryOperation {
	if op == nil {
		op = &xxx_SetDeliveryOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Delivery = o.Delivery
	return op
}

func (o *SetDeliveryRequest) xxx_FromOp(ctx context.Context, op *xxx_SetDeliveryOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Delivery = op.Delivery
}
func (o *SetDeliveryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetDeliveryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDeliveryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetDeliveryResponse structure represents the Delivery operation response
type SetDeliveryResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Delivery return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetDeliveryResponse) xxx_ToOp(ctx context.Context, op *xxx_SetDeliveryOperation) *xxx_SetDeliveryOperation {
	if op == nil {
		op = &xxx_SetDeliveryOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetDeliveryResponse) xxx_FromOp(ctx context.Context, op *xxx_SetDeliveryOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetDeliveryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetDeliveryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDeliveryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetTraceOperation structure represents the Trace operation
type xxx_GetTraceOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Trace  int32          `idl:"name:plTrace" json:"trace"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTraceOperation) OpNum() int { return 15 }

func (o *xxx_GetTraceOperation) OpName() string { return "/IMSMQMessage/v0/Trace" }

func (o *xxx_GetTraceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTraceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetTraceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetTraceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTraceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plTrace {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.Trace); err != nil {
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

func (o *xxx_GetTraceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plTrace {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.Trace); err != nil {
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

// GetTraceRequest structure represents the Trace operation request
type GetTraceRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetTraceRequest) xxx_ToOp(ctx context.Context, op *xxx_GetTraceOperation) *xxx_GetTraceOperation {
	if op == nil {
		op = &xxx_GetTraceOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetTraceRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTraceOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetTraceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetTraceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTraceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetTraceResponse structure represents the Trace operation response
type GetTraceResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat `idl:"name:That" json:"that"`
	Trace int32          `idl:"name:plTrace" json:"trace"`
	// Return: The Trace return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTraceResponse) xxx_ToOp(ctx context.Context, op *xxx_GetTraceOperation) *xxx_GetTraceOperation {
	if op == nil {
		op = &xxx_GetTraceOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Trace = o.Trace
	op.Return = o.Return
	return op
}

func (o *GetTraceResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTraceOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Trace = op.Trace
	o.Return = op.Return
}
func (o *GetTraceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetTraceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTraceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetTraceOperation structure represents the Trace operation
type xxx_SetTraceOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Trace  int32          `idl:"name:lTrace" json:"trace"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetTraceOperation) OpNum() int { return 16 }

func (o *xxx_SetTraceOperation) OpName() string { return "/IMSMQMessage/v0/Trace" }

func (o *xxx_SetTraceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTraceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lTrace {in} (1:(int32))
	{
		if err := w.WriteData(o.Trace); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTraceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lTrace {in} (1:(int32))
	{
		if err := w.ReadData(&o.Trace); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTraceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTraceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetTraceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetTraceRequest structure represents the Trace operation request
type SetTraceRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This  *dcom.ORPCThis `idl:"name:This" json:"this"`
	Trace int32          `idl:"name:lTrace" json:"trace"`
}

func (o *SetTraceRequest) xxx_ToOp(ctx context.Context, op *xxx_SetTraceOperation) *xxx_SetTraceOperation {
	if op == nil {
		op = &xxx_SetTraceOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Trace = o.Trace
	return op
}

func (o *SetTraceRequest) xxx_FromOp(ctx context.Context, op *xxx_SetTraceOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Trace = op.Trace
}
func (o *SetTraceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetTraceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetTraceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetTraceResponse structure represents the Trace operation response
type SetTraceResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Trace return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetTraceResponse) xxx_ToOp(ctx context.Context, op *xxx_SetTraceOperation) *xxx_SetTraceOperation {
	if op == nil {
		op = &xxx_SetTraceOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetTraceResponse) xxx_FromOp(ctx context.Context, op *xxx_SetTraceOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetTraceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetTraceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetTraceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPriorityOperation structure represents the Priority operation
type xxx_GetPriorityOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Priority int32          `idl:"name:plPriority" json:"priority"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPriorityOperation) OpNum() int { return 17 }

func (o *xxx_GetPriorityOperation) OpName() string { return "/IMSMQMessage/v0/Priority" }

func (o *xxx_GetPriorityOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPriorityOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPriorityOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetPriorityOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPriorityOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plPriority {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.Priority); err != nil {
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

func (o *xxx_GetPriorityOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plPriority {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.Priority); err != nil {
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

// GetPriorityRequest structure represents the Priority operation request
type GetPriorityRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPriorityRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPriorityOperation) *xxx_GetPriorityOperation {
	if op == nil {
		op = &xxx_GetPriorityOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetPriorityRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPriorityOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPriorityRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPriorityRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPriorityOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPriorityResponse structure represents the Priority operation response
type GetPriorityResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Priority int32          `idl:"name:plPriority" json:"priority"`
	// Return: The Priority return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPriorityResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPriorityOperation) *xxx_GetPriorityOperation {
	if op == nil {
		op = &xxx_GetPriorityOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Priority = o.Priority
	op.Return = o.Return
	return op
}

func (o *GetPriorityResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPriorityOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Priority = op.Priority
	o.Return = op.Return
}
func (o *GetPriorityResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPriorityResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPriorityOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetPriorityOperation structure represents the Priority operation
type xxx_SetPriorityOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Priority int32          `idl:"name:lPriority" json:"priority"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetPriorityOperation) OpNum() int { return 18 }

func (o *xxx_SetPriorityOperation) OpName() string { return "/IMSMQMessage/v0/Priority" }

func (o *xxx_SetPriorityOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPriorityOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lPriority {in} (1:(int32))
	{
		if err := w.WriteData(o.Priority); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPriorityOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lPriority {in} (1:(int32))
	{
		if err := w.ReadData(&o.Priority); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPriorityOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPriorityOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetPriorityOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetPriorityRequest structure represents the Priority operation request
type SetPriorityRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	Priority int32          `idl:"name:lPriority" json:"priority"`
}

func (o *SetPriorityRequest) xxx_ToOp(ctx context.Context, op *xxx_SetPriorityOperation) *xxx_SetPriorityOperation {
	if op == nil {
		op = &xxx_SetPriorityOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Priority = o.Priority
	return op
}

func (o *SetPriorityRequest) xxx_FromOp(ctx context.Context, op *xxx_SetPriorityOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Priority = op.Priority
}
func (o *SetPriorityRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetPriorityRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPriorityOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetPriorityResponse structure represents the Priority operation response
type SetPriorityResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Priority return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetPriorityResponse) xxx_ToOp(ctx context.Context, op *xxx_SetPriorityOperation) *xxx_SetPriorityOperation {
	if op == nil {
		op = &xxx_SetPriorityOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetPriorityResponse) xxx_FromOp(ctx context.Context, op *xxx_SetPriorityOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetPriorityResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetPriorityResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPriorityOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetJournalOperation structure represents the Journal operation
type xxx_GetJournalOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Journal int32          `idl:"name:plJournal" json:"journal"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetJournalOperation) OpNum() int { return 19 }

func (o *xxx_GetJournalOperation) OpName() string { return "/IMSMQMessage/v0/Journal" }

func (o *xxx_GetJournalOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetJournalOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetJournalOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetJournalOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetJournalOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plJournal {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.Journal); err != nil {
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

func (o *xxx_GetJournalOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plJournal {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.Journal); err != nil {
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

// GetJournalRequest structure represents the Journal operation request
type GetJournalRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetJournalRequest) xxx_ToOp(ctx context.Context, op *xxx_GetJournalOperation) *xxx_GetJournalOperation {
	if op == nil {
		op = &xxx_GetJournalOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetJournalRequest) xxx_FromOp(ctx context.Context, op *xxx_GetJournalOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetJournalRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetJournalRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetJournalOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetJournalResponse structure represents the Journal operation response
type GetJournalResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Journal int32          `idl:"name:plJournal" json:"journal"`
	// Return: The Journal return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetJournalResponse) xxx_ToOp(ctx context.Context, op *xxx_GetJournalOperation) *xxx_GetJournalOperation {
	if op == nil {
		op = &xxx_GetJournalOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Journal = o.Journal
	op.Return = o.Return
	return op
}

func (o *GetJournalResponse) xxx_FromOp(ctx context.Context, op *xxx_GetJournalOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Journal = op.Journal
	o.Return = op.Return
}
func (o *GetJournalResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetJournalResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetJournalOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetJournalOperation structure represents the Journal operation
type xxx_SetJournalOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Journal int32          `idl:"name:lJournal" json:"journal"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetJournalOperation) OpNum() int { return 20 }

func (o *xxx_SetJournalOperation) OpName() string { return "/IMSMQMessage/v0/Journal" }

func (o *xxx_SetJournalOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetJournalOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lJournal {in} (1:(int32))
	{
		if err := w.WriteData(o.Journal); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetJournalOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lJournal {in} (1:(int32))
	{
		if err := w.ReadData(&o.Journal); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetJournalOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetJournalOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetJournalOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetJournalRequest structure represents the Journal operation request
type SetJournalRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	Journal int32          `idl:"name:lJournal" json:"journal"`
}

func (o *SetJournalRequest) xxx_ToOp(ctx context.Context, op *xxx_SetJournalOperation) *xxx_SetJournalOperation {
	if op == nil {
		op = &xxx_SetJournalOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Journal = o.Journal
	return op
}

func (o *SetJournalRequest) xxx_FromOp(ctx context.Context, op *xxx_SetJournalOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Journal = op.Journal
}
func (o *SetJournalRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetJournalRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetJournalOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetJournalResponse structure represents the Journal operation response
type SetJournalResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Journal return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetJournalResponse) xxx_ToOp(ctx context.Context, op *xxx_SetJournalOperation) *xxx_SetJournalOperation {
	if op == nil {
		op = &xxx_SetJournalOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetJournalResponse) xxx_FromOp(ctx context.Context, op *xxx_SetJournalOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetJournalResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetJournalResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetJournalOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetResponseQueueInfoOperation structure represents the ResponseQueueInfo operation
type xxx_GetResponseQueueInfoOperation struct {
	This     *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Response *mqac.QueueInfo `idl:"name:ppqinfoResponse" json:"response"`
	Return   int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetResponseQueueInfoOperation) OpNum() int { return 21 }

func (o *xxx_GetResponseQueueInfoOperation) OpName() string {
	return "/IMSMQMessage/v0/ResponseQueueInfo"
}

func (o *xxx_GetResponseQueueInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResponseQueueInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetResponseQueueInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetResponseQueueInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResponseQueueInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppqinfoResponse {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQQueueInfo}(interface))
	{
		if o.Response != nil {
			_ptr_ppqinfoResponse := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Response != nil {
					if err := o.Response.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.QueueInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Response, _ptr_ppqinfoResponse); err != nil {
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

func (o *xxx_GetResponseQueueInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppqinfoResponse {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQQueueInfo}(interface))
	{
		_ptr_ppqinfoResponse := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Response == nil {
				o.Response = &mqac.QueueInfo{}
			}
			if err := o.Response.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppqinfoResponse := func(ptr interface{}) { o.Response = *ptr.(**mqac.QueueInfo) }
		if err := w.ReadPointer(&o.Response, _s_ppqinfoResponse, _ptr_ppqinfoResponse); err != nil {
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

// GetResponseQueueInfoRequest structure represents the ResponseQueueInfo operation request
type GetResponseQueueInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetResponseQueueInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_GetResponseQueueInfoOperation) *xxx_GetResponseQueueInfoOperation {
	if op == nil {
		op = &xxx_GetResponseQueueInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetResponseQueueInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetResponseQueueInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetResponseQueueInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetResponseQueueInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetResponseQueueInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetResponseQueueInfoResponse structure represents the ResponseQueueInfo operation response
type GetResponseQueueInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Response *mqac.QueueInfo `idl:"name:ppqinfoResponse" json:"response"`
	// Return: The ResponseQueueInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetResponseQueueInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_GetResponseQueueInfoOperation) *xxx_GetResponseQueueInfoOperation {
	if op == nil {
		op = &xxx_GetResponseQueueInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Response = o.Response
	op.Return = o.Return
	return op
}

func (o *GetResponseQueueInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetResponseQueueInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Response = op.Response
	o.Return = op.Return
}
func (o *GetResponseQueueInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetResponseQueueInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetResponseQueueInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetByRefResponseQueueInfoOperation structure represents the ResponseQueueInfo operation
type xxx_SetByRefResponseQueueInfoOperation struct {
	This              *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat  `idl:"name:That" json:"that"`
	QueueInfoResponse *mqac.QueueInfo `idl:"name:pqinfoResponse" json:"queue_info_response"`
	Return            int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_SetByRefResponseQueueInfoOperation) OpNum() int { return 22 }

func (o *xxx_SetByRefResponseQueueInfoOperation) OpName() string {
	return "/IMSMQMessage/v0/ResponseQueueInfo"
}

func (o *xxx_SetByRefResponseQueueInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetByRefResponseQueueInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pqinfoResponse {in} (1:{pointer=ref}*(1))(2:{alias=IMSMQQueueInfo}(interface))
	{
		if o.QueueInfoResponse != nil {
			_ptr_pqinfoResponse := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.QueueInfoResponse != nil {
					if err := o.QueueInfoResponse.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.QueueInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.QueueInfoResponse, _ptr_pqinfoResponse); err != nil {
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
	return nil
}

func (o *xxx_SetByRefResponseQueueInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pqinfoResponse {in} (1:{pointer=ref}*(1))(2:{alias=IMSMQQueueInfo}(interface))
	{
		_ptr_pqinfoResponse := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.QueueInfoResponse == nil {
				o.QueueInfoResponse = &mqac.QueueInfo{}
			}
			if err := o.QueueInfoResponse.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pqinfoResponse := func(ptr interface{}) { o.QueueInfoResponse = *ptr.(**mqac.QueueInfo) }
		if err := w.ReadPointer(&o.QueueInfoResponse, _s_pqinfoResponse, _ptr_pqinfoResponse); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetByRefResponseQueueInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetByRefResponseQueueInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetByRefResponseQueueInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetByRefResponseQueueInfoRequest structure represents the ResponseQueueInfo operation request
type SetByRefResponseQueueInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This              *dcom.ORPCThis  `idl:"name:This" json:"this"`
	QueueInfoResponse *mqac.QueueInfo `idl:"name:pqinfoResponse" json:"queue_info_response"`
}

func (o *SetByRefResponseQueueInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_SetByRefResponseQueueInfoOperation) *xxx_SetByRefResponseQueueInfoOperation {
	if op == nil {
		op = &xxx_SetByRefResponseQueueInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.QueueInfoResponse = o.QueueInfoResponse
	return op
}

func (o *SetByRefResponseQueueInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_SetByRefResponseQueueInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.QueueInfoResponse = op.QueueInfoResponse
}
func (o *SetByRefResponseQueueInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetByRefResponseQueueInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetByRefResponseQueueInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetByRefResponseQueueInfoResponse structure represents the ResponseQueueInfo operation response
type SetByRefResponseQueueInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ResponseQueueInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetByRefResponseQueueInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_SetByRefResponseQueueInfoOperation) *xxx_SetByRefResponseQueueInfoOperation {
	if op == nil {
		op = &xxx_SetByRefResponseQueueInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetByRefResponseQueueInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_SetByRefResponseQueueInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetByRefResponseQueueInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetByRefResponseQueueInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetByRefResponseQueueInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAppSpecificOperation structure represents the AppSpecific operation
type xxx_GetAppSpecificOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	AppSpecific int32          `idl:"name:plAppSpecific" json:"app_specific"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAppSpecificOperation) OpNum() int { return 23 }

func (o *xxx_GetAppSpecificOperation) OpName() string { return "/IMSMQMessage/v0/AppSpecific" }

func (o *xxx_GetAppSpecificOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAppSpecificOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetAppSpecificOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetAppSpecificOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAppSpecificOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plAppSpecific {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.AppSpecific); err != nil {
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

func (o *xxx_GetAppSpecificOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plAppSpecific {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.AppSpecific); err != nil {
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

// GetAppSpecificRequest structure represents the AppSpecific operation request
type GetAppSpecificRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetAppSpecificRequest) xxx_ToOp(ctx context.Context, op *xxx_GetAppSpecificOperation) *xxx_GetAppSpecificOperation {
	if op == nil {
		op = &xxx_GetAppSpecificOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetAppSpecificRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAppSpecificOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetAppSpecificRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetAppSpecificRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAppSpecificOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAppSpecificResponse structure represents the AppSpecific operation response
type GetAppSpecificResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	AppSpecific int32          `idl:"name:plAppSpecific" json:"app_specific"`
	// Return: The AppSpecific return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAppSpecificResponse) xxx_ToOp(ctx context.Context, op *xxx_GetAppSpecificOperation) *xxx_GetAppSpecificOperation {
	if op == nil {
		op = &xxx_GetAppSpecificOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.AppSpecific = o.AppSpecific
	op.Return = o.Return
	return op
}

func (o *GetAppSpecificResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAppSpecificOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.AppSpecific = op.AppSpecific
	o.Return = op.Return
}
func (o *GetAppSpecificResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetAppSpecificResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAppSpecificOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetAppSpecificOperation structure represents the AppSpecific operation
type xxx_SetAppSpecificOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	AppSpecific int32          `idl:"name:lAppSpecific" json:"app_specific"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetAppSpecificOperation) OpNum() int { return 24 }

func (o *xxx_SetAppSpecificOperation) OpName() string { return "/IMSMQMessage/v0/AppSpecific" }

func (o *xxx_SetAppSpecificOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAppSpecificOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lAppSpecific {in} (1:(int32))
	{
		if err := w.WriteData(o.AppSpecific); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAppSpecificOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lAppSpecific {in} (1:(int32))
	{
		if err := w.ReadData(&o.AppSpecific); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAppSpecificOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAppSpecificOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetAppSpecificOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetAppSpecificRequest structure represents the AppSpecific operation request
type SetAppSpecificRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	AppSpecific int32          `idl:"name:lAppSpecific" json:"app_specific"`
}

func (o *SetAppSpecificRequest) xxx_ToOp(ctx context.Context, op *xxx_SetAppSpecificOperation) *xxx_SetAppSpecificOperation {
	if op == nil {
		op = &xxx_SetAppSpecificOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.AppSpecific = o.AppSpecific
	return op
}

func (o *SetAppSpecificRequest) xxx_FromOp(ctx context.Context, op *xxx_SetAppSpecificOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.AppSpecific = op.AppSpecific
}
func (o *SetAppSpecificRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetAppSpecificRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAppSpecificOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetAppSpecificResponse structure represents the AppSpecific operation response
type SetAppSpecificResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AppSpecific return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetAppSpecificResponse) xxx_ToOp(ctx context.Context, op *xxx_SetAppSpecificOperation) *xxx_SetAppSpecificOperation {
	if op == nil {
		op = &xxx_SetAppSpecificOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetAppSpecificResponse) xxx_FromOp(ctx context.Context, op *xxx_SetAppSpecificOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetAppSpecificResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetAppSpecificResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAppSpecificOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSourceMachineGUIDOperation structure represents the SourceMachineGuid operation
type xxx_GetSourceMachineGUIDOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	SourceMachine *oaut.String   `idl:"name:pbstrGuidSrcMachine" json:"source_machine"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSourceMachineGUIDOperation) OpNum() int { return 25 }

func (o *xxx_GetSourceMachineGUIDOperation) OpName() string {
	return "/IMSMQMessage/v0/SourceMachineGuid"
}

func (o *xxx_GetSourceMachineGUIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSourceMachineGUIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSourceMachineGUIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSourceMachineGUIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSourceMachineGUIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrGuidSrcMachine {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SourceMachine != nil {
			_ptr_pbstrGuidSrcMachine := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SourceMachine != nil {
					if err := o.SourceMachine.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SourceMachine, _ptr_pbstrGuidSrcMachine); err != nil {
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

func (o *xxx_GetSourceMachineGUIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrGuidSrcMachine {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrGuidSrcMachine := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SourceMachine == nil {
				o.SourceMachine = &oaut.String{}
			}
			if err := o.SourceMachine.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrGuidSrcMachine := func(ptr interface{}) { o.SourceMachine = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SourceMachine, _s_pbstrGuidSrcMachine, _ptr_pbstrGuidSrcMachine); err != nil {
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

// GetSourceMachineGUIDRequest structure represents the SourceMachineGuid operation request
type GetSourceMachineGUIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSourceMachineGUIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSourceMachineGUIDOperation) *xxx_GetSourceMachineGUIDOperation {
	if op == nil {
		op = &xxx_GetSourceMachineGUIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetSourceMachineGUIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSourceMachineGUIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSourceMachineGUIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSourceMachineGUIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSourceMachineGUIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSourceMachineGUIDResponse structure represents the SourceMachineGuid operation response
type GetSourceMachineGUIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbstrGuidSrcMachine: A pointer to a BSTR that contains the GUID identifier of the
	// computer that sent the message. The string MUST adhere to the following format, specified
	// using ABNF.
	SourceMachine *oaut.String `idl:"name:pbstrGuidSrcMachine" json:"source_machine"`
	// Return: The SourceMachineGuid return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSourceMachineGUIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSourceMachineGUIDOperation) *xxx_GetSourceMachineGUIDOperation {
	if op == nil {
		op = &xxx_GetSourceMachineGUIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.SourceMachine = o.SourceMachine
	op.Return = o.Return
	return op
}

func (o *GetSourceMachineGUIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSourceMachineGUIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SourceMachine = op.SourceMachine
	o.Return = op.Return
}
func (o *GetSourceMachineGUIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSourceMachineGUIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSourceMachineGUIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetBodyLengthOperation structure represents the BodyLength operation
type xxx_GetBodyLengthOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	BodyLength int32          `idl:"name:pcbBody" json:"body_length"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetBodyLengthOperation) OpNum() int { return 26 }

func (o *xxx_GetBodyLengthOperation) OpName() string { return "/IMSMQMessage/v0/BodyLength" }

func (o *xxx_GetBodyLengthOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBodyLengthOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetBodyLengthOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetBodyLengthOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBodyLengthOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pcbBody {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.BodyLength); err != nil {
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

func (o *xxx_GetBodyLengthOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pcbBody {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.BodyLength); err != nil {
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

// GetBodyLengthRequest structure represents the BodyLength operation request
type GetBodyLengthRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetBodyLengthRequest) xxx_ToOp(ctx context.Context, op *xxx_GetBodyLengthOperation) *xxx_GetBodyLengthOperation {
	if op == nil {
		op = &xxx_GetBodyLengthOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetBodyLengthRequest) xxx_FromOp(ctx context.Context, op *xxx_GetBodyLengthOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetBodyLengthRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetBodyLengthRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetBodyLengthOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetBodyLengthResponse structure represents the BodyLength operation response
type GetBodyLengthResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pcbBody: A pointer to a long integer that contains the number of bytes in the body
	// of the message.
	BodyLength int32 `idl:"name:pcbBody" json:"body_length"`
	// Return: The BodyLength return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetBodyLengthResponse) xxx_ToOp(ctx context.Context, op *xxx_GetBodyLengthOperation) *xxx_GetBodyLengthOperation {
	if op == nil {
		op = &xxx_GetBodyLengthOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.BodyLength = o.BodyLength
	op.Return = o.Return
	return op
}

func (o *GetBodyLengthResponse) xxx_FromOp(ctx context.Context, op *xxx_GetBodyLengthOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.BodyLength = op.BodyLength
	o.Return = op.Return
}
func (o *GetBodyLengthResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetBodyLengthResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetBodyLengthOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetBodyOperation structure represents the Body operation
type xxx_GetBodyOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Body   *oaut.Variant  `idl:"name:pvarBody" json:"body"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetBodyOperation) OpNum() int { return 27 }

func (o *xxx_GetBodyOperation) OpName() string { return "/IMSMQMessage/v0/Body" }

func (o *xxx_GetBodyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBodyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetBodyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetBodyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBodyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pvarBody {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Body != nil {
			_ptr_pvarBody := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Body != nil {
					if err := o.Body.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Body, _ptr_pvarBody); err != nil {
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

func (o *xxx_GetBodyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pvarBody {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pvarBody := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Body == nil {
				o.Body = &oaut.Variant{}
			}
			if err := o.Body.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pvarBody := func(ptr interface{}) { o.Body = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Body, _s_pvarBody, _ptr_pvarBody); err != nil {
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

// GetBodyRequest structure represents the Body operation request
type GetBodyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetBodyRequest) xxx_ToOp(ctx context.Context, op *xxx_GetBodyOperation) *xxx_GetBodyOperation {
	if op == nil {
		op = &xxx_GetBodyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetBodyRequest) xxx_FromOp(ctx context.Context, op *xxx_GetBodyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetBodyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetBodyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetBodyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetBodyResponse structure represents the Body operation response
type GetBodyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	Body *oaut.Variant  `idl:"name:pvarBody" json:"body"`
	// Return: The Body return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetBodyResponse) xxx_ToOp(ctx context.Context, op *xxx_GetBodyOperation) *xxx_GetBodyOperation {
	if op == nil {
		op = &xxx_GetBodyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Body = o.Body
	op.Return = o.Return
	return op
}

func (o *GetBodyResponse) xxx_FromOp(ctx context.Context, op *xxx_GetBodyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Body = op.Body
	o.Return = op.Return
}
func (o *GetBodyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetBodyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetBodyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetBodyOperation structure represents the Body operation
type xxx_SetBodyOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Body   *oaut.Variant  `idl:"name:varBody" json:"body"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetBodyOperation) OpNum() int { return 28 }

func (o *xxx_SetBodyOperation) OpName() string { return "/IMSMQMessage/v0/Body" }

func (o *xxx_SetBodyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetBodyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// varBody {in} (1:{alias=VARIANT}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.Body != nil {
			if err := o.Body.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetBodyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// varBody {in} (1:{alias=VARIANT,pointer=ref}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.Body == nil {
			o.Body = &oaut.Variant{}
		}
		if err := o.Body.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetBodyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetBodyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetBodyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetBodyRequest structure represents the Body operation request
type SetBodyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	Body *oaut.Variant  `idl:"name:varBody" json:"body"`
}

func (o *SetBodyRequest) xxx_ToOp(ctx context.Context, op *xxx_SetBodyOperation) *xxx_SetBodyOperation {
	if op == nil {
		op = &xxx_SetBodyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Body = o.Body
	return op
}

func (o *SetBodyRequest) xxx_FromOp(ctx context.Context, op *xxx_SetBodyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Body = op.Body
}
func (o *SetBodyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetBodyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetBodyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetBodyResponse structure represents the Body operation response
type SetBodyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Body return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetBodyResponse) xxx_ToOp(ctx context.Context, op *xxx_SetBodyOperation) *xxx_SetBodyOperation {
	if op == nil {
		op = &xxx_SetBodyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetBodyResponse) xxx_FromOp(ctx context.Context, op *xxx_SetBodyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetBodyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetBodyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetBodyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAdminQueueInfoOperation structure represents the AdminQueueInfo operation
type xxx_GetAdminQueueInfoOperation struct {
	This   *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Admin  *mqac.QueueInfo `idl:"name:ppqinfoAdmin" json:"admin"`
	Return int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAdminQueueInfoOperation) OpNum() int { return 29 }

func (o *xxx_GetAdminQueueInfoOperation) OpName() string { return "/IMSMQMessage/v0/AdminQueueInfo" }

func (o *xxx_GetAdminQueueInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAdminQueueInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetAdminQueueInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetAdminQueueInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAdminQueueInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppqinfoAdmin {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQQueueInfo}(interface))
	{
		if o.Admin != nil {
			_ptr_ppqinfoAdmin := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Admin != nil {
					if err := o.Admin.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.QueueInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Admin, _ptr_ppqinfoAdmin); err != nil {
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

func (o *xxx_GetAdminQueueInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppqinfoAdmin {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQQueueInfo}(interface))
	{
		_ptr_ppqinfoAdmin := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Admin == nil {
				o.Admin = &mqac.QueueInfo{}
			}
			if err := o.Admin.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppqinfoAdmin := func(ptr interface{}) { o.Admin = *ptr.(**mqac.QueueInfo) }
		if err := w.ReadPointer(&o.Admin, _s_ppqinfoAdmin, _ptr_ppqinfoAdmin); err != nil {
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

// GetAdminQueueInfoRequest structure represents the AdminQueueInfo operation request
type GetAdminQueueInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetAdminQueueInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_GetAdminQueueInfoOperation) *xxx_GetAdminQueueInfoOperation {
	if op == nil {
		op = &xxx_GetAdminQueueInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetAdminQueueInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAdminQueueInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetAdminQueueInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetAdminQueueInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAdminQueueInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAdminQueueInfoResponse structure represents the AdminQueueInfo operation response
type GetAdminQueueInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Admin *mqac.QueueInfo `idl:"name:ppqinfoAdmin" json:"admin"`
	// Return: The AdminQueueInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAdminQueueInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_GetAdminQueueInfoOperation) *xxx_GetAdminQueueInfoOperation {
	if op == nil {
		op = &xxx_GetAdminQueueInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Admin = o.Admin
	op.Return = o.Return
	return op
}

func (o *GetAdminQueueInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAdminQueueInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Admin = op.Admin
	o.Return = op.Return
}
func (o *GetAdminQueueInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetAdminQueueInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAdminQueueInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetByRefAdminQueueInfoOperation structure represents the AdminQueueInfo operation
type xxx_SetByRefAdminQueueInfoOperation struct {
	This           *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat  `idl:"name:That" json:"that"`
	QueueInfoAdmin *mqac.QueueInfo `idl:"name:pqinfoAdmin" json:"queue_info_admin"`
	Return         int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_SetByRefAdminQueueInfoOperation) OpNum() int { return 30 }

func (o *xxx_SetByRefAdminQueueInfoOperation) OpName() string {
	return "/IMSMQMessage/v0/AdminQueueInfo"
}

func (o *xxx_SetByRefAdminQueueInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetByRefAdminQueueInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pqinfoAdmin {in} (1:{pointer=ref}*(1))(2:{alias=IMSMQQueueInfo}(interface))
	{
		if o.QueueInfoAdmin != nil {
			_ptr_pqinfoAdmin := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.QueueInfoAdmin != nil {
					if err := o.QueueInfoAdmin.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.QueueInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.QueueInfoAdmin, _ptr_pqinfoAdmin); err != nil {
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
	return nil
}

func (o *xxx_SetByRefAdminQueueInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pqinfoAdmin {in} (1:{pointer=ref}*(1))(2:{alias=IMSMQQueueInfo}(interface))
	{
		_ptr_pqinfoAdmin := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.QueueInfoAdmin == nil {
				o.QueueInfoAdmin = &mqac.QueueInfo{}
			}
			if err := o.QueueInfoAdmin.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pqinfoAdmin := func(ptr interface{}) { o.QueueInfoAdmin = *ptr.(**mqac.QueueInfo) }
		if err := w.ReadPointer(&o.QueueInfoAdmin, _s_pqinfoAdmin, _ptr_pqinfoAdmin); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetByRefAdminQueueInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetByRefAdminQueueInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetByRefAdminQueueInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetByRefAdminQueueInfoRequest structure represents the AdminQueueInfo operation request
type SetByRefAdminQueueInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This           *dcom.ORPCThis  `idl:"name:This" json:"this"`
	QueueInfoAdmin *mqac.QueueInfo `idl:"name:pqinfoAdmin" json:"queue_info_admin"`
}

func (o *SetByRefAdminQueueInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_SetByRefAdminQueueInfoOperation) *xxx_SetByRefAdminQueueInfoOperation {
	if op == nil {
		op = &xxx_SetByRefAdminQueueInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.QueueInfoAdmin = o.QueueInfoAdmin
	return op
}

func (o *SetByRefAdminQueueInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_SetByRefAdminQueueInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.QueueInfoAdmin = op.QueueInfoAdmin
}
func (o *SetByRefAdminQueueInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetByRefAdminQueueInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetByRefAdminQueueInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetByRefAdminQueueInfoResponse structure represents the AdminQueueInfo operation response
type SetByRefAdminQueueInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AdminQueueInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetByRefAdminQueueInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_SetByRefAdminQueueInfoOperation) *xxx_SetByRefAdminQueueInfoOperation {
	if op == nil {
		op = &xxx_SetByRefAdminQueueInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetByRefAdminQueueInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_SetByRefAdminQueueInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetByRefAdminQueueInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetByRefAdminQueueInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetByRefAdminQueueInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetIDOperation structure represents the Id operation
type xxx_GetIDOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	MessageID *oaut.Variant  `idl:"name:pvarMsgId" json:"message_id"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetIDOperation) OpNum() int { return 31 }

func (o *xxx_GetIDOperation) OpName() string { return "/IMSMQMessage/v0/Id" }

func (o *xxx_GetIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pvarMsgId {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.MessageID != nil {
			_ptr_pvarMsgId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MessageID != nil {
					if err := o.MessageID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MessageID, _ptr_pvarMsgId); err != nil {
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

func (o *xxx_GetIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pvarMsgId {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pvarMsgId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MessageID == nil {
				o.MessageID = &oaut.Variant{}
			}
			if err := o.MessageID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pvarMsgId := func(ptr interface{}) { o.MessageID = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.MessageID, _s_pvarMsgId, _ptr_pvarMsgId); err != nil {
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

// GetIDRequest structure represents the Id operation request
type GetIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetIDOperation) *xxx_GetIDOperation {
	if op == nil {
		op = &xxx_GetIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetIDResponse structure represents the Id operation response
type GetIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pvarMsgId: A pointer to a VARIANT that contains a 20-byte array (VT_ARRAY | VT_UI1) that contains the unique message identifier.
	MessageID *oaut.Variant `idl:"name:pvarMsgId" json:"message_id"`
	// Return: The Id return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetIDOperation) *xxx_GetIDOperation {
	if op == nil {
		op = &xxx_GetIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.MessageID = o.MessageID
	op.Return = o.Return
	return op
}

func (o *GetIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MessageID = op.MessageID
	o.Return = op.Return
}
func (o *GetIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetCorrelationIDOperation structure represents the CorrelationId operation
type xxx_GetCorrelationIDOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	MessageID *oaut.Variant  `idl:"name:pvarMsgId" json:"message_id"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCorrelationIDOperation) OpNum() int { return 32 }

func (o *xxx_GetCorrelationIDOperation) OpName() string { return "/IMSMQMessage/v0/CorrelationId" }

func (o *xxx_GetCorrelationIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCorrelationIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetCorrelationIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetCorrelationIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCorrelationIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pvarMsgId {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.MessageID != nil {
			_ptr_pvarMsgId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MessageID != nil {
					if err := o.MessageID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MessageID, _ptr_pvarMsgId); err != nil {
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

func (o *xxx_GetCorrelationIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pvarMsgId {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pvarMsgId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MessageID == nil {
				o.MessageID = &oaut.Variant{}
			}
			if err := o.MessageID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pvarMsgId := func(ptr interface{}) { o.MessageID = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.MessageID, _s_pvarMsgId, _ptr_pvarMsgId); err != nil {
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

// GetCorrelationIDRequest structure represents the CorrelationId operation request
type GetCorrelationIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetCorrelationIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetCorrelationIDOperation) *xxx_GetCorrelationIDOperation {
	if op == nil {
		op = &xxx_GetCorrelationIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetCorrelationIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCorrelationIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetCorrelationIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetCorrelationIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCorrelationIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCorrelationIDResponse structure represents the CorrelationId operation response
type GetCorrelationIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	MessageID *oaut.Variant  `idl:"name:pvarMsgId" json:"message_id"`
	// Return: The CorrelationId return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCorrelationIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetCorrelationIDOperation) *xxx_GetCorrelationIDOperation {
	if op == nil {
		op = &xxx_GetCorrelationIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.MessageID = o.MessageID
	op.Return = o.Return
	return op
}

func (o *GetCorrelationIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCorrelationIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MessageID = op.MessageID
	o.Return = op.Return
}
func (o *GetCorrelationIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetCorrelationIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCorrelationIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetCorrelationIDOperation structure represents the CorrelationId operation
type xxx_SetCorrelationIDOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	MessageID *oaut.Variant  `idl:"name:varMsgId" json:"message_id"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetCorrelationIDOperation) OpNum() int { return 33 }

func (o *xxx_SetCorrelationIDOperation) OpName() string { return "/IMSMQMessage/v0/CorrelationId" }

func (o *xxx_SetCorrelationIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCorrelationIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// varMsgId {in} (1:{alias=VARIANT}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.MessageID != nil {
			if err := o.MessageID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCorrelationIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// varMsgId {in} (1:{alias=VARIANT,pointer=ref}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.MessageID == nil {
			o.MessageID = &oaut.Variant{}
		}
		if err := o.MessageID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCorrelationIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCorrelationIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetCorrelationIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetCorrelationIDRequest structure represents the CorrelationId operation request
type SetCorrelationIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	MessageID *oaut.Variant  `idl:"name:varMsgId" json:"message_id"`
}

func (o *SetCorrelationIDRequest) xxx_ToOp(ctx context.Context, op *xxx_SetCorrelationIDOperation) *xxx_SetCorrelationIDOperation {
	if op == nil {
		op = &xxx_SetCorrelationIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MessageID = o.MessageID
	return op
}

func (o *SetCorrelationIDRequest) xxx_FromOp(ctx context.Context, op *xxx_SetCorrelationIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MessageID = op.MessageID
}
func (o *SetCorrelationIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetCorrelationIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetCorrelationIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetCorrelationIDResponse structure represents the CorrelationId operation response
type SetCorrelationIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CorrelationId return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetCorrelationIDResponse) xxx_ToOp(ctx context.Context, op *xxx_SetCorrelationIDOperation) *xxx_SetCorrelationIDOperation {
	if op == nil {
		op = &xxx_SetCorrelationIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetCorrelationIDResponse) xxx_FromOp(ctx context.Context, op *xxx_SetCorrelationIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetCorrelationIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetCorrelationIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetCorrelationIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAckOperation structure represents the Ack operation
type xxx_GetAckOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Ack    int32          `idl:"name:plAck" json:"ack"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAckOperation) OpNum() int { return 34 }

func (o *xxx_GetAckOperation) OpName() string { return "/IMSMQMessage/v0/Ack" }

func (o *xxx_GetAckOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAckOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetAckOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetAckOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAckOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plAck {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.Ack); err != nil {
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

func (o *xxx_GetAckOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plAck {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.Ack); err != nil {
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

// GetAckRequest structure represents the Ack operation request
type GetAckRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetAckRequest) xxx_ToOp(ctx context.Context, op *xxx_GetAckOperation) *xxx_GetAckOperation {
	if op == nil {
		op = &xxx_GetAckOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetAckRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAckOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetAckRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetAckRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAckOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAckResponse structure represents the Ack operation response
type GetAckResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	Ack  int32          `idl:"name:plAck" json:"ack"`
	// Return: The Ack return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAckResponse) xxx_ToOp(ctx context.Context, op *xxx_GetAckOperation) *xxx_GetAckOperation {
	if op == nil {
		op = &xxx_GetAckOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Ack = o.Ack
	op.Return = o.Return
	return op
}

func (o *GetAckResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAckOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Ack = op.Ack
	o.Return = op.Return
}
func (o *GetAckResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetAckResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAckOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetAckOperation structure represents the Ack operation
type xxx_SetAckOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Ack    int32          `idl:"name:lAck" json:"ack"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetAckOperation) OpNum() int { return 35 }

func (o *xxx_SetAckOperation) OpName() string { return "/IMSMQMessage/v0/Ack" }

func (o *xxx_SetAckOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAckOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lAck {in} (1:(int32))
	{
		if err := w.WriteData(o.Ack); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAckOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lAck {in} (1:(int32))
	{
		if err := w.ReadData(&o.Ack); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAckOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAckOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetAckOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetAckRequest structure represents the Ack operation request
type SetAckRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	Ack  int32          `idl:"name:lAck" json:"ack"`
}

func (o *SetAckRequest) xxx_ToOp(ctx context.Context, op *xxx_SetAckOperation) *xxx_SetAckOperation {
	if op == nil {
		op = &xxx_SetAckOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Ack = o.Ack
	return op
}

func (o *SetAckRequest) xxx_FromOp(ctx context.Context, op *xxx_SetAckOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Ack = op.Ack
}
func (o *SetAckRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetAckRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAckOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetAckResponse structure represents the Ack operation response
type SetAckResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Ack return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetAckResponse) xxx_ToOp(ctx context.Context, op *xxx_SetAckOperation) *xxx_SetAckOperation {
	if op == nil {
		op = &xxx_SetAckOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetAckResponse) xxx_FromOp(ctx context.Context, op *xxx_SetAckOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetAckResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetAckResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAckOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetLabelOperation structure represents the Label operation
type xxx_GetLabelOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Label  *oaut.String   `idl:"name:pbstrLabel" json:"label"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetLabelOperation) OpNum() int { return 36 }

func (o *xxx_GetLabelOperation) OpName() string { return "/IMSMQMessage/v0/Label" }

func (o *xxx_GetLabelOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLabelOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetLabelOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetLabelOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLabelOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrLabel {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Label != nil {
			_ptr_pbstrLabel := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Label != nil {
					if err := o.Label.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Label, _ptr_pbstrLabel); err != nil {
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

func (o *xxx_GetLabelOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrLabel {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrLabel := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Label == nil {
				o.Label = &oaut.String{}
			}
			if err := o.Label.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrLabel := func(ptr interface{}) { o.Label = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Label, _s_pbstrLabel, _ptr_pbstrLabel); err != nil {
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

// GetLabelRequest structure represents the Label operation request
type GetLabelRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetLabelRequest) xxx_ToOp(ctx context.Context, op *xxx_GetLabelOperation) *xxx_GetLabelOperation {
	if op == nil {
		op = &xxx_GetLabelOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetLabelRequest) xxx_FromOp(ctx context.Context, op *xxx_GetLabelOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetLabelRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetLabelRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLabelOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetLabelResponse structure represents the Label operation response
type GetLabelResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat `idl:"name:That" json:"that"`
	Label *oaut.String   `idl:"name:pbstrLabel" json:"label"`
	// Return: The Label return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetLabelResponse) xxx_ToOp(ctx context.Context, op *xxx_GetLabelOperation) *xxx_GetLabelOperation {
	if op == nil {
		op = &xxx_GetLabelOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Label = o.Label
	op.Return = o.Return
	return op
}

func (o *GetLabelResponse) xxx_FromOp(ctx context.Context, op *xxx_GetLabelOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Label = op.Label
	o.Return = op.Return
}
func (o *GetLabelResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetLabelResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLabelOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetLabelOperation structure represents the Label operation
type xxx_SetLabelOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Label  *oaut.String   `idl:"name:bstrLabel" json:"label"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetLabelOperation) OpNum() int { return 37 }

func (o *xxx_SetLabelOperation) OpName() string { return "/IMSMQMessage/v0/Label" }

func (o *xxx_SetLabelOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLabelOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrLabel {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Label != nil {
			_ptr_bstrLabel := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Label != nil {
					if err := o.Label.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Label, _ptr_bstrLabel); err != nil {
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
	return nil
}

func (o *xxx_SetLabelOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrLabel {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrLabel := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Label == nil {
				o.Label = &oaut.String{}
			}
			if err := o.Label.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrLabel := func(ptr interface{}) { o.Label = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Label, _s_bstrLabel, _ptr_bstrLabel); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLabelOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLabelOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetLabelOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetLabelRequest structure represents the Label operation request
type SetLabelRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This  *dcom.ORPCThis `idl:"name:This" json:"this"`
	Label *oaut.String   `idl:"name:bstrLabel" json:"label"`
}

func (o *SetLabelRequest) xxx_ToOp(ctx context.Context, op *xxx_SetLabelOperation) *xxx_SetLabelOperation {
	if op == nil {
		op = &xxx_SetLabelOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Label = o.Label
	return op
}

func (o *SetLabelRequest) xxx_FromOp(ctx context.Context, op *xxx_SetLabelOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Label = op.Label
}
func (o *SetLabelRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetLabelRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetLabelOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetLabelResponse structure represents the Label operation response
type SetLabelResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Label return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetLabelResponse) xxx_ToOp(ctx context.Context, op *xxx_SetLabelOperation) *xxx_SetLabelOperation {
	if op == nil {
		op = &xxx_SetLabelOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetLabelResponse) xxx_FromOp(ctx context.Context, op *xxx_SetLabelOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetLabelResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetLabelResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetLabelOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMaxTimeToReachQueueOperation structure represents the MaxTimeToReachQueue operation
type xxx_GetMaxTimeToReachQueueOperation struct {
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	MaxTimeToReachQueue int32          `idl:"name:plMaxTimeToReachQueue" json:"max_time_to_reach_queue"`
	Return              int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMaxTimeToReachQueueOperation) OpNum() int { return 38 }

func (o *xxx_GetMaxTimeToReachQueueOperation) OpName() string {
	return "/IMSMQMessage/v0/MaxTimeToReachQueue"
}

func (o *xxx_GetMaxTimeToReachQueueOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMaxTimeToReachQueueOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetMaxTimeToReachQueueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetMaxTimeToReachQueueOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMaxTimeToReachQueueOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plMaxTimeToReachQueue {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.MaxTimeToReachQueue); err != nil {
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

func (o *xxx_GetMaxTimeToReachQueueOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plMaxTimeToReachQueue {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.MaxTimeToReachQueue); err != nil {
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

// GetMaxTimeToReachQueueRequest structure represents the MaxTimeToReachQueue operation request
type GetMaxTimeToReachQueueRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetMaxTimeToReachQueueRequest) xxx_ToOp(ctx context.Context, op *xxx_GetMaxTimeToReachQueueOperation) *xxx_GetMaxTimeToReachQueueOperation {
	if op == nil {
		op = &xxx_GetMaxTimeToReachQueueOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetMaxTimeToReachQueueRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMaxTimeToReachQueueOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetMaxTimeToReachQueueRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetMaxTimeToReachQueueRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMaxTimeToReachQueueOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMaxTimeToReachQueueResponse structure represents the MaxTimeToReachQueue operation response
type GetMaxTimeToReachQueueResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	MaxTimeToReachQueue int32          `idl:"name:plMaxTimeToReachQueue" json:"max_time_to_reach_queue"`
	// Return: The MaxTimeToReachQueue return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMaxTimeToReachQueueResponse) xxx_ToOp(ctx context.Context, op *xxx_GetMaxTimeToReachQueueOperation) *xxx_GetMaxTimeToReachQueueOperation {
	if op == nil {
		op = &xxx_GetMaxTimeToReachQueueOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.MaxTimeToReachQueue = o.MaxTimeToReachQueue
	op.Return = o.Return
	return op
}

func (o *GetMaxTimeToReachQueueResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMaxTimeToReachQueueOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MaxTimeToReachQueue = op.MaxTimeToReachQueue
	o.Return = op.Return
}
func (o *GetMaxTimeToReachQueueResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetMaxTimeToReachQueueResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMaxTimeToReachQueueOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetMaxTimeToReachQueueOperation structure represents the MaxTimeToReachQueue operation
type xxx_SetMaxTimeToReachQueueOperation struct {
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	MaxTimeToReachQueue int32          `idl:"name:lMaxTimeToReachQueue" json:"max_time_to_reach_queue"`
	Return              int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetMaxTimeToReachQueueOperation) OpNum() int { return 39 }

func (o *xxx_SetMaxTimeToReachQueueOperation) OpName() string {
	return "/IMSMQMessage/v0/MaxTimeToReachQueue"
}

func (o *xxx_SetMaxTimeToReachQueueOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMaxTimeToReachQueueOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lMaxTimeToReachQueue {in} (1:(int32))
	{
		if err := w.WriteData(o.MaxTimeToReachQueue); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMaxTimeToReachQueueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lMaxTimeToReachQueue {in} (1:(int32))
	{
		if err := w.ReadData(&o.MaxTimeToReachQueue); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMaxTimeToReachQueueOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMaxTimeToReachQueueOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetMaxTimeToReachQueueOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetMaxTimeToReachQueueRequest structure represents the MaxTimeToReachQueue operation request
type SetMaxTimeToReachQueueRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	MaxTimeToReachQueue int32          `idl:"name:lMaxTimeToReachQueue" json:"max_time_to_reach_queue"`
}

func (o *SetMaxTimeToReachQueueRequest) xxx_ToOp(ctx context.Context, op *xxx_SetMaxTimeToReachQueueOperation) *xxx_SetMaxTimeToReachQueueOperation {
	if op == nil {
		op = &xxx_SetMaxTimeToReachQueueOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MaxTimeToReachQueue = o.MaxTimeToReachQueue
	return op
}

func (o *SetMaxTimeToReachQueueRequest) xxx_FromOp(ctx context.Context, op *xxx_SetMaxTimeToReachQueueOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MaxTimeToReachQueue = op.MaxTimeToReachQueue
}
func (o *SetMaxTimeToReachQueueRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetMaxTimeToReachQueueRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMaxTimeToReachQueueOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetMaxTimeToReachQueueResponse structure represents the MaxTimeToReachQueue operation response
type SetMaxTimeToReachQueueResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The MaxTimeToReachQueue return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetMaxTimeToReachQueueResponse) xxx_ToOp(ctx context.Context, op *xxx_SetMaxTimeToReachQueueOperation) *xxx_SetMaxTimeToReachQueueOperation {
	if op == nil {
		op = &xxx_SetMaxTimeToReachQueueOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetMaxTimeToReachQueueResponse) xxx_FromOp(ctx context.Context, op *xxx_SetMaxTimeToReachQueueOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetMaxTimeToReachQueueResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetMaxTimeToReachQueueResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMaxTimeToReachQueueOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMaxTimeToReceiveOperation structure represents the MaxTimeToReceive operation
type xxx_GetMaxTimeToReceiveOperation struct {
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	MaxTimeToReceive int32          `idl:"name:plMaxTimeToReceive" json:"max_time_to_receive"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMaxTimeToReceiveOperation) OpNum() int { return 40 }

func (o *xxx_GetMaxTimeToReceiveOperation) OpName() string {
	return "/IMSMQMessage/v0/MaxTimeToReceive"
}

func (o *xxx_GetMaxTimeToReceiveOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMaxTimeToReceiveOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetMaxTimeToReceiveOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetMaxTimeToReceiveOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMaxTimeToReceiveOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plMaxTimeToReceive {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.MaxTimeToReceive); err != nil {
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

func (o *xxx_GetMaxTimeToReceiveOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plMaxTimeToReceive {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.MaxTimeToReceive); err != nil {
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

// GetMaxTimeToReceiveRequest structure represents the MaxTimeToReceive operation request
type GetMaxTimeToReceiveRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetMaxTimeToReceiveRequest) xxx_ToOp(ctx context.Context, op *xxx_GetMaxTimeToReceiveOperation) *xxx_GetMaxTimeToReceiveOperation {
	if op == nil {
		op = &xxx_GetMaxTimeToReceiveOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetMaxTimeToReceiveRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMaxTimeToReceiveOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetMaxTimeToReceiveRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetMaxTimeToReceiveRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMaxTimeToReceiveOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMaxTimeToReceiveResponse structure represents the MaxTimeToReceive operation response
type GetMaxTimeToReceiveResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	MaxTimeToReceive int32          `idl:"name:plMaxTimeToReceive" json:"max_time_to_receive"`
	// Return: The MaxTimeToReceive return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMaxTimeToReceiveResponse) xxx_ToOp(ctx context.Context, op *xxx_GetMaxTimeToReceiveOperation) *xxx_GetMaxTimeToReceiveOperation {
	if op == nil {
		op = &xxx_GetMaxTimeToReceiveOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.MaxTimeToReceive = o.MaxTimeToReceive
	op.Return = o.Return
	return op
}

func (o *GetMaxTimeToReceiveResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMaxTimeToReceiveOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MaxTimeToReceive = op.MaxTimeToReceive
	o.Return = op.Return
}
func (o *GetMaxTimeToReceiveResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetMaxTimeToReceiveResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMaxTimeToReceiveOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetMaxTimeToReceiveOperation structure represents the MaxTimeToReceive operation
type xxx_SetMaxTimeToReceiveOperation struct {
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	MaxTimeToReceive int32          `idl:"name:lMaxTimeToReceive" json:"max_time_to_receive"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetMaxTimeToReceiveOperation) OpNum() int { return 41 }

func (o *xxx_SetMaxTimeToReceiveOperation) OpName() string {
	return "/IMSMQMessage/v0/MaxTimeToReceive"
}

func (o *xxx_SetMaxTimeToReceiveOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMaxTimeToReceiveOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lMaxTimeToReceive {in} (1:(int32))
	{
		if err := w.WriteData(o.MaxTimeToReceive); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMaxTimeToReceiveOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lMaxTimeToReceive {in} (1:(int32))
	{
		if err := w.ReadData(&o.MaxTimeToReceive); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMaxTimeToReceiveOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMaxTimeToReceiveOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetMaxTimeToReceiveOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetMaxTimeToReceiveRequest structure represents the MaxTimeToReceive operation request
type SetMaxTimeToReceiveRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	MaxTimeToReceive int32          `idl:"name:lMaxTimeToReceive" json:"max_time_to_receive"`
}

func (o *SetMaxTimeToReceiveRequest) xxx_ToOp(ctx context.Context, op *xxx_SetMaxTimeToReceiveOperation) *xxx_SetMaxTimeToReceiveOperation {
	if op == nil {
		op = &xxx_SetMaxTimeToReceiveOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MaxTimeToReceive = o.MaxTimeToReceive
	return op
}

func (o *SetMaxTimeToReceiveRequest) xxx_FromOp(ctx context.Context, op *xxx_SetMaxTimeToReceiveOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MaxTimeToReceive = op.MaxTimeToReceive
}
func (o *SetMaxTimeToReceiveRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetMaxTimeToReceiveRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMaxTimeToReceiveOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetMaxTimeToReceiveResponse structure represents the MaxTimeToReceive operation response
type SetMaxTimeToReceiveResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The MaxTimeToReceive return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetMaxTimeToReceiveResponse) xxx_ToOp(ctx context.Context, op *xxx_SetMaxTimeToReceiveOperation) *xxx_SetMaxTimeToReceiveOperation {
	if op == nil {
		op = &xxx_SetMaxTimeToReceiveOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetMaxTimeToReceiveResponse) xxx_FromOp(ctx context.Context, op *xxx_SetMaxTimeToReceiveOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetMaxTimeToReceiveResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetMaxTimeToReceiveResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMaxTimeToReceiveOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetHashAlgorithmOperation structure represents the HashAlgorithm operation
type xxx_GetHashAlgorithmOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	HashAlgorithm int32          `idl:"name:plHashAlg" json:"hash_algorithm"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetHashAlgorithmOperation) OpNum() int { return 42 }

func (o *xxx_GetHashAlgorithmOperation) OpName() string { return "/IMSMQMessage/v0/HashAlgorithm" }

func (o *xxx_GetHashAlgorithmOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetHashAlgorithmOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetHashAlgorithmOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetHashAlgorithmOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetHashAlgorithmOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plHashAlg {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.HashAlgorithm); err != nil {
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

func (o *xxx_GetHashAlgorithmOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plHashAlg {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.HashAlgorithm); err != nil {
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

// GetHashAlgorithmRequest structure represents the HashAlgorithm operation request
type GetHashAlgorithmRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetHashAlgorithmRequest) xxx_ToOp(ctx context.Context, op *xxx_GetHashAlgorithmOperation) *xxx_GetHashAlgorithmOperation {
	if op == nil {
		op = &xxx_GetHashAlgorithmOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetHashAlgorithmRequest) xxx_FromOp(ctx context.Context, op *xxx_GetHashAlgorithmOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetHashAlgorithmRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetHashAlgorithmRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetHashAlgorithmOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetHashAlgorithmResponse structure represents the HashAlgorithm operation response
type GetHashAlgorithmResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	HashAlgorithm int32          `idl:"name:plHashAlg" json:"hash_algorithm"`
	// Return: The HashAlgorithm return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetHashAlgorithmResponse) xxx_ToOp(ctx context.Context, op *xxx_GetHashAlgorithmOperation) *xxx_GetHashAlgorithmOperation {
	if op == nil {
		op = &xxx_GetHashAlgorithmOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.HashAlgorithm = o.HashAlgorithm
	op.Return = o.Return
	return op
}

func (o *GetHashAlgorithmResponse) xxx_FromOp(ctx context.Context, op *xxx_GetHashAlgorithmOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.HashAlgorithm = op.HashAlgorithm
	o.Return = op.Return
}
func (o *GetHashAlgorithmResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetHashAlgorithmResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetHashAlgorithmOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetHashAlgorithmOperation structure represents the HashAlgorithm operation
type xxx_SetHashAlgorithmOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	HashAlgorithm int32          `idl:"name:lHashAlg" json:"hash_algorithm"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetHashAlgorithmOperation) OpNum() int { return 43 }

func (o *xxx_SetHashAlgorithmOperation) OpName() string { return "/IMSMQMessage/v0/HashAlgorithm" }

func (o *xxx_SetHashAlgorithmOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetHashAlgorithmOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lHashAlg {in} (1:(int32))
	{
		if err := w.WriteData(o.HashAlgorithm); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetHashAlgorithmOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lHashAlg {in} (1:(int32))
	{
		if err := w.ReadData(&o.HashAlgorithm); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetHashAlgorithmOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetHashAlgorithmOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetHashAlgorithmOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetHashAlgorithmRequest structure represents the HashAlgorithm operation request
type SetHashAlgorithmRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	HashAlgorithm int32          `idl:"name:lHashAlg" json:"hash_algorithm"`
}

func (o *SetHashAlgorithmRequest) xxx_ToOp(ctx context.Context, op *xxx_SetHashAlgorithmOperation) *xxx_SetHashAlgorithmOperation {
	if op == nil {
		op = &xxx_SetHashAlgorithmOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.HashAlgorithm = o.HashAlgorithm
	return op
}

func (o *SetHashAlgorithmRequest) xxx_FromOp(ctx context.Context, op *xxx_SetHashAlgorithmOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.HashAlgorithm = op.HashAlgorithm
}
func (o *SetHashAlgorithmRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetHashAlgorithmRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetHashAlgorithmOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetHashAlgorithmResponse structure represents the HashAlgorithm operation response
type SetHashAlgorithmResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The HashAlgorithm return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetHashAlgorithmResponse) xxx_ToOp(ctx context.Context, op *xxx_SetHashAlgorithmOperation) *xxx_SetHashAlgorithmOperation {
	if op == nil {
		op = &xxx_SetHashAlgorithmOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetHashAlgorithmResponse) xxx_FromOp(ctx context.Context, op *xxx_SetHashAlgorithmOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetHashAlgorithmResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetHashAlgorithmResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetHashAlgorithmOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetEncryptAlgorithmOperation structure represents the EncryptAlgorithm operation
type xxx_GetEncryptAlgorithmOperation struct {
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	EncryptAlgorithm int32          `idl:"name:plEncryptAlg" json:"encrypt_algorithm"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetEncryptAlgorithmOperation) OpNum() int { return 44 }

func (o *xxx_GetEncryptAlgorithmOperation) OpName() string {
	return "/IMSMQMessage/v0/EncryptAlgorithm"
}

func (o *xxx_GetEncryptAlgorithmOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEncryptAlgorithmOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetEncryptAlgorithmOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetEncryptAlgorithmOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEncryptAlgorithmOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plEncryptAlg {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.EncryptAlgorithm); err != nil {
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

func (o *xxx_GetEncryptAlgorithmOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plEncryptAlg {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.EncryptAlgorithm); err != nil {
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

// GetEncryptAlgorithmRequest structure represents the EncryptAlgorithm operation request
type GetEncryptAlgorithmRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetEncryptAlgorithmRequest) xxx_ToOp(ctx context.Context, op *xxx_GetEncryptAlgorithmOperation) *xxx_GetEncryptAlgorithmOperation {
	if op == nil {
		op = &xxx_GetEncryptAlgorithmOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetEncryptAlgorithmRequest) xxx_FromOp(ctx context.Context, op *xxx_GetEncryptAlgorithmOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetEncryptAlgorithmRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetEncryptAlgorithmRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEncryptAlgorithmOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetEncryptAlgorithmResponse structure represents the EncryptAlgorithm operation response
type GetEncryptAlgorithmResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	EncryptAlgorithm int32          `idl:"name:plEncryptAlg" json:"encrypt_algorithm"`
	// Return: The EncryptAlgorithm return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetEncryptAlgorithmResponse) xxx_ToOp(ctx context.Context, op *xxx_GetEncryptAlgorithmOperation) *xxx_GetEncryptAlgorithmOperation {
	if op == nil {
		op = &xxx_GetEncryptAlgorithmOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.EncryptAlgorithm = o.EncryptAlgorithm
	op.Return = o.Return
	return op
}

func (o *GetEncryptAlgorithmResponse) xxx_FromOp(ctx context.Context, op *xxx_GetEncryptAlgorithmOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.EncryptAlgorithm = op.EncryptAlgorithm
	o.Return = op.Return
}
func (o *GetEncryptAlgorithmResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetEncryptAlgorithmResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEncryptAlgorithmOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetEncryptAlgorithmOperation structure represents the EncryptAlgorithm operation
type xxx_SetEncryptAlgorithmOperation struct {
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	EncryptAlgorithm int32          `idl:"name:lEncryptAlg" json:"encrypt_algorithm"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetEncryptAlgorithmOperation) OpNum() int { return 45 }

func (o *xxx_SetEncryptAlgorithmOperation) OpName() string {
	return "/IMSMQMessage/v0/EncryptAlgorithm"
}

func (o *xxx_SetEncryptAlgorithmOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEncryptAlgorithmOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lEncryptAlg {in} (1:(int32))
	{
		if err := w.WriteData(o.EncryptAlgorithm); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEncryptAlgorithmOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lEncryptAlg {in} (1:(int32))
	{
		if err := w.ReadData(&o.EncryptAlgorithm); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEncryptAlgorithmOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEncryptAlgorithmOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetEncryptAlgorithmOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetEncryptAlgorithmRequest structure represents the EncryptAlgorithm operation request
type SetEncryptAlgorithmRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	EncryptAlgorithm int32          `idl:"name:lEncryptAlg" json:"encrypt_algorithm"`
}

func (o *SetEncryptAlgorithmRequest) xxx_ToOp(ctx context.Context, op *xxx_SetEncryptAlgorithmOperation) *xxx_SetEncryptAlgorithmOperation {
	if op == nil {
		op = &xxx_SetEncryptAlgorithmOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.EncryptAlgorithm = o.EncryptAlgorithm
	return op
}

func (o *SetEncryptAlgorithmRequest) xxx_FromOp(ctx context.Context, op *xxx_SetEncryptAlgorithmOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.EncryptAlgorithm = op.EncryptAlgorithm
}
func (o *SetEncryptAlgorithmRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetEncryptAlgorithmRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetEncryptAlgorithmOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetEncryptAlgorithmResponse structure represents the EncryptAlgorithm operation response
type SetEncryptAlgorithmResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The EncryptAlgorithm return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetEncryptAlgorithmResponse) xxx_ToOp(ctx context.Context, op *xxx_SetEncryptAlgorithmOperation) *xxx_SetEncryptAlgorithmOperation {
	if op == nil {
		op = &xxx_SetEncryptAlgorithmOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetEncryptAlgorithmResponse) xxx_FromOp(ctx context.Context, op *xxx_SetEncryptAlgorithmOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetEncryptAlgorithmResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetEncryptAlgorithmResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetEncryptAlgorithmOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSentTimeOperation structure represents the SentTime operation
type xxx_GetSentTimeOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	SentTime *oaut.Variant  `idl:"name:pvarSentTime" json:"sent_time"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSentTimeOperation) OpNum() int { return 46 }

func (o *xxx_GetSentTimeOperation) OpName() string { return "/IMSMQMessage/v0/SentTime" }

func (o *xxx_GetSentTimeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSentTimeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSentTimeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSentTimeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSentTimeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pvarSentTime {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.SentTime != nil {
			_ptr_pvarSentTime := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SentTime != nil {
					if err := o.SentTime.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SentTime, _ptr_pvarSentTime); err != nil {
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

func (o *xxx_GetSentTimeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pvarSentTime {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pvarSentTime := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SentTime == nil {
				o.SentTime = &oaut.Variant{}
			}
			if err := o.SentTime.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pvarSentTime := func(ptr interface{}) { o.SentTime = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.SentTime, _s_pvarSentTime, _ptr_pvarSentTime); err != nil {
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

// GetSentTimeRequest structure represents the SentTime operation request
type GetSentTimeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSentTimeRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSentTimeOperation) *xxx_GetSentTimeOperation {
	if op == nil {
		op = &xxx_GetSentTimeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetSentTimeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSentTimeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSentTimeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSentTimeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSentTimeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSentTimeResponse structure represents the SentTime operation response
type GetSentTimeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pvarSentTime: A pointer to a VARIANT that contains a VT_DATE that contains the UTC
	// date and time stamp of when the message was sent.
	SentTime *oaut.Variant `idl:"name:pvarSentTime" json:"sent_time"`
	// Return: The SentTime return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSentTimeResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSentTimeOperation) *xxx_GetSentTimeOperation {
	if op == nil {
		op = &xxx_GetSentTimeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.SentTime = o.SentTime
	op.Return = o.Return
	return op
}

func (o *GetSentTimeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSentTimeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SentTime = op.SentTime
	o.Return = op.Return
}
func (o *GetSentTimeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSentTimeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSentTimeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetArrivedTimeOperation structure represents the ArrivedTime operation
type xxx_GetArrivedTimeOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ArrivedTime *oaut.Variant  `idl:"name:plArrivedTime" json:"arrived_time"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetArrivedTimeOperation) OpNum() int { return 47 }

func (o *xxx_GetArrivedTimeOperation) OpName() string { return "/IMSMQMessage/v0/ArrivedTime" }

func (o *xxx_GetArrivedTimeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetArrivedTimeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetArrivedTimeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetArrivedTimeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetArrivedTimeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plArrivedTime {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.ArrivedTime != nil {
			_ptr_plArrivedTime := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ArrivedTime != nil {
					if err := o.ArrivedTime.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ArrivedTime, _ptr_plArrivedTime); err != nil {
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

func (o *xxx_GetArrivedTimeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plArrivedTime {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_plArrivedTime := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ArrivedTime == nil {
				o.ArrivedTime = &oaut.Variant{}
			}
			if err := o.ArrivedTime.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_plArrivedTime := func(ptr interface{}) { o.ArrivedTime = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.ArrivedTime, _s_plArrivedTime, _ptr_plArrivedTime); err != nil {
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

// GetArrivedTimeRequest structure represents the ArrivedTime operation request
type GetArrivedTimeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetArrivedTimeRequest) xxx_ToOp(ctx context.Context, op *xxx_GetArrivedTimeOperation) *xxx_GetArrivedTimeOperation {
	if op == nil {
		op = &xxx_GetArrivedTimeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetArrivedTimeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetArrivedTimeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetArrivedTimeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetArrivedTimeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetArrivedTimeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetArrivedTimeResponse structure represents the ArrivedTime operation response
type GetArrivedTimeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// plArrivedTime: A pointer to a VARIANT that contains a VT_DATE that contains the UTC
	// date and time stamp of when the message arrived in the queue.
	ArrivedTime *oaut.Variant `idl:"name:plArrivedTime" json:"arrived_time"`
	// Return: The ArrivedTime return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetArrivedTimeResponse) xxx_ToOp(ctx context.Context, op *xxx_GetArrivedTimeOperation) *xxx_GetArrivedTimeOperation {
	if op == nil {
		op = &xxx_GetArrivedTimeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ArrivedTime = o.ArrivedTime
	op.Return = o.Return
	return op
}

func (o *GetArrivedTimeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetArrivedTimeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ArrivedTime = op.ArrivedTime
	o.Return = op.Return
}
func (o *GetArrivedTimeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetArrivedTimeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetArrivedTimeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDestinationQueueInfoOperation structure represents the DestinationQueueInfo operation
type xxx_GetDestinationQueueInfoOperation struct {
	This        *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Destination *mqac.QueueInfo `idl:"name:ppqinfoDest" json:"destination"`
	Return      int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDestinationQueueInfoOperation) OpNum() int { return 48 }

func (o *xxx_GetDestinationQueueInfoOperation) OpName() string {
	return "/IMSMQMessage/v0/DestinationQueueInfo"
}

func (o *xxx_GetDestinationQueueInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDestinationQueueInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetDestinationQueueInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetDestinationQueueInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDestinationQueueInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppqinfoDest {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQQueueInfo}(interface))
	{
		if o.Destination != nil {
			_ptr_ppqinfoDest := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Destination != nil {
					if err := o.Destination.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.QueueInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Destination, _ptr_ppqinfoDest); err != nil {
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

func (o *xxx_GetDestinationQueueInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppqinfoDest {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQQueueInfo}(interface))
	{
		_ptr_ppqinfoDest := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Destination == nil {
				o.Destination = &mqac.QueueInfo{}
			}
			if err := o.Destination.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppqinfoDest := func(ptr interface{}) { o.Destination = *ptr.(**mqac.QueueInfo) }
		if err := w.ReadPointer(&o.Destination, _s_ppqinfoDest, _ptr_ppqinfoDest); err != nil {
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

// GetDestinationQueueInfoRequest structure represents the DestinationQueueInfo operation request
type GetDestinationQueueInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetDestinationQueueInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_GetDestinationQueueInfoOperation) *xxx_GetDestinationQueueInfoOperation {
	if op == nil {
		op = &xxx_GetDestinationQueueInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetDestinationQueueInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDestinationQueueInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetDestinationQueueInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetDestinationQueueInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDestinationQueueInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDestinationQueueInfoResponse structure represents the DestinationQueueInfo operation response
type GetDestinationQueueInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppqinfoDest: A pointer to a pointer to an IMSMQQueueInfo interface that upon successful
	// completion will contain an MSMQQueue object that represents the destination queue
	// for the message.
	Destination *mqac.QueueInfo `idl:"name:ppqinfoDest" json:"destination"`
	// Return: The DestinationQueueInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDestinationQueueInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_GetDestinationQueueInfoOperation) *xxx_GetDestinationQueueInfoOperation {
	if op == nil {
		op = &xxx_GetDestinationQueueInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Destination = o.Destination
	op.Return = o.Return
	return op
}

func (o *GetDestinationQueueInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDestinationQueueInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Destination = op.Destination
	o.Return = op.Return
}
func (o *GetDestinationQueueInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetDestinationQueueInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDestinationQueueInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSenderCertificateOperation structure represents the SenderCertificate operation
type xxx_GetSenderCertificateOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	SenderCert *oaut.Variant  `idl:"name:pvarSenderCert" json:"sender_cert"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSenderCertificateOperation) OpNum() int { return 49 }

func (o *xxx_GetSenderCertificateOperation) OpName() string {
	return "/IMSMQMessage/v0/SenderCertificate"
}

func (o *xxx_GetSenderCertificateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSenderCertificateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSenderCertificateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSenderCertificateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSenderCertificateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pvarSenderCert {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.SenderCert != nil {
			_ptr_pvarSenderCert := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SenderCert != nil {
					if err := o.SenderCert.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SenderCert, _ptr_pvarSenderCert); err != nil {
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

func (o *xxx_GetSenderCertificateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pvarSenderCert {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pvarSenderCert := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SenderCert == nil {
				o.SenderCert = &oaut.Variant{}
			}
			if err := o.SenderCert.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pvarSenderCert := func(ptr interface{}) { o.SenderCert = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.SenderCert, _s_pvarSenderCert, _ptr_pvarSenderCert); err != nil {
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

// GetSenderCertificateRequest structure represents the SenderCertificate operation request
type GetSenderCertificateRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSenderCertificateRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSenderCertificateOperation) *xxx_GetSenderCertificateOperation {
	if op == nil {
		op = &xxx_GetSenderCertificateOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetSenderCertificateRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSenderCertificateOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSenderCertificateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSenderCertificateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSenderCertificateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSenderCertificateResponse structure represents the SenderCertificate operation response
type GetSenderCertificateResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	SenderCert *oaut.Variant  `idl:"name:pvarSenderCert" json:"sender_cert"`
	// Return: The SenderCertificate return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSenderCertificateResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSenderCertificateOperation) *xxx_GetSenderCertificateOperation {
	if op == nil {
		op = &xxx_GetSenderCertificateOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.SenderCert = o.SenderCert
	op.Return = o.Return
	return op
}

func (o *GetSenderCertificateResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSenderCertificateOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SenderCert = op.SenderCert
	o.Return = op.Return
}
func (o *GetSenderCertificateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSenderCertificateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSenderCertificateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSenderCertificateOperation structure represents the SenderCertificate operation
type xxx_SetSenderCertificateOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	SenderCert *oaut.Variant  `idl:"name:varSenderCert" json:"sender_cert"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSenderCertificateOperation) OpNum() int { return 50 }

func (o *xxx_SetSenderCertificateOperation) OpName() string {
	return "/IMSMQMessage/v0/SenderCertificate"
}

func (o *xxx_SetSenderCertificateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSenderCertificateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// varSenderCert {in} (1:{alias=VARIANT}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.SenderCert != nil {
			if err := o.SenderCert.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSenderCertificateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// varSenderCert {in} (1:{alias=VARIANT,pointer=ref}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.SenderCert == nil {
			o.SenderCert = &oaut.Variant{}
		}
		if err := o.SenderCert.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSenderCertificateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSenderCertificateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetSenderCertificateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetSenderCertificateRequest structure represents the SenderCertificate operation request
type SetSenderCertificateRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	SenderCert *oaut.Variant  `idl:"name:varSenderCert" json:"sender_cert"`
}

func (o *SetSenderCertificateRequest) xxx_ToOp(ctx context.Context, op *xxx_SetSenderCertificateOperation) *xxx_SetSenderCertificateOperation {
	if op == nil {
		op = &xxx_SetSenderCertificateOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.SenderCert = o.SenderCert
	return op
}

func (o *SetSenderCertificateRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSenderCertificateOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SenderCert = op.SenderCert
}
func (o *SetSenderCertificateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetSenderCertificateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSenderCertificateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSenderCertificateResponse structure represents the SenderCertificate operation response
type SetSenderCertificateResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SenderCertificate return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSenderCertificateResponse) xxx_ToOp(ctx context.Context, op *xxx_SetSenderCertificateOperation) *xxx_SetSenderCertificateOperation {
	if op == nil {
		op = &xxx_SetSenderCertificateOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetSenderCertificateResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSenderCertificateOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSenderCertificateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetSenderCertificateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSenderCertificateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSenderIDOperation structure represents the SenderId operation
type xxx_GetSenderIDOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	SenderID *oaut.Variant  `idl:"name:pvarSenderId" json:"sender_id"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSenderIDOperation) OpNum() int { return 51 }

func (o *xxx_GetSenderIDOperation) OpName() string { return "/IMSMQMessage/v0/SenderId" }

func (o *xxx_GetSenderIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSenderIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSenderIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSenderIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSenderIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pvarSenderId {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.SenderID != nil {
			_ptr_pvarSenderId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SenderID != nil {
					if err := o.SenderID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SenderID, _ptr_pvarSenderId); err != nil {
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

func (o *xxx_GetSenderIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pvarSenderId {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pvarSenderId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SenderID == nil {
				o.SenderID = &oaut.Variant{}
			}
			if err := o.SenderID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pvarSenderId := func(ptr interface{}) { o.SenderID = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.SenderID, _s_pvarSenderId, _ptr_pvarSenderId); err != nil {
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

// GetSenderIDRequest structure represents the SenderId operation request
type GetSenderIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSenderIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSenderIDOperation) *xxx_GetSenderIDOperation {
	if op == nil {
		op = &xxx_GetSenderIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetSenderIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSenderIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSenderIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSenderIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSenderIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSenderIDResponse structure represents the SenderId operation response
type GetSenderIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	SenderID *oaut.Variant  `idl:"name:pvarSenderId" json:"sender_id"`
	// Return: The SenderId return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSenderIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSenderIDOperation) *xxx_GetSenderIDOperation {
	if op == nil {
		op = &xxx_GetSenderIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.SenderID = o.SenderID
	op.Return = o.Return
	return op
}

func (o *GetSenderIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSenderIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SenderID = op.SenderID
	o.Return = op.Return
}
func (o *GetSenderIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSenderIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSenderIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSenderIDTypeOperation structure represents the SenderIdType operation
type xxx_GetSenderIDTypeOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	SenderIDType int32          `idl:"name:plSenderIdType" json:"sender_id_type"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSenderIDTypeOperation) OpNum() int { return 52 }

func (o *xxx_GetSenderIDTypeOperation) OpName() string { return "/IMSMQMessage/v0/SenderIdType" }

func (o *xxx_GetSenderIDTypeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSenderIDTypeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSenderIDTypeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSenderIDTypeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSenderIDTypeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plSenderIdType {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.SenderIDType); err != nil {
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

func (o *xxx_GetSenderIDTypeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plSenderIdType {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.SenderIDType); err != nil {
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

// GetSenderIDTypeRequest structure represents the SenderIdType operation request
type GetSenderIDTypeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSenderIDTypeRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSenderIDTypeOperation) *xxx_GetSenderIDTypeOperation {
	if op == nil {
		op = &xxx_GetSenderIDTypeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetSenderIDTypeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSenderIDTypeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSenderIDTypeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSenderIDTypeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSenderIDTypeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSenderIDTypeResponse structure represents the SenderIdType operation response
type GetSenderIDTypeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	SenderIDType int32          `idl:"name:plSenderIdType" json:"sender_id_type"`
	// Return: The SenderIdType return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSenderIDTypeResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSenderIDTypeOperation) *xxx_GetSenderIDTypeOperation {
	if op == nil {
		op = &xxx_GetSenderIDTypeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.SenderIDType = o.SenderIDType
	op.Return = o.Return
	return op
}

func (o *GetSenderIDTypeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSenderIDTypeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SenderIDType = op.SenderIDType
	o.Return = op.Return
}
func (o *GetSenderIDTypeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSenderIDTypeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSenderIDTypeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSenderIDTypeOperation structure represents the SenderIdType operation
type xxx_SetSenderIDTypeOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	SenderIDType int32          `idl:"name:lSenderIdType" json:"sender_id_type"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSenderIDTypeOperation) OpNum() int { return 53 }

func (o *xxx_SetSenderIDTypeOperation) OpName() string { return "/IMSMQMessage/v0/SenderIdType" }

func (o *xxx_SetSenderIDTypeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSenderIDTypeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lSenderIdType {in} (1:(int32))
	{
		if err := w.WriteData(o.SenderIDType); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSenderIDTypeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lSenderIdType {in} (1:(int32))
	{
		if err := w.ReadData(&o.SenderIDType); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSenderIDTypeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSenderIDTypeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetSenderIDTypeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetSenderIDTypeRequest structure represents the SenderIdType operation request
type SetSenderIDTypeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	SenderIDType int32          `idl:"name:lSenderIdType" json:"sender_id_type"`
}

func (o *SetSenderIDTypeRequest) xxx_ToOp(ctx context.Context, op *xxx_SetSenderIDTypeOperation) *xxx_SetSenderIDTypeOperation {
	if op == nil {
		op = &xxx_SetSenderIDTypeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.SenderIDType = o.SenderIDType
	return op
}

func (o *SetSenderIDTypeRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSenderIDTypeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SenderIDType = op.SenderIDType
}
func (o *SetSenderIDTypeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetSenderIDTypeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSenderIDTypeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSenderIDTypeResponse structure represents the SenderIdType operation response
type SetSenderIDTypeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SenderIdType return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSenderIDTypeResponse) xxx_ToOp(ctx context.Context, op *xxx_SetSenderIDTypeOperation) *xxx_SetSenderIDTypeOperation {
	if op == nil {
		op = &xxx_SetSenderIDTypeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetSenderIDTypeResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSenderIDTypeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSenderIDTypeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetSenderIDTypeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSenderIDTypeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SendOperation structure represents the Send operation
type xxx_SendOperation struct {
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	DestinationQueue *mqac.Queue    `idl:"name:DestinationQueue" json:"destination_queue"`
	Transaction      *oaut.Variant  `idl:"name:Transaction" json:"transaction"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SendOperation) OpNum() int { return 54 }

func (o *xxx_SendOperation) OpName() string { return "/IMSMQMessage/v0/Send" }

func (o *xxx_SendOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// DestinationQueue {in} (1:{pointer=ref}*(1))(2:{alias=IMSMQQueue}(interface))
	{
		if o.DestinationQueue != nil {
			_ptr_DestinationQueue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DestinationQueue != nil {
					if err := o.DestinationQueue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Queue{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DestinationQueue, _ptr_DestinationQueue); err != nil {
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
	// Transaction {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Transaction != nil {
			_ptr_Transaction := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Transaction != nil {
					if err := o.Transaction.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Transaction, _ptr_Transaction); err != nil {
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
	return nil
}

func (o *xxx_SendOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// DestinationQueue {in} (1:{pointer=ref}*(1))(2:{alias=IMSMQQueue}(interface))
	{
		_ptr_DestinationQueue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DestinationQueue == nil {
				o.DestinationQueue = &mqac.Queue{}
			}
			if err := o.DestinationQueue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_DestinationQueue := func(ptr interface{}) { o.DestinationQueue = *ptr.(**mqac.Queue) }
		if err := w.ReadPointer(&o.DestinationQueue, _s_DestinationQueue, _ptr_DestinationQueue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Transaction {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_Transaction := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Transaction == nil {
				o.Transaction = &oaut.Variant{}
			}
			if err := o.Transaction.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Transaction := func(ptr interface{}) { o.Transaction = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Transaction, _s_Transaction, _ptr_Transaction); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SendOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SendRequest structure represents the Send operation request
type SendRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DestinationQueue: A pointer to an IDispatch interface that can reference either an
	// MSMQQueue object instance or an MSMQDestination object instance.
	DestinationQueue *mqac.Queue `idl:"name:DestinationQueue" json:"destination_queue"`
	// Transaction: A pointer to a VARIANT that contains either:
	//
	//   - A VT_DISPATCH or a VT_DISPATCH | VT_BYREF that points to an MSMQTransaction object.
	//
	// * A VT_I4 that references one of the following enumeration values as defined in section
	// 2.2.2.1 ( 3ff1f343-d6e9-4bab-848f-6f2b39209762 ).
	Transaction *oaut.Variant `idl:"name:Transaction" json:"transaction"`
}

func (o *SendRequest) xxx_ToOp(ctx context.Context, op *xxx_SendOperation) *xxx_SendOperation {
	if op == nil {
		op = &xxx_SendOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.DestinationQueue = o.DestinationQueue
	op.Transaction = o.Transaction
	return op
}

func (o *SendRequest) xxx_FromOp(ctx context.Context, op *xxx_SendOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DestinationQueue = op.DestinationQueue
	o.Transaction = op.Transaction
}
func (o *SendRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SendRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SendOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SendResponse structure represents the Send operation response
type SendResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Send return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SendResponse) xxx_ToOp(ctx context.Context, op *xxx_SendOperation) *xxx_SendOperation {
	if op == nil {
		op = &xxx_SendOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SendResponse) xxx_FromOp(ctx context.Context, op *xxx_SendOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SendResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SendResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SendOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AttachCurrentSecurityContextOperation structure represents the AttachCurrentSecurityContext operation
type xxx_AttachCurrentSecurityContextOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_AttachCurrentSecurityContextOperation) OpNum() int { return 55 }

func (o *xxx_AttachCurrentSecurityContextOperation) OpName() string {
	return "/IMSMQMessage/v0/AttachCurrentSecurityContext"
}

func (o *xxx_AttachCurrentSecurityContextOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AttachCurrentSecurityContextOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AttachCurrentSecurityContextOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_AttachCurrentSecurityContextOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AttachCurrentSecurityContextOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AttachCurrentSecurityContextOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// AttachCurrentSecurityContextRequest structure represents the AttachCurrentSecurityContext operation request
type AttachCurrentSecurityContextRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *AttachCurrentSecurityContextRequest) xxx_ToOp(ctx context.Context, op *xxx_AttachCurrentSecurityContextOperation) *xxx_AttachCurrentSecurityContextOperation {
	if op == nil {
		op = &xxx_AttachCurrentSecurityContextOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *AttachCurrentSecurityContextRequest) xxx_FromOp(ctx context.Context, op *xxx_AttachCurrentSecurityContextOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *AttachCurrentSecurityContextRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AttachCurrentSecurityContextRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AttachCurrentSecurityContextOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AttachCurrentSecurityContextResponse structure represents the AttachCurrentSecurityContext operation response
type AttachCurrentSecurityContextResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AttachCurrentSecurityContext return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AttachCurrentSecurityContextResponse) xxx_ToOp(ctx context.Context, op *xxx_AttachCurrentSecurityContextOperation) *xxx_AttachCurrentSecurityContextOperation {
	if op == nil {
		op = &xxx_AttachCurrentSecurityContextOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *AttachCurrentSecurityContextResponse) xxx_FromOp(ctx context.Context, op *xxx_AttachCurrentSecurityContextOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *AttachCurrentSecurityContextResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AttachCurrentSecurityContextResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AttachCurrentSecurityContextOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
