package imsmqmessage2

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

// IMSMQMessage2 server interface.
type Message2Server interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The Class method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return the represented Message.Class.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetClass(context.Context, *GetClassRequest) (*GetClassResponse, error)

	// PrivLevel operation.
	GetPrivacyLevel(context.Context, *GetPrivacyLevelRequest) (*GetPrivacyLevelResponse, error)

	// PrivLevel operation.
	SetPrivacyLevel(context.Context, *SetPrivacyLevelRequest) (*SetPrivacyLevelResponse, error)

	// AuthLevel operation.
	GetAuthLevel(context.Context, *GetAuthLevelRequest) (*GetAuthLevelResponse, error)

	// AuthLevel operation.
	SetAuthLevel(context.Context, *SetAuthLevelRequest) (*SetAuthLevelResponse, error)

	// The IsAuthenticated method is received by the server in an RPC_REQUEST packet. In
	// response, the server MUST return a BOOLEAN flag indicating whether the message was
	// authenticated by the Queue Manager that received the message.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetIsAuthenticated(context.Context, *GetIsAuthenticatedRequest) (*GetIsAuthenticatedResponse, error)

	// Delivery operation.
	GetDelivery(context.Context, *GetDeliveryRequest) (*GetDeliveryResponse, error)

	// Delivery operation.
	SetDelivery(context.Context, *SetDeliveryRequest) (*SetDeliveryResponse, error)

	// Trace operation.
	GetTrace(context.Context, *GetTraceRequest) (*GetTraceResponse, error)

	// Trace operation.
	SetTrace(context.Context, *SetTraceRequest) (*SetTraceResponse, error)

	// Priority operation.
	GetPriority(context.Context, *GetPriorityRequest) (*GetPriorityResponse, error)

	// Priority operation.
	SetPriority(context.Context, *SetPriorityRequest) (*SetPriorityResponse, error)

	// Journal operation.
	GetJournal(context.Context, *GetJournalRequest) (*GetJournalResponse, error)

	// Journal operation.
	SetJournal(context.Context, *SetJournalRequest) (*SetJournalResponse, error)

	// ResponseQueueInfo_v1 operation.
	GetResponseQueueInfoV1(context.Context, *GetResponseQueueInfoV1Request) (*GetResponseQueueInfoV1Response, error)

	// ResponseQueueInfo_v1 operation.
	SetResponseQueueInfoV1(context.Context, *SetResponseQueueInfoV1Request) (*SetResponseQueueInfoV1Response, error)

	// AppSpecific operation.
	GetAppSpecific(context.Context, *GetAppSpecificRequest) (*GetAppSpecificResponse, error)

	// AppSpecific operation.
	SetAppSpecific(context.Context, *SetAppSpecificRequest) (*SetAppSpecificResponse, error)

	// The SourceMachineGuid method is received by the server in an RPC_REQUEST packet.
	// In response, the server MUST return the represented Message.SourceMachineIdentifier.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetSourceMachineGUID(context.Context, *GetSourceMachineGUIDRequest) (*GetSourceMachineGUIDResponse, error)

	// The BodyLength method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return the number of bytes in the represented Message.Body.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetBodyLength(context.Context, *GetBodyLengthRequest) (*GetBodyLengthResponse, error)

	// Body operation.
	GetBody(context.Context, *GetBodyRequest) (*GetBodyResponse, error)

	// Body operation.
	SetBody(context.Context, *SetBodyRequest) (*SetBodyResponse, error)

	// AdminQueueInfo_v1 operation.
	GetAdminQueueInfoV1(context.Context, *GetAdminQueueInfoV1Request) (*GetAdminQueueInfoV1Response, error)

	// AdminQueueInfo_v1 operation.
	SetAdminQueueInfoV1(context.Context, *SetAdminQueueInfoV1Request) (*SetAdminQueueInfoV1Response, error)

	// The Id method is received by the server in an RPC_REQUEST packet. In response, the
	// server MUST return the represented Message.Identifier.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetID(context.Context, *GetIDRequest) (*GetIDResponse, error)

	// CorrelationId operation.
	GetCorrelationID(context.Context, *GetCorrelationIDRequest) (*GetCorrelationIDResponse, error)

	// CorrelationId operation.
	SetCorrelationID(context.Context, *SetCorrelationIDRequest) (*SetCorrelationIDResponse, error)

	// Ack operation.
	GetAck(context.Context, *GetAckRequest) (*GetAckResponse, error)

	// Ack operation.
	SetAck(context.Context, *SetAckRequest) (*SetAckResponse, error)

	// Label operation.
	GetLabel(context.Context, *GetLabelRequest) (*GetLabelResponse, error)

	// Label operation.
	SetLabel(context.Context, *SetLabelRequest) (*SetLabelResponse, error)

	// MaxTimeToReachQueue operation.
	GetMaxTimeToReachQueue(context.Context, *GetMaxTimeToReachQueueRequest) (*GetMaxTimeToReachQueueResponse, error)

	// MaxTimeToReachQueue operation.
	SetMaxTimeToReachQueue(context.Context, *SetMaxTimeToReachQueueRequest) (*SetMaxTimeToReachQueueResponse, error)

	// MaxTimeToReceive operation.
	GetMaxTimeToReceive(context.Context, *GetMaxTimeToReceiveRequest) (*GetMaxTimeToReceiveResponse, error)

	// MaxTimeToReceive operation.
	SetMaxTimeToReceive(context.Context, *SetMaxTimeToReceiveRequest) (*SetMaxTimeToReceiveResponse, error)

	// HashAlgorithm operation.
	GetHashAlgorithm(context.Context, *GetHashAlgorithmRequest) (*GetHashAlgorithmResponse, error)

	// HashAlgorithm operation.
	SetHashAlgorithm(context.Context, *SetHashAlgorithmRequest) (*SetHashAlgorithmResponse, error)

	// EncryptAlgorithm operation.
	GetEncryptAlgorithm(context.Context, *GetEncryptAlgorithmRequest) (*GetEncryptAlgorithmResponse, error)

	// EncryptAlgorithm operation.
	SetEncryptAlgorithm(context.Context, *SetEncryptAlgorithmRequest) (*SetEncryptAlgorithmResponse, error)

	// The SentTime method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return the represented Message.SentTime.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetSentTime(context.Context, *GetSentTimeRequest) (*GetSentTimeResponse, error)

	// The ArrivedTime method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return the represented Message.ArrivalTime.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetArrivedTime(context.Context, *GetArrivedTimeRequest) (*GetArrivedTimeResponse, error)

	// The DestinationQueueInfo method is received by the server in an RPC_REQUEST packet.
	// In response, the server MUST return an IMSMQQueueInfo interface pointer to an MSMQQueueInfo
	// object that represents the Queue identified by the represented Message.DestinationQueueFormatName.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetDestinationQueueInfo(context.Context, *GetDestinationQueueInfoRequest) (*GetDestinationQueueInfoResponse, error)

	// SenderCertificate operation.
	GetSenderCertificate(context.Context, *GetSenderCertificateRequest) (*GetSenderCertificateResponse, error)

	// SenderCertificate operation.
	SetSenderCertificate(context.Context, *SetSenderCertificateRequest) (*SetSenderCertificateResponse, error)

	// SenderId operation.
	GetSenderID(context.Context, *GetSenderIDRequest) (*GetSenderIDResponse, error)

	// SenderIdType operation.
	GetSenderIDType(context.Context, *GetSenderIDTypeRequest) (*GetSenderIDTypeResponse, error)

	// SenderIdType operation.
	SetSenderIDType(context.Context, *SetSenderIDTypeRequest) (*SetSenderIDTypeResponse, error)

	// The Send method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST send a message.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	Send(context.Context, *SendRequest) (*SendResponse, error)

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
	AttachCurrentSecurityContext(context.Context, *AttachCurrentSecurityContextRequest) (*AttachCurrentSecurityContextResponse, error)

	// The SenderVersion method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return a numeric value that indicates the version of transfer used
	// to send the message.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetSenderVersion(context.Context, *GetSenderVersionRequest) (*GetSenderVersionResponse, error)

	// Extension operation.
	GetExtension(context.Context, *GetExtensionRequest) (*GetExtensionResponse, error)

	// Extension operation.
	SetExtension(context.Context, *SetExtensionRequest) (*SetExtensionResponse, error)

	// ConnectorTypeGuid operation.
	GetConnectorTypeGUID(context.Context, *GetConnectorTypeGUIDRequest) (*GetConnectorTypeGUIDResponse, error)

	// ConnectorTypeGuid operation.
	SetConnectorTypeGUID(context.Context, *SetConnectorTypeGUIDRequest) (*SetConnectorTypeGUIDResponse, error)

	// The TransactionStatusQueueInfo method is received by the server in an RPC_REQUEST
	// packet. In response, the server MUST return an IMSMQQueueInfo interface pointer to
	// an MSMQQueueInfo object that represents the queue identified by the represented Message.TransactionStatusQueueFormatName.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetTransactionStatusQueueInfo(context.Context, *GetTransactionStatusQueueInfoRequest) (*GetTransactionStatusQueueInfoResponse, error)

	// DestinationSymmetricKey operation.
	GetDestinationSymmetricKey(context.Context, *GetDestinationSymmetricKeyRequest) (*GetDestinationSymmetricKeyResponse, error)

	// DestinationSymmetricKey operation.
	SetDestinationSymmetricKey(context.Context, *SetDestinationSymmetricKeyRequest) (*SetDestinationSymmetricKeyResponse, error)

	// Signature operation.
	GetSignature(context.Context, *GetSignatureRequest) (*GetSignatureResponse, error)

	// Signature operation.
	SetSignature(context.Context, *SetSignatureRequest) (*SetSignatureResponse, error)

	// AuthenticationProviderType operation.
	GetAuthenticationProviderType(context.Context, *GetAuthenticationProviderTypeRequest) (*GetAuthenticationProviderTypeResponse, error)

	// AuthenticationProviderType operation.
	SetAuthenticationProviderType(context.Context, *SetAuthenticationProviderTypeRequest) (*SetAuthenticationProviderTypeResponse, error)

	// AuthenticationProviderName operation.
	GetAuthenticationProviderName(context.Context, *GetAuthenticationProviderNameRequest) (*GetAuthenticationProviderNameResponse, error)

	// AuthenticationProviderName operation.
	SetAuthenticationProviderName(context.Context, *SetAuthenticationProviderNameRequest) (*SetAuthenticationProviderNameResponse, error)

	// SenderId operation.
	SetSenderID(context.Context, *SetSenderIDRequest) (*SetSenderIDResponse, error)

	// MsgClass operation.
	GetMessageClass(context.Context, *GetMessageClassRequest) (*GetMessageClassResponse, error)

	// MsgClass operation.
	SetMessageClass(context.Context, *SetMessageClassRequest) (*SetMessageClassResponse, error)

	// Properties operation.
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)

	// The TransactionId method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return the represented Message.TransactionalMessageSequenceIdentifier.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetTransactionID(context.Context, *GetTransactionIDRequest) (*GetTransactionIDResponse, error)

	// The IsFirstInTransaction method is received by the server in an RPC_REQUEST packet.
	// In response, the server MUST return the represented Message.FirstInTransaction.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetIsFirstInTransaction(context.Context, *GetIsFirstInTransactionRequest) (*GetIsFirstInTransactionResponse, error)

	// The IsLastInTransaction method is received by the server in an RPC_REQUEST packet.
	// In response, the server MUST return the represented Message.LastInTransaction.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetIsLastInTransaction(context.Context, *GetIsLastInTransactionRequest) (*GetIsLastInTransactionResponse, error)

	// ResponseQueueInfo operation.
	GetResponseQueueInfo(context.Context, *GetResponseQueueInfoRequest) (*GetResponseQueueInfoResponse, error)

	// ResponseQueueInfo operation.
	SetByRefResponseQueueInfo(context.Context, *SetByRefResponseQueueInfoRequest) (*SetByRefResponseQueueInfoResponse, error)

	// AdminQueueInfo operation.
	GetAdminQueueInfo(context.Context, *GetAdminQueueInfoRequest) (*GetAdminQueueInfoResponse, error)

	// AdminQueueInfo operation.
	SetByRefAdminQueueInfo(context.Context, *SetByRefAdminQueueInfoRequest) (*SetByRefAdminQueueInfoResponse, error)

	// The ReceivedAuthenticationLevel method is received by the server in an RPC_REQUEST
	// packet. In response, the server MUST return the represented Message.AuthenticationLevel.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetReceivedAuthenticationLevel(context.Context, *GetReceivedAuthenticationLevelRequest) (*GetReceivedAuthenticationLevelResponse, error)
}

func RegisterMessage2Server(conn dcerpc.Conn, o Message2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewMessage2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Message2SyntaxV0_0))...)
}

func NewMessage2ServerHandle(o Message2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Message2ServerHandle(ctx, o, opNum, r)
	}
}

func Message2ServerHandle(ctx context.Context, o Message2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Class
		op := &xxx_GetClassOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetClassRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetClass(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // PrivLevel
		op := &xxx_GetPrivacyLevelOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPrivacyLevelRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPrivacyLevel(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // PrivLevel
		op := &xxx_SetPrivacyLevelOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetPrivacyLevelRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetPrivacyLevel(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // AuthLevel
		op := &xxx_GetAuthLevelOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAuthLevelRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAuthLevel(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // AuthLevel
		op := &xxx_SetAuthLevelOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetAuthLevelRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetAuthLevel(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // IsAuthenticated
		op := &xxx_GetIsAuthenticatedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsAuthenticatedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsAuthenticated(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // Delivery
		op := &xxx_GetDeliveryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDeliveryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDelivery(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // Delivery
		op := &xxx_SetDeliveryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetDeliveryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetDelivery(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // Trace
		op := &xxx_GetTraceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTraceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTrace(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // Trace
		op := &xxx_SetTraceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetTraceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetTrace(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // Priority
		op := &xxx_GetPriorityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPriorityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPriority(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // Priority
		op := &xxx_SetPriorityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetPriorityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetPriority(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // Journal
		op := &xxx_GetJournalOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetJournalRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetJournal(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // Journal
		op := &xxx_SetJournalOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetJournalRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetJournal(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // ResponseQueueInfo_v1
		op := &xxx_GetResponseQueueInfoV1Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetResponseQueueInfoV1Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetResponseQueueInfoV1(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // ResponseQueueInfo_v1
		op := &xxx_SetResponseQueueInfoV1Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetResponseQueueInfoV1Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetResponseQueueInfoV1(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // AppSpecific
		op := &xxx_GetAppSpecificOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAppSpecificRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAppSpecific(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // AppSpecific
		op := &xxx_SetAppSpecificOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetAppSpecificRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetAppSpecific(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // SourceMachineGuid
		op := &xxx_GetSourceMachineGUIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSourceMachineGUIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSourceMachineGUID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 26: // BodyLength
		op := &xxx_GetBodyLengthOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetBodyLengthRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetBodyLength(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // Body
		op := &xxx_GetBodyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetBodyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetBody(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 28: // Body
		op := &xxx_SetBodyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetBodyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetBody(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 29: // AdminQueueInfo_v1
		op := &xxx_GetAdminQueueInfoV1Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAdminQueueInfoV1Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAdminQueueInfoV1(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // AdminQueueInfo_v1
		op := &xxx_SetAdminQueueInfoV1Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetAdminQueueInfoV1Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetAdminQueueInfoV1(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 31: // Id
		op := &xxx_GetIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 32: // CorrelationId
		op := &xxx_GetCorrelationIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCorrelationIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCorrelationID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 33: // CorrelationId
		op := &xxx_SetCorrelationIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetCorrelationIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetCorrelationID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 34: // Ack
		op := &xxx_GetAckOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAckRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAck(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 35: // Ack
		op := &xxx_SetAckOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetAckRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetAck(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 36: // Label
		op := &xxx_GetLabelOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLabelRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLabel(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 37: // Label
		op := &xxx_SetLabelOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetLabelRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetLabel(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 38: // MaxTimeToReachQueue
		op := &xxx_GetMaxTimeToReachQueueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMaxTimeToReachQueueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMaxTimeToReachQueue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 39: // MaxTimeToReachQueue
		op := &xxx_SetMaxTimeToReachQueueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMaxTimeToReachQueueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMaxTimeToReachQueue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 40: // MaxTimeToReceive
		op := &xxx_GetMaxTimeToReceiveOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMaxTimeToReceiveRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMaxTimeToReceive(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 41: // MaxTimeToReceive
		op := &xxx_SetMaxTimeToReceiveOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMaxTimeToReceiveRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMaxTimeToReceive(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 42: // HashAlgorithm
		op := &xxx_GetHashAlgorithmOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetHashAlgorithmRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetHashAlgorithm(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 43: // HashAlgorithm
		op := &xxx_SetHashAlgorithmOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetHashAlgorithmRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetHashAlgorithm(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 44: // EncryptAlgorithm
		op := &xxx_GetEncryptAlgorithmOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetEncryptAlgorithmRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetEncryptAlgorithm(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 45: // EncryptAlgorithm
		op := &xxx_SetEncryptAlgorithmOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetEncryptAlgorithmRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetEncryptAlgorithm(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 46: // SentTime
		op := &xxx_GetSentTimeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSentTimeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSentTime(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 47: // ArrivedTime
		op := &xxx_GetArrivedTimeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetArrivedTimeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetArrivedTime(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 48: // DestinationQueueInfo
		op := &xxx_GetDestinationQueueInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDestinationQueueInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDestinationQueueInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 49: // SenderCertificate
		op := &xxx_GetSenderCertificateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSenderCertificateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSenderCertificate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 50: // SenderCertificate
		op := &xxx_SetSenderCertificateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSenderCertificateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSenderCertificate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 51: // SenderId
		op := &xxx_GetSenderIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSenderIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSenderID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 52: // SenderIdType
		op := &xxx_GetSenderIDTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSenderIDTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSenderIDType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 53: // SenderIdType
		op := &xxx_SetSenderIDTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSenderIDTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSenderIDType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 54: // Send
		op := &xxx_SendOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SendRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Send(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 55: // AttachCurrentSecurityContext
		op := &xxx_AttachCurrentSecurityContextOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AttachCurrentSecurityContextRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AttachCurrentSecurityContext(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 56: // SenderVersion
		op := &xxx_GetSenderVersionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSenderVersionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSenderVersion(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 57: // Extension
		op := &xxx_GetExtensionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetExtensionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetExtension(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 58: // Extension
		op := &xxx_SetExtensionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetExtensionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetExtension(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 59: // ConnectorTypeGuid
		op := &xxx_GetConnectorTypeGUIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetConnectorTypeGUIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetConnectorTypeGUID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 60: // ConnectorTypeGuid
		op := &xxx_SetConnectorTypeGUIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetConnectorTypeGUIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetConnectorTypeGUID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 61: // TransactionStatusQueueInfo
		op := &xxx_GetTransactionStatusQueueInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTransactionStatusQueueInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTransactionStatusQueueInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 62: // DestinationSymmetricKey
		op := &xxx_GetDestinationSymmetricKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDestinationSymmetricKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDestinationSymmetricKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 63: // DestinationSymmetricKey
		op := &xxx_SetDestinationSymmetricKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetDestinationSymmetricKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetDestinationSymmetricKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 64: // Signature
		op := &xxx_GetSignatureOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSignatureRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSignature(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 65: // Signature
		op := &xxx_SetSignatureOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSignatureRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSignature(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 66: // AuthenticationProviderType
		op := &xxx_GetAuthenticationProviderTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAuthenticationProviderTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAuthenticationProviderType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 67: // AuthenticationProviderType
		op := &xxx_SetAuthenticationProviderTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetAuthenticationProviderTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetAuthenticationProviderType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 68: // AuthenticationProviderName
		op := &xxx_GetAuthenticationProviderNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAuthenticationProviderNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAuthenticationProviderName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 69: // AuthenticationProviderName
		op := &xxx_SetAuthenticationProviderNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetAuthenticationProviderNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetAuthenticationProviderName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 70: // SenderId
		op := &xxx_SetSenderIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSenderIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSenderID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 71: // MsgClass
		op := &xxx_GetMessageClassOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMessageClassRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMessageClass(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 72: // MsgClass
		op := &xxx_SetMessageClassOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMessageClassRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMessageClass(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 73: // Properties
		op := &xxx_GetPropertiesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPropertiesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetProperties(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 74: // TransactionId
		op := &xxx_GetTransactionIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTransactionIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTransactionID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 75: // IsFirstInTransaction
		op := &xxx_GetIsFirstInTransactionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsFirstInTransactionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsFirstInTransaction(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 76: // IsLastInTransaction
		op := &xxx_GetIsLastInTransactionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsLastInTransactionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsLastInTransaction(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 77: // ResponseQueueInfo
		op := &xxx_GetResponseQueueInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetResponseQueueInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetResponseQueueInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 78: // ResponseQueueInfo
		op := &xxx_SetByRefResponseQueueInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetByRefResponseQueueInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetByRefResponseQueueInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 79: // AdminQueueInfo
		op := &xxx_GetAdminQueueInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAdminQueueInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAdminQueueInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 80: // AdminQueueInfo
		op := &xxx_SetByRefAdminQueueInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetByRefAdminQueueInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetByRefAdminQueueInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 81: // ReceivedAuthenticationLevel
		op := &xxx_GetReceivedAuthenticationLevelOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetReceivedAuthenticationLevelRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetReceivedAuthenticationLevel(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMSMQMessage2
type UnimplementedMessage2Server struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedMessage2Server) GetClass(context.Context, *GetClassRequest) (*GetClassResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetPrivacyLevel(context.Context, *GetPrivacyLevelRequest) (*GetPrivacyLevelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) SetPrivacyLevel(context.Context, *SetPrivacyLevelRequest) (*SetPrivacyLevelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetAuthLevel(context.Context, *GetAuthLevelRequest) (*GetAuthLevelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) SetAuthLevel(context.Context, *SetAuthLevelRequest) (*SetAuthLevelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetIsAuthenticated(context.Context, *GetIsAuthenticatedRequest) (*GetIsAuthenticatedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetDelivery(context.Context, *GetDeliveryRequest) (*GetDeliveryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) SetDelivery(context.Context, *SetDeliveryRequest) (*SetDeliveryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetTrace(context.Context, *GetTraceRequest) (*GetTraceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) SetTrace(context.Context, *SetTraceRequest) (*SetTraceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetPriority(context.Context, *GetPriorityRequest) (*GetPriorityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) SetPriority(context.Context, *SetPriorityRequest) (*SetPriorityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetJournal(context.Context, *GetJournalRequest) (*GetJournalResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) SetJournal(context.Context, *SetJournalRequest) (*SetJournalResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetResponseQueueInfoV1(context.Context, *GetResponseQueueInfoV1Request) (*GetResponseQueueInfoV1Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) SetResponseQueueInfoV1(context.Context, *SetResponseQueueInfoV1Request) (*SetResponseQueueInfoV1Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetAppSpecific(context.Context, *GetAppSpecificRequest) (*GetAppSpecificResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) SetAppSpecific(context.Context, *SetAppSpecificRequest) (*SetAppSpecificResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetSourceMachineGUID(context.Context, *GetSourceMachineGUIDRequest) (*GetSourceMachineGUIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetBodyLength(context.Context, *GetBodyLengthRequest) (*GetBodyLengthResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetBody(context.Context, *GetBodyRequest) (*GetBodyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) SetBody(context.Context, *SetBodyRequest) (*SetBodyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetAdminQueueInfoV1(context.Context, *GetAdminQueueInfoV1Request) (*GetAdminQueueInfoV1Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) SetAdminQueueInfoV1(context.Context, *SetAdminQueueInfoV1Request) (*SetAdminQueueInfoV1Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetID(context.Context, *GetIDRequest) (*GetIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetCorrelationID(context.Context, *GetCorrelationIDRequest) (*GetCorrelationIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) SetCorrelationID(context.Context, *SetCorrelationIDRequest) (*SetCorrelationIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetAck(context.Context, *GetAckRequest) (*GetAckResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) SetAck(context.Context, *SetAckRequest) (*SetAckResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetLabel(context.Context, *GetLabelRequest) (*GetLabelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) SetLabel(context.Context, *SetLabelRequest) (*SetLabelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetMaxTimeToReachQueue(context.Context, *GetMaxTimeToReachQueueRequest) (*GetMaxTimeToReachQueueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) SetMaxTimeToReachQueue(context.Context, *SetMaxTimeToReachQueueRequest) (*SetMaxTimeToReachQueueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetMaxTimeToReceive(context.Context, *GetMaxTimeToReceiveRequest) (*GetMaxTimeToReceiveResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) SetMaxTimeToReceive(context.Context, *SetMaxTimeToReceiveRequest) (*SetMaxTimeToReceiveResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetHashAlgorithm(context.Context, *GetHashAlgorithmRequest) (*GetHashAlgorithmResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) SetHashAlgorithm(context.Context, *SetHashAlgorithmRequest) (*SetHashAlgorithmResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetEncryptAlgorithm(context.Context, *GetEncryptAlgorithmRequest) (*GetEncryptAlgorithmResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) SetEncryptAlgorithm(context.Context, *SetEncryptAlgorithmRequest) (*SetEncryptAlgorithmResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetSentTime(context.Context, *GetSentTimeRequest) (*GetSentTimeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetArrivedTime(context.Context, *GetArrivedTimeRequest) (*GetArrivedTimeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetDestinationQueueInfo(context.Context, *GetDestinationQueueInfoRequest) (*GetDestinationQueueInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetSenderCertificate(context.Context, *GetSenderCertificateRequest) (*GetSenderCertificateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) SetSenderCertificate(context.Context, *SetSenderCertificateRequest) (*SetSenderCertificateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetSenderID(context.Context, *GetSenderIDRequest) (*GetSenderIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetSenderIDType(context.Context, *GetSenderIDTypeRequest) (*GetSenderIDTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) SetSenderIDType(context.Context, *SetSenderIDTypeRequest) (*SetSenderIDTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) Send(context.Context, *SendRequest) (*SendResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) AttachCurrentSecurityContext(context.Context, *AttachCurrentSecurityContextRequest) (*AttachCurrentSecurityContextResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetSenderVersion(context.Context, *GetSenderVersionRequest) (*GetSenderVersionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetExtension(context.Context, *GetExtensionRequest) (*GetExtensionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) SetExtension(context.Context, *SetExtensionRequest) (*SetExtensionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetConnectorTypeGUID(context.Context, *GetConnectorTypeGUIDRequest) (*GetConnectorTypeGUIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) SetConnectorTypeGUID(context.Context, *SetConnectorTypeGUIDRequest) (*SetConnectorTypeGUIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetTransactionStatusQueueInfo(context.Context, *GetTransactionStatusQueueInfoRequest) (*GetTransactionStatusQueueInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetDestinationSymmetricKey(context.Context, *GetDestinationSymmetricKeyRequest) (*GetDestinationSymmetricKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) SetDestinationSymmetricKey(context.Context, *SetDestinationSymmetricKeyRequest) (*SetDestinationSymmetricKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetSignature(context.Context, *GetSignatureRequest) (*GetSignatureResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) SetSignature(context.Context, *SetSignatureRequest) (*SetSignatureResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetAuthenticationProviderType(context.Context, *GetAuthenticationProviderTypeRequest) (*GetAuthenticationProviderTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) SetAuthenticationProviderType(context.Context, *SetAuthenticationProviderTypeRequest) (*SetAuthenticationProviderTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetAuthenticationProviderName(context.Context, *GetAuthenticationProviderNameRequest) (*GetAuthenticationProviderNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) SetAuthenticationProviderName(context.Context, *SetAuthenticationProviderNameRequest) (*SetAuthenticationProviderNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) SetSenderID(context.Context, *SetSenderIDRequest) (*SetSenderIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetMessageClass(context.Context, *GetMessageClassRequest) (*GetMessageClassResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) SetMessageClass(context.Context, *SetMessageClassRequest) (*SetMessageClassResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetTransactionID(context.Context, *GetTransactionIDRequest) (*GetTransactionIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetIsFirstInTransaction(context.Context, *GetIsFirstInTransactionRequest) (*GetIsFirstInTransactionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetIsLastInTransaction(context.Context, *GetIsLastInTransactionRequest) (*GetIsLastInTransactionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetResponseQueueInfo(context.Context, *GetResponseQueueInfoRequest) (*GetResponseQueueInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) SetByRefResponseQueueInfo(context.Context, *SetByRefResponseQueueInfoRequest) (*SetByRefResponseQueueInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetAdminQueueInfo(context.Context, *GetAdminQueueInfoRequest) (*GetAdminQueueInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) SetByRefAdminQueueInfo(context.Context, *SetByRefAdminQueueInfoRequest) (*SetByRefAdminQueueInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessage2Server) GetReceivedAuthenticationLevel(context.Context, *GetReceivedAuthenticationLevelRequest) (*GetReceivedAuthenticationLevelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ Message2Server = (*UnimplementedMessage2Server)(nil)
