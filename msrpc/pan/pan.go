// The pan package implements the PAN client protocol.
//
// # Introduction
//
// The Print System Asynchronous Notification Protocol is used asynchronously by clients
// to receive print status notifications from a server and to send back responses to
// those notifications. A set of notifications and responses are defined together as
// a notification type.
//
// This protocol is based on remote procedure call (RPC) [C706] [MS-RPCE]. The interfaces
// and methods defined by this protocol provide a transport mechanism for arbitrary
// notification types.
//
// The Print System Asynchronous Notification Protocol defines a notification type called
// AsyncUI. The AsyncUI notification type enables a notification source on a server
// to request the display of an informative alert on a client, the client to send back
// user input requested by the alert, and the notification source to request the execution
// of code that is resident on the client.
//
// # Overview
//
// The Print System Asynchronous Notification Protocol serves two primary functions:
//
// * A print server ( 4751a17a-76f3-4a7a-aef3-fdfc5d1bc7a5#gt_59fb3ddc-63cf-45df-8a90-46a6af9e00cb
// ) sending status notifications ( 4751a17a-76f3-4a7a-aef3-fdfc5d1bc7a5#gt_4571dc27-4115-4cdf-8dc3-f8fe410a9966
// ) to a print client ( 4751a17a-76f3-4a7a-aef3-fdfc5d1bc7a5#gt_3b2da3d1-c159-4399-a6dd-dfd5f76fa2f5
// ).
//
// * A print server receiving the client's response ( 4751a17a-76f3-4a7a-aef3-fdfc5d1bc7a5#gt_115e9f9c-314e-49fc-977d-238daf619828
// ) to the notifications.
//
// This protocol has two modes of operation:
//
// * bidirectional communication mode ( 4751a17a-76f3-4a7a-aef3-fdfc5d1bc7a5#gt_b1d50fe5-0a8c-44fe-8802-3382d1af77e5
// )
//
// * unidirectional communication mode ( 4751a17a-76f3-4a7a-aef3-fdfc5d1bc7a5#gt_d203482a-15ef-4528-9064-68b0f2e0c5fa
// )
//
// In bidirectional communication mode, data can flow in two directions between a server
// and client. After a client registers with a server, the client requests a bidirectional
// notification channel from the server. The client uses the channel to request predefined
// print status notifications from the server. When the client subsequently receives
// a notification, the client also uses the channel to send a response back to the server.
//
// In bidirectional communication mode, if multiple clients open the same bidirectional
// notification channel and attempt to respond to the channel's initial notification,
// the server accepts only the first response received and continues to send further
// notifications only to the client whose response was accepted. Subsequent exchanges
// of notifications and responses on the channel take place only between the server
// and that client. Because bidirectional notifications each require a response, the
// initial notification intended to be transmitted on a bidirectional notification channel
// cannot be discarded before a registered client sends a response on that channel (or
// the channel is closed).
//
// The following diagram shows bidirectional communication.
package pan

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
	GoPackage = "pan"
)
