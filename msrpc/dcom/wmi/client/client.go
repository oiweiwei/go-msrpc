package wmi

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
	ienumwbemclassobject "github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi/ienumwbemclassobject/v0"
	iwbembackuprestore "github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi/iwbembackuprestore/v0"
	iwbembackuprestoreex "github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi/iwbembackuprestoreex/v0"
	iwbemcallresult "github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi/iwbemcallresult/v0"
	iwbemfetchsmartenum "github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi/iwbemfetchsmartenum/v0"
	iwbemlevel1login "github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi/iwbemlevel1login/v0"
	iwbemloginclientid "github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi/iwbemloginclientid/v0"
	iwbemloginhelper "github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi/iwbemloginhelper/v0"
	iwbemobjectsink "github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi/iwbemobjectsink/v0"
	iwbemrefreshingservices "github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi/iwbemrefreshingservices/v0"
	iwbemremoterefresher "github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi/iwbemremoterefresher/v0"
	iwbemservices "github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi/iwbemservices/v0"
	iwbemwcosmartenum "github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi/iwbemwcosmartenum/v0"
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
	_ = iwbemobjectsink.GoPackage
	_ = ienumwbemclassobject.GoPackage
	_ = iwbemcallresult.GoPackage
	_ = iwbemservices.GoPackage
	_ = iwbembackuprestore.GoPackage
	_ = iwbembackuprestoreex.GoPackage
	_ = iwbemremoterefresher.GoPackage
	_ = iwbemrefreshingservices.GoPackage
	_ = iwbemwcosmartenum.GoPackage
	_ = iwbemfetchsmartenum.GoPackage
	_ = iwbemloginclientid.GoPackage
	_ = iwbemlevel1login.GoPackage
	_ = iwbemloginhelper.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/wmi"
)

// dcom/wmi client set.
type Client interface {

	// DCOM common interfaces
	RemoteUnknown() iremunknown.RemoteUnknownClient
	RemoteUnknown2() iremunknown2.RemoteUnknown2Client

	// Package specific interfaces
	ObjectSink() iwbemobjectsink.ObjectSinkClient
	EnumClassObject() ienumwbemclassobject.EnumClassObjectClient
	CallResult() iwbemcallresult.CallResultClient
	Services() iwbemservices.ServicesClient
	BackupRestore() iwbembackuprestore.BackupRestoreClient
	BackupRestoreEx() iwbembackuprestoreex.BackupRestoreExClient
	RemoteRefresher() iwbemremoterefresher.RemoteRefresherClient
	RefreshingServices() iwbemrefreshingservices.RefreshingServicesClient
	WCOSmartEnum() iwbemwcosmartenum.WCOSmartEnumClient
	FetchSmartEnum() iwbemfetchsmartenum.FetchSmartEnumClient
	LoginClientID() iwbemloginclientid.LoginClientIDClient
	Level1Login() iwbemlevel1login.Level1LoginClient
	LoginHelper() iwbemloginhelper.LoginHelperClient
	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) Client
}
type xxx_DefaultClient struct {
	cc dcerpc.Conn

	dcomClient dcom_client.Client

	objectSink         iwbemobjectsink.ObjectSinkClient
	enumClassObject    ienumwbemclassobject.EnumClassObjectClient
	callResult         iwbemcallresult.CallResultClient
	services           iwbemservices.ServicesClient
	backupRestore      iwbembackuprestore.BackupRestoreClient
	backupRestoreEx    iwbembackuprestoreex.BackupRestoreExClient
	remoteRefresher    iwbemremoterefresher.RemoteRefresherClient
	refreshingServices iwbemrefreshingservices.RefreshingServicesClient
	wcoSmartEnum       iwbemwcosmartenum.WCOSmartEnumClient
	fetchSmartEnum     iwbemfetchsmartenum.FetchSmartEnumClient
	loginClientID      iwbemloginclientid.LoginClientIDClient
	level1Login        iwbemlevel1login.Level1LoginClient
	loginHelper        iwbemloginhelper.LoginHelperClient
}

func (o *xxx_DefaultClient) RemoteUnknown() iremunknown.RemoteUnknownClient {
	return o.dcomClient.RemoteUnknown()
}

func (o *xxx_DefaultClient) RemoteUnknown2() iremunknown2.RemoteUnknown2Client {
	return o.dcomClient.RemoteUnknown2()
}

func (o *xxx_DefaultClient) ObjectSink() iwbemobjectsink.ObjectSinkClient {
	return o.objectSink
}

func (o *xxx_DefaultClient) EnumClassObject() ienumwbemclassobject.EnumClassObjectClient {
	return o.enumClassObject
}

func (o *xxx_DefaultClient) CallResult() iwbemcallresult.CallResultClient {
	return o.callResult
}

func (o *xxx_DefaultClient) Services() iwbemservices.ServicesClient {
	return o.services
}

func (o *xxx_DefaultClient) BackupRestore() iwbembackuprestore.BackupRestoreClient {
	return o.backupRestore
}

func (o *xxx_DefaultClient) BackupRestoreEx() iwbembackuprestoreex.BackupRestoreExClient {
	return o.backupRestoreEx
}

func (o *xxx_DefaultClient) RemoteRefresher() iwbemremoterefresher.RemoteRefresherClient {
	return o.remoteRefresher
}

func (o *xxx_DefaultClient) RefreshingServices() iwbemrefreshingservices.RefreshingServicesClient {
	return o.refreshingServices
}

func (o *xxx_DefaultClient) WCOSmartEnum() iwbemwcosmartenum.WCOSmartEnumClient {
	return o.wcoSmartEnum
}

func (o *xxx_DefaultClient) FetchSmartEnum() iwbemfetchsmartenum.FetchSmartEnumClient {
	return o.fetchSmartEnum
}

func (o *xxx_DefaultClient) LoginClientID() iwbemloginclientid.LoginClientIDClient {
	return o.loginClientID
}

func (o *xxx_DefaultClient) Level1Login() iwbemlevel1login.Level1LoginClient {
	return o.level1Login
}

func (o *xxx_DefaultClient) LoginHelper() iwbemloginhelper.LoginHelperClient {
	return o.loginHelper
}

func NewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Client, error) {

	opts = append(opts,
		dcerpc.WithAbstractSyntax(iwbemobjectsink.ObjectSinkSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ienumwbemclassobject.EnumClassObjectSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwbemcallresult.CallResultSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwbemservices.ServicesSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwbembackuprestore.BackupRestoreSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwbembackuprestoreex.BackupRestoreExSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwbemremoterefresher.RemoteRefresherSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwbemrefreshingservices.RefreshingServicesSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwbemwcosmartenum.WCOSmartEnumSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwbemfetchsmartenum.FetchSmartEnumSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwbemloginclientid.LoginClientIDSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwbemlevel1login.Level1LoginSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwbemloginhelper.LoginHelperSyntaxV0_0),
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

	objectSinkSubConn, err := sub.SubConn(ctx, iwbemobjectsink.ObjectSinkSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		objectSinkSubConn = sub
	}

	o.objectSink, err = iwbemobjectsink.NewObjectSinkClient(ctx, objectSinkSubConn, append(opts, dcerpc.WithNoBind(objectSinkSubConn))...)

	enumClassObjectSubConn, err := sub.SubConn(ctx, ienumwbemclassobject.EnumClassObjectSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		enumClassObjectSubConn = sub
	}

	o.enumClassObject, err = ienumwbemclassobject.NewEnumClassObjectClient(ctx, enumClassObjectSubConn, append(opts, dcerpc.WithNoBind(enumClassObjectSubConn))...)

	callResultSubConn, err := sub.SubConn(ctx, iwbemcallresult.CallResultSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		callResultSubConn = sub
	}

	o.callResult, err = iwbemcallresult.NewCallResultClient(ctx, callResultSubConn, append(opts, dcerpc.WithNoBind(callResultSubConn))...)

	servicesSubConn, err := sub.SubConn(ctx, iwbemservices.ServicesSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		servicesSubConn = sub
	}

	o.services, err = iwbemservices.NewServicesClient(ctx, servicesSubConn, append(opts, dcerpc.WithNoBind(servicesSubConn))...)

	backupRestoreSubConn, err := sub.SubConn(ctx, iwbembackuprestore.BackupRestoreSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		backupRestoreSubConn = sub
	}

	o.backupRestore, err = iwbembackuprestore.NewBackupRestoreClient(ctx, backupRestoreSubConn, append(opts, dcerpc.WithNoBind(backupRestoreSubConn))...)

	backupRestoreExSubConn, err := sub.SubConn(ctx, iwbembackuprestoreex.BackupRestoreExSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		backupRestoreExSubConn = sub
	}

	o.backupRestoreEx, err = iwbembackuprestoreex.NewBackupRestoreExClient(ctx, backupRestoreExSubConn, append(opts, dcerpc.WithNoBind(backupRestoreExSubConn))...)

	remoteRefresherSubConn, err := sub.SubConn(ctx, iwbemremoterefresher.RemoteRefresherSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		remoteRefresherSubConn = sub
	}

	o.remoteRefresher, err = iwbemremoterefresher.NewRemoteRefresherClient(ctx, remoteRefresherSubConn, append(opts, dcerpc.WithNoBind(remoteRefresherSubConn))...)

	refreshingServicesSubConn, err := sub.SubConn(ctx, iwbemrefreshingservices.RefreshingServicesSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		refreshingServicesSubConn = sub
	}

	o.refreshingServices, err = iwbemrefreshingservices.NewRefreshingServicesClient(ctx, refreshingServicesSubConn, append(opts, dcerpc.WithNoBind(refreshingServicesSubConn))...)

	wcoSmartEnumSubConn, err := sub.SubConn(ctx, iwbemwcosmartenum.WCOSmartEnumSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		wcoSmartEnumSubConn = sub
	}

	o.wcoSmartEnum, err = iwbemwcosmartenum.NewWCOSmartEnumClient(ctx, wcoSmartEnumSubConn, append(opts, dcerpc.WithNoBind(wcoSmartEnumSubConn))...)

	fetchSmartEnumSubConn, err := sub.SubConn(ctx, iwbemfetchsmartenum.FetchSmartEnumSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		fetchSmartEnumSubConn = sub
	}

	o.fetchSmartEnum, err = iwbemfetchsmartenum.NewFetchSmartEnumClient(ctx, fetchSmartEnumSubConn, append(opts, dcerpc.WithNoBind(fetchSmartEnumSubConn))...)

	loginClientIDSubConn, err := sub.SubConn(ctx, iwbemloginclientid.LoginClientIDSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		loginClientIDSubConn = sub
	}

	o.loginClientID, err = iwbemloginclientid.NewLoginClientIDClient(ctx, loginClientIDSubConn, append(opts, dcerpc.WithNoBind(loginClientIDSubConn))...)

	level1LoginSubConn, err := sub.SubConn(ctx, iwbemlevel1login.Level1LoginSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		level1LoginSubConn = sub
	}

	o.level1Login, err = iwbemlevel1login.NewLevel1LoginClient(ctx, level1LoginSubConn, append(opts, dcerpc.WithNoBind(level1LoginSubConn))...)

	loginHelperSubConn, err := sub.SubConn(ctx, iwbemloginhelper.LoginHelperSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		loginHelperSubConn = sub
	}

	o.loginHelper, err = iwbemloginhelper.NewLoginHelperClient(ctx, loginHelperSubConn, append(opts, dcerpc.WithNoBind(loginHelperSubConn))...)
	return o, nil
}

func (o *xxx_DefaultClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultClient) IPID(ctx context.Context, ipid *dcom.IPID) Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultClient{
		dcomClient:         o.dcomClient.IPID(ctx, ipid),
		objectSink:         o.objectSink.IPID(ctx, ipid),
		enumClassObject:    o.enumClassObject.IPID(ctx, ipid),
		callResult:         o.callResult.IPID(ctx, ipid),
		services:           o.services.IPID(ctx, ipid),
		backupRestore:      o.backupRestore.IPID(ctx, ipid),
		backupRestoreEx:    o.backupRestoreEx.IPID(ctx, ipid),
		remoteRefresher:    o.remoteRefresher.IPID(ctx, ipid),
		refreshingServices: o.refreshingServices.IPID(ctx, ipid),
		wcoSmartEnum:       o.wcoSmartEnum.IPID(ctx, ipid),
		fetchSmartEnum:     o.fetchSmartEnum.IPID(ctx, ipid),
		loginClientID:      o.loginClientID.IPID(ctx, ipid),
		level1Login:        o.level1Login.IPID(ctx, ipid),
		loginHelper:        o.loginHelper.IPID(ctx, ipid),
		cc:                 o.cc,
	}
}
