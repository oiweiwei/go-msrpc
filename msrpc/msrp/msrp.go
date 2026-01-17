// The msrp package implements the MSRP client protocol.
//
// # Introduction
//
// This document specifies the Messenger Service Remote Protocol. The Messenger Service
// Remote Protocol is a set of remote procedure call (RPC) interfaces that instruct
// a server (referred to in this document as a "message server") to perform one or more
// of the following tasks:
//
// * Deliver messages to a local or remote message server for display to a console user.
//
// * Manage the names for which the message server receives messages.
//
// The message server does not maintain client state information.
//
// It is recommended that this protocol not be implemented due to the lack of security
// features in the protocol, as described in section 5.1.
//
// # Overview
//
// The Messenger Service Remote Protocol suite is designed to perform the following
// functions:
//
// * Receive and display short text messages to the console user. (This function is
// referred to in this document as the "messaging protocol".)
//
// * Manage the names for which a message server ( 1bb1da94-fc7e-484c-ad2a-4e0d257a50ce#gt_18623b87-0e55-44bc-a5d3-d49388e1716a
// ) receives messages. (This function is referred to in this document as the "name
// management protocol".)
//
// The name management protocol portion of the Messenger Service Remote Protocol is
// used to manage the set of names for which the message server accepts messages. The
// operations in this protocol are very simple, consisting of add, remove, and enumeration
// methods. The messaging protocol portion of the Messenger Service Remote Protocol
// actually has several forms and runs over mailslots over Server Message Block Protocol,
// as specified in [MS-SMB] and RPC dynamic endpoints over User Datagram Protocol (UDP)
// (as specified in [RFC768]). For how the message client selects the transport that
// is used for the messaging protocol, see section 3.2.
//
// Typically, the Messenger Service Remote Protocol is used to send a short text message
// from a server, such as a file server or print server, to a client machine; for example,
// to indicate that a print job has completed or that a file server is shutting down
// and all of its clients need to save their work and disconnect. As such, the roles
// of client and server are reversed from typical protocols, with the message server
// (recipient) of the messages often being the workstation machine and the message client
// (sender) being a server-class machine.
package msrp

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
	GoPackage = "msrp"
)
