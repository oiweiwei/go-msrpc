package mqac

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
	iconnectionpoint "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/iconnectionpoint/v0"
	iconnectionpointcontainer "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/iconnectionpointcontainer/v0"
	ienumconnectionpoints "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/ienumconnectionpoints/v0"
	ienumconnections "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/ienumconnections/v0"
	imsmqapplication "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqapplication/v0"
	imsmqapplication2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqapplication2/v0"
	imsmqapplication3 "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqapplication3/v0"
	imsmqcollection "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqcollection/v0"
	imsmqcoordinatedtransactiondispenser "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqcoordinatedtransactiondispenser/v0"
	imsmqcoordinatedtransactiondispenser2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqcoordinatedtransactiondispenser2/v0"
	imsmqcoordinatedtransactiondispenser3 "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqcoordinatedtransactiondispenser3/v0"
	imsmqdestination "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqdestination/v0"
	imsmqevent2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqevent2/v0"
	imsmqmanagement "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqmanagement/v0"
	imsmqmessage "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqmessage/v0"
	imsmqmessage2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqmessage2/v0"
	imsmqmessage3 "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqmessage3/v0"
	imsmqmessage4 "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqmessage4/v0"
	imsmqoutgoingqueuemanagement "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqoutgoingqueuemanagement/v0"
	imsmqprivatedestination "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqprivatedestination/v0"
	imsmqprivateevent "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqprivateevent/v0"
	imsmqquery "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqquery/v0"
	imsmqquery2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqquery2/v0"
	imsmqquery3 "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqquery3/v0"
	imsmqquery4 "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqquery4/v0"
	imsmqqueue "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqqueue/v0"
	imsmqqueue2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqqueue2/v0"
	imsmqqueue3 "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqqueue3/v0"
	imsmqqueue4 "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqqueue4/v0"
	imsmqqueueinfo "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqqueueinfo/v0"
	imsmqqueueinfo2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqqueueinfo2/v0"
	imsmqqueueinfo3 "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqqueueinfo3/v0"
	imsmqqueueinfo4 "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqqueueinfo4/v0"
	imsmqqueueinfos "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqqueueinfos/v0"
	imsmqqueueinfos2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqqueueinfos2/v0"
	imsmqqueueinfos3 "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqqueueinfos3/v0"
	imsmqqueueinfos4 "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqqueueinfos4/v0"
	imsmqqueuemanagement "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqqueuemanagement/v0"
	imsmqtransaction "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqtransaction/v0"
	imsmqtransaction2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqtransaction2/v0"
	imsmqtransaction3 "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqtransaction3/v0"
	imsmqtransactiondispenser "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqtransactiondispenser/v0"
	imsmqtransactiondispenser2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqtransactiondispenser2/v0"
	imsmqtransactiondispenser3 "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqtransactiondispenser3/v0"
	itransaction "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/itransaction/v0"
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
	_ = itransaction.GoPackage
	_ = ienumconnections.GoPackage
	_ = iconnectionpoint.GoPackage
	_ = ienumconnectionpoints.GoPackage
	_ = iconnectionpointcontainer.GoPackage
	_ = imsmqquery.GoPackage
	_ = imsmqquery2.GoPackage
	_ = imsmqquery3.GoPackage
	_ = imsmqquery4.GoPackage
	_ = imsmqmessage.GoPackage
	_ = imsmqmessage2.GoPackage
	_ = imsmqmessage3.GoPackage
	_ = imsmqmessage4.GoPackage
	_ = imsmqqueue.GoPackage
	_ = imsmqqueue2.GoPackage
	_ = imsmqqueue3.GoPackage
	_ = imsmqqueue4.GoPackage
	_ = imsmqprivateevent.GoPackage
	_ = imsmqevent2.GoPackage
	_ = imsmqqueueinfo.GoPackage
	_ = imsmqqueueinfo2.GoPackage
	_ = imsmqqueueinfo3.GoPackage
	_ = imsmqqueueinfo4.GoPackage
	_ = imsmqqueueinfos.GoPackage
	_ = imsmqqueueinfos2.GoPackage
	_ = imsmqqueueinfos3.GoPackage
	_ = imsmqqueueinfos4.GoPackage
	_ = imsmqtransaction.GoPackage
	_ = imsmqtransaction2.GoPackage
	_ = imsmqtransaction3.GoPackage
	_ = imsmqcoordinatedtransactiondispenser.GoPackage
	_ = imsmqcoordinatedtransactiondispenser2.GoPackage
	_ = imsmqcoordinatedtransactiondispenser3.GoPackage
	_ = imsmqtransactiondispenser.GoPackage
	_ = imsmqtransactiondispenser2.GoPackage
	_ = imsmqtransactiondispenser3.GoPackage
	_ = imsmqapplication.GoPackage
	_ = imsmqapplication2.GoPackage
	_ = imsmqapplication3.GoPackage
	_ = imsmqdestination.GoPackage
	_ = imsmqprivatedestination.GoPackage
	_ = imsmqcollection.GoPackage
	_ = imsmqmanagement.GoPackage
	_ = imsmqoutgoingqueuemanagement.GoPackage
	_ = imsmqqueuemanagement.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/mqac"
)

// dcom/mqac client set.
type Client interface {

	// DCOM common interfaces
	RemoteUnknown() iremunknown.RemoteUnknownClient
	RemoteUnknown2() iremunknown2.RemoteUnknown2Client

	// Package specific interfaces
	ITransaction() itransaction.ITransactionClient
	EnumConnections() ienumconnections.EnumConnectionsClient
	ConnectionPoint() iconnectionpoint.ConnectionPointClient
	EnumConnectionPoints() ienumconnectionpoints.EnumConnectionPointsClient
	ConnectionPointContainer() iconnectionpointcontainer.ConnectionPointContainerClient
	Query() imsmqquery.QueryClient
	Query2() imsmqquery2.Query2Client
	Query3() imsmqquery3.Query3Client
	Query4() imsmqquery4.Query4Client
	Message() imsmqmessage.MessageClient
	Message2() imsmqmessage2.Message2Client
	Message3() imsmqmessage3.Message3Client
	Message4() imsmqmessage4.Message4Client
	Queue() imsmqqueue.QueueClient
	Queue2() imsmqqueue2.Queue2Client
	Queue3() imsmqqueue3.Queue3Client
	Queue4() imsmqqueue4.Queue4Client
	PrivateEvent() imsmqprivateevent.PrivateEventClient
	Event2() imsmqevent2.Event2Client
	QueueInfo() imsmqqueueinfo.QueueInfoClient
	QueueInfo2() imsmqqueueinfo2.QueueInfo2Client
	QueueInfo3() imsmqqueueinfo3.QueueInfo3Client
	QueueInfo4() imsmqqueueinfo4.QueueInfo4Client
	QueueInfos() imsmqqueueinfos.QueueInfosClient
	QueueInfos2() imsmqqueueinfos2.QueueInfos2Client
	QueueInfos3() imsmqqueueinfos3.QueueInfos3Client
	QueueInfos4() imsmqqueueinfos4.QueueInfos4Client
	Transaction() imsmqtransaction.TransactionClient
	Transaction2() imsmqtransaction2.Transaction2Client
	Transaction3() imsmqtransaction3.Transaction3Client
	CoordinatedTransactionDispenser() imsmqcoordinatedtransactiondispenser.CoordinatedTransactionDispenserClient
	CoordinatedTransactionDispenser2() imsmqcoordinatedtransactiondispenser2.CoordinatedTransactionDispenser2Client
	CoordinatedTransactionDispenser3() imsmqcoordinatedtransactiondispenser3.CoordinatedTransactionDispenser3Client
	TransactionDispenser() imsmqtransactiondispenser.TransactionDispenserClient
	TransactionDispenser2() imsmqtransactiondispenser2.TransactionDispenser2Client
	TransactionDispenser3() imsmqtransactiondispenser3.TransactionDispenser3Client
	Application() imsmqapplication.ApplicationClient
	Application2() imsmqapplication2.Application2Client
	Application3() imsmqapplication3.Application3Client
	Destination() imsmqdestination.DestinationClient
	PrivateDestination() imsmqprivatedestination.PrivateDestinationClient
	Collection() imsmqcollection.CollectionClient
	Management() imsmqmanagement.ManagementClient
	OutgoingQueueManagement() imsmqoutgoingqueuemanagement.OutgoingQueueManagementClient
	QueueManagement() imsmqqueuemanagement.QueueManagementClient
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

	itransaction                     itransaction.ITransactionClient
	enumConnections                  ienumconnections.EnumConnectionsClient
	connectionPoint                  iconnectionpoint.ConnectionPointClient
	enumConnectionPoints             ienumconnectionpoints.EnumConnectionPointsClient
	connectionPointContainer         iconnectionpointcontainer.ConnectionPointContainerClient
	query                            imsmqquery.QueryClient
	query2                           imsmqquery2.Query2Client
	query3                           imsmqquery3.Query3Client
	query4                           imsmqquery4.Query4Client
	message                          imsmqmessage.MessageClient
	message2                         imsmqmessage2.Message2Client
	message3                         imsmqmessage3.Message3Client
	message4                         imsmqmessage4.Message4Client
	queue                            imsmqqueue.QueueClient
	queue2                           imsmqqueue2.Queue2Client
	queue3                           imsmqqueue3.Queue3Client
	queue4                           imsmqqueue4.Queue4Client
	privateEvent                     imsmqprivateevent.PrivateEventClient
	event2                           imsmqevent2.Event2Client
	queueInfo                        imsmqqueueinfo.QueueInfoClient
	queueInfo2                       imsmqqueueinfo2.QueueInfo2Client
	queueInfo3                       imsmqqueueinfo3.QueueInfo3Client
	queueInfo4                       imsmqqueueinfo4.QueueInfo4Client
	queueInfos                       imsmqqueueinfos.QueueInfosClient
	queueInfos2                      imsmqqueueinfos2.QueueInfos2Client
	queueInfos3                      imsmqqueueinfos3.QueueInfos3Client
	queueInfos4                      imsmqqueueinfos4.QueueInfos4Client
	transaction                      imsmqtransaction.TransactionClient
	transaction2                     imsmqtransaction2.Transaction2Client
	transaction3                     imsmqtransaction3.Transaction3Client
	coordinatedTransactionDispenser  imsmqcoordinatedtransactiondispenser.CoordinatedTransactionDispenserClient
	coordinatedTransactionDispenser2 imsmqcoordinatedtransactiondispenser2.CoordinatedTransactionDispenser2Client
	coordinatedTransactionDispenser3 imsmqcoordinatedtransactiondispenser3.CoordinatedTransactionDispenser3Client
	transactionDispenser             imsmqtransactiondispenser.TransactionDispenserClient
	transactionDispenser2            imsmqtransactiondispenser2.TransactionDispenser2Client
	transactionDispenser3            imsmqtransactiondispenser3.TransactionDispenser3Client
	application                      imsmqapplication.ApplicationClient
	application2                     imsmqapplication2.Application2Client
	application3                     imsmqapplication3.Application3Client
	destination                      imsmqdestination.DestinationClient
	privateDestination               imsmqprivatedestination.PrivateDestinationClient
	collection                       imsmqcollection.CollectionClient
	management                       imsmqmanagement.ManagementClient
	outgoingQueueManagement          imsmqoutgoingqueuemanagement.OutgoingQueueManagementClient
	queueManagement                  imsmqqueuemanagement.QueueManagementClient
}

func (o *xxx_DefaultClient) RemoteUnknown() iremunknown.RemoteUnknownClient {
	return o.dcomClient.RemoteUnknown()
}

func (o *xxx_DefaultClient) RemoteUnknown2() iremunknown2.RemoteUnknown2Client {
	return o.dcomClient.RemoteUnknown2()
}

func (o *xxx_DefaultClient) ITransaction() itransaction.ITransactionClient {
	return o.itransaction
}

func (o *xxx_DefaultClient) EnumConnections() ienumconnections.EnumConnectionsClient {
	return o.enumConnections
}

func (o *xxx_DefaultClient) ConnectionPoint() iconnectionpoint.ConnectionPointClient {
	return o.connectionPoint
}

func (o *xxx_DefaultClient) EnumConnectionPoints() ienumconnectionpoints.EnumConnectionPointsClient {
	return o.enumConnectionPoints
}

func (o *xxx_DefaultClient) ConnectionPointContainer() iconnectionpointcontainer.ConnectionPointContainerClient {
	return o.connectionPointContainer
}

func (o *xxx_DefaultClient) Query() imsmqquery.QueryClient {
	return o.query
}

func (o *xxx_DefaultClient) Query2() imsmqquery2.Query2Client {
	return o.query2
}

func (o *xxx_DefaultClient) Query3() imsmqquery3.Query3Client {
	return o.query3
}

func (o *xxx_DefaultClient) Query4() imsmqquery4.Query4Client {
	return o.query4
}

func (o *xxx_DefaultClient) Message() imsmqmessage.MessageClient {
	return o.message
}

func (o *xxx_DefaultClient) Message2() imsmqmessage2.Message2Client {
	return o.message2
}

func (o *xxx_DefaultClient) Message3() imsmqmessage3.Message3Client {
	return o.message3
}

func (o *xxx_DefaultClient) Message4() imsmqmessage4.Message4Client {
	return o.message4
}

func (o *xxx_DefaultClient) Queue() imsmqqueue.QueueClient {
	return o.queue
}

func (o *xxx_DefaultClient) Queue2() imsmqqueue2.Queue2Client {
	return o.queue2
}

func (o *xxx_DefaultClient) Queue3() imsmqqueue3.Queue3Client {
	return o.queue3
}

func (o *xxx_DefaultClient) Queue4() imsmqqueue4.Queue4Client {
	return o.queue4
}

func (o *xxx_DefaultClient) PrivateEvent() imsmqprivateevent.PrivateEventClient {
	return o.privateEvent
}

func (o *xxx_DefaultClient) Event2() imsmqevent2.Event2Client {
	return o.event2
}

func (o *xxx_DefaultClient) QueueInfo() imsmqqueueinfo.QueueInfoClient {
	return o.queueInfo
}

func (o *xxx_DefaultClient) QueueInfo2() imsmqqueueinfo2.QueueInfo2Client {
	return o.queueInfo2
}

func (o *xxx_DefaultClient) QueueInfo3() imsmqqueueinfo3.QueueInfo3Client {
	return o.queueInfo3
}

func (o *xxx_DefaultClient) QueueInfo4() imsmqqueueinfo4.QueueInfo4Client {
	return o.queueInfo4
}

func (o *xxx_DefaultClient) QueueInfos() imsmqqueueinfos.QueueInfosClient {
	return o.queueInfos
}

func (o *xxx_DefaultClient) QueueInfos2() imsmqqueueinfos2.QueueInfos2Client {
	return o.queueInfos2
}

func (o *xxx_DefaultClient) QueueInfos3() imsmqqueueinfos3.QueueInfos3Client {
	return o.queueInfos3
}

func (o *xxx_DefaultClient) QueueInfos4() imsmqqueueinfos4.QueueInfos4Client {
	return o.queueInfos4
}

func (o *xxx_DefaultClient) Transaction() imsmqtransaction.TransactionClient {
	return o.transaction
}

func (o *xxx_DefaultClient) Transaction2() imsmqtransaction2.Transaction2Client {
	return o.transaction2
}

func (o *xxx_DefaultClient) Transaction3() imsmqtransaction3.Transaction3Client {
	return o.transaction3
}

func (o *xxx_DefaultClient) CoordinatedTransactionDispenser() imsmqcoordinatedtransactiondispenser.CoordinatedTransactionDispenserClient {
	return o.coordinatedTransactionDispenser
}

func (o *xxx_DefaultClient) CoordinatedTransactionDispenser2() imsmqcoordinatedtransactiondispenser2.CoordinatedTransactionDispenser2Client {
	return o.coordinatedTransactionDispenser2
}

func (o *xxx_DefaultClient) CoordinatedTransactionDispenser3() imsmqcoordinatedtransactiondispenser3.CoordinatedTransactionDispenser3Client {
	return o.coordinatedTransactionDispenser3
}

func (o *xxx_DefaultClient) TransactionDispenser() imsmqtransactiondispenser.TransactionDispenserClient {
	return o.transactionDispenser
}

func (o *xxx_DefaultClient) TransactionDispenser2() imsmqtransactiondispenser2.TransactionDispenser2Client {
	return o.transactionDispenser2
}

func (o *xxx_DefaultClient) TransactionDispenser3() imsmqtransactiondispenser3.TransactionDispenser3Client {
	return o.transactionDispenser3
}

func (o *xxx_DefaultClient) Application() imsmqapplication.ApplicationClient {
	return o.application
}

func (o *xxx_DefaultClient) Application2() imsmqapplication2.Application2Client {
	return o.application2
}

func (o *xxx_DefaultClient) Application3() imsmqapplication3.Application3Client {
	return o.application3
}

func (o *xxx_DefaultClient) Destination() imsmqdestination.DestinationClient {
	return o.destination
}

func (o *xxx_DefaultClient) PrivateDestination() imsmqprivatedestination.PrivateDestinationClient {
	return o.privateDestination
}

func (o *xxx_DefaultClient) Collection() imsmqcollection.CollectionClient {
	return o.collection
}

func (o *xxx_DefaultClient) Management() imsmqmanagement.ManagementClient {
	return o.management
}

func (o *xxx_DefaultClient) OutgoingQueueManagement() imsmqoutgoingqueuemanagement.OutgoingQueueManagementClient {
	return o.outgoingQueueManagement
}

func (o *xxx_DefaultClient) QueueManagement() imsmqqueuemanagement.QueueManagementClient {
	return o.queueManagement
}

func NewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Client, error) {

	opts = append(opts,
		dcerpc.WithAbstractSyntax(itransaction.ITransactionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ienumconnections.EnumConnectionsSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iconnectionpoint.ConnectionPointSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ienumconnectionpoints.EnumConnectionPointsSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iconnectionpointcontainer.ConnectionPointContainerSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqquery.QuerySyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqquery2.Query2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqquery3.Query3SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqquery4.Query4SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqmessage.MessageSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqmessage2.Message2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqmessage3.Message3SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqmessage4.Message4SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqqueue.QueueSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqqueue2.Queue2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqqueue3.Queue3SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqqueue4.Queue4SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqprivateevent.PrivateEventSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqevent2.Event2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqqueueinfo.QueueInfoSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqqueueinfo2.QueueInfo2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqqueueinfo3.QueueInfo3SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqqueueinfo4.QueueInfo4SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqqueueinfos.QueueInfosSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqqueueinfos2.QueueInfos2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqqueueinfos3.QueueInfos3SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqqueueinfos4.QueueInfos4SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqtransaction.TransactionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqtransaction2.Transaction2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqtransaction3.Transaction3SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqcoordinatedtransactiondispenser.CoordinatedTransactionDispenserSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqcoordinatedtransactiondispenser2.CoordinatedTransactionDispenser2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqcoordinatedtransactiondispenser3.CoordinatedTransactionDispenser3SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqtransactiondispenser.TransactionDispenserSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqtransactiondispenser2.TransactionDispenser2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqtransactiondispenser3.TransactionDispenser3SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqapplication.ApplicationSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqapplication2.Application2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqapplication3.Application3SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqdestination.DestinationSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqprivatedestination.PrivateDestinationSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqcollection.CollectionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqmanagement.ManagementSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqoutgoingqueuemanagement.OutgoingQueueManagementSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqqueuemanagement.QueueManagementSyntaxV0_0),
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

	itransactionSubConn, err := sub.SubConn(ctx, itransaction.ITransactionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		itransactionSubConn = sub
	}

	o.itransaction, err = itransaction.NewITransactionClient(ctx, itransactionSubConn, append(opts, dcerpc.WithNoBind(itransactionSubConn))...)

	enumConnectionsSubConn, err := sub.SubConn(ctx, ienumconnections.EnumConnectionsSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		enumConnectionsSubConn = sub
	}

	o.enumConnections, err = ienumconnections.NewEnumConnectionsClient(ctx, enumConnectionsSubConn, append(opts, dcerpc.WithNoBind(enumConnectionsSubConn))...)

	connectionPointSubConn, err := sub.SubConn(ctx, iconnectionpoint.ConnectionPointSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		connectionPointSubConn = sub
	}

	o.connectionPoint, err = iconnectionpoint.NewConnectionPointClient(ctx, connectionPointSubConn, append(opts, dcerpc.WithNoBind(connectionPointSubConn))...)

	enumConnectionPointsSubConn, err := sub.SubConn(ctx, ienumconnectionpoints.EnumConnectionPointsSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		enumConnectionPointsSubConn = sub
	}

	o.enumConnectionPoints, err = ienumconnectionpoints.NewEnumConnectionPointsClient(ctx, enumConnectionPointsSubConn, append(opts, dcerpc.WithNoBind(enumConnectionPointsSubConn))...)

	connectionPointContainerSubConn, err := sub.SubConn(ctx, iconnectionpointcontainer.ConnectionPointContainerSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		connectionPointContainerSubConn = sub
	}

	o.connectionPointContainer, err = iconnectionpointcontainer.NewConnectionPointContainerClient(ctx, connectionPointContainerSubConn, append(opts, dcerpc.WithNoBind(connectionPointContainerSubConn))...)

	querySubConn, err := sub.SubConn(ctx, imsmqquery.QuerySyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		querySubConn = sub
	}

	o.query, err = imsmqquery.NewQueryClient(ctx, querySubConn, append(opts, dcerpc.WithNoBind(querySubConn))...)

	query2SubConn, err := sub.SubConn(ctx, imsmqquery2.Query2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		query2SubConn = sub
	}

	o.query2, err = imsmqquery2.NewQuery2Client(ctx, query2SubConn, append(opts, dcerpc.WithNoBind(query2SubConn))...)

	query3SubConn, err := sub.SubConn(ctx, imsmqquery3.Query3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		query3SubConn = sub
	}

	o.query3, err = imsmqquery3.NewQuery3Client(ctx, query3SubConn, append(opts, dcerpc.WithNoBind(query3SubConn))...)

	query4SubConn, err := sub.SubConn(ctx, imsmqquery4.Query4SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		query4SubConn = sub
	}

	o.query4, err = imsmqquery4.NewQuery4Client(ctx, query4SubConn, append(opts, dcerpc.WithNoBind(query4SubConn))...)

	messageSubConn, err := sub.SubConn(ctx, imsmqmessage.MessageSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		messageSubConn = sub
	}

	o.message, err = imsmqmessage.NewMessageClient(ctx, messageSubConn, append(opts, dcerpc.WithNoBind(messageSubConn))...)

	message2SubConn, err := sub.SubConn(ctx, imsmqmessage2.Message2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		message2SubConn = sub
	}

	o.message2, err = imsmqmessage2.NewMessage2Client(ctx, message2SubConn, append(opts, dcerpc.WithNoBind(message2SubConn))...)

	message3SubConn, err := sub.SubConn(ctx, imsmqmessage3.Message3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		message3SubConn = sub
	}

	o.message3, err = imsmqmessage3.NewMessage3Client(ctx, message3SubConn, append(opts, dcerpc.WithNoBind(message3SubConn))...)

	message4SubConn, err := sub.SubConn(ctx, imsmqmessage4.Message4SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		message4SubConn = sub
	}

	o.message4, err = imsmqmessage4.NewMessage4Client(ctx, message4SubConn, append(opts, dcerpc.WithNoBind(message4SubConn))...)

	queueSubConn, err := sub.SubConn(ctx, imsmqqueue.QueueSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		queueSubConn = sub
	}

	o.queue, err = imsmqqueue.NewQueueClient(ctx, queueSubConn, append(opts, dcerpc.WithNoBind(queueSubConn))...)

	queue2SubConn, err := sub.SubConn(ctx, imsmqqueue2.Queue2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		queue2SubConn = sub
	}

	o.queue2, err = imsmqqueue2.NewQueue2Client(ctx, queue2SubConn, append(opts, dcerpc.WithNoBind(queue2SubConn))...)

	queue3SubConn, err := sub.SubConn(ctx, imsmqqueue3.Queue3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		queue3SubConn = sub
	}

	o.queue3, err = imsmqqueue3.NewQueue3Client(ctx, queue3SubConn, append(opts, dcerpc.WithNoBind(queue3SubConn))...)

	queue4SubConn, err := sub.SubConn(ctx, imsmqqueue4.Queue4SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		queue4SubConn = sub
	}

	o.queue4, err = imsmqqueue4.NewQueue4Client(ctx, queue4SubConn, append(opts, dcerpc.WithNoBind(queue4SubConn))...)

	privateEventSubConn, err := sub.SubConn(ctx, imsmqprivateevent.PrivateEventSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		privateEventSubConn = sub
	}

	o.privateEvent, err = imsmqprivateevent.NewPrivateEventClient(ctx, privateEventSubConn, append(opts, dcerpc.WithNoBind(privateEventSubConn))...)

	event2SubConn, err := sub.SubConn(ctx, imsmqevent2.Event2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		event2SubConn = sub
	}

	o.event2, err = imsmqevent2.NewEvent2Client(ctx, event2SubConn, append(opts, dcerpc.WithNoBind(event2SubConn))...)

	queueInfoSubConn, err := sub.SubConn(ctx, imsmqqueueinfo.QueueInfoSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		queueInfoSubConn = sub
	}

	o.queueInfo, err = imsmqqueueinfo.NewQueueInfoClient(ctx, queueInfoSubConn, append(opts, dcerpc.WithNoBind(queueInfoSubConn))...)

	queueInfo2SubConn, err := sub.SubConn(ctx, imsmqqueueinfo2.QueueInfo2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		queueInfo2SubConn = sub
	}

	o.queueInfo2, err = imsmqqueueinfo2.NewQueueInfo2Client(ctx, queueInfo2SubConn, append(opts, dcerpc.WithNoBind(queueInfo2SubConn))...)

	queueInfo3SubConn, err := sub.SubConn(ctx, imsmqqueueinfo3.QueueInfo3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		queueInfo3SubConn = sub
	}

	o.queueInfo3, err = imsmqqueueinfo3.NewQueueInfo3Client(ctx, queueInfo3SubConn, append(opts, dcerpc.WithNoBind(queueInfo3SubConn))...)

	queueInfo4SubConn, err := sub.SubConn(ctx, imsmqqueueinfo4.QueueInfo4SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		queueInfo4SubConn = sub
	}

	o.queueInfo4, err = imsmqqueueinfo4.NewQueueInfo4Client(ctx, queueInfo4SubConn, append(opts, dcerpc.WithNoBind(queueInfo4SubConn))...)

	queueInfosSubConn, err := sub.SubConn(ctx, imsmqqueueinfos.QueueInfosSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		queueInfosSubConn = sub
	}

	o.queueInfos, err = imsmqqueueinfos.NewQueueInfosClient(ctx, queueInfosSubConn, append(opts, dcerpc.WithNoBind(queueInfosSubConn))...)

	queueInfos2SubConn, err := sub.SubConn(ctx, imsmqqueueinfos2.QueueInfos2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		queueInfos2SubConn = sub
	}

	o.queueInfos2, err = imsmqqueueinfos2.NewQueueInfos2Client(ctx, queueInfos2SubConn, append(opts, dcerpc.WithNoBind(queueInfos2SubConn))...)

	queueInfos3SubConn, err := sub.SubConn(ctx, imsmqqueueinfos3.QueueInfos3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		queueInfos3SubConn = sub
	}

	o.queueInfos3, err = imsmqqueueinfos3.NewQueueInfos3Client(ctx, queueInfos3SubConn, append(opts, dcerpc.WithNoBind(queueInfos3SubConn))...)

	queueInfos4SubConn, err := sub.SubConn(ctx, imsmqqueueinfos4.QueueInfos4SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		queueInfos4SubConn = sub
	}

	o.queueInfos4, err = imsmqqueueinfos4.NewQueueInfos4Client(ctx, queueInfos4SubConn, append(opts, dcerpc.WithNoBind(queueInfos4SubConn))...)

	transactionSubConn, err := sub.SubConn(ctx, imsmqtransaction.TransactionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		transactionSubConn = sub
	}

	o.transaction, err = imsmqtransaction.NewTransactionClient(ctx, transactionSubConn, append(opts, dcerpc.WithNoBind(transactionSubConn))...)

	transaction2SubConn, err := sub.SubConn(ctx, imsmqtransaction2.Transaction2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		transaction2SubConn = sub
	}

	o.transaction2, err = imsmqtransaction2.NewTransaction2Client(ctx, transaction2SubConn, append(opts, dcerpc.WithNoBind(transaction2SubConn))...)

	transaction3SubConn, err := sub.SubConn(ctx, imsmqtransaction3.Transaction3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		transaction3SubConn = sub
	}

	o.transaction3, err = imsmqtransaction3.NewTransaction3Client(ctx, transaction3SubConn, append(opts, dcerpc.WithNoBind(transaction3SubConn))...)

	coordinatedTransactionDispenserSubConn, err := sub.SubConn(ctx, imsmqcoordinatedtransactiondispenser.CoordinatedTransactionDispenserSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		coordinatedTransactionDispenserSubConn = sub
	}

	o.coordinatedTransactionDispenser, err = imsmqcoordinatedtransactiondispenser.NewCoordinatedTransactionDispenserClient(ctx, coordinatedTransactionDispenserSubConn, append(opts, dcerpc.WithNoBind(coordinatedTransactionDispenserSubConn))...)

	coordinatedTransactionDispenser2SubConn, err := sub.SubConn(ctx, imsmqcoordinatedtransactiondispenser2.CoordinatedTransactionDispenser2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		coordinatedTransactionDispenser2SubConn = sub
	}

	o.coordinatedTransactionDispenser2, err = imsmqcoordinatedtransactiondispenser2.NewCoordinatedTransactionDispenser2Client(ctx, coordinatedTransactionDispenser2SubConn, append(opts, dcerpc.WithNoBind(coordinatedTransactionDispenser2SubConn))...)

	coordinatedTransactionDispenser3SubConn, err := sub.SubConn(ctx, imsmqcoordinatedtransactiondispenser3.CoordinatedTransactionDispenser3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		coordinatedTransactionDispenser3SubConn = sub
	}

	o.coordinatedTransactionDispenser3, err = imsmqcoordinatedtransactiondispenser3.NewCoordinatedTransactionDispenser3Client(ctx, coordinatedTransactionDispenser3SubConn, append(opts, dcerpc.WithNoBind(coordinatedTransactionDispenser3SubConn))...)

	transactionDispenserSubConn, err := sub.SubConn(ctx, imsmqtransactiondispenser.TransactionDispenserSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		transactionDispenserSubConn = sub
	}

	o.transactionDispenser, err = imsmqtransactiondispenser.NewTransactionDispenserClient(ctx, transactionDispenserSubConn, append(opts, dcerpc.WithNoBind(transactionDispenserSubConn))...)

	transactionDispenser2SubConn, err := sub.SubConn(ctx, imsmqtransactiondispenser2.TransactionDispenser2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		transactionDispenser2SubConn = sub
	}

	o.transactionDispenser2, err = imsmqtransactiondispenser2.NewTransactionDispenser2Client(ctx, transactionDispenser2SubConn, append(opts, dcerpc.WithNoBind(transactionDispenser2SubConn))...)

	transactionDispenser3SubConn, err := sub.SubConn(ctx, imsmqtransactiondispenser3.TransactionDispenser3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		transactionDispenser3SubConn = sub
	}

	o.transactionDispenser3, err = imsmqtransactiondispenser3.NewTransactionDispenser3Client(ctx, transactionDispenser3SubConn, append(opts, dcerpc.WithNoBind(transactionDispenser3SubConn))...)

	applicationSubConn, err := sub.SubConn(ctx, imsmqapplication.ApplicationSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		applicationSubConn = sub
	}

	o.application, err = imsmqapplication.NewApplicationClient(ctx, applicationSubConn, append(opts, dcerpc.WithNoBind(applicationSubConn))...)

	application2SubConn, err := sub.SubConn(ctx, imsmqapplication2.Application2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		application2SubConn = sub
	}

	o.application2, err = imsmqapplication2.NewApplication2Client(ctx, application2SubConn, append(opts, dcerpc.WithNoBind(application2SubConn))...)

	application3SubConn, err := sub.SubConn(ctx, imsmqapplication3.Application3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		application3SubConn = sub
	}

	o.application3, err = imsmqapplication3.NewApplication3Client(ctx, application3SubConn, append(opts, dcerpc.WithNoBind(application3SubConn))...)

	destinationSubConn, err := sub.SubConn(ctx, imsmqdestination.DestinationSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		destinationSubConn = sub
	}

	o.destination, err = imsmqdestination.NewDestinationClient(ctx, destinationSubConn, append(opts, dcerpc.WithNoBind(destinationSubConn))...)

	privateDestinationSubConn, err := sub.SubConn(ctx, imsmqprivatedestination.PrivateDestinationSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		privateDestinationSubConn = sub
	}

	o.privateDestination, err = imsmqprivatedestination.NewPrivateDestinationClient(ctx, privateDestinationSubConn, append(opts, dcerpc.WithNoBind(privateDestinationSubConn))...)

	collectionSubConn, err := sub.SubConn(ctx, imsmqcollection.CollectionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		collectionSubConn = sub
	}

	o.collection, err = imsmqcollection.NewCollectionClient(ctx, collectionSubConn, append(opts, dcerpc.WithNoBind(collectionSubConn))...)

	managementSubConn, err := sub.SubConn(ctx, imsmqmanagement.ManagementSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		managementSubConn = sub
	}

	o.management, err = imsmqmanagement.NewManagementClient(ctx, managementSubConn, append(opts, dcerpc.WithNoBind(managementSubConn))...)

	outgoingQueueManagementSubConn, err := sub.SubConn(ctx, imsmqoutgoingqueuemanagement.OutgoingQueueManagementSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		outgoingQueueManagementSubConn = sub
	}

	o.outgoingQueueManagement, err = imsmqoutgoingqueuemanagement.NewOutgoingQueueManagementClient(ctx, outgoingQueueManagementSubConn, append(opts, dcerpc.WithNoBind(outgoingQueueManagementSubConn))...)

	queueManagementSubConn, err := sub.SubConn(ctx, imsmqqueuemanagement.QueueManagementSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		queueManagementSubConn = sub
	}

	o.queueManagement, err = imsmqqueuemanagement.NewQueueManagementClient(ctx, queueManagementSubConn, append(opts, dcerpc.WithNoBind(queueManagementSubConn))...)
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
		dcomClient:                       o.dcomClient.IPID(ctx, ipid),
		itransaction:                     o.itransaction.IPID(ctx, ipid),
		enumConnections:                  o.enumConnections.IPID(ctx, ipid),
		connectionPoint:                  o.connectionPoint.IPID(ctx, ipid),
		enumConnectionPoints:             o.enumConnectionPoints.IPID(ctx, ipid),
		connectionPointContainer:         o.connectionPointContainer.IPID(ctx, ipid),
		query:                            o.query.IPID(ctx, ipid),
		query2:                           o.query2.IPID(ctx, ipid),
		query3:                           o.query3.IPID(ctx, ipid),
		query4:                           o.query4.IPID(ctx, ipid),
		message:                          o.message.IPID(ctx, ipid),
		message2:                         o.message2.IPID(ctx, ipid),
		message3:                         o.message3.IPID(ctx, ipid),
		message4:                         o.message4.IPID(ctx, ipid),
		queue:                            o.queue.IPID(ctx, ipid),
		queue2:                           o.queue2.IPID(ctx, ipid),
		queue3:                           o.queue3.IPID(ctx, ipid),
		queue4:                           o.queue4.IPID(ctx, ipid),
		privateEvent:                     o.privateEvent.IPID(ctx, ipid),
		event2:                           o.event2.IPID(ctx, ipid),
		queueInfo:                        o.queueInfo.IPID(ctx, ipid),
		queueInfo2:                       o.queueInfo2.IPID(ctx, ipid),
		queueInfo3:                       o.queueInfo3.IPID(ctx, ipid),
		queueInfo4:                       o.queueInfo4.IPID(ctx, ipid),
		queueInfos:                       o.queueInfos.IPID(ctx, ipid),
		queueInfos2:                      o.queueInfos2.IPID(ctx, ipid),
		queueInfos3:                      o.queueInfos3.IPID(ctx, ipid),
		queueInfos4:                      o.queueInfos4.IPID(ctx, ipid),
		transaction:                      o.transaction.IPID(ctx, ipid),
		transaction2:                     o.transaction2.IPID(ctx, ipid),
		transaction3:                     o.transaction3.IPID(ctx, ipid),
		coordinatedTransactionDispenser:  o.coordinatedTransactionDispenser.IPID(ctx, ipid),
		coordinatedTransactionDispenser2: o.coordinatedTransactionDispenser2.IPID(ctx, ipid),
		coordinatedTransactionDispenser3: o.coordinatedTransactionDispenser3.IPID(ctx, ipid),
		transactionDispenser:             o.transactionDispenser.IPID(ctx, ipid),
		transactionDispenser2:            o.transactionDispenser2.IPID(ctx, ipid),
		transactionDispenser3:            o.transactionDispenser3.IPID(ctx, ipid),
		application:                      o.application.IPID(ctx, ipid),
		application2:                     o.application2.IPID(ctx, ipid),
		application3:                     o.application3.IPID(ctx, ipid),
		destination:                      o.destination.IPID(ctx, ipid),
		privateDestination:               o.privateDestination.IPID(ctx, ipid),
		collection:                       o.collection.IPID(ctx, ipid),
		management:                       o.management.IPID(ctx, ipid),
		outgoingQueueManagement:          o.outgoingQueueManagement.IPID(ctx, ipid),
		queueManagement:                  o.queueManagement.IPID(ctx, ipid),
		cc:                               o.cc,
	}
}
