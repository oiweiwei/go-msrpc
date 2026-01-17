// The bkrp package implements the BKRP client protocol.
//
// # Introduction
//
// The BackupKey Remote Protocol is used by clients to encrypt and decrypt sensitive
// data (such as cryptographic keys) with the help of a server. Data encrypted using
// this protocol can be decrypted only by the server, and the client can safely write
// such encrypted data to storage that is not specially protected. In Windows, this
// protocol is used to provide encryption of user secrets through the Data Protection
// Application Program Interface (DPAPI) in an Active Directory Domain.
//
// Familiarity with cryptography and Public Key Infrastructure (PKI) concepts (such
// as asymmetric and symmetric cryptography, digital certificate concepts, and cryptographic
// key exchange) is required for a complete understanding of this specification. For
// more information about cryptography and PKI concepts, see [CRYPTO].
//
// # Overview
//
// The BackupKey Remote Protocol provides a method of protecting a secret value so that
// the value can be stored in a potentially insecure location, while still being recoverable
// by an authorized user. The protocol does this by encrypting the secret with the assistance
// of a server, in a process known as wrapping. When an authorized user wants to access
// the secret, the user authenticates to the server and submits the wrapped data to
// the server. The server then extracts the original secret in a process known as unwrapping,
// and returns it to the user.
//
// As the name indicates, this protocol was designed specifically to wrap and unwrap
// cryptographic keys. Within the Windows implementation, this protocol is used by the
// Data Protection Application Program Interface (DPAPI) on a client in an Active Directory
// domain and a Domain Controller (DC) in the same domain to wrap cryptographic keys.
// However, all of this protocol's variants will wrap arbitrary secrets. Nothing in
// the protocol requires the secrets to be cryptographic keys or to have any particular
// structure, other than a limitation that is imposed on the length of the secret in
// certain cases. This limitation is specified in section 2.2.2.2.
//
// The BackupKey Remote Protocol consists of two subprotocols, each of which enables
// the client to perform a wrapping operation and a corresponding unwrapping operation.
// In the ServerWrap subprotocol, both the wrapping and unwrapping operations are performed
// through a protocol exchange with a server supporting this subprotocol. On the other
// hand, the server side of the ClientWrap subprotocol consists of a key retrieval method
// and an unwrapping method. Thus, a client can perform the unwrapping operation of
// the ClientWrap subprotocol only through a protocol exchange with a server that supports
// this subprotocol. However, a client can perform the wrapping operation of the ClientWrap
// subprotocol purely locally using public key cryptography, provided that it has in
// the past retrieved a key from a server that supports this subprotocol.
//
// A BackupKey Remote Protocol client or server can implement either or both of these
// subprotocols, and in each case it can implement the entire subprotocol or only the
// unwrapping operation. However, a client or server has to always support unwrapping
// any secrets whose wrapping it performed or enabled. Thus, a server that supports
// ServerWrap wrapping has to also support ServerWrap unwrapping, and a server that
// supports ClientWrap key retrieval has to also support ClientWrap unwrapping. Similarly,
// a client that supports the wrapping operation of either subprotocol has to also support
// the corresponding unwrapping operation.
//
// It is important to note that a BackupKey Remote Protocol server does not actually
// perform remote backup of secrets. Instead, the server wraps each secret and returns
// it to the client. The client is responsible for storing the secret until it is needed
// again, at which point the client can request the server to unwrap the secret.
//
// The BackupKey Remote Protocol uses remote procedure call (RPC) [C706] with the security
// provider extensions for user impersonation and connection encryption and authentication
// specified in [MS-RPCE]. Named pipes over the Server Message Block (SMB) Protocol
// are used as transport. SPNEGO [RFC4178] [MS-SPNG] is used to negotiate an authentication
// mechanism between client and server.
package bkrp

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
	GoPackage = "bkrp"
)
