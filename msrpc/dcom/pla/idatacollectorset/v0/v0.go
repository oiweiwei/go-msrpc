package idatacollectorset

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
	pla "github.com/oiweiwei/go-msrpc/msrpc/dcom/pla"
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
	_ = pla.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/pla"
)

var (
	// IDataCollectorSet interface identifier 03837520-098b-11d8-9414-505054503030
	DataCollectorSetIID = &dcom.IID{Data1: 0x03837520, Data2: 0x098b, Data3: 0x11d8, Data4: []byte{0x94, 0x14, 0x50, 0x50, 0x54, 0x50, 0x30, 0x30}}
	// Syntax UUID
	DataCollectorSetSyntaxUUID = &uuid.UUID{TimeLow: 0x3837520, TimeMid: 0x98b, TimeHiAndVersion: 0x11d8, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0x14, Node: [6]uint8{0x50, 0x50, 0x54, 0x50, 0x30, 0x30}}
	// Syntax ID
	DataCollectorSetSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: DataCollectorSetSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IDataCollectorSet interface.
type DataCollectorSetClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// The DataCollectors (Get) method retrieves the DataCollectors property.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetDataCollectors(context.Context, *GetDataCollectorsRequest, ...dcerpc.CallOption) (*GetDataCollectorsResponse, error)

	// Duration operation.
	GetDuration(context.Context, *GetDurationRequest, ...dcerpc.CallOption) (*GetDurationResponse, error)

	// Duration operation.
	SetDuration(context.Context, *SetDurationRequest, ...dcerpc.CallOption) (*SetDurationResponse, error)

	// Description operation.
	GetDescription(context.Context, *GetDescriptionRequest, ...dcerpc.CallOption) (*GetDescriptionResponse, error)

	// Description operation.
	SetDescription(context.Context, *SetDescriptionRequest, ...dcerpc.CallOption) (*SetDescriptionResponse, error)

	// The DescriptionUnresolved (Get) method retrieves the DescriptionUnresolved property.
	//
	// This method MUST return the description as originally set by the IDataCollectorSet::Description
	// method, as specified in section 3.2.4.1.5.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetDescriptionUnresolved(context.Context, *GetDescriptionUnresolvedRequest, ...dcerpc.CallOption) (*GetDescriptionUnresolvedResponse, error)

	// DisplayName operation.
	GetDisplayName(context.Context, *GetDisplayNameRequest, ...dcerpc.CallOption) (*GetDisplayNameResponse, error)

	// DisplayName operation.
	SetDisplayName(context.Context, *SetDisplayNameRequest, ...dcerpc.CallOption) (*SetDisplayNameResponse, error)

	// The DisplayNameUnresolved (Get) method retrieves the DisplayNameUnresolved property.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetDisplayNameUnresolved(context.Context, *GetDisplayNameUnresolvedRequest, ...dcerpc.CallOption) (*GetDisplayNameUnresolvedResponse, error)

	// Keywords operation.
	GetKeywords(context.Context, *GetKeywordsRequest, ...dcerpc.CallOption) (*GetKeywordsResponse, error)

	// Keywords operation.
	SetKeywords(context.Context, *SetKeywordsRequest, ...dcerpc.CallOption) (*SetKeywordsResponse, error)

	// LatestOutputLocation operation.
	GetLatestOutputLocation(context.Context, *GetLatestOutputLocationRequest, ...dcerpc.CallOption) (*GetLatestOutputLocationResponse, error)

	// LatestOutputLocation operation.
	SetLatestOutputLocation(context.Context, *SetLatestOutputLocationRequest, ...dcerpc.CallOption) (*SetLatestOutputLocationResponse, error)

	// Name operation.
	GetName(context.Context, *GetNameRequest, ...dcerpc.CallOption) (*GetNameResponse, error)

	// OutputLocation operation.
	GetOutputLocation(context.Context, *GetOutputLocationRequest, ...dcerpc.CallOption) (*GetOutputLocationResponse, error)

	// RootPath operation.
	GetRootPath(context.Context, *GetRootPathRequest, ...dcerpc.CallOption) (*GetRootPathResponse, error)

	// RootPath operation.
	SetRootPath(context.Context, *SetRootPathRequest, ...dcerpc.CallOption) (*SetRootPathResponse, error)

	// Segment operation.
	GetSegment(context.Context, *GetSegmentRequest, ...dcerpc.CallOption) (*GetSegmentResponse, error)

	// Segment operation.
	SetSegment(context.Context, *SetSegmentRequest, ...dcerpc.CallOption) (*SetSegmentResponse, error)

	// SegmentMaxDuration operation.
	GetSegmentMaxDuration(context.Context, *GetSegmentMaxDurationRequest, ...dcerpc.CallOption) (*GetSegmentMaxDurationResponse, error)

	// SegmentMaxDuration operation.
	SetSegmentMaxDuration(context.Context, *SetSegmentMaxDurationRequest, ...dcerpc.CallOption) (*SetSegmentMaxDurationResponse, error)

	// SegmentMaxSize operation.
	GetSegmentMaxSize(context.Context, *GetSegmentMaxSizeRequest, ...dcerpc.CallOption) (*GetSegmentMaxSizeResponse, error)

	// SegmentMaxSize operation.
	SetSegmentMaxSize(context.Context, *SetSegmentMaxSizeRequest, ...dcerpc.CallOption) (*SetSegmentMaxSizeResponse, error)

	// SerialNumber operation.
	GetSerialNumber(context.Context, *GetSerialNumberRequest, ...dcerpc.CallOption) (*GetSerialNumberResponse, error)

	// SerialNumber operation.
	SetSerialNumber(context.Context, *SetSerialNumberRequest, ...dcerpc.CallOption) (*SetSerialNumberResponse, error)

	// The Server (Get) method retrieves the Server property.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetServer(context.Context, *GetServerRequest, ...dcerpc.CallOption) (*GetServerResponse, error)

	// The Status (Get) method retrieves the Status property.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetStatus(context.Context, *GetStatusRequest, ...dcerpc.CallOption) (*GetStatusResponse, error)

	// Subdirectory operation.
	GetSubdirectory(context.Context, *GetSubdirectoryRequest, ...dcerpc.CallOption) (*GetSubdirectoryResponse, error)

	// Subdirectory operation.
	SetSubdirectory(context.Context, *SetSubdirectoryRequest, ...dcerpc.CallOption) (*SetSubdirectoryResponse, error)

	// SubdirectoryFormat operation.
	GetSubdirectoryFormat(context.Context, *GetSubdirectoryFormatRequest, ...dcerpc.CallOption) (*GetSubdirectoryFormatResponse, error)

	// SubdirectoryFormat operation.
	SetSubdirectoryFormat(context.Context, *SetSubdirectoryFormatRequest, ...dcerpc.CallOption) (*SetSubdirectoryFormatResponse, error)

	// SubdirectoryFormatPattern operation.
	GetSubdirectoryFormatPattern(context.Context, *GetSubdirectoryFormatPatternRequest, ...dcerpc.CallOption) (*GetSubdirectoryFormatPatternResponse, error)

	// SubdirectoryFormatPattern operation.
	SetSubdirectoryFormatPattern(context.Context, *SetSubdirectoryFormatPatternRequest, ...dcerpc.CallOption) (*SetSubdirectoryFormatPatternResponse, error)

	// Task operation.
	GetTask(context.Context, *GetTaskRequest, ...dcerpc.CallOption) (*GetTaskResponse, error)

	// Task operation.
	SetTask(context.Context, *SetTaskRequest, ...dcerpc.CallOption) (*SetTaskResponse, error)

	// TaskRunAsSelf operation.
	GetTaskRunAsSelf(context.Context, *GetTaskRunAsSelfRequest, ...dcerpc.CallOption) (*GetTaskRunAsSelfResponse, error)

	// TaskRunAsSelf operation.
	SetTaskRunAsSelf(context.Context, *SetTaskRunAsSelfRequest, ...dcerpc.CallOption) (*SetTaskRunAsSelfResponse, error)

	// TaskArguments operation.
	GetTaskArguments(context.Context, *GetTaskArgumentsRequest, ...dcerpc.CallOption) (*GetTaskArgumentsResponse, error)

	// TaskArguments operation.
	SetTaskArguments(context.Context, *SetTaskArgumentsRequest, ...dcerpc.CallOption) (*SetTaskArgumentsResponse, error)

	// TaskUserTextArguments operation.
	GetTaskUserTextArguments(context.Context, *GetTaskUserTextArgumentsRequest, ...dcerpc.CallOption) (*GetTaskUserTextArgumentsResponse, error)

	// TaskUserTextArguments operation.
	SetTaskUserTextArguments(context.Context, *SetTaskUserTextArgumentsRequest, ...dcerpc.CallOption) (*SetTaskUserTextArgumentsResponse, error)

	// The Schedules (Get) method retrieves the Schedules property, as specified in the
	// property table in section 3.2.4.1.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetSchedules(context.Context, *GetSchedulesRequest, ...dcerpc.CallOption) (*GetSchedulesResponse, error)

	// SchedulesEnabled operation.
	GetSchedulesEnabled(context.Context, *GetSchedulesEnabledRequest, ...dcerpc.CallOption) (*GetSchedulesEnabledResponse, error)

	// SchedulesEnabled operation.
	SetSchedulesEnabled(context.Context, *SetSchedulesEnabledRequest, ...dcerpc.CallOption) (*SetSchedulesEnabledResponse, error)

	// The UserAccount (Get) method retrieves the UserAccount property.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetUserAccount(context.Context, *GetUserAccountRequest, ...dcerpc.CallOption) (*GetUserAccountResponse, error)

	// Xml operation.
	GetXML(context.Context, *GetXMLRequest, ...dcerpc.CallOption) (*GetXMLResponse, error)

	// Security operation.
	GetSecurity(context.Context, *GetSecurityRequest, ...dcerpc.CallOption) (*GetSecurityResponse, error)

	// Security operation.
	SetSecurity(context.Context, *SetSecurityRequest, ...dcerpc.CallOption) (*SetSecurityResponse, error)

	// StopOnCompletion operation.
	GetStopOnCompletion(context.Context, *GetStopOnCompletionRequest, ...dcerpc.CallOption) (*GetStopOnCompletionResponse, error)

	// StopOnCompletion operation.
	SetStopOnCompletion(context.Context, *SetStopOnCompletionRequest, ...dcerpc.CallOption) (*SetStopOnCompletionResponse, error)

	// The DataManager (Get) method retrieves the DataManager property, as specified in
	// the property table in section 3.2.4.1.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetDataManager(context.Context, *GetDataManagerRequest, ...dcerpc.CallOption) (*GetDataManagerResponse, error)

	// The SetCredentials method specifies the user account under which the data collector
	// set runs. If both parameters are set to NULL, the Performance Logs and Alerts Protocol
	// MUST clear the existing cached credentials. If no credentials are cached, the data
	// collector set runs with the credentials for the 'LocalSystem' account. Changing the
	// credentials under which the data collector set runs will impact future runs of the
	// data collector set, irrespective of whether they are started locally or remotely.
	// <16>
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	SetCredentials(context.Context, *SetCredentialsRequest, ...dcerpc.CallOption) (*SetCredentialsResponse, error)

	// Query operation.
	Query(context.Context, *QueryRequest, ...dcerpc.CallOption) (*QueryResponse, error)

	// The Commit method updates, validates, or saves a data collector set, or flushes the
	// event trace data collectors of a data collector set.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	Commit(context.Context, *CommitRequest, ...dcerpc.CallOption) (*CommitResponse, error)

	// The Delete method removes the data collector set from storage if it is not running.
	//
	// This method has no parameters.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	Delete(context.Context, *DeleteRequest, ...dcerpc.CallOption) (*DeleteResponse, error)

	// The Start method manually starts the data collector set.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	Start(context.Context, *StartRequest, ...dcerpc.CallOption) (*StartResponse, error)

	// The Stop method manually stops the data collector set.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	Stop(context.Context, *StopRequest, ...dcerpc.CallOption) (*StopResponse, error)

	// SetXml operation.
	SetXML(context.Context, *SetXMLRequest, ...dcerpc.CallOption) (*SetXMLResponse, error)

	// The SetValue method sets a user-defined value. The SetValue method takes a name and
	// value pair as BSTRs. This pair can either be used as metadata for the data collector
	// set, in which case it has no effect on the execution of the data collector set, or
	// it can be used as arguments for task execution. For example, when a performance counter
	// crosses a specified threshold, a scheduled task can run. In the event that there
	// exists a key/value pair that matches a task argument, PLA will substitute the value
	// as an argument to pass into the execution of that task. For more information about
	// using the Value field as task argument, see section 3.2.4.1.39.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	SetValue(context.Context, *SetValueRequest, ...dcerpc.CallOption) (*SetValueResponse, error)

	// The GetValue method retrieves a user-defined value. The GetValue method retrieves
	// a name and value pair that was set using the SetValue method. This pair can either
	// be used as metadata for the data collector set, in which case it has no effect on
	// the execution of the data collector set, or it can be used as arguments for task
	// execution. For example, when a performance counter crosses a specified threshold,
	// a scheduled task can run. In the event that there exists a key/value pair that matches
	// a task argument, PLA will substitute the value as an argument to pass into the execution
	// of that task. For more information about using the Value field as task argument,
	// see section 3.2.4.1.39.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetValue(context.Context, *GetValueRequest, ...dcerpc.CallOption) (*GetValueResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) DataCollectorSetClient
}

type xxx_DefaultDataCollectorSetClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultDataCollectorSetClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultDataCollectorSetClient) GetDataCollectors(ctx context.Context, in *GetDataCollectorsRequest, opts ...dcerpc.CallOption) (*GetDataCollectorsResponse, error) {
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
	out := &GetDataCollectorsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) GetDuration(ctx context.Context, in *GetDurationRequest, opts ...dcerpc.CallOption) (*GetDurationResponse, error) {
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
	out := &GetDurationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) SetDuration(ctx context.Context, in *SetDurationRequest, opts ...dcerpc.CallOption) (*SetDurationResponse, error) {
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
	out := &SetDurationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) GetDescription(ctx context.Context, in *GetDescriptionRequest, opts ...dcerpc.CallOption) (*GetDescriptionResponse, error) {
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
	out := &GetDescriptionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) SetDescription(ctx context.Context, in *SetDescriptionRequest, opts ...dcerpc.CallOption) (*SetDescriptionResponse, error) {
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
	out := &SetDescriptionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) GetDescriptionUnresolved(ctx context.Context, in *GetDescriptionUnresolvedRequest, opts ...dcerpc.CallOption) (*GetDescriptionUnresolvedResponse, error) {
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
	out := &GetDescriptionUnresolvedResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) GetDisplayName(ctx context.Context, in *GetDisplayNameRequest, opts ...dcerpc.CallOption) (*GetDisplayNameResponse, error) {
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
	out := &GetDisplayNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) SetDisplayName(ctx context.Context, in *SetDisplayNameRequest, opts ...dcerpc.CallOption) (*SetDisplayNameResponse, error) {
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
	out := &SetDisplayNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) GetDisplayNameUnresolved(ctx context.Context, in *GetDisplayNameUnresolvedRequest, opts ...dcerpc.CallOption) (*GetDisplayNameUnresolvedResponse, error) {
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
	out := &GetDisplayNameUnresolvedResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) GetKeywords(ctx context.Context, in *GetKeywordsRequest, opts ...dcerpc.CallOption) (*GetKeywordsResponse, error) {
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
	out := &GetKeywordsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) SetKeywords(ctx context.Context, in *SetKeywordsRequest, opts ...dcerpc.CallOption) (*SetKeywordsResponse, error) {
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
	out := &SetKeywordsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) GetLatestOutputLocation(ctx context.Context, in *GetLatestOutputLocationRequest, opts ...dcerpc.CallOption) (*GetLatestOutputLocationResponse, error) {
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
	out := &GetLatestOutputLocationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) SetLatestOutputLocation(ctx context.Context, in *SetLatestOutputLocationRequest, opts ...dcerpc.CallOption) (*SetLatestOutputLocationResponse, error) {
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
	out := &SetLatestOutputLocationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) GetName(ctx context.Context, in *GetNameRequest, opts ...dcerpc.CallOption) (*GetNameResponse, error) {
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
	out := &GetNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) GetOutputLocation(ctx context.Context, in *GetOutputLocationRequest, opts ...dcerpc.CallOption) (*GetOutputLocationResponse, error) {
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
	out := &GetOutputLocationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) GetRootPath(ctx context.Context, in *GetRootPathRequest, opts ...dcerpc.CallOption) (*GetRootPathResponse, error) {
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
	out := &GetRootPathResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) SetRootPath(ctx context.Context, in *SetRootPathRequest, opts ...dcerpc.CallOption) (*SetRootPathResponse, error) {
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
	out := &SetRootPathResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) GetSegment(ctx context.Context, in *GetSegmentRequest, opts ...dcerpc.CallOption) (*GetSegmentResponse, error) {
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
	out := &GetSegmentResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) SetSegment(ctx context.Context, in *SetSegmentRequest, opts ...dcerpc.CallOption) (*SetSegmentResponse, error) {
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
	out := &SetSegmentResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) GetSegmentMaxDuration(ctx context.Context, in *GetSegmentMaxDurationRequest, opts ...dcerpc.CallOption) (*GetSegmentMaxDurationResponse, error) {
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
	out := &GetSegmentMaxDurationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) SetSegmentMaxDuration(ctx context.Context, in *SetSegmentMaxDurationRequest, opts ...dcerpc.CallOption) (*SetSegmentMaxDurationResponse, error) {
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
	out := &SetSegmentMaxDurationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) GetSegmentMaxSize(ctx context.Context, in *GetSegmentMaxSizeRequest, opts ...dcerpc.CallOption) (*GetSegmentMaxSizeResponse, error) {
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
	out := &GetSegmentMaxSizeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) SetSegmentMaxSize(ctx context.Context, in *SetSegmentMaxSizeRequest, opts ...dcerpc.CallOption) (*SetSegmentMaxSizeResponse, error) {
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
	out := &SetSegmentMaxSizeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) GetSerialNumber(ctx context.Context, in *GetSerialNumberRequest, opts ...dcerpc.CallOption) (*GetSerialNumberResponse, error) {
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
	out := &GetSerialNumberResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) SetSerialNumber(ctx context.Context, in *SetSerialNumberRequest, opts ...dcerpc.CallOption) (*SetSerialNumberResponse, error) {
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
	out := &SetSerialNumberResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) GetServer(ctx context.Context, in *GetServerRequest, opts ...dcerpc.CallOption) (*GetServerResponse, error) {
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
	out := &GetServerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) GetStatus(ctx context.Context, in *GetStatusRequest, opts ...dcerpc.CallOption) (*GetStatusResponse, error) {
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
	out := &GetStatusResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) GetSubdirectory(ctx context.Context, in *GetSubdirectoryRequest, opts ...dcerpc.CallOption) (*GetSubdirectoryResponse, error) {
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
	out := &GetSubdirectoryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) SetSubdirectory(ctx context.Context, in *SetSubdirectoryRequest, opts ...dcerpc.CallOption) (*SetSubdirectoryResponse, error) {
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
	out := &SetSubdirectoryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) GetSubdirectoryFormat(ctx context.Context, in *GetSubdirectoryFormatRequest, opts ...dcerpc.CallOption) (*GetSubdirectoryFormatResponse, error) {
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
	out := &GetSubdirectoryFormatResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) SetSubdirectoryFormat(ctx context.Context, in *SetSubdirectoryFormatRequest, opts ...dcerpc.CallOption) (*SetSubdirectoryFormatResponse, error) {
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
	out := &SetSubdirectoryFormatResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) GetSubdirectoryFormatPattern(ctx context.Context, in *GetSubdirectoryFormatPatternRequest, opts ...dcerpc.CallOption) (*GetSubdirectoryFormatPatternResponse, error) {
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
	out := &GetSubdirectoryFormatPatternResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) SetSubdirectoryFormatPattern(ctx context.Context, in *SetSubdirectoryFormatPatternRequest, opts ...dcerpc.CallOption) (*SetSubdirectoryFormatPatternResponse, error) {
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
	out := &SetSubdirectoryFormatPatternResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) GetTask(ctx context.Context, in *GetTaskRequest, opts ...dcerpc.CallOption) (*GetTaskResponse, error) {
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
	out := &GetTaskResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) SetTask(ctx context.Context, in *SetTaskRequest, opts ...dcerpc.CallOption) (*SetTaskResponse, error) {
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
	out := &SetTaskResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) GetTaskRunAsSelf(ctx context.Context, in *GetTaskRunAsSelfRequest, opts ...dcerpc.CallOption) (*GetTaskRunAsSelfResponse, error) {
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
	out := &GetTaskRunAsSelfResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) SetTaskRunAsSelf(ctx context.Context, in *SetTaskRunAsSelfRequest, opts ...dcerpc.CallOption) (*SetTaskRunAsSelfResponse, error) {
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
	out := &SetTaskRunAsSelfResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) GetTaskArguments(ctx context.Context, in *GetTaskArgumentsRequest, opts ...dcerpc.CallOption) (*GetTaskArgumentsResponse, error) {
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
	out := &GetTaskArgumentsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) SetTaskArguments(ctx context.Context, in *SetTaskArgumentsRequest, opts ...dcerpc.CallOption) (*SetTaskArgumentsResponse, error) {
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
	out := &SetTaskArgumentsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) GetTaskUserTextArguments(ctx context.Context, in *GetTaskUserTextArgumentsRequest, opts ...dcerpc.CallOption) (*GetTaskUserTextArgumentsResponse, error) {
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
	out := &GetTaskUserTextArgumentsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) SetTaskUserTextArguments(ctx context.Context, in *SetTaskUserTextArgumentsRequest, opts ...dcerpc.CallOption) (*SetTaskUserTextArgumentsResponse, error) {
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
	out := &SetTaskUserTextArgumentsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) GetSchedules(ctx context.Context, in *GetSchedulesRequest, opts ...dcerpc.CallOption) (*GetSchedulesResponse, error) {
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
	out := &GetSchedulesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) GetSchedulesEnabled(ctx context.Context, in *GetSchedulesEnabledRequest, opts ...dcerpc.CallOption) (*GetSchedulesEnabledResponse, error) {
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
	out := &GetSchedulesEnabledResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) SetSchedulesEnabled(ctx context.Context, in *SetSchedulesEnabledRequest, opts ...dcerpc.CallOption) (*SetSchedulesEnabledResponse, error) {
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
	out := &SetSchedulesEnabledResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) GetUserAccount(ctx context.Context, in *GetUserAccountRequest, opts ...dcerpc.CallOption) (*GetUserAccountResponse, error) {
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
	out := &GetUserAccountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) GetXML(ctx context.Context, in *GetXMLRequest, opts ...dcerpc.CallOption) (*GetXMLResponse, error) {
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
	out := &GetXMLResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) GetSecurity(ctx context.Context, in *GetSecurityRequest, opts ...dcerpc.CallOption) (*GetSecurityResponse, error) {
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
	out := &GetSecurityResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) SetSecurity(ctx context.Context, in *SetSecurityRequest, opts ...dcerpc.CallOption) (*SetSecurityResponse, error) {
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
	out := &SetSecurityResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) GetStopOnCompletion(ctx context.Context, in *GetStopOnCompletionRequest, opts ...dcerpc.CallOption) (*GetStopOnCompletionResponse, error) {
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
	out := &GetStopOnCompletionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) SetStopOnCompletion(ctx context.Context, in *SetStopOnCompletionRequest, opts ...dcerpc.CallOption) (*SetStopOnCompletionResponse, error) {
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
	out := &SetStopOnCompletionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) GetDataManager(ctx context.Context, in *GetDataManagerRequest, opts ...dcerpc.CallOption) (*GetDataManagerResponse, error) {
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
	out := &GetDataManagerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) SetCredentials(ctx context.Context, in *SetCredentialsRequest, opts ...dcerpc.CallOption) (*SetCredentialsResponse, error) {
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
	out := &SetCredentialsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) Query(ctx context.Context, in *QueryRequest, opts ...dcerpc.CallOption) (*QueryResponse, error) {
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
	out := &QueryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) Commit(ctx context.Context, in *CommitRequest, opts ...dcerpc.CallOption) (*CommitResponse, error) {
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
	out := &CommitResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) Delete(ctx context.Context, in *DeleteRequest, opts ...dcerpc.CallOption) (*DeleteResponse, error) {
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
	out := &DeleteResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) Start(ctx context.Context, in *StartRequest, opts ...dcerpc.CallOption) (*StartResponse, error) {
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
	out := &StartResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) Stop(ctx context.Context, in *StopRequest, opts ...dcerpc.CallOption) (*StopResponse, error) {
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
	out := &StopResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) SetXML(ctx context.Context, in *SetXMLRequest, opts ...dcerpc.CallOption) (*SetXMLResponse, error) {
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
	out := &SetXMLResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) SetValue(ctx context.Context, in *SetValueRequest, opts ...dcerpc.CallOption) (*SetValueResponse, error) {
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
	out := &SetValueResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) GetValue(ctx context.Context, in *GetValueRequest, opts ...dcerpc.CallOption) (*GetValueResponse, error) {
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
	out := &GetValueResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorSetClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultDataCollectorSetClient) IPID(ctx context.Context, ipid *dcom.IPID) DataCollectorSetClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultDataCollectorSetClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}
func NewDataCollectorSetClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (DataCollectorSetClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(DataCollectorSetSyntaxV0_0))...)
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
	return &xxx_DefaultDataCollectorSetClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_GetDataCollectorsOperation structure represents the DataCollectors operation
type xxx_GetDataCollectorsOperation struct {
	This       *dcom.ORPCThis               `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat               `idl:"name:That" json:"that"`
	Collectors *pla.DataCollectorCollection `idl:"name:collectors" json:"collectors"`
	Return     int32                        `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDataCollectorsOperation) OpNum() int { return 7 }

func (o *xxx_GetDataCollectorsOperation) OpName() string {
	return "/IDataCollectorSet/v0/DataCollectors"
}

func (o *xxx_GetDataCollectorsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDataCollectorsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetDataCollectorsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetDataCollectorsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDataCollectorsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// collectors {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IDataCollectorCollection}(interface))
	{
		if o.Collectors != nil {
			_ptr_collectors := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Collectors != nil {
					if err := o.Collectors.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&pla.DataCollectorCollection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Collectors, _ptr_collectors); err != nil {
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

func (o *xxx_GetDataCollectorsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// collectors {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IDataCollectorCollection}(interface))
	{
		_ptr_collectors := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Collectors == nil {
				o.Collectors = &pla.DataCollectorCollection{}
			}
			if err := o.Collectors.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_collectors := func(ptr interface{}) { o.Collectors = *ptr.(**pla.DataCollectorCollection) }
		if err := w.ReadPointer(&o.Collectors, _s_collectors, _ptr_collectors); err != nil {
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

// GetDataCollectorsRequest structure represents the DataCollectors operation request
type GetDataCollectorsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetDataCollectorsRequest) xxx_ToOp(ctx context.Context) *xxx_GetDataCollectorsOperation {
	if o == nil {
		return &xxx_GetDataCollectorsOperation{}
	}
	return &xxx_GetDataCollectorsOperation{
		This: o.This,
	}
}

func (o *GetDataCollectorsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDataCollectorsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetDataCollectorsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetDataCollectorsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDataCollectorsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDataCollectorsResponse structure represents the DataCollectors operation response
type GetDataCollectorsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// collectors: Receives the pointer to the data collector collection object.
	Collectors *pla.DataCollectorCollection `idl:"name:collectors" json:"collectors"`
	// Return: The DataCollectors return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDataCollectorsResponse) xxx_ToOp(ctx context.Context) *xxx_GetDataCollectorsOperation {
	if o == nil {
		return &xxx_GetDataCollectorsOperation{}
	}
	return &xxx_GetDataCollectorsOperation{
		That:       o.That,
		Collectors: o.Collectors,
		Return:     o.Return,
	}
}

func (o *GetDataCollectorsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDataCollectorsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Collectors = op.Collectors
	o.Return = op.Return
}
func (o *GetDataCollectorsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetDataCollectorsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDataCollectorsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDurationOperation structure represents the Duration operation
type xxx_GetDurationOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Seconds uint32         `idl:"name:seconds" json:"seconds"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDurationOperation) OpNum() int { return 8 }

func (o *xxx_GetDurationOperation) OpName() string { return "/IDataCollectorSet/v0/Duration" }

func (o *xxx_GetDurationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDurationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetDurationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetDurationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDurationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// seconds {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Seconds); err != nil {
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

func (o *xxx_GetDurationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// seconds {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Seconds); err != nil {
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

// GetDurationRequest structure represents the Duration operation request
type GetDurationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetDurationRequest) xxx_ToOp(ctx context.Context) *xxx_GetDurationOperation {
	if o == nil {
		return &xxx_GetDurationOperation{}
	}
	return &xxx_GetDurationOperation{
		This: o.This,
	}
}

func (o *GetDurationRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDurationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetDurationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetDurationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDurationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDurationResponse structure represents the Duration operation response
type GetDurationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Seconds uint32         `idl:"name:seconds" json:"seconds"`
	// Return: The Duration return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDurationResponse) xxx_ToOp(ctx context.Context) *xxx_GetDurationOperation {
	if o == nil {
		return &xxx_GetDurationOperation{}
	}
	return &xxx_GetDurationOperation{
		That:    o.That,
		Seconds: o.Seconds,
		Return:  o.Return,
	}
}

func (o *GetDurationResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDurationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Seconds = op.Seconds
	o.Return = op.Return
}
func (o *GetDurationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetDurationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDurationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetDurationOperation structure represents the Duration operation
type xxx_SetDurationOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Seconds uint32         `idl:"name:seconds" json:"seconds"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetDurationOperation) OpNum() int { return 9 }

func (o *xxx_SetDurationOperation) OpName() string { return "/IDataCollectorSet/v0/Duration" }

func (o *xxx_SetDurationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDurationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// seconds {in} (1:(uint32))
	{
		if err := w.WriteData(o.Seconds); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDurationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// seconds {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Seconds); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDurationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDurationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetDurationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetDurationRequest structure represents the Duration operation request
type SetDurationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	Seconds uint32         `idl:"name:seconds" json:"seconds"`
}

func (o *SetDurationRequest) xxx_ToOp(ctx context.Context) *xxx_SetDurationOperation {
	if o == nil {
		return &xxx_SetDurationOperation{}
	}
	return &xxx_SetDurationOperation{
		This:    o.This,
		Seconds: o.Seconds,
	}
}

func (o *SetDurationRequest) xxx_FromOp(ctx context.Context, op *xxx_SetDurationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Seconds = op.Seconds
}
func (o *SetDurationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetDurationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDurationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetDurationResponse structure represents the Duration operation response
type SetDurationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Duration return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetDurationResponse) xxx_ToOp(ctx context.Context) *xxx_SetDurationOperation {
	if o == nil {
		return &xxx_SetDurationOperation{}
	}
	return &xxx_SetDurationOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetDurationResponse) xxx_FromOp(ctx context.Context, op *xxx_SetDurationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetDurationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetDurationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDurationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDescriptionOperation structure represents the Description operation
type xxx_GetDescriptionOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	Description *oaut.String   `idl:"name:description" json:"description"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDescriptionOperation) OpNum() int { return 10 }

func (o *xxx_GetDescriptionOperation) OpName() string { return "/IDataCollectorSet/v0/Description" }

func (o *xxx_GetDescriptionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDescriptionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetDescriptionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetDescriptionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDescriptionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// description {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Description != nil {
			_ptr_description := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Description != nil {
					if err := o.Description.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Description, _ptr_description); err != nil {
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

func (o *xxx_GetDescriptionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// description {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_description := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Description == nil {
				o.Description = &oaut.String{}
			}
			if err := o.Description.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_description := func(ptr interface{}) { o.Description = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Description, _s_description, _ptr_description); err != nil {
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

// GetDescriptionRequest structure represents the Description operation request
type GetDescriptionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetDescriptionRequest) xxx_ToOp(ctx context.Context) *xxx_GetDescriptionOperation {
	if o == nil {
		return &xxx_GetDescriptionOperation{}
	}
	return &xxx_GetDescriptionOperation{
		This: o.This,
	}
}

func (o *GetDescriptionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDescriptionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetDescriptionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetDescriptionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDescriptionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDescriptionResponse structure represents the Description operation response
type GetDescriptionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	Description *oaut.String   `idl:"name:description" json:"description"`
	// Return: The Description return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDescriptionResponse) xxx_ToOp(ctx context.Context) *xxx_GetDescriptionOperation {
	if o == nil {
		return &xxx_GetDescriptionOperation{}
	}
	return &xxx_GetDescriptionOperation{
		That:        o.That,
		Description: o.Description,
		Return:      o.Return,
	}
}

func (o *GetDescriptionResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDescriptionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Description = op.Description
	o.Return = op.Return
}
func (o *GetDescriptionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetDescriptionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDescriptionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetDescriptionOperation structure represents the Description operation
type xxx_SetDescriptionOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	Description *oaut.String   `idl:"name:description" json:"description"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetDescriptionOperation) OpNum() int { return 11 }

func (o *xxx_SetDescriptionOperation) OpName() string { return "/IDataCollectorSet/v0/Description" }

func (o *xxx_SetDescriptionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDescriptionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// description {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Description != nil {
			_ptr_description := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Description != nil {
					if err := o.Description.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Description, _ptr_description); err != nil {
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

func (o *xxx_SetDescriptionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// description {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_description := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Description == nil {
				o.Description = &oaut.String{}
			}
			if err := o.Description.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_description := func(ptr interface{}) { o.Description = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Description, _s_description, _ptr_description); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDescriptionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDescriptionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetDescriptionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetDescriptionRequest structure represents the Description operation request
type SetDescriptionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	Description *oaut.String   `idl:"name:description" json:"description"`
}

func (o *SetDescriptionRequest) xxx_ToOp(ctx context.Context) *xxx_SetDescriptionOperation {
	if o == nil {
		return &xxx_SetDescriptionOperation{}
	}
	return &xxx_SetDescriptionOperation{
		This:        o.This,
		Description: o.Description,
	}
}

func (o *SetDescriptionRequest) xxx_FromOp(ctx context.Context, op *xxx_SetDescriptionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Description = op.Description
}
func (o *SetDescriptionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetDescriptionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDescriptionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetDescriptionResponse structure represents the Description operation response
type SetDescriptionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Description return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetDescriptionResponse) xxx_ToOp(ctx context.Context) *xxx_SetDescriptionOperation {
	if o == nil {
		return &xxx_SetDescriptionOperation{}
	}
	return &xxx_SetDescriptionOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetDescriptionResponse) xxx_FromOp(ctx context.Context, op *xxx_SetDescriptionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetDescriptionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetDescriptionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDescriptionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDescriptionUnresolvedOperation structure represents the DescriptionUnresolved operation
type xxx_GetDescriptionUnresolvedOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	Description *oaut.String   `idl:"name:Descr" json:"description"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDescriptionUnresolvedOperation) OpNum() int { return 12 }

func (o *xxx_GetDescriptionUnresolvedOperation) OpName() string {
	return "/IDataCollectorSet/v0/DescriptionUnresolved"
}

func (o *xxx_GetDescriptionUnresolvedOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDescriptionUnresolvedOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetDescriptionUnresolvedOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetDescriptionUnresolvedOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDescriptionUnresolvedOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// Descr {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Description != nil {
			_ptr_Descr := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Description != nil {
					if err := o.Description.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Description, _ptr_Descr); err != nil {
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

func (o *xxx_GetDescriptionUnresolvedOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// Descr {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_Descr := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Description == nil {
				o.Description = &oaut.String{}
			}
			if err := o.Description.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Descr := func(ptr interface{}) { o.Description = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Description, _s_Descr, _ptr_Descr); err != nil {
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

// GetDescriptionUnresolvedRequest structure represents the DescriptionUnresolved operation request
type GetDescriptionUnresolvedRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetDescriptionUnresolvedRequest) xxx_ToOp(ctx context.Context) *xxx_GetDescriptionUnresolvedOperation {
	if o == nil {
		return &xxx_GetDescriptionUnresolvedOperation{}
	}
	return &xxx_GetDescriptionUnresolvedOperation{
		This: o.This,
	}
}

func (o *GetDescriptionUnresolvedRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDescriptionUnresolvedOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetDescriptionUnresolvedRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetDescriptionUnresolvedRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDescriptionUnresolvedOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDescriptionUnresolvedResponse structure represents the DescriptionUnresolved operation response
type GetDescriptionUnresolvedResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Descr: Receives the description of the data collector set in its raw form.
	Description *oaut.String `idl:"name:Descr" json:"description"`
	// Return: The DescriptionUnresolved return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDescriptionUnresolvedResponse) xxx_ToOp(ctx context.Context) *xxx_GetDescriptionUnresolvedOperation {
	if o == nil {
		return &xxx_GetDescriptionUnresolvedOperation{}
	}
	return &xxx_GetDescriptionUnresolvedOperation{
		That:        o.That,
		Description: o.Description,
		Return:      o.Return,
	}
}

func (o *GetDescriptionUnresolvedResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDescriptionUnresolvedOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Description = op.Description
	o.Return = op.Return
}
func (o *GetDescriptionUnresolvedResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetDescriptionUnresolvedResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDescriptionUnresolvedOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDisplayNameOperation structure represents the DisplayName operation
type xxx_GetDisplayNameOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	DisplayName *oaut.String   `idl:"name:DisplayName" json:"display_name"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDisplayNameOperation) OpNum() int { return 13 }

func (o *xxx_GetDisplayNameOperation) OpName() string { return "/IDataCollectorSet/v0/DisplayName" }

func (o *xxx_GetDisplayNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDisplayNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetDisplayNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetDisplayNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDisplayNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// DisplayName {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.DisplayName != nil {
			_ptr_DisplayName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DisplayName != nil {
					if err := o.DisplayName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DisplayName, _ptr_DisplayName); err != nil {
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

func (o *xxx_GetDisplayNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// DisplayName {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_DisplayName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DisplayName == nil {
				o.DisplayName = &oaut.String{}
			}
			if err := o.DisplayName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_DisplayName := func(ptr interface{}) { o.DisplayName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.DisplayName, _s_DisplayName, _ptr_DisplayName); err != nil {
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

// GetDisplayNameRequest structure represents the DisplayName operation request
type GetDisplayNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetDisplayNameRequest) xxx_ToOp(ctx context.Context) *xxx_GetDisplayNameOperation {
	if o == nil {
		return &xxx_GetDisplayNameOperation{}
	}
	return &xxx_GetDisplayNameOperation{
		This: o.This,
	}
}

func (o *GetDisplayNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDisplayNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetDisplayNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetDisplayNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDisplayNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDisplayNameResponse structure represents the DisplayName operation response
type GetDisplayNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	DisplayName *oaut.String   `idl:"name:DisplayName" json:"display_name"`
	// Return: The DisplayName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDisplayNameResponse) xxx_ToOp(ctx context.Context) *xxx_GetDisplayNameOperation {
	if o == nil {
		return &xxx_GetDisplayNameOperation{}
	}
	return &xxx_GetDisplayNameOperation{
		That:        o.That,
		DisplayName: o.DisplayName,
		Return:      o.Return,
	}
}

func (o *GetDisplayNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDisplayNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DisplayName = op.DisplayName
	o.Return = op.Return
}
func (o *GetDisplayNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetDisplayNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDisplayNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetDisplayNameOperation structure represents the DisplayName operation
type xxx_SetDisplayNameOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	DisplayName *oaut.String   `idl:"name:DisplayName" json:"display_name"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetDisplayNameOperation) OpNum() int { return 14 }

func (o *xxx_SetDisplayNameOperation) OpName() string { return "/IDataCollectorSet/v0/DisplayName" }

func (o *xxx_SetDisplayNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDisplayNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// DisplayName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.DisplayName != nil {
			_ptr_DisplayName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DisplayName != nil {
					if err := o.DisplayName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DisplayName, _ptr_DisplayName); err != nil {
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

func (o *xxx_SetDisplayNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// DisplayName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_DisplayName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DisplayName == nil {
				o.DisplayName = &oaut.String{}
			}
			if err := o.DisplayName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_DisplayName := func(ptr interface{}) { o.DisplayName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.DisplayName, _s_DisplayName, _ptr_DisplayName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDisplayNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDisplayNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetDisplayNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetDisplayNameRequest structure represents the DisplayName operation request
type SetDisplayNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	DisplayName *oaut.String   `idl:"name:DisplayName" json:"display_name"`
}

func (o *SetDisplayNameRequest) xxx_ToOp(ctx context.Context) *xxx_SetDisplayNameOperation {
	if o == nil {
		return &xxx_SetDisplayNameOperation{}
	}
	return &xxx_SetDisplayNameOperation{
		This:        o.This,
		DisplayName: o.DisplayName,
	}
}

func (o *SetDisplayNameRequest) xxx_FromOp(ctx context.Context, op *xxx_SetDisplayNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DisplayName = op.DisplayName
}
func (o *SetDisplayNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetDisplayNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDisplayNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetDisplayNameResponse structure represents the DisplayName operation response
type SetDisplayNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DisplayName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetDisplayNameResponse) xxx_ToOp(ctx context.Context) *xxx_SetDisplayNameOperation {
	if o == nil {
		return &xxx_SetDisplayNameOperation{}
	}
	return &xxx_SetDisplayNameOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetDisplayNameResponse) xxx_FromOp(ctx context.Context, op *xxx_SetDisplayNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetDisplayNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetDisplayNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDisplayNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDisplayNameUnresolvedOperation structure represents the DisplayNameUnresolved operation
type xxx_GetDisplayNameUnresolvedOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Name   *oaut.String   `idl:"name:name" json:"name"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDisplayNameUnresolvedOperation) OpNum() int { return 15 }

func (o *xxx_GetDisplayNameUnresolvedOperation) OpName() string {
	return "/IDataCollectorSet/v0/DisplayNameUnresolved"
}

func (o *xxx_GetDisplayNameUnresolvedOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDisplayNameUnresolvedOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetDisplayNameUnresolvedOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetDisplayNameUnresolvedOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDisplayNameUnresolvedOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// name {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Name != nil {
			_ptr_name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Name != nil {
					if err := o.Name.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Name, _ptr_name); err != nil {
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

func (o *xxx_GetDisplayNameUnresolvedOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// name {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Name == nil {
				o.Name = &oaut.String{}
			}
			if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_name := func(ptr interface{}) { o.Name = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Name, _s_name, _ptr_name); err != nil {
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

// GetDisplayNameUnresolvedRequest structure represents the DisplayNameUnresolved operation request
type GetDisplayNameUnresolvedRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetDisplayNameUnresolvedRequest) xxx_ToOp(ctx context.Context) *xxx_GetDisplayNameUnresolvedOperation {
	if o == nil {
		return &xxx_GetDisplayNameUnresolvedOperation{}
	}
	return &xxx_GetDisplayNameUnresolvedOperation{
		This: o.This,
	}
}

func (o *GetDisplayNameUnresolvedRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDisplayNameUnresolvedOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetDisplayNameUnresolvedRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetDisplayNameUnresolvedRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDisplayNameUnresolvedOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDisplayNameUnresolvedResponse structure represents the DisplayNameUnresolved operation response
type GetDisplayNameUnresolvedResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// name: Receives the display name of the data collector set in its original form.
	Name *oaut.String `idl:"name:name" json:"name"`
	// Return: The DisplayNameUnresolved return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDisplayNameUnresolvedResponse) xxx_ToOp(ctx context.Context) *xxx_GetDisplayNameUnresolvedOperation {
	if o == nil {
		return &xxx_GetDisplayNameUnresolvedOperation{}
	}
	return &xxx_GetDisplayNameUnresolvedOperation{
		That:   o.That,
		Name:   o.Name,
		Return: o.Return,
	}
}

func (o *GetDisplayNameUnresolvedResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDisplayNameUnresolvedOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Name = op.Name
	o.Return = op.Return
}
func (o *GetDisplayNameUnresolvedResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetDisplayNameUnresolvedResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDisplayNameUnresolvedOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetKeywordsOperation structure represents the Keywords operation
type xxx_GetKeywordsOperation struct {
	This     *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Keywords *oaut.SafeArray `idl:"name:keywords" json:"keywords"`
	Return   int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetKeywordsOperation) OpNum() int { return 16 }

func (o *xxx_GetKeywordsOperation) OpName() string { return "/IDataCollectorSet/v0/Keywords" }

func (o *xxx_GetKeywordsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetKeywordsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetKeywordsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetKeywordsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetKeywordsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// keywords {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.Keywords != nil {
			_ptr_keywords := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Keywords != nil {
					if err := o.Keywords.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Keywords, _ptr_keywords); err != nil {
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

func (o *xxx_GetKeywordsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// keywords {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_keywords := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Keywords == nil {
				o.Keywords = &oaut.SafeArray{}
			}
			if err := o.Keywords.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_keywords := func(ptr interface{}) { o.Keywords = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.Keywords, _s_keywords, _ptr_keywords); err != nil {
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

// GetKeywordsRequest structure represents the Keywords operation request
type GetKeywordsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetKeywordsRequest) xxx_ToOp(ctx context.Context) *xxx_GetKeywordsOperation {
	if o == nil {
		return &xxx_GetKeywordsOperation{}
	}
	return &xxx_GetKeywordsOperation{
		This: o.This,
	}
}

func (o *GetKeywordsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetKeywordsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetKeywordsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetKeywordsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetKeywordsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetKeywordsResponse structure represents the Keywords operation response
type GetKeywordsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Keywords *oaut.SafeArray `idl:"name:keywords" json:"keywords"`
	// Return: The Keywords return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetKeywordsResponse) xxx_ToOp(ctx context.Context) *xxx_GetKeywordsOperation {
	if o == nil {
		return &xxx_GetKeywordsOperation{}
	}
	return &xxx_GetKeywordsOperation{
		That:     o.That,
		Keywords: o.Keywords,
		Return:   o.Return,
	}
}

func (o *GetKeywordsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetKeywordsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Keywords = op.Keywords
	o.Return = op.Return
}
func (o *GetKeywordsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetKeywordsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetKeywordsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetKeywordsOperation structure represents the Keywords operation
type xxx_SetKeywordsOperation struct {
	This     *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Keywords *oaut.SafeArray `idl:"name:keywords" json:"keywords"`
	Return   int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_SetKeywordsOperation) OpNum() int { return 17 }

func (o *xxx_SetKeywordsOperation) OpName() string { return "/IDataCollectorSet/v0/Keywords" }

func (o *xxx_SetKeywordsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetKeywordsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// keywords {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.Keywords != nil {
			_ptr_keywords := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Keywords != nil {
					if err := o.Keywords.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Keywords, _ptr_keywords); err != nil {
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

func (o *xxx_SetKeywordsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// keywords {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_keywords := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Keywords == nil {
				o.Keywords = &oaut.SafeArray{}
			}
			if err := o.Keywords.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_keywords := func(ptr interface{}) { o.Keywords = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.Keywords, _s_keywords, _ptr_keywords); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetKeywordsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetKeywordsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetKeywordsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetKeywordsRequest structure represents the Keywords operation request
type SetKeywordsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis  `idl:"name:This" json:"this"`
	Keywords *oaut.SafeArray `idl:"name:keywords" json:"keywords"`
}

func (o *SetKeywordsRequest) xxx_ToOp(ctx context.Context) *xxx_SetKeywordsOperation {
	if o == nil {
		return &xxx_SetKeywordsOperation{}
	}
	return &xxx_SetKeywordsOperation{
		This:     o.This,
		Keywords: o.Keywords,
	}
}

func (o *SetKeywordsRequest) xxx_FromOp(ctx context.Context, op *xxx_SetKeywordsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Keywords = op.Keywords
}
func (o *SetKeywordsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetKeywordsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetKeywordsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetKeywordsResponse structure represents the Keywords operation response
type SetKeywordsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Keywords return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetKeywordsResponse) xxx_ToOp(ctx context.Context) *xxx_SetKeywordsOperation {
	if o == nil {
		return &xxx_SetKeywordsOperation{}
	}
	return &xxx_SetKeywordsOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetKeywordsResponse) xxx_FromOp(ctx context.Context, op *xxx_SetKeywordsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetKeywordsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetKeywordsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetKeywordsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetLatestOutputLocationOperation structure represents the LatestOutputLocation operation
type xxx_GetLatestOutputLocationOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Path   *oaut.String   `idl:"name:path" json:"path"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetLatestOutputLocationOperation) OpNum() int { return 18 }

func (o *xxx_GetLatestOutputLocationOperation) OpName() string {
	return "/IDataCollectorSet/v0/LatestOutputLocation"
}

func (o *xxx_GetLatestOutputLocationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLatestOutputLocationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetLatestOutputLocationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetLatestOutputLocationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLatestOutputLocationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetLatestOutputLocationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// GetLatestOutputLocationRequest structure represents the LatestOutputLocation operation request
type GetLatestOutputLocationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetLatestOutputLocationRequest) xxx_ToOp(ctx context.Context) *xxx_GetLatestOutputLocationOperation {
	if o == nil {
		return &xxx_GetLatestOutputLocationOperation{}
	}
	return &xxx_GetLatestOutputLocationOperation{
		This: o.This,
	}
}

func (o *GetLatestOutputLocationRequest) xxx_FromOp(ctx context.Context, op *xxx_GetLatestOutputLocationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetLatestOutputLocationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetLatestOutputLocationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLatestOutputLocationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetLatestOutputLocationResponse structure represents the LatestOutputLocation operation response
type GetLatestOutputLocationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	Path *oaut.String   `idl:"name:path" json:"path"`
	// Return: The LatestOutputLocation return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetLatestOutputLocationResponse) xxx_ToOp(ctx context.Context) *xxx_GetLatestOutputLocationOperation {
	if o == nil {
		return &xxx_GetLatestOutputLocationOperation{}
	}
	return &xxx_GetLatestOutputLocationOperation{
		That:   o.That,
		Path:   o.Path,
		Return: o.Return,
	}
}

func (o *GetLatestOutputLocationResponse) xxx_FromOp(ctx context.Context, op *xxx_GetLatestOutputLocationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Path = op.Path
	o.Return = op.Return
}
func (o *GetLatestOutputLocationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetLatestOutputLocationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLatestOutputLocationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetLatestOutputLocationOperation structure represents the LatestOutputLocation operation
type xxx_SetLatestOutputLocationOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Path   *oaut.String   `idl:"name:path" json:"path"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetLatestOutputLocationOperation) OpNum() int { return 19 }

func (o *xxx_SetLatestOutputLocationOperation) OpName() string {
	return "/IDataCollectorSet/v0/LatestOutputLocation"
}

func (o *xxx_SetLatestOutputLocationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLatestOutputLocationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetLatestOutputLocationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_SetLatestOutputLocationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLatestOutputLocationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetLatestOutputLocationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetLatestOutputLocationRequest structure represents the LatestOutputLocation operation request
type SetLatestOutputLocationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	Path *oaut.String   `idl:"name:path" json:"path"`
}

func (o *SetLatestOutputLocationRequest) xxx_ToOp(ctx context.Context) *xxx_SetLatestOutputLocationOperation {
	if o == nil {
		return &xxx_SetLatestOutputLocationOperation{}
	}
	return &xxx_SetLatestOutputLocationOperation{
		This: o.This,
		Path: o.Path,
	}
}

func (o *SetLatestOutputLocationRequest) xxx_FromOp(ctx context.Context, op *xxx_SetLatestOutputLocationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Path = op.Path
}
func (o *SetLatestOutputLocationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetLatestOutputLocationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetLatestOutputLocationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetLatestOutputLocationResponse structure represents the LatestOutputLocation operation response
type SetLatestOutputLocationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The LatestOutputLocation return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetLatestOutputLocationResponse) xxx_ToOp(ctx context.Context) *xxx_SetLatestOutputLocationOperation {
	if o == nil {
		return &xxx_SetLatestOutputLocationOperation{}
	}
	return &xxx_SetLatestOutputLocationOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetLatestOutputLocationResponse) xxx_FromOp(ctx context.Context, op *xxx_SetLatestOutputLocationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetLatestOutputLocationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetLatestOutputLocationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetLatestOutputLocationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNameOperation structure represents the Name operation
type xxx_GetNameOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Name   *oaut.String   `idl:"name:name" json:"name"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNameOperation) OpNum() int { return 20 }

func (o *xxx_GetNameOperation) OpName() string { return "/IDataCollectorSet/v0/Name" }

func (o *xxx_GetNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// name {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Name != nil {
			_ptr_name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Name != nil {
					if err := o.Name.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Name, _ptr_name); err != nil {
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

func (o *xxx_GetNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// name {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Name == nil {
				o.Name = &oaut.String{}
			}
			if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_name := func(ptr interface{}) { o.Name = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Name, _s_name, _ptr_name); err != nil {
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

// GetNameRequest structure represents the Name operation request
type GetNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetNameRequest) xxx_ToOp(ctx context.Context) *xxx_GetNameOperation {
	if o == nil {
		return &xxx_GetNameOperation{}
	}
	return &xxx_GetNameOperation{
		This: o.This,
	}
}

func (o *GetNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNameResponse structure represents the Name operation response
type GetNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	Name *oaut.String   `idl:"name:name" json:"name"`
	// Return: The Name return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNameResponse) xxx_ToOp(ctx context.Context) *xxx_GetNameOperation {
	if o == nil {
		return &xxx_GetNameOperation{}
	}
	return &xxx_GetNameOperation{
		That:   o.That,
		Name:   o.Name,
		Return: o.Return,
	}
}

func (o *GetNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Name = op.Name
	o.Return = op.Return
}
func (o *GetNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetOutputLocationOperation structure represents the OutputLocation operation
type xxx_GetOutputLocationOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Path   *oaut.String   `idl:"name:path" json:"path"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetOutputLocationOperation) OpNum() int { return 21 }

func (o *xxx_GetOutputLocationOperation) OpName() string {
	return "/IDataCollectorSet/v0/OutputLocation"
}

func (o *xxx_GetOutputLocationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOutputLocationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetOutputLocationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetOutputLocationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOutputLocationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetOutputLocationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// GetOutputLocationRequest structure represents the OutputLocation operation request
type GetOutputLocationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetOutputLocationRequest) xxx_ToOp(ctx context.Context) *xxx_GetOutputLocationOperation {
	if o == nil {
		return &xxx_GetOutputLocationOperation{}
	}
	return &xxx_GetOutputLocationOperation{
		This: o.This,
	}
}

func (o *GetOutputLocationRequest) xxx_FromOp(ctx context.Context, op *xxx_GetOutputLocationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetOutputLocationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetOutputLocationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetOutputLocationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetOutputLocationResponse structure represents the OutputLocation operation response
type GetOutputLocationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	Path *oaut.String   `idl:"name:path" json:"path"`
	// Return: The OutputLocation return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetOutputLocationResponse) xxx_ToOp(ctx context.Context) *xxx_GetOutputLocationOperation {
	if o == nil {
		return &xxx_GetOutputLocationOperation{}
	}
	return &xxx_GetOutputLocationOperation{
		That:   o.That,
		Path:   o.Path,
		Return: o.Return,
	}
}

func (o *GetOutputLocationResponse) xxx_FromOp(ctx context.Context, op *xxx_GetOutputLocationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Path = op.Path
	o.Return = op.Return
}
func (o *GetOutputLocationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetOutputLocationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetOutputLocationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetRootPathOperation structure represents the RootPath operation
type xxx_GetRootPathOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Folder *oaut.String   `idl:"name:folder" json:"folder"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetRootPathOperation) OpNum() int { return 22 }

func (o *xxx_GetRootPathOperation) OpName() string { return "/IDataCollectorSet/v0/RootPath" }

func (o *xxx_GetRootPathOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRootPathOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetRootPathOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetRootPathOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRootPathOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// folder {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Folder != nil {
			_ptr_folder := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Folder != nil {
					if err := o.Folder.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Folder, _ptr_folder); err != nil {
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

func (o *xxx_GetRootPathOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// folder {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_folder := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Folder == nil {
				o.Folder = &oaut.String{}
			}
			if err := o.Folder.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_folder := func(ptr interface{}) { o.Folder = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Folder, _s_folder, _ptr_folder); err != nil {
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

// GetRootPathRequest structure represents the RootPath operation request
type GetRootPathRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetRootPathRequest) xxx_ToOp(ctx context.Context) *xxx_GetRootPathOperation {
	if o == nil {
		return &xxx_GetRootPathOperation{}
	}
	return &xxx_GetRootPathOperation{
		This: o.This,
	}
}

func (o *GetRootPathRequest) xxx_FromOp(ctx context.Context, op *xxx_GetRootPathOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetRootPathRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetRootPathRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRootPathOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetRootPathResponse structure represents the RootPath operation response
type GetRootPathResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Folder *oaut.String   `idl:"name:folder" json:"folder"`
	// Return: The RootPath return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetRootPathResponse) xxx_ToOp(ctx context.Context) *xxx_GetRootPathOperation {
	if o == nil {
		return &xxx_GetRootPathOperation{}
	}
	return &xxx_GetRootPathOperation{
		That:   o.That,
		Folder: o.Folder,
		Return: o.Return,
	}
}

func (o *GetRootPathResponse) xxx_FromOp(ctx context.Context, op *xxx_GetRootPathOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Folder = op.Folder
	o.Return = op.Return
}
func (o *GetRootPathResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetRootPathResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRootPathOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetRootPathOperation structure represents the RootPath operation
type xxx_SetRootPathOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Folder *oaut.String   `idl:"name:folder" json:"folder"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetRootPathOperation) OpNum() int { return 23 }

func (o *xxx_SetRootPathOperation) OpName() string { return "/IDataCollectorSet/v0/RootPath" }

func (o *xxx_SetRootPathOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetRootPathOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// folder {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Folder != nil {
			_ptr_folder := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Folder != nil {
					if err := o.Folder.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Folder, _ptr_folder); err != nil {
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

func (o *xxx_SetRootPathOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// folder {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_folder := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Folder == nil {
				o.Folder = &oaut.String{}
			}
			if err := o.Folder.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_folder := func(ptr interface{}) { o.Folder = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Folder, _s_folder, _ptr_folder); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetRootPathOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetRootPathOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetRootPathOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetRootPathRequest structure represents the RootPath operation request
type SetRootPathRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	Folder *oaut.String   `idl:"name:folder" json:"folder"`
}

func (o *SetRootPathRequest) xxx_ToOp(ctx context.Context) *xxx_SetRootPathOperation {
	if o == nil {
		return &xxx_SetRootPathOperation{}
	}
	return &xxx_SetRootPathOperation{
		This:   o.This,
		Folder: o.Folder,
	}
}

func (o *SetRootPathRequest) xxx_FromOp(ctx context.Context, op *xxx_SetRootPathOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Folder = op.Folder
}
func (o *SetRootPathRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetRootPathRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetRootPathOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetRootPathResponse structure represents the RootPath operation response
type SetRootPathResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RootPath return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetRootPathResponse) xxx_ToOp(ctx context.Context) *xxx_SetRootPathOperation {
	if o == nil {
		return &xxx_SetRootPathOperation{}
	}
	return &xxx_SetRootPathOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetRootPathResponse) xxx_FromOp(ctx context.Context, op *xxx_SetRootPathOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetRootPathResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetRootPathResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetRootPathOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSegmentOperation structure represents the Segment operation
type xxx_GetSegmentOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Segment int16          `idl:"name:segment" json:"segment"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSegmentOperation) OpNum() int { return 24 }

func (o *xxx_GetSegmentOperation) OpName() string { return "/IDataCollectorSet/v0/Segment" }

func (o *xxx_GetSegmentOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSegmentOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSegmentOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSegmentOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSegmentOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// segment {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Segment); err != nil {
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

func (o *xxx_GetSegmentOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// segment {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Segment); err != nil {
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

// GetSegmentRequest structure represents the Segment operation request
type GetSegmentRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSegmentRequest) xxx_ToOp(ctx context.Context) *xxx_GetSegmentOperation {
	if o == nil {
		return &xxx_GetSegmentOperation{}
	}
	return &xxx_GetSegmentOperation{
		This: o.This,
	}
}

func (o *GetSegmentRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSegmentOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSegmentRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetSegmentRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSegmentOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSegmentResponse structure represents the Segment operation response
type GetSegmentResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Segment int16          `idl:"name:segment" json:"segment"`
	// Return: The Segment return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSegmentResponse) xxx_ToOp(ctx context.Context) *xxx_GetSegmentOperation {
	if o == nil {
		return &xxx_GetSegmentOperation{}
	}
	return &xxx_GetSegmentOperation{
		That:    o.That,
		Segment: o.Segment,
		Return:  o.Return,
	}
}

func (o *GetSegmentResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSegmentOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Segment = op.Segment
	o.Return = op.Return
}
func (o *GetSegmentResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetSegmentResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSegmentOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSegmentOperation structure represents the Segment operation
type xxx_SetSegmentOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Segment int16          `idl:"name:segment" json:"segment"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSegmentOperation) OpNum() int { return 25 }

func (o *xxx_SetSegmentOperation) OpName() string { return "/IDataCollectorSet/v0/Segment" }

func (o *xxx_SetSegmentOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSegmentOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// segment {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Segment); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSegmentOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// segment {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Segment); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSegmentOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSegmentOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetSegmentOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetSegmentRequest structure represents the Segment operation request
type SetSegmentRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	Segment int16          `idl:"name:segment" json:"segment"`
}

func (o *SetSegmentRequest) xxx_ToOp(ctx context.Context) *xxx_SetSegmentOperation {
	if o == nil {
		return &xxx_SetSegmentOperation{}
	}
	return &xxx_SetSegmentOperation{
		This:    o.This,
		Segment: o.Segment,
	}
}

func (o *SetSegmentRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSegmentOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Segment = op.Segment
}
func (o *SetSegmentRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetSegmentRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSegmentOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSegmentResponse structure represents the Segment operation response
type SetSegmentResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Segment return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSegmentResponse) xxx_ToOp(ctx context.Context) *xxx_SetSegmentOperation {
	if o == nil {
		return &xxx_SetSegmentOperation{}
	}
	return &xxx_SetSegmentOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetSegmentResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSegmentOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSegmentResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetSegmentResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSegmentOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSegmentMaxDurationOperation structure represents the SegmentMaxDuration operation
type xxx_GetSegmentMaxDurationOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Seconds uint32         `idl:"name:seconds" json:"seconds"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSegmentMaxDurationOperation) OpNum() int { return 26 }

func (o *xxx_GetSegmentMaxDurationOperation) OpName() string {
	return "/IDataCollectorSet/v0/SegmentMaxDuration"
}

func (o *xxx_GetSegmentMaxDurationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSegmentMaxDurationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSegmentMaxDurationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSegmentMaxDurationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSegmentMaxDurationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// seconds {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Seconds); err != nil {
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

func (o *xxx_GetSegmentMaxDurationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// seconds {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Seconds); err != nil {
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

// GetSegmentMaxDurationRequest structure represents the SegmentMaxDuration operation request
type GetSegmentMaxDurationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSegmentMaxDurationRequest) xxx_ToOp(ctx context.Context) *xxx_GetSegmentMaxDurationOperation {
	if o == nil {
		return &xxx_GetSegmentMaxDurationOperation{}
	}
	return &xxx_GetSegmentMaxDurationOperation{
		This: o.This,
	}
}

func (o *GetSegmentMaxDurationRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSegmentMaxDurationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSegmentMaxDurationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetSegmentMaxDurationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSegmentMaxDurationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSegmentMaxDurationResponse structure represents the SegmentMaxDuration operation response
type GetSegmentMaxDurationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Seconds uint32         `idl:"name:seconds" json:"seconds"`
	// Return: The SegmentMaxDuration return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSegmentMaxDurationResponse) xxx_ToOp(ctx context.Context) *xxx_GetSegmentMaxDurationOperation {
	if o == nil {
		return &xxx_GetSegmentMaxDurationOperation{}
	}
	return &xxx_GetSegmentMaxDurationOperation{
		That:    o.That,
		Seconds: o.Seconds,
		Return:  o.Return,
	}
}

func (o *GetSegmentMaxDurationResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSegmentMaxDurationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Seconds = op.Seconds
	o.Return = op.Return
}
func (o *GetSegmentMaxDurationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetSegmentMaxDurationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSegmentMaxDurationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSegmentMaxDurationOperation structure represents the SegmentMaxDuration operation
type xxx_SetSegmentMaxDurationOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Seconds uint32         `idl:"name:seconds" json:"seconds"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSegmentMaxDurationOperation) OpNum() int { return 27 }

func (o *xxx_SetSegmentMaxDurationOperation) OpName() string {
	return "/IDataCollectorSet/v0/SegmentMaxDuration"
}

func (o *xxx_SetSegmentMaxDurationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSegmentMaxDurationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// seconds {in} (1:(uint32))
	{
		if err := w.WriteData(o.Seconds); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSegmentMaxDurationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// seconds {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Seconds); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSegmentMaxDurationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSegmentMaxDurationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetSegmentMaxDurationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetSegmentMaxDurationRequest structure represents the SegmentMaxDuration operation request
type SetSegmentMaxDurationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	Seconds uint32         `idl:"name:seconds" json:"seconds"`
}

func (o *SetSegmentMaxDurationRequest) xxx_ToOp(ctx context.Context) *xxx_SetSegmentMaxDurationOperation {
	if o == nil {
		return &xxx_SetSegmentMaxDurationOperation{}
	}
	return &xxx_SetSegmentMaxDurationOperation{
		This:    o.This,
		Seconds: o.Seconds,
	}
}

func (o *SetSegmentMaxDurationRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSegmentMaxDurationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Seconds = op.Seconds
}
func (o *SetSegmentMaxDurationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetSegmentMaxDurationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSegmentMaxDurationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSegmentMaxDurationResponse structure represents the SegmentMaxDuration operation response
type SetSegmentMaxDurationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SegmentMaxDuration return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSegmentMaxDurationResponse) xxx_ToOp(ctx context.Context) *xxx_SetSegmentMaxDurationOperation {
	if o == nil {
		return &xxx_SetSegmentMaxDurationOperation{}
	}
	return &xxx_SetSegmentMaxDurationOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetSegmentMaxDurationResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSegmentMaxDurationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSegmentMaxDurationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetSegmentMaxDurationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSegmentMaxDurationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSegmentMaxSizeOperation structure represents the SegmentMaxSize operation
type xxx_GetSegmentMaxSizeOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Size   uint32         `idl:"name:size" json:"size"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSegmentMaxSizeOperation) OpNum() int { return 28 }

func (o *xxx_GetSegmentMaxSizeOperation) OpName() string {
	return "/IDataCollectorSet/v0/SegmentMaxSize"
}

func (o *xxx_GetSegmentMaxSizeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSegmentMaxSizeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSegmentMaxSizeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSegmentMaxSizeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSegmentMaxSizeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// size {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Size); err != nil {
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

func (o *xxx_GetSegmentMaxSizeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// size {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Size); err != nil {
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

// GetSegmentMaxSizeRequest structure represents the SegmentMaxSize operation request
type GetSegmentMaxSizeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSegmentMaxSizeRequest) xxx_ToOp(ctx context.Context) *xxx_GetSegmentMaxSizeOperation {
	if o == nil {
		return &xxx_GetSegmentMaxSizeOperation{}
	}
	return &xxx_GetSegmentMaxSizeOperation{
		This: o.This,
	}
}

func (o *GetSegmentMaxSizeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSegmentMaxSizeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSegmentMaxSizeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetSegmentMaxSizeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSegmentMaxSizeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSegmentMaxSizeResponse structure represents the SegmentMaxSize operation response
type GetSegmentMaxSizeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	Size uint32         `idl:"name:size" json:"size"`
	// Return: The SegmentMaxSize return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSegmentMaxSizeResponse) xxx_ToOp(ctx context.Context) *xxx_GetSegmentMaxSizeOperation {
	if o == nil {
		return &xxx_GetSegmentMaxSizeOperation{}
	}
	return &xxx_GetSegmentMaxSizeOperation{
		That:   o.That,
		Size:   o.Size,
		Return: o.Return,
	}
}

func (o *GetSegmentMaxSizeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSegmentMaxSizeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Size = op.Size
	o.Return = op.Return
}
func (o *GetSegmentMaxSizeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetSegmentMaxSizeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSegmentMaxSizeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSegmentMaxSizeOperation structure represents the SegmentMaxSize operation
type xxx_SetSegmentMaxSizeOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Size   uint32         `idl:"name:size" json:"size"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSegmentMaxSizeOperation) OpNum() int { return 29 }

func (o *xxx_SetSegmentMaxSizeOperation) OpName() string {
	return "/IDataCollectorSet/v0/SegmentMaxSize"
}

func (o *xxx_SetSegmentMaxSizeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSegmentMaxSizeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// size {in} (1:(uint32))
	{
		if err := w.WriteData(o.Size); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSegmentMaxSizeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// size {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Size); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSegmentMaxSizeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSegmentMaxSizeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetSegmentMaxSizeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetSegmentMaxSizeRequest structure represents the SegmentMaxSize operation request
type SetSegmentMaxSizeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	Size uint32         `idl:"name:size" json:"size"`
}

func (o *SetSegmentMaxSizeRequest) xxx_ToOp(ctx context.Context) *xxx_SetSegmentMaxSizeOperation {
	if o == nil {
		return &xxx_SetSegmentMaxSizeOperation{}
	}
	return &xxx_SetSegmentMaxSizeOperation{
		This: o.This,
		Size: o.Size,
	}
}

func (o *SetSegmentMaxSizeRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSegmentMaxSizeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Size = op.Size
}
func (o *SetSegmentMaxSizeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetSegmentMaxSizeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSegmentMaxSizeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSegmentMaxSizeResponse structure represents the SegmentMaxSize operation response
type SetSegmentMaxSizeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SegmentMaxSize return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSegmentMaxSizeResponse) xxx_ToOp(ctx context.Context) *xxx_SetSegmentMaxSizeOperation {
	if o == nil {
		return &xxx_SetSegmentMaxSizeOperation{}
	}
	return &xxx_SetSegmentMaxSizeOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetSegmentMaxSizeResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSegmentMaxSizeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSegmentMaxSizeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetSegmentMaxSizeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSegmentMaxSizeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSerialNumberOperation structure represents the SerialNumber operation
type xxx_GetSerialNumberOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Index  uint32         `idl:"name:index" json:"index"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSerialNumberOperation) OpNum() int { return 30 }

func (o *xxx_GetSerialNumberOperation) OpName() string { return "/IDataCollectorSet/v0/SerialNumber" }

func (o *xxx_GetSerialNumberOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSerialNumberOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSerialNumberOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSerialNumberOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSerialNumberOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// index {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Index); err != nil {
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

func (o *xxx_GetSerialNumberOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// index {out, retval} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Index); err != nil {
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

// GetSerialNumberRequest structure represents the SerialNumber operation request
type GetSerialNumberRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSerialNumberRequest) xxx_ToOp(ctx context.Context) *xxx_GetSerialNumberOperation {
	if o == nil {
		return &xxx_GetSerialNumberOperation{}
	}
	return &xxx_GetSerialNumberOperation{
		This: o.This,
	}
}

func (o *GetSerialNumberRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSerialNumberOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSerialNumberRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetSerialNumberRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSerialNumberOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSerialNumberResponse structure represents the SerialNumber operation response
type GetSerialNumberResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat `idl:"name:That" json:"that"`
	Index uint32         `idl:"name:index" json:"index"`
	// Return: The SerialNumber return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSerialNumberResponse) xxx_ToOp(ctx context.Context) *xxx_GetSerialNumberOperation {
	if o == nil {
		return &xxx_GetSerialNumberOperation{}
	}
	return &xxx_GetSerialNumberOperation{
		That:   o.That,
		Index:  o.Index,
		Return: o.Return,
	}
}

func (o *GetSerialNumberResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSerialNumberOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Index = op.Index
	o.Return = op.Return
}
func (o *GetSerialNumberResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetSerialNumberResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSerialNumberOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSerialNumberOperation structure represents the SerialNumber operation
type xxx_SetSerialNumberOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Index  uint32         `idl:"name:index" json:"index"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSerialNumberOperation) OpNum() int { return 31 }

func (o *xxx_SetSerialNumberOperation) OpName() string { return "/IDataCollectorSet/v0/SerialNumber" }

func (o *xxx_SetSerialNumberOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSerialNumberOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// index {in} (1:(uint32))
	{
		if err := w.WriteData(o.Index); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSerialNumberOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// index {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Index); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSerialNumberOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSerialNumberOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetSerialNumberOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetSerialNumberRequest structure represents the SerialNumber operation request
type SetSerialNumberRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This  *dcom.ORPCThis `idl:"name:This" json:"this"`
	Index uint32         `idl:"name:index" json:"index"`
}

func (o *SetSerialNumberRequest) xxx_ToOp(ctx context.Context) *xxx_SetSerialNumberOperation {
	if o == nil {
		return &xxx_SetSerialNumberOperation{}
	}
	return &xxx_SetSerialNumberOperation{
		This:  o.This,
		Index: o.Index,
	}
}

func (o *SetSerialNumberRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSerialNumberOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Index = op.Index
}
func (o *SetSerialNumberRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetSerialNumberRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSerialNumberOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSerialNumberResponse structure represents the SerialNumber operation response
type SetSerialNumberResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SerialNumber return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSerialNumberResponse) xxx_ToOp(ctx context.Context) *xxx_SetSerialNumberOperation {
	if o == nil {
		return &xxx_SetSerialNumberOperation{}
	}
	return &xxx_SetSerialNumberOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetSerialNumberResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSerialNumberOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSerialNumberResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetSerialNumberResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSerialNumberOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetServerOperation structure represents the Server operation
type xxx_GetServerOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Server *oaut.String   `idl:"name:server" json:"server"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetServerOperation) OpNum() int { return 32 }

func (o *xxx_GetServerOperation) OpName() string { return "/IDataCollectorSet/v0/Server" }

func (o *xxx_GetServerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetServerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetServerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// server {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Server != nil {
			_ptr_server := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Server != nil {
					if err := o.Server.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Server, _ptr_server); err != nil {
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

func (o *xxx_GetServerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// server {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_server := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Server == nil {
				o.Server = &oaut.String{}
			}
			if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_server := func(ptr interface{}) { o.Server = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Server, _s_server, _ptr_server); err != nil {
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

// GetServerRequest structure represents the Server operation request
type GetServerRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetServerRequest) xxx_ToOp(ctx context.Context) *xxx_GetServerOperation {
	if o == nil {
		return &xxx_GetServerOperation{}
	}
	return &xxx_GetServerOperation{
		This: o.This,
	}
}

func (o *GetServerRequest) xxx_FromOp(ctx context.Context, op *xxx_GetServerOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetServerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetServerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetServerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetServerResponse structure represents the Server operation response
type GetServerResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// server: Receives the name of the server where the data collector set runs.
	Server *oaut.String `idl:"name:server" json:"server"`
	// Return: The Server return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetServerResponse) xxx_ToOp(ctx context.Context) *xxx_GetServerOperation {
	if o == nil {
		return &xxx_GetServerOperation{}
	}
	return &xxx_GetServerOperation{
		That:   o.That,
		Server: o.Server,
		Return: o.Return,
	}
}

func (o *GetServerResponse) xxx_FromOp(ctx context.Context, op *xxx_GetServerOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Server = op.Server
	o.Return = op.Return
}
func (o *GetServerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetServerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetServerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetStatusOperation structure represents the Status operation
type xxx_GetStatusOperation struct {
	This   *dcom.ORPCThis             `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat             `idl:"name:That" json:"that"`
	Status pla.DataCollectorSetStatus `idl:"name:status" json:"status"`
	Return int32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_GetStatusOperation) OpNum() int { return 33 }

func (o *xxx_GetStatusOperation) OpName() string { return "/IDataCollectorSet/v0/Status" }

func (o *xxx_GetStatusOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetStatusOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetStatusOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetStatusOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetStatusOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// status {out, retval} (1:{pointer=ref}*(1))(2:{alias=DataCollectorSetStatus}(enum))
	{
		if err := w.WriteData(uint16(o.Status)); err != nil {
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

func (o *xxx_GetStatusOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// status {out, retval} (1:{pointer=ref}*(1))(2:{alias=DataCollectorSetStatus}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Status)); err != nil {
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

// GetStatusRequest structure represents the Status operation request
type GetStatusRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetStatusRequest) xxx_ToOp(ctx context.Context) *xxx_GetStatusOperation {
	if o == nil {
		return &xxx_GetStatusOperation{}
	}
	return &xxx_GetStatusOperation{
		This: o.This,
	}
}

func (o *GetStatusRequest) xxx_FromOp(ctx context.Context, op *xxx_GetStatusOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetStatusRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetStatusRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetStatusOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetStatusResponse structure represents the Status operation response
type GetStatusResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// status: Receives the status of the data collector set. The status of the data collector
	// set is defined by one of the DataCollectorSetStatus enumeration values.
	Status pla.DataCollectorSetStatus `idl:"name:status" json:"status"`
	// Return: The Status return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetStatusResponse) xxx_ToOp(ctx context.Context) *xxx_GetStatusOperation {
	if o == nil {
		return &xxx_GetStatusOperation{}
	}
	return &xxx_GetStatusOperation{
		That:   o.That,
		Status: o.Status,
		Return: o.Return,
	}
}

func (o *GetStatusResponse) xxx_FromOp(ctx context.Context, op *xxx_GetStatusOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Status = op.Status
	o.Return = op.Return
}
func (o *GetStatusResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetStatusResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetStatusOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSubdirectoryOperation structure represents the Subdirectory operation
type xxx_GetSubdirectoryOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Folder *oaut.String   `idl:"name:folder" json:"folder"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSubdirectoryOperation) OpNum() int { return 34 }

func (o *xxx_GetSubdirectoryOperation) OpName() string { return "/IDataCollectorSet/v0/Subdirectory" }

func (o *xxx_GetSubdirectoryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubdirectoryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSubdirectoryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSubdirectoryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubdirectoryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// folder {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Folder != nil {
			_ptr_folder := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Folder != nil {
					if err := o.Folder.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Folder, _ptr_folder); err != nil {
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

func (o *xxx_GetSubdirectoryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// folder {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_folder := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Folder == nil {
				o.Folder = &oaut.String{}
			}
			if err := o.Folder.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_folder := func(ptr interface{}) { o.Folder = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Folder, _s_folder, _ptr_folder); err != nil {
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

// GetSubdirectoryRequest structure represents the Subdirectory operation request
type GetSubdirectoryRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSubdirectoryRequest) xxx_ToOp(ctx context.Context) *xxx_GetSubdirectoryOperation {
	if o == nil {
		return &xxx_GetSubdirectoryOperation{}
	}
	return &xxx_GetSubdirectoryOperation{
		This: o.This,
	}
}

func (o *GetSubdirectoryRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSubdirectoryOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSubdirectoryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetSubdirectoryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSubdirectoryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSubdirectoryResponse structure represents the Subdirectory operation response
type GetSubdirectoryResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Folder *oaut.String   `idl:"name:folder" json:"folder"`
	// Return: The Subdirectory return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSubdirectoryResponse) xxx_ToOp(ctx context.Context) *xxx_GetSubdirectoryOperation {
	if o == nil {
		return &xxx_GetSubdirectoryOperation{}
	}
	return &xxx_GetSubdirectoryOperation{
		That:   o.That,
		Folder: o.Folder,
		Return: o.Return,
	}
}

func (o *GetSubdirectoryResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSubdirectoryOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Folder = op.Folder
	o.Return = op.Return
}
func (o *GetSubdirectoryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetSubdirectoryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSubdirectoryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSubdirectoryOperation structure represents the Subdirectory operation
type xxx_SetSubdirectoryOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Folder *oaut.String   `idl:"name:folder" json:"folder"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSubdirectoryOperation) OpNum() int { return 35 }

func (o *xxx_SetSubdirectoryOperation) OpName() string { return "/IDataCollectorSet/v0/Subdirectory" }

func (o *xxx_SetSubdirectoryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubdirectoryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// folder {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Folder != nil {
			_ptr_folder := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Folder != nil {
					if err := o.Folder.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Folder, _ptr_folder); err != nil {
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

func (o *xxx_SetSubdirectoryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// folder {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_folder := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Folder == nil {
				o.Folder = &oaut.String{}
			}
			if err := o.Folder.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_folder := func(ptr interface{}) { o.Folder = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Folder, _s_folder, _ptr_folder); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubdirectoryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubdirectoryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetSubdirectoryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetSubdirectoryRequest structure represents the Subdirectory operation request
type SetSubdirectoryRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	Folder *oaut.String   `idl:"name:folder" json:"folder"`
}

func (o *SetSubdirectoryRequest) xxx_ToOp(ctx context.Context) *xxx_SetSubdirectoryOperation {
	if o == nil {
		return &xxx_SetSubdirectoryOperation{}
	}
	return &xxx_SetSubdirectoryOperation{
		This:   o.This,
		Folder: o.Folder,
	}
}

func (o *SetSubdirectoryRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSubdirectoryOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Folder = op.Folder
}
func (o *SetSubdirectoryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetSubdirectoryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSubdirectoryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSubdirectoryResponse structure represents the Subdirectory operation response
type SetSubdirectoryResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Subdirectory return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSubdirectoryResponse) xxx_ToOp(ctx context.Context) *xxx_SetSubdirectoryOperation {
	if o == nil {
		return &xxx_SetSubdirectoryOperation{}
	}
	return &xxx_SetSubdirectoryOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetSubdirectoryResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSubdirectoryOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSubdirectoryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetSubdirectoryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSubdirectoryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSubdirectoryFormatOperation structure represents the SubdirectoryFormat operation
type xxx_GetSubdirectoryFormatOperation struct {
	This   *dcom.ORPCThis     `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat     `idl:"name:That" json:"that"`
	Format pla.AutoPathFormat `idl:"name:format" json:"format"`
	Return int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSubdirectoryFormatOperation) OpNum() int { return 36 }

func (o *xxx_GetSubdirectoryFormatOperation) OpName() string {
	return "/IDataCollectorSet/v0/SubdirectoryFormat"
}

func (o *xxx_GetSubdirectoryFormatOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubdirectoryFormatOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSubdirectoryFormatOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSubdirectoryFormatOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubdirectoryFormatOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// format {out, retval} (1:{pointer=ref}*(1))(2:{alias=AutoPathFormat}(enum))
	{
		if err := w.WriteData(uint16(o.Format)); err != nil {
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

func (o *xxx_GetSubdirectoryFormatOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// format {out, retval} (1:{pointer=ref}*(1))(2:{alias=AutoPathFormat}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Format)); err != nil {
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

// GetSubdirectoryFormatRequest structure represents the SubdirectoryFormat operation request
type GetSubdirectoryFormatRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSubdirectoryFormatRequest) xxx_ToOp(ctx context.Context) *xxx_GetSubdirectoryFormatOperation {
	if o == nil {
		return &xxx_GetSubdirectoryFormatOperation{}
	}
	return &xxx_GetSubdirectoryFormatOperation{
		This: o.This,
	}
}

func (o *GetSubdirectoryFormatRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSubdirectoryFormatOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSubdirectoryFormatRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetSubdirectoryFormatRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSubdirectoryFormatOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSubdirectoryFormatResponse structure represents the SubdirectoryFormat operation response
type GetSubdirectoryFormatResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat     `idl:"name:That" json:"that"`
	Format pla.AutoPathFormat `idl:"name:format" json:"format"`
	// Return: The SubdirectoryFormat return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSubdirectoryFormatResponse) xxx_ToOp(ctx context.Context) *xxx_GetSubdirectoryFormatOperation {
	if o == nil {
		return &xxx_GetSubdirectoryFormatOperation{}
	}
	return &xxx_GetSubdirectoryFormatOperation{
		That:   o.That,
		Format: o.Format,
		Return: o.Return,
	}
}

func (o *GetSubdirectoryFormatResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSubdirectoryFormatOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Format = op.Format
	o.Return = op.Return
}
func (o *GetSubdirectoryFormatResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetSubdirectoryFormatResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSubdirectoryFormatOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSubdirectoryFormatOperation structure represents the SubdirectoryFormat operation
type xxx_SetSubdirectoryFormatOperation struct {
	This   *dcom.ORPCThis     `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat     `idl:"name:That" json:"that"`
	Format pla.AutoPathFormat `idl:"name:format" json:"format"`
	Return int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSubdirectoryFormatOperation) OpNum() int { return 37 }

func (o *xxx_SetSubdirectoryFormatOperation) OpName() string {
	return "/IDataCollectorSet/v0/SubdirectoryFormat"
}

func (o *xxx_SetSubdirectoryFormatOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubdirectoryFormatOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// format {in} (1:{alias=AutoPathFormat}(enum))
	{
		if err := w.WriteData(uint16(o.Format)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubdirectoryFormatOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// format {in} (1:{alias=AutoPathFormat}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Format)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubdirectoryFormatOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubdirectoryFormatOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetSubdirectoryFormatOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetSubdirectoryFormatRequest structure represents the SubdirectoryFormat operation request
type SetSubdirectoryFormatRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This   *dcom.ORPCThis     `idl:"name:This" json:"this"`
	Format pla.AutoPathFormat `idl:"name:format" json:"format"`
}

func (o *SetSubdirectoryFormatRequest) xxx_ToOp(ctx context.Context) *xxx_SetSubdirectoryFormatOperation {
	if o == nil {
		return &xxx_SetSubdirectoryFormatOperation{}
	}
	return &xxx_SetSubdirectoryFormatOperation{
		This:   o.This,
		Format: o.Format,
	}
}

func (o *SetSubdirectoryFormatRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSubdirectoryFormatOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Format = op.Format
}
func (o *SetSubdirectoryFormatRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetSubdirectoryFormatRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSubdirectoryFormatOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSubdirectoryFormatResponse structure represents the SubdirectoryFormat operation response
type SetSubdirectoryFormatResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SubdirectoryFormat return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSubdirectoryFormatResponse) xxx_ToOp(ctx context.Context) *xxx_SetSubdirectoryFormatOperation {
	if o == nil {
		return &xxx_SetSubdirectoryFormatOperation{}
	}
	return &xxx_SetSubdirectoryFormatOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetSubdirectoryFormatResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSubdirectoryFormatOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSubdirectoryFormatResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetSubdirectoryFormatResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSubdirectoryFormatOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSubdirectoryFormatPatternOperation structure represents the SubdirectoryFormatPattern operation
type xxx_GetSubdirectoryFormatPatternOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Pattern *oaut.String   `idl:"name:pattern" json:"pattern"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSubdirectoryFormatPatternOperation) OpNum() int { return 38 }

func (o *xxx_GetSubdirectoryFormatPatternOperation) OpName() string {
	return "/IDataCollectorSet/v0/SubdirectoryFormatPattern"
}

func (o *xxx_GetSubdirectoryFormatPatternOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubdirectoryFormatPatternOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSubdirectoryFormatPatternOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSubdirectoryFormatPatternOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubdirectoryFormatPatternOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pattern {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Pattern != nil {
			_ptr_pattern := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Pattern != nil {
					if err := o.Pattern.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Pattern, _ptr_pattern); err != nil {
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

func (o *xxx_GetSubdirectoryFormatPatternOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pattern {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pattern := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Pattern == nil {
				o.Pattern = &oaut.String{}
			}
			if err := o.Pattern.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pattern := func(ptr interface{}) { o.Pattern = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Pattern, _s_pattern, _ptr_pattern); err != nil {
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

// GetSubdirectoryFormatPatternRequest structure represents the SubdirectoryFormatPattern operation request
type GetSubdirectoryFormatPatternRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSubdirectoryFormatPatternRequest) xxx_ToOp(ctx context.Context) *xxx_GetSubdirectoryFormatPatternOperation {
	if o == nil {
		return &xxx_GetSubdirectoryFormatPatternOperation{}
	}
	return &xxx_GetSubdirectoryFormatPatternOperation{
		This: o.This,
	}
}

func (o *GetSubdirectoryFormatPatternRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSubdirectoryFormatPatternOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSubdirectoryFormatPatternRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetSubdirectoryFormatPatternRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSubdirectoryFormatPatternOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSubdirectoryFormatPatternResponse structure represents the SubdirectoryFormatPattern operation response
type GetSubdirectoryFormatPatternResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Pattern *oaut.String   `idl:"name:pattern" json:"pattern"`
	// Return: The SubdirectoryFormatPattern return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSubdirectoryFormatPatternResponse) xxx_ToOp(ctx context.Context) *xxx_GetSubdirectoryFormatPatternOperation {
	if o == nil {
		return &xxx_GetSubdirectoryFormatPatternOperation{}
	}
	return &xxx_GetSubdirectoryFormatPatternOperation{
		That:    o.That,
		Pattern: o.Pattern,
		Return:  o.Return,
	}
}

func (o *GetSubdirectoryFormatPatternResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSubdirectoryFormatPatternOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Pattern = op.Pattern
	o.Return = op.Return
}
func (o *GetSubdirectoryFormatPatternResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetSubdirectoryFormatPatternResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSubdirectoryFormatPatternOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSubdirectoryFormatPatternOperation structure represents the SubdirectoryFormatPattern operation
type xxx_SetSubdirectoryFormatPatternOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Pattern *oaut.String   `idl:"name:pattern" json:"pattern"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSubdirectoryFormatPatternOperation) OpNum() int { return 39 }

func (o *xxx_SetSubdirectoryFormatPatternOperation) OpName() string {
	return "/IDataCollectorSet/v0/SubdirectoryFormatPattern"
}

func (o *xxx_SetSubdirectoryFormatPatternOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubdirectoryFormatPatternOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pattern {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Pattern != nil {
			_ptr_pattern := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Pattern != nil {
					if err := o.Pattern.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Pattern, _ptr_pattern); err != nil {
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

func (o *xxx_SetSubdirectoryFormatPatternOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pattern {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pattern := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Pattern == nil {
				o.Pattern = &oaut.String{}
			}
			if err := o.Pattern.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pattern := func(ptr interface{}) { o.Pattern = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Pattern, _s_pattern, _ptr_pattern); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubdirectoryFormatPatternOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubdirectoryFormatPatternOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetSubdirectoryFormatPatternOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetSubdirectoryFormatPatternRequest structure represents the SubdirectoryFormatPattern operation request
type SetSubdirectoryFormatPatternRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	Pattern *oaut.String   `idl:"name:pattern" json:"pattern"`
}

func (o *SetSubdirectoryFormatPatternRequest) xxx_ToOp(ctx context.Context) *xxx_SetSubdirectoryFormatPatternOperation {
	if o == nil {
		return &xxx_SetSubdirectoryFormatPatternOperation{}
	}
	return &xxx_SetSubdirectoryFormatPatternOperation{
		This:    o.This,
		Pattern: o.Pattern,
	}
}

func (o *SetSubdirectoryFormatPatternRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSubdirectoryFormatPatternOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Pattern = op.Pattern
}
func (o *SetSubdirectoryFormatPatternRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetSubdirectoryFormatPatternRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSubdirectoryFormatPatternOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSubdirectoryFormatPatternResponse structure represents the SubdirectoryFormatPattern operation response
type SetSubdirectoryFormatPatternResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SubdirectoryFormatPattern return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSubdirectoryFormatPatternResponse) xxx_ToOp(ctx context.Context) *xxx_SetSubdirectoryFormatPatternOperation {
	if o == nil {
		return &xxx_SetSubdirectoryFormatPatternOperation{}
	}
	return &xxx_SetSubdirectoryFormatPatternOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetSubdirectoryFormatPatternResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSubdirectoryFormatPatternOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSubdirectoryFormatPatternResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetSubdirectoryFormatPatternResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSubdirectoryFormatPatternOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetTaskOperation structure represents the Task operation
type xxx_GetTaskOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Task   *oaut.String   `idl:"name:task" json:"task"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTaskOperation) OpNum() int { return 40 }

func (o *xxx_GetTaskOperation) OpName() string { return "/IDataCollectorSet/v0/Task" }

func (o *xxx_GetTaskOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTaskOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetTaskOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetTaskOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTaskOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// task {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Task != nil {
			_ptr_task := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Task != nil {
					if err := o.Task.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Task, _ptr_task); err != nil {
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

func (o *xxx_GetTaskOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// task {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_task := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Task == nil {
				o.Task = &oaut.String{}
			}
			if err := o.Task.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_task := func(ptr interface{}) { o.Task = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Task, _s_task, _ptr_task); err != nil {
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

// GetTaskRequest structure represents the Task operation request
type GetTaskRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetTaskRequest) xxx_ToOp(ctx context.Context) *xxx_GetTaskOperation {
	if o == nil {
		return &xxx_GetTaskOperation{}
	}
	return &xxx_GetTaskOperation{
		This: o.This,
	}
}

func (o *GetTaskRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTaskOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetTaskRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetTaskRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTaskOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetTaskResponse structure represents the Task operation response
type GetTaskResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	Task *oaut.String   `idl:"name:task" json:"task"`
	// Return: The Task return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTaskResponse) xxx_ToOp(ctx context.Context) *xxx_GetTaskOperation {
	if o == nil {
		return &xxx_GetTaskOperation{}
	}
	return &xxx_GetTaskOperation{
		That:   o.That,
		Task:   o.Task,
		Return: o.Return,
	}
}

func (o *GetTaskResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTaskOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Task = op.Task
	o.Return = op.Return
}
func (o *GetTaskResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetTaskResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTaskOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetTaskOperation structure represents the Task operation
type xxx_SetTaskOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Task   *oaut.String   `idl:"name:task" json:"task"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetTaskOperation) OpNum() int { return 41 }

func (o *xxx_SetTaskOperation) OpName() string { return "/IDataCollectorSet/v0/Task" }

func (o *xxx_SetTaskOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTaskOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// task {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Task != nil {
			_ptr_task := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Task != nil {
					if err := o.Task.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Task, _ptr_task); err != nil {
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

func (o *xxx_SetTaskOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// task {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_task := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Task == nil {
				o.Task = &oaut.String{}
			}
			if err := o.Task.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_task := func(ptr interface{}) { o.Task = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Task, _s_task, _ptr_task); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTaskOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTaskOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetTaskOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetTaskRequest structure represents the Task operation request
type SetTaskRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	Task *oaut.String   `idl:"name:task" json:"task"`
}

func (o *SetTaskRequest) xxx_ToOp(ctx context.Context) *xxx_SetTaskOperation {
	if o == nil {
		return &xxx_SetTaskOperation{}
	}
	return &xxx_SetTaskOperation{
		This: o.This,
		Task: o.Task,
	}
}

func (o *SetTaskRequest) xxx_FromOp(ctx context.Context, op *xxx_SetTaskOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Task = op.Task
}
func (o *SetTaskRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetTaskRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetTaskOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetTaskResponse structure represents the Task operation response
type SetTaskResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Task return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetTaskResponse) xxx_ToOp(ctx context.Context) *xxx_SetTaskOperation {
	if o == nil {
		return &xxx_SetTaskOperation{}
	}
	return &xxx_SetTaskOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetTaskResponse) xxx_FromOp(ctx context.Context, op *xxx_SetTaskOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetTaskResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetTaskResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetTaskOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetTaskRunAsSelfOperation structure represents the TaskRunAsSelf operation
type xxx_GetTaskRunAsSelfOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	RunAsSelf int16          `idl:"name:RunAsSelf" json:"run_as_self"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTaskRunAsSelfOperation) OpNum() int { return 42 }

func (o *xxx_GetTaskRunAsSelfOperation) OpName() string { return "/IDataCollectorSet/v0/TaskRunAsSelf" }

func (o *xxx_GetTaskRunAsSelfOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTaskRunAsSelfOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetTaskRunAsSelfOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetTaskRunAsSelfOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTaskRunAsSelfOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// RunAsSelf {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.RunAsSelf); err != nil {
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

func (o *xxx_GetTaskRunAsSelfOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// RunAsSelf {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.RunAsSelf); err != nil {
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

// GetTaskRunAsSelfRequest structure represents the TaskRunAsSelf operation request
type GetTaskRunAsSelfRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetTaskRunAsSelfRequest) xxx_ToOp(ctx context.Context) *xxx_GetTaskRunAsSelfOperation {
	if o == nil {
		return &xxx_GetTaskRunAsSelfOperation{}
	}
	return &xxx_GetTaskRunAsSelfOperation{
		This: o.This,
	}
}

func (o *GetTaskRunAsSelfRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTaskRunAsSelfOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetTaskRunAsSelfRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetTaskRunAsSelfRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTaskRunAsSelfOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetTaskRunAsSelfResponse structure represents the TaskRunAsSelf operation response
type GetTaskRunAsSelfResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	RunAsSelf int16          `idl:"name:RunAsSelf" json:"run_as_self"`
	// Return: The TaskRunAsSelf return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTaskRunAsSelfResponse) xxx_ToOp(ctx context.Context) *xxx_GetTaskRunAsSelfOperation {
	if o == nil {
		return &xxx_GetTaskRunAsSelfOperation{}
	}
	return &xxx_GetTaskRunAsSelfOperation{
		That:      o.That,
		RunAsSelf: o.RunAsSelf,
		Return:    o.Return,
	}
}

func (o *GetTaskRunAsSelfResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTaskRunAsSelfOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.RunAsSelf = op.RunAsSelf
	o.Return = op.Return
}
func (o *GetTaskRunAsSelfResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetTaskRunAsSelfResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTaskRunAsSelfOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetTaskRunAsSelfOperation structure represents the TaskRunAsSelf operation
type xxx_SetTaskRunAsSelfOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	RunAsSelf int16          `idl:"name:RunAsSelf" json:"run_as_self"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetTaskRunAsSelfOperation) OpNum() int { return 43 }

func (o *xxx_SetTaskRunAsSelfOperation) OpName() string { return "/IDataCollectorSet/v0/TaskRunAsSelf" }

func (o *xxx_SetTaskRunAsSelfOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTaskRunAsSelfOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// RunAsSelf {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.RunAsSelf); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTaskRunAsSelfOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// RunAsSelf {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.RunAsSelf); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTaskRunAsSelfOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTaskRunAsSelfOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetTaskRunAsSelfOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetTaskRunAsSelfRequest structure represents the TaskRunAsSelf operation request
type SetTaskRunAsSelfRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	RunAsSelf int16          `idl:"name:RunAsSelf" json:"run_as_self"`
}

func (o *SetTaskRunAsSelfRequest) xxx_ToOp(ctx context.Context) *xxx_SetTaskRunAsSelfOperation {
	if o == nil {
		return &xxx_SetTaskRunAsSelfOperation{}
	}
	return &xxx_SetTaskRunAsSelfOperation{
		This:      o.This,
		RunAsSelf: o.RunAsSelf,
	}
}

func (o *SetTaskRunAsSelfRequest) xxx_FromOp(ctx context.Context, op *xxx_SetTaskRunAsSelfOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.RunAsSelf = op.RunAsSelf
}
func (o *SetTaskRunAsSelfRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetTaskRunAsSelfRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetTaskRunAsSelfOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetTaskRunAsSelfResponse structure represents the TaskRunAsSelf operation response
type SetTaskRunAsSelfResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The TaskRunAsSelf return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetTaskRunAsSelfResponse) xxx_ToOp(ctx context.Context) *xxx_SetTaskRunAsSelfOperation {
	if o == nil {
		return &xxx_SetTaskRunAsSelfOperation{}
	}
	return &xxx_SetTaskRunAsSelfOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetTaskRunAsSelfResponse) xxx_FromOp(ctx context.Context, op *xxx_SetTaskRunAsSelfOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetTaskRunAsSelfResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetTaskRunAsSelfResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetTaskRunAsSelfOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetTaskArgumentsOperation structure represents the TaskArguments operation
type xxx_GetTaskArgumentsOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Task   *oaut.String   `idl:"name:task" json:"task"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTaskArgumentsOperation) OpNum() int { return 44 }

func (o *xxx_GetTaskArgumentsOperation) OpName() string { return "/IDataCollectorSet/v0/TaskArguments" }

func (o *xxx_GetTaskArgumentsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTaskArgumentsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetTaskArgumentsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetTaskArgumentsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTaskArgumentsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// task {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Task != nil {
			_ptr_task := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Task != nil {
					if err := o.Task.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Task, _ptr_task); err != nil {
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

func (o *xxx_GetTaskArgumentsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// task {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_task := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Task == nil {
				o.Task = &oaut.String{}
			}
			if err := o.Task.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_task := func(ptr interface{}) { o.Task = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Task, _s_task, _ptr_task); err != nil {
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

// GetTaskArgumentsRequest structure represents the TaskArguments operation request
type GetTaskArgumentsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetTaskArgumentsRequest) xxx_ToOp(ctx context.Context) *xxx_GetTaskArgumentsOperation {
	if o == nil {
		return &xxx_GetTaskArgumentsOperation{}
	}
	return &xxx_GetTaskArgumentsOperation{
		This: o.This,
	}
}

func (o *GetTaskArgumentsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTaskArgumentsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetTaskArgumentsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetTaskArgumentsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTaskArgumentsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetTaskArgumentsResponse structure represents the TaskArguments operation response
type GetTaskArgumentsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	Task *oaut.String   `idl:"name:task" json:"task"`
	// Return: The TaskArguments return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTaskArgumentsResponse) xxx_ToOp(ctx context.Context) *xxx_GetTaskArgumentsOperation {
	if o == nil {
		return &xxx_GetTaskArgumentsOperation{}
	}
	return &xxx_GetTaskArgumentsOperation{
		That:   o.That,
		Task:   o.Task,
		Return: o.Return,
	}
}

func (o *GetTaskArgumentsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTaskArgumentsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Task = op.Task
	o.Return = op.Return
}
func (o *GetTaskArgumentsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetTaskArgumentsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTaskArgumentsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetTaskArgumentsOperation structure represents the TaskArguments operation
type xxx_SetTaskArgumentsOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Task   *oaut.String   `idl:"name:task" json:"task"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetTaskArgumentsOperation) OpNum() int { return 45 }

func (o *xxx_SetTaskArgumentsOperation) OpName() string { return "/IDataCollectorSet/v0/TaskArguments" }

func (o *xxx_SetTaskArgumentsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTaskArgumentsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// task {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Task != nil {
			_ptr_task := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Task != nil {
					if err := o.Task.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Task, _ptr_task); err != nil {
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

func (o *xxx_SetTaskArgumentsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// task {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_task := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Task == nil {
				o.Task = &oaut.String{}
			}
			if err := o.Task.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_task := func(ptr interface{}) { o.Task = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Task, _s_task, _ptr_task); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTaskArgumentsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTaskArgumentsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetTaskArgumentsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetTaskArgumentsRequest structure represents the TaskArguments operation request
type SetTaskArgumentsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	Task *oaut.String   `idl:"name:task" json:"task"`
}

func (o *SetTaskArgumentsRequest) xxx_ToOp(ctx context.Context) *xxx_SetTaskArgumentsOperation {
	if o == nil {
		return &xxx_SetTaskArgumentsOperation{}
	}
	return &xxx_SetTaskArgumentsOperation{
		This: o.This,
		Task: o.Task,
	}
}

func (o *SetTaskArgumentsRequest) xxx_FromOp(ctx context.Context, op *xxx_SetTaskArgumentsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Task = op.Task
}
func (o *SetTaskArgumentsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetTaskArgumentsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetTaskArgumentsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetTaskArgumentsResponse structure represents the TaskArguments operation response
type SetTaskArgumentsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The TaskArguments return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetTaskArgumentsResponse) xxx_ToOp(ctx context.Context) *xxx_SetTaskArgumentsOperation {
	if o == nil {
		return &xxx_SetTaskArgumentsOperation{}
	}
	return &xxx_SetTaskArgumentsOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetTaskArgumentsResponse) xxx_FromOp(ctx context.Context, op *xxx_SetTaskArgumentsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetTaskArgumentsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetTaskArgumentsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetTaskArgumentsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetTaskUserTextArgumentsOperation structure represents the TaskUserTextArguments operation
type xxx_GetTaskUserTextArgumentsOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	UserText *oaut.String   `idl:"name:UserText" json:"user_text"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTaskUserTextArgumentsOperation) OpNum() int { return 46 }

func (o *xxx_GetTaskUserTextArgumentsOperation) OpName() string {
	return "/IDataCollectorSet/v0/TaskUserTextArguments"
}

func (o *xxx_GetTaskUserTextArgumentsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTaskUserTextArgumentsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetTaskUserTextArgumentsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetTaskUserTextArgumentsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTaskUserTextArgumentsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// UserText {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.UserText != nil {
			_ptr_UserText := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.UserText != nil {
					if err := o.UserText.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.UserText, _ptr_UserText); err != nil {
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

func (o *xxx_GetTaskUserTextArgumentsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// UserText {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_UserText := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.UserText == nil {
				o.UserText = &oaut.String{}
			}
			if err := o.UserText.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_UserText := func(ptr interface{}) { o.UserText = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.UserText, _s_UserText, _ptr_UserText); err != nil {
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

// GetTaskUserTextArgumentsRequest structure represents the TaskUserTextArguments operation request
type GetTaskUserTextArgumentsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetTaskUserTextArgumentsRequest) xxx_ToOp(ctx context.Context) *xxx_GetTaskUserTextArgumentsOperation {
	if o == nil {
		return &xxx_GetTaskUserTextArgumentsOperation{}
	}
	return &xxx_GetTaskUserTextArgumentsOperation{
		This: o.This,
	}
}

func (o *GetTaskUserTextArgumentsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTaskUserTextArgumentsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetTaskUserTextArgumentsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetTaskUserTextArgumentsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTaskUserTextArgumentsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetTaskUserTextArgumentsResponse structure represents the TaskUserTextArguments operation response
type GetTaskUserTextArgumentsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	UserText *oaut.String   `idl:"name:UserText" json:"user_text"`
	// Return: The TaskUserTextArguments return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTaskUserTextArgumentsResponse) xxx_ToOp(ctx context.Context) *xxx_GetTaskUserTextArgumentsOperation {
	if o == nil {
		return &xxx_GetTaskUserTextArgumentsOperation{}
	}
	return &xxx_GetTaskUserTextArgumentsOperation{
		That:     o.That,
		UserText: o.UserText,
		Return:   o.Return,
	}
}

func (o *GetTaskUserTextArgumentsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTaskUserTextArgumentsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.UserText = op.UserText
	o.Return = op.Return
}
func (o *GetTaskUserTextArgumentsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetTaskUserTextArgumentsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTaskUserTextArgumentsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetTaskUserTextArgumentsOperation structure represents the TaskUserTextArguments operation
type xxx_SetTaskUserTextArgumentsOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	UserText *oaut.String   `idl:"name:UserText" json:"user_text"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetTaskUserTextArgumentsOperation) OpNum() int { return 47 }

func (o *xxx_SetTaskUserTextArgumentsOperation) OpName() string {
	return "/IDataCollectorSet/v0/TaskUserTextArguments"
}

func (o *xxx_SetTaskUserTextArgumentsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTaskUserTextArgumentsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// UserText {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.UserText != nil {
			_ptr_UserText := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.UserText != nil {
					if err := o.UserText.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.UserText, _ptr_UserText); err != nil {
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

func (o *xxx_SetTaskUserTextArgumentsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// UserText {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_UserText := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.UserText == nil {
				o.UserText = &oaut.String{}
			}
			if err := o.UserText.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_UserText := func(ptr interface{}) { o.UserText = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.UserText, _s_UserText, _ptr_UserText); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTaskUserTextArgumentsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTaskUserTextArgumentsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetTaskUserTextArgumentsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetTaskUserTextArgumentsRequest structure represents the TaskUserTextArguments operation request
type SetTaskUserTextArgumentsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	UserText *oaut.String   `idl:"name:UserText" json:"user_text"`
}

func (o *SetTaskUserTextArgumentsRequest) xxx_ToOp(ctx context.Context) *xxx_SetTaskUserTextArgumentsOperation {
	if o == nil {
		return &xxx_SetTaskUserTextArgumentsOperation{}
	}
	return &xxx_SetTaskUserTextArgumentsOperation{
		This:     o.This,
		UserText: o.UserText,
	}
}

func (o *SetTaskUserTextArgumentsRequest) xxx_FromOp(ctx context.Context, op *xxx_SetTaskUserTextArgumentsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.UserText = op.UserText
}
func (o *SetTaskUserTextArgumentsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetTaskUserTextArgumentsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetTaskUserTextArgumentsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetTaskUserTextArgumentsResponse structure represents the TaskUserTextArguments operation response
type SetTaskUserTextArgumentsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The TaskUserTextArguments return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetTaskUserTextArgumentsResponse) xxx_ToOp(ctx context.Context) *xxx_SetTaskUserTextArgumentsOperation {
	if o == nil {
		return &xxx_SetTaskUserTextArgumentsOperation{}
	}
	return &xxx_SetTaskUserTextArgumentsOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetTaskUserTextArgumentsResponse) xxx_FromOp(ctx context.Context, op *xxx_SetTaskUserTextArgumentsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetTaskUserTextArgumentsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetTaskUserTextArgumentsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetTaskUserTextArgumentsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSchedulesOperation structure represents the Schedules operation
type xxx_GetSchedulesOperation struct {
	This      *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat          `idl:"name:That" json:"that"`
	Schedules *pla.ScheduleCollection `idl:"name:ppSchedules" json:"schedules"`
	Return    int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSchedulesOperation) OpNum() int { return 48 }

func (o *xxx_GetSchedulesOperation) OpName() string { return "/IDataCollectorSet/v0/Schedules" }

func (o *xxx_GetSchedulesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSchedulesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSchedulesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSchedulesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSchedulesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppSchedules {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IScheduleCollection}(interface))
	{
		if o.Schedules != nil {
			_ptr_ppSchedules := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Schedules != nil {
					if err := o.Schedules.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&pla.ScheduleCollection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Schedules, _ptr_ppSchedules); err != nil {
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

func (o *xxx_GetSchedulesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppSchedules {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IScheduleCollection}(interface))
	{
		_ptr_ppSchedules := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Schedules == nil {
				o.Schedules = &pla.ScheduleCollection{}
			}
			if err := o.Schedules.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppSchedules := func(ptr interface{}) { o.Schedules = *ptr.(**pla.ScheduleCollection) }
		if err := w.ReadPointer(&o.Schedules, _s_ppSchedules, _ptr_ppSchedules); err != nil {
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

// GetSchedulesRequest structure represents the Schedules operation request
type GetSchedulesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSchedulesRequest) xxx_ToOp(ctx context.Context) *xxx_GetSchedulesOperation {
	if o == nil {
		return &xxx_GetSchedulesOperation{}
	}
	return &xxx_GetSchedulesOperation{
		This: o.This,
	}
}

func (o *GetSchedulesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSchedulesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSchedulesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetSchedulesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSchedulesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSchedulesResponse structure represents the Schedules operation response
type GetSchedulesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppSchedules: Receives a pointer to the schedule collection object. Schedules are
	// created with the CreateSchedule method.
	Schedules *pla.ScheduleCollection `idl:"name:ppSchedules" json:"schedules"`
	// Return: The Schedules return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSchedulesResponse) xxx_ToOp(ctx context.Context) *xxx_GetSchedulesOperation {
	if o == nil {
		return &xxx_GetSchedulesOperation{}
	}
	return &xxx_GetSchedulesOperation{
		That:      o.That,
		Schedules: o.Schedules,
		Return:    o.Return,
	}
}

func (o *GetSchedulesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSchedulesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Schedules = op.Schedules
	o.Return = op.Return
}
func (o *GetSchedulesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetSchedulesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSchedulesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSchedulesEnabledOperation structure represents the SchedulesEnabled operation
type xxx_GetSchedulesEnabledOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Enabled int16          `idl:"name:enabled" json:"enabled"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSchedulesEnabledOperation) OpNum() int { return 49 }

func (o *xxx_GetSchedulesEnabledOperation) OpName() string {
	return "/IDataCollectorSet/v0/SchedulesEnabled"
}

func (o *xxx_GetSchedulesEnabledOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSchedulesEnabledOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSchedulesEnabledOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSchedulesEnabledOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSchedulesEnabledOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// enabled {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Enabled); err != nil {
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

func (o *xxx_GetSchedulesEnabledOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// enabled {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Enabled); err != nil {
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

// GetSchedulesEnabledRequest structure represents the SchedulesEnabled operation request
type GetSchedulesEnabledRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSchedulesEnabledRequest) xxx_ToOp(ctx context.Context) *xxx_GetSchedulesEnabledOperation {
	if o == nil {
		return &xxx_GetSchedulesEnabledOperation{}
	}
	return &xxx_GetSchedulesEnabledOperation{
		This: o.This,
	}
}

func (o *GetSchedulesEnabledRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSchedulesEnabledOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSchedulesEnabledRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetSchedulesEnabledRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSchedulesEnabledOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSchedulesEnabledResponse structure represents the SchedulesEnabled operation response
type GetSchedulesEnabledResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Enabled int16          `idl:"name:enabled" json:"enabled"`
	// Return: The SchedulesEnabled return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSchedulesEnabledResponse) xxx_ToOp(ctx context.Context) *xxx_GetSchedulesEnabledOperation {
	if o == nil {
		return &xxx_GetSchedulesEnabledOperation{}
	}
	return &xxx_GetSchedulesEnabledOperation{
		That:    o.That,
		Enabled: o.Enabled,
		Return:  o.Return,
	}
}

func (o *GetSchedulesEnabledResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSchedulesEnabledOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Enabled = op.Enabled
	o.Return = op.Return
}
func (o *GetSchedulesEnabledResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetSchedulesEnabledResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSchedulesEnabledOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSchedulesEnabledOperation structure represents the SchedulesEnabled operation
type xxx_SetSchedulesEnabledOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Enabled int16          `idl:"name:enabled" json:"enabled"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSchedulesEnabledOperation) OpNum() int { return 50 }

func (o *xxx_SetSchedulesEnabledOperation) OpName() string {
	return "/IDataCollectorSet/v0/SchedulesEnabled"
}

func (o *xxx_SetSchedulesEnabledOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSchedulesEnabledOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// enabled {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Enabled); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSchedulesEnabledOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// enabled {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Enabled); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSchedulesEnabledOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSchedulesEnabledOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetSchedulesEnabledOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetSchedulesEnabledRequest structure represents the SchedulesEnabled operation request
type SetSchedulesEnabledRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	Enabled int16          `idl:"name:enabled" json:"enabled"`
}

func (o *SetSchedulesEnabledRequest) xxx_ToOp(ctx context.Context) *xxx_SetSchedulesEnabledOperation {
	if o == nil {
		return &xxx_SetSchedulesEnabledOperation{}
	}
	return &xxx_SetSchedulesEnabledOperation{
		This:    o.This,
		Enabled: o.Enabled,
	}
}

func (o *SetSchedulesEnabledRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSchedulesEnabledOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Enabled = op.Enabled
}
func (o *SetSchedulesEnabledRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetSchedulesEnabledRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSchedulesEnabledOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSchedulesEnabledResponse structure represents the SchedulesEnabled operation response
type SetSchedulesEnabledResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SchedulesEnabled return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSchedulesEnabledResponse) xxx_ToOp(ctx context.Context) *xxx_SetSchedulesEnabledOperation {
	if o == nil {
		return &xxx_SetSchedulesEnabledOperation{}
	}
	return &xxx_SetSchedulesEnabledOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetSchedulesEnabledResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSchedulesEnabledOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSchedulesEnabledResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetSchedulesEnabledResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSchedulesEnabledOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetUserAccountOperation structure represents the UserAccount operation
type xxx_GetUserAccountOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	User   *oaut.String   `idl:"name:user" json:"user"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetUserAccountOperation) OpNum() int { return 51 }

func (o *xxx_GetUserAccountOperation) OpName() string { return "/IDataCollectorSet/v0/UserAccount" }

func (o *xxx_GetUserAccountOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUserAccountOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetUserAccountOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetUserAccountOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUserAccountOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// user {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.User != nil {
			_ptr_user := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.User != nil {
					if err := o.User.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUserAccountOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// user {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_user := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.User == nil {
				o.User = &oaut.String{}
			}
			if err := o.User.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_user := func(ptr interface{}) { o.User = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.User, _s_user, _ptr_user); err != nil {
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

// GetUserAccountRequest structure represents the UserAccount operation request
type GetUserAccountRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetUserAccountRequest) xxx_ToOp(ctx context.Context) *xxx_GetUserAccountOperation {
	if o == nil {
		return &xxx_GetUserAccountOperation{}
	}
	return &xxx_GetUserAccountOperation{
		This: o.This,
	}
}

func (o *GetUserAccountRequest) xxx_FromOp(ctx context.Context, op *xxx_GetUserAccountOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetUserAccountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetUserAccountRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetUserAccountOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetUserAccountResponse structure represents the UserAccount operation response
type GetUserAccountResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// user: Receives the user account under which the data collector set will run. The
	// value of the string that is returned will be that which was passed as a parameter
	// to the IDataCollectorSet::SetCredentials method if that method was called after the
	// last IDataCollectorSet::Query method was called. If the SetCredentials method was
	// not called, then the string that is returned will be that of the user which the data
	// collector set is registered to run as. In such a case, if the machine is joined to
	// a domain and the user is a member of that domain, then the form of the returned string
	// will be domain\user or user@domain.
	User *oaut.String `idl:"name:user" json:"user"`
	// Return: The UserAccount return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetUserAccountResponse) xxx_ToOp(ctx context.Context) *xxx_GetUserAccountOperation {
	if o == nil {
		return &xxx_GetUserAccountOperation{}
	}
	return &xxx_GetUserAccountOperation{
		That:   o.That,
		User:   o.User,
		Return: o.Return,
	}
}

func (o *GetUserAccountResponse) xxx_FromOp(ctx context.Context, op *xxx_GetUserAccountOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.User = op.User
	o.Return = op.Return
}
func (o *GetUserAccountResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetUserAccountResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetUserAccountOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetXMLOperation structure represents the Xml operation
type xxx_GetXMLOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	XML    *oaut.String   `idl:"name:xml" json:"xml"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetXMLOperation) OpNum() int { return 52 }

func (o *xxx_GetXMLOperation) OpName() string { return "/IDataCollectorSet/v0/Xml" }

func (o *xxx_GetXMLOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetXMLOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetXMLOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetXMLOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetXMLOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// xml {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.XML != nil {
			_ptr_xml := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.XML != nil {
					if err := o.XML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.XML, _ptr_xml); err != nil {
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

func (o *xxx_GetXMLOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// xml {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_xml := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.XML == nil {
				o.XML = &oaut.String{}
			}
			if err := o.XML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_xml := func(ptr interface{}) { o.XML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.XML, _s_xml, _ptr_xml); err != nil {
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

// GetXMLRequest structure represents the Xml operation request
type GetXMLRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetXMLRequest) xxx_ToOp(ctx context.Context) *xxx_GetXMLOperation {
	if o == nil {
		return &xxx_GetXMLOperation{}
	}
	return &xxx_GetXMLOperation{
		This: o.This,
	}
}

func (o *GetXMLRequest) xxx_FromOp(ctx context.Context, op *xxx_GetXMLOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetXMLRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetXMLRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetXMLOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetXMLResponse structure represents the Xml operation response
type GetXMLResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	XML  *oaut.String   `idl:"name:xml" json:"xml"`
	// Return: The Xml return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetXMLResponse) xxx_ToOp(ctx context.Context) *xxx_GetXMLOperation {
	if o == nil {
		return &xxx_GetXMLOperation{}
	}
	return &xxx_GetXMLOperation{
		That:   o.That,
		XML:    o.XML,
		Return: o.Return,
	}
}

func (o *GetXMLResponse) xxx_FromOp(ctx context.Context, op *xxx_GetXMLOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.XML = op.XML
	o.Return = op.Return
}
func (o *GetXMLResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetXMLResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetXMLOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSecurityOperation structure represents the Security operation
type xxx_GetSecurityOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Security *oaut.String   `idl:"name:pbstrSecurity" json:"security"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSecurityOperation) OpNum() int { return 53 }

func (o *xxx_GetSecurityOperation) OpName() string { return "/IDataCollectorSet/v0/Security" }

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

func (o *xxx_GetSecurityOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pbstrSecurity {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Security != nil {
			_ptr_pbstrSecurity := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Security != nil {
					if err := o.Security.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Security, _ptr_pbstrSecurity); err != nil {
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
	// pbstrSecurity {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrSecurity := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Security == nil {
				o.Security = &oaut.String{}
			}
			if err := o.Security.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrSecurity := func(ptr interface{}) { o.Security = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Security, _s_pbstrSecurity, _ptr_pbstrSecurity); err != nil {
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

// GetSecurityRequest structure represents the Security operation request
type GetSecurityRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSecurityRequest) xxx_ToOp(ctx context.Context) *xxx_GetSecurityOperation {
	if o == nil {
		return &xxx_GetSecurityOperation{}
	}
	return &xxx_GetSecurityOperation{
		This: o.This,
	}
}

func (o *GetSecurityRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSecurityOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSecurityRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetSecurityRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSecurityOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSecurityResponse structure represents the Security operation response
type GetSecurityResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Security *oaut.String   `idl:"name:pbstrSecurity" json:"security"`
	// Return: The Security return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSecurityResponse) xxx_ToOp(ctx context.Context) *xxx_GetSecurityOperation {
	if o == nil {
		return &xxx_GetSecurityOperation{}
	}
	return &xxx_GetSecurityOperation{
		That:     o.That,
		Security: o.Security,
		Return:   o.Return,
	}
}

func (o *GetSecurityResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSecurityOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Security = op.Security
	o.Return = op.Return
}
func (o *GetSecurityResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetSecurityResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSecurityOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSecurityOperation structure represents the Security operation
type xxx_SetSecurityOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Security *oaut.String   `idl:"name:bstrSecurity" json:"security"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSecurityOperation) OpNum() int { return 54 }

func (o *xxx_SetSecurityOperation) OpName() string { return "/IDataCollectorSet/v0/Security" }

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
	// bstrSecurity {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Security != nil {
			_ptr_bstrSecurity := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Security != nil {
					if err := o.Security.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Security, _ptr_bstrSecurity); err != nil {
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

func (o *xxx_SetSecurityOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrSecurity {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrSecurity := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Security == nil {
				o.Security = &oaut.String{}
			}
			if err := o.Security.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrSecurity := func(ptr interface{}) { o.Security = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Security, _s_bstrSecurity, _ptr_bstrSecurity); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
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

func (o *xxx_SetSecurityOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetSecurityRequest structure represents the Security operation request
type SetSecurityRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	Security *oaut.String   `idl:"name:bstrSecurity" json:"security"`
}

func (o *SetSecurityRequest) xxx_ToOp(ctx context.Context) *xxx_SetSecurityOperation {
	if o == nil {
		return &xxx_SetSecurityOperation{}
	}
	return &xxx_SetSecurityOperation{
		This:     o.This,
		Security: o.Security,
	}
}

func (o *SetSecurityRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSecurityOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Security = op.Security
}
func (o *SetSecurityRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetSecurityRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSecurityOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSecurityResponse structure represents the Security operation response
type SetSecurityResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Security return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSecurityResponse) xxx_ToOp(ctx context.Context) *xxx_SetSecurityOperation {
	if o == nil {
		return &xxx_SetSecurityOperation{}
	}
	return &xxx_SetSecurityOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetSecurityResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSecurityOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSecurityResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetSecurityResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSecurityOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetStopOnCompletionOperation structure represents the StopOnCompletion operation
type xxx_GetStopOnCompletionOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Stop   int16          `idl:"name:Stop" json:"stop"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetStopOnCompletionOperation) OpNum() int { return 55 }

func (o *xxx_GetStopOnCompletionOperation) OpName() string {
	return "/IDataCollectorSet/v0/StopOnCompletion"
}

func (o *xxx_GetStopOnCompletionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetStopOnCompletionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetStopOnCompletionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetStopOnCompletionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetStopOnCompletionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// Stop {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Stop); err != nil {
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

func (o *xxx_GetStopOnCompletionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// Stop {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Stop); err != nil {
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

// GetStopOnCompletionRequest structure represents the StopOnCompletion operation request
type GetStopOnCompletionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetStopOnCompletionRequest) xxx_ToOp(ctx context.Context) *xxx_GetStopOnCompletionOperation {
	if o == nil {
		return &xxx_GetStopOnCompletionOperation{}
	}
	return &xxx_GetStopOnCompletionOperation{
		This: o.This,
	}
}

func (o *GetStopOnCompletionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetStopOnCompletionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetStopOnCompletionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetStopOnCompletionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetStopOnCompletionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetStopOnCompletionResponse structure represents the StopOnCompletion operation response
type GetStopOnCompletionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	Stop int16          `idl:"name:Stop" json:"stop"`
	// Return: The StopOnCompletion return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetStopOnCompletionResponse) xxx_ToOp(ctx context.Context) *xxx_GetStopOnCompletionOperation {
	if o == nil {
		return &xxx_GetStopOnCompletionOperation{}
	}
	return &xxx_GetStopOnCompletionOperation{
		That:   o.That,
		Stop:   o.Stop,
		Return: o.Return,
	}
}

func (o *GetStopOnCompletionResponse) xxx_FromOp(ctx context.Context, op *xxx_GetStopOnCompletionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Stop = op.Stop
	o.Return = op.Return
}
func (o *GetStopOnCompletionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetStopOnCompletionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetStopOnCompletionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetStopOnCompletionOperation structure represents the StopOnCompletion operation
type xxx_SetStopOnCompletionOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Stop   int16          `idl:"name:Stop" json:"stop"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetStopOnCompletionOperation) OpNum() int { return 56 }

func (o *xxx_SetStopOnCompletionOperation) OpName() string {
	return "/IDataCollectorSet/v0/StopOnCompletion"
}

func (o *xxx_SetStopOnCompletionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetStopOnCompletionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Stop {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Stop); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetStopOnCompletionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Stop {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Stop); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetStopOnCompletionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetStopOnCompletionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetStopOnCompletionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetStopOnCompletionRequest structure represents the StopOnCompletion operation request
type SetStopOnCompletionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	Stop int16          `idl:"name:Stop" json:"stop"`
}

func (o *SetStopOnCompletionRequest) xxx_ToOp(ctx context.Context) *xxx_SetStopOnCompletionOperation {
	if o == nil {
		return &xxx_SetStopOnCompletionOperation{}
	}
	return &xxx_SetStopOnCompletionOperation{
		This: o.This,
		Stop: o.Stop,
	}
}

func (o *SetStopOnCompletionRequest) xxx_FromOp(ctx context.Context, op *xxx_SetStopOnCompletionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Stop = op.Stop
}
func (o *SetStopOnCompletionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetStopOnCompletionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetStopOnCompletionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetStopOnCompletionResponse structure represents the StopOnCompletion operation response
type SetStopOnCompletionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The StopOnCompletion return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetStopOnCompletionResponse) xxx_ToOp(ctx context.Context) *xxx_SetStopOnCompletionOperation {
	if o == nil {
		return &xxx_SetStopOnCompletionOperation{}
	}
	return &xxx_SetStopOnCompletionOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetStopOnCompletionResponse) xxx_FromOp(ctx context.Context, op *xxx_SetStopOnCompletionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetStopOnCompletionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetStopOnCompletionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetStopOnCompletionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDataManagerOperation structure represents the DataManager operation
type xxx_GetDataManagerOperation struct {
	This        *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat   `idl:"name:That" json:"that"`
	DataManager *pla.DataManager `idl:"name:DataManager" json:"data_manager"`
	Return      int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDataManagerOperation) OpNum() int { return 57 }

func (o *xxx_GetDataManagerOperation) OpName() string { return "/IDataCollectorSet/v0/DataManager" }

func (o *xxx_GetDataManagerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDataManagerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetDataManagerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetDataManagerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDataManagerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// DataManager {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IDataManager}(interface))
	{
		if o.DataManager != nil {
			_ptr_DataManager := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DataManager != nil {
					if err := o.DataManager.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&pla.DataManager{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DataManager, _ptr_DataManager); err != nil {
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

func (o *xxx_GetDataManagerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// DataManager {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IDataManager}(interface))
	{
		_ptr_DataManager := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DataManager == nil {
				o.DataManager = &pla.DataManager{}
			}
			if err := o.DataManager.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_DataManager := func(ptr interface{}) { o.DataManager = *ptr.(**pla.DataManager) }
		if err := w.ReadPointer(&o.DataManager, _s_DataManager, _ptr_DataManager); err != nil {
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

// GetDataManagerRequest structure represents the DataManager operation request
type GetDataManagerRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetDataManagerRequest) xxx_ToOp(ctx context.Context) *xxx_GetDataManagerOperation {
	if o == nil {
		return &xxx_GetDataManagerOperation{}
	}
	return &xxx_GetDataManagerOperation{
		This: o.This,
	}
}

func (o *GetDataManagerRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDataManagerOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetDataManagerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetDataManagerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDataManagerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDataManagerResponse structure represents the DataManager operation response
type GetDataManagerResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// DataManager: Receives a pointer to the data manager object.
	DataManager *pla.DataManager `idl:"name:DataManager" json:"data_manager"`
	// Return: The DataManager return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDataManagerResponse) xxx_ToOp(ctx context.Context) *xxx_GetDataManagerOperation {
	if o == nil {
		return &xxx_GetDataManagerOperation{}
	}
	return &xxx_GetDataManagerOperation{
		That:        o.That,
		DataManager: o.DataManager,
		Return:      o.Return,
	}
}

func (o *GetDataManagerResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDataManagerOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DataManager = op.DataManager
	o.Return = op.Return
}
func (o *GetDataManagerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetDataManagerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDataManagerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetCredentialsOperation structure represents the SetCredentials operation
type xxx_SetCredentialsOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	User     *oaut.String   `idl:"name:user" json:"user"`
	Password *oaut.String   `idl:"name:password" json:"password"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetCredentialsOperation) OpNum() int { return 58 }

func (o *xxx_SetCredentialsOperation) OpName() string { return "/IDataCollectorSet/v0/SetCredentials" }

func (o *xxx_SetCredentialsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCredentialsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// user {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.User != nil {
			_ptr_user := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.User != nil {
					if err := o.User.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
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
	// password {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Password != nil {
			_ptr_password := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Password != nil {
					if err := o.Password.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
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
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCredentialsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// user {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_user := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.User == nil {
				o.User = &oaut.String{}
			}
			if err := o.User.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_user := func(ptr interface{}) { o.User = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.User, _s_user, _ptr_user); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// password {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_password := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Password == nil {
				o.Password = &oaut.String{}
			}
			if err := o.Password.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_password := func(ptr interface{}) { o.Password = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Password, _s_password, _ptr_password); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCredentialsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCredentialsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetCredentialsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetCredentialsRequest structure represents the SetCredentials operation request
type SetCredentialsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// user: Supplies the user account under which the data collector set runs. The user
	// name can be specified in the following forms: domain\user or user@domain.
	User *oaut.String `idl:"name:user" json:"user"`
	// password:  Supplies the password of the user account.
	Password *oaut.String `idl:"name:password" json:"password"`
}

func (o *SetCredentialsRequest) xxx_ToOp(ctx context.Context) *xxx_SetCredentialsOperation {
	if o == nil {
		return &xxx_SetCredentialsOperation{}
	}
	return &xxx_SetCredentialsOperation{
		This:     o.This,
		User:     o.User,
		Password: o.Password,
	}
}

func (o *SetCredentialsRequest) xxx_FromOp(ctx context.Context, op *xxx_SetCredentialsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.User = op.User
	o.Password = op.Password
}
func (o *SetCredentialsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetCredentialsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetCredentialsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetCredentialsResponse structure represents the SetCredentials operation response
type SetCredentialsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetCredentials return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetCredentialsResponse) xxx_ToOp(ctx context.Context) *xxx_SetCredentialsOperation {
	if o == nil {
		return &xxx_SetCredentialsOperation{}
	}
	return &xxx_SetCredentialsOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetCredentialsResponse) xxx_FromOp(ctx context.Context, op *xxx_SetCredentialsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetCredentialsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetCredentialsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetCredentialsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryOperation structure represents the Query operation
type xxx_QueryOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Name   *oaut.String   `idl:"name:name" json:"name"`
	Server *oaut.String   `idl:"name:server;pointer:unique" json:"server"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryOperation) OpNum() int { return 59 }

func (o *xxx_QueryOperation) OpName() string { return "/IDataCollectorSet/v0/Query" }

func (o *xxx_QueryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// name {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Name != nil {
			_ptr_name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Name != nil {
					if err := o.Name.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Name, _ptr_name); err != nil {
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
	// server {in} (1:{pointer=unique, alias=BSTR, pointers=[unique]}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Server != nil {
			_ptr_server := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Server != nil {
					if err := o.Server.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Server, _ptr_server); err != nil {
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

func (o *xxx_QueryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// name {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Name == nil {
				o.Name = &oaut.String{}
			}
			if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_name := func(ptr interface{}) { o.Name = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Name, _s_name, _ptr_name); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// server {in} (1:{pointer=unique, alias=BSTR, pointers=[unique]}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_server := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Server == nil {
				o.Server = &oaut.String{}
			}
			if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_server := func(ptr interface{}) { o.Server = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Server, _s_server, _ptr_server); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_QueryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// QueryRequest structure represents the Query operation request
type QueryRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	Name   *oaut.String   `idl:"name:name" json:"name"`
	Server *oaut.String   `idl:"name:server;pointer:unique" json:"server"`
}

func (o *QueryRequest) xxx_ToOp(ctx context.Context) *xxx_QueryOperation {
	if o == nil {
		return &xxx_QueryOperation{}
	}
	return &xxx_QueryOperation{
		This:   o.This,
		Name:   o.Name,
		Server: o.Server,
	}
}

func (o *QueryRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Name = op.Name
	o.Server = op.Server
}
func (o *QueryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryResponse structure represents the Query operation response
type QueryResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Query return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryResponse) xxx_ToOp(ctx context.Context) *xxx_QueryOperation {
	if o == nil {
		return &xxx_QueryOperation{}
	}
	return &xxx_QueryOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *QueryResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *QueryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CommitOperation structure represents the Commit operation
type xxx_CommitOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	Name       *oaut.String   `idl:"name:name" json:"name"`
	Server     *oaut.String   `idl:"name:server;pointer:unique" json:"server"`
	Mode       pla.CommitMode `idl:"name:mode" json:"mode"`
	Validation *pla.ValueMap  `idl:"name:validation" json:"validation"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CommitOperation) OpNum() int { return 60 }

func (o *xxx_CommitOperation) OpName() string { return "/IDataCollectorSet/v0/Commit" }

func (o *xxx_CommitOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommitOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// name {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Name != nil {
			_ptr_name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Name != nil {
					if err := o.Name.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Name, _ptr_name); err != nil {
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
	// server {in} (1:{pointer=unique, alias=BSTR, pointers=[unique]}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Server != nil {
			_ptr_server := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Server != nil {
					if err := o.Server.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Server, _ptr_server); err != nil {
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
	// mode {in} (1:{alias=CommitMode}(enum))
	{
		if err := w.WriteData(uint16(o.Mode)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommitOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// name {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Name == nil {
				o.Name = &oaut.String{}
			}
			if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_name := func(ptr interface{}) { o.Name = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Name, _s_name, _ptr_name); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// server {in} (1:{pointer=unique, alias=BSTR, pointers=[unique]}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_server := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Server == nil {
				o.Server = &oaut.String{}
			}
			if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_server := func(ptr interface{}) { o.Server = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Server, _s_server, _ptr_server); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// mode {in} (1:{alias=CommitMode}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Mode)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommitOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommitOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// validation {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IValueMap}(interface))
	{
		if o.Validation != nil {
			_ptr_validation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Validation != nil {
					if err := o.Validation.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&pla.ValueMap{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Validation, _ptr_validation); err != nil {
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

func (o *xxx_CommitOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// validation {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IValueMap}(interface))
	{
		_ptr_validation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Validation == nil {
				o.Validation = &pla.ValueMap{}
			}
			if err := o.Validation.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_validation := func(ptr interface{}) { o.Validation = *ptr.(**pla.ValueMap) }
		if err := w.ReadPointer(&o.Validation, _s_validation, _ptr_validation); err != nil {
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

// CommitRequest structure represents the Commit operation request
type CommitRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// name: Supplies a unique name used to identify a data collector set.
	Name *oaut.String `idl:"name:name" json:"name"`
	// server: Supplies a name string that MUST be less than 1024 characters in length.
	// PLA MUST return this name from the IDataCollectorSet::Server(Get) method defined
	// in section 3.2.4.1.26; otherwise, this string is not used by PLA.
	Server *oaut.String `idl:"name:server;pointer:unique" json:"server"`
	// mode: Supplies the actions to perform. See section 2.2.2.3 for a description of possible
	// actions.
	Mode pla.CommitMode `idl:"name:mode" json:"mode"`
}

func (o *CommitRequest) xxx_ToOp(ctx context.Context) *xxx_CommitOperation {
	if o == nil {
		return &xxx_CommitOperation{}
	}
	return &xxx_CommitOperation{
		This:   o.This,
		Name:   o.Name,
		Server: o.Server,
		Mode:   o.Mode,
	}
}

func (o *CommitRequest) xxx_FromOp(ctx context.Context, op *xxx_CommitOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Name = op.Name
	o.Server = op.Server
	o.Mode = op.Mode
}
func (o *CommitRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CommitRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CommitOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CommitResponse structure represents the Commit operation response
type CommitResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// validation: Receives a validation value map with a list of properties for this data
	// collector set (and its encapsulated objects) that are invalid or ignored. For each
	// property of the data collector set and its associated objects, passed in by the client,
	// that could not be set, the server MUST return in an IValueMap. Each IValueMapItem
	// in the IValueMap represents a property of the data collector set and its encapsulated
	// objects that could not be set by the server. The Names property of the IValueMapItem
	// represents the property name, while the Values property of the IValueMap represents
	// the HRESULT describing the specific property corresponding to that property. The
	// ValueMapType property of the IValueMap is plaValidation; more information can be
	// found in section 2.2.11. Note that the client MAY choose to ignore any warnings or
	// errors that are returned by the server; however, if it does so, the data collector
	// set might not be executed by the server as the client expects. The client MUST ignore
	// the validation property if the Commit method fails. The validation property will
	// be NULL, and an error code returned, in the case that the data collector set contains
	// an IApiTracingDataCollectorSet and there are array elements that are part of both
	// the IApiTracingDataCollectorSet::IncludeApis array and IApiTracingDataCollectorSet::ExcludeApis
	// array (in other words, there is overlap between the two arrays).
	Validation *pla.ValueMap `idl:"name:validation" json:"validation"`
	// Return: The Commit return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CommitResponse) xxx_ToOp(ctx context.Context) *xxx_CommitOperation {
	if o == nil {
		return &xxx_CommitOperation{}
	}
	return &xxx_CommitOperation{
		That:       o.That,
		Validation: o.Validation,
		Return:     o.Return,
	}
}

func (o *CommitResponse) xxx_FromOp(ctx context.Context, op *xxx_CommitOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Validation = op.Validation
	o.Return = op.Return
}
func (o *CommitResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CommitResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CommitOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteOperation structure represents the Delete operation
type xxx_DeleteOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteOperation) OpNum() int { return 61 }

func (o *xxx_DeleteOperation) OpName() string { return "/IDataCollectorSet/v0/Delete" }

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

func (o *xxx_DeleteOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_DeleteOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteRequest structure represents the Delete operation request
type DeleteRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *DeleteRequest) xxx_ToOp(ctx context.Context) *xxx_DeleteOperation {
	if o == nil {
		return &xxx_DeleteOperation{}
	}
	return &xxx_DeleteOperation{
		This: o.This,
	}
}

func (o *DeleteRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *DeleteRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteResponse structure represents the Delete operation response
type DeleteResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Delete return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteResponse) xxx_ToOp(ctx context.Context) *xxx_DeleteOperation {
	if o == nil {
		return &xxx_DeleteOperation{}
	}
	return &xxx_DeleteOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *DeleteResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_StartOperation structure represents the Start operation
type xxx_StartOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	Synchronous int16          `idl:"name:Synchronous" json:"synchronous"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_StartOperation) OpNum() int { return 62 }

func (o *xxx_StartOperation) OpName() string { return "/IDataCollectorSet/v0/Start" }

func (o *xxx_StartOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Synchronous {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Synchronous); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Synchronous {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Synchronous); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_StartOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// StartRequest structure represents the Start operation request
type StartRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// Synchronous:  Supplies a Boolean indicating whether the start operation MUST be
	// synchronous or asynchronous. In asynchronous mode, the method returns after queuing
	// or failing to queue the data collector set start. The HRESULT returned from the Start
	// function only specifies whether the queuing operation failed or succeeded. If the
	// queuing operation succeeded, (S_OK) is returned even if the data collector set did
	// not start. The only method for detecting that the asynchronous start failed is to
	// poll the Status property to check if the data collector set is still running. In
	// synchronous mode, the method returns when the data collector set actually starts,
	// and the HRESULT only returns (S_OK) if the data collector set started successfully.
	Synchronous int16 `idl:"name:Synchronous" json:"synchronous"`
}

func (o *StartRequest) xxx_ToOp(ctx context.Context) *xxx_StartOperation {
	if o == nil {
		return &xxx_StartOperation{}
	}
	return &xxx_StartOperation{
		This:        o.This,
		Synchronous: o.Synchronous,
	}
}

func (o *StartRequest) xxx_FromOp(ctx context.Context, op *xxx_StartOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Synchronous = op.Synchronous
}
func (o *StartRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *StartRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StartOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// StartResponse structure represents the Start operation response
type StartResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Start return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *StartResponse) xxx_ToOp(ctx context.Context) *xxx_StartOperation {
	if o == nil {
		return &xxx_StartOperation{}
	}
	return &xxx_StartOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *StartResponse) xxx_FromOp(ctx context.Context, op *xxx_StartOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *StartResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *StartResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StartOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_StopOperation structure represents the Stop operation
type xxx_StopOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	Synchronous int16          `idl:"name:Synchronous" json:"synchronous"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_StopOperation) OpNum() int { return 63 }

func (o *xxx_StopOperation) OpName() string { return "/IDataCollectorSet/v0/Stop" }

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
	// Synchronous {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Synchronous); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Synchronous {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Synchronous); err != nil {
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

func (o *xxx_StopOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// StopRequest structure represents the Stop operation request
type StopRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// Synchronous: Supplies a Boolean indicating whether the stop operation MUST be synchronous
	// or asynchronous. In synchronous mode, the method returns when the data collector
	// set actually stops. In asynchronous mode, the method returns after queuing or failing
	// to queue the data collector set stop. The HRESULT returned from the Stop function
	// only specifies whether the queuing operation failed or succeeded. If the queuing
	// operation succeeded, (S_OK) is returned even if the data collector set did not stop.
	// The only method for detecting that the asynchronous stop failed is to poll the Status
	// property to check if the data collector set is still running. In synchronous mode,
	// the method returns when the data collector set actually stops and the HRESULT only
	// returns (S_OK) if the data collector set stopped successfully.
	Synchronous int16 `idl:"name:Synchronous" json:"synchronous"`
}

func (o *StopRequest) xxx_ToOp(ctx context.Context) *xxx_StopOperation {
	if o == nil {
		return &xxx_StopOperation{}
	}
	return &xxx_StopOperation{
		This:        o.This,
		Synchronous: o.Synchronous,
	}
}

func (o *StopRequest) xxx_FromOp(ctx context.Context, op *xxx_StopOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Synchronous = op.Synchronous
}
func (o *StopRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *StopRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StopOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// StopResponse structure represents the Stop operation response
type StopResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Stop return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *StopResponse) xxx_ToOp(ctx context.Context) *xxx_StopOperation {
	if o == nil {
		return &xxx_StopOperation{}
	}
	return &xxx_StopOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *StopResponse) xxx_FromOp(ctx context.Context, op *xxx_StopOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *StopResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *StopResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StopOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetXMLOperation structure represents the SetXml operation
type xxx_SetXMLOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	XML        *oaut.String   `idl:"name:xml" json:"xml"`
	Validation *pla.ValueMap  `idl:"name:validation" json:"validation"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetXMLOperation) OpNum() int { return 64 }

func (o *xxx_SetXMLOperation) OpName() string { return "/IDataCollectorSet/v0/SetXml" }

func (o *xxx_SetXMLOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetXMLOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// xml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.XML != nil {
			_ptr_xml := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.XML != nil {
					if err := o.XML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.XML, _ptr_xml); err != nil {
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

func (o *xxx_SetXMLOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// xml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_xml := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.XML == nil {
				o.XML = &oaut.String{}
			}
			if err := o.XML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_xml := func(ptr interface{}) { o.XML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.XML, _s_xml, _ptr_xml); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetXMLOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetXMLOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// validation {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IValueMap}(interface))
	{
		if o.Validation != nil {
			_ptr_validation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Validation != nil {
					if err := o.Validation.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&pla.ValueMap{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Validation, _ptr_validation); err != nil {
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

func (o *xxx_SetXMLOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// validation {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IValueMap}(interface))
	{
		_ptr_validation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Validation == nil {
				o.Validation = &pla.ValueMap{}
			}
			if err := o.Validation.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_validation := func(ptr interface{}) { o.Validation = *ptr.(**pla.ValueMap) }
		if err := w.ReadPointer(&o.Validation, _s_validation, _ptr_validation); err != nil {
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

// SetXMLRequest structure represents the SetXml operation request
type SetXMLRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	XML  *oaut.String   `idl:"name:xml" json:"xml"`
}

func (o *SetXMLRequest) xxx_ToOp(ctx context.Context) *xxx_SetXMLOperation {
	if o == nil {
		return &xxx_SetXMLOperation{}
	}
	return &xxx_SetXMLOperation{
		This: o.This,
		XML:  o.XML,
	}
}

func (o *SetXMLRequest) xxx_FromOp(ctx context.Context, op *xxx_SetXMLOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.XML = op.XML
}
func (o *SetXMLRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetXMLRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetXMLOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetXMLResponse structure represents the SetXml operation response
type SetXMLResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	Validation *pla.ValueMap  `idl:"name:validation" json:"validation"`
	// Return: The SetXml return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetXMLResponse) xxx_ToOp(ctx context.Context) *xxx_SetXMLOperation {
	if o == nil {
		return &xxx_SetXMLOperation{}
	}
	return &xxx_SetXMLOperation{
		That:       o.That,
		Validation: o.Validation,
		Return:     o.Return,
	}
}

func (o *SetXMLResponse) xxx_FromOp(ctx context.Context, op *xxx_SetXMLOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Validation = op.Validation
	o.Return = op.Return
}
func (o *SetXMLResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetXMLResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetXMLOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetValueOperation structure represents the SetValue operation
type xxx_SetValueOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Key    *oaut.String   `idl:"name:key" json:"key"`
	Value  *oaut.String   `idl:"name:value" json:"value"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetValueOperation) OpNum() int { return 65 }

func (o *xxx_SetValueOperation) OpName() string { return "/IDataCollectorSet/v0/SetValue" }

func (o *xxx_SetValueOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetValueOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// key {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Key != nil {
			_ptr_key := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Key != nil {
					if err := o.Key.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Key, _ptr_key); err != nil {
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
	// value {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Value != nil {
			_ptr_value := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Value != nil {
					if err := o.Value.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
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
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetValueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// key {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_key := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Key == nil {
				o.Key = &oaut.String{}
			}
			if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_key := func(ptr interface{}) { o.Key = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Key, _s_key, _ptr_key); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// value {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_value := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Value == nil {
				o.Value = &oaut.String{}
			}
			if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_value := func(ptr interface{}) { o.Value = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Value, _s_value, _ptr_value); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetValueOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetValueOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetValueOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetValueRequest structure represents the SetValue operation request
type SetValueRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// key: Supplies the name of the value. The key cannot be the empty BSTR, but any other
	// BSTR is a valid value. This key is associated with, and can be used to retrieve,
	// the value field.
	Key *oaut.String `idl:"name:key" json:"key"`
	// value: Supplies the value associated with the key. This is any BSTR that is metadata
	// that needs to be associated with the data collector set or be passed as an argument
	// when a specified task executes.
	Value *oaut.String `idl:"name:value" json:"value"`
}

func (o *SetValueRequest) xxx_ToOp(ctx context.Context) *xxx_SetValueOperation {
	if o == nil {
		return &xxx_SetValueOperation{}
	}
	return &xxx_SetValueOperation{
		This:  o.This,
		Key:   o.Key,
		Value: o.Value,
	}
}

func (o *SetValueRequest) xxx_FromOp(ctx context.Context, op *xxx_SetValueOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Key = op.Key
	o.Value = op.Value
}
func (o *SetValueRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetValueRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetValueOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetValueResponse structure represents the SetValue operation response
type SetValueResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetValue return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetValueResponse) xxx_ToOp(ctx context.Context) *xxx_SetValueOperation {
	if o == nil {
		return &xxx_SetValueOperation{}
	}
	return &xxx_SetValueOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetValueResponse) xxx_FromOp(ctx context.Context, op *xxx_SetValueOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetValueResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetValueResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetValueOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetValueOperation structure represents the GetValue operation
type xxx_GetValueOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Key    *oaut.String   `idl:"name:key" json:"key"`
	Value  *oaut.String   `idl:"name:value" json:"value"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetValueOperation) OpNum() int { return 66 }

func (o *xxx_GetValueOperation) OpName() string { return "/IDataCollectorSet/v0/GetValue" }

func (o *xxx_GetValueOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetValueOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// key {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Key != nil {
			_ptr_key := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Key != nil {
					if err := o.Key.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Key, _ptr_key); err != nil {
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

func (o *xxx_GetValueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// key {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_key := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Key == nil {
				o.Key = &oaut.String{}
			}
			if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_key := func(ptr interface{}) { o.Key = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Key, _s_key, _ptr_key); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetValueOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetValueOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// value {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Value != nil {
			_ptr_value := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Value != nil {
					if err := o.Value.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
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

func (o *xxx_GetValueOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// value {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_value := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Value == nil {
				o.Value = &oaut.String{}
			}
			if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_value := func(ptr interface{}) { o.Value = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Value, _s_value, _ptr_value); err != nil {
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

// GetValueRequest structure represents the GetValue operation request
type GetValueRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// key: Supplies the key of the value to retrieve. The key cannot be null or the empty
	// BSTR. Any other BSTR is a valid value.
	Key *oaut.String `idl:"name:key" json:"key"`
}

func (o *GetValueRequest) xxx_ToOp(ctx context.Context) *xxx_GetValueOperation {
	if o == nil {
		return &xxx_GetValueOperation{}
	}
	return &xxx_GetValueOperation{
		This: o.This,
		Key:  o.Key,
	}
}

func (o *GetValueRequest) xxx_FromOp(ctx context.Context, op *xxx_GetValueOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Key = op.Key
}
func (o *GetValueRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetValueRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetValueOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetValueResponse structure represents the GetValue operation response
type GetValueResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// value: Receives the value associated with the key. The value that is returned, associated
	// with the key that was passed as a parameter into this method, was set by calling
	// the SetValue method.
	Value *oaut.String `idl:"name:value" json:"value"`
	// Return: The GetValue return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetValueResponse) xxx_ToOp(ctx context.Context) *xxx_GetValueOperation {
	if o == nil {
		return &xxx_GetValueOperation{}
	}
	return &xxx_GetValueOperation{
		That:   o.That,
		Value:  o.Value,
		Return: o.Return,
	}
}

func (o *GetValueResponse) xxx_FromOp(ctx context.Context, op *xxx_GetValueOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Value = op.Value
	o.Return = op.Return
}
func (o *GetValueResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetValueResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetValueOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
