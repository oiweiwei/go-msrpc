// The rprn package implements the RPRN client protocol.
//
// # Introduction
//
// The Print System Remote Protocol is a synchronous Remote Procedure Call (RPC)-based
// protocol used by a print client to send print jobs to a print server, to direct their
// processing, to add or remove print queues, and to perform other print system management
// functions.
//
// An enhanced replacement for this protocol is specified in [MS-PAR] The Print System
// Asynchronous Remote Protocol. It provides asynchronous forms of the synchronous operations
// supported by the Print System Remote Protocol and extends them with additional asynchronous
// operations. It also provides for higher-level authentication in RPC calls between
// client and server.
//
// # Overview
//
// The Print System Remote Protocol (RPRN) is based on the Remote Procedure Call (RPC)
// protocol, as defined in in [C706] and [MS-RPCE]. It supports synchronous printing
// and spooling operations between a client and a server and includes print job control
// and print system management.
//
// An enhanced replacement for this protocol is defined in [MS-PAR] The Print System
// Asynchronous Remote Protocol. It provides asynchronous forms of the synchronous operations
// supported by the Print System Remote Protocol and extends them with additional asynchronous
// operations. It also provides for higher-level authentication in RPC calls between
// client and server (see [MS-PAR] sections 3.1.3 and 3.2.3.
//
// The Print System Remote Protocol provides the following functions:
//
// * Management of the print system ( 831cd729-be7c-451e-b729-bd8d84ce4d24#gt_b49fcfa1-e15c-4dce-be33-d74d5bfda785
// ) of a print server ( 831cd729-be7c-451e-b729-bd8d84ce4d24#gt_59fb3ddc-63cf-45df-8a90-46a6af9e00cb
// ) from a client.
//
// * Communication of print job ( 831cd729-be7c-451e-b729-bd8d84ce4d24#gt_12a6e569-e97c-4761-92f0-e397f8d5125f
// ) data from a client to a print server.
//
// * Notifications to the client of changes in the print server's print system.
//
// Server processing instructions are specified by the parameters that are used in the
// protocol methods. These parameters include:
//
// * Printer driver ( 831cd729-be7c-451e-b729-bd8d84ce4d24#gt_1a48eebd-e72c-494d-b8cb-84dfb7bc3b65
// ) configuration information.
//
// * The spool file ( 831cd729-be7c-451e-b729-bd8d84ce4d24#gt_4c0e7da9-88a8-4430-abd7-27e7250b8180
// ) format for the print data that is sent by the client.
//
// * The access level ( 831cd729-be7c-451e-b729-bd8d84ce4d24#gt_54c8a9d7-fa94-444b-b8d8-f6485bad223f
// ) of the connection.
//
// * The target print queue ( 831cd729-be7c-451e-b729-bd8d84ce4d24#gt_569f1f1c-f426-46fa-91d2-3d1eb0b19aa1
// ) name for name-based methods.
//
// * A handle to the target print queue for handle-based methods.
//
// Status information is communicated back to the client in the return codes from calls
// that are made to the print server.
//
// The following sections give an overview of these functions.
package rprn

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
	GoPackage = "rprn"
)
