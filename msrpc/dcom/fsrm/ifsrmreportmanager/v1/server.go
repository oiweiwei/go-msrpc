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
		in := &EnumReportJobsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumReportJobs(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // CreateReportJob
		in := &CreateReportJobRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateReportJob(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // GetReportJob
		in := &GetReportJobRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetReportJob(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // GetOutputDirectory
		in := &GetOutputDirectoryRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetOutputDirectory(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // SetOutputDirectory
		in := &SetOutputDirectoryRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetOutputDirectory(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // IsFilterValidForReportType
		in := &IsFilterValidForReportTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.IsFilterValidForReportType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // GetDefaultFilter
		in := &GetDefaultFilterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDefaultFilter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // SetDefaultFilter
		in := &SetDefaultFilterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetDefaultFilter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // GetReportSizeLimit
		in := &GetReportSizeLimitRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetReportSizeLimit(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // SetReportSizeLimit
		in := &SetReportSizeLimitRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetReportSizeLimit(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
