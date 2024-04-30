// The tsgu package implements the TSGU client protocol.
//
// # Introduction
//
// The Remote Desktop Gateway Server Protocol (RDGSP Protocol)<1>  is used primarily
// for tunneling client to server traffic across firewalls when the Remote Desktop Gateway
// (RDG)<2> server is deployed in the neutral zone of a network. The primary consumer
// of the Terminal Services Gateway Server Protocol is the Remote Desktop Protocol:
// Basic Connectivity and Graphics Remoting [MS-RDPBCGR].
//
// The RDGSP Protocol uses either Hypertext Transfer Protocol (HTTP) or remote procedure
// call (RPC) over HTTP as the transport for establishing the main channel. The protocol
// uses User Datagram Protocol (UDP) as the transport for establishing the side channel
// which is established only when the main channel uses HTTP.
//
// # Overview
//
// The RDGSP Protocol is designed for remote connections from RDG clients originating
// on the Internet to target servers behind a firewall.<3> The protocol establishes
// a connection, called a tunnel (2), from an RDG client to an RDG server in the neutral
// zone. The RDG client uses the tunnel to establish a channel between the RDG client
// and the target server with the RDG server acting as a proxy. Data transfer between
// the RDG client and the target server occurs by using the channel. The tunnel and
// channel maintain active connections.
//
// The RDG client establishes one main channel to the target server. The RDG client
// can establish zero or more side channels depending on the requirements of the Remote
// Desktop Protocol: UDP Transport Extension Protocol [MS-RDPEUDP].
//
// The RDGSP Protocol uses one of the following transports for establishing the main
// channel:
//
// * Remote Procedure Call over HTTP Protocol (RPC over HTTP)
//
// * HTTP ( 76796f19-9e6f-48b9-8b8f-4ef9f197056b#gt_d72f1494-4917-4e9e-a9fd-b8f1b2758dcd
// )
//
// The RDGSP Protocol uses the UDP transport for establishing the side channel.
//
// In this specification, information that is common to all three transport types (RPC
// over HTTP, HTTP, and UDP) is provided at the beginning of each main section and details
// for each transport type are defined in transport-specific subsections that follow
// the main section. The subsections are distinguished as follows:
//
// * Details specific to the RDGHTTP Protocol are documented in subsections that include
// the phrase "HTTP Transport" in the title.
//
// * Details specific to the RDGUDP Protocol are documented in subsections that include
// the phrase "UDP Transport" in the title.
package tsgu

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
	GoPackage = "tsgu"
)
