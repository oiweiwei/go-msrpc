package ifsrmreportjob

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	ifsrmobject "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmobject/v0"
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
	_ = ifsrmobject.GoPackage
)

// IFsrmReportJob server interface.
type ReportJobServer interface {

	// IFsrmObject base class.
	ifsrmobject.ObjectServer

	// Task operation.
	GetTask(context.Context, *GetTaskRequest) (*GetTaskResponse, error)

	// Task operation.
	SetTask(context.Context, *SetTaskRequest) (*SetTaskResponse, error)

	// NamespaceRoots operation.
	GetNamespaceRoots(context.Context, *GetNamespaceRootsRequest) (*GetNamespaceRootsResponse, error)

	// NamespaceRoots operation.
	SetNamespaceRoots(context.Context, *SetNamespaceRootsRequest) (*SetNamespaceRootsResponse, error)

	// Formats operation.
	GetFormats(context.Context, *GetFormatsRequest) (*GetFormatsResponse, error)

	// Formats operation.
	SetFormats(context.Context, *SetFormatsRequest) (*SetFormatsResponse, error)

	// MailTo operation.
	GetMailTo(context.Context, *GetMailToRequest) (*GetMailToResponse, error)

	// MailTo operation.
	SetMailTo(context.Context, *SetMailToRequest) (*SetMailToResponse, error)

	// RunningStatus operation.
	GetRunningStatus(context.Context, *GetRunningStatusRequest) (*GetRunningStatusResponse, error)

	// LastRun operation.
	GetLastRun(context.Context, *GetLastRunRequest) (*GetLastRunResponse, error)

	// LastError operation.
	GetLastError(context.Context, *GetLastErrorRequest) (*GetLastErrorResponse, error)

	// The LastGeneratedInDirectory (get) retrieves the last generated in directory for
	// the report job and returns S_OK upon successful completion.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+-----------------------------+
	//	|         RETURN          |                             |
	//	|       VALUE/CODE        |         DESCRIPTION         |
	//	|                         |                             |
	//	+-------------------------+-----------------------------+
	//	+-------------------------+-----------------------------+
	//	| 0x80070057 E_INVALIDARG | The path parameter is NULL. |
	//	+-------------------------+-----------------------------+
	GetLastGeneratedInDirectory(context.Context, *GetLastGeneratedInDirectoryRequest) (*GetLastGeneratedInDirectoryResponse, error)

	// The EnumReports method enumerates all the reports configured for the report job and
	// returns S_OK upon successful completion.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+--------------------------------+
	//	|         RETURN          |                                |
	//	|       VALUE/CODE        |          DESCRIPTION           |
	//	|                         |                                |
	//	+-------------------------+--------------------------------+
	//	+-------------------------+--------------------------------+
	//	| 0x80070057 E_INVALIDARG | The reports parameter is NULL. |
	//	+-------------------------+--------------------------------+
	EnumReports(context.Context, *EnumReportsRequest) (*EnumReportsResponse, error)

	// The CreateReport method adds a report of the specified type to a report job object.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                    RETURN                    |                                                                                  |
	//	|                  VALUE/CODE                  |                                   DESCRIPTION                                    |
	//	|                                              |                                                                                  |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045332 FSRM_E_REPORT_TYPE_ALREADY_EXISTS | A report of the specified type already exists for the report job.                |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                      | The reportType parameter is not a valid value. If reportType is                  |
	//	|                                              | FsrmReportType_AutomaticClassification or FsrmReportType_Expiration, the         |
	//	|                                              | parameter MUST be considered invalid. <61>                                       |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	CreateReport(context.Context, *CreateReportRequest) (*CreateReportResponse, error)

	// Run operation.
	Run(context.Context, *RunRequest) (*RunResponse, error)

	// WaitForCompletion operation.
	WaitForCompletion(context.Context, *WaitForCompletionRequest) (*WaitForCompletionResponse, error)

	// Cancel operation.
	Cancel(context.Context, *CancelRequest) (*CancelResponse, error)
}

func RegisterReportJobServer(conn dcerpc.Conn, o ReportJobServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewReportJobServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ReportJobSyntaxV1_0))...)
}

func NewReportJobServerHandle(o ReportJobServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ReportJobServerHandle(ctx, o, opNum, r)
	}
}

func ReportJobServerHandle(ctx context.Context, o ReportJobServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 12 {
		// IFsrmObject base method.
		return ifsrmobject.ObjectServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 12: // Task
		op := &xxx_GetTaskOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTaskRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTask(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // Task
		op := &xxx_SetTaskOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetTaskRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetTask(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // NamespaceRoots
		op := &xxx_GetNamespaceRootsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNamespaceRootsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNamespaceRoots(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // NamespaceRoots
		op := &xxx_SetNamespaceRootsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetNamespaceRootsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetNamespaceRoots(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // Formats
		op := &xxx_GetFormatsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFormatsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFormats(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // Formats
		op := &xxx_SetFormatsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFormatsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFormats(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // MailTo
		op := &xxx_GetMailToOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMailToRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMailTo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // MailTo
		op := &xxx_SetMailToOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMailToRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMailTo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // RunningStatus
		op := &xxx_GetRunningStatusOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRunningStatusRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRunningStatus(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // LastRun
		op := &xxx_GetLastRunOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLastRunRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLastRun(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // LastError
		op := &xxx_GetLastErrorOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLastErrorRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLastError(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // LastGeneratedInDirectory
		op := &xxx_GetLastGeneratedInDirectoryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLastGeneratedInDirectoryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLastGeneratedInDirectory(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // EnumReports
		op := &xxx_EnumReportsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumReportsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumReports(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // CreateReport
		op := &xxx_CreateReportOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateReportRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateReport(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 26: // Run
		op := &xxx_RunOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RunRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Run(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // WaitForCompletion
		op := &xxx_WaitForCompletionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &WaitForCompletionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.WaitForCompletion(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 28: // Cancel
		op := &xxx_CancelOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CancelRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Cancel(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmReportJob
type UnimplementedReportJobServer struct {
	ifsrmobject.UnimplementedObjectServer
}

func (UnimplementedReportJobServer) GetTask(context.Context, *GetTaskRequest) (*GetTaskResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportJobServer) SetTask(context.Context, *SetTaskRequest) (*SetTaskResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportJobServer) GetNamespaceRoots(context.Context, *GetNamespaceRootsRequest) (*GetNamespaceRootsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportJobServer) SetNamespaceRoots(context.Context, *SetNamespaceRootsRequest) (*SetNamespaceRootsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportJobServer) GetFormats(context.Context, *GetFormatsRequest) (*GetFormatsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportJobServer) SetFormats(context.Context, *SetFormatsRequest) (*SetFormatsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportJobServer) GetMailTo(context.Context, *GetMailToRequest) (*GetMailToResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportJobServer) SetMailTo(context.Context, *SetMailToRequest) (*SetMailToResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportJobServer) GetRunningStatus(context.Context, *GetRunningStatusRequest) (*GetRunningStatusResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportJobServer) GetLastRun(context.Context, *GetLastRunRequest) (*GetLastRunResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportJobServer) GetLastError(context.Context, *GetLastErrorRequest) (*GetLastErrorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportJobServer) GetLastGeneratedInDirectory(context.Context, *GetLastGeneratedInDirectoryRequest) (*GetLastGeneratedInDirectoryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportJobServer) EnumReports(context.Context, *EnumReportsRequest) (*EnumReportsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportJobServer) CreateReport(context.Context, *CreateReportRequest) (*CreateReportResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportJobServer) Run(context.Context, *RunRequest) (*RunResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportJobServer) WaitForCompletion(context.Context, *WaitForCompletionRequest) (*WaitForCompletionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportJobServer) Cancel(context.Context, *CancelRequest) (*CancelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ReportJobServer = (*UnimplementedReportJobServer)(nil)
