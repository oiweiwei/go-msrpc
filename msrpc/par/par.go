// The par package implements the PAR client protocol.
//
// # Introduction
//
// The Print System Asynchronous Remote Protocol supports printing and spooling operations
// between a client and server, including print job control and print system management.
// It is designed to be used asynchronously by clients whose implementations permit
// them to continue execution without waiting for a method call to return.
//
// This protocol is based on the Remote Procedure Call (RPC) protocol [C706] [MS-RPCE].
// Its functionality is parallel to the Print System Remote Protocol [MS-RPRN], which
// is used synchronously.
//
// # Overview
//
// The Print System Asynchronous Remote Protocol provides the following functions:
//
// * Management of the print system ( a98de8d9-8e9e-4cd7-92f7-f9269206a41e#gt_b49fcfa1-e15c-4dce-be33-d74d5bfda785
// ) of a print server ( a98de8d9-8e9e-4cd7-92f7-f9269206a41e#gt_59fb3ddc-63cf-45df-8a90-46a6af9e00cb
// ) from a client.
//
// * Communication of print job ( a98de8d9-8e9e-4cd7-92f7-f9269206a41e#gt_12a6e569-e97c-4761-92f0-e397f8d5125f
// ) data from a client to a print server.
//
// * Notifications to the client of changes in the print server's print system.
//
// Server processing instructions are specified by the parameters that are used in the
// protocol methods. These parameters include:
//
// * Printer driver ( a98de8d9-8e9e-4cd7-92f7-f9269206a41e#gt_1a48eebd-e72c-494d-b8cb-84dfb7bc3b65
// ) configuration information.
//
// * The spool file format ( a98de8d9-8e9e-4cd7-92f7-f9269206a41e#gt_4b720910-6e5b-495d-b1ce-19087ae36922
// ) for the print data that is sent by the client.
//
// * The access level ( a98de8d9-8e9e-4cd7-92f7-f9269206a41e#gt_54c8a9d7-fa94-444b-b8d8-f6485bad223f
// ) of the connection.
//
// * The target print queue ( a98de8d9-8e9e-4cd7-92f7-f9269206a41e#gt_569f1f1c-f426-46fa-91d2-3d1eb0b19aa1
// ) name for name-based methods.
//
// * A handle to the target print queue for handle-based methods.
//
// Status information is communicated back to the client in the return codes from calls
// that are made to the print server.
//
// The following sections give an overview of these functions.
package par

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
	GoPackage = "par"
)
