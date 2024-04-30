// The mqmr package implements the MQMR client protocol.
//
// # Introduction
//
// The Message Queuing (MSMQ): Queue Manager Management Protocol is a remote procedure
// call (RPC)-based protocol used for management operations on the MSMQ server, including
// monitoring the MSMQ installation and the queues.
//
// Operations that a client can perform using this protocol include:
//
// * Getting information on MSMQ installation and queues.
//
// * Performing actions on an MSMQ installation.
//
// * Performing actions on a queue.
//
// # Overview
//
// The Message Queuing (MSMQ): Queue Manager Management Protocol allows an MSMQ client
// application to perform management operations on an MSMQ server.
//
// This protocol can be used to get the following information:
//
// * Queue ( a59a1c97-3eea-4af5-b31d-937739889af4#gt_c1a6400d-703b-4f9a-a74c-40f1487978d9
// ) properties, such as:
//
// * The path name ( a59a1c97-3eea-4af5-b31d-937739889af4#gt_78c4af57-f564-4aa9-a40d-f54ea6bd8766
// ) of a queue.
//
// * The format name ( a59a1c97-3eea-4af5-b31d-937739889af4#gt_390ae273-7109-44eb-981f-aa157e6b37c0
// ) of a queue.
//
// * Whether a queue is (or is not) located on a computer, or whether it is a transactional
// queue ( a59a1c97-3eea-4af5-b31d-937739889af4#gt_49b6c0e7-59fc-40af-b627-df18c392327b
// ) or a foreign queue ( a59a1c97-3eea-4af5-b31d-937739889af4#gt_ee553c1a-b390-42d5-a785-2412a31f98fb
// ).
//
// * The retransmit interval for messages in an outgoing queue ( a59a1c97-3eea-4af5-b31d-937739889af4#gt_311107bb-e41a-4794-82b6-9ebf83172cb5
// ) for which no order acknowledgment has been received.
//
// * The number of subqueues ( a59a1c97-3eea-4af5-b31d-937739889af4#gt_328fe8aa-d006-44dd-8cc8-dba7c862aaf8
// ) in a specified queue. <1> ( 47ca62cf-676e-4d68-841f-9b936124e75e#Appendix_A_1 )
//
// * The names of the subqueues in a specified queue. <2> ( 47ca62cf-676e-4d68-841f-9b936124e75e#Appendix_A_2
// )
//
// * The version and build information for the computer operating system and the MSMQ
// installation.
//
// * Current queue state, such as:
//
// * The number of messages in a queue or in a queue journal ( a59a1c97-3eea-4af5-b31d-937739889af4#gt_89cac287-8784-4fdf-893f-739cd0ef3785
// ).
//
// * The number of message bytes in a queue or in a queue journal.
//
// * The connection state of an outgoing queue.
//
// * The list of the active queues ( a59a1c97-3eea-4af5-b31d-937739889af4#gt_9f320bf9-052c-4c1c-b2e1-00ffc0067f46
// ) on a computer.
//
// * The name of the current MSMQ Directory Service server ( a59a1c97-3eea-4af5-b31d-937739889af4#gt_0e26e115-867e-4959-8a8b-5624b9d8a119
// ) for a computer.
//
// * Whether a queue manager ( a59a1c97-3eea-4af5-b31d-937739889af4#gt_476f10ed-08f0-4887-b583-59d5cf909979
// ) on a computer is disconnected from the network.
//
// * The list of the path names of all the private queues ( a59a1c97-3eea-4af5-b31d-937739889af4#gt_1a7f2b81-359e-4d2f-b4b4-f30bc7dd0d8f
// ) registered on a computer.
//
// * Auditing information, such as:
//
// * The connection state history of a queue. <3> ( 47ca62cf-676e-4d68-841f-9b936124e75e#Appendix_A_3
// )
//
// * The number of messages sent from a computer to a queue for which no order acknowledgment
// has been received.
//
// * The number of messages sent from a computer to a queue for which an order acknowledgment
// has been received, but a receive acknowledgment message has not been received.
//
// * The date and time when the last order acknowledgment for a message sent from a
// computer to a queue was received.
//
// * The time when MSMQ will attempt to retransmit a message from a computer to a queue.
//
// * The number of times that the last message in the corresponding outgoing queue on
// a computer was sent.
//
// * The number of times that the last order acknowledgment for a message sent from
// a computer to a queue has been received.
//
// * The number of message bytes stored in all the queues on a computer. <4> ( 47ca62cf-676e-4d68-841f-9b936124e75e#Appendix_A_4
// )
//
// * Sequence information, such as:
//
// * The address or a list of possible addresses for routing messages to the destination
// queue in the next hop.
//
// * The next message to be sent from a computer to a queue.
//
// * The last message that was sent from a computer to a queue for which no order acknowledgment
// has been received.
//
// * The first message sent from a computer to a queue for which no order acknowledgment
// has been received.
//
// * An array of arrays of information on the transactional messages ( a59a1c97-3eea-4af5-b31d-937739889af4#gt_3b80e01d-5155-4378-b692-82122af044aa
// ) sent from all source computers to a queue on a target computer. Each element of
// the overall array is an array (vector) containing one of the following pieces of
// information for all of the source computers.
//
// * The format names used to open a queue when the last messages were sent.
//
// * The globally unique identifiers (GUIDs) ( a59a1c97-3eea-4af5-b31d-937739889af4#gt_f49694cc-c350-462d-ab8e-816f0103c6c1
// ) of the sending queue managers.
//
// * The last sequence identifiers.
//
// * The sequence numbers of the last messages sent to a queue by one or more sending
// queue managers.
//
// * The times when each sending queue manager last accessed a queue.
//
// * The number of times that the last messages were rejected.
//
// The protocol can also be used to perform actions on a computer, such as:
//
// * Connecting the queue manager on a computer to a network and an MSMQ Directory Service
// server.
//
// * Disconnecting the queue manager on a computer from a network and an MSMQ Directory
// Service server.
//
// * Deleting empty message files.
//
// The protocol can also be used to perform actions on a queue, such as:
//
// * Pausing the sending of messages from a computer. The queue manager will not send
// messages to the applicable destination queue until a resume action is initiated.
//
// * Restarting the sending of messages after a pause action is initiated.
//
// * Resending the pending transaction sequence (as specified in [MS-MQQB] ( ../ms-mqqb/85498b96-f2c8-43b3-a108-c9d6269dc4af
// ) ).
//
// This is an RPC-based protocol. The server does not maintain client state information.
// The protocol operation is stateless.
//
// This is a simple request-response protocol. For each received method request, the
// server executes the requested method and returns a completion status to the client.
// This is a stateless protocol; each method call is independent of any previous method
// calls.
package mqmr

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
)

var (
	// import guard
	GoPackage = "mqmr"
)
