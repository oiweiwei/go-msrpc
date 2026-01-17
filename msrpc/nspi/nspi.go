// The nspi package implements the NSPI client protocol.
//
// # Introduction
//
// The Name Service Provider Interface (NSPI) Protocol provides messaging clients a
// way to access and manipulate addressing data stored by a server. This protocol consists
// of an abstract data model and a single remote procedure call (RPC) interface to manipulate
// data in that model.
//
// # Overview
//
// Messaging clients that implement a browsable address book need a way to communicate
// with a data store that holds addressing data to access and manipulate that data.
// The NSPI Protocol enables communication between a messaging client and a data store.<1>
//
// The NSPI Protocol is a protocol layer that uses the remote procedure call (RPC) protocol
// as a transport, with a series of interface methods as specified in this document,
// that clients can use to communicate with an NSPI server.
//
// The following diagram is a graphical representation of a typical communication sequence
// between a messaging client and an NSPI server.
package nspi

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
	GoPackage = "nspi"
)
