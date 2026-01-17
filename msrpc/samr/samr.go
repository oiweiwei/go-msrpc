// The samr package implements the SAMR client protocol.
//
// # Introduction
//
// The Security Account Manager (SAM) Remote Protocol (Client-to-Server) provides management
// functionality for an account store or directory containing users and groups. The
// goal of this protocol is to enable IT administrators and end users to manage users,
// groups, and computers.
//
// SAM Remote Protocol specifies the behavior for both local and remote domains by using
// the Active Directory data model as a common data model for both scenarios. SAM Remote
// Protocol also specifies the differences in behavior between these two scenarios.
//
// # Overview
//
// The Security Account Manager (SAM) Remote Protocol (Client-to-Server) provides management
// functionality for an account store or directory containing users and groups. SAM
// Remote Protocol specifies the behavior for both local and remote domains by using
// the Active Directory data model as a common data model for both scenarios, see [MS-AUTHSOD]
// section 1.1.1.5.1. Implementers should familiarize themselves with the following
// documents: Windows System Overview [MS-SYS-ARCHIVE], Windows Protocols Overview [MS-WPO],
// and Active Directory Technical Specification [MS-ADTS].
//
// The goal of this protocol is to enable IT administrators and end users to manage
// users, groups, and computers. IT administrators and their delegates generally have
// full access control to these entities, and consequently can manage the entities'
// life cycles. End users are allowed to make changes to their own data (in most cases,
// limited to just their passwords).
//
// SAM Remote Protocol achieves its goal by enabling the creation, reading, updating,
// and deleting of security principal information. These security principals could be
// in any account store; typically they are in an Active Directory service ([MS-ADTS])
// and in a computer-local security account database. Normative differences in the protocol
// between the two cases local and remote domains are indicated by referring to the
// configuration of the server as a DC or non-DC configuration, where DC stands for
// domain controller (DC).
//
// It is helpful to consider the following two perspectives when understanding and implementing
// this protocol:
//
// * Object-based perspective (section 1.3.1 ( 8aaff2f7-1edd-41a0-ab58-4807ac6124c5
// ) )
//
// * Method-based perspective (section 1.3.2 ( d7b62596-4a46-4556-92dc-3aba6d517907
// ) )
package samr

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
	GoPackage = "samr"
)
