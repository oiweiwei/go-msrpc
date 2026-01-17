// The rdpesc package implements the RDPESC client protocol.
//
// # Introduction
//
// This document specifies an extension (including virtual channels) to the Remote Desktop
// Protocol: File System Virtual Channel Extension for supporting smart card reader-like
// devices.
//
// # Overview
//
// The following figure illustrates a baseline for terminology related to clients and
// servers.
package rdpesc

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
	GoPackage = "dcom/rdpesc"
)
