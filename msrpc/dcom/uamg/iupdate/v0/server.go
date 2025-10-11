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

	// The IUpdate::AutoSelectOnWebSites (opnum 9) method retrieves whether the update is
	// recommended to be selected automatically in website UIs.
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
	// The server SHOULD return the value of the AutoSelectOnWebSites ADM element.
	GetAutoSelectOnWebSites(context.Context, *GetAutoSelectOnWebSitesRequest) (*GetAutoSelectOnWebSitesResponse, error)

	// The IUpdate::BundledUpdates (opnum 10) method retrieves a collection of the updates
	// bundled by the update.
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
	// The server SHOULD return the value of the BundledUpdates ADM element.
	GetBundledUpdates(context.Context, *GetBundledUpdatesRequest) (*GetBundledUpdatesResponse, error)

	// The IUpdate::CanRequireSource (opnum 11) method retrieves whether the installation
	// of the update can require source media.
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
	// The server SHOULD return the value of the CanRequireSource ADM element.
	GetCanRequireSource(context.Context, *GetCanRequireSourceRequest) (*GetCanRequireSourceResponse, error)

	// The IUpdate::Categories (opnum 12) method retrieves a collection of the categories
	// to which the update belongs.
	//
	// The IUpdateHistoryEntry2::Categories (opnum 22) method retrieves a collection of
	// the update categories to which an update belongs.
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
	// The server SHOULD return the value of the Categories ADM element.
	GetCategories(context.Context, *GetCategoriesRequest) (*GetCategoriesResponse, error)

	// The IUpdate::Deadline (opnum 13) method retrieves the date-time before which the
	// update is installed.
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
	// If the HasDeadline ADM element is set to TRUE, the server SHOULD set the vt field
	// of retval to VT_DATE and return the value of the Deadline ADM element in the date
	// field of retval.
	//
	// Otherwise, if the HasDeadline ADM element is set to FALSE, the server SHOULD set
	// the vt field of retval to VT_EMPTY.
	GetDeadline(context.Context, *GetDeadlineRequest) (*GetDeadlineResponse, error)

	// The IUpdate::DeltaCompressedContentAvailable (opnum 14) method retrieves whether
	// delta-compressed content is available for the update.
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
	// The server SHOULD return the value of the DeltaCompressedContentAvailable ADM element.
	GetDeltaCompressedContentAvailable(context.Context, *GetDeltaCompressedContentAvailableRequest) (*GetDeltaCompressedContentAvailableResponse, error)

	// The IUpdate::DeltaCompressedContentPreferred (opnum 15) method retrieves whether
	// delta-compressed content is preferred for the update.
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
	// The server SHOULD return the value of the DeltaCompressedContentPreferred ADM element.
	GetDeltaCompressedContentPreferred(context.Context, *GetDeltaCompressedContentPreferredRequest) (*GetDeltaCompressedContentPreferredResponse, error)

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

	// The  IUpdate::EulaAccepted (opnum 17) method retrieves whether the software license
	// terms associated with the update have been accepted by a user.
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
	// If the value of the HasEula ADM element is TRUE, the server SHOULD return the value
	// of the EulaAccepted ADM element. Otherwise, if the value of the HasEula ADM element
	// is FALSE, the server SHOULD return VARIANT_TRUE.
	GetEulaAccepted(context.Context, *GetEulaAcceptedRequest) (*GetEulaAcceptedResponse, error)

	// The IUpdate::EulaText (opnum 18) method retrieves the text of the software license
	// terms associated with the update.
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
	// If the value of the HasEula ADM element is TRUE, the server SHOULD return the value
	// of the EulaText ADM element. Otherwise, if the value of the HasEula ADM element is
	// FALSE, the server SHOULD return NULL or the empty string.
	GetEulaText(context.Context, *GetEulaTextRequest) (*GetEulaTextResponse, error)

	// The IUpdate::HandlerID (opnum 19) method retrieves the URI identifying the handler
	// for the update.
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
	// The server SHOULD return the value of the HandlerID ADM element.
	GetHandlerID(context.Context, *GetHandlerIDRequest) (*GetHandlerIDResponse, error)

	// The IUpdate::Identity (opnum 20) method retrieves the unique identity of the update.
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
	// The server SHOULD return the value of the Identity ADM element.
	GetIdentity(context.Context, *GetIdentityRequest) (*GetIdentityResponse, error)

	// The ICategory::Image (opnum 12) method retrieves an IImageInformation interface instance
	// that contains information on the image associated with the update category.
	//
	// The IUpdate::Image (opnum 21) method retrieves a localized image associated with
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
	// This method SHOULD return the value of the Image ADM element.
	GetImage(context.Context, *GetImageRequest) (*GetImageResponse, error)

	// The IUpdate::InstallationBehavior (opnum 22) method retrieves a description of the
	// installation behavior of the update.
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
	// The server SHOULD return the value of the InstallationBehavior ADM element.
	GetInstallationBehavior(context.Context, *GetInstallationBehaviorRequest) (*GetInstallationBehaviorResponse, error)

	// The IUpdate::IsBeta (opnum 23) method retrieves whether the update is a beta release.
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
	// The server SHOULD return the value of the IsBeta ADM element.
	GetIsBeta(context.Context, *GetIsBetaRequest) (*GetIsBetaResponse, error)

	// The IUpdate::IsDownloaded (opnum 24) method retrieves whether the update is currently
	// cached by the update agent.
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
	// The server SHOULD return the value of the IsDownloaded ADM element.
	GetIsDownloaded(context.Context, *GetIsDownloadedRequest) (*GetIsDownloadedResponse, error)

	// The IUpdate::IsHidden (opnum 25) method retrieves whether the update has been hidden
	// by a user.
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
	// The server SHOULD return the value of the IsHidden ADM element.
	GetIsHidden(context.Context, *GetIsHiddenRequest) (*GetIsHiddenResponse, error)

	// Opnum26NotUsedOnWire operation.
	// Opnum26NotUsedOnWire

	// The IUpdate::IsInstalled (opnum 27) method retrieves whether the update is installed.
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
	// The server SHOULD return the value of the IsInstalled ADM element.
	GetIsInstalled(context.Context, *GetIsInstalledRequest) (*GetIsInstalledResponse, error)

	// The IUpdate::IsMandatory (opnum 28) method retrieves whether the installation of
	// the update is mandatory.
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
	// The server SHOULD return the value of the IsMandatory ADM element.
	GetIsMandatory(context.Context, *GetIsMandatoryRequest) (*GetIsMandatoryResponse, error)

	// The IUpdate::IsUninstallable (opnum 29) method retrieves whether the update can be
	// uninstalled.
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
	// The server SHOULD return the value of the IsUninstallable ADM element.
	GetIsUninstallable(context.Context, *GetIsUninstallableRequest) (*GetIsUninstallableResponse, error)

	// The IUpdate::Languages (opnum 30) method retrieves a collection of languages supported
	// by the update.
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
	// The server SHOULD return the value of the Languages ADM element.
	GetLanguages(context.Context, *GetLanguagesRequest) (*GetLanguagesResponse, error)

	// The IUpdate::LastDeploymentChangeTime (opnum 31) method retrieves the date-time when
	// the update was last published.
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
	// The server SHOULD return the value of the LastDeploymentChangeTime ADM element.
	GetLastDeploymentChangeTime(context.Context, *GetLastDeploymentChangeTimeRequest) (*GetLastDeploymentChangeTimeResponse, error)

	// The IUpdate::MaxDownloadSize (opnum 32) method retrieves the maximum download size
	// of the update.
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
	// The server SHOULD return the value of the MaxDownloadSize ADM element.
	GetMaxDownloadSize(context.Context, *GetMaxDownloadSizeRequest) (*GetMaxDownloadSizeResponse, error)

	// The IUpdate::MinDownloadSize (opnum 33) method retrieves the minimum download size
	// of the update.
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
	// The server SHOULD return the value of the MinDownloadSize ADM element.
	GetMinDownloadSize(context.Context, *GetMinDownloadSizeRequest) (*GetMinDownloadSizeResponse, error)

	// The IUpdate::MoreInfoUrls (opnum 34) method retrieves a collection of URLs of documents
	// that contain more information about the update.
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
	// The server SHOULD return the value of the MoreInfoUrls ADM element.
	GetMoreInfoUrls(context.Context, *GetMoreInfoUrlsRequest) (*GetMoreInfoUrlsResponse, error)

	// The IUpdate::MsrcSeverity (opnum 35) method retrieves a rating of the severity of
	// the problem that the update fixes.
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
	// The server SHOULD return the value of the MsrcSeverity ADM element.
	GetMsrcSeverity(context.Context, *GetMsrcSeverityRequest) (*GetMsrcSeverityResponse, error)

	// The IUpdate::RecommendedCpuSpeed (opnum 36) method retrieves the CPU speed recommended
	// for the installation of the update.
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
	// The server SHOULD return the value of the RecommendedCpuSpeed ADM element.
	GetRecommendedCpuSpeed(context.Context, *GetRecommendedCpuSpeedRequest) (*GetRecommendedCpuSpeedResponse, error)

	// The IUpdate::RecommendedHardDiskSpace (opnum 37) method retrieves the free hard disk
	// space recommended for the installation of the update.
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
	// The server SHOULD return the value of the RecommendedHardDiskSpace ADM element.
	GetRecommendedHardDiskSpace(context.Context, *GetRecommendedHardDiskSpaceRequest) (*GetRecommendedHardDiskSpaceResponse, error)

	// The IUpdate::RecommendedMemory (opnum 38) method retrieves the amount of memory recommended
	// for the installation of the update.
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
	// The server SHOULD return the value of the RecommendedMemory ADM element.
	GetRecommendedMemory(context.Context, *GetRecommendedMemoryRequest) (*GetRecommendedMemoryResponse, error)

	// The IUpdate::ReleaseNotes (opnum 39) method retrieves the localized release notes
	// for the update.
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
	// The server SHOULD return the value of the ReleaseNotes ADM element.
	GetReleaseNotes(context.Context, *GetReleaseNotesRequest) (*GetReleaseNotesResponse, error)

	// The IUpdate::SecurityBulletinIDs (opnum 40) method retrieves the security bulletin
	// IDs associated with the update.
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
	// The server SHOULD return the value of the SecurityBulletinIDs ADM element.
	GetSecurityBulletinIDs(context.Context, *GetSecurityBulletinIDsRequest) (*GetSecurityBulletinIDsResponse, error)

	// The IUpdate::SupersededUpdateIDs (opnum 41) method retrieves a collection of the
	// IDs of updates that are superseded by the update.
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
	// The server SHOULD return the value of the SupersededUpdateIDs ADM element.
	GetSupersededUpdateIDs(context.Context, *GetSupersededUpdateIDsRequest) (*GetSupersededUpdateIDsResponse, error)

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

	// The ICategory::Type (opnum 15) method retrieves the type of the update category.
	//
	// The IUpdate::Type (opnum 43) method retrieves the type of the update.
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
	// This method SHOULD return the value of the Type ADM element.
	GetType(context.Context, *GetTypeRequest) (*GetTypeResponse, error)

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

	// The IUpdate::UninstallationBehavior (opnum 45) method retrieves a description of
	// the uninstallation behavior for the update.
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
	// The server SHOULD return the value of the UninstallationBehavior ADM element.
	GetUninstallationBehavior(context.Context, *GetUninstallationBehaviorRequest) (*GetUninstallationBehaviorResponse, error)

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

	// The IUpdate::KBArticleIDs (opnum 47) method retrieves a list of Knowledge Base article
	// IDs associated with the update.
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
	// The server SHOULD return the value of the KBArticleIDs ADM element.
	GetKbArticleIDs(context.Context, *GetKbArticleIDsRequest) (*GetKbArticleIDsResponse, error)

	// Opnum48NotUsedOnWire operation.
	// Opnum48NotUsedOnWire

	// The IUpdate::DeploymentAction (opnum 49) method retrieves the deployment action for
	// the update.
	//
	// The DeploymentAction enumeration defines values that describe the server-defined
	// action to be taken on the update by the update agent.
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
	// The server SHOULD return the value of the DeploymentAction ADM element.
	GetDeploymentAction(context.Context, *GetDeploymentActionRequest) (*GetDeploymentActionResponse, error)

	// Opnum50NotUsedOnWire operation.
	// Opnum50NotUsedOnWire

	// The IUpdate::DownloadPriority (opnum 51) method retrieves the download priority for
	// the update.
	//
	// The DownloadPriority enumeration defines values describing the recommended download
	// priority of an update. It is up to the update agent implementation to define the
	// relative differences between download priority levels.<1>
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
	// The server SHOULD return the value of the DownloadPriority ADM element.
	GetDownloadPriority(context.Context, *GetDownloadPriorityRequest) (*GetDownloadPriorityResponse, error)

	// The IUpdate::DownloadContents (opnum 52) method retrieves a collection of download
	// content URLs for the update.
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
	// The server SHOULD return the value of the DownloadContents ADM element.
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
