package ifsrmreportscheduler

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
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
	_ = idispatch.GoPackage
)

// IFsrmReportScheduler server interface.
type ReportSchedulerServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The VerifyNamespaces method checks that all namespaces passed in exist and are valid
	// paths.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                  |
	//	|              VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045304 FSRM_E_PATH_NOT_FOUND      | The specified path could not be found.                                           |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045306 FSRM_E_INVALID_PATH        | The supplied path is not valid.                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045315 FSRM_E_VOLUME_UNSUPPORTED  | The volume is not an NTFS volume.                                                |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8004531F FSRM_E_FILE_SYSTEM_CORRUPT | The file system is corrupt.                                                      |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG               | This code is returned for the following reasons: The namespacesSafeArray         |
	//	|                                       | parameter is NULL. The namespacesSafeArray parameter does not contain a variant  |
	//	|                                       | array of BSTRs.                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	VerifyNamespaces(context.Context, *VerifyNamespacesRequest) (*VerifyNamespacesResponse, error)

	// The CreateScheduleTask method creates a task that is capable of triggering the generation
	// of a File Server Resource Manager Protocol report job.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN              |                                                                                  |
	//	|            VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                  |                                                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045303 FSRM_E_ALREADY_EXISTS | The object already exists.                                                       |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045311 FSRM_E_NOT_SUPPORTED  | This function is not supported for this object.                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045316 FSRM_E_UNEXPECTED     | An unexpected error occurred; check the application event log. You might get     |
	//	|                                  | this error if the XML is malformed.                                              |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG          | This code is returned for the following reasons: The taskName parameter is NULL. |
	//	|                                  | The namespacesSafeArray parameter is NULL. The serializedTask parameter is NULL. |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	CreateScheduleTask(context.Context, *CreateScheduleTaskRequest) (*CreateScheduleTaskResponse, error)

	// The ModifyScheduleTask method modifies the specified task that is capable of triggering
	// the generation of a File Server Resource Manager Protocol report job.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	|           RETURN            |                                                                                  |
	//	|         VALUE/CODE          |                                   DESCRIPTION                                    |
	//	|                             |                                                                                  |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND | The specified task could not be found.                                           |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG     | This code is returned for the following reasons: The taskName parameter is NULL. |
	//	|                             | The serializedTask parameter is NULL. The taskName parameter is not a valid      |
	//	|                             | value.                                                                           |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	ModifyScheduleTask(context.Context, *ModifyScheduleTaskRequest) (*ModifyScheduleTaskResponse, error)

	// The DeleteScheduleTask method deletes the specified task so that it no longer triggers
	// the generation of a File Server Resource Manager Protocol report job.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-----------------------------+----------------------------------------+
	//	|           RETURN            |                                        |
	//	|         VALUE/CODE          |              DESCRIPTION               |
	//	|                             |                                        |
	//	+-----------------------------+----------------------------------------+
	//	+-----------------------------+----------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND | The specified task could not be found. |
	//	+-----------------------------+----------------------------------------+
	//	| 0x80070057 E_INVALIDARG     | The taskName parameter is NULL.        |
	//	+-----------------------------+----------------------------------------+
	DeleteScheduleTask(context.Context, *DeleteScheduleTaskRequest) (*DeleteScheduleTaskResponse, error)
}

func RegisterReportSchedulerServer(conn dcerpc.Conn, o ReportSchedulerServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewReportSchedulerServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ReportSchedulerSyntaxV1_0))...)
}

func NewReportSchedulerServerHandle(o ReportSchedulerServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ReportSchedulerServerHandle(ctx, o, opNum, r)
	}
}

func ReportSchedulerServerHandle(ctx context.Context, o ReportSchedulerServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // VerifyNamespaces
		op := &xxx_VerifyNamespacesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &VerifyNamespacesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.VerifyNamespaces(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // CreateScheduleTask
		op := &xxx_CreateScheduleTaskOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateScheduleTaskRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateScheduleTask(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // ModifyScheduleTask
		op := &xxx_ModifyScheduleTaskOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ModifyScheduleTaskRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ModifyScheduleTask(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // DeleteScheduleTask
		op := &xxx_DeleteScheduleTaskOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteScheduleTaskRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteScheduleTask(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmReportScheduler
type UnimplementedReportSchedulerServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedReportSchedulerServer) VerifyNamespaces(context.Context, *VerifyNamespacesRequest) (*VerifyNamespacesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportSchedulerServer) CreateScheduleTask(context.Context, *CreateScheduleTaskRequest) (*CreateScheduleTaskResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportSchedulerServer) ModifyScheduleTask(context.Context, *ModifyScheduleTaskRequest) (*ModifyScheduleTaskResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportSchedulerServer) DeleteScheduleTask(context.Context, *DeleteScheduleTaskRequest) (*DeleteScheduleTaskResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ReportSchedulerServer = (*UnimplementedReportSchedulerServer)(nil)
