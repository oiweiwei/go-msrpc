// The nrpc package implements the NRPC client protocol.
//
// # Introduction
//
// The Netlogon Remote Protocol is a remote procedure call (RPC) interface that is used
// for user and machine authentication on domain-based networks. The Netlogon Remote
// Protocol RPC interface is also used to replicate the database for backup domain controllers
// (BDCs).
//
// The Netlogon Remote Protocol is used to maintain domain relationships from the members
// of a domain to the domain controller (DC), among DCs for a domain, and between DCs
// across domains. This RPC interface is used to discover and manage these relationships.
//
// # Overview
//
// The Netlogon Remote Protocol is used for secure communication between machines in
// a domain and domain controllers (DCs) (both domain members and DCs). The communication
// is secured by using a shared session key computed between the client and the DC that
// is engaged in the secure communication. The session key is computed by using a preconfigured
// shared secret that is known to the client and the DC.
//
// The Netlogon Remote Protocol client and server can only run on domain-joined systems
// and are started during boot. When a system is unjoined from the domain, then the
// client and server are stopped and will not be started during boot.
//
// The following sections describe the scenarios in which this protocol is used. It
// provides an overview about the general purpose of this protocol and the flow of its
// operations.
package nrpc

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
	GoPackage = "nrpc"
)
