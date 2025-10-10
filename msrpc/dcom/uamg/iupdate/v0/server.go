package iupdate

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

// IUpdate server interface.
type UpdateServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	GetTitle(context.Context, *GetTitleRequest) (*GetTitleResponse, error)

	GetAutoSelectOnWebSites(context.Context, *GetAutoSelectOnWebSitesRequest) (*GetAutoSelectOnWebSitesResponse, error)

	GetBundledUpdates(context.Context, *GetBundledUpdatesRequest) (*GetBundledUpdatesResponse, error)

	GetCanRequireSource(context.Context, *GetCanRequireSourceRequest) (*GetCanRequireSourceResponse, error)

	GetCategories(context.Context, *GetCategoriesRequest) (*GetCategoriesResponse, error)

	GetDeadline(context.Context, *GetDeadlineRequest) (*GetDeadlineResponse, error)

	GetDeltaCompressedContentAvailable(context.Context, *GetDeltaCompressedContentAvailableRequest) (*GetDeltaCompressedContentAvailableResponse, error)

	GetDeltaCompressedContentPreferred(context.Context, *GetDeltaCompressedContentPreferredRequest) (*GetDeltaCompressedContentPreferredResponse, error)

	GetDescription(context.Context, *GetDescriptionRequest) (*GetDescriptionResponse, error)

	GetEulaAccepted(context.Context, *GetEulaAcceptedRequest) (*GetEulaAcceptedResponse, error)

	GetEulaText(context.Context, *GetEulaTextRequest) (*GetEulaTextResponse, error)

	GetHandlerID(context.Context, *GetHandlerIDRequest) (*GetHandlerIDResponse, error)

	GetIdentity(context.Context, *GetIdentityRequest) (*GetIdentityResponse, error)

	GetImage(context.Context, *GetImageRequest) (*GetImageResponse, error)

	GetInstallationBehavior(context.Context, *GetInstallationBehaviorRequest) (*GetInstallationBehaviorResponse, error)

	GetIsBeta(context.Context, *GetIsBetaRequest) (*GetIsBetaResponse, error)

	GetIsDownloaded(context.Context, *GetIsDownloadedRequest) (*GetIsDownloadedResponse, error)

	GetIsHidden(context.Context, *GetIsHiddenRequest) (*GetIsHiddenResponse, error)

	// Opnum26NotUsedOnWire operation.
	// Opnum26NotUsedOnWire

	GetIsInstalled(context.Context, *GetIsInstalledRequest) (*GetIsInstalledResponse, error)

	GetIsMandatory(context.Context, *GetIsMandatoryRequest) (*GetIsMandatoryResponse, error)

	GetIsUninstallable(context.Context, *GetIsUninstallableRequest) (*GetIsUninstallableResponse, error)

	GetLanguages(context.Context, *GetLanguagesRequest) (*GetLanguagesResponse, error)

	GetLastDeploymentChangeTime(context.Context, *GetLastDeploymentChangeTimeRequest) (*GetLastDeploymentChangeTimeResponse, error)

	GetMaxDownloadSize(context.Context, *GetMaxDownloadSizeRequest) (*GetMaxDownloadSizeResponse, error)

	GetMinDownloadSize(context.Context, *GetMinDownloadSizeRequest) (*GetMinDownloadSizeResponse, error)

	GetMoreInfoUrls(context.Context, *GetMoreInfoUrlsRequest) (*GetMoreInfoUrlsResponse, error)

	GetMsrcSeverity(context.Context, *GetMsrcSeverityRequest) (*GetMsrcSeverityResponse, error)

	GetRecommendedCpuSpeed(context.Context, *GetRecommendedCpuSpeedRequest) (*GetRecommendedCpuSpeedResponse, error)

	GetRecommendedHardDiskSpace(context.Context, *GetRecommendedHardDiskSpaceRequest) (*GetRecommendedHardDiskSpaceResponse, error)

	GetRecommendedMemory(context.Context, *GetRecommendedMemoryRequest) (*GetRecommendedMemoryResponse, error)

	GetReleaseNotes(context.Context, *GetReleaseNotesRequest) (*GetReleaseNotesResponse, error)

	GetSecurityBulletinIDs(context.Context, *GetSecurityBulletinIDsRequest) (*GetSecurityBulletinIDsResponse, error)

	GetSupersededUpdateIDs(context.Context, *GetSupersededUpdateIDsRequest) (*GetSupersededUpdateIDsResponse, error)

	GetSupportURL(context.Context, *GetSupportURLRequest) (*GetSupportURLResponse, error)

	GetType(context.Context, *GetTypeRequest) (*GetTypeResponse, error)

	GetUninstallationNotes(context.Context, *GetUninstallationNotesRequest) (*GetUninstallationNotesResponse, error)

	GetUninstallationBehavior(context.Context, *GetUninstallationBehaviorRequest) (*GetUninstallationBehaviorResponse, error)

	GetUninstallationSteps(context.Context, *GetUninstallationStepsRequest) (*GetUninstallationStepsResponse, error)

	GetKbArticleIDs(context.Context, *GetKbArticleIDsRequest) (*GetKbArticleIDsResponse, error)

	// Opnum48NotUsedOnWire operation.
	// Opnum48NotUsedOnWire

	GetDeploymentAction(context.Context, *GetDeploymentActionRequest) (*GetDeploymentActionResponse, error)

	// Opnum50NotUsedOnWire operation.
	// Opnum50NotUsedOnWire

	GetDownloadPriority(context.Context, *GetDownloadPriorityRequest) (*GetDownloadPriorityResponse, error)

	GetDownloadContents(context.Context, *GetDownloadContentsRequest) (*GetDownloadContentsResponse, error)
}

func RegisterUpdateServer(conn dcerpc.Conn, o UpdateServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewUpdateServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(UpdateSyntaxV0_0))...)
}

func NewUpdateServerHandle(o UpdateServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return UpdateServerHandle(ctx, o, opNum, r)
	}
}

func UpdateServerHandle(ctx context.Context, o UpdateServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Title
		op := &xxx_GetTitleOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTitleRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTitle(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // AutoSelectOnWebSites
		op := &xxx_GetAutoSelectOnWebSitesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAutoSelectOnWebSitesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAutoSelectOnWebSites(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // BundledUpdates
		op := &xxx_GetBundledUpdatesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetBundledUpdatesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetBundledUpdates(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // CanRequireSource
		op := &xxx_GetCanRequireSourceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCanRequireSourceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCanRequireSource(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // Categories
		op := &xxx_GetCategoriesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCategoriesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCategories(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // Deadline
		op := &xxx_GetDeadlineOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDeadlineRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDeadline(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // DeltaCompressedContentAvailable
		op := &xxx_GetDeltaCompressedContentAvailableOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDeltaCompressedContentAvailableRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDeltaCompressedContentAvailable(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // DeltaCompressedContentPreferred
		op := &xxx_GetDeltaCompressedContentPreferredOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDeltaCompressedContentPreferredRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDeltaCompressedContentPreferred(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // Description
		op := &xxx_GetDescriptionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDescriptionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDescription(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // EulaAccepted
		op := &xxx_GetEulaAcceptedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetEulaAcceptedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetEulaAccepted(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // EulaText
		op := &xxx_GetEulaTextOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetEulaTextRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetEulaText(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // HandlerID
		op := &xxx_GetHandlerIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetHandlerIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetHandlerID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // Identity
		op := &xxx_GetIdentityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIdentityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIdentity(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // Image
		op := &xxx_GetImageOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetImageRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetImage(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // InstallationBehavior
		op := &xxx_GetInstallationBehaviorOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetInstallationBehaviorRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetInstallationBehavior(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // IsBeta
		op := &xxx_GetIsBetaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsBetaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsBeta(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // IsDownloaded
		op := &xxx_GetIsDownloadedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsDownloadedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsDownloaded(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // IsHidden
		op := &xxx_GetIsHiddenOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsHiddenRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsHidden(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // Opnum26NotUsedOnWire
		// Opnum26NotUsedOnWire
		return nil, nil
	case 26: // IsInstalled
		op := &xxx_GetIsInstalledOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsInstalledRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsInstalled(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // IsMandatory
		op := &xxx_GetIsMandatoryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsMandatoryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsMandatory(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 28: // IsUninstallable
		op := &xxx_GetIsUninstallableOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsUninstallableRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsUninstallable(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 29: // Languages
		op := &xxx_GetLanguagesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLanguagesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLanguages(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // LastDeploymentChangeTime
		op := &xxx_GetLastDeploymentChangeTimeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLastDeploymentChangeTimeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLastDeploymentChangeTime(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 31: // MaxDownloadSize
		op := &xxx_GetMaxDownloadSizeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMaxDownloadSizeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMaxDownloadSize(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 32: // MinDownloadSize
		op := &xxx_GetMinDownloadSizeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMinDownloadSizeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMinDownloadSize(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 33: // MoreInfoUrls
		op := &xxx_GetMoreInfoUrlsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMoreInfoUrlsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMoreInfoUrls(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 34: // MsrcSeverity
		op := &xxx_GetMsrcSeverityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMsrcSeverityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMsrcSeverity(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 35: // RecommendedCpuSpeed
		op := &xxx_GetRecommendedCpuSpeedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRecommendedCpuSpeedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRecommendedCpuSpeed(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 36: // RecommendedHardDiskSpace
		op := &xxx_GetRecommendedHardDiskSpaceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRecommendedHardDiskSpaceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRecommendedHardDiskSpace(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 37: // RecommendedMemory
		op := &xxx_GetRecommendedMemoryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRecommendedMemoryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRecommendedMemory(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 38: // ReleaseNotes
		op := &xxx_GetReleaseNotesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetReleaseNotesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetReleaseNotes(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 39: // SecurityBulletinIDs
		op := &xxx_GetSecurityBulletinIDsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSecurityBulletinIDsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSecurityBulletinIDs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 40: // SupersededUpdateIDs
		op := &xxx_GetSupersededUpdateIDsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSupersededUpdateIDsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSupersededUpdateIDs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 41: // SupportUrl
		op := &xxx_GetSupportURLOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSupportURLRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSupportURL(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 42: // Type
		op := &xxx_GetTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 43: // UninstallationNotes
		op := &xxx_GetUninstallationNotesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetUninstallationNotesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetUninstallationNotes(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 44: // UninstallationBehavior
		op := &xxx_GetUninstallationBehaviorOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetUninstallationBehaviorRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetUninstallationBehavior(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 45: // UninstallationSteps
		op := &xxx_GetUninstallationStepsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetUninstallationStepsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetUninstallationSteps(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 46: // KBArticleIDs
		op := &xxx_GetKbArticleIDsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetKbArticleIDsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetKbArticleIDs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 47: // Opnum48NotUsedOnWire
		// Opnum48NotUsedOnWire
		return nil, nil
	case 48: // DeploymentAction
		op := &xxx_GetDeploymentActionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDeploymentActionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDeploymentAction(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 49: // Opnum50NotUsedOnWire
		// Opnum50NotUsedOnWire
		return nil, nil
	case 50: // DownloadPriority
		op := &xxx_GetDownloadPriorityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDownloadPriorityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDownloadPriority(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 51: // DownloadContents
		op := &xxx_GetDownloadContentsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDownloadContentsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDownloadContents(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IUpdate
type UnimplementedUpdateServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedUpdateServer) GetTitle(context.Context, *GetTitleRequest) (*GetTitleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetAutoSelectOnWebSites(context.Context, *GetAutoSelectOnWebSitesRequest) (*GetAutoSelectOnWebSitesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetBundledUpdates(context.Context, *GetBundledUpdatesRequest) (*GetBundledUpdatesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetCanRequireSource(context.Context, *GetCanRequireSourceRequest) (*GetCanRequireSourceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetCategories(context.Context, *GetCategoriesRequest) (*GetCategoriesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetDeadline(context.Context, *GetDeadlineRequest) (*GetDeadlineResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetDeltaCompressedContentAvailable(context.Context, *GetDeltaCompressedContentAvailableRequest) (*GetDeltaCompressedContentAvailableResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetDeltaCompressedContentPreferred(context.Context, *GetDeltaCompressedContentPreferredRequest) (*GetDeltaCompressedContentPreferredResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetDescription(context.Context, *GetDescriptionRequest) (*GetDescriptionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetEulaAccepted(context.Context, *GetEulaAcceptedRequest) (*GetEulaAcceptedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetEulaText(context.Context, *GetEulaTextRequest) (*GetEulaTextResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetHandlerID(context.Context, *GetHandlerIDRequest) (*GetHandlerIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetIdentity(context.Context, *GetIdentityRequest) (*GetIdentityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetImage(context.Context, *GetImageRequest) (*GetImageResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetInstallationBehavior(context.Context, *GetInstallationBehaviorRequest) (*GetInstallationBehaviorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetIsBeta(context.Context, *GetIsBetaRequest) (*GetIsBetaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetIsDownloaded(context.Context, *GetIsDownloadedRequest) (*GetIsDownloadedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetIsHidden(context.Context, *GetIsHiddenRequest) (*GetIsHiddenResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetIsInstalled(context.Context, *GetIsInstalledRequest) (*GetIsInstalledResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetIsMandatory(context.Context, *GetIsMandatoryRequest) (*GetIsMandatoryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetIsUninstallable(context.Context, *GetIsUninstallableRequest) (*GetIsUninstallableResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetLanguages(context.Context, *GetLanguagesRequest) (*GetLanguagesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetLastDeploymentChangeTime(context.Context, *GetLastDeploymentChangeTimeRequest) (*GetLastDeploymentChangeTimeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetMaxDownloadSize(context.Context, *GetMaxDownloadSizeRequest) (*GetMaxDownloadSizeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetMinDownloadSize(context.Context, *GetMinDownloadSizeRequest) (*GetMinDownloadSizeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetMoreInfoUrls(context.Context, *GetMoreInfoUrlsRequest) (*GetMoreInfoUrlsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetMsrcSeverity(context.Context, *GetMsrcSeverityRequest) (*GetMsrcSeverityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetRecommendedCpuSpeed(context.Context, *GetRecommendedCpuSpeedRequest) (*GetRecommendedCpuSpeedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetRecommendedHardDiskSpace(context.Context, *GetRecommendedHardDiskSpaceRequest) (*GetRecommendedHardDiskSpaceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetRecommendedMemory(context.Context, *GetRecommendedMemoryRequest) (*GetRecommendedMemoryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetReleaseNotes(context.Context, *GetReleaseNotesRequest) (*GetReleaseNotesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetSecurityBulletinIDs(context.Context, *GetSecurityBulletinIDsRequest) (*GetSecurityBulletinIDsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetSupersededUpdateIDs(context.Context, *GetSupersededUpdateIDsRequest) (*GetSupersededUpdateIDsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetSupportURL(context.Context, *GetSupportURLRequest) (*GetSupportURLResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetType(context.Context, *GetTypeRequest) (*GetTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetUninstallationNotes(context.Context, *GetUninstallationNotesRequest) (*GetUninstallationNotesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetUninstallationBehavior(context.Context, *GetUninstallationBehaviorRequest) (*GetUninstallationBehaviorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetUninstallationSteps(context.Context, *GetUninstallationStepsRequest) (*GetUninstallationStepsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetKbArticleIDs(context.Context, *GetKbArticleIDsRequest) (*GetKbArticleIDsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetDeploymentAction(context.Context, *GetDeploymentActionRequest) (*GetDeploymentActionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetDownloadPriority(context.Context, *GetDownloadPriorityRequest) (*GetDownloadPriorityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServer) GetDownloadContents(context.Context, *GetDownloadContentsRequest) (*GetDownloadContentsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ UpdateServer = (*UnimplementedUpdateServer)(nil)
