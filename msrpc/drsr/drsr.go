// The drsr package implements the DRSR client protocol.
//
// # Introduction
//
// The Directory Replication Service (DRS) Remote Protocol is an RPC protocol for replication
// and management of data in Active Directory.
//
// The protocol consists of two RPC interfaces named drsuapi and dsaop. The name of
// each drsuapi method begins with "IDL_DRS", while the name of each dsaop method begins
// with "IDL_DSA". This protocol was originally implemented in Windows 2000 Server operating
// system and is available in all subsequent server releases. It is not available in
// Windows NT 3.1 operating system, Windows NT 3.51 operating system, or Windows NT
// 4.0 operating system.
//
// Some functionality exposed by these RPC protocols is also available using the Lightweight
// Directory Access Protocol (LDAP) protocol ([MS-ADTS] section 3.1.1.3); the overlap
// is described in section 1.4.
//
// The special typographical conventions used in this document are described in section
// 3.2.
//
// State is included in the state model for this specification only as necessitated
// by the requirement that a licensee implementation of the protocols of applicable
// Windows Server releases has to be capable of receiving messages and responding in
// the same manner as applicable Windows Server releases. Behavior is specified in terms
// of request message received, processing based on current state, resulting state transformation,
// and response message sent. Unless otherwise specified, all behaviors are required
// elements of the protocol. Any specified behavior not explicitly qualified with MAY
// or SHOULD is to be treated as if it were specified as a MUST behavior.
//
// # AD LDS for Windows Client operating systems
//
// Note that information that is applicable to Active Directory Lightweight Directory
// Services (AD LDS) on applicable Windows Server releases is also generally applicable
// to AD LDS on Windows clients. For more information, see [MS-ADTS] section 1.
//
// # Pervasive Concepts
//
// The following concepts are pervasive throughout this specification.
//
// This specification uses [KNUTH1] section 2.3.4.2 as a reference for the graph-related
// terms oriented tree, root, vertex, arc, initial vertex, and final vertex.
//
// # Overview
//
// This document specifies the Directory Replication Service (DRS) Remote Protocol,
// an RPC protocol for replication between domain controllers and management of Active
// Directory. The protocol consists of two RPC interfaces, named drsuapi and dsaop.
// The name of each drsuapi method begins with "IDL_DRS", while the name of each dsaop
// method begins with "IDL_DSA".
package drsr

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
	GoPackage = "drsr"
)
