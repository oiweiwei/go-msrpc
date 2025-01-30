package ifsrmreportmanager

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	fsrm "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
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
	_ = dcom.GoPackage
	_ = idispatch.GoPackage
	_ = fsrm.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

var (
	// IFsrmReportManager interface identifier 27b899fe-6ffa-4481-a184-d3daade8a02b
	ReportManagerIID = &dcom.IID{Data1: 0x27b899fe, Data2: 0x6ffa, Data3: 0x4481, Data4: []byte{0xa1, 0x84, 0xd3, 0xda, 0xad, 0xe8, 0xa0, 0x2b}}
	// Syntax UUID
	ReportManagerSyntaxUUID = &uuid.UUID{TimeLow: 0x27b899fe, TimeMid: 0x6ffa, TimeHiAndVersion: 0x4481, ClockSeqHiAndReserved: 0xa1, ClockSeqLow: 0x84, Node: [6]uint8{0xd3, 0xda, 0xad, 0xe8, 0xa0, 0x2b}}
	// Syntax ID
	ReportManagerSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: ReportManagerSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// IFsrmReportManager interface.
type ReportManagerClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

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
	EnumReportJobs(context.Context, *EnumReportJobsRequest, ...dcerpc.CallOption) (*EnumReportJobsResponse, error)

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
	CreateReportJob(context.Context, *CreateReportJobRequest, ...dcerpc.CallOption) (*CreateReportJobResponse, error)

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
	GetReportJob(context.Context, *GetReportJobRequest, ...dcerpc.CallOption) (*GetReportJobResponse, error)

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
	GetOutputDirectory(context.Context, *GetOutputDirectoryRequest, ...dcerpc.CallOption) (*GetOutputDirectoryResponse, error)

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
	SetOutputDirectory(context.Context, *SetOutputDirectoryRequest, ...dcerpc.CallOption) (*SetOutputDirectoryResponse, error)

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
	IsFilterValidForReportType(context.Context, *IsFilterValidForReportTypeRequest, ...dcerpc.CallOption) (*IsFilterValidForReportTypeResponse, error)

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
	GetDefaultFilter(context.Context, *GetDefaultFilterRequest, ...dcerpc.CallOption) (*GetDefaultFilterResponse, error)

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
	SetDefaultFilter(context.Context, *SetDefaultFilterRequest, ...dcerpc.CallOption) (*SetDefaultFilterResponse, error)

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
	GetReportSizeLimit(context.Context, *GetReportSizeLimitRequest, ...dcerpc.CallOption) (*GetReportSizeLimitResponse, error)

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
	SetReportSizeLimit(context.Context, *SetReportSizeLimitRequest, ...dcerpc.CallOption) (*SetReportSizeLimitResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ReportManagerClient
}

type xxx_DefaultReportManagerClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultReportManagerClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultReportManagerClient) EnumReportJobs(ctx context.Context, in *EnumReportJobsRequest, opts ...dcerpc.CallOption) (*EnumReportJobsResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumReportJobsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultReportManagerClient) CreateReportJob(ctx context.Context, in *CreateReportJobRequest, opts ...dcerpc.CallOption) (*CreateReportJobResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateReportJobResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultReportManagerClient) GetReportJob(ctx context.Context, in *GetReportJobRequest, opts ...dcerpc.CallOption) (*GetReportJobResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetReportJobResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultReportManagerClient) GetOutputDirectory(ctx context.Context, in *GetOutputDirectoryRequest, opts ...dcerpc.CallOption) (*GetOutputDirectoryResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetOutputDirectoryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultReportManagerClient) SetOutputDirectory(ctx context.Context, in *SetOutputDirectoryRequest, opts ...dcerpc.CallOption) (*SetOutputDirectoryResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetOutputDirectoryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultReportManagerClient) IsFilterValidForReportType(ctx context.Context, in *IsFilterValidForReportTypeRequest, opts ...dcerpc.CallOption) (*IsFilterValidForReportTypeResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &IsFilterValidForReportTypeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultReportManagerClient) GetDefaultFilter(ctx context.Context, in *GetDefaultFilterRequest, opts ...dcerpc.CallOption) (*GetDefaultFilterResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetDefaultFilterResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultReportManagerClient) SetDefaultFilter(ctx context.Context, in *SetDefaultFilterRequest, opts ...dcerpc.CallOption) (*SetDefaultFilterResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetDefaultFilterResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultReportManagerClient) GetReportSizeLimit(ctx context.Context, in *GetReportSizeLimitRequest, opts ...dcerpc.CallOption) (*GetReportSizeLimitResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetReportSizeLimitResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultReportManagerClient) SetReportSizeLimit(ctx context.Context, in *SetReportSizeLimitRequest, opts ...dcerpc.CallOption) (*SetReportSizeLimitResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetReportSizeLimitResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultReportManagerClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultReportManagerClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultReportManagerClient) IPID(ctx context.Context, ipid *dcom.IPID) ReportManagerClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultReportManagerClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewReportManagerClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ReportManagerClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ReportManagerSyntaxV1_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := idispatch.NewDispatchClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultReportManagerClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_EnumReportJobsOperation structure represents the EnumReportJobs operation
type xxx_EnumReportJobsOperation struct {
	This       *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Options    fsrm.EnumOptions `idl:"name:options" json:"options"`
	ReportJobs *fsrm.Collection `idl:"name:reportJobs" json:"report_jobs"`
	Return     int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumReportJobsOperation) OpNum() int { return 7 }

func (o *xxx_EnumReportJobsOperation) OpName() string { return "/IFsrmReportManager/v1/EnumReportJobs" }

func (o *xxx_EnumReportJobsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumReportJobsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// options {in, default_value={0}} (1:{alias=FsrmEnumOptions}(enum))
	{
		if err := w.WriteEnum(uint16(o.Options)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumReportJobsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// options {in, default_value={0}} (1:{alias=FsrmEnumOptions}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Options)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumReportJobsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumReportJobsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// reportJobs {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCollection}(interface))
	{
		if o.ReportJobs != nil {
			_ptr_reportJobs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReportJobs != nil {
					if err := o.ReportJobs.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.Collection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReportJobs, _ptr_reportJobs); err != nil {
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

func (o *xxx_EnumReportJobsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// reportJobs {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCollection}(interface))
	{
		_ptr_reportJobs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReportJobs == nil {
				o.ReportJobs = &fsrm.Collection{}
			}
			if err := o.ReportJobs.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_reportJobs := func(ptr interface{}) { o.ReportJobs = *ptr.(**fsrm.Collection) }
		if err := w.ReadPointer(&o.ReportJobs, _s_reportJobs, _ptr_reportJobs); err != nil {
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

// EnumReportJobsRequest structure represents the EnumReportJobs operation request
type EnumReportJobsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// options: Contains the FsrmEnumOptions (section 2.2.1.2.5) to use when enumerating
	// the report jobs.
	Options fsrm.EnumOptions `idl:"name:options" json:"options"`
}

func (o *EnumReportJobsRequest) xxx_ToOp(ctx context.Context) *xxx_EnumReportJobsOperation {
	if o == nil {
		return &xxx_EnumReportJobsOperation{}
	}
	return &xxx_EnumReportJobsOperation{
		This:    o.This,
		Options: o.Options,
	}
}

func (o *EnumReportJobsRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumReportJobsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Options = op.Options
}
func (o *EnumReportJobsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumReportJobsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumReportJobsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumReportJobsResponse structure represents the EnumReportJobs operation response
type EnumReportJobsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// reportJobs: Pointer to an IFsrmCollection interface pointer (section 3.2.4.2.1) that
	// upon completion contains a pointer to every report job on the server. The caller
	// MUST release the collection when the caller is done with it.
	ReportJobs *fsrm.Collection `idl:"name:reportJobs" json:"report_jobs"`
	// Return: The EnumReportJobs return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumReportJobsResponse) xxx_ToOp(ctx context.Context) *xxx_EnumReportJobsOperation {
	if o == nil {
		return &xxx_EnumReportJobsOperation{}
	}
	return &xxx_EnumReportJobsOperation{
		That:       o.That,
		ReportJobs: o.ReportJobs,
		Return:     o.Return,
	}
}

func (o *EnumReportJobsResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumReportJobsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReportJobs = op.ReportJobs
	o.Return = op.Return
}
func (o *EnumReportJobsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumReportJobsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumReportJobsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateReportJobOperation structure represents the CreateReportJob operation
type xxx_CreateReportJobOperation struct {
	This      *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ReportJob *fsrm.ReportJob `idl:"name:reportJob" json:"report_job"`
	Return    int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateReportJobOperation) OpNum() int { return 8 }

func (o *xxx_CreateReportJobOperation) OpName() string {
	return "/IFsrmReportManager/v1/CreateReportJob"
}

func (o *xxx_CreateReportJobOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateReportJobOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateReportJobOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateReportJobOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateReportJobOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// reportJob {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmReportJob}(interface))
	{
		if o.ReportJob != nil {
			_ptr_reportJob := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReportJob != nil {
					if err := o.ReportJob.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.ReportJob{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReportJob, _ptr_reportJob); err != nil {
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

func (o *xxx_CreateReportJobOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// reportJob {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmReportJob}(interface))
	{
		_ptr_reportJob := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReportJob == nil {
				o.ReportJob = &fsrm.ReportJob{}
			}
			if err := o.ReportJob.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_reportJob := func(ptr interface{}) { o.ReportJob = *ptr.(**fsrm.ReportJob) }
		if err := w.ReadPointer(&o.ReportJob, _s_reportJob, _ptr_reportJob); err != nil {
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

// CreateReportJobRequest structure represents the CreateReportJob operation request
type CreateReportJobRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *CreateReportJobRequest) xxx_ToOp(ctx context.Context) *xxx_CreateReportJobOperation {
	if o == nil {
		return &xxx_CreateReportJobOperation{}
	}
	return &xxx_CreateReportJobOperation{
		This: o.This,
	}
}

func (o *CreateReportJobRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateReportJobOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *CreateReportJobRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateReportJobRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateReportJobOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateReportJobResponse structure represents the CreateReportJob operation response
type CreateReportJobResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// reportJob: Pointer to an IFsrmReportJob interface pointer (section 3.2.4.2.33) that
	// upon completion contains a pointer to the newly created report job. To have the report
	// job added to the server's List of Persisted Report Jobs (section 3.2.1.5), the caller
	// MUST call Commit (section 3.2.4.2.34.1). The caller MUST release the report job when
	// the caller is done with it.
	ReportJob *fsrm.ReportJob `idl:"name:reportJob" json:"report_job"`
	// Return: The CreateReportJob return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateReportJobResponse) xxx_ToOp(ctx context.Context) *xxx_CreateReportJobOperation {
	if o == nil {
		return &xxx_CreateReportJobOperation{}
	}
	return &xxx_CreateReportJobOperation{
		That:      o.That,
		ReportJob: o.ReportJob,
		Return:    o.Return,
	}
}

func (o *CreateReportJobResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateReportJobOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReportJob = op.ReportJob
	o.Return = op.Return
}
func (o *CreateReportJobResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateReportJobResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateReportJobOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetReportJobOperation structure represents the GetReportJob operation
type xxx_GetReportJobOperation struct {
	This      *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat  `idl:"name:That" json:"that"`
	TaskName  *oaut.String    `idl:"name:taskName" json:"task_name"`
	ReportJob *fsrm.ReportJob `idl:"name:reportJob" json:"report_job"`
	Return    int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetReportJobOperation) OpNum() int { return 9 }

func (o *xxx_GetReportJobOperation) OpName() string { return "/IFsrmReportManager/v1/GetReportJob" }

func (o *xxx_GetReportJobOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReportJobOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// taskName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.TaskName != nil {
			_ptr_taskName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.TaskName != nil {
					if err := o.TaskName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.TaskName, _ptr_taskName); err != nil {
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

func (o *xxx_GetReportJobOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// taskName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_taskName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.TaskName == nil {
				o.TaskName = &oaut.String{}
			}
			if err := o.TaskName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_taskName := func(ptr interface{}) { o.TaskName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.TaskName, _s_taskName, _ptr_taskName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReportJobOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReportJobOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// reportJob {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmReportJob}(interface))
	{
		if o.ReportJob != nil {
			_ptr_reportJob := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReportJob != nil {
					if err := o.ReportJob.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.ReportJob{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReportJob, _ptr_reportJob); err != nil {
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

func (o *xxx_GetReportJobOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// reportJob {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmReportJob}(interface))
	{
		_ptr_reportJob := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReportJob == nil {
				o.ReportJob = &fsrm.ReportJob{}
			}
			if err := o.ReportJob.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_reportJob := func(ptr interface{}) { o.ReportJob = *ptr.(**fsrm.ReportJob) }
		if err := w.ReadPointer(&o.ReportJob, _s_reportJob, _ptr_reportJob); err != nil {
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

// GetReportJobRequest structure represents the GetReportJob operation request
type GetReportJobRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// taskName: Contains the task name for which the server will return the associated
	// report job object. The maximum length of this string MUST be 230 characters.
	TaskName *oaut.String `idl:"name:taskName" json:"task_name"`
}

func (o *GetReportJobRequest) xxx_ToOp(ctx context.Context) *xxx_GetReportJobOperation {
	if o == nil {
		return &xxx_GetReportJobOperation{}
	}
	return &xxx_GetReportJobOperation{
		This:     o.This,
		TaskName: o.TaskName,
	}
}

func (o *GetReportJobRequest) xxx_FromOp(ctx context.Context, op *xxx_GetReportJobOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.TaskName = op.TaskName
}
func (o *GetReportJobRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetReportJobRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetReportJobOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetReportJobResponse structure represents the GetReportJob operation response
type GetReportJobResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// reportJob: Pointer to an IFsrmReportJob interface pointer (section 3.2.4.2.34) that
	// upon completion contains a pointer to the report job object for the task name specified.
	// The caller MUST release the report job when the caller is done with it.
	ReportJob *fsrm.ReportJob `idl:"name:reportJob" json:"report_job"`
	// Return: The GetReportJob return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetReportJobResponse) xxx_ToOp(ctx context.Context) *xxx_GetReportJobOperation {
	if o == nil {
		return &xxx_GetReportJobOperation{}
	}
	return &xxx_GetReportJobOperation{
		That:      o.That,
		ReportJob: o.ReportJob,
		Return:    o.Return,
	}
}

func (o *GetReportJobResponse) xxx_FromOp(ctx context.Context, op *xxx_GetReportJobOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReportJob = op.ReportJob
	o.Return = op.Return
}
func (o *GetReportJobResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetReportJobResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetReportJobOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetOutputDirectoryOperation structure represents the GetOutputDirectory operation
type xxx_GetOutputDirectoryOperation struct {
	This    *dcom.ORPCThis               `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat               `idl:"name:That" json:"that"`
	Context fsrm.ReportGenerationContext `idl:"name:context" json:"context"`
	Path    *oaut.String                 `idl:"name:path" json:"path"`
	Return  int32                        `idl:"name:Return" json:"return"`
}

func (o *xxx_GetOutputDirectoryOperation) OpNum() int { return 10 }

func (o *xxx_GetOutputDirectoryOperation) OpName() string {
	return "/IFsrmReportManager/v1/GetOutputDirectory"
}

func (o *xxx_GetOutputDirectoryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOutputDirectoryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// context {in} (1:{alias=FsrmReportGenerationContext}(enum))
	{
		if err := w.WriteEnum(uint16(o.Context)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOutputDirectoryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// context {in} (1:{alias=FsrmReportGenerationContext}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Context)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOutputDirectoryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOutputDirectoryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// path {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Path != nil {
			_ptr_path := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Path != nil {
					if err := o.Path.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOutputDirectoryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// path {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_path := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Path == nil {
				o.Path = &oaut.String{}
			}
			if err := o.Path.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_path := func(ptr interface{}) { o.Path = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Path, _s_path, _ptr_path); err != nil {
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

// GetOutputDirectoryRequest structure represents the GetOutputDirectory operation request
type GetOutputDirectoryRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// context: Contains the value from the FsrmReportGenerationContext (section 2.2.1.2.15)
	// enumeration of the context to get.
	Context fsrm.ReportGenerationContext `idl:"name:context" json:"context"`
}

func (o *GetOutputDirectoryRequest) xxx_ToOp(ctx context.Context) *xxx_GetOutputDirectoryOperation {
	if o == nil {
		return &xxx_GetOutputDirectoryOperation{}
	}
	return &xxx_GetOutputDirectoryOperation{
		This:    o.This,
		Context: o.Context,
	}
}

func (o *GetOutputDirectoryRequest) xxx_FromOp(ctx context.Context, op *xxx_GetOutputDirectoryOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Context = op.Context
}
func (o *GetOutputDirectoryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetOutputDirectoryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetOutputDirectoryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetOutputDirectoryResponse structure represents the GetOutputDirectory operation response
type GetOutputDirectoryResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// path: Pointer to a variable that upon completion contains the output directory where
	// storage reports generated with the specified context will be stored.
	Path *oaut.String `idl:"name:path" json:"path"`
	// Return: The GetOutputDirectory return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetOutputDirectoryResponse) xxx_ToOp(ctx context.Context) *xxx_GetOutputDirectoryOperation {
	if o == nil {
		return &xxx_GetOutputDirectoryOperation{}
	}
	return &xxx_GetOutputDirectoryOperation{
		That:   o.That,
		Path:   o.Path,
		Return: o.Return,
	}
}

func (o *GetOutputDirectoryResponse) xxx_FromOp(ctx context.Context, op *xxx_GetOutputDirectoryOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Path = op.Path
	o.Return = op.Return
}
func (o *GetOutputDirectoryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetOutputDirectoryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetOutputDirectoryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetOutputDirectoryOperation structure represents the SetOutputDirectory operation
type xxx_SetOutputDirectoryOperation struct {
	This    *dcom.ORPCThis               `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat               `idl:"name:That" json:"that"`
	Context fsrm.ReportGenerationContext `idl:"name:context" json:"context"`
	Path    *oaut.String                 `idl:"name:path" json:"path"`
	Return  int32                        `idl:"name:Return" json:"return"`
}

func (o *xxx_SetOutputDirectoryOperation) OpNum() int { return 11 }

func (o *xxx_SetOutputDirectoryOperation) OpName() string {
	return "/IFsrmReportManager/v1/SetOutputDirectory"
}

func (o *xxx_SetOutputDirectoryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOutputDirectoryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// context {in} (1:{alias=FsrmReportGenerationContext}(enum))
	{
		if err := w.WriteEnum(uint16(o.Context)); err != nil {
			return err
		}
	}
	// path {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Path != nil {
			_ptr_path := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Path != nil {
					if err := o.Path.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
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
	return nil
}

func (o *xxx_SetOutputDirectoryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// context {in} (1:{alias=FsrmReportGenerationContext}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Context)); err != nil {
			return err
		}
	}
	// path {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_path := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Path == nil {
				o.Path = &oaut.String{}
			}
			if err := o.Path.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_path := func(ptr interface{}) { o.Path = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Path, _s_path, _ptr_path); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOutputDirectoryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOutputDirectoryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_SetOutputDirectoryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
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

// SetOutputDirectoryRequest structure represents the SetOutputDirectory operation request
type SetOutputDirectoryRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// context: Contains the value from the FsrmReportGenerationContext (section 2.2.1.2.15)
	// enumeration of the context to set.
	Context fsrm.ReportGenerationContext `idl:"name:context" json:"context"`
	// path: Contains the output directory where storage reports generated with the specified
	// context will be stored. The maximum length of this string MUST be 150 characters.
	Path *oaut.String `idl:"name:path" json:"path"`
}

func (o *SetOutputDirectoryRequest) xxx_ToOp(ctx context.Context) *xxx_SetOutputDirectoryOperation {
	if o == nil {
		return &xxx_SetOutputDirectoryOperation{}
	}
	return &xxx_SetOutputDirectoryOperation{
		This:    o.This,
		Context: o.Context,
		Path:    o.Path,
	}
}

func (o *SetOutputDirectoryRequest) xxx_FromOp(ctx context.Context, op *xxx_SetOutputDirectoryOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Context = op.Context
	o.Path = op.Path
}
func (o *SetOutputDirectoryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetOutputDirectoryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetOutputDirectoryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetOutputDirectoryResponse structure represents the SetOutputDirectory operation response
type SetOutputDirectoryResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetOutputDirectory return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetOutputDirectoryResponse) xxx_ToOp(ctx context.Context) *xxx_SetOutputDirectoryOperation {
	if o == nil {
		return &xxx_SetOutputDirectoryOperation{}
	}
	return &xxx_SetOutputDirectoryOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetOutputDirectoryResponse) xxx_FromOp(ctx context.Context, op *xxx_SetOutputDirectoryOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetOutputDirectoryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetOutputDirectoryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetOutputDirectoryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_IsFilterValidForReportTypeOperation structure represents the IsFilterValidForReportType operation
type xxx_IsFilterValidForReportTypeOperation struct {
	This       *dcom.ORPCThis    `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat    `idl:"name:That" json:"that"`
	ReportType fsrm.ReportType   `idl:"name:reportType" json:"report_type"`
	Filter     fsrm.ReportFilter `idl:"name:filter" json:"filter"`
	Valid      int16             `idl:"name:valid" json:"valid"`
	Return     int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_IsFilterValidForReportTypeOperation) OpNum() int { return 12 }

func (o *xxx_IsFilterValidForReportTypeOperation) OpName() string {
	return "/IFsrmReportManager/v1/IsFilterValidForReportType"
}

func (o *xxx_IsFilterValidForReportTypeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsFilterValidForReportTypeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// reportType {in} (1:{alias=FsrmReportType}(enum))
	{
		if err := w.WriteEnum(uint16(o.ReportType)); err != nil {
			return err
		}
	}
	// filter {in} (1:{alias=FsrmReportFilter}(enum))
	{
		if err := w.WriteEnum(uint16(o.Filter)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsFilterValidForReportTypeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// reportType {in} (1:{alias=FsrmReportType}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.ReportType)); err != nil {
			return err
		}
	}
	// filter {in} (1:{alias=FsrmReportFilter}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Filter)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsFilterValidForReportTypeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsFilterValidForReportTypeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// valid {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Valid); err != nil {
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

func (o *xxx_IsFilterValidForReportTypeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// valid {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Valid); err != nil {
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

// IsFilterValidForReportTypeRequest structure represents the IsFilterValidForReportType operation request
type IsFilterValidForReportTypeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// reportType: Contains the value from the FsrmReportType (section 2.2.1.2.10) enumeration.
	ReportType fsrm.ReportType `idl:"name:reportType" json:"report_type"`
	// filter: Contains the value from the FsrmReportFilter (section 2.2.1.2.16) enumeration.
	Filter fsrm.ReportFilter `idl:"name:filter" json:"filter"`
}

func (o *IsFilterValidForReportTypeRequest) xxx_ToOp(ctx context.Context) *xxx_IsFilterValidForReportTypeOperation {
	if o == nil {
		return &xxx_IsFilterValidForReportTypeOperation{}
	}
	return &xxx_IsFilterValidForReportTypeOperation{
		This:       o.This,
		ReportType: o.ReportType,
		Filter:     o.Filter,
	}
}

func (o *IsFilterValidForReportTypeRequest) xxx_FromOp(ctx context.Context, op *xxx_IsFilterValidForReportTypeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ReportType = op.ReportType
	o.Filter = op.Filter
}
func (o *IsFilterValidForReportTypeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *IsFilterValidForReportTypeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsFilterValidForReportTypeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IsFilterValidForReportTypeResponse structure represents the IsFilterValidForReportType operation response
type IsFilterValidForReportTypeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// valid: Pointer to a variable that upon completion contains the Boolean indication
	// of whether the specified filter is configurable for the specified report type.
	Valid int16 `idl:"name:valid" json:"valid"`
	// Return: The IsFilterValidForReportType return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *IsFilterValidForReportTypeResponse) xxx_ToOp(ctx context.Context) *xxx_IsFilterValidForReportTypeOperation {
	if o == nil {
		return &xxx_IsFilterValidForReportTypeOperation{}
	}
	return &xxx_IsFilterValidForReportTypeOperation{
		That:   o.That,
		Valid:  o.Valid,
		Return: o.Return,
	}
}

func (o *IsFilterValidForReportTypeResponse) xxx_FromOp(ctx context.Context, op *xxx_IsFilterValidForReportTypeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Valid = op.Valid
	o.Return = op.Return
}
func (o *IsFilterValidForReportTypeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *IsFilterValidForReportTypeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsFilterValidForReportTypeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDefaultFilterOperation structure represents the GetDefaultFilter operation
type xxx_GetDefaultFilterOperation struct {
	This        *dcom.ORPCThis    `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat    `idl:"name:That" json:"that"`
	ReportType  fsrm.ReportType   `idl:"name:reportType" json:"report_type"`
	Filter      fsrm.ReportFilter `idl:"name:filter" json:"filter"`
	FilterValue *oaut.Variant     `idl:"name:filterValue" json:"filter_value"`
	Return      int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDefaultFilterOperation) OpNum() int { return 13 }

func (o *xxx_GetDefaultFilterOperation) OpName() string {
	return "/IFsrmReportManager/v1/GetDefaultFilter"
}

func (o *xxx_GetDefaultFilterOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDefaultFilterOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// reportType {in} (1:{alias=FsrmReportType}(enum))
	{
		if err := w.WriteEnum(uint16(o.ReportType)); err != nil {
			return err
		}
	}
	// filter {in} (1:{alias=FsrmReportFilter}(enum))
	{
		if err := w.WriteEnum(uint16(o.Filter)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDefaultFilterOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// reportType {in} (1:{alias=FsrmReportType}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.ReportType)); err != nil {
			return err
		}
	}
	// filter {in} (1:{alias=FsrmReportFilter}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Filter)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDefaultFilterOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDefaultFilterOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// filterValue {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.FilterValue != nil {
			_ptr_filterValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FilterValue != nil {
					if err := o.FilterValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FilterValue, _ptr_filterValue); err != nil {
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

func (o *xxx_GetDefaultFilterOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// filterValue {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_filterValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FilterValue == nil {
				o.FilterValue = &oaut.Variant{}
			}
			if err := o.FilterValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_filterValue := func(ptr interface{}) { o.FilterValue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.FilterValue, _s_filterValue, _ptr_filterValue); err != nil {
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

// GetDefaultFilterRequest structure represents the GetDefaultFilter operation request
type GetDefaultFilterRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// reportType: Contains the value from the FsrmReportType (section 2.2.1.2.10) enumeration.
	ReportType fsrm.ReportType `idl:"name:reportType" json:"report_type"`
	// filter: Contains the value from the FsrmReportFilter (section 2.2.1.2.16) enumeration.
	Filter fsrm.ReportFilter `idl:"name:filter" json:"filter"`
}

func (o *GetDefaultFilterRequest) xxx_ToOp(ctx context.Context) *xxx_GetDefaultFilterOperation {
	if o == nil {
		return &xxx_GetDefaultFilterOperation{}
	}
	return &xxx_GetDefaultFilterOperation{
		This:       o.This,
		ReportType: o.ReportType,
		Filter:     o.Filter,
	}
}

func (o *GetDefaultFilterRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDefaultFilterOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ReportType = op.ReportType
	o.Filter = op.Filter
}
func (o *GetDefaultFilterRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetDefaultFilterRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDefaultFilterOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDefaultFilterResponse structure represents the GetDefaultFilter operation response
type GetDefaultFilterResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// filterValue: Pointer to a variable that upon completion contains the value of the
	// specified filter for the specified report type.
	FilterValue *oaut.Variant `idl:"name:filterValue" json:"filter_value"`
	// Return: The GetDefaultFilter return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDefaultFilterResponse) xxx_ToOp(ctx context.Context) *xxx_GetDefaultFilterOperation {
	if o == nil {
		return &xxx_GetDefaultFilterOperation{}
	}
	return &xxx_GetDefaultFilterOperation{
		That:        o.That,
		FilterValue: o.FilterValue,
		Return:      o.Return,
	}
}

func (o *GetDefaultFilterResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDefaultFilterOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FilterValue = op.FilterValue
	o.Return = op.Return
}
func (o *GetDefaultFilterResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetDefaultFilterResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDefaultFilterOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetDefaultFilterOperation structure represents the SetDefaultFilter operation
type xxx_SetDefaultFilterOperation struct {
	This        *dcom.ORPCThis    `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat    `idl:"name:That" json:"that"`
	ReportType  fsrm.ReportType   `idl:"name:reportType" json:"report_type"`
	Filter      fsrm.ReportFilter `idl:"name:filter" json:"filter"`
	FilterValue *oaut.Variant     `idl:"name:filterValue" json:"filter_value"`
	Return      int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_SetDefaultFilterOperation) OpNum() int { return 14 }

func (o *xxx_SetDefaultFilterOperation) OpName() string {
	return "/IFsrmReportManager/v1/SetDefaultFilter"
}

func (o *xxx_SetDefaultFilterOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDefaultFilterOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// reportType {in} (1:{alias=FsrmReportType}(enum))
	{
		if err := w.WriteEnum(uint16(o.ReportType)); err != nil {
			return err
		}
	}
	// filter {in} (1:{alias=FsrmReportFilter}(enum))
	{
		if err := w.WriteEnum(uint16(o.Filter)); err != nil {
			return err
		}
	}
	// filterValue {in} (1:{alias=VARIANT}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.FilterValue != nil {
			if err := o.FilterValue.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDefaultFilterOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// reportType {in} (1:{alias=FsrmReportType}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.ReportType)); err != nil {
			return err
		}
	}
	// filter {in} (1:{alias=FsrmReportFilter}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Filter)); err != nil {
			return err
		}
	}
	// filterValue {in} (1:{alias=VARIANT,pointer=ref}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.FilterValue == nil {
			o.FilterValue = &oaut.Variant{}
		}
		if err := o.FilterValue.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDefaultFilterOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDefaultFilterOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_SetDefaultFilterOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
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

// SetDefaultFilterRequest structure represents the SetDefaultFilter operation request
type SetDefaultFilterRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// reportType: Contains the value from the FsrmReportType (section 2.2.1.2.10) enumeration.
	ReportType fsrm.ReportType `idl:"name:reportType" json:"report_type"`
	// filter: Contains the value from the FsrmReportFilter (section 2.2.1.2.16) enumeration.
	Filter fsrm.ReportFilter `idl:"name:filter" json:"filter"`
	// filterValue: Contains the value of the specified filter for the specified report
	// type.
	FilterValue *oaut.Variant `idl:"name:filterValue" json:"filter_value"`
}

func (o *SetDefaultFilterRequest) xxx_ToOp(ctx context.Context) *xxx_SetDefaultFilterOperation {
	if o == nil {
		return &xxx_SetDefaultFilterOperation{}
	}
	return &xxx_SetDefaultFilterOperation{
		This:        o.This,
		ReportType:  o.ReportType,
		Filter:      o.Filter,
		FilterValue: o.FilterValue,
	}
}

func (o *SetDefaultFilterRequest) xxx_FromOp(ctx context.Context, op *xxx_SetDefaultFilterOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ReportType = op.ReportType
	o.Filter = op.Filter
	o.FilterValue = op.FilterValue
}
func (o *SetDefaultFilterRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetDefaultFilterRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDefaultFilterOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetDefaultFilterResponse structure represents the SetDefaultFilter operation response
type SetDefaultFilterResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetDefaultFilter return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetDefaultFilterResponse) xxx_ToOp(ctx context.Context) *xxx_SetDefaultFilterOperation {
	if o == nil {
		return &xxx_SetDefaultFilterOperation{}
	}
	return &xxx_SetDefaultFilterOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetDefaultFilterResponse) xxx_FromOp(ctx context.Context, op *xxx_SetDefaultFilterOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetDefaultFilterResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetDefaultFilterResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDefaultFilterOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetReportSizeLimitOperation structure represents the GetReportSizeLimit operation
type xxx_GetReportSizeLimitOperation struct {
	This       *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Limit      fsrm.ReportLimit `idl:"name:limit" json:"limit"`
	LimitValue *oaut.Variant    `idl:"name:limitValue" json:"limit_value"`
	Return     int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetReportSizeLimitOperation) OpNum() int { return 15 }

func (o *xxx_GetReportSizeLimitOperation) OpName() string {
	return "/IFsrmReportManager/v1/GetReportSizeLimit"
}

func (o *xxx_GetReportSizeLimitOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReportSizeLimitOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// limit {in} (1:{alias=FsrmReportLimit}(enum))
	{
		if err := w.WriteEnum(uint16(o.Limit)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReportSizeLimitOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// limit {in} (1:{alias=FsrmReportLimit}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Limit)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReportSizeLimitOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReportSizeLimitOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// limitValue {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.LimitValue != nil {
			_ptr_limitValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.LimitValue != nil {
					if err := o.LimitValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.LimitValue, _ptr_limitValue); err != nil {
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

func (o *xxx_GetReportSizeLimitOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// limitValue {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_limitValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.LimitValue == nil {
				o.LimitValue = &oaut.Variant{}
			}
			if err := o.LimitValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_limitValue := func(ptr interface{}) { o.LimitValue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.LimitValue, _s_limitValue, _ptr_limitValue); err != nil {
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

// GetReportSizeLimitRequest structure represents the GetReportSizeLimit operation request
type GetReportSizeLimitRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// limit: Contains the value from the FsrmReportLimit (section 2.2.1.2.17) enumeration
	// indicating the report size limit for which the value is returned.
	Limit fsrm.ReportLimit `idl:"name:limit" json:"limit"`
}

func (o *GetReportSizeLimitRequest) xxx_ToOp(ctx context.Context) *xxx_GetReportSizeLimitOperation {
	if o == nil {
		return &xxx_GetReportSizeLimitOperation{}
	}
	return &xxx_GetReportSizeLimitOperation{
		This:  o.This,
		Limit: o.Limit,
	}
}

func (o *GetReportSizeLimitRequest) xxx_FromOp(ctx context.Context, op *xxx_GetReportSizeLimitOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Limit = op.Limit
}
func (o *GetReportSizeLimitRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetReportSizeLimitRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetReportSizeLimitOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetReportSizeLimitResponse structure represents the GetReportSizeLimit operation response
type GetReportSizeLimitResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// limitValue: Pointer to a variable that upon completion contains the report size limit
	// of the specified report limit.
	LimitValue *oaut.Variant `idl:"name:limitValue" json:"limit_value"`
	// Return: The GetReportSizeLimit return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetReportSizeLimitResponse) xxx_ToOp(ctx context.Context) *xxx_GetReportSizeLimitOperation {
	if o == nil {
		return &xxx_GetReportSizeLimitOperation{}
	}
	return &xxx_GetReportSizeLimitOperation{
		That:       o.That,
		LimitValue: o.LimitValue,
		Return:     o.Return,
	}
}

func (o *GetReportSizeLimitResponse) xxx_FromOp(ctx context.Context, op *xxx_GetReportSizeLimitOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.LimitValue = op.LimitValue
	o.Return = op.Return
}
func (o *GetReportSizeLimitResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetReportSizeLimitResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetReportSizeLimitOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetReportSizeLimitOperation structure represents the SetReportSizeLimit operation
type xxx_SetReportSizeLimitOperation struct {
	This       *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Limit      fsrm.ReportLimit `idl:"name:limit" json:"limit"`
	LimitValue *oaut.Variant    `idl:"name:limitValue" json:"limit_value"`
	Return     int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_SetReportSizeLimitOperation) OpNum() int { return 16 }

func (o *xxx_SetReportSizeLimitOperation) OpName() string {
	return "/IFsrmReportManager/v1/SetReportSizeLimit"
}

func (o *xxx_SetReportSizeLimitOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetReportSizeLimitOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// limit {in} (1:{alias=FsrmReportLimit}(enum))
	{
		if err := w.WriteEnum(uint16(o.Limit)); err != nil {
			return err
		}
	}
	// limitValue {in} (1:{alias=VARIANT}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.LimitValue != nil {
			if err := o.LimitValue.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetReportSizeLimitOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// limit {in} (1:{alias=FsrmReportLimit}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Limit)); err != nil {
			return err
		}
	}
	// limitValue {in} (1:{alias=VARIANT,pointer=ref}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.LimitValue == nil {
			o.LimitValue = &oaut.Variant{}
		}
		if err := o.LimitValue.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetReportSizeLimitOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetReportSizeLimitOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_SetReportSizeLimitOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
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

// SetReportSizeLimitRequest structure represents the SetReportSizeLimit operation request
type SetReportSizeLimitRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// limit: Contains the value from the FsrmReportLimit (section 2.2.1.2.17) enumeration
	// indicating the report size limit that is set.
	Limit fsrm.ReportLimit `idl:"name:limit" json:"limit"`
	// limitValue: Contains the value of the specified report size limit.
	LimitValue *oaut.Variant `idl:"name:limitValue" json:"limit_value"`
}

func (o *SetReportSizeLimitRequest) xxx_ToOp(ctx context.Context) *xxx_SetReportSizeLimitOperation {
	if o == nil {
		return &xxx_SetReportSizeLimitOperation{}
	}
	return &xxx_SetReportSizeLimitOperation{
		This:       o.This,
		Limit:      o.Limit,
		LimitValue: o.LimitValue,
	}
}

func (o *SetReportSizeLimitRequest) xxx_FromOp(ctx context.Context, op *xxx_SetReportSizeLimitOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Limit = op.Limit
	o.LimitValue = op.LimitValue
}
func (o *SetReportSizeLimitRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetReportSizeLimitRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetReportSizeLimitOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetReportSizeLimitResponse structure represents the SetReportSizeLimit operation response
type SetReportSizeLimitResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetReportSizeLimit return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetReportSizeLimitResponse) xxx_ToOp(ctx context.Context) *xxx_SetReportSizeLimitOperation {
	if o == nil {
		return &xxx_SetReportSizeLimitOperation{}
	}
	return &xxx_SetReportSizeLimitOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetReportSizeLimitResponse) xxx_FromOp(ctx context.Context, op *xxx_SetReportSizeLimitOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetReportSizeLimitResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetReportSizeLimitResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetReportSizeLimitOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
