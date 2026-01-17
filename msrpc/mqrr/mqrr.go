// The mqrr package implements the MQRR client protocol.
//
// # Introduction
//
// This document specifies the Message Queuing (MSMQ): Queue Manager Remote Read Protocol,
// a remote procedure call (RPC)-based protocol that is used by Microsoft Message Queuing
// (MSMQ) clients to read or reject a message from a queue, to move a message between
// queues, and to purge all messages from a queue.
//
// # Overview
//
// Microsoft Message Queuing (MSMQ) is a communications service that provides asynchronous
// and reliable message passing between client applications running on different hosts.
// In MSMQ, clients send application messages to a queue and/or consume application
// messages from a queue. The queue provides persistence of the messages, enabling them
// to survive across application restarts and allowing the sending and receiving client
// applications to send and receive messages asynchronously from each other.
//
// Queues are typically hosted by a communications service called a queue manager. By
// hosting the queue manager in a separate service from the client applications, applications
// can communicate even if they never execute at the same time by exchanging messages
// via a queue hosted by the queue manager.
//
// The queue manager can execute on a different node than the client applications. When
// this scenario occurs, a protocol is required to insert messages into the queue, and
// another protocol is needed to consume messages from the queue. The Message Queuing
// (MSMQ): Queue Manager Remote Read Protocol provides a protocol for consuming messages
// from a remote queue.
package mqrr

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
	GoPackage = "mqrr"
)
