// The cmpo package implements the CMPO client protocol.
//
// # Introduction
//
// This document specifies the MSDTC Connection Manager: OleTx Transports Protocol.
// The MSDTC Connection Manager: OleTx Transports Protocol is a remote procedure call
// (RPC) interface for establishing duplex sessions between two partners and for exchanging
// messages between them. The MSDTC Connection Manager: OleTx Transports Protocol is
// a framing and message transport protocol and, as such, is designed to have other
// protocols layered over the basic session, messaging, and security services that it
// provides.
//
// # Overview
//
// The MSDTC Connection Manager: OleTx Transports Protocol is a peer-to-peer messaging
// protocol layered over a bidirectional pair of RPC connections. Although there is
// asymmetry in the setup and teardown of a session, the peers (or partners) are considered
// equal for the purposes of sending messages to each other.
//
// Together, the pair of RPC connections between the partners is called a session.
package cmpo

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
	GoPackage = "cmpo"
)
