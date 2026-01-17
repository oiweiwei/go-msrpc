package pla

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	dcom_client "github.com/oiweiwei/go-msrpc/msrpc/dcom/client"
	iremunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iremunknown/v0"
	iremunknown2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/iremunknown2/v0"
	ialertdatacollector "github.com/oiweiwei/go-msrpc/msrpc/dcom/pla/ialertdatacollector/v0"
	iapitracingdatacollector "github.com/oiweiwei/go-msrpc/msrpc/dcom/pla/iapitracingdatacollector/v0"
	iconfigurationdatacollector "github.com/oiweiwei/go-msrpc/msrpc/dcom/pla/iconfigurationdatacollector/v0"
	idatacollector "github.com/oiweiwei/go-msrpc/msrpc/dcom/pla/idatacollector/v0"
	idatacollectorcollection "github.com/oiweiwei/go-msrpc/msrpc/dcom/pla/idatacollectorcollection/v0"
	idatacollectorset "github.com/oiweiwei/go-msrpc/msrpc/dcom/pla/idatacollectorset/v0"
	idatacollectorsetcollection "github.com/oiweiwei/go-msrpc/msrpc/dcom/pla/idatacollectorsetcollection/v0"
	idatamanager "github.com/oiweiwei/go-msrpc/msrpc/dcom/pla/idatamanager/v0"
	ifolderaction "github.com/oiweiwei/go-msrpc/msrpc/dcom/pla/ifolderaction/v0"
	ifolderactioncollection "github.com/oiweiwei/go-msrpc/msrpc/dcom/pla/ifolderactioncollection/v0"
	iperformancecounterdatacollector "github.com/oiweiwei/go-msrpc/msrpc/dcom/pla/iperformancecounterdatacollector/v0"
	ischedule "github.com/oiweiwei/go-msrpc/msrpc/dcom/pla/ischedule/v0"
	ischedulecollection "github.com/oiweiwei/go-msrpc/msrpc/dcom/pla/ischedulecollection/v0"
	itracedatacollector "github.com/oiweiwei/go-msrpc/msrpc/dcom/pla/itracedatacollector/v0"
	itracedataprovider "github.com/oiweiwei/go-msrpc/msrpc/dcom/pla/itracedataprovider/v0"
	itracedataprovidercollection "github.com/oiweiwei/go-msrpc/msrpc/dcom/pla/itracedataprovidercollection/v0"
	ivaluemap "github.com/oiweiwei/go-msrpc/msrpc/dcom/pla/ivaluemap/v0"
	ivaluemapitem "github.com/oiweiwei/go-msrpc/msrpc/dcom/pla/ivaluemapitem/v0"
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
	_ = dcom_client.GoPackage
	_ = dcom.GoPackage
	_ = iremunknown.GoPackage
	_ = iremunknown2.GoPackage
	_ = idatacollectorset.GoPackage
	_ = idatamanager.GoPackage
	_ = ifolderaction.GoPackage
	_ = ifolderactioncollection.GoPackage
	_ = idatacollector.GoPackage
	_ = iperformancecounterdatacollector.GoPackage
	_ = iconfigurationdatacollector.GoPackage
	_ = ialertdatacollector.GoPackage
	_ = itracedatacollector.GoPackage
	_ = iapitracingdatacollector.GoPackage
	_ = itracedataprovider.GoPackage
	_ = ischedule.GoPackage
	_ = itracedataprovidercollection.GoPackage
	_ = ischedulecollection.GoPackage
	_ = idatacollectorcollection.GoPackage
	_ = idatacollectorsetcollection.GoPackage
	_ = ivaluemapitem.GoPackage
	_ = ivaluemap.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/pla"
)

// dcom/pla client set.
type Client interface {

	// DCOM common interfaces
	RemoteUnknown() iremunknown.RemoteUnknownClient
	RemoteUnknown2() iremunknown2.RemoteUnknown2Client

	// Package specific interfaces
	DataCollectorSet() idatacollectorset.DataCollectorSetClient
	DataManager() idatamanager.DataManagerClient
	FolderAction() ifolderaction.FolderActionClient
	FolderActionCollection() ifolderactioncollection.FolderActionCollectionClient
	DataCollector() idatacollector.DataCollectorClient
	PerformanceCounterDataCollector() iperformancecounterdatacollector.PerformanceCounterDataCollectorClient
	ConfigurationDataCollector() iconfigurationdatacollector.ConfigurationDataCollectorClient
	AlertDataCollector() ialertdatacollector.AlertDataCollectorClient
	TraceDataCollector() itracedatacollector.TraceDataCollectorClient
	APITracingDataCollector() iapitracingdatacollector.APITracingDataCollectorClient
	TraceDataProvider() itracedataprovider.TraceDataProviderClient
	Schedule() ischedule.ScheduleClient
	TraceDataProviderCollection() itracedataprovidercollection.TraceDataProviderCollectionClient
	ScheduleCollection() ischedulecollection.ScheduleCollectionClient
	DataCollectorCollection() idatacollectorcollection.DataCollectorCollectionClient
	DataCollectorSetCollection() idatacollectorsetcollection.DataCollectorSetCollectionClient
	ValueMapItem() ivaluemapitem.ValueMapItemClient
	ValueMap() ivaluemap.ValueMapClient
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

	dataCollectorSet                idatacollectorset.DataCollectorSetClient
	dataManager                     idatamanager.DataManagerClient
	folderAction                    ifolderaction.FolderActionClient
	folderActionCollection          ifolderactioncollection.FolderActionCollectionClient
	dataCollector                   idatacollector.DataCollectorClient
	performanceCounterDataCollector iperformancecounterdatacollector.PerformanceCounterDataCollectorClient
	configurationDataCollector      iconfigurationdatacollector.ConfigurationDataCollectorClient
	alertDataCollector              ialertdatacollector.AlertDataCollectorClient
	traceDataCollector              itracedatacollector.TraceDataCollectorClient
	apiTracingDataCollector         iapitracingdatacollector.APITracingDataCollectorClient
	traceDataProvider               itracedataprovider.TraceDataProviderClient
	schedule                        ischedule.ScheduleClient
	traceDataProviderCollection     itracedataprovidercollection.TraceDataProviderCollectionClient
	scheduleCollection              ischedulecollection.ScheduleCollectionClient
	dataCollectorCollection         idatacollectorcollection.DataCollectorCollectionClient
	dataCollectorSetCollection      idatacollectorsetcollection.DataCollectorSetCollectionClient
	valueMapItem                    ivaluemapitem.ValueMapItemClient
	valueMap                        ivaluemap.ValueMapClient
}

func (o *xxx_DefaultClient) RemoteUnknown() iremunknown.RemoteUnknownClient {
	return o.dcomClient.RemoteUnknown()
}

func (o *xxx_DefaultClient) RemoteUnknown2() iremunknown2.RemoteUnknown2Client {
	return o.dcomClient.RemoteUnknown2()
}

func (o *xxx_DefaultClient) DataCollectorSet() idatacollectorset.DataCollectorSetClient {
	return o.dataCollectorSet
}

func (o *xxx_DefaultClient) DataManager() idatamanager.DataManagerClient {
	return o.dataManager
}

func (o *xxx_DefaultClient) FolderAction() ifolderaction.FolderActionClient {
	return o.folderAction
}

func (o *xxx_DefaultClient) FolderActionCollection() ifolderactioncollection.FolderActionCollectionClient {
	return o.folderActionCollection
}

func (o *xxx_DefaultClient) DataCollector() idatacollector.DataCollectorClient {
	return o.dataCollector
}

func (o *xxx_DefaultClient) PerformanceCounterDataCollector() iperformancecounterdatacollector.PerformanceCounterDataCollectorClient {
	return o.performanceCounterDataCollector
}

func (o *xxx_DefaultClient) ConfigurationDataCollector() iconfigurationdatacollector.ConfigurationDataCollectorClient {
	return o.configurationDataCollector
}

func (o *xxx_DefaultClient) AlertDataCollector() ialertdatacollector.AlertDataCollectorClient {
	return o.alertDataCollector
}

func (o *xxx_DefaultClient) TraceDataCollector() itracedatacollector.TraceDataCollectorClient {
	return o.traceDataCollector
}

func (o *xxx_DefaultClient) APITracingDataCollector() iapitracingdatacollector.APITracingDataCollectorClient {
	return o.apiTracingDataCollector
}

func (o *xxx_DefaultClient) TraceDataProvider() itracedataprovider.TraceDataProviderClient {
	return o.traceDataProvider
}

func (o *xxx_DefaultClient) Schedule() ischedule.ScheduleClient {
	return o.schedule
}

func (o *xxx_DefaultClient) TraceDataProviderCollection() itracedataprovidercollection.TraceDataProviderCollectionClient {
	return o.traceDataProviderCollection
}

func (o *xxx_DefaultClient) ScheduleCollection() ischedulecollection.ScheduleCollectionClient {
	return o.scheduleCollection
}

func (o *xxx_DefaultClient) DataCollectorCollection() idatacollectorcollection.DataCollectorCollectionClient {
	return o.dataCollectorCollection
}

func (o *xxx_DefaultClient) DataCollectorSetCollection() idatacollectorsetcollection.DataCollectorSetCollectionClient {
	return o.dataCollectorSetCollection
}

func (o *xxx_DefaultClient) ValueMapItem() ivaluemapitem.ValueMapItemClient {
	return o.valueMapItem
}

func (o *xxx_DefaultClient) ValueMap() ivaluemap.ValueMapClient {
	return o.valueMap
}

func NewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Client, error) {

	opts = append(opts,
		dcerpc.WithAbstractSyntax(idatacollectorset.DataCollectorSetSyntaxV0_0),
		dcerpc.WithAbstractSyntax(idatamanager.DataManagerSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifolderaction.FolderActionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifolderactioncollection.FolderActionCollectionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(idatacollector.DataCollectorSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iperformancecounterdatacollector.PerformanceCounterDataCollectorSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iconfigurationdatacollector.ConfigurationDataCollectorSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ialertdatacollector.AlertDataCollectorSyntaxV0_0),
		dcerpc.WithAbstractSyntax(itracedatacollector.TraceDataCollectorSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iapitracingdatacollector.APITracingDataCollectorSyntaxV0_0),
		dcerpc.WithAbstractSyntax(itracedataprovider.TraceDataProviderSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ischedule.ScheduleSyntaxV0_0),
		dcerpc.WithAbstractSyntax(itracedataprovidercollection.TraceDataProviderCollectionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ischedulecollection.ScheduleCollectionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(idatacollectorcollection.DataCollectorCollectionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(idatacollectorsetcollection.DataCollectorSetCollectionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivaluemapitem.ValueMapItemSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivaluemap.ValueMapSyntaxV0_0),
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

	dataCollectorSetSubConn, err := sub.SubConn(ctx, idatacollectorset.DataCollectorSetSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		dataCollectorSetSubConn = sub
	}

	o.dataCollectorSet, err = idatacollectorset.NewDataCollectorSetClient(ctx, dataCollectorSetSubConn, append(opts, dcerpc.WithNoBind(dataCollectorSetSubConn))...)

	dataManagerSubConn, err := sub.SubConn(ctx, idatamanager.DataManagerSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		dataManagerSubConn = sub
	}

	o.dataManager, err = idatamanager.NewDataManagerClient(ctx, dataManagerSubConn, append(opts, dcerpc.WithNoBind(dataManagerSubConn))...)

	folderActionSubConn, err := sub.SubConn(ctx, ifolderaction.FolderActionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		folderActionSubConn = sub
	}

	o.folderAction, err = ifolderaction.NewFolderActionClient(ctx, folderActionSubConn, append(opts, dcerpc.WithNoBind(folderActionSubConn))...)

	folderActionCollectionSubConn, err := sub.SubConn(ctx, ifolderactioncollection.FolderActionCollectionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		folderActionCollectionSubConn = sub
	}

	o.folderActionCollection, err = ifolderactioncollection.NewFolderActionCollectionClient(ctx, folderActionCollectionSubConn, append(opts, dcerpc.WithNoBind(folderActionCollectionSubConn))...)

	dataCollectorSubConn, err := sub.SubConn(ctx, idatacollector.DataCollectorSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		dataCollectorSubConn = sub
	}

	o.dataCollector, err = idatacollector.NewDataCollectorClient(ctx, dataCollectorSubConn, append(opts, dcerpc.WithNoBind(dataCollectorSubConn))...)

	performanceCounterDataCollectorSubConn, err := sub.SubConn(ctx, iperformancecounterdatacollector.PerformanceCounterDataCollectorSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		performanceCounterDataCollectorSubConn = sub
	}

	o.performanceCounterDataCollector, err = iperformancecounterdatacollector.NewPerformanceCounterDataCollectorClient(ctx, performanceCounterDataCollectorSubConn, append(opts, dcerpc.WithNoBind(performanceCounterDataCollectorSubConn))...)

	configurationDataCollectorSubConn, err := sub.SubConn(ctx, iconfigurationdatacollector.ConfigurationDataCollectorSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		configurationDataCollectorSubConn = sub
	}

	o.configurationDataCollector, err = iconfigurationdatacollector.NewConfigurationDataCollectorClient(ctx, configurationDataCollectorSubConn, append(opts, dcerpc.WithNoBind(configurationDataCollectorSubConn))...)

	alertDataCollectorSubConn, err := sub.SubConn(ctx, ialertdatacollector.AlertDataCollectorSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		alertDataCollectorSubConn = sub
	}

	o.alertDataCollector, err = ialertdatacollector.NewAlertDataCollectorClient(ctx, alertDataCollectorSubConn, append(opts, dcerpc.WithNoBind(alertDataCollectorSubConn))...)

	traceDataCollectorSubConn, err := sub.SubConn(ctx, itracedatacollector.TraceDataCollectorSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		traceDataCollectorSubConn = sub
	}

	o.traceDataCollector, err = itracedatacollector.NewTraceDataCollectorClient(ctx, traceDataCollectorSubConn, append(opts, dcerpc.WithNoBind(traceDataCollectorSubConn))...)

	apiTracingDataCollectorSubConn, err := sub.SubConn(ctx, iapitracingdatacollector.APITracingDataCollectorSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		apiTracingDataCollectorSubConn = sub
	}

	o.apiTracingDataCollector, err = iapitracingdatacollector.NewAPITracingDataCollectorClient(ctx, apiTracingDataCollectorSubConn, append(opts, dcerpc.WithNoBind(apiTracingDataCollectorSubConn))...)

	traceDataProviderSubConn, err := sub.SubConn(ctx, itracedataprovider.TraceDataProviderSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		traceDataProviderSubConn = sub
	}

	o.traceDataProvider, err = itracedataprovider.NewTraceDataProviderClient(ctx, traceDataProviderSubConn, append(opts, dcerpc.WithNoBind(traceDataProviderSubConn))...)

	scheduleSubConn, err := sub.SubConn(ctx, ischedule.ScheduleSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		scheduleSubConn = sub
	}

	o.schedule, err = ischedule.NewScheduleClient(ctx, scheduleSubConn, append(opts, dcerpc.WithNoBind(scheduleSubConn))...)

	traceDataProviderCollectionSubConn, err := sub.SubConn(ctx, itracedataprovidercollection.TraceDataProviderCollectionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		traceDataProviderCollectionSubConn = sub
	}

	o.traceDataProviderCollection, err = itracedataprovidercollection.NewTraceDataProviderCollectionClient(ctx, traceDataProviderCollectionSubConn, append(opts, dcerpc.WithNoBind(traceDataProviderCollectionSubConn))...)

	scheduleCollectionSubConn, err := sub.SubConn(ctx, ischedulecollection.ScheduleCollectionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		scheduleCollectionSubConn = sub
	}

	o.scheduleCollection, err = ischedulecollection.NewScheduleCollectionClient(ctx, scheduleCollectionSubConn, append(opts, dcerpc.WithNoBind(scheduleCollectionSubConn))...)

	dataCollectorCollectionSubConn, err := sub.SubConn(ctx, idatacollectorcollection.DataCollectorCollectionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		dataCollectorCollectionSubConn = sub
	}

	o.dataCollectorCollection, err = idatacollectorcollection.NewDataCollectorCollectionClient(ctx, dataCollectorCollectionSubConn, append(opts, dcerpc.WithNoBind(dataCollectorCollectionSubConn))...)

	dataCollectorSetCollectionSubConn, err := sub.SubConn(ctx, idatacollectorsetcollection.DataCollectorSetCollectionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		dataCollectorSetCollectionSubConn = sub
	}

	o.dataCollectorSetCollection, err = idatacollectorsetcollection.NewDataCollectorSetCollectionClient(ctx, dataCollectorSetCollectionSubConn, append(opts, dcerpc.WithNoBind(dataCollectorSetCollectionSubConn))...)

	valueMapItemSubConn, err := sub.SubConn(ctx, ivaluemapitem.ValueMapItemSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		valueMapItemSubConn = sub
	}

	o.valueMapItem, err = ivaluemapitem.NewValueMapItemClient(ctx, valueMapItemSubConn, append(opts, dcerpc.WithNoBind(valueMapItemSubConn))...)

	valueMapSubConn, err := sub.SubConn(ctx, ivaluemap.ValueMapSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		valueMapSubConn = sub
	}

	o.valueMap, err = ivaluemap.NewValueMapClient(ctx, valueMapSubConn, append(opts, dcerpc.WithNoBind(valueMapSubConn))...)
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
		dcomClient:                      o.dcomClient.IPID(ctx, ipid),
		dataCollectorSet:                o.dataCollectorSet.IPID(ctx, ipid),
		dataManager:                     o.dataManager.IPID(ctx, ipid),
		folderAction:                    o.folderAction.IPID(ctx, ipid),
		folderActionCollection:          o.folderActionCollection.IPID(ctx, ipid),
		dataCollector:                   o.dataCollector.IPID(ctx, ipid),
		performanceCounterDataCollector: o.performanceCounterDataCollector.IPID(ctx, ipid),
		configurationDataCollector:      o.configurationDataCollector.IPID(ctx, ipid),
		alertDataCollector:              o.alertDataCollector.IPID(ctx, ipid),
		traceDataCollector:              o.traceDataCollector.IPID(ctx, ipid),
		apiTracingDataCollector:         o.apiTracingDataCollector.IPID(ctx, ipid),
		traceDataProvider:               o.traceDataProvider.IPID(ctx, ipid),
		schedule:                        o.schedule.IPID(ctx, ipid),
		traceDataProviderCollection:     o.traceDataProviderCollection.IPID(ctx, ipid),
		scheduleCollection:              o.scheduleCollection.IPID(ctx, ipid),
		dataCollectorCollection:         o.dataCollectorCollection.IPID(ctx, ipid),
		dataCollectorSetCollection:      o.dataCollectorSetCollection.IPID(ctx, ipid),
		valueMapItem:                    o.valueMapItem.IPID(ctx, ipid),
		valueMap:                        o.valueMap.IPID(ctx, ipid),
		cc:                              o.cc,
	}
}
