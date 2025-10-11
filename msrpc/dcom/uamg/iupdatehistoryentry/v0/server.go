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

	// The IUpdateHistoryEntry::Operation (opnum 8) method retrieves the operation performed.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the Operation ADM element.
	GetOperation(context.Context, *GetOperationRequest) (*GetOperationResponse, error)

	// The ISearchResult::ResultCode (opnum 8) method retrieves the result code for the
	// search.
	//
	// The IUpdateHistoryEntry::ResultCode (opnum 9) method describes the result of the
	// operation.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the ResultCode ADM element.
	GetResultCode(context.Context, *GetResultCodeRequest) (*GetResultCodeResponse, error)

	// The IUpdateException::HResult (opnum 9) method retrieves the HRESULT describing the
	// error.
	//
	// The IUpdateHistoryEntry::HResult (opnum 10) method retrieves the HRESULT from the
	// operation.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the HResult ADM element.
	GetHResult(context.Context, *GetHResultRequest) (*GetHResultResponse, error)

	// The IUpdateHistoryEntry::Date (opnum 11) method retrieves the date-time that the
	// operation was performed.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the result of the Date ADM element.
	GetDate(context.Context, *GetDateRequest) (*GetDateResponse, error)

	// The IUpdateHistoryEntry::UpdateIdentity (opnum 12) method retrieves the identity
	// of the update for which the operation was performed.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the result of the UpdateIdentity ADM element.
	GetUpdateIdentity(context.Context, *GetUpdateIdentityRequest) (*GetUpdateIdentityResponse, error)

	// The IUpdate::Title (opnum 8) method retrieves the localized title of the update.
	//
	// The IUpdateHistoryEntry::Title (opnum 13) method retrieves the title of the update
	// for which the operation was performed.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The server SHOULD return the value of the Title ADM element.
	GetTitle(context.Context, *GetTitleRequest) (*GetTitleResponse, error)

	// The ICategory::Description (opnum 11) method retrieves the description of the update
	// category.
	//
	// The IUpdateHistoryEntry::Description (opnum 14) method retrieves the description
	// of the update for which the operation is performed.
	//
	// The IUpdate::Description (opnum 16) method retrieves the localized description for
	// the update.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the Description ADM element.
	GetDescription(context.Context, *GetDescriptionRequest) (*GetDescriptionResponse, error)

	// The IUpdateHistoryEntry::UnmappedResultCode (opnum 15) method retrieves an unmapped
	// result code from the operation. If the operation produced a result code that is not
	// an HRESULT, the unmapped result code SHOULD be set to that original result code.
	// Otherwise, the unmapped result code SHOULD be 0x00000000.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the UnmappedResultCode ADM element.
	GetUnmappedResultCode(context.Context, *GetUnmappedResultCodeRequest) (*GetUnmappedResultCodeResponse, error)

	// The IUpdateSearcher::ClientApplicationID (opnum 10) method retrieves the string used
	// to identify the current client application.
	//
	// The IUpdateSession::ClientApplicationID (opnum 9) method sets the identifier of the
	// calling application.
	//
	// The IUpdateHistoryEntry::ClientApplicationID (opnum 16) method retrieves the ID of
	// the application that initiated the operation.
	//
	// The IUpdateServiceManager2::ClientApplicationID (opnum 16) method sets a string that
	// identifies the client application that is using this interface.
	//
	// The IUpdateSession::ClientApplicationID (opnum 8) method retrieves the identifier
	// of the calling application.
	//
	// The IUpdateSearcher::ClientApplicationID (opnum 11) method sets the string used to
	// identify the current client application.
	//
	// The IUpdateServiceManager2::ClientApplicationID (opnum 15) method retrieves a string
	// that identifies the client application that is using this interface.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the ClientApplicationID ADM element.
	GetClientApplicationID(context.Context, *GetClientApplicationIDRequest) (*GetClientApplicationIDResponse, error)

	// The ServerSelection enumeration defines values that describe the type of server to
	// use for an update search operation.
	//
	// The IUpdateSearcher::ServerSelection (opnum 14) method retrieves the type of update
	// server used to search against.
	//
	// The IUpdateSearcher::ServerSelection (Opnum 15) method sets the type of update server
	// used to search against.
	//
	// The IUpdateHistoryEntry::ServerSelection (opnum 17) method describes the type of
	// update service that provided the update for which the operation was performed.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the ServerSelection ADM element.
	GetServerSelection(context.Context, *GetServerSelectionRequest) (*GetServerSelectionResponse, error)

	// The IUpdateService::ServiceID (opnum 16) method retrieves the unique identifier for
	// the update service.
	//
	// The IUpdateHistoryEntry::ServiceID (opnum 18) method retrieves the unique identifier
	// of the update service that provided the update for which the operation was performed.
	//
	// The IUpdateSearcher::ServiceID (opnum 24) method retrieves the unique identifier
	// of the update server used to search against.
	//
	// The IUpdateSearcher::ServiceID (opnum 25) method sets the unique identifier of the
	// update server used to search against.
	//
	// The IUpdateServiceRegistration::ServiceID (opnum 9) method retrieves the service
	// identifier.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the ServiceID ADM element.
	GetServiceID(context.Context, *GetServiceIDRequest) (*GetServiceIDResponse, error)

	// The IUpdate::UninstallationSteps (opnum 46) method retrieves a list of localized
	// uninstallation steps for the update.
	//
	// The IUpdateHistoryEntry::UninstallationSteps (opnum 19) method retrieves the uninstallation
	// steps for the update for which the operation was performed.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The server SHOULD return the value of the UninstallationSteps ADM element.
	GetUninstallationSteps(context.Context, *GetUninstallationStepsRequest) (*GetUninstallationStepsResponse, error)

	// The IUpdateHistoryEntry::UninstallationNotes (opnum 20) method retrieves the uninstallation
	// notes for the update for which the operation was performed.
	//
	// The IUpdate::UninstallationNotes (opnum 44) method retrieves the localized uninstallation
	// notes for the update.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the result of the UninstallationNotes ADM element.
	GetUninstallationNotes(context.Context, *GetUninstallationNotesRequest) (*GetUninstallationNotesResponse, error)

	// The IUpdate::SupportUrl (opnum 42) method retrieves a URL for a document containing
	// support information for the update.
	//
	// The IUpdateHistoryEntry::SupportUrl (opnum 21) method retrieves the support URL for
	// the update for which the operation was performed.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The server SHOULD return the value of the SupportUrl ADM element.
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
