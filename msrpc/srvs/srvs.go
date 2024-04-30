// The srvs package implements the SRVS client protocol.
//
// # Introduction
//
// This document specifies the Server Service Remote Protocol. The Server Service Remote
// Protocol is a remote procedure call (RPC)–based protocol that is used for remotely
// enabling file and printer sharing and named pipe access to the server through the
// Server Message Block (SMB) Protocol, as specified in [MS-SMB]. The protocol is also
// used for remote administration of servers that are running Windows.
//
// # Overview
//
// The Server Service Remote Protocol is designed for remotely querying and configuring
// a Server Message Block (SMB) server on a remote computer. By using this protocol,
// a client can query and configure information on the server such as active connections,
// sessions, shares, files, and transport protocols. Clients can also query and configure
// the server itself, for instance by setting the server's type, changing the services
// that are running on the server, or getting a list of all servers of a specific type
// in a domain.
//
// A server can be configured to present different resources based on the name the client
// connects with, allowing it to appear as multiple, distinct servers. This is achieved
// by scoping a share to a specific name, and hosting all of the names on the same server.
//
// The server can also configure one or more aliases, identifying that multiple distinct
// names present the same resources. For example, the administrator could choose to
// expose the same shares for the name "server" and "server.example.com" by creating
// an alias indicating that "server. example.com" is the same as "server". The SMB client
// will connect using the name provided by the calling applications, and is not aware
// whether the name is the server's default machine name, an additionally configured
// name, or an alias. For more information, see the example in section 4.3.
//
// This is an RPC-based protocol. The server does not maintain client state information.
// No sequence of method calls is imposed on this protocol, with the exception of net
// share deletion, which requires a two-phase commit, net file get information, and
// net file close.
package srvs

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
	GoPackage = "srvs"
)
