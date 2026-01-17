// The lrec package implements the LREC client protocol.
//
// # Introduction
//
// The Live Remote Event Capture (LREC) Protocol allows a management station to monitor
// events on a target system across a network. The protocol consists of two components:
//
// * A WS-Management-based control channel for starting and stopping an event capture.
//
// * A remote procedure call (RPC) ( 59d7a0e2-342c-4dc3-bc27-88e9c4aa0415#gt_8a7f6700-8311-45bc-af10-82e10accd331
// ) -based data channel for retrieving events as they are logged on the remote system.
//
// Together, these components can be used to support monitoring scenarios and provide
// a "first line of defense" for troubleshooting scenarios, especially when the remote
// system does not support the ability to locally log events.
//
// # Overview
//
// The Live Remote Event Capture (LREC) protocol allows a client to connect to a server
// to monitor critical information and detect issues as they occur on the server. For
// example, to detect under-provisioned servers, an administrator can use this protocol
// to remotely see the events that indicate under-provisioning which are logged as high
// memory utilization. The remote visibility into the event logging enables the administrator
// to re-balance the load on the server, immediately observe the fix, and continue to
// make improvements as necessary.
//
// In the LREC protocol, information is sent over the network to a client as a sequential
// stream of records each of which is referred to as an event. Multiple software components
// and applications on the server can report events using the protocol. These are referred
// to event providers. Each event provider is identified by a unique "provider ID" and
// its event types are described in a provider manifest organized in any implementation-specific
// way, such as the XML schema specified in [MSDN-EvntManifest].
//
// Event providers can define multiple kinds of events. For example, a user logging
// on could be one kind of event and a user logging off could be another. When a provider
// reports an event, it specifies an event provider-specific Event Type ID (referred
// to as an event ID) that indicates the specific kind of event being reported. The
// event ID is reused whenever another event of the same type is reported. Therefore,
// each event type is uniquely identified by a provider ID and an event ID.
//
// Different server configurations and application workloads have varying requirements
// for monitoring and troubleshooting. To ensure flexible support for these scenarios,
// multiple event providers can be added into an event session to enable simultaneous
// event recording. When using multiple event providers, two techniques in particular
// enable the broad coverage of a session containing many event providers, yet narrow
// the number of observed events:
//
// * The server filters events based on the "error level" or criticality of the events.
//
// * The server filters events based on keywords, such as authentication, input/output,
// or user interface.
//
// In the LREC protocol, an event session is configured and started using a WS-Management-based
// control channel. When the session is started, the server initializes an RPC endpoint
// and the client connects to the server using the RPC endpoint to receive reported
// events. When the client is finished observing reported events, the client stops the
// session using the WS-Management-based control channel. When all event sessions are
// stopped, the RPC endpoint is removed.
package lrec

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
)

var (
	// import guard
	GoPackage = "lrec"
)
