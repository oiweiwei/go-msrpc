// The wkst package implements the WKST client protocol.
//
// # Introduction
//
// The Workstation Service Remote Protocol is used to perform tasks on a computer remotely
// on a network, including:
//
// * Configuring properties and behavior of a Server Message Block ( 3acf0e02-9bbd-4ce0-a7a0-586bc72d3ef4#gt_09dbec39-5e75-4d9a-babf-1c9f1d499625
// ) network redirector ( 3acf0e02-9bbd-4ce0-a7a0-586bc72d3ef4#gt_15c89cb5-6767-49e7-b461-66acaf6c06c8
// ) (SMB network redirector).
//
// * Managing domain ( 3acf0e02-9bbd-4ce0-a7a0-586bc72d3ef4#gt_b0276eb2-4e65-4cf1-a718-e0920a614aca
// ) membership and computer names ( 3acf0e02-9bbd-4ce0-a7a0-586bc72d3ef4#gt_9a7bc8b3-3374-4608-8f73-be20a90b898b
// ).
//
// * Gathering information, such as the number of enabled transport protocols and the
// number of currently logged-on users.
//
// This protocol is based on the Remote Procedure Call (RPC) protocol [C706] [MS-RPCE].
//
// # Overview
//
// The Workstation Service Remote Protocol is designed for remotely querying and configuring
// certain aspects of an SMB network redirector on a remote computer. For example, an
// implementer can use this protocol to query the computer name or major and minor version
// numbers of the operating system running on a remote computer.
//
// An implementer can also use the protocol to configure the behavior of an SMB network
// redirector. For example, an implementer can use this protocol to configure the following:
//
// * The number of seconds the SMB network redirector maintains an inactive SMB connection
// ( 3acf0e02-9bbd-4ce0-a7a0-586bc72d3ef4#gt_e1d88514-18e6-4e2e-a459-20d5e17e9078 )
// to a remote computer's resource before closing it.
//
// * The number of simultaneous network commands that can be sent to the SMB network
// redirector.
//
// * The number of seconds the SMB network redirector waits before disconnecting an
// inactive SMB session ( 3acf0e02-9bbd-4ce0-a7a0-586bc72d3ef4#gt_ee1ec898-536f-41c4-9d90-b4f7d981fd67
// ).
//
// The protocol is also designed to enumerate all the users currently logged onto a
// remote computer, and to enumerate the transport protocols currently enabled for use
// by the SMB network redirector on a remote computer. When enumerating currently logged-on
// users or transport protocols, the protocol does not guarantee that all logged-on
// users or transport protocols are enumerated. The protocol also does not guarantee
// that the enumerated users or transport protocols are not duplicated.
//
// The protocol can also be used to manage domain membership and the computer names
// of a computer on a network. For example, this protocol can be used to configure the
// following:
//
// * The primary name of a computer
//
// * Alternate names of a computer
//
// * The domain membership of a computer
//
// This is an RPC-based protocol. This protocol contains no protocol-specific state
// that is stored across protocol messages and only operates on state accessible through
// other protocols and local services. Some methods manipulate the server state and
// the state at a domain controller (DC) during message processing. This state is not
// part of this protocol but is exposed by other protocols.
//
// This is a simple request-response protocol. For every method call that the server
// (2) receives, it executes the method and returns a completion. The client simply
// returns the completion status to the caller. Each method call is independent of any
// previous method call.
package wkst

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
	GoPackage = "wkst"
)
