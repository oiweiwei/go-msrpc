// The even package implements the EVEN client protocol.
//
// # Introduction
//
// The EventLog Remoting Protocol is an RPC-based protocol that exposes remote procedure
// call (RPC) methods for reading events in both live event logs and backup event logs
// on remote computers. The protocol also specifies how to get general information on
// a log, such as the number of records in the log, the oldest records in the log, and
// if the log is full. The protocol can also be used for clearing and backing up both
// types of event logs.
//
// Note  Early releases of the EventLog Remoting Protocol have never been assigned a
// version number. However, newer releases of the EventLog Remoting Protocol have version
// numbers. For example, the version released with Windows Vista operating system is
// version 6.0.
package even

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
	GoPackage = "even"
)
