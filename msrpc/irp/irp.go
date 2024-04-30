// The irp package implements the IRP client protocol.
//
// # Introduction
//
// The Internet Information Services (IIS) Inetinfo Remote Protocol is a remote procedure
// call (RPC)–based client/server protocol that is used for managing Internet protocol
// servers such as those hosted by Microsoft Internet Information Services (IIS). Managed
// servers can include servers for HTTP, FTP, SMTP, or other Internet protocols. For
// more information on IIS, see [MSDN-IIS].
//
// The universally unique identifier (UUID) for the IIS Inetinfo Remote Protocol interface
// is {82ad4280-036b-11cf-972c-00aa006887b0}.
//
// The version for this interface is 2.0.
//
// # Overview
//
// The Internet Information Services (IIS) Inetinfo Remote Protocol provides functions
// that allow remote administration and statistics gathering from an Internet protocol
// server such as a server implementing the HTTP or FTP protocol. The protocol provides
// methods for gathering statistical data on users, sites, requests, and performance.
// For more information about HTTP and securing HTTP connections, see [RFC2068] and
// [RFC2818].
//
// The server does not maintain client state information. Although some client call
// sequences might be logically related, the protocol operation is stateless.
package irp

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
	GoPackage = "irp"
)
