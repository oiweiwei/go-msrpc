// The scmr package implements the SCMR client protocol.
//
// # Introduction
//
// The Service Control Manager Remote Protocol is a remote procedure call (RPC)–based
// client/server protocol that is used for remotely managing the Service Control Manager
// (SCM). The SCM is an RPC server that enables service configuration and control of
// service programs. For more information, see [MSDN-WINSVC].
//
// # Overview
//
// The Service Control Manager Remote Protocol is a client/server protocol used for
// configuring and controlling service programs running on a remote computer. A remote
// service management session begins with the client initiating the connection request
// to the server. If the server grants the request, the connection is established. The
// client can then make multiple requests to modify, query the configuration, or start
// and stop services on the server by using the same session until the session is terminated.
//
// A typical Service Control Manager Remote Protocol session involves the client connecting
// to the server and requesting to open the SCM on the server. If the server accepts
// the request, it responds with an RPC context handle to the client. The client uses
// this RPC context handle to operate on the server. This usually involves sending another
// request to the server and specifying the type of operation to perform and any specific
// parameters associated with that operation. If the server accepts this request, it
// attempts to perform the specified operation and responds to the client with the result
// of the operation. After the client is finished operating on the server, it terminates
// the protocol by sending a request to close the RPC context handle.
//
// The Service Control Manager Remote Protocol maintains an internal database to store
// service program configurations and state. The Service Control Manager Protocol has
// exclusive access to this internal database. On one operating system instance there
// is only one SCM and one corresponding SCM database. Any updates to this internal
// database are made only through the Service Control Manager Remote Protocol. SCM takes
// care of serializing all concurrent accesses to the SCM database. The SCM database
// is resident in memory; it is recreated every time the SCM restarts (after each reboot).
// Part of the SCM database is retrieved from persistent storage (all information regarding
// registered services) and partially nonpersistent (current active state of the services).
// The persistent information is modified by the SCM when a service is added, configured,
// or deleted. Any attempt to directly modify the persistent part of the database directly
// in the persistent storage is not a supported scenario and will result in possible
// inconsistencies. Finally, if SCM were to be forcefully terminated, the operating
// system will shut down and restart.
package scmr

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
	GoPackage = "scmr"
)
