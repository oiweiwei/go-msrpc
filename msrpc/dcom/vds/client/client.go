package vds

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
	ienumvdsobject "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ienumvdsobject/v0"
	ivdsadvanceddisk "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsadvanceddisk/v0"
	ivdsadvanceddisk2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsadvanceddisk2/v0"
	ivdsadvanceddisk3 "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsadvanceddisk3/v0"
	ivdsadvisesink "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsadvisesink/v0"
	ivdsasync "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsasync/v0"
	ivdscreatepartitionex "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdscreatepartitionex/v0"
	ivdsdisk "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsdisk/v0"
	ivdsdisk2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsdisk2/v0"
	ivdsdisk3 "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsdisk3/v0"
	ivdsdiskonline "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsdiskonline/v0"
	ivdsdiskpartitionmf "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsdiskpartitionmf/v0"
	ivdsdiskpartitionmf2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsdiskpartitionmf2/v0"
	ivdshbaport "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdshbaport/v0"
	ivdshwprovider "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdshwprovider/v0"
	ivdsiscsiinitiatoradapter "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsiscsiinitiatoradapter/v0"
	ivdsiscsiinitiatorportal "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsiscsiinitiatorportal/v0"
	ivdsopenvdisk "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsopenvdisk/v0"
	ivdspack "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdspack/v0"
	ivdspack2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdspack2/v0"
	ivdsprovider "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsprovider/v0"
	ivdsremovable "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsremovable/v0"
	ivdsservice "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsservice/v0"
	ivdsservicehba "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsservicehba/v0"
	ivdsserviceinitialization "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsserviceinitialization/v0"
	ivdsserviceiscsi "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsserviceiscsi/v0"
	ivdsserviceloader "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsserviceloader/v0"
	ivdsservicesan "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsservicesan/v0"
	ivdsservicesw "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsservicesw/v0"
	ivdsserviceuninstalldisk "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsserviceuninstalldisk/v0"
	ivdssubsystemimporttarget "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdssubsystemimporttarget/v0"
	ivdsswprovider "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsswprovider/v0"
	ivdsvdisk "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsvdisk/v0"
	ivdsvdprovider "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsvdprovider/v0"
	ivdsvolume "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsvolume/v0"
	ivdsvolume2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsvolume2/v0"
	ivdsvolumemf "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsvolumemf/v0"
	ivdsvolumemf2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsvolumemf2/v0"
	ivdsvolumemf3 "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsvolumemf3/v0"
	ivdsvolumeonline "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsvolumeonline/v0"
	ivdsvolumeplex "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsvolumeplex/v0"
	ivdsvolumeshrink "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds/ivdsvolumeshrink/v0"
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
	_ = ienumvdsobject.GoPackage
	_ = ivdsadvisesink.GoPackage
	_ = ivdsasync.GoPackage
	_ = ivdsserviceloader.GoPackage
	_ = ivdsservice.GoPackage
	_ = ivdsserviceinitialization.GoPackage
	_ = ivdsserviceuninstalldisk.GoPackage
	_ = ivdsservicehba.GoPackage
	_ = ivdsserviceiscsi.GoPackage
	_ = ivdsservicesan.GoPackage
	_ = ivdsservicesw.GoPackage
	_ = ivdshbaport.GoPackage
	_ = ivdsiscsiinitiatoradapter.GoPackage
	_ = ivdsiscsiinitiatorportal.GoPackage
	_ = ivdsprovider.GoPackage
	_ = ivdsswprovider.GoPackage
	_ = ivdshwprovider.GoPackage
	_ = ivdsvdprovider.GoPackage
	_ = ivdssubsystemimporttarget.GoPackage
	_ = ivdspack.GoPackage
	_ = ivdspack2.GoPackage
	_ = ivdsdisk.GoPackage
	_ = ivdsdisk2.GoPackage
	_ = ivdsdisk3.GoPackage
	_ = ivdsadvanceddisk.GoPackage
	_ = ivdsadvanceddisk2.GoPackage
	_ = ivdsadvanceddisk3.GoPackage
	_ = ivdscreatepartitionex.GoPackage
	_ = ivdsdiskonline.GoPackage
	_ = ivdsdiskpartitionmf.GoPackage
	_ = ivdsdiskpartitionmf2.GoPackage
	_ = ivdsremovable.GoPackage
	_ = ivdsvolume.GoPackage
	_ = ivdsvolume2.GoPackage
	_ = ivdsvolumemf.GoPackage
	_ = ivdsvolumemf2.GoPackage
	_ = ivdsvolumemf3.GoPackage
	_ = ivdsvolumeshrink.GoPackage
	_ = ivdsvolumeonline.GoPackage
	_ = ivdsvolumeplex.GoPackage
	_ = ivdsvdisk.GoPackage
	_ = ivdsopenvdisk.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/vds"
)

// dcom/vds client set.
type Client interface {

	// DCOM common interfaces
	RemoteUnknown() iremunknown.RemoteUnknownClient
	RemoteUnknown2() iremunknown2.RemoteUnknown2Client

	// Package specific interfaces
	EnumObject() ienumvdsobject.EnumObjectClient
	AdviseSink() ivdsadvisesink.AdviseSinkClient
	Async() ivdsasync.AsyncClient
	ServiceLoader() ivdsserviceloader.ServiceLoaderClient
	Service() ivdsservice.ServiceClient
	ServiceInitialization() ivdsserviceinitialization.ServiceInitializationClient
	ServiceUninstallDisk() ivdsserviceuninstalldisk.ServiceUninstallDiskClient
	ServiceHBA() ivdsservicehba.ServiceHBAClient
	ServiceISCSI() ivdsserviceiscsi.ServiceISCSIClient
	ServiceSAN() ivdsservicesan.ServiceSANClient
	ServiceSw() ivdsservicesw.ServiceSwClient
	HBAPort() ivdshbaport.HBAPortClient
	ISCSIInitiatorAdapter() ivdsiscsiinitiatoradapter.ISCSIInitiatorAdapterClient
	ISCSIInitiatorPortal() ivdsiscsiinitiatorportal.ISCSIInitiatorPortalClient
	Provider() ivdsprovider.ProviderClient
	SwProvider() ivdsswprovider.SwProviderClient
	HwProvider() ivdshwprovider.HwProviderClient
	VDiskProvider() ivdsvdprovider.VDiskProviderClient
	SubSystemImportTarget() ivdssubsystemimporttarget.SubSystemImportTargetClient
	Pack() ivdspack.PackClient
	Pack2() ivdspack2.Pack2Client
	Disk() ivdsdisk.DiskClient
	Disk2() ivdsdisk2.Disk2Client
	Disk3() ivdsdisk3.Disk3Client
	AdvancedDisk() ivdsadvanceddisk.AdvancedDiskClient
	AdvancedDisk2() ivdsadvanceddisk2.AdvancedDisk2Client
	AdvancedDisk3() ivdsadvanceddisk3.AdvancedDisk3Client
	CreatePartitionEx() ivdscreatepartitionex.CreatePartitionExClient
	DiskOnline() ivdsdiskonline.DiskOnlineClient
	DiskPartitionMF() ivdsdiskpartitionmf.DiskPartitionMFClient
	DiskPartitionMF2() ivdsdiskpartitionmf2.DiskPartitionMF2Client
	Removable() ivdsremovable.RemovableClient
	Volume() ivdsvolume.VolumeClient
	Volume2() ivdsvolume2.Volume2Client
	VolumeMF() ivdsvolumemf.VolumeMFClient
	VolumeMF2() ivdsvolumemf2.VolumeMF2Client
	VolumeMF3() ivdsvolumemf3.VolumeMF3Client
	VolumeShrink() ivdsvolumeshrink.VolumeShrinkClient
	VolumeOnline() ivdsvolumeonline.VolumeOnlineClient
	VolumePlex() ivdsvolumeplex.VolumePlexClient
	VDisk() ivdsvdisk.VDiskClient
	OpenVDisk() ivdsopenvdisk.OpenVDiskClient
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

	enumObject            ienumvdsobject.EnumObjectClient
	adviseSink            ivdsadvisesink.AdviseSinkClient
	async                 ivdsasync.AsyncClient
	serviceLoader         ivdsserviceloader.ServiceLoaderClient
	service               ivdsservice.ServiceClient
	serviceInitialization ivdsserviceinitialization.ServiceInitializationClient
	serviceUninstallDisk  ivdsserviceuninstalldisk.ServiceUninstallDiskClient
	serviceHBA            ivdsservicehba.ServiceHBAClient
	serviceISCSI          ivdsserviceiscsi.ServiceISCSIClient
	serviceSAN            ivdsservicesan.ServiceSANClient
	serviceSw             ivdsservicesw.ServiceSwClient
	hbaPort               ivdshbaport.HBAPortClient
	iscsiInitiatorAdapter ivdsiscsiinitiatoradapter.ISCSIInitiatorAdapterClient
	iscsiInitiatorPortal  ivdsiscsiinitiatorportal.ISCSIInitiatorPortalClient
	provider              ivdsprovider.ProviderClient
	swProvider            ivdsswprovider.SwProviderClient
	hwProvider            ivdshwprovider.HwProviderClient
	vDiskProvider         ivdsvdprovider.VDiskProviderClient
	subSystemImportTarget ivdssubsystemimporttarget.SubSystemImportTargetClient
	pack                  ivdspack.PackClient
	pack2                 ivdspack2.Pack2Client
	disk                  ivdsdisk.DiskClient
	disk2                 ivdsdisk2.Disk2Client
	disk3                 ivdsdisk3.Disk3Client
	advancedDisk          ivdsadvanceddisk.AdvancedDiskClient
	advancedDisk2         ivdsadvanceddisk2.AdvancedDisk2Client
	advancedDisk3         ivdsadvanceddisk3.AdvancedDisk3Client
	createPartitionEx     ivdscreatepartitionex.CreatePartitionExClient
	diskOnline            ivdsdiskonline.DiskOnlineClient
	diskPartitionMF       ivdsdiskpartitionmf.DiskPartitionMFClient
	diskPartitionMF2      ivdsdiskpartitionmf2.DiskPartitionMF2Client
	removable             ivdsremovable.RemovableClient
	volume                ivdsvolume.VolumeClient
	volume2               ivdsvolume2.Volume2Client
	volumeMF              ivdsvolumemf.VolumeMFClient
	volumeMF2             ivdsvolumemf2.VolumeMF2Client
	volumeMF3             ivdsvolumemf3.VolumeMF3Client
	volumeShrink          ivdsvolumeshrink.VolumeShrinkClient
	volumeOnline          ivdsvolumeonline.VolumeOnlineClient
	volumePlex            ivdsvolumeplex.VolumePlexClient
	vDisk                 ivdsvdisk.VDiskClient
	openVDisk             ivdsopenvdisk.OpenVDiskClient
}

func (o *xxx_DefaultClient) RemoteUnknown() iremunknown.RemoteUnknownClient {
	return o.dcomClient.RemoteUnknown()
}

func (o *xxx_DefaultClient) RemoteUnknown2() iremunknown2.RemoteUnknown2Client {
	return o.dcomClient.RemoteUnknown2()
}

func (o *xxx_DefaultClient) EnumObject() ienumvdsobject.EnumObjectClient {
	return o.enumObject
}

func (o *xxx_DefaultClient) AdviseSink() ivdsadvisesink.AdviseSinkClient {
	return o.adviseSink
}

func (o *xxx_DefaultClient) Async() ivdsasync.AsyncClient {
	return o.async
}

func (o *xxx_DefaultClient) ServiceLoader() ivdsserviceloader.ServiceLoaderClient {
	return o.serviceLoader
}

func (o *xxx_DefaultClient) Service() ivdsservice.ServiceClient {
	return o.service
}

func (o *xxx_DefaultClient) ServiceInitialization() ivdsserviceinitialization.ServiceInitializationClient {
	return o.serviceInitialization
}

func (o *xxx_DefaultClient) ServiceUninstallDisk() ivdsserviceuninstalldisk.ServiceUninstallDiskClient {
	return o.serviceUninstallDisk
}

func (o *xxx_DefaultClient) ServiceHBA() ivdsservicehba.ServiceHBAClient {
	return o.serviceHBA
}

func (o *xxx_DefaultClient) ServiceISCSI() ivdsserviceiscsi.ServiceISCSIClient {
	return o.serviceISCSI
}

func (o *xxx_DefaultClient) ServiceSAN() ivdsservicesan.ServiceSANClient {
	return o.serviceSAN
}

func (o *xxx_DefaultClient) ServiceSw() ivdsservicesw.ServiceSwClient {
	return o.serviceSw
}

func (o *xxx_DefaultClient) HBAPort() ivdshbaport.HBAPortClient {
	return o.hbaPort
}

func (o *xxx_DefaultClient) ISCSIInitiatorAdapter() ivdsiscsiinitiatoradapter.ISCSIInitiatorAdapterClient {
	return o.iscsiInitiatorAdapter
}

func (o *xxx_DefaultClient) ISCSIInitiatorPortal() ivdsiscsiinitiatorportal.ISCSIInitiatorPortalClient {
	return o.iscsiInitiatorPortal
}

func (o *xxx_DefaultClient) Provider() ivdsprovider.ProviderClient {
	return o.provider
}

func (o *xxx_DefaultClient) SwProvider() ivdsswprovider.SwProviderClient {
	return o.swProvider
}

func (o *xxx_DefaultClient) HwProvider() ivdshwprovider.HwProviderClient {
	return o.hwProvider
}

func (o *xxx_DefaultClient) VDiskProvider() ivdsvdprovider.VDiskProviderClient {
	return o.vDiskProvider
}

func (o *xxx_DefaultClient) SubSystemImportTarget() ivdssubsystemimporttarget.SubSystemImportTargetClient {
	return o.subSystemImportTarget
}

func (o *xxx_DefaultClient) Pack() ivdspack.PackClient {
	return o.pack
}

func (o *xxx_DefaultClient) Pack2() ivdspack2.Pack2Client {
	return o.pack2
}

func (o *xxx_DefaultClient) Disk() ivdsdisk.DiskClient {
	return o.disk
}

func (o *xxx_DefaultClient) Disk2() ivdsdisk2.Disk2Client {
	return o.disk2
}

func (o *xxx_DefaultClient) Disk3() ivdsdisk3.Disk3Client {
	return o.disk3
}

func (o *xxx_DefaultClient) AdvancedDisk() ivdsadvanceddisk.AdvancedDiskClient {
	return o.advancedDisk
}

func (o *xxx_DefaultClient) AdvancedDisk2() ivdsadvanceddisk2.AdvancedDisk2Client {
	return o.advancedDisk2
}

func (o *xxx_DefaultClient) AdvancedDisk3() ivdsadvanceddisk3.AdvancedDisk3Client {
	return o.advancedDisk3
}

func (o *xxx_DefaultClient) CreatePartitionEx() ivdscreatepartitionex.CreatePartitionExClient {
	return o.createPartitionEx
}

func (o *xxx_DefaultClient) DiskOnline() ivdsdiskonline.DiskOnlineClient {
	return o.diskOnline
}

func (o *xxx_DefaultClient) DiskPartitionMF() ivdsdiskpartitionmf.DiskPartitionMFClient {
	return o.diskPartitionMF
}

func (o *xxx_DefaultClient) DiskPartitionMF2() ivdsdiskpartitionmf2.DiskPartitionMF2Client {
	return o.diskPartitionMF2
}

func (o *xxx_DefaultClient) Removable() ivdsremovable.RemovableClient {
	return o.removable
}

func (o *xxx_DefaultClient) Volume() ivdsvolume.VolumeClient {
	return o.volume
}

func (o *xxx_DefaultClient) Volume2() ivdsvolume2.Volume2Client {
	return o.volume2
}

func (o *xxx_DefaultClient) VolumeMF() ivdsvolumemf.VolumeMFClient {
	return o.volumeMF
}

func (o *xxx_DefaultClient) VolumeMF2() ivdsvolumemf2.VolumeMF2Client {
	return o.volumeMF2
}

func (o *xxx_DefaultClient) VolumeMF3() ivdsvolumemf3.VolumeMF3Client {
	return o.volumeMF3
}

func (o *xxx_DefaultClient) VolumeShrink() ivdsvolumeshrink.VolumeShrinkClient {
	return o.volumeShrink
}

func (o *xxx_DefaultClient) VolumeOnline() ivdsvolumeonline.VolumeOnlineClient {
	return o.volumeOnline
}

func (o *xxx_DefaultClient) VolumePlex() ivdsvolumeplex.VolumePlexClient {
	return o.volumePlex
}

func (o *xxx_DefaultClient) VDisk() ivdsvdisk.VDiskClient {
	return o.vDisk
}

func (o *xxx_DefaultClient) OpenVDisk() ivdsopenvdisk.OpenVDiskClient {
	return o.openVDisk
}

func NewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Client, error) {

	opts = append(opts,
		dcerpc.WithAbstractSyntax(ienumvdsobject.EnumObjectSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsadvisesink.AdviseSinkSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsasync.AsyncSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsserviceloader.ServiceLoaderSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsservice.ServiceSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsserviceinitialization.ServiceInitializationSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsserviceuninstalldisk.ServiceUninstallDiskSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsservicehba.ServiceHBASyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsserviceiscsi.ServiceISCSISyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsservicesan.ServiceSANSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsservicesw.ServiceSwSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdshbaport.HBAPortSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsiscsiinitiatoradapter.ISCSIInitiatorAdapterSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsiscsiinitiatorportal.ISCSIInitiatorPortalSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsprovider.ProviderSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsswprovider.SwProviderSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdshwprovider.HwProviderSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsvdprovider.VDiskProviderSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdssubsystemimporttarget.SubSystemImportTargetSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdspack.PackSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdspack2.Pack2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsdisk.DiskSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsdisk2.Disk2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsdisk3.Disk3SyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsadvanceddisk.AdvancedDiskSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsadvanceddisk2.AdvancedDisk2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsadvanceddisk3.AdvancedDisk3SyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdscreatepartitionex.CreatePartitionExSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsdiskonline.DiskOnlineSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsdiskpartitionmf.DiskPartitionMFSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsdiskpartitionmf2.DiskPartitionMF2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsremovable.RemovableSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsvolume.VolumeSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsvolume2.Volume2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsvolumemf.VolumeMFSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsvolumemf2.VolumeMF2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsvolumemf3.VolumeMF3SyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsvolumeshrink.VolumeShrinkSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsvolumeonline.VolumeOnlineSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsvolumeplex.VolumePlexSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsvdisk.VDiskSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivdsopenvdisk.OpenVDiskSyntaxV0_0),
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

	enumObjectSubConn, err := sub.SubConn(ctx, ienumvdsobject.EnumObjectSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		enumObjectSubConn = sub
	}

	o.enumObject, err = ienumvdsobject.NewEnumObjectClient(ctx, enumObjectSubConn, append(opts, dcerpc.WithNoBind(enumObjectSubConn))...)

	adviseSinkSubConn, err := sub.SubConn(ctx, ivdsadvisesink.AdviseSinkSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		adviseSinkSubConn = sub
	}

	o.adviseSink, err = ivdsadvisesink.NewAdviseSinkClient(ctx, adviseSinkSubConn, append(opts, dcerpc.WithNoBind(adviseSinkSubConn))...)

	asyncSubConn, err := sub.SubConn(ctx, ivdsasync.AsyncSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		asyncSubConn = sub
	}

	o.async, err = ivdsasync.NewAsyncClient(ctx, asyncSubConn, append(opts, dcerpc.WithNoBind(asyncSubConn))...)

	serviceLoaderSubConn, err := sub.SubConn(ctx, ivdsserviceloader.ServiceLoaderSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		serviceLoaderSubConn = sub
	}

	o.serviceLoader, err = ivdsserviceloader.NewServiceLoaderClient(ctx, serviceLoaderSubConn, append(opts, dcerpc.WithNoBind(serviceLoaderSubConn))...)

	serviceSubConn, err := sub.SubConn(ctx, ivdsservice.ServiceSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		serviceSubConn = sub
	}

	o.service, err = ivdsservice.NewServiceClient(ctx, serviceSubConn, append(opts, dcerpc.WithNoBind(serviceSubConn))...)

	serviceInitializationSubConn, err := sub.SubConn(ctx, ivdsserviceinitialization.ServiceInitializationSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		serviceInitializationSubConn = sub
	}

	o.serviceInitialization, err = ivdsserviceinitialization.NewServiceInitializationClient(ctx, serviceInitializationSubConn, append(opts, dcerpc.WithNoBind(serviceInitializationSubConn))...)

	serviceUninstallDiskSubConn, err := sub.SubConn(ctx, ivdsserviceuninstalldisk.ServiceUninstallDiskSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		serviceUninstallDiskSubConn = sub
	}

	o.serviceUninstallDisk, err = ivdsserviceuninstalldisk.NewServiceUninstallDiskClient(ctx, serviceUninstallDiskSubConn, append(opts, dcerpc.WithNoBind(serviceUninstallDiskSubConn))...)

	serviceHBASubConn, err := sub.SubConn(ctx, ivdsservicehba.ServiceHBASyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		serviceHBASubConn = sub
	}

	o.serviceHBA, err = ivdsservicehba.NewServiceHBAClient(ctx, serviceHBASubConn, append(opts, dcerpc.WithNoBind(serviceHBASubConn))...)

	serviceISCSISubConn, err := sub.SubConn(ctx, ivdsserviceiscsi.ServiceISCSISyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		serviceISCSISubConn = sub
	}

	o.serviceISCSI, err = ivdsserviceiscsi.NewServiceISCSIClient(ctx, serviceISCSISubConn, append(opts, dcerpc.WithNoBind(serviceISCSISubConn))...)

	serviceSANSubConn, err := sub.SubConn(ctx, ivdsservicesan.ServiceSANSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		serviceSANSubConn = sub
	}

	o.serviceSAN, err = ivdsservicesan.NewServiceSANClient(ctx, serviceSANSubConn, append(opts, dcerpc.WithNoBind(serviceSANSubConn))...)

	serviceSwSubConn, err := sub.SubConn(ctx, ivdsservicesw.ServiceSwSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		serviceSwSubConn = sub
	}

	o.serviceSw, err = ivdsservicesw.NewServiceSwClient(ctx, serviceSwSubConn, append(opts, dcerpc.WithNoBind(serviceSwSubConn))...)

	hbaPortSubConn, err := sub.SubConn(ctx, ivdshbaport.HBAPortSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		hbaPortSubConn = sub
	}

	o.hbaPort, err = ivdshbaport.NewHBAPortClient(ctx, hbaPortSubConn, append(opts, dcerpc.WithNoBind(hbaPortSubConn))...)

	iscsiInitiatorAdapterSubConn, err := sub.SubConn(ctx, ivdsiscsiinitiatoradapter.ISCSIInitiatorAdapterSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		iscsiInitiatorAdapterSubConn = sub
	}

	o.iscsiInitiatorAdapter, err = ivdsiscsiinitiatoradapter.NewISCSIInitiatorAdapterClient(ctx, iscsiInitiatorAdapterSubConn, append(opts, dcerpc.WithNoBind(iscsiInitiatorAdapterSubConn))...)

	iscsiInitiatorPortalSubConn, err := sub.SubConn(ctx, ivdsiscsiinitiatorportal.ISCSIInitiatorPortalSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		iscsiInitiatorPortalSubConn = sub
	}

	o.iscsiInitiatorPortal, err = ivdsiscsiinitiatorportal.NewISCSIInitiatorPortalClient(ctx, iscsiInitiatorPortalSubConn, append(opts, dcerpc.WithNoBind(iscsiInitiatorPortalSubConn))...)

	providerSubConn, err := sub.SubConn(ctx, ivdsprovider.ProviderSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		providerSubConn = sub
	}

	o.provider, err = ivdsprovider.NewProviderClient(ctx, providerSubConn, append(opts, dcerpc.WithNoBind(providerSubConn))...)

	swProviderSubConn, err := sub.SubConn(ctx, ivdsswprovider.SwProviderSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		swProviderSubConn = sub
	}

	o.swProvider, err = ivdsswprovider.NewSwProviderClient(ctx, swProviderSubConn, append(opts, dcerpc.WithNoBind(swProviderSubConn))...)

	hwProviderSubConn, err := sub.SubConn(ctx, ivdshwprovider.HwProviderSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		hwProviderSubConn = sub
	}

	o.hwProvider, err = ivdshwprovider.NewHwProviderClient(ctx, hwProviderSubConn, append(opts, dcerpc.WithNoBind(hwProviderSubConn))...)

	vDiskProviderSubConn, err := sub.SubConn(ctx, ivdsvdprovider.VDiskProviderSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		vDiskProviderSubConn = sub
	}

	o.vDiskProvider, err = ivdsvdprovider.NewVDiskProviderClient(ctx, vDiskProviderSubConn, append(opts, dcerpc.WithNoBind(vDiskProviderSubConn))...)

	subSystemImportTargetSubConn, err := sub.SubConn(ctx, ivdssubsystemimporttarget.SubSystemImportTargetSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		subSystemImportTargetSubConn = sub
	}

	o.subSystemImportTarget, err = ivdssubsystemimporttarget.NewSubSystemImportTargetClient(ctx, subSystemImportTargetSubConn, append(opts, dcerpc.WithNoBind(subSystemImportTargetSubConn))...)

	packSubConn, err := sub.SubConn(ctx, ivdspack.PackSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		packSubConn = sub
	}

	o.pack, err = ivdspack.NewPackClient(ctx, packSubConn, append(opts, dcerpc.WithNoBind(packSubConn))...)

	pack2SubConn, err := sub.SubConn(ctx, ivdspack2.Pack2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		pack2SubConn = sub
	}

	o.pack2, err = ivdspack2.NewPack2Client(ctx, pack2SubConn, append(opts, dcerpc.WithNoBind(pack2SubConn))...)

	diskSubConn, err := sub.SubConn(ctx, ivdsdisk.DiskSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		diskSubConn = sub
	}

	o.disk, err = ivdsdisk.NewDiskClient(ctx, diskSubConn, append(opts, dcerpc.WithNoBind(diskSubConn))...)

	disk2SubConn, err := sub.SubConn(ctx, ivdsdisk2.Disk2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		disk2SubConn = sub
	}

	o.disk2, err = ivdsdisk2.NewDisk2Client(ctx, disk2SubConn, append(opts, dcerpc.WithNoBind(disk2SubConn))...)

	disk3SubConn, err := sub.SubConn(ctx, ivdsdisk3.Disk3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		disk3SubConn = sub
	}

	o.disk3, err = ivdsdisk3.NewDisk3Client(ctx, disk3SubConn, append(opts, dcerpc.WithNoBind(disk3SubConn))...)

	advancedDiskSubConn, err := sub.SubConn(ctx, ivdsadvanceddisk.AdvancedDiskSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		advancedDiskSubConn = sub
	}

	o.advancedDisk, err = ivdsadvanceddisk.NewAdvancedDiskClient(ctx, advancedDiskSubConn, append(opts, dcerpc.WithNoBind(advancedDiskSubConn))...)

	advancedDisk2SubConn, err := sub.SubConn(ctx, ivdsadvanceddisk2.AdvancedDisk2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		advancedDisk2SubConn = sub
	}

	o.advancedDisk2, err = ivdsadvanceddisk2.NewAdvancedDisk2Client(ctx, advancedDisk2SubConn, append(opts, dcerpc.WithNoBind(advancedDisk2SubConn))...)

	advancedDisk3SubConn, err := sub.SubConn(ctx, ivdsadvanceddisk3.AdvancedDisk3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		advancedDisk3SubConn = sub
	}

	o.advancedDisk3, err = ivdsadvanceddisk3.NewAdvancedDisk3Client(ctx, advancedDisk3SubConn, append(opts, dcerpc.WithNoBind(advancedDisk3SubConn))...)

	createPartitionExSubConn, err := sub.SubConn(ctx, ivdscreatepartitionex.CreatePartitionExSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		createPartitionExSubConn = sub
	}

	o.createPartitionEx, err = ivdscreatepartitionex.NewCreatePartitionExClient(ctx, createPartitionExSubConn, append(opts, dcerpc.WithNoBind(createPartitionExSubConn))...)

	diskOnlineSubConn, err := sub.SubConn(ctx, ivdsdiskonline.DiskOnlineSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		diskOnlineSubConn = sub
	}

	o.diskOnline, err = ivdsdiskonline.NewDiskOnlineClient(ctx, diskOnlineSubConn, append(opts, dcerpc.WithNoBind(diskOnlineSubConn))...)

	diskPartitionMFSubConn, err := sub.SubConn(ctx, ivdsdiskpartitionmf.DiskPartitionMFSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		diskPartitionMFSubConn = sub
	}

	o.diskPartitionMF, err = ivdsdiskpartitionmf.NewDiskPartitionMFClient(ctx, diskPartitionMFSubConn, append(opts, dcerpc.WithNoBind(diskPartitionMFSubConn))...)

	diskPartitionMF2SubConn, err := sub.SubConn(ctx, ivdsdiskpartitionmf2.DiskPartitionMF2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		diskPartitionMF2SubConn = sub
	}

	o.diskPartitionMF2, err = ivdsdiskpartitionmf2.NewDiskPartitionMF2Client(ctx, diskPartitionMF2SubConn, append(opts, dcerpc.WithNoBind(diskPartitionMF2SubConn))...)

	removableSubConn, err := sub.SubConn(ctx, ivdsremovable.RemovableSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		removableSubConn = sub
	}

	o.removable, err = ivdsremovable.NewRemovableClient(ctx, removableSubConn, append(opts, dcerpc.WithNoBind(removableSubConn))...)

	volumeSubConn, err := sub.SubConn(ctx, ivdsvolume.VolumeSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		volumeSubConn = sub
	}

	o.volume, err = ivdsvolume.NewVolumeClient(ctx, volumeSubConn, append(opts, dcerpc.WithNoBind(volumeSubConn))...)

	volume2SubConn, err := sub.SubConn(ctx, ivdsvolume2.Volume2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		volume2SubConn = sub
	}

	o.volume2, err = ivdsvolume2.NewVolume2Client(ctx, volume2SubConn, append(opts, dcerpc.WithNoBind(volume2SubConn))...)

	volumeMFSubConn, err := sub.SubConn(ctx, ivdsvolumemf.VolumeMFSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		volumeMFSubConn = sub
	}

	o.volumeMF, err = ivdsvolumemf.NewVolumeMFClient(ctx, volumeMFSubConn, append(opts, dcerpc.WithNoBind(volumeMFSubConn))...)

	volumeMF2SubConn, err := sub.SubConn(ctx, ivdsvolumemf2.VolumeMF2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		volumeMF2SubConn = sub
	}

	o.volumeMF2, err = ivdsvolumemf2.NewVolumeMF2Client(ctx, volumeMF2SubConn, append(opts, dcerpc.WithNoBind(volumeMF2SubConn))...)

	volumeMF3SubConn, err := sub.SubConn(ctx, ivdsvolumemf3.VolumeMF3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		volumeMF3SubConn = sub
	}

	o.volumeMF3, err = ivdsvolumemf3.NewVolumeMF3Client(ctx, volumeMF3SubConn, append(opts, dcerpc.WithNoBind(volumeMF3SubConn))...)

	volumeShrinkSubConn, err := sub.SubConn(ctx, ivdsvolumeshrink.VolumeShrinkSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		volumeShrinkSubConn = sub
	}

	o.volumeShrink, err = ivdsvolumeshrink.NewVolumeShrinkClient(ctx, volumeShrinkSubConn, append(opts, dcerpc.WithNoBind(volumeShrinkSubConn))...)

	volumeOnlineSubConn, err := sub.SubConn(ctx, ivdsvolumeonline.VolumeOnlineSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		volumeOnlineSubConn = sub
	}

	o.volumeOnline, err = ivdsvolumeonline.NewVolumeOnlineClient(ctx, volumeOnlineSubConn, append(opts, dcerpc.WithNoBind(volumeOnlineSubConn))...)

	volumePlexSubConn, err := sub.SubConn(ctx, ivdsvolumeplex.VolumePlexSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		volumePlexSubConn = sub
	}

	o.volumePlex, err = ivdsvolumeplex.NewVolumePlexClient(ctx, volumePlexSubConn, append(opts, dcerpc.WithNoBind(volumePlexSubConn))...)

	vDiskSubConn, err := sub.SubConn(ctx, ivdsvdisk.VDiskSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		vDiskSubConn = sub
	}

	o.vDisk, err = ivdsvdisk.NewVDiskClient(ctx, vDiskSubConn, append(opts, dcerpc.WithNoBind(vDiskSubConn))...)

	openVDiskSubConn, err := sub.SubConn(ctx, ivdsopenvdisk.OpenVDiskSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		openVDiskSubConn = sub
	}

	o.openVDisk, err = ivdsopenvdisk.NewOpenVDiskClient(ctx, openVDiskSubConn, append(opts, dcerpc.WithNoBind(openVDiskSubConn))...)
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
		dcomClient:            o.dcomClient.IPID(ctx, ipid),
		enumObject:            o.enumObject.IPID(ctx, ipid),
		adviseSink:            o.adviseSink.IPID(ctx, ipid),
		async:                 o.async.IPID(ctx, ipid),
		serviceLoader:         o.serviceLoader.IPID(ctx, ipid),
		service:               o.service.IPID(ctx, ipid),
		serviceInitialization: o.serviceInitialization.IPID(ctx, ipid),
		serviceUninstallDisk:  o.serviceUninstallDisk.IPID(ctx, ipid),
		serviceHBA:            o.serviceHBA.IPID(ctx, ipid),
		serviceISCSI:          o.serviceISCSI.IPID(ctx, ipid),
		serviceSAN:            o.serviceSAN.IPID(ctx, ipid),
		serviceSw:             o.serviceSw.IPID(ctx, ipid),
		hbaPort:               o.hbaPort.IPID(ctx, ipid),
		iscsiInitiatorAdapter: o.iscsiInitiatorAdapter.IPID(ctx, ipid),
		iscsiInitiatorPortal:  o.iscsiInitiatorPortal.IPID(ctx, ipid),
		provider:              o.provider.IPID(ctx, ipid),
		swProvider:            o.swProvider.IPID(ctx, ipid),
		hwProvider:            o.hwProvider.IPID(ctx, ipid),
		vDiskProvider:         o.vDiskProvider.IPID(ctx, ipid),
		subSystemImportTarget: o.subSystemImportTarget.IPID(ctx, ipid),
		pack:                  o.pack.IPID(ctx, ipid),
		pack2:                 o.pack2.IPID(ctx, ipid),
		disk:                  o.disk.IPID(ctx, ipid),
		disk2:                 o.disk2.IPID(ctx, ipid),
		disk3:                 o.disk3.IPID(ctx, ipid),
		advancedDisk:          o.advancedDisk.IPID(ctx, ipid),
		advancedDisk2:         o.advancedDisk2.IPID(ctx, ipid),
		advancedDisk3:         o.advancedDisk3.IPID(ctx, ipid),
		createPartitionEx:     o.createPartitionEx.IPID(ctx, ipid),
		diskOnline:            o.diskOnline.IPID(ctx, ipid),
		diskPartitionMF:       o.diskPartitionMF.IPID(ctx, ipid),
		diskPartitionMF2:      o.diskPartitionMF2.IPID(ctx, ipid),
		removable:             o.removable.IPID(ctx, ipid),
		volume:                o.volume.IPID(ctx, ipid),
		volume2:               o.volume2.IPID(ctx, ipid),
		volumeMF:              o.volumeMF.IPID(ctx, ipid),
		volumeMF2:             o.volumeMF2.IPID(ctx, ipid),
		volumeMF3:             o.volumeMF3.IPID(ctx, ipid),
		volumeShrink:          o.volumeShrink.IPID(ctx, ipid),
		volumeOnline:          o.volumeOnline.IPID(ctx, ipid),
		volumePlex:            o.volumePlex.IPID(ctx, ipid),
		vDisk:                 o.vDisk.IPID(ctx, ipid),
		openVDisk:             o.openVDisk.IPID(ctx, ipid),
		cc:                    o.cc,
	}
}
