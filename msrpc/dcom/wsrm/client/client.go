package wsrm

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
	iresourcemanager "github.com/oiweiwei/go-msrpc/msrpc/dcom/wsrm/iresourcemanager/v0"
	iresourcemanager2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/wsrm/iresourcemanager2/v0"
	iwrmaccounting "github.com/oiweiwei/go-msrpc/msrpc/dcom/wsrm/iwrmaccounting/v0"
	iwrmcalendar "github.com/oiweiwei/go-msrpc/msrpc/dcom/wsrm/iwrmcalendar/v0"
	iwrmconfig "github.com/oiweiwei/go-msrpc/msrpc/dcom/wsrm/iwrmconfig/v0"
	iwrmmachinegroup "github.com/oiweiwei/go-msrpc/msrpc/dcom/wsrm/iwrmmachinegroup/v0"
	iwrmpolicy "github.com/oiweiwei/go-msrpc/msrpc/dcom/wsrm/iwrmpolicy/v0"
	iwrmprotocol "github.com/oiweiwei/go-msrpc/msrpc/dcom/wsrm/iwrmprotocol/v0"
	iwrmremotesessionmgmt "github.com/oiweiwei/go-msrpc/msrpc/dcom/wsrm/iwrmremotesessionmgmt/v0"
	iwrmresourcegroup "github.com/oiweiwei/go-msrpc/msrpc/dcom/wsrm/iwrmresourcegroup/v0"
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
	_ = iresourcemanager.GoPackage
	_ = iwrmcalendar.GoPackage
	_ = iwrmpolicy.GoPackage
	_ = iwrmresourcegroup.GoPackage
	_ = iwrmaccounting.GoPackage
	_ = iwrmconfig.GoPackage
	_ = iwrmprotocol.GoPackage
	_ = iwrmmachinegroup.GoPackage
	_ = iresourcemanager2.GoPackage
	_ = iwrmremotesessionmgmt.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/wsrm"
)

// dcom/wsrm client set.
type Client interface {

	// DCOM common interfaces
	RemoteUnknown() iremunknown.RemoteUnknownClient
	RemoteUnknown2() iremunknown2.RemoteUnknown2Client

	// Package specific interfaces
	ResourceManager() iresourcemanager.ResourceManagerClient
	IwrmCalendar() iwrmcalendar.IwrmCalendarClient
	IwrmPolicy() iwrmpolicy.IwrmPolicyClient
	IwrmResourceGroup() iwrmresourcegroup.IwrmResourceGroupClient
	IwrmAccounting() iwrmaccounting.IwrmAccountingClient
	IwrmConfig() iwrmconfig.IwrmConfigClient
	IwrmProtocol() iwrmprotocol.IwrmProtocolClient
	IwrmMachineGroup() iwrmmachinegroup.IwrmMachineGroupClient
	ResourceManager2() iresourcemanager2.ResourceManager2Client
	IwrmRemoteSessionManagement() iwrmremotesessionmgmt.IwrmRemoteSessionManagementClient
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

	resourceManager             iresourcemanager.ResourceManagerClient
	iwrmCalendar                iwrmcalendar.IwrmCalendarClient
	iwrmPolicy                  iwrmpolicy.IwrmPolicyClient
	iwrmResourceGroup           iwrmresourcegroup.IwrmResourceGroupClient
	iwrmAccounting              iwrmaccounting.IwrmAccountingClient
	iwrmConfig                  iwrmconfig.IwrmConfigClient
	iwrmProtocol                iwrmprotocol.IwrmProtocolClient
	iwrmMachineGroup            iwrmmachinegroup.IwrmMachineGroupClient
	resourceManager2            iresourcemanager2.ResourceManager2Client
	iwrmRemoteSessionManagement iwrmremotesessionmgmt.IwrmRemoteSessionManagementClient
}

func (o *xxx_DefaultClient) RemoteUnknown() iremunknown.RemoteUnknownClient {
	return o.dcomClient.RemoteUnknown()
}

func (o *xxx_DefaultClient) RemoteUnknown2() iremunknown2.RemoteUnknown2Client {
	return o.dcomClient.RemoteUnknown2()
}

func (o *xxx_DefaultClient) ResourceManager() iresourcemanager.ResourceManagerClient {
	return o.resourceManager
}

func (o *xxx_DefaultClient) IwrmCalendar() iwrmcalendar.IwrmCalendarClient {
	return o.iwrmCalendar
}

func (o *xxx_DefaultClient) IwrmPolicy() iwrmpolicy.IwrmPolicyClient {
	return o.iwrmPolicy
}

func (o *xxx_DefaultClient) IwrmResourceGroup() iwrmresourcegroup.IwrmResourceGroupClient {
	return o.iwrmResourceGroup
}

func (o *xxx_DefaultClient) IwrmAccounting() iwrmaccounting.IwrmAccountingClient {
	return o.iwrmAccounting
}

func (o *xxx_DefaultClient) IwrmConfig() iwrmconfig.IwrmConfigClient {
	return o.iwrmConfig
}

func (o *xxx_DefaultClient) IwrmProtocol() iwrmprotocol.IwrmProtocolClient {
	return o.iwrmProtocol
}

func (o *xxx_DefaultClient) IwrmMachineGroup() iwrmmachinegroup.IwrmMachineGroupClient {
	return o.iwrmMachineGroup
}

func (o *xxx_DefaultClient) ResourceManager2() iresourcemanager2.ResourceManager2Client {
	return o.resourceManager2
}

func (o *xxx_DefaultClient) IwrmRemoteSessionManagement() iwrmremotesessionmgmt.IwrmRemoteSessionManagementClient {
	return o.iwrmRemoteSessionManagement
}

func NewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Client, error) {

	opts = append(opts,
		dcerpc.WithAbstractSyntax(iresourcemanager.ResourceManagerSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwrmcalendar.IwrmCalendarSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwrmpolicy.IwrmPolicySyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwrmresourcegroup.IwrmResourceGroupSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwrmaccounting.IwrmAccountingSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwrmconfig.IwrmConfigSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwrmprotocol.IwrmProtocolSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwrmmachinegroup.IwrmMachineGroupSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iresourcemanager2.ResourceManager2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwrmremotesessionmgmt.IwrmRemoteSessionManagementSyntaxV0_0),
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

	resourceManagerSubConn, err := sub.SubConn(ctx, iresourcemanager.ResourceManagerSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		resourceManagerSubConn = sub
	}

	o.resourceManager, err = iresourcemanager.NewResourceManagerClient(ctx, resourceManagerSubConn, append(opts, dcerpc.WithNoBind(resourceManagerSubConn))...)

	iwrmCalendarSubConn, err := sub.SubConn(ctx, iwrmcalendar.IwrmCalendarSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		iwrmCalendarSubConn = sub
	}

	o.iwrmCalendar, err = iwrmcalendar.NewIwrmCalendarClient(ctx, iwrmCalendarSubConn, append(opts, dcerpc.WithNoBind(iwrmCalendarSubConn))...)

	iwrmPolicySubConn, err := sub.SubConn(ctx, iwrmpolicy.IwrmPolicySyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		iwrmPolicySubConn = sub
	}

	o.iwrmPolicy, err = iwrmpolicy.NewIwrmPolicyClient(ctx, iwrmPolicySubConn, append(opts, dcerpc.WithNoBind(iwrmPolicySubConn))...)

	iwrmResourceGroupSubConn, err := sub.SubConn(ctx, iwrmresourcegroup.IwrmResourceGroupSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		iwrmResourceGroupSubConn = sub
	}

	o.iwrmResourceGroup, err = iwrmresourcegroup.NewIwrmResourceGroupClient(ctx, iwrmResourceGroupSubConn, append(opts, dcerpc.WithNoBind(iwrmResourceGroupSubConn))...)

	iwrmAccountingSubConn, err := sub.SubConn(ctx, iwrmaccounting.IwrmAccountingSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		iwrmAccountingSubConn = sub
	}

	o.iwrmAccounting, err = iwrmaccounting.NewIwrmAccountingClient(ctx, iwrmAccountingSubConn, append(opts, dcerpc.WithNoBind(iwrmAccountingSubConn))...)

	iwrmConfigSubConn, err := sub.SubConn(ctx, iwrmconfig.IwrmConfigSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		iwrmConfigSubConn = sub
	}

	o.iwrmConfig, err = iwrmconfig.NewIwrmConfigClient(ctx, iwrmConfigSubConn, append(opts, dcerpc.WithNoBind(iwrmConfigSubConn))...)

	iwrmProtocolSubConn, err := sub.SubConn(ctx, iwrmprotocol.IwrmProtocolSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		iwrmProtocolSubConn = sub
	}

	o.iwrmProtocol, err = iwrmprotocol.NewIwrmProtocolClient(ctx, iwrmProtocolSubConn, append(opts, dcerpc.WithNoBind(iwrmProtocolSubConn))...)

	iwrmMachineGroupSubConn, err := sub.SubConn(ctx, iwrmmachinegroup.IwrmMachineGroupSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		iwrmMachineGroupSubConn = sub
	}

	o.iwrmMachineGroup, err = iwrmmachinegroup.NewIwrmMachineGroupClient(ctx, iwrmMachineGroupSubConn, append(opts, dcerpc.WithNoBind(iwrmMachineGroupSubConn))...)

	resourceManager2SubConn, err := sub.SubConn(ctx, iresourcemanager2.ResourceManager2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		resourceManager2SubConn = sub
	}

	o.resourceManager2, err = iresourcemanager2.NewResourceManager2Client(ctx, resourceManager2SubConn, append(opts, dcerpc.WithNoBind(resourceManager2SubConn))...)

	iwrmRemoteSessionManagementSubConn, err := sub.SubConn(ctx, iwrmremotesessionmgmt.IwrmRemoteSessionManagementSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		iwrmRemoteSessionManagementSubConn = sub
	}

	o.iwrmRemoteSessionManagement, err = iwrmremotesessionmgmt.NewIwrmRemoteSessionManagementClient(ctx, iwrmRemoteSessionManagementSubConn, append(opts, dcerpc.WithNoBind(iwrmRemoteSessionManagementSubConn))...)
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
		dcomClient:                  o.dcomClient.IPID(ctx, ipid),
		resourceManager:             o.resourceManager.IPID(ctx, ipid),
		iwrmCalendar:                o.iwrmCalendar.IPID(ctx, ipid),
		iwrmPolicy:                  o.iwrmPolicy.IPID(ctx, ipid),
		iwrmResourceGroup:           o.iwrmResourceGroup.IPID(ctx, ipid),
		iwrmAccounting:              o.iwrmAccounting.IPID(ctx, ipid),
		iwrmConfig:                  o.iwrmConfig.IPID(ctx, ipid),
		iwrmProtocol:                o.iwrmProtocol.IPID(ctx, ipid),
		iwrmMachineGroup:            o.iwrmMachineGroup.IPID(ctx, ipid),
		resourceManager2:            o.resourceManager2.IPID(ctx, ipid),
		iwrmRemoteSessionManagement: o.iwrmRemoteSessionManagement.IPID(ctx, ipid),
		cc:                          o.cc,
	}
}
