package uamg

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	dcom_client "github.com/oiweiwei/go-msrpc/msrpc/dcom/client"
	iremunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iremunknown/v0"
	iremunknown2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/iremunknown2/v0"
	iautomaticupdates "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iautomaticupdates/v0"
	iautomaticupdates2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iautomaticupdates2/v0"
	iautomaticupdatesresults "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iautomaticupdatesresults/v0"
	icategory "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/icategory/v0"
	icategorycollection "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/icategorycollection/v0"
	iimageinformation "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iimageinformation/v0"
	iinstallationbehavior "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iinstallationbehavior/v0"
	isearchresult "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/isearchresult/v0"
	istringcollection "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/istringcollection/v0"
	iupdate "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdate/v0"
	iupdate2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdate2/v0"
	iupdate3 "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdate3/v0"
	iupdate4 "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdate4/v0"
	iupdate5 "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdate5/v0"
	iupdatecollection "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdatecollection/v0"
	iupdatedownloadcontent "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdatedownloadcontent/v0"
	iupdatedownloadcontent2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdatedownloadcontent2/v0"
	iupdatedownloadcontentcollection "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdatedownloadcontentcollection/v0"
	iupdateexception "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdateexception/v0"
	iupdateexceptioncollection "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdateexceptioncollection/v0"
	iupdatehistoryentry "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdatehistoryentry/v0"
	iupdatehistoryentry2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdatehistoryentry2/v0"
	iupdatehistoryentrycollection "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdatehistoryentrycollection/v0"
	iupdateidentity "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdateidentity/v0"
	iupdatesearcher "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdatesearcher/v0"
	iupdatesearcher2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdatesearcher2/v0"
	iupdatesearcher3 "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdatesearcher3/v0"
	iupdateservice "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdateservice/v0"
	iupdateservice2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdateservice2/v0"
	iupdateservicecollection "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdateservicecollection/v0"
	iupdateservicemanager "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdateservicemanager/v0"
	iupdateservicemanager2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdateservicemanager2/v0"
	iupdateserviceregistration "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdateserviceregistration/v0"
	iupdatesession "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdatesession/v0"
	iupdatesession2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdatesession2/v0"
	iupdatesession3 "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdatesession3/v0"
	iwindowsdriverupdate "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iwindowsdriverupdate/v0"
	iwindowsdriverupdate2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iwindowsdriverupdate2/v0"
	iwindowsdriverupdate3 "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iwindowsdriverupdate3/v0"
	iwindowsdriverupdate4 "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iwindowsdriverupdate4/v0"
	iwindowsdriverupdate5 "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iwindowsdriverupdate5/v0"
	iwindowsdriverupdateentry "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iwindowsdriverupdateentry/v0"
	iwindowsdriverupdateentrycollection "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iwindowsdriverupdateentrycollection/v0"
	iwindowsupdateagentinfo "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iwindowsupdateagentinfo/v0"
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
	_ = dcom_client.GoPackage
	_ = dcom.GoPackage
	_ = iremunknown.GoPackage
	_ = iremunknown2.GoPackage
	_ = istringcollection.GoPackage
	_ = iwindowsupdateagentinfo.GoPackage
	_ = iautomaticupdatesresults.GoPackage
	_ = iautomaticupdates.GoPackage
	_ = iautomaticupdates2.GoPackage
	_ = iupdateidentity.GoPackage
	_ = iimageinformation.GoPackage
	_ = icategory.GoPackage
	_ = icategorycollection.GoPackage
	_ = iinstallationbehavior.GoPackage
	_ = iupdatedownloadcontent.GoPackage
	_ = iupdatedownloadcontent2.GoPackage
	_ = iupdatedownloadcontentcollection.GoPackage
	_ = iupdate.GoPackage
	_ = iwindowsdriverupdate.GoPackage
	_ = iupdate2.GoPackage
	_ = iupdate3.GoPackage
	_ = iupdate4.GoPackage
	_ = iupdate5.GoPackage
	_ = iwindowsdriverupdateentry.GoPackage
	_ = iwindowsdriverupdateentrycollection.GoPackage
	_ = iwindowsdriverupdate2.GoPackage
	_ = iwindowsdriverupdate3.GoPackage
	_ = iwindowsdriverupdate4.GoPackage
	_ = iwindowsdriverupdate5.GoPackage
	_ = iupdatecollection.GoPackage
	_ = iupdateexception.GoPackage
	_ = iupdateexceptioncollection.GoPackage
	_ = isearchresult.GoPackage
	_ = iupdatehistoryentry.GoPackage
	_ = iupdatehistoryentry2.GoPackage
	_ = iupdatehistoryentrycollection.GoPackage
	_ = iupdatesearcher.GoPackage
	_ = iupdatesearcher2.GoPackage
	_ = iupdatesearcher3.GoPackage
	_ = iupdatesession.GoPackage
	_ = iupdatesession2.GoPackage
	_ = iupdatesession3.GoPackage
	_ = iupdateservice.GoPackage
	_ = iupdateservice2.GoPackage
	_ = iupdateservicecollection.GoPackage
	_ = iupdateserviceregistration.GoPackage
	_ = iupdateservicemanager.GoPackage
	_ = iupdateservicemanager2.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/uamg"
)

// dcom/uamg client set.
type Client interface {

	// DCOM common interfaces
	RemoteUnknown() iremunknown.RemoteUnknownClient
	RemoteUnknown2() iremunknown2.RemoteUnknown2Client

	// Package specific interfaces
	StringCollection() istringcollection.StringCollectionClient
	WindowsUpdateAgentInfo() iwindowsupdateagentinfo.WindowsUpdateAgentInfoClient
	AutomaticUpdatesResults() iautomaticupdatesresults.AutomaticUpdatesResultsClient
	AutomaticUpdates() iautomaticupdates.AutomaticUpdatesClient
	AutomaticUpdates2() iautomaticupdates2.AutomaticUpdates2Client
	UpdateIdentity() iupdateidentity.UpdateIdentityClient
	ImageInformation() iimageinformation.ImageInformationClient
	Category() icategory.CategoryClient
	CategoryCollection() icategorycollection.CategoryCollectionClient
	InstallationBehavior() iinstallationbehavior.InstallationBehaviorClient
	UpdateDownloadContent() iupdatedownloadcontent.UpdateDownloadContentClient
	UpdateDownloadContent2() iupdatedownloadcontent2.UpdateDownloadContent2Client
	UpdateDownloadContentCollection() iupdatedownloadcontentcollection.UpdateDownloadContentCollectionClient
	Update() iupdate.UpdateClient
	WindowsDriverUpdate() iwindowsdriverupdate.WindowsDriverUpdateClient
	Update2() iupdate2.Update2Client
	Update3() iupdate3.Update3Client
	Update4() iupdate4.Update4Client
	Update5() iupdate5.Update5Client
	WindowsDriverUpdateEntry() iwindowsdriverupdateentry.WindowsDriverUpdateEntryClient
	WindowsDriverUpdateEntryCollection() iwindowsdriverupdateentrycollection.WindowsDriverUpdateEntryCollectionClient
	WindowsDriverUpdate2() iwindowsdriverupdate2.WindowsDriverUpdate2Client
	WindowsDriverUpdate3() iwindowsdriverupdate3.WindowsDriverUpdate3Client
	WindowsDriverUpdate4() iwindowsdriverupdate4.WindowsDriverUpdate4Client
	WindowsDriverUpdate5() iwindowsdriverupdate5.WindowsDriverUpdate5Client
	UpdateCollection() iupdatecollection.UpdateCollectionClient
	UpdateException() iupdateexception.UpdateExceptionClient
	UpdateExceptionCollection() iupdateexceptioncollection.UpdateExceptionCollectionClient
	SearchResult() isearchresult.SearchResultClient
	UpdateHistoryEntry() iupdatehistoryentry.UpdateHistoryEntryClient
	UpdateHistoryEntry2() iupdatehistoryentry2.UpdateHistoryEntry2Client
	UpdateHistoryEntryCollection() iupdatehistoryentrycollection.UpdateHistoryEntryCollectionClient
	UpdateSearcher() iupdatesearcher.UpdateSearcherClient
	UpdateSearcher2() iupdatesearcher2.UpdateSearcher2Client
	UpdateSearcher3() iupdatesearcher3.UpdateSearcher3Client
	UpdateSession() iupdatesession.UpdateSessionClient
	UpdateSession2() iupdatesession2.UpdateSession2Client
	UpdateSession3() iupdatesession3.UpdateSession3Client
	UpdateService() iupdateservice.UpdateServiceClient
	UpdateService2() iupdateservice2.UpdateService2Client
	UpdateServiceCollection() iupdateservicecollection.UpdateServiceCollectionClient
	UpdateServiceRegistration() iupdateserviceregistration.UpdateServiceRegistrationClient
	UpdateServiceManager() iupdateservicemanager.UpdateServiceManagerClient
	UpdateServiceManager2() iupdateservicemanager2.UpdateServiceManager2Client
	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) Client
}
type xxx_DefaultClient struct {
	cc dcerpc.Conn

	dcomClient dcom_client.Client

	stringCollection                   istringcollection.StringCollectionClient
	windowsUpdateAgentInfo             iwindowsupdateagentinfo.WindowsUpdateAgentInfoClient
	automaticUpdatesResults            iautomaticupdatesresults.AutomaticUpdatesResultsClient
	automaticUpdates                   iautomaticupdates.AutomaticUpdatesClient
	automaticUpdates2                  iautomaticupdates2.AutomaticUpdates2Client
	updateIdentity                     iupdateidentity.UpdateIdentityClient
	imageInformation                   iimageinformation.ImageInformationClient
	category                           icategory.CategoryClient
	categoryCollection                 icategorycollection.CategoryCollectionClient
	installationBehavior               iinstallationbehavior.InstallationBehaviorClient
	updateDownloadContent              iupdatedownloadcontent.UpdateDownloadContentClient
	updateDownloadContent2             iupdatedownloadcontent2.UpdateDownloadContent2Client
	updateDownloadContentCollection    iupdatedownloadcontentcollection.UpdateDownloadContentCollectionClient
	update                             iupdate.UpdateClient
	windowsDriverUpdate                iwindowsdriverupdate.WindowsDriverUpdateClient
	update2                            iupdate2.Update2Client
	update3                            iupdate3.Update3Client
	update4                            iupdate4.Update4Client
	update5                            iupdate5.Update5Client
	windowsDriverUpdateEntry           iwindowsdriverupdateentry.WindowsDriverUpdateEntryClient
	windowsDriverUpdateEntryCollection iwindowsdriverupdateentrycollection.WindowsDriverUpdateEntryCollectionClient
	windowsDriverUpdate2               iwindowsdriverupdate2.WindowsDriverUpdate2Client
	windowsDriverUpdate3               iwindowsdriverupdate3.WindowsDriverUpdate3Client
	windowsDriverUpdate4               iwindowsdriverupdate4.WindowsDriverUpdate4Client
	windowsDriverUpdate5               iwindowsdriverupdate5.WindowsDriverUpdate5Client
	updateCollection                   iupdatecollection.UpdateCollectionClient
	updateException                    iupdateexception.UpdateExceptionClient
	updateExceptionCollection          iupdateexceptioncollection.UpdateExceptionCollectionClient
	searchResult                       isearchresult.SearchResultClient
	updateHistoryEntry                 iupdatehistoryentry.UpdateHistoryEntryClient
	updateHistoryEntry2                iupdatehistoryentry2.UpdateHistoryEntry2Client
	updateHistoryEntryCollection       iupdatehistoryentrycollection.UpdateHistoryEntryCollectionClient
	updateSearcher                     iupdatesearcher.UpdateSearcherClient
	updateSearcher2                    iupdatesearcher2.UpdateSearcher2Client
	updateSearcher3                    iupdatesearcher3.UpdateSearcher3Client
	updateSession                      iupdatesession.UpdateSessionClient
	updateSession2                     iupdatesession2.UpdateSession2Client
	updateSession3                     iupdatesession3.UpdateSession3Client
	updateService                      iupdateservice.UpdateServiceClient
	updateService2                     iupdateservice2.UpdateService2Client
	updateServiceCollection            iupdateservicecollection.UpdateServiceCollectionClient
	updateServiceRegistration          iupdateserviceregistration.UpdateServiceRegistrationClient
	updateServiceManager               iupdateservicemanager.UpdateServiceManagerClient
	updateServiceManager2              iupdateservicemanager2.UpdateServiceManager2Client
}

func (o *xxx_DefaultClient) RemoteUnknown() iremunknown.RemoteUnknownClient {
	return o.dcomClient.RemoteUnknown()
}

func (o *xxx_DefaultClient) RemoteUnknown2() iremunknown2.RemoteUnknown2Client {
	return o.dcomClient.RemoteUnknown2()
}

func (o *xxx_DefaultClient) StringCollection() istringcollection.StringCollectionClient {
	return o.stringCollection
}

func (o *xxx_DefaultClient) WindowsUpdateAgentInfo() iwindowsupdateagentinfo.WindowsUpdateAgentInfoClient {
	return o.windowsUpdateAgentInfo
}

func (o *xxx_DefaultClient) AutomaticUpdatesResults() iautomaticupdatesresults.AutomaticUpdatesResultsClient {
	return o.automaticUpdatesResults
}

func (o *xxx_DefaultClient) AutomaticUpdates() iautomaticupdates.AutomaticUpdatesClient {
	return o.automaticUpdates
}

func (o *xxx_DefaultClient) AutomaticUpdates2() iautomaticupdates2.AutomaticUpdates2Client {
	return o.automaticUpdates2
}

func (o *xxx_DefaultClient) UpdateIdentity() iupdateidentity.UpdateIdentityClient {
	return o.updateIdentity
}

func (o *xxx_DefaultClient) ImageInformation() iimageinformation.ImageInformationClient {
	return o.imageInformation
}

func (o *xxx_DefaultClient) Category() icategory.CategoryClient {
	return o.category
}

func (o *xxx_DefaultClient) CategoryCollection() icategorycollection.CategoryCollectionClient {
	return o.categoryCollection
}

func (o *xxx_DefaultClient) InstallationBehavior() iinstallationbehavior.InstallationBehaviorClient {
	return o.installationBehavior
}

func (o *xxx_DefaultClient) UpdateDownloadContent() iupdatedownloadcontent.UpdateDownloadContentClient {
	return o.updateDownloadContent
}

func (o *xxx_DefaultClient) UpdateDownloadContent2() iupdatedownloadcontent2.UpdateDownloadContent2Client {
	return o.updateDownloadContent2
}

func (o *xxx_DefaultClient) UpdateDownloadContentCollection() iupdatedownloadcontentcollection.UpdateDownloadContentCollectionClient {
	return o.updateDownloadContentCollection
}

func (o *xxx_DefaultClient) Update() iupdate.UpdateClient {
	return o.update
}

func (o *xxx_DefaultClient) WindowsDriverUpdate() iwindowsdriverupdate.WindowsDriverUpdateClient {
	return o.windowsDriverUpdate
}

func (o *xxx_DefaultClient) Update2() iupdate2.Update2Client {
	return o.update2
}

func (o *xxx_DefaultClient) Update3() iupdate3.Update3Client {
	return o.update3
}

func (o *xxx_DefaultClient) Update4() iupdate4.Update4Client {
	return o.update4
}

func (o *xxx_DefaultClient) Update5() iupdate5.Update5Client {
	return o.update5
}

func (o *xxx_DefaultClient) WindowsDriverUpdateEntry() iwindowsdriverupdateentry.WindowsDriverUpdateEntryClient {
	return o.windowsDriverUpdateEntry
}

func (o *xxx_DefaultClient) WindowsDriverUpdateEntryCollection() iwindowsdriverupdateentrycollection.WindowsDriverUpdateEntryCollectionClient {
	return o.windowsDriverUpdateEntryCollection
}

func (o *xxx_DefaultClient) WindowsDriverUpdate2() iwindowsdriverupdate2.WindowsDriverUpdate2Client {
	return o.windowsDriverUpdate2
}

func (o *xxx_DefaultClient) WindowsDriverUpdate3() iwindowsdriverupdate3.WindowsDriverUpdate3Client {
	return o.windowsDriverUpdate3
}

func (o *xxx_DefaultClient) WindowsDriverUpdate4() iwindowsdriverupdate4.WindowsDriverUpdate4Client {
	return o.windowsDriverUpdate4
}

func (o *xxx_DefaultClient) WindowsDriverUpdate5() iwindowsdriverupdate5.WindowsDriverUpdate5Client {
	return o.windowsDriverUpdate5
}

func (o *xxx_DefaultClient) UpdateCollection() iupdatecollection.UpdateCollectionClient {
	return o.updateCollection
}

func (o *xxx_DefaultClient) UpdateException() iupdateexception.UpdateExceptionClient {
	return o.updateException
}

func (o *xxx_DefaultClient) UpdateExceptionCollection() iupdateexceptioncollection.UpdateExceptionCollectionClient {
	return o.updateExceptionCollection
}

func (o *xxx_DefaultClient) SearchResult() isearchresult.SearchResultClient {
	return o.searchResult
}

func (o *xxx_DefaultClient) UpdateHistoryEntry() iupdatehistoryentry.UpdateHistoryEntryClient {
	return o.updateHistoryEntry
}

func (o *xxx_DefaultClient) UpdateHistoryEntry2() iupdatehistoryentry2.UpdateHistoryEntry2Client {
	return o.updateHistoryEntry2
}

func (o *xxx_DefaultClient) UpdateHistoryEntryCollection() iupdatehistoryentrycollection.UpdateHistoryEntryCollectionClient {
	return o.updateHistoryEntryCollection
}

func (o *xxx_DefaultClient) UpdateSearcher() iupdatesearcher.UpdateSearcherClient {
	return o.updateSearcher
}

func (o *xxx_DefaultClient) UpdateSearcher2() iupdatesearcher2.UpdateSearcher2Client {
	return o.updateSearcher2
}

func (o *xxx_DefaultClient) UpdateSearcher3() iupdatesearcher3.UpdateSearcher3Client {
	return o.updateSearcher3
}

func (o *xxx_DefaultClient) UpdateSession() iupdatesession.UpdateSessionClient {
	return o.updateSession
}

func (o *xxx_DefaultClient) UpdateSession2() iupdatesession2.UpdateSession2Client {
	return o.updateSession2
}

func (o *xxx_DefaultClient) UpdateSession3() iupdatesession3.UpdateSession3Client {
	return o.updateSession3
}

func (o *xxx_DefaultClient) UpdateService() iupdateservice.UpdateServiceClient {
	return o.updateService
}

func (o *xxx_DefaultClient) UpdateService2() iupdateservice2.UpdateService2Client {
	return o.updateService2
}

func (o *xxx_DefaultClient) UpdateServiceCollection() iupdateservicecollection.UpdateServiceCollectionClient {
	return o.updateServiceCollection
}

func (o *xxx_DefaultClient) UpdateServiceRegistration() iupdateserviceregistration.UpdateServiceRegistrationClient {
	return o.updateServiceRegistration
}

func (o *xxx_DefaultClient) UpdateServiceManager() iupdateservicemanager.UpdateServiceManagerClient {
	return o.updateServiceManager
}

func (o *xxx_DefaultClient) UpdateServiceManager2() iupdateservicemanager2.UpdateServiceManager2Client {
	return o.updateServiceManager2
}

func NewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Client, error) {

	opts = append(opts,
		dcerpc.WithAbstractSyntax(istringcollection.StringCollectionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwindowsupdateagentinfo.WindowsUpdateAgentInfoSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iautomaticupdatesresults.AutomaticUpdatesResultsSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iautomaticupdates.AutomaticUpdatesSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iautomaticupdates2.AutomaticUpdates2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(iupdateidentity.UpdateIdentitySyntaxV0_0),
		dcerpc.WithAbstractSyntax(iimageinformation.ImageInformationSyntaxV0_0),
		dcerpc.WithAbstractSyntax(icategory.CategorySyntaxV0_0),
		dcerpc.WithAbstractSyntax(icategorycollection.CategoryCollectionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iinstallationbehavior.InstallationBehaviorSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iupdatedownloadcontent.UpdateDownloadContentSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iupdatedownloadcontent2.UpdateDownloadContent2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(iupdatedownloadcontentcollection.UpdateDownloadContentCollectionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iupdate.UpdateSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwindowsdriverupdate.WindowsDriverUpdateSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iupdate2.Update2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(iupdate3.Update3SyntaxV0_0),
		dcerpc.WithAbstractSyntax(iupdate4.Update4SyntaxV0_0),
		dcerpc.WithAbstractSyntax(iupdate5.Update5SyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwindowsdriverupdateentry.WindowsDriverUpdateEntrySyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwindowsdriverupdateentrycollection.WindowsDriverUpdateEntryCollectionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwindowsdriverupdate2.WindowsDriverUpdate2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwindowsdriverupdate3.WindowsDriverUpdate3SyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwindowsdriverupdate4.WindowsDriverUpdate4SyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwindowsdriverupdate5.WindowsDriverUpdate5SyntaxV0_0),
		dcerpc.WithAbstractSyntax(iupdatecollection.UpdateCollectionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iupdateexception.UpdateExceptionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iupdateexceptioncollection.UpdateExceptionCollectionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(isearchresult.SearchResultSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iupdatehistoryentry.UpdateHistoryEntrySyntaxV0_0),
		dcerpc.WithAbstractSyntax(iupdatehistoryentry2.UpdateHistoryEntry2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(iupdatehistoryentrycollection.UpdateHistoryEntryCollectionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iupdatesearcher.UpdateSearcherSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iupdatesearcher2.UpdateSearcher2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(iupdatesearcher3.UpdateSearcher3SyntaxV0_0),
		dcerpc.WithAbstractSyntax(iupdatesession.UpdateSessionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iupdatesession2.UpdateSession2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(iupdatesession3.UpdateSession3SyntaxV0_0),
		dcerpc.WithAbstractSyntax(iupdateservice.UpdateServiceSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iupdateservice2.UpdateService2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(iupdateservicecollection.UpdateServiceCollectionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iupdateserviceregistration.UpdateServiceRegistrationSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iupdateservicemanager.UpdateServiceManagerSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iupdateservicemanager2.UpdateServiceManager2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(iremunknown.RemoteUnknownSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iremunknown2.RemoteUnknown2SyntaxV0_0),
	)

	cc, err := cc.Bind(ctx, opts...)
	if err != nil {
		return nil, err
	}

	o := &xxx_DefaultClient{cc: cc}

	dcomClient, err := dcom_client.NewClient(ctx, cc, append(opts, dcerpc.WithNoBind(cc))...)
	if err != nil {
		return nil, err
	}
	o.dcomClient = dcomClient

	sub, ok := cc.(dcerpc.SubConn)
	if !ok {
		return nil, fmt.Errorf("sub-conn is not supported")
	}

	stringCollectionSubConn, err := sub.SubConn(ctx, istringcollection.StringCollectionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		stringCollectionSubConn = sub
	}

	o.stringCollection, err = istringcollection.NewStringCollectionClient(ctx, stringCollectionSubConn, append(opts, dcerpc.WithNoBind(stringCollectionSubConn))...)

	windowsUpdateAgentInfoSubConn, err := sub.SubConn(ctx, iwindowsupdateagentinfo.WindowsUpdateAgentInfoSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		windowsUpdateAgentInfoSubConn = sub
	}

	o.windowsUpdateAgentInfo, err = iwindowsupdateagentinfo.NewWindowsUpdateAgentInfoClient(ctx, windowsUpdateAgentInfoSubConn, append(opts, dcerpc.WithNoBind(windowsUpdateAgentInfoSubConn))...)

	automaticUpdatesResultsSubConn, err := sub.SubConn(ctx, iautomaticupdatesresults.AutomaticUpdatesResultsSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		automaticUpdatesResultsSubConn = sub
	}

	o.automaticUpdatesResults, err = iautomaticupdatesresults.NewAutomaticUpdatesResultsClient(ctx, automaticUpdatesResultsSubConn, append(opts, dcerpc.WithNoBind(automaticUpdatesResultsSubConn))...)

	automaticUpdatesSubConn, err := sub.SubConn(ctx, iautomaticupdates.AutomaticUpdatesSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		automaticUpdatesSubConn = sub
	}

	o.automaticUpdates, err = iautomaticupdates.NewAutomaticUpdatesClient(ctx, automaticUpdatesSubConn, append(opts, dcerpc.WithNoBind(automaticUpdatesSubConn))...)

	automaticUpdates2SubConn, err := sub.SubConn(ctx, iautomaticupdates2.AutomaticUpdates2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		automaticUpdates2SubConn = sub
	}

	o.automaticUpdates2, err = iautomaticupdates2.NewAutomaticUpdates2Client(ctx, automaticUpdates2SubConn, append(opts, dcerpc.WithNoBind(automaticUpdates2SubConn))...)

	updateIdentitySubConn, err := sub.SubConn(ctx, iupdateidentity.UpdateIdentitySyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		updateIdentitySubConn = sub
	}

	o.updateIdentity, err = iupdateidentity.NewUpdateIdentityClient(ctx, updateIdentitySubConn, append(opts, dcerpc.WithNoBind(updateIdentitySubConn))...)

	imageInformationSubConn, err := sub.SubConn(ctx, iimageinformation.ImageInformationSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imageInformationSubConn = sub
	}

	o.imageInformation, err = iimageinformation.NewImageInformationClient(ctx, imageInformationSubConn, append(opts, dcerpc.WithNoBind(imageInformationSubConn))...)

	categorySubConn, err := sub.SubConn(ctx, icategory.CategorySyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		categorySubConn = sub
	}

	o.category, err = icategory.NewCategoryClient(ctx, categorySubConn, append(opts, dcerpc.WithNoBind(categorySubConn))...)

	categoryCollectionSubConn, err := sub.SubConn(ctx, icategorycollection.CategoryCollectionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		categoryCollectionSubConn = sub
	}

	o.categoryCollection, err = icategorycollection.NewCategoryCollectionClient(ctx, categoryCollectionSubConn, append(opts, dcerpc.WithNoBind(categoryCollectionSubConn))...)

	installationBehaviorSubConn, err := sub.SubConn(ctx, iinstallationbehavior.InstallationBehaviorSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		installationBehaviorSubConn = sub
	}

	o.installationBehavior, err = iinstallationbehavior.NewInstallationBehaviorClient(ctx, installationBehaviorSubConn, append(opts, dcerpc.WithNoBind(installationBehaviorSubConn))...)

	updateDownloadContentSubConn, err := sub.SubConn(ctx, iupdatedownloadcontent.UpdateDownloadContentSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		updateDownloadContentSubConn = sub
	}

	o.updateDownloadContent, err = iupdatedownloadcontent.NewUpdateDownloadContentClient(ctx, updateDownloadContentSubConn, append(opts, dcerpc.WithNoBind(updateDownloadContentSubConn))...)

	updateDownloadContent2SubConn, err := sub.SubConn(ctx, iupdatedownloadcontent2.UpdateDownloadContent2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		updateDownloadContent2SubConn = sub
	}

	o.updateDownloadContent2, err = iupdatedownloadcontent2.NewUpdateDownloadContent2Client(ctx, updateDownloadContent2SubConn, append(opts, dcerpc.WithNoBind(updateDownloadContent2SubConn))...)

	updateDownloadContentCollectionSubConn, err := sub.SubConn(ctx, iupdatedownloadcontentcollection.UpdateDownloadContentCollectionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		updateDownloadContentCollectionSubConn = sub
	}

	o.updateDownloadContentCollection, err = iupdatedownloadcontentcollection.NewUpdateDownloadContentCollectionClient(ctx, updateDownloadContentCollectionSubConn, append(opts, dcerpc.WithNoBind(updateDownloadContentCollectionSubConn))...)

	updateSubConn, err := sub.SubConn(ctx, iupdate.UpdateSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		updateSubConn = sub
	}

	o.update, err = iupdate.NewUpdateClient(ctx, updateSubConn, append(opts, dcerpc.WithNoBind(updateSubConn))...)

	windowsDriverUpdateSubConn, err := sub.SubConn(ctx, iwindowsdriverupdate.WindowsDriverUpdateSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		windowsDriverUpdateSubConn = sub
	}

	o.windowsDriverUpdate, err = iwindowsdriverupdate.NewWindowsDriverUpdateClient(ctx, windowsDriverUpdateSubConn, append(opts, dcerpc.WithNoBind(windowsDriverUpdateSubConn))...)

	update2SubConn, err := sub.SubConn(ctx, iupdate2.Update2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		update2SubConn = sub
	}

	o.update2, err = iupdate2.NewUpdate2Client(ctx, update2SubConn, append(opts, dcerpc.WithNoBind(update2SubConn))...)

	update3SubConn, err := sub.SubConn(ctx, iupdate3.Update3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		update3SubConn = sub
	}

	o.update3, err = iupdate3.NewUpdate3Client(ctx, update3SubConn, append(opts, dcerpc.WithNoBind(update3SubConn))...)

	update4SubConn, err := sub.SubConn(ctx, iupdate4.Update4SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		update4SubConn = sub
	}

	o.update4, err = iupdate4.NewUpdate4Client(ctx, update4SubConn, append(opts, dcerpc.WithNoBind(update4SubConn))...)

	update5SubConn, err := sub.SubConn(ctx, iupdate5.Update5SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		update5SubConn = sub
	}

	o.update5, err = iupdate5.NewUpdate5Client(ctx, update5SubConn, append(opts, dcerpc.WithNoBind(update5SubConn))...)

	windowsDriverUpdateEntrySubConn, err := sub.SubConn(ctx, iwindowsdriverupdateentry.WindowsDriverUpdateEntrySyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		windowsDriverUpdateEntrySubConn = sub
	}

	o.windowsDriverUpdateEntry, err = iwindowsdriverupdateentry.NewWindowsDriverUpdateEntryClient(ctx, windowsDriverUpdateEntrySubConn, append(opts, dcerpc.WithNoBind(windowsDriverUpdateEntrySubConn))...)

	windowsDriverUpdateEntryCollectionSubConn, err := sub.SubConn(ctx, iwindowsdriverupdateentrycollection.WindowsDriverUpdateEntryCollectionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		windowsDriverUpdateEntryCollectionSubConn = sub
	}

	o.windowsDriverUpdateEntryCollection, err = iwindowsdriverupdateentrycollection.NewWindowsDriverUpdateEntryCollectionClient(ctx, windowsDriverUpdateEntryCollectionSubConn, append(opts, dcerpc.WithNoBind(windowsDriverUpdateEntryCollectionSubConn))...)

	windowsDriverUpdate2SubConn, err := sub.SubConn(ctx, iwindowsdriverupdate2.WindowsDriverUpdate2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		windowsDriverUpdate2SubConn = sub
	}

	o.windowsDriverUpdate2, err = iwindowsdriverupdate2.NewWindowsDriverUpdate2Client(ctx, windowsDriverUpdate2SubConn, append(opts, dcerpc.WithNoBind(windowsDriverUpdate2SubConn))...)

	windowsDriverUpdate3SubConn, err := sub.SubConn(ctx, iwindowsdriverupdate3.WindowsDriverUpdate3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		windowsDriverUpdate3SubConn = sub
	}

	o.windowsDriverUpdate3, err = iwindowsdriverupdate3.NewWindowsDriverUpdate3Client(ctx, windowsDriverUpdate3SubConn, append(opts, dcerpc.WithNoBind(windowsDriverUpdate3SubConn))...)

	windowsDriverUpdate4SubConn, err := sub.SubConn(ctx, iwindowsdriverupdate4.WindowsDriverUpdate4SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		windowsDriverUpdate4SubConn = sub
	}

	o.windowsDriverUpdate4, err = iwindowsdriverupdate4.NewWindowsDriverUpdate4Client(ctx, windowsDriverUpdate4SubConn, append(opts, dcerpc.WithNoBind(windowsDriverUpdate4SubConn))...)

	windowsDriverUpdate5SubConn, err := sub.SubConn(ctx, iwindowsdriverupdate5.WindowsDriverUpdate5SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		windowsDriverUpdate5SubConn = sub
	}

	o.windowsDriverUpdate5, err = iwindowsdriverupdate5.NewWindowsDriverUpdate5Client(ctx, windowsDriverUpdate5SubConn, append(opts, dcerpc.WithNoBind(windowsDriverUpdate5SubConn))...)

	updateCollectionSubConn, err := sub.SubConn(ctx, iupdatecollection.UpdateCollectionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		updateCollectionSubConn = sub
	}

	o.updateCollection, err = iupdatecollection.NewUpdateCollectionClient(ctx, updateCollectionSubConn, append(opts, dcerpc.WithNoBind(updateCollectionSubConn))...)

	updateExceptionSubConn, err := sub.SubConn(ctx, iupdateexception.UpdateExceptionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		updateExceptionSubConn = sub
	}

	o.updateException, err = iupdateexception.NewUpdateExceptionClient(ctx, updateExceptionSubConn, append(opts, dcerpc.WithNoBind(updateExceptionSubConn))...)

	updateExceptionCollectionSubConn, err := sub.SubConn(ctx, iupdateexceptioncollection.UpdateExceptionCollectionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		updateExceptionCollectionSubConn = sub
	}

	o.updateExceptionCollection, err = iupdateexceptioncollection.NewUpdateExceptionCollectionClient(ctx, updateExceptionCollectionSubConn, append(opts, dcerpc.WithNoBind(updateExceptionCollectionSubConn))...)

	searchResultSubConn, err := sub.SubConn(ctx, isearchresult.SearchResultSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		searchResultSubConn = sub
	}

	o.searchResult, err = isearchresult.NewSearchResultClient(ctx, searchResultSubConn, append(opts, dcerpc.WithNoBind(searchResultSubConn))...)

	updateHistoryEntrySubConn, err := sub.SubConn(ctx, iupdatehistoryentry.UpdateHistoryEntrySyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		updateHistoryEntrySubConn = sub
	}

	o.updateHistoryEntry, err = iupdatehistoryentry.NewUpdateHistoryEntryClient(ctx, updateHistoryEntrySubConn, append(opts, dcerpc.WithNoBind(updateHistoryEntrySubConn))...)

	updateHistoryEntry2SubConn, err := sub.SubConn(ctx, iupdatehistoryentry2.UpdateHistoryEntry2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		updateHistoryEntry2SubConn = sub
	}

	o.updateHistoryEntry2, err = iupdatehistoryentry2.NewUpdateHistoryEntry2Client(ctx, updateHistoryEntry2SubConn, append(opts, dcerpc.WithNoBind(updateHistoryEntry2SubConn))...)

	updateHistoryEntryCollectionSubConn, err := sub.SubConn(ctx, iupdatehistoryentrycollection.UpdateHistoryEntryCollectionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		updateHistoryEntryCollectionSubConn = sub
	}

	o.updateHistoryEntryCollection, err = iupdatehistoryentrycollection.NewUpdateHistoryEntryCollectionClient(ctx, updateHistoryEntryCollectionSubConn, append(opts, dcerpc.WithNoBind(updateHistoryEntryCollectionSubConn))...)

	updateSearcherSubConn, err := sub.SubConn(ctx, iupdatesearcher.UpdateSearcherSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		updateSearcherSubConn = sub
	}

	o.updateSearcher, err = iupdatesearcher.NewUpdateSearcherClient(ctx, updateSearcherSubConn, append(opts, dcerpc.WithNoBind(updateSearcherSubConn))...)

	updateSearcher2SubConn, err := sub.SubConn(ctx, iupdatesearcher2.UpdateSearcher2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		updateSearcher2SubConn = sub
	}

	o.updateSearcher2, err = iupdatesearcher2.NewUpdateSearcher2Client(ctx, updateSearcher2SubConn, append(opts, dcerpc.WithNoBind(updateSearcher2SubConn))...)

	updateSearcher3SubConn, err := sub.SubConn(ctx, iupdatesearcher3.UpdateSearcher3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		updateSearcher3SubConn = sub
	}

	o.updateSearcher3, err = iupdatesearcher3.NewUpdateSearcher3Client(ctx, updateSearcher3SubConn, append(opts, dcerpc.WithNoBind(updateSearcher3SubConn))...)

	updateSessionSubConn, err := sub.SubConn(ctx, iupdatesession.UpdateSessionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		updateSessionSubConn = sub
	}

	o.updateSession, err = iupdatesession.NewUpdateSessionClient(ctx, updateSessionSubConn, append(opts, dcerpc.WithNoBind(updateSessionSubConn))...)

	updateSession2SubConn, err := sub.SubConn(ctx, iupdatesession2.UpdateSession2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		updateSession2SubConn = sub
	}

	o.updateSession2, err = iupdatesession2.NewUpdateSession2Client(ctx, updateSession2SubConn, append(opts, dcerpc.WithNoBind(updateSession2SubConn))...)

	updateSession3SubConn, err := sub.SubConn(ctx, iupdatesession3.UpdateSession3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		updateSession3SubConn = sub
	}

	o.updateSession3, err = iupdatesession3.NewUpdateSession3Client(ctx, updateSession3SubConn, append(opts, dcerpc.WithNoBind(updateSession3SubConn))...)

	updateServiceSubConn, err := sub.SubConn(ctx, iupdateservice.UpdateServiceSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		updateServiceSubConn = sub
	}

	o.updateService, err = iupdateservice.NewUpdateServiceClient(ctx, updateServiceSubConn, append(opts, dcerpc.WithNoBind(updateServiceSubConn))...)

	updateService2SubConn, err := sub.SubConn(ctx, iupdateservice2.UpdateService2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		updateService2SubConn = sub
	}

	o.updateService2, err = iupdateservice2.NewUpdateService2Client(ctx, updateService2SubConn, append(opts, dcerpc.WithNoBind(updateService2SubConn))...)

	updateServiceCollectionSubConn, err := sub.SubConn(ctx, iupdateservicecollection.UpdateServiceCollectionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		updateServiceCollectionSubConn = sub
	}

	o.updateServiceCollection, err = iupdateservicecollection.NewUpdateServiceCollectionClient(ctx, updateServiceCollectionSubConn, append(opts, dcerpc.WithNoBind(updateServiceCollectionSubConn))...)

	updateServiceRegistrationSubConn, err := sub.SubConn(ctx, iupdateserviceregistration.UpdateServiceRegistrationSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		updateServiceRegistrationSubConn = sub
	}

	o.updateServiceRegistration, err = iupdateserviceregistration.NewUpdateServiceRegistrationClient(ctx, updateServiceRegistrationSubConn, append(opts, dcerpc.WithNoBind(updateServiceRegistrationSubConn))...)

	updateServiceManagerSubConn, err := sub.SubConn(ctx, iupdateservicemanager.UpdateServiceManagerSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		updateServiceManagerSubConn = sub
	}

	o.updateServiceManager, err = iupdateservicemanager.NewUpdateServiceManagerClient(ctx, updateServiceManagerSubConn, append(opts, dcerpc.WithNoBind(updateServiceManagerSubConn))...)

	updateServiceManager2SubConn, err := sub.SubConn(ctx, iupdateservicemanager2.UpdateServiceManager2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		updateServiceManager2SubConn = sub
	}

	o.updateServiceManager2, err = iupdateservicemanager2.NewUpdateServiceManager2Client(ctx, updateServiceManager2SubConn, append(opts, dcerpc.WithNoBind(updateServiceManager2SubConn))...)
	return o, nil
}

func (o *xxx_DefaultClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultClient) IPID(ctx context.Context, ipid *dcom.IPID) Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultClient{
		dcomClient:                         o.dcomClient.IPID(ctx, ipid),
		stringCollection:                   o.stringCollection.IPID(ctx, ipid),
		windowsUpdateAgentInfo:             o.windowsUpdateAgentInfo.IPID(ctx, ipid),
		automaticUpdatesResults:            o.automaticUpdatesResults.IPID(ctx, ipid),
		automaticUpdates:                   o.automaticUpdates.IPID(ctx, ipid),
		automaticUpdates2:                  o.automaticUpdates2.IPID(ctx, ipid),
		updateIdentity:                     o.updateIdentity.IPID(ctx, ipid),
		imageInformation:                   o.imageInformation.IPID(ctx, ipid),
		category:                           o.category.IPID(ctx, ipid),
		categoryCollection:                 o.categoryCollection.IPID(ctx, ipid),
		installationBehavior:               o.installationBehavior.IPID(ctx, ipid),
		updateDownloadContent:              o.updateDownloadContent.IPID(ctx, ipid),
		updateDownloadContent2:             o.updateDownloadContent2.IPID(ctx, ipid),
		updateDownloadContentCollection:    o.updateDownloadContentCollection.IPID(ctx, ipid),
		update:                             o.update.IPID(ctx, ipid),
		windowsDriverUpdate:                o.windowsDriverUpdate.IPID(ctx, ipid),
		update2:                            o.update2.IPID(ctx, ipid),
		update3:                            o.update3.IPID(ctx, ipid),
		update4:                            o.update4.IPID(ctx, ipid),
		update5:                            o.update5.IPID(ctx, ipid),
		windowsDriverUpdateEntry:           o.windowsDriverUpdateEntry.IPID(ctx, ipid),
		windowsDriverUpdateEntryCollection: o.windowsDriverUpdateEntryCollection.IPID(ctx, ipid),
		windowsDriverUpdate2:               o.windowsDriverUpdate2.IPID(ctx, ipid),
		windowsDriverUpdate3:               o.windowsDriverUpdate3.IPID(ctx, ipid),
		windowsDriverUpdate4:               o.windowsDriverUpdate4.IPID(ctx, ipid),
		windowsDriverUpdate5:               o.windowsDriverUpdate5.IPID(ctx, ipid),
		updateCollection:                   o.updateCollection.IPID(ctx, ipid),
		updateException:                    o.updateException.IPID(ctx, ipid),
		updateExceptionCollection:          o.updateExceptionCollection.IPID(ctx, ipid),
		searchResult:                       o.searchResult.IPID(ctx, ipid),
		updateHistoryEntry:                 o.updateHistoryEntry.IPID(ctx, ipid),
		updateHistoryEntry2:                o.updateHistoryEntry2.IPID(ctx, ipid),
		updateHistoryEntryCollection:       o.updateHistoryEntryCollection.IPID(ctx, ipid),
		updateSearcher:                     o.updateSearcher.IPID(ctx, ipid),
		updateSearcher2:                    o.updateSearcher2.IPID(ctx, ipid),
		updateSearcher3:                    o.updateSearcher3.IPID(ctx, ipid),
		updateSession:                      o.updateSession.IPID(ctx, ipid),
		updateSession2:                     o.updateSession2.IPID(ctx, ipid),
		updateSession3:                     o.updateSession3.IPID(ctx, ipid),
		updateService:                      o.updateService.IPID(ctx, ipid),
		updateService2:                     o.updateService2.IPID(ctx, ipid),
		updateServiceCollection:            o.updateServiceCollection.IPID(ctx, ipid),
		updateServiceRegistration:          o.updateServiceRegistration.IPID(ctx, ipid),
		updateServiceManager:               o.updateServiceManager.IPID(ctx, ipid),
		updateServiceManager2:              o.updateServiceManager2.IPID(ctx, ipid),
		cc:                                 o.cc,
	}
}
