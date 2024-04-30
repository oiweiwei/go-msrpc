// The mqmp package implements the MQMP client protocol.
//
// # Introduction
//
// The Message Queuing (MSMQ): Queue Manager Client Protocol is an RPC-based protocol,
// which enables communication between an application and an MSMQ supporting server
// or a remote MSMQ queue manager. Operations that an MSMQ application performs using
// this protocol include:
//
// * Managing private queues ( 102555d1-0dbf-4b2e-b78c-e388823d252c#gt_1a7f2b81-359e-4d2f-b4b4-f30bc7dd0d8f
// ) that are local queues ( 102555d1-0dbf-4b2e-b78c-e388823d252c#gt_a78ebed7-0657-43fa-bcc8-489aa88ba33f
// ).
//
// * Opening and closing local queue handles and remote queue ( 102555d1-0dbf-4b2e-b78c-e388823d252c#gt_91d29587-3e24-439b-8f2c-c2a43be99afc
// ) handles.
//
// * Enlisting, committing, and aborting internal transactions ( 102555d1-0dbf-4b2e-b78c-e388823d252c#gt_fcccf89d-d9c6-44ed-9f8a-13a204fe35b3
// ).
//
// * Enlisting the queue manager ( 102555d1-0dbf-4b2e-b78c-e388823d252c#gt_476f10ed-08f0-4887-b583-59d5cf909979
// ) in external transactions ( 102555d1-0dbf-4b2e-b78c-e388823d252c#gt_04fa7698-4cbc-4e38-bcc3-58135b87cbe0
// ).
//
// * Purging queues ( 102555d1-0dbf-4b2e-b78c-e388823d252c#gt_c1a6400d-703b-4f9a-a74c-40f1487978d9
// ).
//
// * Creating cursors ( 102555d1-0dbf-4b2e-b78c-e388823d252c#gt_aa5e9c2d-16c1-4301-8bfe-18a0913ed275
// ) for local queues and remote queues.
//
// * Sending messages ( 102555d1-0dbf-4b2e-b78c-e388823d252c#gt_85c78cf0-1fb6-4e5d-85f5-a2e9f58a6b9e
// ).
//
// * Reading messages.
//
// # Overview
//
// This protocol provides a means for applications to communicate with a supporting
// server. An MSMQ application uses this protocol to perform basic message queuing operations
// on a supporting server, such as creating queues, altering queue properties, sending
// messages, and receiving messages. An MSMQ application also uses this protocol to
// communicate with a remote MSMQ queue manager to open and close remote queues.
package mqmp

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
	GoPackage = "mqmp"
)
