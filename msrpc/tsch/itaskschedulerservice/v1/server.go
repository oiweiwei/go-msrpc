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
		op := &xxx_HighestVersionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &HighestVersionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.HighestVersion(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // SchRpcRegisterTask
		op := &xxx_RegisterTaskOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RegisterTaskRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RegisterTask(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // SchRpcRetrieveTask
		op := &xxx_RetrieveTaskOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RetrieveTaskRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RetrieveTask(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // SchRpcCreateFolder
		op := &xxx_CreateFolderOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateFolderRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateFolder(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // SchRpcSetSecurity
		op := &xxx_SetSecurityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSecurityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSecurity(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // SchRpcGetSecurity
		op := &xxx_GetSecurityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSecurityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSecurity(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // SchRpcEnumFolders
		op := &xxx_EnumFoldersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumFoldersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumFolders(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // SchRpcEnumTasks
		op := &xxx_EnumTasksOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumTasksRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumTasks(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // SchRpcEnumInstances
		op := &xxx_EnumInstancesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumInstancesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumInstances(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // SchRpcGetInstanceInfo
		op := &xxx_GetInstanceInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetInstanceInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetInstanceInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // SchRpcStopInstance
		op := &xxx_StopInstanceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StopInstanceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.StopInstance(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // SchRpcStop
		op := &xxx_StopOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StopRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Stop(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // SchRpcRun
		op := &xxx_RunOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RunRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Run(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // SchRpcDelete
		op := &xxx_DeleteOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Delete(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // SchRpcRename
		op := &xxx_RenameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RenameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Rename(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // SchRpcScheduledRuntimes
		op := &xxx_ScheduledRuntimesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ScheduledRuntimesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ScheduledRuntimes(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // SchRpcGetLastRunInfo
		op := &xxx_GetLastRunInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLastRunInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLastRunInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // SchRpcGetTaskInfo
		op := &xxx_GetTaskInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTaskInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTaskInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // SchRpcGetNumberOfMissedRuns
		op := &xxx_GetNumberOfMissedRunsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNumberOfMissedRunsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNumberOfMissedRuns(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // SchRpcEnableTask
		op := &xxx_EnableTaskOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnableTaskRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnableTask(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented ITaskSchedulerService
type UnimplementedTaskSchedulerServiceServer struct {
}

func (UnimplementedTaskSchedulerServiceServer) HighestVersion(context.Context, *HighestVersionRequest) (*HighestVersionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTaskSchedulerServiceServer) RegisterTask(context.Context, *RegisterTaskRequest) (*RegisterTaskResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTaskSchedulerServiceServer) RetrieveTask(context.Context, *RetrieveTaskRequest) (*RetrieveTaskResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTaskSchedulerServiceServer) CreateFolder(context.Context, *CreateFolderRequest) (*CreateFolderResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTaskSchedulerServiceServer) SetSecurity(context.Context, *SetSecurityRequest) (*SetSecurityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTaskSchedulerServiceServer) GetSecurity(context.Context, *GetSecurityRequest) (*GetSecurityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTaskSchedulerServiceServer) EnumFolders(context.Context, *EnumFoldersRequest) (*EnumFoldersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTaskSchedulerServiceServer) EnumTasks(context.Context, *EnumTasksRequest) (*EnumTasksResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTaskSchedulerServiceServer) EnumInstances(context.Context, *EnumInstancesRequest) (*EnumInstancesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTaskSchedulerServiceServer) GetInstanceInfo(context.Context, *GetInstanceInfoRequest) (*GetInstanceInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTaskSchedulerServiceServer) StopInstance(context.Context, *StopInstanceRequest) (*StopInstanceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTaskSchedulerServiceServer) Stop(context.Context, *StopRequest) (*StopResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTaskSchedulerServiceServer) Run(context.Context, *RunRequest) (*RunResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTaskSchedulerServiceServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTaskSchedulerServiceServer) Rename(context.Context, *RenameRequest) (*RenameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTaskSchedulerServiceServer) ScheduledRuntimes(context.Context, *ScheduledRuntimesRequest) (*ScheduledRuntimesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTaskSchedulerServiceServer) GetLastRunInfo(context.Context, *GetLastRunInfoRequest) (*GetLastRunInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTaskSchedulerServiceServer) GetTaskInfo(context.Context, *GetTaskInfoRequest) (*GetTaskInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTaskSchedulerServiceServer) GetNumberOfMissedRuns(context.Context, *GetNumberOfMissedRunsRequest) (*GetNumberOfMissedRunsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTaskSchedulerServiceServer) EnableTask(context.Context, *EnableTaskRequest) (*EnableTaskResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ TaskSchedulerServiceServer = (*UnimplementedTaskSchedulerServiceServer)(nil)
