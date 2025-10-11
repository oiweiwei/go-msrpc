package imsmqqueueinfo3

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

// IMSMQQueueInfo3 server interface.
type QueueInfo3Server interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The QueueGuid method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns the Queue.Identifier that uniquely identifies the referenced queue.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	GetQueueGUID(context.Context, *GetQueueGUIDRequest) (*GetQueueGUIDResponse, error)

	// ServiceTypeGuid operation.
	GetServiceTypeGUID(context.Context, *GetServiceTypeGUIDRequest) (*GetServiceTypeGUIDResponse, error)

	// ServiceTypeGuid operation.
	SetServiceTypeGUID(context.Context, *SetServiceTypeGUIDRequest) (*SetServiceTypeGUIDResponse, error)

	// Label operation.
	GetLabel(context.Context, *GetLabelRequest) (*GetLabelResponse, error)

	// Label operation.
	SetLabel(context.Context, *SetLabelRequest) (*SetLabelResponse, error)

	// PathName operation.
	GetPathName(context.Context, *GetPathNameRequest) (*GetPathNameResponse, error)

	// PathName operation.
	SetPathName(context.Context, *SetPathNameRequest) (*SetPathNameResponse, error)

	// FormatName operation.
	GetFormatName(context.Context, *GetFormatNameRequest) (*GetFormatNameResponse, error)

	// FormatName operation.
	SetFormatName(context.Context, *SetFormatNameRequest) (*SetFormatNameResponse, error)

	// The IsTransactional method is received by the server in an RPC_REQUEST packet. In
	// response, the server returns a value that indicates whether the referenced queue
	// is transactional or nontransactional.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	//
	// When the server processes this call, it MUST abide by the following contract:
	//
	// * If IsRefreshed is FALSE call Refresh ( e4072abd-b433-4b22-a4ed-8f654d182705 ) (section
	// 3.10.4.1.29).
	//
	// * If Refresh returns a value other than S_OK (0x00000000), return the HRESULT returned
	// by Refresh and take no further action.
	//
	// * The Queue ( ../ms-mqdmpr/2e026a09-999e-478f-8c4c-5344b661e579 ).Transactional instance
	// variable is True, and ( IsRefreshed is True or IsQueueCreated is True):
	//
	// * Set the pisTransactional output parameter to MQ_TRANSACTIONAL (0x0001).
	//
	// * Else:
	//
	// * Set the pisTransactional output parameter to MQ_TRANSACTIONAL_NONE (0x0000).
	//
	// * Return S_OK (0x00000000), and take no further action.
	GetIsTransactional(context.Context, *GetIsTransactionalRequest) (*GetIsTransactionalResponse, error)

	// PrivLevel operation.
	GetPrivacyLevel(context.Context, *GetPrivacyLevelRequest) (*GetPrivacyLevelResponse, error)

	// PrivLevel operation.
	SetPrivacyLevel(context.Context, *SetPrivacyLevelRequest) (*SetPrivacyLevelResponse, error)

	// Journal operation.
	GetJournal(context.Context, *GetJournalRequest) (*GetJournalResponse, error)

	// Journal operation.
	SetJournal(context.Context, *SetJournalRequest) (*SetJournalResponse, error)

	// Quota operation.
	GetQuota(context.Context, *GetQuotaRequest) (*GetQuotaResponse, error)

	// Quota operation.
	SetQuota(context.Context, *SetQuotaRequest) (*SetQuotaResponse, error)

	// BasePriority operation.
	GetBasePriority(context.Context, *GetBasePriorityRequest) (*GetBasePriorityResponse, error)

	// BasePriority operation.
	SetBasePriority(context.Context, *SetBasePriorityRequest) (*SetBasePriorityResponse, error)

	// The CreateTime method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns the date and time when the referenced queue was created.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	GetCreateTime(context.Context, *GetCreateTimeRequest) (*GetCreateTimeResponse, error)

	// The ModifyTime method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns the latest date and time when one of the properties of the referenced
	// queue was updated.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	GetModifyTime(context.Context, *GetModifyTimeRequest) (*GetModifyTimeResponse, error)

	// Authenticate operation.
	GetAuthenticate(context.Context, *GetAuthenticateRequest) (*GetAuthenticateResponse, error)

	// Authenticate operation.
	SetAuthenticate(context.Context, *SetAuthenticateRequest) (*SetAuthenticateResponse, error)

	// JournalQuota operation.
	GetJournalQuota(context.Context, *GetJournalQuotaRequest) (*GetJournalQuotaResponse, error)

	// JournalQuota operation.
	SetJournalQuota(context.Context, *SetJournalQuotaRequest) (*SetJournalQuotaResponse, error)

	// The IsWorldReadable method is received by the server in an RPC_REQUEST packet. In
	// response, the server returns a BOOLEAN that indicates whether the referenced queue
	// is accessible to everyone, or only to the owner and the system administrators. This
	// can be computed through the security descriptor in the Queue.Security attribute.
	// The owner is the security principal that has MQSEC_TAKE_QUEUE_OWNERSHIP permissions
	// for the Queue, as defined by the security descriptor in the refQueue.Security attribute.
	//
	// Return Values: The method MUST return S_OK (0x00000000).
	GetIsWorldReadable(context.Context, *GetIsWorldReadableRequest) (*GetIsWorldReadableResponse, error)

	// The Create method is received by the server in an RPC_REQUEST packet. In response,
	// the server creates a new public or private ApplicationQueue.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	Create(context.Context, *CreateRequest) (*CreateResponse, error)

	// The Delete method is received by the server in an RPC_REQUEST packet. In response,
	// the server deletes the referenced queue.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)

	// Open operation.
	Open(context.Context, *OpenRequest) (*OpenResponse, error)

	// The Refresh method is received by the server in an RPC_REQUEST packet. In response,
	// the server refreshes the properties of the MSMQQueueInfo object with the values stored
	// in the directory (for public queues) or in the local QueueManager (for private queues).
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error)

	// The Update method is received by the server in an RPC_REQUEST packet. In response,
	// the server updates the directory or the local QueueManager with the current values
	// of the MSMQQueueInfo object's properties.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)

	// The PathNameDNS method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns the DNS path name that identifies the referenced queue.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetPathNameDNS(context.Context, *GetPathNameDNSRequest) (*GetPathNameDNSResponse, error)

	// Properties operation.
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)

	// Security operation.
	GetSecurity(context.Context, *GetSecurityRequest) (*GetSecurityResponse, error)

	// Security operation.
	SetSecurity(context.Context, *SetSecurityRequest) (*SetSecurityResponse, error)

	// The IsTransactional2 method is received by the server in an RPC_REQUEST packet. In
	// response, the server returns a value that indicates whether the referenced queue
	// is transactional or nontransactional.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetIsTransactional2(context.Context, *GetIsTransactional2Request) (*GetIsTransactional2Response, error)

	// The IsWorldReadable2 method is received by the server in an RPC_REQUEST packet. In
	// response, the server returns a BOOLEAN, which indicates whether the referenced queue
	// is accessible to everyone, or only to the owner and the system administrators.
	//
	// Return Values: The method MUST return S_OK (0x00000000).
	GetIsWorldReadable2(context.Context, *GetIsWorldReadable2Request) (*GetIsWorldReadable2Response, error)

	// MulticastAddress operation.
	GetMulticastAddress(context.Context, *GetMulticastAddressRequest) (*GetMulticastAddressResponse, error)

	// MulticastAddress operation.
	SetMulticastAddress(context.Context, *SetMulticastAddressRequest) (*SetMulticastAddressResponse, error)

	// ADsPath operation.
	GetADsPath(context.Context, *GetADsPathRequest) (*GetADsPathResponse, error)
}

func RegisterQueueInfo3Server(conn dcerpc.Conn, o QueueInfo3Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewQueueInfo3ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(QueueInfo3SyntaxV0_0))...)
}

func NewQueueInfo3ServerHandle(o QueueInfo3Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return QueueInfo3ServerHandle(ctx, o, opNum, r)
	}
}

func QueueInfo3ServerHandle(ctx context.Context, o QueueInfo3Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // QueueGuid
		op := &xxx_GetQueueGUIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetQueueGUIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetQueueGUID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // ServiceTypeGuid
		op := &xxx_GetServiceTypeGUIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetServiceTypeGUIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetServiceTypeGUID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // ServiceTypeGuid
		op := &xxx_SetServiceTypeGUIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetServiceTypeGUIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetServiceTypeGUID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // Label
		op := &xxx_GetLabelOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLabelRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLabel(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // Label
		op := &xxx_SetLabelOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetLabelRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetLabel(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // PathName
		op := &xxx_GetPathNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPathNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPathName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // PathName
		op := &xxx_SetPathNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetPathNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetPathName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // FormatName
		op := &xxx_GetFormatNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFormatNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFormatName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // FormatName
		op := &xxx_SetFormatNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFormatNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFormatName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // IsTransactional
		op := &xxx_GetIsTransactionalOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsTransactionalRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsTransactional(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // PrivLevel
		op := &xxx_GetPrivacyLevelOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPrivacyLevelRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPrivacyLevel(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // PrivLevel
		op := &xxx_SetPrivacyLevelOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetPrivacyLevelRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetPrivacyLevel(ctx, req)
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
	case 21: // Quota
		op := &xxx_GetQuotaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetQuotaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetQuota(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // Quota
		op := &xxx_SetQuotaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetQuotaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetQuota(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // BasePriority
		op := &xxx_GetBasePriorityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetBasePriorityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetBasePriority(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // BasePriority
		op := &xxx_SetBasePriorityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetBasePriorityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetBasePriority(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // CreateTime
		op := &xxx_GetCreateTimeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCreateTimeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCreateTime(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 26: // ModifyTime
		op := &xxx_GetModifyTimeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetModifyTimeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetModifyTime(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // Authenticate
		op := &xxx_GetAuthenticateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAuthenticateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAuthenticate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 28: // Authenticate
		op := &xxx_SetAuthenticateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetAuthenticateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetAuthenticate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 29: // JournalQuota
		op := &xxx_GetJournalQuotaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetJournalQuotaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetJournalQuota(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // JournalQuota
		op := &xxx_SetJournalQuotaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetJournalQuotaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetJournalQuota(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 31: // IsWorldReadable
		op := &xxx_GetIsWorldReadableOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsWorldReadableRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsWorldReadable(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 32: // Create
		op := &xxx_CreateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Create(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 33: // Delete
		op := &xxx_DeleteOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Delete(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 34: // Open
		op := &xxx_OpenOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Open(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 35: // Refresh
		op := &xxx_RefreshOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RefreshRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Refresh(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 36: // Update
		op := &xxx_UpdateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UpdateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Update(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 37: // PathNameDNS
		op := &xxx_GetPathNameDNSOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPathNameDNSRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPathNameDNS(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 38: // Properties
		op := &xxx_GetPropertiesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPropertiesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetProperties(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 39: // Security
		op := &xxx_GetSecurityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSecurityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSecurity(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 40: // Security
		op := &xxx_SetSecurityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSecurityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSecurity(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 41: // IsTransactional2
		op := &xxx_GetIsTransactional2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsTransactional2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsTransactional2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 42: // IsWorldReadable2
		op := &xxx_GetIsWorldReadable2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsWorldReadable2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsWorldReadable2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 43: // MulticastAddress
		op := &xxx_GetMulticastAddressOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMulticastAddressRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMulticastAddress(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 44: // MulticastAddress
		op := &xxx_SetMulticastAddressOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMulticastAddressRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMulticastAddress(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 45: // ADsPath
		op := &xxx_GetADsPathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetADsPathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetADsPath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMSMQQueueInfo3
type UnimplementedQueueInfo3Server struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedQueueInfo3Server) GetQueueGUID(context.Context, *GetQueueGUIDRequest) (*GetQueueGUIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) GetServiceTypeGUID(context.Context, *GetServiceTypeGUIDRequest) (*GetServiceTypeGUIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) SetServiceTypeGUID(context.Context, *SetServiceTypeGUIDRequest) (*SetServiceTypeGUIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) GetLabel(context.Context, *GetLabelRequest) (*GetLabelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) SetLabel(context.Context, *SetLabelRequest) (*SetLabelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) GetPathName(context.Context, *GetPathNameRequest) (*GetPathNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) SetPathName(context.Context, *SetPathNameRequest) (*SetPathNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) GetFormatName(context.Context, *GetFormatNameRequest) (*GetFormatNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) SetFormatName(context.Context, *SetFormatNameRequest) (*SetFormatNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) GetIsTransactional(context.Context, *GetIsTransactionalRequest) (*GetIsTransactionalResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) GetPrivacyLevel(context.Context, *GetPrivacyLevelRequest) (*GetPrivacyLevelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) SetPrivacyLevel(context.Context, *SetPrivacyLevelRequest) (*SetPrivacyLevelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) GetJournal(context.Context, *GetJournalRequest) (*GetJournalResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) SetJournal(context.Context, *SetJournalRequest) (*SetJournalResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) GetQuota(context.Context, *GetQuotaRequest) (*GetQuotaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) SetQuota(context.Context, *SetQuotaRequest) (*SetQuotaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) GetBasePriority(context.Context, *GetBasePriorityRequest) (*GetBasePriorityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) SetBasePriority(context.Context, *SetBasePriorityRequest) (*SetBasePriorityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) GetCreateTime(context.Context, *GetCreateTimeRequest) (*GetCreateTimeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) GetModifyTime(context.Context, *GetModifyTimeRequest) (*GetModifyTimeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) GetAuthenticate(context.Context, *GetAuthenticateRequest) (*GetAuthenticateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) SetAuthenticate(context.Context, *SetAuthenticateRequest) (*SetAuthenticateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) GetJournalQuota(context.Context, *GetJournalQuotaRequest) (*GetJournalQuotaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) SetJournalQuota(context.Context, *SetJournalQuotaRequest) (*SetJournalQuotaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) GetIsWorldReadable(context.Context, *GetIsWorldReadableRequest) (*GetIsWorldReadableResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) Open(context.Context, *OpenRequest) (*OpenResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) GetPathNameDNS(context.Context, *GetPathNameDNSRequest) (*GetPathNameDNSResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) GetSecurity(context.Context, *GetSecurityRequest) (*GetSecurityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) SetSecurity(context.Context, *SetSecurityRequest) (*SetSecurityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) GetIsTransactional2(context.Context, *GetIsTransactional2Request) (*GetIsTransactional2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) GetIsWorldReadable2(context.Context, *GetIsWorldReadable2Request) (*GetIsWorldReadable2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) GetMulticastAddress(context.Context, *GetMulticastAddressRequest) (*GetMulticastAddressResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) SetMulticastAddress(context.Context, *SetMulticastAddressRequest) (*SetMulticastAddressResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueInfo3Server) GetADsPath(context.Context, *GetADsPathRequest) (*GetADsPathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ QueueInfo3Server = (*UnimplementedQueueInfo3Server)(nil)
