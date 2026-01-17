// The claims package implements the CLAIMS client protocol.
//
// # Introduction
//
// This is the primary specification for Active Directory, both Active Directory Domain
// Services (AD DS) and Active Directory Lightweight Directory Services (AD LDS). When
// the specification does not refer specifically to AD DS or AD LDS, it applies to both.
// The state model for this specification is prerequisite to the other specifications
// for Active Directory: [MS-DRSR] and [MS-SRPL].
//
// When no operating system version information is specified, information in this document
// applies to all relevant versions of Windows. Similarly, when no DC functional level
// is specified, information in this document applies to all DC functional levels.
//
// The information in this specification is applicable to the following Microsoft products
// or supplemental software. References to product versions include released service
// packs.
//
// # Overview
//
// This is the primary specification for Active Directory. The state model for this
// specification is prerequisite to the other specifications for Active Directory: [MS-DRSR]
// and [MS-SRPL].
//
// Active Directory is either deployed as AD DS or as AD LDS. This document describes
// both forms. When the specification does not refer specifically to AD DS or AD LDS,
// it applies to both.
//
// The remainder of this section describes the structure of this document.
//
// The basic state model is specified in section 3.1.1.1. The basic state model is prerequisite
// to the remainder of the document. Section 3.1.1.1 also includes descriptive content
// to introduce key concepts and refer to places in the document where the full specification
// is given.
//
// The schema completes the state model and is specified in section 3.1.1.2. The schema
// is prerequisite to the remainder of the document.
//
// Active Directory is a server for LDAP. Section 3.1.1.3 specifies the extensions and
// variations of LDAP that are supported by Active Directory.
//
// LDAP is an access protocol that determines very little about the behavior of the
// data being accessed. Section 3.1.1.4 specifies read (LDAP Search) behaviors, and
// section 3.1.1.5 specifies update (LDAP Add, Modify, Modify DN, Delete) behaviors.
// Section 3.1.1.6 specifies background tasks required due to write operations, to the
// extent that those tasks are exposed by protocols.
//
// One of the update behaviors is the maintenance of the change log for use by Windows
// NT 4.0 operating system backup domain controller (BDC) replication [MS-NRPC] section
// 3.6. The maintenance of this change log is specified in section 3.1.1.7.
//
// The security services that Active Directory offers clients of LDAP are specified
// in section 5.1.
//
// Active Directory contains a number of objects, visible through LDAP, that have special
// significance to the system. Section 6.1 specifies these objects.
//
// A server running Active Directory is part of a distributed system that performs replication.
// The Knowledge Consistency Checker (KCC) is a component that is used to create spanning
// trees for DC-to-DC replication and is specified in section 6.2.
//
// A server running Active Directory is responsible for publishing the services that
// it offers, in order to eliminate the administrative burden of configuring clients
// to use particular servers running Active Directory. A server running Active Directory
// also implements the server side of the LDAP ping and mailslot ping protocols to aid
// clients in selecting among all the servers offering the same service. Section 6.3
// specifies how a server running Active Directory publishes its services, and how a
// client needing some service can use this publication plus the LDAP ping or mailslot
// ping to locate a suitable server.
//
// Computers in a network with Active Directory can be put into a state called "domain
// joined"; when in this state, the computer can authenticate itself. Section 6.4 specifies
// both the state in Active Directory and the state on a computer required for the domain
// joined state.
//
// Each type of data stored in Active Directory has an associated function that compares
// two values to determine if they are equal and, if not, which is greater. Section
// 3.1.1.2 specifies all but one of these functions; the methodology for comparing two
// Unicode strings is specified in section 6.5.
package claims

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
	GoPackage = "adts/claims"
)
