package itaskschedulerservice

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
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
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "tsch"
)

var (
	// Syntax UUID
	TaskSchedulerServiceSyntaxUUID = &uuid.UUID{TimeLow: 0x86d35949, TimeMid: 0x83c9, TimeHiAndVersion: 0x4044, ClockSeqHiAndReserved: 0xb4, ClockSeqLow: 0x24, Node: [6]uint8{0xdb, 0x36, 0x32, 0x31, 0xfd, 0xc}}
	// Syntax ID
	TaskSchedulerServiceSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: TaskSchedulerServiceSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// ITaskSchedulerService interface.
type TaskSchedulerServiceClient interface {

	// The SchRpcHighestVersion method MUST return the highest version of the Task Scheduler
	// Remoting Protocol that is supported by the server.<65>
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	HighestVersion(context.Context, *HighestVersionRequest, ...dcerpc.CallOption) (*HighestVersionResponse, error)

	// The SchRpcRegisterTask method MUST register a task with the server.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in section 2 of [MS-ERREF].
	RegisterTask(context.Context, *RegisterTaskRequest, ...dcerpc.CallOption) (*RegisterTaskResponse, error)

	// The SchRpcRetrieveTask method MUST return a task definition.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	RetrieveTask(context.Context, *RetrieveTaskRequest, ...dcerpc.CallOption) (*RetrieveTaskResponse, error)

	// The SchRpcCreateFolder method creates a new folder.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	CreateFolder(context.Context, *CreateFolderRequest, ...dcerpc.CallOption) (*CreateFolderResponse, error)

	// The SchRpcSetSecurity method MUST set a security descriptor on a task or folder.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	SetSecurity(context.Context, *SetSecurityRequest, ...dcerpc.CallOption) (*SetSecurityResponse, error)

	// The SchRpcGetSecurity method MUST get the security descriptor associated with a task
	// or folder.
	//
	// HRESULT SchRpcGetSecurity( [in, string] const wchar_t* path,
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	GetSecurity(context.Context, *GetSecurityRequest, ...dcerpc.CallOption) (*GetSecurityResponse, error)

	// The SchRpcEnumFolders method MUST retrieve a list of folders on the server.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	EnumFolders(context.Context, *EnumFoldersRequest, ...dcerpc.CallOption) (*EnumFoldersResponse, error)

	// The SchRpcEnumTasks method MUST return the list of tasks in a specific folder.
	//
	// MUST returns S_FALSE if there are more tasks to enumerate.
	//
	// Note that if client requests items 1-10 and then 11-20, the second call can return
	// duplicate entries if the task list has changed in between calls.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	EnumTasks(context.Context, *EnumTasksRequest, ...dcerpc.CallOption) (*EnumTasksResponse, error)

	// The SchRpcEnumInstances method MUST return a task's list of instances that are currently
	// running.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	EnumInstances(context.Context, *EnumInstancesRequest, ...dcerpc.CallOption) (*EnumInstancesResponse, error)

	// The SchRpcGetInstanceInfo method MUST get information about an instance of a running
	// task.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	GetInstanceInfo(context.Context, *GetInstanceInfoRequest, ...dcerpc.CallOption) (*GetInstanceInfoResponse, error)

	// The SchRpcStopInstance method MUST stop a specified instance of a task.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	StopInstance(context.Context, *StopInstanceRequest, ...dcerpc.CallOption) (*StopInstanceResponse, error)

	// The SchRpcStop MUST stop all currently running instances of a task specified by a
	// path.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	Stop(context.Context, *StopRequest, ...dcerpc.CallOption) (*StopResponse, error)

	// The SchRpcRun method MUST run a task specified by a path.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	Run(context.Context, *RunRequest, ...dcerpc.CallOption) (*RunResponse, error)

	// The SchRpcDelete method MUST delete a task or folder in the task store.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	Delete(context.Context, *DeleteRequest, ...dcerpc.CallOption) (*DeleteResponse, error)

	// The SchRpcRename method renames a folder or task.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	Rename(context.Context, *RenameRequest, ...dcerpc.CallOption) (*RenameResponse, error)

	// The SchRpcScheduledRuntimes method MUST return scheduled run times.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	ScheduledRuntimes(context.Context, *ScheduledRuntimesRequest, ...dcerpc.CallOption) (*ScheduledRuntimesResponse, error)

	// The SchRpcGetLastRunInfo method MUST return information about the task's last run.
	//
	// Return Values: For more information about return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	GetLastRunInfo(context.Context, *GetLastRunInfoRequest, ...dcerpc.CallOption) (*GetLastRunInfoResponse, error)

	// The SchRpcGetTaskInfo method MUST return information about a specified task.
	//
	// Return Values: For more information about return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	GetTaskInfo(context.Context, *GetTaskInfoRequest, ...dcerpc.CallOption) (*GetTaskInfoResponse, error)

	// The SchRpcGetNumberOfMissedRuns MUST return the number of times a task was scheduled
	// to run but did not due to the server being unavailable (for example, powered off).
	//
	// Return Values: For more information about return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	GetNumberOfMissedRuns(context.Context, *GetNumberOfMissedRunsRequest, ...dcerpc.CallOption) (*GetNumberOfMissedRunsResponse, error)

	// The SchRpcEnableTask method MUST enable or disable a task.
	//
	// Return Values: For more information about return codes, see [MS-ERREF] section 2.
	EnableTask(context.Context, *EnableTaskRequest, ...dcerpc.CallOption) (*EnableTaskResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// TaskUserCred structure represents TASK_USER_CRED RPC structure.
//
// The TASK_USER_CRED structure contains user credentials and is passed to the server
// during task registration, as specified in section 3.2.5.4.2.
type TaskUserCred struct {
	// userId:  Contains the user name identifying the account under which the task can
	// execute. The user name is a string recognized by Windows authentication, as specified
	// in [MS-AUTHSOD] section 1.1.1.2.<9>
	UserID string `idl:"name:userId;string" json:"user_id"`
	// password:  Contains the password associated with the user specified in the preceding
	// userId field above, represented as a string. For information about security considerations,
	// see section 5.1.
	Password string `idl:"name:password;string" json:"password"`
	// flags:  The flags field contains individual bit flags that are structured as shown
	// in the following table. The client MUST set the undefined bits to 0. The undefined
	// bits are ignored on receipt.
	//
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
	//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | D |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//
	//
	//	+-------------------+----------------------------------------------------------------------------------+
	//	|                   |                                                                                  |
	//	|       VALUE       |                                   DESCRIPTION                                    |
	//	|                   |                                                                                  |
	//	+-------------------+----------------------------------------------------------------------------------+
	//	+-------------------+----------------------------------------------------------------------------------+
	//	| D credFlagDefault | If set to 1, the server can use these credentials if no other credentials are    |
	//	|                   | provided. If set to 0, the server can use these credentials instead of the       |
	//	|                   | credentials specified in the task configuration.                                 |
	//	+-------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:flags" json:"flags"`
}

func (o *TaskUserCred) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TaskUserCred) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.UserID != "" {
		_ptr_userId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.UserID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.UserID, _ptr_userId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Password != "" {
		_ptr_password := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Password); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Password, _ptr_password); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	return nil
}
func (o *TaskUserCred) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_userId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.UserID); err != nil {
			return err
		}
		return nil
	})
	_s_userId := func(ptr interface{}) { o.UserID = *ptr.(*string) }
	if err := w.ReadPointer(&o.UserID, _s_userId, _ptr_userId); err != nil {
		return err
	}
	_ptr_password := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Password); err != nil {
			return err
		}
		return nil
	})
	_s_password := func(ptr interface{}) { o.Password = *ptr.(*string) }
	if err := w.ReadPointer(&o.Password, _s_password, _ptr_password); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	return nil
}

// TaskXMLErrorInfo structure represents TASK_XML_ERROR_INFO RPC structure.
//
// The TASK_XML_ERROR_INFO structure is a return value from the SchRpcRegisterTask (Opnum
// 1) method that contains additional information about an error in the XML task definition.
type TaskXMLErrorInfo struct {
	// line:  Contains the line number where parsing failed. Carriage returns in the XML
	// separate the string into lines. The first line is "line one".
	Line uint32 `idl:"name:line" json:"line"`
	// column:  Contains the column where parsing failed. Indicates the character within
	// the line, where the first character is "column one".
	Column uint32 `idl:"name:column" json:"column"`
	// node:  Contains the attribute or element name that caused the error, or which is
	// missing.
	Node string `idl:"name:node;string" json:"node"`
	// value:  When this field is not empty, it contains the invalid value that caused the
	// error.
	Value string `idl:"name:value;string" json:"value"`
}

func (o *TaskXMLErrorInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TaskXMLErrorInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Line); err != nil {
		return err
	}
	if err := w.WriteData(o.Column); err != nil {
		return err
	}
	if o.Node != "" {
		_ptr_node := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Node); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Node, _ptr_node); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Value != "" {
		_ptr_value := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Value); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Value, _ptr_value); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *TaskXMLErrorInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Line); err != nil {
		return err
	}
	if err := w.ReadData(&o.Column); err != nil {
		return err
	}
	_ptr_node := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Node); err != nil {
			return err
		}
		return nil
	})
	_s_node := func(ptr interface{}) { o.Node = *ptr.(*string) }
	if err := w.ReadPointer(&o.Node, _s_node, _ptr_node); err != nil {
		return err
	}
	_ptr_value := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Value); err != nil {
			return err
		}
		return nil
	})
	_s_value := func(ptr interface{}) { o.Value = *ptr.(*string) }
	if err := w.ReadPointer(&o.Value, _s_value, _ptr_value); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultTaskSchedulerServiceClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultTaskSchedulerServiceClient) HighestVersion(ctx context.Context, in *HighestVersionRequest, opts ...dcerpc.CallOption) (*HighestVersionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &HighestVersionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTaskSchedulerServiceClient) RegisterTask(ctx context.Context, in *RegisterTaskRequest, opts ...dcerpc.CallOption) (*RegisterTaskResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RegisterTaskResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTaskSchedulerServiceClient) RetrieveTask(ctx context.Context, in *RetrieveTaskRequest, opts ...dcerpc.CallOption) (*RetrieveTaskResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RetrieveTaskResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTaskSchedulerServiceClient) CreateFolder(ctx context.Context, in *CreateFolderRequest, opts ...dcerpc.CallOption) (*CreateFolderResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateFolderResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTaskSchedulerServiceClient) SetSecurity(ctx context.Context, in *SetSecurityRequest, opts ...dcerpc.CallOption) (*SetSecurityResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetSecurityResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTaskSchedulerServiceClient) GetSecurity(ctx context.Context, in *GetSecurityRequest, opts ...dcerpc.CallOption) (*GetSecurityResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetSecurityResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTaskSchedulerServiceClient) EnumFolders(ctx context.Context, in *EnumFoldersRequest, opts ...dcerpc.CallOption) (*EnumFoldersResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumFoldersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTaskSchedulerServiceClient) EnumTasks(ctx context.Context, in *EnumTasksRequest, opts ...dcerpc.CallOption) (*EnumTasksResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumTasksResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTaskSchedulerServiceClient) EnumInstances(ctx context.Context, in *EnumInstancesRequest, opts ...dcerpc.CallOption) (*EnumInstancesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumInstancesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTaskSchedulerServiceClient) GetInstanceInfo(ctx context.Context, in *GetInstanceInfoRequest, opts ...dcerpc.CallOption) (*GetInstanceInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetInstanceInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTaskSchedulerServiceClient) StopInstance(ctx context.Context, in *StopInstanceRequest, opts ...dcerpc.CallOption) (*StopInstanceResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &StopInstanceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTaskSchedulerServiceClient) Stop(ctx context.Context, in *StopRequest, opts ...dcerpc.CallOption) (*StopResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &StopResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTaskSchedulerServiceClient) Run(ctx context.Context, in *RunRequest, opts ...dcerpc.CallOption) (*RunResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RunResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTaskSchedulerServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...dcerpc.CallOption) (*DeleteResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DeleteResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTaskSchedulerServiceClient) Rename(ctx context.Context, in *RenameRequest, opts ...dcerpc.CallOption) (*RenameResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RenameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTaskSchedulerServiceClient) ScheduledRuntimes(ctx context.Context, in *ScheduledRuntimesRequest, opts ...dcerpc.CallOption) (*ScheduledRuntimesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ScheduledRuntimesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTaskSchedulerServiceClient) GetLastRunInfo(ctx context.Context, in *GetLastRunInfoRequest, opts ...dcerpc.CallOption) (*GetLastRunInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetLastRunInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTaskSchedulerServiceClient) GetTaskInfo(ctx context.Context, in *GetTaskInfoRequest, opts ...dcerpc.CallOption) (*GetTaskInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetTaskInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTaskSchedulerServiceClient) GetNumberOfMissedRuns(ctx context.Context, in *GetNumberOfMissedRunsRequest, opts ...dcerpc.CallOption) (*GetNumberOfMissedRunsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetNumberOfMissedRunsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTaskSchedulerServiceClient) EnableTask(ctx context.Context, in *EnableTaskRequest, opts ...dcerpc.CallOption) (*EnableTaskResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnableTaskResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTaskSchedulerServiceClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultTaskSchedulerServiceClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewTaskSchedulerServiceClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (TaskSchedulerServiceClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(TaskSchedulerServiceSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultTaskSchedulerServiceClient{cc: cc}, nil
}

// xxx_HighestVersionOperation structure represents the SchRpcHighestVersion operation
type xxx_HighestVersionOperation struct {
	Version uint32 `idl:"name:pVersion" json:"version"`
	Return  int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_HighestVersionOperation) OpNum() int { return 0 }

func (o *xxx_HighestVersionOperation) OpName() string {
	return "/ITaskSchedulerService/v1/SchRpcHighestVersion"
}

func (o *xxx_HighestVersionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_HighestVersionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_HighestVersionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_HighestVersionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_HighestVersionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Version); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_HighestVersionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Version); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// HighestVersionRequest structure represents the SchRpcHighestVersion operation request
type HighestVersionRequest struct {
}

func (o *HighestVersionRequest) xxx_ToOp(ctx context.Context, op *xxx_HighestVersionOperation) *xxx_HighestVersionOperation {
	if op == nil {
		op = &xxx_HighestVersionOperation{}
	}
	if o == nil {
		return op
	}
	return op
}

func (o *HighestVersionRequest) xxx_FromOp(ctx context.Context, op *xxx_HighestVersionOperation) {
	if o == nil {
		return
	}
}
func (o *HighestVersionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *HighestVersionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_HighestVersionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// HighestVersionResponse structure represents the SchRpcHighestVersion operation response
type HighestVersionResponse struct {
	// pVersion: The server MUST return the highest version of the Task Scheduler Remoting
	// Protocol that is supported by the server. This version MUST be in the format 0xMMMMmmmm,
	// where 'M' stands for the major version hex digits and 'm' stands for the minor version
	// hex digits. pVersion SHOULD be set to 0x00010002 or 0x00010003.
	//
	//	+------------+----------------------------------------+
	//	|            |                                        |
	//	|   VALUE    |                MEANING                 |
	//	|            |                                        |
	//	+------------+----------------------------------------+
	//	+------------+----------------------------------------+
	//	| 0x00010002 | Version 1.2 of the protocol is in use. |
	//	+------------+----------------------------------------+
	//	| 0x00010003 | Version 1.3 of the protocol is in use. |
	//	+------------+----------------------------------------+
	//	| 0x00010004 | Version 1.4 of the protocol is in use. |
	//	+------------+----------------------------------------+
	Version uint32 `idl:"name:pVersion" json:"version"`
	// Return: The SchRpcHighestVersion return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *HighestVersionResponse) xxx_ToOp(ctx context.Context, op *xxx_HighestVersionOperation) *xxx_HighestVersionOperation {
	if op == nil {
		op = &xxx_HighestVersionOperation{}
	}
	if o == nil {
		return op
	}
	o.Version = op.Version
	o.Return = op.Return
	return op
}

func (o *HighestVersionResponse) xxx_FromOp(ctx context.Context, op *xxx_HighestVersionOperation) {
	if o == nil {
		return
	}
	o.Version = op.Version
	o.Return = op.Return
}
func (o *HighestVersionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *HighestVersionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_HighestVersionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RegisterTaskOperation structure represents the SchRpcRegisterTask operation
type xxx_RegisterTaskOperation struct {
	Path       string            `idl:"name:path;string;pointer:unique" json:"path"`
	XML        string            `idl:"name:xml;string" json:"xml"`
	Flags      uint32            `idl:"name:flags" json:"flags"`
	SDDL       string            `idl:"name:sddl;string;pointer:unique" json:"sddl"`
	LogonType  uint32            `idl:"name:logonType" json:"logon_type"`
	CredsCount uint32            `idl:"name:cCreds" json:"creds_count"`
	Creds      []*TaskUserCred   `idl:"name:pCreds;size_is:(cCreds);pointer:unique" json:"creds"`
	ActualPath string            `idl:"name:pActualPath;string" json:"actual_path"`
	ErrorInfo  *TaskXMLErrorInfo `idl:"name:pErrorInfo" json:"error_info"`
	Return     int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_RegisterTaskOperation) OpNum() int { return 1 }

func (o *xxx_RegisterTaskOperation) OpName() string {
	return "/ITaskSchedulerService/v1/SchRpcRegisterTask"
}

func (o *xxx_RegisterTaskOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Creds != nil && o.CredsCount == 0 {
		o.CredsCount = uint32(len(o.Creds))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterTaskOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// path {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Path != "" {
			_ptr_path := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_path); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// xml {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.XML); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// sddl {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.SDDL != "" {
			_ptr_sddl := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.SDDL); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.SDDL, _ptr_sddl); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// logonType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.LogonType); err != nil {
			return err
		}
	}
	// cCreds {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.CredsCount); err != nil {
			return err
		}
	}
	// pCreds {in} (1:{pointer=unique}*(1))(2:{alias=TASK_USER_CRED}[dim:0,size_is=cCreds](struct))
	{
		if o.Creds != nil || o.CredsCount > 0 {
			_ptr_pCreds := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.CredsCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.Creds {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.Creds[i1] != nil {
						if err := o.Creds[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&TaskUserCred{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.Creds); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&TaskUserCred{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Creds, _ptr_pCreds); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterTaskOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// path {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_path := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
				return err
			}
			return nil
		})
		_s_path := func(ptr interface{}) { o.Path = *ptr.(*string) }
		if err := w.ReadPointer(&o.Path, _s_path, _ptr_path); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// xml {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.XML); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// sddl {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_sddl := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.SDDL); err != nil {
				return err
			}
			return nil
		})
		_s_sddl := func(ptr interface{}) { o.SDDL = *ptr.(*string) }
		if err := w.ReadPointer(&o.SDDL, _s_sddl, _ptr_sddl); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// logonType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.LogonType); err != nil {
			return err
		}
	}
	// cCreds {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.CredsCount); err != nil {
			return err
		}
	}
	// pCreds {in} (1:{pointer=unique}*(1))(2:{alias=TASK_USER_CRED}[dim:0,size_is=cCreds](struct))
	{
		_ptr_pCreds := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Creds", sizeInfo[0])
			}
			o.Creds = make([]*TaskUserCred, sizeInfo[0])
			for i1 := range o.Creds {
				i1 := i1
				if o.Creds[i1] == nil {
					o.Creds[i1] = &TaskUserCred{}
				}
				if err := o.Creds[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pCreds := func(ptr interface{}) { o.Creds = *ptr.(*[]*TaskUserCred) }
		if err := w.ReadPointer(&o.Creds, _s_pCreds, _ptr_pCreds); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterTaskOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterTaskOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pActualPath {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		if o.ActualPath != "" {
			_ptr_pActualPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ActualPath); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ActualPath, _ptr_pActualPath); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pErrorInfo {out} (1:{pointer=ref}*(2))(2:{alias=PTASK_XML_ERROR_INFO}*(1))(3:{alias=TASK_XML_ERROR_INFO}(struct))
	{
		if o.ErrorInfo != nil {
			_ptr_pErrorInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ErrorInfo != nil {
					if err := o.ErrorInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&TaskXMLErrorInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ErrorInfo, _ptr_pErrorInfo); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterTaskOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pActualPath {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		_ptr_pActualPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ActualPath); err != nil {
				return err
			}
			return nil
		})
		_s_pActualPath := func(ptr interface{}) { o.ActualPath = *ptr.(*string) }
		if err := w.ReadPointer(&o.ActualPath, _s_pActualPath, _ptr_pActualPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pErrorInfo {out} (1:{pointer=ref}*(2))(2:{alias=PTASK_XML_ERROR_INFO,pointer=ref}*(1))(3:{alias=TASK_XML_ERROR_INFO}(struct))
	{
		_ptr_pErrorInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ErrorInfo == nil {
				o.ErrorInfo = &TaskXMLErrorInfo{}
			}
			if err := o.ErrorInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pErrorInfo := func(ptr interface{}) { o.ErrorInfo = *ptr.(**TaskXMLErrorInfo) }
		if err := w.ReadPointer(&o.ErrorInfo, _s_pErrorInfo, _ptr_pErrorInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RegisterTaskRequest structure represents the SchRpcRegisterTask operation request
type RegisterTaskRequest struct {
	// path: MUST contain the full path associated (or to be associated) with a task as
	// specified in section 2.3.11. A NULL path indicates that the server SHOULD either
	// take the path from the task definition or generate a path itself.
	Path string `idl:"name:path;string;pointer:unique" json:"path"`
	// xml: MUST contain the task definition in XML format as specified in section 2.5.
	XML string `idl:"name:xml;string" json:"xml"`
	// flags: The flags field MUST contain individual bit flags that MUST have one or more
	// of the following values.
	//
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+-----+-----+-----+-----+-----+-----+
	//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 |  6  |  7  |  8  |  9  |  3  |  1  |
	//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |     |     |     |     |  0  |     |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+-----+-----+-----+-----+-----+-----+
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+-----+-----+-----+-----+-----+-----+
	//	| 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | I R | D A | D I | U P | C R | V O |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+-----+-----+-----+-----+-----+-----+
	//
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                                      |                                                                                  |
	//	|                VALUE                 |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| VO TASK_VALIDATE_ONLY                | If set to 1, the server MUST only validate the task definition and not actually  |
	//	|                                      | register the task.                                                               |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| CR TASK_CREATE                       | If set to 1, the server MUST register the task if the task does not already      |
	//	|                                      | exist.                                                                           |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| UP TASK_UPDATE                       | If set to 1, the server MUST update the task if the task already exists.         |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| DI TASK_DISABLE                      | If set to 1, the server MUST disable the task.                                   |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| DA TASK_DONT_ADD_PRINCIPAL_ACE       | If set to 1, the server MUST NOT add an 'Allow' access control entry (ACE) for   |
	//	|                                      | the task principal to the task's security descriptor.                            |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| IR TASK_IGNORE_REGISTRATION_TRIGGERS | If set to 1, the server MUST NOT start the task after registering it, even if it |
	//	|                                      | has registration triggers.                                                       |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:flags" json:"flags"`
	// sddl: MUST be a security descriptor in SDDL format as specified in [MS-DTYP]. A NULL
	// value MUST specify that the task inherits its security descriptor from its folder,
	// when creating a new task, or that the task's security descriptor does not change
	// when updating an existing task.
	SDDL string `idl:"name:sddl;string;pointer:unique" json:"sddl"`
	// logonType: MUST contain a TASK_LOGON_TYPE value (section 2.3.9). The server MUST
	// reject invalid values. The specified logonType SHOULD correspond to the type of principal
	// used in the xml task definition, or supplied in pCreds.
	LogonType uint32 `idl:"name:logonType" json:"logon_type"`
	// cCreds: MUST contain the number of credentials passed in pCreds parameter. Client
	// MUST set to 0 or 1, and the server MUST return an error if not 0 or 1.
	CredsCount uint32 `idl:"name:cCreds" json:"creds_count"`
	// pCreds: MUST be an array of credentials for the task with cCreds entries. pCreds
	// MUST be NULL if cCreds is 0. See section 2.3.8 for details of the TASK_USER_CRED
	// structure.
	Creds []*TaskUserCred `idl:"name:pCreds;size_is:(cCreds);pointer:unique" json:"creds"`
}

func (o *RegisterTaskRequest) xxx_ToOp(ctx context.Context, op *xxx_RegisterTaskOperation) *xxx_RegisterTaskOperation {
	if op == nil {
		op = &xxx_RegisterTaskOperation{}
	}
	if o == nil {
		return op
	}
	o.Path = op.Path
	o.XML = op.XML
	o.Flags = op.Flags
	o.SDDL = op.SDDL
	o.LogonType = op.LogonType
	o.CredsCount = op.CredsCount
	o.Creds = op.Creds
	return op
}

func (o *RegisterTaskRequest) xxx_FromOp(ctx context.Context, op *xxx_RegisterTaskOperation) {
	if o == nil {
		return
	}
	o.Path = op.Path
	o.XML = op.XML
	o.Flags = op.Flags
	o.SDDL = op.SDDL
	o.LogonType = op.LogonType
	o.CredsCount = op.CredsCount
	o.Creds = op.Creds
}
func (o *RegisterTaskRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RegisterTaskRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterTaskOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RegisterTaskResponse structure represents the SchRpcRegisterTask operation response
type RegisterTaskResponse struct {
	// pActualPath: If this parameter is non-NULL, the server MUST return the task's actual
	// path.
	ActualPath string `idl:"name:pActualPath;string" json:"actual_path"`
	// pErrorInfo: If this parameter is non-NULL and the XML task definition is invalid,
	// the server MUST return additional error information.
	ErrorInfo *TaskXMLErrorInfo `idl:"name:pErrorInfo" json:"error_info"`
	// Return: The SchRpcRegisterTask return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RegisterTaskResponse) xxx_ToOp(ctx context.Context, op *xxx_RegisterTaskOperation) *xxx_RegisterTaskOperation {
	if op == nil {
		op = &xxx_RegisterTaskOperation{}
	}
	if o == nil {
		return op
	}
	o.ActualPath = op.ActualPath
	o.ErrorInfo = op.ErrorInfo
	o.Return = op.Return
	return op
}

func (o *RegisterTaskResponse) xxx_FromOp(ctx context.Context, op *xxx_RegisterTaskOperation) {
	if o == nil {
		return
	}
	o.ActualPath = op.ActualPath
	o.ErrorInfo = op.ErrorInfo
	o.Return = op.Return
}
func (o *RegisterTaskResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RegisterTaskResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterTaskOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RetrieveTaskOperation structure represents the SchRpcRetrieveTask operation
type xxx_RetrieveTaskOperation struct {
	Path            string `idl:"name:path;string" json:"path"`
	LanguagesBuffer string `idl:"name:lpcwszLanguagesBuffer;string" json:"languages_buffer"`
	LanguagesLength uint32 `idl:"name:pulNumLanguages" json:"languages_length"`
	XML             string `idl:"name:pXml;string" json:"xml"`
	Return          int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_RetrieveTaskOperation) OpNum() int { return 2 }

func (o *xxx_RetrieveTaskOperation) OpName() string {
	return "/ITaskSchedulerService/v1/SchRpcRetrieveTask"
}

func (o *xxx_RetrieveTaskOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetrieveTaskOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// path {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
			return err
		}
	}
	// lpcwszLanguagesBuffer {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.LanguagesBuffer); err != nil {
			return err
		}
	}
	// pulNumLanguages {in} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.LanguagesLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetrieveTaskOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// path {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
			return err
		}
	}
	// lpcwszLanguagesBuffer {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.LanguagesBuffer); err != nil {
			return err
		}
	}
	// pulNumLanguages {in} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.LanguagesLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetrieveTaskOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetrieveTaskOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pXml {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		if o.XML != "" {
			_ptr_pXml := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.XML); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.XML, _ptr_pXml); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetrieveTaskOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pXml {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		_ptr_pXml := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.XML); err != nil {
				return err
			}
			return nil
		})
		_s_pXml := func(ptr interface{}) { o.XML = *ptr.(*string) }
		if err := w.ReadPointer(&o.XML, _s_pXml, _ptr_pXml); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RetrieveTaskRequest structure represents the SchRpcRetrieveTask operation request
type RetrieveTaskRequest struct {
	// path: MUST contain the full path associated with an existing task as specified in
	// section 2.3.11.
	Path string `idl:"name:path;string" json:"path"`
	// lpcwszLanguagesBuffer: If non-NULL, this parameter MUST contain a list of language
	// names separated by the '\' character. The language names MUST be taken from the "String
	// name" column of the table.
	LanguagesBuffer string `idl:"name:lpcwszLanguagesBuffer;string" json:"languages_buffer"`
	// pulNumLanguages: The client SHOULD specify the number of language names in languagesBuffer.
	// The server MUST ignore this parameter.
	LanguagesLength uint32 `idl:"name:pulNumLanguages" json:"languages_length"`
}

func (o *RetrieveTaskRequest) xxx_ToOp(ctx context.Context, op *xxx_RetrieveTaskOperation) *xxx_RetrieveTaskOperation {
	if op == nil {
		op = &xxx_RetrieveTaskOperation{}
	}
	if o == nil {
		return op
	}
	o.Path = op.Path
	o.LanguagesBuffer = op.LanguagesBuffer
	o.LanguagesLength = op.LanguagesLength
	return op
}

func (o *RetrieveTaskRequest) xxx_FromOp(ctx context.Context, op *xxx_RetrieveTaskOperation) {
	if o == nil {
		return
	}
	o.Path = op.Path
	o.LanguagesBuffer = op.LanguagesBuffer
	o.LanguagesLength = op.LanguagesLength
}
func (o *RetrieveTaskRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RetrieveTaskRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RetrieveTaskOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RetrieveTaskResponse structure represents the SchRpcRetrieveTask operation response
type RetrieveTaskResponse struct {
	// pXml: MUST contain the task definition in XML format, localized using the language
	// names contained in the languagesBuffer parameter. For more information about XML
	// localization, see section 2.5.8.
	XML string `idl:"name:pXml;string" json:"xml"`
	// Return: The SchRpcRetrieveTask return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RetrieveTaskResponse) xxx_ToOp(ctx context.Context, op *xxx_RetrieveTaskOperation) *xxx_RetrieveTaskOperation {
	if op == nil {
		op = &xxx_RetrieveTaskOperation{}
	}
	if o == nil {
		return op
	}
	o.XML = op.XML
	o.Return = op.Return
	return op
}

func (o *RetrieveTaskResponse) xxx_FromOp(ctx context.Context, op *xxx_RetrieveTaskOperation) {
	if o == nil {
		return
	}
	o.XML = op.XML
	o.Return = op.Return
}
func (o *RetrieveTaskResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RetrieveTaskResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RetrieveTaskOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateFolderOperation structure represents the SchRpcCreateFolder operation
type xxx_CreateFolderOperation struct {
	Path   string `idl:"name:path;string" json:"path"`
	SDDL   string `idl:"name:sddl;string;pointer:unique" json:"sddl"`
	Flags  uint32 `idl:"name:flags" json:"flags"`
	Return int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateFolderOperation) OpNum() int { return 3 }

func (o *xxx_CreateFolderOperation) OpName() string {
	return "/ITaskSchedulerService/v1/SchRpcCreateFolder"
}

func (o *xxx_CreateFolderOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateFolderOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// path {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
			return err
		}
	}
	// sddl {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.SDDL != "" {
			_ptr_sddl := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.SDDL); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.SDDL, _ptr_sddl); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateFolderOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// path {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
			return err
		}
	}
	// sddl {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_sddl := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.SDDL); err != nil {
				return err
			}
			return nil
		})
		_s_sddl := func(ptr interface{}) { o.SDDL = *ptr.(*string) }
		if err := w.ReadPointer(&o.SDDL, _s_sddl, _ptr_sddl); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateFolderOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateFolderOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateFolderOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateFolderRequest structure represents the SchRpcCreateFolder operation request
type CreateFolderRequest struct {
	// path: MUST contain the full path to be associated with a folder (section 2.3.11).
	Path string `idl:"name:path;string" json:"path"`
	// sddl: If non-NULL, MUST be a security descriptor in SDDL format as specified in [MS-DTYP].
	SDDL string `idl:"name:sddl;string;pointer:unique" json:"sddl"`
	// flags: Unused, MUST be 0.
	Flags uint32 `idl:"name:flags" json:"flags"`
}

func (o *CreateFolderRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateFolderOperation) *xxx_CreateFolderOperation {
	if op == nil {
		op = &xxx_CreateFolderOperation{}
	}
	if o == nil {
		return op
	}
	o.Path = op.Path
	o.SDDL = op.SDDL
	o.Flags = op.Flags
	return op
}

func (o *CreateFolderRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateFolderOperation) {
	if o == nil {
		return
	}
	o.Path = op.Path
	o.SDDL = op.SDDL
	o.Flags = op.Flags
}
func (o *CreateFolderRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateFolderRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateFolderOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateFolderResponse structure represents the SchRpcCreateFolder operation response
type CreateFolderResponse struct {
	// Return: The SchRpcCreateFolder return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateFolderResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateFolderOperation) *xxx_CreateFolderOperation {
	if op == nil {
		op = &xxx_CreateFolderOperation{}
	}
	if o == nil {
		return op
	}
	o.Return = op.Return
	return op
}

func (o *CreateFolderResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateFolderOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *CreateFolderResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateFolderResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateFolderOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSecurityOperation structure represents the SchRpcSetSecurity operation
type xxx_SetSecurityOperation struct {
	Path   string `idl:"name:path;string" json:"path"`
	SDDL   string `idl:"name:sddl;string" json:"sddl"`
	Flags  uint32 `idl:"name:flags" json:"flags"`
	Return int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSecurityOperation) OpNum() int { return 4 }

func (o *xxx_SetSecurityOperation) OpName() string {
	return "/ITaskSchedulerService/v1/SchRpcSetSecurity"
}

func (o *xxx_SetSecurityOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSecurityOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// path {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
			return err
		}
	}
	// sddl {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.SDDL); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSecurityOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// path {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
			return err
		}
	}
	// sddl {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.SDDL); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSecurityOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSecurityOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSecurityOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetSecurityRequest structure represents the SchRpcSetSecurity operation request
type SetSecurityRequest struct {
	// path: MUST contain the full path associated with a task or folder in the format specified
	// in section 2.3.11).
	Path string `idl:"name:path;string" json:"path"`
	// sddl: MUST be a security descriptor in SDDL format as specified in [MS-DTYP].
	SDDL string `idl:"name:sddl;string" json:"sddl"`
	// flags: The flags field MUST contain individual bit flags that MUST have one or more
	// of the following values:
	//
	//	+---+-----+-----+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+-----+---+---+---+---+
	//	| 0 |  1  |  2  | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 |  7  | 8 | 9 | 3 | 1 |
	//	|   |     |     |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |     |   |   | 0 |   |
	//	+---+-----+-----+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+-----+---+---+---+---+
	//	+---+-----+-----+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+-----+---+---+---+---+
	//	| 0 | F T | F F | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | D A | 0 | 0 | 0 | 0 |
	//	+---+-----+-----+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+-----+---+---+---+---+
	//
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|                                |                                                                                  |
	//	|             VALUE              |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| DA TASK_DONT_ADD_PRINCIPAL_ACE | If set to 1, the server MUST NOT add the 'Allow' ACE for the task principal to   |
	//	|                                | the security descriptor.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| FF SCH_FLAG_FOLDER             | If set to 1, the server MUST apply the security descriptor to folders.           |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| FT SCH_FLAG_TASK               | If set to 1, the server MUST apply the security descriptor to tasks.             |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:flags" json:"flags"`
}

func (o *SetSecurityRequest) xxx_ToOp(ctx context.Context, op *xxx_SetSecurityOperation) *xxx_SetSecurityOperation {
	if op == nil {
		op = &xxx_SetSecurityOperation{}
	}
	if o == nil {
		return op
	}
	o.Path = op.Path
	o.SDDL = op.SDDL
	o.Flags = op.Flags
	return op
}

func (o *SetSecurityRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSecurityOperation) {
	if o == nil {
		return
	}
	o.Path = op.Path
	o.SDDL = op.SDDL
	o.Flags = op.Flags
}
func (o *SetSecurityRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetSecurityRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSecurityOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSecurityResponse structure represents the SchRpcSetSecurity operation response
type SetSecurityResponse struct {
	// Return: The SchRpcSetSecurity return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSecurityResponse) xxx_ToOp(ctx context.Context, op *xxx_SetSecurityOperation) *xxx_SetSecurityOperation {
	if op == nil {
		op = &xxx_SetSecurityOperation{}
	}
	if o == nil {
		return op
	}
	o.Return = op.Return
	return op
}

func (o *SetSecurityResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSecurityOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetSecurityResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetSecurityResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSecurityOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSecurityOperation structure represents the SchRpcGetSecurity operation
type xxx_GetSecurityOperation struct {
	Path                string `idl:"name:path;string" json:"path"`
	SecurityInformation uint32 `idl:"name:securityInformation" json:"security_information"`
	SDDL                string `idl:"name:sddl;string" json:"sddl"`
	Return              int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSecurityOperation) OpNum() int { return 5 }

func (o *xxx_GetSecurityOperation) OpName() string {
	return "/ITaskSchedulerService/v1/SchRpcGetSecurity"
}

func (o *xxx_GetSecurityOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSecurityOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// path {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
			return err
		}
	}
	// securityInformation {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SecurityInformation); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSecurityOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// path {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
			return err
		}
	}
	// securityInformation {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SecurityInformation); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSecurityOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSecurityOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// sddl {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		if o.SDDL != "" {
			_ptr_sddl := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.SDDL); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.SDDL, _ptr_sddl); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSecurityOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// sddl {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		_ptr_sddl := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.SDDL); err != nil {
				return err
			}
			return nil
		})
		_s_sddl := func(ptr interface{}) { o.SDDL = *ptr.(*string) }
		if err := w.ReadPointer(&o.SDDL, _s_sddl, _ptr_sddl); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetSecurityRequest structure represents the SchRpcGetSecurity operation request
type GetSecurityRequest struct {
	// path: MUST be the full path associated with a task or folder in the format specified
	// in section 2.3.11).
	Path string `idl:"name:path;string" json:"path"`
	// securityInformation: MUST contain security information in the format of a SECURITY_INFORMATION
	// structure. The SECURITY_INFORMATION structure is defined in [MS-DTYP] section 2.4.7.
	SecurityInformation uint32 `idl:"name:securityInformation" json:"security_information"`
}

func (o *GetSecurityRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSecurityOperation) *xxx_GetSecurityOperation {
	if op == nil {
		op = &xxx_GetSecurityOperation{}
	}
	if o == nil {
		return op
	}
	o.Path = op.Path
	o.SecurityInformation = op.SecurityInformation
	return op
}

func (o *GetSecurityRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSecurityOperation) {
	if o == nil {
		return
	}
	o.Path = op.Path
	o.SecurityInformation = op.SecurityInformation
}
func (o *GetSecurityRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSecurityRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSecurityOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSecurityResponse structure represents the SchRpcGetSecurity operation response
type GetSecurityResponse struct {
	// sddl: MUST point to a buffer that will receive security information in string format.
	// The string format is specified in [MS-DTYP] section 2.5.1.
	SDDL string `idl:"name:sddl;string" json:"sddl"`
	// Return: The SchRpcGetSecurity return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSecurityResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSecurityOperation) *xxx_GetSecurityOperation {
	if op == nil {
		op = &xxx_GetSecurityOperation{}
	}
	if o == nil {
		return op
	}
	o.SDDL = op.SDDL
	o.Return = op.Return
	return op
}

func (o *GetSecurityResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSecurityOperation) {
	if o == nil {
		return
	}
	o.SDDL = op.SDDL
	o.Return = op.Return
}
func (o *GetSecurityResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSecurityResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSecurityOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumFoldersOperation structure represents the SchRpcEnumFolders operation
type xxx_EnumFoldersOperation struct {
	Path           string   `idl:"name:path;string" json:"path"`
	Flags          uint32   `idl:"name:flags" json:"flags"`
	StartIndex     uint32   `idl:"name:pStartIndex" json:"start_index"`
	RequestedCount uint32   `idl:"name:cRequested" json:"requested_count"`
	NamesCount     uint32   `idl:"name:pcNames" json:"names_count"`
	Names          []string `idl:"name:pNames;size_is:(, pcNames);string" json:"names"`
	Return         int32    `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumFoldersOperation) OpNum() int { return 6 }

func (o *xxx_EnumFoldersOperation) OpName() string {
	return "/ITaskSchedulerService/v1/SchRpcEnumFolders"
}

func (o *xxx_EnumFoldersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumFoldersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// path {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pStartIndex {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.StartIndex); err != nil {
			return err
		}
	}
	// cRequested {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequestedCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumFoldersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// path {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pStartIndex {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.StartIndex); err != nil {
			return err
		}
	}
	// cRequested {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequestedCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumFoldersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Names != nil && o.NamesCount == 0 {
		o.NamesCount = uint32(len(o.Names))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumFoldersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pStartIndex {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.StartIndex); err != nil {
			return err
		}
	}
	// pcNames {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NamesCount); err != nil {
			return err
		}
	}
	// pNames {out} (1:{string, pointer=ref}*(2))(2:{string, alias=TASK_NAMES}*(1)[dim:0,size_is=pcNames]*(1)[dim:0,string,null](wchar))
	{
		if o.Names != nil || o.NamesCount > 0 {
			_ptr_pNames := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.NamesCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.Names {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.Names[i1] != "" {
						_ptr_pNames := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
							if err := ndr.WriteUTF16NString(ctx, w, o.Names[i1]); err != nil {
								return err
							}
							return nil
						})
						if err := w.WritePointer(&o.Names[i1], _ptr_pNames); err != nil {
							return err
						}
					} else {
						if err := w.WritePointer(nil); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.Names); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Names, _ptr_pNames); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumFoldersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pStartIndex {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.StartIndex); err != nil {
			return err
		}
	}
	// pcNames {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NamesCount); err != nil {
			return err
		}
	}
	// pNames {out} (1:{string, pointer=ref}*(2))(2:{string, alias=TASK_NAMES,pointer=ref}*(1)[dim:0,size_is=pcNames]*(1)[dim:0,string,null](wchar))
	{
		_ptr_pNames := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Names", sizeInfo[0])
			}
			o.Names = make([]string, sizeInfo[0])
			for i1 := range o.Names {
				i1 := i1
				_ptr_pNames := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
					if err := ndr.ReadUTF16NString(ctx, w, &o.Names[i1]); err != nil {
						return err
					}
					return nil
				})
				_s_pNames := func(ptr interface{}) { o.Names[i1] = *ptr.(*string) }
				if err := w.ReadPointer(&o.Names[i1], _s_pNames, _ptr_pNames); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pNames := func(ptr interface{}) { o.Names = *ptr.(*[]string) }
		if err := w.ReadPointer(&o.Names, _s_pNames, _ptr_pNames); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumFoldersRequest structure represents the SchRpcEnumFolders operation request
type EnumFoldersRequest struct {
	// path: MUST contain the full path associated with a folder using the format specified
	// in section 2.3.11.
	Path string `idl:"name:path;string" json:"path"`
	// flags: All bits except TASK_ENUM_HIDDEN MUST be set to zero and the server MUST return
	// an error if undefined bits are set to 1. This field has one or more of the following
	// values:
	//
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
	//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | H |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//
	// Where the bits are defined as:
	//
	//	+--------------------+----------------------------------------------------------------------------------+
	//	|                    |                                                                                  |
	//	|       VALUE        |                                   DESCRIPTION                                    |
	//	|                    |                                                                                  |
	//	+--------------------+----------------------------------------------------------------------------------+
	//	+--------------------+----------------------------------------------------------------------------------+
	//	| H TASK_ENUM_HIDDEN | If set to 1, the server MUST include hidden tasks in the enumeration, otherwise  |
	//	|                    | the server MUST exclude hidden tasks from the enumeration.                       |
	//	+--------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:flags" json:"flags"`
	// pStartIndex: MUST contain the index at which to start enumeration. If the server
	// returns S_FALSE, the server MUST update startIndex to contain the index at which
	// the enumeration MUST resume.
	StartIndex uint32 `idl:"name:pStartIndex" json:"start_index"`
	// cRequested: MUST contain the number of entries requested. The server MUST NOT return
	// more than cRequested entries.
	RequestedCount uint32 `idl:"name:cRequested" json:"requested_count"`
}

func (o *EnumFoldersRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumFoldersOperation) *xxx_EnumFoldersOperation {
	if op == nil {
		op = &xxx_EnumFoldersOperation{}
	}
	if o == nil {
		return op
	}
	o.Path = op.Path
	o.Flags = op.Flags
	o.StartIndex = op.StartIndex
	o.RequestedCount = op.RequestedCount
	return op
}

func (o *EnumFoldersRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumFoldersOperation) {
	if o == nil {
		return
	}
	o.Path = op.Path
	o.Flags = op.Flags
	o.StartIndex = op.StartIndex
	o.RequestedCount = op.RequestedCount
}
func (o *EnumFoldersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumFoldersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumFoldersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumFoldersResponse structure represents the SchRpcEnumFolders operation response
type EnumFoldersResponse struct {
	// pStartIndex: MUST contain the index at which to start enumeration. If the server
	// returns S_FALSE, the server MUST update startIndex to contain the index at which
	// the enumeration MUST resume.
	StartIndex uint32 `idl:"name:pStartIndex" json:"start_index"`
	// pcNames: MUST contain a count of enumerated subfolder names contained in pNames.
	NamesCount uint32 `idl:"name:pcNames" json:"names_count"`
	// pNames: Buffer MUST contain returned folder names.
	Names []string `idl:"name:pNames;size_is:(, pcNames);string" json:"names"`
	// Return: The SchRpcEnumFolders return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumFoldersResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumFoldersOperation) *xxx_EnumFoldersOperation {
	if op == nil {
		op = &xxx_EnumFoldersOperation{}
	}
	if o == nil {
		return op
	}
	o.StartIndex = op.StartIndex
	o.NamesCount = op.NamesCount
	o.Names = op.Names
	o.Return = op.Return
	return op
}

func (o *EnumFoldersResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumFoldersOperation) {
	if o == nil {
		return
	}
	o.StartIndex = op.StartIndex
	o.NamesCount = op.NamesCount
	o.Names = op.Names
	o.Return = op.Return
}
func (o *EnumFoldersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumFoldersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumFoldersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumTasksOperation structure represents the SchRpcEnumTasks operation
type xxx_EnumTasksOperation struct {
	Path           string   `idl:"name:path;string" json:"path"`
	Flags          uint32   `idl:"name:flags" json:"flags"`
	StartIndex     uint32   `idl:"name:startIndex" json:"start_index"`
	RequestedCount uint32   `idl:"name:cRequested" json:"requested_count"`
	NamesCount     uint32   `idl:"name:pcNames" json:"names_count"`
	Names          []string `idl:"name:pNames;size_is:(, pcNames);string" json:"names"`
	Return         int32    `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumTasksOperation) OpNum() int { return 7 }

func (o *xxx_EnumTasksOperation) OpName() string { return "/ITaskSchedulerService/v1/SchRpcEnumTasks" }

func (o *xxx_EnumTasksOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumTasksOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// path {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// startIndex {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.StartIndex); err != nil {
			return err
		}
	}
	// cRequested {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequestedCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumTasksOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// path {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// startIndex {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.StartIndex); err != nil {
			return err
		}
	}
	// cRequested {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequestedCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumTasksOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Names != nil && o.NamesCount == 0 {
		o.NamesCount = uint32(len(o.Names))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumTasksOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// startIndex {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.StartIndex); err != nil {
			return err
		}
	}
	// pcNames {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NamesCount); err != nil {
			return err
		}
	}
	// pNames {out} (1:{string, pointer=ref}*(2))(2:{string, alias=TASK_NAMES}*(1)[dim:0,size_is=pcNames]*(1)[dim:0,string,null](wchar))
	{
		if o.Names != nil || o.NamesCount > 0 {
			_ptr_pNames := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.NamesCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.Names {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.Names[i1] != "" {
						_ptr_pNames := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
							if err := ndr.WriteUTF16NString(ctx, w, o.Names[i1]); err != nil {
								return err
							}
							return nil
						})
						if err := w.WritePointer(&o.Names[i1], _ptr_pNames); err != nil {
							return err
						}
					} else {
						if err := w.WritePointer(nil); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.Names); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Names, _ptr_pNames); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumTasksOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// startIndex {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.StartIndex); err != nil {
			return err
		}
	}
	// pcNames {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NamesCount); err != nil {
			return err
		}
	}
	// pNames {out} (1:{string, pointer=ref}*(2))(2:{string, alias=TASK_NAMES,pointer=ref}*(1)[dim:0,size_is=pcNames]*(1)[dim:0,string,null](wchar))
	{
		_ptr_pNames := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Names", sizeInfo[0])
			}
			o.Names = make([]string, sizeInfo[0])
			for i1 := range o.Names {
				i1 := i1
				_ptr_pNames := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
					if err := ndr.ReadUTF16NString(ctx, w, &o.Names[i1]); err != nil {
						return err
					}
					return nil
				})
				_s_pNames := func(ptr interface{}) { o.Names[i1] = *ptr.(*string) }
				if err := w.ReadPointer(&o.Names[i1], _s_pNames, _ptr_pNames); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pNames := func(ptr interface{}) { o.Names = *ptr.(*[]string) }
		if err := w.ReadPointer(&o.Names, _s_pNames, _ptr_pNames); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumTasksRequest structure represents the SchRpcEnumTasks operation request
type EnumTasksRequest struct {
	// path: MUST contain the full path associated with a folder as specified in section
	// 2.3.11.
	Path string `idl:"name:path;string" json:"path"`
	// flags: The flags parameter MUST contain individual bit flags that MUST have one or
	// more of the following values:
	//
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
	//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | H |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//
	//
	//	+--------------------+----------------------------------------------------------------------------------+
	//	|                    |                                                                                  |
	//	|       VALUE        |                                   DESCRIPTION                                    |
	//	|                    |                                                                                  |
	//	+--------------------+----------------------------------------------------------------------------------+
	//	+--------------------+----------------------------------------------------------------------------------+
	//	| H TASK_ENUM_HIDDEN | If set to 1, the server MUST include hidden tasks in the enumeration, otherwise  |
	//	|                    | the server MUST exclude hidden tasks from the enumeration.                       |
	//	+--------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:flags" json:"flags"`
	// startIndex: MUST contain the index at which to start enumeration. If the server returns
	// S_FALSE, the server MUST update startIndex to contain the index at which the enumeration
	// MUST resume.
	StartIndex uint32 `idl:"name:startIndex" json:"start_index"`
	// cRequested: MUST contain the number of entries requested. The server MUST NOT return
	// more than cRequested entries.
	RequestedCount uint32 `idl:"name:cRequested" json:"requested_count"`
}

func (o *EnumTasksRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumTasksOperation) *xxx_EnumTasksOperation {
	if op == nil {
		op = &xxx_EnumTasksOperation{}
	}
	if o == nil {
		return op
	}
	o.Path = op.Path
	o.Flags = op.Flags
	o.StartIndex = op.StartIndex
	o.RequestedCount = op.RequestedCount
	return op
}

func (o *EnumTasksRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumTasksOperation) {
	if o == nil {
		return
	}
	o.Path = op.Path
	o.Flags = op.Flags
	o.StartIndex = op.StartIndex
	o.RequestedCount = op.RequestedCount
}
func (o *EnumTasksRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumTasksRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumTasksOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumTasksResponse structure represents the SchRpcEnumTasks operation response
type EnumTasksResponse struct {
	// startIndex: MUST contain the index at which to start enumeration. If the server returns
	// S_FALSE, the server MUST update startIndex to contain the index at which the enumeration
	// MUST resume.
	StartIndex uint32 `idl:"name:startIndex" json:"start_index"`
	// pcNames: The server MUST set pcNames to equal the number of enumerated tasks returned
	// in the pNames parameter.
	NamesCount uint32 `idl:"name:pcNames" json:"names_count"`
	// pNames: Buffer that MUST contain returned task names.
	Names []string `idl:"name:pNames;size_is:(, pcNames);string" json:"names"`
	// Return: The SchRpcEnumTasks return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumTasksResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumTasksOperation) *xxx_EnumTasksOperation {
	if op == nil {
		op = &xxx_EnumTasksOperation{}
	}
	if o == nil {
		return op
	}
	o.StartIndex = op.StartIndex
	o.NamesCount = op.NamesCount
	o.Names = op.Names
	o.Return = op.Return
	return op
}

func (o *EnumTasksResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumTasksOperation) {
	if o == nil {
		return
	}
	o.StartIndex = op.StartIndex
	o.NamesCount = op.NamesCount
	o.Names = op.Names
	o.Return = op.Return
}
func (o *EnumTasksResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumTasksResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumTasksOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumInstancesOperation structure represents the SchRpcEnumInstances operation
type xxx_EnumInstancesOperation struct {
	Path       string       `idl:"name:path;string;pointer:unique" json:"path"`
	Flags      uint32       `idl:"name:flags" json:"flags"`
	GUIDsCount uint32       `idl:"name:pcGuids" json:"guids_count"`
	GUIDs      []*dtyp.GUID `idl:"name:pGuids;size_is:(, pcGuids)" json:"guids"`
	Return     int32        `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumInstancesOperation) OpNum() int { return 8 }

func (o *xxx_EnumInstancesOperation) OpName() string {
	return "/ITaskSchedulerService/v1/SchRpcEnumInstances"
}

func (o *xxx_EnumInstancesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumInstancesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// path {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Path != "" {
			_ptr_path := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_path); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumInstancesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// path {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_path := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
				return err
			}
			return nil
		})
		_s_path := func(ptr interface{}) { o.Path = *ptr.(*string) }
		if err := w.ReadPointer(&o.Path, _s_path, _ptr_path); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumInstancesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.GUIDs != nil && o.GUIDsCount == 0 {
		o.GUIDsCount = uint32(len(o.GUIDs))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumInstancesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pcGuids {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.GUIDsCount); err != nil {
			return err
		}
	}
	// pGuids {out} (1:{pointer=ref}*(2)*(1))(2:{alias=GUID}[dim:0,size_is=pcGuids](struct))
	{
		if o.GUIDs != nil || o.GUIDsCount > 0 {
			_ptr_pGuids := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.GUIDsCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.GUIDs {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.GUIDs[i1] != nil {
						if err := o.GUIDs[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.GUIDs); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.GUIDs, _ptr_pGuids); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumInstancesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pcGuids {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.GUIDsCount); err != nil {
			return err
		}
	}
	// pGuids {out} (1:{pointer=ref}*(2)*(1))(2:{alias=GUID}[dim:0,size_is=pcGuids](struct))
	{
		_ptr_pGuids := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.GUIDs", sizeInfo[0])
			}
			o.GUIDs = make([]*dtyp.GUID, sizeInfo[0])
			for i1 := range o.GUIDs {
				i1 := i1
				if o.GUIDs[i1] == nil {
					o.GUIDs[i1] = &dtyp.GUID{}
				}
				if err := o.GUIDs[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pGuids := func(ptr interface{}) { o.GUIDs = *ptr.(*[]*dtyp.GUID) }
		if err := w.ReadPointer(&o.GUIDs, _s_pGuids, _ptr_pGuids); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumInstancesRequest structure represents the SchRpcEnumInstances operation request
type EnumInstancesRequest struct {
	// path: MUST contain the full path to a task in the format specified in section 2.3.11.
	// If NULL is specified, all instances for all tasks MUST be returned.
	Path string `idl:"name:path;string;pointer:unique" json:"path"`
	// flags: The flags field MUST contain individual bit flags that MUST have one or more
	// of the following values:
	//
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
	//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | H |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//
	//
	//	+--------------------+---------------------------------------------------------------+
	//	|                    |                                                               |
	//	|       VALUE        |                          DESCRIPTION                          |
	//	|                    |                                                               |
	//	+--------------------+---------------------------------------------------------------+
	//	+--------------------+---------------------------------------------------------------+
	//	| H TASK_ENUM_HIDDEN | If set to 1, hidden tasks MUST be included in the result set. |
	//	+--------------------+---------------------------------------------------------------+
	Flags uint32 `idl:"name:flags" json:"flags"`
}

func (o *EnumInstancesRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumInstancesOperation) *xxx_EnumInstancesOperation {
	if op == nil {
		op = &xxx_EnumInstancesOperation{}
	}
	if o == nil {
		return op
	}
	o.Path = op.Path
	o.Flags = op.Flags
	return op
}

func (o *EnumInstancesRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumInstancesOperation) {
	if o == nil {
		return
	}
	o.Path = op.Path
	o.Flags = op.Flags
}
func (o *EnumInstancesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumInstancesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumInstancesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumInstancesResponse structure represents the SchRpcEnumInstances operation response
type EnumInstancesResponse struct {
	// pcGuids: MUST contain the number of instances.
	GUIDsCount uint32 `idl:"name:pcGuids" json:"guids_count"`
	// pGuids: MUST be an array of GUIDs.
	GUIDs []*dtyp.GUID `idl:"name:pGuids;size_is:(, pcGuids)" json:"guids"`
	// Return: The SchRpcEnumInstances return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumInstancesResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumInstancesOperation) *xxx_EnumInstancesOperation {
	if op == nil {
		op = &xxx_EnumInstancesOperation{}
	}
	if o == nil {
		return op
	}
	o.GUIDsCount = op.GUIDsCount
	o.GUIDs = op.GUIDs
	o.Return = op.Return
	return op
}

func (o *EnumInstancesResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumInstancesOperation) {
	if o == nil {
		return
	}
	o.GUIDsCount = op.GUIDsCount
	o.GUIDs = op.GUIDs
	o.Return = op.Return
}
func (o *EnumInstancesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumInstancesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumInstancesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetInstanceInfoOperation structure represents the SchRpcGetInstanceInfo operation
type xxx_GetInstanceInfoOperation struct {
	GUID                *dtyp.GUID   `idl:"name:guid" json:"guid"`
	Path                string       `idl:"name:pPath;string" json:"path"`
	State               uint32       `idl:"name:pState" json:"state"`
	CurrentAction       string       `idl:"name:pCurrentAction;string" json:"current_action"`
	Info                string       `idl:"name:pInfo;string" json:"info"`
	GroupInstancesCount uint32       `idl:"name:pcGroupInstances" json:"group_instances_count"`
	GroupInstances      []*dtyp.GUID `idl:"name:pGroupInstances;size_is:(, pcGroupInstances)" json:"group_instances"`
	EnginePID           uint32       `idl:"name:pEnginePID" json:"engine_pid"`
	Return              int32        `idl:"name:Return" json:"return"`
}

func (o *xxx_GetInstanceInfoOperation) OpNum() int { return 9 }

func (o *xxx_GetInstanceInfoOperation) OpName() string {
	return "/ITaskSchedulerService/v1/SchRpcGetInstanceInfo"
}

func (o *xxx_GetInstanceInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInstanceInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// guid {in} (1:{alias=GUID}(struct))
	{
		if o.GUID != nil {
			if err := o.GUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetInstanceInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// guid {in} (1:{alias=GUID}(struct))
	{
		if o.GUID == nil {
			o.GUID = &dtyp.GUID{}
		}
		if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInstanceInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.GroupInstances != nil && o.GroupInstancesCount == 0 {
		o.GroupInstancesCount = uint32(len(o.GroupInstances))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInstanceInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pPath {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		if o.Path != "" {
			_ptr_pPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_pPath); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pState {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.State); err != nil {
			return err
		}
	}
	// pCurrentAction {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		if o.CurrentAction != "" {
			_ptr_pCurrentAction := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.CurrentAction); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.CurrentAction, _ptr_pCurrentAction); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pInfo {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		if o.Info != "" {
			_ptr_pInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Info); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Info, _ptr_pInfo); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pcGroupInstances {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.GroupInstancesCount); err != nil {
			return err
		}
	}
	// pGroupInstances {out} (1:{pointer=ref}*(2)*(1))(2:{alias=GUID}[dim:0,size_is=pcGroupInstances](struct))
	{
		if o.GroupInstances != nil || o.GroupInstancesCount > 0 {
			_ptr_pGroupInstances := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.GroupInstancesCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.GroupInstances {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.GroupInstances[i1] != nil {
						if err := o.GroupInstances[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.GroupInstances); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.GroupInstances, _ptr_pGroupInstances); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pEnginePID {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.EnginePID); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInstanceInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pPath {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		_ptr_pPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
				return err
			}
			return nil
		})
		_s_pPath := func(ptr interface{}) { o.Path = *ptr.(*string) }
		if err := w.ReadPointer(&o.Path, _s_pPath, _ptr_pPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pState {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.State); err != nil {
			return err
		}
	}
	// pCurrentAction {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		_ptr_pCurrentAction := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.CurrentAction); err != nil {
				return err
			}
			return nil
		})
		_s_pCurrentAction := func(ptr interface{}) { o.CurrentAction = *ptr.(*string) }
		if err := w.ReadPointer(&o.CurrentAction, _s_pCurrentAction, _ptr_pCurrentAction); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pInfo {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		_ptr_pInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Info); err != nil {
				return err
			}
			return nil
		})
		_s_pInfo := func(ptr interface{}) { o.Info = *ptr.(*string) }
		if err := w.ReadPointer(&o.Info, _s_pInfo, _ptr_pInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pcGroupInstances {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.GroupInstancesCount); err != nil {
			return err
		}
	}
	// pGroupInstances {out} (1:{pointer=ref}*(2)*(1))(2:{alias=GUID}[dim:0,size_is=pcGroupInstances](struct))
	{
		_ptr_pGroupInstances := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.GroupInstances", sizeInfo[0])
			}
			o.GroupInstances = make([]*dtyp.GUID, sizeInfo[0])
			for i1 := range o.GroupInstances {
				i1 := i1
				if o.GroupInstances[i1] == nil {
					o.GroupInstances[i1] = &dtyp.GUID{}
				}
				if err := o.GroupInstances[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pGroupInstances := func(ptr interface{}) { o.GroupInstances = *ptr.(*[]*dtyp.GUID) }
		if err := w.ReadPointer(&o.GroupInstances, _s_pGroupInstances, _ptr_pGroupInstances); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pEnginePID {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.EnginePID); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetInstanceInfoRequest structure represents the SchRpcGetInstanceInfo operation request
type GetInstanceInfoRequest struct {
	// guid: MUST contain the GUID of the running task instance.
	GUID *dtyp.GUID `idl:"name:guid" json:"guid"`
}

func (o *GetInstanceInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_GetInstanceInfoOperation) *xxx_GetInstanceInfoOperation {
	if op == nil {
		op = &xxx_GetInstanceInfoOperation{}
	}
	if o == nil {
		return op
	}
	o.GUID = op.GUID
	return op
}

func (o *GetInstanceInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetInstanceInfoOperation) {
	if o == nil {
		return
	}
	o.GUID = op.GUID
}
func (o *GetInstanceInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetInstanceInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetInstanceInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetInstanceInfoResponse structure represents the SchRpcGetInstanceInfo operation response
type GetInstanceInfoResponse struct {
	// pPath: MUST be the location where a string containing the task's path is to be returned
	// in the format specified in section 2.3.11. If NULL, specifies that the path is not
	// requested.
	Path string `idl:"name:pPath;string" json:"path"`
	// pState: Location where the state of the instance (section 2.3.13) is to be returned.
	// If NULL, specifies that the state is not requested.
	State uint32 `idl:"name:pState" json:"state"`
	// pCurrentAction: MUST be the location where the name (id) of the action the task is
	// currently executing is to be returned. If NULL, specifies that the current action
	// is not requested.
	CurrentAction string `idl:"name:pCurrentAction;string" json:"current_action"`
	// pInfo: Unused. If non-NULL, the server MUST set the string to NULL.
	Info string `idl:"name:pInfo;string" json:"info"`
	// pcGroupInstances: Unused.
	GroupInstancesCount uint32 `idl:"name:pcGroupInstances" json:"group_instances_count"`
	// pGroupInstances: Unused. If non-NULL, the server MUST set the GUID to NULL.
	GroupInstances []*dtyp.GUID `idl:"name:pGroupInstances;size_is:(, pcGroupInstances)" json:"group_instances"`
	// pEnginePID: MUST be the location where the Process ID of the process executing the
	// task is to be returned. If NULL, specifies that the Process ID is not requested.
	EnginePID uint32 `idl:"name:pEnginePID" json:"engine_pid"`
	// Return: The SchRpcGetInstanceInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetInstanceInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_GetInstanceInfoOperation) *xxx_GetInstanceInfoOperation {
	if op == nil {
		op = &xxx_GetInstanceInfoOperation{}
	}
	if o == nil {
		return op
	}
	o.Path = op.Path
	o.State = op.State
	o.CurrentAction = op.CurrentAction
	o.Info = op.Info
	o.GroupInstancesCount = op.GroupInstancesCount
	o.GroupInstances = op.GroupInstances
	o.EnginePID = op.EnginePID
	o.Return = op.Return
	return op
}

func (o *GetInstanceInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetInstanceInfoOperation) {
	if o == nil {
		return
	}
	o.Path = op.Path
	o.State = op.State
	o.CurrentAction = op.CurrentAction
	o.Info = op.Info
	o.GroupInstancesCount = op.GroupInstancesCount
	o.GroupInstances = op.GroupInstances
	o.EnginePID = op.EnginePID
	o.Return = op.Return
}
func (o *GetInstanceInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetInstanceInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetInstanceInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_StopInstanceOperation structure represents the SchRpcStopInstance operation
type xxx_StopInstanceOperation struct {
	GUID   *dtyp.GUID `idl:"name:guid" json:"guid"`
	Flags  uint32     `idl:"name:flags" json:"flags"`
	Return int32      `idl:"name:Return" json:"return"`
}

func (o *xxx_StopInstanceOperation) OpNum() int { return 10 }

func (o *xxx_StopInstanceOperation) OpName() string {
	return "/ITaskSchedulerService/v1/SchRpcStopInstance"
}

func (o *xxx_StopInstanceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopInstanceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// guid {in} (1:{alias=GUID}(struct))
	{
		if o.GUID != nil {
			if err := o.GUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopInstanceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// guid {in} (1:{alias=GUID}(struct))
	{
		if o.GUID == nil {
			o.GUID = &dtyp.GUID{}
		}
		if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopInstanceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopInstanceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopInstanceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// StopInstanceRequest structure represents the SchRpcStopInstance operation request
type StopInstanceRequest struct {
	// guid: MUST contain the task instance GUID.
	GUID *dtyp.GUID `idl:"name:guid" json:"guid"`
	// flags: Unused. The client MUST send 0 and the server MUST return an error if nonzero.
	Flags uint32 `idl:"name:flags" json:"flags"`
}

func (o *StopInstanceRequest) xxx_ToOp(ctx context.Context, op *xxx_StopInstanceOperation) *xxx_StopInstanceOperation {
	if op == nil {
		op = &xxx_StopInstanceOperation{}
	}
	if o == nil {
		return op
	}
	o.GUID = op.GUID
	o.Flags = op.Flags
	return op
}

func (o *StopInstanceRequest) xxx_FromOp(ctx context.Context, op *xxx_StopInstanceOperation) {
	if o == nil {
		return
	}
	o.GUID = op.GUID
	o.Flags = op.Flags
}
func (o *StopInstanceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *StopInstanceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StopInstanceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// StopInstanceResponse structure represents the SchRpcStopInstance operation response
type StopInstanceResponse struct {
	// Return: The SchRpcStopInstance return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *StopInstanceResponse) xxx_ToOp(ctx context.Context, op *xxx_StopInstanceOperation) *xxx_StopInstanceOperation {
	if op == nil {
		op = &xxx_StopInstanceOperation{}
	}
	if o == nil {
		return op
	}
	o.Return = op.Return
	return op
}

func (o *StopInstanceResponse) xxx_FromOp(ctx context.Context, op *xxx_StopInstanceOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *StopInstanceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *StopInstanceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StopInstanceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_StopOperation structure represents the SchRpcStop operation
type xxx_StopOperation struct {
	Path   string `idl:"name:path;string;pointer:unique" json:"path"`
	Flags  uint32 `idl:"name:flags" json:"flags"`
	Return int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_StopOperation) OpNum() int { return 11 }

func (o *xxx_StopOperation) OpName() string { return "/ITaskSchedulerService/v1/SchRpcStop" }

func (o *xxx_StopOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// path {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Path != "" {
			_ptr_path := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_path); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// path {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_path := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
				return err
			}
			return nil
		})
		_s_path := func(ptr interface{}) { o.Path = *ptr.(*string) }
		if err := w.ReadPointer(&o.Path, _s_path, _ptr_path); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// StopRequest structure represents the SchRpcStop operation request
type StopRequest struct {
	// path: MUST contain the full path to a task using the format specified in section
	// 2.3.11.
	Path string `idl:"name:path;string;pointer:unique" json:"path"`
	// flags: Unused. The client MUST set 0, and the server MUST return an error if nonzero.
	Flags uint32 `idl:"name:flags" json:"flags"`
}

func (o *StopRequest) xxx_ToOp(ctx context.Context, op *xxx_StopOperation) *xxx_StopOperation {
	if op == nil {
		op = &xxx_StopOperation{}
	}
	if o == nil {
		return op
	}
	o.Path = op.Path
	o.Flags = op.Flags
	return op
}

func (o *StopRequest) xxx_FromOp(ctx context.Context, op *xxx_StopOperation) {
	if o == nil {
		return
	}
	o.Path = op.Path
	o.Flags = op.Flags
}
func (o *StopRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *StopRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StopOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// StopResponse structure represents the SchRpcStop operation response
type StopResponse struct {
	// Return: The SchRpcStop return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *StopResponse) xxx_ToOp(ctx context.Context, op *xxx_StopOperation) *xxx_StopOperation {
	if op == nil {
		op = &xxx_StopOperation{}
	}
	if o == nil {
		return op
	}
	o.Return = op.Return
	return op
}

func (o *StopResponse) xxx_FromOp(ctx context.Context, op *xxx_StopOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *StopResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *StopResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StopOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RunOperation structure represents the SchRpcRun operation
type xxx_RunOperation struct {
	Path      string     `idl:"name:path;string" json:"path"`
	ArgsCount uint32     `idl:"name:cArgs" json:"args_count"`
	Args      []string   `idl:"name:pArgs;size_is:(cArgs);string;pointer:unique" json:"args"`
	Flags     uint32     `idl:"name:flags" json:"flags"`
	SessionID uint32     `idl:"name:sessionId" json:"session_id"`
	User      string     `idl:"name:user;string;pointer:unique" json:"user"`
	GUID      *dtyp.GUID `idl:"name:pGuid" json:"guid"`
	Return    int32      `idl:"name:Return" json:"return"`
}

func (o *xxx_RunOperation) OpNum() int { return 12 }

func (o *xxx_RunOperation) OpName() string { return "/ITaskSchedulerService/v1/SchRpcRun" }

func (o *xxx_RunOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Args != nil && o.ArgsCount == 0 {
		o.ArgsCount = uint32(len(o.Args))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RunOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// path {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
			return err
		}
	}
	// cArgs {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ArgsCount); err != nil {
			return err
		}
	}
	// pArgs {in} (1:{string, pointer=unique}*(1)[dim:0,size_is=cArgs]*(1)[dim:0,string,null](wchar))
	{
		if o.Args != nil || o.ArgsCount > 0 {
			_ptr_pArgs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ArgsCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.Args {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.Args[i1] != "" {
						_ptr_pArgs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
							if err := ndr.WriteUTF16NString(ctx, w, o.Args[i1]); err != nil {
								return err
							}
							return nil
						})
						if err := w.WritePointer(&o.Args[i1], _ptr_pArgs); err != nil {
							return err
						}
					} else {
						if err := w.WritePointer(nil); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.Args); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Args, _ptr_pArgs); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// sessionId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SessionID); err != nil {
			return err
		}
	}
	// user {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.User != "" {
			_ptr_user := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.User); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.User, _ptr_user); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RunOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// path {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
			return err
		}
	}
	// cArgs {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ArgsCount); err != nil {
			return err
		}
	}
	// pArgs {in} (1:{string, pointer=unique}*(1)[dim:0,size_is=cArgs]*(1)[dim:0,string,null](wchar))
	{
		_ptr_pArgs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Args", sizeInfo[0])
			}
			o.Args = make([]string, sizeInfo[0])
			for i1 := range o.Args {
				i1 := i1
				_ptr_pArgs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
					if err := ndr.ReadUTF16NString(ctx, w, &o.Args[i1]); err != nil {
						return err
					}
					return nil
				})
				_s_pArgs := func(ptr interface{}) { o.Args[i1] = *ptr.(*string) }
				if err := w.ReadPointer(&o.Args[i1], _s_pArgs, _ptr_pArgs); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pArgs := func(ptr interface{}) { o.Args = *ptr.(*[]string) }
		if err := w.ReadPointer(&o.Args, _s_pArgs, _ptr_pArgs); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// sessionId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SessionID); err != nil {
			return err
		}
	}
	// user {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_user := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.User); err != nil {
				return err
			}
			return nil
		})
		_s_user := func(ptr interface{}) { o.User = *ptr.(*string) }
		if err := w.ReadPointer(&o.User, _s_user, _ptr_user); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RunOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RunOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pGuid {out} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.GUID != nil {
			if err := o.GUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RunOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pGuid {out} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.GUID == nil {
			o.GUID = &dtyp.GUID{}
		}
		if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RunRequest structure represents the SchRpcRun operation request
type RunRequest struct {
	// path: MUST be the full path to a task using the format specified in section 2.3.11.
	Path string `idl:"name:path;string" json:"path"`
	// cArgs: MUST be the number of strings supplied in pArgs.
	ArgsCount uint32 `idl:"name:cArgs" json:"args_count"`
	// pArgs: MUST be an array of strings of size cArgs. This parameter MUST supply string
	// values for parameter substitution, as specified in section 2.5.9.
	Args []string `idl:"name:pArgs;size_is:(cArgs);string;pointer:unique" json:"args"`
	// flags: The flags field MUST contain individual bit flags that MUST have one or more
	// of the following values.
	//
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+-----+-----+-----+-----+
	//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 |  8  |  9  |  3  |  1  |
	//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |     |     |  0  |     |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+-----+-----+-----+-----+
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+-----+-----+-----+-----+
	//	| 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | U S | S I | I C | A S |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+-----+-----+-----+-----+
	//
	// Flags
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|                                |                                                                                  |
	//	|             VALUE              |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| AS TASK_RUN_AS_SELF            | If set to 1, the server MUST run the task in the context of the caller.          |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| IC TASK_RUN_IGNORE_CONSTRAINTS | If set to 1, the server MUST ignore the conditions in the task definition.       |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| SI TASK_RUN_USE_SESSION_ID     | If set to 1, the server MUST run the task in the login session specified by the  |
	//	|                                | sessionId parameter.                                                             |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| US TASK_RUN_USER_SID           | If set to 1, the userId parameter MUST contain a SID string.                     |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:flags" json:"flags"`
	// sessionId: MUST specify a terminal server session in which to run the task. The server
	// MUST ignore this parameter unless the TASK_RUN_USE_SESSION_ID bit in the flags parameter
	// is set. For more information on terminal server sessions, see [MSDN-WSI].
	SessionID uint32 `idl:"name:sessionId" json:"session_id"`
	// user: If non-NULL, MUST specify the user context under which to run the task. If
	// the TASK_RUN_USER_SID bit in the flags parameter is set, userID MUST contain a SID
	// string. Otherwise, userID MUST contain an account name. If the TASK_RUN_AS_SELF bit
	// in the flag parameter is set, the server MUST ignore the userId parameter.
	User string `idl:"name:user;string;pointer:unique" json:"user"`
}

func (o *RunRequest) xxx_ToOp(ctx context.Context, op *xxx_RunOperation) *xxx_RunOperation {
	if op == nil {
		op = &xxx_RunOperation{}
	}
	if o == nil {
		return op
	}
	o.Path = op.Path
	o.ArgsCount = op.ArgsCount
	o.Args = op.Args
	o.Flags = op.Flags
	o.SessionID = op.SessionID
	o.User = op.User
	return op
}

func (o *RunRequest) xxx_FromOp(ctx context.Context, op *xxx_RunOperation) {
	if o == nil {
		return
	}
	o.Path = op.Path
	o.ArgsCount = op.ArgsCount
	o.Args = op.Args
	o.Flags = op.Flags
	o.SessionID = op.SessionID
	o.User = op.User
}
func (o *RunRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RunRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RunOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RunResponse structure represents the SchRpcRun operation response
type RunResponse struct {
	// pGuid: MUST contain a GUID for the task instance created as result of this call.
	GUID *dtyp.GUID `idl:"name:pGuid" json:"guid"`
	// Return: The SchRpcRun return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RunResponse) xxx_ToOp(ctx context.Context, op *xxx_RunOperation) *xxx_RunOperation {
	if op == nil {
		op = &xxx_RunOperation{}
	}
	if o == nil {
		return op
	}
	o.GUID = op.GUID
	o.Return = op.Return
	return op
}

func (o *RunResponse) xxx_FromOp(ctx context.Context, op *xxx_RunOperation) {
	if o == nil {
		return
	}
	o.GUID = op.GUID
	o.Return = op.Return
}
func (o *RunResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RunResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RunOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteOperation structure represents the SchRpcDelete operation
type xxx_DeleteOperation struct {
	Path   string `idl:"name:path;string" json:"path"`
	Flags  uint32 `idl:"name:flags" json:"flags"`
	Return int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteOperation) OpNum() int { return 13 }

func (o *xxx_DeleteOperation) OpName() string { return "/ITaskSchedulerService/v1/SchRpcDelete" }

func (o *xxx_DeleteOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// path {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// path {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DeleteRequest structure represents the SchRpcDelete operation request
type DeleteRequest struct {
	// path: MUST be the full path to the task or folder to delete using the format specified
	// in section 2.3.11.
	Path string `idl:"name:path;string" json:"path"`
	// flags: Unused. The client MUST set to zero and the server MUST return an error if
	// nonzero.
	Flags uint32 `idl:"name:flags" json:"flags"`
}

func (o *DeleteRequest) xxx_ToOp(ctx context.Context, op *xxx_DeleteOperation) *xxx_DeleteOperation {
	if op == nil {
		op = &xxx_DeleteOperation{}
	}
	if o == nil {
		return op
	}
	o.Path = op.Path
	o.Flags = op.Flags
	return op
}

func (o *DeleteRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteOperation) {
	if o == nil {
		return
	}
	o.Path = op.Path
	o.Flags = op.Flags
}
func (o *DeleteRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeleteRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteResponse structure represents the SchRpcDelete operation response
type DeleteResponse struct {
	// Return: The SchRpcDelete return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteResponse) xxx_ToOp(ctx context.Context, op *xxx_DeleteOperation) *xxx_DeleteOperation {
	if op == nil {
		op = &xxx_DeleteOperation{}
	}
	if o == nil {
		return op
	}
	o.Return = op.Return
	return op
}

func (o *DeleteResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *DeleteResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeleteResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RenameOperation structure represents the SchRpcRename operation
type xxx_RenameOperation struct {
	Path    string `idl:"name:path;string" json:"path"`
	NewName string `idl:"name:newName;string" json:"new_name"`
	Flags   uint32 `idl:"name:flags" json:"flags"`
	Return  int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_RenameOperation) OpNum() int { return 14 }

func (o *xxx_RenameOperation) OpName() string { return "/ITaskSchedulerService/v1/SchRpcRename" }

func (o *xxx_RenameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RenameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// path {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
			return err
		}
	}
	// newName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.NewName); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RenameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// path {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
			return err
		}
	}
	// newName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.NewName); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RenameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RenameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RenameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RenameRequest structure represents the SchRpcRename operation request
type RenameRequest struct {
	// path: MUST be the full path to the task or to a folder to rename. The path MUST be
	// in the format specified in section 2.3.11.
	Path string `idl:"name:path;string" json:"path"`
	// newName: The new name of the task.
	NewName string `idl:"name:newName;string" json:"new_name"`
	// flags: Reserved. The client MUST set this parameter to zero.
	Flags uint32 `idl:"name:flags" json:"flags"`
}

func (o *RenameRequest) xxx_ToOp(ctx context.Context, op *xxx_RenameOperation) *xxx_RenameOperation {
	if op == nil {
		op = &xxx_RenameOperation{}
	}
	if o == nil {
		return op
	}
	o.Path = op.Path
	o.NewName = op.NewName
	o.Flags = op.Flags
	return op
}

func (o *RenameRequest) xxx_FromOp(ctx context.Context, op *xxx_RenameOperation) {
	if o == nil {
		return
	}
	o.Path = op.Path
	o.NewName = op.NewName
	o.Flags = op.Flags
}
func (o *RenameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RenameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RenameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RenameResponse structure represents the SchRpcRename operation response
type RenameResponse struct {
	// Return: The SchRpcRename return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RenameResponse) xxx_ToOp(ctx context.Context, op *xxx_RenameOperation) *xxx_RenameOperation {
	if op == nil {
		op = &xxx_RenameOperation{}
	}
	if o == nil {
		return op
	}
	o.Return = op.Return
	return op
}

func (o *RenameResponse) xxx_FromOp(ctx context.Context, op *xxx_RenameOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RenameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RenameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RenameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ScheduledRuntimesOperation structure represents the SchRpcScheduledRuntimes operation
type xxx_ScheduledRuntimesOperation struct {
	Path           string             `idl:"name:path;string" json:"path"`
	Start          *dtyp.SystemTime   `idl:"name:start;pointer:unique" json:"start"`
	End            *dtyp.SystemTime   `idl:"name:end;pointer:unique" json:"end"`
	Flags          uint32             `idl:"name:flags" json:"flags"`
	RequestedCount uint32             `idl:"name:cRequested" json:"requested_count"`
	RuntimesCount  uint32             `idl:"name:pcRuntimes" json:"runtimes_count"`
	Runtimes       []*dtyp.SystemTime `idl:"name:pRuntimes;size_is:(, pcRuntimes)" json:"runtimes"`
	Return         int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_ScheduledRuntimesOperation) OpNum() int { return 15 }

func (o *xxx_ScheduledRuntimesOperation) OpName() string {
	return "/ITaskSchedulerService/v1/SchRpcScheduledRuntimes"
}

func (o *xxx_ScheduledRuntimesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ScheduledRuntimesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// path {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
			return err
		}
	}
	// start {in} (1:{pointer=unique, alias=PSYSTEMTIME}*(1))(2:{alias=SYSTEMTIME}(struct))
	{
		if o.Start != nil {
			_ptr_start := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Start != nil {
					if err := o.Start.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.SystemTime{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Start, _ptr_start); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// end {in} (1:{pointer=unique, alias=PSYSTEMTIME}*(1))(2:{alias=SYSTEMTIME}(struct))
	{
		if o.End != nil {
			_ptr_end := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.End != nil {
					if err := o.End.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.SystemTime{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.End, _ptr_end); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// cRequested {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequestedCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ScheduledRuntimesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// path {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
			return err
		}
	}
	// start {in} (1:{pointer=unique, alias=PSYSTEMTIME}*(1))(2:{alias=SYSTEMTIME}(struct))
	{
		_ptr_start := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Start == nil {
				o.Start = &dtyp.SystemTime{}
			}
			if err := o.Start.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_start := func(ptr interface{}) { o.Start = *ptr.(**dtyp.SystemTime) }
		if err := w.ReadPointer(&o.Start, _s_start, _ptr_start); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// end {in} (1:{pointer=unique, alias=PSYSTEMTIME}*(1))(2:{alias=SYSTEMTIME}(struct))
	{
		_ptr_end := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.End == nil {
				o.End = &dtyp.SystemTime{}
			}
			if err := o.End.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_end := func(ptr interface{}) { o.End = *ptr.(**dtyp.SystemTime) }
		if err := w.ReadPointer(&o.End, _s_end, _ptr_end); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// cRequested {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequestedCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ScheduledRuntimesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Runtimes != nil && o.RuntimesCount == 0 {
		o.RuntimesCount = uint32(len(o.Runtimes))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ScheduledRuntimesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pcRuntimes {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RuntimesCount); err != nil {
			return err
		}
	}
	// pRuntimes {out} (1:{pointer=ref}*(2))(2:{alias=PSYSTEMTIME}*(1))(3:{alias=SYSTEMTIME}[dim:0,size_is=pcRuntimes](struct))
	{
		if o.Runtimes != nil || o.RuntimesCount > 0 {
			_ptr_pRuntimes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.RuntimesCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.Runtimes {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.Runtimes[i1] != nil {
						if err := o.Runtimes[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dtyp.SystemTime{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.Runtimes); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&dtyp.SystemTime{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Runtimes, _ptr_pRuntimes); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ScheduledRuntimesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pcRuntimes {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RuntimesCount); err != nil {
			return err
		}
	}
	// pRuntimes {out} (1:{pointer=ref}*(2))(2:{alias=PSYSTEMTIME,pointer=ref}*(1))(3:{alias=SYSTEMTIME}[dim:0,size_is=pcRuntimes](struct))
	{
		_ptr_pRuntimes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Runtimes", sizeInfo[0])
			}
			o.Runtimes = make([]*dtyp.SystemTime, sizeInfo[0])
			for i1 := range o.Runtimes {
				i1 := i1
				if o.Runtimes[i1] == nil {
					o.Runtimes[i1] = &dtyp.SystemTime{}
				}
				if err := o.Runtimes[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pRuntimes := func(ptr interface{}) { o.Runtimes = *ptr.(*[]*dtyp.SystemTime) }
		if err := w.ReadPointer(&o.Runtimes, _s_pRuntimes, _ptr_pRuntimes); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ScheduledRuntimesRequest structure represents the SchRpcScheduledRuntimes operation request
type ScheduledRuntimesRequest struct {
	// path: MUST contain the full path to a task using the format specified in section
	// 2.3.11.
	Path string `idl:"name:path;string" json:"path"`
	// start: If non-NULL, MUST specify the start of a time interval. If NULL, the server
	// MUST calculate scheduled runtimes from the start of time, where the start of time
	// is the smallest time value that the specific operating system implements.
	Start *dtyp.SystemTime `idl:"name:start;pointer:unique" json:"start"`
	// end: If non-NULL, MUST specify the end of a time interval. If NULL, the server MUST
	// calculate scheduled runtimes to the end of time, where the end of time is the largest
	// time value that the specific operating system implements.
	End *dtyp.SystemTime `idl:"name:end;pointer:unique" json:"end"`
	// flags: Unused. The client MUST specify 0, and the server MUST return an error if
	// nonzero.
	Flags uint32 `idl:"name:flags" json:"flags"`
	// cRequested: MUST contain the number of scheduled runtimes requested.
	RequestedCount uint32 `idl:"name:cRequested" json:"requested_count"`
}

func (o *ScheduledRuntimesRequest) xxx_ToOp(ctx context.Context, op *xxx_ScheduledRuntimesOperation) *xxx_ScheduledRuntimesOperation {
	if op == nil {
		op = &xxx_ScheduledRuntimesOperation{}
	}
	if o == nil {
		return op
	}
	o.Path = op.Path
	o.Start = op.Start
	o.End = op.End
	o.Flags = op.Flags
	o.RequestedCount = op.RequestedCount
	return op
}

func (o *ScheduledRuntimesRequest) xxx_FromOp(ctx context.Context, op *xxx_ScheduledRuntimesOperation) {
	if o == nil {
		return
	}
	o.Path = op.Path
	o.Start = op.Start
	o.End = op.End
	o.Flags = op.Flags
	o.RequestedCount = op.RequestedCount
}
func (o *ScheduledRuntimesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ScheduledRuntimesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ScheduledRuntimesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ScheduledRuntimesResponse structure represents the SchRpcScheduledRuntimes operation response
type ScheduledRuntimesResponse struct {
	// pcRuntimes: MUST contain the number of runtimes actually returned. The server MUST
	// NOT return more than cRequested runtimes.
	RuntimesCount uint32 `idl:"name:pcRuntimes" json:"runtimes_count"`
	// pRuntimes: MUST be a pointer to an array of scheduled runtimes. The server MUST return
	// the first pcRuntimes runtimes in the specified time interval.
	Runtimes []*dtyp.SystemTime `idl:"name:pRuntimes;size_is:(, pcRuntimes)" json:"runtimes"`
	// Return: The SchRpcScheduledRuntimes return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ScheduledRuntimesResponse) xxx_ToOp(ctx context.Context, op *xxx_ScheduledRuntimesOperation) *xxx_ScheduledRuntimesOperation {
	if op == nil {
		op = &xxx_ScheduledRuntimesOperation{}
	}
	if o == nil {
		return op
	}
	o.RuntimesCount = op.RuntimesCount
	o.Runtimes = op.Runtimes
	o.Return = op.Return
	return op
}

func (o *ScheduledRuntimesResponse) xxx_FromOp(ctx context.Context, op *xxx_ScheduledRuntimesOperation) {
	if o == nil {
		return
	}
	o.RuntimesCount = op.RuntimesCount
	o.Runtimes = op.Runtimes
	o.Return = op.Return
}
func (o *ScheduledRuntimesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ScheduledRuntimesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ScheduledRuntimesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetLastRunInfoOperation structure represents the SchRpcGetLastRunInfo operation
type xxx_GetLastRunInfoOperation struct {
	Path           string           `idl:"name:path;string" json:"path"`
	LastRuntime    *dtyp.SystemTime `idl:"name:pLastRuntime" json:"last_runtime"`
	LastReturnCode uint32           `idl:"name:pLastReturnCode" json:"last_return_code"`
	Return         int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetLastRunInfoOperation) OpNum() int { return 16 }

func (o *xxx_GetLastRunInfoOperation) OpName() string {
	return "/ITaskSchedulerService/v1/SchRpcGetLastRunInfo"
}

func (o *xxx_GetLastRunInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLastRunInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// path {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLastRunInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// path {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLastRunInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLastRunInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pLastRuntime {out} (1:{alias=PSYSTEMTIME}*(1))(2:{alias=SYSTEMTIME}(struct))
	{
		if o.LastRuntime != nil {
			if err := o.LastRuntime.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.SystemTime{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pLastReturnCode {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.LastReturnCode); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLastRunInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pLastRuntime {out} (1:{alias=PSYSTEMTIME,pointer=ref}*(1))(2:{alias=SYSTEMTIME}(struct))
	{
		if o.LastRuntime == nil {
			o.LastRuntime = &dtyp.SystemTime{}
		}
		if err := o.LastRuntime.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pLastReturnCode {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.LastReturnCode); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetLastRunInfoRequest structure represents the SchRpcGetLastRunInfo operation request
type GetLastRunInfoRequest struct {
	// path: MUST contain the full path to a task using the format specified in section
	// 2.3.11.
	Path string `idl:"name:path;string" json:"path"`
}

func (o *GetLastRunInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_GetLastRunInfoOperation) *xxx_GetLastRunInfoOperation {
	if op == nil {
		op = &xxx_GetLastRunInfoOperation{}
	}
	if o == nil {
		return op
	}
	o.Path = op.Path
	return op
}

func (o *GetLastRunInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetLastRunInfoOperation) {
	if o == nil {
		return
	}
	o.Path = op.Path
}
func (o *GetLastRunInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetLastRunInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLastRunInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetLastRunInfoResponse structure represents the SchRpcGetLastRunInfo operation response
type GetLastRunInfoResponse struct {
	// pLastRuntime: The server MUST return an error if this parameter is NULL. The server
	// MUST return the time when the task last started running, or zero if the task has
	// never started.
	LastRuntime *dtyp.SystemTime `idl:"name:pLastRuntime" json:"last_runtime"`
	// pLastReturnCode: The server MUST return an error if this parameter is NULL. The server
	// MUST return the exit code from the task's last execution, or zero if the task has
	// never finished execution.
	LastReturnCode uint32 `idl:"name:pLastReturnCode" json:"last_return_code"`
	// Return: The SchRpcGetLastRunInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetLastRunInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_GetLastRunInfoOperation) *xxx_GetLastRunInfoOperation {
	if op == nil {
		op = &xxx_GetLastRunInfoOperation{}
	}
	if o == nil {
		return op
	}
	o.LastRuntime = op.LastRuntime
	o.LastReturnCode = op.LastReturnCode
	o.Return = op.Return
	return op
}

func (o *GetLastRunInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetLastRunInfoOperation) {
	if o == nil {
		return
	}
	o.LastRuntime = op.LastRuntime
	o.LastReturnCode = op.LastReturnCode
	o.Return = op.Return
}
func (o *GetLastRunInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetLastRunInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLastRunInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetTaskInfoOperation structure represents the SchRpcGetTaskInfo operation
type xxx_GetTaskInfoOperation struct {
	Path    string `idl:"name:path;string" json:"path"`
	Flags   uint32 `idl:"name:flags" json:"flags"`
	Enabled uint32 `idl:"name:pEnabled" json:"enabled"`
	State   uint32 `idl:"name:pState" json:"state"`
	Return  int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTaskInfoOperation) OpNum() int { return 17 }

func (o *xxx_GetTaskInfoOperation) OpName() string {
	return "/ITaskSchedulerService/v1/SchRpcGetTaskInfo"
}

func (o *xxx_GetTaskInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTaskInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// path {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTaskInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// path {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTaskInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTaskInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pEnabled {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Enabled); err != nil {
			return err
		}
	}
	// pState {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.State); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTaskInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pEnabled {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Enabled); err != nil {
			return err
		}
	}
	// pState {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.State); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetTaskInfoRequest structure represents the SchRpcGetTaskInfo operation request
type GetTaskInfoRequest struct {
	// path: MUST contain the full path to a task using the format specified in section
	// 2.3.11.
	Path string `idl:"name:path;string" json:"path"`
	// flags: The flags field MUST contain individual bit flags that MUST have one or more
	// of the following values
	//
	//	+---+---+---+-----+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 1 | 2 |  3  | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
	//	|   |   |   |     |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
	//	+---+---+---+-----+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	+---+---+---+-----+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 0 | 0 | F S | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
	//	+---+---+---+-----+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//
	//
	//	+-------------------+-------------------------------------------------------+
	//	|                   |                                                       |
	//	|       VALUE       |                      DESCRIPTION                      |
	//	|                   |                                                       |
	//	+-------------------+-------------------------------------------------------+
	//	+-------------------+-------------------------------------------------------+
	//	| FS SCH_FLAG_STATE | If set to 1, the server MUST retrieve the TASK_STATE. |
	//	+-------------------+-------------------------------------------------------+
	Flags uint32 `idl:"name:flags" json:"flags"`
}

func (o *GetTaskInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_GetTaskInfoOperation) *xxx_GetTaskInfoOperation {
	if op == nil {
		op = &xxx_GetTaskInfoOperation{}
	}
	if o == nil {
		return op
	}
	o.Path = op.Path
	o.Flags = op.Flags
	return op
}

func (o *GetTaskInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTaskInfoOperation) {
	if o == nil {
		return
	}
	o.Path = op.Path
	o.Flags = op.Flags
}
func (o *GetTaskInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetTaskInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTaskInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetTaskInfoResponse structure represents the SchRpcGetTaskInfo operation response
type GetTaskInfoResponse struct {
	// pEnabled: MUST be a pointer to a Boolean that indicates whether the task is currently
	// enabled.
	Enabled uint32 `idl:"name:pEnabled" json:"enabled"`
	// pState: If non-NULL and the SCH_FLAG_STATE bit in the flags parameter is set, the
	// server MUST return the current state of the task in the format specified in section
	// 2.3.13. Otherwise, the server MUST ignore this parameter.
	State uint32 `idl:"name:pState" json:"state"`
	// Return: The SchRpcGetTaskInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTaskInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_GetTaskInfoOperation) *xxx_GetTaskInfoOperation {
	if op == nil {
		op = &xxx_GetTaskInfoOperation{}
	}
	if o == nil {
		return op
	}
	o.Enabled = op.Enabled
	o.State = op.State
	o.Return = op.Return
	return op
}

func (o *GetTaskInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTaskInfoOperation) {
	if o == nil {
		return
	}
	o.Enabled = op.Enabled
	o.State = op.State
	o.Return = op.Return
}
func (o *GetTaskInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetTaskInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTaskInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNumberOfMissedRunsOperation structure represents the SchRpcGetNumberOfMissedRuns operation
type xxx_GetNumberOfMissedRunsOperation struct {
	Path               string `idl:"name:path;string" json:"path"`
	NumberOfMissedRuns uint32 `idl:"name:pNumberOfMissedRuns" json:"number_of_missed_runs"`
	Return             int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNumberOfMissedRunsOperation) OpNum() int { return 18 }

func (o *xxx_GetNumberOfMissedRunsOperation) OpName() string {
	return "/ITaskSchedulerService/v1/SchRpcGetNumberOfMissedRuns"
}

func (o *xxx_GetNumberOfMissedRunsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNumberOfMissedRunsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// path {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNumberOfMissedRunsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// path {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNumberOfMissedRunsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNumberOfMissedRunsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pNumberOfMissedRuns {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NumberOfMissedRuns); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNumberOfMissedRunsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pNumberOfMissedRuns {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NumberOfMissedRuns); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetNumberOfMissedRunsRequest structure represents the SchRpcGetNumberOfMissedRuns operation request
type GetNumberOfMissedRunsRequest struct {
	// path: MUST contain the full path to a task, in the format specified in section 2.3.11.
	Path string `idl:"name:path;string" json:"path"`
}

func (o *GetNumberOfMissedRunsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNumberOfMissedRunsOperation) *xxx_GetNumberOfMissedRunsOperation {
	if op == nil {
		op = &xxx_GetNumberOfMissedRunsOperation{}
	}
	if o == nil {
		return op
	}
	o.Path = op.Path
	return op
}

func (o *GetNumberOfMissedRunsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNumberOfMissedRunsOperation) {
	if o == nil {
		return
	}
	o.Path = op.Path
}
func (o *GetNumberOfMissedRunsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNumberOfMissedRunsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNumberOfMissedRunsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNumberOfMissedRunsResponse structure represents the SchRpcGetNumberOfMissedRuns operation response
type GetNumberOfMissedRunsResponse struct {
	// pNumberOfMissedRuns: MUST be the address of a DWORD that receives the number of times
	// a task was scheduled to run but did not.
	NumberOfMissedRuns uint32 `idl:"name:pNumberOfMissedRuns" json:"number_of_missed_runs"`
	// Return: The SchRpcGetNumberOfMissedRuns return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNumberOfMissedRunsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNumberOfMissedRunsOperation) *xxx_GetNumberOfMissedRunsOperation {
	if op == nil {
		op = &xxx_GetNumberOfMissedRunsOperation{}
	}
	if o == nil {
		return op
	}
	o.NumberOfMissedRuns = op.NumberOfMissedRuns
	o.Return = op.Return
	return op
}

func (o *GetNumberOfMissedRunsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNumberOfMissedRunsOperation) {
	if o == nil {
		return
	}
	o.NumberOfMissedRuns = op.NumberOfMissedRuns
	o.Return = op.Return
}
func (o *GetNumberOfMissedRunsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNumberOfMissedRunsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNumberOfMissedRunsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnableTaskOperation structure represents the SchRpcEnableTask operation
type xxx_EnableTaskOperation struct {
	Path    string `idl:"name:path;string" json:"path"`
	Enabled uint32 `idl:"name:enabled" json:"enabled"`
	Return  int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_EnableTaskOperation) OpNum() int { return 19 }

func (o *xxx_EnableTaskOperation) OpName() string {
	return "/ITaskSchedulerService/v1/SchRpcEnableTask"
}

func (o *xxx_EnableTaskOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnableTaskOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// path {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
			return err
		}
	}
	// enabled {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Enabled); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnableTaskOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// path {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
			return err
		}
	}
	// enabled {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Enabled); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnableTaskOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnableTaskOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnableTaskOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnableTaskRequest structure represents the SchRpcEnableTask operation request
type EnableTaskRequest struct {
	// path: MUST contain the full path to the task, in the format specified in section
	// 2.3.11.
	Path string `idl:"name:path;string" json:"path"`
	// enabled: If TRUE, the server MUST enable the task. Otherwise, the server MUST disable
	// the task.
	Enabled uint32 `idl:"name:enabled" json:"enabled"`
}

func (o *EnableTaskRequest) xxx_ToOp(ctx context.Context, op *xxx_EnableTaskOperation) *xxx_EnableTaskOperation {
	if op == nil {
		op = &xxx_EnableTaskOperation{}
	}
	if o == nil {
		return op
	}
	o.Path = op.Path
	o.Enabled = op.Enabled
	return op
}

func (o *EnableTaskRequest) xxx_FromOp(ctx context.Context, op *xxx_EnableTaskOperation) {
	if o == nil {
		return
	}
	o.Path = op.Path
	o.Enabled = op.Enabled
}
func (o *EnableTaskRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnableTaskRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnableTaskOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnableTaskResponse structure represents the SchRpcEnableTask operation response
type EnableTaskResponse struct {
	// Return: The SchRpcEnableTask return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnableTaskResponse) xxx_ToOp(ctx context.Context, op *xxx_EnableTaskOperation) *xxx_EnableTaskOperation {
	if op == nil {
		op = &xxx_EnableTaskOperation{}
	}
	if o == nil {
		return op
	}
	o.Return = op.Return
	return op
}

func (o *EnableTaskResponse) xxx_FromOp(ctx context.Context, op *xxx_EnableTaskOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *EnableTaskResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnableTaskResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnableTaskOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
