package iupdatehistoryentry

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

// IUpdateHistoryEntry server interface.
type UpdateHistoryEntryServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	GetOperation(context.Context, *GetOperationRequest) (*GetOperationResponse, error)

	GetResultCode(context.Context, *GetResultCodeRequest) (*GetResultCodeResponse, error)

	GetHResult(context.Context, *GetHResultRequest) (*GetHResultResponse, error)

	GetDate(context.Context, *GetDateRequest) (*GetDateResponse, error)

	GetUpdateIdentity(context.Context, *GetUpdateIdentityRequest) (*GetUpdateIdentityResponse, error)

	GetTitle(context.Context, *GetTitleRequest) (*GetTitleResponse, error)

	GetDescription(context.Context, *GetDescriptionRequest) (*GetDescriptionResponse, error)

	GetUnmappedResultCode(context.Context, *GetUnmappedResultCodeRequest) (*GetUnmappedResultCodeResponse, error)

	GetClientApplicationID(context.Context, *GetClientApplicationIDRequest) (*GetClientApplicationIDResponse, error)

	GetServerSelection(context.Context, *GetServerSelectionRequest) (*GetServerSelectionResponse, error)

	GetServiceID(context.Context, *GetServiceIDRequest) (*GetServiceIDResponse, error)

	GetUninstallationSteps(context.Context, *GetUninstallationStepsRequest) (*GetUninstallationStepsResponse, error)

	GetUninstallationNotes(context.Context, *GetUninstallationNotesRequest) (*GetUninstallationNotesResponse, error)

	GetSupportURL(context.Context, *GetSupportURLRequest) (*GetSupportURLResponse, error)
}

func RegisterUpdateHistoryEntryServer(conn dcerpc.Conn, o UpdateHistoryEntryServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewUpdateHistoryEntryServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(UpdateHistoryEntrySyntaxV0_0))...)
}

func NewUpdateHistoryEntryServerHandle(o UpdateHistoryEntryServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return UpdateHistoryEntryServerHandle(ctx, o, opNum, r)
	}
}

func UpdateHistoryEntryServerHandle(ctx context.Context, o UpdateHistoryEntryServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Operation
		op := &xxx_GetOperationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetOperationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetOperation(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // ResultCode
		op := &xxx_GetResultCodeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetResultCodeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetResultCode(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // HResult
		op := &xxx_GetHResultOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetHResultRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetHResult(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // Date
		op := &xxx_GetDateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // UpdateIdentity
		op := &xxx_GetUpdateIdentityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetUpdateIdentityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetUpdateIdentity(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // Title
		op := &xxx_GetTitleOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTitleRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTitle(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // Description
		op := &xxx_GetDescriptionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDescriptionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDescription(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // UnmappedResultCode
		op := &xxx_GetUnmappedResultCodeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetUnmappedResultCodeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetUnmappedResultCode(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // ClientApplicationID
		op := &xxx_GetClientApplicationIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetClientApplicationIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetClientApplicationID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // ServerSelection
		op := &xxx_GetServerSelectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetServerSelectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetServerSelection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // ServiceID
		op := &xxx_GetServiceIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetServiceIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetServiceID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // UninstallationSteps
		op := &xxx_GetUninstallationStepsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetUninstallationStepsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetUninstallationSteps(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // UninstallationNotes
		op := &xxx_GetUninstallationNotesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetUninstallationNotesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetUninstallationNotes(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // SupportUrl
		op := &xxx_GetSupportURLOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSupportURLRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSupportURL(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IUpdateHistoryEntry
type UnimplementedUpdateHistoryEntryServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedUpdateHistoryEntryServer) GetOperation(context.Context, *GetOperationRequest) (*GetOperationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateHistoryEntryServer) GetResultCode(context.Context, *GetResultCodeRequest) (*GetResultCodeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateHistoryEntryServer) GetHResult(context.Context, *GetHResultRequest) (*GetHResultResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateHistoryEntryServer) GetDate(context.Context, *GetDateRequest) (*GetDateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateHistoryEntryServer) GetUpdateIdentity(context.Context, *GetUpdateIdentityRequest) (*GetUpdateIdentityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateHistoryEntryServer) GetTitle(context.Context, *GetTitleRequest) (*GetTitleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateHistoryEntryServer) GetDescription(context.Context, *GetDescriptionRequest) (*GetDescriptionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateHistoryEntryServer) GetUnmappedResultCode(context.Context, *GetUnmappedResultCodeRequest) (*GetUnmappedResultCodeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateHistoryEntryServer) GetClientApplicationID(context.Context, *GetClientApplicationIDRequest) (*GetClientApplicationIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateHistoryEntryServer) GetServerSelection(context.Context, *GetServerSelectionRequest) (*GetServerSelectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateHistoryEntryServer) GetServiceID(context.Context, *GetServiceIDRequest) (*GetServiceIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateHistoryEntryServer) GetUninstallationSteps(context.Context, *GetUninstallationStepsRequest) (*GetUninstallationStepsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateHistoryEntryServer) GetUninstallationNotes(context.Context, *GetUninstallationNotesRequest) (*GetUninstallationNotesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateHistoryEntryServer) GetSupportURL(context.Context, *GetSupportURLRequest) (*GetSupportURLResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ UpdateHistoryEntryServer = (*UnimplementedUpdateHistoryEntryServer)(nil)
