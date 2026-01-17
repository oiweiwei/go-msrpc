// The lsat package implements the LSAT client protocol.
//
// # Introduction
//
// The Local Security Authority (Translation Methods) Remote Protocol is implemented
// in Windows products to translate identifiers for security principals between human-readable
// and machine-readable forms. This translation can be used in scenarios such as human
// management of resource access.
//
// # Overview
//
// The purpose of this protocol is to translate human-readable names to security identifiers
// (SIDs), as specified in [MS-DTYP] section 2.4.2, and vice versa. The syntax of human-readable
// names is specified in section 3.1.4.5. For example, this protocol can be used to
// translate "corp\John" to "S-1-5-21-397955417-626881126-188441444-1555", and vice
// versa.
//
// This translation is needed for different scenarios, such as managing access to resources.
// In Windows, access to resources is controlled by a security descriptor attached to
// the resource. This security descriptor contains a list of SIDs indicating the security
// principals and the kind of access allowed or denied for those principals. In order
// for humans to manage access to resources, translation occurs between SIDs (persisted
// within security descriptors) and human-readable identifiers (in the user interface).
//
// This protocol is intended for use between any two machines, most frequently between
// a domain member and a domain controller (DC) for that domain. This protocol can also
// be used between domain controllers for trusting domains or forests.
//
// This protocol is a simple request/response-based remote procedure call (RPC) protocol.
// There are no long-lived sessions, although clients are free to cache the RPC connection
// and reuse it over time. A sample sequence of requests and responses is shown in section
// 4.
package lsat

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
	GoPackage = "lsat"
)
