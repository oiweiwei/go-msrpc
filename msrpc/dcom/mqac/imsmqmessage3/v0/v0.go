package imsmqmessage3

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
	// IMSMQMessage3 interface identifier eba96b1a-2168-11d3-898c-00e02c074f6b
	Message3IID = &dcom.IID{Data1: 0xeba96b1a, Data2: 0x2168, Data3: 0x11d3, Data4: []byte{0x89, 0x8c, 0x00, 0xe0, 0x2c, 0x07, 0x4f, 0x6b}}
	// Syntax UUID
	Message3SyntaxUUID = &uuid.UUID{TimeLow: 0xeba96b1a, TimeMid: 0x2168, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x8c, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	// Syntax ID
	Message3SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: Message3SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IMSMQMessage3 interface.
type Message3Client interface {

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

	// ResponseQueueInfo_v1 operation.
	GetResponseQueueInfoV1(context.Context, *GetResponseQueueInfoV1Request, ...dcerpc.CallOption) (*GetResponseQueueInfoV1Response, error)

	// ResponseQueueInfo_v1 operation.
	SetResponseQueueInfoV1(context.Context, *SetResponseQueueInfoV1Request, ...dcerpc.CallOption) (*SetResponseQueueInfoV1Response, error)

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

	// AdminQueueInfo_v1 operation.
	GetAdminQueueInfoV1(context.Context, *GetAdminQueueInfoV1Request, ...dcerpc.CallOption) (*GetAdminQueueInfoV1Response, error)

	// AdminQueueInfo_v1 operation.
	SetAdminQueueInfoV1(context.Context, *SetAdminQueueInfoV1Request, ...dcerpc.CallOption) (*SetAdminQueueInfoV1Response, error)

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

	// The SenderVersion method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return a numeric value that indicates the version of transfer used
	// to send the message.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetSenderVersion(context.Context, *GetSenderVersionRequest, ...dcerpc.CallOption) (*GetSenderVersionResponse, error)

	// Extension operation.
	GetExtension(context.Context, *GetExtensionRequest, ...dcerpc.CallOption) (*GetExtensionResponse, error)

	// Extension operation.
	SetExtension(context.Context, *SetExtensionRequest, ...dcerpc.CallOption) (*SetExtensionResponse, error)

	// ConnectorTypeGuid operation.
	GetConnectorTypeGUID(context.Context, *GetConnectorTypeGUIDRequest, ...dcerpc.CallOption) (*GetConnectorTypeGUIDResponse, error)

	// ConnectorTypeGuid operation.
	SetConnectorTypeGUID(context.Context, *SetConnectorTypeGUIDRequest, ...dcerpc.CallOption) (*SetConnectorTypeGUIDResponse, error)

	// The TransactionStatusQueueInfo method is received by the server in an RPC_REQUEST
	// packet. In response, the server MUST return an IMSMQQueueInfo interface pointer to
	// an MSMQQueueInfo object that represents the queue identified by the represented Message.TransactionStatusQueueFormatName.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetTransactionStatusQueueInfo(context.Context, *GetTransactionStatusQueueInfoRequest, ...dcerpc.CallOption) (*GetTransactionStatusQueueInfoResponse, error)

	// DestinationSymmetricKey operation.
	GetDestinationSymmetricKey(context.Context, *GetDestinationSymmetricKeyRequest, ...dcerpc.CallOption) (*GetDestinationSymmetricKeyResponse, error)

	// DestinationSymmetricKey operation.
	SetDestinationSymmetricKey(context.Context, *SetDestinationSymmetricKeyRequest, ...dcerpc.CallOption) (*SetDestinationSymmetricKeyResponse, error)

	// Signature operation.
	GetSignature(context.Context, *GetSignatureRequest, ...dcerpc.CallOption) (*GetSignatureResponse, error)

	// Signature operation.
	SetSignature(context.Context, *SetSignatureRequest, ...dcerpc.CallOption) (*SetSignatureResponse, error)

	// AuthenticationProviderType operation.
	GetAuthenticationProviderType(context.Context, *GetAuthenticationProviderTypeRequest, ...dcerpc.CallOption) (*GetAuthenticationProviderTypeResponse, error)

	// AuthenticationProviderType operation.
	SetAuthenticationProviderType(context.Context, *SetAuthenticationProviderTypeRequest, ...dcerpc.CallOption) (*SetAuthenticationProviderTypeResponse, error)

	// AuthenticationProviderName operation.
	GetAuthenticationProviderName(context.Context, *GetAuthenticationProviderNameRequest, ...dcerpc.CallOption) (*GetAuthenticationProviderNameResponse, error)

	// AuthenticationProviderName operation.
	SetAuthenticationProviderName(context.Context, *SetAuthenticationProviderNameRequest, ...dcerpc.CallOption) (*SetAuthenticationProviderNameResponse, error)

	// SenderId operation.
	SetSenderID(context.Context, *SetSenderIDRequest, ...dcerpc.CallOption) (*SetSenderIDResponse, error)

	// MsgClass operation.
	GetMessageClass(context.Context, *GetMessageClassRequest, ...dcerpc.CallOption) (*GetMessageClassResponse, error)

	// MsgClass operation.
	SetMessageClass(context.Context, *SetMessageClassRequest, ...dcerpc.CallOption) (*SetMessageClassResponse, error)

	// Properties operation.
	GetProperties(context.Context, *GetPropertiesRequest, ...dcerpc.CallOption) (*GetPropertiesResponse, error)

	// The TransactionId method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return the represented Message.TransactionalMessageSequenceIdentifier.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetTransactionID(context.Context, *GetTransactionIDRequest, ...dcerpc.CallOption) (*GetTransactionIDResponse, error)

	// The IsFirstInTransaction method is received by the server in an RPC_REQUEST packet.
	// In response, the server MUST return the represented Message.FirstInTransaction.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetIsFirstInTransaction(context.Context, *GetIsFirstInTransactionRequest, ...dcerpc.CallOption) (*GetIsFirstInTransactionResponse, error)

	// The IsLastInTransaction method is received by the server in an RPC_REQUEST packet.
	// In response, the server MUST return the represented Message.LastInTransaction.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetIsLastInTransaction(context.Context, *GetIsLastInTransactionRequest, ...dcerpc.CallOption) (*GetIsLastInTransactionResponse, error)

	// ResponseQueueInfo_v2 operation.
	GetResponseQueueInfoV2(context.Context, *GetResponseQueueInfoV2Request, ...dcerpc.CallOption) (*GetResponseQueueInfoV2Response, error)

	// ResponseQueueInfo_v2 operation.
	SetResponseQueueInfoV2(context.Context, *SetResponseQueueInfoV2Request, ...dcerpc.CallOption) (*SetResponseQueueInfoV2Response, error)

	// AdminQueueInfo_v2 operation.
	GetAdminQueueInfoV2(context.Context, *GetAdminQueueInfoV2Request, ...dcerpc.CallOption) (*GetAdminQueueInfoV2Response, error)

	// AdminQueueInfo_v2 operation.
	SetAdminQueueInfoV2(context.Context, *SetAdminQueueInfoV2Request, ...dcerpc.CallOption) (*SetAdminQueueInfoV2Response, error)

	// The ReceivedAuthenticationLevel method is received by the server in an RPC_REQUEST
	// packet. In response, the server MUST return the represented Message.AuthenticationLevel.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetReceivedAuthenticationLevel(context.Context, *GetReceivedAuthenticationLevelRequest, ...dcerpc.CallOption) (*GetReceivedAuthenticationLevelResponse, error)

	// ResponseQueueInfo operation.
	GetResponseQueueInfo(context.Context, *GetResponseQueueInfoRequest, ...dcerpc.CallOption) (*GetResponseQueueInfoResponse, error)

	// ResponseQueueInfo operation.
	SetByRefResponseQueueInfo(context.Context, *SetByRefResponseQueueInfoRequest, ...dcerpc.CallOption) (*SetByRefResponseQueueInfoResponse, error)

	// AdminQueueInfo operation.
	GetAdminQueueInfo(context.Context, *GetAdminQueueInfoRequest, ...dcerpc.CallOption) (*GetAdminQueueInfoResponse, error)

	// AdminQueueInfo operation.
	SetByRefAdminQueueInfo(context.Context, *SetByRefAdminQueueInfoRequest, ...dcerpc.CallOption) (*SetByRefAdminQueueInfoResponse, error)

	// ResponseDestination operation.
	GetResponseDestination(context.Context, *GetResponseDestinationRequest, ...dcerpc.CallOption) (*GetResponseDestinationResponse, error)

	// ResponseDestination operation.
	SetByRefResponseDestination(context.Context, *SetByRefResponseDestinationRequest, ...dcerpc.CallOption) (*SetByRefResponseDestinationResponse, error)

	// The Destination method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return an IDispatch interface pointer to an MSMQDestination object
	// that represents 0 or more queues identified by the represented Message.DestinationMultiQueueFormatName.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetDestination(context.Context, *GetDestinationRequest, ...dcerpc.CallOption) (*GetDestinationResponse, error)

	// The LookupId method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return the represented Message.LookupIdentifier.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetLookupID(context.Context, *GetLookupIDRequest, ...dcerpc.CallOption) (*GetLookupIDResponse, error)

	// The IsAuthenticated2 method is received by the server in an RPC_REQUEST packet. In
	// response, the server MUST return a Boolean flag indicating if the message was authenticated
	// by the Queue Manager that received the Message. This provides identical behavior
	// to the method IMSMQMessage4::IsAuthenticated (section 3.17.4.1.6).
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetIsAuthenticated2(context.Context, *GetIsAuthenticated2Request, ...dcerpc.CallOption) (*GetIsAuthenticated2Response, error)

	// The IsFirstInTransaction2 method is received by the server in an RPC_REQUEST packet.
	// In response, the server MUST return the represented Message.FirstInTransaction.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetIsFirstInTransaction2(context.Context, *GetIsFirstInTransaction2Request, ...dcerpc.CallOption) (*GetIsFirstInTransaction2Response, error)

	// The IsLastInTransaction2 method is received by the server in an RPC_REQUEST packet.
	// In response, the server MUST return the represented Message.LastInTransaction.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetIsLastInTransaction2(context.Context, *GetIsLastInTransaction2Request, ...dcerpc.CallOption) (*GetIsLastInTransaction2Response, error)

	// The AttachCurrentSecurityContext2 method is received by the server in an RPC_REQUEST
	// packet. In response, the server caches the relevant information required to sign
	// a message on behalf of the client, including Message.SenderIdentifier and Message.SenderCertificate.
	// This method is provided purely as an optimization to allow the client to reduce lookups
	// of the security information about the calling client each time the message is sent.
	// The represented Message.SenderIdentifier and Message.SenderCertificate property values
	// MUST NOT be updated as a result of calling this method.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	AttachCurrentSecurityContext2(context.Context, *AttachCurrentSecurityContext2Request, ...dcerpc.CallOption) (*AttachCurrentSecurityContext2Response, error)

	// The SoapEnvelope method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return the represented Message.SoapEnvelope.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetSOAPEnvelope(context.Context, *GetSOAPEnvelopeRequest, ...dcerpc.CallOption) (*GetSOAPEnvelopeResponse, error)

	// The CompoundMessage method is received by the server in an RPC_REQUEST packet. In
	// response, the server MUST return the represented Message.SoapCompoundMessage.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetCompoundMessage(context.Context, *GetCompoundMessageRequest, ...dcerpc.CallOption) (*GetCompoundMessageResponse, error)

	// The SoapHeader method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST set the SOAPHeader instance variable.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	SetSOAPHeader(context.Context, *SetSOAPHeaderRequest, ...dcerpc.CallOption) (*SetSOAPHeaderResponse, error)

	// The SoapBody method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST set the SOAPBody instance variable.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	SetSOAPBody(context.Context, *SetSOAPBodyRequest, ...dcerpc.CallOption) (*SetSOAPBodyResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) Message3Client
}

type xxx_DefaultMessage3Client struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultMessage3Client) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultMessage3Client) GetClass(ctx context.Context, in *GetClassRequest, opts ...dcerpc.CallOption) (*GetClassResponse, error) {
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

func (o *xxx_DefaultMessage3Client) GetPrivacyLevel(ctx context.Context, in *GetPrivacyLevelRequest, opts ...dcerpc.CallOption) (*GetPrivacyLevelResponse, error) {
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

func (o *xxx_DefaultMessage3Client) SetPrivacyLevel(ctx context.Context, in *SetPrivacyLevelRequest, opts ...dcerpc.CallOption) (*SetPrivacyLevelResponse, error) {
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

func (o *xxx_DefaultMessage3Client) GetAuthLevel(ctx context.Context, in *GetAuthLevelRequest, opts ...dcerpc.CallOption) (*GetAuthLevelResponse, error) {
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

func (o *xxx_DefaultMessage3Client) SetAuthLevel(ctx context.Context, in *SetAuthLevelRequest, opts ...dcerpc.CallOption) (*SetAuthLevelResponse, error) {
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

func (o *xxx_DefaultMessage3Client) GetIsAuthenticated(ctx context.Context, in *GetIsAuthenticatedRequest, opts ...dcerpc.CallOption) (*GetIsAuthenticatedResponse, error) {
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

func (o *xxx_DefaultMessage3Client) GetDelivery(ctx context.Context, in *GetDeliveryRequest, opts ...dcerpc.CallOption) (*GetDeliveryResponse, error) {
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

func (o *xxx_DefaultMessage3Client) SetDelivery(ctx context.Context, in *SetDeliveryRequest, opts ...dcerpc.CallOption) (*SetDeliveryResponse, error) {
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

func (o *xxx_DefaultMessage3Client) GetTrace(ctx context.Context, in *GetTraceRequest, opts ...dcerpc.CallOption) (*GetTraceResponse, error) {
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

func (o *xxx_DefaultMessage3Client) SetTrace(ctx context.Context, in *SetTraceRequest, opts ...dcerpc.CallOption) (*SetTraceResponse, error) {
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

func (o *xxx_DefaultMessage3Client) GetPriority(ctx context.Context, in *GetPriorityRequest, opts ...dcerpc.CallOption) (*GetPriorityResponse, error) {
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

func (o *xxx_DefaultMessage3Client) SetPriority(ctx context.Context, in *SetPriorityRequest, opts ...dcerpc.CallOption) (*SetPriorityResponse, error) {
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

func (o *xxx_DefaultMessage3Client) GetJournal(ctx context.Context, in *GetJournalRequest, opts ...dcerpc.CallOption) (*GetJournalResponse, error) {
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

func (o *xxx_DefaultMessage3Client) SetJournal(ctx context.Context, in *SetJournalRequest, opts ...dcerpc.CallOption) (*SetJournalResponse, error) {
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

func (o *xxx_DefaultMessage3Client) GetResponseQueueInfoV1(ctx context.Context, in *GetResponseQueueInfoV1Request, opts ...dcerpc.CallOption) (*GetResponseQueueInfoV1Response, error) {
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
	out := &GetResponseQueueInfoV1Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) SetResponseQueueInfoV1(ctx context.Context, in *SetResponseQueueInfoV1Request, opts ...dcerpc.CallOption) (*SetResponseQueueInfoV1Response, error) {
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
	out := &SetResponseQueueInfoV1Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) GetAppSpecific(ctx context.Context, in *GetAppSpecificRequest, opts ...dcerpc.CallOption) (*GetAppSpecificResponse, error) {
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

func (o *xxx_DefaultMessage3Client) SetAppSpecific(ctx context.Context, in *SetAppSpecificRequest, opts ...dcerpc.CallOption) (*SetAppSpecificResponse, error) {
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

func (o *xxx_DefaultMessage3Client) GetSourceMachineGUID(ctx context.Context, in *GetSourceMachineGUIDRequest, opts ...dcerpc.CallOption) (*GetSourceMachineGUIDResponse, error) {
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

func (o *xxx_DefaultMessage3Client) GetBodyLength(ctx context.Context, in *GetBodyLengthRequest, opts ...dcerpc.CallOption) (*GetBodyLengthResponse, error) {
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

func (o *xxx_DefaultMessage3Client) GetBody(ctx context.Context, in *GetBodyRequest, opts ...dcerpc.CallOption) (*GetBodyResponse, error) {
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

func (o *xxx_DefaultMessage3Client) SetBody(ctx context.Context, in *SetBodyRequest, opts ...dcerpc.CallOption) (*SetBodyResponse, error) {
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

func (o *xxx_DefaultMessage3Client) GetAdminQueueInfoV1(ctx context.Context, in *GetAdminQueueInfoV1Request, opts ...dcerpc.CallOption) (*GetAdminQueueInfoV1Response, error) {
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
	out := &GetAdminQueueInfoV1Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) SetAdminQueueInfoV1(ctx context.Context, in *SetAdminQueueInfoV1Request, opts ...dcerpc.CallOption) (*SetAdminQueueInfoV1Response, error) {
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
	out := &SetAdminQueueInfoV1Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) GetID(ctx context.Context, in *GetIDRequest, opts ...dcerpc.CallOption) (*GetIDResponse, error) {
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

func (o *xxx_DefaultMessage3Client) GetCorrelationID(ctx context.Context, in *GetCorrelationIDRequest, opts ...dcerpc.CallOption) (*GetCorrelationIDResponse, error) {
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

func (o *xxx_DefaultMessage3Client) SetCorrelationID(ctx context.Context, in *SetCorrelationIDRequest, opts ...dcerpc.CallOption) (*SetCorrelationIDResponse, error) {
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

func (o *xxx_DefaultMessage3Client) GetAck(ctx context.Context, in *GetAckRequest, opts ...dcerpc.CallOption) (*GetAckResponse, error) {
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

func (o *xxx_DefaultMessage3Client) SetAck(ctx context.Context, in *SetAckRequest, opts ...dcerpc.CallOption) (*SetAckResponse, error) {
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

func (o *xxx_DefaultMessage3Client) GetLabel(ctx context.Context, in *GetLabelRequest, opts ...dcerpc.CallOption) (*GetLabelResponse, error) {
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

func (o *xxx_DefaultMessage3Client) SetLabel(ctx context.Context, in *SetLabelRequest, opts ...dcerpc.CallOption) (*SetLabelResponse, error) {
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

func (o *xxx_DefaultMessage3Client) GetMaxTimeToReachQueue(ctx context.Context, in *GetMaxTimeToReachQueueRequest, opts ...dcerpc.CallOption) (*GetMaxTimeToReachQueueResponse, error) {
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

func (o *xxx_DefaultMessage3Client) SetMaxTimeToReachQueue(ctx context.Context, in *SetMaxTimeToReachQueueRequest, opts ...dcerpc.CallOption) (*SetMaxTimeToReachQueueResponse, error) {
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

func (o *xxx_DefaultMessage3Client) GetMaxTimeToReceive(ctx context.Context, in *GetMaxTimeToReceiveRequest, opts ...dcerpc.CallOption) (*GetMaxTimeToReceiveResponse, error) {
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

func (o *xxx_DefaultMessage3Client) SetMaxTimeToReceive(ctx context.Context, in *SetMaxTimeToReceiveRequest, opts ...dcerpc.CallOption) (*SetMaxTimeToReceiveResponse, error) {
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

func (o *xxx_DefaultMessage3Client) GetHashAlgorithm(ctx context.Context, in *GetHashAlgorithmRequest, opts ...dcerpc.CallOption) (*GetHashAlgorithmResponse, error) {
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

func (o *xxx_DefaultMessage3Client) SetHashAlgorithm(ctx context.Context, in *SetHashAlgorithmRequest, opts ...dcerpc.CallOption) (*SetHashAlgorithmResponse, error) {
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

func (o *xxx_DefaultMessage3Client) GetEncryptAlgorithm(ctx context.Context, in *GetEncryptAlgorithmRequest, opts ...dcerpc.CallOption) (*GetEncryptAlgorithmResponse, error) {
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

func (o *xxx_DefaultMessage3Client) SetEncryptAlgorithm(ctx context.Context, in *SetEncryptAlgorithmRequest, opts ...dcerpc.CallOption) (*SetEncryptAlgorithmResponse, error) {
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

func (o *xxx_DefaultMessage3Client) GetSentTime(ctx context.Context, in *GetSentTimeRequest, opts ...dcerpc.CallOption) (*GetSentTimeResponse, error) {
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

func (o *xxx_DefaultMessage3Client) GetArrivedTime(ctx context.Context, in *GetArrivedTimeRequest, opts ...dcerpc.CallOption) (*GetArrivedTimeResponse, error) {
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

func (o *xxx_DefaultMessage3Client) GetDestinationQueueInfo(ctx context.Context, in *GetDestinationQueueInfoRequest, opts ...dcerpc.CallOption) (*GetDestinationQueueInfoResponse, error) {
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

func (o *xxx_DefaultMessage3Client) GetSenderCertificate(ctx context.Context, in *GetSenderCertificateRequest, opts ...dcerpc.CallOption) (*GetSenderCertificateResponse, error) {
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

func (o *xxx_DefaultMessage3Client) SetSenderCertificate(ctx context.Context, in *SetSenderCertificateRequest, opts ...dcerpc.CallOption) (*SetSenderCertificateResponse, error) {
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

func (o *xxx_DefaultMessage3Client) GetSenderID(ctx context.Context, in *GetSenderIDRequest, opts ...dcerpc.CallOption) (*GetSenderIDResponse, error) {
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

func (o *xxx_DefaultMessage3Client) GetSenderIDType(ctx context.Context, in *GetSenderIDTypeRequest, opts ...dcerpc.CallOption) (*GetSenderIDTypeResponse, error) {
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

func (o *xxx_DefaultMessage3Client) SetSenderIDType(ctx context.Context, in *SetSenderIDTypeRequest, opts ...dcerpc.CallOption) (*SetSenderIDTypeResponse, error) {
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

func (o *xxx_DefaultMessage3Client) Send(ctx context.Context, in *SendRequest, opts ...dcerpc.CallOption) (*SendResponse, error) {
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

func (o *xxx_DefaultMessage3Client) AttachCurrentSecurityContext(ctx context.Context, in *AttachCurrentSecurityContextRequest, opts ...dcerpc.CallOption) (*AttachCurrentSecurityContextResponse, error) {
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

func (o *xxx_DefaultMessage3Client) GetSenderVersion(ctx context.Context, in *GetSenderVersionRequest, opts ...dcerpc.CallOption) (*GetSenderVersionResponse, error) {
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
	out := &GetSenderVersionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) GetExtension(ctx context.Context, in *GetExtensionRequest, opts ...dcerpc.CallOption) (*GetExtensionResponse, error) {
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
	out := &GetExtensionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) SetExtension(ctx context.Context, in *SetExtensionRequest, opts ...dcerpc.CallOption) (*SetExtensionResponse, error) {
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
	out := &SetExtensionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) GetConnectorTypeGUID(ctx context.Context, in *GetConnectorTypeGUIDRequest, opts ...dcerpc.CallOption) (*GetConnectorTypeGUIDResponse, error) {
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
	out := &GetConnectorTypeGUIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) SetConnectorTypeGUID(ctx context.Context, in *SetConnectorTypeGUIDRequest, opts ...dcerpc.CallOption) (*SetConnectorTypeGUIDResponse, error) {
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
	out := &SetConnectorTypeGUIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) GetTransactionStatusQueueInfo(ctx context.Context, in *GetTransactionStatusQueueInfoRequest, opts ...dcerpc.CallOption) (*GetTransactionStatusQueueInfoResponse, error) {
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
	out := &GetTransactionStatusQueueInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) GetDestinationSymmetricKey(ctx context.Context, in *GetDestinationSymmetricKeyRequest, opts ...dcerpc.CallOption) (*GetDestinationSymmetricKeyResponse, error) {
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
	out := &GetDestinationSymmetricKeyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) SetDestinationSymmetricKey(ctx context.Context, in *SetDestinationSymmetricKeyRequest, opts ...dcerpc.CallOption) (*SetDestinationSymmetricKeyResponse, error) {
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
	out := &SetDestinationSymmetricKeyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) GetSignature(ctx context.Context, in *GetSignatureRequest, opts ...dcerpc.CallOption) (*GetSignatureResponse, error) {
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
	out := &GetSignatureResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) SetSignature(ctx context.Context, in *SetSignatureRequest, opts ...dcerpc.CallOption) (*SetSignatureResponse, error) {
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
	out := &SetSignatureResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) GetAuthenticationProviderType(ctx context.Context, in *GetAuthenticationProviderTypeRequest, opts ...dcerpc.CallOption) (*GetAuthenticationProviderTypeResponse, error) {
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
	out := &GetAuthenticationProviderTypeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) SetAuthenticationProviderType(ctx context.Context, in *SetAuthenticationProviderTypeRequest, opts ...dcerpc.CallOption) (*SetAuthenticationProviderTypeResponse, error) {
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
	out := &SetAuthenticationProviderTypeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) GetAuthenticationProviderName(ctx context.Context, in *GetAuthenticationProviderNameRequest, opts ...dcerpc.CallOption) (*GetAuthenticationProviderNameResponse, error) {
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
	out := &GetAuthenticationProviderNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) SetAuthenticationProviderName(ctx context.Context, in *SetAuthenticationProviderNameRequest, opts ...dcerpc.CallOption) (*SetAuthenticationProviderNameResponse, error) {
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
	out := &SetAuthenticationProviderNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) SetSenderID(ctx context.Context, in *SetSenderIDRequest, opts ...dcerpc.CallOption) (*SetSenderIDResponse, error) {
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
	out := &SetSenderIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) GetMessageClass(ctx context.Context, in *GetMessageClassRequest, opts ...dcerpc.CallOption) (*GetMessageClassResponse, error) {
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
	out := &GetMessageClassResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) SetMessageClass(ctx context.Context, in *SetMessageClassRequest, opts ...dcerpc.CallOption) (*SetMessageClassResponse, error) {
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
	out := &SetMessageClassResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) GetProperties(ctx context.Context, in *GetPropertiesRequest, opts ...dcerpc.CallOption) (*GetPropertiesResponse, error) {
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
	out := &GetPropertiesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) GetTransactionID(ctx context.Context, in *GetTransactionIDRequest, opts ...dcerpc.CallOption) (*GetTransactionIDResponse, error) {
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
	out := &GetTransactionIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) GetIsFirstInTransaction(ctx context.Context, in *GetIsFirstInTransactionRequest, opts ...dcerpc.CallOption) (*GetIsFirstInTransactionResponse, error) {
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
	out := &GetIsFirstInTransactionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) GetIsLastInTransaction(ctx context.Context, in *GetIsLastInTransactionRequest, opts ...dcerpc.CallOption) (*GetIsLastInTransactionResponse, error) {
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
	out := &GetIsLastInTransactionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) GetResponseQueueInfoV2(ctx context.Context, in *GetResponseQueueInfoV2Request, opts ...dcerpc.CallOption) (*GetResponseQueueInfoV2Response, error) {
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
	out := &GetResponseQueueInfoV2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) SetResponseQueueInfoV2(ctx context.Context, in *SetResponseQueueInfoV2Request, opts ...dcerpc.CallOption) (*SetResponseQueueInfoV2Response, error) {
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
	out := &SetResponseQueueInfoV2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) GetAdminQueueInfoV2(ctx context.Context, in *GetAdminQueueInfoV2Request, opts ...dcerpc.CallOption) (*GetAdminQueueInfoV2Response, error) {
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
	out := &GetAdminQueueInfoV2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) SetAdminQueueInfoV2(ctx context.Context, in *SetAdminQueueInfoV2Request, opts ...dcerpc.CallOption) (*SetAdminQueueInfoV2Response, error) {
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
	out := &SetAdminQueueInfoV2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) GetReceivedAuthenticationLevel(ctx context.Context, in *GetReceivedAuthenticationLevelRequest, opts ...dcerpc.CallOption) (*GetReceivedAuthenticationLevelResponse, error) {
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
	out := &GetReceivedAuthenticationLevelResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) GetResponseQueueInfo(ctx context.Context, in *GetResponseQueueInfoRequest, opts ...dcerpc.CallOption) (*GetResponseQueueInfoResponse, error) {
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

func (o *xxx_DefaultMessage3Client) SetByRefResponseQueueInfo(ctx context.Context, in *SetByRefResponseQueueInfoRequest, opts ...dcerpc.CallOption) (*SetByRefResponseQueueInfoResponse, error) {
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

func (o *xxx_DefaultMessage3Client) GetAdminQueueInfo(ctx context.Context, in *GetAdminQueueInfoRequest, opts ...dcerpc.CallOption) (*GetAdminQueueInfoResponse, error) {
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

func (o *xxx_DefaultMessage3Client) SetByRefAdminQueueInfo(ctx context.Context, in *SetByRefAdminQueueInfoRequest, opts ...dcerpc.CallOption) (*SetByRefAdminQueueInfoResponse, error) {
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

func (o *xxx_DefaultMessage3Client) GetResponseDestination(ctx context.Context, in *GetResponseDestinationRequest, opts ...dcerpc.CallOption) (*GetResponseDestinationResponse, error) {
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
	out := &GetResponseDestinationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) SetByRefResponseDestination(ctx context.Context, in *SetByRefResponseDestinationRequest, opts ...dcerpc.CallOption) (*SetByRefResponseDestinationResponse, error) {
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
	out := &SetByRefResponseDestinationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) GetDestination(ctx context.Context, in *GetDestinationRequest, opts ...dcerpc.CallOption) (*GetDestinationResponse, error) {
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
	out := &GetDestinationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) GetLookupID(ctx context.Context, in *GetLookupIDRequest, opts ...dcerpc.CallOption) (*GetLookupIDResponse, error) {
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
	out := &GetLookupIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) GetIsAuthenticated2(ctx context.Context, in *GetIsAuthenticated2Request, opts ...dcerpc.CallOption) (*GetIsAuthenticated2Response, error) {
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
	out := &GetIsAuthenticated2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) GetIsFirstInTransaction2(ctx context.Context, in *GetIsFirstInTransaction2Request, opts ...dcerpc.CallOption) (*GetIsFirstInTransaction2Response, error) {
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
	out := &GetIsFirstInTransaction2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) GetIsLastInTransaction2(ctx context.Context, in *GetIsLastInTransaction2Request, opts ...dcerpc.CallOption) (*GetIsLastInTransaction2Response, error) {
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
	out := &GetIsLastInTransaction2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) AttachCurrentSecurityContext2(ctx context.Context, in *AttachCurrentSecurityContext2Request, opts ...dcerpc.CallOption) (*AttachCurrentSecurityContext2Response, error) {
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
	out := &AttachCurrentSecurityContext2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) GetSOAPEnvelope(ctx context.Context, in *GetSOAPEnvelopeRequest, opts ...dcerpc.CallOption) (*GetSOAPEnvelopeResponse, error) {
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
	out := &GetSOAPEnvelopeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) GetCompoundMessage(ctx context.Context, in *GetCompoundMessageRequest, opts ...dcerpc.CallOption) (*GetCompoundMessageResponse, error) {
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
	out := &GetCompoundMessageResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) SetSOAPHeader(ctx context.Context, in *SetSOAPHeaderRequest, opts ...dcerpc.CallOption) (*SetSOAPHeaderResponse, error) {
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
	out := &SetSOAPHeaderResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) SetSOAPBody(ctx context.Context, in *SetSOAPBodyRequest, opts ...dcerpc.CallOption) (*SetSOAPBodyResponse, error) {
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
	out := &SetSOAPBodyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessage3Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultMessage3Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultMessage3Client) IPID(ctx context.Context, ipid *dcom.IPID) Message3Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultMessage3Client{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewMessage3Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Message3Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(Message3SyntaxV0_0))...)
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
	return &xxx_DefaultMessage3Client{
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

func (o *xxx_GetClassOperation) OpName() string { return "/IMSMQMessage3/v0/Class" }

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

func (o *xxx_GetPrivacyLevelOperation) OpName() string { return "/IMSMQMessage3/v0/PrivLevel" }

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

func (o *xxx_SetPrivacyLevelOperation) OpName() string { return "/IMSMQMessage3/v0/PrivLevel" }

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

func (o *xxx_GetAuthLevelOperation) OpName() string { return "/IMSMQMessage3/v0/AuthLevel" }

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

func (o *xxx_SetAuthLevelOperation) OpName() string { return "/IMSMQMessage3/v0/AuthLevel" }

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

func (o *xxx_GetIsAuthenticatedOperation) OpName() string { return "/IMSMQMessage3/v0/IsAuthenticated" }

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

func (o *xxx_GetDeliveryOperation) OpName() string { return "/IMSMQMessage3/v0/Delivery" }

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

func (o *xxx_SetDeliveryOperation) OpName() string { return "/IMSMQMessage3/v0/Delivery" }

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

func (o *xxx_GetTraceOperation) OpName() string { return "/IMSMQMessage3/v0/Trace" }

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

func (o *xxx_SetTraceOperation) OpName() string { return "/IMSMQMessage3/v0/Trace" }

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

func (o *xxx_GetPriorityOperation) OpName() string { return "/IMSMQMessage3/v0/Priority" }

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

func (o *xxx_SetPriorityOperation) OpName() string { return "/IMSMQMessage3/v0/Priority" }

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

func (o *xxx_GetJournalOperation) OpName() string { return "/IMSMQMessage3/v0/Journal" }

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

func (o *xxx_SetJournalOperation) OpName() string { return "/IMSMQMessage3/v0/Journal" }

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

// xxx_GetResponseQueueInfoV1Operation structure represents the ResponseQueueInfo_v1 operation
type xxx_GetResponseQueueInfoV1Operation struct {
	This     *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Response *mqac.QueueInfo `idl:"name:ppqinfoResponse" json:"response"`
	Return   int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetResponseQueueInfoV1Operation) OpNum() int { return 21 }

func (o *xxx_GetResponseQueueInfoV1Operation) OpName() string {
	return "/IMSMQMessage3/v0/ResponseQueueInfo_v1"
}

func (o *xxx_GetResponseQueueInfoV1Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResponseQueueInfoV1Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetResponseQueueInfoV1Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetResponseQueueInfoV1Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResponseQueueInfoV1Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetResponseQueueInfoV1Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// GetResponseQueueInfoV1Request structure represents the ResponseQueueInfo_v1 operation request
type GetResponseQueueInfoV1Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetResponseQueueInfoV1Request) xxx_ToOp(ctx context.Context, op *xxx_GetResponseQueueInfoV1Operation) *xxx_GetResponseQueueInfoV1Operation {
	if op == nil {
		op = &xxx_GetResponseQueueInfoV1Operation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetResponseQueueInfoV1Request) xxx_FromOp(ctx context.Context, op *xxx_GetResponseQueueInfoV1Operation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetResponseQueueInfoV1Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetResponseQueueInfoV1Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetResponseQueueInfoV1Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetResponseQueueInfoV1Response structure represents the ResponseQueueInfo_v1 operation response
type GetResponseQueueInfoV1Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Response *mqac.QueueInfo `idl:"name:ppqinfoResponse" json:"response"`
	// Return: The ResponseQueueInfo_v1 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetResponseQueueInfoV1Response) xxx_ToOp(ctx context.Context, op *xxx_GetResponseQueueInfoV1Operation) *xxx_GetResponseQueueInfoV1Operation {
	if op == nil {
		op = &xxx_GetResponseQueueInfoV1Operation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Response = o.Response
	op.Return = o.Return
	return op
}

func (o *GetResponseQueueInfoV1Response) xxx_FromOp(ctx context.Context, op *xxx_GetResponseQueueInfoV1Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Response = op.Response
	o.Return = op.Return
}
func (o *GetResponseQueueInfoV1Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetResponseQueueInfoV1Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetResponseQueueInfoV1Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetResponseQueueInfoV1Operation structure represents the ResponseQueueInfo_v1 operation
type xxx_SetResponseQueueInfoV1Operation struct {
	This              *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat  `idl:"name:That" json:"that"`
	QueueInfoResponse *mqac.QueueInfo `idl:"name:pqinfoResponse" json:"queue_info_response"`
	Return            int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_SetResponseQueueInfoV1Operation) OpNum() int { return 22 }

func (o *xxx_SetResponseQueueInfoV1Operation) OpName() string {
	return "/IMSMQMessage3/v0/ResponseQueueInfo_v1"
}

func (o *xxx_SetResponseQueueInfoV1Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetResponseQueueInfoV1Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetResponseQueueInfoV1Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_SetResponseQueueInfoV1Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetResponseQueueInfoV1Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetResponseQueueInfoV1Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetResponseQueueInfoV1Request structure represents the ResponseQueueInfo_v1 operation request
type SetResponseQueueInfoV1Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This              *dcom.ORPCThis  `idl:"name:This" json:"this"`
	QueueInfoResponse *mqac.QueueInfo `idl:"name:pqinfoResponse" json:"queue_info_response"`
}

func (o *SetResponseQueueInfoV1Request) xxx_ToOp(ctx context.Context, op *xxx_SetResponseQueueInfoV1Operation) *xxx_SetResponseQueueInfoV1Operation {
	if op == nil {
		op = &xxx_SetResponseQueueInfoV1Operation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.QueueInfoResponse = o.QueueInfoResponse
	return op
}

func (o *SetResponseQueueInfoV1Request) xxx_FromOp(ctx context.Context, op *xxx_SetResponseQueueInfoV1Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.QueueInfoResponse = op.QueueInfoResponse
}
func (o *SetResponseQueueInfoV1Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetResponseQueueInfoV1Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetResponseQueueInfoV1Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetResponseQueueInfoV1Response structure represents the ResponseQueueInfo_v1 operation response
type SetResponseQueueInfoV1Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ResponseQueueInfo_v1 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetResponseQueueInfoV1Response) xxx_ToOp(ctx context.Context, op *xxx_SetResponseQueueInfoV1Operation) *xxx_SetResponseQueueInfoV1Operation {
	if op == nil {
		op = &xxx_SetResponseQueueInfoV1Operation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetResponseQueueInfoV1Response) xxx_FromOp(ctx context.Context, op *xxx_SetResponseQueueInfoV1Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetResponseQueueInfoV1Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetResponseQueueInfoV1Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetResponseQueueInfoV1Operation{}
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

func (o *xxx_GetAppSpecificOperation) OpName() string { return "/IMSMQMessage3/v0/AppSpecific" }

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

func (o *xxx_SetAppSpecificOperation) OpName() string { return "/IMSMQMessage3/v0/AppSpecific" }

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
	return "/IMSMQMessage3/v0/SourceMachineGuid"
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

func (o *xxx_GetBodyLengthOperation) OpName() string { return "/IMSMQMessage3/v0/BodyLength" }

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

func (o *xxx_GetBodyOperation) OpName() string { return "/IMSMQMessage3/v0/Body" }

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

func (o *xxx_SetBodyOperation) OpName() string { return "/IMSMQMessage3/v0/Body" }

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

// xxx_GetAdminQueueInfoV1Operation structure represents the AdminQueueInfo_v1 operation
type xxx_GetAdminQueueInfoV1Operation struct {
	This   *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Admin  *mqac.QueueInfo `idl:"name:ppqinfoAdmin" json:"admin"`
	Return int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAdminQueueInfoV1Operation) OpNum() int { return 29 }

func (o *xxx_GetAdminQueueInfoV1Operation) OpName() string {
	return "/IMSMQMessage3/v0/AdminQueueInfo_v1"
}

func (o *xxx_GetAdminQueueInfoV1Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAdminQueueInfoV1Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetAdminQueueInfoV1Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetAdminQueueInfoV1Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAdminQueueInfoV1Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetAdminQueueInfoV1Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// GetAdminQueueInfoV1Request structure represents the AdminQueueInfo_v1 operation request
type GetAdminQueueInfoV1Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetAdminQueueInfoV1Request) xxx_ToOp(ctx context.Context, op *xxx_GetAdminQueueInfoV1Operation) *xxx_GetAdminQueueInfoV1Operation {
	if op == nil {
		op = &xxx_GetAdminQueueInfoV1Operation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetAdminQueueInfoV1Request) xxx_FromOp(ctx context.Context, op *xxx_GetAdminQueueInfoV1Operation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetAdminQueueInfoV1Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetAdminQueueInfoV1Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAdminQueueInfoV1Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAdminQueueInfoV1Response structure represents the AdminQueueInfo_v1 operation response
type GetAdminQueueInfoV1Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Admin *mqac.QueueInfo `idl:"name:ppqinfoAdmin" json:"admin"`
	// Return: The AdminQueueInfo_v1 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAdminQueueInfoV1Response) xxx_ToOp(ctx context.Context, op *xxx_GetAdminQueueInfoV1Operation) *xxx_GetAdminQueueInfoV1Operation {
	if op == nil {
		op = &xxx_GetAdminQueueInfoV1Operation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Admin = o.Admin
	op.Return = o.Return
	return op
}

func (o *GetAdminQueueInfoV1Response) xxx_FromOp(ctx context.Context, op *xxx_GetAdminQueueInfoV1Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Admin = op.Admin
	o.Return = op.Return
}
func (o *GetAdminQueueInfoV1Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetAdminQueueInfoV1Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAdminQueueInfoV1Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetAdminQueueInfoV1Operation structure represents the AdminQueueInfo_v1 operation
type xxx_SetAdminQueueInfoV1Operation struct {
	This           *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat  `idl:"name:That" json:"that"`
	QueueInfoAdmin *mqac.QueueInfo `idl:"name:pqinfoAdmin" json:"queue_info_admin"`
	Return         int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_SetAdminQueueInfoV1Operation) OpNum() int { return 30 }

func (o *xxx_SetAdminQueueInfoV1Operation) OpName() string {
	return "/IMSMQMessage3/v0/AdminQueueInfo_v1"
}

func (o *xxx_SetAdminQueueInfoV1Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAdminQueueInfoV1Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetAdminQueueInfoV1Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_SetAdminQueueInfoV1Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAdminQueueInfoV1Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetAdminQueueInfoV1Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetAdminQueueInfoV1Request structure represents the AdminQueueInfo_v1 operation request
type SetAdminQueueInfoV1Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This           *dcom.ORPCThis  `idl:"name:This" json:"this"`
	QueueInfoAdmin *mqac.QueueInfo `idl:"name:pqinfoAdmin" json:"queue_info_admin"`
}

func (o *SetAdminQueueInfoV1Request) xxx_ToOp(ctx context.Context, op *xxx_SetAdminQueueInfoV1Operation) *xxx_SetAdminQueueInfoV1Operation {
	if op == nil {
		op = &xxx_SetAdminQueueInfoV1Operation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.QueueInfoAdmin = o.QueueInfoAdmin
	return op
}

func (o *SetAdminQueueInfoV1Request) xxx_FromOp(ctx context.Context, op *xxx_SetAdminQueueInfoV1Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.QueueInfoAdmin = op.QueueInfoAdmin
}
func (o *SetAdminQueueInfoV1Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetAdminQueueInfoV1Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAdminQueueInfoV1Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetAdminQueueInfoV1Response structure represents the AdminQueueInfo_v1 operation response
type SetAdminQueueInfoV1Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AdminQueueInfo_v1 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetAdminQueueInfoV1Response) xxx_ToOp(ctx context.Context, op *xxx_SetAdminQueueInfoV1Operation) *xxx_SetAdminQueueInfoV1Operation {
	if op == nil {
		op = &xxx_SetAdminQueueInfoV1Operation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetAdminQueueInfoV1Response) xxx_FromOp(ctx context.Context, op *xxx_SetAdminQueueInfoV1Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetAdminQueueInfoV1Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetAdminQueueInfoV1Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAdminQueueInfoV1Operation{}
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

func (o *xxx_GetIDOperation) OpName() string { return "/IMSMQMessage3/v0/Id" }

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

func (o *xxx_GetCorrelationIDOperation) OpName() string { return "/IMSMQMessage3/v0/CorrelationId" }

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

func (o *xxx_SetCorrelationIDOperation) OpName() string { return "/IMSMQMessage3/v0/CorrelationId" }

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

func (o *xxx_GetAckOperation) OpName() string { return "/IMSMQMessage3/v0/Ack" }

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

func (o *xxx_SetAckOperation) OpName() string { return "/IMSMQMessage3/v0/Ack" }

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

func (o *xxx_GetLabelOperation) OpName() string { return "/IMSMQMessage3/v0/Label" }

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

func (o *xxx_SetLabelOperation) OpName() string { return "/IMSMQMessage3/v0/Label" }

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
	return "/IMSMQMessage3/v0/MaxTimeToReachQueue"
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
	return "/IMSMQMessage3/v0/MaxTimeToReachQueue"
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
	return "/IMSMQMessage3/v0/MaxTimeToReceive"
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
	return "/IMSMQMessage3/v0/MaxTimeToReceive"
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

func (o *xxx_GetHashAlgorithmOperation) OpName() string { return "/IMSMQMessage3/v0/HashAlgorithm" }

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

func (o *xxx_SetHashAlgorithmOperation) OpName() string { return "/IMSMQMessage3/v0/HashAlgorithm" }

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
	return "/IMSMQMessage3/v0/EncryptAlgorithm"
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
	return "/IMSMQMessage3/v0/EncryptAlgorithm"
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

func (o *xxx_GetSentTimeOperation) OpName() string { return "/IMSMQMessage3/v0/SentTime" }

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

func (o *xxx_GetArrivedTimeOperation) OpName() string { return "/IMSMQMessage3/v0/ArrivedTime" }

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
	This        *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Destination *mqac.QueueInfo3 `idl:"name:ppqinfoDest" json:"destination"`
	Return      int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDestinationQueueInfoOperation) OpNum() int { return 48 }

func (o *xxx_GetDestinationQueueInfoOperation) OpName() string {
	return "/IMSMQMessage3/v0/DestinationQueueInfo"
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
	// ppqinfoDest {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQQueueInfo3}(interface))
	{
		if o.Destination != nil {
			_ptr_ppqinfoDest := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Destination != nil {
					if err := o.Destination.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.QueueInfo3{}).MarshalNDR(ctx, w); err != nil {
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
	// ppqinfoDest {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQQueueInfo3}(interface))
	{
		_ptr_ppqinfoDest := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Destination == nil {
				o.Destination = &mqac.QueueInfo3{}
			}
			if err := o.Destination.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppqinfoDest := func(ptr interface{}) { o.Destination = *ptr.(**mqac.QueueInfo3) }
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
	Destination *mqac.QueueInfo3 `idl:"name:ppqinfoDest" json:"destination"`
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
	return "/IMSMQMessage3/v0/SenderCertificate"
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
	return "/IMSMQMessage3/v0/SenderCertificate"
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

func (o *xxx_GetSenderIDOperation) OpName() string { return "/IMSMQMessage3/v0/SenderId" }

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

func (o *xxx_GetSenderIDTypeOperation) OpName() string { return "/IMSMQMessage3/v0/SenderIdType" }

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

func (o *xxx_SetSenderIDTypeOperation) OpName() string { return "/IMSMQMessage3/v0/SenderIdType" }

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
	DestinationQueue *oaut.Dispatch `idl:"name:DestinationQueue" json:"destination_queue"`
	Transaction      *oaut.Variant  `idl:"name:Transaction" json:"transaction"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SendOperation) OpNum() int { return 54 }

func (o *xxx_SendOperation) OpName() string { return "/IMSMQMessage3/v0/Send" }

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
	// DestinationQueue {in} (1:{pointer=ref}*(1))(2:{alias=IDispatch}(interface))
	{
		if o.DestinationQueue != nil {
			_ptr_DestinationQueue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DestinationQueue != nil {
					if err := o.DestinationQueue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Dispatch{}).MarshalNDR(ctx, w); err != nil {
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
	// DestinationQueue {in} (1:{pointer=ref}*(1))(2:{alias=IDispatch}(interface))
	{
		_ptr_DestinationQueue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DestinationQueue == nil {
				o.DestinationQueue = &oaut.Dispatch{}
			}
			if err := o.DestinationQueue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_DestinationQueue := func(ptr interface{}) { o.DestinationQueue = *ptr.(**oaut.Dispatch) }
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
	DestinationQueue *oaut.Dispatch `idl:"name:DestinationQueue" json:"destination_queue"`
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
	return "/IMSMQMessage3/v0/AttachCurrentSecurityContext"
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

// xxx_GetSenderVersionOperation structure represents the SenderVersion operation
type xxx_GetSenderVersionOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	SenderVersion int32          `idl:"name:plSenderVersion" json:"sender_version"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSenderVersionOperation) OpNum() int { return 56 }

func (o *xxx_GetSenderVersionOperation) OpName() string { return "/IMSMQMessage3/v0/SenderVersion" }

func (o *xxx_GetSenderVersionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSenderVersionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSenderVersionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSenderVersionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSenderVersionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plSenderVersion {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.SenderVersion); err != nil {
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

func (o *xxx_GetSenderVersionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plSenderVersion {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.SenderVersion); err != nil {
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

// GetSenderVersionRequest structure represents the SenderVersion operation request
type GetSenderVersionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSenderVersionRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSenderVersionOperation) *xxx_GetSenderVersionOperation {
	if op == nil {
		op = &xxx_GetSenderVersionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetSenderVersionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSenderVersionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSenderVersionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSenderVersionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSenderVersionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSenderVersionResponse structure represents the SenderVersion operation response
type GetSenderVersionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// plSenderVersion: A pointer to a long integer that specifies the version of transfer
	// used to send the message.
	SenderVersion int32 `idl:"name:plSenderVersion" json:"sender_version"`
	// Return: The SenderVersion return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSenderVersionResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSenderVersionOperation) *xxx_GetSenderVersionOperation {
	if op == nil {
		op = &xxx_GetSenderVersionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.SenderVersion = o.SenderVersion
	op.Return = o.Return
	return op
}

func (o *GetSenderVersionResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSenderVersionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SenderVersion = op.SenderVersion
	o.Return = op.Return
}
func (o *GetSenderVersionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSenderVersionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSenderVersionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetExtensionOperation structure represents the Extension operation
type xxx_GetExtensionOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Extension *oaut.Variant  `idl:"name:pvarExtension" json:"extension"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetExtensionOperation) OpNum() int { return 57 }

func (o *xxx_GetExtensionOperation) OpName() string { return "/IMSMQMessage3/v0/Extension" }

func (o *xxx_GetExtensionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetExtensionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetExtensionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetExtensionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetExtensionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pvarExtension {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Extension != nil {
			_ptr_pvarExtension := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Extension != nil {
					if err := o.Extension.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Extension, _ptr_pvarExtension); err != nil {
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

func (o *xxx_GetExtensionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pvarExtension {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pvarExtension := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Extension == nil {
				o.Extension = &oaut.Variant{}
			}
			if err := o.Extension.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pvarExtension := func(ptr interface{}) { o.Extension = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Extension, _s_pvarExtension, _ptr_pvarExtension); err != nil {
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

// GetExtensionRequest structure represents the Extension operation request
type GetExtensionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetExtensionRequest) xxx_ToOp(ctx context.Context, op *xxx_GetExtensionOperation) *xxx_GetExtensionOperation {
	if op == nil {
		op = &xxx_GetExtensionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetExtensionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetExtensionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetExtensionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetExtensionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetExtensionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetExtensionResponse structure represents the Extension operation response
type GetExtensionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Extension *oaut.Variant  `idl:"name:pvarExtension" json:"extension"`
	// Return: The Extension return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetExtensionResponse) xxx_ToOp(ctx context.Context, op *xxx_GetExtensionOperation) *xxx_GetExtensionOperation {
	if op == nil {
		op = &xxx_GetExtensionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Extension = o.Extension
	op.Return = o.Return
	return op
}

func (o *GetExtensionResponse) xxx_FromOp(ctx context.Context, op *xxx_GetExtensionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Extension = op.Extension
	o.Return = op.Return
}
func (o *GetExtensionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetExtensionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetExtensionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetExtensionOperation structure represents the Extension operation
type xxx_SetExtensionOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Extension *oaut.Variant  `idl:"name:varExtension" json:"extension"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetExtensionOperation) OpNum() int { return 58 }

func (o *xxx_SetExtensionOperation) OpName() string { return "/IMSMQMessage3/v0/Extension" }

func (o *xxx_SetExtensionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExtensionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// varExtension {in} (1:{alias=VARIANT}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.Extension != nil {
			if err := o.Extension.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_SetExtensionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// varExtension {in} (1:{alias=VARIANT,pointer=ref}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.Extension == nil {
			o.Extension = &oaut.Variant{}
		}
		if err := o.Extension.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExtensionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExtensionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetExtensionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetExtensionRequest structure represents the Extension operation request
type SetExtensionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	Extension *oaut.Variant  `idl:"name:varExtension" json:"extension"`
}

func (o *SetExtensionRequest) xxx_ToOp(ctx context.Context, op *xxx_SetExtensionOperation) *xxx_SetExtensionOperation {
	if op == nil {
		op = &xxx_SetExtensionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Extension = o.Extension
	return op
}

func (o *SetExtensionRequest) xxx_FromOp(ctx context.Context, op *xxx_SetExtensionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Extension = op.Extension
}
func (o *SetExtensionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetExtensionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetExtensionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetExtensionResponse structure represents the Extension operation response
type SetExtensionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Extension return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetExtensionResponse) xxx_ToOp(ctx context.Context, op *xxx_SetExtensionOperation) *xxx_SetExtensionOperation {
	if op == nil {
		op = &xxx_SetExtensionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetExtensionResponse) xxx_FromOp(ctx context.Context, op *xxx_SetExtensionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetExtensionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetExtensionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetExtensionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetConnectorTypeGUIDOperation structure represents the ConnectorTypeGuid operation
type xxx_GetConnectorTypeGUIDOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	ConnectorType *oaut.String   `idl:"name:pbstrGuidConnectorType" json:"connector_type"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetConnectorTypeGUIDOperation) OpNum() int { return 59 }

func (o *xxx_GetConnectorTypeGUIDOperation) OpName() string {
	return "/IMSMQMessage3/v0/ConnectorTypeGuid"
}

func (o *xxx_GetConnectorTypeGUIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConnectorTypeGUIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetConnectorTypeGUIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetConnectorTypeGUIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConnectorTypeGUIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrGuidConnectorType {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ConnectorType != nil {
			_ptr_pbstrGuidConnectorType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ConnectorType != nil {
					if err := o.ConnectorType.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConnectorType, _ptr_pbstrGuidConnectorType); err != nil {
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

func (o *xxx_GetConnectorTypeGUIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrGuidConnectorType {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrGuidConnectorType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConnectorType == nil {
				o.ConnectorType = &oaut.String{}
			}
			if err := o.ConnectorType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrGuidConnectorType := func(ptr interface{}) { o.ConnectorType = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ConnectorType, _s_pbstrGuidConnectorType, _ptr_pbstrGuidConnectorType); err != nil {
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

// GetConnectorTypeGUIDRequest structure represents the ConnectorTypeGuid operation request
type GetConnectorTypeGUIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetConnectorTypeGUIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetConnectorTypeGUIDOperation) *xxx_GetConnectorTypeGUIDOperation {
	if op == nil {
		op = &xxx_GetConnectorTypeGUIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetConnectorTypeGUIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetConnectorTypeGUIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetConnectorTypeGUIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetConnectorTypeGUIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetConnectorTypeGUIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetConnectorTypeGUIDResponse structure represents the ConnectorTypeGuid operation response
type GetConnectorTypeGUIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	ConnectorType *oaut.String   `idl:"name:pbstrGuidConnectorType" json:"connector_type"`
	// Return: The ConnectorTypeGuid return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetConnectorTypeGUIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetConnectorTypeGUIDOperation) *xxx_GetConnectorTypeGUIDOperation {
	if op == nil {
		op = &xxx_GetConnectorTypeGUIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ConnectorType = o.ConnectorType
	op.Return = o.Return
	return op
}

func (o *GetConnectorTypeGUIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetConnectorTypeGUIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ConnectorType = op.ConnectorType
	o.Return = op.Return
}
func (o *GetConnectorTypeGUIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetConnectorTypeGUIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetConnectorTypeGUIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetConnectorTypeGUIDOperation structure represents the ConnectorTypeGuid operation
type xxx_SetConnectorTypeGUIDOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	ConnectorType *oaut.String   `idl:"name:bstrGuidConnectorType" json:"connector_type"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetConnectorTypeGUIDOperation) OpNum() int { return 60 }

func (o *xxx_SetConnectorTypeGUIDOperation) OpName() string {
	return "/IMSMQMessage3/v0/ConnectorTypeGuid"
}

func (o *xxx_SetConnectorTypeGUIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetConnectorTypeGUIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrGuidConnectorType {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ConnectorType != nil {
			_ptr_bstrGuidConnectorType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ConnectorType != nil {
					if err := o.ConnectorType.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConnectorType, _ptr_bstrGuidConnectorType); err != nil {
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

func (o *xxx_SetConnectorTypeGUIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrGuidConnectorType {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrGuidConnectorType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConnectorType == nil {
				o.ConnectorType = &oaut.String{}
			}
			if err := o.ConnectorType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrGuidConnectorType := func(ptr interface{}) { o.ConnectorType = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ConnectorType, _s_bstrGuidConnectorType, _ptr_bstrGuidConnectorType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetConnectorTypeGUIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetConnectorTypeGUIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetConnectorTypeGUIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetConnectorTypeGUIDRequest structure represents the ConnectorTypeGuid operation request
type SetConnectorTypeGUIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	ConnectorType *oaut.String   `idl:"name:bstrGuidConnectorType" json:"connector_type"`
}

func (o *SetConnectorTypeGUIDRequest) xxx_ToOp(ctx context.Context, op *xxx_SetConnectorTypeGUIDOperation) *xxx_SetConnectorTypeGUIDOperation {
	if op == nil {
		op = &xxx_SetConnectorTypeGUIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ConnectorType = o.ConnectorType
	return op
}

func (o *SetConnectorTypeGUIDRequest) xxx_FromOp(ctx context.Context, op *xxx_SetConnectorTypeGUIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ConnectorType = op.ConnectorType
}
func (o *SetConnectorTypeGUIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetConnectorTypeGUIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetConnectorTypeGUIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetConnectorTypeGUIDResponse structure represents the ConnectorTypeGuid operation response
type SetConnectorTypeGUIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ConnectorTypeGuid return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetConnectorTypeGUIDResponse) xxx_ToOp(ctx context.Context, op *xxx_SetConnectorTypeGUIDOperation) *xxx_SetConnectorTypeGUIDOperation {
	if op == nil {
		op = &xxx_SetConnectorTypeGUIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetConnectorTypeGUIDResponse) xxx_FromOp(ctx context.Context, op *xxx_SetConnectorTypeGUIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetConnectorTypeGUIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetConnectorTypeGUIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetConnectorTypeGUIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetTransactionStatusQueueInfoOperation structure represents the TransactionStatusQueueInfo operation
type xxx_GetTransactionStatusQueueInfoOperation struct {
	This              *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat   `idl:"name:That" json:"that"`
	TransactionStatus *mqac.QueueInfo3 `idl:"name:ppqinfoXactStatus" json:"transaction_status"`
	Return            int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTransactionStatusQueueInfoOperation) OpNum() int { return 61 }

func (o *xxx_GetTransactionStatusQueueInfoOperation) OpName() string {
	return "/IMSMQMessage3/v0/TransactionStatusQueueInfo"
}

func (o *xxx_GetTransactionStatusQueueInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTransactionStatusQueueInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetTransactionStatusQueueInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetTransactionStatusQueueInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTransactionStatusQueueInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppqinfoXactStatus {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQQueueInfo3}(interface))
	{
		if o.TransactionStatus != nil {
			_ptr_ppqinfoXactStatus := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.TransactionStatus != nil {
					if err := o.TransactionStatus.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.QueueInfo3{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.TransactionStatus, _ptr_ppqinfoXactStatus); err != nil {
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

func (o *xxx_GetTransactionStatusQueueInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppqinfoXactStatus {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQQueueInfo3}(interface))
	{
		_ptr_ppqinfoXactStatus := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.TransactionStatus == nil {
				o.TransactionStatus = &mqac.QueueInfo3{}
			}
			if err := o.TransactionStatus.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppqinfoXactStatus := func(ptr interface{}) { o.TransactionStatus = *ptr.(**mqac.QueueInfo3) }
		if err := w.ReadPointer(&o.TransactionStatus, _s_ppqinfoXactStatus, _ptr_ppqinfoXactStatus); err != nil {
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

// GetTransactionStatusQueueInfoRequest structure represents the TransactionStatusQueueInfo operation request
type GetTransactionStatusQueueInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetTransactionStatusQueueInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_GetTransactionStatusQueueInfoOperation) *xxx_GetTransactionStatusQueueInfoOperation {
	if op == nil {
		op = &xxx_GetTransactionStatusQueueInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetTransactionStatusQueueInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTransactionStatusQueueInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetTransactionStatusQueueInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetTransactionStatusQueueInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTransactionStatusQueueInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetTransactionStatusQueueInfoResponse structure represents the TransactionStatusQueueInfo operation response
type GetTransactionStatusQueueInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppqinfoXactStatus: A pointer to an IMSMQQueueInfo interface pointer that upon successful
	// completion will contain an MSMQQueue object representing the transaction status queue
	// for the message.
	TransactionStatus *mqac.QueueInfo3 `idl:"name:ppqinfoXactStatus" json:"transaction_status"`
	// Return: The TransactionStatusQueueInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTransactionStatusQueueInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_GetTransactionStatusQueueInfoOperation) *xxx_GetTransactionStatusQueueInfoOperation {
	if op == nil {
		op = &xxx_GetTransactionStatusQueueInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.TransactionStatus = o.TransactionStatus
	op.Return = o.Return
	return op
}

func (o *GetTransactionStatusQueueInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTransactionStatusQueueInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TransactionStatus = op.TransactionStatus
	o.Return = op.Return
}
func (o *GetTransactionStatusQueueInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetTransactionStatusQueueInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTransactionStatusQueueInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDestinationSymmetricKeyOperation structure represents the DestinationSymmetricKey operation
type xxx_GetDestinationSymmetricKeyOperation struct {
	This                    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                    *dcom.ORPCThat `idl:"name:That" json:"that"`
	DestinationSymmetricKey *oaut.Variant  `idl:"name:pvarDestSymmKey" json:"destination_symmetric_key"`
	Return                  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDestinationSymmetricKeyOperation) OpNum() int { return 62 }

func (o *xxx_GetDestinationSymmetricKeyOperation) OpName() string {
	return "/IMSMQMessage3/v0/DestinationSymmetricKey"
}

func (o *xxx_GetDestinationSymmetricKeyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDestinationSymmetricKeyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetDestinationSymmetricKeyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetDestinationSymmetricKeyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDestinationSymmetricKeyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pvarDestSymmKey {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.DestinationSymmetricKey != nil {
			_ptr_pvarDestSymmKey := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DestinationSymmetricKey != nil {
					if err := o.DestinationSymmetricKey.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DestinationSymmetricKey, _ptr_pvarDestSymmKey); err != nil {
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

func (o *xxx_GetDestinationSymmetricKeyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pvarDestSymmKey {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pvarDestSymmKey := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DestinationSymmetricKey == nil {
				o.DestinationSymmetricKey = &oaut.Variant{}
			}
			if err := o.DestinationSymmetricKey.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pvarDestSymmKey := func(ptr interface{}) { o.DestinationSymmetricKey = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.DestinationSymmetricKey, _s_pvarDestSymmKey, _ptr_pvarDestSymmKey); err != nil {
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

// GetDestinationSymmetricKeyRequest structure represents the DestinationSymmetricKey operation request
type GetDestinationSymmetricKeyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetDestinationSymmetricKeyRequest) xxx_ToOp(ctx context.Context, op *xxx_GetDestinationSymmetricKeyOperation) *xxx_GetDestinationSymmetricKeyOperation {
	if op == nil {
		op = &xxx_GetDestinationSymmetricKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetDestinationSymmetricKeyRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDestinationSymmetricKeyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetDestinationSymmetricKeyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetDestinationSymmetricKeyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDestinationSymmetricKeyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDestinationSymmetricKeyResponse structure represents the DestinationSymmetricKey operation response
type GetDestinationSymmetricKeyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That                    *dcom.ORPCThat `idl:"name:That" json:"that"`
	DestinationSymmetricKey *oaut.Variant  `idl:"name:pvarDestSymmKey" json:"destination_symmetric_key"`
	// Return: The DestinationSymmetricKey return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDestinationSymmetricKeyResponse) xxx_ToOp(ctx context.Context, op *xxx_GetDestinationSymmetricKeyOperation) *xxx_GetDestinationSymmetricKeyOperation {
	if op == nil {
		op = &xxx_GetDestinationSymmetricKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.DestinationSymmetricKey = o.DestinationSymmetricKey
	op.Return = o.Return
	return op
}

func (o *GetDestinationSymmetricKeyResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDestinationSymmetricKeyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DestinationSymmetricKey = op.DestinationSymmetricKey
	o.Return = op.Return
}
func (o *GetDestinationSymmetricKeyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetDestinationSymmetricKeyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDestinationSymmetricKeyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetDestinationSymmetricKeyOperation structure represents the DestinationSymmetricKey operation
type xxx_SetDestinationSymmetricKeyOperation struct {
	This                    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                    *dcom.ORPCThat `idl:"name:That" json:"that"`
	DestinationSymmetricKey *oaut.Variant  `idl:"name:varDestSymmKey" json:"destination_symmetric_key"`
	Return                  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetDestinationSymmetricKeyOperation) OpNum() int { return 63 }

func (o *xxx_SetDestinationSymmetricKeyOperation) OpName() string {
	return "/IMSMQMessage3/v0/DestinationSymmetricKey"
}

func (o *xxx_SetDestinationSymmetricKeyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDestinationSymmetricKeyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// varDestSymmKey {in} (1:{alias=VARIANT}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.DestinationSymmetricKey != nil {
			if err := o.DestinationSymmetricKey.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_SetDestinationSymmetricKeyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// varDestSymmKey {in} (1:{alias=VARIANT,pointer=ref}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.DestinationSymmetricKey == nil {
			o.DestinationSymmetricKey = &oaut.Variant{}
		}
		if err := o.DestinationSymmetricKey.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDestinationSymmetricKeyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDestinationSymmetricKeyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetDestinationSymmetricKeyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetDestinationSymmetricKeyRequest structure represents the DestinationSymmetricKey operation request
type SetDestinationSymmetricKeyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                    *dcom.ORPCThis `idl:"name:This" json:"this"`
	DestinationSymmetricKey *oaut.Variant  `idl:"name:varDestSymmKey" json:"destination_symmetric_key"`
}

func (o *SetDestinationSymmetricKeyRequest) xxx_ToOp(ctx context.Context, op *xxx_SetDestinationSymmetricKeyOperation) *xxx_SetDestinationSymmetricKeyOperation {
	if op == nil {
		op = &xxx_SetDestinationSymmetricKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.DestinationSymmetricKey = o.DestinationSymmetricKey
	return op
}

func (o *SetDestinationSymmetricKeyRequest) xxx_FromOp(ctx context.Context, op *xxx_SetDestinationSymmetricKeyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DestinationSymmetricKey = op.DestinationSymmetricKey
}
func (o *SetDestinationSymmetricKeyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetDestinationSymmetricKeyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDestinationSymmetricKeyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetDestinationSymmetricKeyResponse structure represents the DestinationSymmetricKey operation response
type SetDestinationSymmetricKeyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DestinationSymmetricKey return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetDestinationSymmetricKeyResponse) xxx_ToOp(ctx context.Context, op *xxx_SetDestinationSymmetricKeyOperation) *xxx_SetDestinationSymmetricKeyOperation {
	if op == nil {
		op = &xxx_SetDestinationSymmetricKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetDestinationSymmetricKeyResponse) xxx_FromOp(ctx context.Context, op *xxx_SetDestinationSymmetricKeyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetDestinationSymmetricKeyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetDestinationSymmetricKeyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDestinationSymmetricKeyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSignatureOperation structure represents the Signature operation
type xxx_GetSignatureOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Signature *oaut.Variant  `idl:"name:pvarSignature" json:"signature"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSignatureOperation) OpNum() int { return 64 }

func (o *xxx_GetSignatureOperation) OpName() string { return "/IMSMQMessage3/v0/Signature" }

func (o *xxx_GetSignatureOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSignatureOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSignatureOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSignatureOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSignatureOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pvarSignature {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Signature != nil {
			_ptr_pvarSignature := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Signature != nil {
					if err := o.Signature.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Signature, _ptr_pvarSignature); err != nil {
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

func (o *xxx_GetSignatureOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pvarSignature {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pvarSignature := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Signature == nil {
				o.Signature = &oaut.Variant{}
			}
			if err := o.Signature.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pvarSignature := func(ptr interface{}) { o.Signature = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Signature, _s_pvarSignature, _ptr_pvarSignature); err != nil {
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

// GetSignatureRequest structure represents the Signature operation request
type GetSignatureRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSignatureRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSignatureOperation) *xxx_GetSignatureOperation {
	if op == nil {
		op = &xxx_GetSignatureOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetSignatureRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSignatureOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSignatureRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSignatureRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSignatureOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSignatureResponse structure represents the Signature operation response
type GetSignatureResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Signature *oaut.Variant  `idl:"name:pvarSignature" json:"signature"`
	// Return: The Signature return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSignatureResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSignatureOperation) *xxx_GetSignatureOperation {
	if op == nil {
		op = &xxx_GetSignatureOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Signature = o.Signature
	op.Return = o.Return
	return op
}

func (o *GetSignatureResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSignatureOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Signature = op.Signature
	o.Return = op.Return
}
func (o *GetSignatureResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSignatureResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSignatureOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSignatureOperation structure represents the Signature operation
type xxx_SetSignatureOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Signature *oaut.Variant  `idl:"name:varSignature" json:"signature"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSignatureOperation) OpNum() int { return 65 }

func (o *xxx_SetSignatureOperation) OpName() string { return "/IMSMQMessage3/v0/Signature" }

func (o *xxx_SetSignatureOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSignatureOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// varSignature {in} (1:{alias=VARIANT}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.Signature != nil {
			if err := o.Signature.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_SetSignatureOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// varSignature {in} (1:{alias=VARIANT,pointer=ref}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.Signature == nil {
			o.Signature = &oaut.Variant{}
		}
		if err := o.Signature.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSignatureOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSignatureOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetSignatureOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetSignatureRequest structure represents the Signature operation request
type SetSignatureRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	Signature *oaut.Variant  `idl:"name:varSignature" json:"signature"`
}

func (o *SetSignatureRequest) xxx_ToOp(ctx context.Context, op *xxx_SetSignatureOperation) *xxx_SetSignatureOperation {
	if op == nil {
		op = &xxx_SetSignatureOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Signature = o.Signature
	return op
}

func (o *SetSignatureRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSignatureOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Signature = op.Signature
}
func (o *SetSignatureRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetSignatureRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSignatureOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSignatureResponse structure represents the Signature operation response
type SetSignatureResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Signature return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSignatureResponse) xxx_ToOp(ctx context.Context, op *xxx_SetSignatureOperation) *xxx_SetSignatureOperation {
	if op == nil {
		op = &xxx_SetSignatureOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetSignatureResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSignatureOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSignatureResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetSignatureResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSignatureOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAuthenticationProviderTypeOperation structure represents the AuthenticationProviderType operation
type xxx_GetAuthenticationProviderTypeOperation struct {
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	AuthProviderType int32          `idl:"name:plAuthProvType" json:"auth_provider_type"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAuthenticationProviderTypeOperation) OpNum() int { return 66 }

func (o *xxx_GetAuthenticationProviderTypeOperation) OpName() string {
	return "/IMSMQMessage3/v0/AuthenticationProviderType"
}

func (o *xxx_GetAuthenticationProviderTypeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAuthenticationProviderTypeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetAuthenticationProviderTypeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetAuthenticationProviderTypeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAuthenticationProviderTypeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plAuthProvType {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.AuthProviderType); err != nil {
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

func (o *xxx_GetAuthenticationProviderTypeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plAuthProvType {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.AuthProviderType); err != nil {
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

// GetAuthenticationProviderTypeRequest structure represents the AuthenticationProviderType operation request
type GetAuthenticationProviderTypeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetAuthenticationProviderTypeRequest) xxx_ToOp(ctx context.Context, op *xxx_GetAuthenticationProviderTypeOperation) *xxx_GetAuthenticationProviderTypeOperation {
	if op == nil {
		op = &xxx_GetAuthenticationProviderTypeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetAuthenticationProviderTypeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAuthenticationProviderTypeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetAuthenticationProviderTypeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetAuthenticationProviderTypeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAuthenticationProviderTypeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAuthenticationProviderTypeResponse structure represents the AuthenticationProviderType operation response
type GetAuthenticationProviderTypeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	AuthProviderType int32          `idl:"name:plAuthProvType" json:"auth_provider_type"`
	// Return: The AuthenticationProviderType return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAuthenticationProviderTypeResponse) xxx_ToOp(ctx context.Context, op *xxx_GetAuthenticationProviderTypeOperation) *xxx_GetAuthenticationProviderTypeOperation {
	if op == nil {
		op = &xxx_GetAuthenticationProviderTypeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.AuthProviderType = o.AuthProviderType
	op.Return = o.Return
	return op
}

func (o *GetAuthenticationProviderTypeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAuthenticationProviderTypeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.AuthProviderType = op.AuthProviderType
	o.Return = op.Return
}
func (o *GetAuthenticationProviderTypeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetAuthenticationProviderTypeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAuthenticationProviderTypeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetAuthenticationProviderTypeOperation structure represents the AuthenticationProviderType operation
type xxx_SetAuthenticationProviderTypeOperation struct {
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	AuthProviderType int32          `idl:"name:lAuthProvType" json:"auth_provider_type"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetAuthenticationProviderTypeOperation) OpNum() int { return 67 }

func (o *xxx_SetAuthenticationProviderTypeOperation) OpName() string {
	return "/IMSMQMessage3/v0/AuthenticationProviderType"
}

func (o *xxx_SetAuthenticationProviderTypeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAuthenticationProviderTypeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lAuthProvType {in} (1:(int32))
	{
		if err := w.WriteData(o.AuthProviderType); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAuthenticationProviderTypeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lAuthProvType {in} (1:(int32))
	{
		if err := w.ReadData(&o.AuthProviderType); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAuthenticationProviderTypeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAuthenticationProviderTypeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetAuthenticationProviderTypeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetAuthenticationProviderTypeRequest structure represents the AuthenticationProviderType operation request
type SetAuthenticationProviderTypeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	AuthProviderType int32          `idl:"name:lAuthProvType" json:"auth_provider_type"`
}

func (o *SetAuthenticationProviderTypeRequest) xxx_ToOp(ctx context.Context, op *xxx_SetAuthenticationProviderTypeOperation) *xxx_SetAuthenticationProviderTypeOperation {
	if op == nil {
		op = &xxx_SetAuthenticationProviderTypeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.AuthProviderType = o.AuthProviderType
	return op
}

func (o *SetAuthenticationProviderTypeRequest) xxx_FromOp(ctx context.Context, op *xxx_SetAuthenticationProviderTypeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.AuthProviderType = op.AuthProviderType
}
func (o *SetAuthenticationProviderTypeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetAuthenticationProviderTypeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAuthenticationProviderTypeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetAuthenticationProviderTypeResponse structure represents the AuthenticationProviderType operation response
type SetAuthenticationProviderTypeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AuthenticationProviderType return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetAuthenticationProviderTypeResponse) xxx_ToOp(ctx context.Context, op *xxx_SetAuthenticationProviderTypeOperation) *xxx_SetAuthenticationProviderTypeOperation {
	if op == nil {
		op = &xxx_SetAuthenticationProviderTypeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetAuthenticationProviderTypeResponse) xxx_FromOp(ctx context.Context, op *xxx_SetAuthenticationProviderTypeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetAuthenticationProviderTypeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetAuthenticationProviderTypeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAuthenticationProviderTypeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAuthenticationProviderNameOperation structure represents the AuthenticationProviderName operation
type xxx_GetAuthenticationProviderNameOperation struct {
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	AuthProviderName *oaut.String   `idl:"name:pbstrAuthProvName" json:"auth_provider_name"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAuthenticationProviderNameOperation) OpNum() int { return 68 }

func (o *xxx_GetAuthenticationProviderNameOperation) OpName() string {
	return "/IMSMQMessage3/v0/AuthenticationProviderName"
}

func (o *xxx_GetAuthenticationProviderNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAuthenticationProviderNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetAuthenticationProviderNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetAuthenticationProviderNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAuthenticationProviderNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrAuthProvName {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.AuthProviderName != nil {
			_ptr_pbstrAuthProvName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.AuthProviderName != nil {
					if err := o.AuthProviderName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.AuthProviderName, _ptr_pbstrAuthProvName); err != nil {
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

func (o *xxx_GetAuthenticationProviderNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrAuthProvName {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrAuthProvName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.AuthProviderName == nil {
				o.AuthProviderName = &oaut.String{}
			}
			if err := o.AuthProviderName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrAuthProvName := func(ptr interface{}) { o.AuthProviderName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.AuthProviderName, _s_pbstrAuthProvName, _ptr_pbstrAuthProvName); err != nil {
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

// GetAuthenticationProviderNameRequest structure represents the AuthenticationProviderName operation request
type GetAuthenticationProviderNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetAuthenticationProviderNameRequest) xxx_ToOp(ctx context.Context, op *xxx_GetAuthenticationProviderNameOperation) *xxx_GetAuthenticationProviderNameOperation {
	if op == nil {
		op = &xxx_GetAuthenticationProviderNameOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetAuthenticationProviderNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAuthenticationProviderNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetAuthenticationProviderNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetAuthenticationProviderNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAuthenticationProviderNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAuthenticationProviderNameResponse structure represents the AuthenticationProviderName operation response
type GetAuthenticationProviderNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	AuthProviderName *oaut.String   `idl:"name:pbstrAuthProvName" json:"auth_provider_name"`
	// Return: The AuthenticationProviderName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAuthenticationProviderNameResponse) xxx_ToOp(ctx context.Context, op *xxx_GetAuthenticationProviderNameOperation) *xxx_GetAuthenticationProviderNameOperation {
	if op == nil {
		op = &xxx_GetAuthenticationProviderNameOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.AuthProviderName = o.AuthProviderName
	op.Return = o.Return
	return op
}

func (o *GetAuthenticationProviderNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAuthenticationProviderNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.AuthProviderName = op.AuthProviderName
	o.Return = op.Return
}
func (o *GetAuthenticationProviderNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetAuthenticationProviderNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAuthenticationProviderNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetAuthenticationProviderNameOperation structure represents the AuthenticationProviderName operation
type xxx_SetAuthenticationProviderNameOperation struct {
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	AuthProviderName *oaut.String   `idl:"name:bstrAuthProvName" json:"auth_provider_name"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetAuthenticationProviderNameOperation) OpNum() int { return 69 }

func (o *xxx_SetAuthenticationProviderNameOperation) OpName() string {
	return "/IMSMQMessage3/v0/AuthenticationProviderName"
}

func (o *xxx_SetAuthenticationProviderNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAuthenticationProviderNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrAuthProvName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.AuthProviderName != nil {
			_ptr_bstrAuthProvName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.AuthProviderName != nil {
					if err := o.AuthProviderName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.AuthProviderName, _ptr_bstrAuthProvName); err != nil {
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

func (o *xxx_SetAuthenticationProviderNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrAuthProvName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrAuthProvName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.AuthProviderName == nil {
				o.AuthProviderName = &oaut.String{}
			}
			if err := o.AuthProviderName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrAuthProvName := func(ptr interface{}) { o.AuthProviderName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.AuthProviderName, _s_bstrAuthProvName, _ptr_bstrAuthProvName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAuthenticationProviderNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAuthenticationProviderNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetAuthenticationProviderNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetAuthenticationProviderNameRequest structure represents the AuthenticationProviderName operation request
type SetAuthenticationProviderNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	AuthProviderName *oaut.String   `idl:"name:bstrAuthProvName" json:"auth_provider_name"`
}

func (o *SetAuthenticationProviderNameRequest) xxx_ToOp(ctx context.Context, op *xxx_SetAuthenticationProviderNameOperation) *xxx_SetAuthenticationProviderNameOperation {
	if op == nil {
		op = &xxx_SetAuthenticationProviderNameOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.AuthProviderName = o.AuthProviderName
	return op
}

func (o *SetAuthenticationProviderNameRequest) xxx_FromOp(ctx context.Context, op *xxx_SetAuthenticationProviderNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.AuthProviderName = op.AuthProviderName
}
func (o *SetAuthenticationProviderNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetAuthenticationProviderNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAuthenticationProviderNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetAuthenticationProviderNameResponse structure represents the AuthenticationProviderName operation response
type SetAuthenticationProviderNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AuthenticationProviderName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetAuthenticationProviderNameResponse) xxx_ToOp(ctx context.Context, op *xxx_SetAuthenticationProviderNameOperation) *xxx_SetAuthenticationProviderNameOperation {
	if op == nil {
		op = &xxx_SetAuthenticationProviderNameOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetAuthenticationProviderNameResponse) xxx_FromOp(ctx context.Context, op *xxx_SetAuthenticationProviderNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetAuthenticationProviderNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetAuthenticationProviderNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAuthenticationProviderNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSenderIDOperation structure represents the SenderId operation
type xxx_SetSenderIDOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	SenderID *oaut.Variant  `idl:"name:varSenderId" json:"sender_id"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSenderIDOperation) OpNum() int { return 70 }

func (o *xxx_SetSenderIDOperation) OpName() string { return "/IMSMQMessage3/v0/SenderId" }

func (o *xxx_SetSenderIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSenderIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// varSenderId {in} (1:{alias=VARIANT}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.SenderID != nil {
			if err := o.SenderID.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_SetSenderIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// varSenderId {in} (1:{alias=VARIANT,pointer=ref}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.SenderID == nil {
			o.SenderID = &oaut.Variant{}
		}
		if err := o.SenderID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSenderIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSenderIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetSenderIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetSenderIDRequest structure represents the SenderId operation request
type SetSenderIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	SenderID *oaut.Variant  `idl:"name:varSenderId" json:"sender_id"`
}

func (o *SetSenderIDRequest) xxx_ToOp(ctx context.Context, op *xxx_SetSenderIDOperation) *xxx_SetSenderIDOperation {
	if op == nil {
		op = &xxx_SetSenderIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.SenderID = o.SenderID
	return op
}

func (o *SetSenderIDRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSenderIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SenderID = op.SenderID
}
func (o *SetSenderIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetSenderIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSenderIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSenderIDResponse structure represents the SenderId operation response
type SetSenderIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SenderId return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSenderIDResponse) xxx_ToOp(ctx context.Context, op *xxx_SetSenderIDOperation) *xxx_SetSenderIDOperation {
	if op == nil {
		op = &xxx_SetSenderIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetSenderIDResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSenderIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSenderIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetSenderIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSenderIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMessageClassOperation structure represents the MsgClass operation
type xxx_GetMessageClassOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	MessageClass int32          `idl:"name:plMsgClass" json:"message_class"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMessageClassOperation) OpNum() int { return 71 }

func (o *xxx_GetMessageClassOperation) OpName() string { return "/IMSMQMessage3/v0/MsgClass" }

func (o *xxx_GetMessageClassOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMessageClassOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetMessageClassOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetMessageClassOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMessageClassOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plMsgClass {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.MessageClass); err != nil {
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

func (o *xxx_GetMessageClassOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plMsgClass {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.MessageClass); err != nil {
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

// GetMessageClassRequest structure represents the MsgClass operation request
type GetMessageClassRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetMessageClassRequest) xxx_ToOp(ctx context.Context, op *xxx_GetMessageClassOperation) *xxx_GetMessageClassOperation {
	if op == nil {
		op = &xxx_GetMessageClassOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetMessageClassRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMessageClassOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetMessageClassRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetMessageClassRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMessageClassOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMessageClassResponse structure represents the MsgClass operation response
type GetMessageClassResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	MessageClass int32          `idl:"name:plMsgClass" json:"message_class"`
	// Return: The MsgClass return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMessageClassResponse) xxx_ToOp(ctx context.Context, op *xxx_GetMessageClassOperation) *xxx_GetMessageClassOperation {
	if op == nil {
		op = &xxx_GetMessageClassOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.MessageClass = o.MessageClass
	op.Return = o.Return
	return op
}

func (o *GetMessageClassResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMessageClassOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MessageClass = op.MessageClass
	o.Return = op.Return
}
func (o *GetMessageClassResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetMessageClassResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMessageClassOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetMessageClassOperation structure represents the MsgClass operation
type xxx_SetMessageClassOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	MessageClass int32          `idl:"name:lMsgClass" json:"message_class"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetMessageClassOperation) OpNum() int { return 72 }

func (o *xxx_SetMessageClassOperation) OpName() string { return "/IMSMQMessage3/v0/MsgClass" }

func (o *xxx_SetMessageClassOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMessageClassOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lMsgClass {in} (1:(int32))
	{
		if err := w.WriteData(o.MessageClass); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMessageClassOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lMsgClass {in} (1:(int32))
	{
		if err := w.ReadData(&o.MessageClass); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMessageClassOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMessageClassOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetMessageClassOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetMessageClassRequest structure represents the MsgClass operation request
type SetMessageClassRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	MessageClass int32          `idl:"name:lMsgClass" json:"message_class"`
}

func (o *SetMessageClassRequest) xxx_ToOp(ctx context.Context, op *xxx_SetMessageClassOperation) *xxx_SetMessageClassOperation {
	if op == nil {
		op = &xxx_SetMessageClassOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MessageClass = o.MessageClass
	return op
}

func (o *SetMessageClassRequest) xxx_FromOp(ctx context.Context, op *xxx_SetMessageClassOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MessageClass = op.MessageClass
}
func (o *SetMessageClassRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetMessageClassRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMessageClassOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetMessageClassResponse structure represents the MsgClass operation response
type SetMessageClassResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The MsgClass return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetMessageClassResponse) xxx_ToOp(ctx context.Context, op *xxx_SetMessageClassOperation) *xxx_SetMessageClassOperation {
	if op == nil {
		op = &xxx_SetMessageClassOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetMessageClassResponse) xxx_FromOp(ctx context.Context, op *xxx_SetMessageClassOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetMessageClassResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetMessageClassResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMessageClassOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPropertiesOperation structure represents the Properties operation
type xxx_GetPropertiesOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	Properties *oaut.Dispatch `idl:"name:ppcolProperties" json:"properties"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPropertiesOperation) OpNum() int { return 73 }

func (o *xxx_GetPropertiesOperation) OpName() string { return "/IMSMQMessage3/v0/Properties" }

func (o *xxx_GetPropertiesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPropertiesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetPropertiesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppcolProperties {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IDispatch}(interface))
	{
		if o.Properties != nil {
			_ptr_ppcolProperties := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Properties != nil {
					if err := o.Properties.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Dispatch{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Properties, _ptr_ppcolProperties); err != nil {
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

func (o *xxx_GetPropertiesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppcolProperties {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IDispatch}(interface))
	{
		_ptr_ppcolProperties := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Properties == nil {
				o.Properties = &oaut.Dispatch{}
			}
			if err := o.Properties.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppcolProperties := func(ptr interface{}) { o.Properties = *ptr.(**oaut.Dispatch) }
		if err := w.ReadPointer(&o.Properties, _s_ppcolProperties, _ptr_ppcolProperties); err != nil {
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

// GetPropertiesRequest structure represents the Properties operation request
type GetPropertiesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPropertiesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPropertiesOperation) *xxx_GetPropertiesOperation {
	if op == nil {
		op = &xxx_GetPropertiesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetPropertiesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPropertiesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPropertiesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPropertiesResponse structure represents the Properties operation response
type GetPropertiesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	Properties *oaut.Dispatch `idl:"name:ppcolProperties" json:"properties"`
	// Return: The Properties return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPropertiesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPropertiesOperation) *xxx_GetPropertiesOperation {
	if op == nil {
		op = &xxx_GetPropertiesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Properties = o.Properties
	op.Return = o.Return
	return op
}

func (o *GetPropertiesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Properties = op.Properties
	o.Return = op.Return
}
func (o *GetPropertiesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPropertiesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetTransactionIDOperation structure represents the TransactionId operation
type xxx_GetTransactionIDOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	TransactionID *oaut.Variant  `idl:"name:pvarXactId" json:"transaction_id"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTransactionIDOperation) OpNum() int { return 74 }

func (o *xxx_GetTransactionIDOperation) OpName() string { return "/IMSMQMessage3/v0/TransactionId" }

func (o *xxx_GetTransactionIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTransactionIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetTransactionIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetTransactionIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTransactionIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pvarXactId {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.TransactionID != nil {
			_ptr_pvarXactId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.TransactionID != nil {
					if err := o.TransactionID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.TransactionID, _ptr_pvarXactId); err != nil {
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

func (o *xxx_GetTransactionIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pvarXactId {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pvarXactId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.TransactionID == nil {
				o.TransactionID = &oaut.Variant{}
			}
			if err := o.TransactionID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pvarXactId := func(ptr interface{}) { o.TransactionID = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.TransactionID, _s_pvarXactId, _ptr_pvarXactId); err != nil {
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

// GetTransactionIDRequest structure represents the TransactionId operation request
type GetTransactionIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetTransactionIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetTransactionIDOperation) *xxx_GetTransactionIDOperation {
	if op == nil {
		op = &xxx_GetTransactionIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetTransactionIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTransactionIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetTransactionIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetTransactionIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTransactionIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetTransactionIDResponse structure represents the TransactionId operation response
type GetTransactionIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pvarXactId: A pointer to a VARIANT that contains a 20-byte array (VT_ARRAY | VT_UI1) containing the transaction identifier. The format for this ID MUST follow the structure as defined in [MS-MQMQ] section 2.2.8, where the Lineage part is the GUID of the sending computer and the Uniquifier part is a transaction sequence number.
	TransactionID *oaut.Variant `idl:"name:pvarXactId" json:"transaction_id"`
	// Return: The TransactionId return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTransactionIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetTransactionIDOperation) *xxx_GetTransactionIDOperation {
	if op == nil {
		op = &xxx_GetTransactionIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.TransactionID = o.TransactionID
	op.Return = o.Return
	return op
}

func (o *GetTransactionIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTransactionIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TransactionID = op.TransactionID
	o.Return = op.Return
}
func (o *GetTransactionIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetTransactionIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTransactionIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetIsFirstInTransactionOperation structure represents the IsFirstInTransaction operation
type xxx_GetIsFirstInTransactionOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	IsFirstInTransaction int16          `idl:"name:pisFirstInXact" json:"is_first_in_transaction"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetIsFirstInTransactionOperation) OpNum() int { return 75 }

func (o *xxx_GetIsFirstInTransactionOperation) OpName() string {
	return "/IMSMQMessage3/v0/IsFirstInTransaction"
}

func (o *xxx_GetIsFirstInTransactionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsFirstInTransactionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetIsFirstInTransactionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetIsFirstInTransactionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsFirstInTransactionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pisFirstInXact {out, retval} (1:{pointer=ref}*(1)(int16))
	{
		if err := w.WriteData(o.IsFirstInTransaction); err != nil {
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

func (o *xxx_GetIsFirstInTransactionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pisFirstInXact {out, retval} (1:{pointer=ref}*(1)(int16))
	{
		if err := w.ReadData(&o.IsFirstInTransaction); err != nil {
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

// GetIsFirstInTransactionRequest structure represents the IsFirstInTransaction operation request
type GetIsFirstInTransactionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetIsFirstInTransactionRequest) xxx_ToOp(ctx context.Context, op *xxx_GetIsFirstInTransactionOperation) *xxx_GetIsFirstInTransactionOperation {
	if op == nil {
		op = &xxx_GetIsFirstInTransactionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetIsFirstInTransactionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetIsFirstInTransactionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetIsFirstInTransactionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetIsFirstInTransactionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsFirstInTransactionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetIsFirstInTransactionResponse structure represents the IsFirstInTransaction operation response
type GetIsFirstInTransactionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pisFirstInXact: A pointer to a short that specifies whether the message was the first
	// sent by the transaction.
	IsFirstInTransaction int16 `idl:"name:pisFirstInXact" json:"is_first_in_transaction"`
	// Return: The IsFirstInTransaction return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetIsFirstInTransactionResponse) xxx_ToOp(ctx context.Context, op *xxx_GetIsFirstInTransactionOperation) *xxx_GetIsFirstInTransactionOperation {
	if op == nil {
		op = &xxx_GetIsFirstInTransactionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.IsFirstInTransaction = o.IsFirstInTransaction
	op.Return = o.Return
	return op
}

func (o *GetIsFirstInTransactionResponse) xxx_FromOp(ctx context.Context, op *xxx_GetIsFirstInTransactionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.IsFirstInTransaction = op.IsFirstInTransaction
	o.Return = op.Return
}
func (o *GetIsFirstInTransactionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetIsFirstInTransactionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsFirstInTransactionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetIsLastInTransactionOperation structure represents the IsLastInTransaction operation
type xxx_GetIsLastInTransactionOperation struct {
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	IsLastInTransaction int16          `idl:"name:pisLastInXact" json:"is_last_in_transaction"`
	Return              int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetIsLastInTransactionOperation) OpNum() int { return 76 }

func (o *xxx_GetIsLastInTransactionOperation) OpName() string {
	return "/IMSMQMessage3/v0/IsLastInTransaction"
}

func (o *xxx_GetIsLastInTransactionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsLastInTransactionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetIsLastInTransactionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetIsLastInTransactionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsLastInTransactionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pisLastInXact {out, retval} (1:{pointer=ref}*(1)(int16))
	{
		if err := w.WriteData(o.IsLastInTransaction); err != nil {
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

func (o *xxx_GetIsLastInTransactionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pisLastInXact {out, retval} (1:{pointer=ref}*(1)(int16))
	{
		if err := w.ReadData(&o.IsLastInTransaction); err != nil {
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

// GetIsLastInTransactionRequest structure represents the IsLastInTransaction operation request
type GetIsLastInTransactionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetIsLastInTransactionRequest) xxx_ToOp(ctx context.Context, op *xxx_GetIsLastInTransactionOperation) *xxx_GetIsLastInTransactionOperation {
	if op == nil {
		op = &xxx_GetIsLastInTransactionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetIsLastInTransactionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetIsLastInTransactionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetIsLastInTransactionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetIsLastInTransactionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsLastInTransactionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetIsLastInTransactionResponse structure represents the IsLastInTransaction operation response
type GetIsLastInTransactionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pisLastInXact: A pointer to a short that specifies whether the message was the last
	// sent by the transaction.
	IsLastInTransaction int16 `idl:"name:pisLastInXact" json:"is_last_in_transaction"`
	// Return: The IsLastInTransaction return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetIsLastInTransactionResponse) xxx_ToOp(ctx context.Context, op *xxx_GetIsLastInTransactionOperation) *xxx_GetIsLastInTransactionOperation {
	if op == nil {
		op = &xxx_GetIsLastInTransactionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.IsLastInTransaction = o.IsLastInTransaction
	op.Return = o.Return
	return op
}

func (o *GetIsLastInTransactionResponse) xxx_FromOp(ctx context.Context, op *xxx_GetIsLastInTransactionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.IsLastInTransaction = op.IsLastInTransaction
	o.Return = op.Return
}
func (o *GetIsLastInTransactionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetIsLastInTransactionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsLastInTransactionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetResponseQueueInfoV2Operation structure represents the ResponseQueueInfo_v2 operation
type xxx_GetResponseQueueInfoV2Operation struct {
	This     *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Response *mqac.QueueInfo2 `idl:"name:ppqinfoResponse" json:"response"`
	Return   int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetResponseQueueInfoV2Operation) OpNum() int { return 77 }

func (o *xxx_GetResponseQueueInfoV2Operation) OpName() string {
	return "/IMSMQMessage3/v0/ResponseQueueInfo_v2"
}

func (o *xxx_GetResponseQueueInfoV2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResponseQueueInfoV2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetResponseQueueInfoV2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetResponseQueueInfoV2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResponseQueueInfoV2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppqinfoResponse {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQQueueInfo2}(interface))
	{
		if o.Response != nil {
			_ptr_ppqinfoResponse := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Response != nil {
					if err := o.Response.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.QueueInfo2{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetResponseQueueInfoV2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppqinfoResponse {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQQueueInfo2}(interface))
	{
		_ptr_ppqinfoResponse := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Response == nil {
				o.Response = &mqac.QueueInfo2{}
			}
			if err := o.Response.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppqinfoResponse := func(ptr interface{}) { o.Response = *ptr.(**mqac.QueueInfo2) }
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

// GetResponseQueueInfoV2Request structure represents the ResponseQueueInfo_v2 operation request
type GetResponseQueueInfoV2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetResponseQueueInfoV2Request) xxx_ToOp(ctx context.Context, op *xxx_GetResponseQueueInfoV2Operation) *xxx_GetResponseQueueInfoV2Operation {
	if op == nil {
		op = &xxx_GetResponseQueueInfoV2Operation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetResponseQueueInfoV2Request) xxx_FromOp(ctx context.Context, op *xxx_GetResponseQueueInfoV2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetResponseQueueInfoV2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetResponseQueueInfoV2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetResponseQueueInfoV2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetResponseQueueInfoV2Response structure represents the ResponseQueueInfo_v2 operation response
type GetResponseQueueInfoV2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Response *mqac.QueueInfo2 `idl:"name:ppqinfoResponse" json:"response"`
	// Return: The ResponseQueueInfo_v2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetResponseQueueInfoV2Response) xxx_ToOp(ctx context.Context, op *xxx_GetResponseQueueInfoV2Operation) *xxx_GetResponseQueueInfoV2Operation {
	if op == nil {
		op = &xxx_GetResponseQueueInfoV2Operation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Response = o.Response
	op.Return = o.Return
	return op
}

func (o *GetResponseQueueInfoV2Response) xxx_FromOp(ctx context.Context, op *xxx_GetResponseQueueInfoV2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Response = op.Response
	o.Return = op.Return
}
func (o *GetResponseQueueInfoV2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetResponseQueueInfoV2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetResponseQueueInfoV2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetResponseQueueInfoV2Operation structure represents the ResponseQueueInfo_v2 operation
type xxx_SetResponseQueueInfoV2Operation struct {
	This              *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat   `idl:"name:That" json:"that"`
	QueueInfoResponse *mqac.QueueInfo2 `idl:"name:pqinfoResponse" json:"queue_info_response"`
	Return            int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_SetResponseQueueInfoV2Operation) OpNum() int { return 78 }

func (o *xxx_SetResponseQueueInfoV2Operation) OpName() string {
	return "/IMSMQMessage3/v0/ResponseQueueInfo_v2"
}

func (o *xxx_SetResponseQueueInfoV2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetResponseQueueInfoV2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pqinfoResponse {in} (1:{pointer=ref}*(1))(2:{alias=IMSMQQueueInfo2}(interface))
	{
		if o.QueueInfoResponse != nil {
			_ptr_pqinfoResponse := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.QueueInfoResponse != nil {
					if err := o.QueueInfoResponse.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.QueueInfo2{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_SetResponseQueueInfoV2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pqinfoResponse {in} (1:{pointer=ref}*(1))(2:{alias=IMSMQQueueInfo2}(interface))
	{
		_ptr_pqinfoResponse := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.QueueInfoResponse == nil {
				o.QueueInfoResponse = &mqac.QueueInfo2{}
			}
			if err := o.QueueInfoResponse.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pqinfoResponse := func(ptr interface{}) { o.QueueInfoResponse = *ptr.(**mqac.QueueInfo2) }
		if err := w.ReadPointer(&o.QueueInfoResponse, _s_pqinfoResponse, _ptr_pqinfoResponse); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetResponseQueueInfoV2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetResponseQueueInfoV2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetResponseQueueInfoV2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetResponseQueueInfoV2Request structure represents the ResponseQueueInfo_v2 operation request
type SetResponseQueueInfoV2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This              *dcom.ORPCThis   `idl:"name:This" json:"this"`
	QueueInfoResponse *mqac.QueueInfo2 `idl:"name:pqinfoResponse" json:"queue_info_response"`
}

func (o *SetResponseQueueInfoV2Request) xxx_ToOp(ctx context.Context, op *xxx_SetResponseQueueInfoV2Operation) *xxx_SetResponseQueueInfoV2Operation {
	if op == nil {
		op = &xxx_SetResponseQueueInfoV2Operation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.QueueInfoResponse = o.QueueInfoResponse
	return op
}

func (o *SetResponseQueueInfoV2Request) xxx_FromOp(ctx context.Context, op *xxx_SetResponseQueueInfoV2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.QueueInfoResponse = op.QueueInfoResponse
}
func (o *SetResponseQueueInfoV2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetResponseQueueInfoV2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetResponseQueueInfoV2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetResponseQueueInfoV2Response structure represents the ResponseQueueInfo_v2 operation response
type SetResponseQueueInfoV2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ResponseQueueInfo_v2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetResponseQueueInfoV2Response) xxx_ToOp(ctx context.Context, op *xxx_SetResponseQueueInfoV2Operation) *xxx_SetResponseQueueInfoV2Operation {
	if op == nil {
		op = &xxx_SetResponseQueueInfoV2Operation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetResponseQueueInfoV2Response) xxx_FromOp(ctx context.Context, op *xxx_SetResponseQueueInfoV2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetResponseQueueInfoV2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetResponseQueueInfoV2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetResponseQueueInfoV2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAdminQueueInfoV2Operation structure represents the AdminQueueInfo_v2 operation
type xxx_GetAdminQueueInfoV2Operation struct {
	This   *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Admin  *mqac.QueueInfo2 `idl:"name:ppqinfoAdmin" json:"admin"`
	Return int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAdminQueueInfoV2Operation) OpNum() int { return 79 }

func (o *xxx_GetAdminQueueInfoV2Operation) OpName() string {
	return "/IMSMQMessage3/v0/AdminQueueInfo_v2"
}

func (o *xxx_GetAdminQueueInfoV2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAdminQueueInfoV2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetAdminQueueInfoV2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetAdminQueueInfoV2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAdminQueueInfoV2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppqinfoAdmin {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQQueueInfo2}(interface))
	{
		if o.Admin != nil {
			_ptr_ppqinfoAdmin := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Admin != nil {
					if err := o.Admin.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.QueueInfo2{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetAdminQueueInfoV2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppqinfoAdmin {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQQueueInfo2}(interface))
	{
		_ptr_ppqinfoAdmin := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Admin == nil {
				o.Admin = &mqac.QueueInfo2{}
			}
			if err := o.Admin.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppqinfoAdmin := func(ptr interface{}) { o.Admin = *ptr.(**mqac.QueueInfo2) }
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

// GetAdminQueueInfoV2Request structure represents the AdminQueueInfo_v2 operation request
type GetAdminQueueInfoV2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetAdminQueueInfoV2Request) xxx_ToOp(ctx context.Context, op *xxx_GetAdminQueueInfoV2Operation) *xxx_GetAdminQueueInfoV2Operation {
	if op == nil {
		op = &xxx_GetAdminQueueInfoV2Operation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetAdminQueueInfoV2Request) xxx_FromOp(ctx context.Context, op *xxx_GetAdminQueueInfoV2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetAdminQueueInfoV2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetAdminQueueInfoV2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAdminQueueInfoV2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAdminQueueInfoV2Response structure represents the AdminQueueInfo_v2 operation response
type GetAdminQueueInfoV2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Admin *mqac.QueueInfo2 `idl:"name:ppqinfoAdmin" json:"admin"`
	// Return: The AdminQueueInfo_v2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAdminQueueInfoV2Response) xxx_ToOp(ctx context.Context, op *xxx_GetAdminQueueInfoV2Operation) *xxx_GetAdminQueueInfoV2Operation {
	if op == nil {
		op = &xxx_GetAdminQueueInfoV2Operation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Admin = o.Admin
	op.Return = o.Return
	return op
}

func (o *GetAdminQueueInfoV2Response) xxx_FromOp(ctx context.Context, op *xxx_GetAdminQueueInfoV2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Admin = op.Admin
	o.Return = op.Return
}
func (o *GetAdminQueueInfoV2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetAdminQueueInfoV2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAdminQueueInfoV2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetAdminQueueInfoV2Operation structure represents the AdminQueueInfo_v2 operation
type xxx_SetAdminQueueInfoV2Operation struct {
	This           *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat   `idl:"name:That" json:"that"`
	QueueInfoAdmin *mqac.QueueInfo2 `idl:"name:pqinfoAdmin" json:"queue_info_admin"`
	Return         int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_SetAdminQueueInfoV2Operation) OpNum() int { return 80 }

func (o *xxx_SetAdminQueueInfoV2Operation) OpName() string {
	return "/IMSMQMessage3/v0/AdminQueueInfo_v2"
}

func (o *xxx_SetAdminQueueInfoV2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAdminQueueInfoV2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pqinfoAdmin {in} (1:{pointer=ref}*(1))(2:{alias=IMSMQQueueInfo2}(interface))
	{
		if o.QueueInfoAdmin != nil {
			_ptr_pqinfoAdmin := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.QueueInfoAdmin != nil {
					if err := o.QueueInfoAdmin.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.QueueInfo2{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_SetAdminQueueInfoV2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pqinfoAdmin {in} (1:{pointer=ref}*(1))(2:{alias=IMSMQQueueInfo2}(interface))
	{
		_ptr_pqinfoAdmin := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.QueueInfoAdmin == nil {
				o.QueueInfoAdmin = &mqac.QueueInfo2{}
			}
			if err := o.QueueInfoAdmin.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pqinfoAdmin := func(ptr interface{}) { o.QueueInfoAdmin = *ptr.(**mqac.QueueInfo2) }
		if err := w.ReadPointer(&o.QueueInfoAdmin, _s_pqinfoAdmin, _ptr_pqinfoAdmin); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAdminQueueInfoV2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAdminQueueInfoV2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetAdminQueueInfoV2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetAdminQueueInfoV2Request structure represents the AdminQueueInfo_v2 operation request
type SetAdminQueueInfoV2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This           *dcom.ORPCThis   `idl:"name:This" json:"this"`
	QueueInfoAdmin *mqac.QueueInfo2 `idl:"name:pqinfoAdmin" json:"queue_info_admin"`
}

func (o *SetAdminQueueInfoV2Request) xxx_ToOp(ctx context.Context, op *xxx_SetAdminQueueInfoV2Operation) *xxx_SetAdminQueueInfoV2Operation {
	if op == nil {
		op = &xxx_SetAdminQueueInfoV2Operation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.QueueInfoAdmin = o.QueueInfoAdmin
	return op
}

func (o *SetAdminQueueInfoV2Request) xxx_FromOp(ctx context.Context, op *xxx_SetAdminQueueInfoV2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.QueueInfoAdmin = op.QueueInfoAdmin
}
func (o *SetAdminQueueInfoV2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetAdminQueueInfoV2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAdminQueueInfoV2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetAdminQueueInfoV2Response structure represents the AdminQueueInfo_v2 operation response
type SetAdminQueueInfoV2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AdminQueueInfo_v2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetAdminQueueInfoV2Response) xxx_ToOp(ctx context.Context, op *xxx_SetAdminQueueInfoV2Operation) *xxx_SetAdminQueueInfoV2Operation {
	if op == nil {
		op = &xxx_SetAdminQueueInfoV2Operation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetAdminQueueInfoV2Response) xxx_FromOp(ctx context.Context, op *xxx_SetAdminQueueInfoV2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetAdminQueueInfoV2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetAdminQueueInfoV2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAdminQueueInfoV2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetReceivedAuthenticationLevelOperation structure represents the ReceivedAuthenticationLevel operation
type xxx_GetReceivedAuthenticationLevelOperation struct {
	This                        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReceivedAuthenticationLevel int16          `idl:"name:psReceivedAuthenticationLevel" json:"received_authentication_level"`
	Return                      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetReceivedAuthenticationLevelOperation) OpNum() int { return 81 }

func (o *xxx_GetReceivedAuthenticationLevelOperation) OpName() string {
	return "/IMSMQMessage3/v0/ReceivedAuthenticationLevel"
}

func (o *xxx_GetReceivedAuthenticationLevelOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReceivedAuthenticationLevelOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetReceivedAuthenticationLevelOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetReceivedAuthenticationLevelOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReceivedAuthenticationLevelOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// psReceivedAuthenticationLevel {out, retval} (1:{pointer=ref}*(1)(int16))
	{
		if err := w.WriteData(o.ReceivedAuthenticationLevel); err != nil {
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

func (o *xxx_GetReceivedAuthenticationLevelOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// psReceivedAuthenticationLevel {out, retval} (1:{pointer=ref}*(1)(int16))
	{
		if err := w.ReadData(&o.ReceivedAuthenticationLevel); err != nil {
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

// GetReceivedAuthenticationLevelRequest structure represents the ReceivedAuthenticationLevel operation request
type GetReceivedAuthenticationLevelRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetReceivedAuthenticationLevelRequest) xxx_ToOp(ctx context.Context, op *xxx_GetReceivedAuthenticationLevelOperation) *xxx_GetReceivedAuthenticationLevelOperation {
	if op == nil {
		op = &xxx_GetReceivedAuthenticationLevelOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetReceivedAuthenticationLevelRequest) xxx_FromOp(ctx context.Context, op *xxx_GetReceivedAuthenticationLevelOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetReceivedAuthenticationLevelRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetReceivedAuthenticationLevelRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetReceivedAuthenticationLevelOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetReceivedAuthenticationLevelResponse structure represents the ReceivedAuthenticationLevel operation response
type GetReceivedAuthenticationLevelResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// psReceivedAuthenticationLevel: A pointer to a short that identifies if the message
	// was authenticated and what digital signature was used. This parameter corresponds
	// to the MQMSGAUTHENTICATION enum as defined in MQMSGAUTHENTICATION (section 2.2.2.17).
	ReceivedAuthenticationLevel int16 `idl:"name:psReceivedAuthenticationLevel" json:"received_authentication_level"`
	// Return: The ReceivedAuthenticationLevel return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetReceivedAuthenticationLevelResponse) xxx_ToOp(ctx context.Context, op *xxx_GetReceivedAuthenticationLevelOperation) *xxx_GetReceivedAuthenticationLevelOperation {
	if op == nil {
		op = &xxx_GetReceivedAuthenticationLevelOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReceivedAuthenticationLevel = o.ReceivedAuthenticationLevel
	op.Return = o.Return
	return op
}

func (o *GetReceivedAuthenticationLevelResponse) xxx_FromOp(ctx context.Context, op *xxx_GetReceivedAuthenticationLevelOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReceivedAuthenticationLevel = op.ReceivedAuthenticationLevel
	o.Return = op.Return
}
func (o *GetReceivedAuthenticationLevelResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetReceivedAuthenticationLevelResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetReceivedAuthenticationLevelOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetResponseQueueInfoOperation structure represents the ResponseQueueInfo operation
type xxx_GetResponseQueueInfoOperation struct {
	This     *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Response *mqac.QueueInfo3 `idl:"name:ppqinfoResponse" json:"response"`
	Return   int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetResponseQueueInfoOperation) OpNum() int { return 82 }

func (o *xxx_GetResponseQueueInfoOperation) OpName() string {
	return "/IMSMQMessage3/v0/ResponseQueueInfo"
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
	// ppqinfoResponse {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQQueueInfo3}(interface))
	{
		if o.Response != nil {
			_ptr_ppqinfoResponse := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Response != nil {
					if err := o.Response.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.QueueInfo3{}).MarshalNDR(ctx, w); err != nil {
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
	// ppqinfoResponse {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQQueueInfo3}(interface))
	{
		_ptr_ppqinfoResponse := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Response == nil {
				o.Response = &mqac.QueueInfo3{}
			}
			if err := o.Response.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppqinfoResponse := func(ptr interface{}) { o.Response = *ptr.(**mqac.QueueInfo3) }
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
	That     *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Response *mqac.QueueInfo3 `idl:"name:ppqinfoResponse" json:"response"`
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
	This              *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat   `idl:"name:That" json:"that"`
	QueueInfoResponse *mqac.QueueInfo3 `idl:"name:pqinfoResponse" json:"queue_info_response"`
	Return            int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_SetByRefResponseQueueInfoOperation) OpNum() int { return 83 }

func (o *xxx_SetByRefResponseQueueInfoOperation) OpName() string {
	return "/IMSMQMessage3/v0/ResponseQueueInfo"
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
	// pqinfoResponse {in} (1:{pointer=ref}*(1))(2:{alias=IMSMQQueueInfo3}(interface))
	{
		if o.QueueInfoResponse != nil {
			_ptr_pqinfoResponse := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.QueueInfoResponse != nil {
					if err := o.QueueInfoResponse.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.QueueInfo3{}).MarshalNDR(ctx, w); err != nil {
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
	// pqinfoResponse {in} (1:{pointer=ref}*(1))(2:{alias=IMSMQQueueInfo3}(interface))
	{
		_ptr_pqinfoResponse := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.QueueInfoResponse == nil {
				o.QueueInfoResponse = &mqac.QueueInfo3{}
			}
			if err := o.QueueInfoResponse.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pqinfoResponse := func(ptr interface{}) { o.QueueInfoResponse = *ptr.(**mqac.QueueInfo3) }
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
	This              *dcom.ORPCThis   `idl:"name:This" json:"this"`
	QueueInfoResponse *mqac.QueueInfo3 `idl:"name:pqinfoResponse" json:"queue_info_response"`
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

// xxx_GetAdminQueueInfoOperation structure represents the AdminQueueInfo operation
type xxx_GetAdminQueueInfoOperation struct {
	This   *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Admin  *mqac.QueueInfo3 `idl:"name:ppqinfoAdmin" json:"admin"`
	Return int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAdminQueueInfoOperation) OpNum() int { return 84 }

func (o *xxx_GetAdminQueueInfoOperation) OpName() string { return "/IMSMQMessage3/v0/AdminQueueInfo" }

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
	// ppqinfoAdmin {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQQueueInfo3}(interface))
	{
		if o.Admin != nil {
			_ptr_ppqinfoAdmin := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Admin != nil {
					if err := o.Admin.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.QueueInfo3{}).MarshalNDR(ctx, w); err != nil {
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
	// ppqinfoAdmin {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQQueueInfo3}(interface))
	{
		_ptr_ppqinfoAdmin := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Admin == nil {
				o.Admin = &mqac.QueueInfo3{}
			}
			if err := o.Admin.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppqinfoAdmin := func(ptr interface{}) { o.Admin = *ptr.(**mqac.QueueInfo3) }
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
	That  *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Admin *mqac.QueueInfo3 `idl:"name:ppqinfoAdmin" json:"admin"`
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
	This           *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat   `idl:"name:That" json:"that"`
	QueueInfoAdmin *mqac.QueueInfo3 `idl:"name:pqinfoAdmin" json:"queue_info_admin"`
	Return         int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_SetByRefAdminQueueInfoOperation) OpNum() int { return 85 }

func (o *xxx_SetByRefAdminQueueInfoOperation) OpName() string {
	return "/IMSMQMessage3/v0/AdminQueueInfo"
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
	// pqinfoAdmin {in} (1:{pointer=ref}*(1))(2:{alias=IMSMQQueueInfo3}(interface))
	{
		if o.QueueInfoAdmin != nil {
			_ptr_pqinfoAdmin := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.QueueInfoAdmin != nil {
					if err := o.QueueInfoAdmin.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.QueueInfo3{}).MarshalNDR(ctx, w); err != nil {
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
	// pqinfoAdmin {in} (1:{pointer=ref}*(1))(2:{alias=IMSMQQueueInfo3}(interface))
	{
		_ptr_pqinfoAdmin := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.QueueInfoAdmin == nil {
				o.QueueInfoAdmin = &mqac.QueueInfo3{}
			}
			if err := o.QueueInfoAdmin.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pqinfoAdmin := func(ptr interface{}) { o.QueueInfoAdmin = *ptr.(**mqac.QueueInfo3) }
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
	This           *dcom.ORPCThis   `idl:"name:This" json:"this"`
	QueueInfoAdmin *mqac.QueueInfo3 `idl:"name:pqinfoAdmin" json:"queue_info_admin"`
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

// xxx_GetResponseDestinationOperation structure represents the ResponseDestination operation
type xxx_GetResponseDestinationOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Respons *oaut.Dispatch `idl:"name:ppdestRespons" json:"respons"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetResponseDestinationOperation) OpNum() int { return 86 }

func (o *xxx_GetResponseDestinationOperation) OpName() string {
	return "/IMSMQMessage3/v0/ResponseDestination"
}

func (o *xxx_GetResponseDestinationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResponseDestinationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetResponseDestinationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetResponseDestinationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResponseDestinationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppdestRespons {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IDispatch}(interface))
	{
		if o.Respons != nil {
			_ptr_ppdestRespons := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Respons != nil {
					if err := o.Respons.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Dispatch{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Respons, _ptr_ppdestRespons); err != nil {
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

func (o *xxx_GetResponseDestinationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppdestRespons {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IDispatch}(interface))
	{
		_ptr_ppdestRespons := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Respons == nil {
				o.Respons = &oaut.Dispatch{}
			}
			if err := o.Respons.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppdestRespons := func(ptr interface{}) { o.Respons = *ptr.(**oaut.Dispatch) }
		if err := w.ReadPointer(&o.Respons, _s_ppdestRespons, _ptr_ppdestRespons); err != nil {
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

// GetResponseDestinationRequest structure represents the ResponseDestination operation request
type GetResponseDestinationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetResponseDestinationRequest) xxx_ToOp(ctx context.Context, op *xxx_GetResponseDestinationOperation) *xxx_GetResponseDestinationOperation {
	if op == nil {
		op = &xxx_GetResponseDestinationOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetResponseDestinationRequest) xxx_FromOp(ctx context.Context, op *xxx_GetResponseDestinationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetResponseDestinationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetResponseDestinationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetResponseDestinationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetResponseDestinationResponse structure represents the ResponseDestination operation response
type GetResponseDestinationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Respons *oaut.Dispatch `idl:"name:ppdestRespons" json:"respons"`
	// Return: The ResponseDestination return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetResponseDestinationResponse) xxx_ToOp(ctx context.Context, op *xxx_GetResponseDestinationOperation) *xxx_GetResponseDestinationOperation {
	if op == nil {
		op = &xxx_GetResponseDestinationOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Respons = o.Respons
	op.Return = o.Return
	return op
}

func (o *GetResponseDestinationResponse) xxx_FromOp(ctx context.Context, op *xxx_GetResponseDestinationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Respons = op.Respons
	o.Return = op.Return
}
func (o *GetResponseDestinationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetResponseDestinationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetResponseDestinationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetByRefResponseDestinationOperation structure represents the ResponseDestination operation
type xxx_SetByRefResponseDestinationOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Response *oaut.Dispatch `idl:"name:pdestResponse" json:"response"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetByRefResponseDestinationOperation) OpNum() int { return 87 }

func (o *xxx_SetByRefResponseDestinationOperation) OpName() string {
	return "/IMSMQMessage3/v0/ResponseDestination"
}

func (o *xxx_SetByRefResponseDestinationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetByRefResponseDestinationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pdestResponse {in} (1:{pointer=ref}*(1))(2:{alias=IDispatch}(interface))
	{
		if o.Response != nil {
			_ptr_pdestResponse := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Response != nil {
					if err := o.Response.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Dispatch{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Response, _ptr_pdestResponse); err != nil {
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

func (o *xxx_SetByRefResponseDestinationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pdestResponse {in} (1:{pointer=ref}*(1))(2:{alias=IDispatch}(interface))
	{
		_ptr_pdestResponse := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Response == nil {
				o.Response = &oaut.Dispatch{}
			}
			if err := o.Response.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pdestResponse := func(ptr interface{}) { o.Response = *ptr.(**oaut.Dispatch) }
		if err := w.ReadPointer(&o.Response, _s_pdestResponse, _ptr_pdestResponse); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetByRefResponseDestinationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetByRefResponseDestinationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetByRefResponseDestinationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetByRefResponseDestinationRequest structure represents the ResponseDestination operation request
type SetByRefResponseDestinationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	Response *oaut.Dispatch `idl:"name:pdestResponse" json:"response"`
}

func (o *SetByRefResponseDestinationRequest) xxx_ToOp(ctx context.Context, op *xxx_SetByRefResponseDestinationOperation) *xxx_SetByRefResponseDestinationOperation {
	if op == nil {
		op = &xxx_SetByRefResponseDestinationOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Response = o.Response
	return op
}

func (o *SetByRefResponseDestinationRequest) xxx_FromOp(ctx context.Context, op *xxx_SetByRefResponseDestinationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Response = op.Response
}
func (o *SetByRefResponseDestinationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetByRefResponseDestinationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetByRefResponseDestinationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetByRefResponseDestinationResponse structure represents the ResponseDestination operation response
type SetByRefResponseDestinationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ResponseDestination return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetByRefResponseDestinationResponse) xxx_ToOp(ctx context.Context, op *xxx_SetByRefResponseDestinationOperation) *xxx_SetByRefResponseDestinationOperation {
	if op == nil {
		op = &xxx_SetByRefResponseDestinationOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetByRefResponseDestinationResponse) xxx_FromOp(ctx context.Context, op *xxx_SetByRefResponseDestinationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetByRefResponseDestinationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetByRefResponseDestinationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetByRefResponseDestinationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDestinationOperation structure represents the Destination operation
type xxx_GetDestinationOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	Destination *oaut.Dispatch `idl:"name:ppdestDestination" json:"destination"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDestinationOperation) OpNum() int { return 88 }

func (o *xxx_GetDestinationOperation) OpName() string { return "/IMSMQMessage3/v0/Destination" }

func (o *xxx_GetDestinationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDestinationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetDestinationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetDestinationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDestinationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppdestDestination {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IDispatch}(interface))
	{
		if o.Destination != nil {
			_ptr_ppdestDestination := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Destination != nil {
					if err := o.Destination.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Dispatch{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Destination, _ptr_ppdestDestination); err != nil {
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

func (o *xxx_GetDestinationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppdestDestination {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IDispatch}(interface))
	{
		_ptr_ppdestDestination := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Destination == nil {
				o.Destination = &oaut.Dispatch{}
			}
			if err := o.Destination.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppdestDestination := func(ptr interface{}) { o.Destination = *ptr.(**oaut.Dispatch) }
		if err := w.ReadPointer(&o.Destination, _s_ppdestDestination, _ptr_ppdestDestination); err != nil {
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

// GetDestinationRequest structure represents the Destination operation request
type GetDestinationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetDestinationRequest) xxx_ToOp(ctx context.Context, op *xxx_GetDestinationOperation) *xxx_GetDestinationOperation {
	if op == nil {
		op = &xxx_GetDestinationOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetDestinationRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDestinationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetDestinationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetDestinationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDestinationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDestinationResponse structure represents the Destination operation response
type GetDestinationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppdestDestination: A pointer to a pointer to an IDispatch interface that upon successful
	// completion will contain an MSMQDestination object representing the set of queues
	// to send the message to.
	Destination *oaut.Dispatch `idl:"name:ppdestDestination" json:"destination"`
	// Return: The Destination return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDestinationResponse) xxx_ToOp(ctx context.Context, op *xxx_GetDestinationOperation) *xxx_GetDestinationOperation {
	if op == nil {
		op = &xxx_GetDestinationOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Destination = o.Destination
	op.Return = o.Return
	return op
}

func (o *GetDestinationResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDestinationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Destination = op.Destination
	o.Return = op.Return
}
func (o *GetDestinationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetDestinationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDestinationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetLookupIDOperation structure represents the LookupId operation
type xxx_GetLookupIDOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	LookupID *oaut.Variant  `idl:"name:pvarLookupId" json:"lookup_id"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetLookupIDOperation) OpNum() int { return 89 }

func (o *xxx_GetLookupIDOperation) OpName() string { return "/IMSMQMessage3/v0/LookupId" }

func (o *xxx_GetLookupIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLookupIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetLookupIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetLookupIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLookupIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pvarLookupId {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.LookupID != nil {
			_ptr_pvarLookupId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.LookupID != nil {
					if err := o.LookupID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.LookupID, _ptr_pvarLookupId); err != nil {
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

func (o *xxx_GetLookupIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pvarLookupId {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pvarLookupId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.LookupID == nil {
				o.LookupID = &oaut.Variant{}
			}
			if err := o.LookupID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pvarLookupId := func(ptr interface{}) { o.LookupID = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.LookupID, _s_pvarLookupId, _ptr_pvarLookupId); err != nil {
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

// GetLookupIDRequest structure represents the LookupId operation request
type GetLookupIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetLookupIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetLookupIDOperation) *xxx_GetLookupIDOperation {
	if op == nil {
		op = &xxx_GetLookupIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetLookupIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetLookupIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetLookupIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetLookupIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLookupIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetLookupIDResponse structure represents the LookupId operation response
type GetLookupIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pvarLookupId: A pointer to a VARIANT that upon successful completion will contain
	// a BSTR that contains the 64-bit lookup identifier for the message.
	LookupID *oaut.Variant `idl:"name:pvarLookupId" json:"lookup_id"`
	// Return: The LookupId return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetLookupIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetLookupIDOperation) *xxx_GetLookupIDOperation {
	if op == nil {
		op = &xxx_GetLookupIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.LookupID = o.LookupID
	op.Return = o.Return
	return op
}

func (o *GetLookupIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetLookupIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.LookupID = op.LookupID
	o.Return = op.Return
}
func (o *GetLookupIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetLookupIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLookupIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetIsAuthenticated2Operation structure represents the IsAuthenticated2 operation
type xxx_GetIsAuthenticated2Operation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	IsAuthenticated int16          `idl:"name:pisAuthenticated" json:"is_authenticated"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetIsAuthenticated2Operation) OpNum() int { return 90 }

func (o *xxx_GetIsAuthenticated2Operation) OpName() string {
	return "/IMSMQMessage3/v0/IsAuthenticated2"
}

func (o *xxx_GetIsAuthenticated2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsAuthenticated2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetIsAuthenticated2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetIsAuthenticated2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsAuthenticated2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pisAuthenticated {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
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

func (o *xxx_GetIsAuthenticated2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pisAuthenticated {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
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

// GetIsAuthenticated2Request structure represents the IsAuthenticated2 operation request
type GetIsAuthenticated2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetIsAuthenticated2Request) xxx_ToOp(ctx context.Context, op *xxx_GetIsAuthenticated2Operation) *xxx_GetIsAuthenticated2Operation {
	if op == nil {
		op = &xxx_GetIsAuthenticated2Operation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetIsAuthenticated2Request) xxx_FromOp(ctx context.Context, op *xxx_GetIsAuthenticated2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetIsAuthenticated2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetIsAuthenticated2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsAuthenticated2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetIsAuthenticated2Response structure represents the IsAuthenticated2 operation response
type GetIsAuthenticated2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pisAuthenticated: A pointer to a VARIANT_BOOL that specifies whether the message
	// was authenticated.
	IsAuthenticated int16 `idl:"name:pisAuthenticated" json:"is_authenticated"`
	// Return: The IsAuthenticated2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetIsAuthenticated2Response) xxx_ToOp(ctx context.Context, op *xxx_GetIsAuthenticated2Operation) *xxx_GetIsAuthenticated2Operation {
	if op == nil {
		op = &xxx_GetIsAuthenticated2Operation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.IsAuthenticated = o.IsAuthenticated
	op.Return = o.Return
	return op
}

func (o *GetIsAuthenticated2Response) xxx_FromOp(ctx context.Context, op *xxx_GetIsAuthenticated2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.IsAuthenticated = op.IsAuthenticated
	o.Return = op.Return
}
func (o *GetIsAuthenticated2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetIsAuthenticated2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsAuthenticated2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetIsFirstInTransaction2Operation structure represents the IsFirstInTransaction2 operation
type xxx_GetIsFirstInTransaction2Operation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	IsFirstInTransaction int16          `idl:"name:pisFirstInXact" json:"is_first_in_transaction"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetIsFirstInTransaction2Operation) OpNum() int { return 91 }

func (o *xxx_GetIsFirstInTransaction2Operation) OpName() string {
	return "/IMSMQMessage3/v0/IsFirstInTransaction2"
}

func (o *xxx_GetIsFirstInTransaction2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsFirstInTransaction2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetIsFirstInTransaction2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetIsFirstInTransaction2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsFirstInTransaction2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pisFirstInXact {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.IsFirstInTransaction); err != nil {
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

func (o *xxx_GetIsFirstInTransaction2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pisFirstInXact {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.IsFirstInTransaction); err != nil {
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

// GetIsFirstInTransaction2Request structure represents the IsFirstInTransaction2 operation request
type GetIsFirstInTransaction2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetIsFirstInTransaction2Request) xxx_ToOp(ctx context.Context, op *xxx_GetIsFirstInTransaction2Operation) *xxx_GetIsFirstInTransaction2Operation {
	if op == nil {
		op = &xxx_GetIsFirstInTransaction2Operation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetIsFirstInTransaction2Request) xxx_FromOp(ctx context.Context, op *xxx_GetIsFirstInTransaction2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetIsFirstInTransaction2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetIsFirstInTransaction2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsFirstInTransaction2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetIsFirstInTransaction2Response structure represents the IsFirstInTransaction2 operation response
type GetIsFirstInTransaction2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pisFirstInXact: A pointer to a VARIANT_BOOL that specifies whether the message was
	// the first sent by the transaction.
	IsFirstInTransaction int16 `idl:"name:pisFirstInXact" json:"is_first_in_transaction"`
	// Return: The IsFirstInTransaction2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetIsFirstInTransaction2Response) xxx_ToOp(ctx context.Context, op *xxx_GetIsFirstInTransaction2Operation) *xxx_GetIsFirstInTransaction2Operation {
	if op == nil {
		op = &xxx_GetIsFirstInTransaction2Operation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.IsFirstInTransaction = o.IsFirstInTransaction
	op.Return = o.Return
	return op
}

func (o *GetIsFirstInTransaction2Response) xxx_FromOp(ctx context.Context, op *xxx_GetIsFirstInTransaction2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.IsFirstInTransaction = op.IsFirstInTransaction
	o.Return = op.Return
}
func (o *GetIsFirstInTransaction2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetIsFirstInTransaction2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsFirstInTransaction2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetIsLastInTransaction2Operation structure represents the IsLastInTransaction2 operation
type xxx_GetIsLastInTransaction2Operation struct {
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	IsLastInTransaction int16          `idl:"name:pisLastInXact" json:"is_last_in_transaction"`
	Return              int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetIsLastInTransaction2Operation) OpNum() int { return 92 }

func (o *xxx_GetIsLastInTransaction2Operation) OpName() string {
	return "/IMSMQMessage3/v0/IsLastInTransaction2"
}

func (o *xxx_GetIsLastInTransaction2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsLastInTransaction2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetIsLastInTransaction2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetIsLastInTransaction2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsLastInTransaction2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pisLastInXact {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.IsLastInTransaction); err != nil {
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

func (o *xxx_GetIsLastInTransaction2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pisLastInXact {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.IsLastInTransaction); err != nil {
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

// GetIsLastInTransaction2Request structure represents the IsLastInTransaction2 operation request
type GetIsLastInTransaction2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetIsLastInTransaction2Request) xxx_ToOp(ctx context.Context, op *xxx_GetIsLastInTransaction2Operation) *xxx_GetIsLastInTransaction2Operation {
	if op == nil {
		op = &xxx_GetIsLastInTransaction2Operation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetIsLastInTransaction2Request) xxx_FromOp(ctx context.Context, op *xxx_GetIsLastInTransaction2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetIsLastInTransaction2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetIsLastInTransaction2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsLastInTransaction2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetIsLastInTransaction2Response structure represents the IsLastInTransaction2 operation response
type GetIsLastInTransaction2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pisLastInXact: A pointer to a VARIANT_BOOL that specifies whether the message was
	// the last sent by the transaction.
	IsLastInTransaction int16 `idl:"name:pisLastInXact" json:"is_last_in_transaction"`
	// Return: The IsLastInTransaction2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetIsLastInTransaction2Response) xxx_ToOp(ctx context.Context, op *xxx_GetIsLastInTransaction2Operation) *xxx_GetIsLastInTransaction2Operation {
	if op == nil {
		op = &xxx_GetIsLastInTransaction2Operation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.IsLastInTransaction = o.IsLastInTransaction
	op.Return = o.Return
	return op
}

func (o *GetIsLastInTransaction2Response) xxx_FromOp(ctx context.Context, op *xxx_GetIsLastInTransaction2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.IsLastInTransaction = op.IsLastInTransaction
	o.Return = op.Return
}
func (o *GetIsLastInTransaction2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetIsLastInTransaction2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsLastInTransaction2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AttachCurrentSecurityContext2Operation structure represents the AttachCurrentSecurityContext2 operation
type xxx_AttachCurrentSecurityContext2Operation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_AttachCurrentSecurityContext2Operation) OpNum() int { return 93 }

func (o *xxx_AttachCurrentSecurityContext2Operation) OpName() string {
	return "/IMSMQMessage3/v0/AttachCurrentSecurityContext2"
}

func (o *xxx_AttachCurrentSecurityContext2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AttachCurrentSecurityContext2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AttachCurrentSecurityContext2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_AttachCurrentSecurityContext2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AttachCurrentSecurityContext2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AttachCurrentSecurityContext2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// AttachCurrentSecurityContext2Request structure represents the AttachCurrentSecurityContext2 operation request
type AttachCurrentSecurityContext2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *AttachCurrentSecurityContext2Request) xxx_ToOp(ctx context.Context, op *xxx_AttachCurrentSecurityContext2Operation) *xxx_AttachCurrentSecurityContext2Operation {
	if op == nil {
		op = &xxx_AttachCurrentSecurityContext2Operation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *AttachCurrentSecurityContext2Request) xxx_FromOp(ctx context.Context, op *xxx_AttachCurrentSecurityContext2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *AttachCurrentSecurityContext2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AttachCurrentSecurityContext2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AttachCurrentSecurityContext2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AttachCurrentSecurityContext2Response structure represents the AttachCurrentSecurityContext2 operation response
type AttachCurrentSecurityContext2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AttachCurrentSecurityContext2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AttachCurrentSecurityContext2Response) xxx_ToOp(ctx context.Context, op *xxx_AttachCurrentSecurityContext2Operation) *xxx_AttachCurrentSecurityContext2Operation {
	if op == nil {
		op = &xxx_AttachCurrentSecurityContext2Operation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *AttachCurrentSecurityContext2Response) xxx_FromOp(ctx context.Context, op *xxx_AttachCurrentSecurityContext2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *AttachCurrentSecurityContext2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AttachCurrentSecurityContext2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AttachCurrentSecurityContext2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSOAPEnvelopeOperation structure represents the SoapEnvelope operation
type xxx_GetSOAPEnvelopeOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	SOAPEnvelope *oaut.String   `idl:"name:pbstrSoapEnvelope" json:"soap_envelope"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSOAPEnvelopeOperation) OpNum() int { return 94 }

func (o *xxx_GetSOAPEnvelopeOperation) OpName() string { return "/IMSMQMessage3/v0/SoapEnvelope" }

func (o *xxx_GetSOAPEnvelopeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSOAPEnvelopeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSOAPEnvelopeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSOAPEnvelopeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSOAPEnvelopeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrSoapEnvelope {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SOAPEnvelope != nil {
			_ptr_pbstrSoapEnvelope := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SOAPEnvelope != nil {
					if err := o.SOAPEnvelope.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SOAPEnvelope, _ptr_pbstrSoapEnvelope); err != nil {
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

func (o *xxx_GetSOAPEnvelopeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrSoapEnvelope {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrSoapEnvelope := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SOAPEnvelope == nil {
				o.SOAPEnvelope = &oaut.String{}
			}
			if err := o.SOAPEnvelope.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrSoapEnvelope := func(ptr interface{}) { o.SOAPEnvelope = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SOAPEnvelope, _s_pbstrSoapEnvelope, _ptr_pbstrSoapEnvelope); err != nil {
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

// GetSOAPEnvelopeRequest structure represents the SoapEnvelope operation request
type GetSOAPEnvelopeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSOAPEnvelopeRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSOAPEnvelopeOperation) *xxx_GetSOAPEnvelopeOperation {
	if op == nil {
		op = &xxx_GetSOAPEnvelopeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetSOAPEnvelopeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSOAPEnvelopeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSOAPEnvelopeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSOAPEnvelopeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSOAPEnvelopeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSOAPEnvelopeResponse structure represents the SoapEnvelope operation response
type GetSOAPEnvelopeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbstrSoapEnvelope: A pointer to a BSTR containing a string of Unicode characters
	// that contains the SOAP envelope.
	SOAPEnvelope *oaut.String `idl:"name:pbstrSoapEnvelope" json:"soap_envelope"`
	// Return: The SoapEnvelope return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSOAPEnvelopeResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSOAPEnvelopeOperation) *xxx_GetSOAPEnvelopeOperation {
	if op == nil {
		op = &xxx_GetSOAPEnvelopeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.SOAPEnvelope = o.SOAPEnvelope
	op.Return = o.Return
	return op
}

func (o *GetSOAPEnvelopeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSOAPEnvelopeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SOAPEnvelope = op.SOAPEnvelope
	o.Return = op.Return
}
func (o *GetSOAPEnvelopeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSOAPEnvelopeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSOAPEnvelopeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetCompoundMessageOperation structure represents the CompoundMessage operation
type xxx_GetCompoundMessageOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	CompoundMessage *oaut.Variant  `idl:"name:pvarCompoundMessage" json:"compound_message"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCompoundMessageOperation) OpNum() int { return 95 }

func (o *xxx_GetCompoundMessageOperation) OpName() string { return "/IMSMQMessage3/v0/CompoundMessage" }

func (o *xxx_GetCompoundMessageOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCompoundMessageOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetCompoundMessageOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetCompoundMessageOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCompoundMessageOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pvarCompoundMessage {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.CompoundMessage != nil {
			_ptr_pvarCompoundMessage := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CompoundMessage != nil {
					if err := o.CompoundMessage.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CompoundMessage, _ptr_pvarCompoundMessage); err != nil {
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

func (o *xxx_GetCompoundMessageOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pvarCompoundMessage {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pvarCompoundMessage := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.CompoundMessage == nil {
				o.CompoundMessage = &oaut.Variant{}
			}
			if err := o.CompoundMessage.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pvarCompoundMessage := func(ptr interface{}) { o.CompoundMessage = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.CompoundMessage, _s_pvarCompoundMessage, _ptr_pvarCompoundMessage); err != nil {
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

// GetCompoundMessageRequest structure represents the CompoundMessage operation request
type GetCompoundMessageRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetCompoundMessageRequest) xxx_ToOp(ctx context.Context, op *xxx_GetCompoundMessageOperation) *xxx_GetCompoundMessageOperation {
	if op == nil {
		op = &xxx_GetCompoundMessageOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetCompoundMessageRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCompoundMessageOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetCompoundMessageRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetCompoundMessageRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCompoundMessageOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCompoundMessageResponse structure represents the CompoundMessage operation response
type GetCompoundMessageResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pvarCompoundMessage: A pointer to a VARIANT that contains an array of BYTES (VT_ARRAY | VT_UI1) that contains the entire contents of the SRMP message, including the SOAP envelope and all SOAP attachments.
	CompoundMessage *oaut.Variant `idl:"name:pvarCompoundMessage" json:"compound_message"`
	// Return: The CompoundMessage return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCompoundMessageResponse) xxx_ToOp(ctx context.Context, op *xxx_GetCompoundMessageOperation) *xxx_GetCompoundMessageOperation {
	if op == nil {
		op = &xxx_GetCompoundMessageOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.CompoundMessage = o.CompoundMessage
	op.Return = o.Return
	return op
}

func (o *GetCompoundMessageResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCompoundMessageOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.CompoundMessage = op.CompoundMessage
	o.Return = op.Return
}
func (o *GetCompoundMessageResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetCompoundMessageResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCompoundMessageOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSOAPHeaderOperation structure represents the SoapHeader operation
type xxx_SetSOAPHeaderOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	SOAPHeader *oaut.String   `idl:"name:bstrSoapHeader" json:"soap_header"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSOAPHeaderOperation) OpNum() int { return 96 }

func (o *xxx_SetSOAPHeaderOperation) OpName() string { return "/IMSMQMessage3/v0/SoapHeader" }

func (o *xxx_SetSOAPHeaderOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSOAPHeaderOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrSoapHeader {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SOAPHeader != nil {
			_ptr_bstrSoapHeader := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SOAPHeader != nil {
					if err := o.SOAPHeader.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SOAPHeader, _ptr_bstrSoapHeader); err != nil {
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

func (o *xxx_SetSOAPHeaderOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrSoapHeader {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrSoapHeader := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SOAPHeader == nil {
				o.SOAPHeader = &oaut.String{}
			}
			if err := o.SOAPHeader.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrSoapHeader := func(ptr interface{}) { o.SOAPHeader = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SOAPHeader, _s_bstrSoapHeader, _ptr_bstrSoapHeader); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSOAPHeaderOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSOAPHeaderOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetSOAPHeaderOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetSOAPHeaderRequest structure represents the SoapHeader operation request
type SetSOAPHeaderRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrSoapHeader: A BSTR that contains a string representation of the XML elements
	// that will be appended to the end of the SOAP header within the message.
	SOAPHeader *oaut.String `idl:"name:bstrSoapHeader" json:"soap_header"`
}

func (o *SetSOAPHeaderRequest) xxx_ToOp(ctx context.Context, op *xxx_SetSOAPHeaderOperation) *xxx_SetSOAPHeaderOperation {
	if op == nil {
		op = &xxx_SetSOAPHeaderOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.SOAPHeader = o.SOAPHeader
	return op
}

func (o *SetSOAPHeaderRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSOAPHeaderOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SOAPHeader = op.SOAPHeader
}
func (o *SetSOAPHeaderRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetSOAPHeaderRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSOAPHeaderOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSOAPHeaderResponse structure represents the SoapHeader operation response
type SetSOAPHeaderResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SoapHeader return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSOAPHeaderResponse) xxx_ToOp(ctx context.Context, op *xxx_SetSOAPHeaderOperation) *xxx_SetSOAPHeaderOperation {
	if op == nil {
		op = &xxx_SetSOAPHeaderOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetSOAPHeaderResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSOAPHeaderOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSOAPHeaderResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetSOAPHeaderResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSOAPHeaderOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSOAPBodyOperation structure represents the SoapBody operation
type xxx_SetSOAPBodyOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	SOAPBody *oaut.String   `idl:"name:bstrSoapBody" json:"soap_body"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSOAPBodyOperation) OpNum() int { return 97 }

func (o *xxx_SetSOAPBodyOperation) OpName() string { return "/IMSMQMessage3/v0/SoapBody" }

func (o *xxx_SetSOAPBodyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSOAPBodyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrSoapBody {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SOAPBody != nil {
			_ptr_bstrSoapBody := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SOAPBody != nil {
					if err := o.SOAPBody.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SOAPBody, _ptr_bstrSoapBody); err != nil {
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

func (o *xxx_SetSOAPBodyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrSoapBody {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrSoapBody := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SOAPBody == nil {
				o.SOAPBody = &oaut.String{}
			}
			if err := o.SOAPBody.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrSoapBody := func(ptr interface{}) { o.SOAPBody = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SOAPBody, _s_bstrSoapBody, _ptr_bstrSoapBody); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSOAPBodyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSOAPBodyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetSOAPBodyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetSOAPBodyRequest structure represents the SoapBody operation request
type SetSOAPBodyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrSoapBody: A BSTR that contains a string representation of the XML elements that
	// will be appended to the end of the SOAP body within the message.
	SOAPBody *oaut.String `idl:"name:bstrSoapBody" json:"soap_body"`
}

func (o *SetSOAPBodyRequest) xxx_ToOp(ctx context.Context, op *xxx_SetSOAPBodyOperation) *xxx_SetSOAPBodyOperation {
	if op == nil {
		op = &xxx_SetSOAPBodyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.SOAPBody = o.SOAPBody
	return op
}

func (o *SetSOAPBodyRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSOAPBodyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SOAPBody = op.SOAPBody
}
func (o *SetSOAPBodyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetSOAPBodyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSOAPBodyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSOAPBodyResponse structure represents the SoapBody operation response
type SetSOAPBodyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SoapBody return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSOAPBodyResponse) xxx_ToOp(ctx context.Context, op *xxx_SetSOAPBodyOperation) *xxx_SetSOAPBodyOperation {
	if op == nil {
		op = &xxx_SetSOAPBodyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetSOAPBodyResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSOAPBodyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSOAPBodyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetSOAPBodyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSOAPBodyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
