// The wdsc package implements the WDSC client protocol.
//
// # Introduction
//
// The Windows Deployment Services (WDS) Control Protocol specifies an RPC interface
// that provides the ability to remotely invoke services provided by WDS Server. It
// is a client/server protocol which uses RPC as a transport. The protocol provides
// a generic invocation mechanism to send requests to a server and receive replies.
//
// # Overview
//
// Windows Deployment Services (WDS) Control Protocol is a generic client/server protocol
// which is used to invoke services provided by Service Providers in WDS Server. The
// WDS Control Protocol is a simple protocol with no state shared across multiple calls.
// Each call is considered one complete request.
//
// A typical service invocation involves the following:
//
// *
//
// The client has already obtained the name of the WDS Server ( c9ecc8ee-2046-4ef2-b2ec-329fde7f8b01#gt_8ed48787-eb06-4b86-8b2e-19347bfbf07b
// ) , Endpoint GUID ( c9ecc8ee-2046-4ef2-b2ec-329fde7f8b01#gt_3996e5d9-beae-47cc-bf9d-dd0c570fbff5
// ) for Service Provider and OpCode ( c9ecc8ee-2046-4ef2-b2ec-329fde7f8b01#gt_1b22a3a6-f2e9-438d-9b61-274ec834a114
// ) for the operation being invoked.
//
// *
//
// The client constructs a request by packaging required variables (as specified in
// section 2.2.1 ( 787da555-e0af-4e81-a8a8-619368b072ff ) ), Endpoint GUID and OpCode.
//
// *
//
// The WDS Control Protocol sends the request to the server by using RPC ( c9ecc8ee-2046-4ef2-b2ec-329fde7f8b01#gt_8a7f6700-8311-45bc-af10-82e10accd331
// ) interface (as specified in section 3 ( 169ee4aa-8d4a-4fda-9463-9f321f81e88f ) ).
//
// *
//
// The WDS Server dispatches the request to the appropriate Service Provider based on
// Endpoint GUID.
//
// *
//
// Based on the Endpoint GUID and OpCode in the request, Service Provider will:
//
// *
//
// Validate that the client has appropriate rights to perform the operation.
//
// *
//
// Unpack the variables stored in the packet.
//
// *
//
// Perform the requested operation.
//
// *
//
// Package the results in pre-determined variables and complete the RPC request.
//
// *
//
// The client will check for success or failure of the request (as specified in section
// 3.1.4.2 ( 0926d80d-f68c-4df3-bab3-2d4741901367 ) and 3.2 ( 49d0a2b0-5cc2-4204-9434-1de6562c428a
// ) ).
//
// *
//
// Unpack the variables from the reply packet and process the results.
//
// The following diagram shows a client making a request to the WDS Server:
package wdsc

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
	GoPackage = "wdsc"
)
