// The rrp package implements the RRP client protocol.
//
// # Introduction
//
// The Windows Remote Registry Protocol is a remote procedure call (RPC)–based client/server
// protocol that is used for remotely managing a hierarchical Data Store such as the
// Windows registry. For more information, see [MSWINREG].
//
// # Overview
//
// The Windows Remote Registry Protocol is a client/server protocol that is used for
// remotely managing a hierarchical Data Store with lightly typed elements. The layout
// and specifics of such a store is specified in section 3.1.1.
//
// A remote registry management session begins with the client initiating the connection
// request to the server. If the server grants the request, the connection is established.
// The client can then make multiple requests to read or modify the registry on the
// server by using the same session until the session is terminated.
//
// A typical remote registry session involves the client connecting to the server and
// requesting to open a registry key on the server. If the server accepts the request,
// it responds with an RPC context handle that refers to the key. The client uses this
// RPC context handle to operate on that key. This usually involves sending another
// request to the server specifying the type of operation to perform and any specific
// parameters that are associated with that operation. If the server accepts this request,
// it attempts to change the state of the key based on the request and responds to the
// client with the result of the operation. When the client is finished operating on
// the server keys, it terminates the protocol by sending a request to close the RPC
// context handle.
package rrp

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
	GoPackage = "rrp"
)
