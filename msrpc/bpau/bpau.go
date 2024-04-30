// The bpau package implements the BPAU client protocol.
//
// # Introduction
//
// The Background Intelligent Transfer Service (BITS) Peer-Caching: Peer Authentication
// Protocol provides authentication for computers in a domain in support of the BITS
// Peer-Caching: Content Retrieval Protocol, as specified in [MS-BPCR]. Peer authentication
// exchanges X.509 Certificate Authority between computers and associates each certificate
// with a Kerberos principal in the domain.
//
// # Overview
//
// The BITS Peer-Caching: Peer Authentication Protocol allows hosts in an Active Directory
// domain to exchange self-signed X.509 certificates with enough information to associate
// those certificates securely with a domain account.
//
// Peer authentication is intended for use by hosts that implement the BITS Peer-Caching:
// Content Retrieval Protocol, as specified in [MS-BPCR].
//
// Peer authentication uses the Kerberos security system for authentication, allowing
// each host to do the following:
//
// * Verify that the peer is allowed to participate in content retrieval.
//
// * Associate the received certificate with the peer's Kerberos identity in a trustworthy
// way.
//
// This protocol is used as part of a distributed peer-to-peer cache of URL content
// for use by the Background Intelligent Transfer Service (BITS) component. (For more
// information on BITS, see [MSDN-BITS].) Peer authentication ensures that peer clients
// and servers are members of the same domain, or in domains with bidirectional trust.
package bpau

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
	GoPackage = "bpau"
)
