// The oxabref package implements the OXABREF client protocol.
//
// # Introduction
//
// The Address Book Name Service Provider Interface (NSPI) Referral Protocol defines
// a remote procedure call (RPC) service that supplies a caller with the name of an
// NSPI server. Additionally, this protocol can return the Domain Name System (DNS)
// fully qualified domain name (FQDN) of a mailbox server, given the distinguished name
// (DN) of that server.
//
// # Overview
//
// This protocol enables clients to retrieve the network name of a server from a name
// service provider interface (NSPI) referral server. Clients use this protocol before
// performing any NSPI requests, in order to retrieve the name of the NSPI server to
// connect to. This gives the NSPI referral server the ability to control which NSPI
// server an NSPI client will connect to, for purposes including but not limited to
// balancing the client load across multiple NSPI servers, choosing the best version
// of NSPI server for that particular client, or satisfying network requirements that
// are not discernible by the client. Clients also use this protocol to retrieve the
// FQDN of the mailbox server, when only the DN the mailbox server is known. Figure
// 1 shows the request to the NSPI referral server for the name of the NSPI server and
// the server’s response to the client. Figure 2 shows the request to the NSPI referral
// server for the FQDN of the mailbox server and the server’s response to the client.
//
// # Figure 1 Client retrieving NSPI server name from the NSPI referral server
//
// Figure 2 Client retrieving mailbox server name from the NSPI referral server
package oxabref

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
	GoPackage = "oxabref"
)
