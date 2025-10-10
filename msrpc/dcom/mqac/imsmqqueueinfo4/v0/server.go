package imsmqqueueinfo4

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

// IMSMQQueueInfo4 server interface.
type ImsmqQueueInfo4Server interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// QueueGuid operation.
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

	// IsTransactional operation.
	GetIsTransactional(context.Context, *GetIsTransactionalRequest) (*GetIsTransactionalResponse, error)

	// PrivLevel operation.
	GetPrivLevel(context.Context, *GetPrivLevelRequest) (*GetPrivLevelResponse, error)

	// PrivLevel operation.
	SetPrivLevel(context.Context, *SetPrivLevelRequest) (*SetPrivLevelResponse, error)

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

	// CreateTime operation.
	GetCreateTime(context.Context, *GetCreateTimeRequest) (*GetCreateTimeResponse, error)

	// ModifyTime operation.
	GetModifyTime(context.Context, *GetModifyTimeRequest) (*GetModifyTimeResponse, error)

	// Authenticate operation.
	GetAuthenticate(context.Context, *GetAuthenticateRequest) (*GetAuthenticateResponse, error)

	// Authenticate operation.
	SetAuthenticate(context.Context, *SetAuthenticateRequest) (*SetAuthenticateResponse, error)

	// JournalQuota operation.
	GetJournalQuota(context.Context, *GetJournalQuotaRequest) (*GetJournalQuotaResponse, error)

	// JournalQuota operation.
	SetJournalQuota(context.Context, *SetJournalQuotaRequest) (*SetJournalQuotaResponse, error)

	// IsWorldReadable operation.
	GetIsWorldReadable(context.Context, *GetIsWorldReadableRequest) (*GetIsWorldReadableResponse, error)

	// Create operation.
	Create(context.Context, *CreateRequest) (*CreateResponse, error)

	// Delete operation.
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)

	// Open operation.
	Open(context.Context, *OpenRequest) (*OpenResponse, error)

	// Refresh operation.
	Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error)

	// Update operation.
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)

	// PathNameDNS operation.
	GetPathNameDNS(context.Context, *GetPathNameDNSRequest) (*GetPathNameDNSResponse, error)

	// Properties operation.
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)

	// Security operation.
	GetSecurity(context.Context, *GetSecurityRequest) (*GetSecurityResponse, error)

	// Security operation.
	SetSecurity(context.Context, *SetSecurityRequest) (*SetSecurityResponse, error)

	// IsTransactional2 operation.
	GetIsTransactional2(context.Context, *GetIsTransactional2Request) (*GetIsTransactional2Response, error)

	// IsWorldReadable2 operation.
	GetIsWorldReadable2(context.Context, *GetIsWorldReadable2Request) (*GetIsWorldReadable2Response, error)

	// MulticastAddress operation.
	GetMulticastAddress(context.Context, *GetMulticastAddressRequest) (*GetMulticastAddressResponse, error)

	// MulticastAddress operation.
	SetMulticastAddress(context.Context, *SetMulticastAddressRequest) (*SetMulticastAddressResponse, error)

	// ADsPath operation.
	GetADSPath(context.Context, *GetADSPathRequest) (*GetADSPathResponse, error)
}

func RegisterImsmqQueueInfo4Server(conn dcerpc.Conn, o ImsmqQueueInfo4Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewImsmqQueueInfo4ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ImsmqQueueInfo4SyntaxV0_0))...)
}

func NewImsmqQueueInfo4ServerHandle(o ImsmqQueueInfo4Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ImsmqQueueInfo4ServerHandle(ctx, o, opNum, r)
	}
}

func ImsmqQueueInfo4ServerHandle(ctx context.Context, o ImsmqQueueInfo4Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
		op := &xxx_GetPrivLevelOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPrivLevelRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPrivLevel(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // PrivLevel
		op := &xxx_SetPrivLevelOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetPrivLevelRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetPrivLevel(ctx, req)
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
		op := &xxx_GetADSPathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetADSPathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetADSPath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMSMQQueueInfo4
type UnimplementedImsmqQueueInfo4Server struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedImsmqQueueInfo4Server) GetQueueGUID(context.Context, *GetQueueGUIDRequest) (*GetQueueGUIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) GetServiceTypeGUID(context.Context, *GetServiceTypeGUIDRequest) (*GetServiceTypeGUIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) SetServiceTypeGUID(context.Context, *SetServiceTypeGUIDRequest) (*SetServiceTypeGUIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) GetLabel(context.Context, *GetLabelRequest) (*GetLabelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) SetLabel(context.Context, *SetLabelRequest) (*SetLabelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) GetPathName(context.Context, *GetPathNameRequest) (*GetPathNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) SetPathName(context.Context, *SetPathNameRequest) (*SetPathNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) GetFormatName(context.Context, *GetFormatNameRequest) (*GetFormatNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) SetFormatName(context.Context, *SetFormatNameRequest) (*SetFormatNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) GetIsTransactional(context.Context, *GetIsTransactionalRequest) (*GetIsTransactionalResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) GetPrivLevel(context.Context, *GetPrivLevelRequest) (*GetPrivLevelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) SetPrivLevel(context.Context, *SetPrivLevelRequest) (*SetPrivLevelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) GetJournal(context.Context, *GetJournalRequest) (*GetJournalResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) SetJournal(context.Context, *SetJournalRequest) (*SetJournalResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) GetQuota(context.Context, *GetQuotaRequest) (*GetQuotaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) SetQuota(context.Context, *SetQuotaRequest) (*SetQuotaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) GetBasePriority(context.Context, *GetBasePriorityRequest) (*GetBasePriorityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) SetBasePriority(context.Context, *SetBasePriorityRequest) (*SetBasePriorityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) GetCreateTime(context.Context, *GetCreateTimeRequest) (*GetCreateTimeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) GetModifyTime(context.Context, *GetModifyTimeRequest) (*GetModifyTimeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) GetAuthenticate(context.Context, *GetAuthenticateRequest) (*GetAuthenticateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) SetAuthenticate(context.Context, *SetAuthenticateRequest) (*SetAuthenticateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) GetJournalQuota(context.Context, *GetJournalQuotaRequest) (*GetJournalQuotaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) SetJournalQuota(context.Context, *SetJournalQuotaRequest) (*SetJournalQuotaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) GetIsWorldReadable(context.Context, *GetIsWorldReadableRequest) (*GetIsWorldReadableResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) Open(context.Context, *OpenRequest) (*OpenResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) GetPathNameDNS(context.Context, *GetPathNameDNSRequest) (*GetPathNameDNSResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) GetSecurity(context.Context, *GetSecurityRequest) (*GetSecurityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) SetSecurity(context.Context, *SetSecurityRequest) (*SetSecurityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) GetIsTransactional2(context.Context, *GetIsTransactional2Request) (*GetIsTransactional2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) GetIsWorldReadable2(context.Context, *GetIsWorldReadable2Request) (*GetIsWorldReadable2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) GetMulticastAddress(context.Context, *GetMulticastAddressRequest) (*GetMulticastAddressResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) SetMulticastAddress(context.Context, *SetMulticastAddressRequest) (*SetMulticastAddressResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfo4Server) GetADSPath(context.Context, *GetADSPathRequest) (*GetADSPathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ImsmqQueueInfo4Server = (*UnimplementedImsmqQueueInfo4Server)(nil)
