// The gkdi package implements the GKDI client protocol.
//
// # Introduction
//
// The Group Key Distribution Protocol is used by clients to obtain cryptographic keys
// that correspond to arbitrary security descriptors that can be evaluated by an Active
// Directory domain controller (DC). These keys can then be used by the client for various
// purposes, including encrypting data such that it can only be decrypted by a desired
// set of security principals.
//
// Familiarity with cryptography concepts such as asymmetric and symmetric cryptography
// is required for a complete understanding of this specification. For more information
// about cryptography concepts, see [CRYPTO].
//
// # Overview
//
// The Group Key Distribution Protocol is used to obtain cryptographic keys corresponding
// to arbitrary security descriptors that can be evaluated by an Active Directory DC.
// It can be used to obtain symmetric as well as asymmetric keys for each of such security
// descriptors. One possible use of this protocol is to obtain shared keys for a set
// of security principals that are defined by the client in the form of a security descriptor.
// Based on an evaluation of the client's security context and the security descriptor,
// the server can return an error, a public key, or a seed key that can be used to derive
// both the symmetric and asymmetric keys. Whenever the server returns a key of any
// type, it also returns metadata that includes a unique identifier for the key.
//
// The Group Key Distribution Protocol utilizes a single remote procedure call (RPC)
// method that is described in section 3.1.4.1. Conceptually, this method can be used
// by a client for two types of requests:
//
// *
//
// Requesting the most recent key for a security descriptor: after evaluating the client's
// security context against the specified security descriptor, the server will return
// a seed key, a public key, or an error.
//
// *
//
// Requesting a specific key for a security descriptor, or the key used for a set of
// security principals at a specific time in the past: after evaluating the client's
// security context against the specified security descriptor, the server will return
// either a seed key or an error.
//
// Active Directory domain controllers with a DC functional level of DS_BEHAVIOR_WIN2012
// or higher can serve as Group Key Distribution Protocol servers. Clients can locate
// Group Key Distribution Protocol servers by using the DC locator functionality, as
// specified in section 3.2.4.2. These servers store a small amount of state in Active
// Directory (sections 2.3 and 3.1.1), which consists of configuration information and
// one or more root key objects. Other than this, Group Key Distribution Protocol servers
// retain no state across RPC calls.
package gkdi

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
	GoPackage = "gkdi"
)
