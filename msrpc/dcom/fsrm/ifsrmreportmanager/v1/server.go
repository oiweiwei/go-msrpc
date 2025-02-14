package ifsrmreportmanager

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
	_ = idispatch.GoPackage
)

// IFsrmReportManager server interface.
type ReportManagerServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The EnumReportJobs method returns a collection of all the report jobs from the List
	// of Persisted Report Jobs (section 3.2.1.5) and those from the List of Non-Persisted
	// Report Jobs (section 3.2.1.5) on the server.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG         | This return code is returned for the following reasons: The reportJobs parameter |
	//	|                                 | is NULL.                                                                         |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045311 FSRM_E_NOT_SUPPORTED | The options parameter contains FsrmEnumOptions (section 2.2.1.2.5) values that   |
	//	|                                 | are not valid.                                                                   |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	EnumReportJobs(context.Context, *EnumReportJobsRequest) (*EnumReportJobsResponse, error)

	// The CreateReportJob method creates a Non-Persisted Report Job Instance (section 3.2.1.5.1.2)
	// and returns S_OK upon successful completion.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+----------------------------------+
	//	|         RETURN          |                                  |
	//	|       VALUE/CODE        |           DESCRIPTION            |
	//	|                         |                                  |
	//	+-------------------------+----------------------------------+
	//	+-------------------------+----------------------------------+
	//	| 0x80070057 E_INVALIDARG | The reportJob parameter is NULL. |
	//	+-------------------------+----------------------------------+
	CreateReportJob(context.Context, *CreateReportJobRequest) (*CreateReportJobResponse, error)

	// The GetReportJob method returns the report job associated with the specified task
	// name from the List of Persisted Report Jobs (section 3.2.1.5).
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
	//	| 0x80045301 FSRM_E_NOT_FOUND | The specified report job could not be found.                                     |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG     | This code is returned for the following reasons: The taskName parameter is NULL. |
	//	|                             | The reportJob parameter is NULL.                                                 |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	GetReportJob(context.Context, *GetReportJobRequest) (*GetReportJobResponse, error)

	// The GetOutputDirectory method returns the output directory where storage reports
	// generated with the specified context will be stored.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	|         RETURN          |                                                                                  |
	//	|       VALUE/CODE        |                                   DESCRIPTION                                    |
	//	|                         |                                                                                  |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG | This code is returned for the following reasons: The context parameter is not a  |
	//	|                         | valid FsrmReportGenerationContext (section 2.2.1.2.15) value. The path parameter |
	//	|                         | is NULL.                                                                         |
	//	+-------------------------+----------------------------------------------------------------------------------+
	GetOutputDirectory(context.Context, *GetOutputDirectoryRequest) (*GetOutputDirectoryResponse, error)

	// The SetOutputDirectory method sets the output directory where storage reports generated
	// with the specified context will be stored.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	|         RETURN          |                                                                                  |
	//	|       VALUE/CODE        |                                   DESCRIPTION                                    |
	//	|                         |                                                                                  |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG | This code is returned for the following reasons: The path parameter is NULL.     |
	//	|                         | The content of the path parameter exceeds the maximum length of 150 characters.  |
	//	|                         | The context parameter is not a valid FsrmReportGenerationContext (section        |
	//	|                         | 2.2.1.2.15) value.                                                               |
	//	+-------------------------+----------------------------------------------------------------------------------+
	SetOutputDirectory(context.Context, *SetOutputDirectoryRequest) (*SetOutputDirectoryResponse, error)

	// The IsFilterValidForReportType method returns an indication of whether a specified
	// report filter is configurable for the specified report type.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	|         RETURN          |                                                                                  |
	//	|       VALUE/CODE        |                                   DESCRIPTION                                    |
	//	|                         |                                                                                  |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG | This code is returned for the following reasons: The reportType parameter is     |
	//	|                         | not a valid FsrmReportType 2.2.1.2.10 value. The filter parameter is not a valid |
	//	|                         | FsrmReportFilter (section 2.2.1.2.16) value. The valid parameter is NULL.        |
	//	+-------------------------+----------------------------------------------------------------------------------+
	IsFilterValidForReportType(context.Context, *IsFilterValidForReportTypeRequest) (*IsFilterValidForReportTypeResponse, error)

	// The GetDefaultFilter method returns the current value of the specified report filter
	// for the specified report type.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	|         RETURN          |                                                                                  |
	//	|       VALUE/CODE        |                                   DESCRIPTION                                    |
	//	|                         |                                                                                  |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG | This code is returned for the following reasons: The reportType parameter is not |
	//	|                         | a valid FsrmReportType (section 2.2.1.2.10) value. The filter parameter is not a |
	//	|                         | valid FsrmReportFilter (section 2.2.1.2.16) value. The filterValue parameter is  |
	//	|                         | NULL.                                                                            |
	//	+-------------------------+----------------------------------------------------------------------------------+
	GetDefaultFilter(context.Context, *GetDefaultFilterRequest) (*GetDefaultFilterResponse, error)

	// The SetDefaultFilter method sets the value of the specified report filter for the
	// specified report type.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	|         RETURN          |                                                                                  |
	//	|       VALUE/CODE        |                                   DESCRIPTION                                    |
	//	|                         |                                                                                  |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG | This code is returned for the following reasons: The reportType parameter is not |
	//	|                         | a valid FsrmReportType (section 2.2.1.2.10) value. The filter parameter is not a |
	//	|                         | valid FsrmReportFilter (section 2.2.1.2.16) value. The filterValue parameter is  |
	//	|                         | not a valid value for the specified report filter.                               |
	//	+-------------------------+----------------------------------------------------------------------------------+
	SetDefaultFilter(context.Context, *SetDefaultFilterRequest) (*SetDefaultFilterResponse, error)

	// The GetReportSizeLimit method returns the report size limit value of the specified
	// report limit.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	|         RETURN          |                                                                                  |
	//	|       VALUE/CODE        |                                   DESCRIPTION                                    |
	//	|                         |                                                                                  |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG | This code is returned for the following reasons: The limit parameter is not a    |
	//	|                         | valid FsrmReportLimit (section 2.2.1.2.17) value. The limitValue parameter is    |
	//	|                         | NULL.                                                                            |
	//	+-------------------------+----------------------------------------------------------------------------------+
	GetReportSizeLimit(context.Context, *GetReportSizeLimitRequest) (*GetReportSizeLimitResponse, error)

	// The SetReportSizeLimit method sets the report size limit of the specified report
	// limit.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	|         RETURN          |                                                                                  |
	//	|       VALUE/CODE        |                                   DESCRIPTION                                    |
	//	|                         |                                                                                  |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG | This code is returned for the following reasons: The limit parameter is not a    |
	//	|                         | valid FsrmReportLimit (section 2.2.1.2.17) value. The limitValue parameter is    |
	//	|                         | not a valid value; it could not be converted to a positive, long number.         |
	//	+-------------------------+----------------------------------------------------------------------------------+
	SetReportSizeLimit(context.Context, *SetReportSizeLimitRequest) (*SetReportSizeLimitResponse, error)
}

func RegisterReportManagerServer(conn dcerpc.Conn, o ReportManagerServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewReportManagerServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ReportManagerSyntaxV1_0))...)
}

func NewReportManagerServerHandle(o ReportManagerServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ReportManagerServerHandle(ctx, o, opNum, r)
	}
}

func ReportManagerServerHandle(ctx context.Context, o ReportManagerServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // EnumReportJobs
		op := &xxx_EnumReportJobsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumReportJobsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumReportJobs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // CreateReportJob
		op := &xxx_CreateReportJobOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateReportJobRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateReportJob(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // GetReportJob
		op := &xxx_GetReportJobOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetReportJobRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetReportJob(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // GetOutputDirectory
		op := &xxx_GetOutputDirectoryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetOutputDirectoryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetOutputDirectory(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // SetOutputDirectory
		op := &xxx_SetOutputDirectoryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetOutputDirectoryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetOutputDirectory(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // IsFilterValidForReportType
		op := &xxx_IsFilterValidForReportTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IsFilterValidForReportTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.IsFilterValidForReportType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // GetDefaultFilter
		op := &xxx_GetDefaultFilterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDefaultFilterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDefaultFilter(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // SetDefaultFilter
		op := &xxx_SetDefaultFilterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetDefaultFilterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetDefaultFilter(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // GetReportSizeLimit
		op := &xxx_GetReportSizeLimitOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetReportSizeLimitRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetReportSizeLimit(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // SetReportSizeLimit
		op := &xxx_SetReportSizeLimitOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetReportSizeLimitRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetReportSizeLimit(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmReportManager
type UnimplementedReportManagerServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedReportManagerServer) EnumReportJobs(context.Context, *EnumReportJobsRequest) (*EnumReportJobsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportManagerServer) CreateReportJob(context.Context, *CreateReportJobRequest) (*CreateReportJobResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportManagerServer) GetReportJob(context.Context, *GetReportJobRequest) (*GetReportJobResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportManagerServer) GetOutputDirectory(context.Context, *GetOutputDirectoryRequest) (*GetOutputDirectoryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportManagerServer) SetOutputDirectory(context.Context, *SetOutputDirectoryRequest) (*SetOutputDirectoryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportManagerServer) IsFilterValidForReportType(context.Context, *IsFilterValidForReportTypeRequest) (*IsFilterValidForReportTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportManagerServer) GetDefaultFilter(context.Context, *GetDefaultFilterRequest) (*GetDefaultFilterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportManagerServer) SetDefaultFilter(context.Context, *SetDefaultFilterRequest) (*SetDefaultFilterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportManagerServer) GetReportSizeLimit(context.Context, *GetReportSizeLimitRequest) (*GetReportSizeLimitResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportManagerServer) SetReportSizeLimit(context.Context, *SetReportSizeLimitRequest) (*SetReportSizeLimitResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ReportManagerServer = (*UnimplementedReportManagerServer)(nil)
