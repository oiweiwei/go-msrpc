// The brwsa package implements the BRWSA client protocol.
//
// # Introduction
//
// This document specifies the Common Internet File System (CIFS) Browser Auxiliary
// Protocol Specification. This protocol is used by the master browser server and domain
// master browser server as defined in [MS-BRWS]. The master browser server uses this
// protocol to query configuration information for the domains from the domain master
// browser server. The protocol operation is stateless.
//
// # Overview
//
// The main objective of the CIFS Browser Auxiliary Protocol is to provide a method
// for the master browser server of a subnet to query specific additional information
// from the domain master browser server for a given domain. Selection of the master
// browser server and domain master browser server and the roles that these servers
// play are as specified in [MS-BRWS].
package brwsa

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
	GoPackage = "brwsa"
)
