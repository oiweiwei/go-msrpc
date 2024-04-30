// The sch package implements the SCH client protocol.
//
// # Introduction
//
// This specification describes how the Task Scheduler Remoting Protocol is used to
// register and configure a task and to inquire about the status of running tasks on
// a remote machine.
//
// # Overview
//
// The Task Scheduler Remoting Protocol is used to register and configure tasks or to
// query the status of running tasks on a remote server. The Task Scheduler Remoting
// Protocol primarily consists of three separate remote procedure call (RPC) interfaces:
//
// * Net Schedule ( ATSvc ( 4d44c426-fad2-4cc7-9677-bfcd235dca33 ) )
//
// * Task Scheduler Agent ( SASec ( 7849c5ca-a8df-4c1d-8565-41a9b979a63d ) )
//
// * Windows Vista operating system Task Remote Protocol ( ITaskSchedulerService ( eb12c947-7e20-4a30-a528-85bc433cec44
// ) )
//
// All three interfaces use RPC as their transport to configure and manage tasks remotely.
//
// The three interfaces represent a continuum of increasing functionality, with ATSvc
// providing rudimentary functionality and ITaskSchedulerService providing the most
// functionality. Historically, the ATSvc interface is the oldest.<1> The three interfaces
// are not independent—they operate on the task store, shared persistent storage for
// tasks.
//
// In the ATSvc interface (see section 3.2.5.2), a task can be anything that can be
// specified on a command line for execution on the server. The client can specify execution
// at a given time or repeated execution on particular days of the week or month. In
// addition to creating tasks with NetrJobAdd (section 3.2.5.2.1), the interface includes
// methods for deleting tasks (section 3.2.5.2.2), enumerating tasks (section 3.2.5.2.3),
// and querying the status of a task (section 3.2.5.2.4).
//
// The SASec interface (section 3.2.5.3), only includes methods for manipulating account
// information, because most SASec-created task configuration is stored in the file
// system using the .JOB file format (section 2.4). Clients add, delete, enumerate,
// and query tasks using a remote file system protocol as specified in section 3.2.5.
// The .JOB file format provides more features than the ATSvc interface for specifying
// tasks.<2>
//
// Clients use the Windows Remote Registry Protocol Specification, as specified in [MS-RRP],
// to discover the path of the remote directory that stores the tasks as .JOB files
// (section 3.2.5.3). Clients use the SASec interface methods to supply security credentials
// for the remote task execution.
//
// In contrast, the ITaskSchedulerService interface (section 3.2.5.4) includes methods
// for creating, deleting, enumerating, and querying tasks. The remote registry and
// file system protocols are not used. The ITaskSchedulerService interface uses XML
// to specify task configuration (section 2.5). The XML schema provides more features
// than the .JOB File Format for specifying tasks.
package sch

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
	GoPackage = "sch"
)
