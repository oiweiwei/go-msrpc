// The mqqp package implements the MQQP client protocol.
//
// # Introduction
//
// This document specifies the Message Queuing (MSMQ): Queue Manager to Queue Manager
// Protocol. The Queue Manager to Queue Manager Protocol is an RPC-based protocol used
// by the queue manager and runtime library to read and purge messages from a remote
// queue.
//
// # Overview
//
// Message queuing is a communications service that provides asynchronous and reliable
// message passing between client applications, including those client applications
// running on different hosts. In message queuing, clients send messages to a queue
// and consume application messages from a queue. The queue provides persistence of
// the messages, enabling them to survive across application restarts, and allowing
// the sending and receiving client applications to operate asynchronously from each
// other.
//
// Queues are typically hosted by a communications service called a queue manager. By
// hosting the queue manager in a separate service from the client applications, applications
// can communicate by exchanging messages via a queue hosted by the queue manager, even
// if the client applications never execute at the same time.
//
// The queue manager can perform operations on a remote queue. When this scenario occurs,
// a protocol is required to insert messages into the remote queue, and another protocol
// is required to consume messages from the remote queue. The Message Queuing (MSMQ):
// Queue Manager to Queue Manager Protocol provides a protocol for consuming messages
// from a remote queue.
//
// The Queue Manager to Queue Manager Protocol is used only to read messages from a
// queue or to purge messages from the queue. Reading a message also implies deleting
// the message after it is read, as specified in Queue Operations (section 1.3.3).
package mqqp

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
	GoPackage = "mqqp"
)
