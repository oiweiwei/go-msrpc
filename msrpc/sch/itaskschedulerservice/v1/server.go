package itaskschedulerservice

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

// ITaskSchedulerService server interface.
type TaskSchedulerServiceServer interface {

	// The SchRpcHighestVersion method MUST return the highest version of the Task Scheduler
	// Remoting Protocol that is supported by the server.<65>
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	HighestVersion(context.Context, *HighestVersionRequest) (*HighestVersionResponse, error)

	// The SchRpcRegisterTask method MUST register a task with the server.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in section 2 of [MS-ERREF].
	RegisterTask(context.Context, *RegisterTaskRequest) (*RegisterTaskResponse, error)

	// The SchRpcRetrieveTask method MUST return a task definition.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	RetrieveTask(context.Context, *RetrieveTaskRequest) (*RetrieveTaskResponse, error)

	// The SchRpcCreateFolder method creates a new folder.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	CreateFolder(context.Context, *CreateFolderRequest) (*CreateFolderResponse, error)

	// The SchRpcSetSecurity method MUST set a security descriptor on a task or folder.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	SetSecurity(context.Context, *SetSecurityRequest) (*SetSecurityResponse, error)

	// The SchRpcGetSecurity method MUST get the security descriptor associated with a task
	// or folder.
	//
	// HRESULT SchRpcGetSecurity( [in, string] const wchar_t* path,
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	GetSecurity(context.Context, *GetSecurityRequest) (*GetSecurityResponse, error)

	// The SchRpcEnumFolders method MUST retrieve a list of folders on the server.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	EnumFolders(context.Context, *EnumFoldersRequest) (*EnumFoldersResponse, error)

	// The SchRpcEnumTasks method MUST return the list of tasks in a specific folder.
	//
	// MUST returns S_FALSE if there are more tasks to enumerate.
	//
	// Note that if client requests items 1-10 and then 11-20, the second call can return
	// duplicate entries if the task list has changed in between calls.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	EnumTasks(context.Context, *EnumTasksRequest) (*EnumTasksResponse, error)

	// The SchRpcEnumInstances method MUST return a task's list of instances that are currently
	// running.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	EnumInstances(context.Context, *EnumInstancesRequest) (*EnumInstancesResponse, error)

	// The SchRpcGetInstanceInfo method MUST get information about an instance of a running
	// task.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	GetInstanceInfo(context.Context, *GetInstanceInfoRequest) (*GetInstanceInfoResponse, error)

	// The SchRpcStopInstance method MUST stop a specified instance of a task.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	StopInstance(context.Context, *StopInstanceRequest) (*StopInstanceResponse, error)

	// The SchRpcStop MUST stop all currently running instances of a task specified by a
	// path.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	Stop(context.Context, *StopRequest) (*StopResponse, error)

	// The SchRpcRun method MUST run a task specified by a path.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	Run(context.Context, *RunRequest) (*RunResponse, error)

	// The SchRpcDelete method MUST delete a task or folder in the task store.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)

	// The SchRpcRename method renames a folder or task.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	Rename(context.Context, *RenameRequest) (*RenameResponse, error)

	// The SchRpcScheduledRuntimes method MUST return scheduled run times.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	ScheduledRuntimes(context.Context, *ScheduledRuntimesRequest) (*ScheduledRuntimesResponse, error)

	// The SchRpcGetLastRunInfo method MUST return information about the task's last run.
	//
	// Return Values: For more information about return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	GetLastRunInfo(context.Context, *GetLastRunInfoRequest) (*GetLastRunInfoResponse, error)

	// The SchRpcGetTaskInfo method MUST return information about a specified task.
	//
	// Return Values: For more information about return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	GetTaskInfo(context.Context, *GetTaskInfoRequest) (*GetTaskInfoResponse, error)

	// The SchRpcGetNumberOfMissedRuns MUST return the number of times a task was scheduled
	// to run but did not due to the server being unavailable (for example, powered off).
	//
	// Return Values: For more information about return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	GetNumberOfMissedRuns(context.Context, *GetNumberOfMissedRunsRequest) (*GetNumberOfMissedRunsResponse, error)

	// The SchRpcEnableTask method MUST enable or disable a task.
	//
	// Return Values: For more information about return codes, see [MS-ERREF] section 2.
	EnableTask(context.Context, *EnableTaskRequest) (*EnableTaskResponse, error)
}

func RegisterTaskSchedulerServiceServer(conn dcerpc.Conn, o TaskSchedulerServiceServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewTaskSchedulerServiceServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(TaskSchedulerServiceSyntaxV1_0))...)
}

func NewTaskSchedulerServiceServerHandle(o TaskSchedulerServiceServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return TaskSchedulerServiceServerHandle(ctx, o, opNum, r)
	}
}

func TaskSchedulerServiceServerHandle(ctx context.Context, o TaskSchedulerServiceServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // SchRpcHighestVersion
		in := &HighestVersionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.HighestVersion(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 1: // SchRpcRegisterTask
		in := &RegisterTaskRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RegisterTask(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 2: // SchRpcRetrieveTask
		in := &RetrieveTaskRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RetrieveTask(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 3: // SchRpcCreateFolder
		in := &CreateFolderRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateFolder(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // SchRpcSetSecurity
		in := &SetSecurityRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSecurity(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // SchRpcGetSecurity
		in := &GetSecurityRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSecurity(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // SchRpcEnumFolders
		in := &EnumFoldersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumFolders(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // SchRpcEnumTasks
		in := &EnumTasksRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumTasks(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // SchRpcEnumInstances
		in := &EnumInstancesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumInstances(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // SchRpcGetInstanceInfo
		in := &GetInstanceInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetInstanceInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // SchRpcStopInstance
		in := &StopInstanceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.StopInstance(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // SchRpcStop
		in := &StopRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Stop(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // SchRpcRun
		in := &RunRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Run(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // SchRpcDelete
		in := &DeleteRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Delete(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // SchRpcRename
		in := &RenameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Rename(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // SchRpcScheduledRuntimes
		in := &ScheduledRuntimesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ScheduledRuntimes(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // SchRpcGetLastRunInfo
		in := &GetLastRunInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetLastRunInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // SchRpcGetTaskInfo
		in := &GetTaskInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTaskInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // SchRpcGetNumberOfMissedRuns
		in := &GetNumberOfMissedRunsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetNumberOfMissedRuns(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // SchRpcEnableTask
		in := &EnableTaskRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnableTask(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
