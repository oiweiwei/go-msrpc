// The dssp package implements the DSSP client protocol.
//
// # Introduction
//
// The Directory Services Setup Remote Protocol is a client/server-based remote procedure
// call (RPC) protocol. The protocol exposes an RPC interface that a client can call
// to obtain domain-related computer state and configuration information.
//
// # Overview
//
// This protocol provides a remote procedure call (RPC) interface for querying domain-related
// computer state and configuration data. The client end of the Directory Services Setup
// Remote Protocol is an application that issues method calls on the RPC interface.
// The server end of the Directory Services Setup Remote Protocol obtains and replies
// to the client with the requested data about the computer on which the server is running.
// If the client connects to and requests information about a domain controller (DC)
// for the directory service, this data includes the status of any pending promotion
// or demotion of that DC.
package dssp

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
	GoPackage = "dssp"
)
