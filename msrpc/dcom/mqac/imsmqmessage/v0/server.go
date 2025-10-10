package imsmqmessage

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

// IMSMQMessage server interface.
type ImsmqMessageServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// Class operation.
	GetClass(context.Context, *GetClassRequest) (*GetClassResponse, error)

	// PrivLevel operation.
	GetPrivLevel(context.Context, *GetPrivLevelRequest) (*GetPrivLevelResponse, error)

	// PrivLevel operation.
	SetPrivLevel(context.Context, *SetPrivLevelRequest) (*SetPrivLevelResponse, error)

	// AuthLevel operation.
	GetAuthLevel(context.Context, *GetAuthLevelRequest) (*GetAuthLevelResponse, error)

	// AuthLevel operation.
	SetAuthLevel(context.Context, *SetAuthLevelRequest) (*SetAuthLevelResponse, error)

	// IsAuthenticated operation.
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

	// ResponseQueueInfo operation.
	GetResponseQueueInfo(context.Context, *GetResponseQueueInfoRequest) (*GetResponseQueueInfoResponse, error)

	// ResponseQueueInfo operation.
	SetByRefResponseQueueInfo(context.Context, *SetByRefResponseQueueInfoRequest) (*SetByRefResponseQueueInfoResponse, error)

	// AppSpecific operation.
	GetAppSpecific(context.Context, *GetAppSpecificRequest) (*GetAppSpecificResponse, error)

	// AppSpecific operation.
	SetAppSpecific(context.Context, *SetAppSpecificRequest) (*SetAppSpecificResponse, error)

	// SourceMachineGuid operation.
	GetSourceMachineGUID(context.Context, *GetSourceMachineGUIDRequest) (*GetSourceMachineGUIDResponse, error)

	// BodyLength operation.
	GetBodyLength(context.Context, *GetBodyLengthRequest) (*GetBodyLengthResponse, error)

	// Body operation.
	GetBody(context.Context, *GetBodyRequest) (*GetBodyResponse, error)

	// Body operation.
	SetBody(context.Context, *SetBodyRequest) (*SetBodyResponse, error)

	// AdminQueueInfo operation.
	GetAdminQueueInfo(context.Context, *GetAdminQueueInfoRequest) (*GetAdminQueueInfoResponse, error)

	// AdminQueueInfo operation.
	SetByRefAdminQueueInfo(context.Context, *SetByRefAdminQueueInfoRequest) (*SetByRefAdminQueueInfoResponse, error)

	// Id operation.
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

	// SentTime operation.
	GetSentTime(context.Context, *GetSentTimeRequest) (*GetSentTimeResponse, error)

	// ArrivedTime operation.
	GetArrivedTime(context.Context, *GetArrivedTimeRequest) (*GetArrivedTimeResponse, error)

	// DestinationQueueInfo operation.
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

	// Send operation.
	Send(context.Context, *SendRequest) (*SendResponse, error)

	// AttachCurrentSecurityContext operation.
	AttachCurrentSecurityContext(context.Context, *AttachCurrentSecurityContextRequest) (*AttachCurrentSecurityContextResponse, error)
}

func RegisterImsmqMessageServer(conn dcerpc.Conn, o ImsmqMessageServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewImsmqMessageServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ImsmqMessageSyntaxV0_0))...)
}

func NewImsmqMessageServerHandle(o ImsmqMessageServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ImsmqMessageServerHandle(ctx, o, opNum, r)
	}
}

func ImsmqMessageServerHandle(ctx context.Context, o ImsmqMessageServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
		op := &xxx_GetPrivLevelOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPrivLevelRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPrivLevel(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // PrivLevel
		op := &xxx_SetPrivLevelOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetPrivLevelRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetPrivLevel(ctx, req)
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
	case 21: // ResponseQueueInfo
		op := &xxx_GetResponseQueueInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetResponseQueueInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetResponseQueueInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // ResponseQueueInfo
		op := &xxx_SetByRefResponseQueueInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetByRefResponseQueueInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetByRefResponseQueueInfo(ctx, req)
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
	case 29: // AdminQueueInfo
		op := &xxx_GetAdminQueueInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAdminQueueInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAdminQueueInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // AdminQueueInfo
		op := &xxx_SetByRefAdminQueueInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetByRefAdminQueueInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetByRefAdminQueueInfo(ctx, req)
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
	}
	return nil, nil
}

// Unimplemented IMSMQMessage
type UnimplementedImsmqMessageServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedImsmqMessageServer) GetClass(context.Context, *GetClassRequest) (*GetClassResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) GetPrivLevel(context.Context, *GetPrivLevelRequest) (*GetPrivLevelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) SetPrivLevel(context.Context, *SetPrivLevelRequest) (*SetPrivLevelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) GetAuthLevel(context.Context, *GetAuthLevelRequest) (*GetAuthLevelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) SetAuthLevel(context.Context, *SetAuthLevelRequest) (*SetAuthLevelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) GetIsAuthenticated(context.Context, *GetIsAuthenticatedRequest) (*GetIsAuthenticatedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) GetDelivery(context.Context, *GetDeliveryRequest) (*GetDeliveryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) SetDelivery(context.Context, *SetDeliveryRequest) (*SetDeliveryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) GetTrace(context.Context, *GetTraceRequest) (*GetTraceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) SetTrace(context.Context, *SetTraceRequest) (*SetTraceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) GetPriority(context.Context, *GetPriorityRequest) (*GetPriorityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) SetPriority(context.Context, *SetPriorityRequest) (*SetPriorityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) GetJournal(context.Context, *GetJournalRequest) (*GetJournalResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) SetJournal(context.Context, *SetJournalRequest) (*SetJournalResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) GetResponseQueueInfo(context.Context, *GetResponseQueueInfoRequest) (*GetResponseQueueInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) SetByRefResponseQueueInfo(context.Context, *SetByRefResponseQueueInfoRequest) (*SetByRefResponseQueueInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) GetAppSpecific(context.Context, *GetAppSpecificRequest) (*GetAppSpecificResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) SetAppSpecific(context.Context, *SetAppSpecificRequest) (*SetAppSpecificResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) GetSourceMachineGUID(context.Context, *GetSourceMachineGUIDRequest) (*GetSourceMachineGUIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) GetBodyLength(context.Context, *GetBodyLengthRequest) (*GetBodyLengthResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) GetBody(context.Context, *GetBodyRequest) (*GetBodyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) SetBody(context.Context, *SetBodyRequest) (*SetBodyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) GetAdminQueueInfo(context.Context, *GetAdminQueueInfoRequest) (*GetAdminQueueInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) SetByRefAdminQueueInfo(context.Context, *SetByRefAdminQueueInfoRequest) (*SetByRefAdminQueueInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) GetID(context.Context, *GetIDRequest) (*GetIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) GetCorrelationID(context.Context, *GetCorrelationIDRequest) (*GetCorrelationIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) SetCorrelationID(context.Context, *SetCorrelationIDRequest) (*SetCorrelationIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) GetAck(context.Context, *GetAckRequest) (*GetAckResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) SetAck(context.Context, *SetAckRequest) (*SetAckResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) GetLabel(context.Context, *GetLabelRequest) (*GetLabelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) SetLabel(context.Context, *SetLabelRequest) (*SetLabelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) GetMaxTimeToReachQueue(context.Context, *GetMaxTimeToReachQueueRequest) (*GetMaxTimeToReachQueueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) SetMaxTimeToReachQueue(context.Context, *SetMaxTimeToReachQueueRequest) (*SetMaxTimeToReachQueueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) GetMaxTimeToReceive(context.Context, *GetMaxTimeToReceiveRequest) (*GetMaxTimeToReceiveResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) SetMaxTimeToReceive(context.Context, *SetMaxTimeToReceiveRequest) (*SetMaxTimeToReceiveResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) GetHashAlgorithm(context.Context, *GetHashAlgorithmRequest) (*GetHashAlgorithmResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) SetHashAlgorithm(context.Context, *SetHashAlgorithmRequest) (*SetHashAlgorithmResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) GetEncryptAlgorithm(context.Context, *GetEncryptAlgorithmRequest) (*GetEncryptAlgorithmResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) SetEncryptAlgorithm(context.Context, *SetEncryptAlgorithmRequest) (*SetEncryptAlgorithmResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) GetSentTime(context.Context, *GetSentTimeRequest) (*GetSentTimeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) GetArrivedTime(context.Context, *GetArrivedTimeRequest) (*GetArrivedTimeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) GetDestinationQueueInfo(context.Context, *GetDestinationQueueInfoRequest) (*GetDestinationQueueInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) GetSenderCertificate(context.Context, *GetSenderCertificateRequest) (*GetSenderCertificateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) SetSenderCertificate(context.Context, *SetSenderCertificateRequest) (*SetSenderCertificateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) GetSenderID(context.Context, *GetSenderIDRequest) (*GetSenderIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) GetSenderIDType(context.Context, *GetSenderIDTypeRequest) (*GetSenderIDTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) SetSenderIDType(context.Context, *SetSenderIDTypeRequest) (*SetSenderIDTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) Send(context.Context, *SendRequest) (*SendResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqMessageServer) AttachCurrentSecurityContext(context.Context, *AttachCurrentSecurityContextRequest) (*AttachCurrentSecurityContextResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ImsmqMessageServer = (*UnimplementedImsmqMessageServer)(nil)
