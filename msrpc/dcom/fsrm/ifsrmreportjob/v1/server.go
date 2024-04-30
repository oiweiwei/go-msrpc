package ifsrmreportjob

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
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
		in := &GetTaskRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTask(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // Task
		in := &SetTaskRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetTask(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // NamespaceRoots
		in := &GetNamespaceRootsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetNamespaceRoots(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // NamespaceRoots
		in := &SetNamespaceRootsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetNamespaceRoots(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // Formats
		in := &GetFormatsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFormats(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // Formats
		in := &SetFormatsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetFormats(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // MailTo
		in := &GetMailToRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMailTo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // MailTo
		in := &SetMailToRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetMailTo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // RunningStatus
		in := &GetRunningStatusRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetRunningStatus(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // LastRun
		in := &GetLastRunRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetLastRun(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 22: // LastError
		in := &GetLastErrorRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetLastError(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // LastGeneratedInDirectory
		in := &GetLastGeneratedInDirectoryRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetLastGeneratedInDirectory(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // EnumReports
		in := &EnumReportsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumReports(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // CreateReport
		in := &CreateReportRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateReport(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 26: // Run
		in := &RunRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Run(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 27: // WaitForCompletion
		in := &WaitForCompletionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.WaitForCompletion(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 28: // Cancel
		in := &CancelRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Cancel(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
