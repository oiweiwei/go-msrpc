package mqac

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
	_ = (*errors.Error)(nil)
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
	ImsmqQuery() imsmqquery.ImsmqQueryClient
	ImsmqQuery2() imsmqquery2.ImsmqQuery2Client
	ImsmqQuery3() imsmqquery3.ImsmqQuery3Client
	ImsmqQuery4() imsmqquery4.ImsmqQuery4Client
	ImsmqMessage() imsmqmessage.ImsmqMessageClient
	ImsmqMessage2() imsmqmessage2.ImsmqMessage2Client
	ImsmqMessage3() imsmqmessage3.ImsmqMessage3Client
	ImsmqMessage4() imsmqmessage4.ImsmqMessage4Client
	ImsmqQueue() imsmqqueue.ImsmqQueueClient
	ImsmqQueue2() imsmqqueue2.ImsmqQueue2Client
	ImsmqQueue3() imsmqqueue3.ImsmqQueue3Client
	ImsmqQueue4() imsmqqueue4.ImsmqQueue4Client
	ImsmqPrivateEvent() imsmqprivateevent.ImsmqPrivateEventClient
	ImsmqEvent2() imsmqevent2.ImsmqEvent2Client
	ImsmqQueueInfo() imsmqqueueinfo.ImsmqQueueInfoClient
	ImsmqQueueInfo2() imsmqqueueinfo2.ImsmqQueueInfo2Client
	ImsmqQueueInfo3() imsmqqueueinfo3.ImsmqQueueInfo3Client
	ImsmqQueueInfo4() imsmqqueueinfo4.ImsmqQueueInfo4Client
	ImsmqQueueInfos() imsmqqueueinfos.ImsmqQueueInfosClient
	ImsmqQueueInfos2() imsmqqueueinfos2.ImsmqQueueInfos2Client
	ImsmqQueueInfos3() imsmqqueueinfos3.ImsmqQueueInfos3Client
	ImsmqQueueInfos4() imsmqqueueinfos4.ImsmqQueueInfos4Client
	ImsmqTransaction() imsmqtransaction.ImsmqTransactionClient
	ImsmqTransaction2() imsmqtransaction2.ImsmqTransaction2Client
	ImsmqTransaction3() imsmqtransaction3.ImsmqTransaction3Client
	ImsmqCoordinatedTransactionDispenser() imsmqcoordinatedtransactiondispenser.ImsmqCoordinatedTransactionDispenserClient
	ImsmqCoordinatedTransactionDispenser2() imsmqcoordinatedtransactiondispenser2.ImsmqCoordinatedTransactionDispenser2Client
	ImsmqCoordinatedTransactionDispenser3() imsmqcoordinatedtransactiondispenser3.ImsmqCoordinatedTransactionDispenser3Client
	ImsmqTransactionDispenser() imsmqtransactiondispenser.ImsmqTransactionDispenserClient
	ImsmqTransactionDispenser2() imsmqtransactiondispenser2.ImsmqTransactionDispenser2Client
	ImsmqTransactionDispenser3() imsmqtransactiondispenser3.ImsmqTransactionDispenser3Client
	ImsmqApplication() imsmqapplication.ImsmqApplicationClient
	ImsmqApplication2() imsmqapplication2.ImsmqApplication2Client
	ImsmqApplication3() imsmqapplication3.ImsmqApplication3Client
	ImsmqDestination() imsmqdestination.ImsmqDestinationClient
	ImsmqPrivateDestination() imsmqprivatedestination.ImsmqPrivateDestinationClient
	ImsmqCollection() imsmqcollection.ImsmqCollectionClient
	ImsmqManagement() imsmqmanagement.ImsmqManagementClient
	ImsmqOutgoingQueueManagement() imsmqoutgoingqueuemanagement.ImsmqOutgoingQueueManagementClient
	ImsmqQueueManagement() imsmqqueuemanagement.ImsmqQueueManagementClient
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

	itransaction                          itransaction.ITransactionClient
	enumConnections                       ienumconnections.EnumConnectionsClient
	connectionPoint                       iconnectionpoint.ConnectionPointClient
	enumConnectionPoints                  ienumconnectionpoints.EnumConnectionPointsClient
	connectionPointContainer              iconnectionpointcontainer.ConnectionPointContainerClient
	imsmqQuery                            imsmqquery.ImsmqQueryClient
	imsmqQuery2                           imsmqquery2.ImsmqQuery2Client
	imsmqQuery3                           imsmqquery3.ImsmqQuery3Client
	imsmqQuery4                           imsmqquery4.ImsmqQuery4Client
	imsmqMessage                          imsmqmessage.ImsmqMessageClient
	imsmqMessage2                         imsmqmessage2.ImsmqMessage2Client
	imsmqMessage3                         imsmqmessage3.ImsmqMessage3Client
	imsmqMessage4                         imsmqmessage4.ImsmqMessage4Client
	imsmqQueue                            imsmqqueue.ImsmqQueueClient
	imsmqQueue2                           imsmqqueue2.ImsmqQueue2Client
	imsmqQueue3                           imsmqqueue3.ImsmqQueue3Client
	imsmqQueue4                           imsmqqueue4.ImsmqQueue4Client
	imsmqPrivateEvent                     imsmqprivateevent.ImsmqPrivateEventClient
	imsmqEvent2                           imsmqevent2.ImsmqEvent2Client
	imsmqQueueInfo                        imsmqqueueinfo.ImsmqQueueInfoClient
	imsmqQueueInfo2                       imsmqqueueinfo2.ImsmqQueueInfo2Client
	imsmqQueueInfo3                       imsmqqueueinfo3.ImsmqQueueInfo3Client
	imsmqQueueInfo4                       imsmqqueueinfo4.ImsmqQueueInfo4Client
	imsmqQueueInfos                       imsmqqueueinfos.ImsmqQueueInfosClient
	imsmqQueueInfos2                      imsmqqueueinfos2.ImsmqQueueInfos2Client
	imsmqQueueInfos3                      imsmqqueueinfos3.ImsmqQueueInfos3Client
	imsmqQueueInfos4                      imsmqqueueinfos4.ImsmqQueueInfos4Client
	imsmqTransaction                      imsmqtransaction.ImsmqTransactionClient
	imsmqTransaction2                     imsmqtransaction2.ImsmqTransaction2Client
	imsmqTransaction3                     imsmqtransaction3.ImsmqTransaction3Client
	imsmqCoordinatedTransactionDispenser  imsmqcoordinatedtransactiondispenser.ImsmqCoordinatedTransactionDispenserClient
	imsmqCoordinatedTransactionDispenser2 imsmqcoordinatedtransactiondispenser2.ImsmqCoordinatedTransactionDispenser2Client
	imsmqCoordinatedTransactionDispenser3 imsmqcoordinatedtransactiondispenser3.ImsmqCoordinatedTransactionDispenser3Client
	imsmqTransactionDispenser             imsmqtransactiondispenser.ImsmqTransactionDispenserClient
	imsmqTransactionDispenser2            imsmqtransactiondispenser2.ImsmqTransactionDispenser2Client
	imsmqTransactionDispenser3            imsmqtransactiondispenser3.ImsmqTransactionDispenser3Client
	imsmqApplication                      imsmqapplication.ImsmqApplicationClient
	imsmqApplication2                     imsmqapplication2.ImsmqApplication2Client
	imsmqApplication3                     imsmqapplication3.ImsmqApplication3Client
	imsmqDestination                      imsmqdestination.ImsmqDestinationClient
	imsmqPrivateDestination               imsmqprivatedestination.ImsmqPrivateDestinationClient
	imsmqCollection                       imsmqcollection.ImsmqCollectionClient
	imsmqManagement                       imsmqmanagement.ImsmqManagementClient
	imsmqOutgoingQueueManagement          imsmqoutgoingqueuemanagement.ImsmqOutgoingQueueManagementClient
	imsmqQueueManagement                  imsmqqueuemanagement.ImsmqQueueManagementClient
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

func (o *xxx_DefaultClient) ImsmqQuery() imsmqquery.ImsmqQueryClient {
	return o.imsmqQuery
}

func (o *xxx_DefaultClient) ImsmqQuery2() imsmqquery2.ImsmqQuery2Client {
	return o.imsmqQuery2
}

func (o *xxx_DefaultClient) ImsmqQuery3() imsmqquery3.ImsmqQuery3Client {
	return o.imsmqQuery3
}

func (o *xxx_DefaultClient) ImsmqQuery4() imsmqquery4.ImsmqQuery4Client {
	return o.imsmqQuery4
}

func (o *xxx_DefaultClient) ImsmqMessage() imsmqmessage.ImsmqMessageClient {
	return o.imsmqMessage
}

func (o *xxx_DefaultClient) ImsmqMessage2() imsmqmessage2.ImsmqMessage2Client {
	return o.imsmqMessage2
}

func (o *xxx_DefaultClient) ImsmqMessage3() imsmqmessage3.ImsmqMessage3Client {
	return o.imsmqMessage3
}

func (o *xxx_DefaultClient) ImsmqMessage4() imsmqmessage4.ImsmqMessage4Client {
	return o.imsmqMessage4
}

func (o *xxx_DefaultClient) ImsmqQueue() imsmqqueue.ImsmqQueueClient {
	return o.imsmqQueue
}

func (o *xxx_DefaultClient) ImsmqQueue2() imsmqqueue2.ImsmqQueue2Client {
	return o.imsmqQueue2
}

func (o *xxx_DefaultClient) ImsmqQueue3() imsmqqueue3.ImsmqQueue3Client {
	return o.imsmqQueue3
}

func (o *xxx_DefaultClient) ImsmqQueue4() imsmqqueue4.ImsmqQueue4Client {
	return o.imsmqQueue4
}

func (o *xxx_DefaultClient) ImsmqPrivateEvent() imsmqprivateevent.ImsmqPrivateEventClient {
	return o.imsmqPrivateEvent
}

func (o *xxx_DefaultClient) ImsmqEvent2() imsmqevent2.ImsmqEvent2Client {
	return o.imsmqEvent2
}

func (o *xxx_DefaultClient) ImsmqQueueInfo() imsmqqueueinfo.ImsmqQueueInfoClient {
	return o.imsmqQueueInfo
}

func (o *xxx_DefaultClient) ImsmqQueueInfo2() imsmqqueueinfo2.ImsmqQueueInfo2Client {
	return o.imsmqQueueInfo2
}

func (o *xxx_DefaultClient) ImsmqQueueInfo3() imsmqqueueinfo3.ImsmqQueueInfo3Client {
	return o.imsmqQueueInfo3
}

func (o *xxx_DefaultClient) ImsmqQueueInfo4() imsmqqueueinfo4.ImsmqQueueInfo4Client {
	return o.imsmqQueueInfo4
}

func (o *xxx_DefaultClient) ImsmqQueueInfos() imsmqqueueinfos.ImsmqQueueInfosClient {
	return o.imsmqQueueInfos
}

func (o *xxx_DefaultClient) ImsmqQueueInfos2() imsmqqueueinfos2.ImsmqQueueInfos2Client {
	return o.imsmqQueueInfos2
}

func (o *xxx_DefaultClient) ImsmqQueueInfos3() imsmqqueueinfos3.ImsmqQueueInfos3Client {
	return o.imsmqQueueInfos3
}

func (o *xxx_DefaultClient) ImsmqQueueInfos4() imsmqqueueinfos4.ImsmqQueueInfos4Client {
	return o.imsmqQueueInfos4
}

func (o *xxx_DefaultClient) ImsmqTransaction() imsmqtransaction.ImsmqTransactionClient {
	return o.imsmqTransaction
}

func (o *xxx_DefaultClient) ImsmqTransaction2() imsmqtransaction2.ImsmqTransaction2Client {
	return o.imsmqTransaction2
}

func (o *xxx_DefaultClient) ImsmqTransaction3() imsmqtransaction3.ImsmqTransaction3Client {
	return o.imsmqTransaction3
}

func (o *xxx_DefaultClient) ImsmqCoordinatedTransactionDispenser() imsmqcoordinatedtransactiondispenser.ImsmqCoordinatedTransactionDispenserClient {
	return o.imsmqCoordinatedTransactionDispenser
}

func (o *xxx_DefaultClient) ImsmqCoordinatedTransactionDispenser2() imsmqcoordinatedtransactiondispenser2.ImsmqCoordinatedTransactionDispenser2Client {
	return o.imsmqCoordinatedTransactionDispenser2
}

func (o *xxx_DefaultClient) ImsmqCoordinatedTransactionDispenser3() imsmqcoordinatedtransactiondispenser3.ImsmqCoordinatedTransactionDispenser3Client {
	return o.imsmqCoordinatedTransactionDispenser3
}

func (o *xxx_DefaultClient) ImsmqTransactionDispenser() imsmqtransactiondispenser.ImsmqTransactionDispenserClient {
	return o.imsmqTransactionDispenser
}

func (o *xxx_DefaultClient) ImsmqTransactionDispenser2() imsmqtransactiondispenser2.ImsmqTransactionDispenser2Client {
	return o.imsmqTransactionDispenser2
}

func (o *xxx_DefaultClient) ImsmqTransactionDispenser3() imsmqtransactiondispenser3.ImsmqTransactionDispenser3Client {
	return o.imsmqTransactionDispenser3
}

func (o *xxx_DefaultClient) ImsmqApplication() imsmqapplication.ImsmqApplicationClient {
	return o.imsmqApplication
}

func (o *xxx_DefaultClient) ImsmqApplication2() imsmqapplication2.ImsmqApplication2Client {
	return o.imsmqApplication2
}

func (o *xxx_DefaultClient) ImsmqApplication3() imsmqapplication3.ImsmqApplication3Client {
	return o.imsmqApplication3
}

func (o *xxx_DefaultClient) ImsmqDestination() imsmqdestination.ImsmqDestinationClient {
	return o.imsmqDestination
}

func (o *xxx_DefaultClient) ImsmqPrivateDestination() imsmqprivatedestination.ImsmqPrivateDestinationClient {
	return o.imsmqPrivateDestination
}

func (o *xxx_DefaultClient) ImsmqCollection() imsmqcollection.ImsmqCollectionClient {
	return o.imsmqCollection
}

func (o *xxx_DefaultClient) ImsmqManagement() imsmqmanagement.ImsmqManagementClient {
	return o.imsmqManagement
}

func (o *xxx_DefaultClient) ImsmqOutgoingQueueManagement() imsmqoutgoingqueuemanagement.ImsmqOutgoingQueueManagementClient {
	return o.imsmqOutgoingQueueManagement
}

func (o *xxx_DefaultClient) ImsmqQueueManagement() imsmqqueuemanagement.ImsmqQueueManagementClient {
	return o.imsmqQueueManagement
}

func NewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Client, error) {

	opts = append(opts,
		dcerpc.WithAbstractSyntax(itransaction.ITransactionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ienumconnections.EnumConnectionsSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iconnectionpoint.ConnectionPointSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ienumconnectionpoints.EnumConnectionPointsSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iconnectionpointcontainer.ConnectionPointContainerSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqquery.ImsmqQuerySyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqquery2.ImsmqQuery2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqquery3.ImsmqQuery3SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqquery4.ImsmqQuery4SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqmessage.ImsmqMessageSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqmessage2.ImsmqMessage2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqmessage3.ImsmqMessage3SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqmessage4.ImsmqMessage4SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqqueue.ImsmqQueueSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqqueue2.ImsmqQueue2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqqueue3.ImsmqQueue3SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqqueue4.ImsmqQueue4SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqprivateevent.ImsmqPrivateEventSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqevent2.ImsmqEvent2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqqueueinfo.ImsmqQueueInfoSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqqueueinfo2.ImsmqQueueInfo2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqqueueinfo3.ImsmqQueueInfo3SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqqueueinfo4.ImsmqQueueInfo4SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqqueueinfos.ImsmqQueueInfosSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqqueueinfos2.ImsmqQueueInfos2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqqueueinfos3.ImsmqQueueInfos3SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqqueueinfos4.ImsmqQueueInfos4SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqtransaction.ImsmqTransactionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqtransaction2.ImsmqTransaction2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqtransaction3.ImsmqTransaction3SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqcoordinatedtransactiondispenser.ImsmqCoordinatedTransactionDispenserSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqcoordinatedtransactiondispenser2.ImsmqCoordinatedTransactionDispenser2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqcoordinatedtransactiondispenser3.ImsmqCoordinatedTransactionDispenser3SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqtransactiondispenser.ImsmqTransactionDispenserSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqtransactiondispenser2.ImsmqTransactionDispenser2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqtransactiondispenser3.ImsmqTransactionDispenser3SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqapplication.ImsmqApplicationSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqapplication2.ImsmqApplication2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqapplication3.ImsmqApplication3SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqdestination.ImsmqDestinationSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqprivatedestination.ImsmqPrivateDestinationSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqcollection.ImsmqCollectionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqmanagement.ImsmqManagementSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqoutgoingqueuemanagement.ImsmqOutgoingQueueManagementSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsmqqueuemanagement.ImsmqQueueManagementSyntaxV0_0),
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

	imsmqQuerySubConn, err := sub.SubConn(ctx, imsmqquery.ImsmqQuerySyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqQuerySubConn = sub
	}

	o.imsmqQuery, err = imsmqquery.NewImsmqQueryClient(ctx, imsmqQuerySubConn, append(opts, dcerpc.WithNoBind(imsmqQuerySubConn))...)

	imsmqQuery2SubConn, err := sub.SubConn(ctx, imsmqquery2.ImsmqQuery2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqQuery2SubConn = sub
	}

	o.imsmqQuery2, err = imsmqquery2.NewImsmqQuery2Client(ctx, imsmqQuery2SubConn, append(opts, dcerpc.WithNoBind(imsmqQuery2SubConn))...)

	imsmqQuery3SubConn, err := sub.SubConn(ctx, imsmqquery3.ImsmqQuery3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqQuery3SubConn = sub
	}

	o.imsmqQuery3, err = imsmqquery3.NewImsmqQuery3Client(ctx, imsmqQuery3SubConn, append(opts, dcerpc.WithNoBind(imsmqQuery3SubConn))...)

	imsmqQuery4SubConn, err := sub.SubConn(ctx, imsmqquery4.ImsmqQuery4SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqQuery4SubConn = sub
	}

	o.imsmqQuery4, err = imsmqquery4.NewImsmqQuery4Client(ctx, imsmqQuery4SubConn, append(opts, dcerpc.WithNoBind(imsmqQuery4SubConn))...)

	imsmqMessageSubConn, err := sub.SubConn(ctx, imsmqmessage.ImsmqMessageSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqMessageSubConn = sub
	}

	o.imsmqMessage, err = imsmqmessage.NewImsmqMessageClient(ctx, imsmqMessageSubConn, append(opts, dcerpc.WithNoBind(imsmqMessageSubConn))...)

	imsmqMessage2SubConn, err := sub.SubConn(ctx, imsmqmessage2.ImsmqMessage2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqMessage2SubConn = sub
	}

	o.imsmqMessage2, err = imsmqmessage2.NewImsmqMessage2Client(ctx, imsmqMessage2SubConn, append(opts, dcerpc.WithNoBind(imsmqMessage2SubConn))...)

	imsmqMessage3SubConn, err := sub.SubConn(ctx, imsmqmessage3.ImsmqMessage3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqMessage3SubConn = sub
	}

	o.imsmqMessage3, err = imsmqmessage3.NewImsmqMessage3Client(ctx, imsmqMessage3SubConn, append(opts, dcerpc.WithNoBind(imsmqMessage3SubConn))...)

	imsmqMessage4SubConn, err := sub.SubConn(ctx, imsmqmessage4.ImsmqMessage4SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqMessage4SubConn = sub
	}

	o.imsmqMessage4, err = imsmqmessage4.NewImsmqMessage4Client(ctx, imsmqMessage4SubConn, append(opts, dcerpc.WithNoBind(imsmqMessage4SubConn))...)

	imsmqQueueSubConn, err := sub.SubConn(ctx, imsmqqueue.ImsmqQueueSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqQueueSubConn = sub
	}

	o.imsmqQueue, err = imsmqqueue.NewImsmqQueueClient(ctx, imsmqQueueSubConn, append(opts, dcerpc.WithNoBind(imsmqQueueSubConn))...)

	imsmqQueue2SubConn, err := sub.SubConn(ctx, imsmqqueue2.ImsmqQueue2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqQueue2SubConn = sub
	}

	o.imsmqQueue2, err = imsmqqueue2.NewImsmqQueue2Client(ctx, imsmqQueue2SubConn, append(opts, dcerpc.WithNoBind(imsmqQueue2SubConn))...)

	imsmqQueue3SubConn, err := sub.SubConn(ctx, imsmqqueue3.ImsmqQueue3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqQueue3SubConn = sub
	}

	o.imsmqQueue3, err = imsmqqueue3.NewImsmqQueue3Client(ctx, imsmqQueue3SubConn, append(opts, dcerpc.WithNoBind(imsmqQueue3SubConn))...)

	imsmqQueue4SubConn, err := sub.SubConn(ctx, imsmqqueue4.ImsmqQueue4SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqQueue4SubConn = sub
	}

	o.imsmqQueue4, err = imsmqqueue4.NewImsmqQueue4Client(ctx, imsmqQueue4SubConn, append(opts, dcerpc.WithNoBind(imsmqQueue4SubConn))...)

	imsmqPrivateEventSubConn, err := sub.SubConn(ctx, imsmqprivateevent.ImsmqPrivateEventSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqPrivateEventSubConn = sub
	}

	o.imsmqPrivateEvent, err = imsmqprivateevent.NewImsmqPrivateEventClient(ctx, imsmqPrivateEventSubConn, append(opts, dcerpc.WithNoBind(imsmqPrivateEventSubConn))...)

	imsmqEvent2SubConn, err := sub.SubConn(ctx, imsmqevent2.ImsmqEvent2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqEvent2SubConn = sub
	}

	o.imsmqEvent2, err = imsmqevent2.NewImsmqEvent2Client(ctx, imsmqEvent2SubConn, append(opts, dcerpc.WithNoBind(imsmqEvent2SubConn))...)

	imsmqQueueInfoSubConn, err := sub.SubConn(ctx, imsmqqueueinfo.ImsmqQueueInfoSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqQueueInfoSubConn = sub
	}

	o.imsmqQueueInfo, err = imsmqqueueinfo.NewImsmqQueueInfoClient(ctx, imsmqQueueInfoSubConn, append(opts, dcerpc.WithNoBind(imsmqQueueInfoSubConn))...)

	imsmqQueueInfo2SubConn, err := sub.SubConn(ctx, imsmqqueueinfo2.ImsmqQueueInfo2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqQueueInfo2SubConn = sub
	}

	o.imsmqQueueInfo2, err = imsmqqueueinfo2.NewImsmqQueueInfo2Client(ctx, imsmqQueueInfo2SubConn, append(opts, dcerpc.WithNoBind(imsmqQueueInfo2SubConn))...)

	imsmqQueueInfo3SubConn, err := sub.SubConn(ctx, imsmqqueueinfo3.ImsmqQueueInfo3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqQueueInfo3SubConn = sub
	}

	o.imsmqQueueInfo3, err = imsmqqueueinfo3.NewImsmqQueueInfo3Client(ctx, imsmqQueueInfo3SubConn, append(opts, dcerpc.WithNoBind(imsmqQueueInfo3SubConn))...)

	imsmqQueueInfo4SubConn, err := sub.SubConn(ctx, imsmqqueueinfo4.ImsmqQueueInfo4SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqQueueInfo4SubConn = sub
	}

	o.imsmqQueueInfo4, err = imsmqqueueinfo4.NewImsmqQueueInfo4Client(ctx, imsmqQueueInfo4SubConn, append(opts, dcerpc.WithNoBind(imsmqQueueInfo4SubConn))...)

	imsmqQueueInfosSubConn, err := sub.SubConn(ctx, imsmqqueueinfos.ImsmqQueueInfosSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqQueueInfosSubConn = sub
	}

	o.imsmqQueueInfos, err = imsmqqueueinfos.NewImsmqQueueInfosClient(ctx, imsmqQueueInfosSubConn, append(opts, dcerpc.WithNoBind(imsmqQueueInfosSubConn))...)

	imsmqQueueInfos2SubConn, err := sub.SubConn(ctx, imsmqqueueinfos2.ImsmqQueueInfos2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqQueueInfos2SubConn = sub
	}

	o.imsmqQueueInfos2, err = imsmqqueueinfos2.NewImsmqQueueInfos2Client(ctx, imsmqQueueInfos2SubConn, append(opts, dcerpc.WithNoBind(imsmqQueueInfos2SubConn))...)

	imsmqQueueInfos3SubConn, err := sub.SubConn(ctx, imsmqqueueinfos3.ImsmqQueueInfos3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqQueueInfos3SubConn = sub
	}

	o.imsmqQueueInfos3, err = imsmqqueueinfos3.NewImsmqQueueInfos3Client(ctx, imsmqQueueInfos3SubConn, append(opts, dcerpc.WithNoBind(imsmqQueueInfos3SubConn))...)

	imsmqQueueInfos4SubConn, err := sub.SubConn(ctx, imsmqqueueinfos4.ImsmqQueueInfos4SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqQueueInfos4SubConn = sub
	}

	o.imsmqQueueInfos4, err = imsmqqueueinfos4.NewImsmqQueueInfos4Client(ctx, imsmqQueueInfos4SubConn, append(opts, dcerpc.WithNoBind(imsmqQueueInfos4SubConn))...)

	imsmqTransactionSubConn, err := sub.SubConn(ctx, imsmqtransaction.ImsmqTransactionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqTransactionSubConn = sub
	}

	o.imsmqTransaction, err = imsmqtransaction.NewImsmqTransactionClient(ctx, imsmqTransactionSubConn, append(opts, dcerpc.WithNoBind(imsmqTransactionSubConn))...)

	imsmqTransaction2SubConn, err := sub.SubConn(ctx, imsmqtransaction2.ImsmqTransaction2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqTransaction2SubConn = sub
	}

	o.imsmqTransaction2, err = imsmqtransaction2.NewImsmqTransaction2Client(ctx, imsmqTransaction2SubConn, append(opts, dcerpc.WithNoBind(imsmqTransaction2SubConn))...)

	imsmqTransaction3SubConn, err := sub.SubConn(ctx, imsmqtransaction3.ImsmqTransaction3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqTransaction3SubConn = sub
	}

	o.imsmqTransaction3, err = imsmqtransaction3.NewImsmqTransaction3Client(ctx, imsmqTransaction3SubConn, append(opts, dcerpc.WithNoBind(imsmqTransaction3SubConn))...)

	imsmqCoordinatedTransactionDispenserSubConn, err := sub.SubConn(ctx, imsmqcoordinatedtransactiondispenser.ImsmqCoordinatedTransactionDispenserSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqCoordinatedTransactionDispenserSubConn = sub
	}

	o.imsmqCoordinatedTransactionDispenser, err = imsmqcoordinatedtransactiondispenser.NewImsmqCoordinatedTransactionDispenserClient(ctx, imsmqCoordinatedTransactionDispenserSubConn, append(opts, dcerpc.WithNoBind(imsmqCoordinatedTransactionDispenserSubConn))...)

	imsmqCoordinatedTransactionDispenser2SubConn, err := sub.SubConn(ctx, imsmqcoordinatedtransactiondispenser2.ImsmqCoordinatedTransactionDispenser2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqCoordinatedTransactionDispenser2SubConn = sub
	}

	o.imsmqCoordinatedTransactionDispenser2, err = imsmqcoordinatedtransactiondispenser2.NewImsmqCoordinatedTransactionDispenser2Client(ctx, imsmqCoordinatedTransactionDispenser2SubConn, append(opts, dcerpc.WithNoBind(imsmqCoordinatedTransactionDispenser2SubConn))...)

	imsmqCoordinatedTransactionDispenser3SubConn, err := sub.SubConn(ctx, imsmqcoordinatedtransactiondispenser3.ImsmqCoordinatedTransactionDispenser3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqCoordinatedTransactionDispenser3SubConn = sub
	}

	o.imsmqCoordinatedTransactionDispenser3, err = imsmqcoordinatedtransactiondispenser3.NewImsmqCoordinatedTransactionDispenser3Client(ctx, imsmqCoordinatedTransactionDispenser3SubConn, append(opts, dcerpc.WithNoBind(imsmqCoordinatedTransactionDispenser3SubConn))...)

	imsmqTransactionDispenserSubConn, err := sub.SubConn(ctx, imsmqtransactiondispenser.ImsmqTransactionDispenserSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqTransactionDispenserSubConn = sub
	}

	o.imsmqTransactionDispenser, err = imsmqtransactiondispenser.NewImsmqTransactionDispenserClient(ctx, imsmqTransactionDispenserSubConn, append(opts, dcerpc.WithNoBind(imsmqTransactionDispenserSubConn))...)

	imsmqTransactionDispenser2SubConn, err := sub.SubConn(ctx, imsmqtransactiondispenser2.ImsmqTransactionDispenser2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqTransactionDispenser2SubConn = sub
	}

	o.imsmqTransactionDispenser2, err = imsmqtransactiondispenser2.NewImsmqTransactionDispenser2Client(ctx, imsmqTransactionDispenser2SubConn, append(opts, dcerpc.WithNoBind(imsmqTransactionDispenser2SubConn))...)

	imsmqTransactionDispenser3SubConn, err := sub.SubConn(ctx, imsmqtransactiondispenser3.ImsmqTransactionDispenser3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqTransactionDispenser3SubConn = sub
	}

	o.imsmqTransactionDispenser3, err = imsmqtransactiondispenser3.NewImsmqTransactionDispenser3Client(ctx, imsmqTransactionDispenser3SubConn, append(opts, dcerpc.WithNoBind(imsmqTransactionDispenser3SubConn))...)

	imsmqApplicationSubConn, err := sub.SubConn(ctx, imsmqapplication.ImsmqApplicationSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqApplicationSubConn = sub
	}

	o.imsmqApplication, err = imsmqapplication.NewImsmqApplicationClient(ctx, imsmqApplicationSubConn, append(opts, dcerpc.WithNoBind(imsmqApplicationSubConn))...)

	imsmqApplication2SubConn, err := sub.SubConn(ctx, imsmqapplication2.ImsmqApplication2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqApplication2SubConn = sub
	}

	o.imsmqApplication2, err = imsmqapplication2.NewImsmqApplication2Client(ctx, imsmqApplication2SubConn, append(opts, dcerpc.WithNoBind(imsmqApplication2SubConn))...)

	imsmqApplication3SubConn, err := sub.SubConn(ctx, imsmqapplication3.ImsmqApplication3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqApplication3SubConn = sub
	}

	o.imsmqApplication3, err = imsmqapplication3.NewImsmqApplication3Client(ctx, imsmqApplication3SubConn, append(opts, dcerpc.WithNoBind(imsmqApplication3SubConn))...)

	imsmqDestinationSubConn, err := sub.SubConn(ctx, imsmqdestination.ImsmqDestinationSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqDestinationSubConn = sub
	}

	o.imsmqDestination, err = imsmqdestination.NewImsmqDestinationClient(ctx, imsmqDestinationSubConn, append(opts, dcerpc.WithNoBind(imsmqDestinationSubConn))...)

	imsmqPrivateDestinationSubConn, err := sub.SubConn(ctx, imsmqprivatedestination.ImsmqPrivateDestinationSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqPrivateDestinationSubConn = sub
	}

	o.imsmqPrivateDestination, err = imsmqprivatedestination.NewImsmqPrivateDestinationClient(ctx, imsmqPrivateDestinationSubConn, append(opts, dcerpc.WithNoBind(imsmqPrivateDestinationSubConn))...)

	imsmqCollectionSubConn, err := sub.SubConn(ctx, imsmqcollection.ImsmqCollectionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqCollectionSubConn = sub
	}

	o.imsmqCollection, err = imsmqcollection.NewImsmqCollectionClient(ctx, imsmqCollectionSubConn, append(opts, dcerpc.WithNoBind(imsmqCollectionSubConn))...)

	imsmqManagementSubConn, err := sub.SubConn(ctx, imsmqmanagement.ImsmqManagementSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqManagementSubConn = sub
	}

	o.imsmqManagement, err = imsmqmanagement.NewImsmqManagementClient(ctx, imsmqManagementSubConn, append(opts, dcerpc.WithNoBind(imsmqManagementSubConn))...)

	imsmqOutgoingQueueManagementSubConn, err := sub.SubConn(ctx, imsmqoutgoingqueuemanagement.ImsmqOutgoingQueueManagementSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqOutgoingQueueManagementSubConn = sub
	}

	o.imsmqOutgoingQueueManagement, err = imsmqoutgoingqueuemanagement.NewImsmqOutgoingQueueManagementClient(ctx, imsmqOutgoingQueueManagementSubConn, append(opts, dcerpc.WithNoBind(imsmqOutgoingQueueManagementSubConn))...)

	imsmqQueueManagementSubConn, err := sub.SubConn(ctx, imsmqqueuemanagement.ImsmqQueueManagementSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsmqQueueManagementSubConn = sub
	}

	o.imsmqQueueManagement, err = imsmqqueuemanagement.NewImsmqQueueManagementClient(ctx, imsmqQueueManagementSubConn, append(opts, dcerpc.WithNoBind(imsmqQueueManagementSubConn))...)
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
		dcomClient:                            o.dcomClient.IPID(ctx, ipid),
		itransaction:                          o.itransaction.IPID(ctx, ipid),
		enumConnections:                       o.enumConnections.IPID(ctx, ipid),
		connectionPoint:                       o.connectionPoint.IPID(ctx, ipid),
		enumConnectionPoints:                  o.enumConnectionPoints.IPID(ctx, ipid),
		connectionPointContainer:              o.connectionPointContainer.IPID(ctx, ipid),
		imsmqQuery:                            o.imsmqQuery.IPID(ctx, ipid),
		imsmqQuery2:                           o.imsmqQuery2.IPID(ctx, ipid),
		imsmqQuery3:                           o.imsmqQuery3.IPID(ctx, ipid),
		imsmqQuery4:                           o.imsmqQuery4.IPID(ctx, ipid),
		imsmqMessage:                          o.imsmqMessage.IPID(ctx, ipid),
		imsmqMessage2:                         o.imsmqMessage2.IPID(ctx, ipid),
		imsmqMessage3:                         o.imsmqMessage3.IPID(ctx, ipid),
		imsmqMessage4:                         o.imsmqMessage4.IPID(ctx, ipid),
		imsmqQueue:                            o.imsmqQueue.IPID(ctx, ipid),
		imsmqQueue2:                           o.imsmqQueue2.IPID(ctx, ipid),
		imsmqQueue3:                           o.imsmqQueue3.IPID(ctx, ipid),
		imsmqQueue4:                           o.imsmqQueue4.IPID(ctx, ipid),
		imsmqPrivateEvent:                     o.imsmqPrivateEvent.IPID(ctx, ipid),
		imsmqEvent2:                           o.imsmqEvent2.IPID(ctx, ipid),
		imsmqQueueInfo:                        o.imsmqQueueInfo.IPID(ctx, ipid),
		imsmqQueueInfo2:                       o.imsmqQueueInfo2.IPID(ctx, ipid),
		imsmqQueueInfo3:                       o.imsmqQueueInfo3.IPID(ctx, ipid),
		imsmqQueueInfo4:                       o.imsmqQueueInfo4.IPID(ctx, ipid),
		imsmqQueueInfos:                       o.imsmqQueueInfos.IPID(ctx, ipid),
		imsmqQueueInfos2:                      o.imsmqQueueInfos2.IPID(ctx, ipid),
		imsmqQueueInfos3:                      o.imsmqQueueInfos3.IPID(ctx, ipid),
		imsmqQueueInfos4:                      o.imsmqQueueInfos4.IPID(ctx, ipid),
		imsmqTransaction:                      o.imsmqTransaction.IPID(ctx, ipid),
		imsmqTransaction2:                     o.imsmqTransaction2.IPID(ctx, ipid),
		imsmqTransaction3:                     o.imsmqTransaction3.IPID(ctx, ipid),
		imsmqCoordinatedTransactionDispenser:  o.imsmqCoordinatedTransactionDispenser.IPID(ctx, ipid),
		imsmqCoordinatedTransactionDispenser2: o.imsmqCoordinatedTransactionDispenser2.IPID(ctx, ipid),
		imsmqCoordinatedTransactionDispenser3: o.imsmqCoordinatedTransactionDispenser3.IPID(ctx, ipid),
		imsmqTransactionDispenser:             o.imsmqTransactionDispenser.IPID(ctx, ipid),
		imsmqTransactionDispenser2:            o.imsmqTransactionDispenser2.IPID(ctx, ipid),
		imsmqTransactionDispenser3:            o.imsmqTransactionDispenser3.IPID(ctx, ipid),
		imsmqApplication:                      o.imsmqApplication.IPID(ctx, ipid),
		imsmqApplication2:                     o.imsmqApplication2.IPID(ctx, ipid),
		imsmqApplication3:                     o.imsmqApplication3.IPID(ctx, ipid),
		imsmqDestination:                      o.imsmqDestination.IPID(ctx, ipid),
		imsmqPrivateDestination:               o.imsmqPrivateDestination.IPID(ctx, ipid),
		imsmqCollection:                       o.imsmqCollection.IPID(ctx, ipid),
		imsmqManagement:                       o.imsmqManagement.IPID(ctx, ipid),
		imsmqOutgoingQueueManagement:          o.imsmqOutgoingQueueManagement.IPID(ctx, ipid),
		imsmqQueueManagement:                  o.imsmqQueueManagement.IPID(ctx, ipid),
		cc:                                    o.cc,
	}
}
