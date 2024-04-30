// The fsrvp package implements the FSRVP client protocol.
//
// # Introduction
//
// The File Server Remote VSS Protocol (FSRVP) is a remote procedure call (RPC)-based
// protocol that is used for creating shadow copies of file shares on a remote computer.
// This protocol facilitates the backup applications' tasks in performing application-consistent
// backup and restore of VSS-aware applications storing data on network file shares.
//
// # Overview
//
// The File Server Remote VSS Protocol is designed to remotely create shadow copies
// of file shares hosted on a file server. This facilitates applications hosting their
// data on a file server to back up and restore their application state. The client-side
// implementation of this protocol typically runs on an application server and the server-side
// implementation runs on a file server. This protocol is modeled in such way that the
// client-side and server-side implementation can be integrated with existing volume
// shadow copy creation utilities.
package fsrvp

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
	GoPackage = "fsrvp"
)
