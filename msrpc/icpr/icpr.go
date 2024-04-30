// The icpr package implements the ICPR client protocol.
//
// # Introduction
//
// This document specifies the ICertPassage Remote Protocol. This protocol is a subset
// of the Windows Client Certificate Enrollment Protocol, as specified in [MS-WCCE].
// The difference between this protocol and the Windows Client Certificate Enrollment
// Protocol is that this protocol only allows the client to enroll certificates, whereas
// the Windows Client Certificate Enrollment Protocol provides enrollment and additional
// functionality, such as the capability to read certification authority (CA) data and
// configuration information. Reading and understanding the Windows Client Certificate
// Enrollment Protocol, as specified in [MS-WCCE], is essential to understanding the
// ICertPassage Remote Protocol.
//
// # Overview
//
// The ICertPassage Remote Protocol exposes a Remote Procedure Call (RPC) (as specified
// in [MS-RPCE]) interface that allows a client to interact with a certification authority
// (CA) to request and receive X.509 certificates (as specified in [X509]) from the
// CA. The ICertPassage Remote Protocol only provides certificate enrollment functionality.
// The Windows Client Certificate Enrollment Protocol (as specified in [MS-WCCE]) provides
// a larger set of functionality, including reading CA data and configuration information.
// The certificate enrollment process and protocol overview are as specified in [MS-WCCE]
// section 1.3.
//
// The ICertPassage interface defines one method: CertServerRequest (section 3.2.4.1.1).
package icpr

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
	GoPackage = "icpr"
)
