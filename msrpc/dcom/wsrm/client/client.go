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
	Calendar() iwrmcalendar.CalendarClient
	Policy() iwrmpolicy.PolicyClient
	ResourceGroup() iwrmresourcegroup.ResourceGroupClient
	Accounting() iwrmaccounting.AccountingClient
	Config() iwrmconfig.ConfigClient
	Protocol() iwrmprotocol.ProtocolClient
	MachineGroup() iwrmmachinegroup.MachineGroupClient
	ResourceManager2() iresourcemanager2.ResourceManager2Client
	RemoteSessionManagement() iwrmremotesessionmgmt.RemoteSessionManagementClient
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

	resourceManager         iresourcemanager.ResourceManagerClient
	calendar                iwrmcalendar.CalendarClient
	policy                  iwrmpolicy.PolicyClient
	resourceGroup           iwrmresourcegroup.ResourceGroupClient
	accounting              iwrmaccounting.AccountingClient
	config                  iwrmconfig.ConfigClient
	protocol                iwrmprotocol.ProtocolClient
	machineGroup            iwrmmachinegroup.MachineGroupClient
	resourceManager2        iresourcemanager2.ResourceManager2Client
	remoteSessionManagement iwrmremotesessionmgmt.RemoteSessionManagementClient
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

func (o *xxx_DefaultClient) Calendar() iwrmcalendar.CalendarClient {
	return o.calendar
}

func (o *xxx_DefaultClient) Policy() iwrmpolicy.PolicyClient {
	return o.policy
}

func (o *xxx_DefaultClient) ResourceGroup() iwrmresourcegroup.ResourceGroupClient {
	return o.resourceGroup
}

func (o *xxx_DefaultClient) Accounting() iwrmaccounting.AccountingClient {
	return o.accounting
}

func (o *xxx_DefaultClient) Config() iwrmconfig.ConfigClient {
	return o.config
}

func (o *xxx_DefaultClient) Protocol() iwrmprotocol.ProtocolClient {
	return o.protocol
}

func (o *xxx_DefaultClient) MachineGroup() iwrmmachinegroup.MachineGroupClient {
	return o.machineGroup
}

func (o *xxx_DefaultClient) ResourceManager2() iresourcemanager2.ResourceManager2Client {
	return o.resourceManager2
}

func (o *xxx_DefaultClient) RemoteSessionManagement() iwrmremotesessionmgmt.RemoteSessionManagementClient {
	return o.remoteSessionManagement
}

func NewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Client, error) {

	opts = append(opts,
		dcerpc.WithAbstractSyntax(iresourcemanager.ResourceManagerSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwrmcalendar.CalendarSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwrmpolicy.PolicySyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwrmresourcegroup.ResourceGroupSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwrmaccounting.AccountingSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwrmconfig.ConfigSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwrmprotocol.ProtocolSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwrmmachinegroup.MachineGroupSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iresourcemanager2.ResourceManager2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwrmremotesessionmgmt.RemoteSessionManagementSyntaxV0_0),
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

	calendarSubConn, err := sub.SubConn(ctx, iwrmcalendar.CalendarSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		calendarSubConn = sub
	}

	o.calendar, err = iwrmcalendar.NewCalendarClient(ctx, calendarSubConn, append(opts, dcerpc.WithNoBind(calendarSubConn))...)

	policySubConn, err := sub.SubConn(ctx, iwrmpolicy.PolicySyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		policySubConn = sub
	}

	o.policy, err = iwrmpolicy.NewPolicyClient(ctx, policySubConn, append(opts, dcerpc.WithNoBind(policySubConn))...)

	resourceGroupSubConn, err := sub.SubConn(ctx, iwrmresourcegroup.ResourceGroupSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		resourceGroupSubConn = sub
	}

	o.resourceGroup, err = iwrmresourcegroup.NewResourceGroupClient(ctx, resourceGroupSubConn, append(opts, dcerpc.WithNoBind(resourceGroupSubConn))...)

	accountingSubConn, err := sub.SubConn(ctx, iwrmaccounting.AccountingSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		accountingSubConn = sub
	}

	o.accounting, err = iwrmaccounting.NewAccountingClient(ctx, accountingSubConn, append(opts, dcerpc.WithNoBind(accountingSubConn))...)

	configSubConn, err := sub.SubConn(ctx, iwrmconfig.ConfigSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		configSubConn = sub
	}

	o.config, err = iwrmconfig.NewConfigClient(ctx, configSubConn, append(opts, dcerpc.WithNoBind(configSubConn))...)

	protocolSubConn, err := sub.SubConn(ctx, iwrmprotocol.ProtocolSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		protocolSubConn = sub
	}

	o.protocol, err = iwrmprotocol.NewProtocolClient(ctx, protocolSubConn, append(opts, dcerpc.WithNoBind(protocolSubConn))...)

	machineGroupSubConn, err := sub.SubConn(ctx, iwrmmachinegroup.MachineGroupSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		machineGroupSubConn = sub
	}

	o.machineGroup, err = iwrmmachinegroup.NewMachineGroupClient(ctx, machineGroupSubConn, append(opts, dcerpc.WithNoBind(machineGroupSubConn))...)

	resourceManager2SubConn, err := sub.SubConn(ctx, iresourcemanager2.ResourceManager2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		resourceManager2SubConn = sub
	}

	o.resourceManager2, err = iresourcemanager2.NewResourceManager2Client(ctx, resourceManager2SubConn, append(opts, dcerpc.WithNoBind(resourceManager2SubConn))...)

	remoteSessionManagementSubConn, err := sub.SubConn(ctx, iwrmremotesessionmgmt.RemoteSessionManagementSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		remoteSessionManagementSubConn = sub
	}

	o.remoteSessionManagement, err = iwrmremotesessionmgmt.NewRemoteSessionManagementClient(ctx, remoteSessionManagementSubConn, append(opts, dcerpc.WithNoBind(remoteSessionManagementSubConn))...)
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
		dcomClient:              o.dcomClient.IPID(ctx, ipid),
		resourceManager:         o.resourceManager.IPID(ctx, ipid),
		calendar:                o.calendar.IPID(ctx, ipid),
		policy:                  o.policy.IPID(ctx, ipid),
		resourceGroup:           o.resourceGroup.IPID(ctx, ipid),
		accounting:              o.accounting.IPID(ctx, ipid),
		config:                  o.config.IPID(ctx, ipid),
		protocol:                o.protocol.IPID(ctx, ipid),
		machineGroup:            o.machineGroup.IPID(ctx, ipid),
		resourceManager2:        o.resourceManager2.IPID(ctx, ipid),
		remoteSessionManagement: o.remoteSessionManagement.IPID(ctx, ipid),
		cc:                      o.cc,
	}
}
