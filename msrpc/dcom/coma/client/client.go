package coma

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
	ialternatelaunch "github.com/oiweiwei/go-msrpc/msrpc/dcom/coma/ialternatelaunch/v0"
	icapabilitysupport "github.com/oiweiwei/go-msrpc/msrpc/dcom/coma/icapabilitysupport/v0"
	icatalog64bitsupport "github.com/oiweiwei/go-msrpc/msrpc/dcom/coma/icatalog64bitsupport/v0"
	icatalogsession "github.com/oiweiwei/go-msrpc/msrpc/dcom/coma/icatalogsession/v0"
	icatalogtableinfo "github.com/oiweiwei/go-msrpc/msrpc/dcom/coma/icatalogtableinfo/v0"
	icatalogtableread "github.com/oiweiwei/go-msrpc/msrpc/dcom/coma/icatalogtableread/v0"
	icatalogtablewrite "github.com/oiweiwei/go-msrpc/msrpc/dcom/coma/icatalogtablewrite/v0"
	icatalogutils "github.com/oiweiwei/go-msrpc/msrpc/dcom/coma/icatalogutils/v0"
	icatalogutils2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/coma/icatalogutils2/v0"
	icontainercontrol "github.com/oiweiwei/go-msrpc/msrpc/dcom/coma/icontainercontrol/v0"
	icontainercontrol2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/coma/icontainercontrol2/v0"
	iexport "github.com/oiweiwei/go-msrpc/msrpc/dcom/coma/iexport/v0"
	iexport2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/coma/iexport2/v0"
	iimport "github.com/oiweiwei/go-msrpc/msrpc/dcom/coma/iimport/v0"
	iimport2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/coma/iimport2/v0"
	iregister "github.com/oiweiwei/go-msrpc/msrpc/dcom/coma/iregister/v0"
	iregister2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/coma/iregister2/v0"
	ireplicationutil "github.com/oiweiwei/go-msrpc/msrpc/dcom/coma/ireplicationutil/v0"
	iremunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iremunknown/v0"
	iremunknown2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/iremunknown2/v0"
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
	_ = icatalogsession.GoPackage
	_ = icatalog64bitsupport.GoPackage
	_ = icatalogtableinfo.GoPackage
	_ = icatalogtableread.GoPackage
	_ = icatalogtablewrite.GoPackage
	_ = iregister.GoPackage
	_ = iregister2.GoPackage
	_ = iimport.GoPackage
	_ = iimport2.GoPackage
	_ = iexport.GoPackage
	_ = iexport2.GoPackage
	_ = ialternatelaunch.GoPackage
	_ = icatalogutils.GoPackage
	_ = icatalogutils2.GoPackage
	_ = icapabilitysupport.GoPackage
	_ = icontainercontrol.GoPackage
	_ = icontainercontrol2.GoPackage
	_ = ireplicationutil.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/coma"
)

// dcom/coma client set.
type Client interface {

	// DCOM common interfaces
	RemoteUnknown() iremunknown.RemoteUnknownClient
	RemoteUnknown2() iremunknown2.RemoteUnknown2Client

	// Package specific interfaces
	CatalogSession() icatalogsession.CatalogSessionClient
	Catalog64BitSupport() icatalog64bitsupport.Catalog64BitSupportClient
	CatalogTableInfo() icatalogtableinfo.CatalogTableInfoClient
	CatalogTableRead() icatalogtableread.CatalogTableReadClient
	CatalogTableWrite() icatalogtablewrite.CatalogTableWriteClient
	Register() iregister.RegisterClient
	Register2() iregister2.Register2Client
	Import() iimport.ImportClient
	Import2() iimport2.Import2Client
	Export() iexport.ExportClient
	Export2() iexport2.Export2Client
	AlternateLaunch() ialternatelaunch.AlternateLaunchClient
	CatalogUtils() icatalogutils.CatalogUtilsClient
	CatalogUtils2() icatalogutils2.CatalogUtils2Client
	CapabilitySupport() icapabilitysupport.CapabilitySupportClient
	ContainerControl() icontainercontrol.ContainerControlClient
	ContainerControl2() icontainercontrol2.ContainerControl2Client
	ReplicationUtil() ireplicationutil.ReplicationUtilClient
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

	catalogSession      icatalogsession.CatalogSessionClient
	catalog64BitSupport icatalog64bitsupport.Catalog64BitSupportClient
	catalogTableInfo    icatalogtableinfo.CatalogTableInfoClient
	catalogTableRead    icatalogtableread.CatalogTableReadClient
	catalogTableWrite   icatalogtablewrite.CatalogTableWriteClient
	register            iregister.RegisterClient
	register2           iregister2.Register2Client
	import1             iimport.ImportClient
	import2             iimport2.Import2Client
	export              iexport.ExportClient
	export2             iexport2.Export2Client
	alternateLaunch     ialternatelaunch.AlternateLaunchClient
	catalogUtils        icatalogutils.CatalogUtilsClient
	catalogUtils2       icatalogutils2.CatalogUtils2Client
	capabilitySupport   icapabilitysupport.CapabilitySupportClient
	containerControl    icontainercontrol.ContainerControlClient
	containerControl2   icontainercontrol2.ContainerControl2Client
	replicationUtil     ireplicationutil.ReplicationUtilClient
}

func (o *xxx_DefaultClient) RemoteUnknown() iremunknown.RemoteUnknownClient {
	return o.dcomClient.RemoteUnknown()
}

func (o *xxx_DefaultClient) RemoteUnknown2() iremunknown2.RemoteUnknown2Client {
	return o.dcomClient.RemoteUnknown2()
}

func (o *xxx_DefaultClient) CatalogSession() icatalogsession.CatalogSessionClient {
	return o.catalogSession
}

func (o *xxx_DefaultClient) Catalog64BitSupport() icatalog64bitsupport.Catalog64BitSupportClient {
	return o.catalog64BitSupport
}

func (o *xxx_DefaultClient) CatalogTableInfo() icatalogtableinfo.CatalogTableInfoClient {
	return o.catalogTableInfo
}

func (o *xxx_DefaultClient) CatalogTableRead() icatalogtableread.CatalogTableReadClient {
	return o.catalogTableRead
}

func (o *xxx_DefaultClient) CatalogTableWrite() icatalogtablewrite.CatalogTableWriteClient {
	return o.catalogTableWrite
}

func (o *xxx_DefaultClient) Register() iregister.RegisterClient {
	return o.register
}

func (o *xxx_DefaultClient) Register2() iregister2.Register2Client {
	return o.register2
}

func (o *xxx_DefaultClient) Import() iimport.ImportClient {
	return o.import1
}

func (o *xxx_DefaultClient) Import2() iimport2.Import2Client {
	return o.import2
}

func (o *xxx_DefaultClient) Export() iexport.ExportClient {
	return o.export
}

func (o *xxx_DefaultClient) Export2() iexport2.Export2Client {
	return o.export2
}

func (o *xxx_DefaultClient) AlternateLaunch() ialternatelaunch.AlternateLaunchClient {
	return o.alternateLaunch
}

func (o *xxx_DefaultClient) CatalogUtils() icatalogutils.CatalogUtilsClient {
	return o.catalogUtils
}

func (o *xxx_DefaultClient) CatalogUtils2() icatalogutils2.CatalogUtils2Client {
	return o.catalogUtils2
}

func (o *xxx_DefaultClient) CapabilitySupport() icapabilitysupport.CapabilitySupportClient {
	return o.capabilitySupport
}

func (o *xxx_DefaultClient) ContainerControl() icontainercontrol.ContainerControlClient {
	return o.containerControl
}

func (o *xxx_DefaultClient) ContainerControl2() icontainercontrol2.ContainerControl2Client {
	return o.containerControl2
}

func (o *xxx_DefaultClient) ReplicationUtil() ireplicationutil.ReplicationUtilClient {
	return o.replicationUtil
}

func NewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Client, error) {

	opts = append(opts,
		dcerpc.WithAbstractSyntax(icatalogsession.CatalogSessionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(icatalog64bitsupport.Catalog64BitSupportSyntaxV0_0),
		dcerpc.WithAbstractSyntax(icatalogtableinfo.CatalogTableInfoSyntaxV0_0),
		dcerpc.WithAbstractSyntax(icatalogtableread.CatalogTableReadSyntaxV0_0),
		dcerpc.WithAbstractSyntax(icatalogtablewrite.CatalogTableWriteSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iregister.RegisterSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iregister2.Register2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(iimport.ImportSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iimport2.Import2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(iexport.ExportSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iexport2.Export2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(ialternatelaunch.AlternateLaunchSyntaxV0_0),
		dcerpc.WithAbstractSyntax(icatalogutils.CatalogUtilsSyntaxV0_0),
		dcerpc.WithAbstractSyntax(icatalogutils2.CatalogUtils2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(icapabilitysupport.CapabilitySupportSyntaxV0_0),
		dcerpc.WithAbstractSyntax(icontainercontrol.ContainerControlSyntaxV0_0),
		dcerpc.WithAbstractSyntax(icontainercontrol2.ContainerControl2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(ireplicationutil.ReplicationUtilSyntaxV0_0),
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

	catalogSessionSubConn, err := sub.SubConn(ctx, icatalogsession.CatalogSessionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		catalogSessionSubConn = sub
	}

	o.catalogSession, err = icatalogsession.NewCatalogSessionClient(ctx, catalogSessionSubConn, append(opts, dcerpc.WithNoBind(catalogSessionSubConn))...)

	catalog64BitSupportSubConn, err := sub.SubConn(ctx, icatalog64bitsupport.Catalog64BitSupportSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		catalog64BitSupportSubConn = sub
	}

	o.catalog64BitSupport, err = icatalog64bitsupport.NewCatalog64BitSupportClient(ctx, catalog64BitSupportSubConn, append(opts, dcerpc.WithNoBind(catalog64BitSupportSubConn))...)

	catalogTableInfoSubConn, err := sub.SubConn(ctx, icatalogtableinfo.CatalogTableInfoSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		catalogTableInfoSubConn = sub
	}

	o.catalogTableInfo, err = icatalogtableinfo.NewCatalogTableInfoClient(ctx, catalogTableInfoSubConn, append(opts, dcerpc.WithNoBind(catalogTableInfoSubConn))...)

	catalogTableReadSubConn, err := sub.SubConn(ctx, icatalogtableread.CatalogTableReadSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		catalogTableReadSubConn = sub
	}

	o.catalogTableRead, err = icatalogtableread.NewCatalogTableReadClient(ctx, catalogTableReadSubConn, append(opts, dcerpc.WithNoBind(catalogTableReadSubConn))...)

	catalogTableWriteSubConn, err := sub.SubConn(ctx, icatalogtablewrite.CatalogTableWriteSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		catalogTableWriteSubConn = sub
	}

	o.catalogTableWrite, err = icatalogtablewrite.NewCatalogTableWriteClient(ctx, catalogTableWriteSubConn, append(opts, dcerpc.WithNoBind(catalogTableWriteSubConn))...)

	registerSubConn, err := sub.SubConn(ctx, iregister.RegisterSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		registerSubConn = sub
	}

	o.register, err = iregister.NewRegisterClient(ctx, registerSubConn, append(opts, dcerpc.WithNoBind(registerSubConn))...)

	register2SubConn, err := sub.SubConn(ctx, iregister2.Register2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		register2SubConn = sub
	}

	o.register2, err = iregister2.NewRegister2Client(ctx, register2SubConn, append(opts, dcerpc.WithNoBind(register2SubConn))...)

	import1SubConn, err := sub.SubConn(ctx, iimport.ImportSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		import1SubConn = sub
	}

	o.import1, err = iimport.NewImportClient(ctx, import1SubConn, append(opts, dcerpc.WithNoBind(import1SubConn))...)

	import2SubConn, err := sub.SubConn(ctx, iimport2.Import2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		import2SubConn = sub
	}

	o.import2, err = iimport2.NewImport2Client(ctx, import2SubConn, append(opts, dcerpc.WithNoBind(import2SubConn))...)

	exportSubConn, err := sub.SubConn(ctx, iexport.ExportSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		exportSubConn = sub
	}

	o.export, err = iexport.NewExportClient(ctx, exportSubConn, append(opts, dcerpc.WithNoBind(exportSubConn))...)

	export2SubConn, err := sub.SubConn(ctx, iexport2.Export2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		export2SubConn = sub
	}

	o.export2, err = iexport2.NewExport2Client(ctx, export2SubConn, append(opts, dcerpc.WithNoBind(export2SubConn))...)

	alternateLaunchSubConn, err := sub.SubConn(ctx, ialternatelaunch.AlternateLaunchSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		alternateLaunchSubConn = sub
	}

	o.alternateLaunch, err = ialternatelaunch.NewAlternateLaunchClient(ctx, alternateLaunchSubConn, append(opts, dcerpc.WithNoBind(alternateLaunchSubConn))...)

	catalogUtilsSubConn, err := sub.SubConn(ctx, icatalogutils.CatalogUtilsSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		catalogUtilsSubConn = sub
	}

	o.catalogUtils, err = icatalogutils.NewCatalogUtilsClient(ctx, catalogUtilsSubConn, append(opts, dcerpc.WithNoBind(catalogUtilsSubConn))...)

	catalogUtils2SubConn, err := sub.SubConn(ctx, icatalogutils2.CatalogUtils2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		catalogUtils2SubConn = sub
	}

	o.catalogUtils2, err = icatalogutils2.NewCatalogUtils2Client(ctx, catalogUtils2SubConn, append(opts, dcerpc.WithNoBind(catalogUtils2SubConn))...)

	capabilitySupportSubConn, err := sub.SubConn(ctx, icapabilitysupport.CapabilitySupportSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		capabilitySupportSubConn = sub
	}

	o.capabilitySupport, err = icapabilitysupport.NewCapabilitySupportClient(ctx, capabilitySupportSubConn, append(opts, dcerpc.WithNoBind(capabilitySupportSubConn))...)

	containerControlSubConn, err := sub.SubConn(ctx, icontainercontrol.ContainerControlSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		containerControlSubConn = sub
	}

	o.containerControl, err = icontainercontrol.NewContainerControlClient(ctx, containerControlSubConn, append(opts, dcerpc.WithNoBind(containerControlSubConn))...)

	containerControl2SubConn, err := sub.SubConn(ctx, icontainercontrol2.ContainerControl2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		containerControl2SubConn = sub
	}

	o.containerControl2, err = icontainercontrol2.NewContainerControl2Client(ctx, containerControl2SubConn, append(opts, dcerpc.WithNoBind(containerControl2SubConn))...)

	replicationUtilSubConn, err := sub.SubConn(ctx, ireplicationutil.ReplicationUtilSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		replicationUtilSubConn = sub
	}

	o.replicationUtil, err = ireplicationutil.NewReplicationUtilClient(ctx, replicationUtilSubConn, append(opts, dcerpc.WithNoBind(replicationUtilSubConn))...)
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
		dcomClient:          o.dcomClient.IPID(ctx, ipid),
		catalogSession:      o.catalogSession.IPID(ctx, ipid),
		catalog64BitSupport: o.catalog64BitSupport.IPID(ctx, ipid),
		catalogTableInfo:    o.catalogTableInfo.IPID(ctx, ipid),
		catalogTableRead:    o.catalogTableRead.IPID(ctx, ipid),
		catalogTableWrite:   o.catalogTableWrite.IPID(ctx, ipid),
		register:            o.register.IPID(ctx, ipid),
		register2:           o.register2.IPID(ctx, ipid),
		import1:             o.import1.IPID(ctx, ipid),
		import2:             o.import2.IPID(ctx, ipid),
		export:              o.export.IPID(ctx, ipid),
		export2:             o.export2.IPID(ctx, ipid),
		alternateLaunch:     o.alternateLaunch.IPID(ctx, ipid),
		catalogUtils:        o.catalogUtils.IPID(ctx, ipid),
		catalogUtils2:       o.catalogUtils2.IPID(ctx, ipid),
		capabilitySupport:   o.capabilitySupport.IPID(ctx, ipid),
		containerControl:    o.containerControl.IPID(ctx, ipid),
		containerControl2:   o.containerControl2.IPID(ctx, ipid),
		replicationUtil:     o.replicationUtil.IPID(ctx, ipid),
		cc:                  o.cc,
	}
}
