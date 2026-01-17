// The pcq package implements the PCQ client protocol.
//
// # Introduction
//
// The Performance Counter Query Protocol is a remote procedure call (RPC)–based protocol
// that is used for browsing performance counters and retrieving performance counter
// values from a server.
//
// # Overview
//
// To effectively manage systems, administrators need the capability to query for performance
// counter data on the health or state of a particular application or system. Software
// components that are designed with performance counters are therefore easier to manage
// and diagnose. The Performance Counter Query Protocol enables system administrators
// to query performance counters on a remote server.
//
// The Performance Counter Query Protocol is used to retrieve performance counter information
// from a server. The protocol allows a client to enumerate the performance counters
// that are available on the server. The server can use the protocol to return performance
// counter information, such as localized counter names and description strings, performance
// counter types (for more information, see [MSDN-COUNT]), and instance information
// if there are multiple instances of a performance counter. The client can also use
// the protocol to establish a query on the server and add or remove performance counters
// to it. The client can then repeatedly retrieve performance counter data that is associated
// with the query by using the protocol.
package pcq

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
	GoPackage = "pcq"
)
