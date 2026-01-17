// The eerr package implements the EERR client protocol.
//
// # Introduction
//
// This specification for encoding extended error information assumes that the reader
// has familiarity with the concepts and the requirements that are detailed in [MS-RPCE]
// and [C706].
//
// The purpose of the encoding that this specification defines is to allow a software
// agent on one network node to communicate a rich (or extended) error to a software
// agent on another network node. This specification does not define how an extended
// error is transmitted between network nodes. A protocol outside this specification
// is used for that purpose. This specification only defines the encoding rules for
// an extended error.
//
// # Overview
//
// In complex distributed systems, a situation can arise where an error encountered
// on one network node has to be communicated to another network node. A protocol that
// is used to transmit data between network nodes usually has some provisions to transmit
// errors in its messages, but often the error that is being communicated is a single
// unsigned integer or a single unsigned integer plus a short string. As the complexity
// of the system and/or the number of network nodes that are involved grows, a single
// unsigned integer and/or a short string might prove insufficient for quick and efficient
// troubleshooting of all possible scenarios.
//
// This specification defines an encoding for a rich, structured error called an extended
// error. After the extended error is encoded, it has to be transmitted between network
// nodes by a protocol outside this specification.
//
// The extended error itself is used for troubleshooting a malfunctioning system and
// is intended to be used by a human reader or an automated failure diagnostic system.
// This specification does not prescribe what the extended error should be; it specifies
// the fields and field values that are used for encoding the extended error (see section
// 2). Protocols and systems are free to create and encode any extended error a support
// engineer or an expert user of the system might find useful to troubleshoot a malfunctioning
// system.
package eerr

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
	GoPackage = "eerr"
)
