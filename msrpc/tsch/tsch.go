// The tsch package implements the TSCH client protocol.
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
package tsch

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
	GoPackage = "tsch"
)

// ATEnum structure represents AT_ENUM RPC structure.
//
// The format of the AT_ENUM structure is almost identical to the AT_INFO structure
// with only the addition of the JobId field:
type ATEnum struct {
	// JobId:  Identifier of the task. This is a 32-bit integer that uniquely identifies
	// the AT job in the server system.
	//
	// All other fields are interpreted according to the corresponding specification given
	// in AT_INFO (section 2.3.4) for all other fields.
	JobID       uint32 `idl:"name:JobId" json:"job_id"`
	JobTime     uint64 `idl:"name:JobTime" json:"job_time"`
	DaysOfMonth uint32 `idl:"name:DaysOfMonth" json:"days_of_month"`
	DaysOfWeek  uint8  `idl:"name:DaysOfWeek" json:"days_of_week"`
	Flags       uint8  `idl:"name:Flags" json:"flags"`
	Command     string `idl:"name:Command" json:"command"`
}

func (o *ATEnum) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ATEnum) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.JobID); err != nil {
		return err
	}
	if err := w.WriteData(ndr.Uint3264(o.JobTime)); err != nil {
		return err
	}
	if err := w.WriteData(o.DaysOfMonth); err != nil {
		return err
	}
	if err := w.WriteData(o.DaysOfWeek); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.Command != "" {
		_ptr_Command := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.Command); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Command, _ptr_Command); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ATEnum) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.JobID); err != nil {
		return err
	}
	if err := w.ReadData((*ndr.Uint3264)(&o.JobTime)); err != nil {
		return err
	}
	if err := w.ReadData(&o.DaysOfMonth); err != nil {
		return err
	}
	if err := w.ReadData(&o.DaysOfWeek); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_Command := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.Command); err != nil {
			return err
		}
		return nil
	})
	_s_Command := func(ptr interface{}) { o.Command = *ptr.(*string) }
	if err := w.ReadPointer(&o.Command, _s_Command, _ptr_Command); err != nil {
		return err
	}
	return nil
}

// ATInfo structure represents AT_INFO RPC structure.
//
// The client uses the AT_INFO structure to configure a task in the ATSvc NetrJobAdd
// (section 3.2.5.2.1) method. The server returns the AT_INFO structure with information
// about a task in the ATSvc NetrJobGetInfo (section 3.2.5.2.4) method. The format of
// the AT_INFO structure is as follows.
type ATInfo struct {
	// JobTime:  This field is the time of day the task is to run, expressed as milliseconds
	// after midnight. The value is in the range of 0 to 86399999 (24*60*60*1000-1). JobTime
	// is present on the wire as a 32-bit unsigned integer.<5>
	JobTime uint64 `idl:"name:JobTime" json:"job_time"`
	// DaysOfMonth:  Contains individual bit flags that specify that the task is to run
	// on that day of the month. Bits that do not correspond to a real calendar day (for
	// example, bit 2 in February, which means the 30th) are ignored. See also the JOB_ADD_CURRENT_DATE
	// flag below. The mapping of bit to day of month is as follows.
	//
	//	+---+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+---+---+---+---+---+---+---+---+---+
	//	| 0 |  1  |  2  |  3  |  4  |  5  |  6  |  7  |  8  |  9  |  1  |  1  |  2  |  3  |  4  |  5  |  6  |  7  |  8  |  9  |  2  |  1  |  2  | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
	//	|   |     |     |     |     |     |     |     |     |     |  0  |     |     |     |     |     |     |     |     |     |  0  |     |     |   |   |   |   |   |   |   | 0 |   |
	//	+---+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+---+---+---+---+---+---+---+---+---+
	//	+---+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+---+---+---+---+---+---+---+---+---+
	//	| x | 3 1 | 3 0 | 2 9 | 2 8 | 2 7 | 2 6 | 2 5 | 2 4 | 2 3 | 2 2 | 2 1 | 2 0 | 1 9 | 1 8 | 1 7 | 1 6 | 1 5 | 1 4 | 1 3 | 1 2 | 1 1 | 1 0 | 9 | 8 | 7 | 6 | 5 | 4 | 3 | 2 | 1 |
	//	+---+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+---+---+---+---+---+---+---+---+---+
	//
	// 1-31: Any bit set to 1 specifies that the task can be run on that day of the month.
	// More than one bit can be set to 1.
	//
	// x: Unused. Can be set to zero when sent and ignored on receipt.<6>
	DaysOfMonth uint32 `idl:"name:DaysOfMonth" json:"days_of_month"`
	// DaysOfWeek:  Contains individual bit flags that specify the day of the week on which
	// the task is to run. The mapping of bit to day of week is as follows.
	//
	//	+---+-----+-----+-----+-----+-----+-----+-----+
	//	|   |     |     |     |     |     |     |     |
	//	| 0 |  1  |  2  |  3  |  4  |  5  |  6  |  7  |
	//	|   |     |     |     |     |     |     |     |
	//	+---+-----+-----+-----+-----+-----+-----+-----+
	//	+---+-----+-----+-----+-----+-----+-----+-----+
	//	| X | S U | S A | F R | T H | W E | T U | M O |
	//	+---+-----+-----+-----+-----+-----+-----+-----+
	//
	//
	//	+--------------+-------------------------------------------------------------------+
	//	|              |                                                                   |
	//	|    VALUE     |                            DESCRIPTION                            |
	//	|              |                                                                   |
	//	+--------------+-------------------------------------------------------------------+
	//	+--------------+-------------------------------------------------------------------+
	//	| X            | Unused. Can be set to zero when sent and ignored upon receipt.<7> |
	//	+--------------+-------------------------------------------------------------------+
	//	| MO Monday    | If set to 1, specifies that the task can be run on Monday.        |
	//	+--------------+-------------------------------------------------------------------+
	//	| TU Tuesday   | If set to 1, specifies that the task can be run on Tuesday.       |
	//	+--------------+-------------------------------------------------------------------+
	//	| WE Wednesday | If set to 1, specifies that the task can be run on Wednesday.     |
	//	+--------------+-------------------------------------------------------------------+
	//	| TH Thursday  | If set to 1, specifies that the task can be run on Thursday.      |
	//	+--------------+-------------------------------------------------------------------+
	//	| FR Friday    | If set to 1, specifies that the task can be run on Friday.        |
	//	+--------------+-------------------------------------------------------------------+
	//	| SA Saturday  | If set to 1, indicates that the task can be run on Saturday.      |
	//	+--------------+-------------------------------------------------------------------+
	//	| SU Sunday    | If set to 1, indicates that the task can be run on Sunday.        |
	//	+--------------+-------------------------------------------------------------------+
	DaysOfWeek uint8 `idl:"name:DaysOfWeek" json:"days_of_week"`
	// Flags:  Contains individual bit flags set to zero or more of the following values.
	//
	//	+---+---+---+-----+-----+-----+-----+-----+
	//	|   |   |   |     |     |     |     |     |
	//	| 0 | 1 | 2 |  3  |  4  |  5  |  6  |  7  |
	//	|   |   |   |     |     |     |     |     |
	//	+---+---+---+-----+-----+-----+-----+-----+
	//	+---+---+---+-----+-----+-----+-----+-----+
	//	| X | 0 | 0 | N I | A C | R T | E R | R P |
	//	+---+---+---+-----+-----+-----+-----+-----+
	//
	//
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	|                         |                                                                                  |
	//	|          VALUE          |                                   DESCRIPTION                                    |
	//	|                         |                                                                                  |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| X                       | Unused. MUST be set to zero when sent and MUST be ignored on receipt.<8>         |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| RP JOB_RUN_PERIODICALLY | If set to 1, specifies that the task has a repeating schedule. See Global Timer  |
	//	|                         | (section 3.2.6.1).                                                               |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| ER JOB_EXEC_ERROR       | If an error was encountered during the last time this task tried to execute a    |
	//	|                         | program, this bit is set to 1.                                                   |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| RT JOB_RUNS_TODAY       | If set to 1, indicates that this task is to run today.                           |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| AC JOB_ADD_CURRENT_DATE | If set to 1, indicates that the bit corresponding to the current day (today) bit |
	//	|                         | is added to the DaysOfMonth bit mask.                                            |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| NI JOB_NONINTERACTIVE   | If set to 1, specifies that this task is not to be interactive; that is, it will |
	//	|                         | not interact with the current logged-on user.                                    |
	//	+-------------------------+----------------------------------------------------------------------------------+
	Flags uint8 `idl:"name:Flags" json:"flags"`
	// Command:  This member is a Unicode string that contains the name of a command, batch
	// program, or binary file to execute; or the name of a document to be executed by its
	// associated executable.
	//
	// Note  The DaysOfMonth and DayOfWeek bits can be used simultaneously. For example,
	// setting DaysOfMonth to the fifteenth of the month, DayOfWeek to Tuesday, and JOB_RUN_PERIODICALLY
	// to 1 will cause the job to run on the fifteenth of every month and on Tuesday of
	// every week.
	Command string `idl:"name:Command;string" json:"command"`
}

func (o *ATInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ATInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(ndr.Uint3264(o.JobTime)); err != nil {
		return err
	}
	if err := w.WriteData(o.DaysOfMonth); err != nil {
		return err
	}
	if err := w.WriteData(o.DaysOfWeek); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.Command != "" {
		_ptr_Command := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Command); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Command, _ptr_Command); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ATInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData((*ndr.Uint3264)(&o.JobTime)); err != nil {
		return err
	}
	if err := w.ReadData(&o.DaysOfMonth); err != nil {
		return err
	}
	if err := w.ReadData(&o.DaysOfWeek); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_Command := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Command); err != nil {
			return err
		}
		return nil
	})
	_s_Command := func(ptr interface{}) { o.Command = *ptr.(*string) }
	if err := w.ReadPointer(&o.Command, _s_Command, _ptr_Command); err != nil {
		return err
	}
	return nil
}
